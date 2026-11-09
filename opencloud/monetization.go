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

	"github.com/luckfire/goblox"
	"github.com/luckfire/goblox/auth"
)

var (
	CreateCreatorStoreProduct = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/creator-store-products",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("creator-store-product:write")

	GetCreatorStoreProduct = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/creator-store-products/{productId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("creator-store-product:read")

	UpdateCreatorStoreProduct = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/creator-store-products/{productId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("creator-store-product:write")

	GetSubscription = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/subscription-products/{productId}/subscriptions/{subscriptionId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithAnyScopes("universe.subscription-product.subscription:read", "universe:write")
)

// MonetizationService will handle communciation with the actions related to the API.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type MonetizationService service

type CreatorStoreProductQuantity struct {
	Significand int `json:"significand"`
	Exponent    int `json:"exponent"`
}

type CreatorStoreProductPrice struct {
	CurrencyCode string                      `json:"currencyCode"`
	Quantity     CreatorStoreProductQuantity `json:"quantity"`
}

type CreatorStoreProductRestriction string

const (
	CreatorStoreProductRestrictionUnspecified                 CreatorStoreProductRestriction = "RESTRICTION_UNSPECIFIED"
	CreatorStoreProductRestrictionSoldItemRestricted          CreatorStoreProductRestriction = "SOLD_ITEM_RESTRICTED"
	CreatorStoreProductRestrictionSellerTemporarilyRestricted CreatorStoreProductRestriction = "SELLER_TEMPORARILY_RESTRICTED"
	CreatorStoreProductRestrictionSellerPermanentlyRestricted CreatorStoreProductRestriction = "SELLER_PERMANENTLY_RESTRICTED"
	CreatorStoreProductRestrictionSellerNoLongerActive        CreatorStoreProductRestriction = "SELLER_NO_LONGER_ACTIVE"
)

type CreatorStoreProduct struct {
	Path         string                           `json:"path"`
	BasePrice    CreatorStoreProductPrice         `json:"basePrice"`
	Published    bool                             `json:"published"`
	Restrictions []CreatorStoreProductRestriction `json:"restrictions"`
	Purchasable  bool                             `json:"purchasable"`
	UserSeller   bool                             `json:"userSeller"`
	ModelAssetID string                           `json:"modelAssetId"`
}

type CreatorStoreProductCreate struct {
	BasePrice     *CreatorStoreProductPrice `json:"basePrice,omitempty"`
	PurchasePrice *CreatorStoreProductPrice `json:"purchasePrice,omitempty"`
	Published     *bool                     `json:"published,omitempty"`
	ModelAssetID  *string                   `json:"modelAssetId,omitempty"`
}

// CreateCreatorStoreProduct will create a new store product on the creator store.
//
// Required scopes: creator-store-product:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/CreatorStoreProduct#Cloud_CreateCreatorStoreProduct
//
// [POST] /cloud/v2/creator-store-products
func (s *MonetizationService) CreateCreatorStoreProduct(ctx context.Context, data CreatorStoreProductCreate, reqOpts ...goblox.RequestOpt) (*CreatorStoreProduct, *Response, error) {
	product := new(CreatorStoreProduct)
	resp, err := s.client.Call(ctx, CreateCreatorStoreProduct.Compile(nil), data, product, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return product, resp, nil
}

// GetCreatorStoreProduct will fetch information on a specificed store product.
//
// Required scopes: creator-store-product:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/CreatorStoreProduct#Cloud_GetCreatorStoreProduct
//
// [GET] /cloud/v2/creator-store-products/{product_id}
func (s *MonetizationService) GetCreatorStoreProduct(ctx context.Context, productId string, reqOpts ...goblox.RequestOpt) (*CreatorStoreProduct, *Response, error) {
	product := new(CreatorStoreProduct)
	resp, err := s.client.Call(ctx, GetCreatorStoreProduct.Compile(nil, productId), nil, product, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return product, resp, nil
}

type CreatorStoreProductUpdate struct {
	BasePrice     *CreatorStoreProductPrice `json:"basePrice,omitempty"`
	PurchasePrice *CreatorStoreProductPrice `json:"purchasePrice,omitempty"`
	Published     *bool                     `json:"published,omitempty"`
}

// UpdateCreatorStoreProduct will update information for a specificed store product.
//
// Required scopes: creator-store-product:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/CreatorStoreProduct#Cloud_UpdateCreatorStoreProduct
//
// [PATCH] /cloud/v2/creator-store-products/{product_id}
func (s *MonetizationService) UpdateCreatorStoreProduct(ctx context.Context, productId string, data CreatorStoreProductUpdate, reqOpts ...goblox.RequestOpt) (*CreatorStoreProduct, *Response, error) {
	product := new(CreatorStoreProduct)
	resp, err := s.client.Call(ctx, UpdateCreatorStoreProduct.Compile(nil, productId), data, product, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return product, resp, nil
}

type SubscriptionState string

const (
	SubscriptionStateUnspecified                     SubscriptionState = "STATE_UNSPECIFIED"
	SubscriptionStateSubscribedWillRenew             SubscriptionState = "SUBSCRIBED_WILL_RENEW"
	SubscriptionStateSubscribedWillNotRenew          SubscriptionState = "SUBSCRIBED_WILL_NOT_RENEW"
	SubscriptionStateSubscribedRenewalPaymentPending SubscriptionState = "SUBSCRIBED_RENEWAL_PAYMENT_PENDING"
	SubscriptionStateExpired                         SubscriptionState = "EXPIRED"
)

type SubscriptionExpirationReason string

const (
	SubscriptionExpirationReasonUnspecified         SubscriptionExpirationReason = "EXPIRATION_REASON_UNSPECIFIED"
	SubscriptionExpirationReasonProductInactive     SubscriptionExpirationReason = "PRODUCT_INACTIVE"
	SubscriptionExpirationReasonProductDeleted      SubscriptionExpirationReason = "PRODUCT_DELETED"
	SubscriptionExpirationReasonSubscriberCancelled SubscriptionExpirationReason = "SUBSCRIBER_CANCELLED"
	SubscriptionExpirationReasonSubscriberRefunded  SubscriptionExpirationReason = "SUBSCRIBER_REFUNDED"
	SubscriptionExpirationReasonLapsed              SubscriptionExpirationReason = "LAPSED"
)

type SubscriptionExpirationDetails struct {
	Reason SubscriptionExpirationReason `json:"reason"`
}

type SubscriptionPurchasePlatform string

const (
	SubscriptionPurchasePlatformUnspecified SubscriptionPurchasePlatform = "PURCHASE_PLATFORM_UNSPECIFIED"
	SubscriptionPurchasePlatformDesktop     SubscriptionPurchasePlatform = "DESKTOP"
	SubscriptionPurchasePlatformMobile      SubscriptionPurchasePlatform = "MOBILE"
)

type SubscriptionPaymentProvider string

const (
	SubscriptionPaymentProviderUnspecified  SubscriptionPaymentProvider = "PAYMENT_PROVIDER_UNSPECIFIED"
	SubscriptionPaymentProviderStripe       SubscriptionPaymentProvider = "STRIPE"
	SubscriptionPaymentProviderApple        SubscriptionPaymentProvider = "APPLE"
	SubscriptionPaymentProviderGoogle       SubscriptionPaymentProvider = "GOOGLE"
	SubscriptionPaymentProviderRobloxCredit SubscriptionPaymentProvider = "ROBLOX_CREDIT"
)

type Subscription struct {
	Path              string                        `json:"path"`
	CreateTime        string                        `json:"createTime"`
	UpdateTime        string                        `json:"updateTime"`
	Active            bool                          `json:"active"`
	WillRenew         bool                          `json:"willRenew"`
	LastBillingTime   string                        `json:"lastBillingTime"`
	NextRenewTime     string                        `json:"nextRenewTime"`
	ExpireTime        string                        `json:"expireTime"`
	State             SubscriptionState             `json:"state"`
	ExpirationDetails SubscriptionExpirationDetails `json:"expirationDetails"`
	PurchasePlatform  SubscriptionPurchasePlatform  `json:"purchasePlatform"`
	PaymentProvider   SubscriptionPaymentProvider   `json:"paymentProvider"`
	User              string                        `json:"user"`
}

type SubscriptionView string

const (
	SubscriptionViewUnspecified SubscriptionView = "VIEW_UNSPECIFIED"
	SubscriptionViewBasic       SubscriptionView = "BASIC"
	SubscriptionViewFull        SubscriptionView = "FULL"
)

type SubscriptionOpts struct {
	View *SubscriptionView `url:"view,omitempty"`
}

// GetSubscription will fetch information on a specificed subscription.
//
// Required scopes:
//
// - universe:write
//
// - universe:writeuniverse.subscription-product.subscription:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/Subscription#Cloud_GetSubscription
//
// [GET] /cloud/v2/universes/{universe_id}/subscription-products/{product_id}/subscriptions/{subscription_id}
func (s *MonetizationService) GetSubscription(ctx context.Context, universeId, productId, subscriptionId string, opts *SubscriptionOpts, reqOpts ...goblox.RequestOpt) (*Subscription, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	subscription := new(Subscription)
	resp, err := s.client.Call(ctx, GetSubscription.Compile(q, universeId, productId, subscriptionId), nil, subscription, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return subscription, resp, nil
}

var (
	CreateDeveloperProduct = goblox.NewEndpoint(
		http.MethodPost,
		"/developer-products/v2/universes/{universeId}/developer-products",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("developer-product:write")

	ListDeveloperProducts = goblox.NewEndpoint(
		http.MethodGet,
		"/developer-products/v2/universes/{universeId}/developer-products/creator",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("developer-product:read")

	UpdateDeveloperProduct = goblox.NewEndpoint(
		http.MethodPatch,
		"/developer-products/v2/universes/{universeId}/developer-products/{productId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("developer-product:write")

	GetDeveloperProduct = goblox.NewEndpoint(
		http.MethodGet,
		"/developer-products/v2/universes/{universeId}/developer-products/{productId}/creator",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("developer-product:read")

	CreateGamePass = goblox.NewEndpoint(
		http.MethodPost,
		"/game-passes/v1/universes/{universeId}/game-passes",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("game-pass:write")

	ListGamePasses = goblox.NewEndpoint(
		http.MethodGet,
		"/game-passes/v1/universes/{universeId}/game-passes/creator",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("game-pass:read")

	UpdateGamePass = goblox.NewEndpoint(
		http.MethodPatch,
		"/game-passes/v1/universes/{universeId}/game-passes/{gamePassId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("game-pass:write")

	GetGamePass = goblox.NewEndpoint(
		http.MethodGet,
		"/game-passes/v1/universes/{universeId}/game-passes/{gamePassId}/creator",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("game-pass:read")
)

type CreateDeveloperProductRequest struct {
	Name                     *string  `form:"name"`
	Description              *string  `form:"description"`
	IsForSale                *bool    `form:"isForSale"`
	Price                    *int64   `form:"price"`
	ImageFile                *os.File `form:"imageFile"`
	IsRegionalPricingEnabled *bool    `form:"isRegionalPricingEnabled"`
	IsManagedPricingEnabled  *bool    `form:"isManagedPricingEnabled"`
}

type DeveloperProductPricingFeature string

const (
	DeveloperProductPricingFeatureInvalid           DeveloperProductPricingFeature = "Invalid"
	DeveloperProductPricingFeaturePriceOptimization DeveloperProductPricingFeature = "PriceOptimization"
	DeveloperProductPricingFeatureUserFixedPrice    DeveloperProductPricingFeature = "UserFixedPrice"
	DeveloperProductPricingFeatureRegionalPricing   DeveloperProductPricingFeature = "RegionalPricing"
)

type DeveloperProductPriceInformation struct {
	DefaultPriceInRobux int64                            `json:"defaultPriceInRobux"`
	EnabledFeatures     []DeveloperProductPricingFeature `json:"enabledFeatures"`
}

type DeveloperProduct struct {
	ProductID               int64                            `json:"productId"`
	Name                    string                           `json:"name"`
	Description             string                           `json:"description"`
	IconImageAssetID        int64                            `json:"iconImageAssetId"`
	UniverseID              int64                            `json:"universeId"`
	IsForSale               bool                             `json:"isForSale"`
	StorePageEnabled        bool                             `json:"storePageEnabled"`
	PriceInformation        DeveloperProductPriceInformation `json:"priceInformation"`
	IsImmutable             bool                             `json:"isImmutable"`
	CreatedTimestamp        string                           `json:"createdTimestamp"`
	UpdatedTimestamp        string                           `json:"updatedTimestamp"`
	IsManagedPricingEnabled bool                             `json:"isManagedPricingEnabled"`
}

// CreateDeveloperProduct will create developer product.
//
// Required scopes: developer-product:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#DeveloperProducts_CreateDeveloperProductV2
//
// [POST] /developer-products/v2/universes/{universeId}/developer-products
func (s *MonetizationService) CreateDeveloperProduct(ctx context.Context, universeId string, data CreateDeveloperProductRequest, reqOpts ...goblox.RequestOpt) (*DeveloperProduct, *Response, error) {
	developerProduct := new(DeveloperProduct)
	e := CreateDeveloperProduct.Compile(nil, universeId)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Name != nil {
		_ = writer.WriteField("name", fmt.Sprint(*data.Name))
	}
	if data.Description != nil {
		_ = writer.WriteField("description", fmt.Sprint(*data.Description))
	}
	if data.IsForSale != nil {
		_ = writer.WriteField("isForSale", fmt.Sprint(*data.IsForSale))
	}
	if data.Price != nil {
		_ = writer.WriteField("price", fmt.Sprint(*data.Price))
	}
	if data.ImageFile != nil {
		part, err := writer.CreateFormFile("imageFile", filepath.Base(data.ImageFile.Name()))
		if err != nil {
			return nil, nil, err
		}
		if _, err := io.Copy(part, data.ImageFile); err != nil {
			return nil, nil, err
		}
	}
	if data.IsRegionalPricingEnabled != nil {
		_ = writer.WriteField("isRegionalPricingEnabled", fmt.Sprint(*data.IsRegionalPricingEnabled))
	}
	if data.IsManagedPricingEnabled != nil {
		_ = writer.WriteField("isManagedPricingEnabled", fmt.Sprint(*data.IsManagedPricingEnabled))
	}
	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, writer.FormDataContentType())
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.CallRequest(ctx, e, req, developerProduct, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return developerProduct, resp, nil
}

type ListDeveloperProductsOptions struct {
	PageSize  *int    `url:"pageSize,omitempty"`
	PageToken *string `url:"pageToken,omitempty"`
}

type DeveloperProductList struct {
	DeveloperProducts []DeveloperProduct `json:"developerProducts"`
	NextPageToken     string             `json:"nextPageToken"`
}

// ListDeveloperProducts will list developer products by universe with configuration details.
//
// Required scopes: developer-product:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#DeveloperProducts_ListDeveloperProductConfigsByUniverseV2
//
// [GET] /developer-products/v2/universes/{universeId}/developer-products/creator
func (s *MonetizationService) ListDeveloperProducts(ctx context.Context, universeId string, opts *ListDeveloperProductsOptions, reqOpts ...goblox.RequestOpt) (*DeveloperProductList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	developerProductList := new(DeveloperProductList)
	resp, err := s.client.Call(ctx, ListDeveloperProducts.Compile(q, universeId), nil, developerProductList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return developerProductList, resp, nil
}

type UpdateDeveloperProductRequest struct {
	Name                     *string  `form:"name"`
	Description              *string  `form:"description"`
	IsForSale                *bool    `form:"isForSale"`
	Price                    *int64   `form:"price"`
	ImageFile                *os.File `form:"imageFile"`
	IsRegionalPricingEnabled *bool    `form:"isRegionalPricingEnabled"`
	StorePageEnabled         *bool    `form:"storePageEnabled"`
	IsManagedPricingEnabled  *bool    `form:"isManagedPricingEnabled"`
}

// UpdateDeveloperProduct will update developer product.
//
// Required scopes: developer-product:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#DeveloperProducts_UpdateDeveloperProductV2
//
// [PATCH] /developer-products/v2/universes/{universeId}/developer-products/{productId}
func (s *MonetizationService) UpdateDeveloperProduct(ctx context.Context, universeId, productId string, data UpdateDeveloperProductRequest, reqOpts ...goblox.RequestOpt) (*Response, error) {
	e := UpdateDeveloperProduct.Compile(nil, universeId, productId)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Name != nil {
		_ = writer.WriteField("name", fmt.Sprint(*data.Name))
	}
	if data.Description != nil {
		_ = writer.WriteField("description", fmt.Sprint(*data.Description))
	}
	if data.IsForSale != nil {
		_ = writer.WriteField("isForSale", fmt.Sprint(*data.IsForSale))
	}
	if data.Price != nil {
		_ = writer.WriteField("price", fmt.Sprint(*data.Price))
	}
	if data.ImageFile != nil {
		part, err := writer.CreateFormFile("imageFile", filepath.Base(data.ImageFile.Name()))
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(part, data.ImageFile); err != nil {
			return nil, err
		}
	}
	if data.IsRegionalPricingEnabled != nil {
		_ = writer.WriteField("isRegionalPricingEnabled", fmt.Sprint(*data.IsRegionalPricingEnabled))
	}
	if data.StorePageEnabled != nil {
		_ = writer.WriteField("storePageEnabled", fmt.Sprint(*data.StorePageEnabled))
	}
	if data.IsManagedPricingEnabled != nil {
		_ = writer.WriteField("isManagedPricingEnabled", fmt.Sprint(*data.IsManagedPricingEnabled))
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}

	return s.client.CallRequest(ctx, e, req, nil, reqOpts...)
}

// GetDeveloperProduct will get developer product with configuration details.
//
// Required scopes: developer-product:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/developer-products#DeveloperProducts_GetDeveloperProductConfigV2
//
// [GET] /developer-products/v2/universes/{universeId}/developer-products/{productId}/creator
func (s *MonetizationService) GetDeveloperProduct(ctx context.Context, universeId, productId string, reqOpts ...goblox.RequestOpt) (*DeveloperProduct, *Response, error) {
	developerProduct := new(DeveloperProduct)
	resp, err := s.client.Call(ctx, GetDeveloperProduct.Compile(nil, universeId, productId), nil, developerProduct, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return developerProduct, resp, nil
}

type CreateGamePassRequest struct {
	Name                     *string  `form:"name"`
	Description              *string  `form:"description"`
	ImageFile                *os.File `form:"imageFile"`
	IsForSale                *bool    `form:"isForSale"`
	Price                    *int64   `form:"price"`
	IsRegionalPricingEnabled *bool    `form:"isRegionalPricingEnabled"`
	IsManagedPricingEnabled  *bool    `form:"isManagedPricingEnabled"`
}

type GamePassPricingFeature string

const (
	GamePassPricingFeatureInvalid           GamePassPricingFeature = "Invalid"
	GamePassPricingFeaturePriceOptimization GamePassPricingFeature = "PriceOptimization"
	GamePassPricingFeatureUserFixedPrice    GamePassPricingFeature = "UserFixedPrice"
	GamePassPricingFeatureRegionalPricing   GamePassPricingFeature = "RegionalPricing"
)

type GamePassPriceInformation struct {
	DefaultPriceInRobux int64                    `json:"defaultPriceInRobux"`
	EnabledFeatures     []GamePassPricingFeature `json:"enabledFeatures"`
}

type GamePass struct {
	GamePassID              int64                    `json:"gamePassId"`
	Name                    string                   `json:"name"`
	Description             string                   `json:"description"`
	IsForSale               bool                     `json:"isForSale"`
	IconAssetID             int64                    `json:"iconAssetId"`
	CreatedTimestamp        string                   `json:"createdTimestamp"`
	UpdatedTimestamp        string                   `json:"updatedTimestamp"`
	PriceInformation        GamePassPriceInformation `json:"priceInformation"`
	IsManagedPricingEnabled bool                     `json:"isManagedPricingEnabled"`
}

// CreateGamePass will create game pass.
//
// Required scopes: game-pass:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#GamePasses_CreateGamePass
//
// [POST] /game-passes/v1/universes/{universeId}/game-passes
func (s *MonetizationService) CreateGamePass(ctx context.Context, universeId string, data CreateGamePassRequest, reqOpts ...goblox.RequestOpt) (*GamePass, *Response, error) {
	gamePass := new(GamePass)
	e := CreateGamePass.Compile(nil, universeId)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Name != nil {
		_ = writer.WriteField("name", fmt.Sprint(*data.Name))
	}
	if data.Description != nil {
		_ = writer.WriteField("description", fmt.Sprint(*data.Description))
	}
	if data.ImageFile != nil {
		part, err := writer.CreateFormFile("imageFile", filepath.Base(data.ImageFile.Name()))
		if err != nil {
			return nil, nil, err
		}
		if _, err := io.Copy(part, data.ImageFile); err != nil {
			return nil, nil, err
		}
	}
	if data.IsForSale != nil {
		_ = writer.WriteField("isForSale", fmt.Sprint(*data.IsForSale))
	}
	if data.Price != nil {
		_ = writer.WriteField("price", fmt.Sprint(*data.Price))
	}
	if data.IsRegionalPricingEnabled != nil {
		_ = writer.WriteField("isRegionalPricingEnabled", fmt.Sprint(*data.IsRegionalPricingEnabled))
	}
	if data.IsManagedPricingEnabled != nil {
		_ = writer.WriteField("isManagedPricingEnabled", fmt.Sprint(*data.IsManagedPricingEnabled))
	}
	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, writer.FormDataContentType())
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.CallRequest(ctx, e, req, gamePass, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return gamePass, resp, nil
}

type ListGamePassesOptions struct {
	PageSize  *int    `url:"pageSize,omitempty"`
	PageToken *string `url:"pageToken,omitempty"`
}

type GamePassList struct {
	GamePasses    []GamePass `json:"gamePasses"`
	NextPageToken string     `json:"nextPageToken"`
}

// ListGamePasses will list game passes by universe with configuration details.
//
// Required scopes: game-pass:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#GamePasses_ListGamePassConfigsByUniverse
//
// [GET] /game-passes/v1/universes/{universeId}/game-passes/creator
func (s *MonetizationService) ListGamePasses(ctx context.Context, universeId string, opts *ListGamePassesOptions, reqOpts ...goblox.RequestOpt) (*GamePassList, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	gamePassList := new(GamePassList)
	resp, err := s.client.Call(ctx, ListGamePasses.Compile(q, universeId), nil, gamePassList, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return gamePassList, resp, nil
}

type UpdateGamePassRequest struct {
	Name                     *string  `form:"name"`
	Description              *string  `form:"description"`
	File                     *os.File `form:"file"`
	ImageFile                *os.File `form:"imageFile"`
	IsForSale                *bool    `form:"isForSale"`
	Price                    *int64   `form:"price"`
	IsRegionalPricingEnabled *bool    `form:"isRegionalPricingEnabled"`
	IsManagedPricingEnabled  *bool    `form:"isManagedPricingEnabled"`
}

// UpdateGamePass will update game pass.
//
// Required scopes: game-pass:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#GamePasses_UpdateGamePass
//
// [PATCH] /game-passes/v1/universes/{universeId}/game-passes/{gamePassId}
func (s *MonetizationService) UpdateGamePass(ctx context.Context, universeId, gamePassId string, data UpdateGamePassRequest, reqOpts ...goblox.RequestOpt) (*Response, error) {
	e := UpdateGamePass.Compile(nil, universeId, gamePassId)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if data.Name != nil {
		_ = writer.WriteField("name", fmt.Sprint(*data.Name))
	}
	if data.Description != nil {
		_ = writer.WriteField("description", fmt.Sprint(*data.Description))
	}
	if data.File != nil {
		part, err := writer.CreateFormFile("file", filepath.Base(data.File.Name()))
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(part, data.File); err != nil {
			return nil, err
		}
	}
	if data.ImageFile != nil {
		part, err := writer.CreateFormFile("imageFile", filepath.Base(data.ImageFile.Name()))
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(part, data.ImageFile); err != nil {
			return nil, err
		}
	}
	if data.IsForSale != nil {
		_ = writer.WriteField("isForSale", fmt.Sprint(*data.IsForSale))
	}
	if data.Price != nil {
		_ = writer.WriteField("price", fmt.Sprint(*data.Price))
	}
	if data.IsRegionalPricingEnabled != nil {
		_ = writer.WriteField("isRegionalPricingEnabled", fmt.Sprint(*data.IsRegionalPricingEnabled))
	}
	if data.IsManagedPricingEnabled != nil {
		_ = writer.WriteField("isManagedPricingEnabled", fmt.Sprint(*data.IsManagedPricingEnabled))
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := s.client.NewMultipartRequest(e.Endpoint.Method, e.Path, &body, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}

	return s.client.CallRequest(ctx, e, req, nil, reqOpts...)
}

// GetGamePass will get game pass with configuration details.
//
// Required scopes: game-pass:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/game-passes#GamePasses_GetGamePassConfig
//
// [GET] /game-passes/v1/universes/{universeId}/game-passes/{gamePassId}/creator
func (s *MonetizationService) GetGamePass(ctx context.Context, universeId, gamePassId string, reqOpts ...goblox.RequestOpt) (*GamePass, *Response, error) {
	gamePass := new(GamePass)
	resp, err := s.client.Call(ctx, GetGamePass.Compile(nil, universeId, gamePassId), nil, gamePass, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return gamePass, resp, nil
}
