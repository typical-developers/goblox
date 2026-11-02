package opencloud

import (
	"context"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	GetUniverse = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}",
		auth.KindAPIKey, auth.KindOAuth2,
	)

	UpdateUniverse = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/universes/{universeId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:write")

	PublishUniverseMessage = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}:publishMessage",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe-messaging-service:publish")

	RestartUniverseServers = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}:restartServers",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:write")
)

// UniversesService handles communication with universe endpoints.
type UniversesService service

type UniverseVisibility string

const (
	UniverseVisibilityUnspecified UniverseVisibility = "VISIBILITY_UNSPECIFIED"
	UniverseVisibilityPublic      UniverseVisibility = "PUBLIC"
	UniverseVisibilityPrivate     UniverseVisibility = "PRIVATE"
)

type UniverseAgeRating string

const (
	UniverseAgeRatingUnspecified UniverseAgeRating = "AGE_RATING_UNSPECIFIED"
	UniverseAgeRatingAll         UniverseAgeRating = "AGE_RATING_ALL"
	UniverseAgeRating9Plus       UniverseAgeRating = "AGE_RATING_9_PLUS"
	UniverseAgeRating13Plus      UniverseAgeRating = "AGE_RATING_13_PLUS"
	UniverseAgeRating17Plus      UniverseAgeRating = "AGE_RATING_17_PLUS"
)

type UniverseSocialLink struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
}

type Universe struct {
	Path                    string             `json:"path"`
	CreateTime              string             `json:"createTime"`
	UpdateTime              string             `json:"updateTime"`
	DisplayName             string             `json:"displayName"`
	Description             string             `json:"description"`
	User                    *string            `json:"user"`
	Group                   *string            `json:"group"`
	Visibility              UniverseVisibility `json:"visibility"`
	FacebookSocialLink      UniverseSocialLink `json:"facebookSocialLink"`
	TwitterSocialLink       UniverseSocialLink `json:"twitterSocialLink"`
	YoutubeSocialLink       UniverseSocialLink `json:"youtubeSocialLink"`
	TwitchSocialLink        UniverseSocialLink `json:"twitchSocialLink"`
	DiscordSocialLink       UniverseSocialLink `json:"discordSocialLink"`
	RobloxGroupSocialLink   UniverseSocialLink `json:"robloxgroupSocialLink"`
	GuildedSocialLink       UniverseSocialLink `json:"guildedSocialLink"`
	VoiceChatEnabled        bool               `json:"voiceChatEnabled"`
	AgeRating               UniverseAgeRating  `json:"ageRating"`
	PrivateServerPriceRobux int                `json:"privateServerPriceRobux"`
	DesktopEnabled          bool               `json:"desktopEnabled"`
	MobileEnabled           bool               `json:"mobileEnabled"`
	TabletEnabled           bool               `json:"tabletEnabled"`
	ConsoleEnabled          bool               `json:"consoleEnabled"`
	VREnabled               bool               `json:"vrEnabled"`
}

