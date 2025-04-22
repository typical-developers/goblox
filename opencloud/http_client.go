package opencloud

import http_client "github.com/typical-developers/goblox/internal/http_client"

var HTTP *http_client.HTTPClient

func init() {
	HTTP = http_client.NewHTTPClient("https://apis.roblox.com/cloud")
}

func SetAPIToken(apiToken string) {
	HTTP.SetHeader("x-api-key", apiToken)
}
