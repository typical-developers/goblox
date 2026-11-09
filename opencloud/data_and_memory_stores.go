package opencloud

import (
	"context"
	"net/http"

	"github.com/luckfire/goblox"
	"github.com/luckfire/goblox/auth"
)

var (
	ListDataStores = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/data-stores",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.control:list")

	SnapshotDataStores = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/data-stores:snapshot",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.control:snapshot")

	FlushMemoryStore = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/memory-store:flush",
		auth.KindAPIKey,
	).WithScopes("memory-store:flush")

	CreateMemoryStoreQueueItem = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/memory-store/queues/{queueId}/items",
		auth.KindAPIKey,
	).WithScopes("memory-store.queue:add")

	DiscardMemoryStoreQueueItems = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/memory-store/queues/{queueId}/items:discard",
		auth.KindAPIKey,
	).WithScopes("memory-store.queue:discard")

	ReadMemoryStoreQueueItems = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/memory-store/queues/{queueId}/items:read",
		auth.KindAPIKey,
	).WithScopes("memory-store.queue:dequeue")

	ListMemoryStoreSortedMapItems = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/memory-store/sorted-maps/{sortedMapId}/items",
		auth.KindAPIKey,
	).WithScopes("memory-store.sorted-map:read")

	CreateMemoryStoreSortedMapItem = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/memory-store/sorted-maps/{sortedMapId}/items",
		auth.KindAPIKey,
	).WithScopes("memory-store.sorted-map:write")

	GetMemoryStoreSortedMapItem = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/memory-store/sorted-maps/{sortedMapId}/items/{itemId}",
		auth.KindAPIKey,
	).WithScopes("memory-store.sorted-map:read")

	DeleteMemoryStoreSortedMapItem = goblox.NewEndpoint(
		http.MethodDelete,
		"/cloud/v2/universes/{universeId}/memory-store/sorted-maps/{sortedMapId}/items/{itemId}",
		auth.KindAPIKey,
	).WithScopes("memory-store.sorted-map:write")

	UpdateMemoryStoreSortedMapItem = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/universes/{universeId}/memory-store/sorted-maps/{sortedMapId}/items/{itemId}",
		auth.KindAPIKey,
	).WithScopes("memory-store.sorted-map:write")

	ListOrderedDataStoreEntries = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/ordered-data-stores/{orderedDataStoreId}/scopes/{scopeId}/entries",
		auth.KindAPIKey,
	).WithScopes("universe.ordered-data-store.scope.entry:read")

	CreateOrderedDataStoreEntry = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/ordered-data-stores/{orderedDataStoreId}/scopes/{scopeId}/entries",
		auth.KindAPIKey,
	).WithScopes("universe.ordered-data-store.scope.entry:write")

	GetOrderedDataStoreEntry = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/ordered-data-stores/{orderedDataStoreId}/scopes/{scopeId}/entries/{entryId}",
		auth.KindAPIKey,
	).WithScopes("universe.ordered-data-store.scope.entry:read")

	DeleteOrderedDataStoreEntry = goblox.NewEndpoint(
		http.MethodDelete,
		"/cloud/v2/universes/{universeId}/ordered-data-stores/{orderedDataStoreId}/scopes/{scopeId}/entries/{entryId}",
		auth.KindAPIKey,
	).WithScopes("universe.ordered-data-store.scope.entry:write")

	UpdateOrderedDataStoreEntry = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/universes/{universeId}/ordered-data-stores/{orderedDataStoreId}/scopes/{scopeId}/entries/{entryId}",
		auth.KindAPIKey,
	).WithScopes("universe.ordered-data-store.scope.entry:write")

	IncrementOrderedDataStoreEntry = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/ordered-data-stores/{orderedDataStoreId}/scopes/{scopeId}/entries/{entryId}:increment",
		auth.KindAPIKey,
	).WithScopes("universe.ordered-data-store.scope.entry:write")

	ListDataStoreEntries = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/entries",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.objects:list")

	ListDataStoreEntriesScoped = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/scopes/{scope}/entries",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.objects:list")

	CreateDataStoreEntry = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/entries",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.objects:create")

	CreateDataStoreEntryScoped = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/scopes/{scope}/entries",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.objects:create")

	GetDataStoreEntry = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/entries/{entryId}",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.objects:read")

	GetDataStoreEntryScoped = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/scopes/{scope}/entries/{entryId}",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.objects:read")

	DeleteDataStoreEntry = goblox.NewEndpoint(
		http.MethodDelete,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/entries/{entryId}",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.objects:delete")

	DeleteDataStoreEntryScoped = goblox.NewEndpoint(
		http.MethodDelete,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/scopes/{scope}/entries/{entryId}",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.objects:delete")

	UpdateDataStoreEntry = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/entries/{entryId}",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.objects:update")

	UpdateDataStoreEntryScoped = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/scopes/{scope}/entries/{entryId}",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.objects:update")

	IncrementDataStoreEntry = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/entries/{entryId}:increment",
		auth.KindAPIKey,
	).WithAnyScopes("universe-datastores.objects:create", "universe-datastores.objects:update")

	IncrementDataStoreEntryScoped = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/scopes/{scope}/entries/{entryId}:increment",
		auth.KindAPIKey,
	).WithAnyScopes("universe-datastores.objects:create", "universe-datastores.objects:update")

	ListDataStoreEntryRevisions = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/entries/{entryId}:listRevisions",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.versions:list")

	ListDataStoreEntryRevisionsScoped = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}/scopes/{scope}/entries/{entryId}:listRevisions",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.versions:list")
)

