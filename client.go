package gojikan

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/time/rate"

	"github.com/frisbm/gojikan/internal"
)

type Client struct {
	c         *http.Client
	baseURL   string
	withCache bool
	ttl       *time.Duration

	rateLimit *rate.Limiter
}

func New(opts ...ClientOption) (*Client, error) {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	if c.c != nil && c.withCache {
		return nil, errors.New("cache option can only be used once")
	}
	if c.c == nil {
		c.c = internal.HttpClient(c.withCache, c.ttl)
	}
	if c.baseURL == "" {
		c.baseURL = "https://api.jikan.moe/v4"
	}

	return c, nil
}

type ClientOption func(c *Client)

// WithCache returns a ClientOption that enables caching for the client.
func WithCache(ttl time.Duration) ClientOption {
	return func(c *Client) {
		c.withCache = true
		c.ttl = &ttl
	}
}

// WithHTTPClient returns a ClientOption that allows you to set a custom HTTP client. Cannot be used with WithCache.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.c = client
	}
}

// WithBaseURL returns a ClientOption that allows you to set a custom base URL for the client.
// The base URL will be trimmed of leading and trailing slashes: "/"
func WithBaseURL(baseUrl string) ClientOption {
	return func(c *Client) {
		c.baseURL = strings.Trim(baseUrl, "/")
	}
}

func WithRateLimit(limit, burst int) ClientOption {
	return func(c *Client) {
		if limit <= 0 {
			return
		}
		c.rateLimit = rate.NewLimiter(rate.Limit(limit), burst)
	}
}

func get[T any](ctx context.Context, c *Client, url string) (t T, ferr error) {
	zero := *new(T)

	if c.rateLimit != nil {
		if err := c.rateLimit.Wait(ctx); err != nil {
			return zero, fmt.Errorf("rate limit: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return zero, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.c.Do(req)
	if err != nil {
		return zero, fmt.Errorf("sending request: %w", err)
	}

	defer func(resp *http.Response) {
		if cerr := resp.Body.Close(); cerr != nil {
			ferr = errors.Join(ferr, fmt.Errorf("error closing response body: %w", cerr))
		}
	}(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		errResponse := ErrorResponse{}
		err = json.Unmarshal(body, &errResponse)
		if err != nil {
			return zero, fmt.Errorf("unmarshalling response body: %w for status code: %d", err, resp.StatusCode)
		}
		return zero, &errResponse
	}

	var result T
	err = json.Unmarshal(body, &result)
	if err != nil {
		return zero, fmt.Errorf("unmarshalling response body for type %T: %w", result, err)
	}

	return result, nil
}
