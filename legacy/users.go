package legacy

import (
	"context"
	"net/http"

	"github.com/luckfire/goblox"
	"github.com/luckfire/goblox/auth"
)

var (
	GetBirthdate = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/birthdate",
		auth.KindCookie,
	).WithDomain("users")

	UpdateBirthdate = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/birthdate",
		auth.KindCookie,
	).WithDomain("users")

	GetDescription = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/description",
		auth.KindCookie,
	).WithDomain("users")

	UpdateDescription = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/description",
		auth.KindCookie,
	).WithDomain("users")

	ValidateDisplayName = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/display-names/validate",
		auth.KindCookie, auth.KindNone,
	).WithDomain("users")

	GetGender = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/gender",
		auth.KindCookie,
	).WithDomain("users")

	UpdateGender = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/gender",
		auth.KindCookie,
	).WithDomain("users")

	GetUser = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}",
		auth.KindCookie, auth.KindNone,
	).WithDomain("users")

	ValidateUserDisplayName = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/display-names/validate",
		auth.KindCookie,
	).WithDomain("users")

	SetDisplayName = goblox.NewEndpoint(
		http.MethodPatch,
		"/v1/users/{userId}/display-names",
		auth.KindCookie,
	).WithDomain("users")

	GetUsernameHistory = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/username-history",
		auth.KindCookie, auth.KindNone,
	).WithDomain("users")

	GetAuthenticatedUser = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/authenticated",
		auth.KindCookie,
	).WithDomain("users")

	GetAuthenticatedUserAgeBracket = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/authenticated/age-bracket",
		auth.KindCookie,
	).WithDomain("users")

	GetAuthenticatedUserCountryCode = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/authenticated/country-code",
		auth.KindCookie,
	).WithDomain("users")

	GetAuthenticatedUserRoles = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/authenticated/roles",
		auth.KindCookie,
	).WithDomain("users")

	SearchUsers = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/search",
		auth.KindCookie, auth.KindNone,
	).WithDomain("users")

	GetUsersByUsernames = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/usernames/users",
		auth.KindCookie, auth.KindNone,
	).WithDomain("users")

	GetUsersByIDs = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/users",
		auth.KindCookie, auth.KindNone,
	).WithDomain("users")
)

type UsersService service

type Birthdate struct {
	BirthMonth int `json:"birthMonth"`
	BirthDay   int `json:"birthDay"`
	BirthYear  int `json:"birthYear"`
}

// GetBirthdate will get the authenticated user's birthdate.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_birthdate
//
// [GET] https://users.roblox.com/v1/birthdate
func (s *UsersService) GetBirthdate(ctx context.Context, reqOpts ...goblox.RequestOpt) (*Birthdate, *Response, error) {
	birthdate := new(Birthdate)
	resp, err := s.client.Call(ctx, GetBirthdate.Compile(nil), nil, birthdate, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return birthdate, resp, nil
}

type UpdateBirthdateRequest struct {
	BirthMonth int    `json:"birthMonth"`
	BirthDay   int    `json:"birthDay"`
	BirthYear  int    `json:"birthYear"`
	Password   string `json:"password"`
}

// UpdateBirthdate will update the authenticated user's birthdate.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_post_v1_birthdate
//
// [POST] https://users.roblox.com/v1/birthdate
func (s *UsersService) UpdateBirthdate(ctx context.Context, data UpdateBirthdateRequest, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, UpdateBirthdate.Compile(nil), data, nil, reqOpts...)
}

type descriptionBody struct {
	Description string `json:"description"`
}

// GetDescription will get the authenticated user's description.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_description
//
// [GET] https://users.roblox.com/v1/description
func (s *UsersService) GetDescription(ctx context.Context, reqOpts ...goblox.RequestOpt) (string, *Response, error) {
	result := new(descriptionBody)
	resp, err := s.client.Call(ctx, GetDescription.Compile(nil), nil, result, reqOpts...)
	if err != nil {
		return "", resp, err
	}

	return result.Description, resp, nil
}

// UpdateDescription will update the authenticated user's description and return the new description.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_post_v1_description
//
// [POST] https://users.roblox.com/v1/description
func (s *UsersService) UpdateDescription(ctx context.Context, description string, reqOpts ...goblox.RequestOpt) (string, *Response, error) {
	result := new(descriptionBody)
	resp, err := s.client.Call(ctx, UpdateDescription.Compile(nil), descriptionBody{Description: description}, result, reqOpts...)
	if err != nil {
		return "", resp, err
	}

	return result.Description, resp, nil
}

