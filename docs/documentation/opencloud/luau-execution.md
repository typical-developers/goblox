# Luau Execution
These endpoints let you execute Luau code in one of your experience(s).

::: danger
Scripts can run engine methods regardless of the scopes granted. For example, it can run datastore related methods. Proceed with caution when granting any luau-execution scopes.
:::

## Variables
### `LuauExecutionTaskPathRegex`
This regex is used to parse the URL of a path to grab the `UniverseID`, `PlaceID`, `VersionID`, `SessionID`, and `TaskID`. It is used by the [`TaskInfo`](/documentation/opencloud/luau-execution#luauexecutiontask) utility method, but is exposed for anyone to use it.
```go
regexp.MustCompile(`universes/(?<UniverseID>\d+)\/places\/(?<PlaceID>\d+)\/(versions\/(?<VersionID>\d+)\/)?(luau-execution-sessions\/(?<SessionID>.+)?\/tasks\/(?<TaskID>.+)|(luau-execution-session-tasks\/(?<TaskID>.+)))`)
```

## Constants
### `LuauExecutionState`
```go
type LuauExecutionState string

const (
	LuauExecutionStateUnspecified LuauExecutionState = "STATE_UNSPECIFIED"
	LuauExecutionStateQueued      LuauExecutionState = "QUEUED"
	LuauExecutionStateProcessing  LuauExecutionState = "PROCESSING"
	LuauExecutionStateCancelled   LuauExecutionState = "CANCELLED"
	LuauExecutionStateComplete    LuauExecutionState = "COMPLETE"
	LuauExecutionStateFailed      LuauExecutionState = "FAILED"
)
```
### `LuauExecutionErrorCode`
```go
type LuauExecutionErrorCode string

const (
	LuauExecutionErrorCodeUnspecified             LuauExecutionErrorCode = "ERROR_CODE_UNSPECIFIED"
	LuauExecutionErrorCodeScriptError             LuauExecutionErrorCode = "SCRIPT_ERROR"
	LuauExecutionErrorCodeDeadlineExceeded        LuauExecutionErrorCode = "DEADLINE_EXCEEDED"
	LuauExecutionErrorCodeOutputSizeLimitExceeded LuauExecutionErrorCode = "OUTPUT_SIZE_LIMIT_EXCEEDED"
	LuauExecutionErrorCodeInternalError           LuauExecutionErrorCode = "INTERNAL_ERROR"
)
```
### `StructuredMessageType`
```go
type StructuredMessageType string

const (
	StructuredMessageTypeUnspecified StructuredMessageType = "MESSAGE_TYPE_UNSPECIFIED"
	StructuredMessageTypeOutput      StructuredMessageType = "OUTPUT"
	StructuredMessageTypeInfo        StructuredMessageType = "INFO"
	StructuredMessageTypeWarning     StructuredMessageType = "WARNING"
	StructuredMessageTypeError       StructuredMessageType = "ERROR"
)
```

## Structures
### `LuauExecutionTaskError`
```go
type LuauExecutionTaskError struct {
	Code    LuauExecutionErrorCode `json:"code"`
	Message string                 `json:"message"`
}
```
### `LuauExecutionTaskOutput`
```go
type LuauExecutionTaskOutput struct {
	Results []any `json:"results"`
}
```
### `LuauExecutionTask`
```go
type LuauExecutionTask struct {
	Path                string                   `json:"path"`
	CreateTime          string                   `json:"createTime"`
	UpdateTime          string                   `json:"updateTime"`
	User                string                   `json:"user"`
	State               LuauExecutionState       `json:"state"`
	Script              string                   `json:"script"`
	Timeout             string                   `json:"timeout"`
	Error               *LuauExecutionTaskError  `json:"error"`
	Output              *LuauExecutionTaskOutput `json:"output"`
	BinaryInput         string                   `json:"binaryInput"`
	EnabledBinaryOutput bool                     `json:"enabledBinaryOutput"`
	BinaryOutputURI     string                   `json:"binaryOutputUri"`
}

// TaskInfo will return information from the URL path for the task.
// This is useful for GetLuauExecutionSessionTask when polling the method.
//
// universeId, placeId, and taskId will always be present.
// versionId and sessionId may or may not be nil.
func (t *LuauExecutionTask) TaskInfo() (universeId string, placeId string, versionId *string, sessionId *string, taskId string)
```
### `LuauExecutionTaskCreate`
```go
type LuauExecutionTaskCreate struct {
	Script              string `json:"script,omitempty"`
	Timeout             string `json:"timeout,omitempty"`
	BinaryInput         string `json:"binaryInput,omitempty"`
	EnabledBinaryOutput bool   `json:"enabledBinaryOutput,omitempty"`
}
```
### `LuauExecutionTaskLogStructuredMessage`
```go
type LuauExecutionTaskLogStructuredMessage struct {
	Message     string                `json:"message"`
	CreateTime  string                `json:"createTime"`
	MessageType StructuredMessageType `json:"messageType"`
}
```
### `LuauExecutionTaskLog`
```go
type LuauExecutionTaskLog struct {
	Path    string   `json:"path"`
	Mesages []string `json:"messages"`
}
```
### `LuauExecutionTaskLogs`
```go
type LuauExecutionTaskLogs struct {
	LuauExecutionSessionTaskLogs []LuauExecutionTaskLog `json:"luauExecutionSessionTaskLogs"`
	NextPageToken                string                 `json:"nextPageToken"`
}
```