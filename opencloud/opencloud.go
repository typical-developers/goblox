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

	userAgent string
	BaseURL   *url.URL

	retry        goblox.RetryPolicy
	retryEnabled bool

	Assets             *AssetsService
	Config             *ConfigService
	DataAndMemoryStore *DataAndMemoryStoreService
	LuauExecution      *LuauExecutionService
	Monetization       *MonetizationService
	Places             *PlacesService
	Universes          *UniversesService
	Users              *UsersService
	Groups             *GroupsService
	Ads                *AdsService
	Analytics          *AnalyticsService
	Badges             *BadgesService
	CreatorStore       *CreatorStoreService
	Events             *EventsService
	Localization       *LocalizationService
	Matchmaking        *MatchmakingService
	Secrets            *SecretsService
	ServerManagement   *ServerManagementService
	ServiceEfficiency  *ServiceEfficiencyService
	Thumbnails         *ThumbnailsService
}

type ClientOpts func(*Client)

// NewClient will create a completely new client.
// To authenticate with the Opencloud API, you can use the WithAPIKey or WithOAuthToken methods.
//
// WithAPIKey will use an API key that you created for your experience.
// WithOAuthToken should be used for OAuth applications that are athenticated.
func NewClient(opts ...ClientOpts) *Client {
	c := &Client{
		client:       http.DefaultClient,
		userAgent:    goblox.UserAgent,
		retry:        goblox.DefaultRetryPolicy,
		retryEnabled: true,
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
	c.Places = (*PlacesService)(&c.common)
	c.Universes = (*UniversesService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Groups = (*GroupsService)(&c.common)
	c.Ads = (*AdsService)(&c.common)
	c.Analytics = (*AnalyticsService)(&c.common)
	c.Badges = (*BadgesService)(&c.common)
	c.CreatorStore = (*CreatorStoreService)(&c.common)
	c.Events = (*EventsService)(&c.common)
	c.Localization = (*LocalizationService)(&c.common)
	c.Matchmaking = (*MatchmakingService)(&c.common)
	c.Secrets = (*SecretsService)(&c.common)
	c.ServerManagement = (*ServerManagementService)(&c.common)
	c.ServiceEfficiency = (*ServiceEfficiencyService)(&c.common)
	c.Thumbnails = (*ThumbnailsService)(&c.common)

	return c
}

// WithCredential will let you use a custom credential to authenticate with the Opencloud API.
// You can use [auth.APIKey] or [auth.OAuthToken] to create a credential.
func (c *Client) WithCredential(cred auth.Credential) *Client {
	nc := &Client{
		client:       c.client,
		credential:   cred,
		userAgent:    c.userAgent,
		retry:        c.retry,
		retryEnabled: c.retryEnabled,
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

// WithBaseURL will override with a custom base url.
func WithBaseURL(u *url.URL) ClientOpts {
	return func(c *Client) {
		c.BaseURL = u
	}
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

func (c *Client) config(opts []goblox.RequestOpt) goblox.RequestConfig {
	config := goblox.RequestConfig{Credential: c.credential}

	for _, opt := range opts {
		opt(&config)
	}

	return config
}

func (c *Client) send(ctx context.Context, req *http.Request, _ *goblox.Endpoint, config goblox.RequestConfig, out any) (*Response, error) {
	retryEnabled := c.retryEnabled
	if config.Retry != nil {
		retryEnabled = *config.Retry
	}

	handler, hasHandler := config.Credential.(auth.ResponseHandler)
	replayable := req.Body == nil || req.Body == http.NoBody || req.GetBody != nil

	attempt := 0
	handlerRetried := false
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
				if hook := c.retry.OnRetry; hook != nil {
					hook(attempt+1, delay, resp)
				}

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

// Call send a reqeust.
func (c *Client) Call(ctx context.Context, e *goblox.CompiledEndpoint, body, out any, opts ...goblox.RequestOpt) (*Response, error) {
	config := c.config(opts)

	if err := e.Endpoint.Check(config.Credential); err != nil {
		return nil, err
	}
	if err := e.Endpoint.CheckScopes(ctx, config.Credential); err != nil {
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
	if err := e.Endpoint.CheckScopes(ctx, config.Credential); err != nil {
		return nil, err
	}

	return c.send(ctx, req, e.Endpoint, config, out)
}

// NewRequest will build a new request.
func (c *Client) NewRequest(method, urlString string, body any) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlString)
	if err != nil {
		return nil, err
	}

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

// NewMultipartRequest will build a request that sends multipart data.
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

// Do will send a request.
func (c *Client) Do(ctx context.Context, req *http.Request, v any, opts ...goblox.RequestOpt) (*Response, error) {
	return c.send(ctx, req, nil, c.config(opts), v)
}

type Response struct {
	*http.Response
}

// Ok will check if the status code is 200-299.
func (r *Response) Ok() bool {
	return (r.StatusCode >= 200 && r.StatusCode <= 299)
}

// RateLimit will parse the rate limit from the headers.
func (r *Response) RateLimit() goblox.RateLimit {
	return goblox.ParseRateLimit(r.Header)
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
