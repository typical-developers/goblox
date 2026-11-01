package goblox

import (
	"context"
	"errors"
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
	// Scopes are all required.
	Scopes []string
	// AnyScopes need at least one present, for endpoints whose docs list
	// alternatives (a read scope or a broader write scope, for example).
	AnyScopes []string
}

func NewEndpoint(method, route string, accepts ...auth.Kind) *Endpoint {
	return &Endpoint{
		Method:  method,
		Route:   route,
		Accepts: auth.KindsOf(accepts...),
	}
}

// WithScopes records scopes that are all required.
func (e *Endpoint) WithScopes(scopes ...string) *Endpoint {
	e.Scopes = scopes
	return e
}

// WithAnyScopes records scopes of which at least one is required.
func (e *Endpoint) WithAnyScopes(scopes ...string) *Endpoint {
	e.AnyScopes = scopes
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

	if !e.Accepts.Has(kind) {
		return &UnsupportedAuthError{
			Method:  e.Method,
			Route:   e.Route,
			Have:    kind,
			Accepts: e.Accepts.List(),
		}
	}

	return nil
}

func (e *Endpoint) CheckScopes(ctx context.Context, credential auth.Credential) error {
	if (len(e.Scopes) == 0 && len(e.AnyScopes) == 0) || credential == nil {
		return nil
	}

	scoped, ok := credential.(auth.Scoped)
	if !ok {
		return nil
	}

	granted, err := scoped.GrantedScopes(ctx)
	if err != nil {
		if errors.Is(err, auth.ErrGrantInactive) || ctx.Err() != nil {
			return err
		}
		return nil
	}
	if granted == nil {
		return nil
	}

	have := make(map[string]bool, len(granted))
	for _, s := range granted {
		have[s] = true
	}

	var missing []string
	for _, s := range e.Scopes {
		if !have[s] {
			missing = append(missing, s)
		}
	}
	var anyOf []string
	if len(e.AnyScopes) > 0 {
		found := false
		for _, s := range e.AnyScopes {
			if have[s] {
				found = true
				break
			}
		}
		if !found {
			anyOf = e.AnyScopes
		}
	}

	if len(missing) > 0 || len(anyOf) > 0 {
		return &MissingScopesError{
			Method:  e.Method,
			Route:   e.Route,
			Missing: missing,
			AnyOf:   anyOf,
		}
	}

	return nil
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
