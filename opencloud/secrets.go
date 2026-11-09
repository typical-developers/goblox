package opencloud

import (
	"context"
	"net/http"

	"github.com/luckfire/goblox"
	"github.com/luckfire/goblox/auth"
)

var (
	ListSecrets = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/secrets",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.secret:read")

	CreateSecret = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/secrets",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.secret:write")

	GetSecretsPublicKey = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/secrets/public-key",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.secret:read")

	DeleteSecret = goblox.NewEndpoint(
		http.MethodDelete,
		"/cloud/v2/universes/{universeId}/secrets/{secretId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.secret:write")

	UpdateSecret = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/universes/{universeId}/secrets/{secretId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.secret:write")
)

// SecretsService handles communication with the secrets store endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type SecretsService service

type ListSecretsOptions struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type Secret struct {
	ID         *string `json:"id,omitempty"`
	Secret     *string `json:"secret,omitempty"`
	KeyID      *string `json:"key_id,omitempty"`
	Domain     *string `json:"domain,omitempty"`
	CreateTime *string `json:"create_time,omitempty"`
	UpdateTime *string `json:"update_time,omitempty"`
}

type SecretPaginatedList struct {
	Secrets            []Secret `json:"secrets"`
	NextPageCursor     string   `json:"nextPageCursor"`
	PreviousPageCursor string   `json:"previousPageCursor"`
}

// ListSecrets will list all secrets defined for a universe.
//
// Required scopes: universe.secret:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/universes#get_cloud_v2_universes__universeId__secrets
//
// [GET] /cloud/v2/universes/{universeId}/secrets
func (s *SecretsService) ListSecrets(ctx context.Context, universeId string, opts *ListSecretsOptions, reqOpts ...goblox.RequestOpt) (*SecretPaginatedList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	secretPaginatedList := new(SecretPaginatedList)
	resp, err := s.client.Call(ctx, ListSecrets.Compile(q, universeId), nil, secretPaginatedList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return secretPaginatedList, resp, nil
}

// CreateSecret will create a new secret.
//
// Required scopes: universe.secret:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/universes#post_cloud_v2_universes__universeId__secrets
//
// [POST] /cloud/v2/universes/{universeId}/secrets
func (s *SecretsService) CreateSecret(ctx context.Context, universeId string, data Secret, reqOpts ...goblox.RequestOpt) (*Secret, *Response, error) {
	secret := new(Secret)
	resp, err := s.client.Call(ctx, CreateSecret.Compile(nil, universeId), data, secret, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return secret, resp, nil
}

// GetSecretsPublicKey will retrieve the public key for a universe.
//
// Required scopes: universe.secret:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/universes#get_cloud_v2_universes__universeId__secrets_public_key
//
// [GET] /cloud/v2/universes/{universeId}/secrets/public-key
func (s *SecretsService) GetSecretsPublicKey(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*Secret, *Response, error) {
	secret := new(Secret)
	resp, err := s.client.Call(ctx, GetSecretsPublicKey.Compile(nil, universeId), nil, secret, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return secret, resp, nil
}

// DeleteSecret will permanently delete a secret from a universe.
//
// Required scopes: universe.secret:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/universes#delete_cloud_v2_universes__universeId__secrets__secretId_
//
// [DELETE] /cloud/v2/universes/{universeId}/secrets/{secretId}
func (s *SecretsService) DeleteSecret(ctx context.Context, universeId, secretId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, DeleteSecret.Compile(nil, universeId, secretId), nil, nil, reqOpts...)
}

// UpdateSecret will update an existing secret.
//
// Required scopes: universe.secret:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/universes#patch_cloud_v2_universes__universeId__secrets__secretId_
//
// [PATCH] /cloud/v2/universes/{universeId}/secrets/{secretId}
func (s *SecretsService) UpdateSecret(ctx context.Context, universeId, secretId string, data Secret, reqOpts ...goblox.RequestOpt) (*Secret, *Response, error) {
	secret := new(Secret)
	resp, err := s.client.Call(ctx, UpdateSecret.Compile(nil, universeId, secretId), data, secret, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return secret, resp, nil
}
