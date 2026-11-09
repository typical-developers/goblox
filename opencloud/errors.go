package opencloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/luckfire/goblox"
)

type GatewayError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	// Roblox's API errors are very convoluted, there are multiple
	// different structure types that could get returned.

	goblox.ResponseError

	Message     string
	Details     []json.RawMessage
	FieldErrors map[string][]string
	Gateway     []GatewayError
	TraceID     string
}

func (e *Error) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "goblox: HTTP %d", e.StatusCode)

	switch {
	case e.Message != "":
		fmt.Fprintf(&b, ": %s", e.Message)
	case len(e.Gateway) > 0:
		fmt.Fprintf(&b, ": %s", e.Gateway[0].Message)
	}

	for field, msgs := range e.FieldErrors {
		fmt.Fprintf(&b, "; %s: %s", field, strings.Join(msgs, ", "))
	}

	return b.String()
}

func decodeError(resp *http.Response, body []byte) error {
	e := &Error{ResponseError: goblox.NewResponseError(resp, body)}

	var raw struct {
		Code         string            `json:"code"`
		Error        string            `json:"error"`
		Message      string            `json:"message"`
		Title        string            `json:"title"`
		Details      []json.RawMessage `json:"details"`
		ErrorDetails []json.RawMessage `json:"errorDetails"`
		Errors       json.RawMessage   `json:"errors"` // array (gateway) or object (validation)
		Extensions   struct {
			TraceID string `json:"traceId"`
		} `json:"extensions"`
	}
	if json.Unmarshal(body, &raw) != nil {
		return e // not JSON (an HTML gateway page, an empty body): status only
	}

	e.Message = raw.Message
	if e.Message == "" {
		e.Message = raw.Title
	}
	e.Details = raw.Details
	if len(e.Details) == 0 {
		e.Details = raw.ErrorDetails
	}
	e.TraceID = raw.Extensions.TraceID

	switch t := bytes.TrimSpace(raw.Errors); {
	case len(t) > 0 && t[0] == '[':
		_ = json.Unmarshal(t, &e.Gateway)
	case len(t) > 0 && t[0] == '{':
		_ = json.Unmarshal(t, &e.FieldErrors)
	}

	return e
}
