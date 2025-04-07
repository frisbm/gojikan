package gojikan

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetAnimeFullById(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    AnimeFull
		wantErr error
	}{
		{
			name: "success - Steins;Gate",
			id:   9253,
			want: AnimeFull{
				MalId: 9253,
				Url:   "https://myanimelist.net/anime/9253/Steins_Gate",
				Images: &AnimeImages{
					Jpg: &Jpg{
						ImageUrl:      "https://cdn.myanimelist.net/images/anime/1935/127974.jpg",
						SmallImageUrl: "https://cdn.myanimelist.net/images/anime/1935/127974t.jpg",
						LargeImageUrl: "https://cdn.myanimelist.net/images/anime/1935/127974l.jpg",
					},
					Webp: &Webp{
						ImageUrl:      "https://cdn.myanimelist.net/images/anime/1935/127974.webp",
						SmallImageUrl: "https://cdn.myanimelist.net/images/anime/1935/127974t.webp",
						LargeImageUrl: "https://cdn.myanimelist.net/images/anime/1935/127974l.webp",
					},
				},
				Trailer: &TrailerBase{
					YoutubeId: "27OZc-ku6is",
					Url:       "https://www.youtube.com/watch?v=27OZc-ku6is",
					EmbedUrl:  "https://www.youtube.com/embed/27OZc-ku6is?enablejsapi=1&wmode=opaque&autoplay=1",
				},
				Approved: true,
				Titles: []Title{
					Title{
						Type:  "Default",
						Title: "Steins;Gate",
					},
					Title{
						Type:  "Japanese",
						Title: "STEINS;GATE",
					},
					Title{
						Type:  "English",
						Title: "Steins;Gate",
					},
				},
				Title:         "Steins;Gate",
				TitleEnglish:  "Steins;Gate",
				TitleJapanese: "STEINS;GATE",
				TitleSynonyms: []string{},
				Type:          "TV",
				Source:        "Visual novel",
				Episodes:      24,
				Status:        "Finished Airing",
				Airing:        false,
				Aired: &Daterange{
					From: "2011-04-06T00:00:00+00:00",
					To:   "2011-09-14T00:00:00+00:00",
					Prop: &Prop{
						From: &Date{
							Day:   6,
							Month: 4,
							Year:  2011,
						},
						To: &Date{
							Day:   14,
							Month: 9,
							Year:  2011,
						},
						String: "",
					},
				},
				Duration:   "24 min per ep",
				Rating:     "PG-13 - Teens 13 or older",
				Score:      9.07,
				ScoredBy:   1456435,
				Rank:       3,
				Popularity: 14,
				Members:    2681698,
				Favorites:  195221,
				Synopsis:   "Eccentric scientist Rintarou Okabe has a never-ending thirst for scientific exploration. Together with his ditzy but well-meaning friend Mayuri Shiina and his roommate Itaru Hashida, Okabe founds the Future Gadget Laboratory in the hopes of creating technological innovations that baffle the human psyche. Despite claims of grandeur, the only notable \"gadget\" the trio have created is a microwave that has the mystifying power to turn bananas into green goo.\n\nHowever, when Okabe attends a conference on time travel, he experiences a series of strange events that lead him to believe that there is more to the \"Phone Microwave\" gadget than meets the eye. Apparently able to send text messages into the past using the microwave, Okabe dabbles further with the \"time machine,\" attracting the ire and attention of the mysterious organization SERN.\n\nDue to the novel discovery, Okabe and his friends find themselves in an ever-present danger. As he works to mitigate the damage his invention has caused to the timeline, Okabe fights a battle to not only save his loved ones but also to preserve his degrading sanity.\n\n[Written by MAL Rewrite]",
				Background: "Steins;Gate is based on 5pb. and Nitroplus' visual novel of the same title released in 2009. It serves as the second entry in the Science Adventure series.",
				Season:     "spring",
				Year:       2011,
				Broadcast: &Broadcast{
					Day:      "Wednesdays",
					Time:     "02:05",
					Timezone: "Asia/Tokyo",
					String:   "Wednesdays at 02:05 (JST)",
				},
				Producers: []MalUrl{
					MalUrl{
						MalId: 61,
						Type:  "anime",
						Name:  "Frontier Works",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/61/Frontier_Works",
					},
					MalUrl{
						MalId: 108,
						Type:  "anime",
						Name:  "Media Factory",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/108/Media_Factory",
					},
					MalUrl{
						MalId: 113,
						Type:  "anime",
						Name:  "Kadokawa Shoten",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/113/Kadokawa_Shoten",
					},
					MalUrl{
						MalId: 166,
						Type:  "anime",
						Name:  "Movic",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/166/Movic",
					},
					MalUrl{
						MalId: 238,
						Type:  "anime",
						Name:  "AT-X",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/238/AT-X",
					},
					MalUrl{
						MalId: 352,
						Type:  "anime",
						Name:  "Kadokawa Pictures Japan",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/352/Kadokawa_Pictures_Japan",
					},
					MalUrl{
						MalId: 459,
						Type:  "anime",
						Name:  "Nitroplus",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/459/Nitroplus",
					},
				},
				Licensors: []MalUrl{
					MalUrl{
						MalId: 102,
						Type:  "anime",
						Name:  "Funimation",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/102/Funimation",
					},
				},
				Studios: []MalUrl{
					MalUrl{
						MalId: 314,
						Type:  "anime",
						Name:  "White Fox",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/314/White_Fox",
					},
				},
				Genres: []MalUrl{
					MalUrl{
						MalId: 8,
						Type:  "anime",
						Name:  "Drama",
						Title: "",
						Url:   "https://myanimelist.net/anime/genre/8/Drama",
					},
					MalUrl{
						MalId: 24,
						Type:  "anime",
						Name:  "Sci-Fi",
						Title: "",
						Url:   "https://myanimelist.net/anime/genre/24/Sci-Fi",
					},
					MalUrl{
						MalId: 41,
						Type:  "anime",
						Name:  "Suspense",
						Title: "",
						Url:   "https://myanimelist.net/anime/genre/41/Suspense",
					},
				},
				ExplicitGenres: []MalUrl{}, // p0
				Themes: []MalUrl{
					MalUrl{
						MalId: 40,
						Type:  "anime",
						Name:  "Psychological",
						Title: "",
						Url:   "https://myanimelist.net/anime/genre/40/Psychological",
					},
					MalUrl{
						MalId: 78,
						Type:  "anime",
						Name:  "Time Travel",
						Title: "",
						Url:   "https://myanimelist.net/anime/genre/78/Time_Travel",
					},
				},
				Demographics: []MalUrl{},
				Relations: []AnimeFullRelations{
					AnimeFullRelations{
						Relation: "Sequel",
						Entry: []MalUrl{
							MalUrl{
								MalId: 10863,
								Type:  "anime",
								Name:  "Steins;Gate: Oukoubakko no Poriomania",
								Title: "",
								Url:   "https://myanimelist.net/anime/10863/Steins_Gate__Oukoubakko_no_Poriomania",
							},
						},
					},
					AnimeFullRelations{
						Relation: "Adaptation",
						Entry: []MalUrl{
							MalUrl{
								MalId: 17517,
								Type:  "manga",
								Name:  "Steins;Gate",
								Title: "",
								Url:   "https://myanimelist.net/manga/17517/Steins_Gate",
							},
						},
					},
					AnimeFullRelations{
						Relation: "Alternative Version",
						Entry: []MalUrl{
							MalUrl{
								MalId: 32188,
								Type:  "anime",
								Name:  "Steins;Gate: Kyoukaimenjou no Missing Link - Divide By Zero",
								Title: "",
								Url:   "https://myanimelist.net/anime/32188/Steins_Gate__Kyoukaimenjou_no_Missing_Link_-_Divide_By_Zero",
							},
						},
					},
					AnimeFullRelations{
						Relation: "Alternative Setting",
						Entry: []MalUrl{
							MalUrl{
								MalId: 4975,
								Type:  "anime",
								Name:  "ChäoS;HEAd",
								Title: "",
								Url:   "https://myanimelist.net/anime/4975/ChäoS_HEAd",
							},
							MalUrl{
								MalId: 13599,
								Type:  "anime",
								Name:  "Robotics;Notes",
								Title: "",
								Url:   "https://myanimelist.net/anime/13599/Robotics_Notes",
							},
							MalUrl{
								MalId: 32962,
								Type:  "anime",
								Name:  "Occultic;Nine",
								Title: "",
								Url:   "https://myanimelist.net/anime/32962/Occultic_Nine",
							},
							MalUrl{
								MalId: 30485,
								Type:  "anime",
								Name:  "ChäoS;Child",
								Title: "",
								Url:   "https://myanimelist.net/anime/30485/ChäoS_Child",
							},
						},
					},
					AnimeFullRelations{
						Relation: "Other",
						Entry: []MalUrl{
							MalUrl{
								MalId: 27957,
								Type:  "anime",
								Name:  "Steins;Gate: Soumei Eichi no Cognitive Computing",
								Title: "",
								Url:   "https://myanimelist.net/anime/27957/Steins_Gate__Soumei_Eichi_no_Cognitive_Computing",
							},
						},
					},
				},
				Theme: &AnimeFullTheme{
					Openings: []string{
						"\"Hacking to the Gate\" by Kanako Itou",
					},
					Endings: []string{
						"1: \"Toki Tsukasadoru Juuni no Meiyaku (刻司ル十二ノ盟約)\" by Yui Sakakibara (eps 1-21)",
						"2: \"Fake Verthandi\" by Takeshi Abo (eps 22)",
						"3: \"Sky Clad no Kansokusha (スカイクラッドの観測者)\" by Kanako Itou (eps 23)",
						"4: \"Another Heaven\" by Kanako Itou (eps 24)",
					},
				},
				External: []AnimeFullExternal{
					AnimeFullExternal{
						Name: "Official Site",
						Url:  "http://www.steinsgate.tv/",
					},
					AnimeFullExternal{
						Name: "@sg_anime",
						Url:  "https://twitter.com/sg_anime",
					},
					AnimeFullExternal{
						Name: "AniDB",
						Url:  "https://anidb.net/perl-bin/animedb.pl?show=anime&aid=7729",
					},
					AnimeFullExternal{
						Name: "ANN",
						Url:  "https://www.animenewsnetwork.com/encyclopedia/anime.php?id=11770",
					},
					AnimeFullExternal{
						Name: "Wikipedia",
						Url:  "http://en.wikipedia.org/wiki/Steins;Gate",
					},
					AnimeFullExternal{
						Name: "Wikipedia",
						Url:  "https://ja.wikipedia.org/wiki/STEINS;GATE",
					},
					AnimeFullExternal{
						Name: "Syoboi",
						Url:  "https://cal.syoboi.jp/tid/2142",
					},
				},
				Streaming: []AnimeFullExternal{
					AnimeFullExternal{
						Name: "Crunchyroll",
						Url:  "http://www.crunchyroll.com/series-229050",
					},
					AnimeFullExternal{
						Name: "Netflix",
						Url:  "https://www.netflix.com/",
					},
				},
			},
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
			require.Equal(t, tt.want, got)
		})
	}
}

