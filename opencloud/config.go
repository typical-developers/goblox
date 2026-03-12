package opencloud

import (
	"context"
	"fmt"
	"net/http"
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
func (s *ConfigService) GetConfigWithoutMetadata(ctx context.Context, universeId, repository string) (*Config, *Response, error) {
	u := fmt.Sprintf("/creator-configs-public-api/v1/configs/universes/%s/repositories/%s", universeId, repository)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	config := new(Config)
	resp, err := s.client.Do(ctx, req, config)
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
func (s *ConfigService) GetConfigDraft(ctx context.Context, universeId, repository string) (*ConfigDraft, *Response, error) {
	u := fmt.Sprintf("/creator-configs-public-api/v1/configs/universes/%s/repositories/%s/draft", universeId, repository)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	draft := new(ConfigDraft)
	resp, err := s.client.Do(ctx, req, draft)
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
func (s *ConfigService) PartialUpdateConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftUpdate) (*ConfigDraftUpdated, *Response, error) {
	u := fmt.Sprintf("/creator-configs-public-api/v1/configs/universes/%s/repositories/%s/draft", universeId, repository)

	req, err := s.client.NewRequest(http.MethodPatch, u, data)
	if err != nil {
		return nil, nil, err
	}

	details := new(ConfigDraftUpdated)
	resp, err := s.client.Do(ctx, req, details)
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
func (s *ConfigService) DeleteConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftDelete) (*ConfigDraftDeleted, *Response, error) {
	u := fmt.Sprintf("/creator-configs-public-api/v1/configs/universes/%s/repositories/%s/draft", universeId, repository)

	req, err := s.client.NewRequest(http.MethodDelete, u, data)
	if err != nil {
		return nil, nil, err
	}

	details := new(ConfigDraftDeleted)
	resp, err := s.client.Do(ctx, req, details)
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
func (s *ConfigService) UpdateConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftUpdate) (*ConfigDraftUpdated, *Response, error) {
	u := fmt.Sprintf("/creator-configs-public-api/v1/configs/universes/%s/repositories/%s/draft:overwrite", universeId, repository)

	req, err := s.client.NewRequest(http.MethodPut, u, data)
	if err != nil {
		return nil, nil, err
	}

	details := new(ConfigDraftUpdated)
	resp, err := s.client.Do(ctx, req, details)
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
func (s *ConfigService) GetConfig(ctx context.Context, universeId, repository string) (*FullConfig, *Response, error) {
	u := fmt.Sprintf("/creator-configs-public-api/v1/configs/universes/%s/repositories/%s/full", universeId, repository)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	config := new(FullConfig)
	resp, err := s.client.Do(ctx, req, config)
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
func (s *ConfigService) PublishConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftPublish) (*ConfigDraftPublished, *Response, error) {
	u := fmt.Sprintf("/creator-configs-public-api/v1/configs/universes/%s/repositories/%s/publish", universeId, repository)

	req, err := s.client.NewRequest(http.MethodPost, u, data)
	if err != nil {
		return nil, nil, err
	}

	details := new(ConfigDraftPublished)
	resp, err := s.client.Do(ctx, req, details)
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
func (s *ConfigService) ListRevisionHistory(ctx context.Context, universeId, repository string, opts *ConfigRevisionHistoryOptions) (*ConfigRevisionHistory, *Response, error) {
	u := fmt.Sprintf("/creator-configs-public-api/v1/configs/universes/%s/repositories/%s/revisions", universeId, repository)

	u, err := addOpts(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	revisionHistory := new(ConfigRevisionHistory)
	resp, err := s.client.Do(ctx, req, revisionHistory)
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
func (s *ConfigService) RevisionRestore(ctx context.Context, universeId, repository, revisionId string) (*ConfigRestoreRevision, *Response, error) {
	u := fmt.Sprintf("/creator-configs-public-api/v1/configs/universes/%s/repositories/%s/revisions/%s/restore", universeId, repository, revisionId)

	req, err := s.client.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	details := new(ConfigRestoreRevision)
	resp, err := s.client.Do(ctx, req, details)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}
