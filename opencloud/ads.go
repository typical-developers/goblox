package opencloud

import (
	"context"
	"fmt"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	ListAdvertisableUniverses = goblox.NewEndpoint(
		http.MethodGet,
		"/ads-management/v1/advertisable-universes",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("ad.campaign:read")

	ListAdBillingAccounts = goblox.NewEndpoint(
		http.MethodGet,
		"/ads-management/v1/billing-accounts",
		auth.KindAPIKey,
	).WithScopes("ad.billing:read")

	GetAdBillingAccount = goblox.NewEndpoint(
		http.MethodGet,
		"/ads-management/v1/billing-accounts/{id}",
		auth.KindAPIKey,
	).WithScopes("ad.billing:read")

	GetAdCampaignOptions = goblox.NewEndpoint(
		http.MethodGet,
		"/ads-management/v1/campaign-options",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("ad.campaign:read")

	ListAdCampaigns = goblox.NewEndpoint(
		http.MethodGet,
		"/ads-management/v1/campaigns",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("ad.campaign:read")

	CreateAdCampaign = goblox.NewEndpoint(
		http.MethodPost,
		"/ads-management/v1/campaigns",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("ad.campaign:write")

	GetAdCampaign = goblox.NewEndpoint(
		http.MethodGet,
		"/ads-management/v1/campaigns/{id}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("ad.campaign:read")

	UpdateAdCampaign = goblox.NewEndpoint(
		http.MethodPatch,
		"/ads-management/v1/campaigns/{id}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("ad.campaign:write")

	BatchGetAdCampaignStatus = goblox.NewEndpoint(
		http.MethodPost,
		"/ads-management/v1/campaigns:batchGetStatus",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("ad.campaign:read")

	ListAdCreatives = goblox.NewEndpoint(
		http.MethodGet,
		"/ads-management/v1/creatives",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("ad.campaign:read")
)

// AdsService handles communication with the ads management endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type AdsService service

type AdvertisableUniverse struct {
	UniverseID string `json:"universeId"`
}

type ListAdvertisableUniversesResponse struct {
	AdvertisableUniverses []AdvertisableUniverse `json:"advertisableUniverses"`
}

// ListAdvertisableUniverses will list advertisable universes.
//
// Required scopes: ad.campaign:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/advertising#get_ads_management_v1_advertisable_universes
//
// [GET] /ads-management/v1/advertisable-universes
func (s *AdsService) ListAdvertisableUniverses(ctx context.Context, reqOpts ...goblox.RequestOpt) (*ListAdvertisableUniversesResponse, *Response, error) {
	listAdvertisableUniversesResponse := new(ListAdvertisableUniversesResponse)
	resp, err := s.client.Call(ctx, ListAdvertisableUniverses.Compile(nil), nil, listAdvertisableUniversesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listAdvertisableUniversesResponse, resp, nil
}

type BillingAccountStatus string

const (
	BillingAccountStatusActive   BillingAccountStatus = "ACTIVE"
	BillingAccountStatusDisabled BillingAccountStatus = "DISABLED"
	BillingAccountStatusArchived BillingAccountStatus = "ARCHIVED"
)

type BillingAccountType string

const (
	BillingAccountTypeSelfService BillingAccountType = "SELF_SERVICE"
	BillingAccountTypeManaged     BillingAccountType = "MANAGED"
	BillingAccountTypeInternal    BillingAccountType = "INTERNAL"
)

type BillingAccount struct {
	CreateTime     string               `json:"createTime"`
	DisplayName    string               `json:"displayName"`
	ID             string               `json:"id"`
	OrganizationID string               `json:"organizationId"`
	Status         BillingAccountStatus `json:"status"`
	TimeZone       string               `json:"timeZone"`
	Type           BillingAccountType   `json:"type"`
	UpdateTime     string               `json:"updateTime"`
}

type ListBillingAccountsResponse struct {
	BillingAccounts []BillingAccount `json:"billingAccounts"`
	NextPageToken   string           `json:"nextPageToken"`
}

// ListAdBillingAccounts will list billing accounts.
//
// Required scopes: ad.billing:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/advertising#get_ads_management_v1_billing_accounts
//
// [GET] /ads-management/v1/billing-accounts
func (s *AdsService) ListAdBillingAccounts(ctx context.Context, opts *Options, reqOpts ...goblox.RequestOpt) (*ListBillingAccountsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	listBillingAccountsResponse := new(ListBillingAccountsResponse)
	resp, err := s.client.Call(ctx, ListAdBillingAccounts.Compile(q), nil, listBillingAccountsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listBillingAccountsResponse, resp, nil
}

// GetAdBillingAccount will get a billing account.
//
// Required scopes: ad.billing:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/advertising#get_ads_management_v1_billing_accounts__id_
//
// [GET] /ads-management/v1/billing-accounts/{id}
func (s *AdsService) GetAdBillingAccount(ctx context.Context, id string, reqOpts ...goblox.RequestOpt) (*BillingAccount, *Response, error) {
	billingAccount := new(BillingAccount)
	resp, err := s.client.Call(ctx, GetAdBillingAccount.Compile(nil, id), nil, billingAccount, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return billingAccount, resp, nil
}

type GetAdCampaignOptionsOptionsObjective string

const (
	GetAdCampaignOptionsOptionsObjectivePlays      GetAdCampaignOptionsOptionsObjective = "PLAYS"
	GetAdCampaignOptionsOptionsObjectiveEngagement GetAdCampaignOptionsOptionsObjective = "ENGAGEMENT"
)

type GetAdCampaignOptionsOptions struct {
	UniverseID *string                               `url:"universeId,omitempty"`
	Objective  *GetAdCampaignOptionsOptionsObjective `url:"objective,omitempty"`
}

type AdFormat struct {
	Format string `json:"format"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type PublicUniverseEligibilityReason string

const (
	PublicUniverseEligibilityReasonNoPermission PublicUniverseEligibilityReason = "NO_PERMISSION"
	PublicUniverseEligibilityReasonBlocked      PublicUniverseEligibilityReason = "BLOCKED"
)

type PublicUniverseEligibility struct {
	Eligible   bool                              `json:"eligible"`
	Reasons    []PublicUniverseEligibilityReason `json:"reasons"`
	UniverseID string                            `json:"universeId"`
}

type CampaignOptionsResponseObjective string

const (
	CampaignOptionsResponseObjectivePlays      CampaignOptionsResponseObjective = "PLAYS"
	CampaignOptionsResponseObjectiveEngagement CampaignOptionsResponseObjective = "ENGAGEMENT"
)

type CampaignOptionsResponsePaymentType string

const (
	CampaignOptionsResponsePaymentTypeCreditCard CampaignOptionsResponsePaymentType = "CREDIT_CARD"
	CampaignOptionsResponsePaymentTypeAdsCredit  CampaignOptionsResponsePaymentType = "ADS_CREDIT"
	CampaignOptionsResponsePaymentTypeInvoice    CampaignOptionsResponsePaymentType = "INVOICE"
)

type TargetingDimensionsAgeGroup string

const (
	TargetingDimensionsAgeGroupAge1317   TargetingDimensionsAgeGroup = "AGE_13_17"
	TargetingDimensionsAgeGroupAge1824   TargetingDimensionsAgeGroup = "AGE_18_24"
	TargetingDimensionsAgeGroupAge25Plus TargetingDimensionsAgeGroup = "AGE_25_PLUS"
)

type TargetingDimensionsDevice string

const (
	TargetingDimensionsDevicePhone   TargetingDimensionsDevice = "PHONE"
	TargetingDimensionsDeviceTablet  TargetingDimensionsDevice = "TABLET"
	TargetingDimensionsDeviceDesktop TargetingDimensionsDevice = "DESKTOP"
	TargetingDimensionsDeviceConsole TargetingDimensionsDevice = "CONSOLE"
)

type TargetingDimensions struct {
	AgeGroups []TargetingDimensionsAgeGroup `json:"ageGroups"`
	Countries []string                      `json:"countries"`
	Devices   []TargetingDimensionsDevice   `json:"devices"`
}

type CampaignOptionsResponse struct {
	AdFormats           []AdFormat                           `json:"adFormats"`
	Eligibility         PublicUniverseEligibility            `json:"eligibility"`
	Objectives          []CampaignOptionsResponseObjective   `json:"objectives"`
	PaymentTypes        []CampaignOptionsResponsePaymentType `json:"paymentTypes"`
	TargetingDimensions TargetingDimensions                  `json:"targetingDimensions"`
}

// GetAdCampaignOptions will get campaign options.
//
// Required scopes: ad.campaign:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/advertising#get_ads_management_v1_campaign_options
//
// [GET] /ads-management/v1/campaign-options
func (s *AdsService) GetAdCampaignOptions(ctx context.Context, opts *GetAdCampaignOptionsOptions, reqOpts ...goblox.RequestOpt) (*CampaignOptionsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	campaignOptionsResponse := new(CampaignOptionsResponse)
	resp, err := s.client.Call(ctx, GetAdCampaignOptions.Compile(q), nil, campaignOptionsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return campaignOptionsResponse, resp, nil
}

type AdBidStrategy string

const (
	AdBidStrategyAutomated AdBidStrategy = "AUTOMATED"
)

type AdBid struct {
	Strategy *AdBidStrategy `json:"strategy,omitempty"`
}

type AdBudgetType string

const (
	AdBudgetTypeDaily    AdBudgetType = "DAILY"
	AdBudgetTypeLifetime AdBudgetType = "LIFETIME"
)

type AdBudget struct {
	AmountMicros           *string       `json:"amountMicros,omitempty"`
	ScheduledAmountMicros  *string       `json:"scheduledAmountMicros,omitempty"`
	ScheduledEffectiveTime *string       `json:"scheduledEffectiveTime,omitempty"`
	Type                   *AdBudgetType `json:"type,omitempty"`
}

type AdCampaignDeliveryStatus string

const (
	AdCampaignDeliveryStatusServing    AdCampaignDeliveryStatus = "SERVING"
	AdCampaignDeliveryStatusInReview   AdCampaignDeliveryStatus = "IN_REVIEW"
	AdCampaignDeliveryStatusNotServing AdCampaignDeliveryStatus = "NOT_SERVING"
	AdCampaignDeliveryStatusRejected   AdCampaignDeliveryStatus = "REJECTED"
)

type AdCampaignDeliveryStatusReason string

const (
	AdCampaignDeliveryStatusReasonScheduled           AdCampaignDeliveryStatusReason = "SCHEDULED"
	AdCampaignDeliveryStatusReasonAutoPaused          AdCampaignDeliveryStatusReason = "AUTO_PAUSED"
	AdCampaignDeliveryStatusReasonPaused              AdCampaignDeliveryStatusReason = "PAUSED"
	AdCampaignDeliveryStatusReasonInactive            AdCampaignDeliveryStatusReason = "INACTIVE"
	AdCampaignDeliveryStatusReasonCompleted           AdCampaignDeliveryStatusReason = "COMPLETED"
	AdCampaignDeliveryStatusReasonExpired             AdCampaignDeliveryStatusReason = "EXPIRED"
	AdCampaignDeliveryStatusReasonCancelled           AdCampaignDeliveryStatusReason = "CANCELLED"
	AdCampaignDeliveryStatusReasonModerated           AdCampaignDeliveryStatusReason = "MODERATED"
	AdCampaignDeliveryStatusReasonClickBait           AdCampaignDeliveryStatusReason = "CLICK_BAIT"
	AdCampaignDeliveryStatusReasonGameFiltered        AdCampaignDeliveryStatusReason = "GAME_FILTERED"
	AdCampaignDeliveryStatusReasonPrivate             AdCampaignDeliveryStatusReason = "PRIVATE"
	AdCampaignDeliveryStatusReasonPlaceJoinRestricted AdCampaignDeliveryStatusReason = "PLACE_JOIN_RESTRICTED"
	AdCampaignDeliveryStatusReasonLearning            AdCampaignDeliveryStatusReason = "LEARNING"
)

type AdCampaignObjective string

const (
	AdCampaignObjectivePlays      AdCampaignObjective = "PLAYS"
	AdCampaignObjectiveEngagement AdCampaignObjective = "ENGAGEMENT"
	AdCampaignObjectiveAwareness  AdCampaignObjective = "AWARENESS"
	AdCampaignObjectiveReach      AdCampaignObjective = "REACH"
	AdCampaignObjectiveVideoViews AdCampaignObjective = "VIDEO_VIEWS"
	AdCampaignObjectiveRobuxSpend AdCampaignObjective = "ROBUX_SPEND"
)

type AdCampaignPaymentType string

const (
	AdCampaignPaymentTypeCreditCard AdCampaignPaymentType = "CREDIT_CARD"
	AdCampaignPaymentTypeAdsCredit  AdCampaignPaymentType = "ADS_CREDIT"
	AdCampaignPaymentTypeInvoice    AdCampaignPaymentType = "INVOICE"
)

type AdSchedule struct {
	DurationInDays *int    `json:"durationInDays,omitempty"`
	EndTime        *string `json:"endTime,omitempty"`
	StartTime      *string `json:"startTime,omitempty"`
}

type AdCampaignStatus string

const (
	AdCampaignStatusActive    AdCampaignStatus = "ACTIVE"
	AdCampaignStatusPaused    AdCampaignStatus = "PAUSED"
	AdCampaignStatusCancelled AdCampaignStatus = "CANCELLED"
)

type AdTargetingAgeGroup string

const (
	AdTargetingAgeGroupAge1317   AdTargetingAgeGroup = "AGE_13_17"
	AdTargetingAgeGroupAge1824   AdTargetingAgeGroup = "AGE_18_24"
	AdTargetingAgeGroupAge25Plus AdTargetingAgeGroup = "AGE_25_PLUS"
)

type AdTargetingDevice string

const (
	AdTargetingDevicePhone   AdTargetingDevice = "PHONE"
	AdTargetingDeviceTablet  AdTargetingDevice = "TABLET"
	AdTargetingDeviceDesktop AdTargetingDevice = "DESKTOP"
	AdTargetingDeviceConsole AdTargetingDevice = "CONSOLE"
)

type AdTargeting struct {
	AgeGroups []AdTargetingAgeGroup `json:"ageGroups,omitempty"`
	Countries []string              `json:"countries,omitempty"`
	Devices   []AdTargetingDevice   `json:"devices,omitempty"`
}

type AdCampaign struct {
	Bid                   AdBid                            `json:"bid"`
	BillingAccountID      string                           `json:"billingAccountId"`
	Budget                AdBudget                         `json:"budget"`
	CreateTime            string                           `json:"createTime"`
	CreativeAssetIDs      []string                         `json:"creativeAssetIds"`
	DeliveryStatus        AdCampaignDeliveryStatus         `json:"deliveryStatus"`
	DeliveryStatusReasons []AdCampaignDeliveryStatusReason `json:"deliveryStatusReasons"`
	ID                    string                           `json:"id"`
	Name                  string                           `json:"name"`
	Objective             AdCampaignObjective              `json:"objective"`
	PaymentType           AdCampaignPaymentType            `json:"paymentType"`
	Schedule              AdSchedule                       `json:"schedule"`
	Status                AdCampaignStatus                 `json:"status"`
	TargetUniverseID      string                           `json:"targetUniverseId"`
	Targeting             AdTargeting                      `json:"targeting"`
	UpdateTime            string                           `json:"updateTime"`
}

type ListCampaignsResponse struct {
	Campaigns     []AdCampaign `json:"campaigns"`
	NextPageToken string       `json:"nextPageToken"`
}

// ListAdCampaigns will list campaigns.
//
// Required scopes: ad.campaign:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/advertising#get_ads_management_v1_campaigns
//
// [GET] /ads-management/v1/campaigns
func (s *AdsService) ListAdCampaigns(ctx context.Context, opts *Options, reqOpts ...goblox.RequestOpt) (*ListCampaignsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	listCampaignsResponse := new(ListCampaignsResponse)
	resp, err := s.client.Call(ctx, ListAdCampaigns.Compile(q), nil, listCampaignsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listCampaignsResponse, resp, nil
}

type CreateCampaignRequestObjective string

const (
	CreateCampaignRequestObjectivePlays      CreateCampaignRequestObjective = "PLAYS"
	CreateCampaignRequestObjectiveEngagement CreateCampaignRequestObjective = "ENGAGEMENT"
)

type CreateCampaignRequestPaymentType string

const (
	CreateCampaignRequestPaymentTypeCreditCard CreateCampaignRequestPaymentType = "CREDIT_CARD"
	CreateCampaignRequestPaymentTypeAdsCredit  CreateCampaignRequestPaymentType = "ADS_CREDIT"
	CreateCampaignRequestPaymentTypeInvoice    CreateCampaignRequestPaymentType = "INVOICE"
)

type CreateCampaignRequest struct {
	Bid              *AdBid                            `json:"bid,omitempty"`
	Budget           *AdBudget                         `json:"budget,omitempty"`
	CreativeAssetIDs []string                          `json:"creativeAssetIds,omitempty"`
	Name             *string                           `json:"name,omitempty"`
	Objective        *CreateCampaignRequestObjective   `json:"objective,omitempty"`
	PaymentType      *CreateCampaignRequestPaymentType `json:"paymentType,omitempty"`
	Schedule         *AdSchedule                       `json:"schedule,omitempty"`
	TargetUniverseID *string                           `json:"targetUniverseId,omitempty"`
	Targeting        *AdTargeting                      `json:"targeting,omitempty"`
}

type CreateAdCampaignOptions struct {
	XIdempotencyKey *string `url:"-"` // Required.
}

// CreateAdCampaign will create a campaign.
//
// Required scopes: ad.campaign:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/advertising#post_ads_management_v1_campaigns
//
// [POST] /ads-management/v1/campaigns
func (s *AdsService) CreateAdCampaign(ctx context.Context, data CreateCampaignRequest, opts *CreateAdCampaignOptions, reqOpts ...goblox.RequestOpt) (*AdCampaign, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}
	if opts != nil {
		if opts.XIdempotencyKey != nil {
			reqOpts = append(reqOpts, goblox.WithHeader("x-idempotency-key", fmt.Sprint(*opts.XIdempotencyKey)))
		}
	}

	adCampaign := new(AdCampaign)
	resp, err := s.client.Call(ctx, CreateAdCampaign.Compile(q), data, adCampaign, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return adCampaign, resp, nil
}

// GetAdCampaign will get a campaign.
//
// Required scopes: ad.campaign:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/advertising#get_ads_management_v1_campaigns__id_
//
// [GET] /ads-management/v1/campaigns/{id}
func (s *AdsService) GetAdCampaign(ctx context.Context, id string, reqOpts ...goblox.RequestOpt) (*AdCampaign, *Response, error) {
	adCampaign := new(AdCampaign)
	resp, err := s.client.Call(ctx, GetAdCampaign.Compile(nil, id), nil, adCampaign, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return adCampaign, resp, nil
}

type UpdateBudget struct {
	AmountMicros *string `json:"amountMicros,omitempty"`
}

type UpdateCampaignRequestStatus string

const (
	UpdateCampaignRequestStatusActive    UpdateCampaignRequestStatus = "ACTIVE"
	UpdateCampaignRequestStatusPaused    UpdateCampaignRequestStatus = "PAUSED"
	UpdateCampaignRequestStatusCancelled UpdateCampaignRequestStatus = "CANCELLED"
)

type UpdateCampaignRequest struct {
	Bid              *AdBid                       `json:"bid,omitempty"`
	Budget           *UpdateBudget                `json:"budget,omitempty"`
	CreativeAssetIDs []string                     `json:"creativeAssetIds,omitempty"`
	Name             *string                      `json:"name,omitempty"`
	Schedule         *AdSchedule                  `json:"schedule,omitempty"`
	Status           *UpdateCampaignRequestStatus `json:"status,omitempty"`
	Targeting        *AdTargeting                 `json:"targeting,omitempty"`
}

// UpdateAdCampaign will update a campaign.
//
// Required scopes: ad.campaign:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/advertising#patch_ads_management_v1_campaigns__id_
//
// [PATCH] /ads-management/v1/campaigns/{id}
func (s *AdsService) UpdateAdCampaign(ctx context.Context, id string, data UpdateCampaignRequest, reqOpts ...goblox.RequestOpt) (*AdCampaign, *Response, error) {
	adCampaign := new(AdCampaign)
	resp, err := s.client.Call(ctx, UpdateAdCampaign.Compile(nil, id), data, adCampaign, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return adCampaign, resp, nil
}

type BatchGetStatusRequest struct {
	CampaignIDs []string `json:"campaignIds,omitempty"`
}

type CampaignIDFailureReason string

const (
	CampaignIDFailureReasonNotFound CampaignIDFailureReason = "NOT_FOUND"
)

type CampaignIDFailure struct {
	ID     string                  `json:"id"`
	Reason CampaignIDFailureReason `json:"reason"`
}

type CampaignStatusDeliveryStatus string

const (
	CampaignStatusDeliveryStatusServing    CampaignStatusDeliveryStatus = "SERVING"
	CampaignStatusDeliveryStatusInReview   CampaignStatusDeliveryStatus = "IN_REVIEW"
	CampaignStatusDeliveryStatusNotServing CampaignStatusDeliveryStatus = "NOT_SERVING"
	CampaignStatusDeliveryStatusRejected   CampaignStatusDeliveryStatus = "REJECTED"
)

type CampaignStatusDeliveryStatusReason string

const (
	CampaignStatusDeliveryStatusReasonScheduled           CampaignStatusDeliveryStatusReason = "SCHEDULED"
	CampaignStatusDeliveryStatusReasonAutoPaused          CampaignStatusDeliveryStatusReason = "AUTO_PAUSED"
	CampaignStatusDeliveryStatusReasonPaused              CampaignStatusDeliveryStatusReason = "PAUSED"
	CampaignStatusDeliveryStatusReasonInactive            CampaignStatusDeliveryStatusReason = "INACTIVE"
	CampaignStatusDeliveryStatusReasonCompleted           CampaignStatusDeliveryStatusReason = "COMPLETED"
	CampaignStatusDeliveryStatusReasonExpired             CampaignStatusDeliveryStatusReason = "EXPIRED"
	CampaignStatusDeliveryStatusReasonCancelled           CampaignStatusDeliveryStatusReason = "CANCELLED"
	CampaignStatusDeliveryStatusReasonModerated           CampaignStatusDeliveryStatusReason = "MODERATED"
	CampaignStatusDeliveryStatusReasonClickBait           CampaignStatusDeliveryStatusReason = "CLICK_BAIT"
	CampaignStatusDeliveryStatusReasonGameFiltered        CampaignStatusDeliveryStatusReason = "GAME_FILTERED"
	CampaignStatusDeliveryStatusReasonPrivate             CampaignStatusDeliveryStatusReason = "PRIVATE"
	CampaignStatusDeliveryStatusReasonPlaceJoinRestricted CampaignStatusDeliveryStatusReason = "PLACE_JOIN_RESTRICTED"
	CampaignStatusDeliveryStatusReasonLearning            CampaignStatusDeliveryStatusReason = "LEARNING"
)

type CampaignStatusStatus string

const (
	CampaignStatusStatusActive    CampaignStatusStatus = "ACTIVE"
	CampaignStatusStatusPaused    CampaignStatusStatus = "PAUSED"
	CampaignStatusStatusCancelled CampaignStatusStatus = "CANCELLED"
)

type CampaignStatus struct {
	DeliveryStatus        CampaignStatusDeliveryStatus         `json:"deliveryStatus"`
	DeliveryStatusReasons []CampaignStatusDeliveryStatusReason `json:"deliveryStatusReasons"`
	ID                    string                               `json:"id"`
	Status                CampaignStatusStatus                 `json:"status"`
}

type BatchGetStatusResponse struct {
	Failures []CampaignIDFailure `json:"failures"`
	Statuses []CampaignStatus    `json:"statuses"`
}

// BatchGetAdCampaignStatus will batch get campaign status.
//
// Required scopes: ad.campaign:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/advertising#post_ads_management_v1_campaigns_batchGetStatus
//
// [POST] /ads-management/v1/campaigns:batchGetStatus
func (s *AdsService) BatchGetAdCampaignStatus(ctx context.Context, data BatchGetStatusRequest, reqOpts ...goblox.RequestOpt) (*BatchGetStatusResponse, *Response, error) {
	batchGetStatusResponse := new(BatchGetStatusResponse)
	resp, err := s.client.Call(ctx, BatchGetAdCampaignStatus.Compile(nil), data, batchGetStatusResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return batchGetStatusResponse, resp, nil
}

type ListAdCreativesOptions struct {
	IsArchived  *bool   `url:"isArchived,omitempty"`
	MaxPageSize *int    `url:"maxPageSize,omitempty"`
	PageToken   *string `url:"pageToken,omitempty"`
}

type AdCreativeAssetType string

const (
	AdCreativeAssetTypeImage AdCreativeAssetType = "IMAGE"
)

type AdCreativeModerationStatus string

const (
	AdCreativeModerationStatusPendingReview AdCreativeModerationStatus = "PENDING_REVIEW"
	AdCreativeModerationStatusApproved      AdCreativeModerationStatus = "APPROVED"
	AdCreativeModerationStatusRejected      AdCreativeModerationStatus = "REJECTED"
)

type AdCreativeSource string

const (
	AdCreativeSourceUpload       AdCreativeSource = "UPLOAD"
	AdCreativeSourceInHouseGenAi AdCreativeSource = "IN_HOUSE_GEN_AI"
)

type AdCreative struct {
	AssetID          string                     `json:"assetId"`
	AssetName        string                     `json:"assetName"`
	AssetType        AdCreativeAssetType        `json:"assetType"`
	CreateTime       string                     `json:"createTime"`
	Height           int                        `json:"height"`
	ID               string                     `json:"id"`
	IsArchived       bool                       `json:"isArchived"`
	ModerationStatus AdCreativeModerationStatus `json:"moderationStatus"`
	Source           AdCreativeSource           `json:"source"`
	UniverseID       string                     `json:"universeId"`
	UpdateTime       string                     `json:"updateTime"`
	Width            int                        `json:"width"`
}

type ListCreativesResponse struct {
	Creatives     []AdCreative `json:"creatives"`
	NextPageToken string       `json:"nextPageToken"`
}

// ListAdCreatives will list creatives.
//
// Required scopes: ad.campaign:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/advertising#get_ads_management_v1_creatives
//
// [GET] /ads-management/v1/creatives
func (s *AdsService) ListAdCreatives(ctx context.Context, opts *ListAdCreativesOptions, reqOpts ...goblox.RequestOpt) (*ListCreativesResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	listCreativesResponse := new(ListCreativesResponse)
	resp, err := s.client.Call(ctx, ListAdCreatives.Compile(q), nil, listCreativesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listCreativesResponse, resp, nil
}
