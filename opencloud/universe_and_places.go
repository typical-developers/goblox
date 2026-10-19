package opencloud

import "context"

// UniverseAndPlacesService handles communication with universe and place endpoints.
//
// Deprecated: Use [UniversesService] and [PlacesService] instead.
type UniverseAndPlacesService service

// GetInstance retrieves an instance.
//
// Deprecated: Use [PlacesService.GetInstance] instead.
func (s *UniverseAndPlacesService) GetInstance(ctx context.Context, universeId, placeId, instanceId string) (*Instance, *Response, error) {
	return s.client.Places.GetInstance(ctx, universeId, placeId, instanceId)
}

// UpdateInstance updates an instance's property data.
//
// Deprecated: Use [PlacesService.UpdateInstance] instead.
func (s *UniverseAndPlacesService) UpdateInstance() error {
	return s.client.Places.UpdateInstance()
}

// ListInstanceChildren lists an instance's children.
//
// Deprecated: Use [PlacesService.ListInstanceChildren] instead.
func (s *UniverseAndPlacesService) ListInstanceChildren(ctx context.Context, universeId, placeId, instanceId string, opts *Options) (*InstanceChildren, *Response, error) {
	return s.client.Places.ListInstanceChildren(ctx, universeId, placeId, instanceId, opts)
}

// GetPlace gets information on a specified place.
//
// Deprecated: Use [PlacesService.GetPlace] instead.
func (s *UniverseAndPlacesService) GetPlace(ctx context.Context, universeId, placeId string) (*Place, *Response, error) {
	return s.client.Places.GetPlace(ctx, universeId, placeId)
}

// UpdatePlace updates information for a specified place.
//
// Deprecated: Use [PlacesService.UpdatePlace] instead.
func (s *UniverseAndPlacesService) UpdatePlace(ctx context.Context, universeId, placeId string, data PlaceUpdate, opts *PlaceUpdateOptions) (*Place, *Response, error) {
	return s.client.Places.UpdatePlace(ctx, universeId, placeId, data, opts)
}

// GetUniverse fetches information on a specified universe.
//
// Deprecated: Use [UniversesService.GetUniverse] instead.
func (s *UniverseAndPlacesService) GetUniverse(ctx context.Context, universeId string) (*Universe, *Response, error) {
	return s.client.Universes.GetUniverse(ctx, universeId)
}

// UpdateUniverse updates information for a specified universe.
//
// Deprecated: Use [UniversesService.UpdateUniverse] instead.
func (s *UniverseAndPlacesService) UpdateUniverse(ctx context.Context, universeId string, data UniverseUpdate) (*Universe, *Response, error) {
	return s.client.Universes.UpdateUniverse(ctx, universeId, data)
}

// PublishUniverseMessage publishes a message to all of the universe's live servers.
//
// Deprecated: Use [UniversesService.PublishUniverseMessage] instead.
func (s *UniverseAndPlacesService) PublishUniverseMessage(ctx context.Context, universeId string, data UniverseMessage) (*Response, error) {
	return s.client.Universes.PublishUniverseMessage(ctx, universeId, data)
}

// RestartUniverseServers restarts all active servers for a universe if a new version has been published.
//
// Deprecated: Use [UniversesService.RestartUniverseServers] instead.
func (s *UniverseAndPlacesService) RestartUniverseServers(ctx context.Context, universeId string) (*Response, error) {
	return s.client.Universes.RestartUniverseServers(ctx, universeId)
}
