package opencloud

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
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

	c.BaseURL, _ = url.Parse(baseURL)

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
func (c *Client) WithCredential(cred auth.Credential) *Client {
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
	return c.WithCredential(auth.APIKey(apiKey))
}

// WithOAuthToken will use an OAuth token to authenticate with the Opencloud API.
//
// You can create a new OAuth client at: https://create.roblox.com/dashboard/credentials?activeTab=OAuthTab
func (c *Client) WithOAuthToken(token string) *Client {
	return c.WithCredential(auth.OAuthToken(token))
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

func addOpts(urlString string, opts any) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return urlString, nil
	}

	u, err := url.Parse(urlString)
	if err != nil {
		return urlString, err
	}

	q, err := query.Values(opts)
	if err != nil {
		return urlString, err
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
