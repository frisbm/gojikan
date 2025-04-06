package gojikan

import (
	"errors"
	"fmt"
	"strings"
)

// ErrorResponse represents the error response from the API.
type ErrorResponse struct {
	Status    int    `json:"status"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	ApiError  string `json:"error"`
	ReportUrl string `json:"report_url"`
}

// Error implements the error interface for ErrorResponse.
func (e *ErrorResponse) Error() string {
	var errorParts []string
	if e.Status != 0 {
		errorParts = append(errorParts, fmt.Sprintf("Status: %d", e.Status))
	}
	if e.Type != "" {
		errorParts = append(errorParts, fmt.Sprintf("Type: %s", e.Type))
	}
	if e.Message != "" {
		errorParts = append(errorParts, fmt.Sprintf("Message: %s", e.Message))
	}
	if e.ApiError != "" {
		errorParts = append(errorParts, fmt.Sprintf("ApiError: %s", e.ApiError))
	}
	if e.ReportUrl != "" {
		errorParts = append(errorParts, fmt.Sprintf("ReportUrl: %s", e.ReportUrl))
	}
	return fmt.Sprintf("error response: %s", strings.Join(errorParts, ", "))
}

func (e *ErrorResponse) Is(target error) bool {
	var t *ErrorResponse
	ok := errors.As(target, &t)
	if !ok {
		return false
	}
	return e.Status == t.Status && e.Type == t.Type && e.Message == t.Message && e.ApiError == t.ApiError
}

type Response[T any] struct {
	Data       T           `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Pagination represents the pagination information in the API response.
type Pagination struct {
	LastVisiblePage int  `json:"last_visible_page"`
	HasNextPage     bool `json:"has_next_page"`
}

type AnimeImages struct {
	Jpg  *Jpg  `json:"jpg,omitempty"`
	Webp *Webp `json:"webp,omitempty"`
}

// Webp Available images in WEBP
type Webp struct {
	// Image URL WEBP
	ImageUrl string `json:"image_url,omitempty"`
	// Small Image URL WEBP
	SmallImageUrl string `json:"small_image_url,omitempty"`
	// Image URL WEBP
	LargeImageUrl string `json:"large_image_url,omitempty"`
}

// Jpg Available images in JPG
type Jpg struct {
	// Image URL JPG
	ImageUrl string `json:"image_url,omitempty"`
	// Small Image URL JPG
	SmallImageUrl string `json:"small_image_url,omitempty"`
	// Image URL JPG
	LargeImageUrl string `json:"large_image_url,omitempty"`
}

// TrailerBase Youtube Details
type TrailerBase struct {
	// YouTube ID
	YoutubeId string `json:"youtube_id,omitempty"`
	// YouTube URL
	Url string `json:"url,omitempty"`
	// Parsed Embed URL
	EmbedUrl string `json:"embed_url,omitempty"`
}

// Daterange Date range
type Daterange struct {
	// Date ISO8601
	From string `json:"from,omitempty"`
	// Date ISO8601
	To   string `json:"to,omitempty"`
	Prop *Prop  `json:"prop,omitempty"`
}

// Date Date object
type Date struct {
	// Day
	Day int32 `json:"day,omitempty"`
	// Month
	Month int32 `json:"month,omitempty"`
	// Year
	Year int32 `json:"year,omitempty"`
}

// Prop Date Prop
type Prop struct {
	From *Date `json:"from,omitempty"`
	To   *Date `json:"to,omitempty"`
	// Raw parsed string
	String string `json:"string,omitempty"`
}

// Title Parsed Title Data
type Title struct {
	// Title type
	Type string `json:"type,omitempty"`
	// Title value
	Title string `json:"title,omitempty"`
}

// Broadcast Details
type Broadcast struct {
	// Day of the week
	Day string `json:"day,omitempty"`
	// Time in 24 hour format
	Time string `json:"time,omitempty"`
	// Timezone (Tz Database format https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
	Timezone string `json:"timezone,omitempty"`
	// Raw parsed broadcast string
	String string `json:"string,omitempty"`
}

// MalUrl Parsed URL Data
type MalUrl struct {
	// MyAnimeList ID
	MalId int32 `json:"mal_id,omitempty"`
	// Type of resource
	Type string `json:"type,omitempty"`
	// Resource Name
	Name string `json:"name,omitempty"`
	// Resource Title
	Title string `json:"title,omitempty"`
	// MyAnimeList URL
	Url string `json:"url,omitempty"`
}

// AnimeFullRelations is the relations of the anime
type AnimeFullRelations struct {
	// Relation type
	Relation string   `json:"relation,omitempty"`
	Entry    []MalUrl `json:"entry,omitempty"`
}

type AnimeFullTheme struct {
	Openings []string `json:"openings,omitempty"`
	Endings  []string `json:"endings,omitempty"`
}

type AnimeFullExternal struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}
