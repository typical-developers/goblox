package opencloud

import (
	"context"
	"fmt"
	"net/http"
)

// PlacesService handles communication with place and instance endpoints.
type PlacesService service

type InstanceRunContext string

const (
	InstanceRunContextLegacy InstanceRunContext = "Legacy"
	InstanceRunContextServer InstanceRunContext = "Server"
	InstanceRunContextClient InstanceRunContext = "Client"
	InstanceRunContextPlugin InstanceRunContext = "Plugin"
)

type InstanceFolder any

type InstanceLocalScript struct {
	Enabled    bool               `json:"Enabled"`
	RunContext InstanceRunContext `json:"RunContext"`
	Source     string             `json:"Source"`
}

type InstanceModuleScript struct {
	Source string `json:"Source"`
}

type InstanceScript struct {
	Enabled    bool               `json:"Enabled"`
	RunContext InstanceRunContext `json:"RunContext"`
	Source     string             `json:"Source"`
}

type InstanceDetails struct {
	Folder       *InstanceFolder       `json:"Folder"`
	LocalScript  *InstanceLocalScript  `json:"LocalScript"`
	ModuleScript *InstanceModuleScript `json:"ModuleScript"`
	Script       *InstanceScript       `json:"Script"`
}

type EngineInstance struct {
	Id     string `json:"Id"`
	Parent string `json:"Parent"`
	Name   string `json:"Name"`
}

type Instance struct {
	Path           string         `json:"path"`
	HasChildren    bool           `json:"hasChildren"`
	EngineInstance EngineInstance `json:"engineInstance"`
}

// GetInstance will fetch an instance's property data.
//
// Required scopes: universe.place.instance:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Instance#Cloud_GetInstance
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}/instances/{instance_id}
func (s *PlacesService) GetInstance(ctx context.Context, universeId, placeId, instanceId string) (*Instance, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s/places/%s/instances/%s", universeId, placeId, instanceId)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	instance := new(Instance)
	resp, err := s.client.Do(ctx, req, instance)
	if err != nil {
		return nil, resp, err
	}

	return instance, resp, nil
}

// UpdateInstance will update an instance's property data.
//
// Required scopes: universe.place.instance:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Instance#Cloud_UpdateInstance
//
// [PATCH] /cloud/v2/universes/{universe_id}/places/{place_id}/instances/{instance_id}
func (s *PlacesService) UpdateInstance() error {
	// I'm not entirely sure the purpose of this endpoint, the docs do not help with how this endpoint is used.
	// The main concern is the data being pushed and what it is used for. I'm not sure how to actually go about implementing it.
	// - LuckFire

	return fmt.Errorf("UpdateInstance: Method is not implemented. Need this method? Make a PR on the GitHub!\nhttps://github.com/typical-developers/goblox")
}

type InstanceChildren struct {
	Instances     []Instance `json:"instances"`
	NextPageToken string     `json:"nextPageToken"`
}

// ListInstanceChildren will list and instance's children.
//
// Required scopes: universe.place.instance:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Instance#Cloud_ListInstanceChildren
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}/instances/{instance_id}:listChildren
func (s *PlacesService) ListInstanceChildren(ctx context.Context, universeId, placeId, instanceId string, opts *Options) (*InstanceChildren, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s/places/%s/instances/%s:listChildren", universeId, placeId, instanceId)

	u, err := addOpts(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	instanceChildren := new(InstanceChildren)
	resp, err := s.client.Do(ctx, req, instanceChildren)
	if err != nil {
		return nil, resp, err
	}

	return instanceChildren, resp, nil
}

type Place struct {
	Path        string `json:"path"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	ServerSize  int    `json:"serverSize"`
}

// GetPlace will get information on a specified place.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Place#Cloud_GetPlace
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}
func (s *PlacesService) GetPlace(ctx context.Context, universeId, placeId string) (*Place, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s/places/%s", universeId, placeId)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	place := new(Place)
	resp, err := s.client.Do(ctx, req, place)
	if err != nil {
		return nil, resp, err
	}

	return place, resp, nil
}

type PlaceUpdate struct {
	DisplayName *string `json:"displayName,omitempty"`
	Description *string `json:"description,omitempty"`
	ServerSize  *int    `json:"serverSize,omitempty"`
}

type PlaceUpdateOptions struct {
	UpdateMask *string `url:"updateMask,omitempty"`
}

// UpdatePlace will update information for a specified place.
//
// Required scopes: universe.place:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Place#Cloud_UpdatePlace
//
// [PATCH] /cloud/v2/universes/{universe_id}/places/{place_id}
func (s *PlacesService) UpdatePlace(ctx context.Context, universeId, placeId string, data PlaceUpdate, opts *PlaceUpdateOptions) (*Place, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/universes/%s/places/%s", universeId, placeId)

	u, err := addOpts(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPatch, u, data)
	if err != nil {
		return nil, nil, err
	}

	place := new(Place)
	resp, err := s.client.Do(ctx, req, place)
	if err != nil {
		return nil, resp, err
	}

	return place, resp, nil
}
