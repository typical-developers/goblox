package opencloud

import (
	"context"
	"fmt"
	"net/http"
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
func (s *UniversesService) GetUniverse(ctx context.Context, universeId string) (*Universe, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s", universeId)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	universe := new(Universe)
	resp, err := s.client.Do(ctx, req, universe)
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
func (s *UniversesService) UpdateUniverse(ctx context.Context, universeId string, data UniverseUpdate) (*Universe, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s", universeId)

	req, err := s.client.NewRequest(http.MethodPatch, u, data)
	if err != nil {
		return nil, nil, err
	}

	universe := new(Universe)
	resp, err := s.client.Do(ctx, req, universe)
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
func (s *UniversesService) PublishUniverseMessage(ctx context.Context, universeId string, data UniverseMessage) (*Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s:publishMessage", universeId)

	req, err := s.client.NewRequest(http.MethodPost, u, data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
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
func (s *UniversesService) RestartUniverseServers(ctx context.Context, universeId string) (*Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s:restartServers", universeId)

	req, err := s.client.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
