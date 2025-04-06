package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testEnum string

const testEnumTest testEnum = "test"

func (e testEnum) String() string {
	return string(e)
}

func TestBuildUrl(t *testing.T) {
	tests := []struct {
		name        string
		baseUrl     string
		path        string
		pathParams  map[string]any
		queryParams map[string][]any
		want        string
		wantErr     bool
	}{
		{
			name:    "only a base url",
			baseUrl: "https://example.com",
			want:    "https://example.com/",
		},
		{
			name:    "base url with path",
			baseUrl: "https://example.com",
			path:    "/api/v1",
			want:    "https://example.com/api/v1",
		},
		{
			name:    "base url with path and path params",
			baseUrl: "https://example.com",
			path:    "/api/v1/{int}/{string}/{bool}/{time}/{enum}",
			pathParams: map[string]any{
				"int":    123,
				"string": "test",
				"bool":   true,
				"time":   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				"enum":   testEnumTest,
			},
			want: "https://example.com/api/v1/123/test/true/2025-01-01/test",
		},
		{
			name:    "base url with path and path params and query params",
			baseUrl: "https://example.com",
			path:    "/api/v1",
			queryParams: map[string][]any{
				"oneint":      {1},
				"manyints":    {1, 2, 3},
				"onebool":     {true},
				"manybools":   {true, false, true},
				"onestring":   {"one"},
				"manystrings": {"one", "two", "three"},
				"oneenum":     {testEnumTest},
				"manyenums":   {testEnumTest, testEnumTest},
				"onetime":     {time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
				"manytimes":   {time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)},
				"empty":       {},
			},
			want: "https://example.com/api/v1?manybools=manybools&manyenums=test,test&manyints=1,2,3&manystrings=one,two,three&manytimes=2025-01-01,2025-01-02&onebool&oneenum=test&oneint=1&onestring=one&onetime=2025-01-01",
		},
		{
			name:    "empty base url",
			baseUrl: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildUrl(tt.baseUrl, tt.path, tt.pathParams, tt.queryParams)
			if err != nil {
				require.True(t, tt.wantErr, "unexpected error: %v", err)
				return
			}
			require.False(t, tt.wantErr, "expected no error")
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_marshalParam(t *testing.T) {
	tests := []struct {
		name  string
		param any
		want  string
	}{
		{
			name:  "string",
			param: "test",
			want:  "test",
		},
		{
			name:  "int",
			param: 123,
			want:  "123",
		},
		{
			name:  "bool",
			param: true,
			want:  "true",
		},
		{
			name:  "time",
			param: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:  "2025-01-01",
		},
		{
			name:  "enum",
			param: testEnumTest,
			want:  "test",
		},
		{
			name:  "nil",
			param: nil,
			want:  "",
		},
		{
			name: "something else",
			param: struct {
				Name string
			}{
				Name: "test",
			},
			want: "{test}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := marshalParam(tt.param)
			require.Equal(t, tt.want, got)
		})
	}
}
