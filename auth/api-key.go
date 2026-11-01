package auth

import "net/http"

type APIKey string

func (c APIKey) Kind() Kind {
	return KindAPIKey
}

func (c APIKey) Apply(req *http.Request) error {
	req.Header.Set("X-API-Key", string(c))
	return nil
}
