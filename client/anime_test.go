package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/frisbm/gojikan"
)

func TestClient_GetAnimeFullById(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    gojikan.AnimeFull
		wantErr error
	}{
		{
			name: "success - Cowboy Bebop",
			id:   1,
			want: gojikan.AnimeFull{
				Aired: &gojikan.Daterange{
					From: "1998-04-03T00:00:00+00:00",
					To:   "1999-04-24T00:00:00+00:00",
				},
				Airing:     false,
				Background: "When Cowboy Bebop first aired in spring of 1998 on TV Tokyo, only episodes 2-3, 7-15, and 18 were broadcast, it was concluded with a recap special known as Yose Atsume Blues. This was due to anime censorship having increased following the big controversies over Evangelion, as a result most of the series was pulled from the air due to violent content. Satellite channel WOWOW picked up the series in the fall of that year and aired it in its entirety uncensored. Cowboy Bebop was not a ratings hit in Japan, but sold over 19,000 DVD units in the initial release run, and 81,000 overall. Protagonist Spike Spiegel won Best Male Character, and Megumi Hayashibara won Best Voice Actor for her role as Faye Valentine in the 1999 and 2000 Anime Grand Prix, respectively. Cowboy Bebop's biggest influence has been in the United States, where it premiered on Adult Swim in 2001 with many reruns since. The show's heavy Western influence struck a chord with American viewers, where it became a \"gateway drug\" to anime aimed at adult audiences.",
				Episodes:   26,
				Season:     gojikan.SeasonSpring,
				Title:      "Cowboy Bebop",
				Url:        "https://myanimelist.net/anime/1/Cowboy_Bebop",
				Year:       1998,
			},
		},
		{
			name: "success - Steins;Gate",
			id:   9253,
			want: gojikan.AnimeFull{
				Aired: &gojikan.Daterange{
					From: "2011-04-06T00:00:00+00:00",
					To:   "2011-09-14T00:00:00+00:00",
				},
				Airing:     false,
				Background: "Steins;Gate is based on 5pb. and Nitroplus' visual novel of the same title released in 2009. It serves as the second entry in the Science Adventure series.",
				Episodes:   24,
				Season:     gojikan.SeasonSpring,
				Title:      "Steins;Gate",
				Url:        "https://myanimelist.net/anime/9253/Steins_Gate",
				Year:       2011,
			},
		},
		{
			name: "error - not found",
			id:   2,
			wantErr: &gojikan.ErrorResponse{
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
			c, err := New()
			require.NoError(t, err)
			ctx := context.Background()
			got, err := c.GetAnimeFullById(ctx, tt.id)
			if err != nil {
				require.NotNil(t, tt.wantErr, "unexpected error: %v", err)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.Nil(t, tt.wantErr, "expected no error")
			require.Equal(t, tt.want.Aired.From, got.Aired.From)
			require.Equal(t, tt.want.Aired.To, got.Aired.To)
			require.Equal(t, tt.want.Airing, got.Airing)
			require.Equal(t, tt.want.Background, got.Background)
			require.Equal(t, tt.want.Episodes, got.Episodes)
			require.Equal(t, tt.want.Season, got.Season)
			require.Equal(t, tt.want.Title, got.Title)
			require.Equal(t, tt.want.Url, got.Url)
			require.Equal(t, tt.want.Year, got.Year)
		})
	}
}
