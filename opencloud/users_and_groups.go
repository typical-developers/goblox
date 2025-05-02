package opencloud

import (
	"fmt"
	"net/http"

	http_client "github.com/typical-developers/goblox/internal/http_client"
)

// https://create.roblox.com/docs/en-us/cloud/reference/User#Get-User
func (s *Session) GetUser(userId string) (result *User, err error) {
	path := fmt.Sprintf("/cloud/v2/users/%s", userId)
	res, err := s.Client.Do(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[User](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/User#Generate-User-Thumbnail
func (s *Session) GenerateUserThumbnail(userId string, query *GenerateUserThumbnailQuery) (resuilt *UserThumbnail, err error) {
	path := fmt.Sprintf("/cloud/v2/users/%s:generateThumbnail", userId)
	res, err := s.Client.Do(http.MethodPost, path, nil, query.ConvertToStringMap())
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[UserThumbnail](res)
}

// https://create.roblox.com/docs/cloud/reference/UserRestriction#List-User-Restrictions
func (s *Session) ListUserRestrictions(universeId string, placeId *string) (result *UserRestrictionsList, err error) {
	var path string
	if placeId != nil {
		path = fmt.Sprintf("/cloud/v2/universes/%s/places/%s/user-restrictions", universeId, *placeId)
	} else {
		path = fmt.Sprintf("/cloud/v2/universes/%s/user-restrictions", universeId)
	}

	res, err := s.Client.Do(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[UserRestrictionsList](res)
}

// https://create.roblox.com/docs/cloud/reference/UserRestriction#Get-User-Restriction
func (s *Session) GetUserRestriction(universeId string, placeId *string, userId string) (result *UserRestriction, err error) {
	var path string

	if placeId != nil {
		path = fmt.Sprintf("/cloud/v2/universes/%s/places/%s/user-restrictions/%s", universeId, *placeId, userId)
	} else {
		path = fmt.Sprintf("/cloud/v2/universes/%s/user-restrictions/%s", universeId, userId)
	}

	res, err := s.Client.Do(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[UserRestriction](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/UserRestriction#Update-User-Restriction
func (s *Session) UpdateUserRestriction(universeId string, placeId *string, userId string, data UserRestrictionUpdate) (result *UserRestriction, err error) {
	if err := data.validateUpdate(); err != nil {
		return nil, err
	}

	var path string

	if placeId != nil {
		path = fmt.Sprintf("/cloud/v2/universes/%s/places/%s/user-restrictions/%s", universeId, *placeId, userId)
	} else {
		path = fmt.Sprintf("/cloud/v2/universes/%s/user-restrictions/%s", universeId, userId)
	}

	res, err := s.Client.Do(http.MethodPatch, path, data, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[UserRestriction](res)
}
