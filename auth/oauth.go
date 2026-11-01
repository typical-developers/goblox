package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
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

type OAuthIntrospector struct {
	ClientID, ClientSecret string
	Source                 TokenSource
}

func IntrospectedOAuthToken(token, clientID, clientSecret string) *Introspected {
	o := OAuthToken(token)

	return &Introspected{
		Credential: o,
		Introspector: OAuthIntrospector{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Source:       o.Source,
		},
		TTL: 15 * time.Minute,
	}
}

type oAuthIntrospect struct {
	Active bool   `json:"active"`
	Scope  string `json:"scope"`
	Exp    int64  `json:"exp"`
	Sub    string `json:"sub"`
}

func (i OAuthIntrospector) Introspect(ctx context.Context) (*Grant, error) {
	if i.Source == nil {
		return nil, fmt.Errorf("goblox/auth: OAuthIntrospector has no token source")
	}

	token, err := i.Source.Token(ctx)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost, "https://apis.roblox.com/oauth/v1/token/introspect",
		strings.NewReader(url.Values{
			"token":         {token},
			"client_id":     {i.ClientID},
			"client_secret": {i.ClientSecret},
		}.Encode()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	auth := new(oAuthIntrospect)
	if err := introspect(req, auth); err != nil {
		return nil, err
	}

	grant := &Grant{
		Active:    auth.Active,
		Scopes:    strings.Fields(auth.Scope),
		OwnerID:   auth.Sub,
		CheckedAt: time.Now(),
	}
	if auth.Exp > 0 {
		grant.ExpiresAt = time.Unix(auth.Exp, 0)
	}

	return grant, nil
}
