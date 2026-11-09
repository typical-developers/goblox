package opencloud

import (
	"context"
	"net/http"

	"github.com/luckfire/goblox"
	"github.com/luckfire/goblox/auth"
)

var (
	ListAlertConfigs = goblox.NewEndpoint(
		http.MethodGet,
		"/analytics-alert-control-plane/v1/resource/{resourceType}/id/{resourceId}/alerts",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics.alert:read")

	CreateAlertConfig = goblox.NewEndpoint(
		http.MethodPost,
		"/analytics-alert-control-plane/v1/resource/{resourceType}/id/{resourceId}/alerts",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics.alert:write")

	ListAlertIncidents = goblox.NewEndpoint(
		http.MethodGet,
		"/analytics-alert-control-plane/v1/resource/{resourceType}/id/{resourceId}/alerts/incidents",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics.alert:read")

	DeleteAlertConfig = goblox.NewEndpoint(
		http.MethodDelete,
		"/analytics-alert-control-plane/v1/resource/{resourceType}/id/{resourceId}/alerts/{alertId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics.alert:write")

	UpdateAlertConfig = goblox.NewEndpoint(
		http.MethodPatch,
		"/analytics-alert-control-plane/v1/resource/{resourceType}/id/{resourceId}/alerts/{alertId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics.alert:write")

	QueryUniverseBenchmarks = goblox.NewEndpoint(
		http.MethodPost,
		"/analytics-benchmark/v1/universes/{universeId}/benchmarks",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:read")

	QueryDimensionValues = goblox.NewEndpoint(
		http.MethodPost,
		"/analytics-query-api/v1/universes/{universeId}/dimension-values",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:read")

	QueryMetrics = goblox.NewEndpoint(
		http.MethodPost,
		"/analytics-query-api/v1/universes/{universeId}/metrics",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:read")

	GetDimensionValuesOperation = goblox.NewEndpoint(
		http.MethodGet,
		"/analytics-query-api/v1/universes/{universeId}/operations/dimension-values/{operationId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:read")

	GetMetricsOperation = goblox.NewEndpoint(
		http.MethodGet,
		"/analytics-query-api/v1/universes/{universeId}/operations/metrics/{operationId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:read")

	ListUniverseLogRegexes = goblox.NewEndpoint(
		http.MethodGet,
		"/log-attribute-http-service/v1/universe-regexes",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:read")

	CreateUniverseLogRegex = goblox.NewEndpoint(
		http.MethodPost,
		"/log-attribute-http-service/v1/universe-regexes",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:write")

	DeleteUniverseLogRegex = goblox.NewEndpoint(
		http.MethodDelete,
		"/log-attribute-http-service/v1/universe-regexes/{id}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:write")

	UpdateUniverseLogRegex = goblox.NewEndpoint(
		http.MethodPut,
		"/log-attribute-http-service/v1/universe-regexes/{id}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:write")

	IgnoreUniverseLogRegex = goblox.NewEndpoint(
		http.MethodPut,
		"/log-attribute-http-service/v1/universe-regexes/{id}/ignore",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:write")

	ReorderUniverseLogRegex = goblox.NewEndpoint(
		http.MethodPut,
		"/log-attribute-http-service/v1/universe-regexes/{id}/reorder",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:write")

	GetUniverseLogDetails = goblox.NewEndpoint(
		http.MethodGet,
		"/universe-performance-raqi/v1/events/logging/universe/{universeId}/log-details",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.analytics:read")
)

// AnalyticsService handles communication with the analytics, alerting and log endpoints.
//
// Roblox Open Cloud API Docs: https://create.roblox.com/docs/en-us/cloud
type AnalyticsService service

type AlertFiringStatus string

const (
	AlertFiringStatusOk     AlertFiringStatus = "OK"
	AlertFiringStatusFiring AlertFiringStatus = "Firing"
)

type AlertSeverity string

const (
	AlertSeveritySev0 AlertSeverity = "SEV_0"
	AlertSeveritySev1 AlertSeverity = "SEV_1"
	AlertSeveritySev2 AlertSeverity = "SEV_2"
)

type ListAlertConfigsOptions struct {
	IDs          []string           `url:"ids,omitempty"`
	FiringStatus *AlertFiringStatus `url:"firingStatus,omitempty"`
	Severities   []AlertSeverity    `url:"severities,omitempty"`
	Metrics      []string           `url:"metrics,omitempty"`
}

type AlertResourceType string

const (
	AlertResourceTypeUniverse AlertResourceType = "Universe"
)

type AlertInterval string

const (
	AlertIntervalOneMinute AlertInterval = "OneMinute"
	AlertIntervalHalfHour  AlertInterval = "HalfHour"
	AlertIntervalOneHour   AlertInterval = "OneHour"
	AlertIntervalOneDay    AlertInterval = "OneDay"
)

type AlertQueryFilter struct {
	Dimension *string  `json:"dimension,omitempty"`
	Values    []string `json:"values,omitempty"`
}

type AlertQueryBreakdown struct {
	Dimensions []string `json:"dimensions,omitempty"`
}

type ConditionOperator string

const (
	ConditionOperatorGt  ConditionOperator = "Gt"
	ConditionOperatorGte ConditionOperator = "Gte"
	ConditionOperatorLt  ConditionOperator = "Lt"
	ConditionOperatorLte ConditionOperator = "Lte"
)

type EvaluationMode string

const (
	EvaluationModeAbsolute         EvaluationMode = "Absolute"
	EvaluationModePeriodOverPeriod EvaluationMode = "PeriodOverPeriod"
)

type AlertCondition struct {
	Operator               ConditionOperator `json:"operator"`
	Threshold              float64           `json:"threshold"`
	EvaluationMode         EvaluationMode    `json:"evaluationMode"`
	PeriodOffsetMultiplier int               `json:"periodOffsetMultiplier"`
}

type AlertConfigState string

const (
	AlertConfigStateEnabled        AlertConfigState = "Enabled"
	AlertConfigStateDisabled       AlertConfigState = "Disabled"
	AlertConfigStatePausedByRoblox AlertConfigState = "PausedByRoblox"
	AlertConfigStateSyncing        AlertConfigState = "Syncing"
	AlertConfigStateError          AlertConfigState = "Error"
)

type WebhookReceiver struct {
	WebhookConfigurationID *string `json:"webhookConfigurationId,omitempty"`
}

type WebhookReceiverConfig struct {
	Receivers []WebhookReceiver `json:"receivers,omitempty"`
}

type AlertDetailResponse struct {
	AlertID                string                `json:"alertId"`
	ResourceType           AlertResourceType     `json:"resourceType"`
	ResourceID             string                `json:"resourceId"`
	Name                   string                `json:"name"`
	Metric                 string                `json:"metric"`
	Description            string                `json:"description"`
	Severity               AlertSeverity         `json:"severity"`
	Interval               AlertInterval         `json:"interval"`
	ConsecutiveOccurrences int                   `json:"consecutiveOccurrences"`
	Filter                 []AlertQueryFilter    `json:"filter"`
	Breakdown              []AlertQueryBreakdown `json:"breakdown"`
	Condition              AlertCondition        `json:"condition"`
	ConfigState            AlertConfigState      `json:"configState"`
	WebhookReceiverConfig  WebhookReceiverConfig `json:"webhookReceiverConfig"`
	FiringStatus           AlertFiringStatus     `json:"firingStatus"`
	LastFiredAt            string                `json:"lastFiredAt"`
	CreatedAt              string                `json:"createdAt"`
	LastModifiedAt         string                `json:"lastModifiedAt"`
	LastModifiedBy         string                `json:"lastModifiedBy"`
}

// ListAlertConfigs will return the alert configurations for a resource.
//
// Required scopes: universe.analytics.alert:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#get_analytics_alert_control_plane_v1_resource__resourceType__id__resourceId__alerts
//
// [GET] /analytics-alert-control-plane/v1/resource/{resourceType}/id/{resourceId}/alerts
func (s *AnalyticsService) ListAlertConfigs(ctx context.Context, resourceType, resourceId string, opts *ListAlertConfigsOptions, reqOpts ...goblox.RequestOpt) ([]AlertDetailResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	var result []AlertDetailResponse
	resp, err := s.client.Call(ctx, ListAlertConfigs.Compile(q, resourceType, resourceId), nil, &result, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

type AlertConditionInput struct {
	Operator               *ConditionOperator `json:"operator,omitempty"`
	Threshold              *float64           `json:"threshold,omitempty"`
	EvaluationMode         *EvaluationMode    `json:"evaluationMode,omitempty"`
	PeriodOffsetMultiplier *int               `json:"periodOffsetMultiplier,omitempty"`
}

type CreateAlertConfigRequest struct {
	Name                   *string                `json:"name,omitempty"`
	Metric                 *string                `json:"metric,omitempty"`
	Description            *string                `json:"description,omitempty"`
	Severity               *AlertSeverity         `json:"severity,omitempty"`
	Interval               *AlertInterval         `json:"interval,omitempty"`
	ConsecutiveOccurrences *int                   `json:"consecutiveOccurrences,omitempty"`
	Breakdown              []AlertQueryBreakdown  `json:"breakdown,omitempty"`
	Filter                 []AlertQueryFilter     `json:"filter,omitempty"`
	Condition              *AlertConditionInput   `json:"condition,omitempty"`
	WebhookReceiverConfig  *WebhookReceiverConfig `json:"webhookReceiverConfig,omitempty"`
}

// CreateAlertConfig will create a new alert configuration for a resource.
//
// Required scopes: universe.analytics.alert:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#post_analytics_alert_control_plane_v1_resource__resourceType__id__resourceId__alerts
//
// [POST] /analytics-alert-control-plane/v1/resource/{resourceType}/id/{resourceId}/alerts
func (s *AnalyticsService) CreateAlertConfig(ctx context.Context, resourceType, resourceId string, data CreateAlertConfigRequest, reqOpts ...goblox.RequestOpt) (*AlertDetailResponse, *Response, error) {
	alertDetailResponse := new(AlertDetailResponse)
	resp, err := s.client.Call(ctx, CreateAlertConfig.Compile(nil, resourceType, resourceId), data, alertDetailResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return alertDetailResponse, resp, nil
}

type ListAlertIncidentsOptions struct {
	StartTime  string          `url:"startTime"` // Required.
	EndTime    string          `url:"endTime"`   // Required.
	Metric     *string         `url:"metric,omitempty"`
	AlertIDs   []string        `url:"alertIds,omitempty"`
	Severities []AlertSeverity `url:"severities,omitempty"`
}

type FiringDimensionResponse struct {
	Dimension string `json:"dimension"`
	Value     string `json:"value"`
}

type FiringConditionResponse struct {
	FiringValue     float64                 `json:"firingValue"`
	FiringDimension FiringDimensionResponse `json:"firingDimension"`
}

type FiringMetadataResponse struct {
	Filter          []AlertQueryFilter        `json:"filter"`
	Breakdown       []AlertQueryBreakdown     `json:"breakdown"`
	Condition       AlertCondition            `json:"condition"`
	FiringCondition []FiringConditionResponse `json:"firingCondition"`
}

type AlertConfigSummary struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Metric      string                `json:"metric"`
	Severity    AlertSeverity         `json:"severity"`
	Description string                `json:"description"`
	Interval    AlertInterval         `json:"interval"`
	Filter      []AlertQueryFilter    `json:"filter"`
	Breakdown   []AlertQueryBreakdown `json:"breakdown"`
	Condition   AlertCondition        `json:"condition"`
}

type IncidentDetailResponse struct {
	ID                    string                 `json:"id"`
	ResourceType          AlertResourceType      `json:"resourceType"`
	ResourceID            string                 `json:"resourceId"`
	Status                AlertFiringStatus      `json:"status"`
	OpenedAt              string                 `json:"openedAt"`
	OpenedWindowStartAt   string                 `json:"openedWindowStartAt"`
	ResolvedAt            string                 `json:"resolvedAt"`
	ResolvedWindowStartAt string                 `json:"resolvedWindowStartAt"`
	FirstFiringMetadata   FiringMetadataResponse `json:"firstFiringMetadata"`
	LatestFiringMetadata  FiringMetadataResponse `json:"latestFiringMetadata"`
	AlertConfig           AlertConfigSummary     `json:"alertConfig"`
}

// ListAlertIncidents will return alert incidents for a resource whose active window overlaps the specified time.
//
// Required scopes: universe.analytics.alert:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#get_analytics_alert_control_plane_v1_resource__resourceType__id__resourceId__alerts_incidents
//
// [GET] /analytics-alert-control-plane/v1/resource/{resourceType}/id/{resourceId}/alerts/incidents
func (s *AnalyticsService) ListAlertIncidents(ctx context.Context, resourceType, resourceId string, opts *ListAlertIncidentsOptions, reqOpts ...goblox.RequestOpt) ([]IncidentDetailResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	var result []IncidentDetailResponse
	resp, err := s.client.Call(ctx, ListAlertIncidents.Compile(q, resourceType, resourceId), nil, &result, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// DeleteAlertConfig will delete an alert configuration.
//
// Required scopes: universe.analytics.alert:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#delete_analytics_alert_control_plane_v1_resource__resourceType__id__resourceId__alerts__alertId_
//
// [DELETE] /analytics-alert-control-plane/v1/resource/{resourceType}/id/{resourceId}/alerts/{alertId}
func (s *AnalyticsService) DeleteAlertConfig(ctx context.Context, resourceType, resourceId, alertId string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, DeleteAlertConfig.Compile(nil, resourceType, resourceId, alertId), nil, nil, reqOpts...)
}

type UpdateAlertConfigRequest struct {
	Name                   *string                `json:"name,omitempty"`
	Metric                 *string                `json:"metric,omitempty"`
	Description            *string                `json:"description,omitempty"`
	Severity               *AlertSeverity         `json:"severity,omitempty"`
	Interval               *AlertInterval         `json:"interval,omitempty"`
	ConsecutiveOccurrences *int                   `json:"consecutiveOccurrences,omitempty"`
	Breakdown              []AlertQueryBreakdown  `json:"breakdown,omitempty"`
	Filter                 []AlertQueryFilter     `json:"filter,omitempty"`
	Condition              *AlertConditionInput   `json:"condition,omitempty"`
	WebhookReceiverConfig  *WebhookReceiverConfig `json:"webhookReceiverConfig,omitempty"`
	ConfigState            *AlertConfigState      `json:"configState,omitempty"`
}

// UpdateAlertConfig will update an existing alert configuration.
//
// Required scopes: universe.analytics.alert:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#patch_analytics_alert_control_plane_v1_resource__resourceType__id__resourceId__alerts__alertId_
//
// [PATCH] /analytics-alert-control-plane/v1/resource/{resourceType}/id/{resourceId}/alerts/{alertId}
func (s *AnalyticsService) UpdateAlertConfig(ctx context.Context, resourceType, resourceId, alertId string, data UpdateAlertConfigRequest, reqOpts ...goblox.RequestOpt) (*AlertDetailResponse, *Response, error) {
	alertDetailResponse := new(AlertDetailResponse)
	resp, err := s.client.Call(ctx, UpdateAlertConfig.Compile(nil, resourceType, resourceId, alertId), data, alertDetailResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return alertDetailResponse, resp, nil
}

type BreakdownFilter struct {
	Dimension *string `json:"dimension,omitempty"`
	Value     *string `json:"value,omitempty"`
}

type BenchmarkType string

const (
	BenchmarkTypeBenchmarkTypeInvalid        BenchmarkType = "BENCHMARK_TYPE_INVALID"
	BenchmarkTypeBenchmarkTypeGenre          BenchmarkType = "BENCHMARK_TYPE_GENRE"
	BenchmarkTypeBenchmarkTypeSimilarity     BenchmarkType = "BENCHMARK_TYPE_SIMILARITY"
	BenchmarkTypeBenchmarkTypeTopExperiences BenchmarkType = "BENCHMARK_TYPE_TOP_EXPERIENCES"
)

type BenchmarkQuery struct {
	Metric             *string           `json:"metric,omitempty"`
	StartTime          *string           `json:"startTime,omitempty"`
	EndTime            *string           `json:"endTime,omitempty"`
	Filter             []BreakdownFilter `json:"filter,omitempty"`
	BenchmarkType      *BenchmarkType    `json:"benchmarkType,omitempty"`
	Percentiles        []int             `json:"percentiles,omitempty"`
	BenchmarkVariantID *string           `json:"benchmarkVariantId,omitempty"`
}

type AnalyticsBreakdownValue struct {
	Dimension    string `json:"dimension"`
	Value        string `json:"value"`
	DisplayValue string `json:"displayValue"`
}

type DataPointMetadata struct {
	BenchmarkType BenchmarkType `json:"benchmarkType"`
	Genre         string        `json:"genre"`
}

type AnalyticsDataPointV2 struct {
	Time     string            `json:"time"`
	Value    float64           `json:"value"`
	Metadata DataPointMetadata `json:"metadata"`
}

type AnalyticsMetricValueV2 struct {
	BreakdownValue []AnalyticsBreakdownValue `json:"breakdownValue"`
	DataPoints     []AnalyticsDataPointV2    `json:"dataPoints"`
}

type QueryBenchmarkResult struct {
	Values []AnalyticsMetricValueV2 `json:"values"`
}

type BenchmarkObjective string

const (
	BenchmarkObjectiveBenchmarkObjectiveInvalid  BenchmarkObjective = "BENCHMARK_OBJECTIVE_INVALID"
	BenchmarkObjectiveBenchmarkObjectiveMaximize BenchmarkObjective = "BENCHMARK_OBJECTIVE_MAXIMIZE"
	BenchmarkObjectiveBenchmarkObjectiveMinimize BenchmarkObjective = "BENCHMARK_OBJECTIVE_MINIMIZE"
)

type QueryUniverseBenchmarkResponse struct {
	Result         QueryBenchmarkResult `json:"result"`
	AvailableTypes []BenchmarkType      `json:"availableTypes"`
	Objective      BenchmarkObjective   `json:"objective"`
}

// QueryUniverseBenchmarks will return daily benchmark percentiles for a metric, comparing a universe against
// other experiences on Roblox.
//
// Required scopes: universe.analytics:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#post_analytics_benchmark_v1_universes__universe_id__benchmarks
//
// [POST] /analytics-benchmark/v1/universes/{universeId}/benchmarks
func (s *AnalyticsService) QueryUniverseBenchmarks(ctx context.Context, universeId string, data BenchmarkQuery, reqOpts ...goblox.RequestOpt) (*QueryUniverseBenchmarkResponse, *Response, error) {
	queryUniverseBenchmarkResponse := new(QueryUniverseBenchmarkResponse)
	resp, err := s.client.Call(ctx, QueryUniverseBenchmarks.Compile(nil, universeId), data, queryUniverseBenchmarkResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return queryUniverseBenchmarkResponse, resp, nil
}

type FilterOperation string

const (
	FilterOperationIn                 FilterOperation = "In"
	FilterOperationNotIn              FilterOperation = "NotIn"
	FilterOperationGreaterThan        FilterOperation = "GreaterThan"
	FilterOperationGreaterThanOrEqual FilterOperation = "GreaterThanOrEqual"
	FilterOperationLessThan           FilterOperation = "LessThan"
	FilterOperationLessThanOrEqual    FilterOperation = "LessThanOrEqual"
	FilterOperationMatch              FilterOperation = "Match"
)

type QueryFilter struct {
	Dimension *string          `json:"dimension,omitempty"`
	Values    []string         `json:"values,omitempty"`
	Operation *FilterOperation `json:"operation,omitempty"`
}

type MetricGranularity string

const (
	MetricGranularityOneMinute MetricGranularity = "OneMinute"
	MetricGranularityHalfHour  MetricGranularity = "HalfHour"
	MetricGranularityOneHour   MetricGranularity = "OneHour"
	MetricGranularityOneDay    MetricGranularity = "OneDay"
	MetricGranularityOneWeek   MetricGranularity = "OneWeek"
	MetricGranularityOneMonth  MetricGranularity = "OneMonth"
	MetricGranularityNone      MetricGranularity = "None"
)

type DimensionValuesRequest struct {
	Metric      *string            `json:"metric,omitempty"`
	Dimensions  []string           `json:"dimensions,omitempty"`
	StartTime   *string            `json:"startTime,omitempty"`
	EndTime     *string            `json:"endTime,omitempty"`
	Filter      []QueryFilter      `json:"filter,omitempty"`
	Granularity *MetricGranularity `json:"granularity,omitempty"`
	Limit       *int               `json:"limit,omitempty"`
}

type DimensionValue struct {
	Value        string `json:"value"`
	DisplayValue string `json:"displayValue"`
}

type DimensionValues struct {
	Dimension string           `json:"dimension"`
	Values    []DimensionValue `json:"values"`
}

type DimensionValuesResponse struct {
	Values []DimensionValues `json:"values"`
}

type AnalyticsOperationMetadata struct {
	CreatedTime string `json:"createdTime"`
}

type DimensionValuesOperationResult struct {
	Path     string                     `json:"path"`
	Done     bool                       `json:"done"`
	Response DimensionValuesResponse    `json:"response"`
	Metadata AnalyticsOperationMetadata `json:"metadata"`
}

// QueryDimensionValues will query dimension values for a universe.
//
// Required scopes: universe.analytics:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#post_analytics_query_api_v1_universes__universe_id__dimension_values
//
// [POST] /analytics-query-api/v1/universes/{universeId}/dimension-values
func (s *AnalyticsService) QueryDimensionValues(ctx context.Context, universeId string, data DimensionValuesRequest, reqOpts ...goblox.RequestOpt) (*DimensionValuesOperationResult, *Response, error) {
	dimensionValuesOperationResult := new(DimensionValuesOperationResult)
	resp, err := s.client.Call(ctx, QueryDimensionValues.Compile(nil, universeId), data, dimensionValuesOperationResult, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dimensionValuesOperationResult, resp, nil
}

type QueryRequest struct {
	Metric      *string            `json:"metric,omitempty"`
	Granularity *MetricGranularity `json:"granularity,omitempty"`
	StartTime   *string            `json:"startTime,omitempty"`
	EndTime     *string            `json:"endTime,omitempty"`
	Breakdown   []string           `json:"breakdown,omitempty"`
	Filter      []QueryFilter      `json:"filter,omitempty"`
	Limit       *int               `json:"limit,omitempty"`
}

type BreakdownValue struct {
	Dimension    string `json:"dimension"`
	Value        string `json:"value"`
	DisplayValue string `json:"displayValue"`
}

type DataStatus string

const (
	DataStatusValid                       DataStatus = "Valid"
	DataStatusProjected                   DataStatus = "Projected"
	DataStatusNotStatisticallySignificant DataStatus = "NotStatisticallySignificant"
)

type AnalyticsDataPoint struct {
	Time         string     `json:"time"`
	Value        float64    `json:"value"`
	StringValues []string   `json:"stringValues"`
	Status       DataStatus `json:"status"`
}

type AnalyticsMetricValue struct {
	Breakdowns []BreakdownValue     `json:"breakdowns"`
	DataPoints []AnalyticsDataPoint `json:"dataPoints"`
}

type QueryResponse struct {
	Values []AnalyticsMetricValue `json:"values"`
}

type QueryOperationResult struct {
	Path     string                     `json:"path"`
	Done     bool                       `json:"done"`
	Response QueryResponse              `json:"response"`
	Metadata AnalyticsOperationMetadata `json:"metadata"`
}

// QueryMetrics will query time series metric data for a universe.
//
// Required scopes: universe.analytics:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#post_analytics_query_api_v1_universes__universe_id__metrics
//
// [POST] /analytics-query-api/v1/universes/{universeId}/metrics
func (s *AnalyticsService) QueryMetrics(ctx context.Context, universeId string, data QueryRequest, reqOpts ...goblox.RequestOpt) (*QueryOperationResult, *Response, error) {
	queryOperationResult := new(QueryOperationResult)
	resp, err := s.client.Call(ctx, QueryMetrics.Compile(nil, universeId), data, queryOperationResult, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return queryOperationResult, resp, nil
}

// GetDimensionValuesOperation will retrieve the result of a long-running dimension values operation.
//
// Required scopes: universe.analytics:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#get_analytics_query_api_v1_universes__universe_id__operations_dimension_values__operation_id_
//
// [GET] /analytics-query-api/v1/universes/{universeId}/operations/dimension-values/{operationId}
func (s *AnalyticsService) GetDimensionValuesOperation(ctx context.Context, universeId, operationId string, reqOpts ...goblox.RequestOpt) (*DimensionValuesOperationResult, *Response, error) {
	dimensionValuesOperationResult := new(DimensionValuesOperationResult)
	resp, err := s.client.Call(ctx, GetDimensionValuesOperation.Compile(nil, universeId, operationId), nil, dimensionValuesOperationResult, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return dimensionValuesOperationResult, resp, nil
}

// GetMetricsOperation will retrieve the result of a long-running metrics query operation.
//
// Required scopes: universe.analytics:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#get_analytics_query_api_v1_universes__universe_id__operations_metrics__operation_id_
//
// [GET] /analytics-query-api/v1/universes/{universeId}/operations/metrics/{operationId}
func (s *AnalyticsService) GetMetricsOperation(ctx context.Context, universeId, operationId string, reqOpts ...goblox.RequestOpt) (*QueryOperationResult, *Response, error) {
	queryOperationResult := new(QueryOperationResult)
	resp, err := s.client.Call(ctx, GetMetricsOperation.Compile(nil, universeId, operationId), nil, queryOperationResult, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return queryOperationResult, resp, nil
}

type RegexStatus string

const (
	RegexStatusInvalid        RegexStatus = "Invalid"
	RegexStatusCreatorCreated RegexStatus = "CreatorCreated"
	RegexStatusAutoGenerated  RegexStatus = "AutoGenerated"
	RegexStatusCreatorIgnored RegexStatus = "CreatorIgnored"
)

type ListUniverseLogRegexesOptions struct {
	UniverseID *int64       `url:"universeId,omitempty"`
	Status     *RegexStatus `url:"status,omitempty"`
}

type RegexOperation string

const (
	RegexOperationInvalid RegexOperation = "Invalid"
	RegexOperationIgnore  RegexOperation = "Ignore"
	RegexOperationGroup   RegexOperation = "Group"
)

type UniverseRegexHTTPResponse struct {
	ID             int64          `json:"id"`
	UniverseID     int64          `json:"universeId"`
	RegexStatus    RegexStatus    `json:"regexStatus"`
	Pattern        string         `json:"pattern"`
	Output         string         `json:"output"`
	Order          int            `json:"order"`
	RegexOperation RegexOperation `json:"regexOperation"`
	CreatedTime    string         `json:"createdTime"`
	UpdatedTime    string         `json:"updatedTime"`
	MatchedCount   int            `json:"matchedCount"`
	SampleMessage  string         `json:"sampleMessage"`
}

type FindUniverseRegexesHTTPResponse struct {
	UniverseRegexes []UniverseRegexHTTPResponse `json:"universeRegexes"`
}

// ListUniverseLogRegexes will list universe log regexes.
//
// Required scopes: universe.analytics:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#LogAttributeService_Find
//
// [GET] /log-attribute-http-service/v1/universe-regexes
func (s *AnalyticsService) ListUniverseLogRegexes(ctx context.Context, opts *ListUniverseLogRegexesOptions, reqOpts ...goblox.RequestOpt) (*FindUniverseRegexesHTTPResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	findUniverseRegexesHTTPResponse := new(FindUniverseRegexesHTTPResponse)
	resp, err := s.client.Call(ctx, ListUniverseLogRegexes.Compile(q), nil, findUniverseRegexesHTTPResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return findUniverseRegexesHTTPResponse, resp, nil
}

type CreateUniverseRegexHTTPRequest struct {
	UniverseID     *int64          `json:"universeId,omitempty"`
	Pattern        *string         `json:"pattern,omitempty"`
	Output         *string         `json:"output,omitempty"`
	RegexOperation *RegexOperation `json:"regexOperation,omitempty"`
}

type UniverseRegexMutationHTTPResponse struct {
	Order int `json:"order"`
}

// CreateUniverseLogRegex will create universe log regex.
//
// Required scopes: universe.analytics:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#LogAttributeService_Create
//
// [POST] /log-attribute-http-service/v1/universe-regexes
func (s *AnalyticsService) CreateUniverseLogRegex(ctx context.Context, data CreateUniverseRegexHTTPRequest, reqOpts ...goblox.RequestOpt) (*UniverseRegexMutationHTTPResponse, *Response, error) {
	universeRegexMutationHTTPResponse := new(UniverseRegexMutationHTTPResponse)
	resp, err := s.client.Call(ctx, CreateUniverseLogRegex.Compile(nil), data, universeRegexMutationHTTPResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return universeRegexMutationHTTPResponse, resp, nil
}

// DeleteUniverseLogRegex will delete universe log regex.
//
// Required scopes: universe.analytics:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#LogAttributeService_Delete
//
// [DELETE] /log-attribute-http-service/v1/universe-regexes/{id}
func (s *AnalyticsService) DeleteUniverseLogRegex(ctx context.Context, id string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, DeleteUniverseLogRegex.Compile(nil, id), nil, nil, reqOpts...)
}

type UpdateUniverseRegexHTTPRequest struct {
	Pattern        *string         `json:"pattern,omitempty"`
	Output         *string         `json:"output,omitempty"`
	RegexOperation *RegexOperation `json:"regexOperation,omitempty"`
}

// UpdateUniverseLogRegex will update universe log regex.
//
// Required scopes: universe.analytics:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#LogAttributeService_Update
//
// [PUT] /log-attribute-http-service/v1/universe-regexes/{id}
func (s *AnalyticsService) UpdateUniverseLogRegex(ctx context.Context, id string, data UpdateUniverseRegexHTTPRequest, reqOpts ...goblox.RequestOpt) (*UniverseRegexMutationHTTPResponse, *Response, error) {
	universeRegexMutationHTTPResponse := new(UniverseRegexMutationHTTPResponse)
	resp, err := s.client.Call(ctx, UpdateUniverseLogRegex.Compile(nil, id), data, universeRegexMutationHTTPResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return universeRegexMutationHTTPResponse, resp, nil
}

// IgnoreUniverseLogRegex will ignore universe log regex.
//
// Required scopes: universe.analytics:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#LogAttributeService_Ignore
//
// [PUT] /log-attribute-http-service/v1/universe-regexes/{id}/ignore
func (s *AnalyticsService) IgnoreUniverseLogRegex(ctx context.Context, id string, reqOpts ...goblox.RequestOpt) (*Response, error) {
	return s.client.Call(ctx, IgnoreUniverseLogRegex.Compile(nil, id), nil, nil, reqOpts...)
}

type ReorderUniverseRegexHTTPRequest struct {
	Order *int `json:"order,omitempty"`
}

// ReorderUniverseLogRegex will reorder universe log regex.
//
// Required scopes: universe.analytics:write
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#LogAttributeService_Reorder
//
// [PUT] /log-attribute-http-service/v1/universe-regexes/{id}/reorder
func (s *AnalyticsService) ReorderUniverseLogRegex(ctx context.Context, id string, data ReorderUniverseRegexHTTPRequest, reqOpts ...goblox.RequestOpt) (*UniverseRegexMutationHTTPResponse, *Response, error) {
	universeRegexMutationHTTPResponse := new(UniverseRegexMutationHTTPResponse)
	resp, err := s.client.Call(ctx, ReorderUniverseLogRegex.Compile(nil, id), data, universeRegexMutationHTTPResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return universeRegexMutationHTTPResponse, resp, nil
}

type GetUniverseLogDetailsOptions struct {
	MessageHashIDs        []string `url:"messageHashIds,omitempty"`
	Keyword               *string  `url:"keyword,omitempty"`
	StartTime             *string  `url:"startTime,omitempty"`
	EndTime               *string  `url:"endTime,omitempty"`
	PlaceID               *int64   `url:"placeId,omitempty"`
	PlaceVersions         []int64  `url:"placeVersions,omitempty"`
	FirstSeenPlaceVersion *int64   `url:"firstSeenPlaceVersion,omitempty"`
}

type ErrorLogDetail struct {
	ID                    string `json:"id"`
	LogSeverity           string `json:"logSeverity"`
	LogSource             string `json:"logSource"`
	Os                    string `json:"os"`
	Platform              string `json:"platform"`
	Message               string `json:"message"`
	Stacktrace            string `json:"stacktrace"`
	FirstSeenUTC          string `json:"firstSeenUtc"`
	FirstSeenPlaceVersion int64  `json:"firstSeenPlaceVersion"`
}

type ErrorLogDetailResponse struct {
	Values []ErrorLogDetail `json:"values"`
}

// GetUniverseLogDetails will get logs based on message hash ids, keywords, placeVersions, start/end time.
//
// Required scopes: universe.analytics:read
//
// This endpoint is marked as experimental by Roblox and may change without notice.
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/analytics#ErrorLoggingDetail_GetLogs
//
// [GET] /universe-performance-raqi/v1/events/logging/universe/{universeId}/log-details
func (s *AnalyticsService) GetUniverseLogDetails(ctx context.Context, universeId string, opts *GetUniverseLogDetailsOptions, reqOpts ...goblox.RequestOpt) (*ErrorLogDetailResponse, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	errorLogDetailResponse := new(ErrorLogDetailResponse)
	resp, err := s.client.Call(ctx, GetUniverseLogDetails.Compile(q, universeId), nil, errorLogDetailResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return errorLogDetailResponse, resp, nil
}
