package gojikan

import (
	"bytes"
	"fmt"
	"strings"
)

type stringer interface {
	String() string
}

func unmarshalJson[T stringer](b []byte, values []T) (T, error) {
	b = bytes.Trim(b, "\"")
	if len(b) == 0 {
		return *new(T), nil
	}
	s := string(b)

	for _, value := range values {
		if strings.EqualFold(value.String(), s) {
			return value, nil
		}
	}
	return *new(T), fmt.Errorf("invalid value %s, expected one of %v", s, values)
}

type FilterTopic string

const (
	// FilterTopicAll All topics
	FilterTopicAll FilterTopic = "all"
	// FilterTopicEpisode Episode topics
	FilterTopicEpisode FilterTopic = "episode"
	// FilterTopicOther Other topics
	FilterTopicOther FilterTopic = "other"
)

func (f FilterTopic) String() string {
	return string(f)
}

var filterTopicValues = []FilterTopic{
	FilterTopicAll,
	FilterTopicEpisode,
	FilterTopicOther,
}

func (f *FilterTopic) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, filterTopicValues)
	if err != nil {
		return err
	}
	*f = v
	return nil
}

type FilterDay string

const (
	// FilterDayMonday Monday
	FilterDayMonday FilterDay = "monday"
	// FilterDayTuesday Tuesday
	FilterDayTuesday FilterDay = "tuesday"
	// FilterDayWednesday Wednesday
	FilterDayWednesday FilterDay = "wednesday"
	// FilterDayThursday Thursday
	FilterDayThursday FilterDay = "thursday"
	// FilterDayFriday Friday
	FilterDayFriday FilterDay = "friday"
	// FilterDaySaturday Saturday
	FilterDaySaturday FilterDay = "saturday"
	// FilterDaySunday Sunday
	FilterDaySunday FilterDay = "sunday"
	// FilterDayUnknown Unknown
	FilterDayUnknown FilterDay = "unknown"
	// FilterDayOther Other
	FilterDayOther FilterDay = "other"
)

func (f FilterDay) String() string {
	return string(f)
}

var filterDayValues = []FilterDay{
	FilterDayMonday,
	FilterDayTuesday,
	FilterDayWednesday,
	FilterDayThursday,
	FilterDayFriday,
	FilterDaySaturday,
	FilterDaySunday,
	FilterDayUnknown,
	FilterDayOther,
}

func (f *FilterDay) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, filterDayValues)
	if err != nil {
		return err
	}
	*f = v
	return nil
}

type FilterKids string

const (
	// FilterKidsTrue True
	FilterKidsTrue FilterKids = "true"
	// FilterKidsFalse False
	FilterKidsFalse FilterKids = "false"
)

func (f FilterKids) String() string {
	return string(f)
}

var filterKidsValues = []FilterKids{
	FilterKidsTrue,
	FilterKidsFalse,
}

func (f *FilterKids) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, filterKidsValues)
	if err != nil {
		return err
	}
	*f = v
	return nil
}

type FilterSfw string

const (
	// FilterSfwTrue True
	FilterSfwTrue FilterSfw = "true"
	// FilterSfwFalse False
	FilterSfwFalse FilterSfw = "false"
)

func (f FilterSfw) String() string {
	return string(f)
}

var filterSfwValues = []FilterSfw{
	FilterSfwTrue,
	FilterSfwFalse,
}

func (f *FilterSfw) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, filterSfwValues)
	if err != nil {
		return err
	}
	*f = v
	return nil
}

type FilterEntryType string

const (
	// FilterEntryTypeTV TV
	FilterEntryTypeTV FilterEntryType = "tv"
	// FilterEntryTypeMovie Movie
	FilterEntryTypeMovie FilterEntryType = "movie"
	// FilterEntryTypeOva OVA
	FilterEntryTypeOva FilterEntryType = "ova"
	// FilterEntryTypeSpecial Special
	FilterEntryTypeSpecial FilterEntryType = "special"
	// FilterEntryTypeOna ONA
	FilterEntryTypeOna FilterEntryType = "ona"
	// FilterEntryTypeMusic Music
	FilterEntryTypeMusic FilterEntryType = "music"
)

func (f FilterEntryType) String() string {
	return string(f)
}

var filterEntryTypeValues = []FilterEntryType{
	FilterEntryTypeTV,
	FilterEntryTypeMovie,
	FilterEntryTypeOva,
	FilterEntryTypeSpecial,
	FilterEntryTypeOna,
	FilterEntryTypeMusic,
}

func (f *FilterEntryType) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, filterEntryTypeValues)
	if err != nil {
		return err
	}
	*f = v
	return nil
}

type FilterUserHistoryType string

const (
	// FilterUserHistoryTypeAnime Anime
	FilterUserHistoryTypeAnime FilterUserHistoryType = "anime"
	// FilterUserHistoryTypeManga Manga
	FilterUserHistoryTypeManga FilterUserHistoryType = "manga"
)

