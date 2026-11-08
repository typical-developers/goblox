package legacy

import (
	"context"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	GetGames = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	GetGameMedia = goblox.NewEndpoint(
		http.MethodGet,
		"/v2/games/{universeId}/media",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	GetUserFavoriteGames = goblox.NewEndpoint(
		http.MethodGet,
		"/v2/users/{userId}/favorite/games",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	GetUserGames = goblox.NewEndpoint(
		http.MethodGet,
		"/v2/users/{userId}/games",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	GetGroupGames = goblox.NewEndpoint(
		http.MethodGet,
		"/v2/groups/{groupId}/gamesV2",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	IsGameFavorited = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games/{universeId}/favorites",
		auth.KindCookie,
	).WithDomain("games")

	SetGameFavorite = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/games/{universeId}/favorites",
		auth.KindCookie,
	).WithDomain("games")

	GetGameFavoritesCount = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games/{universeId}/favorites/count",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	GetGameVotes = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games/votes",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	GetUserGameVote = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games/{universeId}/votes/user",
		auth.KindCookie,
	).WithDomain("games")

	SetUserGameVote = goblox.NewEndpoint(
		http.MethodPatch,
		"/v1/games/{universeId}/user-votes",
		auth.KindCookie,
	).WithDomain("games")

	GetGameProductInfo = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games/games-product-info",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	GetPlaceDetails = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games/multiget-place-details",
		auth.KindCookie,
	).WithDomain("games")

	GetPlayabilityStatuses = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games/multiget-playability-status",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	GetGameRecommendations = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games/recommendations/game/{universeId}",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	GetGameServers = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games/{placeId}/servers/{serverType}",
		auth.KindCookie, auth.KindNone,
	).WithDomain("games")

	GetPrivateServers = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/games/{placeId}/private-servers",
		auth.KindCookie,
	).WithDomain("games")

	CreateVIPServer = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/games/vip-servers/{universeId}",
		auth.KindCookie,
	).WithDomain("games")

	GetVIPServer = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/vip-servers/{id}",
		auth.KindCookie,
	).WithDomain("games")

	UpdateVIPServer = goblox.NewEndpoint(
		http.MethodPatch,
		"/v1/vip-servers/{id}",
		auth.KindCookie,
	).WithDomain("games")

	UpdateVIPServerSubscription = goblox.NewEndpoint(
		http.MethodPatch,
		"/v1/vip-servers/{id}/subscription",
		auth.KindCookie,
	).WithDomain("games")

	UpdateVIPServerPermissions = goblox.NewEndpoint(
		http.MethodPatch,
		"/v1/vip-servers/{id}/permissions",
		auth.KindCookie,
	).WithDomain("games")
)

type GamesService service

type GameCreator struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	IsRNVAccount     bool   `json:"isRNVAccount"`
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"`
}

type RefundPolicy struct {
	PolicyText       string `json:"policyText"`
	LearnMoreBaseURL string `json:"learnMoreBaseUrl"`
	Locale           string `json:"locale"`
	ArticleID        string `json:"articleId"`
}

