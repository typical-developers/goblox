package opencloud

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	UpdateBadgeLocalizedDescription = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/badges/{badgeId}/description/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-badge:manage")

	ListBadgeLocalizedIcons = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-game-internationalization/v1/badges/{badgeId}/icons",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-badge:manage")

	DeleteBadgeLocalizedIcon = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-game-internationalization/v1/badges/{badgeId}/icons/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-badge:manage")

	UpdateBadgeLocalizedIcon = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-game-internationalization/v1/badges/{badgeId}/icons/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-badge:manage")

	ListBadgeLocalizedNameDescriptions = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-game-internationalization/v1/badges/{badgeId}/name-description",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-badge:manage")

	DeleteBadgeLocalizedNameDescription = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-game-internationalization/v1/badges/{badgeId}/name-description/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-badge:manage")

	UpdateBadgeLocalizedNameDescription = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/badges/{badgeId}/name-description/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-badge:manage")

	UpdateBadgeLocalizedName = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/badges/{badgeId}/name/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-badge:manage")

	UpdateDeveloperProductLocalizedDescription = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/developer-products/{developerProductId}/description/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-developer-product:manage")

	ListDeveloperProductLocalizedIcons = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-game-internationalization/v1/developer-products/{developerProductId}/icons",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-developer-product:manage")

	DeleteDeveloperProductLocalizedIcon = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-game-internationalization/v1/developer-products/{developerProductId}/icons/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-developer-product:manage")

	UpdateDeveloperProductLocalizedIcon = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-game-internationalization/v1/developer-products/{developerProductId}/icons/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-developer-product:manage")

	ListDeveloperProductLocalizedNameDescriptions = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-game-internationalization/v1/developer-products/{developerProductId}/name-description",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-developer-product:manage")

	DeleteDeveloperProductLocalizedNameDescription = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-game-internationalization/v1/developer-products/{developerProductId}/name-description/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-developer-product:manage")

	UpdateDeveloperProductLocalizedNameDescription = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/developer-products/{developerProductId}/name-description/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-developer-product:manage")

	UpdateDeveloperProductLocalizedName = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/developer-products/{developerProductId}/name/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-developer-product:manage")

	ListGameLocalizedIcons = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-game-internationalization/v1/game-icon/games/{gameId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	DeleteGameLocalizedIcon = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-game-internationalization/v1/game-icon/games/{gameId}/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	UpdateGameLocalizedIcon = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-game-internationalization/v1/game-icon/games/{gameId}/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	UpdateGamePassLocalizedDescription = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/game-passes/{gamePassId}/description/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-game-pass:manage")

	ListGamePassLocalizedIcons = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-game-internationalization/v1/game-passes/{gamePassId}/icons",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-game-pass:manage")

	DeleteGamePassLocalizedIcon = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-game-internationalization/v1/game-passes/{gamePassId}/icons/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-game-pass:manage")

	UpdateGamePassLocalizedIcon = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-game-internationalization/v1/game-passes/{gamePassId}/icons/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-game-pass:manage")

	ListGamePassLocalizedNameDescriptions = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-game-internationalization/v1/game-passes/{gamePassId}/name-description",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-game-pass:manage")

	DeleteGamePassLocalizedNameDescription = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-game-internationalization/v1/game-passes/{gamePassId}/name-description/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-game-pass:manage")

	UpdateGamePassLocalizedNameDescription = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/game-passes/{gamePassId}/name-description/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-game-pass:manage")

	UpdateGamePassLocalizedName = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/game-passes/{gamePassId}/name/language-codes/{languageCode}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-game-pass:manage")

	UpdateGameThumbnailAltText = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-game-internationalization/v1/game-thumbnails/games/{gameId}/language-codes/{languageCode}/alt-text",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	UploadGameThumbnail = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-game-internationalization/v1/game-thumbnails/games/{gameId}/language-codes/{languageCode}/image",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	OrderGameThumbnails = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-game-internationalization/v1/game-thumbnails/games/{gameId}/language-codes/{languageCode}/images/order",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	DeleteGameThumbnail = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-game-internationalization/v1/game-thumbnails/games/{gameId}/language-codes/{languageCode}/images/{imageId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetGameNameDescriptionHistory = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-game-internationalization/v1/name-description/games/translation-history",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	UpdateGameNameDescriptions = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/name-description/games/{gameId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	SetGameSourceLanguage = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/source-language/games/{gameId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	UpdateGameSupportedLanguages = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/supported-languages/games/{gameId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetGameAutomaticTranslationStatus = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-game-internationalization/v1/supported-languages/games/{gameId}/automatic-translation-status",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	UpdateGameLanguageAutomaticTranslationStatus = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/supported-languages/games/{gameId}/languages/{languageCode}/automatic-translation-status",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	UpdateGameLanguageImageTranslationStatus = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/supported-languages/games/{gameId}/languages/{languageCode}/image-translation-status",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	UpdateGameLanguageDisplayInfoAutomaticTranslation = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-game-internationalization/v1/supported-languages/games/{gameId}/languages/{languageCode}/universe-display-info-automatic-translation-settings",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetGameDisplayInfoAutomaticTranslationSettings = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-game-internationalization/v1/supported-languages/games/{gameId}/universe-display-info-automatic-translation-settings",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	CreateGameAutoLocalizationTable = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-localization-tables/v1/autolocalization/games/{gameId}/autolocalizationtable",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	UpdateGameAutoLocalizationSettings = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-localization-tables/v1/autolocalization/games/{gameId}/settings",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetAutoLocalizationMetadata = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-localization-tables/v1/autolocalization/metadata",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetLocalizationTableLimits = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-localization-tables/v1/localization-table/limits",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetLocalizationTableByAssetID = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-localization-tables/v1/localization-table/tables/{assetId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetLocalizationTable = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-localization-tables/v1/localization-table/tables/{tableId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	UpdateLocalizationTable = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-localization-tables/v1/localization-table/tables/{tableId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	ListLocalizationTableEntries = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-localization-tables/v1/localization-table/tables/{tableId}/entries",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetLocalizationTableTranslationHistory = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-localization-tables/v1/localization-table/tables/{tableId}/entries/translation-history",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetLocalizationTableEntryCount = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-localization-tables/v1/localization-table/tables/{tableId}/entry-count",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")
)

// LocalizationService handles communication with the game internationalization and localization table endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type LocalizationService service

type UpdateBadgeDescriptionRequest struct {
	Description *string `json:"description,omitempty"`
}

type UpdateBadgeDescriptionResponse struct {
	Description string `json:"description"`
}

// UpdateBadgeLocalizedDescription will update localized description of a badge.
//
// Required scopes: legacy-badge:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/badges#patch_legacy_game_internationalization_v1_badges__badgeId__description_language_codes__languageCode_
//
// [PATCH] /legacy-game-internationalization/v1/badges/{badgeId}/description/language-codes/{languageCode}
func (s *LocalizationService) UpdateBadgeLocalizedDescription(ctx context.Context, badgeId, languageCode string, data UpdateBadgeDescriptionRequest, reqOpts ...goblox.RequestOpt) (*UpdateBadgeDescriptionResponse, *Response, error) {
	updateBadgeDescriptionResponse := new(UpdateBadgeDescriptionResponse)
	resp, err := s.client.Call(ctx, UpdateBadgeLocalizedDescription.Compile(nil, badgeId, languageCode), data, updateBadgeDescriptionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateBadgeDescriptionResponse, resp, nil
}

type ListBadgeLocalizedIconsOptions struct {
	Width  *int `url:"width,omitempty"`
	Height *int `url:"height,omitempty"`
}

type GetBadgeIconResponseState string

const (
	GetBadgeIconResponseStateApproved      GetBadgeIconResponseState = "Approved"
	GetBadgeIconResponseStatePendingReview GetBadgeIconResponseState = "PendingReview"
	GetBadgeIconResponseStateUnAvailable   GetBadgeIconResponseState = "UnAvailable"
	GetBadgeIconResponseStateRejected      GetBadgeIconResponseState = "Rejected"
	GetBadgeIconResponseStateError         GetBadgeIconResponseState = "Error"
)

type GetBadgeIconResponse struct {
	ImageID      string                    `json:"imageId"`
	ImageURL     string                    `json:"imageUrl"`
	State        GetBadgeIconResponseState `json:"state"`
	LanguageCode string                    `json:"languageCode"`
}

type GetBadgeIconList struct {
	Data []GetBadgeIconResponse `json:"data"`
}

// ListBadgeLocalizedIcons will get all icons for a badge.
//
// Required scopes: legacy-badge:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/badges#get_legacy_game_internationalization_v1_badges__badgeId__icons
//
// [GET] /legacy-game-internationalization/v1/badges/{badgeId}/icons
func (s *LocalizationService) ListBadgeLocalizedIcons(ctx context.Context, badgeId string, opts *ListBadgeLocalizedIconsOptions, reqOpts ...goblox.RequestOpt) (*GetBadgeIconList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getBadgeIconList := new(GetBadgeIconList)
	resp, err := s.client.Call(ctx, ListBadgeLocalizedIcons.Compile(q, badgeId), nil, getBadgeIconList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getBadgeIconList, resp, nil
}

// DeleteBadgeLocalizedIcon will delete a localized icon from a badge.
//
// Required scopes: legacy-badge:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/badges#delete_legacy_game_internationalization_v1_badges__badgeId__icons_language_codes__languageCode_
//
// [DELETE] /legacy-game-internationalization/v1/badges/{badgeId}/icons/language-codes/{languageCode}
func (s *LocalizationService) DeleteBadgeLocalizedIcon(ctx context.Context, badgeId, languageCode string, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, DeleteBadgeLocalizedIcon.Compile(nil, badgeId, languageCode), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type UpdateBadgeLocalizedIconRequest struct {
	Files *os.File `form:"Files"`
}

// UpdateBadgeLocalizedIcon will update a badge's icon.
//
// Required scopes: legacy-badge:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/badges#post_legacy_game_internationalization_v1_badges__badgeId__icons_language_codes__languageCode_
//
// [POST] /legacy-game-internationalization/v1/badges/{badgeId}/icons/language-codes/{languageCode}
func (s *LocalizationService) UpdateBadgeLocalizedIcon(ctx context.Context, badgeId, languageCode string, data UpdateBadgeLocalizedIconRequest, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	e := UpdateBadgeLocalizedIcon.Compile(nil, badgeId, languageCode)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Files != nil {
		part, err := writer.CreateFormFile("Files", filepath.Base(data.Files.Name()))
		if err != nil {
			return nil, nil, err
		}
		if _, err := io.Copy(part, data.Files); err != nil {
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

	resp, err := s.client.CallRequest(ctx, e, req, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type LocalizationNameDescriptionUpdateType string

const (
	LocalizationNameDescriptionUpdateTypeInvalid     LocalizationNameDescriptionUpdateType = "Invalid"
	LocalizationNameDescriptionUpdateTypeName        LocalizationNameDescriptionUpdateType = "Name"
	LocalizationNameDescriptionUpdateTypeDescription LocalizationNameDescriptionUpdateType = "Description"
)

type LocalizationNameDescription struct {
	Name         *string                                `json:"name,omitempty"`
	Description  *string                                `json:"description,omitempty"`
	UpdateType   *LocalizationNameDescriptionUpdateType `json:"updateType,omitempty"`
	LanguageCode *string                                `json:"languageCode,omitempty"`
}

type NameDescriptionList struct {
	Data []LocalizationNameDescription `json:"data"`
}

// ListBadgeLocalizedNameDescriptions will get all names and descriptions of a badge.
//
// Required scopes: legacy-badge:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/badges#get_legacy_game_internationalization_v1_badges__badgeId__name_description
//
// [GET] /legacy-game-internationalization/v1/badges/{badgeId}/name-description
func (s *LocalizationService) ListBadgeLocalizedNameDescriptions(ctx context.Context, badgeId string, reqOpts ...goblox.RequestOpt) (*NameDescriptionList, *Response, error) {
	nameDescriptionList := new(NameDescriptionList)
	resp, err := s.client.Call(ctx, ListBadgeLocalizedNameDescriptions.Compile(nil, badgeId), nil, nameDescriptionList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return nameDescriptionList, resp, nil
}

// DeleteBadgeLocalizedNameDescription will delete localized name and description of a badge.
//
// Required scopes: legacy-badge:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/badges#delete_legacy_game_internationalization_v1_badges__badgeId__name_description_language_codes__languageCode_
//
// [DELETE] /legacy-game-internationalization/v1/badges/{badgeId}/name-description/language-codes/{languageCode}
func (s *LocalizationService) DeleteBadgeLocalizedNameDescription(ctx context.Context, badgeId, languageCode string, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, DeleteBadgeLocalizedNameDescription.Compile(nil, badgeId, languageCode), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type UpdateBadgeNameDescriptionRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateBadgeNameDescriptionResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateBadgeLocalizedNameDescription will update localized name and description of a badge.
//
// Required scopes: legacy-badge:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/badges#patch_legacy_game_internationalization_v1_badges__badgeId__name_description_language_codes__languageCode_
//
// [PATCH] /legacy-game-internationalization/v1/badges/{badgeId}/name-description/language-codes/{languageCode}
func (s *LocalizationService) UpdateBadgeLocalizedNameDescription(ctx context.Context, badgeId, languageCode string, data UpdateBadgeNameDescriptionRequest, reqOpts ...goblox.RequestOpt) (*UpdateBadgeNameDescriptionResponse, *Response, error) {
	updateBadgeNameDescriptionResponse := new(UpdateBadgeNameDescriptionResponse)
	resp, err := s.client.Call(ctx, UpdateBadgeLocalizedNameDescription.Compile(nil, badgeId, languageCode), data, updateBadgeNameDescriptionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateBadgeNameDescriptionResponse, resp, nil
}

type UpdateBadgeNameRequest struct {
	Name *string `json:"name,omitempty"`
}

type UpdateBadgeNameResponse struct {
	Name string `json:"name"`
}

// UpdateBadgeLocalizedName will update localized name of a badge.
//
// Required scopes: legacy-badge:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/badges#patch_legacy_game_internationalization_v1_badges__badgeId__name_language_codes__languageCode_
//
// [PATCH] /legacy-game-internationalization/v1/badges/{badgeId}/name/language-codes/{languageCode}
func (s *LocalizationService) UpdateBadgeLocalizedName(ctx context.Context, badgeId, languageCode string, data UpdateBadgeNameRequest, reqOpts ...goblox.RequestOpt) (*UpdateBadgeNameResponse, *Response, error) {
	updateBadgeNameResponse := new(UpdateBadgeNameResponse)
	resp, err := s.client.Call(ctx, UpdateBadgeLocalizedName.Compile(nil, badgeId, languageCode), data, updateBadgeNameResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateBadgeNameResponse, resp, nil
}

type UpdateDeveloperProductDescriptionRequest struct {
	Description *string `json:"description,omitempty"`
}

type UpdateDeveloperProductDescriptionResponse struct {
	Description string `json:"description"`
}

// UpdateDeveloperProductLocalizedDescription will update localized description of a developer product.
//
// Required scopes: legacy-developer-product:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#patch_legacy_game_internationalization_v1_developer_products__developerProductId__description_language_codes__languageCode_
//
// [PATCH] /legacy-game-internationalization/v1/developer-products/{developerProductId}/description/language-codes/{languageCode}
func (s *LocalizationService) UpdateDeveloperProductLocalizedDescription(ctx context.Context, developerProductId, languageCode string, data UpdateDeveloperProductDescriptionRequest, reqOpts ...goblox.RequestOpt) (*UpdateDeveloperProductDescriptionResponse, *Response, error) {
	updateDeveloperProductDescriptionResponse := new(UpdateDeveloperProductDescriptionResponse)
	resp, err := s.client.Call(ctx, UpdateDeveloperProductLocalizedDescription.Compile(nil, developerProductId, languageCode), data, updateDeveloperProductDescriptionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateDeveloperProductDescriptionResponse, resp, nil
}

type ListDeveloperProductLocalizedIconsOptions struct {
	Width  *int `url:"width,omitempty"`
	Height *int `url:"height,omitempty"`
}

type GetDeveloperProductIconResponseState string

const (
	GetDeveloperProductIconResponseStateApproved      GetDeveloperProductIconResponseState = "Approved"
	GetDeveloperProductIconResponseStatePendingReview GetDeveloperProductIconResponseState = "PendingReview"
	GetDeveloperProductIconResponseStateUnAvailable   GetDeveloperProductIconResponseState = "UnAvailable"
	GetDeveloperProductIconResponseStateRejected      GetDeveloperProductIconResponseState = "Rejected"
	GetDeveloperProductIconResponseStateError         GetDeveloperProductIconResponseState = "Error"
)

type GetDeveloperProductIconResponse struct {
	ImageID      string                               `json:"imageId"`
	ImageURL     string                               `json:"imageUrl"`
	State        GetDeveloperProductIconResponseState `json:"state"`
	LanguageCode string                               `json:"languageCode"`
}

type GetDeveloperProductIconList struct {
	Data []GetDeveloperProductIconResponse `json:"data"`
}

// ListDeveloperProductLocalizedIcons will get all icons for a developer product.
//
// Required scopes: legacy-developer-product:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#get_legacy_game_internationalization_v1_developer_products__developerProductId__icons
//
// [GET] /legacy-game-internationalization/v1/developer-products/{developerProductId}/icons
func (s *LocalizationService) ListDeveloperProductLocalizedIcons(ctx context.Context, developerProductId string, opts *ListDeveloperProductLocalizedIconsOptions, reqOpts ...goblox.RequestOpt) (*GetDeveloperProductIconList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getDeveloperProductIconList := new(GetDeveloperProductIconList)
	resp, err := s.client.Call(ctx, ListDeveloperProductLocalizedIcons.Compile(q, developerProductId), nil, getDeveloperProductIconList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getDeveloperProductIconList, resp, nil
}

// DeleteDeveloperProductLocalizedIcon will delete a localized icon from a developer product.
//
// Required scopes: legacy-developer-product:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#delete_legacy_game_internationalization_v1_developer_products__developerProductId__icons_language_codes__languageCode_
//
// [DELETE] /legacy-game-internationalization/v1/developer-products/{developerProductId}/icons/language-codes/{languageCode}
func (s *LocalizationService) DeleteDeveloperProductLocalizedIcon(ctx context.Context, developerProductId, languageCode string, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, DeleteDeveloperProductLocalizedIcon.Compile(nil, developerProductId, languageCode), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type UpdateDeveloperProductLocalizedIconRequest struct {
	Files *os.File `form:"Files"`
}

// UpdateDeveloperProductLocalizedIcon will update a developer product's icon.
//
// Required scopes: legacy-developer-product:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#post_legacy_game_internationalization_v1_developer_products__developerProductId__icons_language_codes__languageCode_
//
// [POST] /legacy-game-internationalization/v1/developer-products/{developerProductId}/icons/language-codes/{languageCode}
func (s *LocalizationService) UpdateDeveloperProductLocalizedIcon(ctx context.Context, developerProductId, languageCode string, data UpdateDeveloperProductLocalizedIconRequest, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	e := UpdateDeveloperProductLocalizedIcon.Compile(nil, developerProductId, languageCode)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Files != nil {
		part, err := writer.CreateFormFile("Files", filepath.Base(data.Files.Name()))
		if err != nil {
			return nil, nil, err
		}
		if _, err := io.Copy(part, data.Files); err != nil {
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

	resp, err := s.client.CallRequest(ctx, e, req, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

// ListDeveloperProductLocalizedNameDescriptions will get all names and descriptions of a developer product.
//
// Required scopes: legacy-developer-product:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#get_legacy_game_internationalization_v1_developer_products__developerProductId__name_description
//
// [GET] /legacy-game-internationalization/v1/developer-products/{developerProductId}/name-description
func (s *LocalizationService) ListDeveloperProductLocalizedNameDescriptions(ctx context.Context, developerProductId string, reqOpts ...goblox.RequestOpt) (*NameDescriptionList, *Response, error) {
	nameDescriptionList := new(NameDescriptionList)
	resp, err := s.client.Call(ctx, ListDeveloperProductLocalizedNameDescriptions.Compile(nil, developerProductId), nil, nameDescriptionList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return nameDescriptionList, resp, nil
}

// DeleteDeveloperProductLocalizedNameDescription will delete localized name and description of a developer
// product.
//
// Required scopes: legacy-developer-product:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#delete_legacy_game_internationalization_v1_developer_products__developerProductId__name_description_language_codes__languageCode_
//
// [DELETE] /legacy-game-internationalization/v1/developer-products/{developerProductId}/name-description/language-codes/{languageCode}
func (s *LocalizationService) DeleteDeveloperProductLocalizedNameDescription(ctx context.Context, developerProductId, languageCode string, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, DeleteDeveloperProductLocalizedNameDescription.Compile(nil, developerProductId, languageCode), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type UpdateDeveloperProductNameDescriptionRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateDeveloperProductNameDescriptionResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateDeveloperProductLocalizedNameDescription will update localized name and description of a developer
// product.
//
// Required scopes: legacy-developer-product:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#patch_legacy_game_internationalization_v1_developer_products__developerProductId__name_description_language_codes__languageCode_
//
// [PATCH] /legacy-game-internationalization/v1/developer-products/{developerProductId}/name-description/language-codes/{languageCode}
func (s *LocalizationService) UpdateDeveloperProductLocalizedNameDescription(ctx context.Context, developerProductId, languageCode string, data UpdateDeveloperProductNameDescriptionRequest, reqOpts ...goblox.RequestOpt) (*UpdateDeveloperProductNameDescriptionResponse, *Response, error) {
	updateDeveloperProductNameDescriptionResponse := new(UpdateDeveloperProductNameDescriptionResponse)
	resp, err := s.client.Call(ctx, UpdateDeveloperProductLocalizedNameDescription.Compile(nil, developerProductId, languageCode), data, updateDeveloperProductNameDescriptionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateDeveloperProductNameDescriptionResponse, resp, nil
}

type UpdateDeveloperProductNameRequest struct {
	Name *string `json:"name,omitempty"`
}

type UpdateDeveloperProductNameResponse struct {
	Name string `json:"name"`
}

// UpdateDeveloperProductLocalizedName will update localized name of a developer product.
//
// Required scopes: legacy-developer-product:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#patch_legacy_game_internationalization_v1_developer_products__developerProductId__name_language_codes__languageCode_
//
// [PATCH] /legacy-game-internationalization/v1/developer-products/{developerProductId}/name/language-codes/{languageCode}
func (s *LocalizationService) UpdateDeveloperProductLocalizedName(ctx context.Context, developerProductId, languageCode string, data UpdateDeveloperProductNameRequest, reqOpts ...goblox.RequestOpt) (*UpdateDeveloperProductNameResponse, *Response, error) {
	updateDeveloperProductNameResponse := new(UpdateDeveloperProductNameResponse)
	resp, err := s.client.Call(ctx, UpdateDeveloperProductLocalizedName.Compile(nil, developerProductId, languageCode), data, updateDeveloperProductNameResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateDeveloperProductNameResponse, resp, nil
}

type ListGameLocalizedIconsOptions struct {
	Width  *int `url:"width,omitempty"`
	Height *int `url:"height,omitempty"`
}

type GetGameIconResponseState string

const (
	GetGameIconResponseStateApproved      GetGameIconResponseState = "Approved"
	GetGameIconResponseStatePendingReview GetGameIconResponseState = "PendingReview"
	GetGameIconResponseStateUnAvailable   GetGameIconResponseState = "UnAvailable"
	GetGameIconResponseStateRejected      GetGameIconResponseState = "Rejected"
	GetGameIconResponseStateError         GetGameIconResponseState = "Error"
)

type GetGameIconResponse struct {
	ImageID      string                   `json:"imageId"`
	ImageURL     string                   `json:"imageUrl"`
	State        GetGameIconResponseState `json:"state"`
	LanguageCode string                   `json:"languageCode"`
}

type GetGameIconList struct {
	Data []GetGameIconResponse `json:"data"`
}

// ListGameLocalizedIcons will get all icons for a game.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#get_legacy_game_internationalization_v1_game_icon_games__gameId_
//
// [GET] /legacy-game-internationalization/v1/game-icon/games/{gameId}
func (s *LocalizationService) ListGameLocalizedIcons(ctx context.Context, gameId string, opts *ListGameLocalizedIconsOptions, reqOpts ...goblox.RequestOpt) (*GetGameIconList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getGameIconList := new(GetGameIconList)
	resp, err := s.client.Call(ctx, ListGameLocalizedIcons.Compile(q, gameId), nil, getGameIconList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getGameIconList, resp, nil
}

// DeleteGameLocalizedIcon will delete a localized icon from a game.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#delete_legacy_game_internationalization_v1_game_icon_games__gameId__language_codes__languageCode_
//
// [DELETE] /legacy-game-internationalization/v1/game-icon/games/{gameId}/language-codes/{languageCode}
func (s *LocalizationService) DeleteGameLocalizedIcon(ctx context.Context, gameId, languageCode string, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, DeleteGameLocalizedIcon.Compile(nil, gameId, languageCode), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type UpdateGameLocalizedIconRequest struct {
	Files *os.File `form:"Files"`
}

// UpdateGameLocalizedIcon will update a game's icon.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#post_legacy_game_internationalization_v1_game_icon_games__gameId__language_codes__languageCode_
//
// [POST] /legacy-game-internationalization/v1/game-icon/games/{gameId}/language-codes/{languageCode}
func (s *LocalizationService) UpdateGameLocalizedIcon(ctx context.Context, gameId, languageCode string, data UpdateGameLocalizedIconRequest, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	e := UpdateGameLocalizedIcon.Compile(nil, gameId, languageCode)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Files != nil {
		part, err := writer.CreateFormFile("Files", filepath.Base(data.Files.Name()))
		if err != nil {
			return nil, nil, err
		}
		if _, err := io.Copy(part, data.Files); err != nil {
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

	resp, err := s.client.CallRequest(ctx, e, req, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type UpdateGamePassDescriptionRequest struct {
	Description *string `json:"description,omitempty"`
}

type UpdateGamePassDescriptionResponse struct {
	Description string `json:"description"`
}

// UpdateGamePassLocalizedDescription will update localized description of a game pass.
//
// Required scopes: legacy-game-pass:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#patch_legacy_game_internationalization_v1_game_passes__gamePassId__description_language_codes__languageCode_
//
// [PATCH] /legacy-game-internationalization/v1/game-passes/{gamePassId}/description/language-codes/{languageCode}
func (s *LocalizationService) UpdateGamePassLocalizedDescription(ctx context.Context, gamePassId, languageCode string, data UpdateGamePassDescriptionRequest, reqOpts ...goblox.RequestOpt) (*UpdateGamePassDescriptionResponse, *Response, error) {
	updateGamePassDescriptionResponse := new(UpdateGamePassDescriptionResponse)
	resp, err := s.client.Call(ctx, UpdateGamePassLocalizedDescription.Compile(nil, gamePassId, languageCode), data, updateGamePassDescriptionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateGamePassDescriptionResponse, resp, nil
}

type ListGamePassLocalizedIconsOptions struct {
	Width  *int `url:"width,omitempty"`
	Height *int `url:"height,omitempty"`
}

type GetGamePassIconResponseState string

const (
	GetGamePassIconResponseStateApproved      GetGamePassIconResponseState = "Approved"
	GetGamePassIconResponseStatePendingReview GetGamePassIconResponseState = "PendingReview"
	GetGamePassIconResponseStateUnAvailable   GetGamePassIconResponseState = "UnAvailable"
	GetGamePassIconResponseStateRejected      GetGamePassIconResponseState = "Rejected"
	GetGamePassIconResponseStateError         GetGamePassIconResponseState = "Error"
)

type GetGamePassIconResponse struct {
	ImageID      string                       `json:"imageId"`
	ImageURL     string                       `json:"imageUrl"`
	State        GetGamePassIconResponseState `json:"state"`
	LanguageCode string                       `json:"languageCode"`
}

type GetGamePassIconList struct {
	Data []GetGamePassIconResponse `json:"data"`
}

// ListGamePassLocalizedIcons will get all icons for a game pass.
//
// Required scopes: legacy-game-pass:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#get_legacy_game_internationalization_v1_game_passes__gamePassId__icons
//
// [GET] /legacy-game-internationalization/v1/game-passes/{gamePassId}/icons
func (s *LocalizationService) ListGamePassLocalizedIcons(ctx context.Context, gamePassId string, opts *ListGamePassLocalizedIconsOptions, reqOpts ...goblox.RequestOpt) (*GetGamePassIconList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getGamePassIconList := new(GetGamePassIconList)
	resp, err := s.client.Call(ctx, ListGamePassLocalizedIcons.Compile(q, gamePassId), nil, getGamePassIconList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getGamePassIconList, resp, nil
}

// DeleteGamePassLocalizedIcon will delete a localized icon from a game pass.
//
// Required scopes: legacy-game-pass:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#delete_legacy_game_internationalization_v1_game_passes__gamePassId__icons_language_codes__languageCode_
//
// [DELETE] /legacy-game-internationalization/v1/game-passes/{gamePassId}/icons/language-codes/{languageCode}
func (s *LocalizationService) DeleteGamePassLocalizedIcon(ctx context.Context, gamePassId, languageCode string, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, DeleteGamePassLocalizedIcon.Compile(nil, gamePassId, languageCode), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type UpdateGamePassLocalizedIconRequest struct {
	Files *os.File `form:"Files"`
}

// UpdateGamePassLocalizedIcon will update a game pass's icon.
//
// Required scopes: legacy-game-pass:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#post_legacy_game_internationalization_v1_game_passes__gamePassId__icons_language_codes__languageCode_
//
// [POST] /legacy-game-internationalization/v1/game-passes/{gamePassId}/icons/language-codes/{languageCode}
func (s *LocalizationService) UpdateGamePassLocalizedIcon(ctx context.Context, gamePassId, languageCode string, data UpdateGamePassLocalizedIconRequest, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	e := UpdateGamePassLocalizedIcon.Compile(nil, gamePassId, languageCode)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Files != nil {
		part, err := writer.CreateFormFile("Files", filepath.Base(data.Files.Name()))
		if err != nil {
			return nil, nil, err
		}
		if _, err := io.Copy(part, data.Files); err != nil {
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

	resp, err := s.client.CallRequest(ctx, e, req, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

// ListGamePassLocalizedNameDescriptions will get all names and descriptions of a game pass.
//
// Required scopes: legacy-game-pass:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#get_legacy_game_internationalization_v1_game_passes__gamePassId__name_description
//
// [GET] /legacy-game-internationalization/v1/game-passes/{gamePassId}/name-description
func (s *LocalizationService) ListGamePassLocalizedNameDescriptions(ctx context.Context, gamePassId string, reqOpts ...goblox.RequestOpt) (*NameDescriptionList, *Response, error) {
	nameDescriptionList := new(NameDescriptionList)
	resp, err := s.client.Call(ctx, ListGamePassLocalizedNameDescriptions.Compile(nil, gamePassId), nil, nameDescriptionList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return nameDescriptionList, resp, nil
}

// DeleteGamePassLocalizedNameDescription will delete localized name and description of a game pass.
//
// Required scopes: legacy-game-pass:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#delete_legacy_game_internationalization_v1_game_passes__gamePassId__name_description_language_codes__languageCode_
//
// [DELETE] /legacy-game-internationalization/v1/game-passes/{gamePassId}/name-description/language-codes/{languageCode}
func (s *LocalizationService) DeleteGamePassLocalizedNameDescription(ctx context.Context, gamePassId, languageCode string, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, DeleteGamePassLocalizedNameDescription.Compile(nil, gamePassId, languageCode), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type UpdateGamePassNameDescriptionRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateGamePassNameDescriptionResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateGamePassLocalizedNameDescription will update localized name and description of a game pass.
//
// Required scopes: legacy-game-pass:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#patch_legacy_game_internationalization_v1_game_passes__gamePassId__name_description_language_codes__languageCode_
//
// [PATCH] /legacy-game-internationalization/v1/game-passes/{gamePassId}/name-description/language-codes/{languageCode}
func (s *LocalizationService) UpdateGamePassLocalizedNameDescription(ctx context.Context, gamePassId, languageCode string, data UpdateGamePassNameDescriptionRequest, reqOpts ...goblox.RequestOpt) (*UpdateGamePassNameDescriptionResponse, *Response, error) {
	updateGamePassNameDescriptionResponse := new(UpdateGamePassNameDescriptionResponse)
	resp, err := s.client.Call(ctx, UpdateGamePassLocalizedNameDescription.Compile(nil, gamePassId, languageCode), data, updateGamePassNameDescriptionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateGamePassNameDescriptionResponse, resp, nil
}

type UpdateGamePassNameRequest struct {
	Name *string `json:"name,omitempty"`
}

type UpdateGamePassNameResponse struct {
	Name string `json:"name"`
}

// UpdateGamePassLocalizedName will update localized name of a game pass.
//
// Required scopes: legacy-game-pass:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#patch_legacy_game_internationalization_v1_game_passes__gamePassId__name_language_codes__languageCode_
//
// [PATCH] /legacy-game-internationalization/v1/game-passes/{gamePassId}/name/language-codes/{languageCode}
func (s *LocalizationService) UpdateGamePassLocalizedName(ctx context.Context, gamePassId, languageCode string, data UpdateGamePassNameRequest, reqOpts ...goblox.RequestOpt) (*UpdateGamePassNameResponse, *Response, error) {
	updateGamePassNameResponse := new(UpdateGamePassNameResponse)
	resp, err := s.client.Call(ctx, UpdateGamePassLocalizedName.Compile(nil, gamePassId, languageCode), data, updateGamePassNameResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateGamePassNameResponse, resp, nil
}

type UpdateThumbnailAltTextRequest struct {
	ThumbnailID *int64  `json:"thumbnailId,omitempty"`
	AltText     *string `json:"altText,omitempty"`
}

// UpdateGameThumbnailAltText will update the game thumbnail alt text.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#post_legacy_game_internationalization_v1_game_thumbnails_games__gameId__language_codes__languageCode__alt_text
//
// [POST] /legacy-game-internationalization/v1/game-thumbnails/games/{gameId}/language-codes/{languageCode}/alt-text
func (s *LocalizationService) UpdateGameThumbnailAltText(ctx context.Context, gameId, languageCode string, data UpdateThumbnailAltTextRequest, reqOpts ...goblox.RequestOpt) (string, *Response, error) {
	var result string
	resp, err := s.client.Call(ctx, UpdateGameThumbnailAltText.Compile(nil, gameId, languageCode), data, &result, reqOpts...)
	if err != nil {
		return "", resp, err
	}

	return result, resp, nil
}

type UploadGameThumbnailRequest struct {
	Files *os.File `form:"Files"`
}

type UploadImageForGameThumbnailResponse struct {
	MediaAssetID string `json:"mediaAssetId"`
}

// UploadGameThumbnail will upload the game thumbnail.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#post_legacy_game_internationalization_v1_game_thumbnails_games__gameId__language_codes__languageCode__image
//
// [POST] /legacy-game-internationalization/v1/game-thumbnails/games/{gameId}/language-codes/{languageCode}/image
func (s *LocalizationService) UploadGameThumbnail(ctx context.Context, gameId, languageCode string, data UploadGameThumbnailRequest, reqOpts ...goblox.RequestOpt) (*UploadImageForGameThumbnailResponse, *Response, error) {
	uploadImageForGameThumbnailResponse := new(UploadImageForGameThumbnailResponse)
	e := UploadGameThumbnail.Compile(nil, gameId, languageCode)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Files != nil {
		part, err := writer.CreateFormFile("Files", filepath.Base(data.Files.Name()))
		if err != nil {
			return nil, nil, err
		}
		if _, err := io.Copy(part, data.Files); err != nil {
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

	resp, err := s.client.CallRequest(ctx, e, req, uploadImageForGameThumbnailResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return uploadImageForGameThumbnailResponse, resp, nil
}

type SortImageIDsRequest struct {
	MediaAssetIDs []int64 `json:"mediaAssetIds,omitempty"`
}

// OrderGameThumbnails will order the specified image Ids for the game thumbnails.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#post_legacy_game_internationalization_v1_game_thumbnails_games__gameId__language_codes__languageCode__images_order
//
// [POST] /legacy-game-internationalization/v1/game-thumbnails/games/{gameId}/language-codes/{languageCode}/images/order
func (s *LocalizationService) OrderGameThumbnails(ctx context.Context, gameId, languageCode string, data SortImageIDsRequest, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, OrderGameThumbnails.Compile(nil, gameId, languageCode), data, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

// DeleteGameThumbnail will delete the game thumbnail.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#delete_legacy_game_internationalization_v1_game_thumbnails_games__gameId__language_codes__languageCode__images__imageId_
//
// [DELETE] /legacy-game-internationalization/v1/game-thumbnails/games/{gameId}/language-codes/{languageCode}/images/{imageId}
func (s *LocalizationService) DeleteGameThumbnail(ctx context.Context, gameId, languageCode, imageId string, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, DeleteGameThumbnail.Compile(nil, gameId, languageCode, imageId), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type GetNameDescriptionHistoryV2RequestContentType string

const (
	GetNameDescriptionHistoryV2RequestContentTypeUniverseDisplayNames               GetNameDescriptionHistoryV2RequestContentType = "UniverseDisplayNames"
	GetNameDescriptionHistoryV2RequestContentTypeUniverseDisplayDescriptions        GetNameDescriptionHistoryV2RequestContentType = "UniverseDisplayDescriptions"
	GetNameDescriptionHistoryV2RequestContentTypeBadgeDisplayName                   GetNameDescriptionHistoryV2RequestContentType = "BadgeDisplayName"
	GetNameDescriptionHistoryV2RequestContentTypeBadgeDisplayDescription            GetNameDescriptionHistoryV2RequestContentType = "BadgeDisplayDescription"
	GetNameDescriptionHistoryV2RequestContentTypeDeveloperProductDisplayName        GetNameDescriptionHistoryV2RequestContentType = "DeveloperProductDisplayName"
	GetNameDescriptionHistoryV2RequestContentTypeDeveloperProductDisplayDescription GetNameDescriptionHistoryV2RequestContentType = "DeveloperProductDisplayDescription"
	GetNameDescriptionHistoryV2RequestContentTypeGamePassDisplayName                GetNameDescriptionHistoryV2RequestContentType = "GamePassDisplayName"
	GetNameDescriptionHistoryV2RequestContentTypeGamePassDisplayDescription         GetNameDescriptionHistoryV2RequestContentType = "GamePassDisplayDescription"
)

type GetNameDescriptionHistoryV2RequestSortOrder string

const (
	GetNameDescriptionHistoryV2RequestSortOrderAsc  GetNameDescriptionHistoryV2RequestSortOrder = "Asc"
	GetNameDescriptionHistoryV2RequestSortOrderDesc GetNameDescriptionHistoryV2RequestSortOrder = "Desc"
)

type GetNameDescriptionHistoryV2Request struct {
	ContentID    *int64                                         `json:"contentId,omitempty"`
	ContentType  *GetNameDescriptionHistoryV2RequestContentType `json:"contentType,omitempty"`
	LanguageCode *string                                        `json:"languageCode,omitempty"`
	Cursor       *string                                        `json:"cursor,omitempty"`
	Count        *int                                           `json:"count,omitempty"`
	SortOrder    *GetNameDescriptionHistoryV2RequestSortOrder   `json:"sortOrder,omitempty"`
}

type LocalizationTranslatorAgentType string

const (
	LocalizationTranslatorAgentTypeUser       LocalizationTranslatorAgentType = "User"
	LocalizationTranslatorAgentTypeAutomation LocalizationTranslatorAgentType = "Automation"
)

type LocalizationTranslator struct {
	ID        int64                           `json:"id"`
	AgentType LocalizationTranslatorAgentType `json:"agentType"`
}

type LocalizationTranslationHistory struct {
	TranslationText string                 `json:"translationText"`
	Translator      LocalizationTranslator `json:"translator"`
	Created         string                 `json:"created"`
}

type GetNameDescriptionHistoryResponse struct {
	History         []LocalizationTranslationHistory `json:"history"`
	LastEvaluatedID string                           `json:"lastEvaluatedId"`
}

// GetGameNameDescriptionHistory will get the history for name or description in a provided language.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#post_legacy_game_internationalization_v1_name_description_games_translation_history
//
// [POST] /legacy-game-internationalization/v1/name-description/games/translation-history
func (s *LocalizationService) GetGameNameDescriptionHistory(ctx context.Context, data GetNameDescriptionHistoryV2Request, reqOpts ...goblox.RequestOpt) (*GetNameDescriptionHistoryResponse, *Response, error) {
	getNameDescriptionHistoryResponse := new(GetNameDescriptionHistoryResponse)
	resp, err := s.client.Call(ctx, GetGameNameDescriptionHistory.Compile(nil), data, getNameDescriptionHistoryResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getNameDescriptionHistoryResponse, resp, nil
}

type UpdateNameDescriptionsRequest struct {
	Data []LocalizationNameDescription `json:"data,omitempty"`
}

type FailedNameDescription struct {
	LanguageCode string `json:"languageCode"`
	ErrorCode    int    `json:"errorCode"`
}

type UpdateNameDescriptionsResponse struct {
	SuccessOperations []LocalizationNameDescription `json:"successOperations"`
	FailedOperations  []FailedNameDescription       `json:"failedOperations"`
}

// UpdateGameNameDescriptions will update a game's name and/or description in multiple languages.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#patch_legacy_game_internationalization_v1_name_description_games__gameId_
//
// [PATCH] /legacy-game-internationalization/v1/name-description/games/{gameId}
func (s *LocalizationService) UpdateGameNameDescriptions(ctx context.Context, gameId string, data UpdateNameDescriptionsRequest, reqOpts ...goblox.RequestOpt) (*UpdateNameDescriptionsResponse, *Response, error) {
	updateNameDescriptionsResponse := new(UpdateNameDescriptionsResponse)
	resp, err := s.client.Call(ctx, UpdateGameNameDescriptions.Compile(nil, gameId), data, updateNameDescriptionsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateNameDescriptionsResponse, resp, nil
}

type SetGameSourceLanguageOptions struct {
	LanguageCode string `url:"languageCode"` // Required.
}

// SetGameSourceLanguage will set the source language of a game.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#patch_legacy_game_internationalization_v1_source_language_games__gameId_
//
// [PATCH] /legacy-game-internationalization/v1/source-language/games/{gameId}
func (s *LocalizationService) SetGameSourceLanguage(ctx context.Context, gameId string, opts *SetGameSourceLanguageOptions, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, SetGameSourceLanguage.Compile(q, gameId), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type PatchLanguageLanguageCodeType string

const (
	PatchLanguageLanguageCodeTypeLanguage PatchLanguageLanguageCodeType = "Language"
	PatchLanguageLanguageCodeTypeLocale   PatchLanguageLanguageCodeType = "Locale"
)

type PatchLanguage struct {
	LanguageCodeType *PatchLanguageLanguageCodeType `json:"languageCodeType,omitempty"`
	LanguageCode     *string                        `json:"languageCode,omitempty"`
	Delete           *bool                          `json:"delete,omitempty"`
}

// UpdateGameSupportedLanguages will add or remove supported languages for a game.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#patch_legacy_game_internationalization_v1_supported_languages_games__gameId_
//
// [PATCH] /legacy-game-internationalization/v1/supported-languages/games/{gameId}
func (s *LocalizationService) UpdateGameSupportedLanguages(ctx context.Context, gameId string, data []PatchLanguage, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, UpdateGameSupportedLanguages.Compile(nil, gameId), data, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type LanguageOrLocaleSettingsLanguageCodeType string

const (
	LanguageOrLocaleSettingsLanguageCodeTypeLanguage LanguageOrLocaleSettingsLanguageCodeType = "Language"
	LanguageOrLocaleSettingsLanguageCodeTypeLocale   LanguageOrLocaleSettingsLanguageCodeType = "Locale"
)

type LanguageOrLocaleSettings struct {
	LanguageCodeType              LanguageOrLocaleSettingsLanguageCodeType `json:"languageCodeType"`
	LanguageCode                  string                                   `json:"languageCode"`
	IsAutomaticTranslationEnabled bool                                     `json:"isAutomaticTranslationEnabled"`
	IsImageTranslationEnabled     bool                                     `json:"isImageTranslationEnabled"`
}

type LanguageOrLocaleSettingsList struct {
	Data []LanguageOrLocaleSettings `json:"data"`
}

// GetGameAutomaticTranslationStatus will get the automatic translation status of supported languages for a game.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#get_legacy_game_internationalization_v1_supported_languages_games__gameId__automatic_translation_status
//
// [GET] /legacy-game-internationalization/v1/supported-languages/games/{gameId}/automatic-translation-status
func (s *LocalizationService) GetGameAutomaticTranslationStatus(ctx context.Context, gameId string, reqOpts ...goblox.RequestOpt) (*LanguageOrLocaleSettingsList, *Response, error) {
	languageOrLocaleSettingsList := new(LanguageOrLocaleSettingsList)
	resp, err := s.client.Call(ctx, GetGameAutomaticTranslationStatus.Compile(nil, gameId), nil, languageOrLocaleSettingsList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return languageOrLocaleSettingsList, resp, nil
}

type EditAutomaticTranslationStatusForGameAndLanguageResponse struct {
	GameID                        int64  `json:"gameId"`
	LanguageCode                  string `json:"languageCode"`
	IsAutomaticTranslationEnabled bool   `json:"isAutomaticTranslationEnabled"`
}

// UpdateGameLanguageAutomaticTranslationStatus will enable or disable automatic translation for a game and
// language.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#patch_legacy_game_internationalization_v1_supported_languages_games__gameId__languages__languageCode__automatic_translation_status
//
// [PATCH] /legacy-game-internationalization/v1/supported-languages/games/{gameId}/languages/{languageCode}/automatic-translation-status
func (s *LocalizationService) UpdateGameLanguageAutomaticTranslationStatus(ctx context.Context, gameId, languageCode string, data bool, reqOpts ...goblox.RequestOpt) (*EditAutomaticTranslationStatusForGameAndLanguageResponse, *Response, error) {
	editAutomaticTranslationStatusForGameAndLanguageResponse := new(EditAutomaticTranslationStatusForGameAndLanguageResponse)
	resp, err := s.client.Call(ctx, UpdateGameLanguageAutomaticTranslationStatus.Compile(nil, gameId, languageCode), data, editAutomaticTranslationStatusForGameAndLanguageResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return editAutomaticTranslationStatusForGameAndLanguageResponse, resp, nil
}

type EditImageTranslationStatusForGameAndLanguageResponse struct {
	GameID                    int64  `json:"gameId"`
	LanguageCode              string `json:"languageCode"`
	IsImageTranslationEnabled bool   `json:"isImageTranslationEnabled"`
}

// UpdateGameLanguageImageTranslationStatus will enable or disable image translation for a game and language.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#patch_legacy_game_internationalization_v1_supported_languages_games__gameId__languages__languageCode__image_translation_status
//
// [PATCH] /legacy-game-internationalization/v1/supported-languages/games/{gameId}/languages/{languageCode}/image-translation-status
func (s *LocalizationService) UpdateGameLanguageImageTranslationStatus(ctx context.Context, gameId, languageCode string, data bool, reqOpts ...goblox.RequestOpt) (*EditImageTranslationStatusForGameAndLanguageResponse, *Response, error) {
	editImageTranslationStatusForGameAndLanguageResponse := new(EditImageTranslationStatusForGameAndLanguageResponse)
	resp, err := s.client.Call(ctx, UpdateGameLanguageImageTranslationStatus.Compile(nil, gameId, languageCode), data, editImageTranslationStatusForGameAndLanguageResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return editImageTranslationStatusForGameAndLanguageResponse, resp, nil
}

type UpdateUniverseDisplayInfoAutomaticTranslationSettingsResponse struct {
	GameID                                           int64  `json:"gameId"`
	LanguageCode                                     string `json:"languageCode"`
	IsUniverseDisplayInfoAutomaticTranslationEnabled bool   `json:"isUniverseDisplayInfoAutomaticTranslationEnabled"`
}

// UpdateGameLanguageDisplayInfoAutomaticTranslation will update the switch which controls if the
// UniverseDisplayInformation should be automatically translated.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#patch_legacy_game_internationalization_v1_supported_languages_games__gameId__languages__languageCode__universe_display_info_automatic_translation_settings
//
// [PATCH] /legacy-game-internationalization/v1/supported-languages/games/{gameId}/languages/{languageCode}/universe-display-info-automatic-translation-settings
func (s *LocalizationService) UpdateGameLanguageDisplayInfoAutomaticTranslation(ctx context.Context, gameId, languageCode string, data bool, reqOpts ...goblox.RequestOpt) (*UpdateUniverseDisplayInfoAutomaticTranslationSettingsResponse, *Response, error) {
	updateUniverseDisplayInfoAutomaticTranslationSettingsResponse := new(UpdateUniverseDisplayInfoAutomaticTranslationSettingsResponse)
	resp, err := s.client.Call(ctx, UpdateGameLanguageDisplayInfoAutomaticTranslation.Compile(nil, gameId, languageCode), data, updateUniverseDisplayInfoAutomaticTranslationSettingsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateUniverseDisplayInfoAutomaticTranslationSettingsResponse, resp, nil
}

type UniverseDisplayInfoAutomaticTranslationSettings struct {
	LanguageCode                                     string `json:"languageCode"`
	IsUniverseDisplayInfoAutomaticTranslationEnabled bool   `json:"isUniverseDisplayInfoAutomaticTranslationEnabled"`
}

type UniverseDisplayInfoAutomaticTranslationSettingsList struct {
	Data []UniverseDisplayInfoAutomaticTranslationSettings `json:"data"`
}

// GetGameDisplayInfoAutomaticTranslationSettings will get UniverseDisplayInfo automatic translation settings.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#get_legacy_game_internationalization_v1_supported_languages_games__gameId__universe_display_info_automatic_translation_settings
//
// [GET] /legacy-game-internationalization/v1/supported-languages/games/{gameId}/universe-display-info-automatic-translation-settings
func (s *LocalizationService) GetGameDisplayInfoAutomaticTranslationSettings(ctx context.Context, gameId string, reqOpts ...goblox.RequestOpt) (*UniverseDisplayInfoAutomaticTranslationSettingsList, *Response, error) {
	universeDisplayInfoAutomaticTranslationSettingsList := new(UniverseDisplayInfoAutomaticTranslationSettingsList)
	resp, err := s.client.Call(ctx, GetGameDisplayInfoAutomaticTranslationSettings.Compile(nil, gameId), nil, universeDisplayInfoAutomaticTranslationSettingsList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return universeDisplayInfoAutomaticTranslationSettingsList, resp, nil
}

type GameAutolocalizationInformationResponse struct {
	IsAutolocalizationEnabled         bool   `json:"isAutolocalizationEnabled"`
	IsAutomaticEntriesSettingEnabled  bool   `json:"isAutomaticEntriesSettingEnabled"`
	IsAutomaticEntriesDeletionEnabled bool   `json:"isAutomaticEntriesDeletionEnabled"`
	ShouldUseLocalizationTable        bool   `json:"shouldUseLocalizationTable"`
	AutoLocalizationTableID           string `json:"autoLocalizationTableId"`
	SourceLanguage                    string `json:"sourceLanguage"`
	AssetID                           int64  `json:"assetId"`
	ShouldUseImageLocalizationTable   bool   `json:"shouldUseImageLocalizationTable"`
	IsAutoLocalizationForImageEnabled bool   `json:"isAutoLocalizationForImageEnabled"`
}

// CreateGameAutoLocalizationTable will create an auto-localization table for a game.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#post_legacy_localization_tables_v1_autolocalization_games__gameId__autolocalizationtable
//
// [POST] /legacy-localization-tables/v1/autolocalization/games/{gameId}/autolocalizationtable
func (s *LocalizationService) CreateGameAutoLocalizationTable(ctx context.Context, gameId string, reqOpts ...goblox.RequestOpt) (*GameAutolocalizationInformationResponse, *Response, error) {
	gameAutolocalizationInformationResponse := new(GameAutolocalizationInformationResponse)
	resp, err := s.client.Call(ctx, CreateGameAutoLocalizationTable.Compile(nil, gameId), nil, gameAutolocalizationInformationResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return gameAutolocalizationInformationResponse, resp, nil
}

type SetAutolocalizationSettingsForGameRequest struct {
	IsAutolocalizationEnabled          *bool `json:"isAutolocalizationEnabled,omitempty"`
	IsAutomaticEntriesSettingEnabled   *bool `json:"isAutomaticEntriesSettingEnabled,omitempty"`
	IsAutomaticEntriesDeletionsEnabled *bool `json:"isAutomaticEntriesDeletionsEnabled,omitempty"`
	ShouldUseLocalizationTable         *bool `json:"shouldUseLocalizationTable,omitempty"`
	ShouldUseImageLocalizationTable    *bool `json:"shouldUseImageLocalizationTable,omitempty"`
	IsAutoLocalizationForImageEnabled  *bool `json:"isAutoLocalizationForImageEnabled,omitempty"`
}

// UpdateGameAutoLocalizationSettings will set a game's auto-localization related settings.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#patch_legacy_localization_tables_v1_autolocalization_games__gameId__settings
//
// [PATCH] /legacy-localization-tables/v1/autolocalization/games/{gameId}/settings
func (s *LocalizationService) UpdateGameAutoLocalizationSettings(ctx context.Context, gameId string, data SetAutolocalizationSettingsForGameRequest, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, UpdateGameAutoLocalizationSettings.Compile(nil, gameId), data, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type AutoLocalizationMetadataResponse struct {
	IsReactVersionEnabledForAutoLocalizationSettings bool `json:"isReactVersionEnabledForAutoLocalizationSettings"`
	IsTabbedUiEnabledForConfigureLocalizationPage    bool `json:"isTabbedUIEnabledForConfigureLocalizationPage"`
	IsAutomaticTranslationToggleUiEnabled            bool `json:"isAutomaticTranslationToggleUIEnabled"`
	IsAutomaticTranslationQuotaUiEnabled             bool `json:"isAutomaticTranslationQuotaUIEnabled"`
}

// GetAutoLocalizationMetadata will metadata for AutoLocalization Configuration.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#get_legacy_localization_tables_v1_autolocalization_metadata
//
// [GET] /legacy-localization-tables/v1/autolocalization/metadata
func (s *LocalizationService) GetAutoLocalizationMetadata(ctx context.Context, reqOpts ...goblox.RequestOpt) (*AutoLocalizationMetadataResponse, *Response, error) {
	autoLocalizationMetadataResponse := new(AutoLocalizationMetadataResponse)
	resp, err := s.client.Call(ctx, GetAutoLocalizationMetadata.Compile(nil), nil, autoLocalizationMetadataResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return autoLocalizationMetadataResponse, resp, nil
}

type EntryOperationLimits struct {
	MaxContextLength          int `json:"maxContextLength"`
	MaxKeyLength              int `json:"maxKeyLength"`
	MaxSourceLength           int `json:"maxSourceLength"`
	MaxExampleLength          int `json:"maxExampleLength"`
	MaxGameLocationPathLength int `json:"maxGameLocationPathLength"`
}

type TableOperationLimits struct {
	MaxEntriesPerUpdate int `json:"maxEntriesPerUpdate"`
}

type GetLimitsResponse struct {
	EntryOperationLimits EntryOperationLimits `json:"entryOperationLimits"`
	TableOperationLimits TableOperationLimits `json:"tableOperationLimits"`
}

// GetLocalizationTableLimits will get limits for translation table entries operations.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#get_legacy_localization_tables_v1_localization_table_limits
//
// [GET] /legacy-localization-tables/v1/localization-table/limits
func (s *LocalizationService) GetLocalizationTableLimits(ctx context.Context, reqOpts ...goblox.RequestOpt) (*GetLimitsResponse, *Response, error) {
	getLimitsResponse := new(GetLimitsResponse)
	resp, err := s.client.Call(ctx, GetLocalizationTableLimits.Compile(nil), nil, getLimitsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getLimitsResponse, resp, nil
}

type GetTableResponseOwnerType string

const (
	GetTableResponseOwnerTypeUser  GetTableResponseOwnerType = "User"
	GetTableResponseOwnerTypeGroup GetTableResponseOwnerType = "Group"
)

type GetTableResponse struct {
	ID        string                    `json:"id"`
	Name      string                    `json:"name"`
	OwnerType GetTableResponseOwnerType `json:"ownerType"`
	OwnerID   int64                     `json:"ownerId"`
	AssetID   int64                     `json:"assetId"`
}

// GetLocalizationTableByAssetID will get table information by the assetId of the table.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/assets#get_legacy_localization_tables_v1_localization_table_tables__assetId_
//
// [GET] /legacy-localization-tables/v1/localization-table/tables/{assetId}
func (s *LocalizationService) GetLocalizationTableByAssetID(ctx context.Context, assetId string, reqOpts ...goblox.RequestOpt) (*GetTableResponse, *Response, error) {
	getTableResponse := new(GetTableResponse)
	resp, err := s.client.Call(ctx, GetLocalizationTableByAssetID.Compile(nil, assetId), nil, getTableResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getTableResponse, resp, nil
}

// GetLocalizationTable will get table information by the id of the table.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#get_legacy_localization_tables_v1_localization_table_tables__tableId_
//
// [GET] /legacy-localization-tables/v1/localization-table/tables/{tableId}
func (s *LocalizationService) GetLocalizationTable(ctx context.Context, tableId string, reqOpts ...goblox.RequestOpt) (*GetTableResponse, *Response, error) {
	getTableResponse := new(GetTableResponse)
	resp, err := s.client.Call(ctx, GetLocalizationTable.Compile(nil, tableId), nil, getTableResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getTableResponse, resp, nil
}

type EntryIdentifierEntryFormat string

const (
	EntryIdentifierEntryFormatInvalid EntryIdentifierEntryFormat = "Invalid"
	EntryIdentifierEntryFormatLegacy  EntryIdentifierEntryFormat = "Legacy"
	EntryIdentifierEntryFormatIcu     EntryIdentifierEntryFormat = "Icu"
)

type EntryIdentifier struct {
	Key         *string                     `json:"key,omitempty"`
	Context     *string                     `json:"context,omitempty"`
	Source      *string                     `json:"source,omitempty"`
	EntryFormat *EntryIdentifierEntryFormat `json:"entryFormat,omitempty"`
}

type GameLocation struct {
	Path *string `json:"path,omitempty"`
}

type EntryMetadataEntryFormat string

const (
	EntryMetadataEntryFormatInvalid EntryMetadataEntryFormat = "Invalid"
	EntryMetadataEntryFormatLegacy  EntryMetadataEntryFormat = "Legacy"
	EntryMetadataEntryFormatIcu     EntryMetadataEntryFormat = "Icu"
)

type EntryMetadata struct {
	Example       *string                   `json:"example,omitempty"`
	GameLocations []GameLocation            `json:"gameLocations,omitempty"`
	EntryFormat   *EntryMetadataEntryFormat `json:"entryFormat,omitempty"`
}

type LocalizationChangeAgentChangeAgentType string

const (
	LocalizationChangeAgentChangeAgentTypeInvalid    LocalizationChangeAgentChangeAgentType = "Invalid"
	LocalizationChangeAgentChangeAgentTypeUser       LocalizationChangeAgentChangeAgentType = "User"
	LocalizationChangeAgentChangeAgentTypeAutomation LocalizationChangeAgentChangeAgentType = "Automation"
	LocalizationChangeAgentChangeAgentTypeDefault    LocalizationChangeAgentChangeAgentType = "Default"
)

type LocalizationChangeAgentOptionalIDCase string

const (
	LocalizationChangeAgentOptionalIDCaseNone LocalizationChangeAgentOptionalIDCase = "None"
	LocalizationChangeAgentOptionalIDCaseID   LocalizationChangeAgentOptionalIDCase = "Id"
)

type LocalizationChangeAgent struct {
	ChangeAgentType *LocalizationChangeAgentChangeAgentType `json:"ChangeAgentType,omitempty"`
	ID              *string                                 `json:"Id,omitempty"`
	OptionalIDCase  *LocalizationChangeAgentOptionalIDCase  `json:"OptionalIdCase,omitempty"`
}

type PatchTranslation struct {
	Locale          *string                  `json:"locale,omitempty"`
	TranslationText *string                  `json:"translationText,omitempty"`
	Delete          *bool                    `json:"delete,omitempty"`
	ChangeAgent     *LocalizationChangeAgent `json:"changeAgent,omitempty"`
	UpdatedTime     *string                  `json:"updatedTime,omitempty"`
}

type PatchEntry struct {
	Identifier   *EntryIdentifier   `json:"identifier,omitempty"`
	Metadata     *EntryMetadata     `json:"metadata,omitempty"`
	Translations []PatchTranslation `json:"translations,omitempty"`
	Delete       *bool              `json:"delete,omitempty"`
}

type UpdateTableContentsRequest struct {
	Name    *string      `json:"name,omitempty"`
	Entries []PatchEntry `json:"entries,omitempty"`
}

type UpdateLocalizationTableOptions struct {
	GameID *int64 `url:"gameId,omitempty"`
}

type LocalizationError struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type LocalizationTranslatorV2AgentType string

const (
	LocalizationTranslatorV2AgentTypeUser       LocalizationTranslatorV2AgentType = "User"
	LocalizationTranslatorV2AgentTypeAutomation LocalizationTranslatorV2AgentType = "Automation"
)

type LocalizationTranslatorV2 struct {
	ID        int64                             `json:"id"`
	AgentType LocalizationTranslatorV2AgentType `json:"agentType"`
}

type LocalizationTranslation struct {
	Locale          string                   `json:"locale"`
	TranslationText string                   `json:"translationText"`
	Translator      LocalizationTranslatorV2 `json:"translator"`
	UpdatedTime     string                   `json:"updatedTime"`
	FeedbackCount   int                      `json:"feedbackCount"`
}

type FailedEntry struct {
	Error        LocalizationError         `json:"error"`
	Identifier   EntryIdentifier           `json:"identifier"`
	Metadata     EntryMetadata             `json:"metadata"`
	Translations []LocalizationTranslation `json:"translations"`
	CreatedTime  string                    `json:"createdTime"`
}

type ModifiedEntry struct {
	Identifier   EntryIdentifier           `json:"identifier"`
	Translations []LocalizationTranslation `json:"translations"`
}

type UpdateTableContentsResponse struct {
	FailedEntriesAndTranslations   []FailedEntry   `json:"failedEntriesAndTranslations"`
	ModifiedEntriesAndTranslations []ModifiedEntry `json:"modifiedEntriesAndTranslations"`
}

// UpdateLocalizationTable will update the tables contents based on what is provided.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#patch_legacy_localization_tables_v1_localization_table_tables__tableId_
//
// [PATCH] /legacy-localization-tables/v1/localization-table/tables/{tableId}
func (s *LocalizationService) UpdateLocalizationTable(ctx context.Context, tableId string, data UpdateTableContentsRequest, opts *UpdateLocalizationTableOptions, reqOpts ...goblox.RequestOpt) (*UpdateTableContentsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	updateTableContentsResponse := new(UpdateTableContentsResponse)
	resp, err := s.client.Call(ctx, UpdateLocalizationTable.Compile(q, tableId), data, updateTableContentsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateTableContentsResponse, resp, nil
}

type ListLocalizationTableEntriesOptionsEntryFormat string

const (
	ListLocalizationTableEntriesOptionsEntryFormatInvalid ListLocalizationTableEntriesOptionsEntryFormat = "Invalid"
	ListLocalizationTableEntriesOptionsEntryFormatLegacy  ListLocalizationTableEntriesOptionsEntryFormat = "Legacy"
	ListLocalizationTableEntriesOptionsEntryFormatIcu     ListLocalizationTableEntriesOptionsEntryFormat = "Icu"
)

type ListLocalizationTableEntriesOptions struct {
	Cursor      *string                                         `url:"cursor,omitempty"`
	GameID      *int64                                          `url:"gameId,omitempty"`
	EntryFormat *ListLocalizationTableEntriesOptionsEntryFormat `url:"entryFormat,omitempty"`
}

type LocalizationEntry struct {
	Identifier   EntryIdentifier           `json:"identifier"`
	Metadata     EntryMetadata             `json:"metadata"`
	Translations []LocalizationTranslation `json:"translations"`
	CreatedTime  string                    `json:"createdTime"`
}

type GetTableEntriesPagedResponse struct {
	PreviousPageCursor string              `json:"previousPageCursor"`
	NextPageCursor     string              `json:"nextPageCursor"`
	Data               []LocalizationEntry `json:"data"`
}

// ListLocalizationTableEntries will get a batch of entries for a table.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#get_legacy_localization_tables_v1_localization_table_tables__tableId__entries
//
// [GET] /legacy-localization-tables/v1/localization-table/tables/{tableId}/entries
func (s *LocalizationService) ListLocalizationTableEntries(ctx context.Context, tableId string, opts *ListLocalizationTableEntriesOptions, reqOpts ...goblox.RequestOpt) (*GetTableEntriesPagedResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getTableEntriesPagedResponse := new(GetTableEntriesPagedResponse)
	resp, err := s.client.Call(ctx, ListLocalizationTableEntries.Compile(q, tableId), nil, getTableEntriesPagedResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getTableEntriesPagedResponse, resp, nil
}

type CursorEntryIdentifierSortOrder string

const (
	CursorEntryIdentifierSortOrderAsc  CursorEntryIdentifierSortOrder = "Asc"
	CursorEntryIdentifierSortOrderDesc CursorEntryIdentifierSortOrder = "Desc"
)

type CursorEntryIdentifier struct {
	Cursor     *string                         `json:"cursor,omitempty"`
	Identifier *EntryIdentifier                `json:"identifier,omitempty"`
	Count      *int                            `json:"count,omitempty"`
	SortOrder  *CursorEntryIdentifierSortOrder `json:"sortOrder,omitempty"`
}

type GetTableEntriesTranslationHistoryRequestSourceType string

const (
	GetTableEntriesTranslationHistoryRequestSourceTypeInvalid    GetTableEntriesTranslationHistoryRequestSourceType = "Invalid"
	GetTableEntriesTranslationHistoryRequestSourceTypeText       GetTableEntriesTranslationHistoryRequestSourceType = "Text"
	GetTableEntriesTranslationHistoryRequestSourceTypeAssetImage GetTableEntriesTranslationHistoryRequestSourceType = "AssetImage"
)

type GetTableEntriesTranslationHistoryRequest struct {
	Locale     *string                                             `json:"locale,omitempty"`
	Entries    []CursorEntryIdentifier                             `json:"entries,omitempty"`
	SourceType *GetTableEntriesTranslationHistoryRequestSourceType `json:"sourceType,omitempty"`
}

type GetLocalizationTableTranslationHistoryOptions struct {
	GameID *int64 `url:"gameId,omitempty"`
}

type LocalizationTranslationHistoryV2 struct {
	TranslationText string                   `json:"translationText"`
	Translator      LocalizationTranslatorV2 `json:"translator"`
	Created         string                   `json:"created"`
}

type EntryTranslationHistoryPaged struct {
	Identifier EntryIdentifier                    `json:"identifier"`
	History    []LocalizationTranslationHistoryV2 `json:"history"`
	NextCursor string                             `json:"nextCursor"`
}

type FailedEntryTranslationHistoryPaged struct {
	Identifier EntryIdentifier   `json:"identifier"`
	Count      int               `json:"count"`
	Error      LocalizationError `json:"error"`
}

type GetTableEntriesTranslationHistoryResponse struct {
	TableID       string                               `json:"tableId"`
	Locale        string                               `json:"locale"`
	Entries       []EntryTranslationHistoryPaged       `json:"entries"`
	FailedEntries []FailedEntryTranslationHistoryPaged `json:"failedEntries"`
}

// GetLocalizationTableTranslationHistory will get the translation history for each entry passed in.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#post_legacy_localization_tables_v1_localization_table_tables__tableId__entries_translation_history
//
// [POST] /legacy-localization-tables/v1/localization-table/tables/{tableId}/entries/translation-history
func (s *LocalizationService) GetLocalizationTableTranslationHistory(ctx context.Context, tableId string, data GetTableEntriesTranslationHistoryRequest, opts *GetLocalizationTableTranslationHistoryOptions, reqOpts ...goblox.RequestOpt) (*GetTableEntriesTranslationHistoryResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getTableEntriesTranslationHistoryResponse := new(GetTableEntriesTranslationHistoryResponse)
	resp, err := s.client.Call(ctx, GetLocalizationTableTranslationHistory.Compile(q, tableId), data, getTableEntriesTranslationHistoryResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getTableEntriesTranslationHistoryResponse, resp, nil
}

type GetLocalizationTableEntryCountOptionsEntryFormat string

const (
	GetLocalizationTableEntryCountOptionsEntryFormatInvalid GetLocalizationTableEntryCountOptionsEntryFormat = "Invalid"
	GetLocalizationTableEntryCountOptionsEntryFormatLegacy  GetLocalizationTableEntryCountOptionsEntryFormat = "Legacy"
	GetLocalizationTableEntryCountOptionsEntryFormatIcu     GetLocalizationTableEntryCountOptionsEntryFormat = "Icu"
)

type GetLocalizationTableEntryCountOptions struct {
	GameID      *int64                                            `url:"gameId,omitempty"`
	EntryFormat *GetLocalizationTableEntryCountOptionsEntryFormat `url:"entryFormat,omitempty"`
}

type GetTableEntryCountResponse struct {
	ID         string `json:"id"`
	EntryCount int    `json:"entryCount"`
}

// GetLocalizationTableEntryCount will get the number of entries in the specified table.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/localization#get_legacy_localization_tables_v1_localization_table_tables__tableId__entry_count
//
// [GET] /legacy-localization-tables/v1/localization-table/tables/{tableId}/entry-count
func (s *LocalizationService) GetLocalizationTableEntryCount(ctx context.Context, tableId string, opts *GetLocalizationTableEntryCountOptions, reqOpts ...goblox.RequestOpt) (*GetTableEntryCountResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getTableEntryCountResponse := new(GetTableEntryCountResponse)
	resp, err := s.client.Call(ctx, GetLocalizationTableEntryCount.Compile(q, tableId), nil, getTableEntryCountResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getTableEntryCountResponse, resp, nil
}
