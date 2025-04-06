package gojikan

import (
	"context"
	"fmt"

	"github.com/frisbm/gojikan/internal"
)

func (c *Client) GetAnimeFullById(ctx context.Context, id int) (AnimeFull, error) {
	pathParams := map[string]any{
		"id": id,
	}
	queryParams := map[string][]any{}
	url, err := internal.BuildUrl(c.baseURL, getAnimeFullByIdPath, pathParams, queryParams)
	if err != nil {
		return AnimeFull{}, fmt.Errorf("failed to build url: %w", err)
	}
	resp, err := get[Response[AnimeFull]](ctx, c, url)
	if err != nil {
		return AnimeFull{}, fmt.Errorf("failed to get anime full by id: %w", err)
	}
	return resp.Data, nil
}

func (c *Client) GetAnimeById(ctx context.Context, id int) (Anime, error) {
	pathParams := map[string]any{
		"id": id,
	}
	queryParams := map[string][]any{}
	url, err := internal.BuildUrl(c.baseURL, getAnimeByIdPath, pathParams, queryParams)
	if err != nil {
		return Anime{}, fmt.Errorf("failed to build url: %w", err)
	}
	resp, err := get[Response[Anime]](ctx, c, url)
	if err != nil {
		return Anime{}, fmt.Errorf("failed to get anime by id: %w", err)
	}
	return resp.Data, nil
}