type GameDetail struct {
	ID                        int64         `json:"id"`
	RootPlaceID               int64         `json:"rootPlaceId"`
	Name                      string        `json:"name"`
	Description               string        `json:"description"`
	SourceName                string        `json:"sourceName"`
	SourceDescription         string        `json:"sourceDescription"`
	Creator                   GameCreator   `json:"creator"`
	Price                     *int64        `json:"price"`
	AllowedGearGenres         []string      `json:"allowedGearGenres"`
	AllowedGearCategories     []string      `json:"allowedGearCategories"`
	IsGenreEnforced           bool          `json:"isGenreEnforced"`
	CopyingAllowed            bool          `json:"copyingAllowed"`
	Playing                   int64         `json:"playing"`
	Visits                    int64         `json:"visits"`
	MaxPlayers                int           `json:"maxPlayers"`
	Created                   string        `json:"created"`
	Updated                   string        `json:"updated"`
	StudioAccessToApisAllowed bool          `json:"studioAccessToApisAllowed"`
	CreateVipServersAllowed   bool          `json:"createVipServersAllowed"`
	UniverseAvatarType        Enum          `json:"universeAvatarType"`
	Genre                     string        `json:"genre"`
	GenreL1                   string        `json:"genre_l1"`
	GenreL2                   string        `json:"genre_l2"`
	UntranslatedGenreL1       string        `json:"untranslated_genre_l1"`
	IsAllGenre                bool          `json:"isAllGenre"`
	IsFavoritedByUser         bool          `json:"isFavoritedByUser"`
	FavoritedCount            int64         `json:"favoritedCount"`
	LicenseDescription        string        `json:"licenseDescription"`
	RefundLink                string        `json:"refundLink"`
	LocalizedFiatPrice        string        `json:"localizedFiatPrice"`
	RefundPolicy              *RefundPolicy `json:"refundPolicy"`
	CanonicalURLPath          string        `json:"canonicalUrlPath"`
	IsContentRestricted       bool          `json:"isContentRestricted"`
	CreationSource            string        `json:"creationSource"`
}

type GetGamesOptions struct {
	UniverseIDs []int64 `url:"universeIds,comma"` // Required. At most 50.
	Fields      *string `url:"fields,omitempty"`  // Comma-separated field names. All fields are returned when nil.
}

