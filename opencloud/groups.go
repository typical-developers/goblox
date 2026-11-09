package opencloud

import (
	"context"
	"net/http"

	"github.com/luckfire/goblox"
	"github.com/luckfire/goblox/auth"
)

var (
	GetGroup = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/groups/{groupId}",
		auth.KindAPIKey, auth.KindOAuth2,
	)

	ListGroupJoinRequests = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/groups/{groupId}/join-requests",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group:read")

	AcceptGroupJoinRequest = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/groups/{groupId}/join-requests/{userId}:accept",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group:write")

	DeclineGroupJoinRequest = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/groups/{groupId}/join-requests/{userId}:decline",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group:write")

	ListGroupMemberships = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/groups/{groupId}/memberships",
		auth.KindAPIKey, auth.KindOAuth2,
	)

	UpdateGroupMemberships = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/groups/{groupId}/memberships/{userId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group:write")

	ListGroupRoles = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/groups/{groupId}/roles",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group:read")

	GetGroupRoles = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/groups/{groupId}/roles/{roleId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group:read")

	GetGroupShout = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/groups/{groupId}/shout",
		auth.KindAPIKey, auth.KindOAuth2,
	)
)

// GroupsService handles communication with group endpoints.
type GroupsService service

type Group struct {
	Path               string `json:"path"`
	CreateTime         string `json:"createTime"`
	UpdateTime         string `json:"updateTime"`
	ID                 string `json:"id"`
	DisplayName        string `json:"displayName"`
	Description        string `json:"description"`
	Owner              string `json:"owner"`
	MemberCount        int    `json:"memberCount"`
	PublicEntryAllowed bool   `json:"publicEntryAllowed"`
	Locked             bool   `json:"locked"`
	Verified           bool   `json:"verified"`
}

