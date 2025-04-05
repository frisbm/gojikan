package gojikan

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func ptr[T any](v T) *T {
	return &v
}

func TestClientWithResponses_GetAnimeByIdWithResponse(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    *GetAnimeByIdResponse
		wantErr bool
	}{
		{
			name: "success - Cowboy Bebop",
			id:   1,
			want: &GetAnimeByIdResponse{
				JSON200: &GetAnimeByIdResponseJSON200{
					Data: &Anime{
						Aired: &Daterange{
							From: ptr("1998-04-03T00:00:00+00:00"),
							To:   ptr("1999-04-24T00:00:00+00:00"),
						},
						Airing:     ptr(false),
						Background: ptr("When Cowboy Bebop first aired in spring of 1998 on TV Tokyo, only episodes 2-3, 7-15, and 18 were broadcast, it was concluded with a recap special known as Yose Atsume Blues. This was due to anime censorship having increased following the big controversies over Evangelion, as a result most of the series was pulled from the air due to violent content. Satellite channel WOWOW picked up the series in the fall of that year and aired it in its entirety uncensored. Cowboy Bebop was not a ratings hit in Japan, but sold over 19,000 DVD units in the initial release run, and 81,000 overall. Protagonist Spike Spiegel won Best Male Character, and Megumi Hayashibara won Best Voice Actor for her role as Faye Valentine in the 1999 and 2000 Anime Grand Prix, respectively. Cowboy Bebop's biggest influence has been in the United States, where it premiered on Adult Swim in 2001 with many reruns since. The show's heavy Western influence struck a chord with American viewers, where it became a \"gateway drug\" to anime aimed at adult audiences."),
						Episodes:   ptr(26),
						Season:     ptr(AnimeSeasonSpring),
						Title:      ptr("Cowboy Bebop"),
						Url:        ptr("https://myanimelist.net/anime/1/Cowboy_Bebop"),
						Year:       ptr(1998),
					},
				},
			},
		},
		{
			name: "success - Steins;Gate",
			id:   9253,
			want: &GetAnimeByIdResponse{
				JSON200: &GetAnimeByIdResponseJSON200{
					Data: &Anime{
						Aired: &Daterange{
							From: ptr("2011-04-06T00:00:00+00:00"),
							To:   ptr("2011-09-14T00:00:00+00:00"),
						},
						Airing:     ptr(false),
						Background: ptr("Steins;Gate is based on 5pb. and Nitroplus' visual novel of the same title released in 2009. It serves as the second entry in the Science Adventure series."),
						Episodes:   ptr(24),
						Season:     ptr(AnimeSeasonSpring),
						Title:      ptr("Steins;Gate"),
						Url:        ptr("https://myanimelist.net/anime/9253/Steins_Gate"),
						Year:       ptr(2011),
					},
				},
			},
		},
		{
			name:    "error - Not Found",
			id:      2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClientWithResponses("https://api.jikan.moe/v4")
			require.NoError(t, err)
			resp, err := c.GetAnimeByIdWithResponse(context.Background(), tt.id)
			if err != nil {
				require.True(t, tt.wantErr, "unexpected error: %v", err)
				return
			}
			require.False(t, tt.wantErr, "expected no error")
			want := tt.want.JSON200.Data
			got := resp.JSON200.Data
			require.Equal(t, *want.Aired.From, *got.Aired.From)
			require.Equal(t, *want.Aired.To, *got.Aired.To)
			require.Equal(t, *want.Airing, *got.Airing)
			require.Equal(t, *want.Background, *got.Background)
			require.Equal(t, *want.Episodes, *got.Episodes)
			require.Equal(t, *want.Season, *got.Season)
			require.Equal(t, *want.Title, *got.Title)
			require.Equal(t, *want.Url, *got.Url)
			require.Equal(t, *want.Year, *got.Year)
		})
	}
}
