package opencloud

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	baseURL = "https://apis.roblox.com"
)

type service struct {
	client *Client
}

type Client struct {
	client     *http.Client
	common     service
	credential auth.Credential

	BaseURL *url.URL

	Assets             *AssetsService
	Config             *ConfigService
	DataAndMemoryStore *DataAndMemoryStoreService
	LuauExecution      *LuauExecutionService
	Monetization       *MonetizationService
	Places             *PlacesService
	Universes          *UniversesService
	Users              *UsersService
	Groups             *GroupsService

	// Deprecated: Use [UsersService] and [GroupsService] instead.
	UserAndGroups *UserAndGroupsService
	// Deprecated: Use [UniversesService] and [PlacesService].
	UniverseAndPlaces *UniverseAndPlacesService
}

type ClientOpts func(*Client)

// NewClient will create a completely new client.
// To authenticate with the Opencloud API, you can use the WithAPIKey or WithOAuthToken methods.
//
// WithAPIKey will use an API key that you created for your experience.
// WithOAuthToken should be used for OAuth applications that are athenticated.
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

	if c.BaseURL == nil {
		c.BaseURL, _ = url.Parse(baseURL)
	}

	// v1
	c.Assets = (*AssetsService)(&c.common)

	// v2
	c.Config = (*ConfigService)(&c.common)
	c.DataAndMemoryStore = (*DataAndMemoryStoreService)(&c.common)
	c.LuauExecution = (*LuauExecutionService)(&c.common)
	c.Monetization = (*MonetizationService)(&c.common)
	c.UniverseAndPlaces = (*UniverseAndPlacesService)(&c.common)
	c.Places = (*PlacesService)(&c.common)
	c.Universes = (*UniversesService)(&c.common)
	c.UserAndGroups = (*UserAndGroupsService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Groups = (*GroupsService)(&c.common)

	return c
}

// WithCredential will let you use a custom credential to authenticate with the Opencloud API.
// You can use [auth.APIKey] or [auth.OAuthToken] to create a credential.
func (c *Client) withCredential(cred auth.Credential) *Client {
	nc := &Client{
		client:     c.client,
		credential: cred,
		BaseURL:    c.BaseURL,
	}

	return nc.init()
}

// WithAPIKey will use an API key to authenticate with the Opencloud API.
//
// You can create a new API key at: https://create.roblox.com/dashboard/credentials?activeTab=ApiKeysTab
func (c *Client) WithAPIKey(apiKey string) *Client {
	return c.withCredential(auth.APIKey(apiKey))
}

// WithOAuthToken will use an OAuth token to authenticate with the Opencloud API.
//
// You can create a new OAuth client at: https://create.roblox.com/dashboard/credentials?activeTab=OAuthTab
func (c *Client) WithOAuthToken(token string) *Client {
	return c.withCredential(auth.OAuthToken(token))
}

// WithBaseURL will override with a custom base url.
func WithBaseURL(u *url.URL) ClientOpts {
	return func(c *Client) {
		c.BaseURL = u
	}
}

func (c *Client) config(opts []goblox.RequestOpt) goblox.RequestConfig {
	config := goblox.RequestConfig{Credential: c.credential}

	for _, opt := range opts {
		opt(&config)
	}

	return config
}

func (c *Client) send(ctx context.Context, req *http.Request, _ *goblox.Endpoint, config goblox.RequestConfig, out any) (*Response, error) {
	req = req.WithContext(ctx)

	for k, vs := range config.Header {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}

	if config.Credential != nil {
		if err := config.Credential.Apply(req); err != nil {
			return nil, err
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{resp}
	if response.Ok() {
		if out != nil && response.StatusCode != http.StatusNoContent {
			if err := json.NewDecoder(resp.Body).Decode(out); err != nil && !errors.Is(err, io.EOF) {
				return response, err
			}
		}
	}

	return response, nil
}

func (c *Client) Call(ctx context.Context, e *goblox.CompiledEndpoint, body, out any, opts ...goblox.RequestOpt) (*Response, error) {
	config := c.config(opts)

	if err := e.Endpoint.Check(config.Credential); err != nil {
		return nil, err
	}

	req, err := c.NewRequest(e.Endpoint.Method, e.Path, body)
	if err != nil {
		return nil, err
	}

	return c.send(ctx, req, e.Endpoint, config, out)
}

// CallRequest sends a request built by hand (a multipart upload, for example)
// for endpoint e. Like Call, it checks the credential against the endpoint first.
func (c *Client) CallRequest(ctx context.Context, e *goblox.CompiledEndpoint, req *http.Request, out any, opts ...goblox.RequestOpt) (*Response, error) {
	config := c.config(opts)

	if err := e.Endpoint.Check(config.Credential); err != nil {
		return nil, err
	}

	return c.send(ctx, req, e.Endpoint, config, out)
}

func (c *Client) NewRequest(method, urlString string, body any) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlString)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
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

func (c *Client) NewMultipartRequest(method, urlString string, body *bytes.Buffer, contentType string) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlString)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	return req, nil
}

type Response struct {
	*http.Response
}

// Ok will check if the status code is 200-299.
func (r *Response) Ok() bool {
	return (r.StatusCode >= 200 && r.StatusCode <= 299)
}

func (c *Client) Do(ctx context.Context, req *http.Request, v any, opts ...goblox.RequestOpt) (*Response, error) {
	return c.send(ctx, req, nil, c.config(opts), v)
}

// queryValues encodes a struct with `url` tags as query values for Compile.
// A nil pointer gives no values.
func queryValues(opts any) (url.Values, error) {
	v := reflect.ValueOf(opts)
	if !v.IsValid() || (v.Kind() == reflect.Ptr && v.IsNil()) {
		return nil, nil
	}

	return query.Values(opts)
}