// DataAndMemoryStoreService will handle communciation with the actions related to the API.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type DataAndMemoryStoreService service

type DataStore struct {
	Path       string `json:"path"`
	CreateTime string `json:"createTime"`
	ID         string `json:"id"`
}

type DataStoreList struct {
	DataStores    []DataStore `json:"dataStores"`
	NextPageToken string      `json:"nextPageToken"`
}

// ListDataStores will fetch a list of data stores for a specific universe.
//
// Required scopes: universe-datastores.control:list
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/DataStore#Cloud_ListDataStores
//
// [GET] /cloud/v2/universes/{universe_id}/data-stores
func (s *DataAndMemoryStoreService) ListDataStores(ctx context.Context, universeId string, opts *OptionsWithFilter, reqOpts ...goblox.RequestOpt) (*DataStoreList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	dataStoreList := new(DataStoreList)
	resp, err := s.client.Call(ctx, ListDataStores.Compile(q, universeId), nil, dataStoreList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dataStoreList, resp, nil
}

type DataStoreSnapshot struct {
	NewSnapshotTaken   bool   `json:"newSnapshotTaken"`
	LatestSnapshotTime string `json:"latestSnapshotTime"`
}

// SnapshotDataStores will take a snapshot of all data stores for a specific universe.
// After the snapshot is taken, the next write to every key in the experience creaets a versioned backup of data.
//
// Required scopes:
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/DataStore#Cloud_SnapshotDataStores
//
// [POST] /cloud/v2/universes/{universe_id}/data-stores:snapshot
func (s *DataAndMemoryStoreService) SnapshotDataStores(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*DataStoreSnapshot, *Response, error) {
	dataStoreSnapshot := new(DataStoreSnapshot)
	resp, err := s.client.Call(ctx, SnapshotDataStores.Compile(nil, universeId), nil, dataStoreSnapshot, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dataStoreSnapshot, resp, nil
}

type DataStoreEntryState string

const (
	DataStoreEntryStateUnspecified DataStoreEntryState = "STATE_UNSPECIFIED"
	DataStoreEntryStateActive      DataStoreEntryState = "ACTIVE"
	DataStoreEntryStateDeleted     DataStoreEntryState = "DELETED"
)

type DataStoreEntry struct {
	Path                 string              `json:"path"`
	CreateTime           string              `json:"createTime"`
	RevisionID           string              `json:"revisionId"`
	RevisionCreationTime string              `json:"revisionCreationTime"`
	State                DataStoreEntryState `json:"state"`
	Etag                 string              `json:"etag"`
	Value                any                 `json:"value"`
	ID                   string              `json:"id"`
	User                 string              `json:"user"`
}

type DataStoreEntriesList struct {
	DataStoreEntries []DataStoreEntry `json:"dataStoreEntries"`
	NextPageToken    string           `json:"nextPageToken"`
}

type ListDataStoreEntriesOptions struct {
	MaxPageSize *int    `url:"maxPageSize,omitempty"`
	PageToken   *string `url:"pageToken,omitempty"`
	Filter      *string `url:"filter,omitempty"`
	ShowDeleted *bool   `url:"showDeleted,omitempty"`
}

// ListDataStoreEntries will fetch a list of entries for a specific data store under a specific universe.
//
// Required scopes: universe-datastores.objects:list
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/DataStoreEntry#Cloud_ListDataStoreEntries
//
// [GET] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/entries
//
// [GET] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/scopes/{scope_id}/entries
func (s *DataAndMemoryStoreService) ListDataStoreEntries(ctx context.Context, universeId, dataStoreId string, scope *string, opts *ListDataStoreEntriesOptions, reqOpts ...goblox.RequestOpt) (*DataStoreEntriesList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	e := ListDataStoreEntries.Compile(q, universeId, dataStoreId)
	if scope != nil {
		e = ListDataStoreEntriesScoped.Compile(q, universeId, dataStoreId, *scope)
	}

	dataStoreEntriesList := new(DataStoreEntriesList)
	resp, err := s.client.Call(ctx, e, nil, dataStoreEntriesList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dataStoreEntriesList, resp, nil
}

type DataStoreEntryCreate struct {
	Etag       *string         `json:"etag,omitempty"`
	Value      *any            `json:"value,omitempty"`
	Users      *[]string       `json:"users,omitempty"`
	Attributes *map[string]any `json:"attributes,omitempty"`
}

type DataStoreEntryCreateOptions struct {
	ID *string `url:"id,omitempty"`
}

// CreateDataStoreEntry will create a new entry for the specified datastore under a specific universe.
//
// Required scopes: universe-datastores.objects:create
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/DataStoreEntry#Cloud_CreateDataStoreEntry
//
// [POST] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/entries
//
// [POST] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/scopes/{scope_id}/entries
func (s *DataAndMemoryStoreService) CreateDataStoreEntry(ctx context.Context, universeId, dataStoreId string, scope *string, data DataStoreEntryCreate, opts *DataStoreEntryCreateOptions, reqOpts ...goblox.RequestOpt) (*DataStoreEntry, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	e := CreateDataStoreEntry.Compile(q, universeId, dataStoreId)
	if scope != nil {
		e = CreateDataStoreEntryScoped.Compile(q, universeId, dataStoreId, *scope)
	}

	dataStoreEntry := new(DataStoreEntry)
	resp, err := s.client.Call(ctx, e, data, dataStoreEntry, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dataStoreEntry, resp, nil
}

// GetDataStoreEntry will fetch a specified entry for the specified datastore under a specific universe.
//
// Required scopes: universe-datastores.objects:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/DataStoreEntry#Cloud_GetDataStoreEntry
//
// [GET] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/entries/{entry_id}
//
// [GET] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/scopes/{scope_id}/entries/{entry_id}
func (s *DataAndMemoryStoreService) GetDataStoreEntry(ctx context.Context, universeId, datastoreId string, scope *string, entryId string, reqOpts ...goblox.RequestOpt) (*DataStoreEntry, *Response, error) {
	e := GetDataStoreEntry.Compile(nil, universeId, datastoreId, entryId)
	if scope != nil {
		e = GetDataStoreEntryScoped.Compile(nil, universeId, datastoreId, *scope, entryId)
	}

	dataStoreEntry := new(DataStoreEntry)
	resp, err := s.client.Call(ctx, e, nil, dataStoreEntry, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dataStoreEntry, resp, nil
}

// DeleteDataStoreEntry will delete a specified entry for the specified datastore under a specific universe.
//
// Required scopes: universe-datastores.objects:delete
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/DataStoreEntry#Cloud_DeleteDataStoreEntry
//
// [DELETE] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/entries/{entry_id}
//
// [DELETE] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/scopes/{scope_id}/entries/{entry_id}
func (s *DataAndMemoryStoreService) DeleteDataStoreEntry(ctx context.Context, universeId, datastoreId string, scope *string, entryId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	e := DeleteDataStoreEntry.Compile(nil, universeId, datastoreId, entryId)
	if scope != nil {
		e = DeleteDataStoreEntryScoped.Compile(nil, universeId, datastoreId, *scope, entryId)
	}

	return s.client.Call(ctx, e, nil, nil, reqOpts...)
}

type DataStoreEntryUpdate struct {
	Etag       *string         `json:"etag,omitempty"`
	Value      *any            `json:"value,omitempty"`
	Users      *[]string       `json:"users,omitempty"`
	Attributes *map[string]any `json:"attributes,omitempty"`
}

type DataStoreEntryUpdateOpts struct {
	AllowMissing *bool `url:"allowMissing,omitempty"`
}

// UpdateDataStoreEntry will update a specified entry for the specified datastore under a specific universe.
//
// Required scopes: universe-datastores.objects:update
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/DataStoreEntry#Cloud_UpdateDataStoreEntry
//
// [PATCH] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/entries/{entry_id}
//
// [PATCH] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/scopes/{scope_id}/entries/{entry_id}
func (s *DataAndMemoryStoreService) UpdateDataStoreEntry(ctx context.Context, universeId, datastoreId string, scope *string, entryId string, data DataStoreEntryUpdate, opts *DataStoreEntryUpdateOpts, reqOpts ...goblox.RequestOpt) (*DataStoreEntry, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	e := UpdateDataStoreEntry.Compile(q, universeId, datastoreId, entryId)
	if scope != nil {
		e = UpdateDataStoreEntryScoped.Compile(q, universeId, datastoreId, *scope, entryId)
	}

	dataStoreEntry := new(DataStoreEntry)
	resp, err := s.client.Call(ctx, e, data, dataStoreEntry, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dataStoreEntry, resp, nil
}

type DataStoreEntryIncrement struct {
	Amount     *int            `json:"amount,omitempty"`
	Users      *[]string       `json:"users,omitempty"`
	Attributes *map[string]any `json:"attributes,omitempty"`
}

// IncrementDataStoreEntry will increment the value for a specified entry for the specified datastore under a specific universe.
//
// Required scopes:
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/DataStoreEntry#Cloud_IncrementDataStoreEntry
//
// [POST] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/entries/{entry_id}:increment
//
// [POST] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/scopes/{scope_id}/entries/{entry_id}:increment
func (s *DataAndMemoryStoreService) IncrementDataStoreEntry(ctx context.Context, universeId, datastoreId string, scope *string, entryId string, data DataStoreEntryIncrement, reqOpts ...goblox.RequestOpt) (*DataStoreEntry, *Response, error) {
	e := IncrementDataStoreEntry.Compile(nil, universeId, datastoreId, entryId)
	if scope != nil {
		e = IncrementDataStoreEntryScoped.Compile(nil, universeId, datastoreId, *scope, entryId)
	}

	dataStoreEntry := new(DataStoreEntry)
	resp, err := s.client.Call(ctx, e, data, dataStoreEntry, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dataStoreEntry, resp, nil
}

type DataStoryEntryRevisionsList struct {
	DataStoreEntries []DataStoreEntry `json:"dataStoreEntries"`
	NextPageToken    string           `json:"nextPageToken"`
}

// ListDataStoreEntryRevisions will fetch a list of revisions for a specific entry for the specified datastore under a specific universe.
//
// Required scopes: universe-datastores.versions:list
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/DataStoreEntry#Cloud_ListDataStoreEntryRevisions
//
// [GET] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/entries/{entry_id}:listRevisions
//
// [GET] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}/scopes/{scope_id}/entries/{entry_id}:listRevisions
func (s *DataAndMemoryStoreService) ListDataStoreEntryRevisions(ctx context.Context, universeId, datastoreId string, scope *string, entryId string, opts *Options, reqOpts ...goblox.RequestOpt) (*DataStoryEntryRevisionsList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	e := ListDataStoreEntryRevisions.Compile(q, universeId, datastoreId, entryId)
	if scope != nil {
		e = ListDataStoreEntryRevisionsScoped.Compile(q, universeId, datastoreId, *scope, entryId)
	}

	dataStoreEntryRevisionsList := new(DataStoryEntryRevisionsList)
	resp, err := s.client.Call(ctx, e, nil, dataStoreEntryRevisionsList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dataStoreEntryRevisionsList, resp, nil
}

// FlushMemoryStore will asynchronously flush the memory store for a specific universe.
//
// Required scopes: universe.memory-store:flush
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/MemoryStore#Cloud_FlushMemoryStore
//
// [POST] /cloud/v2/universes/{universe_id}/memory-store:flush
func (s *DataAndMemoryStoreService) FlushMemoryStore(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*Operation, *Response, error) {
	operation := new(Operation)
	resp, err := s.client.Call(ctx, FlushMemoryStore.Compile(nil, universeId), nil, operation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return operation, resp, nil
}

type MemoryStoreQueueItem struct {
	Path       string `json:"path"`
	Data       any    `json:"data"`
	Priority   int    `json:"priority"`
	ExpireTime string `json:"expireTime"`
	ID         string `json:"id"`
}

type MemoryStoreQueueItemCreate struct {
	Data     *any    `json:"data,omitempty"`
	Priority *int    `json:"priority,omitempty"`
	TTL      *string `json:"ttl,omitempty"`
}

// CreateMemoryStoreQueueItem will create a new item in the memory store queue for a specific universe.
//
// Required scopes: universe.memory-store.queue:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/MemoryStoreQueueItem#Cloud_CreateMemoryStoreQueueItem
//
// [POST] /cloud/v2/universes/{universe_id}/memory-store/queues/{queue_id}/items
func (s *DataAndMemoryStoreService) CreateMemoryStoreQueueItem(ctx context.Context, universeId, queueId string, data MemoryStoreQueueItemCreate, reqOpts ...goblox.RequestOpt) (*MemoryStoreQueueItem, *Response, error) {
	memoryStoreQueueItem := new(MemoryStoreQueueItem)
	resp, err := s.client.Call(ctx, CreateMemoryStoreQueueItem.Compile(nil, universeId, queueId), data, memoryStoreQueueItem, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return memoryStoreQueueItem, resp, nil
}

type MemoryStoreQueueItemsDiscard struct {
	ReadID string `json:"readId"`
}

// DiscardMemoryStoreQueueItems will discard a specific item(s) in the memory store queue for a specific universe.
//
// Required scopes: universe.memory-store.queue:discard
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/MemoryStoreQueueItem#Cloud_DiscardMemoryStoreQueueItems
//
// [POST] /cloud/v2/universes/{universe_id}/memory-store/queues/{queue_id}/items:discard
func (s *DataAndMemoryStoreService) DiscardMemoryStoreQueueItems(ctx context.Context, universeId, queueId string, data MemoryStoreQueueItemsDiscard, reqOpts ...goblox.RequestOpt) (*Response, error) {
	resp, err := s.client.Call(ctx, DiscardMemoryStoreQueueItems.Compile(nil, universeId, queueId), data, nil, reqOpts...)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

type MemoryStoreQueueItems struct {
	ReadID string                 `json:"readId"`
	Items  []MemoryStoreQueueItem `json:"items"`
}

type MemoryStoreQueueItemsOptions struct {
	Count              *int    `url:"count,omitempty"`
	AllOrNothing       *bool   `url:"allOrNothing,omitempty"`
	InvisibilityWindow *string `url:"invisibilityWindow,omitempty"`
}

// ReadMemoryStoreQueueItems will fetch specific item(s) in the memory store queue for a specific universe.
//
// Required scopes: universe.memory-store.queue:dequeue
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/MemoryStoreQueueItem#Cloud_ReadMemoryStoreQueueItems
//
// [GET] /cloud/v2/universes/{universe_id}/memory-store/queues/{queue_id}/items:read
func (s *DataAndMemoryStoreService) ReadMemoryStoreQueueItems(ctx context.Context, universeId, queueId string, opts *MemoryStoreQueueItemsOptions, reqOpts ...goblox.RequestOpt) (*MemoryStoreQueueItems, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	memoryStoreQueueItems := new(MemoryStoreQueueItems)
	resp, err := s.client.Call(ctx, ReadMemoryStoreQueueItems.Compile(q, universeId, queueId), nil, memoryStoreQueueItems, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return memoryStoreQueueItems, resp, nil
}

type MemoryStoreSortedMapItem struct {
	Path           string  `json:"path"`
	Value          any     `json:"value"`
	Etag           string  `json:"etag"`
	ID             string  `json:"id"`
	StringSortKey  *string `json:"stringSortKey,omitempty"`
	NumericSortKey *int    `json:"numericSortKey,omitempty"`
}

type MemoryStoreSortedMapList struct {
	MemoryStoreSortedMapItems []MemoryStoreSortedMapItem `json:"memoryStoreSortedMapItems"`
	NextPageToken             string                     `json:"nextPageToken"`
}

type MemoryStoreSortedMapItemListOptions struct {
	MaxPageSize *int    `url:"maxPageSize,omitempty"`
	PageToken   *string `url:"pageToken,omitempty"`
	Filter      *string `url:"filter,omitempty"`
	OrderBy     *string `url:"orderBy,omitempty"`
}

// ListMemoryStoreSortedMapItems will fetch a list of memory store map items for a specific sorted map under a specific universe.
//
// Required scopes: universe.memory-store.sorted-map:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/MemoryStoreSortedMapItem#Cloud_ListMemoryStoreSortedMapItems
//
// [GET] /cloud/v2/universes/{universe_id}/memory-store/sorted-maps/{sorted_map_id}/items
func (s *DataAndMemoryStoreService) ListMemoryStoreSortedMapItems(ctx context.Context, universeId, sortedMapId string, opts *MemoryStoreSortedMapItemListOptions, reqOpts ...goblox.RequestOpt) (*MemoryStoreSortedMapList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	memoryStoreSortedMapList := new(MemoryStoreSortedMapList)
	resp, err := s.client.Call(ctx, ListMemoryStoreSortedMapItems.Compile(q, universeId, sortedMapId), nil, memoryStoreSortedMapList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return memoryStoreSortedMapList, resp, nil
}

type MemoryStoreSortedMapItemCreate struct {
	Value         *any    `json:"value,omitempty"`
	TTL           *string `json:"ttl,omitempty"`
	ID            *string `json:"id,omitempty"`
	StringSortKey *string `json:"stringSortKey,omitempty"`
}

type MemoryStoreSortedMapItemCreateOptions struct {
	ID string `url:"id,omitempty"`
}

// CreateMemoryStoreSortedMapItem will create a new item in the memory store sorted map for a specific universe.
//
// Required scopes: universe.memory-store.sorted-map:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/MemoryStoreSortedMapItem#Cloud_CreateMemoryStoreSortedMapItem
//
// [POST] /cloud/v2/universes/{universe_id}/memory-store/sorted-maps/{sorted_map_id}/items
func (s *DataAndMemoryStoreService) CreateMemoryStoreSortedMapItem(ctx context.Context, universeId, sortedMapId string, data MemoryStoreSortedMapItemCreate, opts *MemoryStoreSortedMapItemCreateOptions, reqOpts ...goblox.RequestOpt) (*MemoryStoreSortedMapItem, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	memoryStoreSortedMapItem := new(MemoryStoreSortedMapItem)
	resp, err := s.client.Call(ctx, CreateMemoryStoreSortedMapItem.Compile(q, universeId, sortedMapId), data, memoryStoreSortedMapItem, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return memoryStoreSortedMapItem, resp, nil
}

// GetMemoryStoreSortedMapItem will fetch a specific item in the memory store sorted map for a specific universe.
//
// Required scopes: universe.memory-store.sorted-map:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/MemoryStoreSortedMapItem#Cloud_GetMemoryStoreSortedMapItem
//
// [GET] /cloud/v2/universes/{universe_id}/memory-store/sorted-maps/{sorted_map_id}/items/{item_id}
func (s *DataAndMemoryStoreService) GetMemoryStoreSortedMapItem(ctx context.Context, universeId, sortedMapId, itemId string, reqOpts ...goblox.RequestOpt) (*MemoryStoreSortedMapItem, *Response, error) {
	memoryStoreSortedMapItem := new(MemoryStoreSortedMapItem)
	resp, err := s.client.Call(ctx, GetMemoryStoreSortedMapItem.Compile(nil, universeId, sortedMapId, itemId), nil, memoryStoreSortedMapItem, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return memoryStoreSortedMapItem, resp, nil
}

// DeleteMemoryStoreSortedMapItem will delete a specific item in the memory store sorted map for a specific universe.
//
// Required scopes: universe.memory-store.sorted-map:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/MemoryStoreSortedMapItem#Cloud_DeleteMemoryStoreSortedMapItem
//
// [DELETE] /cloud/v2/universes/{universe_id}/memory-store/sorted-maps/{sorted_map_id}/items/{item_id}
func (s *DataAndMemoryStoreService) DeleteMemoryStoreSortedMapItem(ctx context.Context, universeId, sortedMapId, itemId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	resp, err := s.client.Call(ctx, DeleteMemoryStoreSortedMapItem.Compile(nil, universeId, sortedMapId, itemId), nil, nil, reqOpts...)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

type MemoryStoreSortedMapItemUpdate struct {
	Value         *any    `json:"value,omitempty"`
	TTL           *string `json:"ttl,omitempty"`
	ID            *string `json:"id,omitempty"`
	StringSortKey *string `json:"stringSortKey,omitempty"`
}

type MemoryStoreSortedMapItemUpdateOpts struct {
	ID string `url:"id,omitempty"`
}

// UpdateMemoryStoreSortedMapItem will update a specific item in the memory store sorted map for a specific universe.
//
// Required scopes: universe.memory-store.sorted-map:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/MemoryStoreSortedMapItem#Cloud_UpdateMemoryStoreSortedMapItem
//
// [PATCH] /cloud/v2/universes/{universe_id}/memory-store/sorted-maps/{sorted_map_id}/items/{item_id}
func (s *DataAndMemoryStoreService) UpdateMemoryStoreSortedMapItem(ctx context.Context, universeId, sortedMapId, itemId string, data MemoryStoreSortedMapItemUpdate, opts *MemoryStoreSortedMapItemUpdateOpts, reqOpts ...goblox.RequestOpt) (*MemoryStoreSortedMapItem, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	memoryStoreSortedMapItem := new(MemoryStoreSortedMapItem)
	resp, err := s.client.Call(ctx, UpdateMemoryStoreSortedMapItem.Compile(q, universeId, sortedMapId, itemId), data, memoryStoreSortedMapItem, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return memoryStoreSortedMapItem, resp, nil
}

type OrderedDataStoreEntry struct {
	Path  string `json:"path"`
	Value int    `json:"value"`
	ID    string `json:"id"`
}

type OrderedDataStoreEntryList struct {
	OrderedDataStoreEntries []OrderedDataStoreEntry `json:"orderedDataStoreEntries"`
	NextPageToken           string                  `json:"nextPageToken"`
}

type ListOrderedDataStoreEntriesOptions struct {
	MaxPageSize *int    `url:"maxPageSize,omitempty"`
	PageToken   *string `url:"pageToken,omitempty"`
	OrderBy     *string `url:"orderBy,omitempty"`
	Filter      *string `url:"filter,omitempty"`
}

// -
//
// Required scopes: universe.ordered-data-store.scope.entry:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/OrderedDataStoreEntry#Cloud_ListOrderedDataStoreEntries
//
// [GET] /cloud/v2/universes/{universe_id}/ordered-data-stores/{ordered_data_store_id}/scopes/{scope_id}/entries
func (s *DataAndMemoryStoreService) ListOrderedDataStoreEntries(ctx context.Context, universeId, orderedDataStoreId, scopeId string, opts *ListOrderedDataStoreEntriesOptions, reqOpts ...goblox.RequestOpt) (*OrderedDataStoreEntryList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	orderedDataStoreEntryList := new(OrderedDataStoreEntryList)
	resp, err := s.client.Call(ctx, ListOrderedDataStoreEntries.Compile(q, universeId, orderedDataStoreId, scopeId), nil, orderedDataStoreEntryList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return orderedDataStoreEntryList, resp, nil
}

type OrderedDataStoreEntryCreate struct {
	Value *int `json:"value,omitempty"`
}

type OrderedDataStoreEntryCreateOptions struct {
	ID *string `url:"id,omitempty"`
}

// -
//
// Required scopes: universe.ordered-data-store.scope.entry:write
//
// Roblox Opencloud API Docs:https://create.roblox.com/docs/en-us/cloud/reference/OrderedDataStoreEntry#Cloud_CreateOrderedDataStoreEntry
//
// [POST] /cloud/v2/universes/{universe_id}/ordered-data-stores/{ordered_data_store_id}/scopes/{scope_id}/entries
func (s *DataAndMemoryStoreService) CreateOrderedDataStoreEntry(ctx context.Context, universeId, orderedDataStoreId, scopeId string, data OrderedDataStoreEntryCreate, opts *OrderedDataStoreEntryCreateOptions, reqOpts ...goblox.RequestOpt) (*OrderedDataStoreEntry, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	orderedDataStoreEntry := new(OrderedDataStoreEntry)
	resp, err := s.client.Call(ctx, CreateOrderedDataStoreEntry.Compile(q, universeId, orderedDataStoreId, scopeId), data, orderedDataStoreEntry, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return orderedDataStoreEntry, resp, nil
}

// -
//
// Required scopes: universe.ordered-data-store.scope.entry:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/OrderedDataStoreEntry#Cloud_GetOrderedDataStoreEntry
//
// [GET] /cloud/v2/universes/{universe_id}/ordered-data-stores/{ordered_data_store_id}/scopes/{scope_id}/entries/{entry_id}
func (s *DataAndMemoryStoreService) GetOrderedDataStoreEntry(ctx context.Context, universeId, orderedDataStoreId, scopeId, entryId string, reqOpts ...goblox.RequestOpt) (*OrderedDataStoreEntry, *Response, error) {
	orderedDataStoreEntry := new(OrderedDataStoreEntry)
	resp, err := s.client.Call(ctx, GetOrderedDataStoreEntry.Compile(nil, universeId, orderedDataStoreId, scopeId, entryId), nil, orderedDataStoreEntry, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return orderedDataStoreEntry, resp, nil
}

// -
//
// Required scopes: universe.ordered-data-store.scope.entry:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/OrderedDataStoreEntry#Cloud_DeleteOrderedDataStoreEntry
//
// [DELETE] /cloud/v2/universes/{universe_id}/ordered-data-stores/{ordered_data_store_id}/scopes/{scope_id}/entries/{entry_id}
func (s *DataAndMemoryStoreService) DeleteOrderedDataStoreEntry(ctx context.Context, universeId, orderedDataStoreId, scopeId, entryId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	resp, err := s.client.Call(ctx, DeleteOrderedDataStoreEntry.Compile(nil, universeId, orderedDataStoreId, scopeId, entryId), nil, nil, reqOpts...)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

type OrderedDataStoreEntryUpdate struct {
	Value *int `json:"value,omitempty"`
}

type OrderedDataStoreEntryUpdateOpts struct {
	AllowMissing *bool `url:"allowMissing,omitempty"`
}

// -
//
// Required scopes: universe.ordered-data-store.scope.entry:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/OrderedDataStoreEntry#Cloud_UpdateOrderedDataStoreEntry
//
// [PATCH] /cloud/v2/universes/{universe_id}/ordered-data-stores/{ordered_data_store_id}/scopes/{scope_id}/entries/{entry_id}
func (s *DataAndMemoryStoreService) UpdateOrderedDataStoreEntry(ctx context.Context, universeId, orderedDataStoreId, scopeId, entryId string, data OrderedDataStoreEntryUpdate, opts *OrderedDataStoreEntryUpdateOpts, reqOpts ...goblox.RequestOpt) (*OrderedDataStoreEntry, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	orderedDataStoreEntry := new(OrderedDataStoreEntry)
	resp, err := s.client.Call(ctx, UpdateOrderedDataStoreEntry.Compile(q, universeId, orderedDataStoreId, scopeId, entryId), data, orderedDataStoreEntry, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return orderedDataStoreEntry, resp, nil
}

type OrderedDataStoreEntryIncrement struct {
	Amount *int `json:"amount,omitempty"`
}

// -
//
// Required scopes: universe.ordered-data-store.scope.entry:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/OrderedDataStoreEntry#Cloud_IncrementOrderedDataStoreEntry
//
// [POST] /cloud/v2/universes/{universe_id}/ordered-data-stores/{ordered_data_store_id}/scopes/{scope_id}/entries/{entry_id}:increment
func (s *DataAndMemoryStoreService) IncrementOrderedDataStoreEntry(ctx context.Context, universeId, orderedDataStoreId, scopeId, entryId string, data OrderedDataStoreEntryIncrement, reqOpts ...goblox.RequestOpt) (*OrderedDataStoreEntry, *Response, error) {
	orderedDataStoreEntry := new(OrderedDataStoreEntry)
	resp, err := s.client.Call(ctx, IncrementOrderedDataStoreEntry.Compile(nil, universeId, orderedDataStoreId, scopeId, entryId), data, orderedDataStoreEntry, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return orderedDataStoreEntry, resp, nil
}

var (
	DeleteDataStore = goblox.NewEndpoint(
		http.MethodDelete,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.control:delete")

	UndeleteDataStore = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/universes/{universeId}/data-stores/{dataStoreId}:undelete",
		auth.KindAPIKey,
	).WithScopes("universe-datastores.control:delete")

	GetMemoryStoreFlushOperation = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/memory-store/operations/{operationId}",
		auth.KindAPIKey,
	).WithScopes("memory-store:get")
)

// DeleteDataStore will schedule the specified data store for permanent deletion.
//
// Required scopes: universe-datastores.control:delete
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/storage#Cloud_DeleteDataStore
//
// [DELETE] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}
func (s *DataAndMemoryStoreService) DeleteDataStore(ctx context.Context, universeId, dataStoreId string, reqOpts ...goblox.RequestOpt) (*DataStore, *Response, error) {
	dataStore := new(DataStore)
	resp, err := s.client.Call(ctx, DeleteDataStore.Compile(nil, universeId, dataStoreId), nil, dataStore, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dataStore, resp, nil
}

type UndeleteDataStoreRequest map[string]any

// UndeleteDataStore will restore the data store.
//
// Required scopes: universe-datastores.control:delete
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/storage#Cloud_UndeleteDataStore
//
// [POST] /cloud/v2/universes/{universe_id}/data-stores/{data_store_id}:undelete
func (s *DataAndMemoryStoreService) UndeleteDataStore(ctx context.Context, universeId, dataStoreId string, data UndeleteDataStoreRequest, reqOpts ...goblox.RequestOpt) (*DataStore, *Response, error) {
	dataStore := new(DataStore)
	resp, err := s.client.Call(ctx, UndeleteDataStore.Compile(nil, universeId, dataStoreId), data, dataStore, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dataStore, resp, nil
}

type GetMemoryStoreFlushOperationOptionsScope string

const (
	GetMemoryStoreFlushOperationOptionsScopeLive GetMemoryStoreFlushOperationOptionsScope = "LIVE"
	GetMemoryStoreFlushOperationOptionsScopeTest GetMemoryStoreFlushOperationOptionsScope = "TEST"
)

type GetMemoryStoreFlushOperationOptions struct {
	Scope *GetMemoryStoreFlushOperationOptionsScope `url:"scope,omitempty"`
}

// GetMemoryStoreFlushOperation will retrieve the status of the operation to flush the memory stores of a
// universe.
//
// Required scopes: memory-store:get
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/storage#Cloud_GetMemoryStoreFlushOperation
//
// [GET] /cloud/v2/universes/{universe_id}/memory-store/operations/{operation_id}
func (s *DataAndMemoryStoreService) GetMemoryStoreFlushOperation(ctx context.Context, universeId, operationId string, opts *GetMemoryStoreFlushOperationOptions, reqOpts ...goblox.RequestOpt) (*Operation, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	operation := new(Operation)
	resp, err := s.client.Call(ctx, GetMemoryStoreFlushOperation.Compile(q, universeId, operationId), nil, operation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return operation, resp, nil
}
