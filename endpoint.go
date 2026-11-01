package goblox

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/typical-developers/goblox/auth"
)

type Endpoint struct {
	Domain  string
	Route   string
	Method  string
	Accepts auth.Kinds
	Scopes  []string
}

func NewEndpoint(method, route string, accepts ...auth.Kind) *Endpoint {
	return &Endpoint{
		Method:  method,
		Route:   route,
		Accepts: auth.KindsOf(accepts...),
	}
}

// WithScopes records the scopes the endpoint needs.
func (e *Endpoint) WithScopes(scopes ...string) *Endpoint {
	e.Scopes = scopes
	return e
}

// Check will return an error if the credential isn't in Accepts.
// An endpoint with no accepts will not be checked.
func (e *Endpoint) Check(credential auth.Credential) error {
	if e.Accepts == 0 {
		return nil
	}

	kind := auth.KindNone
	if credential != nil {
		kind = credential.Kind()
	}

	if e.Accepts.Has(kind) {
		return nil
	}

	accepts := make([]string, len(e.Accepts.List()))
	for i, k := range e.Accepts.List() {
		accepts[i] = k.String()
	}

	return fmt.Errorf("goblox: %s %s doesn't accept %s auth (accepts: %s)",
		e.Method, e.Route, kind, strings.Join(accepts, ","),
	)
}

type CompiledEndpoint struct {
	Endpoint *Endpoint
	Path     string
}

// Compile will fill {placeholders} in order from params and append any query valoues.
func (e *Endpoint) Compile(query url.Values, params ...any) *CompiledEndpoint {
	path := e.Route

	for _, p := range params {
		start := strings.Index(path, "{")
		end := strings.Index(path, "}")

		if start == -1 || end < start {
			break
		}

		path = path[:start] + url.PathEscape(fmt.Sprint(p)) + path[end+1:]
	}

	if q := query.Encode(); q != "" {
		path += "?" + q
	}

	return &CompiledEndpoint{
		Endpoint: e,
		Path:     path,
	}
}
