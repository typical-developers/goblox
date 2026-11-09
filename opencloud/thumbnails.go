package opencloud

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/luckfire/goblox"
	"github.com/luckfire/goblox/auth"
)

var (
	FindThumbnailPersonalizations = goblox.NewEndpoint(
		http.MethodGet,
		"/thumbnail-personalization-api/v1/universes/{universeId}/personalization",
		auth.KindAPIKey,
	).WithScopes("universe.thumbnail:read")

	CreateThumbnailPersonalization = goblox.NewEndpoint(
		http.MethodPost,
		"/thumbnail-personalization-api/v1/universes/{universeId}/personalization/create",
		auth.KindAPIKey,
	).WithScopes("universe.thumbnail:write")

	UpdateThumbnailPersonalization = goblox.NewEndpoint(
		http.MethodPost,
		"/thumbnail-personalization-api/v1/universes/{universeId}/personalization/update",
		auth.KindAPIKey,
	).WithScopes("universe.thumbnail:write")

	DeleteHomepageThumbnails = goblox.NewEndpoint(
		http.MethodDelete,
		"/thumbnail-personalization-api/v1/universes/{universeId}/thumbnails",
		auth.KindAPIKey,
	).WithScopes("universe.thumbnail:write")

	GetHomepageThumbnails = goblox.NewEndpoint(
		http.MethodGet,
		"/thumbnail-personalization-api/v1/universes/{universeId}/thumbnails",
		auth.KindAPIKey,
	).WithScopes("universe.thumbnail:read")

	UploadHomepageThumbnails = goblox.NewEndpoint(
		http.MethodPost,
		"/thumbnail-personalization-api/v1/universes/{universeId}/thumbnails/uploads",
		auth.KindAPIKey,
	).WithScopes("universe.thumbnail:write")

	GetHomepageThumbnailUploadStatus = goblox.NewEndpoint(
		http.MethodGet,
		"/thumbnail-personalization-api/v1/universes/{universeId}/thumbnails/uploads/status",
		auth.KindAPIKey,
	).WithScopes("universe.thumbnail:read")

	GetAssetThumbnail3D = goblox.NewEndpoint(
		http.MethodGet,
		"https://thumbnails.roblox.com/v1/assets-thumbnail-3d",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("thumbnail:read")

	GetUserAvatarThumbnail3D = goblox.NewEndpoint(
		http.MethodGet,
		"https://thumbnails.roblox.com/v1/users/avatar-3d",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("thumbnail:read")

	GetOutfitThumbnail3D = goblox.NewEndpoint(
		http.MethodGet,
		"https://thumbnails.roblox.com/v1/users/outfit-3d",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("thumbnail:read")
)

// ThumbnailsService handles communication with the thumbnail endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type ThumbnailsService service

type PersonalizedConfigStatus string

const (
	PersonalizedConfigStatusActive   PersonalizedConfigStatus = "Active"
	PersonalizedConfigStatusInactive PersonalizedConfigStatus = "Inactive"
)

type FindThumbnailPersonalizationsOptions struct {
	Status PersonalizedConfigStatus `url:"status"` // Required.
	Limit  *int                     `url:"limit,omitempty"`
	Cursor *string                  `url:"cursor,omitempty"`
}

type PersonalizedThumbnailStatus string

const (
	PersonalizedThumbnailStatusActive   PersonalizedThumbnailStatus = "Active"
	PersonalizedThumbnailStatusInactive PersonalizedThumbnailStatus = "Inactive"
)

type HomepageThumbnailStatus string

const (
	HomepageThumbnailStatusActive HomepageThumbnailStatus = "Active"
	HomepageThumbnailStatusSpammy HomepageThumbnailStatus = "Spammy"
)

type PersonalizedThumbnail struct {
	HomepageThumbnailID         string                      `json:"homepageThumbnailId"`
	AssetID                     int64                       `json:"assetId"`
	PersonalizedThumbnailStatus PersonalizedThumbnailStatus `json:"personalizedThumbnailStatus"`
	HomepageThumbnailStatus     HomepageThumbnailStatus     `json:"homepageThumbnailStatus"`
}

type ThumbnailPersonalizedConfig struct {
	ID                       string                   `json:"id"`
	PersonalizedConfigStatus PersonalizedConfigStatus `json:"personalizedConfigStatus"`
	CreatedUTC               string                   `json:"createdUtc"`
	Thumbnails               []PersonalizedThumbnail  `json:"thumbnails"`
	LastServedUTC            string                   `json:"lastServedUtc"`
}

type FindPersonalizedConfigsResponse struct {
	PersonalizedConfigs []ThumbnailPersonalizedConfig `json:"personalizedConfigs"`
	NextCursor          string                        `json:"nextCursor"`
}

// FindThumbnailPersonalizations will get a page of thumbnail personalization configurations for a universe.
//
// Required scopes: universe.thumbnail:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/thumbnail-personalization#ThumbnailPersonalizationApi.HomepageThumbnail_FindThumbnailPersonalizations
//
// [GET] /thumbnail-personalization-api/v1/universes/{universeId}/personalization
func (s *ThumbnailsService) FindThumbnailPersonalizations(ctx context.Context, universeId string, opts *FindThumbnailPersonalizationsOptions, reqOpts ...goblox.RequestOpt) (*FindPersonalizedConfigsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	findPersonalizedConfigsResponse := new(FindPersonalizedConfigsResponse)
	resp, err := s.client.Call(ctx, FindThumbnailPersonalizations.Compile(q, universeId), nil, findPersonalizedConfigsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return findPersonalizedConfigsResponse, resp, nil
}

type CreateThumbnailPersonalizationRequest struct {
	HomepageThumbnailIDs []string `json:"homepageThumbnailIds,omitempty"`
}

// CreateThumbnailPersonalization will create a new active thumbnail personalization configuration for a
// universe.
//
// Required scopes: universe.thumbnail:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/thumbnail-personalization#ThumbnailPersonalizationApi.HomepageThumbnail_CreateThumbnailPersonalization
//
// [POST] /thumbnail-personalization-api/v1/universes/{universeId}/personalization/create
func (s *ThumbnailsService) CreateThumbnailPersonalization(ctx context.Context, universeId string, data CreateThumbnailPersonalizationRequest, reqOpts ...goblox.RequestOpt) (string, *Response, error) {
	var result string
	resp, err := s.client.Call(ctx, CreateThumbnailPersonalization.Compile(nil, universeId), data, &result, reqOpts...)
	if err != nil {
		return "", resp, err
	}

	return result, resp, nil
}

type UpdateThumbnailPersonalizationRequest struct {
	HomepageThumbnailIDs []string `json:"homepageThumbnailIds,omitempty"`
	ID                   *string  `json:"id,omitempty"`
}

// UpdateThumbnailPersonalization will replace the homepage thumbnails in an active personalization configuration
// without changing its identifier.
//
// Required scopes: universe.thumbnail:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/thumbnail-personalization#ThumbnailPersonalizationApi.HomepageThumbnail_UpdateThumbnailPersonalization
//
// [POST] /thumbnail-personalization-api/v1/universes/{universeId}/personalization/update
func (s *ThumbnailsService) UpdateThumbnailPersonalization(ctx context.Context, universeId string, data UpdateThumbnailPersonalizationRequest, reqOpts ...goblox.RequestOpt) (string, *Response, error) {
	var result string
	resp, err := s.client.Call(ctx, UpdateThumbnailPersonalization.Compile(nil, universeId), data, &result, reqOpts...)
	if err != nil {
		return "", resp, err
	}

	return result, resp, nil
}

type DeleteHomepageThumbnailsOptions struct {
	HomepageThumbnailIDs []string `url:"homepageThumbnailIds"` // Required.
}

// DeleteHomepageThumbnails will delete homepage thumbnails from a universe.
//
// Required scopes: universe.thumbnail:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/thumbnail-personalization#ThumbnailPersonalizationApi.HomepageThumbnail_DeleteHomepageThumbnails
//
// [DELETE] /thumbnail-personalization-api/v1/universes/{universeId}/thumbnails
func (s *ThumbnailsService) DeleteHomepageThumbnails(ctx context.Context, universeId string, opts *DeleteHomepageThumbnailsOptions, reqOpts ...goblox.RequestOpt) (map[string]bool, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	var result map[string]bool
	resp, err := s.client.Call(ctx, DeleteHomepageThumbnails.Compile(q, universeId), nil, &result, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

type GetHomepageThumbnailsOptions struct {
	Limit      *int    `url:"limit,omitempty"`
	NextCursor *string `url:"nextCursor,omitempty"`
}

type ThumbnailModerationStatus string

const (
	ThumbnailModerationStatusReviewing   ThumbnailModerationStatus = "Reviewing"
	ThumbnailModerationStatusRejected    ThumbnailModerationStatus = "Rejected"
	ThumbnailModerationStatusApproved    ThumbnailModerationStatus = "Approved"
	ThumbnailModerationStatusUnspecified ThumbnailModerationStatus = "Unspecified"
)

type HomepageThumbnail struct {
	HomepageThumbnailID      string                    `json:"homepageThumbnailId"`
	AssetID                  int64                     `json:"assetId"`
	ModerationStatus         ThumbnailModerationStatus `json:"moderationStatus"`
	PersonalizedConfigStatus PersonalizedConfigStatus  `json:"personalizedConfigStatus"`
	HomepageThumbnailStatus  HomepageThumbnailStatus   `json:"homepageThumbnailStatus"`
}

type FindThumbnailsResponse struct {
	HomepageThumbnails []HomepageThumbnail `json:"homepageThumbnails"`
	NextCursor         string              `json:"nextCursor"`
}

// GetHomepageThumbnails will get a page of homepage thumbnails for a universe.
//
// Required scopes: universe.thumbnail:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/thumbnail-personalization#ThumbnailPersonalizationApi.HomepageThumbnail_GetHomepageThumbnails
//
// [GET] /thumbnail-personalization-api/v1/universes/{universeId}/thumbnails
func (s *ThumbnailsService) GetHomepageThumbnails(ctx context.Context, universeId string, opts *GetHomepageThumbnailsOptions, reqOpts ...goblox.RequestOpt) (*FindThumbnailsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	findThumbnailsResponse := new(FindThumbnailsResponse)
	resp, err := s.client.Call(ctx, GetHomepageThumbnails.Compile(q, universeId), nil, findThumbnailsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return findThumbnailsResponse, resp, nil
}

type UploadHomepageThumbnailsRequest struct {
	Files []*os.File `form:"files"`
}

type UploadThumbnailsResponse struct {
	FileToOperationIDDict map[string]string `json:"fileToOperationIdDict"`
}

// UploadHomepageThumbnails will upload one or more images to create homepage thumbnails for a universe.
//
// Required scopes: universe.thumbnail:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/thumbnail-personalization#ThumbnailPersonalizationApi.HomepageThumbnail_UploadHomepageThumbnails
//
// [POST] /thumbnail-personalization-api/v1/universes/{universeId}/thumbnails/uploads
func (s *ThumbnailsService) UploadHomepageThumbnails(ctx context.Context, universeId string, data UploadHomepageThumbnailsRequest, reqOpts ...goblox.RequestOpt) (*UploadThumbnailsResponse, *Response, error) {
	uploadThumbnailsResponse := new(UploadThumbnailsResponse)
	e := UploadHomepageThumbnails.Compile(nil, universeId)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	for _, file := range data.Files {
		part, err := writer.CreateFormFile("files", filepath.Base(file.Name()))
		if err != nil {
			return nil, nil, err
		}
		if _, err := io.Copy(part, file); err != nil {
			return nil, nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, writer.FormDataContentType())
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.CallRequest(ctx, e, req, uploadThumbnailsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return uploadThumbnailsResponse, resp, nil
}

type GetHomepageThumbnailUploadStatusOptions struct {
	OperationIDs []string `url:"operationIds"` // Required.
}

type ThumbnailUploadStatus int

type UploadThumbnailStatus struct {
	HomepageThumbnailID string                    `json:"homepageThumbnailId"`
	AssetID             int64                     `json:"assetId"`
	ModerationStatus    ThumbnailModerationStatus `json:"moderationStatus"`
}

type UploadThumbnailsStatusResponse struct {
	UploadStatus              ThumbnailUploadStatus            `json:"uploadStatus"`
	UploadThumbnailStatusDict map[string]UploadThumbnailStatus `json:"uploadThumbnailStatusDict"`
}

// GetHomepageThumbnailUploadStatus will get the processing status of homepage thumbnail upload operations.
//
// Required scopes: universe.thumbnail:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/thumbnail-personalization#ThumbnailPersonalizationApi.HomepageThumbnail_GetHomepageThumbnailsStatus
//
// [GET] /thumbnail-personalization-api/v1/universes/{universeId}/thumbnails/uploads/status
func (s *ThumbnailsService) GetHomepageThumbnailUploadStatus(ctx context.Context, universeId string, opts *GetHomepageThumbnailUploadStatusOptions, reqOpts ...goblox.RequestOpt) (*UploadThumbnailsStatusResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	uploadThumbnailsStatusResponse := new(UploadThumbnailsStatusResponse)
	resp, err := s.client.Call(ctx, GetHomepageThumbnailUploadStatus.Compile(q, universeId), nil, uploadThumbnailsStatusResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return uploadThumbnailsStatusResponse, resp, nil
}

type GetAssetThumbnail3DOptions struct {
	AssetID       int64  `url:"assetId"` // Required.
	UseGltf       *bool  `url:"useGltf,omitempty"`
	RobloxPlaceID *int64 `url:"-"`
}

type Thumbnail3D struct {
	TargetID int64  `json:"targetId"`
	State    string `json:"state"`
	ImageURL string `json:"imageUrl"`
	Version  string `json:"version"`
}

// GetAssetThumbnail3D will get a 3D thumbnail for an asset.
//
// Required scopes: thumbnail:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/assets#thumbnails_get_v1_assets_thumbnail_3d
//
// [GET] /v1/assets-thumbnail-3d
func (s *ThumbnailsService) GetAssetThumbnail3D(ctx context.Context, opts *GetAssetThumbnail3DOptions, reqOpts ...goblox.RequestOpt) (*Thumbnail3D, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}
	if opts != nil {
		if opts.RobloxPlaceID != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("Roblox-Place-Id", fmt.Sprint(*opts.RobloxPlaceID)))
		}
	}

	thumbnail3D := new(Thumbnail3D)
	resp, err := s.client.Call(ctx, GetAssetThumbnail3D.Compile(q), nil, thumbnail3D, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return thumbnail3D, resp, nil
}

type GetUserAvatarThumbnail3DOptions struct {
	UserID int64 `url:"userId"` // Required.
}

// GetUserAvatarThumbnail3D will get a 3D thumbnail of a user's avatar.
//
// Required scopes: thumbnail:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/avatars#thumbnails_get_v1_users_avatar_3d
//
// [GET] /v1/users/avatar-3d
func (s *ThumbnailsService) GetUserAvatarThumbnail3D(ctx context.Context, opts *GetUserAvatarThumbnail3DOptions, reqOpts ...goblox.RequestOpt) (*Thumbnail3D, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	thumbnail3D := new(Thumbnail3D)
	resp, err := s.client.Call(ctx, GetUserAvatarThumbnail3D.Compile(q), nil, thumbnail3D, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return thumbnail3D, resp, nil
}

type GetOutfitThumbnail3DOptions struct {
	OutfitID int64 `url:"outfitId"` // Required.
}

// GetOutfitThumbnail3D will get a 3D thumbnail of an outfit.
//
// Required scopes: thumbnail:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/thumbnails#thumbnails_get_v1_users_outfit_3d
//
// [GET] /v1/users/outfit-3d
func (s *ThumbnailsService) GetOutfitThumbnail3D(ctx context.Context, opts *GetOutfitThumbnail3DOptions, reqOpts ...goblox.RequestOpt) (*Thumbnail3D, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	thumbnail3D := new(Thumbnail3D)
	resp, err := s.client.Call(ctx, GetOutfitThumbnail3D.Compile(q), nil, thumbnail3D, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return thumbnail3D, resp, nil
}