func (f FilterUserHistoryType) String() string {
	return string(f)
}

var filterUserHistoryTypeValues = []FilterUserHistoryType{
	FilterUserHistoryTypeAnime,
	FilterUserHistoryTypeManga,
}

func (f *FilterUserHistoryType) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, filterUserHistoryTypeValues)
	if err != nil {
		return err
	}
	*f = v
	return nil
}

type FilterUserAnimeListStatus string

const (
	// FilterUserAnimeListStatusAll All
	FilterUserAnimeListStatusAll FilterUserAnimeListStatus = "all"
	// FilterUserAnimeListStatusWatching Watching
	FilterUserAnimeListStatusWatching FilterUserAnimeListStatus = "watching"
	// FilterUserAnimeListStatusCompleted Completed
	FilterUserAnimeListStatusCompleted FilterUserAnimeListStatus = "completed"
	// FilterUserAnimeListStatusOnHold On Hold
	FilterUserAnimeListStatusOnHold FilterUserAnimeListStatus = "onhold"
	// FilterUserAnimeListStatusDropped Dropped
	FilterUserAnimeListStatusDropped FilterUserAnimeListStatus = "dropped"
	// FilterUserAnimeListStatusPlanToWatch Plan to Watch
	FilterUserAnimeListStatusPlanToWatch FilterUserAnimeListStatus = "plantowatch"
)

func (f FilterUserAnimeListStatus) String() string {
	return string(f)
}

var filterUserAnimeListStatusValues = []FilterUserAnimeListStatus{
	FilterUserAnimeListStatusAll,
	FilterUserAnimeListStatusWatching,
	FilterUserAnimeListStatusCompleted,
	FilterUserAnimeListStatusOnHold,
	FilterUserAnimeListStatusDropped,
	FilterUserAnimeListStatusPlanToWatch,
}

func (f *FilterUserAnimeListStatus) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, filterUserAnimeListStatusValues)
	if err != nil {
		return err
	}
	*f = v
	return nil
}

type OrderByAnime string

const (
	// OrderByAnimeMalId MyAnimeList ID
	OrderByAnimeMalId OrderByAnime = "mal_id"
	// OrderByAnimeTitle Title
	OrderByAnimeTitle OrderByAnime = "title"
	// OrderByAnimeStartDate Start Date
	OrderByAnimeStartDate OrderByAnime = "start_date"
	// OrderByAnimeEndDate End Date
	OrderByAnimeEndDate OrderByAnime = "end_date"
	// OrderByAnimeEpisodes Episodes
	OrderByAnimeEpisodes OrderByAnime = "episodes"
	// OrderByAnimeScore Score
	OrderByAnimeScore OrderByAnime = "score"
	// OrderByAnimeScoredBy Scored By
	OrderByAnimeScoredBy OrderByAnime = "scored_by"
	// OrderByAnimeRank Rank
	OrderByAnimeRank OrderByAnime = "rank"
	// OrderByAnimePopularity Popularity
	OrderByAnimePopularity OrderByAnime = "popularity"
	// OrderByAnimeMembers Members
	OrderByAnimeMembers OrderByAnime = "members"
	// OrderByAnimeFavorites Favorites
	OrderByAnimeFavorites OrderByAnime = "favorites"
)

func (o OrderByAnime) String() string {
	return string(o)
}

var orderByAnimeValues = []OrderByAnime{
	OrderByAnimeMalId,
	OrderByAnimeTitle,
	OrderByAnimeStartDate,
	OrderByAnimeEndDate,
	OrderByAnimeEpisodes,
	OrderByAnimeScore,
	OrderByAnimeScoredBy,
	OrderByAnimeRank,
	OrderByAnimePopularity,
	OrderByAnimeMembers,
	OrderByAnimeFavorites,
}

func (o *OrderByAnime) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, orderByAnimeValues)
	if err != nil {
		return err
	}
	*o = v
	return nil
}

type FilterAnimeRating string

const (
	// FilterAnimeRatingG G
	FilterAnimeRatingG FilterAnimeRating = "g"
	// FilterAnimeRatingPG PG
	FilterAnimeRatingPG FilterAnimeRating = "pg"
	// FilterAnimeRatingPG13 PG13
	FilterAnimeRatingPG13 FilterAnimeRating = "pg13"
	// FilterAnimeRatingR17 R17
	FilterAnimeRatingR17 FilterAnimeRating = "r17"
	// FilterAnimeRatingR R
	FilterAnimeRatingR FilterAnimeRating = "r"
	// FilterAnimeRatingRx Rx
	FilterAnimeRatingRx FilterAnimeRating = "rx"
)

func (a FilterAnimeRating) String() string {
	return string(a)
}

