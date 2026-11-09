package opencloud

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/luckfire/goblox"
	"github.com/luckfire/goblox/auth"
)

var (
	GetInstance = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/places/{placeId}/instances/{instanceId}",
		auth.KindAPIKey,
	).WithScopes("universe.place.instance:read")

	ListInstanceChildren = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/places/{placeId}/instances/{instanceId}:listChildren",
		auth.KindAPIKey,
	).WithScopes("universe.place.instance:read")

	GetPlace = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/places/{placeId}",
		auth.KindAPIKey, auth.KindOAuth2,
	)

	UpdatePlace = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/universes/{universeId}/places/{placeId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.place:write")

	UpdateInstance = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/universes/{universeId}/places/{placeId}/instances/{instanceId}",
		auth.KindAPIKey,
	).WithScopes("universe.place.instance:write")
)

// PlacesService handles communication with place and instance endpoints.
type PlacesService service

type InstanceRunContext string

const (
	InstanceRunContextLegacy InstanceRunContext = "Legacy"
	InstanceRunContextServer InstanceRunContext = "Server"
	InstanceRunContextClient InstanceRunContext = "Client"
	InstanceRunContextPlugin InstanceRunContext = "Plugin"
)

type InstanceFolder any

type InstanceLocalScript struct {
	Enabled    bool               `json:"Enabled"`
	RunContext InstanceRunContext `json:"RunContext"`
	Source     string             `json:"Source"`
}

type InstanceModuleScript struct {
	Source string `json:"Source"`
}

type InstanceScript struct {
	Enabled    bool               `json:"Enabled"`
	RunContext InstanceRunContext `json:"RunContext"`
	Source     string             `json:"Source"`
}

type InstanceDetails struct {
	Folder       *InstanceFolder       `json:"Folder"`
	LocalScript  *InstanceLocalScript  `json:"LocalScript"`
	ModuleScript *InstanceModuleScript `json:"ModuleScript"`
	Script       *InstanceScript       `json:"Script"`
}

type EngineInstance struct {
	Id     string `json:"Id"`
	Parent string `json:"Parent"`
	Name   string `json:"Name"`
}

type Instance struct {
	Path           string         `json:"path"`
	HasChildren    bool           `json:"hasChildren"`
	EngineInstance EngineInstance `json:"engineInstance"`
}

