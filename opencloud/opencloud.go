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
)

var (
	baseURL = "https://apis.roblox.com"
)

type service struct {
	client *Client
}

type Client struct {
	client *http.Client
	common service

	BaseURL *url.URL
	APIKey  string

	// Different OpenCloud API services.
	LuauExecution     *LuauExecutionService
	UniverseAndPlaces *UniverseAndPlacesService
	UserAndGroups     *UserAndGroupsService
}

type Response struct {
	*http.Response
}

// NewClient will return a new Opencloud API client.
// Authorizing with a key is mandatory; the API cannot be accessed without authentication.
func NewClient(apiKey string) *Client {
	c := &Client{client: http.DefaultClient}

	c.common.client = c

	c.BaseURL, _ = url.Parse(baseURL)
	c.APIKey = apiKey

	c.LuauExecution = (*LuauExecutionService)(&c.common)
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

	req.Header.Set("X-API-KEY", c.APIKey)

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