var filterAnimeRatingValues = []FilterAnimeRating{
	FilterAnimeRatingG,
	FilterAnimeRatingPG,
	FilterAnimeRatingPG13,
	FilterAnimeRatingR17,
	FilterAnimeRatingR,
	FilterAnimeRatingRx,
}

func (a *FilterAnimeRating) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, filterAnimeRatingValues)
	if err != nil {
		return err
	}
	*a = v
	return nil
}

type AnimeStatus string

const (
	// AnimeStatusAiring Airing
	AnimeStatusAiring AnimeStatus = "airing"
	// AnimeStatusComplete Complete
	AnimeStatusComplete AnimeStatus = "complete"
	// AnimeStatusUpcoming Upcoming
	AnimeStatusUpcoming AnimeStatus = "upcoming"
)

func (a AnimeStatus) String() string {
	return string(a)
}

var animeStatusValues = []AnimeStatus{
	AnimeStatusAiring,
	AnimeStatusComplete,
	AnimeStatusUpcoming,
}

func (a *AnimeStatus) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, animeStatusValues)
	if err != nil {
		return err
	}
	*a = v
	return nil
}

type FilterAnimeType string

const (
	// FilterAnimeTypeTV TV
	FilterAnimeTypeTV FilterAnimeType = "tv"
	// FilterAnimeTypeMovie Movie
	FilterAnimeTypeMovie FilterAnimeType = "movie"
	// FilterAnimeTypeOva OVA
	FilterAnimeTypeOva FilterAnimeType = "ova"
	// FilterAnimeTypeSpecial Special
	FilterAnimeTypeSpecial FilterAnimeType = "special"
	// FilterAnimeTypeOna ONA
	FilterAnimeTypeOna FilterAnimeType = "ona"
	// FilterAnimeTypeMusic Music
	FilterAnimeTypeMusic FilterAnimeType = "music"
	// FilterAnimeTypeCM CM
	FilterAnimeTypeCM FilterAnimeType = "cm"
	// FilterAnimeTypePV PV
	FilterAnimeTypePV FilterAnimeType = "pv"
	// FilterAnimeTypeTVSpecial TV Special
	FilterAnimeTypeTVSpecial FilterAnimeType = "tv_special"
)

func (a FilterAnimeType) String() string {
	return string(a)
}

var filterAnimeTypeValues = []FilterAnimeType{
	FilterAnimeTypeTV,
	FilterAnimeTypeMovie,
	FilterAnimeTypeOva,
	FilterAnimeTypeSpecial,
	FilterAnimeTypeOna,
	FilterAnimeTypeMusic,
	FilterAnimeTypeCM,
	FilterAnimeTypePV,
	FilterAnimeTypeTVSpecial,
}

func (a *FilterAnimeType) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, filterAnimeTypeValues)
	if err != nil {
		return err
	}
	*a = v
	return nil
}

type CharacterOrderBy string

const (
	// CharacterOrderByMalId MyAnimeList ID
	CharacterOrderByMalId CharacterOrderBy = "mal_id"
	// CharacterOrderByName Name
	CharacterOrderByName CharacterOrderBy = "name"
	// CharacterOrderByFavorites Favorites
	CharacterOrderByFavorites CharacterOrderBy = "favorites"
)

func (c CharacterOrderBy) String() string {
	return string(c)
}

var characterOrderByValues = []CharacterOrderBy{
	CharacterOrderByMalId,
	CharacterOrderByName,
	CharacterOrderByFavorites,
}

func (c *CharacterOrderBy) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, characterOrderByValues)
	if err != nil {
		return err
	}
	*c = v
	return nil
}

type QueryCategory string

const (
	// QueryCategoryAnime Anime
	QueryCategoryAnime QueryCategory = "anime"
	// QueryCategoryManga Manga
	QueryCategoryManga QueryCategory = "manga"
	// QueryCategoryActorsAndArtists Actors and Artists
	QueryCategoryActorsAndArtists QueryCategory = "actors_and_artists"
	// QueryCategoryCharacters Characters
	QueryCategoryCharacters QueryCategory = "characters"
	// QueryCategoryCitiesAndNeighborhoods Cities and Neighborhoods
	QueryCategoryCitiesAndNeighborhoods QueryCategory = "cities_and_neighborhoods"
	// QueryCategoryCompanies Companies
	QueryCategoryCompanies QueryCategory = "companies"
	// QueryCategoryConventions Conventions
	QueryCategoryConventions QueryCategory = "conventions"
	// QueryCategoryGames Games
	QueryCategoryGames QueryCategory = "games"
	// QueryCategoryJapan Japan
	QueryCategoryJapan QueryCategory = "japan"
	// QueryCategoryMusic Music
	QueryCategoryMusic QueryCategory = "music"
	// QueryCategoryOther Other
	QueryCategoryOther QueryCategory = "other"
	// QueryCategorySchools Schools
	QueryCategorySchools QueryCategory = "schools"
)

