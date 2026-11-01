package auth

import (
	"context"
	"net/http"
)

type TokenSource interface {
	Token(ctx context.Context) (string, error)
}

type OAuth struct {
	Source TokenSource
}

type OAuthFunc func(ctx context.Context) (string, error)

func (f OAuthFunc) Token(ctx context.Context) (string, error) {
	return f(ctx)
}

type Token string

func (t Token) Token(context.Context) (string, error) {
	return string(t), nil
}

func OAuthToken(token string) OAuth {
	return OAuth{Token(token)}
}

func (c OAuth) Kind() Kind {
	return KindOAuth2
}

func (c OAuth) Apply(req *http.Request) error {
	token, err := c.Source.Token(req.Context())
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}
