package opencloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

var (
	baseURL = "https://apis.roblox.com"
)

type service struct {
	client *Client
}

type APIKeyRoundTripper struct {
	APIKey    string
	Transport http.RoundTripper
}

func (c *APIKeyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-KEY", c.APIKey)
	return c.Transport.RoundTrip(req)
}

type OAuthRoundTripper struct {
	OAuthToken string
	Transport  http.RoundTripper
}

func (c *OAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.OAuthToken))
	return c.Transport.RoundTrip(req)
}

type Client struct {
	client *http.Client
	common service

	BaseURL *url.URL

	// v2 Opencloud API services
	DataAndMemoryStore *DataAndMemoryStoreService
	LuauExecution      *LuauExecutionService
	Monetization       *MonetizationService
	UniverseAndPlaces  *UniverseAndPlacesService
	UserAndGroups      *UserAndGroupsService
}

type Response struct {
	*http.Response
}

// NewClientWithAPIKey will use an API key to authenticate with the Opencloud API.
//
// You can create a new API key at: https://create.roblox.com/dashboard/credentials?activeTab=ApiKeysTab
func NewClientWithAPIKey(apiKey string) *Client {
	c := &Client{
		client: &http.Client{
			Transport: &APIKeyRoundTripper{
				APIKey:    apiKey,
				Transport: http.DefaultTransport,
			},
		},
	}

	return c.init()
}

// TODO: Implement automatic OAuth token refreshing.
//
// NewClientWithOAuth will use an OAuth token to authenticate with the Opencloud API.
//
// The token must be from the user that authenticated.
func NewClientWithOAuth(token string) *Client {
	c := &Client{
		client: &http.Client{
			Transport: &OAuthRoundTripper{
				OAuthToken: token,
				Transport:  http.DefaultTransport,
			},
		},
	}

	return c.init()
}

func (c *Client) init() *Client {
	c.common.client = c

	c.BaseURL, _ = url.Parse(baseURL)

	c.DataAndMemoryStore = (*DataAndMemoryStoreService)(&c.common)
	c.LuauExecution = (*LuauExecutionService)(&c.common)
	c.Monetization = (*MonetizationService)(&c.common)
	c.UniverseAndPlaces = (*UniverseAndPlacesService)(&c.common)
	c.UserAndGroups = (*UserAndGroupsService)(&c.common)

	return c
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

func (c *Client) Do(ctx context.Context, req *http.Request, v any) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return nil, err
	}

	return response, err
}

type Options struct {
	MaxPageSize int    `url:"maxPageSize,omitempty"`
	PageToken   string `url:"pageToken,omitempty"`
}

type OptionsWithFilter struct {
	Options
	Filter string `url:"filter,omitempty"`
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

type OperationResponse map[string]any

type OperationMetadata map[string]any

type Operation struct {
	Path     string            `json:"path"`
	Done     bool              `json:"done"`
	Response OperationResponse `json:"response"`
	Metadata OperationMetadata `json:"metadata"`
}
