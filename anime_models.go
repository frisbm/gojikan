package gojikan

type Anime struct {
	// MyAnimeList ID
	MalId int32 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	Url     string       `json:"url,omitempty"`
	Images  *AnimeImages `json:"images,omitempty"`
	Trailer *TrailerBase `json:"trailer,omitempty"`
	// Whether the entry is pending approval on MAL or not
	Approved bool `json:"approved,omitempty"`
	// All titles
	Titles []Title `json:"titles,omitempty"`
	// Title
	Title string `json:"title,omitempty"`
	// English Title
	TitleEnglish string `json:"title_english,omitempty"`
	// Japanese Title
	TitleJapanese string `json:"title_japanese,omitempty"`
	// Other Titles
	TitleSynonyms []string `json:"title_synonyms,omitempty"`
	// Anime Type
	Type AnimeType `json:"type,omitempty"`
	// Original Material/Source adapted from
	Source string `json:"source,omitempty"`
	// Episode count
	Episodes int32 `json:"episodes,omitempty"`
	// Airing status
	Status AiringStatus `json:"status,omitempty"`
	// Airing boolean
	Airing bool       `json:"airing,omitempty"`
	Aired  *Daterange `json:"aired,omitempty"`
	// Parsed raw duration
	Duration string `json:"duration,omitempty"`
	// Anime audience rating
	Rating AnimeRating `json:"rating,omitempty"`
	// Score
	Score float64 `json:"score,omitempty"`
	// Number of users
	ScoredBy int32 `json:"scored_by,omitempty"`
	// Ranking
	Rank int32 `json:"rank,omitempty"`
	// Popularity
	Popularity int32 `json:"popularity,omitempty"`
	// Number of users who have added this entry to their list
	Members int32 `json:"members,omitempty"`
	// Number of users who have favorited this entry
	Favorites int32 `json:"favorites,omitempty"`
	// Synopsis
	Synopsis string `json:"synopsis,omitempty"`
	// Background
	Background string `json:"background,omitempty"`
	// Season
	Season Season `json:"season,omitempty"`
	// Year
	Year           int32      `json:"year,omitempty"`
	Broadcast      *Broadcast `json:"broadcast,omitempty"`
	Producers      []MalUrl   `json:"producers,omitempty"`
	Licensors      []MalUrl   `json:"licensors,omitempty"`
	Studios        []MalUrl   `json:"studios,omitempty"`
	Genres         []MalUrl   `json:"genres,omitempty"`
	ExplicitGenres []MalUrl   `json:"explicit_genres,omitempty"`
	Themes         []MalUrl   `json:"themes,omitempty"`
	Demographics   []MalUrl   `json:"demographics,omitempty"`
}

type AnimeFull struct {
	// MyAnimeList ID
	MalId int32 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	Url     string       `json:"url,omitempty"`
	Images  *AnimeImages `json:"images,omitempty"`
	Trailer *TrailerBase `json:"trailer,omitempty"`
	// Whether the entry is pending approval on MAL or not
	Approved bool `json:"approved,omitempty"`
	// All titles
	Titles []Title `json:"titles,omitempty"`
	// Title
	Title string `json:"title,omitempty"`
	// English Title
	TitleEnglish string `json:"title_english,omitempty"`
	// Japanese Title
	TitleJapanese string `json:"title_japanese,omitempty"`
	// Other Titles
	TitleSynonyms []string `json:"title_synonyms,omitempty"`
	// Anime Type
	Type AnimeType `json:"type,omitempty"`
	// Original Material/Source adapted from
	Source string `json:"source,omitempty"`
	// Episode count
	Episodes int32 `json:"episodes,omitempty"`
	// Airing status
	Status AiringStatus `json:"status,omitempty"`
	// Airing boolean
	Airing bool       `json:"airing,omitempty"`
	Aired  *Daterange `json:"aired,omitempty"`
	// Parsed raw duration
	Duration string `json:"duration,omitempty"`
	// Anime audience rating
	Rating AnimeRating `json:"rating,omitempty"`
	// Score
	Score float64 `json:"score,omitempty"`
	// Number of users
	ScoredBy int32 `json:"scored_by,omitempty"`
	// Ranking
	Rank int32 `json:"rank,omitempty"`
	// Popularity
	Popularity int32 `json:"popularity,omitempty"`
	// Number of users who have added this entry to their list
	Members int32 `json:"members,omitempty"`
	// Number of users who have favorited this entry
	Favorites int32 `json:"favorites,omitempty"`
	// Synopsis
	Synopsis string `json:"synopsis,omitempty"`
	// Background
	Background string `json:"background,omitempty"`
	// Season
	Season Season `json:"season,omitempty"`
	// Year
	Year           int32                `json:"year,omitempty"`
	Broadcast      *Broadcast           `json:"broadcast,omitempty"`
	Producers      []MalUrl             `json:"producers,omitempty"`
	Licensors      []MalUrl             `json:"licensors,omitempty"`
	Studios        []MalUrl             `json:"studios,omitempty"`
	Genres         []MalUrl             `json:"genres,omitempty"`
	ExplicitGenres []MalUrl             `json:"explicit_genres,omitempty"`
	Themes         []MalUrl             `json:"themes,omitempty"`
	Demographics   []MalUrl             `json:"demographics,omitempty"`
	Relations      []AnimeFullRelations `json:"relations,omitempty"`
	Theme          *AnimeFullTheme      `json:"theme,omitempty"`
	External       []AnimeFullExternal  `json:"external,omitempty"`
	Streaming      []AnimeFullExternal  `json:"streaming,omitempty"`
}

type AnimeCharacters struct {
	Character *AnimeCharactersCharacter `json:"character,omitempty"`
	// Character's Role
	Role        string                       `json:"role,omitempty"`
	VoiceActors []AnimeCharactersVoiceActors `json:"voice_actors,omitempty"`
}

type AnimeCharactersCharacter struct {
	// MyAnimeList ID
	MalId int32 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	Url    string           `json:"url,omitempty"`
	Images *CharacterImages `json:"images,omitempty"`
	// Character Name
	Name string `json:"name,omitempty"`
}

type AnimeCharactersVoiceActors struct {
	Person   *AnimeCharactersPerson `json:"person,omitempty"`
	Language string                 `json:"language,omitempty"`
}

type AnimeCharactersPerson struct {
	MalId  int32         `json:"mal_id,omitempty"`
	Url    string        `json:"url,omitempty"`
	Images *PeopleImages `json:"images,omitempty"`
	Name   string        `json:"name,omitempty"`
}

type CharacterImages struct {
	Jpg  *CharacterImagesJpg  `json:"jpg,omitempty"`
	Webp *CharacterImagesWebp `json:"webp,omitempty"`
}

type CharacterImagesJpg struct {
	// Image URL JPG
	ImageUrl string `json:"image_url,omitempty"`
	// Small Image URL JPG
	SmallImageUrl string `json:"small_image_url,omitempty"`
}

type CharacterImagesWebp struct {
	// Image URL WEBP
	ImageUrl string `json:"image_url,omitempty"`
	// Small Image URL WEBP
	SmallImageUrl string `json:"small_image_url,omitempty"`
}

type PeopleImages struct {
	Jpg *Jpg `json:"jpg,omitempty"`
}

type AnimeStaff struct {
	Person *AnimeStaffPerson `json:"person,omitempty"`
	// Staff Positions
	Positions []string `json:"positions,omitempty"`
}

type AnimeStaffPerson struct {
	// MyAnimeList ID
	MalId int32 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	Url    string        `json:"url,omitempty"`
	Images *PeopleImages `json:"images,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
}