// GetUniverse will fetch information on a specified universe.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Universe#Cloud_GetUniverse
//
// [GET] /cloud/v2/universes/{universe_id}
func (s *UniversesService) GetUniverse(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*Universe, *Response, error) {
	universe := new(Universe)
	resp, err := s.client.Call(ctx, GetUniverse.Compile(nil, universeId), nil, universe, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return universe, resp, nil
}

type UniverseUpdate struct {
	DisplayName             *string             `json:"displayName,omitempty"`
	Description             *string             `json:"description,omitempty"`
	Visibility              *UniverseVisibility `json:"visibility,omitempty"`
	FacebookSocialLink      *UniverseSocialLink `json:"facebookSocialLink,omitempty"`
	TwitterSocialLink       *UniverseSocialLink `json:"twitterSocialLink,omitempty"`
	YoutubeSocialLink       *UniverseSocialLink `json:"youtubeSocialLink,omitempty"`
	TwitchSocialLink        *UniverseSocialLink `json:"twitchSocialLink,omitempty"`
	DiscordSocialLink       *UniverseSocialLink `json:"discordSocialLink,omitempty"`
	RobloxGroupSocialLink   *UniverseSocialLink `json:"robloxgroupSocialLink,omitempty"`
	VoiceChatEnabled        *bool               `json:"voiceChatEnabled,omitempty"`
	PrivateServerPriceRobux *int                `json:"privateServerPriceRobux,omitempty"`
	DesktopEnabled          *bool               `json:"desktopEnabled,omitempty"`
	MobileEnabled           *bool               `json:"mobileEnabled,omitempty"`
	TabletEnabled           *bool               `json:"tabletEnabled,omitempty"`
	ConsoleEnabled          *bool               `json:"consoleEnabled,omitempty"`
	VREnabled               *bool               `json:"vrEnabled,omitempty"`
}

// UpdateUniverse will update information for a specified universe.
//
// Required scopes: universe:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Universe#Cloud_UpdateUniverse
//
// [PATCH] /cloud/v2/universes/{universe_id}
func (s *UniversesService) UpdateUniverse(ctx context.Context, universeId string, data UniverseUpdate, reqOpts ...goblox.RequestOpt) (*Universe, *Response, error) {
	universe := new(Universe)
	resp, err := s.client.Call(ctx, UpdateUniverse.Compile(nil, universeId), data, universe, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return universe, resp, nil
}

type UniverseMessage struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

// PublishUniverseMessage will publish a message to all of the universe's live servers.
//
// Required scopes: universe-messaging-service:publish
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Universe#Cloud_PublishUniverseMessage
//
// [POST] /cloud/v2/universes/{universe_id}:publishMessage
func (s *UniversesService) PublishUniverseMessage(ctx context.Context, universeId string, data UniverseMessage, reqOpts ...goblox.RequestOpt) (*Response, error) {
	resp, err := s.client.Call(ctx, PublishUniverseMessage.Compile(nil, universeId), data, nil, reqOpts...)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RestartUniverseServers will restart all active servers for a specific universe ONLY if a new version of the experience has been published.
//
// Required scopes: universe:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Universe#Cloud_RestartUniverseServers
//
// [POST] /cloud/v2/universes/{universe_id}:restartServers
func (s *UniversesService) RestartUniverseServers(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	resp, err := s.client.Call(ctx, RestartUniverseServers.Compile(nil, universeId), nil, nil, reqOpts...)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

var (
	GenerateSpeechAsset = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}:generateSpeechAsset",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:write", "asset:read", "asset:write")

	TranslateText = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}:translateText",
		auth.KindAPIKey,
	).WithScopes("universe:write")

	GetUniversesPermissions = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-develop/v1/universes/multiget/permissions",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetUniversesTeamCreateSettings = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-develop/v1/universes/multiget/teamcreate",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-team-collaboration:manage")

	ActivateUniverse = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-develop/v1/universes/{universeId}/activate",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	DeactivateUniverse = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-develop/v1/universes/{universeId}/deactivate",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetUniversePermissions = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-develop/v1/universes/{universeId}/permissions",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe:manage")

	GetUniverseTeamCreateSettings = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-develop/v1/universes/{universeId}/teamcreate",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-team-collaboration:manage")

	UpdateUniverseTeamCreateSettings = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-develop/v1/universes/{universeId}/teamcreate",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-team-collaboration:manage")

	RemoveUniverseTeamCreateMember = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-develop/v1/universes/{universeId}/teamcreate/memberships",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-team-collaboration:manage")
)

type GenerateSpeechAssetRequestGeneratedSpeechStyle struct {
	VoiceID *string  `json:"voiceId,omitempty"`
	Pitch   *float64 `json:"pitch,omitempty"`
	Speed   *float64 `json:"speed,omitempty"`
}

