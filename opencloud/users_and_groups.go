package opencloud

import (
	"fmt"
	"net/http"

	http_client "github.com/typical-developers/goblox/internal/http_client"
)

// https://create.roblox.com/docs/en-us/cloud/reference/User#Get-User
func GetUser(userId string) (result *User, err error) {
	path := fmt.Sprintf("/cloud/v2/users/%s", userId)
	res, err := HTTP.Request(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[User](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/User#Generate-User-Thumbnail
func GenerateUserThumbnail(userId string, query *GenerateUserThumbnailQuery) (resuilt *UserThumbnail, err error) {
	path := fmt.Sprintf("/cloud/v2/users/%s:generateThumbnail", userId)
	res, err := HTTP.Request(http.MethodPost, path, nil, query.ConvertToStringMap())
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[UserThumbnail](res)
}
