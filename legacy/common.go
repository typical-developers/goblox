package legacy

import (
	"encoding/json"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

type WrappedResponse[T any] struct {
	Data T `json:"data"`
}

// PageResponse is a page of results from an endpoint that uses cursors.
// An empty NextPageCursor means there are no more pages.
type PageResponse[T any] struct {
	PreviousPageCursor string `json:"previousPageCursor"`
	NextPageCursor     string `json:"nextPageCursor"`
	Data               []T    `json:"data"`
}

type SortOrder string

const (
	SortOrderAsc  SortOrder = "Asc"
	SortOrderDesc SortOrder = "Desc"
)

// PageOptions are the paging options shared by most paged endpoints.
// Limit is usually one of 10, 25, 50, or 100.
type PageOptions struct {
	Limit     *int      `url:"limit,omitempty"`
	Cursor    *string   `url:"cursor,omitempty"`
	SortOrder SortOrder `url:"sortOrder,omitempty"`
}

// Enum is an enum value from a legacy API. Roblox returns some enums by name
// and others by number, so either is kept as a string ("User" or "1").
type Enum string

func (e *Enum) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	if len(b) > 0 && b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*e = Enum(s)
		return nil
	}

	*e = Enum(b)
	return nil
}

// RelatedEntity is a reference to another entity, such as a game's creator or root place.
type RelatedEntity struct {
	ID   int64  `json:"id"`
	Type Enum   `json:"type"`
	Name string `json:"name"`
}

type countResponse struct {
	Count int64 `json:"count"`
}

func Pointer[T any](v T) *T {
	return &v
}

// queryValues encodes a struct with `url` tags as query values for Compile.
// A nil pointer gives no values.
func queryValues(opts any) (url.Values, error) {
	v := reflect.ValueOf(opts)
	if !v.IsValid() || (v.Kind() == reflect.Ptr && v.IsNil()) {
		return nil, nil
	}

	return query.Values(opts)
}
