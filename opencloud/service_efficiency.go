package opencloud

import (
	"context"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	GetAccountForGroup = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/accounts/groups/{groupId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	UpdateAccountForGroup = goblox.NewEndpoint(
		http.MethodPut,
		"/service-efficiency-api/v1/accounts/groups/{groupId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:write")

	GetAccountForUser = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/accounts/users/{userId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	UpdateAccountForUser = goblox.NewEndpoint(
		http.MethodPut,
		"/service-efficiency-api/v1/accounts/users/{userId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:write")

	GetBillingActivitiesByGroup = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/billing_activities/groups/{groupId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	GetBillingActivitiesByUser = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/billing_activities/users/{userId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	GetBillByGroup = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/bills/groups/{groupId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	GetBillByUser = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/bills/users/{userId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	GetConsumerConfigurationsByUniverse = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/configurations/universes/{universeId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	UpdateConsumerConfigurationsByUniverse = goblox.NewEndpoint(
		http.MethodPatch,
		"/service-efficiency-api/v1/configurations/universes/{universeId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:write")

	CheckEligibilityForGroup = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/eligibility/groups/{groupId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	CheckEligibilityForUser = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/eligibility/users/{userId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	IsAllowed = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/is_allowed",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	GetServiceEfficiencyPaymentMethodForGroup = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/payment_profiles/groups/{groupId}/service_efficiency",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	SetServiceEfficiencyPaymentMethodForGroup = goblox.NewEndpoint(
		http.MethodPut,
		"/service-efficiency-api/v1/payment_profiles/groups/{groupId}/service_efficiency",
		auth.KindAPIKey,
	).WithScopes("extended-services:write")

	DeletePaymentProfileForUser = goblox.NewEndpoint(
		http.MethodDelete,
		"/service-efficiency-api/v1/payment_profiles/users/{userId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:write")

	ListPaymentProfilesByUser = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/payment_profiles/users/{userId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	CreatePaymentProfileForUser = goblox.NewEndpoint(
		http.MethodPost,
		"/service-efficiency-api/v1/payment_profiles/users/{userId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:write")

	GetServiceEfficiencyPaymentMethodForUser = goblox.NewEndpoint(
		http.MethodGet,
		"/service-efficiency-api/v1/payment_profiles/users/{userId}/service_efficiency",
		auth.KindAPIKey,
	).WithScopes("extended-services:read")

	SetServiceEfficiencyPaymentMethodForUser = goblox.NewEndpoint(
		http.MethodPut,
		"/service-efficiency-api/v1/payment_profiles/users/{userId}/service_efficiency",
		auth.KindAPIKey,
	).WithScopes("extended-services:write")

	SubmitManualPaymentForGroup = goblox.NewEndpoint(
		http.MethodPost,
		"/service-efficiency-api/v1/payments/groups/{groupId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:write")

	SubmitManualPaymentForUser = goblox.NewEndpoint(
		http.MethodPost,
		"/service-efficiency-api/v1/payments/users/{userId}",
		auth.KindAPIKey,
	).WithScopes("extended-services:write")
)

// ServiceEfficiencyService handles communication with the service efficiency (billing) endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type ServiceEfficiencyService service

type AccountTaxType string

const (
	AccountTaxTypeInvalid    AccountTaxType = "Invalid"
	AccountTaxTypeIndividual AccountTaxType = "Individual"
	AccountTaxTypeBusiness   AccountTaxType = "Business"
)

type TaxIDType string

const (
	TaxIDTypeInvalid                TaxIDType = "Invalid"
	TaxIDTypeUnitedStatesEin        TaxIDType = "UnitedStatesEin"
	TaxIDTypeEuropeanVatNumber      TaxIDType = "EuropeanVatNumber"
	TaxIDTypeUnitedKingdomVatNumber TaxIDType = "UnitedKingdomVatNumber"
	TaxIDTypeCanadianBusinessNumber TaxIDType = "CanadianBusinessNumber"
	TaxIDTypeCanadianGstHstNumber   TaxIDType = "CanadianGstHstNumber"
	TaxIDTypeMexicanRfcNumber       TaxIDType = "MexicanRfcNumber"
)

type ServiceEfficiencyAccount struct {
	AccountTaxType *AccountTaxType `json:"accountTaxType,omitempty"`
	AccountName    *string         `json:"accountName,omitempty"`
	TaxID          *string         `json:"taxId,omitempty"`
	TaxIDType      *TaxIDType      `json:"taxIdType,omitempty"`
}

type GetAccountResponse struct {
	Account ServiceEfficiencyAccount `json:"account"`
}

// GetAccountForGroup will get an Account for a Group.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_GetAccountForGroup
//
// [GET] /service-efficiency-api/v1/accounts/groups/{groupId}
func (s *ServiceEfficiencyService) GetAccountForGroup(ctx context.Context, groupId string, reqOpts ...goblox.RequestOpt) (*GetAccountResponse, *Response, error) {
	getAccountResponse := new(GetAccountResponse)
	resp, err := s.client.Call(ctx, GetAccountForGroup.Compile(nil, groupId), nil, getAccountResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getAccountResponse, resp, nil
}

type UpdateAccountRequest struct {
	Account *ServiceEfficiencyAccount `json:"account,omitempty"`
}

type UpdateAccountResponse struct {
	Account ServiceEfficiencyAccount `json:"account"`
}

// UpdateAccountForGroup will update an Account for a Group.
//
// Required scopes: extended-services:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_UpdateAccountForGroup
//
// [PUT] /service-efficiency-api/v1/accounts/groups/{groupId}
func (s *ServiceEfficiencyService) UpdateAccountForGroup(ctx context.Context, groupId string, data UpdateAccountRequest, reqOpts ...goblox.RequestOpt) (*UpdateAccountResponse, *Response, error) {
	updateAccountResponse := new(UpdateAccountResponse)
	resp, err := s.client.Call(ctx, UpdateAccountForGroup.Compile(nil, groupId), data, updateAccountResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateAccountResponse, resp, nil
}

// GetAccountForUser will get an Account for a User.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_GetAccountForUser
//
// [GET] /service-efficiency-api/v1/accounts/users/{userId}
func (s *ServiceEfficiencyService) GetAccountForUser(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (*GetAccountResponse, *Response, error) {
	getAccountResponse := new(GetAccountResponse)
	resp, err := s.client.Call(ctx, GetAccountForUser.Compile(nil, userId), nil, getAccountResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getAccountResponse, resp, nil
}

// UpdateAccountForUser will update an Account for a User.
//
// Required scopes: extended-services:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_UpdateAccountForUser
//
// [PUT] /service-efficiency-api/v1/accounts/users/{userId}
func (s *ServiceEfficiencyService) UpdateAccountForUser(ctx context.Context, userId string, data UpdateAccountRequest, reqOpts ...goblox.RequestOpt) (*UpdateAccountResponse, *Response, error) {
	updateAccountResponse := new(UpdateAccountResponse)
	resp, err := s.client.Call(ctx, UpdateAccountForUser.Compile(nil, userId), data, updateAccountResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateAccountResponse, resp, nil
}

type GetBillingActivitiesByGroupOptions struct {
	StartYear  int `url:"startYear"`  // Required.
	StartMonth int `url:"startMonth"` // Required.
	EndYear    int `url:"endYear"`    // Required.
	EndMonth   int `url:"endMonth"`   // Required.
}

type ResponseAccountState string

const (
	ResponseAccountStateInvalid   ResponseAccountState = "Invalid"
	ResponseAccountStateNormal    ResponseAccountState = "Normal"
	ResponseAccountStateOverdue   ResponseAccountState = "Overdue"
	ResponseAccountStateSuspended ResponseAccountState = "Suspended"
)

type BillStatus string

const (
	BillStatusInvalid BillStatus = "Invalid"
	BillStatusUnpaid  BillStatus = "Unpaid"
	BillStatusPaid    BillStatus = "Paid"
	BillStatusClosed  BillStatus = "Closed"
)

type ServiceEfficiencyMoney struct {
	CurrencyCode *string `json:"currencyCode,omitempty"`
	Units        *int64  `json:"units,omitempty"`
	Nanos        *int    `json:"nanos,omitempty"`
}

type BillSummary struct {
	Timestamp   int64                  `json:"timestamp"`
	BillID      string                 `json:"billId"`
	BillStatus  BillStatus             `json:"billStatus"`
	TotalAmount ServiceEfficiencyMoney `json:"totalAmount"`
}

type BillingActivitiesResponse struct {
	AccountState       ResponseAccountState   `json:"accountState"`
	BillSummaries      []BillSummary          `json:"billSummaries"`
	CurrentBalance     ServiceEfficiencyMoney `json:"currentBalance"`
	MonthToDateBalance ServiceEfficiencyMoney `json:"monthToDateBalance"`
	ChargingThreshold  ServiceEfficiencyMoney `json:"chargingThreshold"`
}

// GetBillingActivitiesByGroup will get Billing Activities.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_GetBillingActivitiesByGroup
//
// [GET] /service-efficiency-api/v1/billing_activities/groups/{groupId}
func (s *ServiceEfficiencyService) GetBillingActivitiesByGroup(ctx context.Context, groupId string, opts *GetBillingActivitiesByGroupOptions, reqOpts ...goblox.RequestOpt) (*BillingActivitiesResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	billingActivitiesResponse := new(BillingActivitiesResponse)
	resp, err := s.client.Call(ctx, GetBillingActivitiesByGroup.Compile(q, groupId), nil, billingActivitiesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return billingActivitiesResponse, resp, nil
}

type GetBillingActivitiesByUserOptions struct {
	StartYear  int `url:"startYear"`  // Required.
	StartMonth int `url:"startMonth"` // Required.
	EndYear    int `url:"endYear"`    // Required.
	EndMonth   int `url:"endMonth"`   // Required.
}

// GetBillingActivitiesByUser will get Billing Activities.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_GetBillingActivitiesByUser
//
// [GET] /service-efficiency-api/v1/billing_activities/users/{userId}
func (s *ServiceEfficiencyService) GetBillingActivitiesByUser(ctx context.Context, userId string, opts *GetBillingActivitiesByUserOptions, reqOpts ...goblox.RequestOpt) (*BillingActivitiesResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	billingActivitiesResponse := new(BillingActivitiesResponse)
	resp, err := s.client.Call(ctx, GetBillingActivitiesByUser.Compile(q, userId), nil, billingActivitiesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return billingActivitiesResponse, resp, nil
}

type GetBillByGroupOptions struct {
	BillID *string `url:"billId,omitempty"`
	Year   *int    `url:"year,omitempty"`
	Month  *int    `url:"month,omitempty"`
	Date   *int    `url:"date,omitempty"`
	Latest *bool   `url:"latest,omitempty"`
}

type ServiceEfficiencyPrice struct {
	Cost       *ServiceEfficiencyMoney `json:"cost,omitempty"`
	UnitAmount *int64                  `json:"unitAmount,omitempty"`
}

type ResourceUsage struct {
	ResourceID string                 `json:"resourceId"`
	FreeAmount float64                `json:"freeAmount"`
	PaidAmount float64                `json:"paidAmount"`
	UnitCost   ServiceEfficiencyMoney `json:"unitCost"`
	Price      ServiceEfficiencyPrice `json:"price"`
	TotalCost  ServiceEfficiencyMoney `json:"totalCost"`
}

type ServiceUsage struct {
	ServiceID      string                 `json:"serviceId"`
	SubtotalAmount ServiceEfficiencyMoney `json:"subtotalAmount"`
	ResourceUsages []ResourceUsage        `json:"resourceUsages"`
}

type ConsumerUsage struct {
	ConsumerID     int64                  `json:"consumerId"`
	ConsumerName   string                 `json:"consumerName"`
	SubtotalAmount ServiceEfficiencyMoney `json:"subtotalAmount"`
	ServiceUsages  []ServiceUsage         `json:"serviceUsages"`
}

type ServiceEfficiencyAddress struct {
	Country      string `json:"country"`
	PostalCode   string `json:"postalCode"`
	State        string `json:"state"`
	City         string `json:"city"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	Name         string `json:"name"`
}

type TransactionStatus string

const (
	TransactionStatusInvalid    TransactionStatus = "Invalid"
	TransactionStatusPending    TransactionStatus = "Pending"
	TransactionStatusSuccessful TransactionStatus = "Successful"
	TransactionStatusFailed     TransactionStatus = "Failed"
)

type ServiceEfficiencyPayment struct {
	Timestamp int64                  `json:"timestamp"`
	Amount    ServiceEfficiencyMoney `json:"amount"`
	Status    TransactionStatus      `json:"status"`
}

type BillDetailsResponse struct {
	AccountState   ResponseAccountState       `json:"accountState"`
	BillID         string                     `json:"billId"`
	BillStatus     BillStatus                 `json:"billStatus"`
	UsageAmount    ServiceEfficiencyMoney     `json:"usageAmount"`
	TaxAmount      ServiceEfficiencyMoney     `json:"taxAmount"`
	TotalAmount    ServiceEfficiencyMoney     `json:"totalAmount"`
	Usages         []ConsumerUsage            `json:"usages"`
	BillingAddress ServiceEfficiencyAddress   `json:"billingAddress"`
	RobloxAddress  ServiceEfficiencyAddress   `json:"robloxAddress"`
	TaxID          string                     `json:"taxId"`
	TaxIDType      TaxIDType                  `json:"taxIdType"`
	Payments       []ServiceEfficiencyPayment `json:"payments"`
	ConversionRate float64                    `json:"conversionRate"`
}

// GetBillByGroup will get Billing Statement for group owned account.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_GetBillByGroup
//
// [GET] /service-efficiency-api/v1/bills/groups/{groupId}
func (s *ServiceEfficiencyService) GetBillByGroup(ctx context.Context, groupId string, opts *GetBillByGroupOptions, reqOpts ...goblox.RequestOpt) (*BillDetailsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	billDetailsResponse := new(BillDetailsResponse)
	resp, err := s.client.Call(ctx, GetBillByGroup.Compile(q, groupId), nil, billDetailsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return billDetailsResponse, resp, nil
}

type GetBillByUserOptions struct {
	BillID *string `url:"billId,omitempty"`
	Year   *int    `url:"year,omitempty"`
	Month  *int    `url:"month,omitempty"`
	Date   *int    `url:"date,omitempty"`
	Latest *bool   `url:"latest,omitempty"`
}

// GetBillByUser will get Billing Statement for user owned account.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_GetBillByUser
//
// [GET] /service-efficiency-api/v1/bills/users/{userId}
func (s *ServiceEfficiencyService) GetBillByUser(ctx context.Context, userId string, opts *GetBillByUserOptions, reqOpts ...goblox.RequestOpt) (*BillDetailsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	billDetailsResponse := new(BillDetailsResponse)
	resp, err := s.client.Call(ctx, GetBillByUser.Compile(q, userId), nil, billDetailsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return billDetailsResponse, resp, nil
}

type ResourceConfiguration struct {
	ResourceID    *string                 `json:"resourceId,omitempty"`
	Unlocked      *bool                   `json:"unlocked,omitempty"`
	MonthlyBudget *ServiceEfficiencyMoney `json:"monthlyBudget,omitempty"`
	UnitCost      *ServiceEfficiencyMoney `json:"unitCost,omitempty"`
	Price         *ServiceEfficiencyPrice `json:"price,omitempty"`
}

type BudgetLevel string

const (
	BudgetLevelResourceOnly BudgetLevel = "ResourceOnly"
	BudgetLevelServiceOnly  BudgetLevel = "ServiceOnly"
)

type ServiceConfiguration struct {
	ServiceID              *string                 `json:"serviceId,omitempty"`
	MonthlyBudget          *ServiceEfficiencyMoney `json:"monthlyBudget,omitempty"`
	ResourceConfigurations []ResourceConfiguration `json:"resourceConfigurations,omitempty"`
	BudgetLevel            *BudgetLevel            `json:"budgetLevel,omitempty"`
}

type ConsumerConfigurationResponse struct {
	AccountState          ResponseAccountState   `json:"accountState"`
	AccountInfoComplete   bool                   `json:"accountInfoComplete"`
	ServiceConfigurations []ServiceConfiguration `json:"serviceConfigurations"`
}

// GetConsumerConfigurationsByUniverse will get ConsumerConfiguration for a universe.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/service-efficiency#ServiceEfficiencyApi_GetConsumerConfigurationsByUniverse
//
// [GET] /service-efficiency-api/v1/configurations/universes/{universeId}
func (s *ServiceEfficiencyService) GetConsumerConfigurationsByUniverse(ctx context.Context, universeId string, reqOpts ...goblox.RequestOpt) (*ConsumerConfigurationResponse, *Response, error) {
	consumerConfigurationResponse := new(ConsumerConfigurationResponse)
	resp, err := s.client.Call(ctx, GetConsumerConfigurationsByUniverse.Compile(nil, universeId), nil, consumerConfigurationResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return consumerConfigurationResponse, resp, nil
}

type ConsumerConfigurationRequest struct {
	ServiceConfigurations []ServiceConfiguration `json:"serviceConfigurations,omitempty"`
}

// UpdateConsumerConfigurationsByUniverse will update ConsumerConfiguration for a universe.
//
// Required scopes: extended-services:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/service-efficiency#ServiceEfficiencyApi_UpdateConsumerConfigurationsByUniverse
//
// [PATCH] /service-efficiency-api/v1/configurations/universes/{universeId}
func (s *ServiceEfficiencyService) UpdateConsumerConfigurationsByUniverse(ctx context.Context, universeId string, data ConsumerConfigurationRequest, reqOpts ...goblox.RequestOpt) (*ConsumerConfigurationResponse, *Response, error) {
	consumerConfigurationResponse := new(ConsumerConfigurationResponse)
	resp, err := s.client.Call(ctx, UpdateConsumerConfigurationsByUniverse.Compile(nil, universeId), data, consumerConfigurationResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return consumerConfigurationResponse, resp, nil
}

type EligibilityRequirement string

const (
	EligibilityRequirementInvalid          EligibilityRequirement = "Invalid"
	EligibilityRequirementAge              EligibilityRequirement = "Age"
	EligibilityRequirementAccountCountry   EligibilityRequirement = "AccountCountry"
	EligibilityRequirementIDVerified       EligibilityRequirement = "IdVerified"
	EligibilityRequirementEmailVerified    EligibilityRequirement = "EmailVerified"
	EligibilityRequirementModerationStatus EligibilityRequirement = "ModerationStatus"
	EligibilityRequirementPremium          EligibilityRequirement = "Premium"
	EligibilityRequirementUserAgreement    EligibilityRequirement = "UserAgreement"
)

type EligibilityRequirementResult struct {
	EligibilityRequirement EligibilityRequirement `json:"eligibilityRequirement"`
	Eligible               bool                   `json:"eligible"`
}

type CheckEligibilityResponse struct {
	EligibilityRequirementResults []EligibilityRequirementResult `json:"eligibilityRequirementResults"`
}

// CheckEligibilityForGroup will check the Eligibility of a Group.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_CheckEligibilityForGroup
//
// [GET] /service-efficiency-api/v1/eligibility/groups/{groupId}
func (s *ServiceEfficiencyService) CheckEligibilityForGroup(ctx context.Context, groupId string, reqOpts ...goblox.RequestOpt) (*CheckEligibilityResponse, *Response, error) {
	checkEligibilityResponse := new(CheckEligibilityResponse)
	resp, err := s.client.Call(ctx, CheckEligibilityForGroup.Compile(nil, groupId), nil, checkEligibilityResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return checkEligibilityResponse, resp, nil
}

// CheckEligibilityForUser will check the Eligibility of a User.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_CheckEligibilityForUser
//
// [GET] /service-efficiency-api/v1/eligibility/users/{userId}
func (s *ServiceEfficiencyService) CheckEligibilityForUser(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (*CheckEligibilityResponse, *Response, error) {
	checkEligibilityResponse := new(CheckEligibilityResponse)
	resp, err := s.client.Call(ctx, CheckEligibilityForUser.Compile(nil, userId), nil, checkEligibilityResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return checkEligibilityResponse, resp, nil
}

type IsAllowedOptions struct {
	UniverseID *int64 `url:"universeId,omitempty"`
}

// IsAllowed will return whether the current user can access the features in this API, optionally with a
// universe.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_IsAllowed
//
// [GET] /service-efficiency-api/v1/is_allowed
func (s *ServiceEfficiencyService) IsAllowed(ctx context.Context, opts *IsAllowedOptions, reqOpts ...goblox.RequestOpt) (bool, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return false, nil, err
	}

	var result bool
	resp, err := s.client.Call(ctx, IsAllowed.Compile(q), nil, &result, reqOpts...)
	if err != nil {
		return false, resp, err
	}

	return result, resp, nil
}

type PaymentProfile struct {
	PaymentProfileID        string `json:"paymentProfileId"`
	PaymentProfileOwnerType string `json:"paymentProfileOwnerType"`
	CardNetwork             string `json:"cardNetwork"`
	Last4Digits             string `json:"last4Digits"`
	ExpMonth                int    `json:"expMonth"`
	ExpYear                 int    `json:"expYear"`
}

type GetServiceEfficiencyPaymentMethodResponse struct {
	PaymentProfile PaymentProfile `json:"paymentProfile"`
}

// GetServiceEfficiencyPaymentMethodForGroup will get the Service Efficiency Payment Method for a Group.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_GetServiceEfficiencyPaymentMethodForGroup
//
// [GET] /service-efficiency-api/v1/payment_profiles/groups/{groupId}/service_efficiency
func (s *ServiceEfficiencyService) GetServiceEfficiencyPaymentMethodForGroup(ctx context.Context, groupId string, reqOpts ...goblox.RequestOpt) (*GetServiceEfficiencyPaymentMethodResponse, *Response, error) {
	getServiceEfficiencyPaymentMethodResponse := new(GetServiceEfficiencyPaymentMethodResponse)
	resp, err := s.client.Call(ctx, GetServiceEfficiencyPaymentMethodForGroup.Compile(nil, groupId), nil, getServiceEfficiencyPaymentMethodResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getServiceEfficiencyPaymentMethodResponse, resp, nil
}

type SetServiceEfficiencyPaymentMethodForGroupOptions struct {
	PaymentProfileID        string `url:"paymentProfileId"`        // Required.
	PaymentProfileOwnerType string `url:"paymentProfileOwnerType"` // Required.
}

type SetServiceEfficiencyPaymentMethodResponse map[string]any

// SetServiceEfficiencyPaymentMethodForGroup will set the Service Efficiency Payment Method for a Group.
//
// Required scopes: extended-services:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_SetServiceEfficiencyPaymentMethodForGroup
//
// [PUT] /service-efficiency-api/v1/payment_profiles/groups/{groupId}/service_efficiency
func (s *ServiceEfficiencyService) SetServiceEfficiencyPaymentMethodForGroup(ctx context.Context, groupId string, opts *SetServiceEfficiencyPaymentMethodForGroupOptions, reqOpts ...goblox.RequestOpt) (*SetServiceEfficiencyPaymentMethodResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	setServiceEfficiencyPaymentMethodResponse := new(SetServiceEfficiencyPaymentMethodResponse)
	resp, err := s.client.Call(ctx, SetServiceEfficiencyPaymentMethodForGroup.Compile(q, groupId), nil, setServiceEfficiencyPaymentMethodResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return setServiceEfficiencyPaymentMethodResponse, resp, nil
}

type DeletePaymentProfileForUserOptions struct {
	PaymentProfileID string `url:"paymentProfileId"` // Required.
}

type DeletePaymentProfileResponse map[string]any

// DeletePaymentProfileForUser will delete a Payment Profile for a User.
//
// Required scopes: extended-services:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_DeletePaymentProfileForUser
//
// [DELETE] /service-efficiency-api/v1/payment_profiles/users/{userId}
func (s *ServiceEfficiencyService) DeletePaymentProfileForUser(ctx context.Context, userId string, opts *DeletePaymentProfileForUserOptions, reqOpts ...goblox.RequestOpt) (*DeletePaymentProfileResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	deletePaymentProfileResponse := new(DeletePaymentProfileResponse)
	resp, err := s.client.Call(ctx, DeletePaymentProfileForUser.Compile(q, userId), nil, deletePaymentProfileResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return deletePaymentProfileResponse, resp, nil
}

type ListPaymentProfilesResponse struct {
	PaymentProfiles []PaymentProfile `json:"paymentProfiles"`
}

// ListPaymentProfilesByUser will list Payment Profiles of a User.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_ListPaymentProfilesByUser
//
// [GET] /service-efficiency-api/v1/payment_profiles/users/{userId}
func (s *ServiceEfficiencyService) ListPaymentProfilesByUser(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (*ListPaymentProfilesResponse, *Response, error) {
	listPaymentProfilesResponse := new(ListPaymentProfilesResponse)
	resp, err := s.client.Call(ctx, ListPaymentProfilesByUser.Compile(nil, userId), nil, listPaymentProfilesResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listPaymentProfilesResponse, resp, nil
}

type CreatePaymentProfileResponse struct {
	ClientSecret string `json:"clientSecret"`
}

// CreatePaymentProfileForUser will create Payment Profile for a User.
//
// Required scopes: extended-services:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_CreatePaymentProfileForUser
//
// [POST] /service-efficiency-api/v1/payment_profiles/users/{userId}
func (s *ServiceEfficiencyService) CreatePaymentProfileForUser(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (*CreatePaymentProfileResponse, *Response, error) {
	createPaymentProfileResponse := new(CreatePaymentProfileResponse)
	resp, err := s.client.Call(ctx, CreatePaymentProfileForUser.Compile(nil, userId), nil, createPaymentProfileResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return createPaymentProfileResponse, resp, nil
}

// GetServiceEfficiencyPaymentMethodForUser will get the Service Efficiency Payment Method for a User.
//
// Required scopes: extended-services:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_GetServiceEfficiencyPaymentMethodForUser
//
// [GET] /service-efficiency-api/v1/payment_profiles/users/{userId}/service_efficiency
func (s *ServiceEfficiencyService) GetServiceEfficiencyPaymentMethodForUser(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (*GetServiceEfficiencyPaymentMethodResponse, *Response, error) {
	getServiceEfficiencyPaymentMethodResponse := new(GetServiceEfficiencyPaymentMethodResponse)
	resp, err := s.client.Call(ctx, GetServiceEfficiencyPaymentMethodForUser.Compile(nil, userId), nil, getServiceEfficiencyPaymentMethodResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getServiceEfficiencyPaymentMethodResponse, resp, nil
}

type SetServiceEfficiencyPaymentMethodForUserOptions struct {
	PaymentProfileID        string `url:"paymentProfileId"`        // Required.
	PaymentProfileOwnerType string `url:"paymentProfileOwnerType"` // Required.
}

// SetServiceEfficiencyPaymentMethodForUser will set the Service Efficiency Payment Method for a User.
//
// Required scopes: extended-services:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_SetServiceEfficiencyPaymentMethodForUser
//
// [PUT] /service-efficiency-api/v1/payment_profiles/users/{userId}/service_efficiency
func (s *ServiceEfficiencyService) SetServiceEfficiencyPaymentMethodForUser(ctx context.Context, userId string, opts *SetServiceEfficiencyPaymentMethodForUserOptions, reqOpts ...goblox.RequestOpt) (*SetServiceEfficiencyPaymentMethodResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	setServiceEfficiencyPaymentMethodResponse := new(SetServiceEfficiencyPaymentMethodResponse)
	resp, err := s.client.Call(ctx, SetServiceEfficiencyPaymentMethodForUser.Compile(q, userId), nil, setServiceEfficiencyPaymentMethodResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return setServiceEfficiencyPaymentMethodResponse, resp, nil
}

type SubmitManualPaymentRequest struct {
	Amount *ServiceEfficiencyMoney `json:"amount,omitempty"`
	BillID *string                 `json:"billId,omitempty"`
}

type SubmitManualPaymentResponse map[string]any

// SubmitManualPaymentForGroup will submit a Manual Payment for a Group.
//
// Required scopes: extended-services:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_SubmitManualPaymentForGroup
//
// [POST] /service-efficiency-api/v1/payments/groups/{groupId}
func (s *ServiceEfficiencyService) SubmitManualPaymentForGroup(ctx context.Context, groupId string, data SubmitManualPaymentRequest, reqOpts ...goblox.RequestOpt) (*SubmitManualPaymentResponse, *Response, error) {
	submitManualPaymentResponse := new(SubmitManualPaymentResponse)
	resp, err := s.client.Call(ctx, SubmitManualPaymentForGroup.Compile(nil, groupId), data, submitManualPaymentResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return submitManualPaymentResponse, resp, nil
}

// SubmitManualPaymentForUser will submit a Manual Payment for a User.
//
// Required scopes: extended-services:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/accounts#ServiceEfficiencyApi_SubmitManualPaymentForUser
//
// [POST] /service-efficiency-api/v1/payments/users/{userId}
func (s *ServiceEfficiencyService) SubmitManualPaymentForUser(ctx context.Context, userId string, data SubmitManualPaymentRequest, reqOpts ...goblox.RequestOpt) (*SubmitManualPaymentResponse, *Response, error) {
	submitManualPaymentResponse := new(SubmitManualPaymentResponse)
	resp, err := s.client.Call(ctx, SubmitManualPaymentForUser.Compile(nil, userId), data, submitManualPaymentResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return submitManualPaymentResponse, resp, nil
}