func (c QueryCategory) String() string {
	return string(c)
}

var queryCategoryValues = []QueryCategory{
	QueryCategoryAnime,
	QueryCategoryManga,
	QueryCategoryActorsAndArtists,
	QueryCategoryCharacters,
	QueryCategoryCitiesAndNeighborhoods,
	QueryCategoryCompanies,
	QueryCategoryConventions,
	QueryCategoryGames,
	QueryCategoryJapan,
	QueryCategoryMusic,
	QueryCategoryOther,
	QueryCategorySchools,
}

func (c *QueryCategory) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, queryCategoryValues)
	if err != nil {
		return err
	}
	*c = v
	return nil
}

type ClubOrderBy string

const (
	// ClubOrderByMalId MyAnimeList ID
	ClubOrderByMalId ClubOrderBy = "mal_id"
	// ClubOrderByName Name
	ClubOrderByName ClubOrderBy = "name"
	// ClubOrderByMembersCount Members Count
	ClubOrderByMembersCount ClubOrderBy = "members_count"
	// ClubOrderByCreated Created
	ClubOrderByCreated ClubOrderBy = "created"
)

func (c ClubOrderBy) String() string {
	return string(c)
}

var clubOrderByValues = []ClubOrderBy{
	ClubOrderByMalId,
	ClubOrderByName,
	ClubOrderByMembersCount,
	ClubOrderByCreated,
}

func (c *ClubOrderBy) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, clubOrderByValues)
	if err != nil {
		return err
	}
	*c = v
	return nil
}

type ClubType string

const (
	// ClubTypePublic Public
	ClubTypePublic ClubType = "public"
	// ClubTypePrivate Private
	ClubTypePrivate ClubType = "private"
	// ClubTypeSecret Secret
	ClubTypeSecret ClubType = "secret"
)

func (c ClubType) String() string {
	return string(c)
}

var clubTypeValues = []ClubType{
	ClubTypePublic,
	ClubTypePrivate,
	ClubTypeSecret,
}

func (c *ClubType) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, clubTypeValues)
	if err != nil {
		return err
	}
	*c = v
	return nil
}

type Gender string

const (
	// GenderAny Any
	GenderAny Gender = "any"
	// GenderMale Male
	GenderMale Gender = "male"
	// GenderFemale Female
	GenderFemale Gender = "female"
	// GenderNonbinary Nonbinary
	GenderNonbinary Gender = "nonbinary"
)

func (g Gender) String() string {
	return string(g)
}

var genderValues = []Gender{
	GenderAny,
	GenderMale,
	GenderFemale,
	GenderNonbinary,
}

func (g *Gender) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, genderValues)
	if err != nil {
		return err
	}
	*g = v
	return nil
}

type GenreFilter string

const (
	// GenreFilterGenres Genres
	GenreFilterGenres GenreFilter = "genres"
	// GenreFilterExplicitGenres Explicit Genres
	GenreFilterExplicitGenres GenreFilter = "explicit_genres"
	// GenreFilterThemes Themes
	GenreFilterThemes GenreFilter = "themes"
	// GenreFilterDemographics Demographics
	GenreFilterDemographics GenreFilter = "demographics"
)

func (g GenreFilter) String() string {
	return string(g)
}

var genreFilterValues = []GenreFilter{
	GenreFilterGenres,
	GenreFilterExplicitGenres,
	GenreFilterThemes,
	GenreFilterDemographics,
}

func (g *GenreFilter) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, genreFilterValues)
	if err != nil {
		return err
	}
	*g = v
	return nil
}

type MagazineOrderBy string

const (
	// MagazineOrderByMalId MyAnimeList ID
	MagazineOrderByMalId MagazineOrderBy = "mal_id"
	// MagazineOrderByName Name
	MagazineOrderByName MagazineOrderBy = "name"
	// MagazineOrderByCount Count
	MagazineOrderByCount MagazineOrderBy = "count"
)

func (m MagazineOrderBy) String() string {
	return string(m)
}

var magazineOrderByValues = []MagazineOrderBy{
	MagazineOrderByMalId,
	MagazineOrderByName,
	MagazineOrderByCount,
}

func (m *MagazineOrderBy) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, magazineOrderByValues)
	if err != nil {
		return err
	}
	*m = v
	return nil
}

type UserMangaListStatusFilter string

