package opencloud

import (
	"context"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	GetConfigWithoutMetadata = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:read")

	GetConfigDraft = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:read")

	PartialUpdateConfigDraft = goblox.NewEndpoint(
		http.MethodPatch,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:write")

	DeleteConfigDraft = goblox.NewEndpoint(
		http.MethodDelete,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:write")

	UpdateConfigDraft = goblox.NewEndpoint(
		http.MethodPut,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft:overwrite",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:write")

	GetConfig = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/full",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:read")

	PublishConfigDraft = goblox.NewEndpoint(
		http.MethodPost,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/publish",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:write")

	ListRevisionHistory = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/revisions",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:read")

	RevisionRestore = goblox.NewEndpoint(
		http.MethodPost,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/revisions/{revisionId}/restore",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:write")
)

// ConfigService will handle communication with the actions related to the API.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type ConfigService service

type ConfigMetadata struct {
	ConfigVersion int `json:"configVersion"`
}

type Config struct {
	Metadata ConfigMetadata `json:"metadata"`
	Entries  map[string]any `json:"entries"`
}

type ConfigEntry struct {
	Value            any    `json:"value"`
	Description      string `json:"description"`
	LastModifiedTime string `json:"lastModifiedTime"`
	LastAccessedTime string `json:"lastAccessedTime"`
}

type FullConfig struct {
	Metadata ConfigMetadata         `json:"metadata"`
	Entries  map[string]ConfigEntry `json:"entries"`
}

// GetConfigWithoutMetadata will get the currently published config without metadata and decorators.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_GetConfigRepositoryValues
//
// Required scopes: universe:read
//
// [GET] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}
func (s *ConfigService) GetConfigWithoutMetadata(ctx context.Context, universeId, repository string, reqOpts ...goblox.RequestOpt) (*Config, *Response, error) {
	config := new(Config)

	resp, err := s.client.Call(ctx, GetConfigWithoutMetadata.Compile(nil, universeId, repository), nil, config, reqOpts...)
	if err != nil {
		return nil, resp, err
	}
	return config, resp, nil
}

type ConfigDraftEntry struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}

type ConfigDraft struct {
	DraftHash string                      `json:"draftHash"`
	Entries   map[string]ConfigDraftEntry `json:"entries"`
}

