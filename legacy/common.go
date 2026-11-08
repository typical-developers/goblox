package legacy

import (
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

type WrappedResponse[T any] struct {
	Data T `json:"data"`
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
