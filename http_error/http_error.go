package http_error

import (
	"fmt"
	"net/http"
)

type HTTPResponseError struct {
	StatusCode int
	StatusText string
	Body       *[]byte
}

func (e HTTPResponseError) Error() string {
	return fmt.Sprintf("HTTPResponseError: %s", e.StatusText)
}

func NewHTTPResponseError(res *http.Response, body []byte) error {
	statusCode := res.StatusCode
	statusText := res.Status

	return HTTPResponseError{
		StatusCode: statusCode,
		StatusText: statusText,
		Body:       &body,
	}
}
