package opencloud

import "context"

// UserAndGroupsService handles communication with user and group endpoints.
//
// Deprecated: Use [UsersService] and [GroupsService] instead.
type UserAndGroupsService service

// ListAssetQuota will return information regarding the user's asset quota uploads.
//
// Deprecated: Use [UsersService.ListAssetQuota] instead.
func (s *UserAndGroupsService) ListAssetQuota(ctx context.Context, userId string, opts *OptionsWithFilter) (*AssetQuotas, *Response, error) {
	return s.client.Users.ListAssetQuota(ctx, userId, opts)
}

// GetGroup will fetch information on a specificed group.
//
// Deprecated: Use [GroupsService.GetGroup] instead.
func (s *UserAndGroupsService) GetGroup(ctx context.Context, groupId string) (*Group, *Response, error) {
	return s.client.Groups.GetGroup(ctx, groupId)
}

// ListGroupJoinRequests will fetch join requests for a specificed group.
//
// Deprecated: Use [GroupsService.ListGroupJoinRequests] instead.
func (s *UserAndGroupsService) ListGroupJoinRequests(ctx context.Context, groupId string, opts *OptionsWithFilter) (*GroupJoinRequests, *Response, error) {
	return s.client.Groups.ListGroupJoinRequests(ctx, groupId, opts)
}

// AcceptGroupJoinRequest will accept the join request for a user under a specificed group.
//
// Deprecated: Use [GroupsService.AcceptGroupJoinRequest] instead.
func (s *UserAndGroupsService) AcceptGroupJoinRequest(ctx context.Context, groupId, userId string) (*Response, error) {
	return s.client.Groups.AcceptGroupJoinRequest(ctx, groupId, userId)
}

// DeclineGroupJoinRequest will decline the join request for a user under a specificed group.
//
// Deprecated: Use [GroupsService.DeclineGroupJoinRequest] instead.
func (s *UserAndGroupsService) DeclineGroupJoinRequest(ctx context.Context, groupId, userId string) (*Response, error) {
	return s.client.Groups.DeclineGroupJoinRequest(ctx, groupId, userId)
}

// ListGroupMemberships will fetch memberships for a specificed group.
//
// Deprecated: Use [GroupsService.ListGroupMemberships] instead.
func (s *UserAndGroupsService) ListGroupMemberships(ctx context.Context, groupId string, opts *OptionsWithFilter) (*GroupMemberships, *Response, error) {
	return s.client.Groups.ListGroupMemberships(ctx, groupId, opts)
}

// UpdateGroupMemberships will update memberships for a user under a specificed group.
//
// Deprecated: Use [GroupsService.UpdateGroupMemberships] instead.
func (s *UserAndGroupsService) UpdateGroupMemberships(ctx context.Context, groupId, userId string, data GroupMembershipUpdate) (*GroupMembership, *Response, error) {
	return s.client.Groups.UpdateGroupMemberships(ctx, groupId, userId, data)
}

// ListGroupRoles will fetch roles for a specificed group.
//
// Deprecated: Use [GroupsService.ListGroupRoles] instead.
func (s *UserAndGroupsService) ListGroupRoles(ctx context.Context, groupId string, opts *OptionsWithFilter) (*GroupRoles, *Response, error) {
	return s.client.Groups.ListGroupRoles(ctx, groupId, opts)
}

// GetGroupRoles will fetch a specificed role for a specificed group.
//
// Deprecated: Use [GroupsService.GetGroupRoles] instead.
func (s *UserAndGroupsService) GetGroupRoles(ctx context.Context, groupId, roleId string) (*GroupRole, *Response, error) {
	return s.client.Groups.GetGroupRoles(ctx, groupId, roleId)
}

// GetGroupShout will fetch the shout for a specificed group.
//
// Deprecated: Use [GroupsService.GetGroupShout] instead.
func (s *UserAndGroupsService) GetGroupShout(ctx context.Context, groupId string) (*GroupShout, *Response, error) {
	return s.client.Groups.GetGroupShout(ctx, groupId)
}

// ListInventoryItems will fetch inventory items for a specificed user.
//
// Deprecated: Use [UsersService.ListInventoryItems] instead.
func (s *UserAndGroupsService) ListInventoryItems(ctx context.Context, userId string, opts *OptionsWithFilter) (*InventoryItems, *Response, error) {
	return s.client.Users.ListInventoryItems(ctx, userId, opts)
}

// GetUser will fetch information on a specificed user.
//
// Deprecated: Use [UsersService.GetUser] instead.
func (s *UserAndGroupsService) GetUser(ctx context.Context, userId string) (*User, *Response, error) {
	return s.client.Users.GetUser(ctx, userId)
}

// GenerateUserThumbnail will generate and return a url for a user's avatar.
//
// Deprecated: Use [UsersService.GenerateUserThumbnail] instead.
func (s *UserAndGroupsService) GenerateUserThumbnail(ctx context.Context, userId string, opts *UserThumbnailOptions) (*UserThumbnail, *Response, error) {
	return s.client.Users.GenerateUserThumbnail(ctx, userId, opts)
}

// CreateUserNotification will send a notification to a user.
//
// Deprecated: Use [UsersService.CreateUserNotification] instead.
func (s *UserAndGroupsService) CreateUserNotification(ctx context.Context, userId string, data UserNotificationCreate) (*UserNotification, *Response, error) {
	return s.client.Users.CreateUserNotification(ctx, userId, data)
}

// ListUserRestrictions will list restrictions for users thave have been restricted for the specified universe or place.
//
// Deprecated: Use [UsersService.ListUserRestrictions] instead.
func (s *UserAndGroupsService) ListUserRestrictions(ctx context.Context, universeId string, placeId *string, opts *Options) (*UserRestrictions, *Response, error) {
	return s.client.Users.ListUserRestrictions(ctx, universeId, placeId, opts)
}

// GetUserRestriction will get the user restriction for the specified user under the specified universe or place.
//
// Deprecated: Use [UsersService.GetUserRestriction] instead.
func (s *UserAndGroupsService) GetUserRestriction(ctx context.Context, universeId string, placeId *string, userId string) (*UserRestriction, *Response, error) {
	return s.client.Users.GetUserRestriction(ctx, universeId, placeId, userId)
}

// UpdateUserRestriction will update the current user restriction for the specified user under the specified universe or place.
//
// Deprecated: Use [UsersService.UpdateUserRestriction] instead.
func (s *UserAndGroupsService) UpdateUserRestriction(ctx context.Context, universeId string, placeId *string, userId string, data UserRestrictionUpdate, opts *UserRestrictionUpdateOptions) (*UserRestriction, *Response, error) {
	return s.client.Users.UpdateUserRestriction(ctx, universeId, placeId, userId, data, opts)
}

// ListUserRestrictionLogs will list logs of user restrictions for the specified universe or place.
//
// Deprecated: Use [UsersService.ListUserRestrictionLogs] instead.
func (s *UserAndGroupsService) ListUserRestrictionLogs(ctx context.Context, universeId string, placeId *string, opts *OptionsWithFilter) (*UserRestrictionLogs, *Response, error) {
	return s.client.Users.ListUserRestrictionLogs(ctx, universeId, placeId, opts)
}
