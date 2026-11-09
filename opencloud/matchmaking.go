package opencloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/luckfire/goblox"
)

var (
	GetClientStatus = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/client-status",
	)

	SetClientStatus = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/client-status",
	)

	ForecastGameUpdate = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/game-instances/forecast-update",
	)

	GetGameUpdateStatus = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/game-instances/get-update-status",
	)

	LaunchGameUpdate = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/game-instances/launch-update",
	)

	ShutdownGameInstances = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/game-instances/shutdown",
	)

	ShutdownAllGameInstances = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/game-instances/shutdown-all",
	)

	CreatePlayerAttributeDefinition = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/matchmaking/player-attribute",
	)

	DeletePlayerAttributeDefinition = goblox.NewEndpoint(
		http.MethodDelete,
		"/matchmaking-api/v1/matchmaking/player-attribute/{attributeId}",
	)

	UpdatePlayerAttributeDefinition = goblox.NewEndpoint(
		http.MethodPatch,
		"/matchmaking-api/v1/matchmaking/player-attribute/{attributeId}",
	)

	ListPlayerAttributeDefinitions = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/matchmaking/player-attributes/{universeId}",
	)

	CreateScoringConfiguration = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/matchmaking/scoring-configuration",
	)

	GetScoringDefaultWeights = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/matchmaking/scoring-configuration/default-weights",
	)

	GenerateMockServerSignalValues = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/matchmaking/scoring-configuration/generate-mock-servers",
	)

	SetPlaceScoringConfiguration = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/matchmaking/scoring-configuration/place",
	)

	RemovePlaceScoringConfiguration = goblox.NewEndpoint(
		http.MethodDelete,
		"/matchmaking-api/v1/matchmaking/scoring-configuration/place/{placeId}",
	)

	DeleteScoringConfiguration = goblox.NewEndpoint(
		http.MethodDelete,
		"/matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}",
	)

	GetScoringConfiguration = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}",
	)

	UpdateScoringConfiguration = goblox.NewEndpoint(
		http.MethodPatch,
		"/matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}",
	)

	CreateScoringConfigurationSignal = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}/signals",
	)

	DeleteScoringConfigurationSignal = goblox.NewEndpoint(
		http.MethodDelete,
		"/matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}/signals/{signalName}",
	)

	UpdateScoringConfigurationSignal = goblox.NewEndpoint(
		http.MethodPatch,
		"/matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}/signals/{signalName}",
	)

	ListScoringConfigurations = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/matchmaking/scoring-configurations/{universeId}",
	)

	ListPlaceScoringConfigurations = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/matchmaking/scoring-configurations/{universeId}/places",
	)

	CreateServerAttributeDefinition = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/matchmaking/server-attribute",
	)

	DeleteServerAttributeDefinition = goblox.NewEndpoint(
		http.MethodDelete,
		"/matchmaking-api/v1/matchmaking/server-attribute/{attributeId}",
	)

	UpdateServerAttributeDefinition = goblox.NewEndpoint(
		http.MethodPatch,
		"/matchmaking-api/v1/matchmaking/server-attribute/{attributeId}",
	)

	ListServerAttributeDefinitions = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/matchmaking/server-attributes/{universeId}",
	)

	GetMatchmakingFeatureFlags = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/matchmaking/universe/{universeId}/feature-flags",
	)

	DeleteStartPlaceOverride = goblox.NewEndpoint(
		http.MethodDelete,
		"/matchmaking-api/v1/matchmaking/universe/{universeId}/start-place",
	)

	GetStartPlaceOverride = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/matchmaking/universe/{universeId}/start-place",
	)

	SetStartPlaceOverride = goblox.NewEndpoint(
		http.MethodPost,
		"/matchmaking-api/v1/matchmaking/universe/{universeId}/start-place",
	)

	TestStartPlaceOverride = goblox.NewEndpoint(
		http.MethodGet,
		"/matchmaking-api/v1/matchmaking/universe/{universeId}/start-place/test",
	)
)

// MatchmakingService handles communication with the matchmaking endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type MatchmakingService service

type ClientStatusGetRequest struct {
	BrowserTrackerID *int64 `json:"browserTrackerId,omitempty"`
}

// GetClientStatus will get the matchmaking client status.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_client_status
//
// [GET] /matchmaking-api/v1/client-status
func (s *MatchmakingService) GetClientStatus(ctx context.Context, data ClientStatusGetRequest, reqOpts ...goblox.RequestOpt) (json.RawMessage, *Response, error) {
	var raw json.RawMessage
	resp, err := s.client.Call(ctx, GetClientStatus.Compile(nil), data, &raw, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return raw, resp, nil
}

type ClientStatusSetRequest struct {
	Status           *string `json:"status,omitempty"`
	BrowserTrackerID *int64  `json:"browserTrackerId,omitempty"`
}

// SetClientStatus will set the matchmaking client status.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_client_status
//
// [POST] /matchmaking-api/v1/client-status
func (s *MatchmakingService) SetClientStatus(ctx context.Context, data ClientStatusSetRequest, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, SetClientStatus.Compile(nil), data, nil, reqOpts...)
}

type ForecastUpdateRequest struct {
	UniverseID *int64 `json:"universeId,omitempty"`
}

type PlaceSummaryForGameUpdate struct {
	PlaceID                 int64          `json:"placeId"`
	PlayersToBeKicked       int            `json:"playersToBeKicked"`
	TotalPlayers            int            `json:"totalPlayers"`
	InstancesToBeClosed     int            `json:"instancesToBeClosed"`
	TotalInstances          int            `json:"totalInstances"`
	PlaceVersion            int            `json:"placeVersion"`
	LastUpdated             string         `json:"lastUpdated"`
	IsInUniverse            bool           `json:"isInUniverse"`
	VersionToPlayerCounts   map[string]int `json:"versionToPlayerCounts"`
	VersionToInstanceCounts map[string]int `json:"versionToInstanceCounts"`
}

type ForecastUpdateResponse struct {
	PlaceSummaries []PlaceSummaryForGameUpdate `json:"placeSummaries"`
}