// GetGroup will fetch information on a specificed group.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Group#Cloud_GetGroup
//
// [GET] /cloud/v2/groups/{group_id}
func (s *GroupsService) GetGroup(ctx context.Context, groupId string, reqOpts ...goblox.RequestOpt) (*Group, *Response, error) {
	group := new(Group)
	resp, err := s.client.Call(ctx, GetGroup.Compile(nil, groupId), nil, group, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

type GroupJoinRequest struct {
	Path       string `json:"path"`
	CreateTime string `json:"createTime"`
	User       string `json:"user"`
}

type GroupJoinRequests struct {
	GroupJoinRequests []GroupJoinRequest `json:"groupJoinRequests"`
	NextPageToken     string             `json:"nextPageToken"`
}

// ListGroupJoinRequests will fetch join requests for a specificed group.
//
// Required scopes: group:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/GroupJoinRequest#Cloud_ListGroupJoinRequests
//
// [GET] /cloud/v2/groups/{group_id}/join-requests
func (s *GroupsService) ListGroupJoinRequests(ctx context.Context, groupId string, opts *OptionsWithFilter, reqOpts ...goblox.RequestOpt) (*GroupJoinRequests, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	groupJoinRequests := new(GroupJoinRequests)
	resp, err := s.client.Call(ctx, ListGroupJoinRequests.Compile(q, groupId), nil, groupJoinRequests, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupJoinRequests, resp, nil
}

// AcceptGroupJoinRequest will accept the join request for a user under a specificed group.
//
// Required scopes: group:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/GroupJoinRequest#Cloud_AcceptGroupJoinRequest
//
// [POST] /cloud/v2/groups/{group_id}/join-requests/{user_id}:accept
func (s *GroupsService) AcceptGroupJoinRequest(ctx context.Context, groupId, userId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	resp, err := s.client.Call(ctx, AcceptGroupJoinRequest.Compile(nil, groupId, userId), nil, nil, reqOpts...)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeclineGroupJoinRequest will decline the join request for a user under a specificed group.
//
// Required scopes: group:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/GroupJoinRequest#Cloud_DeclineGroupJoinRequest
//
// [POST] /cloud/v2/groups/{group_id}/join-requests/{user_id}:decline
func (s *GroupsService) DeclineGroupJoinRequest(ctx context.Context, groupId, userId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	resp, err := s.client.Call(ctx, DeclineGroupJoinRequest.Compile(nil, groupId, userId), nil, nil, reqOpts...)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

type GroupMembership struct {
	Path       string `json:"path"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	User       string `json:"user"`
	Role       string `json:"role"`
}

type GroupMemberships struct {
	GroupMemberships []GroupMembership `json:"groupMemberships"`
	NextPageToken    string            `json:"nextPageToken"`
}

// ListGroupMemberships will fetch memberships for a specificed group.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/GroupMembership#Cloud_ListGroupMemberships
//
// [GET] /cloud/v2/groups/{group_id}/memberships
func (s *GroupsService) ListGroupMemberships(ctx context.Context, groupId string, opts *OptionsWithFilter, reqOpts ...goblox.RequestOpt) (*GroupMemberships, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	groupMemberships := new(GroupMemberships)
	resp, err := s.client.Call(ctx, ListGroupMemberships.Compile(q, groupId), nil, groupMemberships, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupMemberships, resp, nil
}

type GroupMembershipUpdate struct {
	Role *string `json:"role,omitempty"`
}

// UpdateGroupMemberships will update memberships for a user under a specificed group.
//
// Required scopes: group:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/GroupMembership#Cloud_UpdateGroupMembership
//
// [PATCH] /cloud/v2/groups/{group_id}/memberships/{user_id}
func (s *GroupsService) UpdateGroupMemberships(ctx context.Context, groupId, userId string, data GroupMembershipUpdate, reqOpts ...goblox.RequestOpt) (*GroupMembership, *Response, error) {
	groupMembership := new(GroupMembership)
	resp, err := s.client.Call(ctx, UpdateGroupMemberships.Compile(nil, groupId, userId), data, groupMembership, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupMembership, resp, nil
}

type GroupRolePermissions struct {
	ViewWallPosts         bool `json:"viewWallPosts"`
	CreateWallPosts       bool `json:"createWallPosts"`
	DeleteWallPosts       bool `json:"deleteWallPosts"`
	ViewGroupShout        bool `json:"viewGroupShout"`
	CreateGroupShout      bool `json:"createGroupShout"`
	ChangeRank            bool `json:"changeRank"`
	AcceptRequests        bool `json:"acceptRequests"`
	ExileMembers          bool `json:"exileMembers"`
	ManageRelationships   bool `json:"manageRelationships"`
	ViewAuditLog          bool `json:"viewAuditLog"`
	SpendGroupFunds       bool `json:"spendGroupFunds"`
	AdvertiseGroup        bool `json:"advertiseGroup"`
	CreateAvatarItems     bool `json:"createAvatarItems"`
	ManageAvatarItems     bool `json:"manageAvatarItems"`
	ManageGroupUniverses  bool `json:"manageGroupUniverses"`
	ViewUniverseAnalytics bool `json:"viewUniverseAnalytics"`
	CreateAPIKeys         bool `json:"createApiKeys"`
	ManageAPIKeys         bool `json:"manageApiKeys"`
	BanMembers            bool `json:"banMembers"`
	ViewForums            bool `json:"viewForums"`
	ManageCategories      bool `json:"manageCategories"`
	CreatePosts           bool `json:"createPosts"`
	LockPosts             bool `json:"lockPosts"`
	PinPosts              bool `json:"pinPosts"`
	RemovePosts           bool `json:"removePosts"`
	CreateComments        bool `json:"createComments"`
	RemoveComments        bool `json:"removeComments"`
}

type GroupRole struct {
	Path        string               `json:"path"`
	CreateTime  string               `json:"createTime"`
	UpdateTime  string               `json:"updateTime"`
	ID          string               `json:"id"`
	DisplayName string               `json:"displayName"`
	Description string               `json:"description"`
	Rank        int                  `json:"rank"`
	MemberCount int                  `json:"memberCount"`
	Permissions GroupRolePermissions `json:"permissions"`
}

type GroupRoles struct {
	GroupRoles    []GroupRole `json:"groupRoles"`
	NextPageToken string      `json:"nextPageToken"`
}

// ListGroupRoles will fetch roles for a specificed group.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/GroupRole#Cloud_ListGroupRoles
//
// [GET] /cloud/v2/groups/{group_id}/roles
func (s *GroupsService) ListGroupRoles(ctx context.Context, groupId string, opts *OptionsWithFilter, reqOpts ...goblox.RequestOpt) (*GroupRoles, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	groupRoles := new(GroupRoles)
	resp, err := s.client.Call(ctx, ListGroupRoles.Compile(q, groupId), nil, groupRoles, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupRoles, resp, nil
}

// GetGroupRoles will fetch a specificed role for a specificed group.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/GroupRole#Cloud_GetGroupRole
//
// [GET] /cloud/v2/groups/{group_id}/roles/{role_id}
func (s *GroupsService) GetGroupRoles(ctx context.Context, groupId, roleId string, reqOpts ...goblox.RequestOpt) (*GroupRole, *Response, error) {
	groupRole := new(GroupRole)
	resp, err := s.client.Call(ctx, GetGroupRoles.Compile(nil, groupId, roleId), nil, groupRole, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupRole, resp, nil
}

type GroupShout struct {
	Path       string `json:"path"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	Content    string `json:"content"`
	Poster     string `json:"poster"`
}

// GetGroupShout will fetch the shout for a specificed group.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/GroupShout#Cloud_GetGroupShout
//
// [GET] /cloud/v2/groups/{group_id}/shout
func (s *GroupsService) GetGroupShout(ctx context.Context, groupId string, reqOpts ...goblox.RequestOpt) (*GroupShout, *Response, error) {
	groupShout := new(GroupShout)
	resp, err := s.client.Call(ctx, GetGroupShout.Compile(nil, groupId), nil, groupShout, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupShout, resp, nil
}

var (
	ListGroupForumCategories = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/groups/{groupId}/forum-categories",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group-forum:read")

	ListGroupForumPosts = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/groups/{groupId}/forum-categories/{forumCategoryId}/posts",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group-forum:read")

	ListGroupForumComments = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/groups/{groupId}/forum-categories/{forumCategoryId}/posts/{postId}/comments",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group-forum:read")

	AssignGroupMembershipRole = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/groups/{groupId}/memberships/{membershipId}:assignRole",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group:write")

	UnassignGroupMembershipRole = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/groups/{groupId}/memberships/{membershipId}:unassignRole",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("group:write")

	ListManageableGroups = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-develop/v1/user/groups/canmanage",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-group:manage")

	GetGroupPolicies = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-groups/v1/groups/policies",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-group:manage")

	GetGroupAuditLog = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-groups/v1/groups/{groupId}/audit-log",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-group:manage")

	UpdateGroupDescription = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-groups/v1/groups/{groupId}/description",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-group:manage")

	UpdateGroupNotificationPreference = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-groups/v1/groups/{groupId}/notification-preference",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-user:manage")

	GetGroupSettings = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-groups/v1/groups/{groupId}/settings",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-group:manage")

	UpdateGroupSettings = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-groups/v1/groups/{groupId}/settings",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-group:manage")

	UpdateGroupShout = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-groups/v1/groups/{groupId}/status",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-group:manage")

	ListPendingGroupJoinRequests = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-groups/v1/user/groups/pending",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-group:manage")
)

