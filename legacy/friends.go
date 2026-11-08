package legacy

import (
	"context"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	GetFriendsMetadata = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/metadata",
		auth.KindCookie, auth.KindNone,
	).WithDomain("friends")

	CheckQRSession = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/my/friends/{userId}/check-qr-session",
		auth.KindCookie,
	).WithDomain("friends")

	GetMyFriendsCount = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/my/friends/count",
		auth.KindCookie,
	).WithDomain("friends")

	GetFriendRequests = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/my/friends/requests",
		auth.KindCookie,
	).WithDomain("friends")

	GetNewFriendRequestsCount = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/my/new-friend-requests/count",
		auth.KindCookie,
	).WithDomain("friends")

	GetMyTrustedFriendsCount = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/my/trusted-friends/count",
		auth.KindCookie,
	).WithDomain("friends")

	GetTrustedFriendRequests = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/my/trusted-friends/requests",
		auth.KindCookie,
	).WithDomain("friends")

	GetPendingFriendRequestsCount = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/user/friend-requests/count",
		auth.KindCookie,
	).WithDomain("friends")

	GetPendingTrustedFriendRequestsCount = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/user/trusted-friend-requests/count",
		auth.KindCookie,
	).WithDomain("friends")

	GetFollowers = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{targetUserId}/followers",
		auth.KindCookie, auth.KindNone,
	).WithDomain("friends")

	GetFollowersCount = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{targetUserId}/followers/count",
		auth.KindCookie, auth.KindNone,
	).WithDomain("friends")

	GetFollowings = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{targetUserId}/followings",
		auth.KindCookie, auth.KindNone,
	).WithDomain("friends")

	GetFollowingsCount = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{targetUserId}/followings/count",
		auth.KindCookie, auth.KindNone,
	).WithDomain("friends")

	GetFriends = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/friends",
		auth.KindCookie, auth.KindNone,
	).WithDomain("friends")

	GetFriendsCount = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/friends/count",
		auth.KindCookie, auth.KindNone,
	).WithDomain("friends")

	FindFriends = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/friends/find",
		auth.KindCookie, auth.KindNone,
	).WithDomain("friends")

	GetInactiveFriends = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/friends/inactive",
		auth.KindCookie,
	).WithDomain("friends")

	GetOnlineFriends = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/friends/online",
		auth.KindCookie,
	).WithDomain("friends")

	SearchFriends = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/friends/search",
		auth.KindCookie,
	).WithDomain("friends")

	GetFriendStatuses = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/friends/statuses",
		auth.KindCookie,
	).WithDomain("friends")

	GetTrustedFriendsCount = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/trusted-friends/count",
		auth.KindCookie,
	).WithDomain("friends")

	RefreshQRSession = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/my/friends/refresh-qr-session",
		auth.KindCookie,
	).WithDomain("friends")

	MultiGetAreFriends = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/user/{userId}/multiget-are-friends",
		auth.KindCookie,
	).WithDomain("friends")

	FollowingExists = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/user/following-exists",
		auth.KindCookie,
	).WithDomain("friends")

	FollowUser = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/users/{targetUserId}/follow",
		auth.KindCookie,
	).WithDomain("friends")

	UnfollowUser = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/users/{targetUserId}/unfollow",
		auth.KindCookie,
	).WithDomain("friends")
)

type FriendsService service

// getCount calls an endpoint that returns {"count": n}.
func (s *FriendsService) getCount(ctx context.Context, e *goblox.CompiledEndpoint, opts []goblox.RequestOpt) (int64, *Response, error) {
	result := new(countResponse)
	resp, err := s.client.Call(ctx, e, nil, result, opts...)
	if err != nil {
		return 0, resp, err
	}

	return result.Count, resp, nil
}

type GetFriendsMetadataOptions struct {
	TargetUserID *int64 `url:"targetUserId,omitempty"`
}

type FriendsMetadata struct {
	IsFriendsFilterBarEnabled          bool   `json:"isFriendsFilterBarEnabled"`
	IsFriendsPageSortExperimentEnabled bool   `json:"isFriendsPageSortExperimentEnabled"`
	IsFriendsUserDataStoreCacheEnabled bool   `json:"isFriendsUserDataStoreCacheEnabled"`
	FrequentFriendSortRollout          int    `json:"frequentFriendSortRollout"`
	UserName                           string `json:"userName"`
	DisplayName                        string `json:"displayName"`
}