// ForecastGameUpdate will forecast the outcome of launching an update.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_game_instances_forecast_update
//
// [POST] /matchmaking-api/v1/game-instances/forecast-update
func (s *MatchmakingService) ForecastGameUpdate(ctx context.Context, data ForecastUpdateRequest, reqOpts ...goblox.RequestOpt) (*ForecastUpdateResponse, *Response, error) {
	forecastUpdateResponse := new(ForecastUpdateResponse)
	resp, err := s.client.Call(ctx, ForecastGameUpdate.Compile(nil), data, forecastUpdateResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return forecastUpdateResponse, resp, nil
}

type GetGameUpdateStatusOptions struct {
	UniverseID *int64 `url:"universeId,omitempty"`
}

type PlaceUpdateStatus struct {
	PlaceID                  int64  `json:"placeId"`
	Phase                    string `json:"phase"`
	StartTime                string `json:"startTime"`
	EndTime                  string `json:"endTime"`
	StartPlayersToBeKicked   int    `json:"startPlayersToBeKicked"`
	StartInstancesToBeClosed int    `json:"startInstancesToBeClosed"`
	NumPlayersToBeKicked     int    `json:"numPlayersToBeKicked"`
	InstancesToBeClosed      int    `json:"instancesToBeClosed"`
	PlaceVersion             int    `json:"placeVersion"`
	PlaceVersionsToBeClosed  []int  `json:"placeVersionsToBeClosed"`
}

type GameUpdateStatus struct {
	ID                   string                       `json:"id"`
	UniverseID           int64                        `json:"universeId"`
	StartTime            string                       `json:"startTime"`
	CloseOldVersionsOnly bool                         `json:"closeOldVersionsOnly"`
	BleedOffEndTime      string                       `json:"bleedOffEndTime"`
	PlaceUpdateStatuses  map[string]PlaceUpdateStatus `json:"placeUpdateStatuses"`
	BleedOffServers      bool                         `json:"bleedOffServers"`
}

type GetUpdateStatusResponse struct {
	UpdateStatusList []GameUpdateStatus `json:"updateStatusList"`
}

// GetGameUpdateStatus will get the rollout status of an update.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_game_instances_get_update_status
//
// [GET] /matchmaking-api/v1/game-instances/get-update-status
func (s *MatchmakingService) GetGameUpdateStatus(ctx context.Context, opts *GetGameUpdateStatusOptions, reqOpts ...goblox.RequestOpt) (*GetUpdateStatusResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getUpdateStatusResponse := new(GetUpdateStatusResponse)
	resp, err := s.client.Call(ctx, GetGameUpdateStatus.Compile(q), nil, getUpdateStatusResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getUpdateStatusResponse, resp, nil
}

type LaunchUpdateRequest struct {
	UniverseID              *int64           `json:"universeId,omitempty"`
	PlaceIDs                []int64          `json:"placeIds,omitempty"`
	CloseOldVersionsOnly    *bool            `json:"closeOldVersionsOnly,omitempty"`
	BleedOffServers         *bool            `json:"bleedOffServers,omitempty"`
	BleedOffDurationMinutes *int             `json:"bleedOffDurationMinutes,omitempty"`
	PlaceIDToVersions       map[string][]int `json:"placeIdToVersions,omitempty"`
	Attributes              *string          `json:"attributes,omitempty"`
}

type LaunchUpdateResponse struct {
	UpdateID             string `json:"updateId"`
	NumPlayersToBeKicked int    `json:"numPlayersToBeKicked"`
	InstancesToBeClosed  int    `json:"instancesToBeClosed"`
}

// LaunchGameUpdate will launch a game update.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_game_instances_launch_update
//
// [POST] /matchmaking-api/v1/game-instances/launch-update
func (s *MatchmakingService) LaunchGameUpdate(ctx context.Context, data LaunchUpdateRequest, reqOpts ...goblox.RequestOpt) (*LaunchUpdateResponse, *Response, error) {
	launchUpdateResponse := new(LaunchUpdateResponse)
	resp, err := s.client.Call(ctx, LaunchGameUpdate.Compile(nil), data, launchUpdateResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return launchUpdateResponse, resp, nil
}

type ShutdownGameInstancesRequest struct {
	PlaceID         *int64  `json:"placeId,omitempty"`
	PrivateServerID *int64  `json:"privateServerId,omitempty"`
	GameID          *string `json:"gameId,omitempty"`
}

type ShutdownGameInstancesResponse struct {
	PlaceID         int64  `json:"placeId"`
	PrivateServerID int64  `json:"privateServerId"`
	GameID          string `json:"gameId"`
}

// ShutdownGameInstances will shutdown game instances.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_game_instances_shutdown
//
// [POST] /matchmaking-api/v1/game-instances/shutdown
func (s *MatchmakingService) ShutdownGameInstances(ctx context.Context, data ShutdownGameInstancesRequest, reqOpts ...goblox.RequestOpt) (*ShutdownGameInstancesResponse, *Response, error) {
	shutdownGameInstancesResponse := new(ShutdownGameInstancesResponse)
	resp, err := s.client.Call(ctx, ShutdownGameInstances.Compile(nil), data, shutdownGameInstancesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return shutdownGameInstancesResponse, resp, nil
}

type ShutdownAllGameInstancesRequest struct {
	PlaceID          *int64 `form:"PlaceId"`
	ReplaceInstances *bool  `form:"ReplaceInstances"`
}

type ShutdownAllGameInstancesResponse struct {
	PlaceID          int64 `json:"placeId"`
	ReplaceInstances bool  `json:"replaceInstances"`
}

// ShutdownAllGameInstances will shutdown all game instances.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_game_instances_shutdown_all
//
// [POST] /matchmaking-api/v1/game-instances/shutdown-all
func (s *MatchmakingService) ShutdownAllGameInstances(ctx context.Context, data ShutdownAllGameInstancesRequest, reqOpts ...goblox.RequestOpt) (*ShutdownAllGameInstancesResponse, *Response, error) {
	shutdownAllGameInstancesResponse := new(ShutdownAllGameInstancesResponse)
	e := ShutdownAllGameInstances.Compile(nil)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.PlaceID != nil {
		_ = writer.WriteField("PlaceId", fmt.Sprint(*data.PlaceID))
	}
	if data.ReplaceInstances != nil {
		_ = writer.WriteField("ReplaceInstances", fmt.Sprint(*data.ReplaceInstances))
	}
	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, writer.FormDataContentType())
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.CallRequest(ctx, e, req, shutdownAllGameInstancesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return shutdownAllGameInstancesResponse, resp, nil
}

type MatchmakingAttributeDataType string

const (
	MatchmakingAttributeDataTypeInvalid MatchmakingAttributeDataType = "Invalid"
	MatchmakingAttributeDataTypeBool    MatchmakingAttributeDataType = "Bool"
	MatchmakingAttributeDataTypeNumber  MatchmakingAttributeDataType = "Number"
	MatchmakingAttributeDataTypeString  MatchmakingAttributeDataType = "String"
)

type MatchmakingAttributeValueType string

const (
	MatchmakingAttributeValueTypeInvalid MatchmakingAttributeValueType = "Invalid"
	MatchmakingAttributeValueTypeBool    MatchmakingAttributeValueType = "Bool"
	MatchmakingAttributeValueTypeNumber  MatchmakingAttributeValueType = "Number"
	MatchmakingAttributeValueTypeString  MatchmakingAttributeValueType = "String"
)

type MatchmakingAttributeValue struct {
	BoolValue    *bool                          `json:"boolValue,omitempty"`
	NumericValue *float64                       `json:"numericValue,omitempty"`
	StringValue  *string                        `json:"stringValue,omitempty"`
	Type         *MatchmakingAttributeValueType `json:"type,omitempty"`
}

type MatchmakingAttributeValueLocationCase string

const (
	MatchmakingAttributeValueLocationCaseInvalid           MatchmakingAttributeValueLocationCase = "Invalid"
	MatchmakingAttributeValueLocationCaseDataStoreLocation MatchmakingAttributeValueLocationCase = "DataStoreLocation"
)

type DataStoreLocation struct {
	DataStoreName *string `json:"dataStoreName,omitempty"`
	Scope         *string `json:"scope,omitempty"`
	KeyTemplate   *string `json:"keyTemplate,omitempty"`
	ValuePath     *string `json:"valuePath,omitempty"`
}

type MatchmakingAttributeValueLocation struct {
	LocationCase      *MatchmakingAttributeValueLocationCase `json:"locationCase,omitempty"`
	DataStoreLocation *DataStoreLocation                     `json:"dataStoreLocation,omitempty"`
}

type CreateMatchmakingPlayerAttributeDefinitionRequest struct {
	UniverseID             *int64                             `json:"universeId,omitempty"`
	Name                   *string                            `json:"name,omitempty"`
	DataType               *MatchmakingAttributeDataType      `json:"dataType,omitempty"`
	DefaultValue           *MatchmakingAttributeValue         `json:"defaultValue,omitempty"`
	AttributeValueLocation *MatchmakingAttributeValueLocation `json:"attributeValueLocation,omitempty"`
}

type MatchmakingPlayerAttributeDefinition struct {
	ID                     string                            `json:"id"`
	UniverseID             int64                             `json:"universeId"`
	Name                   string                            `json:"name"`
	DataType               MatchmakingAttributeDataType      `json:"dataType"`
	DefaultValue           MatchmakingAttributeValue         `json:"defaultValue"`
	AttributeValueLocation MatchmakingAttributeValueLocation `json:"attributeValueLocation"`
	CreatedTime            string                            `json:"createdTime"`
	UpdatedTime            string                            `json:"updatedTime"`
}

type CreateMatchmakingPlayerAttributeDefinitionResponse struct {
	PlayerAttributeDefinition MatchmakingPlayerAttributeDefinition `json:"playerAttributeDefinition"`
}

// CreatePlayerAttributeDefinition will create a PlayerAttributeDefinition.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_matchmaking_player_attribute
//
// [POST] /matchmaking-api/v1/matchmaking/player-attribute
func (s *MatchmakingService) CreatePlayerAttributeDefinition(ctx context.Context, data CreateMatchmakingPlayerAttributeDefinitionRequest, reqOpts ...goblox.RequestOpt) (*CreateMatchmakingPlayerAttributeDefinitionResponse, *Response, error) {
	createMatchmakingPlayerAttributeDefinitionResponse := new(CreateMatchmakingPlayerAttributeDefinitionResponse)
	resp, err := s.client.Call(ctx, CreatePlayerAttributeDefinition.Compile(nil), data, createMatchmakingPlayerAttributeDefinitionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return createMatchmakingPlayerAttributeDefinitionResponse, resp, nil
}

type DeleteMatchmakingPlayerAttributeDefinitionResponse map[string]any

// DeletePlayerAttributeDefinition will delete the PlayerAttributeDefinition specified by attributeId.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#delete_matchmaking_api_v1_matchmaking_player_attribute__attributeId_
//
// [DELETE] /matchmaking-api/v1/matchmaking/player-attribute/{attributeId}
func (s *MatchmakingService) DeletePlayerAttributeDefinition(ctx context.Context, attributeId string, reqOpts ...goblox.RequestOpt) (*DeleteMatchmakingPlayerAttributeDefinitionResponse, *Response, error) {
	deleteMatchmakingPlayerAttributeDefinitionResponse := new(DeleteMatchmakingPlayerAttributeDefinitionResponse)
	resp, err := s.client.Call(ctx, DeletePlayerAttributeDefinition.Compile(nil, attributeId), nil, deleteMatchmakingPlayerAttributeDefinitionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return deleteMatchmakingPlayerAttributeDefinitionResponse, resp, nil
}

type UpdateMatchmakingPlayerAttributeDefinitionRequest struct {
	AttributeID            *string                            `json:"attributeId,omitempty"`
	DefaultValue           *MatchmakingAttributeValue         `json:"defaultValue,omitempty"`
	AttributeValueLocation *MatchmakingAttributeValueLocation `json:"attributeValueLocation,omitempty"`
}

type UpdateMatchmakingPlayerAttributeDefinitionResponse struct {
	PlayerAttributeDefinition MatchmakingPlayerAttributeDefinition `json:"playerAttributeDefinition"`
}

// UpdatePlayerAttributeDefinition will update the PlayerAttributeDefinition specified by attributeId.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#patch_matchmaking_api_v1_matchmaking_player_attribute__attributeId_
//
// [PATCH] /matchmaking-api/v1/matchmaking/player-attribute/{attributeId}
func (s *MatchmakingService) UpdatePlayerAttributeDefinition(ctx context.Context, attributeId string, data UpdateMatchmakingPlayerAttributeDefinitionRequest, reqOpts ...goblox.RequestOpt) (*UpdateMatchmakingPlayerAttributeDefinitionResponse, *Response, error) {
	updateMatchmakingPlayerAttributeDefinitionResponse := new(UpdateMatchmakingPlayerAttributeDefinitionResponse)
	resp, err := s.client.Call(ctx, UpdatePlayerAttributeDefinition.Compile(nil, attributeId), data, updateMatchmakingPlayerAttributeDefinitionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateMatchmakingPlayerAttributeDefinitionResponse, resp, nil
}

type ListMatchmakingPlayerAttributeDefinitionsResponse struct {
	PlayerAttributeSchema []MatchmakingPlayerAttributeDefinition `json:"playerAttributeSchema"`
}

// ListPlayerAttributeDefinitions will list all PlayerAttributeDefinitions of a universe.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_matchmaking_player_attributes__universeId_
//
// [GET] /matchmaking-api/v1/matchmaking/player-attributes/{universeId}
func (s *MatchmakingService) ListPlayerAttributeDefinitions(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*ListMatchmakingPlayerAttributeDefinitionsResponse, *Response, error) {
	listMatchmakingPlayerAttributeDefinitionsResponse := new(ListMatchmakingPlayerAttributeDefinitionsResponse)
	resp, err := s.client.Call(ctx, ListPlayerAttributeDefinitions.Compile(nil, universeId), nil, listMatchmakingPlayerAttributeDefinitionsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listMatchmakingPlayerAttributeDefinitionsResponse, resp, nil
}

type CreateMatchmakingScoringConfigurationRequestMatchmakingSignalWeights struct {
	Invalid          *float64 `json:"Invalid,omitempty"`
	Occupancy        *float64 `json:"Occupancy,omitempty"`
	Age              *float64 `json:"Age,omitempty"`
	Language         *float64 `json:"Language,omitempty"`
	Latency          *float64 `json:"Latency,omitempty"`
	PreferredPlayers *float64 `json:"PreferredPlayers,omitempty"`
	VoiceChat        *float64 `json:"VoiceChat,omitempty"`
	DeviceType       *float64 `json:"DeviceType,omitempty"`
	PlayHistory      *float64 `json:"PlayHistory,omitempty"`
	TextChat         *float64 `json:"TextChat,omitempty"`
}

type MatchmakingScoringConfigurationTemplate string

const (
	MatchmakingScoringConfigurationTemplateNotApplicable MatchmakingScoringConfigurationTemplate = "NotApplicable"
	MatchmakingScoringConfigurationTemplateUseDefault    MatchmakingScoringConfigurationTemplate = "UseDefault"
)

type CreateMatchmakingScoringConfigurationRequest struct {
	UniverseID               *int64                                                                `json:"universeId,omitempty"`
	Name                     *string                                                               `json:"name,omitempty"`
	MatchmakingSignalWeights *CreateMatchmakingScoringConfigurationRequestMatchmakingSignalWeights `json:"matchmakingSignalWeights,omitempty"`
	CustomSignalWeights      map[string]float64                                                    `json:"customSignalWeights,omitempty"`
	Template                 *MatchmakingScoringConfigurationTemplate                              `json:"template,omitempty"`
}

type MatchmakingSignalType string

const (
	MatchmakingSignalTypeInvalid          MatchmakingSignalType = "Invalid"
	MatchmakingSignalTypeOccupancy        MatchmakingSignalType = "Occupancy"
	MatchmakingSignalTypeAge              MatchmakingSignalType = "Age"
	MatchmakingSignalTypeLanguage         MatchmakingSignalType = "Language"
	MatchmakingSignalTypeLatency          MatchmakingSignalType = "Latency"
	MatchmakingSignalTypePreferredPlayers MatchmakingSignalType = "PreferredPlayers"
	MatchmakingSignalTypeVoiceChat        MatchmakingSignalType = "VoiceChat"
	MatchmakingSignalTypeDeviceType       MatchmakingSignalType = "DeviceType"
	MatchmakingSignalTypePlayHistory      MatchmakingSignalType = "PlayHistory"
	MatchmakingSignalTypeTextChat         MatchmakingSignalType = "TextChat"
)

type MatchmakingSignalWeightConfiguration struct {
	Weight     float64               `json:"weight"`
	SignalType MatchmakingSignalType `json:"signalType"`
}

type MatchmakingScoringConfigurationSignalWeights struct {
	Invalid          MatchmakingSignalWeightConfiguration `json:"Invalid"`
	Occupancy        MatchmakingSignalWeightConfiguration `json:"Occupancy"`
	Age              MatchmakingSignalWeightConfiguration `json:"Age"`
	Language         MatchmakingSignalWeightConfiguration `json:"Language"`
	Latency          MatchmakingSignalWeightConfiguration `json:"Latency"`
	PreferredPlayers MatchmakingSignalWeightConfiguration `json:"PreferredPlayers"`
	VoiceChat        MatchmakingSignalWeightConfiguration `json:"VoiceChat"`
	DeviceType       MatchmakingSignalWeightConfiguration `json:"DeviceType"`
	PlayHistory      MatchmakingSignalWeightConfiguration `json:"PlayHistory"`
	TextChat         MatchmakingSignalWeightConfiguration `json:"TextChat"`
}

type CustomSignalType string

const (
	CustomSignalTypeInvalid           CustomSignalType = "Invalid"
	CustomSignalTypePlayerCategorical CustomSignalType = "PlayerCategorical"
	CustomSignalTypeServerCategorical CustomSignalType = "ServerCategorical"
	CustomSignalTypeServerNumerical   CustomSignalType = "ServerNumerical"
	CustomSignalTypePlayerNumerical   CustomSignalType = "PlayerNumerical"
)

type MatchmakingAttributeReference struct {
	ID *string `json:"id,omitempty"`
}

type MatchmakingSignalCurveType string

const (
	MatchmakingSignalCurveTypeInvalid        MatchmakingSignalCurveType = "Invalid"
	MatchmakingSignalCurveTypePositiveLinear MatchmakingSignalCurveType = "PositiveLinear"
	MatchmakingSignalCurveTypeNegativeLinear MatchmakingSignalCurveType = "NegativeLinear"
)

type PlayerCategoricalSignalConfiguration struct {
	PlayerAttribute *MatchmakingAttributeReference `json:"playerAttribute,omitempty"`
	CurveType       *MatchmakingSignalCurveType    `json:"curveType,omitempty"`
}

type MatchmakingCategoricalAttributeComparisonType string

const (
	MatchmakingCategoricalAttributeComparisonTypeInvalid            MatchmakingCategoricalAttributeComparisonType = "Invalid"
	MatchmakingCategoricalAttributeComparisonTypeEqualityToPlayer   MatchmakingCategoricalAttributeComparisonType = "EqualityToPlayer"
	MatchmakingCategoricalAttributeComparisonTypeEqualityToConstant MatchmakingCategoricalAttributeComparisonType = "EqualityToConstant"
)

type ServerCategoricalSignalConfiguration struct {
	ServerAttribute *MatchmakingAttributeReference                 `json:"serverAttribute,omitempty"`
	ComparisonType  *MatchmakingCategoricalAttributeComparisonType `json:"comparisonType,omitempty"`
	PlayerAttribute *MatchmakingAttributeReference                 `json:"playerAttribute,omitempty"`
	ConstantValue   *string                                        `json:"constantValue,omitempty"`
}

type MatchmakingNumericalAttributeComparisonType string

const (
	MatchmakingNumericalAttributeComparisonTypeInvalid                     MatchmakingNumericalAttributeComparisonType = "Invalid"
	MatchmakingNumericalAttributeComparisonTypeAbsoluteDifferentToPlayer   MatchmakingNumericalAttributeComparisonType = "AbsoluteDifferentToPlayer"
	MatchmakingNumericalAttributeComparisonTypeAbsoluteDifferentToConstant MatchmakingNumericalAttributeComparisonType = "AbsoluteDifferentToConstant"
)

type ServerNumericalSignalConfiguration struct {
	ServerAttribute       *MatchmakingAttributeReference               `json:"serverAttribute,omitempty"`
	ComparisonType        *MatchmakingNumericalAttributeComparisonType `json:"comparisonType,omitempty"`
	MaxRelevantDifference *float64                                     `json:"maxRelevantDifference,omitempty"`
	ConstantValue         *float64                                     `json:"constantValue,omitempty"`
	PlayerAttribute       *MatchmakingAttributeReference               `json:"playerAttribute,omitempty"`
}

type AttributeAggregationFunction string

const (
	AttributeAggregationFunctionInvalid AttributeAggregationFunction = "Invalid"
	AttributeAggregationFunctionAverage AttributeAggregationFunction = "Average"
	AttributeAggregationFunctionMedian  AttributeAggregationFunction = "Median"
	AttributeAggregationFunctionSum     AttributeAggregationFunction = "Sum"
)

type MatchmakingAttributeAggregation struct {
	AggregationFunction *AttributeAggregationFunction `json:"aggregationFunction,omitempty"`
}

type PlayerNumericalSignalConfiguration struct {
	PlayerAttribute       *MatchmakingAttributeReference               `json:"playerAttribute,omitempty"`
	Aggregation           *MatchmakingAttributeAggregation             `json:"aggregation,omitempty"`
	ComparisonType        *MatchmakingNumericalAttributeComparisonType `json:"comparisonType,omitempty"`
	MaxRelevantDifference *float64                                     `json:"maxRelevantDifference,omitempty"`
	ConstantValue         *float64                                     `json:"constantValue,omitempty"`
}

type CustomSignalConfiguration struct {
	Name                                 *string                               `json:"name,omitempty"`
	Weight                               *float64                              `json:"weight,omitempty"`
	CustomSignalType                     *CustomSignalType                     `json:"customSignalType,omitempty"`
	PlayerCategoricalSignalConfiguration *PlayerCategoricalSignalConfiguration `json:"playerCategoricalSignalConfiguration,omitempty"`
	ServerCategoricalSignalConfiguration *ServerCategoricalSignalConfiguration `json:"serverCategoricalSignalConfiguration,omitempty"`
	ServerNumericalSignalConfiguration   *ServerNumericalSignalConfiguration   `json:"serverNumericalSignalConfiguration,omitempty"`
	PlayerNumericalSignalConfiguration   *PlayerNumericalSignalConfiguration   `json:"playerNumericalSignalConfiguration,omitempty"`
	Description                          *string                               `json:"description,omitempty"`
}

type MatchmakingScoringConfiguration struct {
	ID            string                                       `json:"id"`
	Name          string                                       `json:"name"`
	SignalWeights MatchmakingScoringConfigurationSignalWeights `json:"signalWeights"`
	CustomSignals map[string]CustomSignalConfiguration         `json:"customSignals"`
	CreatedTime   string                                       `json:"createdTime"`
	UpdatedTime   string                                       `json:"updatedTime"`
}

type CreateMatchmakingScoringConfigurationResponse struct {
	ScoringConfiguration MatchmakingScoringConfiguration `json:"scoringConfiguration"`
}

// CreateScoringConfiguration will create a matchmaking scoring configuration.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_matchmaking_scoring_configuration
//
// [POST] /matchmaking-api/v1/matchmaking/scoring-configuration
func (s *MatchmakingService) CreateScoringConfiguration(ctx context.Context, data CreateMatchmakingScoringConfigurationRequest, reqOpts ...goblox.RequestOpt) (*CreateMatchmakingScoringConfigurationResponse, *Response, error) {
	createMatchmakingScoringConfigurationResponse := new(CreateMatchmakingScoringConfigurationResponse)
	resp, err := s.client.Call(ctx, CreateScoringConfiguration.Compile(nil), data, createMatchmakingScoringConfigurationResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return createMatchmakingScoringConfigurationResponse, resp, nil
}

type GetMatchmakingScoringDefaultWeightsResponseWeights struct {
	Invalid          float64 `json:"Invalid"`
	Occupancy        float64 `json:"Occupancy"`
	Age              float64 `json:"Age"`
	Language         float64 `json:"Language"`
	Latency          float64 `json:"Latency"`
	PreferredPlayers float64 `json:"PreferredPlayers"`
	VoiceChat        float64 `json:"VoiceChat"`
	DeviceType       float64 `json:"DeviceType"`
	PlayHistory      float64 `json:"PlayHistory"`
	TextChat         float64 `json:"TextChat"`
}

type GetMatchmakingScoringDefaultWeightsResponse struct {
	Weights GetMatchmakingScoringDefaultWeightsResponseWeights `json:"weights"`
}

// GetScoringDefaultWeights will get the default matchmaking scoring weights.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_matchmaking_scoring_configuration_default_weights
//
// [GET] /matchmaking-api/v1/matchmaking/scoring-configuration/default-weights
func (s *MatchmakingService) GetScoringDefaultWeights(ctx context.Context, reqOpts ...goblox.RequestOpt) (*GetMatchmakingScoringDefaultWeightsResponse, *Response, error) {
	getMatchmakingScoringDefaultWeightsResponse := new(GetMatchmakingScoringDefaultWeightsResponse)
	resp, err := s.client.Call(ctx, GetScoringDefaultWeights.Compile(nil), nil, getMatchmakingScoringDefaultWeightsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getMatchmakingScoringDefaultWeightsResponse, resp, nil
}

type GenerateMockServerSignalValuesOptions struct {
	Count                    *int     `url:"count,omitempty"`
	Capacity                 *int     `url:"capacity,omitempty"`
	PlayerAge                *int     `url:"playerAge,omitempty"`
	PlayerPlayHistory        *float64 `url:"playerPlayHistory,omitempty"`
	IsPlayerVoiceChatEnabled *bool    `url:"isPlayerVoiceChatEnabled,omitempty"`
}

type MockServerSignalValues struct {
	Capacity                 int     `json:"capacity"`
	Occupancy                int     `json:"occupancy"`
	HasPreferredPlayers      bool    `json:"hasPreferredPlayers"`
	PlayerAge                int     `json:"playerAge"`
	ServerAveragePlayerAge   int     `json:"serverAveragePlayerAge"`
	CommonLanguagePlayers    int     `json:"commonLanguagePlayers"`
	Latency                  int     `json:"latency"`
	CommonDevicePlayers      int     `json:"commonDevicePlayers"`
	PlayerPlayHistory        float64 `json:"playerPlayHistory"`
	ServerAveragePlayHistory float64 `json:"serverAveragePlayHistory"`
	CommonVoicePlayers       int     `json:"commonVoicePlayers"`
	CommonTextChatPlayers    int     `json:"commonTextChatPlayers"`
}

type GenerateMockServerSignalValuesResponse struct {
	ExampleGameSignalValues []MockServerSignalValues `json:"exampleGameSignalValues"`
}

// GenerateMockServerSignalValues will generate mock server signal values for testing a scoring configuration.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_matchmaking_scoring_configuration_generate_mock_servers
//
// [GET] /matchmaking-api/v1/matchmaking/scoring-configuration/generate-mock-servers
func (s *MatchmakingService) GenerateMockServerSignalValues(ctx context.Context, opts *GenerateMockServerSignalValuesOptions, reqOpts ...goblox.RequestOpt) (*GenerateMockServerSignalValuesResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	generateMockServerSignalValuesResponse := new(GenerateMockServerSignalValuesResponse)
	resp, err := s.client.Call(ctx, GenerateMockServerSignalValues.Compile(q), nil, generateMockServerSignalValuesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return generateMockServerSignalValuesResponse, resp, nil
}

type SetPlaceMatchmakingScoringConfigurationRequest struct {
	PlaceID                *int64  `json:"placeId,omitempty"`
	ScoringConfigurationID *string `json:"scoringConfigurationId,omitempty"`
}

type SetMatchmakingScoringConfigurationResponse struct {
	ScoringConfiguration MatchmakingScoringConfiguration `json:"scoringConfiguration"`
}

// SetPlaceScoringConfiguration will set a matchmaking scoring configuration for a place.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_matchmaking_scoring_configuration_place
//
// [POST] /matchmaking-api/v1/matchmaking/scoring-configuration/place
func (s *MatchmakingService) SetPlaceScoringConfiguration(ctx context.Context, data SetPlaceMatchmakingScoringConfigurationRequest, reqOpts ...goblox.RequestOpt) (*SetMatchmakingScoringConfigurationResponse, *Response, error) {
	setMatchmakingScoringConfigurationResponse := new(SetMatchmakingScoringConfigurationResponse)
	resp, err := s.client.Call(ctx, SetPlaceScoringConfiguration.Compile(nil), data, setMatchmakingScoringConfigurationResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return setMatchmakingScoringConfigurationResponse, resp, nil
}

type RemovePlaceMatchmakingScoringConfigurationResponse map[string]any

// RemovePlaceScoringConfiguration will remove the matchmaking scoring configuration for a place.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#delete_matchmaking_api_v1_matchmaking_scoring_configuration_place__placeId_
//
// [DELETE] /matchmaking-api/v1/matchmaking/scoring-configuration/place/{placeId}
func (s *MatchmakingService) RemovePlaceScoringConfiguration(ctx context.Context, placeId string, reqOpts ...goblox.RequestOpt) (*RemovePlaceMatchmakingScoringConfigurationResponse, *Response, error) {
	removePlaceMatchmakingScoringConfigurationResponse := new(RemovePlaceMatchmakingScoringConfigurationResponse)
	resp, err := s.client.Call(ctx, RemovePlaceScoringConfiguration.Compile(nil, placeId), nil, removePlaceMatchmakingScoringConfigurationResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return removePlaceMatchmakingScoringConfigurationResponse, resp, nil
}

type DeleteMatchmakingScoringConfigurationResponse map[string]any

// DeleteScoringConfiguration will delete a matchmaking scoring configuration.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#delete_matchmaking_api_v1_matchmaking_scoring_configuration__scoringConfigurationId_
//
// [DELETE] /matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}
func (s *MatchmakingService) DeleteScoringConfiguration(ctx context.Context, scoringConfigurationId string, reqOpts ...goblox.RequestOpt) (*DeleteMatchmakingScoringConfigurationResponse, *Response, error) {
	deleteMatchmakingScoringConfigurationResponse := new(DeleteMatchmakingScoringConfigurationResponse)
	resp, err := s.client.Call(ctx, DeleteScoringConfiguration.Compile(nil, scoringConfigurationId), nil, deleteMatchmakingScoringConfigurationResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return deleteMatchmakingScoringConfigurationResponse, resp, nil
}

type GetMatchmakingScoringConfigurationResponse struct {
	ScoringConfiguration MatchmakingScoringConfiguration `json:"scoringConfiguration"`
}

// GetScoringConfiguration will get a matchmaking scoring configuration.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_matchmaking_scoring_configuration__scoringConfigurationId_
//
// [GET] /matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}
func (s *MatchmakingService) GetScoringConfiguration(ctx context.Context, scoringConfigurationId string, reqOpts ...goblox.RequestOpt) (*GetMatchmakingScoringConfigurationResponse, *Response, error) {
	getMatchmakingScoringConfigurationResponse := new(GetMatchmakingScoringConfigurationResponse)
	resp, err := s.client.Call(ctx, GetScoringConfiguration.Compile(nil, scoringConfigurationId), nil, getMatchmakingScoringConfigurationResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getMatchmakingScoringConfigurationResponse, resp, nil
}

type UpdateMatchmakingScoringConfigurationRequestMatchmakingSignalWeights struct {
	Invalid          *float64 `json:"Invalid,omitempty"`
	Occupancy        *float64 `json:"Occupancy,omitempty"`
	Age              *float64 `json:"Age,omitempty"`
	Language         *float64 `json:"Language,omitempty"`
	Latency          *float64 `json:"Latency,omitempty"`
	PreferredPlayers *float64 `json:"PreferredPlayers,omitempty"`
	VoiceChat        *float64 `json:"VoiceChat,omitempty"`
	DeviceType       *float64 `json:"DeviceType,omitempty"`
	PlayHistory      *float64 `json:"PlayHistory,omitempty"`
	TextChat         *float64 `json:"TextChat,omitempty"`
}

type UpdateMatchmakingScoringConfigurationRequest struct {
	ScoringConfigurationID   *string                                                               `json:"scoringConfigurationId,omitempty"`
	Name                     *string                                                               `json:"name,omitempty"`
	MatchmakingSignalWeights *UpdateMatchmakingScoringConfigurationRequestMatchmakingSignalWeights `json:"matchmakingSignalWeights,omitempty"`
	CustomSignalWeights      map[string]float64                                                    `json:"customSignalWeights,omitempty"`
}

type UpdateMatchmakingScoringConfigurationResponse struct {
	ScoringConfiguration MatchmakingScoringConfiguration `json:"scoringConfiguration"`
}

// UpdateScoringConfiguration will update a matchmaking scoring configuration.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#patch_matchmaking_api_v1_matchmaking_scoring_configuration__scoringConfigurationId_
//
// [PATCH] /matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}
func (s *MatchmakingService) UpdateScoringConfiguration(ctx context.Context, scoringConfigurationId string, data UpdateMatchmakingScoringConfigurationRequest, reqOpts ...goblox.RequestOpt) (*UpdateMatchmakingScoringConfigurationResponse, *Response, error) {
	updateMatchmakingScoringConfigurationResponse := new(UpdateMatchmakingScoringConfigurationResponse)
	resp, err := s.client.Call(ctx, UpdateScoringConfiguration.Compile(nil, scoringConfigurationId), data, updateMatchmakingScoringConfigurationResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateMatchmakingScoringConfigurationResponse, resp, nil
}

type CreateCustomMatchmakingSignalRequest struct {
	ScoringConfigurationID *string                    `json:"scoringConfigurationId,omitempty"`
	SignalConfiguration    *CustomSignalConfiguration `json:"signalConfiguration,omitempty"`
}

type CreateCustomMatchmakingSignalResponse struct {
	ScoringConfiguration MatchmakingScoringConfiguration `json:"scoringConfiguration"`
}

// CreateScoringConfigurationSignal will create a matchmaking scoring configuration signal.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_matchmaking_scoring_configuration__scoringConfigurationId__signals
//
// [POST] /matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}/signals
func (s *MatchmakingService) CreateScoringConfigurationSignal(ctx context.Context, scoringConfigurationId string, data CreateCustomMatchmakingSignalRequest, reqOpts ...goblox.RequestOpt) (*CreateCustomMatchmakingSignalResponse, *Response, error) {
	createCustomMatchmakingSignalResponse := new(CreateCustomMatchmakingSignalResponse)
	resp, err := s.client.Call(ctx, CreateScoringConfigurationSignal.Compile(nil, scoringConfigurationId), data, createCustomMatchmakingSignalResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return createCustomMatchmakingSignalResponse, resp, nil
}

type DeleteCustomMatchmakingSignalResponse struct {
	ScoringConfiguration MatchmakingScoringConfiguration `json:"scoringConfiguration"`
}

// DeleteScoringConfigurationSignal will delete a matchmaking scoring configuration custom signal.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#delete_matchmaking_api_v1_matchmaking_scoring_configuration__scoringConfigurationId__signals__signalName_
//
// [DELETE] /matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}/signals/{signalName}
func (s *MatchmakingService) DeleteScoringConfigurationSignal(ctx context.Context, scoringConfigurationId, signalName string, reqOpts ...goblox.RequestOpt) (*DeleteCustomMatchmakingSignalResponse, *Response, error) {
	deleteCustomMatchmakingSignalResponse := new(DeleteCustomMatchmakingSignalResponse)
	resp, err := s.client.Call(ctx, DeleteScoringConfigurationSignal.Compile(nil, scoringConfigurationId, signalName), nil, deleteCustomMatchmakingSignalResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return deleteCustomMatchmakingSignalResponse, resp, nil
}

type UpdateCustomMatchmakingSignalRequest struct {
	ScoringConfigurationID *string                    `json:"scoringConfigurationId,omitempty"`
	SignalConfiguration    *CustomSignalConfiguration `json:"signalConfiguration,omitempty"`
}

type UpdateCustomMatchmakingSignalResponse struct {
	ScoringConfiguration MatchmakingScoringConfiguration `json:"scoringConfiguration"`
}

// UpdateScoringConfigurationSignal will update a matchmaking scoring configuration signal.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#patch_matchmaking_api_v1_matchmaking_scoring_configuration__scoringConfigurationId__signals__signalName_
//
// [PATCH] /matchmaking-api/v1/matchmaking/scoring-configuration/{scoringConfigurationId}/signals/{signalName}
func (s *MatchmakingService) UpdateScoringConfigurationSignal(ctx context.Context, scoringConfigurationId, signalName string, data UpdateCustomMatchmakingSignalRequest, reqOpts ...goblox.RequestOpt) (*UpdateCustomMatchmakingSignalResponse, *Response, error) {
	updateCustomMatchmakingSignalResponse := new(UpdateCustomMatchmakingSignalResponse)
	resp, err := s.client.Call(ctx, UpdateScoringConfigurationSignal.Compile(nil, scoringConfigurationId, signalName), data, updateCustomMatchmakingSignalResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateCustomMatchmakingSignalResponse, resp, nil
}

type ListMatchmakingScoringConfigurationsResponse struct {
	ScoringConfigurations []MatchmakingScoringConfiguration `json:"scoringConfigurations"`
}

// ListScoringConfigurations will list all matchmaking scoring configurations for a universe.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_matchmaking_scoring_configurations__universeId_
//
// [GET] /matchmaking-api/v1/matchmaking/scoring-configurations/{universeId}
func (s *MatchmakingService) ListScoringConfigurations(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*ListMatchmakingScoringConfigurationsResponse, *Response, error) {
	listMatchmakingScoringConfigurationsResponse := new(ListMatchmakingScoringConfigurationsResponse)
	resp, err := s.client.Call(ctx, ListScoringConfigurations.Compile(nil, universeId), nil, listMatchmakingScoringConfigurationsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listMatchmakingScoringConfigurationsResponse, resp, nil
}

type ListPlaceMatchmakingScoringConfigurationsResponse struct {
	PlaceScoringConfigurations map[string]MatchmakingScoringConfiguration `json:"placeScoringConfigurations"`
}

// ListPlaceScoringConfigurations will list all places with a matchmaking scoring configuration for a universe.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_matchmaking_scoring_configurations__universeId__places
//
// [GET] /matchmaking-api/v1/matchmaking/scoring-configurations/{universeId}/places
func (s *MatchmakingService) ListPlaceScoringConfigurations(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*ListPlaceMatchmakingScoringConfigurationsResponse, *Response, error) {
	listPlaceMatchmakingScoringConfigurationsResponse := new(ListPlaceMatchmakingScoringConfigurationsResponse)
	resp, err := s.client.Call(ctx, ListPlaceScoringConfigurations.Compile(nil, universeId), nil, listPlaceMatchmakingScoringConfigurationsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listPlaceMatchmakingScoringConfigurationsResponse, resp, nil
}

type MatchmakingServerAttributeDefaultValueSourceCase string

const (
	MatchmakingServerAttributeDefaultValueSourceCaseInvalid                  MatchmakingServerAttributeDefaultValueSourceCase = "Invalid"
	MatchmakingServerAttributeDefaultValueSourceCaseConstant                 MatchmakingServerAttributeDefaultValueSourceCase = "Constant"
	MatchmakingServerAttributeDefaultValueSourceCasePlayerAttributeReference MatchmakingServerAttributeDefaultValueSourceCase = "PlayerAttributeReference"
)

type MatchmakingServerAttributeDefaultValue struct {
	SourceCase               *MatchmakingServerAttributeDefaultValueSourceCase `json:"sourceCase,omitempty"`
	Constant                 *MatchmakingAttributeValue                        `json:"constant,omitempty"`
	PlayerAttributeReference *MatchmakingAttributeReference                    `json:"playerAttributeReference,omitempty"`
}

type CreateMatchmakingServerAttributeDefinitionRequest struct {
	UniverseID   *int64                                  `json:"universeId,omitempty"`
	Name         *string                                 `json:"name,omitempty"`
	DataType     *MatchmakingAttributeDataType           `json:"dataType,omitempty"`
	DefaultValue *MatchmakingServerAttributeDefaultValue `json:"defaultValue,omitempty"`
}

type MatchmakingServerAttributeDefinition struct {
	ID           string                                 `json:"id"`
	UniverseID   int64                                  `json:"universeId"`
	Name         string                                 `json:"name"`
	DataType     MatchmakingAttributeDataType           `json:"dataType"`
	DefaultValue MatchmakingServerAttributeDefaultValue `json:"defaultValue"`
	CreatedTime  string                                 `json:"createdTime"`
	UpdatedTime  string                                 `json:"updatedTime"`
}

type CreateMatchmakingServerAttributeDefinitionResponse struct {
	PlayerAttributeDefinition MatchmakingServerAttributeDefinition `json:"playerAttributeDefinition"`
}

// CreateServerAttributeDefinition will create a ServerAttributeDefinition.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_matchmaking_server_attribute
//
// [POST] /matchmaking-api/v1/matchmaking/server-attribute
func (s *MatchmakingService) CreateServerAttributeDefinition(ctx context.Context, data CreateMatchmakingServerAttributeDefinitionRequest, reqOpts ...goblox.RequestOpt) (*CreateMatchmakingServerAttributeDefinitionResponse, *Response, error) {
	createMatchmakingServerAttributeDefinitionResponse := new(CreateMatchmakingServerAttributeDefinitionResponse)
	resp, err := s.client.Call(ctx, CreateServerAttributeDefinition.Compile(nil), data, createMatchmakingServerAttributeDefinitionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return createMatchmakingServerAttributeDefinitionResponse, resp, nil
}

type DeleteMatchmakingServerAttributeDefinitionResponse map[string]any

// DeleteServerAttributeDefinition will delete the ServerAttributeDefinition specified by attributeId.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#delete_matchmaking_api_v1_matchmaking_server_attribute__attributeId_
//
// [DELETE] /matchmaking-api/v1/matchmaking/server-attribute/{attributeId}
func (s *MatchmakingService) DeleteServerAttributeDefinition(ctx context.Context, attributeId string, reqOpts ...goblox.RequestOpt) (*DeleteMatchmakingServerAttributeDefinitionResponse, *Response, error) {
	deleteMatchmakingServerAttributeDefinitionResponse := new(DeleteMatchmakingServerAttributeDefinitionResponse)
	resp, err := s.client.Call(ctx, DeleteServerAttributeDefinition.Compile(nil, attributeId), nil, deleteMatchmakingServerAttributeDefinitionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return deleteMatchmakingServerAttributeDefinitionResponse, resp, nil
}

type UpdateMatchmakingServerAttributeDefinitionRequest struct {
	AttributeID  *string                                 `json:"attributeId,omitempty"`
	DefaultValue *MatchmakingServerAttributeDefaultValue `json:"defaultValue,omitempty"`
}

type UpdateMatchmakingServerAttributeDefinitionResponse struct {
	PlayerAttributeDefinition MatchmakingServerAttributeDefinition `json:"playerAttributeDefinition"`
}

// UpdateServerAttributeDefinition will update the ServerAttributeDefinition specified by attributeId.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#patch_matchmaking_api_v1_matchmaking_server_attribute__attributeId_
//
// [PATCH] /matchmaking-api/v1/matchmaking/server-attribute/{attributeId}
func (s *MatchmakingService) UpdateServerAttributeDefinition(ctx context.Context, attributeId string, data UpdateMatchmakingServerAttributeDefinitionRequest, reqOpts ...goblox.RequestOpt) (*UpdateMatchmakingServerAttributeDefinitionResponse, *Response, error) {
	updateMatchmakingServerAttributeDefinitionResponse := new(UpdateMatchmakingServerAttributeDefinitionResponse)
	resp, err := s.client.Call(ctx, UpdateServerAttributeDefinition.Compile(nil, attributeId), data, updateMatchmakingServerAttributeDefinitionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateMatchmakingServerAttributeDefinitionResponse, resp, nil
}

type ListMatchmakingServerAttributeDefinitionsResponse struct {
	ServerAttributeSchema []MatchmakingServerAttributeDefinition `json:"serverAttributeSchema"`
}

// ListServerAttributeDefinitions will list all ServerAttributeDefinitions of a universe.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_matchmaking_server_attributes__universeId_
//
// [GET] /matchmaking-api/v1/matchmaking/server-attributes/{universeId}
func (s *MatchmakingService) ListServerAttributeDefinitions(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*ListMatchmakingServerAttributeDefinitionsResponse, *Response, error) {
	listMatchmakingServerAttributeDefinitionsResponse := new(ListMatchmakingServerAttributeDefinitionsResponse)
	resp, err := s.client.Call(ctx, ListServerAttributeDefinitions.Compile(nil, universeId), nil, listMatchmakingServerAttributeDefinitionsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listMatchmakingServerAttributeDefinitionsResponse, resp, nil
}

type MatchmakingCustomizationFeatureFlags struct {
	IsMatchmakingCustomizationAllowed            bool `json:"isMatchmakingCustomizationAllowed"`
	IsMatchmakingCustomizationExperimentsAllowed bool `json:"isMatchmakingCustomizationExperimentsAllowed"`
	IsMatchmakingTextChatSignalEnabled           bool `json:"isMatchmakingTextChatSignalEnabled"`
}

type GetMatchmakingCustomizationFeatureFlagsResponse struct {
	FeatureFlags MatchmakingCustomizationFeatureFlags `json:"featureFlags"`
}

// GetMatchmakingFeatureFlags will get feature flags for a customized matchmaking for a given universe.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_matchmaking_universe__universeId__feature_flags
//
// [GET] /matchmaking-api/v1/matchmaking/universe/{universeId}/feature-flags
func (s *MatchmakingService) GetMatchmakingFeatureFlags(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*GetMatchmakingCustomizationFeatureFlagsResponse, *Response, error) {
	getMatchmakingCustomizationFeatureFlagsResponse := new(GetMatchmakingCustomizationFeatureFlagsResponse)
	resp, err := s.client.Call(ctx, GetMatchmakingFeatureFlags.Compile(nil, universeId), nil, getMatchmakingCustomizationFeatureFlagsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getMatchmakingCustomizationFeatureFlagsResponse, resp, nil
}

type DeleteStartPlaceOverrideResponse map[string]any

// DeleteStartPlaceOverride will delete the start place override for a universe.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#delete_matchmaking_api_v1_matchmaking_universe__universeId__start_place
//
// [DELETE] /matchmaking-api/v1/matchmaking/universe/{universeId}/start-place
func (s *MatchmakingService) DeleteStartPlaceOverride(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*DeleteStartPlaceOverrideResponse, *Response, error) {
	deleteStartPlaceOverrideResponse := new(DeleteStartPlaceOverrideResponse)
	resp, err := s.client.Call(ctx, DeleteStartPlaceOverride.Compile(nil, universeId), nil, deleteStartPlaceOverrideResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return deleteStartPlaceOverrideResponse, resp, nil
}

type GetStartPlaceOverrideResponse struct {
	DataStoreLocation DataStoreLocation `json:"dataStoreLocation"`
}

// GetStartPlaceOverride will get the start place override for a universe.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_matchmaking_universe__universeId__start_place
//
// [GET] /matchmaking-api/v1/matchmaking/universe/{universeId}/start-place
func (s *MatchmakingService) GetStartPlaceOverride(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*GetStartPlaceOverrideResponse, *Response, error) {
	getStartPlaceOverrideResponse := new(GetStartPlaceOverrideResponse)
	resp, err := s.client.Call(ctx, GetStartPlaceOverride.Compile(nil, universeId), nil, getStartPlaceOverrideResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getStartPlaceOverrideResponse, resp, nil
}

type SetStartPlaceOverrideRequestBody struct {
	DataStoreLocation *DataStoreLocation `json:"dataStoreLocation,omitempty"`
}

type SetStartPlaceOverrideResponse map[string]any

// SetStartPlaceOverride will set (create or update) the start place override for a universe.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#post_matchmaking_api_v1_matchmaking_universe__universeId__start_place
//
// [POST] /matchmaking-api/v1/matchmaking/universe/{universeId}/start-place
func (s *MatchmakingService) SetStartPlaceOverride(ctx context.Context, universeId string, data SetStartPlaceOverrideRequestBody, reqOpts ...goblox.RequestOpt) (*SetStartPlaceOverrideResponse, *Response, error) {
	setStartPlaceOverrideResponse := new(SetStartPlaceOverrideResponse)
	resp, err := s.client.Call(ctx, SetStartPlaceOverride.Compile(nil, universeId), data, setStartPlaceOverrideResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return setStartPlaceOverrideResponse, resp, nil
}

type TestStartPlaceOverrideOptions struct {
	UserID *int64 `url:"userId,omitempty"`
}

type StartPlaceOverrideTestResult string

const (
	StartPlaceOverrideTestResultInvalid                StartPlaceOverrideTestResult = "Invalid"
	StartPlaceOverrideTestResultFound                  StartPlaceOverrideTestResult = "Found"
	StartPlaceOverrideTestResultNoOverrideConfigured   StartPlaceOverrideTestResult = "NoOverrideConfigured"
	StartPlaceOverrideTestResultDataStoreEntryNotFound StartPlaceOverrideTestResult = "DataStoreEntryNotFound"
	StartPlaceOverrideTestResultDataStoreNotFound      StartPlaceOverrideTestResult = "DataStoreNotFound"
	StartPlaceOverrideTestResultDataStoreError         StartPlaceOverrideTestResult = "DataStoreError"
	StartPlaceOverrideTestResultFieldMissing           StartPlaceOverrideTestResult = "FieldMissing"
	StartPlaceOverrideTestResultValueParseFailure      StartPlaceOverrideTestResult = "ValueParseFailure"
	StartPlaceOverrideTestResultPlaceNotInExperience   StartPlaceOverrideTestResult = "PlaceNotInExperience"
	StartPlaceOverrideTestResultPlaceNotPublished      StartPlaceOverrideTestResult = "PlaceNotPublished"
)

type TestStartPlaceOverrideResponse struct {
	Result  StartPlaceOverrideTestResult `json:"result"`
	PlaceID int64                        `json:"placeId"`
}

// TestStartPlaceOverride will test a user's saved start place in a universe against the place checks gamejoin
// makes.
//
// Required scopes: none
//
// Roblox's spec doesn't list API key or OAuth2 support for this endpoint, so the credential kind isn't checked.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/matchmaking#get_matchmaking_api_v1_matchmaking_universe__universeId__start_place_test
//
// [GET] /matchmaking-api/v1/matchmaking/universe/{universeId}/start-place/test
func (s *MatchmakingService) TestStartPlaceOverride(ctx context.Context, universeId string, opts *TestStartPlaceOverrideOptions, reqOpts ...goblox.RequestOpt) (*TestStartPlaceOverrideResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	testStartPlaceOverrideResponse := new(TestStartPlaceOverrideResponse)
	resp, err := s.client.Call(ctx, TestStartPlaceOverride.Compile(q, universeId), nil, testStartPlaceOverrideResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return testStartPlaceOverrideResponse, resp, nil
}
