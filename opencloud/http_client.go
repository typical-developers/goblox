package opencloud

import (
	"fmt"

	http_client "github.com/typical-developers/goblox/internal/http_client"
)

type Session struct {
	Client *http_client.HTTPClient
}

func NewSession() *Session {
	return &Session{
		Client: http_client.NewHTTPClient("apis.roblox.com"),
	}
}

func (s *Session) SetAPIToken(apiToken string) *Session {
	s.Client.SetHeader("x-api-key", apiToken)
	return s
}

func (s *Session) SetOAuthToken(apiToken string) *Session {
	s.Client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return s
}
