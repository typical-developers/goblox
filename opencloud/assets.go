package opencloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	).WithScopes("asset:write")

	RestoreAsset = goblox.NewEndpoint(
		http.MethodPost,
		"/assets/v1/assets/{assetId}:restore",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:write")

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
	).WithScopes("asset:write")

	GetOperation = goblox.NewEndpoint(
		http.MethodGet,
		"/assets/v1/operations/{operationId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read")

	CreateAsset = goblox.NewEndpoint(
		http.MethodPost,
		"/assets/v1/assets",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:write")

	UpdateAsset = goblox.NewEndpoint(
		http.MethodPatch,
		"/assets/v1/assets/{assetId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:write")
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

var (
	GetAssetDeliveryLocation = goblox.NewEndpoint(
		http.MethodGet,
		"/asset-delivery-api/v1/assetId/{assetId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-asset:manage")

	GetAssetVersionDeliveryLocation = goblox.NewEndpoint(
		http.MethodGet,
		"/asset-delivery-api/v1/assetId/{assetId}/version/{versionNumber}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-asset:manage")

	GrantAssetPermissions = goblox.NewEndpoint(
		http.MethodPatch,
		"/asset-permissions-api/v1/assets/permissions",
		auth.KindAPIKey,
	).WithScopes("asset-permissions:write")
)

type GetAssetDeliveryLocationOptions struct {
	AcceptEncoding                        *string `url:"-"`
	RobloxPlaceID                         *int64  `url:"-"`
	AssetType                             *string `url:"-"`
	AssetFormat                           *string `url:"-"`
	RobloxAssetFormat                     *string `url:"-"`
	RobloxUsageContext                    *string `url:"-"`
	SkipSigningScripts                    *bool   `url:"skipSigningScripts,omitempty"`
	ClientInsert                          *int    `url:"clientInsert,omitempty"`
	Scriptinsert                          *int    `url:"scriptinsert,omitempty"`
	ModulePlaceID                         *int64  `url:"modulePlaceId,omitempty"`
	Serverplaceid                         *int64  `url:"serverplaceid,omitempty"`
	ExpectedAssetType                     *string `url:"expectedAssetType,omitempty"`
	DoNotFallbackToBaselineRepresentation *bool   `url:"doNotFallbackToBaselineRepresentation,omitempty"`
	ContentRepresentationPriorityList     *string `url:"contentRepresentationPriorityList,omitempty"`
	AccessContext                         *string `url:"accessContext,omitempty"`
}

type AssetDeliveryLocationEntry struct {
	AssetFormat string `json:"assetFormat"`
	Location    string `json:"location"`
}

type AssetDeliveryLocation struct {
	Location     string                       `json:"location"`
	Locations    []AssetDeliveryLocationEntry `json:"locations"`
	RequestID    string                       `json:"requestId"`
	IsArchived   bool                         `json:"isArchived"`
	AssetTypeID  int                          `json:"assetTypeId"`
	IsRecordable bool                         `json:"isRecordable"`
}

// GetAssetDeliveryLocation will retrieve an asset by its ID with OpenCloud auth.
//
// Required scopes: legacy-asset:manage
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/assets#get_asset_delivery_api_v1_assetId__assetId_
//
// [GET] /asset-delivery-api/v1/assetId/{assetId}
func (s *AssetsService) GetAssetDeliveryLocation(ctx context.Context, assetId string, opts *GetAssetDeliveryLocationOptions, reqOpts ...goblox.RequestOpt) (*AssetDeliveryLocation, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}
	if opts != nil {
		if opts.AcceptEncoding != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("Accept-Encoding", fmt.Sprint(*opts.AcceptEncoding)))
		}
		if opts.RobloxPlaceID != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("Roblox-Place-Id", fmt.Sprint(*opts.RobloxPlaceID)))
		}
		if opts.AssetType != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("AssetType", fmt.Sprint(*opts.AssetType)))
		}
		if opts.AssetFormat != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("AssetFormat", fmt.Sprint(*opts.AssetFormat)))
		}
		if opts.RobloxAssetFormat != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("Roblox-AssetFormat", fmt.Sprint(*opts.RobloxAssetFormat)))
		}
		if opts.RobloxUsageContext != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("Roblox-Usage-Context", fmt.Sprint(*opts.RobloxUsageContext)))
		}
	}

	assetDeliveryLocation := new(AssetDeliveryLocation)
	resp, err := s.client.Call(ctx, GetAssetDeliveryLocation.Compile(q, assetId), nil, assetDeliveryLocation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return assetDeliveryLocation, resp, nil
}

