package gojikan

import (
	"errors"
	"fmt"
	"strings"
)

// ErrorResponse represents the error response from the API.
type ErrorResponse struct {
	Status    int    `json:"status"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	ApiError  string `json:"error"`
	ReportUrl string `json:"report_url"`
}

// Error implements the error interface for ErrorResponse.
func (e *ErrorResponse) Error() string {
	var errorParts []string
	if e.Status != 0 {
		errorParts = append(errorParts, fmt.Sprintf("Status: %d", e.Status))
	}
	if e.Type != "" {
		errorParts = append(errorParts, fmt.Sprintf("Type: %s", e.Type))
	}
	if e.Message != "" {
		errorParts = append(errorParts, fmt.Sprintf("Message: %s", e.Message))
	}
	if e.ApiError != "" {
		errorParts = append(errorParts, fmt.Sprintf("ApiError: %s", e.ApiError))
	}
	if e.ReportUrl != "" {
		errorParts = append(errorParts, fmt.Sprintf("ReportUrl: %s", e.ReportUrl))
	}
	return fmt.Sprintf("error response: %s", strings.Join(errorParts, ", "))
}

func (e *ErrorResponse) Is(target error) bool {
	var t *ErrorResponse
	ok := errors.As(target, &t)
	if !ok {
		return false
	}
	return e.Status == t.Status && e.Type == t.Type && e.Message == t.Message && e.ApiError == t.ApiError
}

// Pagination represents the pagination information in the API response.
type Pagination struct {
	LastVisiblePage int  `json:"last_visible_page"`
	HasNextPage     bool `json:"has_next_page"`
}
