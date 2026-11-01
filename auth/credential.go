package auth

import "net/http"

type Kind int

func (k Kind) String() string {
	switch k {
	case KindNone:
		return "none"
	case KindAPIKey:
		return "API key"
	case KindOAuth2:
		return "OAuth2"
	case KindCookie:
		return "cookie"
	default:
		return "unknown"
	}
}

const (
	KindNone Kind = iota
	KindAPIKey
	KindOAuth2
	KindCookie
)

type Kinds uint8

func KindsOf(ks ...Kind) Kinds {
	var s Kinds
	for _, k := range ks {
		s |= 1 << k
	}

	return s
}

func (s Kinds) Has(k Kind) bool {
	return s&(1<<k) != 0
}

func (s Kinds) List() []Kind {
	var out []Kind
	for k := KindNone; k <= KindCookie; k++ {
		if s.Has(k) {
			out = append(out, k)
		}
	}

	return out
}

type Credential interface {
	Kind() Kind
	Apply(req *http.Request) error
}

type ResponseHandler interface {
	HandleResponse(resp *http.Response) (retry bool)
}
