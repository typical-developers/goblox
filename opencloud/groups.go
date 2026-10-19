package opencloud

import (
	"context"
	"fmt"
	"net/http"
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
func (s *GroupsService) GetGroup(ctx context.Context, groupId string) (*Group, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/groups/%s", groupId)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(Group)
	resp, err := s.client.Do(ctx, req, group)
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
func (s *GroupsService) ListGroupJoinRequests(ctx context.Context, groupId string, opts *OptionsWithFilter) (*GroupJoinRequests, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/groups/%s/join-requests", groupId)

	u, err := addOpts(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	groupJoinRequests := new(GroupJoinRequests)
	resp, err := s.client.Do(ctx, req, groupJoinRequests)
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
func (s *GroupsService) AcceptGroupJoinRequest(ctx context.Context, groupId, userId string) (*Response, error) {
	u := fmt.Sprintf("/cloud/v2/groups/%s/join-requests/%s:accept", groupId, userId)

	req, err := s.client.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
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
func (s *GroupsService) DeclineGroupJoinRequest(ctx context.Context, groupId, userId string) (*Response, error) {
	u := fmt.Sprintf("/cloud/v2/groups/%s/join-requests/%s:decline", groupId, userId)

	req, err := s.client.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
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
func (s *GroupsService) ListGroupMemberships(ctx context.Context, groupId string, opts *OptionsWithFilter) (*GroupMemberships, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/groups/%s/memberships", groupId)

	u, err := addOpts(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	groupMemberships := new(GroupMemberships)
	resp, err := s.client.Do(ctx, req, groupMemberships)
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
func (s *GroupsService) UpdateGroupMemberships(ctx context.Context, groupId, userId string, data GroupMembershipUpdate) (*GroupMembership, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/groups/%s/memberships/%s", groupId, userId)

	req, err := s.client.NewRequest(http.MethodPatch, u, data)
	if err != nil {
		return nil, nil, err
	}

	groupMembership := new(GroupMembership)
	resp, err := s.client.Do(ctx, req, groupMembership)
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
func (s *GroupsService) ListGroupRoles(ctx context.Context, groupId string, opts *OptionsWithFilter) (*GroupRoles, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/groups/%s/roles", groupId)

	u, err := addOpts(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	groupRoles := new(GroupRoles)
	resp, err := s.client.Do(ctx, req, groupRoles)
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
func (s *GroupsService) GetGroupRoles(ctx context.Context, groupId, roleId string) (*GroupRole, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/groups/%s/roles/%s", groupId, roleId)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	groupRole := new(GroupRole)
	resp, err := s.client.Do(ctx, req, groupRole)
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
func (s *GroupsService) GetGroupShout(ctx context.Context, groupId string) (*GroupShout, *Response, error) {
	u := fmt.Sprintf("/cloud/v2/groups/%s/shout", groupId)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	groupShout := new(GroupShout)
	resp, err := s.client.Do(ctx, req, groupShout)
	if err != nil {
		return nil, resp, err
	}

	return groupShout, resp, nil
}