func TestClient_GetAnimeById(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    Anime
		wantErr error
	}{
		{
			name: "success - Steins;Gate",
			id:   9253,
			want: Anime{
				MalId: 9253,
				Url:   "https://myanimelist.net/anime/9253/Steins_Gate",
				Images: &AnimeImages{
					Jpg: &Jpg{
						ImageUrl:      "https://cdn.myanimelist.net/images/anime/1935/127974.jpg",
						SmallImageUrl: "https://cdn.myanimelist.net/images/anime/1935/127974t.jpg",
						LargeImageUrl: "https://cdn.myanimelist.net/images/anime/1935/127974l.jpg",
					},
					Webp: &Webp{
						ImageUrl:      "https://cdn.myanimelist.net/images/anime/1935/127974.webp",
						SmallImageUrl: "https://cdn.myanimelist.net/images/anime/1935/127974t.webp",
						LargeImageUrl: "https://cdn.myanimelist.net/images/anime/1935/127974l.webp",
					},
				},
				Trailer: &TrailerBase{
					YoutubeId: "27OZc-ku6is",
					Url:       "https://www.youtube.com/watch?v=27OZc-ku6is",
					EmbedUrl:  "https://www.youtube.com/embed/27OZc-ku6is?enablejsapi=1&wmode=opaque&autoplay=1",
				},
				Approved: true,
				Titles: []Title{
					Title{
						Type:  "Default",
						Title: "Steins;Gate",
					},
					Title{
						Type:  "Japanese",
						Title: "STEINS;GATE",
					},
					Title{
						Type:  "English",
						Title: "Steins;Gate",
					},
				},
				Title:         "Steins;Gate",
				TitleEnglish:  "Steins;Gate",
				TitleJapanese: "STEINS;GATE",
				TitleSynonyms: []string{},
				Type:          "TV",
				Source:        "Visual novel",
				Episodes:      24,
				Status:        "Finished Airing",
				Airing:        false,
				Aired: &Daterange{
					From: "2011-04-06T00:00:00+00:00",
					To:   "2011-09-14T00:00:00+00:00",
					Prop: &Prop{
						From: &Date{
							Day:   6,
							Month: 4,
							Year:  2011,
						},
						To: &Date{
							Day:   14,
							Month: 9,
							Year:  2011,
						},
						String: "",
					},
				},
				Duration:   "24 min per ep",
				Rating:     "PG-13 - Teens 13 or older",
				Score:      9.07,
				ScoredBy:   1456435,
				Rank:       3,
				Popularity: 14,
				Members:    2681698,
				Favorites:  195221,
				Synopsis:   "Eccentric scientist Rintarou Okabe has a never-ending thirst for scientific exploration. Together with his ditzy but well-meaning friend Mayuri Shiina and his roommate Itaru Hashida, Okabe founds the Future Gadget Laboratory in the hopes of creating technological innovations that baffle the human psyche. Despite claims of grandeur, the only notable \"gadget\" the trio have created is a microwave that has the mystifying power to turn bananas into green goo.\n\nHowever, when Okabe attends a conference on time travel, he experiences a series of strange events that lead him to believe that there is more to the \"Phone Microwave\" gadget than meets the eye. Apparently able to send text messages into the past using the microwave, Okabe dabbles further with the \"time machine,\" attracting the ire and attention of the mysterious organization SERN.\n\nDue to the novel discovery, Okabe and his friends find themselves in an ever-present danger. As he works to mitigate the damage his invention has caused to the timeline, Okabe fights a battle to not only save his loved ones but also to preserve his degrading sanity.\n\n[Written by MAL Rewrite]",
				Background: "Steins;Gate is based on 5pb. and Nitroplus' visual novel of the same title released in 2009. It serves as the second entry in the Science Adventure series.",
				Season:     "spring",
				Year:       2011,
				Broadcast: &Broadcast{
					Day:      "Wednesdays",
					Time:     "02:05",
					Timezone: "Asia/Tokyo",
					String:   "Wednesdays at 02:05 (JST)",
				},
				Producers: []MalUrl{
					MalUrl{
						MalId: 61,
						Type:  "anime",
						Name:  "Frontier Works",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/61/Frontier_Works",
					},
					MalUrl{
						MalId: 108,
						Type:  "anime",
						Name:  "Media Factory",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/108/Media_Factory",
					},
					MalUrl{
						MalId: 113,
						Type:  "anime",
						Name:  "Kadokawa Shoten",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/113/Kadokawa_Shoten",
					},
					MalUrl{
						MalId: 166,
						Type:  "anime",
						Name:  "Movic",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/166/Movic",
					},
					MalUrl{
						MalId: 238,
						Type:  "anime",
						Name:  "AT-X",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/238/AT-X",
					},
					MalUrl{
						MalId: 352,
						Type:  "anime",
						Name:  "Kadokawa Pictures Japan",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/352/Kadokawa_Pictures_Japan",
					},
					MalUrl{
						MalId: 459,
						Type:  "anime",
						Name:  "Nitroplus",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/459/Nitroplus",
					},
				},
				Licensors: []MalUrl{
					MalUrl{
						MalId: 102,
						Type:  "anime",
						Name:  "Funimation",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/102/Funimation",
					},
				},
				Studios: []MalUrl{
					MalUrl{
						MalId: 314,
						Type:  "anime",
						Name:  "White Fox",
						Title: "",
						Url:   "https://myanimelist.net/anime/producer/314/White_Fox",
					},
				},
				Genres: []MalUrl{
					MalUrl{
						MalId: 8,
						Type:  "anime",
						Name:  "Drama",
						Title: "",
						Url:   "https://myanimelist.net/anime/genre/8/Drama",
					},
					MalUrl{
						MalId: 24,
						Type:  "anime",
						Name:  "Sci-Fi",
						Title: "",
						Url:   "https://myanimelist.net/anime/genre/24/Sci-Fi",
					},
					MalUrl{
						MalId: 41,
						Type:  "anime",
						Name:  "Suspense",
						Title: "",
						Url:   "https://myanimelist.net/anime/genre/41/Suspense",
					},
				},
				ExplicitGenres: []MalUrl{}, // p0
				Themes: []MalUrl{
					MalUrl{
						MalId: 40,
						Type:  "anime",
						Name:  "Psychological",
						Title: "",
						Url:   "https://myanimelist.net/anime/genre/40/Psychological",
					},
					MalUrl{
						MalId: 78,
						Type:  "anime",
						Name:  "Time Travel",
						Title: "",
						Url:   "https://myanimelist.net/anime/genre/78/Time_Travel",
					},
				},
				Demographics: []MalUrl{},
			},
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
			c, err := New()
			require.NoError(t, err)
			ctx := context.Background()
			got, err := c.GetAnimeById(ctx, tt.id)
			if err != nil {
				require.NotNil(t, tt.wantErr, "unexpected error: %v", err)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.Nil(t, tt.wantErr, "expected no error")
			require.Equal(t, tt.want, got)
		})
	}
}