// GetFriendsMetadata will get the friends page metadata and feature flags.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_metadata
//
// [GET] https://friends.roblox.com/v1/metadata
func (s *FriendsService) GetFriendsMetadata(ctx context.Context, opts *GetFriendsMetadataOptions, reqOpts ...goblox.RequestOpt) (*FriendsMetadata, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	metadata := new(FriendsMetadata)
	resp, err := s.client.Call(ctx, GetFriendsMetadata.Compile(q), nil, metadata, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return metadata, resp, nil
}

// CheckQRSession will check whether a QR friending session is active for the user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_my_friends__userId__check_qr_session
//
// [GET] https://friends.roblox.com/v1/my/friends/{userId}/check-qr-session
func (s *FriendsService) CheckQRSession(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (bool, *Response, error) {
	var active bool
	resp, err := s.client.Call(ctx, CheckQRSession.Compile(nil, userId), nil, &active, reqOpts...)
	if err != nil {
		return false, resp, err
	}

	return active, resp, nil
}

// RefreshQRSession will refresh the authenticated user's QR friending session.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_post_v1_my_friends_refresh_qr_session
//
// [POST] https://friends.roblox.com/v1/my/friends/refresh-qr-session
func (s *FriendsService) RefreshQRSession(ctx context.Context, reqOpts ...goblox.RequestOpt) (bool, *Response, error) {
	var result struct {
		Success bool `json:"Success"`
	}
	resp, err := s.client.Call(ctx, RefreshQRSession.Compile(nil), nil, &result, reqOpts...)
	if err != nil {
		return false, resp, err
	}

	return result.Success, resp, nil
}

// GetMyFriendsCount will get the authenticated user's friend count.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_my_friends_count
//
// [GET] https://friends.roblox.com/v1/my/friends/count
func (s *FriendsService) GetMyFriendsCount(ctx context.Context, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	return s.getCount(ctx, GetMyFriendsCount.Compile(nil), reqOpts)
}

type GetFriendRequestsOptions struct {
	Limit             *int    `url:"limit,omitempty"`
	Cursor            *string `url:"cursor,omitempty"`
	SessionID         *string `url:"sessionId,omitempty"`
	FriendRequestSort *int    `url:"friendRequestSort,omitempty"` // 0, 1, or 2.
}

type FriendRequestInfo struct {
	SentAt           string `json:"sentAt"`
	SenderID         int64  `json:"senderId"`
	SourceUniverseID *int64 `json:"sourceUniverseId"`
	OriginSourceType Enum   `json:"originSourceType"`
	ContactName      string `json:"contactName"`
	SenderNickname   string `json:"senderNickname"`
}

type FriendRequest struct {
	FriendRequest          FriendRequestInfo `json:"friendRequest"`
	MutualFriendsList      []string          `json:"mutualFriendsList"`
	ID                     int64             `json:"id"`
	Name                   string            `json:"name"`
	DisplayName            string            `json:"displayName"`
	Description            string            `json:"description"`
	Created                string            `json:"created"`
	IsBanned               bool              `json:"isBanned"`
	ExternalAppDisplayName string            `json:"externalAppDisplayName"`
	HasVerifiedBadge       bool              `json:"hasVerifiedBadge"`
}

// GetFriendRequests will get a page of the authenticated user's incoming friend requests.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_my_friends_requests
//
// [GET] https://friends.roblox.com/v1/my/friends/requests
func (s *FriendsService) GetFriendRequests(ctx context.Context, opts *GetFriendRequestsOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[FriendRequest], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[FriendRequest])
	resp, err := s.client.Call(ctx, GetFriendRequests.Compile(q), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

// GetNewFriendRequestsCount will get the number of new friend requests for the authenticated user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_my_new_friend_requests_count
//
// [GET] https://friends.roblox.com/v1/my/new-friend-requests/count
func (s *FriendsService) GetNewFriendRequestsCount(ctx context.Context, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	return s.getCount(ctx, GetNewFriendRequestsCount.Compile(nil), reqOpts)
}

// GetMyTrustedFriendsCount will get the authenticated user's trusted friend count.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_my_trusted_friends_count
//
// [GET] https://friends.roblox.com/v1/my/trusted-friends/count
func (s *FriendsService) GetMyTrustedFriendsCount(ctx context.Context, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	return s.getCount(ctx, GetMyTrustedFriendsCount.Compile(nil), reqOpts)
}

type GetTrustedFriendRequestsOptions struct {
	Limit     *int    `url:"limit,omitempty"`
	Cursor    *string `url:"cursor,omitempty"`
	SortOrder *int    `url:"sortOrder,omitempty"` // 0 or 1.
}

type TrustedFriendRequest struct {
	SentAt   string `json:"sentAt"`
	SenderID int64  `json:"senderId"`
}

// GetTrustedFriendRequests will get a page of the authenticated user's incoming trusted friend requests.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_my_trusted_friends_requests
//
// [GET] https://friends.roblox.com/v1/my/trusted-friends/requests
func (s *FriendsService) GetTrustedFriendRequests(ctx context.Context, opts *GetTrustedFriendRequestsOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[TrustedFriendRequest], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[TrustedFriendRequest])
	resp, err := s.client.Call(ctx, GetTrustedFriendRequests.Compile(q), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

// GetPendingFriendRequestsCount will get the number of pending friend requests for the authenticated user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_user_friend_requests_count
//
// [GET] https://friends.roblox.com/v1/user/friend-requests/count
func (s *FriendsService) GetPendingFriendRequestsCount(ctx context.Context, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	return s.getCount(ctx, GetPendingFriendRequestsCount.Compile(nil), reqOpts)
}

// GetPendingTrustedFriendRequestsCount will get the number of pending trusted friend requests for the
// authenticated user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_user_trusted_friend_requests_count
//
// [GET] https://friends.roblox.com/v1/user/trusted-friend-requests/count
func (s *FriendsService) GetPendingTrustedFriendRequestsCount(ctx context.Context, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	return s.getCount(ctx, GetPendingTrustedFriendRequestsCount.Compile(nil), reqOpts)
}

// FriendUser is a user returned by the friends and followers endpoints.
type FriendUser struct {
	ID                     int64  `json:"id"`
	Name                   string `json:"name"`
	DisplayName            string `json:"displayName"`
	Description            string `json:"description"`
	Created                string `json:"created"`
	IsOnline               bool   `json:"isOnline"`
	PresenceType           Enum   `json:"presenceType"`
	IsDeleted              bool   `json:"isDeleted"`
	IsBanned               bool   `json:"isBanned"`
	FriendFrequentScore    int64  `json:"friendFrequentScore"`
	FriendFrequentRank     int64  `json:"friendFrequentRank"`
	ExternalAppDisplayName string `json:"externalAppDisplayName"`
	HasVerifiedBadge       bool   `json:"hasVerifiedBadge"`
}

// GetFollowers will get a page of the users following a user.
// Limit can be 10, 18, 25, 50, or 100.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__targetUserId__followers
//
// [GET] https://friends.roblox.com/v1/users/{targetUserId}/followers
func (s *FriendsService) GetFollowers(ctx context.Context, targetUserId string, opts *PageOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[FriendUser], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[FriendUser])
	resp, err := s.client.Call(ctx, GetFollowers.Compile(q, targetUserId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

// GetFollowersCount will get the number of users following a user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__targetUserId__followers_count
//
// [GET] https://friends.roblox.com/v1/users/{targetUserId}/followers/count
func (s *FriendsService) GetFollowersCount(ctx context.Context, targetUserId string, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	return s.getCount(ctx, GetFollowersCount.Compile(nil, targetUserId), reqOpts)
}

// GetFollowings will get a page of the users that a user follows.
// Limit can be 10, 18, 25, 50, or 100.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__targetUserId__followings
//
// [GET] https://friends.roblox.com/v1/users/{targetUserId}/followings
func (s *FriendsService) GetFollowings(ctx context.Context, targetUserId string, opts *PageOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[FriendUser], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[FriendUser])
	resp, err := s.client.Call(ctx, GetFollowings.Compile(q, targetUserId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

// GetFollowingsCount will get the number of users that a user follows.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__targetUserId__followings_count
//
// [GET] https://friends.roblox.com/v1/users/{targetUserId}/followings/count
func (s *FriendsService) GetFollowingsCount(ctx context.Context, targetUserId string, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	return s.getCount(ctx, GetFollowingsCount.Compile(nil, targetUserId), reqOpts)
}

type GetFriendsOptions struct {
	UserSort *int `url:"userSort,omitempty"` // 0, 1, or 2.
}

// GetFriends will get all of a user's friends.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__userId__friends
//
// [GET] https://friends.roblox.com/v1/users/{userId}/friends
func (s *FriendsService) GetFriends(ctx context.Context, userId string, opts *GetFriendsOptions, reqOpts ...goblox.RequestOpt) ([]FriendUser, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	results := new(WrappedResponse[[]FriendUser])
	resp, err := s.client.Call(ctx, GetFriends.Compile(q, userId), nil, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}

// GetFriendsCount will get a user's friend count.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__userId__friends_count
//
// [GET] https://friends.roblox.com/v1/users/{userId}/friends/count
func (s *FriendsService) GetFriendsCount(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	return s.getCount(ctx, GetFriendsCount.Compile(nil, userId), reqOpts)
}

// Friend is a minimal friend returned by the cursored friend endpoints.
type Friend struct {
	ID               int64 `json:"id"`
	HasVerifiedBadge bool  `json:"hasVerifiedBadge"`
}

// FriendsPage is a page of friends. Unlike most endpoints, these use PascalCase keys
// and a HasMore flag.
type FriendsPage struct {
	PreviousCursor string   `json:"PreviousCursor"`
	PageItems      []Friend `json:"PageItems"`
	NextCursor     string   `json:"NextCursor"`
	HasMore        bool     `json:"HasMore"`
}

type FindFriendsOptions struct {
	UserSort        *int    `url:"userSort,omitempty"` // 0, 1, or 2.
	Cursor          *string `url:"cursor,omitempty"`
	Limit           *int    `url:"limit,omitempty"`
	FindFriendsType *int    `url:"findFriendsType,omitempty"` // 0 or 1.
}

// FindFriends will get a page of a user's friends.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__userId__friends_find
//
// [GET] https://friends.roblox.com/v1/users/{userId}/friends/find
func (s *FriendsService) FindFriends(ctx context.Context, userId string, opts *FindFriendsOptions, reqOpts ...goblox.RequestOpt) (*FriendsPage, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(FriendsPage)
	resp, err := s.client.Call(ctx, FindFriends.Compile(q, userId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

// GetInactiveFriends will get a user's inactive friends.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__userId__friends_inactive
//
// [GET] https://friends.roblox.com/v1/users/{userId}/friends/inactive
func (s *FriendsService) GetInactiveFriends(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) ([]FriendUser, *Response, error) {
	results := new(WrappedResponse[[]FriendUser])
	resp, err := s.client.Call(ctx, GetInactiveFriends.Compile(nil, userId), nil, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}

type FriendPresence struct {
	UserPresenceType Enum    `json:"UserPresenceType"`
	UserLocationType Enum    `json:"UserLocationType"`
	LastLocation     string  `json:"lastLocation"`
	PlaceID          *int64  `json:"placeId"`
	RootPlaceID      *int64  `json:"rootPlaceId"`
	GameInstanceID   *string `json:"gameInstanceId"`
	UniverseID       *int64  `json:"universeId"`
	LastOnline       string  `json:"lastOnline"`
}

type OnlineFriend struct {
	ID           int64          `json:"id"`
	Name         string         `json:"name"`
	DisplayName  string         `json:"displayName"`
	UserPresence FriendPresence `json:"userPresence"`
	SortScore    float64        `json:"sortScore"`
}

// GetOnlineFriends will get a user's online friends.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__userId__friends_online
//
// [GET] https://friends.roblox.com/v1/users/{userId}/friends/online
func (s *FriendsService) GetOnlineFriends(ctx context.Context, userId string, opts *GetFriendsOptions, reqOpts ...goblox.RequestOpt) ([]OnlineFriend, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	results := new(WrappedResponse[[]OnlineFriend])
	resp, err := s.client.Call(ctx, GetOnlineFriends.Compile(q, userId), nil, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}

type SearchFriendsOptions struct {
	Query  *string `url:"query,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
	Limit  *int    `url:"limit,omitempty"`
}

// SearchFriends will search a user's friends by name.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__userId__friends_search
//
// [GET] https://friends.roblox.com/v1/users/{userId}/friends/search
func (s *FriendsService) SearchFriends(ctx context.Context, userId string, opts *SearchFriendsOptions, reqOpts ...goblox.RequestOpt) (*FriendsPage, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(FriendsPage)
	resp, err := s.client.Call(ctx, SearchFriends.Compile(q, userId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

type GetFriendStatusesOptions struct {
	UserIDs []int64 `url:"userIds,comma"` // Required.
}

type FriendStatus struct {
	ID     int64 `json:"id"`
	Status Enum  `json:"status"`
}

// GetFriendStatuses will get the friendship status between a user and a list of other users.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__userId__friends_statuses
//
// [GET] https://friends.roblox.com/v1/users/{userId}/friends/statuses
func (s *FriendsService) GetFriendStatuses(ctx context.Context, userId string, opts *GetFriendStatusesOptions, reqOpts ...goblox.RequestOpt) ([]FriendStatus, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	results := new(WrappedResponse[[]FriendStatus])
	resp, err := s.client.Call(ctx, GetFriendStatuses.Compile(q, userId), nil, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}

// GetTrustedFriendsCount will get a user's trusted friend count.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_get_v1_users__userId__trusted_friends_count
//
// [GET] https://friends.roblox.com/v1/users/{userId}/trusted-friends/count
func (s *FriendsService) GetTrustedFriendsCount(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	return s.getCount(ctx, GetTrustedFriendsCount.Compile(nil, userId), reqOpts)
}

type targetUserIDsBody struct {
	TargetUserIDs []int64 `json:"targetUserIds"`
}

// MultiGetAreFriends will check which of the target users are friends with the requesting user, and
// return the IDs of those that are.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_post_v1_user__userId__multiget_are_friends
//
// [POST] https://friends.roblox.com/v1/user/{userId}/multiget-are-friends
func (s *FriendsService) MultiGetAreFriends(ctx context.Context, userId string, targetUserIds []int64, reqOpts ...goblox.RequestOpt) ([]int64, *Response, error) {
	var result struct {
		FriendsID []int64 `json:"friendsId"`
	}
	resp, err := s.client.Call(ctx, MultiGetAreFriends.Compile(nil, userId), targetUserIDsBody{TargetUserIDs: targetUserIds}, &result, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return result.FriendsID, resp, nil
}

type FollowingStatus struct {
	UserID      int64 `json:"userId"`
	IsFollowing bool  `json:"isFollowing"`
	IsFollowed  bool  `json:"isFollowed"`
}

// FollowingExists will check whether the authenticated user follows, and is followed by, each target user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_post_v1_user_following_exists
//
// [POST] https://friends.roblox.com/v1/user/following-exists
func (s *FriendsService) FollowingExists(ctx context.Context, targetUserIds []int64, reqOpts ...goblox.RequestOpt) ([]FollowingStatus, *Response, error) {
	var result struct {
		Followings []FollowingStatus `json:"followings"`
	}
	resp, err := s.client.Call(ctx, FollowingExists.Compile(nil), targetUserIDsBody{TargetUserIDs: targetUserIds}, &result, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return result.Followings, resp, nil
}

// CaptchaToken is a solved captcha, for endpoints that may ask for one.
type CaptchaToken struct {
	CaptchaID       string `json:"captchaId,omitempty"`
	CaptchaToken    string `json:"captchaToken,omitempty"`
	CaptchaProvider string `json:"captchaProvider,omitempty"`
	ChallengeID     string `json:"challengeId,omitempty"`
}

type CaptchaStatus struct {
	Success           bool `json:"success"`
	IsCaptchaRequired bool `json:"isCaptchaRequired"`
}

// FollowUser will make the authenticated user follow the target user.
// Roblox may ask for a captcha, which is reported in the returned CaptchaStatus.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_post_v1_users__targetUserId__follow
//
// [POST] https://friends.roblox.com/v1/users/{targetUserId}/follow
func (s *FriendsService) FollowUser(ctx context.Context, targetUserId string, captcha *CaptchaToken, reqOpts ...goblox.RequestOpt) (*CaptchaStatus, *Response, error) {
	var body any = struct{}{}
	if captcha != nil {
		body = captcha
	}

	status := new(CaptchaStatus)
	resp, err := s.client.Call(ctx, FollowUser.Compile(nil, targetUserId), body, status, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}

// UnfollowUser will make the authenticated user unfollow the target user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/friends#friends_post_v1_users__targetUserId__unfollow
//
// [POST] https://friends.roblox.com/v1/users/{targetUserId}/unfollow
func (s *FriendsService) UnfollowUser(ctx context.Context, targetUserId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, UnfollowUser.Compile(nil, targetUserId), nil, nil, reqOpts...)
}
