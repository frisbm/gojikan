package internal

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

// cachedResponse stores the details of a cached HTTP response.
type cachedResponse struct {
	statusCode int
	header     http.Header
	body       []byte // Store body as bytes to allow re-reading
	expiry     time.Time
}

// CachingTransport is an http.RoundTripper that adds caching for GET requests.
type CachingTransport struct {
	t          http.RoundTripper
	cache      map[string]cachedResponse
	cacheMutex sync.RWMutex
	ttl        time.Duration
}

// NewCachingTransport creates a new CachingTransport.
func NewCachingTransport(transport http.RoundTripper, ttl time.Duration) *CachingTransport {
	return &CachingTransport{
		t:     transport,
		cache: make(map[string]cachedResponse),
		ttl:   ttl,
	}
}

// RoundTrip executes a single HTTP transaction, possibly returning a cached response.
func (ct *CachingTransport) RoundTrip(req *http.Request) (resp *http.Response, ferr error) {
	if req.Method != http.MethodGet {
		return ct.t.RoundTrip(req)
	}
	cacheKey := req.URL.String()

	ct.cacheMutex.RLock()
	cached, found := ct.cache[cacheKey]
	ct.cacheMutex.RUnlock()

	if found && time.Now().Before(cached.expiry) {
		response := &http.Response{
			StatusCode:    cached.statusCode,
			Header:        cached.header.Clone(),
			Body:          io.NopCloser(bytes.NewReader(cached.body)),
			ContentLength: int64(len(cached.body)),
			Request:       req,        // Associate original request
			Proto:         "HTTP/1.1", // Assume HTTP/1.1, adjust if needed
			ProtoMajor:    1,
			ProtoMinor:    1,
		}
		return response, nil
	}

	resp, err := ct.t.RoundTrip(req)
	if err != nil {
		return nil, fmt.Errorf("error in RoundTrip: %w", err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		bodyBytes, readErr := io.ReadAll(resp.Body)
		defer func(resp *http.Response) {
			if cerr := resp.Body.Close(); cerr != nil {
				ferr = errors.Join(ferr, fmt.Errorf("error closing response body: %w", cerr))
			}
		}(resp)

		if readErr != nil {
			resp.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			return resp, nil
		}

		newCached := cachedResponse{
			statusCode: resp.StatusCode,
			header:     resp.Header.Clone(),
			body:       bodyBytes,
			expiry:     time.Now().Add(ct.ttl),
		}

		ct.cacheMutex.Lock()
		ct.cache[cacheKey] = newCached
		ct.cacheMutex.Unlock()
		resp.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	return resp, nil
}

// HttpClient creates an HTTP client with GET request caching.
func HttpClient(withcache bool, ttl *time.Duration) *http.Client {
	baseTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   15 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   15 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConns:          10,
		IdleConnTimeout:       60 * time.Second,
	}
	if !withcache {
		return &http.Client{
			Transport: baseTransport,
			Timeout:   60 * time.Second,
		}
	}

	cachettl := 1_000_000_000 * time.Second
	if ttl != nil {
		cachettl = *ttl
	}

	cachingTransport := NewCachingTransport(baseTransport, cachettl)

	client := &http.Client{
		Transport: cachingTransport,
		Timeout:   60 * time.Second,
	}
	return client
}