type GenerateSpeechAssetRequest struct {
	Text        *string                                         `json:"text,omitempty"`
	SpeechStyle *GenerateSpeechAssetRequestGeneratedSpeechStyle `json:"speechStyle,omitempty"`
}

// GenerateSpeechAsset will generate an English speech audio asset from the specified text.
//
// Required scopes:
//
// - universe:write
//
// - asset:read
//
// - asset:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/assets#Cloud_GenerateSpeechAsset
//
// [POST] /cloud/v2/universes/{universe_id}:generateSpeechAsset
func (s *UniversesService) GenerateSpeechAsset(ctx context.Context, universeId string, data GenerateSpeechAssetRequest, reqOpts ...goblox.RequestOpt) (*Operation, *Response, error) {
	operation := new(Operation)
	resp, err := s.client.Call(ctx, GenerateSpeechAsset.Compile(nil, universeId), data, operation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return operation, resp, nil
}

type TranslateTextRequest struct {
	Text                *string  `json:"text,omitempty"`
	SourceLanguageCode  *string  `json:"sourceLanguageCode,omitempty"`
	TargetLanguageCodes []string `json:"targetLanguageCodes,omitempty"`
}

type TranslateTextResponse struct {
	SourceLanguageCode string            `json:"sourceLanguageCode"`
	Translations       map[string]string `json:"translations"`
}

// TranslateText will translate the provided text from one language to another.
//
// Required scopes: universe:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/generative-ai#Cloud_TranslateText
//
// [POST] /cloud/v2/universes/{universe_id}:translateText
func (s *UniversesService) TranslateText(ctx context.Context, universeId string, data TranslateTextRequest, reqOpts ...goblox.RequestOpt) (*TranslateTextResponse, *Response, error) {
	translateTextResponse := new(TranslateTextResponse)
	resp, err := s.client.Call(ctx, TranslateText.Compile(nil, universeId), data, translateTextResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return translateTextResponse, resp, nil
}

type GetUniversesPermissionsOptions struct {
	IDs []int64 `url:"ids"` // Required.
}

type UniverseIDPermissionsModel struct {
	UniverseID   int64 `json:"universeId"`
	CanManage    bool  `json:"canManage"`
	CanCloudEdit bool  `json:"canCloudEdit"`
}

type UniverseIDPermissionsList struct {
	Data []UniverseIDPermissionsModel `json:"data"`
}

// GetUniversesPermissions will return an array of granted and declined permissions related to the universes with
// the ids in ids for the authenticated user.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/universes#get_legacy_develop_v1_universes_multiget_permissions
//
// [GET] /legacy-develop/v1/universes/multiget/permissions
func (s *UniversesService) GetUniversesPermissions(ctx context.Context, opts *GetUniversesPermissionsOptions, reqOpts ...goblox.RequestOpt) (*UniverseIDPermissionsList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	universeIDPermissionsList := new(UniverseIDPermissionsList)
	resp, err := s.client.Call(ctx, GetUniversesPermissions.Compile(q), nil, universeIDPermissionsList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return universeIDPermissionsList, resp, nil
}

type GetUniversesTeamCreateSettingsOptions struct {
	IDs []int64 `url:"ids"` // Required.
}

type UniverseTeamCreateSettingsModel struct {
	ID        int64 `json:"id"`
	IsEnabled bool  `json:"isEnabled"`
}

type UniverseTeamCreateSettingsList struct {
	Data []UniverseTeamCreateSettingsModel `json:"data"`
}

// GetUniversesTeamCreateSettings will get TeamCreate settings for multiple universes specified by Ids.
//
// Required scopes: legacy-team-collaboration:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/team-create#get_legacy_develop_v1_universes_multiget_teamcreate
//
// [GET] /legacy-develop/v1/universes/multiget/teamcreate
func (s *UniversesService) GetUniversesTeamCreateSettings(ctx context.Context, opts *GetUniversesTeamCreateSettingsOptions, reqOpts ...goblox.RequestOpt) (*UniverseTeamCreateSettingsList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	universeTeamCreateSettingsList := new(UniverseTeamCreateSettingsList)
	resp, err := s.client.Call(ctx, GetUniversesTeamCreateSettings.Compile(q), nil, universeTeamCreateSettingsList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return universeTeamCreateSettingsList, resp, nil
}

// ActivateUniverse will activate a universe.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/universes#post_legacy_develop_v1_universes__universeId__activate
//
// [POST] /legacy-develop/v1/universes/{universeId}/activate
func (s *UniversesService) ActivateUniverse(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, ActivateUniverse.Compile(nil, universeId), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

// DeactivateUniverse will deactivate a universe.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/universes#post_legacy_develop_v1_universes__universeId__deactivate
//
// [POST] /legacy-develop/v1/universes/{universeId}/deactivate
func (s *UniversesService) DeactivateUniverse(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, DeactivateUniverse.Compile(nil, universeId), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type UniversePermissionsModel struct {
	CanManage    bool `json:"canManage"`
	CanCloudEdit bool `json:"canCloudEdit"`
}

// GetUniversePermissions will return list of granted and declined permissions related to the universe with the
// id universeId for authenticated user.
//
// Required scopes: legacy-universe:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/universes#get_legacy_develop_v1_universes__universeId__permissions
//
// [GET] /legacy-develop/v1/universes/{universeId}/permissions
func (s *UniversesService) GetUniversePermissions(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*UniversePermissionsModel, *Response, error) {
	universePermissionsModel := new(UniversePermissionsModel)
	resp, err := s.client.Call(ctx, GetUniversePermissions.Compile(nil, universeId), nil, universePermissionsModel, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return universePermissionsModel, resp, nil
}

type TeamCreateSettingsResponse struct {
	IsEnabled bool `json:"isEnabled"`
}

// GetUniverseTeamCreateSettings will get the Team Create settings for a universe.
//
// Required scopes: legacy-team-collaboration:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/team-create#get_legacy_develop_v1_universes__universeId__teamcreate
//
// [GET] /legacy-develop/v1/universes/{universeId}/teamcreate
func (s *UniversesService) GetUniverseTeamCreateSettings(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*TeamCreateSettingsResponse, *Response, error) {
	teamCreateSettingsResponse := new(TeamCreateSettingsResponse)
	resp, err := s.client.Call(ctx, GetUniverseTeamCreateSettings.Compile(nil, universeId), nil, teamCreateSettingsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return teamCreateSettingsResponse, resp, nil
}

type UpdateTeamCreateSettingsRequest struct {
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

// UpdateUniverseTeamCreateSettings will edit team create settings for a universe.
//
// Required scopes: legacy-team-collaboration:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/team-create#patch_legacy_develop_v1_universes__universeId__teamcreate
//
// [PATCH] /legacy-develop/v1/universes/{universeId}/teamcreate
func (s *UniversesService) UpdateUniverseTeamCreateSettings(ctx context.Context, universeId string, data UpdateTeamCreateSettingsRequest, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, UpdateUniverseTeamCreateSettings.Compile(nil, universeId), data, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type TeamCreateMembershipRequest struct {
	UserID *int64 `json:"userId,omitempty"`
}

// RemoveUniverseTeamCreateMember will remove a user from a TeamCreate permissions list.
//
// Required scopes: legacy-team-collaboration:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/team-create#delete_legacy_develop_v1_universes__universeId__teamcreate_memberships
//
// [DELETE] /legacy-develop/v1/universes/{universeId}/teamcreate/memberships
func (s *UniversesService) RemoveUniverseTeamCreateMember(ctx context.Context, universeId string, data TeamCreateMembershipRequest, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, RemoveUniverseTeamCreateMember.Compile(nil, universeId), data, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}
