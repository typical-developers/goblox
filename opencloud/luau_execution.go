package opencloud

import (
	"context"
	"fmt"
	"net/http"
)

// LuauExecutionService will handle communciation with the actions related to the API.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type LuauExecutionService struct {
	client *Client
}

type LuauExecutionState string

const (
	LuauExecutionStateUnspecified LuauExecutionState = "STATE_UNSPECIFIED"
	LuauExecutionStateQueued      LuauExecutionState = "QUEUED"
	LuauExecutionStateProcessing  LuauExecutionState = "PROCESSING"
	LuauExecutionStateCancelled   LuauExecutionState = "CANCELLED"
	LuauExecutionStateComplete    LuauExecutionState = "COMPLETE"
	LuauExecutionStateFailed      LuauExecutionState = "FAILED"
)

type LuauExecutionErrorCode string

const (
	LuauExecutionErrorCodeUnspecified             LuauExecutionErrorCode = "ERROR_CODE_UNSPECIFIED"
	LuauExecutionErrorCodeScriptError             LuauExecutionErrorCode = "SCRIPT_ERROR"
	LuauExecutionErrorCodeDeadlineExceeded        LuauExecutionErrorCode = "DEADLINE_EXCEEDED"
	LuauExecutionErrorCodeOutputSizeLimitExceeded LuauExecutionErrorCode = "OUTPUT_SIZE_LIMIT_EXCEEDED"
	LuauExecutionErrorCodeInternalError           LuauExecutionErrorCode = "INTERNAL_ERROR"
)

type LuauExecutionTaskError struct {
	Code    LuauExecutionErrorCode `json:"code"`
	Message string                 `json:"message"`
}

type LuauExecutionTaskOutput struct {
	Results []any `json:"results"`
}

type LuauExecutionTask struct {
	Path                string                   `json:"path"`
	CreateTime          string                   `json:"createTime"`
	UpdateTime          string                   `json:"updateTime"`
	User                string                   `json:"user"`
	Status              LuauExecutionState       `json:"state"`
	Script              string                   `json:"script"`
	Timeout             string                   `json:"timeout"`
	Error               *LuauExecutionTaskError  `json:"error"`
	Output              *LuauExecutionTaskOutput `json:"output"`
	BinaryInput         string                   `json:"binaryInput"`
	EnabledBinaryOutput bool                     `json:"enabledBinaryOutput"`
	BinaryOutput        string                   `json:"binaryOutput"`
}

// CreateLuauExecutionSessionTask will execute a Luau script on a specific place.
// If a versionId is provided, the code will be executed on that version of the place.
//
// Required scopes: universe.place.luau-execution-session:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/LuauExecutionSessionTask#Cloud_CreateLuauExecutionSessionTask
//
// [POST] /cloud/v2/universes/{universe_id}/places/{place_id}/luau-execution-session-tasks
//
// [POST] /cloud/v2/universes/{universe_id}/places/{place_id}/luau-execution-session-tasks/{version_id}
func (s *LuauExecutionService) CreateLuauExecutionSessionTask(ctx context.Context, universeId, placeId string, versionId *string) (*LuauExecutionTask, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s/places/%s/luau-execution-session-tasks", universeId, placeId)
	if versionId != nil {
		u = fmt.Sprintf("/cloud/v2/universes/%s/places/%s/luau-execution-session-tasks/%s", universeId, placeId, *versionId)
	}

	req, err := s.client.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	luauExecutionSessionTask := new(LuauExecutionTask)
	resp, err := s.client.Do(ctx, req, luauExecutionSessionTask)
	if err != nil {
		return nil, resp, err
	}

	return luauExecutionSessionTask, resp, nil
}

// GetLuauExecutionSessionTask will fetch the executed Luau task.
//
// Required scopes:
//
// - universe.place.luau-execution-session:read
//
// - universe.place.luau-execution-session:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/LuauExecutionSessionTask#Cloud_GetLuauExecutionSessionTask
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}/versions/{version_id}/luau-execution-sessions/{luau_execution_session_id}/tasks/{task_id}
func (s *LuauExecutionService) GetLuauExecutionSessionTask(ctx context.Context, universeId, placeId, versionId, taskId string) (*LuauExecutionTask, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s/places/%s/versions/%s/luau-execution-sessions/%s/tasks/%s", universeId, placeId, versionId, taskId, taskId)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	luauExecutionSessionTask := new(LuauExecutionTask)
	resp, err := s.client.Do(ctx, req, luauExecutionSessionTask)
	if err != nil {
		return nil, resp, err
	}

	return luauExecutionSessionTask, resp, nil
}

type StructuredMessageType string

const (
	StructuredMessageTypeUnspecified StructuredMessageType = "MESSAGE_TYPE_UNSPECIFIED"
	StructuredMessageTypeOutput      StructuredMessageType = "OUTPUT"
	StructuredMessageTypeInfo        StructuredMessageType = "INFO"
	StructuredMessageTypeWarning     StructuredMessageType = "WARNING"
	StructuredMessageTypeError       StructuredMessageType = "ERROR"
)

type LuauExecutionTaskLogStructuredMessage struct {
	Message     string                `json:"message"`
	CreateTime  string                `json:"createTime"`
	MessageType StructuredMessageType `json:"messageType"`
}

type LuauExecutionTaskLog struct {
	Path    string   `json:"path"`
	Mesages []string `json:"messages"`
}

type LuauExecutionTaskLogs struct {
	LuauExecutionSessionTaskLogs []LuauExecutionTaskLog `json:"luauExecutionSessionTaskLogs"`
	NextPageToken                string                 `json:"nextPageToken"`
}

// ListLuauExecutionSessionTaskLogs will list log chunks generated from a specific Luau task.
//
// Required scopes:
//
// - universe.place.luau-execution-session:read
//
// - universe.place.luau-execution-session:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/LuauExecutionSessionTaskLog#Cloud_ListLuauExecutionSessionTaskLogs
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}/versions/{version_id}/luau-execution-sessions/{luau_execution_session_id}/tasks/{task_id}/logs
func (s *LuauExecutionService) ListLuauExecutionSessionTaskLogs(ctx context.Context, universeId, placeId, versionId, taskId string, opts *Options) (*LuauExecutionTaskLogs, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s/places/%s/versions/%s/luau-execution-sessions/%s/tasks/%s/logs", universeId, placeId, versionId, taskId, taskId)

	u, err := addOpts(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	luauExecutionSessionTaskLogs := new(LuauExecutionTaskLogs)
	resp, err := s.client.Do(ctx, req, luauExecutionSessionTaskLogs)
	if err != nil {
		return nil, resp, err
	}

	return luauExecutionSessionTaskLogs, resp, nil
}