// GetConfigDraft will return the current draft changes for the config.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_GetConfigRepositoryDraft
//
// Required scopes: universe:read
//
// [GET] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft
func (s *ConfigService) GetConfigDraft(ctx context.Context, universeId, repository string, reqOpts ...goblox.RequestOpt) (*ConfigDraft, *Response, error) {
	draft := new(ConfigDraft)
	resp, err := s.client.Call(ctx, GetConfigDraft.Compile(nil, universeId, repository), nil, draft, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return draft, resp, nil
}

type ConfigDraftUpdate struct {
	DraftHash string         `json:"draftHash"`
	Entries   map[string]any `json:"entries"`
}

type ConfigDraftUpdated struct {
	DraftHash string `json:"draftHash"`
}

// PartialUpdateConfigDraft will partially update the draft changes for the config.
// A key that is not included will not be changed.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_UpdateDraft
//
// Required scopes: universe:write
//
// [PATCH] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft
func (s *ConfigService) PartialUpdateConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftUpdate, reqOpts ...goblox.RequestOpt) (*ConfigDraftUpdated, *Response, error) {
	details := new(ConfigDraftUpdated)
	resp, err := s.client.Call(ctx, PartialUpdateConfigDraft.Compile(nil, universeId, repository), data, details, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}

type ConfigDraftDelete struct {
	DraftHash string `json:"draftHash"`
}

type ConfigDraftDeleted struct {
	DraftHash string `json:"draftHash"`
}

// DeleteConfigDraft will reset the current draft.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_DeleteDraft
//
// Required scopes: universe:write
//
// [DELETE] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft
func (s *ConfigService) DeleteConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftDelete, reqOpts ...goblox.RequestOpt) (*ConfigDraftDeleted, *Response, error) {
	details := new(ConfigDraftDeleted)
	resp, err := s.client.Call(ctx, DeleteConfigDraft.Compile(nil, universeId, repository), data, details, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}

// UpdateConfigDraft will overwrite the current draft.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_OverwriteDraft
//
// Required scopes: universe:write
//
// [PUT] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft:overwrite
func (s *ConfigService) UpdateConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftUpdate, reqOpts ...goblox.RequestOpt) (*ConfigDraftUpdated, *Response, error) {
	details := new(ConfigDraftUpdated)
	resp, err := s.client.Call(ctx, UpdateConfigDraft.Compile(nil, universeId, repository), data, details, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}

// GetConfig will get the currently published config.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_GetConfigRepositoryFull
//
// Required scopes: universe:read
//
// [GET] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/full
func (s *ConfigService) GetConfig(ctx context.Context, universeId, repository string, reqOpts ...goblox.RequestOpt) (*FullConfig, *Response, error) {
	config := new(FullConfig)
	resp, err := s.client.Call(ctx, GetConfig.Compile(nil, universeId, repository), nil, config, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}

type ConfigDraftPublish struct {
	DraftHash          string `json:"draftHash"`
	Message            string `json:"message,omitempty"`
	DeploymentStrategy string `json:"deploymentStrategy,omitempty"`
}

type ConfigDraftPublished struct {
	ConfigVersion int `json:"configVersion"`
}

// PublishConfigDraft will publish the current draft changes to the published config.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_PublishDraft
//
// Required scopes: universe:write
//
// [POST] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/publish
func (s *ConfigService) PublishConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftPublish, reqOpts ...goblox.RequestOpt) (*ConfigDraftPublished, *Response, error) {
	details := new(ConfigDraftPublished)
	resp, err := s.client.Call(ctx, PublishConfigDraft.Compile(nil, universeId, repository), data, details, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}

type ConfigRevisionChange struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type ConfigRevision struct {
	RevisionID           string                          `json:"revisionId"`
	Version              int                             `json:"version"`
	Time                 string                          `json:"time"`
	PublishedBy          string                          `json:"publishedBy"`
	Message              string                          `json:"message"`
	DeploymentResult     string                          `json:"deploymentResult"`
	AutoRestoreToVersion int                             `json:"autoRestoreToVersion"`
	Changes              map[string]ConfigRevisionChange `json:"changes"`
}

type ConfigRevisionHistory struct {
	Revisions []ConfigRevision `json:"revisions"`
}

type ConfigRevisionHistoryOptions struct {
	StartTime   *string `url:"startTime,omitempty"`
	EndTime     *string `url:"endTime,omitempty"`
	MaxPageSize *int    `url:"MaxPageSize,omitempty"`
	Skip        *bool   `url:"Skip,omitempty"`
	SearchKey   *string `url:"SearchKey,omitempty"`
	SortOrder   *string `url:"SortOrder,omitempty"`
	SortKey     *string `url:"SortKey,omitempty"`
}

// ListRevisionHistory will list the previous revisions to the config.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_ListRevisions
//
// Required scopes: universe:read
//
// [GET] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/revisions
func (s *ConfigService) ListRevisionHistory(ctx context.Context, universeId, repository string, opts *ConfigRevisionHistoryOptions, reqOpts ...goblox.RequestOpt) (*ConfigRevisionHistory, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	revisionHistory := new(ConfigRevisionHistory)
	resp, err := s.client.Call(ctx, ListRevisionHistory.Compile(q, universeId, repository), nil, revisionHistory, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return revisionHistory, resp, nil
}

type ConfigRestoreRevision struct {
	DraftHash string `json:"draftHash"`
}

// RevisionRestore will stage a revert to a previous revision of the config.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_RestoreRevision
//
// Required scopes: universe:write
//
// [POST] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/revisions/{revisionId}/restore
func (s *ConfigService) RevisionRestore(ctx context.Context, universeId, repository, revisionId string, reqOpts ...goblox.RequestOpt) (*ConfigRestoreRevision, *Response, error) {
	details := new(ConfigRestoreRevision)
	resp, err := s.client.Call(ctx, RevisionRestore.Compile(nil, universeId, repository, revisionId), nil, details, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}