func TestClient_GetAnimeCharacters(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		wantLen int
		want    AnimeCharacters
		wantErr error
	}{
		{
			name:    "success - Steins;Gate",
			id:      9253,
			wantLen: 24,
			want: AnimeCharacters{
				Character: &AnimeCharactersCharacter{
					MalId: 35258,
					Url:   "https://myanimelist.net/character/35258/Itaru_Hashida",
					Images: &CharacterImages{
						Jpg: &CharacterImagesJpg{
							ImageUrl:      "https://cdn.myanimelist.net/images/characters/6/113767.jpg?s=5d160f99286a0891c5e32413a5438622",
							SmallImageUrl: "",
						},
						Webp: &CharacterImagesWebp{
							ImageUrl:      "https://cdn.myanimelist.net/images/characters/6/113767.webp?s=5d160f99286a0891c5e32413a5438622",
							SmallImageUrl: "https://cdn.myanimelist.net/images/characters/6/113767t.webp?s=5d160f99286a0891c5e32413a5438622",
						},
					},
					Name: "Hashida, Itaru",
				},
				Role: "Main",
				VoiceActors: []AnimeCharactersVoiceActors{
					AnimeCharactersVoiceActors{
						Person: &AnimeCharactersPerson{
							MalId: 1,
							Url:   "https://myanimelist.net/people/1/Tomokazu_Seki",
							Images: &PeopleImages{
								Jpg: &Jpg{
									ImageUrl:      "https://cdn.myanimelist.net/images/voiceactors/1/55486.jpg?s=23988e0eb96abf9ac6389a3f7b4b4659",
									SmallImageUrl: "",
									LargeImageUrl: "",
								},
							},
							Name: "Seki, Tomokazu",
						},
						Language: "Japanese",
					},
					AnimeCharactersVoiceActors{
						Person: &AnimeCharactersPerson{
							MalId: 19069,
							Url:   "https://myanimelist.net/people/19069/Tyson_Rinehart",
							Images: &PeopleImages{
								Jpg: &Jpg{
									ImageUrl:      "https://cdn.myanimelist.net/images/voiceactors/1/83649.jpg?s=786c4319139f7280e4c8fc13ac926826",
									SmallImageUrl: "",
									LargeImageUrl: "",
								},
							},
							Name: "Rinehart, Tyson",
						},
						Language: "English",
					},
					AnimeCharactersVoiceActors{
						Person: &AnimeCharactersPerson{
							MalId: 20558,
							Url:   "https://myanimelist.net/people/20558/Paolo_Vivio",
							Images: &PeopleImages{
								Jpg: &Jpg{
									ImageUrl:      "https://cdn.myanimelist.net/images/voiceactors/2/23154.jpg?s=f2a23bba7a45b280351918361560c5a8",
									SmallImageUrl: "",
									LargeImageUrl: "",
								},
							},
							Name: "Vivio, Paolo",
						},
						Language: "Italian",
					},
					AnimeCharactersVoiceActors{
						Person: &AnimeCharactersPerson{
							MalId: 22365,
							Url:   "https://myanimelist.net/people/22365/Jesco_Wirthgen",
							Images: &PeopleImages{
								Jpg: &Jpg{
									ImageUrl:      "https://cdn.myanimelist.net/images/voiceactors/3/25773.jpg?s=ec5f621e2e13fa04e0c23bae727ed122",
									SmallImageUrl: "",
									LargeImageUrl: "",
								},
							},
							Name: "Wirthgen, Jesco",
						},
						Language: "German",
					},
					AnimeCharactersVoiceActors{
						Person: &AnimeCharactersPerson{
							MalId: 28021,
							Url:   "https://myanimelist.net/people/28021/Laurent_Pasquier",
							Images: &PeopleImages{
								Jpg: &Jpg{
									ImageUrl:      "https://cdn.myanimelist.net/images/voiceactors/1/70233.jpg?s=152b3d5d7556f1d88246d0cca4386991",
									SmallImageUrl: "",
									LargeImageUrl: "",
								},
							},
							Name: "Pasquier, Laurent",
						},
						Language: "French",
					},
					AnimeCharactersVoiceActors{
						Person: &AnimeCharactersPerson{
							MalId: 55742,
							Url:   "https://myanimelist.net/people/55742/Bruno_Casemiro",
							Images: &PeopleImages{
								Jpg: &Jpg{
									ImageUrl:      "https://cdn.myanimelist.net/images/voiceactors/1/74419.jpg?s=99f5b9cd255c6b996a9de4417241d665",
									SmallImageUrl: "",
									LargeImageUrl: "",
								},
							},
							Name: "Casemiro, Bruno",
						},
						Language: "Portuguese (BR)",
					},
					AnimeCharactersVoiceActors{
						Person: &AnimeCharactersPerson{
							MalId: 59217,
							Url:   "https://myanimelist.net/people/59217/Iván_Fernández",
							Images: &PeopleImages{
								Jpg: &Jpg{
									ImageUrl:      "https://cdn.myanimelist.net/images/voiceactors/2/70076.jpg?s=605019d10a4d89189111c02d17547752",
									SmallImageUrl: "",
									LargeImageUrl: "",
								},
							},
							Name: "Fernández, Iván",
						},
						Language: "Spanish",
					},
					AnimeCharactersVoiceActors{
						Person: &AnimeCharactersPerson{
							MalId: 7819,
							Url:   "https://myanimelist.net/people/7819/Szabolcs_Seszták",
							Images: &PeopleImages{
								Jpg: &Jpg{
									ImageUrl:      "https://cdn.myanimelist.net/images/voiceactors/3/5071.jpg?s=9db58eb08fd4c54c182e6f16736402df",
									SmallImageUrl: "",
									LargeImageUrl: "",
								},
							},
							Name: "Seszták, Szabolcs",
						},
						Language: "Hungarian",
					},
				},
			},
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
			c, err := New()
			require.NoError(t, err)
			ctx := context.Background()
			got, err := c.GetAnimeCharacters(ctx, tt.id)
			if err != nil {
				require.NotNil(t, tt.wantErr, "unexpected error: %v", err)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.Nil(t, tt.wantErr, "expected no error")
			require.Equal(t, tt.wantLen, len(got))
			require.Equal(t, tt.want, got[0])
		})
	}
}

