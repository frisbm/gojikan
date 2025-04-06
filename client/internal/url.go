package internal

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"
)

func BuildUrl(baseUrl string, path string, pathParams map[string]any, queryParams map[string][]any) (string, error) {
	if baseUrl == "" {
		return "", fmt.Errorf("baseUrl cannot be empty")
	}
	if path == "" {
		path = "/"
	}
	var sb strings.Builder
	sb.WriteString(baseUrl)
	pathKeys := slices.Sorted(maps.Keys(pathParams))
	for _, key := range pathKeys {
		path = strings.ReplaceAll(path, fmt.Sprintf("{%s}", key), marshalParam(pathParams[key]))
	}
	sb.WriteString(path)

	if len(queryParams) > 0 {
		queryKeys := slices.Sorted(maps.Keys(queryParams))
		paramCount := 0
		sb.WriteString("?")

		for _, key := range queryKeys {
			values, ok := queryParams[key]
			if !ok {
				continue
			}
			if len(values) == 0 {
				continue
			}
			if paramCount > 0 {
				sb.WriteString("&")
			}
			if len(values) == 1 {
				param := marshalParam(values[0])
				if param == "true" {
					sb.WriteString(key)
				} else if param == "false" {
					// do nothing
					sb.WriteString("")
				} else {
					sb.WriteString(fmt.Sprintf("%s=%v", key, param))
				}
			}
			if len(values) > 1 {
				var paramValues []string
				for _, value := range values {
					param := marshalParam(value)
					if param == "true" {
						paramValues = append(paramValues, key)
						break
					} else if param == "false" {
						break
					} else {
						paramValues = append(paramValues, param)
					}
				}
				sb.WriteString(fmt.Sprintf("%s=%s", key, strings.Join(paramValues, ",")))
			}
			paramCount++
		}
	}
	return sb.String(), nil
}

type Stringer interface {
	String() string
}

func marshalParam(param any) string {
	if param == nil {
		return ""
	}
	switch v := param.(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case time.Time:
		return v.Format(time.DateOnly)
	default:
		if str, ok := param.(Stringer); ok {
			return str.String()
		}
		return fmt.Sprintf("%v", v)
	}
}
