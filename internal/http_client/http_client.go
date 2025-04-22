package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type HTTPClient struct {
	BaseURL string
	Client  *http.Client
	Headers map[string]string
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		Headers: map[string]string{},
	}
}

func (c *HTTPClient) SetHeader(key string, value string) {
	c.Headers[key] = value
}

func (c *HTTPClient) RemoveHeader(key string) {
	delete(c.Headers, key)
}

func (c *HTTPClient) BuildRequest(method string, path string, body interface{}) (*http.Request, error) {
	url := url.URL{
		Scheme: "https",
		Host:   c.BaseURL,
		Path:   path,
	}

	if body != nil {
		json, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(method, url.String(), bytes.NewReader(json))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")

		return req, nil
	}

	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *HTTPClient) Request(req *http.Request) (result *ResponseResult, err error) {
	// Set headers
	if len(c.Headers) > 0 {
		for k, v := range c.Headers {
			req.Header.Set(k, v)
		}
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		// TODO: Better error handling later.
		return nil, fmt.Errorf("HTTP error: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &ResponseResult{
		StatusCode: res.StatusCode,
		Body:       body,
	}, nil
}

type ResponseResult struct {
	StatusCode int
	Body       []byte
}

func (r *ResponseResult) DecodeResult(result interface{}) error {
	return json.Unmarshal(r.Body, &result)
}
