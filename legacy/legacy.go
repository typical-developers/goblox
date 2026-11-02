package legacy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

type service struct {
	client *Client
}

type Client struct {
	client     *http.Client
	common     service
	credential auth.Credential

	baseDomain string

	Thumbnails *ThumbnailsService
}

type ClientOpts func(*Client)

func NewClient(opts ...ClientOpts) *Client {
	c := &Client{
		client:     http.DefaultClient,
		baseDomain: "roblox.com",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c.init()
}

func (c *Client) init() *Client {
	c.common.client = c

	c.Thumbnails = (*ThumbnailsService)(&c.common)

	return c
}

// WithCredential will let you use a custom credential to authenticate with the Opencloud API.
// You can use [auth.ROBLOSecurityCookie] to create a credential.
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

func (c *Client) resolveURL(e *goblox.CompiledEndpoint) (*url.URL, error) {
	host := e.Endpoint.Domain
	if host == "" {
		return nil, fmt.Errorf("goblox/legacy: %s %s has no domain", e.Endpoint.Method, e.Endpoint.Route)
	}
	if !strings.Contains(host, ".") {
		host += "." + c.baseDomain
	}

	base := &url.URL{Scheme: "https", Host: host}

	// Parse rather than setting Path, so the "?query" from Compile stays a query.
	return base.Parse(e.Path)
}

func (c *Client) config(opts []goblox.RequestOpt) goblox.RequestConfig {
	config := goblox.RequestConfig{Credential: c.credential}

	for _, opt := range opts {
		opt(&config)
	}

	return config
}

func (c *Client) NewRequest(method string, u *url.URL, body any) (*http.Request, error) {
	var reader io.Reader
	if body != nil {
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		if err := encoder.Encode(body); err != nil {
			return nil, err
		}
		reader = bytes.NewReader(buf.Bytes())
	}

	req, err := http.NewRequest(method, u.String(), reader)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) send(ctx context.Context, req *http.Request, config goblox.RequestConfig, out any) (*Response, error) {
	handler, canRetry := config.Credential.(auth.ResponseHandler)

	// Only retry when the body can be replayed.
	canRetry = canRetry && (req.Body == nil || req.Body == http.NoBody || req.GetBody != nil)

	var resp *http.Response
	for attempt := 0; attempt < 2; attempt++ {
		// A fresh copy each attempt, so headers and credentials aren't added twice.
		r := req.Clone(ctx)
		if attempt > 0 && req.GetBody != nil {
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			r.Body = body
		}

		for k, vs := range config.Header {
			for _, v := range vs {
				r.Header.Add(k, v)
			}
		}

		if config.Credential != nil {
			if err := config.Credential.Apply(r); err != nil {
				return nil, err
			}
		}

		var err error
		resp, err = c.client.Do(r)
		if err != nil {
			return nil, err
		}

		// The cookie picks up X-CSRF-Token from a 403 and asks for one retry.
		if attempt == 0 && canRetry && handler.HandleResponse(resp) {
			_, _ = io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			continue
		}
		break
	}
	defer resp.Body.Close()

	response := &Response{resp}
	if response.Ok() && out != nil && resp.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil && !errors.Is(err, io.EOF) {
			return response, err
		}
	}

	return response, nil
}

func (c *Client) Call(ctx context.Context, e *goblox.CompiledEndpoint, body, out any, opts ...goblox.RequestOpt) (*Response, error) {
	config := c.config(opts)

	if err := e.Endpoint.Check(config.Credential); err != nil {
		return nil, err
	}
	if err := e.Endpoint.CheckScopes(ctx, config.Credential); err != nil {
		return nil, err
	}

	u, err := c.resolveURL(e)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest(e.Endpoint.Method, u, body)
	if err != nil {
		return nil, err
	}

	return c.send(ctx, req, config, out)
}

type Response struct {
	*http.Response
}

// Ok will check if the status code is 200-299.
func (r *Response) Ok() bool {
	return r.StatusCode >= 200 && r.StatusCode <= 299
}
