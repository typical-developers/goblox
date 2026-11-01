package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	ErrGrantInactive = errors.New("grant is inactive, expired, or disabled.")
)

type Grant struct {
	Active    bool
	Scopes    []string
	OwnerID   string // The user that owns the key/token.
	CheckedAt time.Time
	ExpiresAt time.Time // zero if unknown or never
}

type Introspector interface {
	Introspect(ctx context.Context) (*Grant, error)
}

type Introspected struct {
	Credential
	Introspector Introspector
	TTL          time.Duration // How often to re-check API keys. No value defaults to 10 minutes.

	grant    *Grant
	mu       sync.Mutex
	inflight chan struct{}
}

func (c *Introspected) GrantedScopes(ctx context.Context) ([]string, error) {
	g, err := c.Grant(ctx)
	if err != nil {
		return nil, err
	}
	return g.Scopes, nil
}

// isFreshGrant will check if the grant has been refreshed within the TTL.
func (c *Introspected) isFreshGrant(g *Grant, now time.Time) bool {
	ttl := c.TTL

	if ttl <= 0 {
		ttl = 10 * time.Minute
	}

	if now.Sub(g.CheckedAt) >= ttl {
		return false
	}

	if !g.ExpiresAt.IsZero() && !now.Before(g.ExpiresAt) {
		return false
	}

	return true
}

// checkActive will determine if the current grant is active.
func (c *Introspected) checkActive(g *Grant) (*Grant, error) {
	if !g.Active {
		return nil, ErrGrantInactive
	}

	return g, nil
}

// Grant returns the cached grant, introspecting if it's missing, stale or expired.
func (c *Introspected) Grant(ctx context.Context) (*Grant, error) {
	for {
		c.mu.Lock()

		if g := c.grant; g != nil && c.isFreshGrant(g, time.Now()) {
			c.mu.Unlock()
			return c.checkActive(g)
		}

		// Wait for an ongoing call to finish introspecting.
		if wait := c.inflight; wait != nil {
			c.mu.Unlock()

			select {
			case <-wait:
				continue
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}

		done := make(chan struct{})
		c.inflight = done
		c.mu.Unlock()

		g, err := c.Introspector.Introspect(ctx)

		c.mu.Lock()
		if err == nil && g != nil {
			if g.CheckedAt.IsZero() {
				g.CheckedAt = time.Now()
			}
			c.grant = g
		}
		c.inflight = nil
		close(done)
		c.mu.Unlock()

		if err != nil {
			return nil, err
		}
		if g == nil {
			return nil, errors.New("introspector returned no grant")
		}
		return c.checkActive(g)
	}
}

// Invalidate drops the cached grant so the next request will introspect.
func (c *Introspected) Invalidate() {
	c.mu.Lock()
	c.grant = nil
	c.mu.Unlock()
}

func (c *Introspected) HandleResponse(resp *http.Response) bool {
	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		c.Invalidate()
	}
	if h, ok := c.Credential.(ResponseHandler); ok {
		return h.HandleResponse(resp)
	}
	return false
}

func introspect(req *http.Request, out any) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status not ok, received status %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(out)
}
