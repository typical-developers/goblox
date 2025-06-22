package opencloud

// --- Response Structures

type OperationResponse map[string]any

type OperationMetadata map[string]any

type Operation struct {
	Path     string            `json:"path"`
	Done     bool              `json:"done"`
	Response OperationResponse `json:"response"`
	Metadata OperationMetadata `json:"metadata"`
}

// --- Query Options

type Options struct {
	MaxPageSize int    `url:"maxPageSize,omitempty"`
	PageToken   string `url:"pageToken,omitempty"`
}

type OptionsWithFilter struct {
	Options
	Filter string `url:"filter,omitempty"`
}
