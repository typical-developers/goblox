package opencloud

import (
	"fmt"
	"net/http"

	http_client "github.com/typical-developers/goblox/internal/http_client"
)

// https://create.roblox.com/docs/en-us/cloud/reference/Instance#Get-Instance
func (s *Session) GetInstance(universeId string, placeId string, instanceId string) (result *Instance, err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s/places/%s/instances/%s", universeId, placeId, instanceId)
	res, err := s.Client.Do(http.MethodGet, path, nil, nil)

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
func (s *Session) UpdateInstance() error {
	return fmt.Errorf("UpdateInstance: Method is not implemented. Need this method? Make a PR on the GitHub!\nhttps://github.com/typical-developers/goblox")
}

// https://create.roblox.com/docs/en-us/cloud/reference/Instance#List-Instance-Children
func (s *Session) ListInstanceChildren(universeId string, placeId string, instanceId string, query *ListInstanceQuery) (result *ListInstance, err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s/places/%s/instances/%s:listChildren", universeId, placeId, instanceId)
	res, err := s.Client.Do(http.MethodGet, path, nil, query.ConvertToStringMap())

	if err != nil {
		return
	}

	return http_client.DecodeResult[ListInstance](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/Place#Get-Place
func (s *Session) GetPlace(universeId string, placeId string) (result *Place, err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s/places/%s", universeId, placeId)
	res, err := s.Client.Do(http.MethodGet, path, nil, nil)

	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[Place](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/Place#Update-Place
func (s *Session) UpdatePlace(universeId string, placeId string, data PlaceUpdate) (result *Place, err error) {
	if err := data.validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/cloud/v2/universes/%s/places/%s", universeId, placeId)
	res, err := s.Client.Do(http.MethodPatch, path, data, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[Place](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/Universe#Get-Universe
func (s *Session) GetUniverse(universeId string) (result *Universe, err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s", universeId)
	res, err := s.Client.Do(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[Universe](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/Universe#Update-Universe
func (s *Session) UpdateUniverse(universeId string, data UniverseUpdate) (result *Universe, err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s", universeId)
	res, err := s.Client.Do(http.MethodPatch, path, data, nil)
	if err != nil {
		return nil, err
	}

	return http_client.DecodeResult[Universe](res)
}

// https://create.roblox.com/docs/en-us/cloud/reference/Universe#Publish-Universe-Message
func (s *Session) PublishUniverseMessage(universeId string, data UniverseMessage) (err error) {
	if err := data.validate(); err != nil {
		return err
	}

	path := fmt.Sprintf("/cloud/v2/universes/%s:publishMessage", universeId)
	_, err = s.Client.Do(http.MethodPost, path, data, nil)
	if err != nil {
		return err
	}

	return nil
}

// https://create.roblox.com/docs/en-us/cloud/reference/Universe#Restart-Universe-Servers
func (s *Session) RestartUniverseServers(universeId string) (err error) {
	path := fmt.Sprintf("/cloud/v2/universes/%s:restartServers", universeId)
	_, err = s.Client.Do(http.MethodPost, path, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