func TestClient_GetAnimeStaff(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		wantLen int
		want    AnimeStaff
		wantErr error
	}{
		{
			name:    "success - Steins;Gate",
			id:      9253,
			wantLen: 160,
			want: AnimeStaff{
				Person: &AnimeStaffPerson{
					MalId: 203,
					Url:   "https://myanimelist.net/people/203/Justin_Cook",
					Images: &PeopleImages{
						Jpg: &Jpg{
							ImageUrl:      "https://cdn.myanimelist.net/images/voiceactors/1/80501.jpg?s=ee808f428c434aee43fe003623f6bf0b",
							SmallImageUrl: "",
							LargeImageUrl: "",
						},
					},
					Name: "Cook, Justin",
				},
				Positions: []string{
					"Producer",
				},
			},
		},
		{
			name: "error - not found",
			id:   2,
			wantErr: &ErrorResponse{
				Status:    404,
				Type:      "BadResponseException",
				Message:   "Resource does not exist",
				ApiError:  "404 on https://myanimelist.net/anime/2/_/stagg",
				ReportUrl: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := New()
			require.NoError(t, err)
			ctx := context.Background()
			got, err := c.GetAnimeStaff(ctx, tt.id)
			if err != nil {
				require.NotNil(t, tt.wantErr, "unexpected error: %v", err)
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.Nil(t, tt.wantErr, "expected no error")
			require.Equal(t, tt.wantLen, len(got))
			require.Equal(t, tt.want, got[0])
		})
	}
}
