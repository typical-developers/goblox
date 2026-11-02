package opencloud

import (
	"context"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	GetGameServerFilterOptions = goblox.NewEndpoint(
		http.MethodGet,
		"/server-management/v1/universes/{universeId}/places/{placeId}/game-servers:filter-options",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:read")

	ListGameServers = goblox.NewEndpoint(
		http.MethodGet,
		"/server-management/v1/universes/{universeId}/places/{placeId}/versions/{versionNumber}/game-servers",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:read")

	ListGameServerLogs = goblox.NewEndpoint(
		http.MethodGet,
		"/server-management/v1/universes/{universeId}/places/{placeId}/versions/{versionNumber}/game-servers/{jobId}/logs",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:read")

	ListRestartStatuses = goblox.NewEndpoint(
		http.MethodGet,
		"/server-management/v1/universes/{universeId}/restarts",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:read")

	LaunchRestart = goblox.NewEndpoint(
		http.MethodPost,
		"/server-management/v1/universes/{universeId}/restarts",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:write")

	ForecastRestart = goblox.NewEndpoint(
		http.MethodGet,
		"/server-management/v1/universes/{universeId}/restarts:forecast",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:read")
)

// ServerManagementService handles communication with the server management endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type ServerManagementService service

type FilterField string

const (
	FilterFieldEngineVersion FilterField = "EngineVersion"
	FilterFieldPlaceVersion  FilterField = "PlaceVersion"
	FilterFieldAll           FilterField = "ALL"
)

type GetGameServerFilterOptionsOptions struct {
	Filter *FilterField `url:"Filter,omitempty"`
}

type ServerManagementFilterType string

const (
	ServerManagementFilterTypeString ServerManagementFilterType = "String"
	ServerManagementFilterTypeNumber ServerManagementFilterType = "Number"
)

type FilterFieldInfo struct {
	Field  FilterField                `json:"field"`
	Type   ServerManagementFilterType `json:"type"`
	Values []any                      `json:"values"`
}

type FilterOptionsResponse struct {
	Filters map[string]FilterFieldInfo `json:"filters"`
}

// GetGameServerFilterOptions will get available filter options for game servers.
//
// Required scopes: universe:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/observability#GameServers_GetFilterOptions
//
// [GET] /server-management/v1/universes/{universeId}/places/{placeId}/game-servers:filter-options
func (s *ServerManagementService) GetGameServerFilterOptions(ctx context.Context, universeId, placeId string, opts *GetGameServerFilterOptionsOptions, reqOpts ...goblox.RequestOpt) (*FilterOptionsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	filterOptionsResponse := new(FilterOptionsResponse)
	resp, err := s.client.Call(ctx, GetGameServerFilterOptions.Compile(q, universeId, placeId), nil, filterOptionsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return filterOptionsResponse, resp, nil
}

type ListGameServersOptions struct {
	MaxPageSize *int    `url:"MaxPageSize,omitempty"`
	PageToken   *string `url:"PageToken,omitempty"`
	OrderBy     *string `url:"OrderBy,omitempty"`
	Filter      *string `url:"Filter,omitempty"`
}

type ServerManagementServerType int

type ServerManagementServerStatus string

const (
	ServerManagementServerStatusActive          ServerManagementServerStatus = "active"
	ServerManagementServerStatusShutDown        ServerManagementServerStatus = "shut_down"
	ServerManagementServerStatusRestarted       ServerManagementServerStatus = "restarted"
	ServerManagementServerStatusRobloxRestarted ServerManagementServerStatus = "roblox_restarted"
	ServerManagementServerStatusCrashed         ServerManagementServerStatus = "crashed"
	ServerManagementServerStatusOutOfMemory     ServerManagementServerStatus = "out_of_memory"
	ServerManagementServerStatusModerated       ServerManagementServerStatus = "moderated"
	ServerManagementServerStatusPending         ServerManagementServerStatus = "pending"
	ServerManagementServerStatusUnspecified     ServerManagementServerStatus = "unspecified"
)

type ServerManagementGameServer struct {
	PlaceID          string                       `json:"placeId"`
	PlaceVersion     string                       `json:"placeVersion"`
	JobID            string                       `json:"jobId"`
	EngineVersion    string                       `json:"engineVersion"`
	CreateTime       string                       `json:"createTime"`
	Uptime           string                       `json:"uptime"`
	MemoryUsageBytes int64                        `json:"memoryUsageBytes"`
	FrameRate        float64                      `json:"frameRate"`
	Occupancy        int                          `json:"occupancy"`
	MaxOccupancy     int                          `json:"maxOccupancy"`
	Type             ServerManagementServerType   `json:"type"`
	Full             bool                         `json:"full"`
	ShutDown         bool                         `json:"shutDown"`
	Status           ServerManagementServerStatus `json:"status"`
	TerminationTime  string                       `json:"terminationTime"`
	PlayerIDs        []int64                      `json:"playerIds"`
}

type ListGameServersResponse struct {
	GameServers               []ServerManagementGameServer `json:"gameServers"`
	NextPageToken             string                       `json:"nextPageToken"`
	PreviousPageToken         string                       `json:"previousPageToken"`
	TotalCount                int                          `json:"totalCount"`
	ShutdownServersFetchError bool                         `json:"shutdownServersFetchError"`
	ActiveServersFetchError   bool                         `json:"activeServersFetchError"`
}

// ListGameServers will list game servers for a specific place version.
//
// Required scopes: universe:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/observability#GameServers_ListGameServers
//
// [GET] /server-management/v1/universes/{universeId}/places/{placeId}/versions/{versionNumber}/game-servers
func (s *ServerManagementService) ListGameServers(ctx context.Context, universeId, placeId, versionNumber string, opts *ListGameServersOptions, reqOpts ...goblox.RequestOpt) (*ListGameServersResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	listGameServersResponse := new(ListGameServersResponse)
	resp, err := s.client.Call(ctx, ListGameServers.Compile(q, universeId, placeId, versionNumber), nil, listGameServersResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listGameServersResponse, resp, nil
}

type ListGameServerLogsOptions struct {
	MaxPageSize *int    `url:"MaxPageSize,omitempty"`
	PageToken   *string `url:"PageToken,omitempty"`
	OrderBy     *string `url:"OrderBy,omitempty"`
	Filter      *string `url:"Filter,omitempty"`
}

type LogSeverity int

type StructuredStackFrame struct {
	Line           int    `json:"line"`
	Function       string `json:"function"`
	ScriptPath     string `json:"scriptPath"`
	ScriptUniqueID string `json:"scriptUniqueId"`
	IsDynamic      bool   `json:"isDynamic"`
}

type GameServerLog struct {
	MessageTimestampMs   string                 `json:"messageTimestampMs"`
	UniverseID           int64                  `json:"universeId"`
	PlaceID              string                 `json:"placeId"`
	PlaceVersion         string                 `json:"placeVersion"`
	JobID                string                 `json:"jobId"`
	Severity             LogSeverity            `json:"severity"`
	Message              string                 `json:"message"`
	StackTrace           string                 `json:"stackTrace"`
	MessageTemplate      string                 `json:"messageTemplate"`
	Context              string                 `json:"context"`
	LoggerFullPath       string                 `json:"loggerFullPath"`
	SkippedCount         int                    `json:"skippedCount"`
	RateLimitedCount     int                    `json:"rateLimitedCount"`
	StructuredStackTrace []StructuredStackFrame `json:"structuredStackTrace"`
}

type ListGameServerLogsResponse struct {
	GameServerLogs []GameServerLog `json:"gameServerLogs"`
	NextPageToken  string          `json:"nextPageToken"`
}

// ListGameServerLogs will list game server logs for a specific server job id.
//
// Required scopes: universe:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/observability#GameServers_ListGameServerLogs
//
// [GET] /server-management/v1/universes/{universeId}/places/{placeId}/versions/{versionNumber}/game-servers/{jobId}/logs
func (s *ServerManagementService) ListGameServerLogs(ctx context.Context, universeId, placeId, versionNumber, jobId string, opts *ListGameServerLogsOptions, reqOpts ...goblox.RequestOpt) (*ListGameServerLogsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	listGameServerLogsResponse := new(ListGameServerLogsResponse)
	resp, err := s.client.Call(ctx, ListGameServerLogs.Compile(q, universeId, placeId, versionNumber, jobId), nil, listGameServerLogsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listGameServerLogsResponse, resp, nil
}

type RestartState string

const (
	RestartStateDelaying   RestartState = "DELAYING"
	RestartStateRestarting RestartState = "RESTARTING"
	RestartStateSucceeded  RestartState = "SUCCEEDED"
)

type PlaceFilter struct {
	Versions              []int `json:"versions,omitempty"`
	ExcludeCurrentVersion *bool `json:"excludeCurrentVersion,omitempty"`
}

type PlaceRestartStatus struct {
	State              RestartState `json:"state"`
	StartTime          string       `json:"startTime"`
	EndTime            string       `json:"endTime"`
	TotalPlayers       int          `json:"totalPlayers"`
	TotalInstances     int          `json:"totalInstances"`
	RemainingPlayers   int          `json:"remainingPlayers"`
	RemainingInstances int          `json:"remainingInstances"`
	Filter             PlaceFilter  `json:"filter"`
	LatestVersion      string       `json:"latestVersion"`
}

type RestartStatus struct {
	UniverseID           string                        `json:"universeId"`
	ScheduledTime        string                        `json:"scheduledTime"`
	StartTime            string                        `json:"startTime"`
	PlaceRestartStatuses map[string]PlaceRestartStatus `json:"placeRestartStatuses"`
}

type ListRestartStatusesResponse struct {
	RestartStatuses map[string]RestartStatus `json:"restartStatuses"`
}

// ListRestartStatuses will list restart statuses for a universe.
//
// Required scopes: universe:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/observability#Restarts_ListRestartStatuses
//
// [GET] /server-management/v1/universes/{universeId}/restarts
func (s *ServerManagementService) ListRestartStatuses(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*ListRestartStatusesResponse, *Response, error) {
	listRestartStatusesResponse := new(ListRestartStatusesResponse)
	resp, err := s.client.Call(ctx, ListRestartStatuses.Compile(nil, universeId), nil, listRestartStatusesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listRestartStatusesResponse, resp, nil
}

type LaunchRestartRequest struct {
	BleedOffDurationMinutes *int                   `json:"bleedOffDurationMinutes,omitempty"`
	Places                  map[string]PlaceFilter `json:"places,omitempty"`
	Attributes              any                    `json:"attributes,omitempty"`
}

type LaunchRestartResponse struct {
	ID                string `json:"id"`
	PlayersImpacted   int    `json:"playersImpacted"`
	InstancesImpacted int    `json:"instancesImpacted"`
}

// LaunchRestart will launch a game server restart for a universe.
//
// Required scopes: universe:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/observability#Restarts_LaunchRestart
//
// [POST] /server-management/v1/universes/{universeId}/restarts
func (s *ServerManagementService) LaunchRestart(ctx context.Context, universeId string, data LaunchRestartRequest, reqOpts ...goblox.RequestOpt) (*LaunchRestartResponse, *Response, error) {
	launchRestartResponse := new(LaunchRestartResponse)
	resp, err := s.client.Call(ctx, LaunchRestart.Compile(nil, universeId), data, launchRestartResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return launchRestartResponse, resp, nil
}

type PlaceSummaryForGameRestart struct {
	PlayersImpacted     int            `json:"playersImpacted"`
	TotalPlayers        int            `json:"totalPlayers"`
	InstancesImpacted   int            `json:"instancesImpacted"`
	TotalInstances      int            `json:"totalInstances"`
	LatestPlaceVersion  string         `json:"latestPlaceVersion"`
	PublishTime         string         `json:"publishTime"`
	IsNotInUniverse     bool           `json:"isNotInUniverse"`
	PlayersPerVersion   map[string]int `json:"playersPerVersion"`
	InstancesPerVersion map[string]int `json:"instancesPerVersion"`
}

type ForecastRestartResponse struct {
	PlaceForecasts map[string]PlaceSummaryForGameRestart `json:"placeForecasts"`
}

// ForecastRestart will forecast the impact of restarting game servers for a universe.
//
// Required scopes: universe:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/observability#Restarts_ForecastRestart
//
// [GET] /server-management/v1/universes/{universeId}/restarts:forecast
func (s *ServerManagementService) ForecastRestart(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*ForecastRestartResponse, *Response, error) {
	forecastRestartResponse := new(ForecastRestartResponse)
	resp, err := s.client.Call(ctx, ForecastRestart.Compile(nil, universeId), nil, forecastRestartResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return forecastRestartResponse, resp, nil
}