const (
	// UserMangaListStatusFilterAll All
	UserMangaListStatusFilterAll UserMangaListStatusFilter = "all"
	// UserMangaListStatusFilterReading Reading
	UserMangaListStatusFilterReading UserMangaListStatusFilter = "reading"
	// UserMangaListStatusFilterCompleted Completed
	UserMangaListStatusFilterCompleted UserMangaListStatusFilter = "completed"
	// UserMangaListStatusFilterOnHold On Hold
	UserMangaListStatusFilterOnHold UserMangaListStatusFilter = "onhold"
	// UserMangaListStatusFilterDropped Dropped
	UserMangaListStatusFilterDropped UserMangaListStatusFilter = "dropped"
	// UserMangaListStatusFilterPlanToRead Plan to Read
	UserMangaListStatusFilterPlanToRead UserMangaListStatusFilter = "plantoread"
)

func (u UserMangaListStatusFilter) String() string {
	return string(u)
}

var userMangaListStatusFilterValues = []UserMangaListStatusFilter{
	UserMangaListStatusFilterAll,
	UserMangaListStatusFilterReading,
	UserMangaListStatusFilterCompleted,
	UserMangaListStatusFilterOnHold,
	UserMangaListStatusFilterDropped,
	UserMangaListStatusFilterPlanToRead,
}

func (u *UserMangaListStatusFilter) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, userMangaListStatusFilterValues)
	if err != nil {
		return err
	}
	*u = v
	return nil
}

type MangaOrderBy string

const (
	// MangaOrderByMalId MyAnimeList ID
	MangaOrderByMalId MangaOrderBy = "mal_id"
	// MangaOrderByTitle Title
	MangaOrderByTitle MangaOrderBy = "title"
	// MangaOrderByStartDate Start Date
	MangaOrderByStartDate MangaOrderBy = "start_date"
	// MangaOrderByEndDate End Date
	MangaOrderByEndDate MangaOrderBy = "end_date"
	// MangaOrderByChapters Chapters
	MangaOrderByChapters MangaOrderBy = "chapters"
	// MangaOrderByVolumes Volumes
	MangaOrderByVolumes MangaOrderBy = "volumes"
	// MangaOrderByScore Score
	MangaOrderByScore MangaOrderBy = "score"
	// MangaOrderByScoredBy Scored By
	MangaOrderByScoredBy MangaOrderBy = "scored_by"
	// MangaOrderByRank Rank
	MangaOrderByRank MangaOrderBy = "rank"
	// MangaOrderByPopularity Popularity
	MangaOrderByPopularity MangaOrderBy = "popularity"
	// MangaOrderByMembers Members
	MangaOrderByMembers MangaOrderBy = "members"
	// MangaOrderByFavorites Favorites
	MangaOrderByFavorites MangaOrderBy = "favorites"
)

func (m MangaOrderBy) String() string {
	return string(m)
}

var mangaOrderByValues = []MangaOrderBy{
	MangaOrderByMalId,
	MangaOrderByTitle,
	MangaOrderByStartDate,
	MangaOrderByEndDate,
	MangaOrderByChapters,
	MangaOrderByVolumes,
	MangaOrderByScore,
	MangaOrderByScoredBy,
	MangaOrderByRank,
	MangaOrderByPopularity,
	MangaOrderByMembers,
	MangaOrderByFavorites,
}

func (m *MangaOrderBy) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, mangaOrderByValues)
	if err != nil {
		return err
	}
	*m = v
	return nil
}

type MangaStatus string

const (
	// MangaStatusPublishing Publishing
	MangaStatusPublishing MangaStatus = "publishing"
	// MangaStatusComplete Complete
	MangaStatusComplete MangaStatus = "complete"
	// MangaStatusHiatus Hiatus
	MangaStatusHiatus MangaStatus = "hiatus"
	// MangaStatusDiscontinued Discontinued
	MangaStatusDiscontinued MangaStatus = "discontinued"
	// MangaStatusUpcoming Upcoming
	MangaStatusUpcoming MangaStatus = "upcoming"
)

func (m MangaStatus) String() string {
	return string(m)
}

var mangaStatusValues = []MangaStatus{
	MangaStatusPublishing,
	MangaStatusComplete,
	MangaStatusHiatus,
	MangaStatusDiscontinued,
	MangaStatusUpcoming,
}

func (m *MangaStatus) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, mangaStatusValues)
	if err != nil {
		return err
	}
	*m = v
	return nil
}

type FilterMangaType string

const (
	// FilterMangaTypeManga Manga
	FilterMangaTypeManga FilterMangaType = "manga"
	// FilterMangaTypeNovel Novel
	FilterMangaTypeNovel FilterMangaType = "novel"
	// FilterMangaTypeLightNovel Light Novel
	FilterMangaTypeLightNovel FilterMangaType = "lightnovel"
	// FilterMangaTypeOneShot One Shot
	FilterMangaTypeOneShot FilterMangaType = "oneshot"
	// FilterMangaTypeDoujin Doujin
	FilterMangaTypeDoujin FilterMangaType = "doujin"
	// FilterMangaTypeManhwa Manhwa
	FilterMangaTypeManhwa FilterMangaType = "manhwa"
	// FilterMangaTypeManhua Manhua
	FilterMangaTypeManhua FilterMangaType = "manhua"
)

