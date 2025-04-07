package gojikan

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var testClient = func() *Client {
	client, err := New(WithRateLimit(1, time.Second, 1))
	if err != nil {
		panic(fmt.Errorf("create client: %w", err))
	}
	return client
}()

func TestClient_GetAnimeFullById(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		wantErr error
	}{
		{
			name: "success - Steins;Gate",
			id:   9253,
		},
		{
			name: "error - not found",
			id:   2,
			wantErr: &ErrorResponse{
				Status:    404,
				Type:      "BadResponseException",
				Message:   "Resource does not exist",
				ApiError:  "404 on https://myanimelist.net/anime/2/",
				ReportUrl: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := testClient.GetAnimeFullById(ctx, tt.id)
			if err != nil {
				require.NotNil(t, tt.wantErr, "unexpected error: %v", err)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.Nil(t, tt.wantErr, "expected no error")
			require.NotZero(t, got)
		})
	}
}

func TestClient_GetAnimeById(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		wantErr error
	}{
		{
			name: "success - Steins;Gate",
			id:   9253,
		},
		{
			name: "error - not found",
			id:   2,
			wantErr: &ErrorResponse{
				Status:    404,
				Type:      "BadResponseException",
				Message:   "Resource does not exist",
				ApiError:  "404 on https://myanimelist.net/anime/2/",
				ReportUrl: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := testClient.GetAnimeById(ctx, tt.id)
			if err != nil {
				require.NotNil(t, tt.wantErr, "unexpected error: %v", err)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.Nil(t, tt.wantErr, "expected no error")
			require.NotZero(t, got)
		})
	}
}

func TestClient_GetAnimeCharacters(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		wantLen int
		wantErr error
	}{
		{
			name:    "success - Steins;Gate",
			id:      9253,
			wantLen: 27,
		},
		{
			name: "error - not found",
			id:   2,
			wantErr: &ErrorResponse{
				Status:    404,
				Type:      "BadResponseException",
				Message:   "Resource does not exist",
				ApiError:  "404 on https://myanimelist.net/anime/2/_/characters",
				ReportUrl: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := testClient.GetAnimeCharacters(ctx, tt.id)
			if err != nil {
				require.NotNil(t, tt.wantErr, "unexpected error: %v", err)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.Nil(t, tt.wantErr, "expected no error")
			require.Equal(t, tt.wantLen, len(got))
			require.NotZero(t, got)
		})
	}
}

func TestClient_GetAnimeStaff(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		wantLen int
		wantErr error
	}{
		{
			name:    "success - Steins;Gate",
			id:      9253,
			wantLen: 160,
		},
		{
			name: "error - not found",
			id:   2,
			wantErr: &ErrorResponse{
				Status:    404,
				Type:      "BadResponseException",
				Message:   "Resource does not exist",
				ApiError:  "404 on https://myanimelist.net/anime/2/_/characters",
				ReportUrl: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := testClient.GetAnimeStaff(ctx, tt.id)
			if err != nil {
				require.NotNil(t, tt.wantErr, "unexpected error: %v", err)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.Nil(t, tt.wantErr, "expected no error")
			require.Equal(t, tt.wantLen, len(got))
			require.NotZero(t, got)
		})
	}
}
