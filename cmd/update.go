package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const ApiDocUrl = "https://raw.githubusercontent.com/jikan-me/jikan-rest/master/storage/api-docs/api-docs.json"

type specReplacement struct {
	old string
	new string
}

var specReplacements = []specReplacement{
	{"\"type\": \"float\"", "\"type\": \"number\""},
	{"\"maximum\": \"5\"", "\"maximum\": 5"},
	{"\"minimum\": \"1\"", "\"minimum\": 1"},
}

func main() {
	ctx := context.Background()
	fmt.Println("Starting OpenAPI spec update and code generation...")
	err := update(ctx)
	if err != nil {
		log.Fatalf("Update process failed: %v", err)
	}
	fmt.Println("Update process completed successfully.")
}

func update(ctx context.Context) (ferr error) {
	var specBuffer bytes.Buffer

	if err := pullSpec(ctx, &specBuffer); err != nil {
		return fmt.Errorf("failed to pull spec: %w", err)
	}
	modifiedSpec := replaceInSpec(specBuffer.String(), specReplacements)
	if err := writeSpec(modifiedSpec); err != nil {
		return fmt.Errorf("failed to write spec: %w", err)
	}
	if err := generate(); err != nil {
		return fmt.Errorf("code generation failed: %w", err)
	}
	if err := fmtCode(); err != nil {
		return fmt.Errorf("code formatting failed: %w", err)
	}
	return nil
}

func fmtCode() (ferr error) {
	cmd := exec.Command("make", "fmt")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		errMsg := fmt.Sprintf("'make fmt' failed with error: %v", err)
		if stderrStr := stderr.String(); stderrStr != "" {
			errMsg += fmt.Sprintf("\nstderr:\n%s", stderrStr)
		}
		return errors.New(errMsg)
	}
	return nil
}

func generate() (ferr error) {
	cmd := exec.Command("go", "generate", "./...")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		errMsg := fmt.Sprintf("'go generate ./...' failed with error: %v", err)
		if stderrStr := stderr.String(); stderrStr != "" {
			errMsg += fmt.Sprintf("\nstderr:\n%s", stderrStr)
		}
		return errors.New(errMsg)
	}
	return nil
}

func writeSpec(spec string) (ferr error) {
	filePath := "api-docs.json"
	if err := os.WriteFile(filePath, []byte(spec), 0644); err != nil {
		return fmt.Errorf("failed to write spec file '%s': %w", filePath, err)
	}
	return nil
}

func replaceInSpec(spec string, replacements []specReplacement) string {
	result := spec
	for _, replacement := range replacements {
		result = strings.ReplaceAll(result, replacement.old, replacement.new)
	}
	return result
}

func pullSpec(ctx context.Context, buf io.Writer) (ferr error) {
	client := httpClient()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ApiDocUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("HTTP request timed out: %w", err)
		}
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("HTTP request cancelled: %w", err)
		}
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			ferr = errors.Join(ferr, fmt.Errorf("failed to close response body: %w", cerr))
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}

	if _, err = io.Copy(buf, resp.Body); err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	return nil
}

func httpClient() *http.Client {
	transport := &http.Transport{
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

	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second,
	}

	return client
}
