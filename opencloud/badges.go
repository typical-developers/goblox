package opencloud

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
	UpdateBadge = goblox.NewEndpoint(
		http.MethodPatch,
		"/legacy-badges/v1/badges/{badgeId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe.badge:write", "legacy-universe.badge:manage-and-spend-robux")

	CreateBadge = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-badges/v1/universes/{universeId}/badges",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe.badge:manage-and-spend-robux")

	UpdateBadgeIcon = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-publish/v1/badges/{badgeId}/icon",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-badge:manage")
)

// BadgesService handles communication with the badge endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type BadgesService service

type UpdateBadgeRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
}

type EmptyResponse map[string]any

// UpdateBadge will update badge configuration.
//
// Required scopes:
//
// - legacy-universe.badge:write
//
// - legacy-universe.badge:manage-and-spend-robux
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/badges#patch_legacy_badges_v1_badges__badgeId_
//
// [PATCH] /legacy-badges/v1/badges/{badgeId}
func (s *BadgesService) UpdateBadge(ctx context.Context, badgeId string, data UpdateBadgeRequest, reqOpts ...goblox.RequestOpt) (*EmptyResponse, *Response, error) {
	emptyResponse := new(EmptyResponse)
	resp, err := s.client.Call(ctx, UpdateBadge.Compile(nil, badgeId), data, emptyResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return emptyResponse, resp, nil
}

type CreateBadgeRequestPaymentSourceType int

type CreateBadgeRequest struct {
	Name              *string                              `form:"name"`
	Description       *string                              `form:"description"`
	PaymentSourceType *CreateBadgeRequestPaymentSourceType `form:"paymentSourceType"`
	Files             *os.File                             `form:"files"`
	ExpectedCost      *int64                               `form:"expectedCost"`
	IsActive          *bool                                `form:"isActive"`
}

type RelatedBadgeAwarderTypeType int

type RelatedBadgeAwarderType struct {
	ID   int64                       `json:"id"`
	Type RelatedBadgeAwarderTypeType `json:"type"`
	Name string                      `json:"name"`
}

type BadgeAwardStatisticsResponse struct {
	PastDayAwardedCount int64   `json:"pastDayAwardedCount"`
	AwardedCount        int64   `json:"awardedCount"`
	WinRatePercentage   float64 `json:"winRatePercentage"`
}

type BadgeResponseV2 struct {
	ID                 int64                        `json:"id"`
	Name               string                       `json:"name"`
	Description        string                       `json:"description"`
	DisplayName        string                       `json:"displayName"`
	DisplayDescription string                       `json:"displayDescription"`
	Enabled            bool                         `json:"enabled"`
	IconImageID        int64                        `json:"iconImageId"`
	DisplayIconImageID int64                        `json:"displayIconImageId"`
	Awarder            RelatedBadgeAwarderType      `json:"awarder"`
	Statistics         BadgeAwardStatisticsResponse `json:"statistics"`
	Created            string                       `json:"created"`
	Updated            string                       `json:"updated"`
}

// CreateBadge will create a new badge.
//
// Required scopes: legacy-universe.badge:manage-and-spend-robux
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/badges#post_legacy_badges_v1_universes__universeId__badges
//
// [POST] /legacy-badges/v1/universes/{universeId}/badges
func (s *BadgesService) CreateBadge(ctx context.Context, universeId string, data CreateBadgeRequest, reqOpts ...goblox.RequestOpt) (*BadgeResponseV2, *Response, error) {
	badgeResponseV2 := new(BadgeResponseV2)
	e := CreateBadge.Compile(nil, universeId)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Name != nil {
		_ = writer.WriteField("name", fmt.Sprint(*data.Name))
	}
	if data.Description != nil {
		_ = writer.WriteField("description", fmt.Sprint(*data.Description))
	}
	if data.PaymentSourceType != nil {
		_ = writer.WriteField("paymentSourceType", fmt.Sprint(*data.PaymentSourceType))
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

	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, writer.FormDataContentType())
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.CallRequest(ctx, e, req, badgeResponseV2, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return badgeResponseV2, resp, nil
}

type UpdateBadgeIconRequest struct {
	Files *os.File `form:"Files"`
}

type UploadResponse struct {
	TargetID int64 `json:"targetId"`
}

// UpdateBadgeIcon will overwrite a badge icon with a new one.
//
// Required scopes: legacy-badge:manage
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/assets#post_legacy_publish_v1_badges__badgeId__icon
//
// [POST] /legacy-publish/v1/badges/{badgeId}/icon
func (s *BadgesService) UpdateBadgeIcon(ctx context.Context, badgeId string, data UpdateBadgeIconRequest, reqOpts ...goblox.RequestOpt) (*UploadResponse, *Response, error) {
	uploadResponse := new(UploadResponse)
	e := UpdateBadgeIcon.Compile(nil, badgeId)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Files != nil {
		part, err := writer.CreateFormFile("Files", filepath.Base(data.Files.Name()))
		if err != nil {
			return nil, nil, err
		}
		if _, err := io.Copy(part, data.Files); err != nil {
			return nil, nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, writer.FormDataContentType())
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.CallRequest(ctx, e, req, uploadResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return uploadResponse, resp, nil
}
