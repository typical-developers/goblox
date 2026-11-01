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
func (c *Client) withCredential(cred auth.Credential) *Client {
	nc := &Client{
		client:     c.client,
		credential: cred,
	}

	return nc.init()
}

// WithCookie will use a .ROBLOSECURITY cookie to authenticate with APIs.
func (c *Client) WithCookie(cookie string) *Client {
	return c.withCredential(auth.NewROBLOSecurityCookie(cookie))
}

func (c *Client) resolveURL(r route) (*string, error) {
	domainURL := url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("https://%s.roblox.com", r.Domain),
		Path:   r.Path,
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