// GetGames will get the details of a list of games, including player and visit counts.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games
//
// [GET] https://games.roblox.com/v1/games
func (s *GamesService) GetGames(ctx context.Context, opts *GetGamesOptions, reqOpts ...goblox.RequestOpt) ([]GameDetail, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	results := new(WrappedResponse[[]GameDetail])
	resp, err := s.client.Call(ctx, GetGames.Compile(q), nil, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}

type GetGameMediaOptions struct {
	FetchAllExperienceRelatedMedia *bool `url:"fetchAllExperienceRelatedMedia,omitempty"`
}

type GameMediaItem struct {
	AssetTypeID int64  `json:"assetTypeId"`
	AssetType   string `json:"assetType"`
	ImageID     *int64 `json:"imageId"`
	VideoHash   string `json:"videoHash"`
	VideoTitle  string `json:"videoTitle"`
	Approved    bool   `json:"approved"`
	AltText     string `json:"altText"`
	VideoID     string `json:"videoId"`
}

// GetGameMedia will get a game's media, such as its thumbnail images and videos.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v2_games__universeId__media
//
// [GET] https://games.roblox.com/v2/games/{universeId}/media
func (s *GamesService) GetGameMedia(ctx context.Context, universeId string, opts *GetGameMediaOptions, reqOpts ...goblox.RequestOpt) ([]GameMediaItem, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	results := new(WrappedResponse[[]GameMediaItem])
	resp, err := s.client.Call(ctx, GetGameMedia.Compile(q, universeId), nil, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}

type GameContentMetadata struct {
	BadgePosition      string `json:"badgePosition"`
	BadgeAnalyticsID   string `json:"badgeAnalyticsId"`
	BadgeType          string `json:"badgeType"`
	BadgeIcon          string `json:"badgeIcon"`
	BadgeComponentType string `json:"badgeComponentType"`
}

// Game is a game in a list of a user's or group's games. Price and ContentMetadata are only
// set for favorite games.
type Game struct {
	ID              int64                `json:"id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	Creator         RelatedEntity        `json:"creator"`
	RootPlace       RelatedEntity        `json:"rootPlace"`
	Created         string               `json:"created"`
	Updated         string               `json:"updated"`
	PlaceVisits     int64                `json:"placeVisits"`
	Price           *int64               `json:"price,omitempty"`
	ContentMetadata *GameContentMetadata `json:"contentMetadata,omitempty"`
}

type ListGamesOptions struct {
	AccessFilter *int      `url:"accessFilter,omitempty"` // 1, 2, or 4.
	Limit        *int      `url:"limit,omitempty"`        // 10, 25, 50, or 100.
	Cursor       *string   `url:"cursor,omitempty"`
	SortOrder    SortOrder `url:"sortOrder,omitempty"`
}

// GetUserFavoriteGames will get a page of a user's favorite games.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v2_users__userId__favorite_games
//
// [GET] https://games.roblox.com/v2/users/{userId}/favorite/games
func (s *GamesService) GetUserFavoriteGames(ctx context.Context, userId string, opts *ListGamesOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[Game], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[Game])
	resp, err := s.client.Call(ctx, GetUserFavoriteGames.Compile(q, userId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

// GetUserGames will get a page of the games created by a user.
// Limit can be 10, 25, or 50.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v2_users__userId__games
//
// [GET] https://games.roblox.com/v2/users/{userId}/games
func (s *GamesService) GetUserGames(ctx context.Context, userId string, opts *ListGamesOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[Game], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[Game])
	resp, err := s.client.Call(ctx, GetUserGames.Compile(q, userId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

// GetGroupGames will get a page of the games created by a group.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v2_groups__groupId__gamesV2
//
// [GET] https://games.roblox.com/v2/groups/{groupId}/gamesV2
func (s *GamesService) GetGroupGames(ctx context.Context, groupId string, opts *ListGamesOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[Game], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[Game])
	resp, err := s.client.Call(ctx, GetGroupGames.Compile(q, groupId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

type gameFavoriteBody struct {
	IsFavorited bool `json:"isFavorited"`
}

// IsGameFavorited will check whether the authenticated user has favorited a game.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games__universeId__favorites
//
// [GET] https://games.roblox.com/v1/games/{universeId}/favorites
func (s *GamesService) IsGameFavorited(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (bool, *Response, error) {
	result := new(gameFavoriteBody)
	resp, err := s.client.Call(ctx, IsGameFavorited.Compile(nil, universeId), nil, result, reqOpts...)
	if err != nil {
		return false, resp, err
	}

	return result.IsFavorited, resp, nil
}

// SetGameFavorite will favorite or unfavorite a game for the authenticated user.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_post_v1_games__universeId__favorites
//
// [POST] https://games.roblox.com/v1/games/{universeId}/favorites
func (s *GamesService) SetGameFavorite(ctx context.Context, universeId string, isFavorited bool, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, SetGameFavorite.Compile(nil, universeId), gameFavoriteBody{IsFavorited: isFavorited}, nil, reqOpts...)
}

// GetGameFavoritesCount will get the number of times a game has been favorited.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games__universeId__favorites_count
//
// [GET] https://games.roblox.com/v1/games/{universeId}/favorites/count
func (s *GamesService) GetGameFavoritesCount(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	var result struct {
		FavoritesCount int64 `json:"favoritesCount"`
	}
	resp, err := s.client.Call(ctx, GetGameFavoritesCount.Compile(nil, universeId), nil, &result, reqOpts...)
	if err != nil {
		return 0, resp, err
	}

	return result.FavoritesCount, resp, nil
}

type GetGameVotesOptions struct {
	UniverseIDs []int64 `url:"universeIds,comma"` // Required.
}

type GameVotes struct {
	ID        int64 `json:"id"`
	UpVotes   int64 `json:"upVotes"`
	DownVotes int64 `json:"downVotes"`
}

// GetGameVotes will get the up and down votes for a list of games.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games_votes
//
// [GET] https://games.roblox.com/v1/games/votes
func (s *GamesService) GetGameVotes(ctx context.Context, opts *GetGameVotesOptions, reqOpts ...goblox.RequestOpt) ([]GameVotes, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	results := new(WrappedResponse[[]GameVotes])
	resp, err := s.client.Call(ctx, GetGameVotes.Compile(q), nil, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}

type UserGameVote struct {
	CanVote              bool   `json:"canVote"`
	UserVote             *bool  `json:"userVote"` // nil when the user hasn't voted.
	ReasonForNotVoteable string `json:"reasonForNotVoteable"`
}

// GetUserGameVote will get the authenticated user's vote on a game.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games__universeId__votes_user
//
// [GET] https://games.roblox.com/v1/games/{universeId}/votes/user
func (s *GamesService) GetUserGameVote(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*UserGameVote, *Response, error) {
	vote := new(UserGameVote)
	resp, err := s.client.Call(ctx, GetUserGameVote.Compile(nil, universeId), nil, vote, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return vote, resp, nil
}

type userGameVoteBody struct {
	Vote *bool `json:"vote"`
}

// SetUserGameVote will set the authenticated user's vote on a game. true is an up vote,
// false is a down vote, and nil removes the vote.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_patch_v1_games__universeId__user_votes
//
// [PATCH] https://games.roblox.com/v1/games/{universeId}/user-votes
func (s *GamesService) SetUserGameVote(ctx context.Context, universeId string, vote *bool, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, SetUserGameVote.Compile(nil, universeId), userGameVoteBody{Vote: vote}, nil, reqOpts...)
}

type GetGameProductInfoOptions struct {
	UniverseIDs []int64 `url:"universeIds,comma"` // Required. At most 100.
}

type FiatPurchaseData struct {
	LocalizedFiatPrice string `json:"localizedFiatPrice"`
	BasePriceID        string `json:"basePriceId"`
}

type GameProduct struct {
	UniverseID       int64             `json:"universeId"`
	IsForSale        bool              `json:"isForSale"`
	ProductID        int64             `json:"productId"`
	Price            *int64            `json:"price"`
	SellerID         int64             `json:"sellerId"`
	FiatPurchaseData *FiatPurchaseData `json:"fiatPurchaseData"`
}

// GetGameProductInfo will get the product info used to purchase a list of paid-access games.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games_games_product_info
//
// [GET] https://games.roblox.com/v1/games/games-product-info
func (s *GamesService) GetGameProductInfo(ctx context.Context, opts *GetGameProductInfoOptions, reqOpts ...goblox.RequestOpt) ([]GameProduct, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	results := new(WrappedResponse[[]GameProduct])
	resp, err := s.client.Call(ctx, GetGameProductInfo.Compile(q), nil, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}

type GetPlaceDetailsOptions struct {
	PlaceIDs []int64 `url:"placeIds,comma"` // Required.
}

type PlaceDetails struct {
	PlaceID             int64             `json:"placeId"`
	Name                string            `json:"name"`
	Description         string            `json:"description"`
	SourceName          string            `json:"sourceName"`
	SourceDescription   string            `json:"sourceDescription"`
	URL                 string            `json:"url"`
	Builder             string            `json:"builder"`
	BuilderID           int64             `json:"builderId"`
	HasVerifiedBadge    bool              `json:"hasVerifiedBadge"`
	IsPlayable          bool              `json:"isPlayable"`
	ReasonProhibited    string            `json:"reasonProhibited"`
	UniverseID          int64             `json:"universeId"`
	UniverseRootPlaceID int64             `json:"universeRootPlaceId"`
	Price               int64             `json:"price"`
	ImageToken          string            `json:"imageToken"`
	FiatPurchaseData    *FiatPurchaseData `json:"fiatPurchaseData"`
}

// GetPlaceDetails will get the details of a list of places, including the universe each belongs to.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games_multiget_place_details
//
// [GET] https://games.roblox.com/v1/games/multiget-place-details
func (s *GamesService) GetPlaceDetails(ctx context.Context, opts *GetPlaceDetailsOptions, reqOpts ...goblox.RequestOpt) ([]PlaceDetails, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	var results []PlaceDetails
	resp, err := s.client.Call(ctx, GetPlaceDetails.Compile(q), nil, &results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results, resp, nil
}

type GetPlayabilityStatusesOptions struct {
	UniverseIDs []int64 `url:"universeIds,comma"` // Required. At most 50.
}

type PlayabilityTreatmentData struct {
	TitleText           string `json:"titleText"`
	BodyText            string `json:"bodyText"`
	PrimaryActionText   string `json:"primaryActionText"`
	SecondaryActionText string `json:"secondaryActionText"`
}

type PlayabilityTreatment struct {
	Treatment string                   `json:"treatment"`
	Data      PlayabilityTreatmentData `json:"data"`
}

type PrivatePlaytestInfo struct {
	IsPlayable        bool `json:"isPlayable"`
	PlayabilityStatus Enum `json:"playabilityStatus"`
}

type PlayabilityStatus struct {
	UniverseID            int64                 `json:"universeId"`
	PlayabilityStatus     Enum                  `json:"playabilityStatus"`
	IsPlayable            bool                  `json:"isPlayable"`
	UnplayableDisplayText string                `json:"unplayableDisplayText"`
	PlayableUxTreatment   *PlayabilityTreatment `json:"playableUxTreatment"`
	UpsellUxTreatment     *PlayabilityTreatment `json:"upsellUxTreatment"`
	DemoModeAvailable     bool                  `json:"demoModeAvailable"`
	PrivatePlaytestInfo   *PrivatePlaytestInfo  `json:"privatePlaytestInfo"`
}

// GetPlayabilityStatuses will get whether the authenticated user can play each of a list of games.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games_multiget_playability_status
//
// [GET] https://games.roblox.com/v1/games/multiget-playability-status
func (s *GamesService) GetPlayabilityStatuses(ctx context.Context, opts *GetPlayabilityStatusesOptions, reqOpts ...goblox.RequestOpt) ([]PlayabilityStatus, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	var results []PlayabilityStatus
	resp, err := s.client.Call(ctx, GetPlayabilityStatuses.Compile(q), nil, &results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results, resp, nil
}

type GetGameRecommendationsOptions struct {
	PaginationKey             *string `url:"PaginationKey,omitempty"` // Leave nil for the first page.
	MaxRows                   *int    `url:"MaxRows,omitempty"`
	IsTruncatedResultsEnabled *bool   `url:"IsTruncatedResultsEnabled,omitempty"`
}

type RecommendedGame struct {
	UniverseID                   int64                `json:"universeId"`
	PlaceID                      int64                `json:"placeId"`
	Name                         string               `json:"name"`
	GameDescription              string               `json:"gameDescription"`
	Genre                        string               `json:"genre"`
	CreatorID                    int64                `json:"creatorId"`
	CreatorName                  string               `json:"creatorName"`
	CreatorType                  string               `json:"creatorType"`
	CreatorHasVerifiedBadge      bool                 `json:"creatorHasVerifiedBadge"`
	TotalUpVotes                 int64                `json:"totalUpVotes"`
	TotalDownVotes               int64                `json:"totalDownVotes"`
	PlayerCount                  int64                `json:"playerCount"`
	ImageToken                   string               `json:"imageToken"`
	IsSponsored                  bool                 `json:"isSponsored"`
	NativeAdData                 string               `json:"nativeAdData"`
	IsShowSponsoredLabel         bool                 `json:"isShowSponsoredLabel"`
	Price                        *int64               `json:"price"`
	AnalyticsIdentifier          string               `json:"analyticsIdentifier"`
	MinimumAge                   int                  `json:"minimumAge"`
	AgeRecommendationDisplayName string               `json:"ageRecommendationDisplayName"`
	ContentMetadata              *GameContentMetadata `json:"contentMetadata"`
	CanonicalURLPath             string               `json:"canonicalUrlPath"`
}

type GameRecommendations struct {
	Games             []RecommendedGame `json:"games"`
	NextPaginationKey *string           `json:"nextPaginationKey"` // nil when there are no more pages.
}

// GetGameRecommendations will get games that are recommended based on a game.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games_recommendations_game__universeId_
//
// [GET] https://games.roblox.com/v1/games/recommendations/game/{universeId}
func (s *GamesService) GetGameRecommendations(ctx context.Context, universeId string, opts *GetGameRecommendationsOptions, reqOpts ...goblox.RequestOpt) (*GameRecommendations, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	recommendations := new(GameRecommendations)
	resp, err := s.client.Call(ctx, GetGameRecommendations.Compile(q, universeId), nil, recommendations, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return recommendations, resp, nil
}

type ServerType string

const (
	ServerTypePublic ServerType = "Public"
	ServerTypeFriend ServerType = "Friend" // Needs a cookie.
)

type GetGameServersOptions struct {
	SortOrder        SortOrder `url:"sortOrder,omitempty"`
	ExcludeFullGames *bool     `url:"excludeFullGames,omitempty"`
	Limit            *int      `url:"limit,omitempty"` // 10, 25, 50, or 100.
	Cursor           *string   `url:"cursor,omitempty"`
}

type ServerPlayer struct {
	PlayerToken string `json:"playerToken"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type GameServer struct {
	ID           string         `json:"id"`
	MaxPlayers   int            `json:"maxPlayers"`
	Playing      int            `json:"playing"`
	PlayerTokens []string       `json:"playerTokens"`
	Players      []ServerPlayer `json:"players"`
	FPS          float64        `json:"fps"`
	Ping         int            `json:"ping"`
}

// GetGameServers will get a page of a place's running servers.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games__placeId__servers__serverType_
//
// [GET] https://games.roblox.com/v1/games/{placeId}/servers/{serverType}
func (s *GamesService) GetGameServers(ctx context.Context, placeId string, serverType ServerType, opts *GetGameServersOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[GameServer], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[GameServer])
	resp, err := s.client.Call(ctx, GetGameServers.Compile(q, placeId, serverType), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

type GetPrivateServersOptions struct {
	Limit  *int    `url:"limit,omitempty"` // 10, 25, 50, or 100.
	Cursor *string `url:"cursor,omitempty"`
}

type ServerOwner struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	DisplayName      string `json:"displayName"`
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"`
}

type PrivateServer struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	VIPServerID  int64          `json:"vipServerId"`
	AccessCode   string         `json:"accessCode"`
	Owner        ServerOwner    `json:"owner"`
	MaxPlayers   int            `json:"maxPlayers"`
	Playing      int            `json:"playing"`
	PlayerTokens []string       `json:"playerTokens"`
	Players      []ServerPlayer `json:"players"`
	FPS          float64        `json:"fps"`
	Ping         int            `json:"ping"`
}

type PrivateServersPage struct {
	GameJoinRestricted bool            `json:"gameJoinRestricted"`
	PreviousPageCursor string          `json:"previousPageCursor"`
	NextPageCursor     string          `json:"nextPageCursor"`
	Data               []PrivateServer `json:"data"`
}

// GetPrivateServers will get a page of a place's private servers that the authenticated user can join.
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_games__placeId__private_servers
//
// [GET] https://games.roblox.com/v1/games/{placeId}/private-servers
func (s *GamesService) GetPrivateServers(ctx context.Context, placeId string, opts *GetPrivateServersOptions, reqOpts ...goblox.RequestOpt) (*PrivateServersPage, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PrivateServersPage)
	resp, err := s.client.Call(ctx, GetPrivateServers.Compile(q, placeId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

type VIPServerGame struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	RootPlace RelatedEntity `json:"rootPlace"`
}

type VIPServerSubscription struct {
	Active               bool   `json:"active"`
	Expired              bool   `json:"expired"`
	ExpirationDate       string `json:"expirationDate"`
	Price                *int64 `json:"price"`
	CanRenew             bool   `json:"canRenew"`
	HasInsufficientFunds bool   `json:"hasInsufficientFunds"`
	HasRecurringProfile  bool   `json:"hasRecurringProfile"`
	HasPriceChanged      bool   `json:"hasPriceChanged"`
}

type VIPServerPermissions struct {
	ClanAllowed    bool          `json:"clanAllowed"`
	EnemyClanID    *int64        `json:"enemyClanId"`
	FriendsAllowed bool          `json:"friendsAllowed"`
	Users          []PartialUser `json:"users"`
}

type VIPServerVoiceSettings struct {
	Enabled bool `json:"enabled"`
}

type VIPServer struct {
	ID            int64                   `json:"id"`
	Name          string                  `json:"name"`
	Game          *VIPServerGame          `json:"game"`
	JoinCode      string                  `json:"joinCode"`
	Active        bool                    `json:"active"`
	Subscription  *VIPServerSubscription  `json:"subscription"`
	Permissions   *VIPServerPermissions   `json:"permissions"`
	VoiceSettings *VIPServerVoiceSettings `json:"voiceSettings"`
	Link          string                  `json:"link"`
}

type CreateVIPServerRequest struct {
	Name                string `json:"name"`
	ExpectedPrice       int64  `json:"expectedPrice"`
	IsPurchaseConfirmed bool   `json:"isPurchaseConfirmed"`
}

// CreateVIPServer will buy a private server for a game, charging the authenticated user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_post_v1_games_vip_servers__universeId_
//
// [POST] https://games.roblox.com/v1/games/vip-servers/{universeId}
func (s *GamesService) CreateVIPServer(ctx context.Context, universeId string, data CreateVIPServerRequest, reqOpts ...goblox.RequestOpt) (*VIPServer, *Response, error) {
	server := new(VIPServer)
	resp, err := s.client.Call(ctx, CreateVIPServer.Compile(nil, universeId), data, server, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return server, resp, nil
}

// GetVIPServer will get a private server's settings. The authenticated user must own it.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_get_v1_vip_servers__id_
//
// [GET] https://games.roblox.com/v1/vip-servers/{id}
func (s *GamesService) GetVIPServer(ctx context.Context, id string, reqOpts ...goblox.RequestOpt) (*VIPServer, *Response, error) {
	server := new(VIPServer)
	resp, err := s.client.Call(ctx, GetVIPServer.Compile(nil, id), nil, server, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return server, resp, nil
}

type UpdateVIPServerRequest struct {
	Name        *string `json:"name,omitempty"`
	NewJoinCode *bool   `json:"newJoinCode,omitempty"` // true generates a new join link.
	Active      *bool   `json:"active,omitempty"`
}

// UpdateVIPServer will update a private server's name, join link, or active state.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_patch_v1_vip_servers__id_
//
// [PATCH] https://games.roblox.com/v1/vip-servers/{id}
func (s *GamesService) UpdateVIPServer(ctx context.Context, id string, data UpdateVIPServerRequest, reqOpts ...goblox.RequestOpt) (*VIPServer, *Response, error) {
	server := new(VIPServer)
	resp, err := s.client.Call(ctx, UpdateVIPServer.Compile(nil, id), data, server, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return server, resp, nil
}

type UpdateVIPServerSubscriptionRequest struct {
	Active *bool  `json:"active,omitempty"`
	Price  *int64 `json:"price,omitempty"`
}

// UpdateVIPServerSubscription will update whether a private server's subscription renews.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_patch_v1_vip_servers__id__subscription
//
// [PATCH] https://games.roblox.com/v1/vip-servers/{id}/subscription
func (s *GamesService) UpdateVIPServerSubscription(ctx context.Context, id string, data UpdateVIPServerSubscriptionRequest, reqOpts ...goblox.RequestOpt) (*VIPServerSubscription, *Response, error) {
	subscription := new(VIPServerSubscription)
	resp, err := s.client.Call(ctx, UpdateVIPServerSubscription.Compile(nil, id), data, subscription, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return subscription, resp, nil
}

type UpdateVIPServerPermissionsRequest struct {
	ClanAllowed    *bool   `json:"clanAllowed,omitempty"`
	EnemyClanID    *int64  `json:"enemyClanId,omitempty"`
	FriendsAllowed *bool   `json:"friendsAllowed,omitempty"`
	UsersToAdd     []int64 `json:"usersToAdd,omitempty"`
	UsersToRemove  []int64 `json:"usersToRemove,omitempty"`
}

// UpdateVIPServerPermissions will update who is allowed to join a private server.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/games#games_patch_v1_vip_servers__id__permissions
//
// [PATCH] https://games.roblox.com/v1/vip-servers/{id}/permissions
func (s *GamesService) UpdateVIPServerPermissions(ctx context.Context, id string, data UpdateVIPServerPermissionsRequest, reqOpts ...goblox.RequestOpt) (*VIPServerPermissions, *Response, error) {
	permissions := new(VIPServerPermissions)
	resp, err := s.client.Call(ctx, UpdateVIPServerPermissions.Compile(nil, id), data, permissions, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return permissions, resp, nil
}
