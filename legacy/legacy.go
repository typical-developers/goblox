package legacy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/typical-developers/goblox/auth"
)

type api struct {
	Domain     string
	LegacyPath *string
}

func (a api) Path(format string, args ...any) route {
	return route{
		api:  a,
		Path: fmt.Sprintf(format, args...),
	}
}

type route struct {
	api

	Path string
	// Indicates if the endpoint is mirrored on apis.roblox.com.
	// If it is, this endpoint will support API Key / OAuth2.
	HasLegacy bool
}

func (r route) WithLegacy() route {
	r.HasLegacy = true
	return r
}

type service struct {
	client *Client
}

type Client struct {
	client     *http.Client
	common     service
	credential auth.Credential
}

type ClientOpts func(*Client)

func NewClient(opts ...ClientOpts) *Client {
	c := &Client{
		client: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c.init()
}

func (c *Client) init() *Client {
	c.common.client = c

	return c
}

// WithCredential will let you use a custom credential to authenticate with the Opencloud API.
// You can use [auth.APIKey], [auth.OAuthToken], or [auth.ROBLOSecurityCookie] to create a credential.
func (c *Client) WithCredential(cred auth.Credential) *Client {
	nc := &Client{
		client:     c.client,
		credential: cred,
	}

	return nc.init()
}

// WithAPIKey will use an API key to authenticate with the Opencloud API.
//
// You can create a new API key at: https://create.roblox.com/dashboard/credentials?activeTab=ApiKeysTab
func (c *Client) WithAPIKey(apiKey string) *Client {
	return c.WithCredential(auth.APIKey(apiKey))
}

// WithOAuthToken will use an OAuth token to authenticate with the Opencloud API.
//
// You can create a new OAuth client at: https://create.roblox.com/dashboard/credentials?activeTab=OAuthTab
func (c *Client) WithOAuthToken(token string) *Client {
	return c.WithCredential(auth.OAuthToken(token))
}

// WithCookie will use a .ROBLOSECURITY cookie to authenticate with APIs.
func (c *Client) WithCookie(cookie string) *Client {
	return c.WithCredential(auth.NewROBLOSecurityCookie(cookie))
}

func (c *Client) resolveURL(r route) (*string, error) {
	domainURL := url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("https://%s.roblox.com", r.Domain),
		Path:   r.Path,
	}

	authType := auth.KindNone
	if c.credential != nil {
		authType = c.credential.Kind()
	}

	switch authType {
	// Some of the legacy endpoitns are mirrored on apis.roblox.com
	// and will allow the usage of an API key or OAuth2 token.
	case auth.KindAPIKey, auth.KindOAuth2:
		if r.HasLegacy {
			if r.LegacyPath == nil {
				return nil, fmt.Errorf("goblox legacy: %s is marked mirrored but %q has no mirror prefix", r.Path, r.Domain)
			}

			domainURL = url.URL{
				Scheme: "https",
				Host:   "https://apis.roblox.com",
				Path:   fmt.Sprintf("/%s%s", *r.LegacyPath, r.Path),
			}

			break
		}

		return nil, fmt.Errorf("goblox legacy: %s is not mirrored on apis.roblox.com but an API Key or OAuth2 credential was provided", r.Path)
	}

	return Pointer(domainURL.String()), nil
}

func (c *Client) NewRequest(method string, r route, body any) (*http.Request, error) {
	uStr, err := c.resolveURL(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, *uStr, nil)
	if err != nil {
		return nil, err
	}

	if body != nil {
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(body)
		if err != nil {
			return nil, err
		}

		req.Body = io.NopCloser(&buf)
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

type Response struct {
	*http.Response
}

func (c *Client) Do(ctx context.Context, req *http.Request, v any) (*Response, error) {
	req = req.WithContext(ctx)
	if c.credential != nil {
		if err := c.credential.Apply(req); err != nil {
			return nil, err
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	return response, err
}