func (m FilterMangaType) String() string {
	return string(m)
}

var filterMangaTypeValues = []FilterMangaType{
	FilterMangaTypeManga,
	FilterMangaTypeNovel,
	FilterMangaTypeLightNovel,
	FilterMangaTypeOneShot,
	FilterMangaTypeDoujin,
	FilterMangaTypeManhwa,
	FilterMangaTypeManhua,
}

func (m *FilterMangaType) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, filterMangaTypeValues)
	if err != nil {
		return err
	}
	*m = v
	return nil
}

type PeopleOrderBy string

const (
	// PeopleOrderByMalId MyAnimeList ID
	PeopleOrderByMalId PeopleOrderBy = "mal_id"
	// PeopleOrderByName Name
	PeopleOrderByName PeopleOrderBy = "name"
	// PeopleOrderByBirthday Birthday
	PeopleOrderByBirthday PeopleOrderBy = "birthday"
	// PeopleOrderByFavorites Favorites
	PeopleOrderByFavorites PeopleOrderBy = "favorites"
)

func (p PeopleOrderBy) String() string {
	return string(p)
}

var peopleOrderByValues = []PeopleOrderBy{
	PeopleOrderByMalId,
	PeopleOrderByName,
	PeopleOrderByBirthday,
	PeopleOrderByFavorites,
}

func (p *PeopleOrderBy) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, peopleOrderByValues)
	if err != nil {
		return err
	}
	*p = v
	return nil
}

type PeopleStatus string

const (
	// PeopleStatusActive Active
	PeopleStatusActive PeopleStatus = "active"
	// PeopleStatusRetired Retired
	PeopleStatusRetired PeopleStatus = "retired"
	// PeopleStatusUnknown Unknown
	PeopleStatusUnknown PeopleStatus = "unknown"
	// PeopleStatusDeceased Deceased
	PeopleStatusDeceased PeopleStatus = "deceased"
	// PeopleStatusOther Other
	PeopleStatusOther PeopleStatus = "other"
)

func (p PeopleStatus) String() string {
	return string(p)
}

var peopleStatusValues = []PeopleStatus{
	PeopleStatusActive,
	PeopleStatusRetired,
	PeopleStatusUnknown,
	PeopleStatusDeceased,
	PeopleStatusOther,
}

func (p *PeopleStatus) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, peopleStatusValues)
	if err != nil {
		return err
	}
	*p = v
	return nil
}

type SortDirection string

const (
	// SortDirectionDesc Descending
	SortDirectionDesc SortDirection = "desc"
	// SortDirectionAsc Ascending
	SortDirectionAsc SortDirection = "asc"
)

func (s SortDirection) String() string {
	return string(s)
}

var sortDirectionValues = []SortDirection{
	SortDirectionDesc,
	SortDirectionAsc,
}

func (s *SortDirection) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, sortDirectionValues)
	if err != nil {
		return err
	}
	*s = v
	return nil
}

type TopAnimeFilter string

const (
	// TopAnimeFilterAiring Airing
	TopAnimeFilterAiring TopAnimeFilter = "airing"
	// TopAnimeFilterUpcoming Upcoming
	TopAnimeFilterUpcoming TopAnimeFilter = "upcoming"
	// TopAnimeFilterByPopularity By Popularity
	TopAnimeFilterByPopularity TopAnimeFilter = "bypopularity"
	// TopAnimeFilterFavorite Favorite
	TopAnimeFilterFavorite TopAnimeFilter = "favorite"
)

func (t TopAnimeFilter) String() string {
	return string(t)
}

var topAnimeFilterValues = []TopAnimeFilter{
	TopAnimeFilterAiring,
	TopAnimeFilterUpcoming,
	TopAnimeFilterByPopularity,
	TopAnimeFilterFavorite,
}

func (t *TopAnimeFilter) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, topAnimeFilterValues)
	if err != nil {
		return err
	}
	*t = v
	return nil
}

type TopMangaFilter string

const (
	// TopMangaFilterPublishing Publishing
	TopMangaFilterPublishing TopMangaFilter = "publishing"
	// TopMangaFilterUpcoming Upcoming
	TopMangaFilterUpcoming TopMangaFilter = "upcoming"
	// TopMangaFilterByPopularity By Popularity
	TopMangaFilterByPopularity TopMangaFilter = "bypopularity"
	// TopMangaFilterFavorite Favorite
	TopMangaFilterFavorite TopMangaFilter = "favorite"
)

func (t TopMangaFilter) String() string {
	return string(t)
}

var topMangaFilterValues = []TopMangaFilter{
	TopMangaFilterPublishing,
	TopMangaFilterUpcoming,
	TopMangaFilterByPopularity,
	TopMangaFilterFavorite,
}