// GetInstance will fetch an instance's property data.
//
// Required scopes: universe.place.instance:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Instance#Cloud_GetInstance
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}/instances/{instance_id}
func (s *PlacesService) GetInstance(ctx context.Context, universeId, placeId, instanceId string, reqOpts ...goblox.RequestOpt) (*Instance, *Response, error) {
	instance := new(Instance)
	resp, err := s.client.Call(ctx, GetInstance.Compile(nil, universeId, placeId, instanceId), nil, instance, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return instance, resp, nil
}

type EngineInstanceUpdate struct {
	Id      *string          `json:"Id,omitempty"`
	Parent  *string          `json:"Parent,omitempty"`
	Name    *string          `json:"Name,omitempty"`
	Details *InstanceDetails `json:"Details,omitempty"`
}

type InstanceUpdate struct {
	EngineInstance *EngineInstanceUpdate `json:"engineInstance,omitempty"`
}

type InstanceUpdateOptions struct {
	UpdateMask *string `url:"updateMask,omitempty"`
}

// UpdateInstance will update an instance's property data, such as a script's Source.
// The update runs as a long-running operation; poll it with [PlacesService.GetUpdateInstanceOperation].
//
// Required scopes: universe.place.instance:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/places#Cloud_UpdateInstance
//
// [PATCH] /cloud/v2/universes/{universe_id}/places/{place_id}/instances/{instance_id}
func (s *PlacesService) UpdateInstance(ctx context.Context, universeId, placeId, instanceId string, data InstanceUpdate, opts *InstanceUpdateOptions, reqOpts ...goblox.RequestOpt) (*Operation, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	operation := new(Operation)
	resp, err := s.client.Call(ctx, UpdateInstance.Compile(q, universeId, placeId, instanceId), data, operation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return operation, resp, nil
}

type InstanceChildren struct {
	Instances     []Instance `json:"instances"`
	NextPageToken string     `json:"nextPageToken"`
}

// ListInstanceChildren will list and instance's children.
//
// Required scopes: universe.place.instance:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Instance#Cloud_ListInstanceChildren
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}/instances/{instance_id}:listChildren
func (s *PlacesService) ListInstanceChildren(ctx context.Context, universeId, placeId, instanceId string, opts *Options, reqOpts ...goblox.RequestOpt) (*InstanceChildren, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	instanceChildren := new(InstanceChildren)
	resp, err := s.client.Call(ctx, ListInstanceChildren.Compile(q, universeId, placeId, instanceId), nil, instanceChildren, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return instanceChildren, resp, nil
}

type Place struct {
	Path        string `json:"path"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	ServerSize  int    `json:"serverSize"`
}

// GetPlace will get information on a specified place.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Place#Cloud_GetPlace
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}
func (s *PlacesService) GetPlace(ctx context.Context, universeId, placeId string, reqOpts ...goblox.RequestOpt) (*Place, *Response, error) {
	place := new(Place)
	resp, err := s.client.Call(ctx, GetPlace.Compile(nil, universeId, placeId), nil, place, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return place, resp, nil
}

type PlaceUpdate struct {
	DisplayName *string `json:"displayName,omitempty"`
	Description *string `json:"description,omitempty"`
	ServerSize  *int    `json:"serverSize,omitempty"`
}

type PlaceUpdateOptions struct {
	UpdateMask *string `url:"updateMask,omitempty"`
}

// UpdatePlace will update information for a specified place.
//
// Required scopes: universe.place:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Place#Cloud_UpdatePlace
//
// [PATCH] /cloud/v2/universes/{universe_id}/places/{place_id}
func (s *PlacesService) UpdatePlace(ctx context.Context, universeId, placeId string, data PlaceUpdate, opts *PlaceUpdateOptions, reqOpts ...goblox.RequestOpt) (*Place, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	place := new(Place)
	resp, err := s.client.Call(ctx, UpdatePlace.Compile(q, universeId, placeId), data, place, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return place, resp, nil
}

var (
	GetUpdateInstanceOperation = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/places/{placeId}/instances/{instanceId}/operations/{operationId}",
		auth.KindAPIKey,
	)

	ListTeamCreateActiveSessionMembers = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-develop/v1/places/{placeId}/teamcreate/active_session/members",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-team-collaboration:manage")

	CloseTeamTest = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-develop/v2/teamtest/{placeId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-team-collaboration:manage")

	GetPlaceContributors = goblox.NewEndpoint(
		http.MethodGet,
		"/place-version-history-api/v1/{placeId}/contributors",
		auth.KindAPIKey,
	).WithScopes("universe.place:read")

	ListPlaceVersionHistory = goblox.NewEndpoint(
		http.MethodGet,
		"/place-version-history-api/v1/{placeId}/history",
		auth.KindAPIKey,
	).WithScopes("universe.place:read")

	PublishPlaceVersion = goblox.NewEndpoint(
		http.MethodPost,
		"/universes/v1/{universeId}/places/{placeId}/versions",
		auth.KindAPIKey,
	).WithScopes("universe-places:write")

	UpdatePlaceVersionNotes = goblox.NewEndpoint(
		http.MethodPost,
		"/place-version-history-api/v1/{placeId}/version/{version}/notes",
		auth.KindAPIKey,
	).WithScopes("universe.place:write")
)

// GetUpdateInstanceOperation will retrieve the status of the operation to update an instance.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/places#Cloud_GetUpdateInstanceOperation
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}/instances/{instance_id}/operations/{operation_id}
func (s *PlacesService) GetUpdateInstanceOperation(ctx context.Context, universeId, placeId, instanceId, operationId string, reqOpts ...goblox.RequestOpt) (*Operation, *Response, error) {
	operation := new(Operation)
	resp, err := s.client.Call(ctx, GetUpdateInstanceOperation.Compile(nil, universeId, placeId, instanceId, operationId), nil, operation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return operation, resp, nil
}

type ListTeamCreateActiveSessionMembersOptionsLimit int

type ListTeamCreateActiveSessionMembersOptions struct {
	Limit  *ListTeamCreateActiveSessionMembersOptionsLimit `url:"limit,omitempty"`
	Cursor *string                                         `url:"cursor,omitempty"`
}

type SkinnyUser struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type SkinnyUserList struct {
	Data []SkinnyUser `json:"data"`
}

// ListTeamCreateActiveSessionMembers will list of users in the active Team Create session.
//
// Required scopes: legacy-team-collaboration:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/places#get_legacy_develop_v1_places__placeId__teamcreate_active_session_members
//
// [GET] /legacy-develop/v1/places/{placeId}/teamcreate/active_session/members
func (s *PlacesService) ListTeamCreateActiveSessionMembers(ctx context.Context, placeId string, opts *ListTeamCreateActiveSessionMembersOptions, reqOpts ...goblox.RequestOpt) (*SkinnyUserList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	skinnyUserList := new(SkinnyUserList)
	resp, err := s.client.Call(ctx, ListTeamCreateActiveSessionMembers.Compile(q, placeId), nil, skinnyUserList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return skinnyUserList, resp, nil
}

type CloseTeamTestOptions struct {
	GameID string `url:"gameId"` // Required.
}

// CloseTeamTest will close a game instance that is being used for team testing.
//
// Required scopes: legacy-team-collaboration:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/places#delete_legacy_develop_v2_teamtest__placeId_
//
// [DELETE] /legacy-develop/v2/teamtest/{placeId}
func (s *PlacesService) CloseTeamTest(ctx context.Context, placeId string, opts *CloseTeamTestOptions, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, CloseTeamTest.Compile(q, placeId), nil, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type GetPlaceContributorsOptions struct {
	Cursor   *string `url:"cursor,omitempty"`
	PageSize *int    `url:"pageSize,omitempty"`
}

type GetPlaceContributorsResponse struct {
	NextCursor   string  `json:"nextCursor"`
	HasMore      bool    `json:"hasMore"`
	Contributors []int64 `json:"contributors"`
}

// GetPlaceContributors will fetch all previous contributors of a place.
//
// Required scopes: universe.place:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/places#PlaceVersion_GetPlaceContributors
//
// [GET] /place-version-history-api/v1/{placeId}/contributors
func (s *PlacesService) GetPlaceContributors(ctx context.Context, placeId string, opts *GetPlaceContributorsOptions, reqOpts ...goblox.RequestOpt) (*GetPlaceContributorsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getPlaceContributorsResponse := new(GetPlaceContributorsResponse)
	resp, err := s.client.Call(ctx, GetPlaceContributors.Compile(q, placeId), nil, getPlaceContributorsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getPlaceContributorsResponse, resp, nil
}

type ListPlaceVersionHistoryOptions struct {
	Cursor      *string `url:"cursor,omitempty"`
	IsPublished *bool   `url:"isPublished,omitempty"`
	HasNotes    *bool   `url:"hasNotes,omitempty"`
	SaveType    *int    `url:"saveType,omitempty"`
	SearchTerm  *string `url:"searchTerm,omitempty"`
	Contributor *int64  `url:"contributor,omitempty"`
	StartTime   *string `url:"startTime,omitempty"`
	EndTime     *string `url:"endTime,omitempty"`
	PageSize    *int    `url:"pageSize,omitempty"`
}

type PlaceSaveType int

type PlacePublishStatus int

type PlaceVersion struct {
	Version       string             `json:"version"`
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	Contributors  []int64            `json:"contributors"`
	SaveType      PlaceSaveType      `json:"saveType"`
	IsPublished   bool               `json:"isPublished"`
	PublishStatus PlacePublishStatus `json:"publishStatus"`
	HasNotes      bool               `json:"hasNotes"`
	CreatedTime   string             `json:"createdTime"`
	CreatedBy     int64              `json:"createdBy"`
}

type GetPlaceVersionHistoryResponse struct {
	NextCursor    string         `json:"nextCursor"`
	HasMore       bool           `json:"hasMore"`
	PlaceVersions []PlaceVersion `json:"placeVersions"`
}

// ListPlaceVersionHistory will get place versions, using various filters.
//
// Required scopes: universe.place:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/places#PlaceVersion_GetPlaceVersionHistory
//
// [GET] /place-version-history-api/v1/{placeId}/history
func (s *PlacesService) ListPlaceVersionHistory(ctx context.Context, placeId string, opts *ListPlaceVersionHistoryOptions, reqOpts ...goblox.RequestOpt) (*GetPlaceVersionHistoryResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getPlaceVersionHistoryResponse := new(GetPlaceVersionHistoryResponse)
	resp, err := s.client.Call(ctx, ListPlaceVersionHistory.Compile(q, placeId), nil, getPlaceVersionHistoryResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getPlaceVersionHistoryResponse, resp, nil
}

type UpdatePlaceVersionNotesRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdatePlaceVersionNotesResponse struct {
	PlaceVersion PlaceVersion `json:"placeVersion"`
}

// UpdatePlaceVersionNotes will modify the notes of a particular version for a place.
//
// Required scopes: universe.place:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/places#PlaceVersion_UpdatePlaceVersionNotes
//
// [POST] /place-version-history-api/v1/{placeId}/version/{version}/notes
func (s *PlacesService) UpdatePlaceVersionNotes(ctx context.Context, placeId, version string, data UpdatePlaceVersionNotesRequest, reqOpts ...goblox.RequestOpt) (*UpdatePlaceVersionNotesResponse, *Response, error) {
	updatePlaceVersionNotesResponse := new(UpdatePlaceVersionNotesResponse)
	resp, err := s.client.Call(ctx, UpdatePlaceVersionNotes.Compile(nil, placeId, version), data, updatePlaceVersionNotesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updatePlaceVersionNotesResponse, resp, nil
}

type PlaceVersionType string

const (
	PlaceVersionTypeSaved     PlaceVersionType = "Saved"
	PlaceVersionTypePublished PlaceVersionType = "Published"
)

type PublishPlaceVersionOptions struct {
	VersionType *PlaceVersionType `url:"versionType,omitempty"`
}

type PlaceVersionPublished struct {
	VersionNumber int64 `json:"versionNumber"`
}

// PublishPlaceVersion will publish a new version of a place from a .rbxl or .rbxlx file.
// Files ending in .rbxlx are sent as XML, everything else as binary.
//
// Required scopes: universe-places:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/places#Places_CreatePlaceVersionApiKey
//
// [POST] /universes/v1/{universeId}/places/{placeId}/versions
func (s *PlacesService) PublishPlaceVersion(ctx context.Context, universeId, placeId string, place *os.File, opts *PublishPlaceVersionOptions, reqOpts ...goblox.RequestOpt) (*PlaceVersionPublished, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	var body bytes.Buffer
	if _, err := io.Copy(&body, place); err != nil {
		return nil, nil, err
	}

	contentType := "application/octet-stream"
	if strings.EqualFold(filepath.Ext(place.Name()), ".rbxlx") {
		contentType = "application/xml"
	}

	e := PublishPlaceVersion.Compile(q, universeId, placeId)
	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, contentType)
	if err != nil {
		return nil, nil, err
	}

	published := new(PlaceVersionPublished)
	resp, err := s.client.CallRequest(ctx, e, req, published, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return published, resp, nil
}
