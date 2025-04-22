package http_client

import "net/http"

type HTTPClient struct {
	BaseURL string
	Client  *http.Client
	Headers map[string]string
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		BaseURL: baseURL,
		Client:  http.DefaultClient,
		Headers: map[string]string{},
	}
}

func (c *HTTPClient) SetHeader(key string, value string) {
	c.Headers[key] = value
}

func (c *HTTPClient) RemoveHeader(key string) {
	delete(c.Headers, key)
}