func (t *TopMangaFilter) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, topMangaFilterValues)
	if err != nil {
		return err
	}
	*t = v
	return nil
}

type TopReviewsType string

const (
	// TopReviewsTypeAnime Anime
	TopReviewsTypeAnime TopReviewsType = "anime"
	// TopReviewsTypeManga Manga
	TopReviewsTypeManga TopReviewsType = "manga"
)

func (t TopReviewsType) String() string {
	return string(t)
}

var topReviewsTypeValues = []TopReviewsType{
	TopReviewsTypeAnime,
	TopReviewsTypeManga,
}

func (t *TopReviewsType) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, topReviewsTypeValues)
	if err != nil {
		return err
	}
	*t = v
	return nil
}

type AnimeType string

const (
	// AnimeTypeTV TV
	AnimeTypeTV AnimeType = "TV"
	// AnimeTypeOVA OVA
	AnimeTypeOVA AnimeType = "OVA"
	// AnimeTypeMovie Movie
	AnimeTypeMovie AnimeType = "Movie"
	// AnimeTypeSpecial Special
	AnimeTypeSpecial AnimeType = "Special"
	// AnimeTypeONA ONA
	AnimeTypeONA AnimeType = "ONA"
	// AnimeTypeMusic Music
	AnimeTypeMusic AnimeType = "Music"
)

func (a AnimeType) String() string {
	return string(a)
}

var animeTypeValues = []AnimeType{
	AnimeTypeTV,
	AnimeTypeOVA,
	AnimeTypeMovie,
	AnimeTypeSpecial,
	AnimeTypeONA,
	AnimeTypeMusic,
}

func (a *AnimeType) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, animeTypeValues)
	if err != nil {
		return err
	}
	*a = v
	return nil
}

type AiringStatus string

const (
	// AiringStatusFinishedAiring Finished Airing
	AiringStatusFinishedAiring AiringStatus = "Finished Airing"
	// AiringStatusCurrentlyAiring Currently Airing
	AiringStatusCurrentlyAiring AiringStatus = "Currently Airing"
	// AiringStatusNotYetAired Not Yet Aired
	AiringStatusNotYetAired AiringStatus = "Not yet aired"
)

func (a AiringStatus) String() string {
	return string(a)
}

var airingStatusValues = []AiringStatus{
	AiringStatusFinishedAiring,
	AiringStatusCurrentlyAiring,
	AiringStatusNotYetAired,
}

func (a *AiringStatus) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, airingStatusValues)
	if err != nil {
		return err
	}
	*a = v
	return nil
}

type AnimeRating string

const (
	// AnimeRatingG G - All Ages
	AnimeRatingG AnimeRating = "G - All Ages"
	// AnimeRatingPG PG - Children
	AnimeRatingPG AnimeRating = "PG - Children"
	// AnimeRatingPG13 PG-13 - Teens 13 or older
	AnimeRatingPG13 AnimeRating = "PG-13 - Teens 13 or older"
	// AnimeRatingR R - 17+ (violence & profanity)
	AnimeRatingR AnimeRating = "R - 17+ (violence & profanity)"
	// AnimeRatingRPlus R+ - Mild Nudity
	AnimeRatingRPlus AnimeRating = "R+ - Mild Nudity"
	// AnimeRatingRx Rx - Hentai
	AnimeRatingRx AnimeRating = "Rx - Hentai"
)

func (a AnimeRating) String() string {
	return string(a)
}

var animeRatingValues = []AnimeRating{
	AnimeRatingG,
	AnimeRatingPG,
	AnimeRatingPG13,
	AnimeRatingR,
	AnimeRatingRPlus,
	AnimeRatingRx,
}

func (a *AnimeRating) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, animeRatingValues)
	if err != nil {
		return err
	}
	*a = v
	return nil
}

type Season string

const (
	// SeasonWinter Winter season
	SeasonWinter Season = "winter"
	// SeasonSpring Spring season
	SeasonSpring Season = "spring"
	// SeasonSummer Summer season
	SeasonSummer Season = "summer"
	// SeasonFall Fall season
	SeasonFall Season = "fall"
)

func (s Season) String() string {
	return string(s)
}

var seasonValues = []Season{
	SeasonWinter,
	SeasonSpring,
	SeasonSummer,
	SeasonFall,
}

func (s *Season) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, seasonValues)
	if err != nil {
		return err
	}
	*s = v
	return nil
}

type ClubCategory string

