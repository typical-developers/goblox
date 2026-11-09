package goblox

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/luckfire/goblox/auth"
)

type MissingScopesError struct {
	Method  string
	Route   string
	Missing []string // required scopes the credential doesn't have
	AnyOf   []string // set when none of the endpoint's AnyScopes are present
}

func (e *MissingScopesError) Error() string {
	var parts []string
	if len(e.Missing) > 0 {
		parts = append(parts, "missing scopes: "+strings.Join(e.Missing, ", "))
	}
	if len(e.AnyOf) > 0 {
		parts = append(parts, "needs one of: "+strings.Join(e.AnyOf, ", "))
	}

	return fmt.Sprintf("goblox: %s %s %s", e.Method, e.Route, strings.Join(parts, "; "))
}

type UnsupportedAuthError struct {
	Method  string
	Route   string
	Have    auth.Kind   // the current auth kind
	Accepts []auth.Kind // the types of auth that the route accepts
}

func (e *UnsupportedAuthError) Error() string {
	accepts := make([]string, len(e.Accepts))
	for i, k := range e.Accepts {
		accepts[i] = k.String()
	}

	return fmt.Sprintf("goblox: %s %s doesn't accept %s auth (accepts: %s)",
		e.Method, e.Route, e.Have, strings.Join(accepts, ", "),
	)
}

type ResponseError struct {
	Response   *http.Response
	StatusCode int
	Body       []byte
	RetryAfter time.Duration
}

func parseRetryAfter(v string) time.Duration {
	if v == "" {
		return 0
	}

	if secs, err := strconv.Atoi(v); err == nil {
		return time.Duration(secs) * time.Second
	}
	if t, err := http.ParseTime(v); err == nil {
		return time.Until(t)
	}

	return 0
}

func NewResponseError(resp *http.Response, body []byte) ResponseError {
	return ResponseError{
		Response:   resp,
		StatusCode: resp.StatusCode,
		Body:       body,
		RetryAfter: parseRetryAfter(resp.Header.Get("Retry-After")),
	}
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("goblox: HTTP %d", e.StatusCode)
}

func (e *ResponseError) Retryable() bool {
	return e.StatusCode == http.StatusTooManyRequests || e.StatusCode >= 500
}
