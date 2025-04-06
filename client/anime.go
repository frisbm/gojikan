package client

import (
	"context"
	"fmt"

	"github.com/frisbm/gojikan"
	"github.com/frisbm/gojikan/client/internal"
)

const getAnimeFullByIdPath = "/anime/{id}/full"

type GetAnimeFullByIdResponse struct {
	Data gojikan.AnimeFull `json:"data"`
}

func (c *Client) GetAnimeFullById(ctx context.Context, id int) (gojikan.AnimeFull, error) {
	pathParams := map[string]any{
		"id": id,
	}
	queryParams := map[string][]any{}
	url, err := internal.BuildUrl(c.baseURL, getAnimeFullByIdPath, pathParams, queryParams)
	if err != nil {
		return gojikan.AnimeFull{}, fmt.Errorf("failed to build url: %w", err)
	}
	resp, err := get[GetAnimeFullByIdResponse](ctx, c, url)
	if err != nil {
		return gojikan.AnimeFull{}, fmt.Errorf("failed to get anime full by id: %w", err)
	}
	return resp.Data, nil
}
