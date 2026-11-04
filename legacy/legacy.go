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
	userAgent  string

	retry        goblox.RetryPolicy
	retryEnabled bool

	Thumbnails *ThumbnailsService
}

type ClientOpts func(*Client)

func NewClient(opts ...ClientOpts) *Client {
	c := &Client{
		client:     http.DefaultClient,
		baseDomain: "roblox.com",
		userAgent:  goblox.UserAgent,
		retry: goblox.RetryPolicy{
			MaxRetries: 3,
		},
		retryEnabled: true,
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

// withCredential will let you use a custom credential to authenticate with the Opencloud API.
func (c *Client) withCredential(cred auth.Credential) *Client {
	nc := &Client{
		client:       c.client,
		credential:   cred,
		baseDomain:   c.baseDomain,
		userAgent:    c.userAgent,
		retry:        c.retry,
		retryEnabled: c.retryEnabled,
	}

	return nc.init()
}

// WithCookie will use a .ROBLOSECURITY cookie to authenticate with APIs.
func (c *Client) WithCookie(cookie string) *Client {
	return c.withCredential(auth.NewROBLOSecurityCookie(cookie))
}

// WithUserAgentApp will let you set the app for the client.
func WithUserAgentApp(app string) ClientOpts {
	return func(c *Client) {
		c.userAgent = goblox.AppUserAgent(app)
	}
}

// WithRetryPolicy sets how this client retries failed requests.
func WithRetryPolicy(p goblox.RetryPolicy) ClientOpts {
	return func(c *Client) {
		c.retry = p
	}
}

// WithRetries sets whether requests are retried by default.
func WithRetries(enabled bool) ClientOpts {
	return func(c *Client) {
		c.retryEnabled = enabled
	}
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
	retryEnabled := c.retryEnabled
	if config.Retry != nil {
		retryEnabled = *config.Retry
	}

	handler, hasHandler := config.Credential.(auth.ResponseHandler)
	replayable := req.Body == nil || req.Body == http.NoBody || req.GetBody != nil
	handlerRetried := false

	attempt := 0
	for {
		r := req.Clone(ctx)
		if req.GetBody != nil {
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			r.Body = body
		}

		r.Header.Set("User-Agent", c.userAgent)
		for k, vs := range config.Header {
			r.Header.Del(k)
			for _, v := range vs {
				r.Header.Add(k, v)
			}
		}

		if config.Credential != nil {
			if err := config.Credential.Apply(r); err != nil {
				return nil, err
			}
		}

		resp, err := c.client.Do(r)

		// The credential's own retry (a cookie's new CSRF token). Does not count towards
		// the retry policy.
		if err == nil && hasHandler && replayable && !handlerRetried && handler.HandleResponse(resp) {
			handlerRetried = true
			_, _ = io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			continue
		}

		// Runs the retry policy if retries are enabled.
		if retryEnabled {
			if delay, retry := c.retry.Next(attempt, r, resp, err); retry {
				if resp != nil {
					_, _ = io.Copy(io.Discard, resp.Body)
					resp.Body.Close()
				}
				if err := goblox.Sleep(ctx, delay); err != nil {
					return nil, err
				}
				attempt++
				continue
			}
		}

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		response := &Response{resp}

		if !response.Ok() {
			body, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
			return response, decodeError(resp, body)
		}

		if out != nil && response.StatusCode != http.StatusNoContent {
			if err := json.NewDecoder(resp.Body).Decode(out); err != nil && !errors.Is(err, io.EOF) {
				return response, err
			}
		}

		return response, nil
	}
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

// RateLimit will parse the rate limit from the headers.
func (r *Response) RateLimit() goblox.RateLimit {
	return goblox.ParseRateLimit(r.Header)
}
