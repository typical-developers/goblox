package opencloud

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	GetAsset = goblox.NewEndpoint(
		http.MethodGet,
		"/assets/v1/assets/{assetId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read")

	ArchiveAsset = goblox.NewEndpoint(
		http.MethodPost,
		"/assets/v1/assets/{assetId}:archive",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read", "asset:write")

	RestoreAsset = goblox.NewEndpoint(
		http.MethodPost,
		"/assets/v1/assets/{assetId}:restore",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read", "asset:write")

	GetAssetVersion = goblox.NewEndpoint(
		http.MethodGet,
		"/assets/v1/assets/{assetId}/versions/{versionNumber}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read")

	GetAssetVersions = goblox.NewEndpoint(
		http.MethodGet,
		"/assets/v1/assets/{assetId}/versions",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read")

	RollbackAssetVersion = goblox.NewEndpoint(
		http.MethodPost,
		"/assets/v1/assets/{assetId}/versions:rollback",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read", "asset:write")

	GetOperation = goblox.NewEndpoint(
		http.MethodGet,
		"/assets/v1/operations/{operationId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read")

	CreateAsset = goblox.NewEndpoint(
		http.MethodPost,
		"/assets/v1/assets",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read", "asset:write")

	UpdateAsset = goblox.NewEndpoint(
		http.MethodPatch,
		"/assets/v1/assets/{assetId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read", "asset:write")
)

// AssetsService will handle communciation with the actions related to the API.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type AssetsService service

type AssetType string

const (
	AssetTypeAudio AssetType = "AUDIO"
	AssetTypeDecal AssetType = "DECAL"
	AssetTypeModel AssetType = "MODEL"
)

type AssetCreator struct {
	UserId  string `json:"userId,omitempty"`
	GroupId string `json:"groupId,omitempty"`
}

type AssetCreationContext struct {
	AssetPrivacy  string       `json:"assetPrivacy"`
	Creator       AssetCreator `json:"creator"`
	ExpectedPrice int          `json:"expectedPrices"`
}

type AssetModerationState string

const (
	AssetModerationStateReviewing AssetModerationState = "Reviewing"
	AssetModerationStateRejected  AssetModerationState = "Rejected"
	AssetModerationStateApproved  AssetModerationState = "Approved"
)

type AssetModerationResult struct {
	ModerationState AssetModerationState `json:"moderationState"`
}

type AssetPreview struct {
	Asset   string `json:"asset"`
	AltText string `json:"altText"`
}

type AssetSocialLink struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
}

type AssetSocialLinks struct {
	FacebookSocialLink *AssetSocialLink `json:"facebookSocialLink,omitempty"`
	TwitterSocialLink  *AssetSocialLink `json:"twitterSocialLink,omitempty"`
	YouTubeSocialLink  *AssetSocialLink `json:"youtubeSocialLink,omitempty"`
	TwitchSocialLink   *AssetSocialLink `json:"twitchSocialLink,omitempty"`
	DiscordSocialLink  *AssetSocialLink `json:"discordSocialLink,omitempty"`
	GitHubSocialLink   *AssetSocialLink `json:"githubSocialLink,omitempty"`
	RobloxSocialLink   *AssetSocialLink `json:"robloxSocialLink,omitempty"`
	GuildedSocialLink  *AssetSocialLink `json:"guildedSocialLink,omitempty"`
	DevForumSocialLink *AssetSocialLink `json:"devForumSocialLink,omitempty"`
}

type Asset struct {
	AssetType          AssetType             `json:"assetType"`
	AssetID            string                `json:"assetId"`
	CreationConext     AssetCreationContext  `json:"creationContext"`
	DisplayName        string                `json:"displayName"`
	Description        string                `json:"description"`
	Path               string                `json:"path"`
	RevisionID         string                `json:"revisionId"`
	RevisionCreateTime string                `json:"revisionCreateTime"`
	ModerationResult   AssetModerationResult `json:"moderationResult"`
	Icon               string                `json:"icon"`
	Previews           []AssetPreview        `json:"previews"`
	State              string                `json:"state"`
	AssetSocialLinks
}

type AssetCreate struct {
	AssetType       *AssetType            `json:"assetType,omitempty"`
	DisplayAnme     *string               `json:"displayName,omitempty"`
	Description     *string               `json:"description,omitempty"`
	CreationContext *AssetCreationContext `json:"creationContext,omitempty"`
}

// CreateAsset will create a new asset.
//
// Required scopes:
//
// - asset:read
//
// - asset:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/reference/cloud/assets/v1#POST-v1-assets
//
// [POST] /assets/v1/assets
func (s *AssetsService) CreateAsset(ctx context.Context, data AssetCreate, asset *os.File, reqOpts ...goblox.RequestOpt) (*Operation, *Response, error) {
	e := CreateAsset.Compile(nil)

	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Writes the request field with the desired JSON data.
	_ = writer.WriteField("request", string(jsonb))

	// Writes the fileContent and type.
	part, err := writer.CreateFormFile("fileContent", filepath.Base(asset.Name()))
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(part, asset)
	if err != nil {
		return nil, nil, err
	}

	header := make([]byte, 512)
	_, err = asset.Read(header)
	if err != nil {
		return nil, nil, err
	}

	_ = writer.WriteField("type", http.DetectContentType(header))
	//  --

	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, writer.FormDataContentType())
	if err != nil {
		return nil, nil, err
	}

	operation := new(Operation)
	resp, err := s.client.CallRequest(ctx, e, req, operation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return operation, resp, nil
}

type AssetGetOptions struct {
	ReadMask string `url:"readMask,omitempty"`
}

// GetAsset will fetch a specificed asset.
//
// Required scopes: asset:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/reference/cloud/assets/v1#GET-v1-assets
//
// [GET] /assets/v1/assets/{asset_id}
func (s *AssetsService) GetAsset(ctx context.Context, assetId string, opts *AssetGetOptions, reqOpts ...goblox.RequestOpt) (*Asset, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	asset := new(Asset)
	resp, err := s.client.Call(ctx, GetAsset.Compile(q, assetId), nil, asset, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return asset, resp, nil
}

type AssetUpdate struct {
	AssetType       *AssetType            `json:"assetType,omitempty"`
	AssetID         *string               `json:"assetId,omitempty"`
	DisplayName     *string               `json:"displayName,omitempty"`
	Description     *string               `json:"description,omitempty"`
	CreationContext *AssetCreationContext `json:"creationContext,omitempty"`
	Previews        []*AssetPreview       `json:"previews,omitempty"`
	AssetSocialLinks
}

type AssetUpdateOptions struct {
	AssetID     *string `url:"assetId,omitempty"`
	UpdateMasdk *string `url:"updateMask,omitempty"`
}

// UpdateAsset will update an asset's metadata or contents.
//
// Required scopes:
//
// - asset:read
//
// - asset:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/reference/cloud/assets/v1#PATCH-v1-assets-_assetId_
//
// [PATCH] /assets/v1/assets/{asset_id}
func (s *AssetsService) UpdateAsset(ctx context.Context, assetId string, data *AssetUpdate, asset *os.File, opts *AssetUpdateOptions, reqOpts ...goblox.RequestOpt) (*Operation, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	e := UpdateAsset.Compile(q, assetId)

	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Writes the request field with the desired JSON data.
	_ = writer.WriteField("request", string(jsonb))

	// Writes the fileContent and type.
	if asset != nil {
		part, err := writer.CreateFormFile("fileContent", filepath.Base(asset.Name()))
		if err != nil {
			return nil, nil, err
		}
		_, err = io.Copy(part, asset)
		if err != nil {
			return nil, nil, err
		}

		header := make([]byte, 512)
		_, err = asset.Read(header)
		if err != nil {
			return nil, nil, err
		}

		_ = writer.WriteField("type", http.DetectContentType(header))
	}
	//  --

	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, writer.FormDataContentType())
	if err != nil {
		return nil, nil, err
	}

	operation := new(Operation)
	resp, err := s.client.CallRequest(ctx, e, req, operation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return operation, resp, nil
}

// ArchiveAsset will archive a specificed asset, making it no longer usable / visible in experiences.
//
// Required scopes:
//
// - asset:read
//
// - asset:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/reference/cloud/assets/v1#POST-v1-assets-{assetId}:archive
//
// [POST]  /assets/v1/assets/{assetId}:archive
func (s *AssetsService) ArchiveAsset(ctx context.Context, assetId string, reqOpts ...goblox.RequestOpt) (*Asset, *Response, error) {
	asset := new(Asset)
	resp, err := s.client.Call(ctx, ArchiveAsset.Compile(nil, assetId), nil, asset, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return asset, resp, nil
}

// RestoreAsset will restore a previously archived asset, making it usable / visible once more in experiences.
//
// Required scopes:
//
// - asset:read
//
// - asset:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/reference/cloud/assets/v1#POST-v1-assets-{assetId}:restore
//
// [POST] /assets/v1/assets/{assetId}:restore
func (s *AssetsService) RestoreAsset(ctx context.Context, assetId string, reqOpts ...goblox.RequestOpt) (*Asset, *Response, error) {
	asset := new(Asset)
	resp, err := s.client.Call(ctx, RestoreAsset.Compile(nil, assetId), nil, asset, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return asset, resp, nil
}

type AssetVersion struct {
	CreationContext  AssetCreationContext  `json:"creationContext"`
	Path             string                `json:"path"`
	ModerationResult AssetModerationResult `json:"moderationResult"`
	Published        bool                  `json:"published"`
}

// GetAssetVersion will fetch an asset's version.
//
// Required scopes: asset:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/reference/cloud/assets/v1#GET-v1-assets-_assetId_-versions-_versionNumber_
//
// [GET] /assets/v1/assets/{assetId}/versions/{versionNumber}
func (s *AssetsService) GetAssetVersion(ctx context.Context, assetId string, versionNumber string, reqOpts ...goblox.RequestOpt) (*AssetVersion, *Response, error) {
	asset := new(AssetVersion)
	resp, err := s.client.Call(ctx, GetAssetVersion.Compile(nil, assetId, versionNumber), nil, asset, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return asset, resp, nil
}

type AssetVersionList struct {
	AssetVersions []AssetVersion `json:"assetVersions"`
	NextPageToken string         `json:"nextPageToken"`
}

// GetAssetVersions will fetch an asset's versions.
//
// Required scopes: asset:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/reference/cloud/assets/v1#GET-assets-v1-assets-_assetId_-versions
//
// [GET] /assets/v1/assets/{assetId}/versions
func (s *AssetsService) GetAssetVersions(ctx context.Context, assetId string, opts *Options, reqOpts ...goblox.RequestOpt) (*AssetVersionList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	asset := new(AssetVersionList)
	resp, err := s.client.Call(ctx, GetAssetVersions.Compile(q, assetId), nil, asset, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return asset, resp, nil
}

type AssetVersionRollback struct {
	AssetVersion *string `json:"assetVersion,omitempty"`
}

// RollbackAssetVersion will rollback an asset's version to a specified version.
//
// Required scopes:
//
// - asset:read
//
// - asset:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/reference/cloud/assets/v1#POST-assets-v1-assets-_assetId_-versions:rollback
//
// [POST] /assets/v1/assets/{assetId}/versions:rollback
func (s *AssetsService) RollbackAssetVersion(ctx context.Context, assetId string, data AssetVersionRollback, reqOpts ...goblox.RequestOpt) (*AssetVersion, *Response, error) {
	asset := new(AssetVersion)
	resp, err := s.client.Call(ctx, RollbackAssetVersion.Compile(nil, assetId), data, asset, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return asset, resp, nil
}

// GetOperation will fetch an asset's operation.
//
// Required scopes: asset:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/reference/cloud/assets/v1#GET-v1-operations-_operationId_
//
// [GET] /assets/v1/operations/{operationId}
func (s *AssetsService) GetOperation(ctx context.Context, operationId string, reqOpts ...goblox.RequestOpt) (*Operation, *Response, error) {
	operation := new(Operation)
	resp, err := s.client.Call(ctx, GetOperation.Compile(nil, operationId), nil, operation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return operation, resp, nil
}
