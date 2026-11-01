package goblox

import (
	"net/http"

	"github.com/typical-developers/goblox/auth"
)

// RequestConfig is what RequestOpts change. Fields are read by Client.Do.
type RequestConfig struct {
	Credential    auth.Credential
	HasCredential bool // an option set Credential, possibly to nil
	Header        http.Header
}

// RequestOpt changes a single request.
type RequestOpt func(*RequestConfig)

// WithCredential uses cred for this request instead of the client's default.
func WithCredential(cred auth.Credential) RequestOpt {
	return func(c *RequestConfig) {
		c.Credential = cred
		c.HasCredential = true
	}
}

// WithAPIKey uses an API key for this request.
func WithAPIKey(key string) RequestOpt {
	return WithCredential(auth.APIKey(key))
}

// WithOAuthToken uses a user's OAuth access token for this request.
func WithOAuthToken(token string) RequestOpt {
	return WithCredential(auth.OAuthToken(token))
}

// WithoutCredential sends this request with no credential.
func WithoutCredential() RequestOpt {
	return WithCredential(nil)
}

// WithHeader adds a header to this request.
func WithHeader(key, value string) RequestOpt {
	return func(c *RequestConfig) {
		if c.Header == nil {
			c.Header = http.Header{}
		}

		c.Header.Add(key, value)
	}
}