type ValidateDisplayNameOptions struct {
	DisplayName string `url:"displayName"` // Required.
	Birthdate   string `url:"birthdate"`   // Required. An ISO 8601 date-time.
}

// ValidateDisplayName will check whether a display name is valid for a new user.
// A nil error means the display name is valid.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_display_names_validate
//
// [GET] https://users.roblox.com/v1/display-names/validate
func (s *UsersService) ValidateDisplayName(ctx context.Context, opts *ValidateDisplayNameOptions, reqOpts ...goblox.RequestOpt) (*Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, err
	}

	return s.client.Call(ctx, ValidateDisplayName.Compile(q), nil, nil, reqOpts...)
}

type Gender int

const (
	GenderUnknown Gender = 1
	GenderMale    Gender = 2
	GenderFemale  Gender = 3
)

type genderBody struct {
	Gender Gender `json:"gender"`
}

// GetGender will get the authenticated user's gender.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_gender
//
// [GET] https://users.roblox.com/v1/gender
func (s *UsersService) GetGender(ctx context.Context, reqOpts ...goblox.RequestOpt) (Gender, *Response, error) {
	result := new(genderBody)
	resp, err := s.client.Call(ctx, GetGender.Compile(nil), nil, result, reqOpts...)
	if err != nil {
		return 0, resp, err
	}

	return result.Gender, resp, nil
}

// UpdateGender will update the authenticated user's gender.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_post_v1_gender
//
// [POST] https://users.roblox.com/v1/gender
func (s *UsersService) UpdateGender(ctx context.Context, gender Gender, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, UpdateGender.Compile(nil), genderBody{Gender: gender}, nil, reqOpts...)
}

type User struct {
	ID                     int64  `json:"id"`
	Name                   string `json:"name"`
	DisplayName            string `json:"displayName"`
	Description            string `json:"description"`
	Created                string `json:"created"`
	IsBanned               bool   `json:"isBanned"`
	ExternalAppDisplayName string `json:"externalAppDisplayName"`
	HasVerifiedBadge       bool   `json:"hasVerifiedBadge"`
}

