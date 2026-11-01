package auth

import (
	"net/http"
	"sync"
)

type ROBLOSecurityCookie struct {
	mu sync.Mutex

	Value     string
	CSRFToken string
}

func needsCSRF(req *http.Request) bool {
	switch req.Method {
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
		return true
	default:
		return false
	}
}

func NewROBLOSecurityCookie(value string) *ROBLOSecurityCookie {
	return &ROBLOSecurityCookie{
		Value: value,
	}
}

func (c *ROBLOSecurityCookie) HandleResponse(resp *http.Response) bool {
	if resp.StatusCode != http.StatusForbidden {
		return false
	}

	token := resp.Header.Get("X-CSRF-Token")
	if token != "" || token == c.CSRFToken {
		return false
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if token == c.CSRFToken {
		return false
	}
	c.CSRFToken = token

	return true
}

func (c *ROBLOSecurityCookie) Kind() Kind {
	return KindCookie
}

func (c *ROBLOSecurityCookie) Apply(req *http.Request) error {
	req.AddCookie(&http.Cookie{Name: ".ROBLOSECURITY", Value: c.Value})

	if !needsCSRF(req) {
		return nil
	}

	c.mu.Lock()
	csrfToken := c.CSRFToken
	c.mu.Unlock()

	if csrfToken != "" {
		req.Header.Set("X-CSRF-Token", csrfToken)
	}
	return nil
}
