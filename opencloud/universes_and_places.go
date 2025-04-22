package opencloud

import (
	"fmt"
	"net/http"

	http_client "github.com/typical-developers/goblox/internal/http_client"
)

// https://create.roblox.com/docs/en-us/cloud/reference/Instance#Get-Instance
func GetInstance(universeId string, placeId string, instanceId string) (result *Instance, err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s/places/%s/instances/%s", universeId, placeId, instanceId)
	res, err := HTTP.Request(http.MethodGet, path, nil, nil)

	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[Instance](res)
}

// TODO: Implement this method.
// I'm not entirely sure the purpose of this endpoint, so I'm not sure how to actually go about implementing it.
// - LuckFire
//
// https://create.roblox.com/docs/en-us/cloud/reference/Instance#Update-Instance
func UpdateInstance() error {
	return fmt.Errorf("UpdateInstance: Method is not implemented. Need this method? Make a PR on the GitHub!\nhttps://github.com/typical-developers/goblox")
}

// https://create.roblox.com/docs/en-us/cloud/reference/Instance#List-Instance-Children
func ListInstanceChildren(universeId string, placeId string, instanceId string, query *ListInstanceQuery) (result *ListInstance, err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s/places/%s/instances/%s:listChildren", universeId, placeId, instanceId)
	res, err := HTTP.Request(http.MethodGet, path, nil, query.ConvertToStringMap())

	if err != nil {
		return
	}

	return http_client.DecodeResult[ListInstance](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/Place#Get-Place
func GetPlace(universeId string, placeId string) (result *Place, err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s/places/%s", universeId, placeId)
	res, err := HTTP.Request(http.MethodGet, path, nil, nil)

	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[Place](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/Place#Update-Place
func UpdatePlace(universeId string, placeId string, data PlaceUpdate) (result *Place, err error) {
	err = data.Validate()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/cloud/v2/universes/%s/places/%s", universeId, placeId)
	res, err := HTTP.Request(http.MethodPatch, path, data, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[Place](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/Universe#Get-Universe
func GetUniverse(universeId string) (result *Universe, err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s", universeId)
	res, err := HTTP.Request(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[Universe](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/Universe#Update-Universe
func UpdateUniverse(universeId string, data UniverseUpdate) (result *Universe, err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s", universeId)
	res, err := HTTP.Request(http.MethodPatch, path, data, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[Universe](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/Universe#Publish-Universe-Message
func PublishUniverseMessage(universeId string, data UniverseMessage) (err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s:publishMessage", universeId)
	_, err = HTTP.Request(http.MethodPost, path, data, nil)
	if err != nil {
		return err
	}

	return nil
}

// https://create.roblox.com/docs/en-us/cloud/reference/Universe#Restart-Universe-Servers
func RestartUniverseServers(universeId string) (err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s:restartServers", universeId)
	_, err = HTTP.Request(http.MethodPost, path, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