type GetAssetVersionDeliveryLocationOptions struct {
	AcceptEncoding                        *string `url:"-"`
	RobloxPlaceID                         *int64  `url:"-"`
	AssetType                             *string `url:"-"`
	AssetFormat                           *string `url:"-"`
	RobloxAssetFormat                     *string `url:"-"`
	RobloxUsageContext                    *string `url:"-"`
	SkipSigningScripts                    *bool   `url:"skipSigningScripts,omitempty"`
	ClientInsert                          *int    `url:"clientInsert,omitempty"`
	Scriptinsert                          *int    `url:"scriptinsert,omitempty"`
	ModulePlaceID                         *int64  `url:"modulePlaceId,omitempty"`
	Serverplaceid                         *int64  `url:"serverplaceid,omitempty"`
	ExpectedAssetType                     *string `url:"expectedAssetType,omitempty"`
	DoNotFallbackToBaselineRepresentation *bool   `url:"doNotFallbackToBaselineRepresentation,omitempty"`
	ContentRepresentationPriorityList     *string `url:"contentRepresentationPriorityList,omitempty"`
	AccessContext                         *string `url:"accessContext,omitempty"`
}

// GetAssetVersionDeliveryLocation will retrieve an asset by its ID and version number with OpenCloud auth.
//
// Required scopes: legacy-asset:manage
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/assets#get_asset_delivery_api_v1_assetId__assetId__version__versionNumber_
//
// [GET] /asset-delivery-api/v1/assetId/{assetId}/version/{versionNumber}
func (s *AssetsService) GetAssetVersionDeliveryLocation(ctx context.Context, assetId, versionNumber string, opts *GetAssetVersionDeliveryLocationOptions, reqOpts ...goblox.RequestOpt) (*AssetDeliveryLocation, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}
	if opts != nil {
		if opts.AcceptEncoding != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("Accept-Encoding", fmt.Sprint(*opts.AcceptEncoding)))
		}
		if opts.RobloxPlaceID != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("Roblox-Place-Id", fmt.Sprint(*opts.RobloxPlaceID)))
		}
		if opts.AssetType != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("AssetType", fmt.Sprint(*opts.AssetType)))
		}
		if opts.AssetFormat != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("AssetFormat", fmt.Sprint(*opts.AssetFormat)))
		}
		if opts.RobloxAssetFormat != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("Roblox-AssetFormat", fmt.Sprint(*opts.RobloxAssetFormat)))
		}
		if opts.RobloxUsageContext != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("Roblox-Usage-Context", fmt.Sprint(*opts.RobloxUsageContext)))
		}
	}

	assetDeliveryLocation := new(AssetDeliveryLocation)
	resp, err := s.client.Call(ctx, GetAssetVersionDeliveryLocation.Compile(q, assetId, versionNumber), nil, assetDeliveryLocation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return assetDeliveryLocation, resp, nil
}

type AssetSubjectType string

const (
	AssetSubjectTypeInvalid      AssetSubjectType = "Invalid"
	AssetSubjectTypeUser         AssetSubjectType = "User"
	AssetSubjectTypeGroup        AssetSubjectType = "Group"
	AssetSubjectTypeGroupRoleset AssetSubjectType = "GroupRoleset"
	AssetSubjectTypeAll          AssetSubjectType = "All"
	AssetSubjectTypeUniverse     AssetSubjectType = "Universe"
)

type AssetAction string

