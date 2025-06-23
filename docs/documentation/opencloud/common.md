# Common
These are some commonly used structures for the Opencloud API.

## Query Parameters
### `Options`
```go
type Options struct {
	MaxPageSize int    `url:"maxPageSize,omitempty"`
	PageToken   string `url:"pageToken,omitempty"`
}
```
### `OptionsWithFilter`
```go
type OptionsWithFilter struct {
	Options
	Filter string `url:"filter,omitempty"`
}
```

## Responses
### `Operation`
```go
type OperationResponse map[string]any

type OperationMetadata map[string]any

type Operation struct {
	Path     string            `json:"path"`
	Done     bool              `json:"done"`
	Response OperationResponse `json:"response"`
	Metadata OperationMetadata `json:"metadata"`
}
```