package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type APIKey string

func (c APIKey) Kind() Kind {
	return KindAPIKey
}

func (c APIKey) Apply(req *http.Request) error {
	req.Header.Set("X-API-Key", string(c))
	return nil
}

type APIKeyIntrospector struct {
	Key APIKey
}

func IntrospectedAPIKey(key string, ttl time.Duration) *Introspected {
	k := APIKey(key)
	return &Introspected{
		Credential: k,
		Introspector: APIKeyIntrospector{
			Key: k,
		},
		TTL: ttl,
	}
}

type apiKeyScope struct {
	Name       string   `json:"name"`
	Operations []string `json:"operations"`
}

type apiKeyIntrospect struct {
	AuthorizedUserID  int64         `json:"authorizedUserId"`
	Enabled           bool          `json:"enabled"`
	Expired           bool          `json:"expired"`
	ExpirationTimeUTC string        `json:"expirationTimeUtc"`
	Scopes            []apiKeyScope `json:"scopes"`
}

type apiKeyIntrospectBody struct {
	APIKey string `json:"apiKey"`
}

func (i APIKeyIntrospector) Introspect(ctx context.Context) (*Grant, error) {
	body, err := json.Marshal(&apiKeyIntrospectBody{
		APIKey: string(i.Key),
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost, "https://apis.roblox.com/api-keys/v1/introspect",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	apiKey := new(apiKeyIntrospect)
	if err := introspect(req, &apiKey); err != nil {
		return nil, err
	}

	now := time.Now()
	grant := &Grant{
		Active:    apiKey.Enabled && !apiKey.Expired,
		CheckedAt: now,
	}
	if apiKey.AuthorizedUserID != 0 {
		grant.OwnerID = strconv.FormatInt(apiKey.AuthorizedUserID, 10)
	}
	if t, err := time.Parse(time.RFC3339, apiKey.ExpirationTimeUTC); err == nil && t.After(now) {
		grant.ExpiresAt = t
	}
	for _, s := range apiKey.Scopes {
		for _, op := range s.Operations {
			grant.Scopes = append(grant.Scopes, s.Name+":"+op)
		}
	}

	return grant, nil
}