const (
	AssetActionInvalid       AssetAction = "Invalid"
	AssetActionEdit          AssetAction = "Edit"
	AssetActionUse           AssetAction = "Use"
	AssetActionDownload      AssetAction = "Download"
	AssetActionCopyFromRcc   AssetAction = "CopyFromRcc"
	AssetActionUpdateFromRcc AssetAction = "UpdateFromRcc"
)

type AssetGrantRequest struct {
	AssetID             *int64 `json:"assetId,omitempty"`
	GrantToDependencies *bool  `json:"grantToDependencies,omitempty"`
	ParentVersionNumber *int   `json:"parentVersionNumber,omitempty"`
}

type BatchGrantPermissionsRequest struct {
	SubjectType           *AssetSubjectType   `json:"subjectType,omitempty"`
	SubjectID             *string             `json:"subjectId,omitempty"`
	Action                *AssetAction        `json:"action,omitempty"`
	Requests              []AssetGrantRequest `json:"requests,omitempty"`
	AssetIDs              []int64             `json:"assetIds,omitempty"`
	EnableDeepAccessCheck *bool               `json:"enableDeepAccessCheck,omitempty"`
}

type ErrorCode string

const (
	ErrorCodeUnknownError                 ErrorCode = "UnknownError"
	ErrorCodeInvalidRequest               ErrorCode = "InvalidRequest"
	ErrorCodeAssetNotFound                ErrorCode = "AssetNotFound"
	ErrorCodeCannotManageAsset            ErrorCode = "CannotManageAsset"
	ErrorCodePublicAssetCannotBeGrantedTo ErrorCode = "PublicAssetCannotBeGrantedTo"
	ErrorCodeCannotManageSubject          ErrorCode = "CannotManageSubject"
	ErrorCodeSubjectNotFound              ErrorCode = "SubjectNotFound"
	ErrorCodeAssetTypeNotEnabled          ErrorCode = "AssetTypeNotEnabled"
	ErrorCodePermissionLimitReached       ErrorCode = "PermissionLimitReached"
	ErrorCodeDependenciesLimitReached     ErrorCode = "DependenciesLimitReached"
	ErrorCodeNotRequestable               ErrorCode = "NotRequestable"
	ErrorCodeRequesterNotAConnection      ErrorCode = "RequesterNotAConnection"
	ErrorCodeAlreadyPending               ErrorCode = "AlreadyPending"
	ErrorCodeAlreadyHasAccess             ErrorCode = "AlreadyHasAccess"
	ErrorCodeRequestNotFound              ErrorCode = "RequestNotFound"
	ErrorCodeInvalidRequestStatus         ErrorCode = "InvalidRequestStatus"
	ErrorCodeRateLimited                  ErrorCode = "RateLimited"
	ErrorCodeCallerNotOwner               ErrorCode = "CallerNotOwner"
	ErrorCodeBatchSizeLimitExceeded       ErrorCode = "BatchSizeLimitExceeded"
)

type GrantPermissionError struct {
	AssetID int64     `json:"assetId"`
	Code    ErrorCode `json:"code"`
}

type BatchGrantPermissionsResponse struct {
	SuccessAssetIDs []int64                `json:"successAssetIds"`
	Errors          []GrantPermissionError `json:"errors"`
}

// GrantAssetPermissions will grant a subject permission to multiple assets.
//
// Required scopes: asset-permissions:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/assets#patch_asset_permissions_api_v1_assets_permissions
//
// [PATCH] /asset-permissions-api/v1/assets/permissions
func (s *AssetsService) GrantAssetPermissions(ctx context.Context, data BatchGrantPermissionsRequest, reqOpts ...goblox.RequestOpt) (*BatchGrantPermissionsResponse, *Response, error) {
	batchGrantPermissionsResponse := new(BatchGrantPermissionsResponse)
	resp, err := s.client.Call(ctx, GrantAssetPermissions.Compile(nil), data, batchGrantPermissionsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return batchGrantPermissionsResponse, resp, nil
}
