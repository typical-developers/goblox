package opencloud

import http_client "github.com/typical-developers/goblox/internal/http_client"

var Client *http_client.HTTPClient

func init() {
	Client = http_client.NewHTTPClient("apis.roblox.com")
}

func SetAPIToken(apiToken string) {
	Client.SetHeader("x-api-key", apiToken)
}