// GetUser will get detailed information about a user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_users__userId_
//
// [GET] https://users.roblox.com/v1/users/{userId}
func (s *UsersService) GetUser(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (*User, *Response, error) {
	user := new(User)
	resp, err := s.client.Call(ctx, GetUser.Compile(nil, userId), nil, user, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

type ValidateUserDisplayNameOptions struct {
	DisplayName string `url:"displayName"` // Required.
}

// ValidateUserDisplayName will check whether a display name is valid for an existing user.
// A nil error means the display name is valid.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_users__userId__display_names_validate
//
// [GET] https://users.roblox.com/v1/users/{userId}/display-names/validate
func (s *UsersService) ValidateUserDisplayName(ctx context.Context, userId string, opts *ValidateUserDisplayNameOptions, reqOpts ...goblox.RequestOpt) (*Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, err
	}

	return s.client.Call(ctx, ValidateUserDisplayName.Compile(q, userId), nil, nil, reqOpts...)
}

type setDisplayNameBody struct {
	NewDisplayName string `json:"newDisplayName"`
}

// SetDisplayName will set the display name for the authenticated user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_patch_v1_users__userId__display_names
//
// [PATCH] https://users.roblox.com/v1/users/{userId}/display-names
func (s *UsersService) SetDisplayName(ctx context.Context, userId string, displayName string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, SetDisplayName.Compile(nil, userId), setDisplayNameBody{NewDisplayName: displayName}, nil, reqOpts...)
}

type PreviousUsername struct {
	Name string `json:"name"`
}

// GetUsernameHistory will get a page of a user's previous usernames.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_users__userId__username_history
//
// [GET] https://users.roblox.com/v1/users/{userId}/username-history
func (s *UsersService) GetUsernameHistory(ctx context.Context, userId string, opts *PageOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[PreviousUsername], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[PreviousUsername])
	resp, err := s.client.Call(ctx, GetUsernameHistory.Compile(q, userId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

type AuthenticatedUser struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

// GetAuthenticatedUser will get the user that the client's cookie belongs to.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_users_authenticated
//
// [GET] https://users.roblox.com/v1/users/authenticated
func (s *UsersService) GetAuthenticatedUser(ctx context.Context, reqOpts ...goblox.RequestOpt) (*AuthenticatedUser, *Response, error) {
	user := new(AuthenticatedUser)
	resp, err := s.client.Call(ctx, GetAuthenticatedUser.Compile(nil), nil, user, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// GetAuthenticatedUserAgeBracket will get the age bracket of the authenticated user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_users_authenticated_age_bracket
//
// [GET] https://users.roblox.com/v1/users/authenticated/age-bracket
func (s *UsersService) GetAuthenticatedUserAgeBracket(ctx context.Context, reqOpts ...goblox.RequestOpt) (int, *Response, error) {
	var result struct {
		AgeBracket int `json:"ageBracket"`
	}
	resp, err := s.client.Call(ctx, GetAuthenticatedUserAgeBracket.Compile(nil), nil, &result, reqOpts...)
	if err != nil {
		return 0, resp, err
	}

	return result.AgeBracket, resp, nil
}

// GetAuthenticatedUserCountryCode will get the country code of the authenticated user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_users_authenticated_country_code
//
// [GET] https://users.roblox.com/v1/users/authenticated/country-code
func (s *UsersService) GetAuthenticatedUserCountryCode(ctx context.Context, reqOpts ...goblox.RequestOpt) (string, *Response, error) {
	var result struct {
		CountryCode string `json:"countryCode"`
	}
	resp, err := s.client.Call(ctx, GetAuthenticatedUserCountryCode.Compile(nil), nil, &result, reqOpts...)
	if err != nil {
		return "", resp, err
	}

	return result.CountryCode, resp, nil
}

// GetAuthenticatedUserRoles will get the public roles of the authenticated user.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_users_authenticated_roles
//
// [GET] https://users.roblox.com/v1/users/authenticated/roles
func (s *UsersService) GetAuthenticatedUserRoles(ctx context.Context, reqOpts ...goblox.RequestOpt) ([]string, *Response, error) {
	var result struct {
		Roles []string `json:"roles"`
	}
	resp, err := s.client.Call(ctx, GetAuthenticatedUserRoles.Compile(nil), nil, &result, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return result.Roles, resp, nil
}

type SearchUsersOptions struct {
	Keyword   string  `url:"keyword"` // Required.
	SessionID *string `url:"sessionId,omitempty"`
	Limit     *int    `url:"limit,omitempty"` // 10, 25, 50, or 100.
	Cursor    *string `url:"cursor,omitempty"`
}

type SearchUser struct {
	ID                int64    `json:"id"`
	Name              string   `json:"name"`
	DisplayName       string   `json:"displayName"`
	PreviousUsernames []string `json:"previousUsernames"`
	HasVerifiedBadge  bool     `json:"hasVerifiedBadge"`
}

// SearchUsers will search for users by keyword.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_get_v1_users_search
//
// [GET] https://users.roblox.com/v1/users/search
func (s *UsersService) SearchUsers(ctx context.Context, opts *SearchUsersOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[SearchUser], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[SearchUser])
	resp, err := s.client.Call(ctx, SearchUsers.Compile(q), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

type GetUsersByUsernamesRequest struct {
	Usernames          []string `json:"usernames"`
	ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
}

type UserByUsername struct {
	RequestedUsername string `json:"requestedUsername"`
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	DisplayName       string `json:"displayName"`
	HasVerifiedBadge  bool   `json:"hasVerifiedBadge"`
}

// GetUsersByUsernames will get users by their usernames, including previous usernames.
// Usernames that don't match a user are left out of the results.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_post_v1_usernames_users
//
// [POST] https://users.roblox.com/v1/usernames/users
func (s *UsersService) GetUsersByUsernames(ctx context.Context, data GetUsersByUsernamesRequest, reqOpts ...goblox.RequestOpt) ([]UserByUsername, *Response, error) {
	results := new(WrappedResponse[[]UserByUsername])
	resp, err := s.client.Call(ctx, GetUsersByUsernames.Compile(nil), data, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}

type GetUsersByIDsRequest struct {
	UserIDs            []int64 `json:"userIds"`
	ExcludeBannedUsers bool    `json:"excludeBannedUsers"`
}

type PartialUser struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	DisplayName      string `json:"displayName"`
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"`
}

// GetUsersByIDs will get users by their IDs.
// IDs that don't match a user are left out of the results.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/users#users_post_v1_users
//
// [POST] https://users.roblox.com/v1/users
func (s *UsersService) GetUsersByIDs(ctx context.Context, data GetUsersByIDsRequest, reqOpts ...goblox.RequestOpt) ([]PartialUser, *Response, error) {
	results := new(WrappedResponse[[]PartialUser])
	resp, err := s.client.Call(ctx, GetUsersByIDs.Compile(nil), data, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}
