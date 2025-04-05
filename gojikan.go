package gojikan

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

const (
	AnimeRatingGAllAges             AnimeRating = "G - All Ages"
	AnimeRatingPG13Teens13OrOlder   AnimeRating = "PG-13 - Teens 13 or older"
	AnimeRatingPGChildren           AnimeRating = "PG - Children"
	AnimeRatingR17ViolenceProfanity AnimeRating = "R - 17+ (violence & profanity)"
	AnimeRatingRMildNudity          AnimeRating = "R+ - Mild Nudity"
	AnimeRatingRxHentai             AnimeRating = "Rx - Hentai"
)
const (
	AnimeSeasonFall   AnimeSeason = "fall"
	AnimeSeasonSpring AnimeSeason = "spring"
	AnimeSeasonSummer AnimeSeason = "summer"
	AnimeSeasonWinter AnimeSeason = "winter"
)
const (
	AnimeStatusCurrentlyAiring AnimeStatus = "Currently Airing"
	AnimeStatusFinishedAiring  AnimeStatus = "Finished Airing"
	AnimeStatusNotYetAired     AnimeStatus = "Not yet aired"
)
const (
	AnimeTypeMovie   AnimeType = "Movie"
	AnimeTypeMusic   AnimeType = "Music"
	AnimeTypeONA     AnimeType = "ONA"
	AnimeTypeOVA     AnimeType = "OVA"
	AnimeTypeSpecial AnimeType = "Special"
	AnimeTypeTV      AnimeType = "TV"
)
const (
	AnimeFullRatingGAllAges             AnimeFullRating = "G - All Ages"
	AnimeFullRatingPG13Teens13OrOlder   AnimeFullRating = "PG-13 - Teens 13 or older"
	AnimeFullRatingPGChildren           AnimeFullRating = "PG - Children"
	AnimeFullRatingR17ViolenceProfanity AnimeFullRating = "R - 17+ (violence & profanity)"
	AnimeFullRatingRMildNudity          AnimeFullRating = "R+ - Mild Nudity"
	AnimeFullRatingRxHentai             AnimeFullRating = "Rx - Hentai"
)
const (
	AnimeFullSeasonFall   AnimeFullSeason = "fall"
	AnimeFullSeasonSpring AnimeFullSeason = "spring"
	AnimeFullSeasonSummer AnimeFullSeason = "summer"
	AnimeFullSeasonWinter AnimeFullSeason = "winter"
)
const (
	AnimeFullStatusCurrentlyAiring AnimeFullStatus = "Currently Airing"
	AnimeFullStatusFinishedAiring  AnimeFullStatus = "Finished Airing"
	AnimeFullStatusNotYetAired     AnimeFullStatus = "Not yet aired"
)
const (
	AnimeFullTypeMovie   AnimeFullType = "Movie"
	AnimeFullTypeMusic   AnimeFullType = "Music"
	AnimeFullTypeONA     AnimeFullType = "ONA"
	AnimeFullTypeOVA     AnimeFullType = "OVA"
	AnimeFullTypeSpecial AnimeFullType = "Special"
	AnimeFullTypeTV      AnimeFullType = "TV"
)
const (
	AnimeSearchQueryOrderbyEndDate    AnimeSearchQueryOrderby = "end_date"
	AnimeSearchQueryOrderbyEpisodes   AnimeSearchQueryOrderby = "episodes"
	AnimeSearchQueryOrderbyFavorites  AnimeSearchQueryOrderby = "favorites"
	AnimeSearchQueryOrderbyMalId      AnimeSearchQueryOrderby = "mal_id"
	AnimeSearchQueryOrderbyMembers    AnimeSearchQueryOrderby = "members"
	AnimeSearchQueryOrderbyPopularity AnimeSearchQueryOrderby = "popularity"
	AnimeSearchQueryOrderbyRank       AnimeSearchQueryOrderby = "rank"
	AnimeSearchQueryOrderbyScore      AnimeSearchQueryOrderby = "score"
	AnimeSearchQueryOrderbyScoredBy   AnimeSearchQueryOrderby = "scored_by"
	AnimeSearchQueryOrderbyStartDate  AnimeSearchQueryOrderby = "start_date"
	AnimeSearchQueryOrderbyTitle      AnimeSearchQueryOrderby = "title"
)
const (
	AnimeSearchQueryRatingG    AnimeSearchQueryRating = "g"
	AnimeSearchQueryRatingPg   AnimeSearchQueryRating = "pg"
	AnimeSearchQueryRatingPg13 AnimeSearchQueryRating = "pg13"
	AnimeSearchQueryRatingR    AnimeSearchQueryRating = "r"
	AnimeSearchQueryRatingR17  AnimeSearchQueryRating = "r17"
	AnimeSearchQueryRatingRx   AnimeSearchQueryRating = "rx"
)
const (
	AnimeSearchQueryStatusAiring   AnimeSearchQueryStatus = "airing"
	AnimeSearchQueryStatusComplete AnimeSearchQueryStatus = "complete"
	AnimeSearchQueryStatusUpcoming AnimeSearchQueryStatus = "upcoming"
)
const (
	AnimeSearchQueryTypeCm        AnimeSearchQueryType = "cm"
	AnimeSearchQueryTypeMovie     AnimeSearchQueryType = "movie"
	AnimeSearchQueryTypeMusic     AnimeSearchQueryType = "music"
	AnimeSearchQueryTypeOna       AnimeSearchQueryType = "ona"
	AnimeSearchQueryTypeOva       AnimeSearchQueryType = "ova"
	AnimeSearchQueryTypePv        AnimeSearchQueryType = "pv"
	AnimeSearchQueryTypeSpecial   AnimeSearchQueryType = "special"
	AnimeSearchQueryTypeTv        AnimeSearchQueryType = "tv"
	AnimeSearchQueryTypeTvSpecial AnimeSearchQueryType = "tv_special"
)
const (
	CharactersSearchQueryOrderbyFavorites CharactersSearchQueryOrderby = "favorites"
	CharactersSearchQueryOrderbyMalId     CharactersSearchQueryOrderby = "mal_id"
	CharactersSearchQueryOrderbyName      CharactersSearchQueryOrderby = "name"
)
const (
	ClubAccessPrivate ClubAccess = "private"
	ClubAccessPublic  ClubAccess = "public"
	ClubAccessSecret  ClubAccess = "secret"
)
const (
	ClubCategoryActorsArtists       ClubCategory = "actors & artists"
	ClubCategoryAnime               ClubCategory = "anime"
	ClubCategoryCharacters          ClubCategory = "characters"
	ClubCategoryCitiesNeighborhoods ClubCategory = "cities & neighborhoods"
	ClubCategoryCompanies           ClubCategory = "companies"
	ClubCategoryConventions         ClubCategory = "conventions"
	ClubCategoryGames               ClubCategory = "games"
	ClubCategoryJapan               ClubCategory = "japan"
	ClubCategoryManga               ClubCategory = "manga"
	ClubCategoryMusic               ClubCategory = "music"
	ClubCategoryOthers              ClubCategory = "others"
	ClubCategorySchools             ClubCategory = "schools"
)
const (
	ClubSearchQueryCategoryActorsAndArtists       ClubSearchQueryCategory = "actors_and_artists"
	ClubSearchQueryCategoryAnime                  ClubSearchQueryCategory = "anime"
	ClubSearchQueryCategoryCharacters             ClubSearchQueryCategory = "characters"
	ClubSearchQueryCategoryCitiesAndNeighborhoods ClubSearchQueryCategory = "cities_and_neighborhoods"
	ClubSearchQueryCategoryCompanies              ClubSearchQueryCategory = "companies"
	ClubSearchQueryCategoryConventions            ClubSearchQueryCategory = "conventions"
	ClubSearchQueryCategoryGames                  ClubSearchQueryCategory = "games"
	ClubSearchQueryCategoryJapan                  ClubSearchQueryCategory = "japan"
	ClubSearchQueryCategoryManga                  ClubSearchQueryCategory = "manga"
	ClubSearchQueryCategoryMusic                  ClubSearchQueryCategory = "music"
	ClubSearchQueryCategoryOther                  ClubSearchQueryCategory = "other"
	ClubSearchQueryCategorySchools                ClubSearchQueryCategory = "schools"
)
const (
	ClubSearchQueryOrderbyCreated      ClubSearchQueryOrderby = "created"
	ClubSearchQueryOrderbyMalId        ClubSearchQueryOrderby = "mal_id"
	ClubSearchQueryOrderbyMembersCount ClubSearchQueryOrderby = "members_count"
	ClubSearchQueryOrderbyName         ClubSearchQueryOrderby = "name"
)
const (
	ClubSearchQueryTypePrivate ClubSearchQueryType = "private"
	ClubSearchQueryTypePublic  ClubSearchQueryType = "public"
	ClubSearchQueryTypeSecret  ClubSearchQueryType = "secret"
)
const (
	GenreQueryFilterDemographics   GenreQueryFilter = "demographics"
	GenreQueryFilterExplicitGenres GenreQueryFilter = "explicit_genres"
	GenreQueryFilterGenres         GenreQueryFilter = "genres"
	GenreQueryFilterThemes         GenreQueryFilter = "themes"
)
const (
	MagazinesQueryOrderbyCount MagazinesQueryOrderby = "count"
	MagazinesQueryOrderbyMalId MagazinesQueryOrderby = "mal_id"
	MagazinesQueryOrderbyName  MagazinesQueryOrderby = "name"
)
const (
	MangaStatusDiscontinued    MangaStatus = "Discontinued"
	MangaStatusFinished        MangaStatus = "Finished"
	MangaStatusNotYetPublished MangaStatus = "Not yet published"
	MangaStatusOnHiatus        MangaStatus = "On Hiatus"
	MangaStatusPublishing      MangaStatus = "Publishing"
)
const (
	MangaTypeDoujinshi  MangaType = "Doujinshi"
	MangaTypeLightNovel MangaType = "Light Novel"
	MangaTypeManga      MangaType = "Manga"
	MangaTypeManhua     MangaType = "Manhua"
	MangaTypeManhwa     MangaType = "Manhwa"
	MangaTypeNovel      MangaType = "Novel"
	MangaTypeOEL        MangaType = "OEL"
	MangaTypeOneShot    MangaType = "One-shot"
)
const (
	MangaFullStatusDiscontinued    MangaFullStatus = "Discontinued"
	MangaFullStatusFinished        MangaFullStatus = "Finished"
	MangaFullStatusNotYetPublished MangaFullStatus = "Not yet published"
	MangaFullStatusOnHiatus        MangaFullStatus = "On Hiatus"
	MangaFullStatusPublishing      MangaFullStatus = "Publishing"
)
const (
	MangaFullTypeDoujinshi  MangaFullType = "Doujinshi"
	MangaFullTypeLightNovel MangaFullType = "Light Novel"
	MangaFullTypeManga      MangaFullType = "Manga"
	MangaFullTypeManhua     MangaFullType = "Manhua"
	MangaFullTypeManhwa     MangaFullType = "Manhwa"
	MangaFullTypeNovel      MangaFullType = "Novel"
	MangaFullTypeOEL        MangaFullType = "OEL"
	MangaFullTypeOneShot    MangaFullType = "One-shot"
)
const (
	MangaSearchQueryOrderbyChapters   MangaSearchQueryOrderby = "chapters"
	MangaSearchQueryOrderbyEndDate    MangaSearchQueryOrderby = "end_date"
	MangaSearchQueryOrderbyFavorites  MangaSearchQueryOrderby = "favorites"
	MangaSearchQueryOrderbyMalId      MangaSearchQueryOrderby = "mal_id"
	MangaSearchQueryOrderbyMembers    MangaSearchQueryOrderby = "members"
	MangaSearchQueryOrderbyPopularity MangaSearchQueryOrderby = "popularity"
	MangaSearchQueryOrderbyRank       MangaSearchQueryOrderby = "rank"
	MangaSearchQueryOrderbyScore      MangaSearchQueryOrderby = "score"
	MangaSearchQueryOrderbyScoredBy   MangaSearchQueryOrderby = "scored_by"
	MangaSearchQueryOrderbyStartDate  MangaSearchQueryOrderby = "start_date"
	MangaSearchQueryOrderbyTitle      MangaSearchQueryOrderby = "title"
	MangaSearchQueryOrderbyVolumes    MangaSearchQueryOrderby = "volumes"
)
const (
	MangaSearchQueryStatusComplete     MangaSearchQueryStatus = "complete"
	MangaSearchQueryStatusDiscontinued MangaSearchQueryStatus = "discontinued"
	MangaSearchQueryStatusHiatus       MangaSearchQueryStatus = "hiatus"
	MangaSearchQueryStatusPublishing   MangaSearchQueryStatus = "publishing"
	MangaSearchQueryStatusUpcoming     MangaSearchQueryStatus = "upcoming"
)
const (
	MangaSearchQueryTypeDoujin     MangaSearchQueryType = "doujin"
	MangaSearchQueryTypeLightnovel MangaSearchQueryType = "lightnovel"
	MangaSearchQueryTypeManga      MangaSearchQueryType = "manga"
	MangaSearchQueryTypeManhua     MangaSearchQueryType = "manhua"
	MangaSearchQueryTypeManhwa     MangaSearchQueryType = "manhwa"
	MangaSearchQueryTypeNovel      MangaSearchQueryType = "novel"
	MangaSearchQueryTypeOneshot    MangaSearchQueryType = "oneshot"
)
const (
	PeopleSearchQueryOrderbyBirthday  PeopleSearchQueryOrderby = "birthday"
	PeopleSearchQueryOrderbyFavorites PeopleSearchQueryOrderby = "favorites"
	PeopleSearchQueryOrderbyMalId     PeopleSearchQueryOrderby = "mal_id"
	PeopleSearchQueryOrderbyName      PeopleSearchQueryOrderby = "name"
)
const (
	ProducersQueryOrderbyCount       ProducersQueryOrderby = "count"
	ProducersQueryOrderbyEstablished ProducersQueryOrderby = "established"
	ProducersQueryOrderbyFavorites   ProducersQueryOrderby = "favorites"
	ProducersQueryOrderbyMalId       ProducersQueryOrderby = "mal_id"
)
const (
	SearchQuerySortAsc  SearchQuerySort = "asc"
	SearchQuerySortDesc SearchQuerySort = "desc"
)
const (
	TopAnimeFilterAiring       TopAnimeFilter = "airing"
	TopAnimeFilterBypopularity TopAnimeFilter = "bypopularity"
	TopAnimeFilterFavorite     TopAnimeFilter = "favorite"
	TopAnimeFilterUpcoming     TopAnimeFilter = "upcoming"
)
const (
	TopMangaFilterBypopularity TopMangaFilter = "bypopularity"
	TopMangaFilterFavorite     TopMangaFilter = "favorite"
	TopMangaFilterPublishing   TopMangaFilter = "publishing"
	TopMangaFilterUpcoming     TopMangaFilter = "upcoming"
)
const (
	TopReviewsTypeEnumAnime TopReviewsTypeEnum = "anime"
	TopReviewsTypeEnumManga TopReviewsTypeEnum = "manga"
)
const (
	UserAnimeListStatusFilterAll         UserAnimeListStatusFilter = "all"
	UserAnimeListStatusFilterCompleted   UserAnimeListStatusFilter = "completed"
	UserAnimeListStatusFilterDropped     UserAnimeListStatusFilter = "dropped"
	UserAnimeListStatusFilterOnhold      UserAnimeListStatusFilter = "onhold"
	UserAnimeListStatusFilterPlantowatch UserAnimeListStatusFilter = "plantowatch"
	UserAnimeListStatusFilterWatching    UserAnimeListStatusFilter = "watching"
)
const (
	UserMangaListStatusFilterAll        UserMangaListStatusFilter = "all"
	UserMangaListStatusFilterCompleted  UserMangaListStatusFilter = "completed"
	UserMangaListStatusFilterDropped    UserMangaListStatusFilter = "dropped"
	UserMangaListStatusFilterOnhold     UserMangaListStatusFilter = "onhold"
	UserMangaListStatusFilterPlantoread UserMangaListStatusFilter = "plantoread"
	UserMangaListStatusFilterReading    UserMangaListStatusFilter = "reading"
)
const (
	UsersSearchQueryGenderAny       UsersSearchQueryGender = "any"
	UsersSearchQueryGenderFemale    UsersSearchQueryGender = "female"
	UsersSearchQueryGenderMale      UsersSearchQueryGender = "male"
	UsersSearchQueryGenderNonbinary UsersSearchQueryGender = "nonbinary"
)
const (
	GetAnimeForumParamsFilterAll     GetAnimeForumParamsFilter = "all"
	GetAnimeForumParamsFilterEpisode GetAnimeForumParamsFilter = "episode"
	GetAnimeForumParamsFilterOther   GetAnimeForumParamsFilter = "other"
)
const (
	GetMangaTopicsParamsFilterAll     GetMangaTopicsParamsFilter = "all"
	GetMangaTopicsParamsFilterEpisode GetMangaTopicsParamsFilter = "episode"
	GetMangaTopicsParamsFilterOther   GetMangaTopicsParamsFilter = "other"
)
const (
	GetSchedulesParamsFilterFriday    GetSchedulesParamsFilter = "friday"
	GetSchedulesParamsFilterMonday    GetSchedulesParamsFilter = "monday"
	GetSchedulesParamsFilterOther     GetSchedulesParamsFilter = "other"
	GetSchedulesParamsFilterSaturday  GetSchedulesParamsFilter = "saturday"
	GetSchedulesParamsFilterSunday    GetSchedulesParamsFilter = "sunday"
	GetSchedulesParamsFilterThursday  GetSchedulesParamsFilter = "thursday"
	GetSchedulesParamsFilterTuesday   GetSchedulesParamsFilter = "tuesday"
	GetSchedulesParamsFilterUnknown   GetSchedulesParamsFilter = "unknown"
	GetSchedulesParamsFilterWednesday GetSchedulesParamsFilter = "wednesday"
)
const (
	GetSchedulesParamsKidsFalse GetSchedulesParamsKids = "false"
	GetSchedulesParamsKidsTrue  GetSchedulesParamsKids = "true"
)
const (
	GetSchedulesParamsSfwFalse GetSchedulesParamsSfw = "false"
	GetSchedulesParamsSfwTrue  GetSchedulesParamsSfw = "true"
)
const (
	GetSeasonNowParamsFilterMovie   GetSeasonNowParamsFilter = "movie"
	GetSeasonNowParamsFilterMusic   GetSeasonNowParamsFilter = "music"
	GetSeasonNowParamsFilterOna     GetSeasonNowParamsFilter = "ona"
	GetSeasonNowParamsFilterOva     GetSeasonNowParamsFilter = "ova"
	GetSeasonNowParamsFilterSpecial GetSeasonNowParamsFilter = "special"
	GetSeasonNowParamsFilterTv      GetSeasonNowParamsFilter = "tv"
)
const (
	GetSeasonUpcomingParamsFilterMovie   GetSeasonUpcomingParamsFilter = "movie"
	GetSeasonUpcomingParamsFilterMusic   GetSeasonUpcomingParamsFilter = "music"
	GetSeasonUpcomingParamsFilterOna     GetSeasonUpcomingParamsFilter = "ona"
	GetSeasonUpcomingParamsFilterOva     GetSeasonUpcomingParamsFilter = "ova"
	GetSeasonUpcomingParamsFilterSpecial GetSeasonUpcomingParamsFilter = "special"
	GetSeasonUpcomingParamsFilterTv      GetSeasonUpcomingParamsFilter = "tv"
)
const (
	GetSeasonParamsFilterMovie   GetSeasonParamsFilter = "movie"
	GetSeasonParamsFilterMusic   GetSeasonParamsFilter = "music"
	GetSeasonParamsFilterOna     GetSeasonParamsFilter = "ona"
	GetSeasonParamsFilterOva     GetSeasonParamsFilter = "ova"
	GetSeasonParamsFilterSpecial GetSeasonParamsFilter = "special"
	GetSeasonParamsFilterTv      GetSeasonParamsFilter = "tv"
)
const (
	GetUserHistoryParamsTypeAnime GetUserHistoryParamsType = "anime"
	GetUserHistoryParamsTypeManga GetUserHistoryParamsType = "manga"
)

type AnimeVideosEpisodesPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type CharacterPicturesDataElement struct {
	ImageUrl      *string `json:"image_url"`
	LargeImageUrl *string `json:"large_image_url"`
}
type PaginationPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type PeopleImagesJpg struct {
	ImageUrl *string `json:"image_url"`
}
type UserImagesWebp struct {
	ImageUrl *string `json:"image_url"`
}
type AnimeFullStreamingElement struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}
type PersonFullVoicesElement struct {
	Anime     *AnimeMeta     `json:"anime,omitempty"`
	Character *CharacterMeta `json:"character,omitempty"`
	Role      *string        `json:"role,omitempty"`
}
type UserFavoritesMangaElement struct {
	Images    *MangaImages `json:"images,omitempty"`
	MalId     *int         `json:"mal_id,omitempty"`
	StartYear *int         `json:"start_year,omitempty"`
	Title     *string      `json:"title,omitempty"`
	Type      *string      `json:"type,omitempty"`
	Url       *string      `json:"url,omitempty"`
}
type CharactersSearchPaginationItems struct {
	Count   *int `json:"count,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	Total   *int `json:"total,omitempty"`
}
type MangaReviewReactions struct {
	Confusing   *int `json:"confusing,omitempty"`
	Creative    *int `json:"creative,omitempty"`
	Funny       *int `json:"funny,omitempty"`
	Informative *int `json:"informative,omitempty"`
	LoveIt      *int `json:"love_it,omitempty"`
	Nice        *int `json:"nice,omitempty"`
	Overall     *int `json:"overall,omitempty"`
	WellWritten *int `json:"well_written,omitempty"`
}
type TrailerImagesImages struct {
	ImageUrl        *string `json:"image_url"`
	LargeImageUrl   *string `json:"large_image_url"`
	MaximumImageUrl *string `json:"maximum_image_url"`
	MediumImageUrl  *string `json:"medium_image_url"`
	SmallImageUrl   *string `json:"small_image_url"`
}
type GetTopReviewsResponseJSON200DataPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type AnimeRelationsDataElement struct {
	Entry    *[]MalUrl `json:"entry,omitempty"`
	Relation *string   `json:"relation,omitempty"`
}
type AnimeVideosDataMusicVideosElement struct {
	Meta  *AnimeVideosDataMusicVideosElementMeta `json:"meta,omitempty"`
	Title *string                                `json:"title,omitempty"`
	Video *Trailer                               `json:"video,omitempty"`
}
type AnimeVideosDataPromoElement struct {
	Title   *string  `json:"title,omitempty"`
	Trailer *Trailer `json:"trailer,omitempty"`
}
type CharacterAnimeDataElement struct {
	Anime *AnimeMeta `json:"anime,omitempty"`
	Role  *string    `json:"role,omitempty"`
}
type ProducerFullExternalElement struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}
type UsersTempDataElementImagesWebp struct {
	ImageUrl *string `json:"image_url,omitempty"`
}
type GetCharacterFullByIdResponseJSON200 struct {
	Data *CharacterFull `json:"data,omitempty"`
}
type GetClubsByIdResponseJSON200 struct {
	Data *Club `json:"data,omitempty"`
}
type AnimeEpisodesPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type AnimeFullTheme struct {
	Endings  *[]string `json:"endings,omitempty"`
	Openings *[]string `json:"openings,omitempty"`
}
type AnimeStatisticsData struct {
	Completed   *int                                `json:"completed,omitempty"`
	Dropped     *int                                `json:"dropped,omitempty"`
	OnHold      *int                                `json:"on_hold,omitempty"`
	PlanToWatch *int                                `json:"plan_to_watch,omitempty"`
	Scores      *[]AnimeStatisticsDataScoresElement `json:"scores,omitempty"`
	Total       *int                                `json:"total,omitempty"`
	Watching    *int                                `json:"watching,omitempty"`
}
type DaterangeProp struct {
	From   *DaterangePropFrom `json:"from,omitempty"`
	String *string            `json:"string"`
	To     *DaterangePropTo   `json:"to,omitempty"`
}
type GetClubMembersResponseJSON200DataElement struct {
	Images   *UserImages `json:"images,omitempty"`
	Url      *string     `json:"url,omitempty"`
	Username *string     `json:"username,omitempty"`
}
type PaginationPlusPaginationItems struct {
	Count   *int `json:"count,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	Total   *int `json:"total,omitempty"`
}
type UsersSearchDataElement struct {
	Images     *UserImages `json:"images,omitempty"`
	LastOnline *string     `json:"last_online,omitempty"`
	Url        *string     `json:"url,omitempty"`
	Username   *string     `json:"username,omitempty"`
}
type ForumDataElement struct {
	AuthorUrl      *string                      `json:"author_url,omitempty"`
	AuthorUsername *string                      `json:"author_username,omitempty"`
	Comments       *int                         `json:"comments,omitempty"`
	Date           *string                      `json:"date,omitempty"`
	LastComment    *ForumDataElementLastComment `json:"last_comment,omitempty"`
	MalId          *int                         `json:"mal_id,omitempty"`
	Title          *string                      `json:"title,omitempty"`
	Url            *string                      `json:"url,omitempty"`
}
type SchedulesPaginationItems struct {
	Count   *int `json:"count,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	Total   *int `json:"total,omitempty"`
}
type UserAboutDataElement struct {
	About *string `json:"about"`
}
type UserUpdatesDataAnimeElement struct {
	Date          *string    `json:"date,omitempty"`
	Entry         *AnimeMeta `json:"entry,omitempty"`
	EpisodesSeen  *int       `json:"episodes_seen"`
	EpisodesTotal *int       `json:"episodes_total"`
	Score         *int       `json:"score"`
	Status        *string    `json:"status,omitempty"`
}
type GetRandomCharactersResponseJSON200 struct {
	Data *Character `json:"data,omitempty"`
}
type GetTopReviewsResponseJSON200 struct {
	Data *GetTopReviewsResponseJSON200Data `json:"data,omitempty"`
}
type AnimeImagesJpg struct {
	ImageUrl      *string `json:"image_url"`
	LargeImageUrl *string `json:"large_image_url"`
	SmallImageUrl *string `json:"small_image_url"`
}
type ClubRelationsData struct {
	Anime      *[]MalUrl `json:"anime,omitempty"`
	Characters *[]MalUrl `json:"characters,omitempty"`
	Manga      *[]MalUrl `json:"manga,omitempty"`
}
type MangaImagesWebp struct {
	ImageUrl      *string `json:"image_url"`
	LargeImageUrl *string `json:"large_image_url"`
	SmallImageUrl *string `json:"small_image_url"`
}
type MangaReviewsPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type PeopleSearchPaginationItems struct {
	Count   *int `json:"count,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	Total   *int `json:"total,omitempty"`
}
type WatchEpisodesPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type WatchPromosPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type AnimeCharactersDataElementVoiceActorsElementPerson struct {
	Images *PeopleImages `json:"images,omitempty"`
	MalId  *int          `json:"mal_id,omitempty"`
	Name   *string       `json:"name,omitempty"`
	Url    *string       `json:"url,omitempty"`
}
type AnimeEpisodesDataElement struct {
	Aired         *string  `json:"aired"`
	Filler        *bool    `json:"filler,omitempty"`
	ForumUrl      *string  `json:"forum_url"`
	MalId         *int     `json:"mal_id,omitempty"`
	Recap         *bool    `json:"recap,omitempty"`
	Score         *float32 `json:"score"`
	Title         *string  `json:"title,omitempty"`
	TitleJapanese *string  `json:"title_japanese"`
	TitleRomanji  *string  `json:"title_romanji"`
	Url           *string  `json:"url"`
}
type DaterangePropTo struct {
	Day   *int `json:"day"`
	Month *int `json:"month"`
	Year  *int `json:"year"`
}
type MangaSearchPaginationItems struct {
	Count   *int `json:"count,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	Total   *int `json:"total,omitempty"`
}
type UserProfileFullStatisticsManga struct {
	ChaptersRead *int     `json:"chapters_read,omitempty"`
	Completed    *int     `json:"completed,omitempty"`
	DaysRead     *float32 `json:"days_read,omitempty"`
	Dropped      *int     `json:"dropped,omitempty"`
	MeanScore    *float32 `json:"mean_score,omitempty"`
	OnHold       *int     `json:"on_hold,omitempty"`
	PlanToRead   *int     `json:"plan_to_read,omitempty"`
	Reading      *int     `json:"reading,omitempty"`
	Reread       *int     `json:"reread,omitempty"`
	TotalEntries *int     `json:"total_entries,omitempty"`
	VolumesRead  *int     `json:"volumes_read,omitempty"`
}
type WatchEpisodesDataElement struct {
	Entry        *AnimeMeta                                 `json:"entry,omitempty"`
	Episodes     *[]WatchEpisodesDataElementEpisodesElement `json:"episodes,omitempty"`
	RegionLocked *bool                                      `json:"region_locked,omitempty"`
}
type GetTopReviews200DataData0Reactions struct {
	Confusing   *int `json:"confusing,omitempty"`
	Creative    *int `json:"creative,omitempty"`
	Funny       *int `json:"funny,omitempty"`
	Informative *int `json:"informative,omitempty"`
	LoveIt      *int `json:"love_it,omitempty"`
	Nice        *int `json:"nice,omitempty"`
	Overall     *int `json:"overall,omitempty"`
	WellWritten *int `json:"well_written,omitempty"`
}
type AnimeSearchPaginationItems struct {
	Count   *int `json:"count,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	Total   *int `json:"total,omitempty"`
}
type AnimeVideosData struct {
	Episodes    *[]AnimeVideosDataEpisodesElement    `json:"episodes,omitempty"`
	MusicVideos *[]AnimeVideosDataMusicVideosElement `json:"music_videos,omitempty"`
	Promo       *[]AnimeVideosDataPromoElement       `json:"promo,omitempty"`
}
type CharacterFullMangaElement struct {
	Manga *MangaMeta `json:"manga,omitempty"`
	Role  *string    `json:"role,omitempty"`
}
type MangaFullExternalElement struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}
type MangaImagesJpg struct {
	ImageUrl      *string `json:"image_url"`
	LargeImageUrl *string `json:"large_image_url"`
	SmallImageUrl *string `json:"small_image_url"`
}
type UserFavoritesAnimeElement struct {
	Images    *AnimeImages `json:"images,omitempty"`
	MalId     *int         `json:"mal_id,omitempty"`
	StartYear *int         `json:"start_year,omitempty"`
	Title     *string      `json:"title,omitempty"`
	Type      *string      `json:"type,omitempty"`
	Url       *string      `json:"url,omitempty"`
}
type UserUpdatesData struct {
	Anime *[]UserUpdatesDataAnimeElement `json:"anime,omitempty"`
	Manga *[]UserUpdatesDataMangaElement `json:"manga,omitempty"`
}
type GetMangaRelationsResponseJSON200 struct {
	Data *[]Relation `json:"data,omitempty"`
}
type AnimeVideosDataMusicVideosElementMeta struct {
	Author *string `json:"author"`
	Title  *string `json:"title"`
}
type MagazinesPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type PersonFullMangaElement struct {
	Manga    *MangaMeta `json:"manga,omitempty"`
	Position *string    `json:"position,omitempty"`
}
type GetAnimeEpisodeByIdResponseJSON200 struct {
	Data *AnimeEpisode `json:"data,omitempty"`
}
type GetRandomAnimeResponseJSON200 struct {
	Data *Anime `json:"data,omitempty"`
}
type AnimeSearchPagination struct {
	HasNextPage     *bool                       `json:"has_next_page,omitempty"`
	Items           *AnimeSearchPaginationItems `json:"items,omitempty"`
	LastVisiblePage *int                        `json:"last_visible_page,omitempty"`
}
type AnimeStaffDataElement struct {
	Person    *AnimeStaffDataElementPerson `json:"person,omitempty"`
	Positions *[]string                    `json:"positions,omitempty"`
}
type ClubStaffDataElement struct {
	Url      *string `json:"url,omitempty"`
	Username *string `json:"username,omitempty"`
}
type MangaCharactersDataElement struct {
	Character *CharacterMeta `json:"character,omitempty"`
	Role      *string        `json:"role,omitempty"`
}
type MangaNewsDataElement struct {
	AuthorUrl      *string       `json:"author_url,omitempty"`
	AuthorUsername *string       `json:"author_username,omitempty"`
	Comments       *int          `json:"comments,omitempty"`
	Date           *string       `json:"date,omitempty"`
	Excerpt        *string       `json:"excerpt,omitempty"`
	ForumUrl       *string       `json:"forum_url,omitempty"`
	Images         *CommonImages `json:"images,omitempty"`
	MalId          *int          `json:"mal_id,omitempty"`
	Title          *string       `json:"title,omitempty"`
	Url            *string       `json:"url,omitempty"`
}
type MangaStatisticsDataScoresElement struct {
	Percentage *float32 `json:"percentage,omitempty"`
	Score      *int     `json:"score,omitempty"`
	Votes      *int     `json:"votes,omitempty"`
}
type AnimeCharactersDataElement struct {
	Character   *AnimeCharactersDataElementCharacter            `json:"character,omitempty"`
	Role        *string                                         `json:"role,omitempty"`
	VoiceActors *[]AnimeCharactersDataElementVoiceActorsElement `json:"voice_actors,omitempty"`
}
type CharacterImagesJpg struct {
	ImageUrl      *string `json:"image_url"`
	SmallImageUrl *string `json:"small_image_url"`
}
type ClubMemberDataElement struct {
	Images   *UserImages `json:"images,omitempty"`
	Url      *string     `json:"url,omitempty"`
	Username *string     `json:"username,omitempty"`
}
type ExternalLinksDataElement struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}
type PicturesVariantsDataElement struct {
	Images *CommonImages `json:"images,omitempty"`
}
type AnimeUserupdatesPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type MangaUserupdatesPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type GetMangaFullByIdResponseJSON200 struct {
	Data *MangaFull `json:"data,omitempty"`
}
type GetRandomUsersResponseJSON200 struct {
	Data *UserProfile `json:"data,omitempty"`
}
type AnimeReviewsPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type CharactersSearchPagination struct {
	HasNextPage     *bool                            `json:"has_next_page,omitempty"`
	Items           *CharactersSearchPaginationItems `json:"items,omitempty"`
	LastVisiblePage *int                             `json:"last_visible_page,omitempty"`
}
type MangaStatisticsData struct {
	Completed  *int                                `json:"completed,omitempty"`
	Dropped    *int                                `json:"dropped,omitempty"`
	OnHold     *int                                `json:"on_hold,omitempty"`
	PlanToRead *int                                `json:"plan_to_read,omitempty"`
	Reading    *int                                `json:"reading,omitempty"`
	Scores     *[]MangaStatisticsDataScoresElement `json:"scores,omitempty"`
	Total      *int                                `json:"total,omitempty"`
}
type ProducersPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type UsersTempDataElement struct {
	About      *string                         `json:"about,omitempty"`
	AnimeStats *UsersTempDataElementAnimeStats `json:"anime_stats,omitempty"`
	Birthday   *string                         `json:"birthday,omitempty"`
	Favorites  *UsersTempDataElementFavorites  `json:"favorites,omitempty"`
	Gender     *string                         `json:"gender,omitempty"`
	Images     *UsersTempDataElementImages     `json:"images,omitempty"`
	Joined     *string                         `json:"joined,omitempty"`
	LastOnline *string                         `json:"last_online,omitempty"`
	Location   *string                         `json:"location,omitempty"`
	MalId      *int                            `json:"mal_id,omitempty"`
	MangaStats *UsersTempDataElementMangaStats `json:"manga_stats,omitempty"`
	Url        *string                         `json:"url,omitempty"`
	Username   *string                         `json:"username,omitempty"`
}
type UsersTempDataElementFavorites struct {
	Anime      *[]EntryMeta `json:"anime,omitempty"`
	Characters *[]EntryMeta `json:"characters,omitempty"`
	Manga      *[]EntryMeta `json:"manga,omitempty"`
	People     *[]EntryMeta `json:"people,omitempty"`
}
type AnimeCharactersDataElementVoiceActorsElement struct {
	Language *string                                             `json:"language,omitempty"`
	Person   *AnimeCharactersDataElementVoiceActorsElementPerson `json:"person,omitempty"`
}
type UserProfileFullExternalElement struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}
type UsersSearchPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type WatchEpisodesDataElementEpisodesElement struct {
	MalId   *string `json:"mal_id,omitempty"`
	Premium *bool   `json:"premium,omitempty"`
	Title   *string `json:"title,omitempty"`
	Url     *string `json:"url,omitempty"`
}
type GetUserProfileResponseJSON200 struct {
	Data *UserProfile `json:"data,omitempty"`
}
type CharacterFullAnimeElement struct {
	Anime *AnimeMeta `json:"anime,omitempty"`
	Role  *string    `json:"role,omitempty"`
}
type MangaReviewsDataElement struct {
	Date          *string                           `json:"date,omitempty"`
	IsPreliminary *bool                             `json:"is_preliminary,omitempty"`
	IsSpoiler     *bool                             `json:"is_spoiler,omitempty"`
	MalId         *int                              `json:"mal_id,omitempty"`
	Reactions     *MangaReviewsDataElementReactions `json:"reactions,omitempty"`
	Review        *string                           `json:"review,omitempty"`
	Score         *int                              `json:"score,omitempty"`
	Tags          *[]string                         `json:"tags,omitempty"`
	Type          *string                           `json:"type,omitempty"`
	Url           *string                           `json:"url,omitempty"`
	User          *UserMeta                         `json:"user,omitempty"`
}
type PersonAnimeDataElement struct {
	Anime    *AnimeMeta `json:"anime,omitempty"`
	Position *string    `json:"position,omitempty"`
}
type GetTopReviews200DataData1Reactions struct {
	Confusing   *int `json:"confusing,omitempty"`
	Creative    *int `json:"creative,omitempty"`
	Funny       *int `json:"funny,omitempty"`
	Informative *int `json:"informative,omitempty"`
	LoveIt      *int `json:"love_it,omitempty"`
	Nice        *int `json:"nice,omitempty"`
	Overall     *int `json:"overall,omitempty"`
	WellWritten *int `json:"well_written,omitempty"`
}
type AnimeFullExternalElement struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}
type TrailerImages1 struct {
	ImageUrl        *string `json:"image_url"`
	LargeImageUrl   *string `json:"large_image_url"`
	MaximumImageUrl *string `json:"maximum_image_url"`
	MediumImageUrl  *string `json:"medium_image_url"`
	SmallImageUrl   *string `json:"small_image_url"`
}
type GetAnimeByIdResponseJSON200 struct {
	Data *Anime `json:"data,omitempty"`
}
type AnimeFullRelationsElement struct {
	Entry    *[]MalUrl `json:"entry,omitempty"`
	Relation *string   `json:"relation,omitempty"`
}
type AnimeStaffDataElementPerson struct {
	Images *PeopleImages `json:"images,omitempty"`
	MalId  *int          `json:"mal_id,omitempty"`
	Name   *string       `json:"name,omitempty"`
	Url    *string       `json:"url,omitempty"`
}
type AnimeVideosDataEpisodesElement struct {
	Episode *string       `json:"episode,omitempty"`
	Images  *CommonImages `json:"images,omitempty"`
	MalId   *int          `json:"mal_id,omitempty"`
	Title   *string       `json:"title,omitempty"`
	Url     *string       `json:"url,omitempty"`
}
type MangaSearchPagination struct {
	HasNextPage     *bool                       `json:"has_next_page,omitempty"`
	Items           *MangaSearchPaginationItems `json:"items,omitempty"`
	LastVisiblePage *int                        `json:"last_visible_page,omitempty"`
}
type PersonVoiceActingRolesDataElement struct {
	Anime     *AnimeMeta     `json:"anime,omitempty"`
	Character *CharacterMeta `json:"character,omitempty"`
	Role      *string        `json:"role,omitempty"`
}
type RecommendationsPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type SchedulesPagination struct {
	HasNextPage     *bool                     `json:"has_next_page,omitempty"`
	Items           *SchedulesPaginationItems `json:"items,omitempty"`
	LastVisiblePage *int                      `json:"last_visible_page,omitempty"`
}
type GetUserReviews200DataData0Reactions struct {
	Confusing   *int `json:"confusing,omitempty"`
	Creative    *int `json:"creative,omitempty"`
	Funny       *int `json:"funny,omitempty"`
	Informative *int `json:"informative,omitempty"`
	LoveIt      *int `json:"love_it,omitempty"`
	Nice        *int `json:"nice,omitempty"`
	Overall     *int `json:"overall,omitempty"`
	WellWritten *int `json:"well_written,omitempty"`
}
type AnimeStatisticsDataScoresElement struct {
	Percentage *float32 `json:"percentage,omitempty"`
	Score      *int     `json:"score,omitempty"`
	Votes      *int     `json:"votes,omitempty"`
}
type AnimeImagesWebp struct {
	ImageUrl      *string `json:"image_url"`
	LargeImageUrl *string `json:"large_image_url"`
	SmallImageUrl *string `json:"small_image_url"`
}
type UserImagesJpg struct {
	ImageUrl *string `json:"image_url"`
}
type UserProfileFullStatisticsAnime struct {
	Completed       *int     `json:"completed,omitempty"`
	DaysWatched     *float32 `json:"days_watched,omitempty"`
	Dropped         *int     `json:"dropped,omitempty"`
	EpisodesWatched *int     `json:"episodes_watched,omitempty"`
	MeanScore       *float32 `json:"mean_score,omitempty"`
	OnHold          *int     `json:"on_hold,omitempty"`
	PlanToWatch     *int     `json:"plan_to_watch,omitempty"`
	Rewatched       *int     `json:"rewatched,omitempty"`
	TotalEntries    *int     `json:"total_entries,omitempty"`
	Watching        *int     `json:"watching,omitempty"`
}
type UserStatisticsData struct {
	Anime *UserStatisticsDataAnime `json:"anime,omitempty"`
	Manga *UserStatisticsDataManga `json:"manga,omitempty"`
}
type GetAnimeFullByIdResponseJSON200 struct {
	Data *AnimeFull `json:"data,omitempty"`
}
type GetPersonFullByIdResponseJSON200 struct {
	Data *PersonFull `json:"data,omitempty"`
}
type GetProducerByIdResponseJSON200 struct {
	Data *Producer `json:"data,omitempty"`
}
type AnimeReviewsDataElementReactions struct {
	Confusing   *int `json:"confusing,omitempty"`
	Creative    *int `json:"creative,omitempty"`
	Funny       *int `json:"funny,omitempty"`
	Informative *int `json:"informative,omitempty"`
	LoveIt      *int `json:"love_it,omitempty"`
	Nice        *int `json:"nice,omitempty"`
	Overall     *int `json:"overall,omitempty"`
	WellWritten *int `json:"well_written,omitempty"`
}
type ForumDataElementLastComment struct {
	AuthorUrl      *string `json:"author_url,omitempty"`
	AuthorUsername *string `json:"author_username,omitempty"`
	Date           *string `json:"date"`
	Url            *string `json:"url,omitempty"`
}
type UserUpdatesDataMangaElement struct {
	ChaptersRead  *int       `json:"chapters_read"`
	ChaptersTotal *int       `json:"chapters_total"`
	Date          *string    `json:"date,omitempty"`
	Entry         *MangaMeta `json:"entry,omitempty"`
	Score         *int       `json:"score"`
	Status        *string    `json:"status,omitempty"`
	VolumesRead   *int       `json:"volumes_read"`
	VolumesTotal  *int       `json:"volumes_total"`
}
type SeasonsDataElement struct {
	Seasons *[]string `json:"seasons,omitempty"`
	Year    *int      `json:"year,omitempty"`
}
type UserFriendsPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type UserStatisticsDataManga struct {
	ChaptersRead *int     `json:"chapters_read,omitempty"`
	Completed    *int     `json:"completed,omitempty"`
	DaysRead     *float32 `json:"days_read,omitempty"`
	Dropped      *int     `json:"dropped,omitempty"`
	MeanScore    *float32 `json:"mean_score,omitempty"`
	OnHold       *int     `json:"on_hold,omitempty"`
	PlanToRead   *int     `json:"plan_to_read,omitempty"`
	Reading      *int     `json:"reading,omitempty"`
	Reread       *int     `json:"reread,omitempty"`
	TotalEntries *int     `json:"total_entries,omitempty"`
	VolumesRead  *int     `json:"volumes_read,omitempty"`
}
type UsersTempDataElementAnimeStats struct {
	Completed       *int     `json:"completed,omitempty"`
	DaysWatched     *float32 `json:"days_watched,omitempty"`
	Dropped         *int     `json:"dropped,omitempty"`
	EpisodesWatched *int     `json:"episodes_watched,omitempty"`
	MeanScore       *float32 `json:"mean_score,omitempty"`
	OnHold          *int     `json:"on_hold,omitempty"`
	PlanToWatch     *int     `json:"plan_to_watch,omitempty"`
	Rewatched       *int     `json:"rewatched,omitempty"`
	TotalEntries    *int     `json:"total_entries,omitempty"`
	Watching        *int     `json:"watching,omitempty"`
}
type GetClubMembersResponseJSON200 struct {
	Data       *[]GetClubMembersResponseJSON200DataElement `json:"data,omitempty"`
	Pagination *GetClubMembersResponseJSON200Pagination    `json:"pagination,omitempty"`
}
type GetUserFullProfileResponseJSON200 struct {
	Data *UserProfileFull `json:"data,omitempty"`
}
type AnimeUserupdatesDataElement struct {
	Date          *string   `json:"date,omitempty"`
	EpisodesSeen  *int      `json:"episodes_seen"`
	EpisodesTotal *int      `json:"episodes_total"`
	Score         *int      `json:"score"`
	Status        *string   `json:"status,omitempty"`
	User          *UserMeta `json:"user,omitempty"`
}
type MangaUserupdatesDataElement struct {
	ChaptersRead  *int      `json:"chapters_read,omitempty"`
	ChaptersTotal *int      `json:"chapters_total,omitempty"`
	Date          *string   `json:"date,omitempty"`
	Score         *int      `json:"score"`
	Status        *string   `json:"status,omitempty"`
	User          *UserMeta `json:"user,omitempty"`
	VolumesRead   *int      `json:"volumes_read,omitempty"`
	VolumesTotal  *int      `json:"volumes_total,omitempty"`
}
type NewsDataElement struct {
	AuthorUrl      *string       `json:"author_url,omitempty"`
	AuthorUsername *string       `json:"author_username,omitempty"`
	Comments       *int          `json:"comments,omitempty"`
	Date           *string       `json:"date,omitempty"`
	Excerpt        *string       `json:"excerpt,omitempty"`
	ForumUrl       *string       `json:"forum_url,omitempty"`
	Images         *CommonImages `json:"images,omitempty"`
	MalId          *int          `json:"mal_id,omitempty"`
	Title          *string       `json:"title,omitempty"`
	Url            *string       `json:"url,omitempty"`
}
type PersonMangaDataElement struct {
	Manga    *MangaMeta `json:"manga,omitempty"`
	Position *string    `json:"position,omitempty"`
}
type UserClubsPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type GetAnimeRelationsResponseJSON200 struct {
	Data *[]Relation `json:"data,omitempty"`
}
type GetRandomMangaResponseJSON200 struct {
	Data *Manga `json:"data,omitempty"`
}
type AnimeNewsDataElement struct {
	AuthorUrl      *string       `json:"author_url,omitempty"`
	AuthorUsername *string       `json:"author_username,omitempty"`
	Comments       *int          `json:"comments,omitempty"`
	Date           *string       `json:"date,omitempty"`
	Excerpt        *string       `json:"excerpt,omitempty"`
	ForumUrl       *string       `json:"forum_url,omitempty"`
	Images         *CommonImages `json:"images,omitempty"`
	MalId          *int          `json:"mal_id,omitempty"`
	Title          *string       `json:"title,omitempty"`
	Url            *string       `json:"url,omitempty"`
}
type AnimeReviewsDataElement struct {
	Date            *string                           `json:"date,omitempty"`
	EpisodesWatched *int                              `json:"episodes_watched,omitempty"`
	IsPreliminary   *bool                             `json:"is_preliminary,omitempty"`
	IsSpoiler       *bool                             `json:"is_spoiler,omitempty"`
	MalId           *int                              `json:"mal_id,omitempty"`
	Reactions       *AnimeReviewsDataElementReactions `json:"reactions,omitempty"`
	Review          *string                           `json:"review,omitempty"`
	Score           *int                              `json:"score,omitempty"`
	Tags            *[]string                         `json:"tags,omitempty"`
	Type            *string                           `json:"type,omitempty"`
	Url             *string                           `json:"url,omitempty"`
	User            *UserMeta                         `json:"user,omitempty"`
}
type ClubsSearchPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type UsersTempDataElementImagesJpg struct {
	ImageUrl *string `json:"image_url,omitempty"`
}
type AnimeNewsPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type AnimeReviewReactions struct {
	Confusing   *int `json:"confusing,omitempty"`
	Creative    *int `json:"creative,omitempty"`
	Funny       *int `json:"funny,omitempty"`
	Informative *int `json:"informative,omitempty"`
	LoveIt      *int `json:"love_it,omitempty"`
	Nice        *int `json:"nice,omitempty"`
	Overall     *int `json:"overall,omitempty"`
	WellWritten *int `json:"well_written,omitempty"`
}
type CommonImagesJpg struct {
	ImageUrl *string `json:"image_url"`
}
type DaterangePropFrom struct {
	Day   *int `json:"day"`
	Month *int `json:"month"`
	Year  *int `json:"year"`
}
type UserFriendsDataElement struct {
	FriendsSince *string   `json:"friends_since,omitempty"`
	LastOnline   *string   `json:"last_online,omitempty"`
	User         *UserMeta `json:"user,omitempty"`
}
type GetUserByIdResponseJSON200 struct {
	Data *UserById `json:"data,omitempty"`
}
type GetUserReviewsResponseJSON200Data struct {
	Data       *[]GetUserReviews_200_Data_Data_Item         `json:"data,omitempty"`
	Pagination *GetUserReviewsResponseJSON200DataPagination `json:"pagination,omitempty"`
}
type GetUserReviewsResponseJSON200DataPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type PaginationPlusPagination struct {
	HasNextPage     *bool                          `json:"has_next_page,omitempty"`
	Items           *PaginationPlusPaginationItems `json:"items,omitempty"`
	LastVisiblePage *int                           `json:"last_visible_page,omitempty"`
}
type GetUserReviewsResponseJSON200 struct {
	Data *GetUserReviewsResponseJSON200Data `json:"data,omitempty"`
}
type GetUserReviews200DataData1Reactions struct {
	Confusing   *int `json:"confusing,omitempty"`
	Creative    *int `json:"creative,omitempty"`
	Funny       *int `json:"funny,omitempty"`
	Informative *int `json:"informative,omitempty"`
	LoveIt      *int `json:"love_it,omitempty"`
	Nice        *int `json:"nice,omitempty"`
	Overall     *int `json:"overall,omitempty"`
	WellWritten *int `json:"well_written,omitempty"`
}
type CharacterVoiceActorsDataElement struct {
	Language *string     `json:"language,omitempty"`
	Person   *PersonMeta `json:"person,omitempty"`
}
type MoreinfoData struct {
	Moreinfo *string `json:"moreinfo"`
}
type PicturesDataElement struct {
	Images *AnimeImages `json:"images,omitempty"`
}
type GetPersonByIdResponseJSON200 struct {
	Data *Person `json:"data,omitempty"`
}
type GetProducerFullByIdResponseJSON200 struct {
	Data *ProducerFull `json:"data,omitempty"`
}
type GetTopReviewsResponseJSON200Data struct {
	Data       *[]GetTopReviews_200_Data_Data_Item         `json:"data,omitempty"`
	Pagination *GetTopReviewsResponseJSON200DataPagination `json:"pagination,omitempty"`
}
type GetUserFavoritesResponseJSON200 struct {
	Data *UserFavorites `json:"data,omitempty"`
}
type AnimeThemesData struct {
	Endings  *[]string `json:"endings,omitempty"`
	Openings *[]string `json:"openings,omitempty"`
}
type AnimeVideosEpisodesDataElement struct {
	Episode *string       `json:"episode,omitempty"`
	Images  *CommonImages `json:"images,omitempty"`
	MalId   *int          `json:"mal_id,omitempty"`
	Title   *string       `json:"title,omitempty"`
	Url     *string       `json:"url,omitempty"`
}
type CharacterMangaDataElement struct {
	Manga *MangaMeta `json:"manga,omitempty"`
	Role  *string    `json:"role,omitempty"`
}
type MangaFullRelationsElement struct {
	Entry    *[]MalUrl `json:"entry,omitempty"`
	Relation *string   `json:"relation,omitempty"`
}
type MangaReviewsDataElementReactions struct {
	Confusing   *int `json:"confusing,omitempty"`
	Creative    *int `json:"creative,omitempty"`
	Funny       *int `json:"funny,omitempty"`
	Informative *int `json:"informative,omitempty"`
	LoveIt      *int `json:"love_it,omitempty"`
	Nice        *int `json:"nice,omitempty"`
	Overall     *int `json:"overall,omitempty"`
	WellWritten *int `json:"well_written,omitempty"`
}
type UsersTempDataElementImages struct {
	Jpg  *UsersTempDataElementImagesJpg  `json:"jpg,omitempty"`
	Webp *UsersTempDataElementImagesWebp `json:"webp,omitempty"`
}
type GetClubMembersResponseJSON200Pagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type GetMangaByIdResponseJSON200 struct {
	Data *Manga `json:"data,omitempty"`
}
type MangaNewsPagination struct {
	HasNextPage     *bool `json:"has_next_page,omitempty"`
	LastVisiblePage *int  `json:"last_visible_page,omitempty"`
}
type UserStatisticsDataAnime struct {
	Completed       *int     `json:"completed,omitempty"`
	DaysWatched     *float32 `json:"days_watched,omitempty"`
	Dropped         *int     `json:"dropped,omitempty"`
	EpisodesWatched *int     `json:"episodes_watched,omitempty"`
	MeanScore       *float32 `json:"mean_score,omitempty"`
	OnHold          *int     `json:"on_hold,omitempty"`
	PlanToWatch     *int     `json:"plan_to_watch,omitempty"`
	Rewatched       *int     `json:"rewatched,omitempty"`
	TotalEntries    *int     `json:"total_entries,omitempty"`
	Watching        *int     `json:"watching,omitempty"`
}
type UsersTempDataElementMangaStats struct {
	ChaptersRead *int     `json:"chapters_read,omitempty"`
	Completed    *int     `json:"completed,omitempty"`
	DaysRead     *float32 `json:"days_read,omitempty"`
	Dropped      *int     `json:"dropped,omitempty"`
	MeanScore    *float32 `json:"mean_score,omitempty"`
	OnHold       *int     `json:"on_hold,omitempty"`
	PlanToRead   *int     `json:"plan_to_read,omitempty"`
	Reading      *int     `json:"reading,omitempty"`
	Reread       *int     `json:"reread,omitempty"`
	TotalEntries *int     `json:"total_entries,omitempty"`
	VolumesRead  *int     `json:"volumes_read,omitempty"`
}
type GetCharacterByIdResponseJSON200 struct {
	Data *Character `json:"data,omitempty"`
}
type PersonFullAnimeElement struct {
	Anime    *AnimeMeta `json:"anime,omitempty"`
	Position *string    `json:"position,omitempty"`
}
type UserClubsDataElement struct {
	MalId *int    `json:"mal_id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Url   *string `json:"url,omitempty"`
}
type UserFavoritesCharactersElement struct {
	Images *CharacterImages `json:"images,omitempty"`
	MalId  *int             `json:"mal_id,omitempty"`
	Name   *string          `json:"name,omitempty"`
	Title  *string          `json:"title,omitempty"`
	Type   *string          `json:"type,omitempty"`
	Url    *string          `json:"url,omitempty"`
}
type AnimeCharactersDataElementCharacter struct {
	Images *CharacterImages `json:"images,omitempty"`
	MalId  *int             `json:"mal_id,omitempty"`
	Name   *string          `json:"name,omitempty"`
	Url    *string          `json:"url,omitempty"`
}
type CharacterFullVoicesElement struct {
	Language *string     `json:"language,omitempty"`
	Person   *PersonMeta `json:"person,omitempty"`
}
type EntryRecommendationsDataElement struct {
	Entry *EntryRecommendations_Data_Entry `json:"entry,omitempty"`
}
type RecommendationsDataElement struct {
	Content *string                            `json:"content,omitempty"`
	Entry   *[]Recommendations_Data_Entry_Item `json:"entry,omitempty"`
	MalId   *string                            `json:"mal_id,omitempty"`
	User    *UserById                          `json:"user,omitempty"`
}
type UserProfileFullStatistics struct {
	Anime *UserProfileFullStatisticsAnime `json:"anime,omitempty"`
	Manga *UserProfileFullStatisticsManga `json:"manga,omitempty"`
}
type GetRandomPeopleResponseJSON200 struct {
	Data *Person `json:"data,omitempty"`
}
type CharacterImagesWebp struct {
	ImageUrl      *string `json:"image_url"`
	SmallImageUrl *string `json:"small_image_url"`
}
type PeopleSearchPagination struct {
	HasNextPage     *bool                        `json:"has_next_page,omitempty"`
	Items           *PeopleSearchPaginationItems `json:"items,omitempty"`
	LastVisiblePage *int                         `json:"last_visible_page,omitempty"`
}
type StreamingLinksDataElement struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}
type WatchPromosDataElement struct {
	Entry   *AnimeMeta `json:"entry,omitempty"`
	Trailer *[]Trailer `json:"trailer,omitempty"`
}
type Anime struct {
	Aired          *Daterange   `json:"aired,omitempty"`
	Airing         *bool        `json:"airing,omitempty"`
	Approved       *bool        `json:"approved,omitempty"`
	Background     *string      `json:"background"`
	Broadcast      *Broadcast   `json:"broadcast,omitempty"`
	Demographics   *[]MalUrl    `json:"demographics,omitempty"`
	Duration       *string      `json:"duration"`
	Episodes       *int         `json:"episodes"`
	ExplicitGenres *[]MalUrl    `json:"explicit_genres,omitempty"`
	Favorites      *int         `json:"favorites"`
	Genres         *[]MalUrl    `json:"genres,omitempty"`
	Images         *AnimeImages `json:"images,omitempty"`
	Licensors      *[]MalUrl    `json:"licensors,omitempty"`
	MalId          *int         `json:"mal_id,omitempty"`
	Members        *int         `json:"members"`
	Popularity     *int         `json:"popularity"`
	Producers      *[]MalUrl    `json:"producers,omitempty"`
	Rank           *int         `json:"rank"`
	Rating         *AnimeRating `json:"rating"`
	Score          *float32     `json:"score"`
	ScoredBy       *int         `json:"scored_by"`
	Season         *AnimeSeason `json:"season"`
	Source         *string      `json:"source"`
	Status         *AnimeStatus `json:"status"`
	Studios        *[]MalUrl    `json:"studios,omitempty"`
	Synopsis       *string      `json:"synopsis"`
	Themes         *[]MalUrl    `json:"themes,omitempty"`
	Title          *string      `json:"title,omitempty"`
	TitleEnglish   *string      `json:"title_english"`
	TitleJapanese  *string      `json:"title_japanese"`
	TitleSynonyms  *[]string    `json:"title_synonyms,omitempty"`
	Titles         *[]Title     `json:"titles,omitempty"`
	Trailer        *TrailerBase `json:"trailer,omitempty"`
	Type           *AnimeType   `json:"type"`
	Url            *string      `json:"url,omitempty"`
	Year           *int         `json:"year"`
}
type AnimeRating string
type AnimeSeason string
type AnimeStatus string
type AnimeType string
type AnimeCharacters struct {
	Data *[]AnimeCharactersDataElement `json:"data,omitempty"`
}
type AnimeEpisode struct {
	Aired         *string `json:"aired"`
	Duration      *int    `json:"duration"`
	Filler        *bool   `json:"filler,omitempty"`
	MalId         *int    `json:"mal_id,omitempty"`
	Recap         *bool   `json:"recap,omitempty"`
	Synopsis      *string `json:"synopsis"`
	Title         *string `json:"title,omitempty"`
	TitleJapanese *string `json:"title_japanese"`
	TitleRomanji  *string `json:"title_romanji"`
	Url           *string `json:"url,omitempty"`
}
type AnimeEpisodes struct {
	Data       *[]AnimeEpisodesDataElement `json:"data,omitempty"`
	Pagination *AnimeEpisodesPagination    `json:"pagination,omitempty"`
}
type AnimeFull struct {
	Aired          *Daterange                   `json:"aired,omitempty"`
	Airing         *bool                        `json:"airing,omitempty"`
	Approved       *bool                        `json:"approved,omitempty"`
	Background     *string                      `json:"background"`
	Broadcast      *Broadcast                   `json:"broadcast,omitempty"`
	Demographics   *[]MalUrl                    `json:"demographics,omitempty"`
	Duration       *string                      `json:"duration"`
	Episodes       *int                         `json:"episodes"`
	ExplicitGenres *[]MalUrl                    `json:"explicit_genres,omitempty"`
	External       *[]AnimeFullExternalElement  `json:"external,omitempty"`
	Favorites      *int                         `json:"favorites"`
	Genres         *[]MalUrl                    `json:"genres,omitempty"`
	Images         *AnimeImages                 `json:"images,omitempty"`
	Licensors      *[]MalUrl                    `json:"licensors,omitempty"`
	MalId          *int                         `json:"mal_id,omitempty"`
	Members        *int                         `json:"members"`
	Popularity     *int                         `json:"popularity"`
	Producers      *[]MalUrl                    `json:"producers,omitempty"`
	Rank           *int                         `json:"rank"`
	Rating         *AnimeFullRating             `json:"rating"`
	Relations      *[]AnimeFullRelationsElement `json:"relations,omitempty"`
	Score          *float32                     `json:"score"`
	ScoredBy       *int                         `json:"scored_by"`
	Season         *AnimeFullSeason             `json:"season"`
	Source         *string                      `json:"source"`
	Status         *AnimeFullStatus             `json:"status"`
	Streaming      *[]AnimeFullStreamingElement `json:"streaming,omitempty"`
	Studios        *[]MalUrl                    `json:"studios,omitempty"`
	Synopsis       *string                      `json:"synopsis"`
	Theme          *AnimeFullTheme              `json:"theme,omitempty"`
	Themes         *[]MalUrl                    `json:"themes,omitempty"`
	Title          *string                      `json:"title,omitempty"`
	TitleEnglish   *string                      `json:"title_english"`
	TitleJapanese  *string                      `json:"title_japanese"`
	TitleSynonyms  *[]string                    `json:"title_synonyms,omitempty"`
	Titles         *[]Title                     `json:"titles,omitempty"`
	Trailer        *TrailerBase                 `json:"trailer,omitempty"`
	Type           *AnimeFullType               `json:"type"`
	Url            *string                      `json:"url,omitempty"`
	Year           *int                         `json:"year"`
}
type AnimeFullRating string
type AnimeFullSeason string
type AnimeFullStatus string
type AnimeFullType string
type AnimeImages struct {
	Jpg  *AnimeImagesJpg  `json:"jpg,omitempty"`
	Webp *AnimeImagesWebp `json:"webp,omitempty"`
}
type AnimeMeta struct {
	Images *AnimeImages `json:"images,omitempty"`
	MalId  *int         `json:"mal_id,omitempty"`
	Title  *string      `json:"title,omitempty"`
	Url    *string      `json:"url,omitempty"`
}
type AnimeNews struct {
	Data       *[]AnimeNewsDataElement `json:"data,omitempty"`
	Pagination *AnimeNewsPagination    `json:"pagination,omitempty"`
}
type AnimeRelations struct {
	Data *[]AnimeRelationsDataElement `json:"data,omitempty"`
}
type AnimeReview struct {
	Date            *string               `json:"date,omitempty"`
	EpisodesWatched *int                  `json:"episodes_watched,omitempty"`
	IsPreliminary   *bool                 `json:"is_preliminary,omitempty"`
	IsSpoiler       *bool                 `json:"is_spoiler,omitempty"`
	MalId           *int                  `json:"mal_id,omitempty"`
	Reactions       *AnimeReviewReactions `json:"reactions,omitempty"`
	Review          *string               `json:"review,omitempty"`
	Score           *int                  `json:"score,omitempty"`
	Tags            *[]string             `json:"tags,omitempty"`
	Type            *string               `json:"type,omitempty"`
	Url             *string               `json:"url,omitempty"`
}
type AnimeReviews struct {
	Data       *[]AnimeReviewsDataElement `json:"data,omitempty"`
	Pagination *AnimeReviewsPagination    `json:"pagination,omitempty"`
}
type AnimeSearch struct {
	Data       *[]Anime               `json:"data,omitempty"`
	Pagination *AnimeSearchPagination `json:"pagination,omitempty"`
}
type AnimeSearchQueryOrderby string
type AnimeSearchQueryRating string
type AnimeSearchQueryStatus string
type AnimeSearchQueryType string
type AnimeStaff struct {
	Data *[]AnimeStaffDataElement `json:"data,omitempty"`
}
type AnimeStatistics struct {
	Data *AnimeStatisticsData `json:"data,omitempty"`
}
type AnimeThemes struct {
	Data *AnimeThemesData `json:"data,omitempty"`
}
type AnimeUserupdates struct {
	Data       *[]AnimeUserupdatesDataElement `json:"data,omitempty"`
	Pagination *AnimeUserupdatesPagination    `json:"pagination,omitempty"`
}
type AnimeVideos struct {
	Data *AnimeVideosData `json:"data,omitempty"`
}
type AnimeVideosEpisodes struct {
	Data       *[]AnimeVideosEpisodesDataElement `json:"data,omitempty"`
	Pagination *AnimeVideosEpisodesPagination    `json:"pagination,omitempty"`
}
type Broadcast struct {
	Day      *string `json:"day"`
	String   *string `json:"string"`
	Time     *string `json:"time"`
	Timezone *string `json:"timezone"`
}
type Character struct {
	About     *string          `json:"about"`
	Favorites *int             `json:"favorites,omitempty"`
	Images    *CharacterImages `json:"images,omitempty"`
	MalId     *int             `json:"mal_id,omitempty"`
	Name      *string          `json:"name,omitempty"`
	NameKanji *string          `json:"name_kanji"`
	Nicknames *[]string        `json:"nicknames,omitempty"`
	Url       *string          `json:"url,omitempty"`
}
type CharacterAnime struct {
	Data *[]CharacterAnimeDataElement `json:"data,omitempty"`
}
type CharacterFull struct {
	About     *string                       `json:"about"`
	Anime     *[]CharacterFullAnimeElement  `json:"anime,omitempty"`
	Favorites *int                          `json:"favorites,omitempty"`
	Images    *CharacterImages              `json:"images,omitempty"`
	MalId     *int                          `json:"mal_id,omitempty"`
	Manga     *[]CharacterFullMangaElement  `json:"manga,omitempty"`
	Name      *string                       `json:"name,omitempty"`
	NameKanji *string                       `json:"name_kanji"`
	Nicknames *[]string                     `json:"nicknames,omitempty"`
	Url       *string                       `json:"url,omitempty"`
	Voices    *[]CharacterFullVoicesElement `json:"voices,omitempty"`
}
type CharacterImages struct {
	Jpg  *CharacterImagesJpg  `json:"jpg,omitempty"`
	Webp *CharacterImagesWebp `json:"webp,omitempty"`
}
type CharacterManga struct {
	Data *[]CharacterMangaDataElement `json:"data,omitempty"`
}
type CharacterMeta struct {
	Images *CharacterImages `json:"images,omitempty"`
	MalId  *int             `json:"mal_id,omitempty"`
	Name   *string          `json:"name,omitempty"`
	Url    *string          `json:"url,omitempty"`
}
type CharacterPictures struct {
	Data *[]CharacterPicturesDataElement `json:"data,omitempty"`
}
type CharacterVoiceActors struct {
	Data *[]CharacterVoiceActorsDataElement `json:"data,omitempty"`
}
type CharactersSearch struct {
	Data       *[]Character                `json:"data,omitempty"`
	Pagination *CharactersSearchPagination `json:"pagination,omitempty"`
}
type CharactersSearchQueryOrderby string
type Club struct {
	Access   *ClubAccess   `json:"access,omitempty"`
	Category *ClubCategory `json:"category,omitempty"`
	Created  *string       `json:"created,omitempty"`
	Images   *CommonImages `json:"images,omitempty"`
	MalId    *int          `json:"mal_id,omitempty"`
	Members  *int          `json:"members,omitempty"`
	Name     *string       `json:"name,omitempty"`
	Url      *string       `json:"url,omitempty"`
}
type ClubAccess string
type ClubCategory string
type ClubMember struct {
	Data *[]ClubMemberDataElement `json:"data,omitempty"`
}
type ClubRelations struct {
	Data *ClubRelationsData `json:"data,omitempty"`
}
type ClubSearchQueryCategory string
type ClubSearchQueryOrderby string
type ClubSearchQueryType string
type ClubStaff struct {
	Data *[]ClubStaffDataElement `json:"data,omitempty"`
}
type ClubsSearch struct {
	Data       *[]Club                `json:"data,omitempty"`
	Pagination *ClubsSearchPagination `json:"pagination,omitempty"`
}
type CommonImages struct {
	Jpg *CommonImagesJpg `json:"jpg,omitempty"`
}
type Daterange struct {
	From *string        `json:"from"`
	Prop *DaterangeProp `json:"prop,omitempty"`
	To   *string        `json:"to"`
}
type EntryMeta struct {
	ImageUrl *string `json:"image_url,omitempty"`
	MalId    *int    `json:"mal_id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Url      *string `json:"url,omitempty"`
}
type EntryRecommendations struct {
	Data *[]EntryRecommendationsDataElement `json:"data,omitempty"`
}
type EntryRecommendations_Data_Entry struct {
	union json.RawMessage
}
type ExternalLinks struct {
	Data *[]ExternalLinksDataElement `json:"data,omitempty"`
}
type Forum struct {
	Data *[]ForumDataElement `json:"data,omitempty"`
}
type Genre struct {
	Count *int    `json:"count,omitempty"`
	MalId *int    `json:"mal_id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Url   *string `json:"url,omitempty"`
}
type GenreQueryFilter string
type Genres struct {
	Data *[]Genre `json:"data,omitempty"`
}
type History struct {
	Date      *string `json:"date,omitempty"`
	Entry     *MalUrl `json:"entry,omitempty"`
	Increment *int    `json:"increment,omitempty"`
}
type Magazine struct {
	Count *int    `json:"count,omitempty"`
	MalId *int    `json:"mal_id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Url   *string `json:"url,omitempty"`
}
type Magazines struct {
	Data       *[]Magazine          `json:"data,omitempty"`
	Pagination *MagazinesPagination `json:"pagination,omitempty"`
}
type MagazinesQueryOrderby string
type MalUrl struct {
	MalId *int    `json:"mal_id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Type  *string `json:"type,omitempty"`
	Url   *string `json:"url,omitempty"`
}
type MalUrl2 struct {
	MalId *int    `json:"mal_id,omitempty"`
	Title *string `json:"title,omitempty"`
	Type  *string `json:"type,omitempty"`
	Url   *string `json:"url,omitempty"`
}
type Manga struct {
	Approved       *bool        `json:"approved,omitempty"`
	Authors        *[]MalUrl    `json:"authors,omitempty"`
	Background     *string      `json:"background"`
	Chapters       *int         `json:"chapters"`
	Demographics   *[]MalUrl    `json:"demographics,omitempty"`
	ExplicitGenres *[]MalUrl    `json:"explicit_genres,omitempty"`
	Favorites      *int         `json:"favorites"`
	Genres         *[]MalUrl    `json:"genres,omitempty"`
	Images         *MangaImages `json:"images,omitempty"`
	MalId          *int         `json:"mal_id,omitempty"`
	Members        *int         `json:"members"`
	Popularity     *int         `json:"popularity"`
	Published      *Daterange   `json:"published,omitempty"`
	Publishing     *bool        `json:"publishing,omitempty"`
	Rank           *int         `json:"rank"`
	Score          *float32     `json:"score"`
	ScoredBy       *int         `json:"scored_by"`
	Serializations *[]MalUrl    `json:"serializations,omitempty"`
	Status         *MangaStatus `json:"status,omitempty"`
	Synopsis       *string      `json:"synopsis"`
	Themes         *[]MalUrl    `json:"themes,omitempty"`
	Title          *string      `json:"title,omitempty"`
	TitleEnglish   *string      `json:"title_english"`
	TitleJapanese  *string      `json:"title_japanese"`
	Titles         *[]Title     `json:"titles,omitempty"`
	Type           *MangaType   `json:"type"`
	Url            *string      `json:"url,omitempty"`
	Volumes        *int         `json:"volumes"`
}
type MangaStatus string
type MangaType string
type MangaCharacters struct {
	Data *[]MangaCharactersDataElement `json:"data,omitempty"`
}
type MangaFull struct {
	Approved       *bool                        `json:"approved,omitempty"`
	Authors        *[]MalUrl                    `json:"authors,omitempty"`
	Background     *string                      `json:"background"`
	Chapters       *int                         `json:"chapters"`
	Demographics   *[]MalUrl                    `json:"demographics,omitempty"`
	ExplicitGenres *[]MalUrl                    `json:"explicit_genres,omitempty"`
	External       *[]MangaFullExternalElement  `json:"external,omitempty"`
	Favorites      *int                         `json:"favorites"`
	Genres         *[]MalUrl                    `json:"genres,omitempty"`
	Images         *MangaImages                 `json:"images,omitempty"`
	MalId          *int                         `json:"mal_id,omitempty"`
	Members        *int                         `json:"members"`
	Popularity     *int                         `json:"popularity"`
	Published      *Daterange                   `json:"published,omitempty"`
	Publishing     *bool                        `json:"publishing,omitempty"`
	Rank           *int                         `json:"rank"`
	Relations      *[]MangaFullRelationsElement `json:"relations,omitempty"`
	Score          *float32                     `json:"score"`
	ScoredBy       *int                         `json:"scored_by"`
	Serializations *[]MalUrl                    `json:"serializations,omitempty"`
	Status         *MangaFullStatus             `json:"status,omitempty"`
	Synopsis       *string                      `json:"synopsis"`
	Themes         *[]MalUrl                    `json:"themes,omitempty"`
	Title          *string                      `json:"title,omitempty"`
	TitleEnglish   *string                      `json:"title_english"`
	TitleJapanese  *string                      `json:"title_japanese"`
	TitleSynonyms  *[]string                    `json:"title_synonyms,omitempty"`
	Titles         *[]Title                     `json:"titles,omitempty"`
	Type           *MangaFullType               `json:"type"`
	Url            *string                      `json:"url,omitempty"`
	Volumes        *int                         `json:"volumes"`
}
type MangaFullStatus string
type MangaFullType string
type MangaImages struct {
	Jpg  *MangaImagesJpg  `json:"jpg,omitempty"`
	Webp *MangaImagesWebp `json:"webp,omitempty"`
}
type MangaMeta struct {
	Images *MangaImages `json:"images,omitempty"`
	MalId  *int         `json:"mal_id,omitempty"`
	Title  *string      `json:"title,omitempty"`
	Url    *string      `json:"url,omitempty"`
}
type MangaNews struct {
	Data       *[]MangaNewsDataElement `json:"data,omitempty"`
	Pagination *MangaNewsPagination    `json:"pagination,omitempty"`
}
type MangaPictures struct {
	Data *[]MangaImages `json:"data,omitempty"`
}
type MangaReview struct {
	Date          *string               `json:"date,omitempty"`
	IsPreliminary *bool                 `json:"is_preliminary,omitempty"`
	IsSpoiler     *bool                 `json:"is_spoiler,omitempty"`
	MalId         *int                  `json:"mal_id,omitempty"`
	Reactions     *MangaReviewReactions `json:"reactions,omitempty"`
	Review        *string               `json:"review,omitempty"`
	Score         *int                  `json:"score,omitempty"`
	Tags          *[]string             `json:"tags,omitempty"`
	Type          *string               `json:"type,omitempty"`
	Url           *string               `json:"url,omitempty"`
}
type MangaReviews struct {
	Data       *[]MangaReviewsDataElement `json:"data,omitempty"`
	Pagination *MangaReviewsPagination    `json:"pagination,omitempty"`
}
type MangaSearch struct {
	Data       *[]Manga               `json:"data,omitempty"`
	Pagination *MangaSearchPagination `json:"pagination,omitempty"`
}
type MangaSearchQueryOrderby string
type MangaSearchQueryStatus string
type MangaSearchQueryType string
type MangaStatistics struct {
	Data *MangaStatisticsData `json:"data,omitempty"`
}
type MangaUserupdates struct {
	Data       *[]MangaUserupdatesDataElement `json:"data,omitempty"`
	Pagination *MangaUserupdatesPagination    `json:"pagination,omitempty"`
}
type Moreinfo struct {
	Data *MoreinfoData `json:"data,omitempty"`
}
type News struct {
	Data *[]NewsDataElement `json:"data,omitempty"`
}
type Pagination struct {
	Pagination *PaginationPagination `json:"pagination,omitempty"`
}
type PaginationPlus struct {
	Pagination *PaginationPlusPagination `json:"pagination,omitempty"`
}
type PeopleImages struct {
	Jpg *PeopleImagesJpg `json:"jpg,omitempty"`
}
type PeopleSearch struct {
	Data       *[]Person               `json:"data,omitempty"`
	Pagination *PeopleSearchPagination `json:"pagination,omitempty"`
}
type PeopleSearchQueryOrderby string
type Person struct {
	About          *string       `json:"about"`
	AlternateNames *[]string     `json:"alternate_names,omitempty"`
	Birthday       *string       `json:"birthday"`
	FamilyName     *string       `json:"family_name"`
	Favorites      *int          `json:"favorites,omitempty"`
	GivenName      *string       `json:"given_name"`
	Images         *PeopleImages `json:"images,omitempty"`
	MalId          *int          `json:"mal_id,omitempty"`
	Name           *string       `json:"name,omitempty"`
	Url            *string       `json:"url,omitempty"`
	WebsiteUrl     *string       `json:"website_url"`
}
type PersonAnime struct {
	Data *[]PersonAnimeDataElement `json:"data,omitempty"`
}
type PersonFull struct {
	About          *string                    `json:"about"`
	AlternateNames *[]string                  `json:"alternate_names,omitempty"`
	Anime          *[]PersonFullAnimeElement  `json:"anime,omitempty"`
	Birthday       *string                    `json:"birthday"`
	FamilyName     *string                    `json:"family_name"`
	Favorites      *int                       `json:"favorites,omitempty"`
	GivenName      *string                    `json:"given_name"`
	Images         *PeopleImages              `json:"images,omitempty"`
	MalId          *int                       `json:"mal_id,omitempty"`
	Manga          *[]PersonFullMangaElement  `json:"manga,omitempty"`
	Name           *string                    `json:"name,omitempty"`
	Url            *string                    `json:"url,omitempty"`
	Voices         *[]PersonFullVoicesElement `json:"voices,omitempty"`
	WebsiteUrl     *string                    `json:"website_url"`
}
type PersonManga struct {
	Data *[]PersonMangaDataElement `json:"data,omitempty"`
}
type PersonMeta struct {
	Images *PeopleImages `json:"images,omitempty"`
	MalId  *int          `json:"mal_id,omitempty"`
	Name   *string       `json:"name,omitempty"`
	Url    *string       `json:"url,omitempty"`
}
type PersonPictures struct {
	Data *[]PeopleImages `json:"data,omitempty"`
}
type PersonVoiceActingRoles struct {
	Data *[]PersonVoiceActingRolesDataElement `json:"data,omitempty"`
}
type Pictures struct {
	Data *[]PicturesDataElement `json:"data,omitempty"`
}
type PicturesVariants struct {
	Data *[]PicturesVariantsDataElement `json:"data,omitempty"`
}
type Producer struct {
	About       *string       `json:"about"`
	Count       *int          `json:"count,omitempty"`
	Established *string       `json:"established"`
	Favorites   *int          `json:"favorites,omitempty"`
	Images      *CommonImages `json:"images,omitempty"`
	MalId       *int          `json:"mal_id,omitempty"`
	Titles      *[]Title      `json:"titles,omitempty"`
	Url         *string       `json:"url,omitempty"`
}
type ProducerFull struct {
	About       *string                        `json:"about"`
	Count       *int                           `json:"count,omitempty"`
	Established *string                        `json:"established"`
	External    *[]ProducerFullExternalElement `json:"external,omitempty"`
	Favorites   *int                           `json:"favorites,omitempty"`
	Images      *CommonImages                  `json:"images,omitempty"`
	MalId       *int                           `json:"mal_id,omitempty"`
	Titles      *[]Title                       `json:"titles,omitempty"`
	Url         *string                        `json:"url,omitempty"`
}
type Producers struct {
	Data       *[]Producer          `json:"data,omitempty"`
	Pagination *ProducersPagination `json:"pagination,omitempty"`
}
type ProducersQueryOrderby string
type Random struct {
	Data *[]Random_Data_Item `json:"data,omitempty"`
}
type Random_Data_Item struct {
	union json.RawMessage
}
type Recommendations struct {
	Data       *[]RecommendationsDataElement `json:"data,omitempty"`
	Pagination *RecommendationsPagination    `json:"pagination,omitempty"`
}
type Recommendations_Data_Entry_Item struct {
	union json.RawMessage
}
type Relation struct {
	Entry    *[]MalUrl `json:"entry,omitempty"`
	Relation *string   `json:"relation,omitempty"`
}
type ReviewsCollection struct {
	Data *[]ReviewsCollection_Data_Item `json:"data,omitempty"`
}
type ReviewsCollection_Data_Item struct {
	union json.RawMessage
}
type Schedules struct {
	Data       *[]Anime             `json:"data,omitempty"`
	Pagination *SchedulesPagination `json:"pagination,omitempty"`
}
type SearchQuerySort string
type Seasons struct {
	Data *[]SeasonsDataElement `json:"data,omitempty"`
}
type StreamingLinks struct {
	Data *[]StreamingLinksDataElement `json:"data,omitempty"`
}
type Title struct {
	Title *string `json:"title,omitempty"`
	Type  *string `json:"type,omitempty"`
}
type TopAnimeFilter string
type TopMangaFilter string
type TopReviewsTypeEnum string
type Trailer struct {
	EmbedUrl  *string         `json:"embed_url"`
	Images    *TrailerImages1 `json:"images,omitempty"`
	Url       *string         `json:"url"`
	YoutubeId *string         `json:"youtube_id"`
}
type TrailerBase struct {
	EmbedUrl  *string `json:"embed_url"`
	Url       *string `json:"url"`
	YoutubeId *string `json:"youtube_id"`
}
type TrailerImages struct {
	Images *TrailerImagesImages `json:"images,omitempty"`
}
type UserAbout struct {
	Data *[]UserAboutDataElement `json:"data,omitempty"`
}
type UserAnimeListStatusFilter string
type UserById struct {
	Url      *string `json:"url,omitempty"`
	Username *string `json:"username,omitempty"`
}
type UserClubs struct {
	Data       *[]UserClubsDataElement `json:"data,omitempty"`
	Pagination *UserClubsPagination    `json:"pagination,omitempty"`
}
type UserFavorites struct {
	Anime      *[]UserFavoritesAnimeElement      `json:"anime,omitempty"`
	Characters *[]UserFavoritesCharactersElement `json:"characters,omitempty"`
	Manga      *[]UserFavoritesMangaElement      `json:"manga,omitempty"`
	People     *[]CharacterMeta                  `json:"people,omitempty"`
}
type UserFriends struct {
	Data       *[]UserFriendsDataElement `json:"data,omitempty"`
	Pagination *UserFriendsPagination    `json:"pagination,omitempty"`
}
type UserHistory struct {
	Data *[]History `json:"data,omitempty"`
}
type UserImages struct {
	Jpg  *UserImagesJpg  `json:"jpg,omitempty"`
	Webp *UserImagesWebp `json:"webp,omitempty"`
}
type UserMangaListStatusFilter string
type UserMeta struct {
	Images   *UserImages `json:"images,omitempty"`
	Url      *string     `json:"url,omitempty"`
	Username *string     `json:"username,omitempty"`
}
type UserProfile struct {
	Birthday   *string     `json:"birthday"`
	Gender     *string     `json:"gender"`
	Images     *UserImages `json:"images,omitempty"`
	Joined     *string     `json:"joined"`
	LastOnline *string     `json:"last_online"`
	Location   *string     `json:"location"`
	MalId      *int        `json:"mal_id"`
	Url        *string     `json:"url,omitempty"`
	Username   *string     `json:"username,omitempty"`
}
type UserProfileFull struct {
	Birthday   *string                           `json:"birthday"`
	External   *[]UserProfileFullExternalElement `json:"external,omitempty"`
	Gender     *string                           `json:"gender"`
	Images     *UserImages                       `json:"images,omitempty"`
	Joined     *string                           `json:"joined"`
	LastOnline *string                           `json:"last_online"`
	Location   *string                           `json:"location"`
	MalId      *int                              `json:"mal_id"`
	Statistics *UserProfileFullStatistics        `json:"statistics,omitempty"`
	Url        *string                           `json:"url,omitempty"`
	Username   *string                           `json:"username,omitempty"`
}
type UserStatistics struct {
	Data *UserStatisticsData `json:"data,omitempty"`
}
type UserUpdates struct {
	Data *UserUpdatesData `json:"data,omitempty"`
}
type UsersSearch struct {
	Data       *[]UsersSearchDataElement `json:"data,omitempty"`
	Pagination *UsersSearchPagination    `json:"pagination,omitempty"`
}
type UsersSearchQueryGender string
type UsersTemp struct {
	Data *[]UsersTempDataElement `json:"data,omitempty"`
}
type WatchEpisodes struct {
	Data       *[]WatchEpisodesDataElement `json:"data,omitempty"`
	Pagination *WatchEpisodesPagination    `json:"pagination,omitempty"`
}
type WatchPromos struct {
	Data       *[]WatchPromosDataElement `json:"data,omitempty"`
	Pagination *WatchPromosPagination    `json:"pagination,omitempty"`
	Title      *string                   `json:"title,omitempty"`
}
type Continuing = bool
type Kids = bool
type Limit = int
type Page = int
type Preliminary = bool
type Sfw = bool
type Spoilers = bool
type Unapproved = bool
type GetAnimeSearchParams struct {
	Unapproved    *Unapproved              `form:"unapproved,omitempty" json:"unapproved,omitempty"`
	Page          *Page                    `form:"page,omitempty" json:"page,omitempty"`
	Limit         *Limit                   `form:"limit,omitempty" json:"limit,omitempty"`
	Q             *string                  `form:"q,omitempty" json:"q,omitempty"`
	Type          *AnimeSearchQueryType    `form:"type,omitempty" json:"type,omitempty"`
	Score         *float32                 `form:"score,omitempty" json:"score,omitempty"`
	MinScore      *float32                 `form:"min_score,omitempty" json:"min_score,omitempty"`
	MaxScore      *float32                 `form:"max_score,omitempty" json:"max_score,omitempty"`
	Status        *AnimeSearchQueryStatus  `form:"status,omitempty" json:"status,omitempty"`
	Rating        *AnimeSearchQueryRating  `form:"rating,omitempty" json:"rating,omitempty"`
	Sfw           *bool                    `form:"sfw,omitempty" json:"sfw,omitempty"`
	Genres        *string                  `form:"genres,omitempty" json:"genres,omitempty"`
	GenresExclude *string                  `form:"genres_exclude,omitempty" json:"genres_exclude,omitempty"`
	OrderBy       *AnimeSearchQueryOrderby `form:"order_by,omitempty" json:"order_by,omitempty"`
	Sort          *SearchQuerySort         `form:"sort,omitempty" json:"sort,omitempty"`
	Letter        *string                  `form:"letter,omitempty" json:"letter,omitempty"`
	Producers     *string                  `form:"producers,omitempty" json:"producers,omitempty"`
	StartDate     *string                  `form:"start_date,omitempty" json:"start_date,omitempty"`
	EndDate       *string                  `form:"end_date,omitempty" json:"end_date,omitempty"`
}
type GetAnimeEpisodesParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetAnimeForumParams struct {
	Filter *GetAnimeForumParamsFilter `form:"filter,omitempty" json:"filter,omitempty"`
}
type GetAnimeForumParamsFilter string
type GetAnimeNewsParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetAnimeReviewsParams struct {
	Page        *Page        `form:"page,omitempty" json:"page,omitempty"`
	Preliminary *Preliminary `form:"preliminary,omitempty" json:"preliminary,omitempty"`
	Spoilers    *Spoilers    `form:"spoilers,omitempty" json:"spoilers,omitempty"`
}
type GetAnimeUserUpdatesParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetAnimeVideosEpisodesParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetCharactersSearchParams struct {
	Page    *Page                         `form:"page,omitempty" json:"page,omitempty"`
	Limit   *Limit                        `form:"limit,omitempty" json:"limit,omitempty"`
	Q       *string                       `form:"q,omitempty" json:"q,omitempty"`
	OrderBy *CharactersSearchQueryOrderby `form:"order_by,omitempty" json:"order_by,omitempty"`
	Sort    *SearchQuerySort              `form:"sort,omitempty" json:"sort,omitempty"`
	Letter  *string                       `form:"letter,omitempty" json:"letter,omitempty"`
}
type GetClubsSearchParams struct {
	Page     *Page                    `form:"page,omitempty" json:"page,omitempty"`
	Limit    *Limit                   `form:"limit,omitempty" json:"limit,omitempty"`
	Q        *string                  `form:"q,omitempty" json:"q,omitempty"`
	Type     *ClubSearchQueryType     `form:"type,omitempty" json:"type,omitempty"`
	Category *ClubSearchQueryCategory `form:"category,omitempty" json:"category,omitempty"`
	OrderBy  *ClubSearchQueryOrderby  `form:"order_by,omitempty" json:"order_by,omitempty"`
	Sort     *SearchQuerySort         `form:"sort,omitempty" json:"sort,omitempty"`
	Letter   *string                  `form:"letter,omitempty" json:"letter,omitempty"`
}
type GetClubMembersParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetAnimeGenresParams struct {
	Filter *GenreQueryFilter `form:"filter,omitempty" json:"filter,omitempty"`
}
type GetMangaGenresParams struct {
	Filter *GenreQueryFilter `form:"filter,omitempty" json:"filter,omitempty"`
}
type GetMagazinesParams struct {
	Page    *Page                  `form:"page,omitempty" json:"page,omitempty"`
	Limit   *Limit                 `form:"limit,omitempty" json:"limit,omitempty"`
	Q       *string                `form:"q,omitempty" json:"q,omitempty"`
	OrderBy *MagazinesQueryOrderby `form:"order_by,omitempty" json:"order_by,omitempty"`
	Sort    *SearchQuerySort       `form:"sort,omitempty" json:"sort,omitempty"`
	Letter  *string                `form:"letter,omitempty" json:"letter,omitempty"`
}
type GetMangaSearchParams struct {
	Unapproved    *Unapproved              `form:"unapproved,omitempty" json:"unapproved,omitempty"`
	Page          *Page                    `form:"page,omitempty" json:"page,omitempty"`
	Limit         *Limit                   `form:"limit,omitempty" json:"limit,omitempty"`
	Q             *string                  `form:"q,omitempty" json:"q,omitempty"`
	Type          *MangaSearchQueryType    `form:"type,omitempty" json:"type,omitempty"`
	Score         *float32                 `form:"score,omitempty" json:"score,omitempty"`
	MinScore      *float32                 `form:"min_score,omitempty" json:"min_score,omitempty"`
	MaxScore      *float32                 `form:"max_score,omitempty" json:"max_score,omitempty"`
	Status        *MangaSearchQueryStatus  `form:"status,omitempty" json:"status,omitempty"`
	Sfw           *bool                    `form:"sfw,omitempty" json:"sfw,omitempty"`
	Genres        *string                  `form:"genres,omitempty" json:"genres,omitempty"`
	GenresExclude *string                  `form:"genres_exclude,omitempty" json:"genres_exclude,omitempty"`
	OrderBy       *MangaSearchQueryOrderby `form:"order_by,omitempty" json:"order_by,omitempty"`
	Sort          *SearchQuerySort         `form:"sort,omitempty" json:"sort,omitempty"`
	Letter        *string                  `form:"letter,omitempty" json:"letter,omitempty"`
	Magazines     *string                  `form:"magazines,omitempty" json:"magazines,omitempty"`
	StartDate     *string                  `form:"start_date,omitempty" json:"start_date,omitempty"`
	EndDate       *string                  `form:"end_date,omitempty" json:"end_date,omitempty"`
}
type GetMangaTopicsParams struct {
	Filter *GetMangaTopicsParamsFilter `form:"filter,omitempty" json:"filter,omitempty"`
}
type GetMangaTopicsParamsFilter string
type GetMangaNewsParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetMangaReviewsParams struct {
	Page        *Page        `form:"page,omitempty" json:"page,omitempty"`
	Preliminary *Preliminary `form:"preliminary,omitempty" json:"preliminary,omitempty"`
	Spoilers    *Spoilers    `form:"spoilers,omitempty" json:"spoilers,omitempty"`
}
type GetMangaUserUpdatesParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetPeopleSearchParams struct {
	Page    *Page                     `form:"page,omitempty" json:"page,omitempty"`
	Limit   *Limit                    `form:"limit,omitempty" json:"limit,omitempty"`
	Q       *string                   `form:"q,omitempty" json:"q,omitempty"`
	OrderBy *PeopleSearchQueryOrderby `form:"order_by,omitempty" json:"order_by,omitempty"`
	Sort    *SearchQuerySort          `form:"sort,omitempty" json:"sort,omitempty"`
	Letter  *string                   `form:"letter,omitempty" json:"letter,omitempty"`
}
type GetProducersParams struct {
	Page    *Page                  `form:"page,omitempty" json:"page,omitempty"`
	Limit   *Limit                 `form:"limit,omitempty" json:"limit,omitempty"`
	Q       *string                `form:"q,omitempty" json:"q,omitempty"`
	OrderBy *ProducersQueryOrderby `form:"order_by,omitempty" json:"order_by,omitempty"`
	Sort    *SearchQuerySort       `form:"sort,omitempty" json:"sort,omitempty"`
	Letter  *string                `form:"letter,omitempty" json:"letter,omitempty"`
}
type GetRecentAnimeRecommendationsParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetRecentMangaRecommendationsParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetRecentAnimeReviewsParams struct {
	Page        *Page        `form:"page,omitempty" json:"page,omitempty"`
	Preliminary *Preliminary `form:"preliminary,omitempty" json:"preliminary,omitempty"`
	Spoilers    *Spoilers    `form:"spoilers,omitempty" json:"spoilers,omitempty"`
}
type GetRecentMangaReviewsParams struct {
	Page        *Page        `form:"page,omitempty" json:"page,omitempty"`
	Preliminary *Preliminary `form:"preliminary,omitempty" json:"preliminary,omitempty"`
	Spoilers    *Spoilers    `form:"spoilers,omitempty" json:"spoilers,omitempty"`
}
type GetSchedulesParams struct {
	Filter     *GetSchedulesParamsFilter `form:"filter,omitempty" json:"filter,omitempty"`
	Kids       *GetSchedulesParamsKids   `form:"kids,omitempty" json:"kids,omitempty"`
	Sfw        *GetSchedulesParamsSfw    `form:"sfw,omitempty" json:"sfw,omitempty"`
	Unapproved *Unapproved               `form:"unapproved,omitempty" json:"unapproved,omitempty"`
	Page       *Page                     `form:"page,omitempty" json:"page,omitempty"`
	Limit      *Limit                    `form:"limit,omitempty" json:"limit,omitempty"`
}
type GetSchedulesParamsFilter string
type GetSchedulesParamsKids string
type GetSchedulesParamsSfw string
type GetSeasonNowParams struct {
	Filter     *GetSeasonNowParamsFilter `form:"filter,omitempty" json:"filter,omitempty"`
	Sfw        *Sfw                      `form:"sfw,omitempty" json:"sfw,omitempty"`
	Unapproved *Unapproved               `form:"unapproved,omitempty" json:"unapproved,omitempty"`
	Continuing *Continuing               `form:"continuing,omitempty" json:"continuing,omitempty"`
	Page       *Page                     `form:"page,omitempty" json:"page,omitempty"`
	Limit      *Limit                    `form:"limit,omitempty" json:"limit,omitempty"`
}
type GetSeasonNowParamsFilter string
type GetSeasonUpcomingParams struct {
	Filter     *GetSeasonUpcomingParamsFilter `form:"filter,omitempty" json:"filter,omitempty"`
	Sfw        *Sfw                           `form:"sfw,omitempty" json:"sfw,omitempty"`
	Unapproved *Unapproved                    `form:"unapproved,omitempty" json:"unapproved,omitempty"`
	Continuing *Continuing                    `form:"continuing,omitempty" json:"continuing,omitempty"`
	Page       *Page                          `form:"page,omitempty" json:"page,omitempty"`
	Limit      *Limit                         `form:"limit,omitempty" json:"limit,omitempty"`
}
type GetSeasonUpcomingParamsFilter string
type GetSeasonParams struct {
	Filter     *GetSeasonParamsFilter `form:"filter,omitempty" json:"filter,omitempty"`
	Sfw        *Sfw                   `form:"sfw,omitempty" json:"sfw,omitempty"`
	Unapproved *Unapproved            `form:"unapproved,omitempty" json:"unapproved,omitempty"`
	Continuing *Continuing            `form:"continuing,omitempty" json:"continuing,omitempty"`
	Page       *Page                  `form:"page,omitempty" json:"page,omitempty"`
	Limit      *Limit                 `form:"limit,omitempty" json:"limit,omitempty"`
}
type GetSeasonParamsFilter string
type GetTopAnimeParams struct {
	Type   *AnimeSearchQueryType   `form:"type,omitempty" json:"type,omitempty"`
	Filter *TopAnimeFilter         `form:"filter,omitempty" json:"filter,omitempty"`
	Rating *AnimeSearchQueryRating `form:"rating,omitempty" json:"rating,omitempty"`
	Sfw    *bool                   `form:"sfw,omitempty" json:"sfw,omitempty"`
	Page   *Page                   `form:"page,omitempty" json:"page,omitempty"`
	Limit  *Limit                  `form:"limit,omitempty" json:"limit,omitempty"`
}
type GetTopCharactersParams struct {
	Page  *Page  `form:"page,omitempty" json:"page,omitempty"`
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`
}
type GetTopMangaParams struct {
	Type   *MangaSearchQueryType `form:"type,omitempty" json:"type,omitempty"`
	Filter *TopMangaFilter       `form:"filter,omitempty" json:"filter,omitempty"`
	Page   *Page                 `form:"page,omitempty" json:"page,omitempty"`
	Limit  *Limit                `form:"limit,omitempty" json:"limit,omitempty"`
}
type GetTopPeopleParams struct {
	Page  *Page  `form:"page,omitempty" json:"page,omitempty"`
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`
}
type GetTopReviewsParams struct {
	Page        *Page               `form:"page,omitempty" json:"page,omitempty"`
	Type        *TopReviewsTypeEnum `form:"type,omitempty" json:"type,omitempty"`
	Preliminary *bool               `form:"preliminary,omitempty" json:"preliminary,omitempty"`
	Spoilers    *bool               `form:"spoilers,omitempty" json:"spoilers,omitempty"`
}
type GetUsersSearchParams struct {
	Page     *Page                   `form:"page,omitempty" json:"page,omitempty"`
	Limit    *Limit                  `form:"limit,omitempty" json:"limit,omitempty"`
	Q        *string                 `form:"q,omitempty" json:"q,omitempty"`
	Gender   *UsersSearchQueryGender `form:"gender,omitempty" json:"gender,omitempty"`
	Location *string                 `form:"location,omitempty" json:"location,omitempty"`
	MaxAge   *int                    `form:"maxAge,omitempty" json:"maxAge,omitempty"`
	MinAge   *int                    `form:"minAge,omitempty" json:"minAge,omitempty"`
}
type GetUserAnimelistParams struct {
	Status *UserAnimeListStatusFilter `form:"status,omitempty" json:"status,omitempty"`
}
type GetUserClubsParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetUserFriendsParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetUserHistoryParams struct {
	Type *GetUserHistoryParamsType `form:"type,omitempty" json:"type,omitempty"`
}
type GetUserHistoryParamsType string
type GetUserMangaListParams struct {
	Status *UserMangaListStatusFilter `form:"status,omitempty" json:"status,omitempty"`
}
type GetUserRecommendationsParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetUserReviewsParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}
type GetWatchRecentPromosParams struct {
	Page *Page `form:"page,omitempty" json:"page,omitempty"`
}

func (t EntryRecommendations_Data_Entry) AsAnimeMeta() (AnimeMeta, error) {
	var body AnimeMeta
	err := json.Unmarshal(t.union, &body)
	return body, err
}
func (t *EntryRecommendations_Data_Entry) FromAnimeMeta(v AnimeMeta) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}
func (t *EntryRecommendations_Data_Entry) MergeAnimeMeta(v AnimeMeta) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}
func (t EntryRecommendations_Data_Entry) AsMangaMeta() (MangaMeta, error) {
	var body MangaMeta
	err := json.Unmarshal(t.union, &body)
	return body, err
}
func (t *EntryRecommendations_Data_Entry) FromMangaMeta(v MangaMeta) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}
func (t *EntryRecommendations_Data_Entry) MergeMangaMeta(v MangaMeta) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}
func (t EntryRecommendations_Data_Entry) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}
func (t *EntryRecommendations_Data_Entry) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
func (t Random_Data_Item) AsAnime() (Anime, error) {
	var body Anime
	err := json.Unmarshal(t.union, &body)
	return body, err
}
func (t *Random_Data_Item) FromAnime(v Anime) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}
func (t *Random_Data_Item) MergeAnime(v Anime) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}
func (t Random_Data_Item) AsManga() (Manga, error) {
	var body Manga
	err := json.Unmarshal(t.union, &body)
	return body, err
}
func (t *Random_Data_Item) FromManga(v Manga) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}
func (t *Random_Data_Item) MergeManga(v Manga) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}
func (t Random_Data_Item) AsCharacter() (Character, error) {
	var body Character
	err := json.Unmarshal(t.union, &body)
	return body, err
}
func (t *Random_Data_Item) FromCharacter(v Character) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}
func (t *Random_Data_Item) MergeCharacter(v Character) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}
func (t Random_Data_Item) AsPerson() (Person, error) {
	var body Person
	err := json.Unmarshal(t.union, &body)
	return body, err
}
func (t *Random_Data_Item) FromPerson(v Person) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}
func (t *Random_Data_Item) MergePerson(v Person) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}
func (t Random_Data_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}
func (t *Random_Data_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
func (t Recommendations_Data_Entry_Item) AsAnimeMeta() (AnimeMeta, error) {
	var body AnimeMeta
	err := json.Unmarshal(t.union, &body)
	return body, err
}
func (t *Recommendations_Data_Entry_Item) FromAnimeMeta(v AnimeMeta) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}
func (t *Recommendations_Data_Entry_Item) MergeAnimeMeta(v AnimeMeta) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}
func (t Recommendations_Data_Entry_Item) AsMangaMeta() (MangaMeta, error) {
	var body MangaMeta
	err := json.Unmarshal(t.union, &body)
	return body, err
}
func (t *Recommendations_Data_Entry_Item) FromMangaMeta(v MangaMeta) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}
func (t *Recommendations_Data_Entry_Item) MergeMangaMeta(v MangaMeta) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}
func (t Recommendations_Data_Entry_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}
func (t *Recommendations_Data_Entry_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
func (t ReviewsCollection_Data_Item) AsAnimeReview() (AnimeReview, error) {
	var body AnimeReview
	err := json.Unmarshal(t.union, &body)
	return body, err
}
func (t *ReviewsCollection_Data_Item) FromAnimeReview(v AnimeReview) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}
func (t *ReviewsCollection_Data_Item) MergeAnimeReview(v AnimeReview) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}
func (t ReviewsCollection_Data_Item) AsMangaReview() (MangaReview, error) {
	var body MangaReview
	err := json.Unmarshal(t.union, &body)
	return body, err
}
func (t *ReviewsCollection_Data_Item) FromMangaReview(v MangaReview) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}
func (t *ReviewsCollection_Data_Item) MergeMangaReview(v MangaReview) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}
func (t ReviewsCollection_Data_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}
func (t *ReviewsCollection_Data_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

type RequestEditorFn func(ctx context.Context, req *http.Request) error
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
type Client struct {
	Server         string
	Client         HttpRequestDoer
	RequestEditors []RequestEditorFn
}
type ClientOption func(*Client) error

func NewClient(server string, opts ...ClientOption) (*Client, error) {
	client := Client{
		Server: server,
	}
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

type ClientInterface interface {
	GetAnimeSearch(ctx context.Context, params *GetAnimeSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeCharacters(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeEpisodes(ctx context.Context, id int, params *GetAnimeEpisodesParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeEpisodeById(ctx context.Context, id int, episode int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeExternal(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeForum(ctx context.Context, id int, params *GetAnimeForumParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeFullById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeMoreInfo(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeNews(ctx context.Context, id int, params *GetAnimeNewsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimePictures(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeRecommendations(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeRelations(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeReviews(ctx context.Context, id int, params *GetAnimeReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeStaff(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeStatistics(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeStreaming(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeThemes(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeUserUpdates(ctx context.Context, id int, params *GetAnimeUserUpdatesParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeVideos(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeVideosEpisodes(ctx context.Context, id int, params *GetAnimeVideosEpisodesParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetCharactersSearch(ctx context.Context, params *GetCharactersSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetCharacterById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetCharacterAnime(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetCharacterFullById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetCharacterManga(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetCharacterPictures(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetCharacterVoiceActors(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetClubsSearch(ctx context.Context, params *GetClubsSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetClubsById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetClubMembers(ctx context.Context, id int, params *GetClubMembersParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetClubRelations(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetClubStaff(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetAnimeGenres(ctx context.Context, params *GetAnimeGenresParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaGenres(ctx context.Context, params *GetMangaGenresParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMagazines(ctx context.Context, params *GetMagazinesParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaSearch(ctx context.Context, params *GetMangaSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaCharacters(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaExternal(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaTopics(ctx context.Context, id int, params *GetMangaTopicsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaFullById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaMoreInfo(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaNews(ctx context.Context, id int, params *GetMangaNewsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaPictures(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaRecommendations(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaRelations(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaReviews(ctx context.Context, id int, params *GetMangaReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaStatistics(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetMangaUserUpdates(ctx context.Context, id int, params *GetMangaUserUpdatesParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetPeopleSearch(ctx context.Context, params *GetPeopleSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetPersonById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetPersonAnime(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetPersonFullById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetPersonManga(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetPersonPictures(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetPersonVoices(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetProducers(ctx context.Context, params *GetProducersParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetProducerById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetProducerExternal(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetProducerFullById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetRandomAnime(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetRandomCharacters(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetRandomManga(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetRandomPeople(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetRandomUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetRecentAnimeRecommendations(ctx context.Context, params *GetRecentAnimeRecommendationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetRecentMangaRecommendations(ctx context.Context, params *GetRecentMangaRecommendationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetRecentAnimeReviews(ctx context.Context, params *GetRecentAnimeReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetRecentMangaReviews(ctx context.Context, params *GetRecentMangaReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetSchedules(ctx context.Context, params *GetSchedulesParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetSeasonsList(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetSeasonNow(ctx context.Context, params *GetSeasonNowParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetSeasonUpcoming(ctx context.Context, params *GetSeasonUpcomingParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetSeason(ctx context.Context, year int, season string, params *GetSeasonParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetTopAnime(ctx context.Context, params *GetTopAnimeParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetTopCharacters(ctx context.Context, params *GetTopCharactersParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetTopManga(ctx context.Context, params *GetTopMangaParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetTopPeople(ctx context.Context, params *GetTopPeopleParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetTopReviews(ctx context.Context, params *GetTopReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUsersSearch(ctx context.Context, params *GetUsersSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserProfile(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserAbout(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserAnimelist(ctx context.Context, username string, params *GetUserAnimelistParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserClubs(ctx context.Context, username string, params *GetUserClubsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserExternal(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserFavorites(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserFriends(ctx context.Context, username string, params *GetUserFriendsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserFullProfile(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserHistory(ctx context.Context, username string, params *GetUserHistoryParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserMangaList(ctx context.Context, username string, params *GetUserMangaListParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserRecommendations(ctx context.Context, username string, params *GetUserRecommendationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserReviews(ctx context.Context, username string, params *GetUserReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserStatistics(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetUserUpdates(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetWatchRecentEpisodes(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetWatchPopularEpisodes(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetWatchRecentPromos(ctx context.Context, params *GetWatchRecentPromosParams, reqEditors ...RequestEditorFn) (*http.Response, error)
	GetWatchPopularPromos(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAnimeSearch(ctx context.Context, params *GetAnimeSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeSearchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeCharacters(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeCharactersRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeEpisodes(ctx context.Context, id int, params *GetAnimeEpisodesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeEpisodesRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeEpisodeById(ctx context.Context, id int, episode int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeEpisodeByIdRequest(c.Server, id, episode)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeExternal(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeExternalRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeForum(ctx context.Context, id int, params *GetAnimeForumParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeForumRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeFullById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeFullByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeMoreInfo(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeMoreInfoRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeNews(ctx context.Context, id int, params *GetAnimeNewsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeNewsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimePictures(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimePicturesRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeRecommendations(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeRecommendationsRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeRelations(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeRelationsRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeReviews(ctx context.Context, id int, params *GetAnimeReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeReviewsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeStaff(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeStaffRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeStatistics(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeStatisticsRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeStreaming(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeStreamingRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeThemes(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeThemesRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeUserUpdates(ctx context.Context, id int, params *GetAnimeUserUpdatesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeUserUpdatesRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeVideos(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeVideosRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeVideosEpisodes(ctx context.Context, id int, params *GetAnimeVideosEpisodesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeVideosEpisodesRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetCharactersSearch(ctx context.Context, params *GetCharactersSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetCharactersSearchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetCharacterById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetCharacterByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetCharacterAnime(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetCharacterAnimeRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetCharacterFullById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetCharacterFullByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetCharacterManga(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetCharacterMangaRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetCharacterPictures(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetCharacterPicturesRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetCharacterVoiceActors(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetCharacterVoiceActorsRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetClubsSearch(ctx context.Context, params *GetClubsSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetClubsSearchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetClubsById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetClubsByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetClubMembers(ctx context.Context, id int, params *GetClubMembersParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetClubMembersRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetClubRelations(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetClubRelationsRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetClubStaff(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetClubStaffRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetAnimeGenres(ctx context.Context, params *GetAnimeGenresParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnimeGenresRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaGenres(ctx context.Context, params *GetMangaGenresParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaGenresRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMagazines(ctx context.Context, params *GetMagazinesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMagazinesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaSearch(ctx context.Context, params *GetMangaSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaSearchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaCharacters(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaCharactersRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaExternal(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaExternalRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaTopics(ctx context.Context, id int, params *GetMangaTopicsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaTopicsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaFullById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaFullByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaMoreInfo(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaMoreInfoRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaNews(ctx context.Context, id int, params *GetMangaNewsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaNewsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaPictures(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaPicturesRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaRecommendations(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaRecommendationsRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaRelations(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaRelationsRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaReviews(ctx context.Context, id int, params *GetMangaReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaReviewsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaStatistics(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaStatisticsRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetMangaUserUpdates(ctx context.Context, id int, params *GetMangaUserUpdatesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMangaUserUpdatesRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetPeopleSearch(ctx context.Context, params *GetPeopleSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPeopleSearchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetPersonById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPersonByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetPersonAnime(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPersonAnimeRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetPersonFullById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPersonFullByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetPersonManga(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPersonMangaRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetPersonPictures(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPersonPicturesRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetPersonVoices(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPersonVoicesRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetProducers(ctx context.Context, params *GetProducersParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProducersRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetProducerById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProducerByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetProducerExternal(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProducerExternalRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetProducerFullById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProducerFullByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetRandomAnime(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRandomAnimeRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetRandomCharacters(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRandomCharactersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetRandomManga(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRandomMangaRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetRandomPeople(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRandomPeopleRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetRandomUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRandomUsersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetRecentAnimeRecommendations(ctx context.Context, params *GetRecentAnimeRecommendationsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRecentAnimeRecommendationsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetRecentMangaRecommendations(ctx context.Context, params *GetRecentMangaRecommendationsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRecentMangaRecommendationsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetRecentAnimeReviews(ctx context.Context, params *GetRecentAnimeReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRecentAnimeReviewsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetRecentMangaReviews(ctx context.Context, params *GetRecentMangaReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRecentMangaReviewsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetSchedules(ctx context.Context, params *GetSchedulesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSchedulesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetSeasonsList(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSeasonsListRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetSeasonNow(ctx context.Context, params *GetSeasonNowParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSeasonNowRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetSeasonUpcoming(ctx context.Context, params *GetSeasonUpcomingParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSeasonUpcomingRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetSeason(ctx context.Context, year int, season string, params *GetSeasonParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSeasonRequest(c.Server, year, season, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetTopAnime(ctx context.Context, params *GetTopAnimeParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTopAnimeRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetTopCharacters(ctx context.Context, params *GetTopCharactersParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTopCharactersRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetTopManga(ctx context.Context, params *GetTopMangaParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTopMangaRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetTopPeople(ctx context.Context, params *GetTopPeopleParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTopPeopleRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetTopReviews(ctx context.Context, params *GetTopReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTopReviewsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUsersSearch(ctx context.Context, params *GetUsersSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsersSearchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserById(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserProfile(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserProfileRequest(c.Server, username)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserAbout(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserAboutRequest(c.Server, username)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserAnimelist(ctx context.Context, username string, params *GetUserAnimelistParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserAnimelistRequest(c.Server, username, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserClubs(ctx context.Context, username string, params *GetUserClubsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserClubsRequest(c.Server, username, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserExternal(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserExternalRequest(c.Server, username)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserFavorites(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserFavoritesRequest(c.Server, username)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserFriends(ctx context.Context, username string, params *GetUserFriendsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserFriendsRequest(c.Server, username, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserFullProfile(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserFullProfileRequest(c.Server, username)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserHistory(ctx context.Context, username string, params *GetUserHistoryParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserHistoryRequest(c.Server, username, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserMangaList(ctx context.Context, username string, params *GetUserMangaListParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserMangaListRequest(c.Server, username, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserRecommendations(ctx context.Context, username string, params *GetUserRecommendationsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserRecommendationsRequest(c.Server, username, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserReviews(ctx context.Context, username string, params *GetUserReviewsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserReviewsRequest(c.Server, username, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserStatistics(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserStatisticsRequest(c.Server, username)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetUserUpdates(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserUpdatesRequest(c.Server, username)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetWatchRecentEpisodes(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetWatchRecentEpisodesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetWatchPopularEpisodes(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetWatchPopularEpisodesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetWatchRecentPromos(ctx context.Context, params *GetWatchRecentPromosParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetWatchRecentPromosRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetWatchPopularPromos(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetWatchPopularPromosRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetAnimeSearchRequest(server string, params *GetAnimeSearchParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Unapproved != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "unapproved", runtime.ParamLocationQuery, *params.Unapproved); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Q != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "q", runtime.ParamLocationQuery, *params.Q); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Type != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Score != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "score", runtime.ParamLocationQuery, *params.Score); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.MinScore != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "min_score", runtime.ParamLocationQuery, *params.MinScore); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.MaxScore != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "max_score", runtime.ParamLocationQuery, *params.MaxScore); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Status != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "status", runtime.ParamLocationQuery, *params.Status); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Rating != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "rating", runtime.ParamLocationQuery, *params.Rating); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sfw != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sfw", runtime.ParamLocationQuery, *params.Sfw); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Genres != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "genres", runtime.ParamLocationQuery, *params.Genres); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.GenresExclude != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "genres_exclude", runtime.ParamLocationQuery, *params.GenresExclude); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.OrderBy != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "order_by", runtime.ParamLocationQuery, *params.OrderBy); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sort != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, *params.Sort); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Letter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "letter", runtime.ParamLocationQuery, *params.Letter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Producers != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "producers", runtime.ParamLocationQuery, *params.Producers); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.StartDate != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "start_date", runtime.ParamLocationQuery, *params.StartDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.EndDate != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "end_date", runtime.ParamLocationQuery, *params.EndDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeCharactersRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/characters", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeEpisodesRequest(server string, id int, params *GetAnimeEpisodesParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/episodes", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeEpisodeByIdRequest(server string, id int, episode int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	var pathParam1 string
	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "episode", runtime.ParamLocationPath, episode)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/episodes/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeExternalRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/external", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeForumRequest(server string, id int, params *GetAnimeForumParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/forum", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Filter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter", runtime.ParamLocationQuery, *params.Filter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeFullByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/full", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeMoreInfoRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/moreinfo", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeNewsRequest(server string, id int, params *GetAnimeNewsParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/news", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimePicturesRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/pictures", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeRecommendationsRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/recommendations", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeRelationsRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/relations", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeReviewsRequest(server string, id int, params *GetAnimeReviewsParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/reviews", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Preliminary != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "preliminary", runtime.ParamLocationQuery, *params.Preliminary); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Spoilers != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "spoilers", runtime.ParamLocationQuery, *params.Spoilers); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeStaffRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/staff", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeStatisticsRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/statistics", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeStreamingRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/streaming", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeThemesRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/themes", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeUserUpdatesRequest(server string, id int, params *GetAnimeUserUpdatesParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/userupdates", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeVideosRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/videos", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeVideosEpisodesRequest(server string, id int, params *GetAnimeVideosEpisodesParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/anime/%s/videos/episodes", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetCharactersSearchRequest(server string, params *GetCharactersSearchParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/characters")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Q != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "q", runtime.ParamLocationQuery, *params.Q); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.OrderBy != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "order_by", runtime.ParamLocationQuery, *params.OrderBy); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sort != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, *params.Sort); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Letter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "letter", runtime.ParamLocationQuery, *params.Letter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetCharacterByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/characters/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetCharacterAnimeRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/characters/%s/anime", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetCharacterFullByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/characters/%s/full", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetCharacterMangaRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/characters/%s/manga", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetCharacterPicturesRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/characters/%s/pictures", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetCharacterVoiceActorsRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/characters/%s/voices", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetClubsSearchRequest(server string, params *GetClubsSearchParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/clubs")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Q != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "q", runtime.ParamLocationQuery, *params.Q); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Type != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Category != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "category", runtime.ParamLocationQuery, *params.Category); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.OrderBy != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "order_by", runtime.ParamLocationQuery, *params.OrderBy); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sort != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, *params.Sort); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Letter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "letter", runtime.ParamLocationQuery, *params.Letter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetClubsByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/clubs/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetClubMembersRequest(server string, id int, params *GetClubMembersParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/clubs/%s/members", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetClubRelationsRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/clubs/%s/relations", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetClubStaffRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/clubs/%s/staff", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetAnimeGenresRequest(server string, params *GetAnimeGenresParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/genres/anime")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Filter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter", runtime.ParamLocationQuery, *params.Filter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaGenresRequest(server string, params *GetMangaGenresParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/genres/manga")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Filter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter", runtime.ParamLocationQuery, *params.Filter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMagazinesRequest(server string, params *GetMagazinesParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/magazines")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Q != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "q", runtime.ParamLocationQuery, *params.Q); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.OrderBy != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "order_by", runtime.ParamLocationQuery, *params.OrderBy); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sort != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, *params.Sort); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Letter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "letter", runtime.ParamLocationQuery, *params.Letter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaSearchRequest(server string, params *GetMangaSearchParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Unapproved != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "unapproved", runtime.ParamLocationQuery, *params.Unapproved); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Q != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "q", runtime.ParamLocationQuery, *params.Q); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Type != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Score != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "score", runtime.ParamLocationQuery, *params.Score); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.MinScore != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "min_score", runtime.ParamLocationQuery, *params.MinScore); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.MaxScore != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "max_score", runtime.ParamLocationQuery, *params.MaxScore); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Status != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "status", runtime.ParamLocationQuery, *params.Status); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sfw != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sfw", runtime.ParamLocationQuery, *params.Sfw); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Genres != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "genres", runtime.ParamLocationQuery, *params.Genres); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.GenresExclude != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "genres_exclude", runtime.ParamLocationQuery, *params.GenresExclude); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.OrderBy != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "order_by", runtime.ParamLocationQuery, *params.OrderBy); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sort != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, *params.Sort); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Letter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "letter", runtime.ParamLocationQuery, *params.Letter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Magazines != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "magazines", runtime.ParamLocationQuery, *params.Magazines); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.StartDate != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "start_date", runtime.ParamLocationQuery, *params.StartDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.EndDate != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "end_date", runtime.ParamLocationQuery, *params.EndDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaCharactersRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/characters", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaExternalRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/external", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaTopicsRequest(server string, id int, params *GetMangaTopicsParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/forum", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Filter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter", runtime.ParamLocationQuery, *params.Filter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaFullByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/full", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaMoreInfoRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/moreinfo", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaNewsRequest(server string, id int, params *GetMangaNewsParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/news", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaPicturesRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/pictures", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaRecommendationsRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/recommendations", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaRelationsRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/relations", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaReviewsRequest(server string, id int, params *GetMangaReviewsParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/reviews", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Preliminary != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "preliminary", runtime.ParamLocationQuery, *params.Preliminary); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Spoilers != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "spoilers", runtime.ParamLocationQuery, *params.Spoilers); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaStatisticsRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/statistics", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetMangaUserUpdatesRequest(server string, id int, params *GetMangaUserUpdatesParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/manga/%s/userupdates", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetPeopleSearchRequest(server string, params *GetPeopleSearchParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/people")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Q != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "q", runtime.ParamLocationQuery, *params.Q); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.OrderBy != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "order_by", runtime.ParamLocationQuery, *params.OrderBy); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sort != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, *params.Sort); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Letter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "letter", runtime.ParamLocationQuery, *params.Letter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetPersonByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/people/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetPersonAnimeRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/people/%s/anime", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetPersonFullByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/people/%s/full", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetPersonMangaRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/people/%s/manga", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetPersonPicturesRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/people/%s/pictures", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetPersonVoicesRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/people/%s/voices", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetProducersRequest(server string, params *GetProducersParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/producers")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Q != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "q", runtime.ParamLocationQuery, *params.Q); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.OrderBy != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "order_by", runtime.ParamLocationQuery, *params.OrderBy); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sort != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, *params.Sort); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Letter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "letter", runtime.ParamLocationQuery, *params.Letter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetProducerByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/producers/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetProducerExternalRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/producers/%s/external", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetProducerFullByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/producers/%s/full", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetRandomAnimeRequest(server string) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/random/anime")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetRandomCharactersRequest(server string) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/random/characters")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetRandomMangaRequest(server string) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/random/manga")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetRandomPeopleRequest(server string) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/random/people")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetRandomUsersRequest(server string) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/random/users")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetRecentAnimeRecommendationsRequest(server string, params *GetRecentAnimeRecommendationsParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/recommendations/anime")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetRecentMangaRecommendationsRequest(server string, params *GetRecentMangaRecommendationsParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/recommendations/manga")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetRecentAnimeReviewsRequest(server string, params *GetRecentAnimeReviewsParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/reviews/anime")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Preliminary != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "preliminary", runtime.ParamLocationQuery, *params.Preliminary); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Spoilers != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "spoilers", runtime.ParamLocationQuery, *params.Spoilers); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetRecentMangaReviewsRequest(server string, params *GetRecentMangaReviewsParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/reviews/manga")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Preliminary != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "preliminary", runtime.ParamLocationQuery, *params.Preliminary); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Spoilers != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "spoilers", runtime.ParamLocationQuery, *params.Spoilers); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetSchedulesRequest(server string, params *GetSchedulesParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/schedules")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Filter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter", runtime.ParamLocationQuery, *params.Filter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Kids != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "kids", runtime.ParamLocationQuery, *params.Kids); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sfw != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sfw", runtime.ParamLocationQuery, *params.Sfw); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Unapproved != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "unapproved", runtime.ParamLocationQuery, *params.Unapproved); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetSeasonsListRequest(server string) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/seasons")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetSeasonNowRequest(server string, params *GetSeasonNowParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/seasons/now")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Filter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter", runtime.ParamLocationQuery, *params.Filter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sfw != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sfw", runtime.ParamLocationQuery, *params.Sfw); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Unapproved != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "unapproved", runtime.ParamLocationQuery, *params.Unapproved); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Continuing != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "continuing", runtime.ParamLocationQuery, *params.Continuing); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetSeasonUpcomingRequest(server string, params *GetSeasonUpcomingParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/seasons/upcoming")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Filter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter", runtime.ParamLocationQuery, *params.Filter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sfw != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sfw", runtime.ParamLocationQuery, *params.Sfw); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Unapproved != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "unapproved", runtime.ParamLocationQuery, *params.Unapproved); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Continuing != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "continuing", runtime.ParamLocationQuery, *params.Continuing); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetSeasonRequest(server string, year int, season string, params *GetSeasonParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "year", runtime.ParamLocationPath, year)
	if err != nil {
		return nil, err
	}
	var pathParam1 string
	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "season", runtime.ParamLocationPath, season)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/seasons/%s/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Filter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter", runtime.ParamLocationQuery, *params.Filter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sfw != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sfw", runtime.ParamLocationQuery, *params.Sfw); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Unapproved != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "unapproved", runtime.ParamLocationQuery, *params.Unapproved); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Continuing != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "continuing", runtime.ParamLocationQuery, *params.Continuing); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetTopAnimeRequest(server string, params *GetTopAnimeParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/top/anime")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Type != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Filter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter", runtime.ParamLocationQuery, *params.Filter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Rating != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "rating", runtime.ParamLocationQuery, *params.Rating); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Sfw != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sfw", runtime.ParamLocationQuery, *params.Sfw); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetTopCharactersRequest(server string, params *GetTopCharactersParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/top/characters")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetTopMangaRequest(server string, params *GetTopMangaParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/top/manga")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Type != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Filter != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "filter", runtime.ParamLocationQuery, *params.Filter); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetTopPeopleRequest(server string, params *GetTopPeopleParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/top/people")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetTopReviewsRequest(server string, params *GetTopReviewsParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/top/reviews")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Type != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Preliminary != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "preliminary", runtime.ParamLocationQuery, *params.Preliminary); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Spoilers != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "spoilers", runtime.ParamLocationQuery, *params.Spoilers); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUsersSearchRequest(server string, params *GetUsersSearchParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Limit != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Q != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "q", runtime.ParamLocationQuery, *params.Q); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Gender != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "gender", runtime.ParamLocationQuery, *params.Gender); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.Location != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "location", runtime.ParamLocationQuery, *params.Location); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.MaxAge != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxAge", runtime.ParamLocationQuery, *params.MaxAge); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		if params.MinAge != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "minAge", runtime.ParamLocationQuery, *params.MinAge); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserByIdRequest(server string, id int) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/userbyid/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserProfileRequest(server string, username string) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserAboutRequest(server string, username string) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/about", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserAnimelistRequest(server string, username string, params *GetUserAnimelistParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/animelist", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Status != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "status", runtime.ParamLocationQuery, *params.Status); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserClubsRequest(server string, username string, params *GetUserClubsParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/clubs", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserExternalRequest(server string, username string) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/external", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserFavoritesRequest(server string, username string) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/favorites", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserFriendsRequest(server string, username string, params *GetUserFriendsParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/friends", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserFullProfileRequest(server string, username string) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/full", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserHistoryRequest(server string, username string, params *GetUserHistoryParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/history", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Type != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserMangaListRequest(server string, username string, params *GetUserMangaListParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/mangalist", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Status != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "status", runtime.ParamLocationQuery, *params.Status); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserRecommendationsRequest(server string, username string, params *GetUserRecommendationsParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/recommendations", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserReviewsRequest(server string, username string, params *GetUserReviewsParams) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/reviews", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserStatisticsRequest(server string, username string) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/statistics", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetUserUpdatesRequest(server string, username string) (*http.Request, error) {
	var err error
	var pathParam0 string
	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "username", runtime.ParamLocationPath, username)
	if err != nil {
		return nil, err
	}
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/users/%s/userupdates", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetWatchRecentEpisodesRequest(server string) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/watch/episodes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetWatchPopularEpisodesRequest(server string) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/watch/episodes/popular")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetWatchRecentPromosRequest(server string, params *GetWatchRecentPromosParams) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/watch/promos")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		queryValues := queryURL.Query()
		if params.Page != nil {
			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}
		}
		queryURL.RawQuery = queryValues.Encode()
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func NewGetWatchPopularPromosRequest(server string) (*http.Request, error) {
	var err error
	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	operationPath := fmt.Sprintf("/watch/promos/popular")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}
	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

type ClientWithResponses struct {
	ClientInterface
}

func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

type ClientWithResponsesInterface interface {
	GetAnimeSearchWithResponse(ctx context.Context, params *GetAnimeSearchParams, reqEditors ...RequestEditorFn) (*GetAnimeSearchResponse, error)
	GetAnimeByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeByIdResponse, error)
	GetAnimeCharactersWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeCharactersResponse, error)
	GetAnimeEpisodesWithResponse(ctx context.Context, id int, params *GetAnimeEpisodesParams, reqEditors ...RequestEditorFn) (*GetAnimeEpisodesResponse, error)
	GetAnimeEpisodeByIdWithResponse(ctx context.Context, id int, episode int, reqEditors ...RequestEditorFn) (*GetAnimeEpisodeByIdResponse, error)
	GetAnimeExternalWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeExternalResponse, error)
	GetAnimeForumWithResponse(ctx context.Context, id int, params *GetAnimeForumParams, reqEditors ...RequestEditorFn) (*GetAnimeForumResponse, error)
	GetAnimeFullByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeFullByIdResponse, error)
	GetAnimeMoreInfoWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeMoreInfoResponse, error)
	GetAnimeNewsWithResponse(ctx context.Context, id int, params *GetAnimeNewsParams, reqEditors ...RequestEditorFn) (*GetAnimeNewsResponse, error)
	GetAnimePicturesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimePicturesResponse, error)
	GetAnimeRecommendationsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeRecommendationsResponse, error)
	GetAnimeRelationsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeRelationsResponse, error)
	GetAnimeReviewsWithResponse(ctx context.Context, id int, params *GetAnimeReviewsParams, reqEditors ...RequestEditorFn) (*GetAnimeReviewsResponse, error)
	GetAnimeStaffWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeStaffResponse, error)
	GetAnimeStatisticsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeStatisticsResponse, error)
	GetAnimeStreamingWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeStreamingResponse, error)
	GetAnimeThemesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeThemesResponse, error)
	GetAnimeUserUpdatesWithResponse(ctx context.Context, id int, params *GetAnimeUserUpdatesParams, reqEditors ...RequestEditorFn) (*GetAnimeUserUpdatesResponse, error)
	GetAnimeVideosWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeVideosResponse, error)
	GetAnimeVideosEpisodesWithResponse(ctx context.Context, id int, params *GetAnimeVideosEpisodesParams, reqEditors ...RequestEditorFn) (*GetAnimeVideosEpisodesResponse, error)
	GetCharactersSearchWithResponse(ctx context.Context, params *GetCharactersSearchParams, reqEditors ...RequestEditorFn) (*GetCharactersSearchResponse, error)
	GetCharacterByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterByIdResponse, error)
	GetCharacterAnimeWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterAnimeResponse, error)
	GetCharacterFullByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterFullByIdResponse, error)
	GetCharacterMangaWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterMangaResponse, error)
	GetCharacterPicturesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterPicturesResponse, error)
	GetCharacterVoiceActorsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterVoiceActorsResponse, error)
	GetClubsSearchWithResponse(ctx context.Context, params *GetClubsSearchParams, reqEditors ...RequestEditorFn) (*GetClubsSearchResponse, error)
	GetClubsByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetClubsByIdResponse, error)
	GetClubMembersWithResponse(ctx context.Context, id int, params *GetClubMembersParams, reqEditors ...RequestEditorFn) (*GetClubMembersResponse, error)
	GetClubRelationsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetClubRelationsResponse, error)
	GetClubStaffWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetClubStaffResponse, error)
	GetAnimeGenresWithResponse(ctx context.Context, params *GetAnimeGenresParams, reqEditors ...RequestEditorFn) (*GetAnimeGenresResponse, error)
	GetMangaGenresWithResponse(ctx context.Context, params *GetMangaGenresParams, reqEditors ...RequestEditorFn) (*GetMangaGenresResponse, error)
	GetMagazinesWithResponse(ctx context.Context, params *GetMagazinesParams, reqEditors ...RequestEditorFn) (*GetMagazinesResponse, error)
	GetMangaSearchWithResponse(ctx context.Context, params *GetMangaSearchParams, reqEditors ...RequestEditorFn) (*GetMangaSearchResponse, error)
	GetMangaByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaByIdResponse, error)
	GetMangaCharactersWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaCharactersResponse, error)
	GetMangaExternalWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaExternalResponse, error)
	GetMangaTopicsWithResponse(ctx context.Context, id int, params *GetMangaTopicsParams, reqEditors ...RequestEditorFn) (*GetMangaTopicsResponse, error)
	GetMangaFullByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaFullByIdResponse, error)
	GetMangaMoreInfoWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaMoreInfoResponse, error)
	GetMangaNewsWithResponse(ctx context.Context, id int, params *GetMangaNewsParams, reqEditors ...RequestEditorFn) (*GetMangaNewsResponse, error)
	GetMangaPicturesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaPicturesResponse, error)
	GetMangaRecommendationsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaRecommendationsResponse, error)
	GetMangaRelationsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaRelationsResponse, error)
	GetMangaReviewsWithResponse(ctx context.Context, id int, params *GetMangaReviewsParams, reqEditors ...RequestEditorFn) (*GetMangaReviewsResponse, error)
	GetMangaStatisticsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaStatisticsResponse, error)
	GetMangaUserUpdatesWithResponse(ctx context.Context, id int, params *GetMangaUserUpdatesParams, reqEditors ...RequestEditorFn) (*GetMangaUserUpdatesResponse, error)
	GetPeopleSearchWithResponse(ctx context.Context, params *GetPeopleSearchParams, reqEditors ...RequestEditorFn) (*GetPeopleSearchResponse, error)
	GetPersonByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonByIdResponse, error)
	GetPersonAnimeWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonAnimeResponse, error)
	GetPersonFullByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonFullByIdResponse, error)
	GetPersonMangaWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonMangaResponse, error)
	GetPersonPicturesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonPicturesResponse, error)
	GetPersonVoicesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonVoicesResponse, error)
	GetProducersWithResponse(ctx context.Context, params *GetProducersParams, reqEditors ...RequestEditorFn) (*GetProducersResponse, error)
	GetProducerByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetProducerByIdResponse, error)
	GetProducerExternalWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetProducerExternalResponse, error)
	GetProducerFullByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetProducerFullByIdResponse, error)
	GetRandomAnimeWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRandomAnimeResponse, error)
	GetRandomCharactersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRandomCharactersResponse, error)
	GetRandomMangaWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRandomMangaResponse, error)
	GetRandomPeopleWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRandomPeopleResponse, error)
	GetRandomUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRandomUsersResponse, error)
	GetRecentAnimeRecommendationsWithResponse(ctx context.Context, params *GetRecentAnimeRecommendationsParams, reqEditors ...RequestEditorFn) (*GetRecentAnimeRecommendationsResponse, error)
	GetRecentMangaRecommendationsWithResponse(ctx context.Context, params *GetRecentMangaRecommendationsParams, reqEditors ...RequestEditorFn) (*GetRecentMangaRecommendationsResponse, error)
	GetRecentAnimeReviewsWithResponse(ctx context.Context, params *GetRecentAnimeReviewsParams, reqEditors ...RequestEditorFn) (*GetRecentAnimeReviewsResponse, error)
	GetRecentMangaReviewsWithResponse(ctx context.Context, params *GetRecentMangaReviewsParams, reqEditors ...RequestEditorFn) (*GetRecentMangaReviewsResponse, error)
	GetSchedulesWithResponse(ctx context.Context, params *GetSchedulesParams, reqEditors ...RequestEditorFn) (*GetSchedulesResponse, error)
	GetSeasonsListWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetSeasonsListResponse, error)
	GetSeasonNowWithResponse(ctx context.Context, params *GetSeasonNowParams, reqEditors ...RequestEditorFn) (*GetSeasonNowResponse, error)
	GetSeasonUpcomingWithResponse(ctx context.Context, params *GetSeasonUpcomingParams, reqEditors ...RequestEditorFn) (*GetSeasonUpcomingResponse, error)
	GetSeasonWithResponse(ctx context.Context, year int, season string, params *GetSeasonParams, reqEditors ...RequestEditorFn) (*GetSeasonResponse, error)
	GetTopAnimeWithResponse(ctx context.Context, params *GetTopAnimeParams, reqEditors ...RequestEditorFn) (*GetTopAnimeResponse, error)
	GetTopCharactersWithResponse(ctx context.Context, params *GetTopCharactersParams, reqEditors ...RequestEditorFn) (*GetTopCharactersResponse, error)
	GetTopMangaWithResponse(ctx context.Context, params *GetTopMangaParams, reqEditors ...RequestEditorFn) (*GetTopMangaResponse, error)
	GetTopPeopleWithResponse(ctx context.Context, params *GetTopPeopleParams, reqEditors ...RequestEditorFn) (*GetTopPeopleResponse, error)
	GetTopReviewsWithResponse(ctx context.Context, params *GetTopReviewsParams, reqEditors ...RequestEditorFn) (*GetTopReviewsResponse, error)
	GetUsersSearchWithResponse(ctx context.Context, params *GetUsersSearchParams, reqEditors ...RequestEditorFn) (*GetUsersSearchResponse, error)
	GetUserByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetUserByIdResponse, error)
	GetUserProfileWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserProfileResponse, error)
	GetUserAboutWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserAboutResponse, error)
	GetUserAnimelistWithResponse(ctx context.Context, username string, params *GetUserAnimelistParams, reqEditors ...RequestEditorFn) (*GetUserAnimelistResponse, error)
	GetUserClubsWithResponse(ctx context.Context, username string, params *GetUserClubsParams, reqEditors ...RequestEditorFn) (*GetUserClubsResponse, error)
	GetUserExternalWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserExternalResponse, error)
	GetUserFavoritesWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserFavoritesResponse, error)
	GetUserFriendsWithResponse(ctx context.Context, username string, params *GetUserFriendsParams, reqEditors ...RequestEditorFn) (*GetUserFriendsResponse, error)
	GetUserFullProfileWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserFullProfileResponse, error)
	GetUserHistoryWithResponse(ctx context.Context, username string, params *GetUserHistoryParams, reqEditors ...RequestEditorFn) (*GetUserHistoryResponse, error)
	GetUserMangaListWithResponse(ctx context.Context, username string, params *GetUserMangaListParams, reqEditors ...RequestEditorFn) (*GetUserMangaListResponse, error)
	GetUserRecommendationsWithResponse(ctx context.Context, username string, params *GetUserRecommendationsParams, reqEditors ...RequestEditorFn) (*GetUserRecommendationsResponse, error)
	GetUserReviewsWithResponse(ctx context.Context, username string, params *GetUserReviewsParams, reqEditors ...RequestEditorFn) (*GetUserReviewsResponse, error)
	GetUserStatisticsWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserStatisticsResponse, error)
	GetUserUpdatesWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserUpdatesResponse, error)
	GetWatchRecentEpisodesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetWatchRecentEpisodesResponse, error)
	GetWatchPopularEpisodesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetWatchPopularEpisodesResponse, error)
	GetWatchRecentPromosWithResponse(ctx context.Context, params *GetWatchRecentPromosParams, reqEditors ...RequestEditorFn) (*GetWatchRecentPromosResponse, error)
	GetWatchPopularPromosWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetWatchPopularPromosResponse, error)
}
type GetAnimeSearchResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeSearch
}

func (r GetAnimeSearchResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeSearchResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetAnimeByIdResponseJSON200
}

func (r GetAnimeByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeCharactersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeCharacters
}

func (r GetAnimeCharactersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeCharactersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeEpisodesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeEpisodes
}

func (r GetAnimeEpisodesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeEpisodesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeEpisodeByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetAnimeEpisodeByIdResponseJSON200
}

func (r GetAnimeEpisodeByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeEpisodeByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeExternalResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ExternalLinks
}

func (r GetAnimeExternalResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeExternalResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeForumResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Forum
}

func (r GetAnimeForumResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeForumResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeFullByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetAnimeFullByIdResponseJSON200
}

func (r GetAnimeFullByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeFullByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeMoreInfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Moreinfo
}

func (r GetAnimeMoreInfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeMoreInfoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeNewsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeNews
}

func (r GetAnimeNewsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeNewsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimePicturesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PicturesVariants
}

func (r GetAnimePicturesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimePicturesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeRecommendationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *EntryRecommendations
}

func (r GetAnimeRecommendationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeRecommendationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeRelationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetAnimeRelationsResponseJSON200
}

func (r GetAnimeRelationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeRelationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeReviewsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeReviews
}

func (r GetAnimeReviewsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeReviewsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeStaffResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeStaff
}

func (r GetAnimeStaffResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeStaffResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeStatisticsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeStatistics
}

func (r GetAnimeStatisticsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeStatisticsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeStreamingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ExternalLinks
}

func (r GetAnimeStreamingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeStreamingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeThemesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeThemes
}

func (r GetAnimeThemesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeThemesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeUserUpdatesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeUserupdates
}

func (r GetAnimeUserUpdatesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeUserUpdatesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeVideosResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeVideos
}

func (r GetAnimeVideosResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeVideosResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeVideosEpisodesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeVideosEpisodes
}

func (r GetAnimeVideosEpisodesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeVideosEpisodesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCharactersSearchResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CharactersSearch
}

func (r GetCharactersSearchResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetCharactersSearchResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCharacterByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetCharacterByIdResponseJSON200
}

func (r GetCharacterByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetCharacterByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCharacterAnimeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CharacterAnime
}

func (r GetCharacterAnimeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetCharacterAnimeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCharacterFullByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetCharacterFullByIdResponseJSON200
}

func (r GetCharacterFullByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetCharacterFullByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCharacterMangaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CharacterManga
}

func (r GetCharacterMangaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetCharacterMangaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCharacterPicturesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CharacterPictures
}

func (r GetCharacterPicturesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetCharacterPicturesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCharacterVoiceActorsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CharacterVoiceActors
}

func (r GetCharacterVoiceActorsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetCharacterVoiceActorsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetClubsSearchResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ClubsSearch
}

func (r GetClubsSearchResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetClubsSearchResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetClubsByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetClubsByIdResponseJSON200
}

func (r GetClubsByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetClubsByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetClubMembersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetClubMembersResponseJSON200
}

func (r GetClubMembersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetClubMembersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetClubRelationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ClubRelations
}

func (r GetClubRelationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetClubRelationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetClubStaffResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ClubStaff
}

func (r GetClubStaffResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetClubStaffResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnimeGenresResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Genres
}

func (r GetAnimeGenresResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetAnimeGenresResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaGenresResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Genres
}

func (r GetMangaGenresResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaGenresResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMagazinesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Magazines
}

func (r GetMagazinesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMagazinesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaSearchResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MangaSearch
}

func (r GetMangaSearchResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaSearchResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetMangaByIdResponseJSON200
}

func (r GetMangaByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaCharactersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MangaCharacters
}

func (r GetMangaCharactersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaCharactersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaExternalResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ExternalLinks
}

func (r GetMangaExternalResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaExternalResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaTopicsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Forum
}

func (r GetMangaTopicsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaTopicsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaFullByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetMangaFullByIdResponseJSON200
}

func (r GetMangaFullByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaFullByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaMoreInfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Moreinfo
}

func (r GetMangaMoreInfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaMoreInfoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaNewsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MangaNews
}

func (r GetMangaNewsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaNewsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaPicturesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MangaPictures
}

func (r GetMangaPicturesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaPicturesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaRecommendationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *EntryRecommendations
}

func (r GetMangaRecommendationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaRecommendationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaRelationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetMangaRelationsResponseJSON200
}

func (r GetMangaRelationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaRelationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaReviewsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MangaReviews
}

func (r GetMangaReviewsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaReviewsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaStatisticsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MangaStatistics
}

func (r GetMangaStatisticsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaStatisticsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMangaUserUpdatesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MangaUserupdates
}

func (r GetMangaUserUpdatesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetMangaUserUpdatesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPeopleSearchResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PeopleSearch
}

func (r GetPeopleSearchResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetPeopleSearchResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPersonByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetPersonByIdResponseJSON200
}

func (r GetPersonByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetPersonByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPersonAnimeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PersonAnime
}

func (r GetPersonAnimeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetPersonAnimeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPersonFullByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetPersonFullByIdResponseJSON200
}

func (r GetPersonFullByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetPersonFullByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPersonMangaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PersonManga
}

func (r GetPersonMangaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetPersonMangaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPersonPicturesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PersonPictures
}

func (r GetPersonPicturesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetPersonPicturesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPersonVoicesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PersonVoiceActingRoles
}

func (r GetPersonVoicesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetPersonVoicesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetProducersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Producers
}

func (r GetProducersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetProducersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetProducerByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetProducerByIdResponseJSON200
}

func (r GetProducerByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetProducerByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetProducerExternalResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ExternalLinks
}

func (r GetProducerExternalResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetProducerExternalResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetProducerFullByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetProducerFullByIdResponseJSON200
}

func (r GetProducerFullByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetProducerFullByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRandomAnimeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetRandomAnimeResponseJSON200
}

func (r GetRandomAnimeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetRandomAnimeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRandomCharactersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetRandomCharactersResponseJSON200
}

func (r GetRandomCharactersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetRandomCharactersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRandomMangaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetRandomMangaResponseJSON200
}

func (r GetRandomMangaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetRandomMangaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRandomPeopleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetRandomPeopleResponseJSON200
}

func (r GetRandomPeopleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetRandomPeopleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRandomUsersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetRandomUsersResponseJSON200
}

func (r GetRandomUsersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetRandomUsersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRecentAnimeRecommendationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Recommendations
}

func (r GetRecentAnimeRecommendationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetRecentAnimeRecommendationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRecentMangaRecommendationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Recommendations
}

func (r GetRecentMangaRecommendationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetRecentMangaRecommendationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRecentAnimeReviewsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
}

func (r GetRecentAnimeReviewsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetRecentAnimeReviewsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRecentMangaReviewsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
}

func (r GetRecentMangaReviewsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetRecentMangaReviewsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSchedulesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Schedules
}

func (r GetSchedulesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetSchedulesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSeasonsListResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Seasons
}

func (r GetSeasonsListResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetSeasonsListResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSeasonNowResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeSearch
}

func (r GetSeasonNowResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetSeasonNowResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSeasonUpcomingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeSearch
}

func (r GetSeasonUpcomingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetSeasonUpcomingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSeasonResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeSearch
}

func (r GetSeasonResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetSeasonResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTopAnimeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnimeSearch
}

func (r GetTopAnimeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetTopAnimeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTopCharactersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CharactersSearch
}

func (r GetTopCharactersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetTopCharactersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTopMangaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MangaSearch
}

func (r GetTopMangaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetTopMangaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTopPeopleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PeopleSearch
}

func (r GetTopPeopleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetTopPeopleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTopReviewsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetTopReviewsResponseJSON200
}
type GetTopReviews200DataData0 struct {
	Anime           *AnimeMeta                          `json:"anime,omitempty"`
	Date            *string                             `json:"date,omitempty"`
	EpisodesWatched *int                                `json:"episodes_watched,omitempty"`
	IsPreliminary   *bool                               `json:"is_preliminary,omitempty"`
	IsSpoiler       *bool                               `json:"is_spoiler,omitempty"`
	MalId           *int                                `json:"mal_id,omitempty"`
	Reactions       *GetTopReviews200DataData0Reactions `json:"reactions,omitempty"`
	Review          *string                             `json:"review,omitempty"`
	Score           *int                                `json:"score,omitempty"`
	Tags            *[]string                           `json:"tags,omitempty"`
	Type            *string                             `json:"type,omitempty"`
	Url             *string                             `json:"url,omitempty"`
	User            *UserMeta                           `json:"user,omitempty"`
}
type GetTopReviews200DataData1 struct {
	Date          *string                             `json:"date,omitempty"`
	IsPreliminary *bool                               `json:"is_preliminary,omitempty"`
	IsSpoiler     *bool                               `json:"is_spoiler,omitempty"`
	MalId         *int                                `json:"mal_id,omitempty"`
	Manga         *MangaMeta                          `json:"manga,omitempty"`
	Reactions     *GetTopReviews200DataData1Reactions `json:"reactions,omitempty"`
	Review        *string                             `json:"review,omitempty"`
	Score         *int                                `json:"score,omitempty"`
	Tags          *[]string                           `json:"tags,omitempty"`
	Type          *string                             `json:"type,omitempty"`
	Url           *string                             `json:"url,omitempty"`
	User          *UserMeta                           `json:"user,omitempty"`
}
type GetTopReviews_200_Data_Data_Item struct {
	union json.RawMessage
}

func (r GetTopReviewsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetTopReviewsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUsersSearchResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UsersSearch
}

func (r GetUsersSearchResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUsersSearchResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetUserByIdResponseJSON200
}

func (r GetUserByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserProfileResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetUserProfileResponseJSON200
}

func (r GetUserProfileResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserProfileResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserAboutResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserAbout
}

func (r GetUserAboutResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserAboutResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserAnimelistResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
}

func (r GetUserAnimelistResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserAnimelistResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserClubsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserClubs
}

func (r GetUserClubsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserClubsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserExternalResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ExternalLinks
}

func (r GetUserExternalResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserExternalResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserFavoritesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetUserFavoritesResponseJSON200
}

func (r GetUserFavoritesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserFavoritesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserFriendsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserFriends
}

func (r GetUserFriendsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserFriendsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserFullProfileResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetUserFullProfileResponseJSON200
}

func (r GetUserFullProfileResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserFullProfileResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserHistoryResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserHistory
}

func (r GetUserHistoryResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserHistoryResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserMangaListResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
}

func (r GetUserMangaListResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserMangaListResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserRecommendationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Recommendations
}

func (r GetUserRecommendationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserRecommendationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserReviewsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetUserReviewsResponseJSON200
}
type GetUserReviews200DataData0 struct {
	Anime           *AnimeMeta                           `json:"anime,omitempty"`
	Date            *string                              `json:"date,omitempty"`
	EpisodesWatched *int                                 `json:"episodes_watched,omitempty"`
	IsPreliminary   *bool                                `json:"is_preliminary,omitempty"`
	IsSpoiler       *bool                                `json:"is_spoiler,omitempty"`
	MalId           *int                                 `json:"mal_id,omitempty"`
	Reactions       *GetUserReviews200DataData0Reactions `json:"reactions,omitempty"`
	Review          *string                              `json:"review,omitempty"`
	Score           *int                                 `json:"score,omitempty"`
	Tags            *[]string                            `json:"tags,omitempty"`
	Type            *string                              `json:"type,omitempty"`
	Url             *string                              `json:"url,omitempty"`
	User            *UserMeta                            `json:"user,omitempty"`
}
type GetUserReviews200DataData1 struct {
	Date          *string                              `json:"date,omitempty"`
	IsPreliminary *bool                                `json:"is_preliminary,omitempty"`
	IsSpoiler     *bool                                `json:"is_spoiler,omitempty"`
	MalId         *int                                 `json:"mal_id,omitempty"`
	Manga         *MangaMeta                           `json:"manga,omitempty"`
	Reactions     *GetUserReviews200DataData1Reactions `json:"reactions,omitempty"`
	Review        *string                              `json:"review,omitempty"`
	Score         *int                                 `json:"score,omitempty"`
	Tags          *[]string                            `json:"tags,omitempty"`
	Type          *string                              `json:"type,omitempty"`
	Url           *string                              `json:"url,omitempty"`
	User          *UserMeta                            `json:"user,omitempty"`
}
type GetUserReviews_200_Data_Data_Item struct {
	union json.RawMessage
}

func (r GetUserReviewsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserReviewsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserStatisticsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserStatistics
}

func (r GetUserStatisticsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserStatisticsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserUpdatesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *UserUpdates
}

func (r GetUserUpdatesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetUserUpdatesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetWatchRecentEpisodesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WatchEpisodes
}

func (r GetWatchRecentEpisodesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetWatchRecentEpisodesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetWatchPopularEpisodesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WatchEpisodes
}

func (r GetWatchPopularEpisodesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetWatchPopularEpisodesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetWatchRecentPromosResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WatchPromos
}

func (r GetWatchRecentPromosResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetWatchRecentPromosResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetWatchPopularPromosResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WatchPromos
}

func (r GetWatchPopularPromosResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetWatchPopularPromosResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetAnimeSearchWithResponse(ctx context.Context, params *GetAnimeSearchParams, reqEditors ...RequestEditorFn) (*GetAnimeSearchResponse, error) {
	rsp, err := c.GetAnimeSearch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeSearchResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeByIdResponse, error) {
	rsp, err := c.GetAnimeById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeByIdResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeCharactersWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeCharactersResponse, error) {
	rsp, err := c.GetAnimeCharacters(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeCharactersResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeEpisodesWithResponse(ctx context.Context, id int, params *GetAnimeEpisodesParams, reqEditors ...RequestEditorFn) (*GetAnimeEpisodesResponse, error) {
	rsp, err := c.GetAnimeEpisodes(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeEpisodesResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeEpisodeByIdWithResponse(ctx context.Context, id int, episode int, reqEditors ...RequestEditorFn) (*GetAnimeEpisodeByIdResponse, error) {
	rsp, err := c.GetAnimeEpisodeById(ctx, id, episode, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeEpisodeByIdResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeExternalWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeExternalResponse, error) {
	rsp, err := c.GetAnimeExternal(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeExternalResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeForumWithResponse(ctx context.Context, id int, params *GetAnimeForumParams, reqEditors ...RequestEditorFn) (*GetAnimeForumResponse, error) {
	rsp, err := c.GetAnimeForum(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeForumResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeFullByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeFullByIdResponse, error) {
	rsp, err := c.GetAnimeFullById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeFullByIdResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeMoreInfoWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeMoreInfoResponse, error) {
	rsp, err := c.GetAnimeMoreInfo(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeMoreInfoResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeNewsWithResponse(ctx context.Context, id int, params *GetAnimeNewsParams, reqEditors ...RequestEditorFn) (*GetAnimeNewsResponse, error) {
	rsp, err := c.GetAnimeNews(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeNewsResponse(rsp)
}
func (c *ClientWithResponses) GetAnimePicturesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimePicturesResponse, error) {
	rsp, err := c.GetAnimePictures(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimePicturesResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeRecommendationsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeRecommendationsResponse, error) {
	rsp, err := c.GetAnimeRecommendations(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeRecommendationsResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeRelationsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeRelationsResponse, error) {
	rsp, err := c.GetAnimeRelations(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeRelationsResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeReviewsWithResponse(ctx context.Context, id int, params *GetAnimeReviewsParams, reqEditors ...RequestEditorFn) (*GetAnimeReviewsResponse, error) {
	rsp, err := c.GetAnimeReviews(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeReviewsResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeStaffWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeStaffResponse, error) {
	rsp, err := c.GetAnimeStaff(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeStaffResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeStatisticsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeStatisticsResponse, error) {
	rsp, err := c.GetAnimeStatistics(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeStatisticsResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeStreamingWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeStreamingResponse, error) {
	rsp, err := c.GetAnimeStreaming(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeStreamingResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeThemesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeThemesResponse, error) {
	rsp, err := c.GetAnimeThemes(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeThemesResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeUserUpdatesWithResponse(ctx context.Context, id int, params *GetAnimeUserUpdatesParams, reqEditors ...RequestEditorFn) (*GetAnimeUserUpdatesResponse, error) {
	rsp, err := c.GetAnimeUserUpdates(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeUserUpdatesResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeVideosWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAnimeVideosResponse, error) {
	rsp, err := c.GetAnimeVideos(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeVideosResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeVideosEpisodesWithResponse(ctx context.Context, id int, params *GetAnimeVideosEpisodesParams, reqEditors ...RequestEditorFn) (*GetAnimeVideosEpisodesResponse, error) {
	rsp, err := c.GetAnimeVideosEpisodes(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeVideosEpisodesResponse(rsp)
}
func (c *ClientWithResponses) GetCharactersSearchWithResponse(ctx context.Context, params *GetCharactersSearchParams, reqEditors ...RequestEditorFn) (*GetCharactersSearchResponse, error) {
	rsp, err := c.GetCharactersSearch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetCharactersSearchResponse(rsp)
}
func (c *ClientWithResponses) GetCharacterByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterByIdResponse, error) {
	rsp, err := c.GetCharacterById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetCharacterByIdResponse(rsp)
}
func (c *ClientWithResponses) GetCharacterAnimeWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterAnimeResponse, error) {
	rsp, err := c.GetCharacterAnime(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetCharacterAnimeResponse(rsp)
}
func (c *ClientWithResponses) GetCharacterFullByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterFullByIdResponse, error) {
	rsp, err := c.GetCharacterFullById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetCharacterFullByIdResponse(rsp)
}
func (c *ClientWithResponses) GetCharacterMangaWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterMangaResponse, error) {
	rsp, err := c.GetCharacterManga(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetCharacterMangaResponse(rsp)
}
func (c *ClientWithResponses) GetCharacterPicturesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterPicturesResponse, error) {
	rsp, err := c.GetCharacterPictures(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetCharacterPicturesResponse(rsp)
}
func (c *ClientWithResponses) GetCharacterVoiceActorsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetCharacterVoiceActorsResponse, error) {
	rsp, err := c.GetCharacterVoiceActors(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetCharacterVoiceActorsResponse(rsp)
}
func (c *ClientWithResponses) GetClubsSearchWithResponse(ctx context.Context, params *GetClubsSearchParams, reqEditors ...RequestEditorFn) (*GetClubsSearchResponse, error) {
	rsp, err := c.GetClubsSearch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetClubsSearchResponse(rsp)
}
func (c *ClientWithResponses) GetClubsByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetClubsByIdResponse, error) {
	rsp, err := c.GetClubsById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetClubsByIdResponse(rsp)
}
func (c *ClientWithResponses) GetClubMembersWithResponse(ctx context.Context, id int, params *GetClubMembersParams, reqEditors ...RequestEditorFn) (*GetClubMembersResponse, error) {
	rsp, err := c.GetClubMembers(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetClubMembersResponse(rsp)
}
func (c *ClientWithResponses) GetClubRelationsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetClubRelationsResponse, error) {
	rsp, err := c.GetClubRelations(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetClubRelationsResponse(rsp)
}
func (c *ClientWithResponses) GetClubStaffWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetClubStaffResponse, error) {
	rsp, err := c.GetClubStaff(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetClubStaffResponse(rsp)
}
func (c *ClientWithResponses) GetAnimeGenresWithResponse(ctx context.Context, params *GetAnimeGenresParams, reqEditors ...RequestEditorFn) (*GetAnimeGenresResponse, error) {
	rsp, err := c.GetAnimeGenres(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnimeGenresResponse(rsp)
}
func (c *ClientWithResponses) GetMangaGenresWithResponse(ctx context.Context, params *GetMangaGenresParams, reqEditors ...RequestEditorFn) (*GetMangaGenresResponse, error) {
	rsp, err := c.GetMangaGenres(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaGenresResponse(rsp)
}
func (c *ClientWithResponses) GetMagazinesWithResponse(ctx context.Context, params *GetMagazinesParams, reqEditors ...RequestEditorFn) (*GetMagazinesResponse, error) {
	rsp, err := c.GetMagazines(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMagazinesResponse(rsp)
}
func (c *ClientWithResponses) GetMangaSearchWithResponse(ctx context.Context, params *GetMangaSearchParams, reqEditors ...RequestEditorFn) (*GetMangaSearchResponse, error) {
	rsp, err := c.GetMangaSearch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaSearchResponse(rsp)
}
func (c *ClientWithResponses) GetMangaByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaByIdResponse, error) {
	rsp, err := c.GetMangaById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaByIdResponse(rsp)
}
func (c *ClientWithResponses) GetMangaCharactersWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaCharactersResponse, error) {
	rsp, err := c.GetMangaCharacters(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaCharactersResponse(rsp)
}
func (c *ClientWithResponses) GetMangaExternalWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaExternalResponse, error) {
	rsp, err := c.GetMangaExternal(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaExternalResponse(rsp)
}
func (c *ClientWithResponses) GetMangaTopicsWithResponse(ctx context.Context, id int, params *GetMangaTopicsParams, reqEditors ...RequestEditorFn) (*GetMangaTopicsResponse, error) {
	rsp, err := c.GetMangaTopics(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaTopicsResponse(rsp)
}
func (c *ClientWithResponses) GetMangaFullByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaFullByIdResponse, error) {
	rsp, err := c.GetMangaFullById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaFullByIdResponse(rsp)
}
func (c *ClientWithResponses) GetMangaMoreInfoWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaMoreInfoResponse, error) {
	rsp, err := c.GetMangaMoreInfo(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaMoreInfoResponse(rsp)
}
func (c *ClientWithResponses) GetMangaNewsWithResponse(ctx context.Context, id int, params *GetMangaNewsParams, reqEditors ...RequestEditorFn) (*GetMangaNewsResponse, error) {
	rsp, err := c.GetMangaNews(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaNewsResponse(rsp)
}
func (c *ClientWithResponses) GetMangaPicturesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaPicturesResponse, error) {
	rsp, err := c.GetMangaPictures(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaPicturesResponse(rsp)
}
func (c *ClientWithResponses) GetMangaRecommendationsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaRecommendationsResponse, error) {
	rsp, err := c.GetMangaRecommendations(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaRecommendationsResponse(rsp)
}
func (c *ClientWithResponses) GetMangaRelationsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaRelationsResponse, error) {
	rsp, err := c.GetMangaRelations(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaRelationsResponse(rsp)
}
func (c *ClientWithResponses) GetMangaReviewsWithResponse(ctx context.Context, id int, params *GetMangaReviewsParams, reqEditors ...RequestEditorFn) (*GetMangaReviewsResponse, error) {
	rsp, err := c.GetMangaReviews(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaReviewsResponse(rsp)
}
func (c *ClientWithResponses) GetMangaStatisticsWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetMangaStatisticsResponse, error) {
	rsp, err := c.GetMangaStatistics(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaStatisticsResponse(rsp)
}
func (c *ClientWithResponses) GetMangaUserUpdatesWithResponse(ctx context.Context, id int, params *GetMangaUserUpdatesParams, reqEditors ...RequestEditorFn) (*GetMangaUserUpdatesResponse, error) {
	rsp, err := c.GetMangaUserUpdates(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMangaUserUpdatesResponse(rsp)
}
func (c *ClientWithResponses) GetPeopleSearchWithResponse(ctx context.Context, params *GetPeopleSearchParams, reqEditors ...RequestEditorFn) (*GetPeopleSearchResponse, error) {
	rsp, err := c.GetPeopleSearch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPeopleSearchResponse(rsp)
}
func (c *ClientWithResponses) GetPersonByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonByIdResponse, error) {
	rsp, err := c.GetPersonById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPersonByIdResponse(rsp)
}
func (c *ClientWithResponses) GetPersonAnimeWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonAnimeResponse, error) {
	rsp, err := c.GetPersonAnime(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPersonAnimeResponse(rsp)
}
func (c *ClientWithResponses) GetPersonFullByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonFullByIdResponse, error) {
	rsp, err := c.GetPersonFullById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPersonFullByIdResponse(rsp)
}
func (c *ClientWithResponses) GetPersonMangaWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonMangaResponse, error) {
	rsp, err := c.GetPersonManga(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPersonMangaResponse(rsp)
}
func (c *ClientWithResponses) GetPersonPicturesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonPicturesResponse, error) {
	rsp, err := c.GetPersonPictures(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPersonPicturesResponse(rsp)
}
func (c *ClientWithResponses) GetPersonVoicesWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetPersonVoicesResponse, error) {
	rsp, err := c.GetPersonVoices(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPersonVoicesResponse(rsp)
}
func (c *ClientWithResponses) GetProducersWithResponse(ctx context.Context, params *GetProducersParams, reqEditors ...RequestEditorFn) (*GetProducersResponse, error) {
	rsp, err := c.GetProducers(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProducersResponse(rsp)
}
func (c *ClientWithResponses) GetProducerByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetProducerByIdResponse, error) {
	rsp, err := c.GetProducerById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProducerByIdResponse(rsp)
}
func (c *ClientWithResponses) GetProducerExternalWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetProducerExternalResponse, error) {
	rsp, err := c.GetProducerExternal(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProducerExternalResponse(rsp)
}
func (c *ClientWithResponses) GetProducerFullByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetProducerFullByIdResponse, error) {
	rsp, err := c.GetProducerFullById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProducerFullByIdResponse(rsp)
}
func (c *ClientWithResponses) GetRandomAnimeWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRandomAnimeResponse, error) {
	rsp, err := c.GetRandomAnime(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRandomAnimeResponse(rsp)
}
func (c *ClientWithResponses) GetRandomCharactersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRandomCharactersResponse, error) {
	rsp, err := c.GetRandomCharacters(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRandomCharactersResponse(rsp)
}
func (c *ClientWithResponses) GetRandomMangaWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRandomMangaResponse, error) {
	rsp, err := c.GetRandomManga(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRandomMangaResponse(rsp)
}
func (c *ClientWithResponses) GetRandomPeopleWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRandomPeopleResponse, error) {
	rsp, err := c.GetRandomPeople(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRandomPeopleResponse(rsp)
}
func (c *ClientWithResponses) GetRandomUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRandomUsersResponse, error) {
	rsp, err := c.GetRandomUsers(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRandomUsersResponse(rsp)
}
func (c *ClientWithResponses) GetRecentAnimeRecommendationsWithResponse(ctx context.Context, params *GetRecentAnimeRecommendationsParams, reqEditors ...RequestEditorFn) (*GetRecentAnimeRecommendationsResponse, error) {
	rsp, err := c.GetRecentAnimeRecommendations(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRecentAnimeRecommendationsResponse(rsp)
}
func (c *ClientWithResponses) GetRecentMangaRecommendationsWithResponse(ctx context.Context, params *GetRecentMangaRecommendationsParams, reqEditors ...RequestEditorFn) (*GetRecentMangaRecommendationsResponse, error) {
	rsp, err := c.GetRecentMangaRecommendations(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRecentMangaRecommendationsResponse(rsp)
}
func (c *ClientWithResponses) GetRecentAnimeReviewsWithResponse(ctx context.Context, params *GetRecentAnimeReviewsParams, reqEditors ...RequestEditorFn) (*GetRecentAnimeReviewsResponse, error) {
	rsp, err := c.GetRecentAnimeReviews(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRecentAnimeReviewsResponse(rsp)
}
func (c *ClientWithResponses) GetRecentMangaReviewsWithResponse(ctx context.Context, params *GetRecentMangaReviewsParams, reqEditors ...RequestEditorFn) (*GetRecentMangaReviewsResponse, error) {
	rsp, err := c.GetRecentMangaReviews(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRecentMangaReviewsResponse(rsp)
}
func (c *ClientWithResponses) GetSchedulesWithResponse(ctx context.Context, params *GetSchedulesParams, reqEditors ...RequestEditorFn) (*GetSchedulesResponse, error) {
	rsp, err := c.GetSchedules(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSchedulesResponse(rsp)
}
func (c *ClientWithResponses) GetSeasonsListWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetSeasonsListResponse, error) {
	rsp, err := c.GetSeasonsList(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSeasonsListResponse(rsp)
}
func (c *ClientWithResponses) GetSeasonNowWithResponse(ctx context.Context, params *GetSeasonNowParams, reqEditors ...RequestEditorFn) (*GetSeasonNowResponse, error) {
	rsp, err := c.GetSeasonNow(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSeasonNowResponse(rsp)
}
func (c *ClientWithResponses) GetSeasonUpcomingWithResponse(ctx context.Context, params *GetSeasonUpcomingParams, reqEditors ...RequestEditorFn) (*GetSeasonUpcomingResponse, error) {
	rsp, err := c.GetSeasonUpcoming(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSeasonUpcomingResponse(rsp)
}
func (c *ClientWithResponses) GetSeasonWithResponse(ctx context.Context, year int, season string, params *GetSeasonParams, reqEditors ...RequestEditorFn) (*GetSeasonResponse, error) {
	rsp, err := c.GetSeason(ctx, year, season, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSeasonResponse(rsp)
}
func (c *ClientWithResponses) GetTopAnimeWithResponse(ctx context.Context, params *GetTopAnimeParams, reqEditors ...RequestEditorFn) (*GetTopAnimeResponse, error) {
	rsp, err := c.GetTopAnime(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTopAnimeResponse(rsp)
}
func (c *ClientWithResponses) GetTopCharactersWithResponse(ctx context.Context, params *GetTopCharactersParams, reqEditors ...RequestEditorFn) (*GetTopCharactersResponse, error) {
	rsp, err := c.GetTopCharacters(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTopCharactersResponse(rsp)
}
func (c *ClientWithResponses) GetTopMangaWithResponse(ctx context.Context, params *GetTopMangaParams, reqEditors ...RequestEditorFn) (*GetTopMangaResponse, error) {
	rsp, err := c.GetTopManga(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTopMangaResponse(rsp)
}
func (c *ClientWithResponses) GetTopPeopleWithResponse(ctx context.Context, params *GetTopPeopleParams, reqEditors ...RequestEditorFn) (*GetTopPeopleResponse, error) {
	rsp, err := c.GetTopPeople(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTopPeopleResponse(rsp)
}
func (c *ClientWithResponses) GetTopReviewsWithResponse(ctx context.Context, params *GetTopReviewsParams, reqEditors ...RequestEditorFn) (*GetTopReviewsResponse, error) {
	rsp, err := c.GetTopReviews(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTopReviewsResponse(rsp)
}
func (c *ClientWithResponses) GetUsersSearchWithResponse(ctx context.Context, params *GetUsersSearchParams, reqEditors ...RequestEditorFn) (*GetUsersSearchResponse, error) {
	rsp, err := c.GetUsersSearch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUsersSearchResponse(rsp)
}
func (c *ClientWithResponses) GetUserByIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetUserByIdResponse, error) {
	rsp, err := c.GetUserById(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserByIdResponse(rsp)
}
func (c *ClientWithResponses) GetUserProfileWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserProfileResponse, error) {
	rsp, err := c.GetUserProfile(ctx, username, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserProfileResponse(rsp)
}
func (c *ClientWithResponses) GetUserAboutWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserAboutResponse, error) {
	rsp, err := c.GetUserAbout(ctx, username, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserAboutResponse(rsp)
}
func (c *ClientWithResponses) GetUserAnimelistWithResponse(ctx context.Context, username string, params *GetUserAnimelistParams, reqEditors ...RequestEditorFn) (*GetUserAnimelistResponse, error) {
	rsp, err := c.GetUserAnimelist(ctx, username, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserAnimelistResponse(rsp)
}
func (c *ClientWithResponses) GetUserClubsWithResponse(ctx context.Context, username string, params *GetUserClubsParams, reqEditors ...RequestEditorFn) (*GetUserClubsResponse, error) {
	rsp, err := c.GetUserClubs(ctx, username, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserClubsResponse(rsp)
}
func (c *ClientWithResponses) GetUserExternalWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserExternalResponse, error) {
	rsp, err := c.GetUserExternal(ctx, username, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserExternalResponse(rsp)
}
func (c *ClientWithResponses) GetUserFavoritesWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserFavoritesResponse, error) {
	rsp, err := c.GetUserFavorites(ctx, username, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserFavoritesResponse(rsp)
}
func (c *ClientWithResponses) GetUserFriendsWithResponse(ctx context.Context, username string, params *GetUserFriendsParams, reqEditors ...RequestEditorFn) (*GetUserFriendsResponse, error) {
	rsp, err := c.GetUserFriends(ctx, username, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserFriendsResponse(rsp)
}
func (c *ClientWithResponses) GetUserFullProfileWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserFullProfileResponse, error) {
	rsp, err := c.GetUserFullProfile(ctx, username, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserFullProfileResponse(rsp)
}
func (c *ClientWithResponses) GetUserHistoryWithResponse(ctx context.Context, username string, params *GetUserHistoryParams, reqEditors ...RequestEditorFn) (*GetUserHistoryResponse, error) {
	rsp, err := c.GetUserHistory(ctx, username, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserHistoryResponse(rsp)
}
func (c *ClientWithResponses) GetUserMangaListWithResponse(ctx context.Context, username string, params *GetUserMangaListParams, reqEditors ...RequestEditorFn) (*GetUserMangaListResponse, error) {
	rsp, err := c.GetUserMangaList(ctx, username, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserMangaListResponse(rsp)
}
func (c *ClientWithResponses) GetUserRecommendationsWithResponse(ctx context.Context, username string, params *GetUserRecommendationsParams, reqEditors ...RequestEditorFn) (*GetUserRecommendationsResponse, error) {
	rsp, err := c.GetUserRecommendations(ctx, username, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserRecommendationsResponse(rsp)
}
func (c *ClientWithResponses) GetUserReviewsWithResponse(ctx context.Context, username string, params *GetUserReviewsParams, reqEditors ...RequestEditorFn) (*GetUserReviewsResponse, error) {
	rsp, err := c.GetUserReviews(ctx, username, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserReviewsResponse(rsp)
}
func (c *ClientWithResponses) GetUserStatisticsWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserStatisticsResponse, error) {
	rsp, err := c.GetUserStatistics(ctx, username, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserStatisticsResponse(rsp)
}
func (c *ClientWithResponses) GetUserUpdatesWithResponse(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*GetUserUpdatesResponse, error) {
	rsp, err := c.GetUserUpdates(ctx, username, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserUpdatesResponse(rsp)
}
func (c *ClientWithResponses) GetWatchRecentEpisodesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetWatchRecentEpisodesResponse, error) {
	rsp, err := c.GetWatchRecentEpisodes(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetWatchRecentEpisodesResponse(rsp)
}
func (c *ClientWithResponses) GetWatchPopularEpisodesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetWatchPopularEpisodesResponse, error) {
	rsp, err := c.GetWatchPopularEpisodes(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetWatchPopularEpisodesResponse(rsp)
}
func (c *ClientWithResponses) GetWatchRecentPromosWithResponse(ctx context.Context, params *GetWatchRecentPromosParams, reqEditors ...RequestEditorFn) (*GetWatchRecentPromosResponse, error) {
	rsp, err := c.GetWatchRecentPromos(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetWatchRecentPromosResponse(rsp)
}
func (c *ClientWithResponses) GetWatchPopularPromosWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetWatchPopularPromosResponse, error) {
	rsp, err := c.GetWatchPopularPromos(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetWatchPopularPromosResponse(rsp)
}
func ParseGetAnimeSearchResponse(rsp *http.Response) (*GetAnimeSearchResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeSearchResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeByIdResponse(rsp *http.Response) (*GetAnimeByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetAnimeByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeCharactersResponse(rsp *http.Response) (*GetAnimeCharactersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeCharactersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeCharacters
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeEpisodesResponse(rsp *http.Response) (*GetAnimeEpisodesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeEpisodesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeEpisodes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeEpisodeByIdResponse(rsp *http.Response) (*GetAnimeEpisodeByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeEpisodeByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetAnimeEpisodeByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeExternalResponse(rsp *http.Response) (*GetAnimeExternalResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeExternalResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ExternalLinks
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeForumResponse(rsp *http.Response) (*GetAnimeForumResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeForumResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Forum
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeFullByIdResponse(rsp *http.Response) (*GetAnimeFullByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeFullByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetAnimeFullByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeMoreInfoResponse(rsp *http.Response) (*GetAnimeMoreInfoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeMoreInfoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Moreinfo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeNewsResponse(rsp *http.Response) (*GetAnimeNewsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeNewsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeNews
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimePicturesResponse(rsp *http.Response) (*GetAnimePicturesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimePicturesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PicturesVariants
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeRecommendationsResponse(rsp *http.Response) (*GetAnimeRecommendationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeRecommendationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest EntryRecommendations
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeRelationsResponse(rsp *http.Response) (*GetAnimeRelationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeRelationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetAnimeRelationsResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeReviewsResponse(rsp *http.Response) (*GetAnimeReviewsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeReviewsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeReviews
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeStaffResponse(rsp *http.Response) (*GetAnimeStaffResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeStaffResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeStaff
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeStatisticsResponse(rsp *http.Response) (*GetAnimeStatisticsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeStatisticsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeStatistics
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeStreamingResponse(rsp *http.Response) (*GetAnimeStreamingResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeStreamingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ExternalLinks
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeThemesResponse(rsp *http.Response) (*GetAnimeThemesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeThemesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeThemes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeUserUpdatesResponse(rsp *http.Response) (*GetAnimeUserUpdatesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeUserUpdatesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeUserupdates
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeVideosResponse(rsp *http.Response) (*GetAnimeVideosResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeVideosResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeVideos
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeVideosEpisodesResponse(rsp *http.Response) (*GetAnimeVideosEpisodesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeVideosEpisodesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeVideosEpisodes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetCharactersSearchResponse(rsp *http.Response) (*GetCharactersSearchResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetCharactersSearchResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CharactersSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetCharacterByIdResponse(rsp *http.Response) (*GetCharacterByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetCharacterByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetCharacterByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetCharacterAnimeResponse(rsp *http.Response) (*GetCharacterAnimeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetCharacterAnimeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CharacterAnime
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetCharacterFullByIdResponse(rsp *http.Response) (*GetCharacterFullByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetCharacterFullByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetCharacterFullByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetCharacterMangaResponse(rsp *http.Response) (*GetCharacterMangaResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetCharacterMangaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CharacterManga
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetCharacterPicturesResponse(rsp *http.Response) (*GetCharacterPicturesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetCharacterPicturesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CharacterPictures
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetCharacterVoiceActorsResponse(rsp *http.Response) (*GetCharacterVoiceActorsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetCharacterVoiceActorsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CharacterVoiceActors
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetClubsSearchResponse(rsp *http.Response) (*GetClubsSearchResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetClubsSearchResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ClubsSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetClubsByIdResponse(rsp *http.Response) (*GetClubsByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetClubsByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetClubsByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetClubMembersResponse(rsp *http.Response) (*GetClubMembersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetClubMembersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetClubMembersResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetClubRelationsResponse(rsp *http.Response) (*GetClubRelationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetClubRelationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ClubRelations
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetClubStaffResponse(rsp *http.Response) (*GetClubStaffResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetClubStaffResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ClubStaff
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetAnimeGenresResponse(rsp *http.Response) (*GetAnimeGenresResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetAnimeGenresResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Genres
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaGenresResponse(rsp *http.Response) (*GetMangaGenresResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaGenresResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Genres
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMagazinesResponse(rsp *http.Response) (*GetMagazinesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMagazinesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Magazines
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaSearchResponse(rsp *http.Response) (*GetMangaSearchResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaSearchResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MangaSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaByIdResponse(rsp *http.Response) (*GetMangaByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetMangaByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaCharactersResponse(rsp *http.Response) (*GetMangaCharactersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaCharactersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MangaCharacters
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaExternalResponse(rsp *http.Response) (*GetMangaExternalResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaExternalResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ExternalLinks
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaTopicsResponse(rsp *http.Response) (*GetMangaTopicsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaTopicsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Forum
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaFullByIdResponse(rsp *http.Response) (*GetMangaFullByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaFullByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetMangaFullByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaMoreInfoResponse(rsp *http.Response) (*GetMangaMoreInfoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaMoreInfoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Moreinfo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaNewsResponse(rsp *http.Response) (*GetMangaNewsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaNewsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MangaNews
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaPicturesResponse(rsp *http.Response) (*GetMangaPicturesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaPicturesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MangaPictures
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaRecommendationsResponse(rsp *http.Response) (*GetMangaRecommendationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaRecommendationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest EntryRecommendations
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaRelationsResponse(rsp *http.Response) (*GetMangaRelationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaRelationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetMangaRelationsResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaReviewsResponse(rsp *http.Response) (*GetMangaReviewsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaReviewsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MangaReviews
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaStatisticsResponse(rsp *http.Response) (*GetMangaStatisticsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaStatisticsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MangaStatistics
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetMangaUserUpdatesResponse(rsp *http.Response) (*GetMangaUserUpdatesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetMangaUserUpdatesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MangaUserupdates
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetPeopleSearchResponse(rsp *http.Response) (*GetPeopleSearchResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetPeopleSearchResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PeopleSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetPersonByIdResponse(rsp *http.Response) (*GetPersonByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetPersonByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetPersonByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetPersonAnimeResponse(rsp *http.Response) (*GetPersonAnimeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetPersonAnimeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PersonAnime
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetPersonFullByIdResponse(rsp *http.Response) (*GetPersonFullByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetPersonFullByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetPersonFullByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetPersonMangaResponse(rsp *http.Response) (*GetPersonMangaResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetPersonMangaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PersonManga
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetPersonPicturesResponse(rsp *http.Response) (*GetPersonPicturesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetPersonPicturesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PersonPictures
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetPersonVoicesResponse(rsp *http.Response) (*GetPersonVoicesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetPersonVoicesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PersonVoiceActingRoles
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetProducersResponse(rsp *http.Response) (*GetProducersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetProducersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Producers
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetProducerByIdResponse(rsp *http.Response) (*GetProducerByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetProducerByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetProducerByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetProducerExternalResponse(rsp *http.Response) (*GetProducerExternalResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetProducerExternalResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ExternalLinks
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetProducerFullByIdResponse(rsp *http.Response) (*GetProducerFullByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetProducerFullByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetProducerFullByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetRandomAnimeResponse(rsp *http.Response) (*GetRandomAnimeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetRandomAnimeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetRandomAnimeResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetRandomCharactersResponse(rsp *http.Response) (*GetRandomCharactersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetRandomCharactersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetRandomCharactersResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetRandomMangaResponse(rsp *http.Response) (*GetRandomMangaResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetRandomMangaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetRandomMangaResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetRandomPeopleResponse(rsp *http.Response) (*GetRandomPeopleResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetRandomPeopleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetRandomPeopleResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetRandomUsersResponse(rsp *http.Response) (*GetRandomUsersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetRandomUsersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetRandomUsersResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetRecentAnimeRecommendationsResponse(rsp *http.Response) (*GetRecentAnimeRecommendationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetRecentAnimeRecommendationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Recommendations
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetRecentMangaRecommendationsResponse(rsp *http.Response) (*GetRecentMangaRecommendationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetRecentMangaRecommendationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Recommendations
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetRecentAnimeReviewsResponse(rsp *http.Response) (*GetRecentAnimeReviewsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetRecentAnimeReviewsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetRecentMangaReviewsResponse(rsp *http.Response) (*GetRecentMangaReviewsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetRecentMangaReviewsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetSchedulesResponse(rsp *http.Response) (*GetSchedulesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetSchedulesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Schedules
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetSeasonsListResponse(rsp *http.Response) (*GetSeasonsListResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetSeasonsListResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Seasons
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetSeasonNowResponse(rsp *http.Response) (*GetSeasonNowResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetSeasonNowResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetSeasonUpcomingResponse(rsp *http.Response) (*GetSeasonUpcomingResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetSeasonUpcomingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetSeasonResponse(rsp *http.Response) (*GetSeasonResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetSeasonResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetTopAnimeResponse(rsp *http.Response) (*GetTopAnimeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetTopAnimeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnimeSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetTopCharactersResponse(rsp *http.Response) (*GetTopCharactersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetTopCharactersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CharactersSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetTopMangaResponse(rsp *http.Response) (*GetTopMangaResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetTopMangaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MangaSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetTopPeopleResponse(rsp *http.Response) (*GetTopPeopleResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetTopPeopleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PeopleSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetTopReviewsResponse(rsp *http.Response) (*GetTopReviewsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetTopReviewsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetTopReviewsResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUsersSearchResponse(rsp *http.Response) (*GetUsersSearchResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUsersSearchResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UsersSearch
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserByIdResponse(rsp *http.Response) (*GetUserByIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetUserByIdResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserProfileResponse(rsp *http.Response) (*GetUserProfileResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserProfileResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetUserProfileResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserAboutResponse(rsp *http.Response) (*GetUserAboutResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserAboutResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserAbout
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserAnimelistResponse(rsp *http.Response) (*GetUserAnimelistResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserAnimelistResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserClubsResponse(rsp *http.Response) (*GetUserClubsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserClubsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserClubs
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserExternalResponse(rsp *http.Response) (*GetUserExternalResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserExternalResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ExternalLinks
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserFavoritesResponse(rsp *http.Response) (*GetUserFavoritesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserFavoritesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetUserFavoritesResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserFriendsResponse(rsp *http.Response) (*GetUserFriendsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserFriendsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserFriends
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserFullProfileResponse(rsp *http.Response) (*GetUserFullProfileResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserFullProfileResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetUserFullProfileResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserHistoryResponse(rsp *http.Response) (*GetUserHistoryResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserHistoryResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserHistory
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserMangaListResponse(rsp *http.Response) (*GetUserMangaListResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserMangaListResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserRecommendationsResponse(rsp *http.Response) (*GetUserRecommendationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserRecommendationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Recommendations
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserReviewsResponse(rsp *http.Response) (*GetUserReviewsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserReviewsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetUserReviewsResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserStatisticsResponse(rsp *http.Response) (*GetUserStatisticsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserStatisticsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserStatistics
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetUserUpdatesResponse(rsp *http.Response) (*GetUserUpdatesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetUserUpdatesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest UserUpdates
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetWatchRecentEpisodesResponse(rsp *http.Response) (*GetWatchRecentEpisodesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetWatchRecentEpisodesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WatchEpisodes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetWatchPopularEpisodesResponse(rsp *http.Response) (*GetWatchPopularEpisodesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetWatchPopularEpisodesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WatchEpisodes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetWatchRecentPromosResponse(rsp *http.Response) (*GetWatchRecentPromosResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetWatchRecentPromosResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WatchPromos
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
func ParseGetWatchPopularPromosResponse(rsp *http.Response) (*GetWatchPopularPromosResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetWatchPopularPromosResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WatchPromos
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	case rsp.StatusCode != 200:
		return nil, fmt.Errorf("unexpected status code: %d", rsp.StatusCode)
	}
	return response, nil
}
