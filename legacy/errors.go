package legacy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/luckfire/goblox"
)

type ErrorDetail struct {
	Code              int    `json:"code"`
	Message           string `json:"message"`
	UserFacingMessage string `json:"userFacingMessage,omitempty"`
	Field             string `json:"field,omitempty"`
}

type Error struct {
	goblox.ResponseError
	Errors []ErrorDetail
}

func (e *Error) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "goblox/legacy: HTTP %d", e.StatusCode)

	for _, d := range e.Errors {
		fmt.Fprintf(&b, "; %d: %s", d.Code, d.Message)
	}

	return b.String()
}

// Has reports whether Roblox returned the given error code. Codes are per
// endpoint, so 3 means one thing on groups and another on users.
func (e *Error) Has(code int) bool {
	for _, d := range e.Errors {
		if d.Code == code {
			return true
		}
	}

	return false
}

func decodeError(resp *http.Response, body []byte) error {
	e := &Error{ResponseError: goblox.NewResponseError(resp, body)}
	var raw struct {
		Errors []ErrorDetail `json:"errors"`
	}
	if json.Unmarshal(body, &raw) == nil {
		e.Errors = raw.Errors
	}
	return e
}
