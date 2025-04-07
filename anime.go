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

func (c *Client) GetAnimeCharacters(ctx context.Context, id int) ([]AnimeCharacters, error) {
	pathParams := map[string]any{
		"id": id,
	}
	queryParams := map[string][]any{}
	url, err := internal.BuildUrl(c.baseURL, getAnimeCharactersPath, pathParams, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to build url: %w", err)
	}
	resp, err := get[Response[[]AnimeCharacters]](ctx, c, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get anime characters: %w", err)
	}
	return resp.Data, nil
}
func (c *Client) GetAnimeStaff(ctx context.Context, id int) ([]AnimeStaff, error) {
	pathParams := map[string]any{
		"id": id,
	}
	queryParams := map[string][]any{}
	url, err := internal.BuildUrl(c.baseURL, getAnimeStaffPath, pathParams, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to build url: %w", err)
	}
	resp, err := get[Response[[]AnimeStaff]](ctx, c, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get anime staff: %w", err)
	}
	return resp.Data, nil
}
func (c *Client) GetAnimeEpisodes(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeEpisodeById(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeNews(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeForum(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeVideos(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeVideosEpisodes(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimePictures(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeStatistics(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeMoreInfo(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeRecommendations(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeUserUpdates(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeReviews(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeRelations(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeThemes(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeExternal(ctx context.Context) (any, error) {
	return nil, nil
}
func (c *Client) GetAnimeStreaming(ctx context.Context) (any, error) {
	return nil, nil
}