type GroupForumCategory struct {
	Path                 string `json:"path"`
	CreateTime           string `json:"createTime"`
	UpdateTime           string `json:"updateTime"`
	GroupForumCategoryID string `json:"groupForumCategoryId"`
	DisplayName          string `json:"displayName"`
	Creator              string `json:"creator"`
	ArchiveTime          string `json:"archiveTime"`
	Archiver             string `json:"archiver"`
}

type ListGroupForumCategoriesResponse struct {
	GroupForumCategories []GroupForumCategory `json:"groupForumCategories"`
	NextPageToken        string               `json:"nextPageToken"`
}

// ListGroupForumCategories will list forum categories in the group.
//
// Required scopes: group-forum:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#Cloud_ListGroupForumCategories
//
// [GET] /cloud/v2/groups/{group_id}/forum-categories
func (s *GroupsService) ListGroupForumCategories(ctx context.Context, groupId string, opts *OptionsWithFilter, reqOpts ...goblox.RequestOpt) (*ListGroupForumCategoriesResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	listGroupForumCategoriesResponse := new(ListGroupForumCategoriesResponse)
	resp, err := s.client.Call(ctx, ListGroupForumCategories.Compile(q, groupId), nil, listGroupForumCategoriesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listGroupForumCategoriesResponse, resp, nil
}

type ListGroupForumPostsOptionsView string

const (
	ListGroupForumPostsOptionsViewViewUnspecified      ListGroupForumPostsOptionsView = "VIEW_UNSPECIFIED"
	ListGroupForumPostsOptionsViewFull                 ListGroupForumPostsOptionsView = "FULL"
	ListGroupForumPostsOptionsViewFullWithFirstComment ListGroupForumPostsOptionsView = "FULL_WITH_FIRST_COMMENT"
)

type ListGroupForumPostsOptions struct {
	MaxPageSize *int                            `url:"maxPageSize,omitempty"`
	PageToken   *string                         `url:"pageToken,omitempty"`
	Filter      *string                         `url:"filter,omitempty"`
	View        *ListGroupForumPostsOptionsView `url:"view,omitempty"`
}

type MessageReaction struct {
	EmoteID string `json:"emoteId"`
	Count   int64  `json:"count"`
}

type ContentMessage struct {
	CreateTime      string            `json:"createTime"`
	UpdateTime      string            `json:"updateTime"`
	PlainText       string            `json:"plainText"`
	Author          string            `json:"author"`
	MessageReaction []MessageReaction `json:"messageReaction"`
}

type GroupForumComment struct {
	Path                string         `json:"path"`
	GroupForumCommentID string         `json:"groupForumCommentId"`
	Message             ContentMessage `json:"message"`
}

type GroupForumPost struct {
	Path             string            `json:"path"`
	CreateTime       string            `json:"createTime"`
	UpdateTime       string            `json:"updateTime"`
	Pinned           bool              `json:"pinned"`
	Locked           bool              `json:"locked"`
	GroupForumPostID string            `json:"groupForumPostId"`
	Title            string            `json:"title"`
	FirstComment     GroupForumComment `json:"firstComment"`
	Author           string            `json:"author"`
	ArchiveTime      string            `json:"archiveTime"`
	Archiver         string            `json:"archiver"`
	CommentCount     int64             `json:"commentCount"`
}

type ListGroupForumPostsResponse struct {
	GroupForumPosts []GroupForumPost `json:"groupForumPosts"`
	NextPageToken   string           `json:"nextPageToken"`
}

// ListGroupForumPosts will list forum posts in the group's forum category.
//
// Required scopes: group-forum:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#Cloud_ListGroupForumPosts
//
// [GET] /cloud/v2/groups/{group_id}/forum-categories/{forum_category_id}/posts
func (s *GroupsService) ListGroupForumPosts(ctx context.Context, groupId, forumCategoryId string, opts *ListGroupForumPostsOptions, reqOpts ...goblox.RequestOpt) (*ListGroupForumPostsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	listGroupForumPostsResponse := new(ListGroupForumPostsResponse)
	resp, err := s.client.Call(ctx, ListGroupForumPosts.Compile(q, groupId, forumCategoryId), nil, listGroupForumPostsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listGroupForumPostsResponse, resp, nil
}

type ListGroupForumCommentsResponse struct {
	GroupForumComments []GroupForumComment `json:"groupForumComments"`
	NextPageToken      string              `json:"nextPageToken"`
}

// ListGroupForumComments will list forum comments on a group's forum post.
//
// Required scopes: group-forum:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#Cloud_ListGroupForumComments
//
// [GET] /cloud/v2/groups/{group_id}/forum-categories/{forum_category_id}/posts/{post_id}/comments
func (s *GroupsService) ListGroupForumComments(ctx context.Context, groupId, forumCategoryId, postId string, opts *OptionsWithFilter, reqOpts ...goblox.RequestOpt) (*ListGroupForumCommentsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	listGroupForumCommentsResponse := new(ListGroupForumCommentsResponse)
	resp, err := s.client.Call(ctx, ListGroupForumComments.Compile(q, groupId, forumCategoryId, postId), nil, listGroupForumCommentsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listGroupForumCommentsResponse, resp, nil
}

type AssignRoleGroupMembershipRequest struct {
	Role *string `json:"role,omitempty"`
}

// AssignGroupMembershipRole will assign a specific role to a user within a group.
//
// Required scopes: group:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#Cloud_AssignRoleGroupMembership
//
// [POST] /cloud/v2/groups/{group_id}/memberships/{membership_id}:assignRole
func (s *GroupsService) AssignGroupMembershipRole(ctx context.Context, groupId, membershipId string, data AssignRoleGroupMembershipRequest, reqOpts ...goblox.RequestOpt) (*GroupMembership, *Response, error) {
	groupMembership := new(GroupMembership)
	resp, err := s.client.Call(ctx, AssignGroupMembershipRole.Compile(nil, groupId, membershipId), data, groupMembership, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupMembership, resp, nil
}

type UnassignRoleGroupMembershipRequest struct {
	Role *string `json:"role,omitempty"`
}

// UnassignGroupMembershipRole will unassign a specific role from a user within a group.
//
// Required scopes: group:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#Cloud_UnassignRoleGroupMembership
//
// [POST] /cloud/v2/groups/{group_id}/memberships/{membership_id}:unassignRole
func (s *GroupsService) UnassignGroupMembershipRole(ctx context.Context, groupId, membershipId string, data UnassignRoleGroupMembershipRequest, reqOpts ...goblox.RequestOpt) (*GroupMembership, *Response, error) {
	groupMembership := new(GroupMembership)
	resp, err := s.client.Call(ctx, UnassignGroupMembershipRole.Compile(nil, groupId, membershipId), data, groupMembership, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupMembership, resp, nil
}

type ManageableGroup struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ManageableGroupList struct {
	Data []ManageableGroup `json:"data"`
}

// ListManageableGroups will get a list of Groups that a user can manage.
//
// Required scopes: legacy-group:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#get_legacy_develop_v1_user_groups_canmanage
//
// [GET] /legacy-develop/v1/user/groups/canmanage
func (s *GroupsService) ListManageableGroups(ctx context.Context, reqOpts ...goblox.RequestOpt) (*ManageableGroupList, *Response, error) {
	manageableGroupList := new(ManageableGroupList)
	resp, err := s.client.Call(ctx, ListManageableGroups.Compile(nil), nil, manageableGroupList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return manageableGroupList, resp, nil
}

type GroupPolicyRequest struct {
	GroupIDs []int64 `json:"groupIds,omitempty"`
}

type GroupPolicyResponse struct {
	CanViewGroup bool  `json:"canViewGroup"`
	GroupID      int64 `json:"groupId"`
}

type GroupPoliciesResponse struct {
	Groups []GroupPolicyResponse `json:"groups"`
}

// GetGroupPolicies will get group policy info used for compliance.
//
// Required scopes: legacy-group:manage
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#post_legacy_groups_v1_groups_policies
//
// [POST] /legacy-groups/v1/groups/policies
func (s *GroupsService) GetGroupPolicies(ctx context.Context, data GroupPolicyRequest, reqOpts ...goblox.RequestOpt) (*GroupPoliciesResponse, *Response, error) {
	groupPoliciesResponse := new(GroupPoliciesResponse)
	resp, err := s.client.Call(ctx, GetGroupPolicies.Compile(nil), data, groupPoliciesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupPoliciesResponse, resp, nil
}

type GetGroupAuditLogOptionsActionType string

const (
	GetGroupAuditLogOptionsActionTypeDeletePost                              GetGroupAuditLogOptionsActionType = "DeletePost"
	GetGroupAuditLogOptionsActionTypeRemoveMember                            GetGroupAuditLogOptionsActionType = "RemoveMember"
	GetGroupAuditLogOptionsActionTypeAcceptJoinRequest                       GetGroupAuditLogOptionsActionType = "AcceptJoinRequest"
	GetGroupAuditLogOptionsActionTypeDeclineJoinRequest                      GetGroupAuditLogOptionsActionType = "DeclineJoinRequest"
	GetGroupAuditLogOptionsActionTypePostStatus                              GetGroupAuditLogOptionsActionType = "PostStatus"
	GetGroupAuditLogOptionsActionTypeChangeRank                              GetGroupAuditLogOptionsActionType = "ChangeRank"
	GetGroupAuditLogOptionsActionTypeBuyAd                                   GetGroupAuditLogOptionsActionType = "BuyAd"
	GetGroupAuditLogOptionsActionTypeSendAllyRequest                         GetGroupAuditLogOptionsActionType = "SendAllyRequest"
	GetGroupAuditLogOptionsActionTypeCreateEnemy                             GetGroupAuditLogOptionsActionType = "CreateEnemy"
	GetGroupAuditLogOptionsActionTypeAcceptAllyRequest                       GetGroupAuditLogOptionsActionType = "AcceptAllyRequest"
	GetGroupAuditLogOptionsActionTypeDeclineAllyRequest                      GetGroupAuditLogOptionsActionType = "DeclineAllyRequest"
	GetGroupAuditLogOptionsActionTypeDeleteAlly                              GetGroupAuditLogOptionsActionType = "DeleteAlly"
	GetGroupAuditLogOptionsActionTypeDeleteEnemy                             GetGroupAuditLogOptionsActionType = "DeleteEnemy"
	GetGroupAuditLogOptionsActionTypeAddGroupPlace                           GetGroupAuditLogOptionsActionType = "AddGroupPlace"
	GetGroupAuditLogOptionsActionTypeRemoveGroupPlace                        GetGroupAuditLogOptionsActionType = "RemoveGroupPlace"
	GetGroupAuditLogOptionsActionTypeCreateItems                             GetGroupAuditLogOptionsActionType = "CreateItems"
	GetGroupAuditLogOptionsActionTypeConfigureItems                          GetGroupAuditLogOptionsActionType = "ConfigureItems"
	GetGroupAuditLogOptionsActionTypeSpendGroupFunds                         GetGroupAuditLogOptionsActionType = "SpendGroupFunds"
	GetGroupAuditLogOptionsActionTypeChangeOwner                             GetGroupAuditLogOptionsActionType = "ChangeOwner"
	GetGroupAuditLogOptionsActionTypeDelete                                  GetGroupAuditLogOptionsActionType = "Delete"
	GetGroupAuditLogOptionsActionTypeAdjustCurrencyAmounts                   GetGroupAuditLogOptionsActionType = "AdjustCurrencyAmounts"
	GetGroupAuditLogOptionsActionTypeAbandon                                 GetGroupAuditLogOptionsActionType = "Abandon"
	GetGroupAuditLogOptionsActionTypeClaim                                   GetGroupAuditLogOptionsActionType = "Claim"
	GetGroupAuditLogOptionsActionTypeRename                                  GetGroupAuditLogOptionsActionType = "Rename"
	GetGroupAuditLogOptionsActionTypeChangeDescription                       GetGroupAuditLogOptionsActionType = "ChangeDescription"
	GetGroupAuditLogOptionsActionTypeInviteToClan                            GetGroupAuditLogOptionsActionType = "InviteToClan"
	GetGroupAuditLogOptionsActionTypeKickFromClan                            GetGroupAuditLogOptionsActionType = "KickFromClan"
	GetGroupAuditLogOptionsActionTypeCancelClanInvite                        GetGroupAuditLogOptionsActionType = "CancelClanInvite"
	GetGroupAuditLogOptionsActionTypeBuyClan                                 GetGroupAuditLogOptionsActionType = "BuyClan"
	GetGroupAuditLogOptionsActionTypeCreateGroupAsset                        GetGroupAuditLogOptionsActionType = "CreateGroupAsset"
	GetGroupAuditLogOptionsActionTypeUpdateGroupAsset                        GetGroupAuditLogOptionsActionType = "UpdateGroupAsset"
	GetGroupAuditLogOptionsActionTypeConfigureGroupAsset                     GetGroupAuditLogOptionsActionType = "ConfigureGroupAsset"
	GetGroupAuditLogOptionsActionTypeRevertGroupAsset                        GetGroupAuditLogOptionsActionType = "RevertGroupAsset"
	GetGroupAuditLogOptionsActionTypeCreateGroupDeveloperProduct             GetGroupAuditLogOptionsActionType = "CreateGroupDeveloperProduct"
	GetGroupAuditLogOptionsActionTypeConfigureGroupGame                      GetGroupAuditLogOptionsActionType = "ConfigureGroupGame"
	GetGroupAuditLogOptionsActionTypeCreateGroupDeveloperSubscriptionProduct GetGroupAuditLogOptionsActionType = "CreateGroupDeveloperSubscriptionProduct"
	GetGroupAuditLogOptionsActionTypeLock                                    GetGroupAuditLogOptionsActionType = "Lock"
	GetGroupAuditLogOptionsActionTypeUnlock                                  GetGroupAuditLogOptionsActionType = "Unlock"
	GetGroupAuditLogOptionsActionTypeCreateGamePass                          GetGroupAuditLogOptionsActionType = "CreateGamePass"
	GetGroupAuditLogOptionsActionTypeCreateBadge                             GetGroupAuditLogOptionsActionType = "CreateBadge"
	GetGroupAuditLogOptionsActionTypeConfigureBadge                          GetGroupAuditLogOptionsActionType = "ConfigureBadge"
	GetGroupAuditLogOptionsActionTypeSavePlace                               GetGroupAuditLogOptionsActionType = "SavePlace"
	GetGroupAuditLogOptionsActionTypePublishPlace                            GetGroupAuditLogOptionsActionType = "PublishPlace"
	GetGroupAuditLogOptionsActionTypeUpdateRolesetRank                       GetGroupAuditLogOptionsActionType = "UpdateRolesetRank"
	GetGroupAuditLogOptionsActionTypeUpdateRolesetData                       GetGroupAuditLogOptionsActionType = "UpdateRolesetData"
	GetGroupAuditLogOptionsActionTypeUpdateRoleSetPosition                   GetGroupAuditLogOptionsActionType = "UpdateRoleSetPosition"
	GetGroupAuditLogOptionsActionTypeBanMember                               GetGroupAuditLogOptionsActionType = "BanMember"
	GetGroupAuditLogOptionsActionTypeUnbanMember                             GetGroupAuditLogOptionsActionType = "UnbanMember"
	GetGroupAuditLogOptionsActionTypeCreateForumCategory                     GetGroupAuditLogOptionsActionType = "CreateForumCategory"
	GetGroupAuditLogOptionsActionTypeUpdateForumCategory                     GetGroupAuditLogOptionsActionType = "UpdateForumCategory"
	GetGroupAuditLogOptionsActionTypeArchiveForumCategory                    GetGroupAuditLogOptionsActionType = "ArchiveForumCategory"
	GetGroupAuditLogOptionsActionTypeDeleteForumCategory                     GetGroupAuditLogOptionsActionType = "DeleteForumCategory"
	GetGroupAuditLogOptionsActionTypeDeleteForumPost                         GetGroupAuditLogOptionsActionType = "DeleteForumPost"
	GetGroupAuditLogOptionsActionTypeDeleteForumComment                      GetGroupAuditLogOptionsActionType = "DeleteForumComment"
	GetGroupAuditLogOptionsActionTypePinForumPost                            GetGroupAuditLogOptionsActionType = "PinForumPost"
	GetGroupAuditLogOptionsActionTypeUnpinForumPost                          GetGroupAuditLogOptionsActionType = "UnpinForumPost"
	GetGroupAuditLogOptionsActionTypeLockForumPost                           GetGroupAuditLogOptionsActionType = "LockForumPost"
	GetGroupAuditLogOptionsActionTypeUnlockForumPost                         GetGroupAuditLogOptionsActionType = "UnlockForumPost"
	GetGroupAuditLogOptionsActionTypeCreateRoleset                           GetGroupAuditLogOptionsActionType = "CreateRoleset"
	GetGroupAuditLogOptionsActionTypeDeleteRoleset                           GetGroupAuditLogOptionsActionType = "DeleteRoleset"
	GetGroupAuditLogOptionsActionTypeCreateCommerceProduct                   GetGroupAuditLogOptionsActionType = "CreateCommerceProduct"
	GetGroupAuditLogOptionsActionTypeSetCommerceProductActive                GetGroupAuditLogOptionsActionType = "SetCommerceProductActive"
	GetGroupAuditLogOptionsActionTypeArchiveCommerceProduct                  GetGroupAuditLogOptionsActionType = "ArchiveCommerceProduct"
	GetGroupAuditLogOptionsActionTypeAcceptCommerceProductBundlingFee        GetGroupAuditLogOptionsActionType = "AcceptCommerceProductBundlingFee"
	GetGroupAuditLogOptionsActionTypeSetCommerceProductInactive              GetGroupAuditLogOptionsActionType = "SetCommerceProductInactive"
	GetGroupAuditLogOptionsActionTypeConnectMerchant                         GetGroupAuditLogOptionsActionType = "ConnectMerchant"
	GetGroupAuditLogOptionsActionTypeDisconnectMerchant                      GetGroupAuditLogOptionsActionType = "DisconnectMerchant"
	GetGroupAuditLogOptionsActionTypeJoinGroup                               GetGroupAuditLogOptionsActionType = "JoinGroup"
	GetGroupAuditLogOptionsActionTypeLeaveGroup                              GetGroupAuditLogOptionsActionType = "LeaveGroup"
	GetGroupAuditLogOptionsActionTypeUpdateGroupIcon                         GetGroupAuditLogOptionsActionType = "UpdateGroupIcon"
	GetGroupAuditLogOptionsActionTypeUpdateGroupCoverPhoto                   GetGroupAuditLogOptionsActionType = "UpdateGroupCoverPhoto"
	GetGroupAuditLogOptionsActionTypeAssignRole                              GetGroupAuditLogOptionsActionType = "AssignRole"
	GetGroupAuditLogOptionsActionTypeUnassignRole                            GetGroupAuditLogOptionsActionType = "UnassignRole"
	GetGroupAuditLogOptionsActionTypePublishAnnouncement                     GetGroupAuditLogOptionsActionType = "PublishAnnouncement"
	GetGroupAuditLogOptionsActionTypeDeleteAnnouncement                      GetGroupAuditLogOptionsActionType = "DeleteAnnouncement"
	GetGroupAuditLogOptionsActionTypeUpdateRoleSetPermissions                GetGroupAuditLogOptionsActionType = "UpdateRoleSetPermissions"
	GetGroupAuditLogOptionsActionTypeUpdateGroupSecuritySettings             GetGroupAuditLogOptionsActionType = "UpdateGroupSecuritySettings"
	GetGroupAuditLogOptionsActionTypeGrantEnterpriseTier                     GetGroupAuditLogOptionsActionType = "GrantEnterpriseTier"
	GetGroupAuditLogOptionsActionTypeRevokeEnterpriseTier                    GetGroupAuditLogOptionsActionType = "RevokeEnterpriseTier"
	GetGroupAuditLogOptionsActionTypeDeactivateGroup                         GetGroupAuditLogOptionsActionType = "DeactivateGroup"
)

type GetGroupAuditLogOptionsLimit int

type GetGroupAuditLogOptionsSortOrder string

const (
	GetGroupAuditLogOptionsSortOrderAsc  GetGroupAuditLogOptionsSortOrder = "Asc"
	GetGroupAuditLogOptionsSortOrderDesc GetGroupAuditLogOptionsSortOrder = "Desc"
)

type GetGroupAuditLogOptions struct {
	ActionType *GetGroupAuditLogOptionsActionType `url:"actionType,omitempty"`
	UserID     *int64                             `url:"userId,omitempty"`
	Limit      *GetGroupAuditLogOptionsLimit      `url:"limit,omitempty"`
	Cursor     *string                            `url:"cursor,omitempty"`
	SortOrder  *GetGroupAuditLogOptionsSortOrder  `url:"sortOrder,omitempty"`
}

type GroupUserModelBuildersClubMembershipType int

type GroupUserModel struct {
	BuildersClubMembershipType GroupUserModelBuildersClubMembershipType `json:"buildersClubMembershipType"`
	HasVerifiedBadge           bool                                     `json:"hasVerifiedBadge"`
	UserID                     int64                                    `json:"userId"`
	Username                   string                                   `json:"username"`
	DisplayName                string                                   `json:"displayName"`
}

type GroupRoleResponseColor int

type GroupRoleResponse struct {
	ID          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Rank        int                    `json:"rank"`
	MemberCount int64                  `json:"memberCount"`
	IsBase      bool                   `json:"isBase"`
	Color       GroupRoleResponseColor `json:"color"`
	IsPrivate   bool                   `json:"isPrivate"`
}

type UserGroupRoleResponse struct {
	User GroupUserModel    `json:"user"`
	Role GroupRoleResponse `json:"role"`
}

type GroupAuditLogResponseItem struct {
	Actor       UserGroupRoleResponse `json:"actor"`
	ActionType  string                `json:"actionType"`
	Description any                   `json:"description"`
	Created     string                `json:"created"`
}

type GroupAuditLogPage struct {
	PreviousPageCursor string                      `json:"previousPageCursor"`
	NextPageCursor     string                      `json:"nextPageCursor"`
	Data               []GroupAuditLogResponseItem `json:"data"`
}

// GetGroupAuditLog will get the Group's audit log.
//
// Required scopes: legacy-group:manage
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#get_legacy_groups_v1_groups__groupId__audit_log
//
// [GET] /legacy-groups/v1/groups/{groupId}/audit-log
func (s *GroupsService) GetGroupAuditLog(ctx context.Context, groupId string, opts *GetGroupAuditLogOptions, reqOpts ...goblox.RequestOpt) (*GroupAuditLogPage, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	groupAuditLogPage := new(GroupAuditLogPage)
	resp, err := s.client.Call(ctx, GetGroupAuditLog.Compile(q, groupId), nil, groupAuditLogPage, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupAuditLogPage, resp, nil
}

type UpdateGroupDescriptionRequest struct {
	Description *string `json:"description,omitempty"`
}

type GroupDescriptionResponse struct {
	NewDescription string `json:"newDescription"`
}

// UpdateGroupDescription will update the group's description.
//
// Required scopes: legacy-group:manage
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#patch_legacy_groups_v1_groups__groupId__description
//
// [PATCH] /legacy-groups/v1/groups/{groupId}/description
func (s *GroupsService) UpdateGroupDescription(ctx context.Context, groupId string, data UpdateGroupDescriptionRequest, reqOpts ...goblox.RequestOpt) (*GroupDescriptionResponse, *Response, error) {
	groupDescriptionResponse := new(GroupDescriptionResponse)
	resp, err := s.client.Call(ctx, UpdateGroupDescription.Compile(nil, groupId), data, groupDescriptionResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupDescriptionResponse, resp, nil
}

type UpdateGroupNotificationPreferenceRequestType int

type UpdateGroupNotificationPreferenceRequest struct {
	NotificationsEnabled *bool                                         `json:"notificationsEnabled,omitempty"`
	Type                 *UpdateGroupNotificationPreferenceRequestType `json:"type,omitempty"`
}

type UpdateGroupNotificationPreferenceResponse int

// UpdateGroupNotificationPreference will update the authenticated user's notification preference for a group.
//
// Required scopes: legacy-user:manage
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#patch_legacy_groups_v1_groups__groupId__notification_preference
//
// [PATCH] /legacy-groups/v1/groups/{groupId}/notification-preference
func (s *GroupsService) UpdateGroupNotificationPreference(ctx context.Context, groupId string, data UpdateGroupNotificationPreferenceRequest, reqOpts ...goblox.RequestOpt) (*UpdateGroupNotificationPreferenceResponse, *Response, error) {
	updateGroupNotificationPreferenceResponse := new(UpdateGroupNotificationPreferenceResponse)
	resp, err := s.client.Call(ctx, UpdateGroupNotificationPreference.Compile(nil, groupId), data, updateGroupNotificationPreferenceResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateGroupNotificationPreferenceResponse, resp, nil
}

type GroupSettingsResponseVerificationLevel int

type GroupSettingsResponseAccountTenureRequirement int

type GroupSettingsResponseSlowmode int

type TierRequirementKey int

type TierRequirement struct {
	Key       TierRequirementKey `json:"key"`
	Satisfied bool               `json:"satisfied"`
}

type TierCapabilities struct {
	IsEligibleForUnrestrictedMessages bool `json:"isEligibleForUnrestrictedMessages"`
}

type CommunityTierInfoResponse struct {
	GroupID           int64             `json:"groupId"`
	CurrentTier       int               `json:"currentTier"`
	PreviousTier      int               `json:"previousTier"`
	TierUpdatedTime   string            `json:"tierUpdatedTime"`
	LastEvaluatedTime string            `json:"lastEvaluatedTime"`
	Requirements      []TierRequirement `json:"requirements"`
	Capabilities      TierCapabilities  `json:"capabilities"`
}

type GroupSettingsResponse struct {
	IsApprovalRequired          bool                                          `json:"isApprovalRequired"`
	IsBuildersClubRequired      bool                                          `json:"isBuildersClubRequired"`
	AreEnemiesAllowed           bool                                          `json:"areEnemiesAllowed"`
	AreGroupFundsVisible        bool                                          `json:"areGroupFundsVisible"`
	AreGroupGamesVisible        bool                                          `json:"areGroupGamesVisible"`
	IsGroupNameChangeEnabled    bool                                          `json:"isGroupNameChangeEnabled"`
	VerificationLevel           GroupSettingsResponseVerificationLevel        `json:"verificationLevel"`
	AccountTenureRequirement    GroupSettingsResponseAccountTenureRequirement `json:"accountTenureRequirement"`
	Slowmode                    GroupSettingsResponseSlowmode                 `json:"slowmode"`
	IsMemberListVisibleToPublic bool                                          `json:"isMemberListVisibleToPublic"`
	IsAutoAssignRoleDisabled    bool                                          `json:"isAutoAssignRoleDisabled"`
	HasLegacyAutoAssignRole     bool                                          `json:"hasLegacyAutoAssignRole"`
	CommunityTier               CommunityTierInfoResponse                     `json:"communityTier"`
}

// GetGroupSettings will get the Group's settings.
//
// Required scopes: legacy-group:manage
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#get_legacy_groups_v1_groups__groupId__settings
//
// [GET] /legacy-groups/v1/groups/{groupId}/settings
func (s *GroupsService) GetGroupSettings(ctx context.Context, groupId string, reqOpts ...goblox.RequestOpt) (*GroupSettingsResponse, *Response, error) {
	groupSettingsResponse := new(GroupSettingsResponse)
	resp, err := s.client.Call(ctx, GetGroupSettings.Compile(nil, groupId), nil, groupSettingsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupSettingsResponse, resp, nil
}

type UpdateGroupSettingsRequestVerificationLevel int

type UpdateGroupSettingsRequestAccountTenureRequirement int

type UpdateGroupSettingsRequestSlowmode int

type UpdateGroupSettingsRequest struct {
	IsApprovalRequired          *bool                                               `json:"isApprovalRequired,omitempty"`
	AreEnemiesAllowed           *bool                                               `json:"areEnemiesAllowed,omitempty"`
	AreGroupFundsVisible        *bool                                               `json:"areGroupFundsVisible,omitempty"`
	AreGroupGamesVisible        *bool                                               `json:"areGroupGamesVisible,omitempty"`
	VerificationLevel           *UpdateGroupSettingsRequestVerificationLevel        `json:"verificationLevel,omitempty"`
	AccountTenureRequirement    *UpdateGroupSettingsRequestAccountTenureRequirement `json:"accountTenureRequirement,omitempty"`
	Slowmode                    *UpdateGroupSettingsRequestSlowmode                 `json:"slowmode,omitempty"`
	IsMemberListVisibleToPublic *bool                                               `json:"isMemberListVisibleToPublic,omitempty"`
}

// UpdateGroupSettings will update the group's settings.
//
// Required scopes: legacy-group:manage
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#patch_legacy_groups_v1_groups__groupId__settings
//
// [PATCH] /legacy-groups/v1/groups/{groupId}/settings
func (s *GroupsService) UpdateGroupSettings(ctx context.Context, groupId string, data UpdateGroupSettingsRequest, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, UpdateGroupSettings.Compile(nil, groupId), data, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type PostGroupStatusRequest struct {
	Message *string `json:"message,omitempty"`
}

type ShoutResponse struct {
	Body    string         `json:"body"`
	Poster  GroupUserModel `json:"poster"`
	Created string         `json:"created"`
	Updated string         `json:"updated"`
}

// UpdateGroupShout will set the group's shout (status).
//
// Required scopes: legacy-group:manage
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#patch_legacy_groups_v1_groups__groupId__status
//
// [PATCH] /legacy-groups/v1/groups/{groupId}/status
func (s *GroupsService) UpdateGroupShout(ctx context.Context, groupId string, data PostGroupStatusRequest, reqOpts ...goblox.RequestOpt) (*ShoutResponse, *Response, error) {
	shoutResponse := new(ShoutResponse)
	resp, err := s.client.Call(ctx, UpdateGroupShout.Compile(nil, groupId), data, shoutResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return shoutResponse, resp, nil
}

type GroupDetailResponse struct {
	ID                 int64                     `json:"id"`
	Name               string                    `json:"name"`
	Description        string                    `json:"description"`
	Owner              GroupUserModel            `json:"owner"`
	Shout              ShoutResponse             `json:"shout"`
	MemberCount        int                       `json:"memberCount"`
	IsBuildersClubOnly bool                      `json:"isBuildersClubOnly"`
	PublicEntryAllowed bool                      `json:"publicEntryAllowed"`
	IsLocked           bool                      `json:"isLocked"`
	IsDeactivated      bool                      `json:"isDeactivated"`
	HasVerifiedBadge   bool                      `json:"hasVerifiedBadge"`
	HasSocialModules   bool                      `json:"hasSocialModules"`
	CommunityTier      CommunityTierInfoResponse `json:"communityTier"`
}

type GroupDetailList struct {
	Data []GroupDetailResponse `json:"data"`
}

// ListPendingGroupJoinRequests will get groups that the authenticated user has requested to join.
//
// Required scopes: legacy-group:manage
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/groups#get_legacy_groups_v1_user_groups_pending
//
// [GET] /legacy-groups/v1/user/groups/pending
func (s *GroupsService) ListPendingGroupJoinRequests(ctx context.Context, reqOpts ...goblox.RequestOpt) (*GroupDetailList, *Response, error) {
	groupDetailList := new(GroupDetailList)
	resp, err := s.client.Call(ctx, ListPendingGroupJoinRequests.Compile(nil), nil, groupDetailList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return groupDetailList, resp, nil
}
