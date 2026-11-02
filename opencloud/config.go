package opencloud

import (
	"context"
	"net/http"

	"github.com/typical-developers/goblox"
	"github.com/typical-developers/goblox/auth"
)

var (
	GetConfigWithoutMetadata = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:read")

	GetConfigDraft = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:read")

	PartialUpdateConfigDraft = goblox.NewEndpoint(
		http.MethodPatch,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:write")

	DeleteConfigDraft = goblox.NewEndpoint(
		http.MethodDelete,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:write")

	UpdateConfigDraft = goblox.NewEndpoint(
		http.MethodPut,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft:overwrite",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:write")

	GetConfig = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/full",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:read")

	PublishConfigDraft = goblox.NewEndpoint(
		http.MethodPost,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/publish",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:write")

	ListRevisionHistory = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/revisions",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:read")

	RevisionRestore = goblox.NewEndpoint(
		http.MethodPost,
		"/creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/revisions/{revisionId}/restore",
		auth.KindAPIKey, auth.KindOAuth2, auth.KindCookie,
	).WithScopes("universe:write")
)

// ConfigService will handle communication with the actions related to the API.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type ConfigService service

type ConfigMetadata struct {
	ConfigVersion int `json:"configVersion"`
}

type Config struct {
	Metadata ConfigMetadata `json:"metadata"`
	Entries  map[string]any `json:"entries"`
}

type ConfigEntry struct {
	Value            any    `json:"value"`
	Description      string `json:"description"`
	LastModifiedTime string `json:"lastModifiedTime"`
	LastAccessedTime string `json:"lastAccessedTime"`
}

type FullConfig struct {
	Metadata ConfigMetadata         `json:"metadata"`
	Entries  map[string]ConfigEntry `json:"entries"`
}

// GetConfigWithoutMetadata will get the currently published config without metadata and decorators.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_GetConfigRepositoryValues
//
// Required scopes: universe:read
//
// [GET] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}
func (s *ConfigService) GetConfigWithoutMetadata(ctx context.Context, universeId, repository string, reqOpts ...goblox.RequestOpt) (*Config, *Response, error) {
	config := new(Config)

	resp, err := s.client.Call(ctx, GetConfigWithoutMetadata.Compile(nil, universeId, repository), nil, config, reqOpts...)
	if err != nil {
		return nil, resp, err
	}
	return config, resp, nil
}

type ConfigDraftEntry struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}

type ConfigDraft struct {
	DraftHash string                      `json:"draftHash"`
	Entries   map[string]ConfigDraftEntry `json:"entries"`
}

