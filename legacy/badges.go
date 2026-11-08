package legacy

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	GetBadge = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/badges/{badgeId}",
		auth.KindCookie, auth.KindNone,
	).WithDomain("badges")

	UpdateBadge = goblox.NewEndpoint(
		http.MethodPatch,
		"/v1/badges/{badgeId}",
		auth.KindCookie,
	).WithDomain("badges")

	GetBadgeMetadata = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/badges/metadata",
		auth.KindCookie, auth.KindNone,
	).WithDomain("badges")

	GetUniverseBadges = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/universes/{universeId}/badges",
		auth.KindCookie, auth.KindNone,
	).WithDomain("badges")

	CreateBadge = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/universes/{universeId}/badges",
		auth.KindCookie,
	).WithDomain("badges")

	GetFreeBadgesQuota = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/universes/{universeId}/free-badges-quota",
		auth.KindCookie,
	).WithDomain("badges")

	GetUserBadges = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/badges",
		auth.KindCookie, auth.KindNone,
	).WithDomain("badges")

	GetUserBadgeAwardedDate = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/badges/{badgeId}/awarded-date",
		auth.KindCookie, auth.KindNone,
	).WithDomain("badges")

	GetUserBadgeAwardedDates = goblox.NewEndpoint(
		http.MethodGet,
		"/v1/users/{userId}/badges/awarded-dates",
		auth.KindCookie, auth.KindNone,
	).WithDomain("badges")

	UpdateBadgeIcon = goblox.NewEndpoint(
		http.MethodPost,
		"/v1/badges/{badgeId}/icon",
		auth.KindCookie,
	).WithDomain("badges")

	RemoveBadge = goblox.NewEndpoint(
		http.MethodDelete,
		"/v1/user/badges/{badgeId}",
		auth.KindCookie,
	).WithDomain("badges")
)

type BadgesService service

type BadgeStatistics struct {
	PastDayAwardedCount int64   `json:"pastDayAwardedCount"`
	AwardedCount        int64   `json:"awardedCount"`
	WinRatePercentage   float64 `json:"winRatePercentage"`
}

type BadgeUniverse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	RootPlaceID int64  `json:"rootPlaceId"`
}

// Badge is a badge. Depending on the endpoint, either AwardingUniverse or Awarder is set.
type Badge struct {
	ID                 int64           `json:"id"`
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	DisplayName        string          `json:"displayName"`
	DisplayDescription string          `json:"displayDescription"`
	Enabled            bool            `json:"enabled"`
	IconImageID        int64           `json:"iconImageId"`
	DisplayIconImageID int64           `json:"displayIconImageId"`
	Created            string          `json:"created"`
	Updated            string          `json:"updated"`
	Statistics         BadgeStatistics `json:"statistics"`
	AwardingUniverse   *BadgeUniverse  `json:"awardingUniverse,omitempty"`
	Awarder            *RelatedEntity  `json:"awarder,omitempty"`
}

// GetBadge will get a badge, including its award statistics and the universe that awards it.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_get_v1_badges__badgeId_
//
// [GET] https://badges.roblox.com/v1/badges/{badgeId}
func (s *BadgesService) GetBadge(ctx context.Context, badgeId string, reqOpts ...goblox.RequestOpt) (*Badge, *Response, error) {
	badge := new(Badge)
	resp, err := s.client.Call(ctx, GetBadge.Compile(nil, badgeId), nil, badge, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return badge, resp, nil
}

type UpdateBadgeRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
}

// UpdateBadge will update a badge's name, description, or enabled state.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_patch_v1_badges__badgeId_
//
// [PATCH] https://badges.roblox.com/v1/badges/{badgeId}
func (s *BadgesService) UpdateBadge(ctx context.Context, badgeId string, data UpdateBadgeRequest, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, UpdateBadge.Compile(nil, badgeId), data, nil, reqOpts...)
}

type BadgeMetadata struct {
	BadgeCreationPrice        int64 `json:"badgeCreationPrice"`
	MaxBadgeNameLength        int   `json:"maxBadgeNameLength"`
	MaxBadgeDescriptionLength int   `json:"maxBadgeDescriptionLength"`
	DefaultBadgeIconImageID   int64 `json:"defaultBadgeIconImageId"` // 0 when there is no default icon.
}