const (
	// ClubCategoryActorsAndArtists Actors and Artists
	ClubCategoryActorsAndArtists ClubCategory = "actors & artists"
	// ClubCategoryAnime Anime
	ClubCategoryAnime ClubCategory = "anime"
	// ClubCategoryCharacters Characters
	ClubCategoryCharacters ClubCategory = "characters"
	// ClubCategoryCitiesAndNeighborhoods Cities and Neighborhoods
	ClubCategoryCitiesAndNeighborhoods ClubCategory = "cities & neighborhoods"
	// ClubCategoryCompanies Companies
	ClubCategoryCompanies ClubCategory = "companies"
	// ClubCategoryConventions Conventions
	ClubCategoryConventions ClubCategory = "conventions"
	// ClubCategoryGames Games
	ClubCategoryGames ClubCategory = "games"
	// ClubCategoryJapan Japan
	ClubCategoryJapan ClubCategory = "japan"
	// ClubCategoryManga Manga
	ClubCategoryManga ClubCategory = "manga"
	// ClubCategoryMusic Music
	ClubCategoryMusic ClubCategory = "music"
	// ClubCategoryOther Other
	ClubCategoryOther ClubCategory = "others"
	// ClubCategorySchools Schools
	ClubCategorySchools ClubCategory = "schools"
)

func (c ClubCategory) String() string {
	return string(c)
}

var clubCategoryValues = []ClubCategory{
	ClubCategoryActorsAndArtists,
	ClubCategoryAnime,
	ClubCategoryCharacters,
	ClubCategoryCitiesAndNeighborhoods,
	ClubCategoryCompanies,
	ClubCategoryConventions,
	ClubCategoryGames,
	ClubCategoryJapan,
	ClubCategoryManga,
	ClubCategoryMusic,
	ClubCategoryOther,
	ClubCategorySchools,
}

func (c *ClubCategory) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, clubCategoryValues)
	if err != nil {
		return err
	}
	*c = v
	return nil
}

type ClubAccess string

const (
	// ClubAccessPublic Public
	ClubAccessPublic ClubAccess = "public"
	// ClubAccessPrivate Private
	ClubAccessPrivate ClubAccess = "private"
	// ClubAccessSecret Secret
	ClubAccessSecret ClubAccess = "secret"
)

func (c ClubAccess) String() string {
	return string(c)
}

var clubAccessValues = []ClubAccess{
	ClubAccessPublic,
	ClubAccessPrivate,
	ClubAccessSecret,
}

func (c *ClubAccess) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, clubAccessValues)
	if err != nil {
		return err
	}
	*c = v
	return nil
}

type MangaType string

const (
	// MangaTypeManga Manga
	MangaTypeManga MangaType = "Manga"
	// MangaTypeNovel Novel
	MangaTypeNovel MangaType = "Novel"
	// MangaTypeLightNovel Light Novel
	MangaTypeLightNovel MangaType = "Light Novel"
	// MangaTypeOneShot One Shot
	MangaTypeOneShot MangaType = "One-shot"
	// MangaTypeDoujinshi Doujinshi
	MangaTypeDoujinshi MangaType = "Doujinshi"
	// MangaTypeManhua Manhua
	MangaTypeManhua MangaType = "Manhua"
	// MangaTypeManhwa Manhwa
	MangaTypeManhwa MangaType = "Manhwa"
	// MangaTypeOEL OEL
	MangaTypeOEL MangaType = "OEL"
)

func (m MangaType) String() string {
	return string(m)
}

var mangaTypeValues = []MangaType{
	MangaTypeManga,
	MangaTypeNovel,
	MangaTypeLightNovel,
	MangaTypeOneShot,
	MangaTypeDoujinshi,
	MangaTypeManhua,
	MangaTypeManhwa,
	MangaTypeOEL,
}

func (m *MangaType) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, mangaTypeValues)
	if err != nil {
		return err
	}
	*m = v
	return nil
}

type MangaStatusType string

const (
	// MangaStatusTypeFinished Finished
	MangaStatusTypeFinished MangaStatusType = "Finished"
	// MangaStatusTypePublishing Publishing
	MangaStatusTypePublishing MangaStatusType = "Publishing"
	// MangaStatusTypeOnHiatus On Hiatus
	MangaStatusTypeOnHiatus MangaStatusType = "On Hiatus"
	// MangaStatusTypeDiscontinued Discontinued
	MangaStatusTypeDiscontinued MangaStatusType = "Discontinued"
	// MangaStatusTypeNotYetPublished Not yet published
	MangaStatusTypeNotYetPublished MangaStatusType = "Not yet published"
)

func (m MangaStatusType) String() string {
	return string(m)
}

var mangaStatusTypeValues = []MangaStatusType{
	MangaStatusTypeFinished,
	MangaStatusTypePublishing,
	MangaStatusTypeOnHiatus,
	MangaStatusTypeDiscontinued,
	MangaStatusTypeNotYetPublished,
}

func (m *MangaStatusType) UnmarshalJSON(b []byte) error {
	v, err := unmarshalJson(b, mangaStatusTypeValues)
	if err != nil {
		return err
	}
	*m = v
	return nil
}
