package opencloud

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	DeleteSave = goblox.NewEndpoint(
		http.MethodDelete,
		"/toolbox-service/v1/saves",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("creator-store-save:write")

	ListSaves = goblox.NewEndpoint(
		http.MethodGet,
		"/toolbox-service/v1/saves",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("creator-store-save:read")

	CreateSave = goblox.NewEndpoint(
		http.MethodPost,
		"/toolbox-service/v1/saves",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("creator-store-save:write")

	BulkDeleteSaves = goblox.NewEndpoint(
		http.MethodPost,
		"/toolbox-service/v1/saves:bulkDelete",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("creator-store-save:write")

	GetCreatorStoreAsset = goblox.NewEndpoint(
		http.MethodGet,
		"/toolbox-service/v2/assets/{id}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("creator-store-product:read")

	SearchCreatorStoreAssetsByQuery = goblox.NewEndpoint(
		http.MethodGet,
		"/toolbox-service/v2/assets:search",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("creator-store-product:read")

	SearchCreatorStoreAssets = goblox.NewEndpoint(
		http.MethodPost,
		"/toolbox-service/v2/assets:search",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("creator-store-product:read")
)

// CreatorStoreService handles communication with the Creator Store (toolbox) endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type CreatorStoreService service

type SearchCategoryType string

const (
	SearchCategoryTypeAudio      SearchCategoryType = "Audio"
	SearchCategoryTypePlace      SearchCategoryType = "Place"
	SearchCategoryTypeModel      SearchCategoryType = "Model"
	SearchCategoryTypeDecal      SearchCategoryType = "Decal"
	SearchCategoryTypePlugin     SearchCategoryType = "Plugin"
	SearchCategoryTypeMeshPart   SearchCategoryType = "MeshPart"
	SearchCategoryTypeVideo      SearchCategoryType = "Video"
	SearchCategoryTypeFontFamily SearchCategoryType = "FontFamily"
)

type DeleteSaveOptions struct {
	TargetType     *SearchCategoryType `url:"targetType,omitempty"`
	TargetID       *int64              `url:"targetId,omitempty"`
	CollectionName *string             `url:"collectionName,omitempty"`
}

// DeleteSave will delete a save.
//
// Required scopes: creator-store-save:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/creator-store#Saves_DeleteSave
//
// [DELETE] /toolbox-service/v1/saves
func (s *CreatorStoreService) DeleteSave(ctx context.Context, opts *DeleteSaveOptions, reqOpts ...goblox.RequestOpt) (*Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, err
	}

	return s.client.Call(ctx, DeleteSave.Compile(q), nil, nil, reqOpts...)
}

type SavesSortCategory string

const (
	SavesSortCategoryName         SavesSortCategory = "Name"
	SavesSortCategoryCreator      SavesSortCategory = "Creator"
	SavesSortCategoryRatings      SavesSortCategory = "Ratings"
	SavesSortCategoryTargetType   SavesSortCategory = "TargetType"
	SavesSortCategoryDateSaved    SavesSortCategory = "DateSaved"
	SavesSortCategoryLastModified SavesSortCategory = "LastModified"
	SavesSortCategoryPrice        SavesSortCategory = "Price"
)

type CreatorStoreSortDirection string

const (
	CreatorStoreSortDirectionNone       CreatorStoreSortDirection = "None"
	CreatorStoreSortDirectionAscending  CreatorStoreSortDirection = "Ascending"
	CreatorStoreSortDirectionDescending CreatorStoreSortDirection = "Descending"
)

type ListSavesOptions struct {
	TargetType      *SearchCategoryType        `url:"targetType,omitempty"`
	TargetID        *int64                     `url:"targetId,omitempty"`
	CollectionName  *string                    `url:"collectionName,omitempty"`
	SortBy          *SavesSortCategory         `url:"sortBy,omitempty"`
	SortDirection   *CreatorStoreSortDirection `url:"sortDirection,omitempty"`
	Limit           *int                       `url:"limit,omitempty"`
	Page            *int                       `url:"page,omitempty"`
	Keyword         *string                    `url:"keyword,omitempty"`
	HideOwnedAssets *bool                      `url:"hideOwnedAssets,omitempty"`
}

type VotingModel struct {
	ShowVotes     bool  `json:"showVotes"`
	UpVotes       int64 `json:"upVotes"`
	DownVotes     int64 `json:"downVotes"`
	CanVote       bool  `json:"canVote"`
	UserVote      bool  `json:"userVote"`
	HasVoted      bool  `json:"hasVoted"`
	VoteCount     int64 `json:"voteCount"`
	UpVotePercent int   `json:"upVotePercent"`
}

type CreatorModelV2 struct {
	Creator                    string `json:"creator"`
	UserID                     int64  `json:"userId"`
	GroupID                    int64  `json:"groupId"`
	Name                       string `json:"name"`
	Verified                   bool   `json:"verified"`
	LatestGroupUpdaterUserID   int64  `json:"latestGroupUpdaterUserId"`
	LatestGroupUpdaterUserName string `json:"latestGroupUpdaterUserName"`
}

type CreatorStoreDecimal struct {
	Significand int64 `json:"significand"`
	Exponent    int   `json:"exponent"`
}

type CreatorStoreMoney struct {
	CurrencyCode string              `json:"currencyCode"`
	Quantity     CreatorStoreDecimal `json:"quantity"`
}

type CreatorStoreAssetProduct struct {
	PurchasePrice CreatorStoreMoney `json:"purchasePrice"`
	Purchasable   bool              `json:"purchasable"`
}

type CreatorStoreAsset struct {
	Voting              VotingModel              `json:"voting"`
	Creator             CreatorModelV2           `json:"creator"`
	CreatorStoreProduct CreatorStoreAssetProduct `json:"creatorStoreProduct"`
	Asset               json.RawMessage          `json:"asset"`
}

type HydratedSave struct {
	Owned             bool              `json:"owned"`
	DateSaved         string            `json:"dateSaved"`
	CreatorStoreAsset CreatorStoreAsset `json:"creatorStoreAsset"`
}

type GetSavesResponse struct {
	TotalCount int            `json:"totalCount"`
	Saves      []HydratedSave `json:"saves"`
}

// ListSaves will get saves from a collection.
//
// Required scopes: creator-store-save:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/creator-store#Saves_GetSaves
//
// [GET] /toolbox-service/v1/saves
func (s *CreatorStoreService) ListSaves(ctx context.Context, opts *ListSavesOptions, reqOpts ...goblox.RequestOpt) (*GetSavesResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	getSavesResponse := new(GetSavesResponse)
	resp, err := s.client.Call(ctx, ListSaves.Compile(q), nil, getSavesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getSavesResponse, resp, nil
}

type CreateSaveRequest struct {
	TargetType     *SearchCategoryType `json:"targetType,omitempty"`
	TargetID       *int64              `json:"targetId,omitempty"`
	CollectionName *string             `json:"collectionName,omitempty"`
}

// CreateSave will create a save.
//
// Required scopes: creator-store-save:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/creator-store#Saves_CreateSave
//
// [POST] /toolbox-service/v1/saves
func (s *CreatorStoreService) CreateSave(ctx context.Context, data CreateSaveRequest, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, CreateSave.Compile(nil), data, nil, reqOpts...)
}

type DeleteSaveRequest struct {
	TargetType     *SearchCategoryType `json:"targetType,omitempty"`
	TargetID       *int64              `json:"targetId,omitempty"`
	CollectionName *string             `json:"collectionName,omitempty"`
}

type BulkDeleteSavesRequest struct {
	Targets []DeleteSaveRequest `json:"targets,omitempty"`
}

type BulkDeleteSavesResponse struct {
	DeletedCount int `json:"deletedCount"`
}

// BulkDeleteSaves will delete multiple saves at once.
//
// Required scopes: creator-store-save:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/creator-store#Saves_BulkDeleteSaves
//
// [POST] /toolbox-service/v1/saves:bulkDelete
func (s *CreatorStoreService) BulkDeleteSaves(ctx context.Context, data BulkDeleteSavesRequest, reqOpts ...goblox.RequestOpt) (*BulkDeleteSavesResponse, *Response, error) {
	bulkDeleteSavesResponse := new(BulkDeleteSavesResponse)
	resp, err := s.client.Call(ctx, BulkDeleteSaves.Compile(nil), data, bulkDeleteSavesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return bulkDeleteSavesResponse, resp, nil
}

// GetCreatorStoreAsset will get details for a single Creator Store asset.
//
// Required scopes: creator-store-product:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/assets#Toolbox_GetAssetDetails
//
// [GET] /toolbox-service/v2/assets/{id}
func (s *CreatorStoreService) GetCreatorStoreAsset(ctx context.Context, id string, reqOpts ...goblox.RequestOpt) (*CreatorStoreAsset, *Response, error) {
	creatorStoreAsset := new(CreatorStoreAsset)
	resp, err := s.client.Call(ctx, GetCreatorStoreAsset.Compile(nil, id), nil, creatorStoreAsset, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return creatorStoreAsset, resp, nil
}

type ModelSubType string

const (
	ModelSubTypeAd           ModelSubType = "Ad"
	ModelSubTypeMaterialPack ModelSubType = "MaterialPack"
	ModelSubTypePackage      ModelSubType = "Package"
)

type CreatorStoreSortCategory string

const (
	CreatorStoreSortCategoryRelevance     CreatorStoreSortCategory = "Relevance"
	CreatorStoreSortCategoryTrending      CreatorStoreSortCategory = "Trending"
	CreatorStoreSortCategoryTop           CreatorStoreSortCategory = "Top"
	CreatorStoreSortCategoryAudioDuration CreatorStoreSortCategory = "AudioDuration"
	CreatorStoreSortCategoryCreateTime    CreatorStoreSortCategory = "CreateTime"
	CreatorStoreSortCategoryUpdatedTime   CreatorStoreSortCategory = "UpdatedTime"
	CreatorStoreSortCategoryRatings       CreatorStoreSortCategory = "Ratings"
)

type SearchAudioTypeModel string

const (
	SearchAudioTypeModelMusic       SearchAudioTypeModel = "Music"
	SearchAudioTypeModelSoundEffect SearchAudioTypeModel = "SoundEffect"
	SearchAudioTypeModelUnknown     SearchAudioTypeModel = "Unknown"
)

type ModelInstanceType string

const (
	ModelInstanceTypeScript    ModelInstanceType = "Script"
	ModelInstanceTypeMeshPart  ModelInstanceType = "MeshPart"
	ModelInstanceTypeDecal     ModelInstanceType = "Decal"
	ModelInstanceTypeAnimation ModelInstanceType = "Animation"
	ModelInstanceTypeAudio     ModelInstanceType = "Audio"
	ModelInstanceTypeTool      ModelInstanceType = "Tool"
)

type SearchView string

const (
	SearchViewIDs  SearchView = "IDs"
	SearchViewCore SearchView = "Core"
	SearchViewFull SearchView = "Full"
)

type MusicChartType string

const (
	MusicChartTypeNone    MusicChartType = "None"
	MusicChartTypeCurrent MusicChartType = "Current"
	MusicChartTypeWeek    MusicChartType = "Week"
	MusicChartTypeMonth   MusicChartType = "Month"
	MusicChartTypeYear    MusicChartType = "Year"
)

type SearchCreatorStoreAssetsByQueryOptions struct {
	SearchCategoryType          *SearchCategoryType        `url:"searchCategoryType,omitempty"`
	Query                       *string                    `url:"query,omitempty"`
	ModelSubTypes               []ModelSubType             `url:"modelSubTypes,omitempty"`
	ExcludedModelSubTypes       []ModelSubType             `url:"excludedModelSubTypes,omitempty"`
	Creator                     *string                    `url:"creator,omitempty"`
	UserID                      *int64                     `url:"userId,omitempty"`
	GroupID                     *int64                     `url:"groupId,omitempty"`
	PageToken                   *string                    `url:"pageToken,omitempty"`
	PageNumber                  *int                       `url:"pageNumber,omitempty"`
	MaxPageSize                 *int                       `url:"maxPageSize,omitempty"`
	SortDirection               *CreatorStoreSortDirection `url:"sortDirection,omitempty"`
	SortCategory                *CreatorStoreSortCategory  `url:"sortCategory,omitempty"`
	AudioMinDurationSeconds     *int                       `url:"audioMinDurationSeconds,omitempty"`
	AudioMaxDurationSeconds     *int                       `url:"audioMaxDurationSeconds,omitempty"`
	AudioArtist                 *string                    `url:"audioArtist,omitempty"`
	AudioAlbum                  *string                    `url:"audioAlbum,omitempty"`
	AudioTypes                  []SearchAudioTypeModel     `url:"audioTypes,omitempty"`
	IncludedInstanceTypes       []ModelInstanceType        `url:"includedInstanceTypes,omitempty"`
	IncludeOnlyVerifiedCreators *bool                      `url:"includeOnlyVerifiedCreators,omitempty"`
	MinPriceCents               *int64                     `url:"minPriceCents,omitempty"`
	MaxPriceCents               *int64                     `url:"maxPriceCents,omitempty"`
	Facets                      []string                   `url:"facets,omitempty"`
	Tags                        []string                   `url:"tags,omitempty"`
	CategoryPath                *string                    `url:"categoryPath,omitempty"`
	SearchView                  *SearchView                `url:"searchView,omitempty"`
	MusicChartType              *MusicChartType            `url:"musicChartType,omitempty"`
	Swimlane                    *string                    `url:"swimlane,omitempty"`
}

type QueryFacets struct {
	AppliedFacets   []string `json:"appliedFacets"`
	AvailableFacets []string `json:"availableFacets"`
}

type QueryCorrection struct {
	SuggestedQuery string `json:"suggestedQuery"`
}

type SearchCreatorStoreAssetsResponse struct {
	NextPageToken      string              `json:"nextPageToken"`
	QueryFacets        QueryFacets         `json:"queryFacets"`
	CreatorStoreAssets []CreatorStoreAsset `json:"creatorStoreAssets"`
	TotalResults       int64               `json:"totalResults"`
	QueryCorrection    QueryCorrection     `json:"queryCorrection"`
	FilteredKeyword    string              `json:"filteredKeyword"`
}

// SearchCreatorStoreAssetsByQuery will search the Creator Store for assets using query parameters.
//
// Required scopes: creator-store-product:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/assets#Toolbox_SearchCreatorStoreAssetsDeprecated
//
// [GET] /toolbox-service/v2/assets:search
func (s *CreatorStoreService) SearchCreatorStoreAssetsByQuery(ctx context.Context, opts *SearchCreatorStoreAssetsByQueryOptions, reqOpts ...goblox.RequestOpt) (*SearchCreatorStoreAssetsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	searchCreatorStoreAssetsResponse := new(SearchCreatorStoreAssetsResponse)
	resp, err := s.client.Call(ctx, SearchCreatorStoreAssetsByQuery.Compile(q), nil, searchCreatorStoreAssetsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return searchCreatorStoreAssetsResponse, resp, nil
}

type AudioSearchFilters struct {
	AudioTypes         []SearchAudioTypeModel `json:"audioTypes,omitempty"`
	MinDurationSeconds *int                   `json:"minDurationSeconds,omitempty"`
	MaxDurationSeconds *int                   `json:"maxDurationSeconds,omitempty"`
	Artist             *string                `json:"artist,omitempty"`
	Album              *string                `json:"album,omitempty"`
	MusicChartType     *MusicChartType        `json:"musicChartType,omitempty"`
}

type ModelSearchFilters struct {
	IncludedSubTypes      []ModelSubType      `json:"includedSubTypes,omitempty"`
	ExcludedSubTypes      []ModelSubType      `json:"excludedSubTypes,omitempty"`
	IncludedInstanceTypes []ModelInstanceType `json:"includedInstanceTypes,omitempty"`
}

type SearchCreatorStoreAssetsRequest struct {
	SearchCategoryType          *SearchCategoryType        `json:"searchCategoryType,omitempty"`
	Query                       *string                    `json:"query,omitempty"`
	Image                       *string                    `json:"image,omitempty"`
	UserID                      *int64                     `json:"userId,omitempty"`
	GroupID                     *int64                     `json:"groupId,omitempty"`
	PageToken                   *string                    `json:"pageToken,omitempty"`
	PageNumber                  *int                       `json:"pageNumber,omitempty"`
	MaxPageSize                 *int                       `json:"maxPageSize,omitempty"`
	SortDirection               *CreatorStoreSortDirection `json:"sortDirection,omitempty"`
	SortCategory                *CreatorStoreSortCategory  `json:"sortCategory,omitempty"`
	IncludeOnlyVerifiedCreators *bool                      `json:"includeOnlyVerifiedCreators,omitempty"`
	MinPriceCents               *int64                     `json:"minPriceCents,omitempty"`
	MaxPriceCents               *int64                     `json:"maxPriceCents,omitempty"`
	Facets                      []string                   `json:"facets,omitempty"`
	Tags                        []string                   `json:"tags,omitempty"`
	CategoryPath                *string                    `json:"categoryPath,omitempty"`
	Swimlane                    *string                    `json:"swimlane,omitempty"`
	SearchView                  *SearchView                `json:"searchView,omitempty"`
	AudioSearchFilters          *AudioSearchFilters        `json:"audioSearchFilters,omitempty"`
	ModelSearchFilters          *ModelSearchFilters        `json:"modelSearchFilters,omitempty"`
}

// SearchCreatorStoreAssets will search the Creator Store for assets.
//
// Required scopes: creator-store-product:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/creator-store#Toolbox_SearchCreatorStoreAssets
//
// [POST] /toolbox-service/v2/assets:search
func (s *CreatorStoreService) SearchCreatorStoreAssets(ctx context.Context, data SearchCreatorStoreAssetsRequest, reqOpts ...goblox.RequestOpt) (*SearchCreatorStoreAssetsResponse, *Response, error) {
	searchCreatorStoreAssetsResponse := new(SearchCreatorStoreAssetsResponse)
	resp, err := s.client.Call(ctx, SearchCreatorStoreAssets.Compile(nil), data, searchCreatorStoreAssetsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return searchCreatorStoreAssetsResponse, resp, nil
}
