package http_client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/typical-developers/goblox/http_error"
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

func (c *HTTPClient) buildRequest(method string, path string, body interface{}) (*http.Request, error) {
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

func (c *HTTPClient) Do(method string, path string, reqBody interface{}, query map[string]string) (result *ResponseResult, err error) {
	req, err := c.buildRequest(method, path, reqBody)
	if err != nil {
		return nil, err
	}

	// Set headers
	if len(c.Headers) > 0 {
		for k, v := range c.Headers {
			req.Header.Set(k, v)
		}
	}

	if len(query) > 0 {
		q := req.URL.Query()
		for k, v := range query {
			q.Set(k, v)
		}

		req.URL.RawQuery = q.Encode()
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, http_error.NewHTTPResponseError(res, resBody)
	}

	return &ResponseResult{
		StatusCode: res.StatusCode,
		Body:       resBody,
	}, nil
}

type ResponseResult struct {
	StatusCode int
	Body       []byte
}

func DecodeResult[T interface{}](r *ResponseResult) (result *T, err error) {
	err = json.Unmarshal(r.Body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