// GetBadgeMetadata will get badge system metadata, such as the creation price and field length limits.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_get_v1_badges_metadata
//
// [GET] https://badges.roblox.com/v1/badges/metadata
func (s *BadgesService) GetBadgeMetadata(ctx context.Context, reqOpts ...goblox.RequestOpt) (*BadgeMetadata, *Response, error) {
	metadata := new(BadgeMetadata)
	resp, err := s.client.Call(ctx, GetBadgeMetadata.Compile(nil), nil, metadata, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return metadata, resp, nil
}

type BadgeSortBy string

const (
	BadgeSortByRank        BadgeSortBy = "Rank"
	BadgeSortByDateCreated BadgeSortBy = "DateCreated"
)

type GetUniverseBadgesOptions struct {
	SortBy    BadgeSortBy `url:"sortBy,omitempty"`
	Limit     *int        `url:"limit,omitempty"` // 10, 25, 50, or 100.
	Cursor    *string     `url:"cursor,omitempty"`
	SortOrder SortOrder   `url:"sortOrder,omitempty"`
}

// GetUniverseBadges will get a page of the badges in a universe.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_get_v1_universes__universeId__badges
//
// [GET] https://badges.roblox.com/v1/universes/{universeId}/badges
func (s *BadgesService) GetUniverseBadges(ctx context.Context, universeId string, opts *GetUniverseBadgesOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[Badge], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[Badge])
	resp, err := s.client.Call(ctx, GetUniverseBadges.Compile(q, universeId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

// BadgePaymentSourceType is the payment source to charge when creating a badge: 1 or 2.
type BadgePaymentSourceType int

type CreateBadgeRequest struct {
	Name              *string                 `form:"name"`
	Description       *string                 `form:"description"`
	PaymentSourceType *BadgePaymentSourceType `form:"paymentSourceType"`
	Files             *os.File                `form:"files"` // The badge icon. The default icon is used when nil.
	ExpectedCost      *int64                  `form:"expectedCost"`
	IsActive          *bool                   `form:"isActive"`
}

// CreateBadge will create a badge in a universe and charge the payment source.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_post_v1_universes__universeId__badges
//
// [POST] https://badges.roblox.com/v1/universes/{universeId}/badges
func (s *BadgesService) CreateBadge(ctx context.Context, universeId string, data CreateBadgeRequest, reqOpts ...goblox.RequestOpt) (*Badge, *Response, error) {
	badge := new(Badge)
	e := CreateBadge.Compile(nil, universeId)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Name != nil {
		_ = writer.WriteField("name", *data.Name)
	}
	if data.Description != nil {
		_ = writer.WriteField("description", *data.Description)
	}
	if data.PaymentSourceType != nil {
		_ = writer.WriteField("paymentSourceType", fmt.Sprint(int(*data.PaymentSourceType)))
	}
	if data.Files != nil {
		part, err := writer.CreateFormFile("files", filepath.Base(data.Files.Name()))
		if err != nil {
			return nil, nil, err
		}
		if _, err := io.Copy(part, data.Files); err != nil {
			return nil, nil, err
		}
	}
	if data.ExpectedCost != nil {
		_ = writer.WriteField("expectedCost", fmt.Sprint(*data.ExpectedCost))
	}
	if data.IsActive != nil {
		_ = writer.WriteField("isActive", fmt.Sprint(*data.IsActive))
	}
	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewMultipartRequest(e, &body, writer.FormDataContentType())
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.CallRequest(ctx, e, req, badge, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return badge, resp, nil
}

// GetFreeBadgesQuota will get the number of free badges a universe has left today (UTC).
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_get_v1_universes__universeId__free_badges_quota
//
// [GET] https://badges.roblox.com/v1/universes/{universeId}/free-badges-quota
func (s *BadgesService) GetFreeBadgesQuota(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	var quota int64
	resp, err := s.client.Call(ctx, GetFreeBadgesQuota.Compile(nil, universeId), nil, &quota, reqOpts...)
	if err != nil {
		return 0, resp, err
	}

	return quota, resp, nil
}

// GetUserBadges will get a page of the badges a user has been awarded.
// Results respect the user's inventory privacy settings.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_get_v1_users__userId__badges
//
// [GET] https://badges.roblox.com/v1/users/{userId}/badges
func (s *BadgesService) GetUserBadges(ctx context.Context, userId string, opts *PageOptions, reqOpts ...goblox.RequestOpt) (*PageResponse[Badge], *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	page := new(PageResponse[Badge])
	resp, err := s.client.Call(ctx, GetUserBadges.Compile(q, userId), nil, page, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

type BadgeAward struct {
	BadgeID     int64  `json:"badgeId"`
	AwardedDate string `json:"awardedDate"`
}

// GetUserBadgeAwardedDate will get when a badge was awarded to a user.
// It returns a nil BadgeAward when the user doesn't have the badge.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_get_v1_users__userId__badges__badgeId__awarded_date
//
// [GET] https://badges.roblox.com/v1/users/{userId}/badges/{badgeId}/awarded-date
func (s *BadgesService) GetUserBadgeAwardedDate(ctx context.Context, userId string, badgeId string, reqOpts ...goblox.RequestOpt) (*BadgeAward, *Response, error) {
	award := new(BadgeAward)
	resp, err := s.client.Call(ctx, GetUserBadgeAwardedDate.Compile(nil, userId, badgeId), nil, award, reqOpts...)
	if err != nil {
		return nil, resp, err
	}
	if resp.StatusCode == http.StatusNoContent {
		return nil, resp, nil
	}

	return award, resp, nil
}

type GetUserBadgeAwardedDatesOptions struct {
	BadgeIDs []int64 `url:"badgeIds,comma"` // Required.
}

// GetUserBadgeAwardedDates will get when each of a list of badges was awarded to a user.
// Badges the user doesn't have are left out of the results.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_get_v1_users__userId__badges_awarded_dates
//
// [GET] https://badges.roblox.com/v1/users/{userId}/badges/awarded-dates
func (s *BadgesService) GetUserBadgeAwardedDates(ctx context.Context, userId string, opts *GetUserBadgeAwardedDatesOptions, reqOpts ...goblox.RequestOpt) ([]BadgeAward, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	results := new(WrappedResponse[[]BadgeAward])
	resp, err := s.client.Call(ctx, GetUserBadgeAwardedDates.Compile(q, userId), nil, results, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return results.Data, resp, nil
}

type UpdateBadgeIconRequest struct {
	Files *os.File `form:"Files"`
}

// UpdateBadgeIcon will replace a badge's icon and return the asset ID of the new icon.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_post_v1_badges__badgeId__icon
//
// [POST] https://badges.roblox.com/v1/badges/{badgeId}/icon
func (s *BadgesService) UpdateBadgeIcon(ctx context.Context, badgeId string, data UpdateBadgeIconRequest, reqOpts ...goblox.RequestOpt) (int64, *Response, error) {
	e := UpdateBadgeIcon.Compile(nil, badgeId)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Files != nil {
		part, err := writer.CreateFormFile("Files", filepath.Base(data.Files.Name()))
		if err != nil {
			return 0, nil, err
		}
		if _, err := io.Copy(part, data.Files); err != nil {
			return 0, nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return 0, nil, err
	}

	req, err := s.client.NewMultipartRequest(e, &body, writer.FormDataContentType())
	if err != nil {
		return 0, nil, err
	}

	var result struct {
		TargetID int64 `json:"targetId"`
	}
	resp, err := s.client.CallRequest(ctx, e, req, &result, reqOpts...)
	if err != nil {
		return 0, resp, err
	}

	return result.TargetID, resp, nil
}

// RemoveBadge will remove a badge from the authenticated user's account.
//
// Roblox API Docs: https://create.roblox.com/docs/cloud/reference/domains/badges#badges_delete_v1_user_badges__badgeId_
//
// [DELETE] https://badges.roblox.com/v1/user/badges/{badgeId}
func (s *BadgesService) RemoveBadge(ctx context.Context, badgeId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, RemoveBadge.Compile(nil, badgeId), nil, nil, reqOpts...)
}
