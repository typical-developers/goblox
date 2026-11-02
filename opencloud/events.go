package opencloud

import (
	"context"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	DeleteGameEvent = goblox.NewEndpoint(
		http.MethodDelete,
		"/virtual-events/v3/game-events/{eventId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.event:write")

	GetGameEvent = goblox.NewEndpoint(
		http.MethodGet,
		"/virtual-events/v3/game-events/{eventId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.event:read")

	UpdateGameEvent = goblox.NewEndpoint(
		http.MethodPatch,
		"/virtual-events/v3/game-events/{eventId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.event:write")

	ListGameEvents = goblox.NewEndpoint(
		http.MethodGet,
		"/virtual-events/v3/universes/{universeId}/game-events",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.event:read")

	CreateGameEvent = goblox.NewEndpoint(
		http.MethodPost,
		"/virtual-events/v3/universes/{universeId}/game-events",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.event:write")
)

// EventsService handles communication with the virtual events endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type EventsService service

// DeleteGameEvent will permanently delete an existing game event.
//
// Required scopes: universe.event:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/virtual-events#GameEvent_DeleteGameEvent
//
// [DELETE] /virtual-events/v3/game-events/{eventId}
func (s *EventsService) DeleteGameEvent(ctx context.Context, eventId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, DeleteGameEvent.Compile(nil, eventId), nil, nil, reqOpts...)
}

type GetGameEventOptions struct {
	Fields *string `url:"fields,omitempty"`
}

type HostType string

const (
	HostTypeUser  HostType = "user"
	HostTypeGroup HostType = "group"
)

type HostResponse struct {
	HostType         HostType `json:"hostType"`
	HostID           int64    `json:"hostId"`
	HostName         string   `json:"hostName"`
	HasVerifiedBadge bool     `json:"hasVerifiedBadge"`
}

type EventVisibility string

const (
	EventVisibilityPrivate   EventVisibility = "private"
	EventVisibilityPublic    EventVisibility = "public"
	EventVisibilityModerated EventVisibility = "moderated"
)

type FeaturingStatus string

const (
	FeaturingStatusEnabled  FeaturingStatus = "enabled"
	FeaturingStatusDisabled FeaturingStatus = "disabled"
)

type EventCategory string

const (
	EventCategoryContentUpdate  EventCategory = "contentUpdate"
	EventCategoryLocationUpdate EventCategory = "locationUpdate"
	EventCategorySystemUpdate   EventCategory = "systemUpdate"
	EventCategoryActivity       EventCategory = "activity"
	EventCategoryNewContent     EventCategory = "newContent"
	EventCategoryItemDrop       EventCategory = "itemDrop"
	EventCategoryNewSeason      EventCategory = "newSeason"
	EventCategoryNewLocation    EventCategory = "newLocation"
	EventCategoryNewMap         EventCategory = "newMap"
	EventCategoryMoreLevels     EventCategory = "moreLevels"
	EventCategoryNewFeature     EventCategory = "newFeature"
	EventCategoryEarlyAccess    EventCategory = "earlyAccess"
	EventCategoryExpansion      EventCategory = "expansion"
	EventCategoryChallenge      EventCategory = "challenge"
	EventCategoryQuest          EventCategory = "quest"
	EventCategoryFestival       EventCategory = "festival"
)

type CategoryResponse struct {
	Category EventCategory `json:"category"`
	Rank     int           `json:"rank"`
}

type ThumbnailResponse struct {
	MediaID int64 `json:"mediaId"`
	Rank    int   `json:"rank"`
}

type RsvpStatus string

const (
	RsvpStatusNone       RsvpStatus = "none"
	RsvpStatusGoing      RsvpStatus = "going"
	RsvpStatusMaybeGoing RsvpStatus = "maybeGoing"
	RsvpStatusNotGoing   RsvpStatus = "notGoing"
)

type GameEventResponse struct {
	ID                   string              `json:"id"`
	Title                string              `json:"title"`
	DisplayTitle         string              `json:"displayTitle"`
	Subtitle             string              `json:"subtitle"`
	DisplaySubtitle      string              `json:"displaySubtitle"`
	Description          string              `json:"description"`
	DisplayDescription   string              `json:"displayDescription"`
	StartTime            string              `json:"startTime"`
	EndTime              string              `json:"endTime"`
	UniverseID           int64               `json:"universeId"`
	PlaceID              int64               `json:"placeId"`
	Host                 HostResponse        `json:"host"`
	Visibility           EventVisibility     `json:"visibility"`
	FeaturingStatus      FeaturingStatus     `json:"featuringStatus"`
	Tagline              string              `json:"tagline"`
	Categories           []CategoryResponse  `json:"categories"`
	Thumbnails           []ThumbnailResponse `json:"thumbnails"`
	AllThumbnailsCreated bool                `json:"allThumbnailsCreated"`
	UserRsvpStatus       RsvpStatus          `json:"userRsvpStatus"`
	CreateTime           string              `json:"createTime"`
	UpdateTime           string              `json:"updateTime"`
}

// GetGameEvent will get a single game event by ID.
//
// Required scopes: universe.event:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/virtual-events#GameEvent_GetGameEvent
//
// [GET] /virtual-events/v3/game-events/{eventId}
func (s *EventsService) GetGameEvent(ctx context.Context, eventId string, opts *GetGameEventOptions, reqOpts ...goblox.RequestOpt) (*GameEventResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	gameEventResponse := new(GameEventResponse)
	resp, err := s.client.Call(ctx, GetGameEvent.Compile(q, eventId), nil, gameEventResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return gameEventResponse, resp, nil
}

type EventRankedCategory struct {
	Category *EventCategory `json:"category,omitempty"`
	Rank     *int           `json:"rank,omitempty"`
}

type EventMedia struct {
	MediaID *int64 `json:"mediaId,omitempty"`
	Rank    *int   `json:"rank,omitempty"`
}

type UpdateGameEventRequest struct {
	Title           *string               `json:"title,omitempty"`
	Subtitle        *string               `json:"subtitle,omitempty"`
	Description     *string               `json:"description,omitempty"`
	StartTime       *string               `json:"startTime,omitempty"`
	EndTime         *string               `json:"endTime,omitempty"`
	Visibility      *EventVisibility      `json:"visibility,omitempty"`
	PlaceID         *int64                `json:"placeId,omitempty"`
	Categories      []EventRankedCategory `json:"categories,omitempty"`
	Thumbnails      []EventMedia          `json:"thumbnails,omitempty"`
	FeaturingStatus *FeaturingStatus      `json:"featuringStatus,omitempty"`
	Tagline         *string               `json:"tagline,omitempty"`
}

// UpdateGameEvent will partially update a game event.
//
// Required scopes: universe.event:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/virtual-events#GameEvent_UpdateGameEvent
//
// [PATCH] /virtual-events/v3/game-events/{eventId}
func (s *EventsService) UpdateGameEvent(ctx context.Context, eventId string, data UpdateGameEventRequest, reqOpts ...goblox.RequestOpt) (*GameEventResponse, *Response, error) {
	gameEventResponse := new(GameEventResponse)
	resp, err := s.client.Call(ctx, UpdateGameEvent.Compile(nil, eventId), data, gameEventResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return gameEventResponse, resp, nil
}

type ListGameEventsOptions struct {
	PageSize     *int             `url:"pageSize,omitempty"`
	PageToken    *string          `url:"pageToken,omitempty"`
	Reverse      *bool            `url:"reverse,omitempty"`
	StartsBefore *string          `url:"startsBefore,omitempty"`
	StartsAfter  *string          `url:"startsAfter,omitempty"`
	EndsBefore   *string          `url:"endsBefore,omitempty"`
	EndsAfter    *string          `url:"endsAfter,omitempty"`
	Visibility   *EventVisibility `url:"visibility,omitempty"`
	Fields       *string          `url:"fields,omitempty"`
}

type PaginatedGameEventsResponse struct {
	GameEvents        []GameEventResponse `json:"gameEvents"`
	NextPageToken     string              `json:"nextPageToken"`
	PreviousPageToken string              `json:"previousPageToken"`
}

// ListGameEvents will list game events under a universe.
//
// Required scopes: universe.event:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/virtual-events#GameEvents_GetGameEvents
//
// [GET] /virtual-events/v3/universes/{universeId}/game-events
func (s *EventsService) ListGameEvents(ctx context.Context, universeId string, opts *ListGameEventsOptions, reqOpts ...goblox.RequestOpt) (*PaginatedGameEventsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	paginatedGameEventsResponse := new(PaginatedGameEventsResponse)
	resp, err := s.client.Call(ctx, ListGameEvents.Compile(q, universeId), nil, paginatedGameEventsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return paginatedGameEventsResponse, resp, nil
}

type CreateGameEventRequest struct {
	Title           *string               `json:"title,omitempty"`
	Subtitle        *string               `json:"subtitle,omitempty"`
	Description     *string               `json:"description,omitempty"`
	StartTime       *string               `json:"startTime,omitempty"`
	EndTime         *string               `json:"endTime,omitempty"`
	Visibility      *EventVisibility      `json:"visibility,omitempty"`
	GroupID         *int64                `json:"groupId,omitempty"`
	PlaceID         *int64                `json:"placeId,omitempty"`
	Categories      []EventRankedCategory `json:"categories,omitempty"`
	Thumbnails      []EventMedia          `json:"thumbnails,omitempty"`
	FeaturingStatus *FeaturingStatus      `json:"featuringStatus,omitempty"`
	Tagline         *string               `json:"tagline,omitempty"`
}

// CreateGameEvent will create a new game event.
//
// Required scopes: universe.event:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/virtual-events#GameEvents_CreateGameEvent
//
// [POST] /virtual-events/v3/universes/{universeId}/game-events
func (s *EventsService) CreateGameEvent(ctx context.Context, universeId string, data CreateGameEventRequest, reqOpts ...goblox.RequestOpt) (*GameEventResponse, *Response, error) {
	gameEventResponse := new(GameEventResponse)
	resp, err := s.client.Call(ctx, CreateGameEvent.Compile(nil, universeId), data, gameEventResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return gameEventResponse, resp, nil
}