// GetConfigDraft will return the current draft changes for the config.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_GetConfigRepositoryDraft
//
// Required scopes: universe:read
//
// [GET] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft
func (s *ConfigService) GetConfigDraft(ctx context.Context, universeId, repository string, reqOpts ...goblox.RequestOpt) (*ConfigDraft, *Response, error) {
	draft := new(ConfigDraft)
	resp, err := s.client.Call(ctx, GetConfigDraft.Compile(nil, universeId, repository), nil, draft, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return draft, resp, nil
}

type ConfigDraftUpdate struct {
	DraftHash string         `json:"draftHash"`
	Entries   map[string]any `json:"entries"`
}

type ConfigDraftUpdated struct {
	DraftHash string `json:"draftHash"`
}

// PartialUpdateConfigDraft will partially update the draft changes for the config.
// A key that is not included will not be changed.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_UpdateDraft
//
// Required scopes: universe:write
//
// [PATCH] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft
func (s *ConfigService) PartialUpdateConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftUpdate, reqOpts ...goblox.RequestOpt) (*ConfigDraftUpdated, *Response, error) {
	details := new(ConfigDraftUpdated)
	resp, err := s.client.Call(ctx, PartialUpdateConfigDraft.Compile(nil, universeId, repository), data, details, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}

type ConfigDraftDelete struct {
	DraftHash string `json:"draftHash"`
}

type ConfigDraftDeleted struct {
	DraftHash string `json:"draftHash"`
}

// DeleteConfigDraft will reset the current draft.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_DeleteDraft
//
// Required scopes: universe:write
//
// [DELETE] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft
func (s *ConfigService) DeleteConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftDelete, reqOpts ...goblox.RequestOpt) (*ConfigDraftDeleted, *Response, error) {
	details := new(ConfigDraftDeleted)
	resp, err := s.client.Call(ctx, DeleteConfigDraft.Compile(nil, universeId, repository), data, details, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}

// UpdateConfigDraft will overwrite the current draft.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_OverwriteDraft
//
// Required scopes: universe:write
//
// [PUT] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/draft:overwrite
func (s *ConfigService) UpdateConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftUpdate, reqOpts ...goblox.RequestOpt) (*ConfigDraftUpdated, *Response, error) {
	details := new(ConfigDraftUpdated)
	resp, err := s.client.Call(ctx, UpdateConfigDraft.Compile(nil, universeId, repository), data, details, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}

// GetConfig will get the currently published config.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_GetConfigRepositoryFull
//
// Required scopes: universe:read
//
// [GET] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/full
func (s *ConfigService) GetConfig(ctx context.Context, universeId, repository string, reqOpts ...goblox.RequestOpt) (*FullConfig, *Response, error) {
	config := new(FullConfig)
	resp, err := s.client.Call(ctx, GetConfig.Compile(nil, universeId, repository), nil, config, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}

type ConfigDraftPublish struct {
	DraftHash          string `json:"draftHash"`
	Message            string `json:"message,omitempty"`
	DeploymentStrategy string `json:"deploymentStrategy,omitempty"`
}

type ConfigDraftPublished struct {
	ConfigVersion int `json:"configVersion"`
}

// PublishConfigDraft will publish the current draft changes to the published config.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_PublishDraft
//
// Required scopes: universe:write
//
// [POST] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/publish
func (s *ConfigService) PublishConfigDraft(ctx context.Context, universeId, repository string, data ConfigDraftPublish, reqOpts ...goblox.RequestOpt) (*ConfigDraftPublished, *Response, error) {
	details := new(ConfigDraftPublished)
	resp, err := s.client.Call(ctx, PublishConfigDraft.Compile(nil, universeId, repository), data, details, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}

type ConfigRevisionChange struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type ConfigRevision struct {
	RevisionID           string                          `json:"revisionId"`
	Version              int                             `json:"version"`
	Time                 string                          `json:"time"`
	PublishedBy          string                          `json:"publishedBy"`
	Message              string                          `json:"message"`
	DeploymentResult     string                          `json:"deploymentResult"`
	AutoRestoreToVersion int                             `json:"autoRestoreToVersion"`
	Changes              map[string]ConfigRevisionChange `json:"changes"`
}

type ConfigRevisionHistory struct {
	Revisions []ConfigRevision `json:"revisions"`
}

type ConfigRevisionHistoryOptions struct {
	StartTime   *string `url:"startTime,omitempty"`
	EndTime     *string `url:"endTime,omitempty"`
	MaxPageSize *int    `url:"MaxPageSize,omitempty"`
	Skip        *bool   `url:"Skip,omitempty"`
	SearchKey   *string `url:"SearchKey,omitempty"`
	SortOrder   *string `url:"SortOrder,omitempty"`
	SortKey     *string `url:"SortKey,omitempty"`
}

// ListRevisionHistory will list the previous revisions to the config.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_ListRevisions
//
// Required scopes: universe:read
//
// [GET] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/revisions
func (s *ConfigService) ListRevisionHistory(ctx context.Context, universeId, repository string, opts *ConfigRevisionHistoryOptions, reqOpts ...goblox.RequestOpt) (*ConfigRevisionHistory, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	revisionHistory := new(ConfigRevisionHistory)
	resp, err := s.client.Call(ctx, ListRevisionHistory.Compile(q, universeId, repository), nil, revisionHistory, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return revisionHistory, resp, nil
}

type ConfigRestoreRevision struct {
	DraftHash string `json:"draftHash"`
}

// RevisionRestore will stage a revert to a previous revision of the config.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/configs#CreatorConfigsPublicApi_RestoreRevision
//
// Required scopes: universe:write
//
// [POST] /creator-configs-public-api/v1/configs/universes/{universeId}/repositories/{repository}/revisions/{revisionId}/restore
func (s *ConfigService) RevisionRestore(ctx context.Context, universeId, repository, revisionId string, reqOpts ...goblox.RequestOpt) (*ConfigRestoreRevision, *Response, error) {
	details := new(ConfigRestoreRevision)
	resp, err := s.client.Call(ctx, RevisionRestore.Compile(nil, universeId, repository, revisionId), nil, details, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return details, resp, nil
}

var (
	ListExperiments = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:read")

	CreateExperiment = goblox.NewEndpoint(
		http.MethodPost,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:write")

	DiscardExperiment = goblox.NewEndpoint(
		http.MethodDelete,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:write")

	GetExperiment = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:read")

	UpdateExperiment = goblox.NewEndpoint(
		http.MethodPatch,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:write")

	GetExperimentStats = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}/stats",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:read")

	CompleteExperiment = goblox.NewEndpoint(
		http.MethodPost,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}:complete",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:write")

	ScheduleExperiment = goblox.NewEndpoint(
		http.MethodPost,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}:schedule",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:write")

	StartExperiment = goblox.NewEndpoint(
		http.MethodPost,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}:start",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:write")

	CalculateExperimentMde = goblox.NewEndpoint(
		http.MethodPost,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments:calculateMde",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:read")

	GetExperimentOperationStatus = goblox.NewEndpoint(
		http.MethodGet,
		"/creator-configs-public-api/v1/experimentation/universes/{universeId}/operations/{operationId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe:read")
)

type ExperimentState string

const (
	ExperimentStateDraft     ExperimentState = "EXPERIMENT_STATE_DRAFT"
	ExperimentStateScheduled ExperimentState = "EXPERIMENT_STATE_SCHEDULED"
	ExperimentStateRunning   ExperimentState = "EXPERIMENT_STATE_RUNNING"
	ExperimentStateCompleted ExperimentState = "EXPERIMENT_STATE_COMPLETED"
	ExperimentStateCancelled ExperimentState = "EXPERIMENT_STATE_CANCELLED"
	ExperimentStateDeleted   ExperimentState = "EXPERIMENT_STATE_DELETED"
)

type ExperimentSortOrder string

const (
	ExperimentSortOrderAscending  ExperimentSortOrder = "SORT_ORDER_ASCENDING"
	ExperimentSortOrderDescending ExperimentSortOrder = "SORT_ORDER_DESCENDING"
)

type ListExperimentSortKey string

const (
	ListExperimentSortKeyName             ListExperimentSortKey = "LIST_EXPERIMENT_SORT_KEY_NAME"
	ListExperimentSortKeyLastModifiedTime ListExperimentSortKey = "LIST_EXPERIMENT_SORT_KEY_LAST_MODIFIED_TIME"
	ListExperimentSortKeyStartTime        ListExperimentSortKey = "LIST_EXPERIMENT_SORT_KEY_START_TIME"
	ListExperimentSortKeyRelevance        ListExperimentSortKey = "LIST_EXPERIMENT_SORT_KEY_RELEVANCE"
)

type ExperimentProductType string

const (
	ExperimentProductTypeInGameConfigs ExperimentProductType = "EXPERIMENT_PRODUCT_TYPE_IN_GAME_CONFIGS"
	ExperimentProductTypeMatchmaking   ExperimentProductType = "EXPERIMENT_PRODUCT_TYPE_MATCHMAKING"
)

type ListExperimentsOptions struct {
	MaxPageSize       *int                   `url:"maxPageSize,omitempty"`
	Skip              *int                   `url:"skip,omitempty"`
	SearchKey         *string                `url:"searchKey,omitempty"`
	StateFilters      []ExperimentState      `url:"stateFilters,omitempty"`
	SortOrder         *ExperimentSortOrder   `url:"sortOrder,omitempty"`
	SortKey           *ListExperimentSortKey `url:"sortKey,omitempty"`
	ProductTypeFilter *ExperimentProductType `url:"productTypeFilter,omitempty"`
}

type ExperimentSummary struct {
	ID                  string                `json:"id"`
	Name                string                `json:"name"`
	Description         string                `json:"description"`
	CreatedTime         string                `json:"createdTime"`
	LastUpdatedTime     string                `json:"lastUpdatedTime"`
	StartedTime         string                `json:"startedTime"`
	StoppedTime         string                `json:"stoppedTime"`
	State               ExperimentState       `json:"state"`
	CreatedBy           string                `json:"createdBy"`
	ExperimentConfigKey string                `json:"experimentConfigKey"`
	ScheduledTime       string                `json:"scheduledTime"`
	DurationSeconds     int64                 `json:"durationSeconds"`
	ProductType         ExperimentProductType `json:"productType"`
}

type ListExperimentsResponse struct {
	Experiments []ExperimentSummary `json:"experiments"`
	Total       int                 `json:"total"`
}

// ListExperiments will list experiments for a universe.
//
// Required scopes: universe:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_ListExperiments
//
// [GET] /creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments
func (s *ConfigService) ListExperiments(ctx context.Context, universeId string, opts *ListExperimentsOptions, reqOpts ...goblox.RequestOpt) (*ListExperimentsResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	listExperimentsResponse := new(ListExperimentsResponse)
	resp, err := s.client.Call(ctx, ListExperiments.Compile(q, universeId), nil, listExperimentsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return listExperimentsResponse, resp, nil
}

type VariantMeta struct {
	VariantID  *string `json:"variantId,omitempty"`
	Label      *string `json:"label,omitempty"`
	IsBaseline *bool   `json:"isBaseline,omitempty"`
	Weight     *int    `json:"weight,omitempty"`
}

type ExperimentConfigEntry struct {
	Key         *string `json:"key,omitempty"`
	EntryValue  any     `json:"entryValue,omitempty"`
	Description *string `json:"description,omitempty"`
}

type SingleConfigExperimentVariant struct {
	VariantMeta *VariantMeta           `json:"variantMeta,omitempty"`
	ConfigEntry *ExperimentConfigEntry `json:"configEntry,omitempty"`
}

type InGameConfigExperimentConfiguration struct {
	Variants []SingleConfigExperimentVariant `json:"variants,omitempty"`
}

type PlaceScoringConfig struct {
	PlaceID                    *int64  `json:"placeId,omitempty"`
	MatchmakingScoringConfigID *string `json:"matchmakingScoringConfigId,omitempty"`
	UsePlatformDefault         *bool   `json:"usePlatformDefault,omitempty"`
}

type MatchmakingExperimentVariant struct {
	VariantMeta             *VariantMeta         `json:"variantMeta,omitempty"`
	PlaceMatchmakingConfigs []PlaceScoringConfig `json:"placeMatchmakingConfigs,omitempty"`
}

type MatchmakingExperimentConfiguration struct {
	Variants []MatchmakingExperimentVariant `json:"variants,omitempty"`
}

type ExperimentConfiguration struct {
	ProductType                         *ExperimentProductType               `json:"productType,omitempty"`
	InGameConfigExperimentConfiguration *InGameConfigExperimentConfiguration `json:"inGameConfigExperimentConfiguration,omitempty"`
	MatchmakingExperimentConfiguration  *MatchmakingExperimentConfiguration  `json:"matchmakingExperimentConfiguration,omitempty"`
}

type ConditionalRuleOperator string

const (
	ConditionalRuleOperatorAnd ConditionalRuleOperator = "OPERATOR_AND"
	ConditionalRuleOperatorOr  ConditionalRuleOperator = "OPERATOR_OR"
	ConditionalRuleOperatorNot ConditionalRuleOperator = "OPERATOR_NOT"
	ConditionalRuleOperatorEq  ConditionalRuleOperator = "OPERATOR_EQ"
	ConditionalRuleOperatorNe  ConditionalRuleOperator = "OPERATOR_NE"
	ConditionalRuleOperatorGt  ConditionalRuleOperator = "OPERATOR_GT"
	ConditionalRuleOperatorLt  ConditionalRuleOperator = "OPERATOR_LT"
	ConditionalRuleOperatorGte ConditionalRuleOperator = "OPERATOR_GTE"
	ConditionalRuleOperatorLte ConditionalRuleOperator = "OPERATOR_LTE"
	ConditionalRuleOperatorIn  ConditionalRuleOperator = "OPERATOR_IN"
	ConditionalRuleOperatorNin ConditionalRuleOperator = "OPERATOR_NIN"
)

type LiteralValueDto struct {
	BooleanValue *bool    `json:"booleanValue,omitempty"`
	IntegerValue *int64   `json:"integerValue,omitempty"`
	DoubleValue  *float64 `json:"doubleValue,omitempty"`
	StringValue  *string  `json:"stringValue,omitempty"`
}

type RpnOperandDto struct {
	AttributeReference *string          `json:"attributeReference,omitempty"`
	LiteralValue       *LiteralValueDto `json:"literalValue,omitempty"`
}

type RpnTokenDto struct {
	Operator *ConditionalRuleOperator `json:"operator,omitempty"`
	Operand  *RpnOperandDto           `json:"operand,omitempty"`
}

type ConditionalRuleDefinition struct {
	Tokens []RpnTokenDto `json:"tokens,omitempty"`
}

type TargetingCriteria struct {
	Rule *ConditionalRuleDefinition `json:"rule,omitempty"`
}

type UniverseExperimentMetric string

const (
	UniverseExperimentMetricAverageSessionTime          UniverseExperimentMetric = "UNIVERSE_EXPERIMENT_METRIC_AVERAGE_SESSION_TIME"
	UniverseExperimentMetricPlaytimePerUser             UniverseExperimentMetric = "UNIVERSE_EXPERIMENT_METRIC_PLAYTIME_PER_USER"
	UniverseExperimentMetricDay1Retention               UniverseExperimentMetric = "UNIVERSE_EXPERIMENT_METRIC_DAY_1_RETENTION"
	UniverseExperimentMetricDay7Retention               UniverseExperimentMetric = "UNIVERSE_EXPERIMENT_METRIC_DAY_7_RETENTION"
	UniverseExperimentMetricPayerConversionRate         UniverseExperimentMetric = "UNIVERSE_EXPERIMENT_METRIC_PAYER_CONVERSION_RATE"
	UniverseExperimentMetricAverageRevenuePerUser       UniverseExperimentMetric = "UNIVERSE_EXPERIMENT_METRIC_AVERAGE_REVENUE_PER_USER"
	UniverseExperimentMetricAverageRevenuePerPayingUser UniverseExperimentMetric = "UNIVERSE_EXPERIMENT_METRIC_AVERAGE_REVENUE_PER_PAYING_USER"
)

type CreateExperimentData struct {
	Name                    *string                   `json:"name,omitempty"`
	Description             *string                   `json:"description,omitempty"`
	ExperimentConfiguration *ExperimentConfiguration  `json:"experimentConfiguration,omitempty"`
	ExposurePercent         *int                      `json:"exposurePercent,omitempty"`
	TargetingCriteria       *TargetingCriteria        `json:"targetingCriteria,omitempty"`
	DurationSeconds         *int64                    `json:"durationSeconds,omitempty"`
	UniverseGoalMetric      *UniverseExperimentMetric `json:"universeGoalMetric,omitempty"`
}

type OperationalStatus string

const (
	OperationalStatusReady      OperationalStatus = "OPERATIONAL_STATUS_READY"
	OperationalStatusCreating   OperationalStatus = "OPERATIONAL_STATUS_CREATING"
	OperationalStatusUpdating   OperationalStatus = "OPERATIONAL_STATUS_UPDATING"
	OperationalStatusStarting   OperationalStatus = "OPERATIONAL_STATUS_STARTING"
	OperationalStatusStopping   OperationalStatus = "OPERATIONAL_STATUS_STOPPING"
	OperationalStatusScheduling OperationalStatus = "OPERATIONAL_STATUS_SCHEDULING"
	OperationalStatusDeleting   OperationalStatus = "OPERATIONAL_STATUS_DELETING"
	OperationalStatusRampingUp  OperationalStatus = "OPERATIONAL_STATUS_RAMPING_UP"
	OperationalStatusSyncing    OperationalStatus = "OPERATIONAL_STATUS_SYNCING"
	OperationalStatusRollingOut OperationalStatus = "OPERATIONAL_STATUS_ROLLING_OUT"
)

type ExperimentAPIErrorType string

const (
	ExperimentAPIErrorTypeInvalidVariantConfiguration                            ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_INVALID_VARIANT_CONFIGURATION"
	ExperimentAPIErrorTypeInvalidVariantLabel                                    ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_INVALID_VARIANT_LABEL"
	ExperimentAPIErrorTypeMustHaveExactlyOneBaselineVariant                      ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MUST_HAVE_EXACTLY_ONE_BASELINE_VARIANT"
	ExperimentAPIErrorTypeMatchmakingOverlappingRuntime                          ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_OVERLAPPING_RUNTIME"
	ExperimentAPIErrorTypeMatchmakingInvalidScoringConfigurations                ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_INVALID_SCORING_CONFIGURATIONS"
	ExperimentAPIErrorTypeMatchmakingUnexpectedProductType                       ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_UNEXPECTED_PRODUCT_TYPE"
	ExperimentAPIErrorTypeMatchmakingEmptyConfiguration                          ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_EMPTY_CONFIGURATION"
	ExperimentAPIErrorTypeMatchmakingEmptyVariantMetadata                        ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_EMPTY_VARIANT_METADATA"
	ExperimentAPIErrorTypeMatchmakingVariantWeightMustBePositive                 ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_VARIANT_WEIGHT_MUST_BE_POSITIVE"
	ExperimentAPIErrorTypeMatchmakingVariantWeightsUnbalanced                    ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_VARIANT_WEIGHTS_UNBALANCED"
	ExperimentAPIErrorTypeMatchmakingPlaceConfigRequiresScoringIDOrDefault       ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_PLACE_CONFIG_REQUIRES_SCORING_ID_OR_DEFAULT"
	ExperimentAPIErrorTypeConfigsMissingVariant                                  ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_CONFIGS_MISSING_VARIANT"
	ExperimentAPIErrorTypeConfigsVariantMissingKey                               ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_CONFIGS_VARIANT_MISSING_KEY"
	ExperimentAPIErrorTypeConfigsVariantMissingValue                             ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_CONFIGS_VARIANT_MISSING_VALUE"
	ExperimentAPIErrorTypeConfigsKeyAlreadyInUse                                 ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_CONFIGS_KEY_ALREADY_IN_USE"
	ExperimentAPIErrorTypeSystemError                                            ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_SYSTEM_ERROR"
	ExperimentAPIErrorTypeExperimentResultsNotFound                              ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_EXPERIMENT_RESULTS_NOT_FOUND"
	ExperimentAPIErrorTypeInvalidTargetingCriteria                               ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_INVALID_TARGETING_CRITERIA"
	ExperimentAPIErrorTypeTargetingNotAllowed                                    ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_TARGETING_NOT_ALLOWED"
	ExperimentAPIErrorTypeInvalidDuration                                        ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_INVALID_DURATION"
	ExperimentAPIErrorTypeConfigsVariantConfigNotFound                           ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_CONFIGS_VARIANT_CONFIG_NOT_FOUND"
	ExperimentAPIErrorTypeInvalidExperimentName                                  ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_INVALID_EXPERIMENT_NAME"
	ExperimentAPIErrorTypeConditionalConfigRequiresNullBaseline                  ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_CONDITIONAL_CONFIG_REQUIRES_NULL_BASELINE"
	ExperimentAPIErrorTypeMatchmakingControlVariantMustMatchCurrentConfiguration ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_CONTROL_VARIANT_MUST_MATCH_CURRENT_CONFIGURATION"
	ExperimentAPIErrorTypeMatchmakingVariantsMustCoverSamePlaces                 ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_VARIANTS_MUST_COVER_SAME_PLACES"
	ExperimentAPIErrorTypeMatchmakingVariantHasDuplicatePlaces                   ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_MATCHMAKING_VARIANT_HAS_DUPLICATE_PLACES"
	ExperimentAPIErrorTypeConfigsKeyHasReservedPrefix                            ExperimentAPIErrorType = "EXPERIMENT_API_ERROR_TYPE_CONFIGS_KEY_HAS_RESERVED_PREFIX"
)

type ExperimentAPIError struct {
	ErrorType    ExperimentAPIErrorType `json:"errorType"`
	ErrorCode    string                 `json:"errorCode"`
	ErrorMessage string                 `json:"errorMessage"`
	ErrorContext map[string]string      `json:"errorContext"`
}

type UniverseMetricConfiguration struct {
	GoalMetric      UniverseExperimentMetric   `json:"goalMetric"`
	LearningMetrics []UniverseExperimentMetric `json:"learningMetrics"`
}

type ProductExperiment struct {
	ID                          string                      `json:"id"`
	Name                        string                      `json:"name"`
	Description                 string                      `json:"description"`
	CreatedTime                 string                      `json:"createdTime"`
	LastUpdatedTime             string                      `json:"lastUpdatedTime"`
	StartedTime                 string                      `json:"startedTime"`
	StoppedTime                 string                      `json:"stoppedTime"`
	State                       ExperimentState             `json:"state"`
	ExperimentConfiguration     ExperimentConfiguration     `json:"experimentConfiguration"`
	ExposurePercent             int                         `json:"exposurePercent"`
	TargetingCriteria           TargetingCriteria           `json:"targetingCriteria"`
	OperationalStatus           OperationalStatus           `json:"operationalStatus"`
	CreatedBy                   string                      `json:"createdBy"`
	ScheduledTime               string                      `json:"scheduledTime"`
	DurationSeconds             int64                       `json:"durationSeconds"`
	UniverseMetricConfiguration UniverseMetricConfiguration `json:"universeMetricConfiguration"`
}

type ExperimentOperation struct {
	OperationID string             `json:"operationId"`
	Status      OperationalStatus  `json:"status"`
	CreatedTime string             `json:"createdTime"`
	Done        bool               `json:"done"`
	Error       ExperimentAPIError `json:"error"`
	Experiment  ProductExperiment  `json:"experiment"`
}

type CreateExperimentResponse struct {
	Operation ExperimentOperation `json:"operation"`
}

// CreateExperiment will create a new experiment for a universe.
//
// Required scopes: universe:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_CreateExperiment
//
// [POST] /creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments
func (s *ConfigService) CreateExperiment(ctx context.Context, universeId string, data CreateExperimentData, reqOpts ...goblox.RequestOpt) (*CreateExperimentResponse, *Response, error) {
	createExperimentResponse := new(CreateExperimentResponse)
	resp, err := s.client.Call(ctx, CreateExperiment.Compile(nil, universeId), data, createExperimentResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return createExperimentResponse, resp, nil
}

type DiscardExperimentResponse struct {
	Operation ExperimentOperation `json:"operation"`
}

// DiscardExperiment will discard an experiment.
//
// Required scopes: universe:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_DiscardExperiment
//
// [DELETE] /creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}
func (s *ConfigService) DiscardExperiment(ctx context.Context, universeId, experimentId string, reqOpts ...goblox.RequestOpt) (*DiscardExperimentResponse, *Response, error) {
	discardExperimentResponse := new(DiscardExperimentResponse)
	resp, err := s.client.Call(ctx, DiscardExperiment.Compile(nil, universeId, experimentId), nil, discardExperimentResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return discardExperimentResponse, resp, nil
}

type GetExperimentResponse struct {
	Experiment ProductExperiment `json:"experiment"`
}

// GetExperiment will get an experiment by ID.
//
// Required scopes: universe:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_GetExperiment
//
// [GET] /creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}
func (s *ConfigService) GetExperiment(ctx context.Context, universeId, experimentId string, reqOpts ...goblox.RequestOpt) (*GetExperimentResponse, *Response, error) {
	getExperimentResponse := new(GetExperimentResponse)
	resp, err := s.client.Call(ctx, GetExperiment.Compile(nil, universeId, experimentId), nil, getExperimentResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getExperimentResponse, resp, nil
}

type UpdateExperimentData struct {
	Name                    *string                   `json:"name,omitempty"`
	Description             *string                   `json:"description,omitempty"`
	ExperimentConfiguration *ExperimentConfiguration  `json:"experimentConfiguration,omitempty"`
	ExposurePercent         *int                      `json:"exposurePercent,omitempty"`
	TargetingCriteria       *TargetingCriteria        `json:"targetingCriteria,omitempty"`
	DurationSeconds         *int64                    `json:"durationSeconds,omitempty"`
	UniverseGoalMetric      *UniverseExperimentMetric `json:"universeGoalMetric,omitempty"`
}

type UpdateExperimentResponse struct {
	Operation ExperimentOperation `json:"operation"`
}

// UpdateExperiment will update an existing experiment.
//
// Required scopes: universe:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_UpdateExperiment
//
// [PATCH] /creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}
func (s *ConfigService) UpdateExperiment(ctx context.Context, universeId, experimentId string, data UpdateExperimentData, reqOpts ...goblox.RequestOpt) (*UpdateExperimentResponse, *Response, error) {
	updateExperimentResponse := new(UpdateExperimentResponse)
	resp, err := s.client.Call(ctx, UpdateExperiment.Compile(nil, universeId, experimentId), data, updateExperimentResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return updateExperimentResponse, resp, nil
}

type ExperimentStats struct {
	IsSrmDetected bool `json:"isSrmDetected"`
}

type ExperimentStatsOperation struct {
	ExperimentID     int64              `json:"experimentId"`
	LastModifiedTime string             `json:"lastModifiedTime"`
	Done             bool               `json:"done"`
	Error            ExperimentAPIError `json:"error"`
	ExperimentStats  ExperimentStats    `json:"experimentStats"`
}

type GetExperimentStatsResponse struct {
	Operation ExperimentStatsOperation `json:"operation"`
}

// GetExperimentStats will get diagnostic stats for an experiment.
//
// Required scopes: universe:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_GetExperimentStats
//
// [GET] /creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}/stats
func (s *ConfigService) GetExperimentStats(ctx context.Context, universeId, experimentId string, reqOpts ...goblox.RequestOpt) (*GetExperimentStatsResponse, *Response, error) {
	getExperimentStatsResponse := new(GetExperimentStatsResponse)
	resp, err := s.client.Call(ctx, GetExperimentStats.Compile(nil, universeId, experimentId), nil, getExperimentStatsResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getExperimentStatsResponse, resp, nil
}

type CompleteExperimentRequest struct {
	VariantID *string `json:"variantId,omitempty"`
}

type CompleteExperimentResponse struct {
	Operation ExperimentOperation `json:"operation"`
}

// CompleteExperiment will complete an experiment by stopping it and rolling out the requested variant.
//
// Required scopes: universe:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_CompleteExperiment
//
// [POST] /creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}:complete
func (s *ConfigService) CompleteExperiment(ctx context.Context, universeId, experimentId string, data CompleteExperimentRequest, reqOpts ...goblox.RequestOpt) (*CompleteExperimentResponse, *Response, error) {
	completeExperimentResponse := new(CompleteExperimentResponse)
	resp, err := s.client.Call(ctx, CompleteExperiment.Compile(nil, universeId, experimentId), data, completeExperimentResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return completeExperimentResponse, resp, nil
}

type ScheduleExperimentRequest struct {
	ScheduledStartTime *string `json:"scheduledStartTime,omitempty"`
}

type ScheduleExperimentResponse struct {
	Operation ExperimentOperation `json:"operation"`
}

// ScheduleExperiment will schedule an experiment to start at a future UTC time.
//
// Required scopes: universe:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_ScheduleExperiment
//
// [POST] /creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}:schedule
func (s *ConfigService) ScheduleExperiment(ctx context.Context, universeId, experimentId string, data ScheduleExperimentRequest, reqOpts ...goblox.RequestOpt) (*ScheduleExperimentResponse, *Response, error) {
	scheduleExperimentResponse := new(ScheduleExperimentResponse)
	resp, err := s.client.Call(ctx, ScheduleExperiment.Compile(nil, universeId, experimentId), data, scheduleExperimentResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return scheduleExperimentResponse, resp, nil
}

type StartExperimentResponse struct {
	Operation ExperimentOperation `json:"operation"`
}

// StartExperiment will start an experiment immediately.
//
// Required scopes: universe:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_StartExperiment
//
// [POST] /creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments/{experimentId}:start
func (s *ConfigService) StartExperiment(ctx context.Context, universeId, experimentId string, reqOpts ...goblox.RequestOpt) (*StartExperimentResponse, *Response, error) {
	startExperimentResponse := new(StartExperimentResponse)
	resp, err := s.client.Call(ctx, StartExperiment.Compile(nil, universeId, experimentId), nil, startExperimentResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return startExperimentResponse, resp, nil
}

type CalculateExperimentMDEData struct {
	ExposurePercent    *int                      `json:"exposurePercent,omitempty"`
	DurationSeconds    *int64                    `json:"durationSeconds,omitempty"`
	BaselineWeight     *float64                  `json:"baselineWeight,omitempty"`
	VariantWeights     []float64                 `json:"variantWeights,omitempty"`
	UniverseGoalMetric *UniverseExperimentMetric `json:"universeGoalMetric,omitempty"`
	TargetingCriteria  *TargetingCriteria        `json:"targetingCriteria,omitempty"`
}

type MDEResult struct {
	TotalSampleSize            int       `json:"totalSampleSize"`
	MDERelativePercentages     []float64 `json:"mdeRelativePercentages"`
	MinimumSampleSizeThreshold int       `json:"minimumSampleSizeThreshold"`
}

type MDEOperation struct {
	OperationID      string             `json:"operationId"`
	LastModifiedTime string             `json:"lastModifiedTime"`
	Done             bool               `json:"done"`
	Error            ExperimentAPIError `json:"error"`
	MDE              MDEResult          `json:"mde"`
}

type CalculateExperimentMDEResponse struct {
	Operation MDEOperation `json:"operation"`
}

// CalculateExperimentMde will calculate the Minimum Detectable Effect (MDE) for a prospective experiment.
//
// Required scopes: universe:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_CalculateExperimentMde
//
// [POST] /creator-configs-public-api/v1/experimentation/universes/{universeId}/experiments:calculateMde
func (s *ConfigService) CalculateExperimentMde(ctx context.Context, universeId string, data CalculateExperimentMDEData, reqOpts ...goblox.RequestOpt) (*CalculateExperimentMDEResponse, *Response, error) {
	calculateExperimentMDEResponse := new(CalculateExperimentMDEResponse)
	resp, err := s.client.Call(ctx, CalculateExperimentMde.Compile(nil, universeId), data, calculateExperimentMDEResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return calculateExperimentMDEResponse, resp, nil
}

type GetExperimentOperationStatusResponse struct {
	Operation ExperimentOperation `json:"operation"`
}

// GetExperimentOperationStatus will get the status of an experiment async operation.
//
// Required scopes: universe:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/experiments#PublicExperimentation_GetExperimentOperationStatus
//
// [GET] /creator-configs-public-api/v1/experimentation/universes/{universeId}/operations/{operationId}
func (s *ConfigService) GetExperimentOperationStatus(ctx context.Context, universeId, operationId string, reqOpts ...goblox.RequestOpt) (*GetExperimentOperationStatusResponse, *Response, error) {
	getExperimentOperationStatusResponse := new(GetExperimentOperationStatusResponse)
	resp, err := s.client.Call(ctx, GetExperimentOperationStatus.Compile(nil, universeId, operationId), nil, getExperimentOperationStatusResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return getExperimentOperationStatusResponse, resp, nil
}
