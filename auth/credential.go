package auth

import "net/http"

type Kind int

const (
	KindNone Kind = iota
	KindAPIKey
	KindOAuth2
	KindCookie
)

type Credential interface {
	Kind() Kind
	Apply(req *http.Request) error
}

type ResponseHandler interface {
	HandleResponse(resp *http.Response) (retry bool)
}
