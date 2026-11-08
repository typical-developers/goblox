package legacy

import (
	"context"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	GetUserPresences = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/presence/users",
		auth.KindCookie, auth.KindNone,
	).WithDomain("presence")
)

type PresenceService service

type UserPresenceType int

const (
	UserPresenceTypeOffline   UserPresenceType = 0
	UserPresenceTypeOnline    UserPresenceType = 1
	UserPresenceTypeInGame    UserPresenceType = 2
	UserPresenceTypeInStudio  UserPresenceType = 3
	UserPresenceTypeInvisible UserPresenceType = 4
)

// UserPresence is a user's presence. Location fields are only filled in when the
// requesting user is allowed to see them, which usually needs a cookie.
type UserPresence struct {
	UserPresenceType UserPresenceType `json:"userPresenceType"`
	LastLocation     string           `json:"lastLocation"`
	PlaceID          *int64           `json:"placeId"`
	RootPlaceID      *int64           `json:"rootPlaceId"`
	GameID           *string          `json:"gameId"`
	UniverseID       *int64           `json:"universeId"`
	UserID           int64            `json:"userId"`
}

type userPresenceRequest struct {
	UserIDs []int64 `json:"userIds"`
}

// GetUserPresences will get the presence of a list of users.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/presence#presence_post_v1_presence_users
//
// [POST] https://presence.roblox.com/v1/presence/users
func (s *PresenceService) GetUserPresences(ctx context.Context, userIds []int64, reqOpts ...goblox.RequestOpt) ([]UserPresence, *Response, error) {
	var result struct {
		UserPresences []UserPresence `json:"userPresences"`
	}
	resp, err := s.client.Call(ctx, GetUserPresences.Compile(nil), userPresenceRequest{UserIDs: userIds}, &result, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return result.UserPresences, resp, nil
}
