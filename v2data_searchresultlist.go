//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"time"
)

// The HTTP method for `execution_url`.
type V2DataSearchResultListDataExecutionSpecHTTPMethod string

// List of values that V2DataSearchResultListDataExecutionSpecHTTPMethod can take
const (
	V2DataSearchResultListDataExecutionSpecHTTPMethodGET  V2DataSearchResultListDataExecutionSpecHTTPMethod = "GET"
	V2DataSearchResultListDataExecutionSpecHTTPMethodPOST V2DataSearchResultListDataExecutionSpecHTTPMethod = "POST"
)

type V2DataSearchResultListMetricsCatalog string

// List of values that V2DataSearchResultListMetricsCatalog can take
const (
	V2DataSearchResultListMetricsCatalogMetrics V2DataSearchResultListMetricsCatalog = "metrics"
)

// How the metric's value is formatted.
type V2DataSearchResultListMetricsMeasureFormat string

// List of values that V2DataSearchResultListMetricsMeasureFormat can take
const (
	V2DataSearchResultListMetricsMeasureFormatCurrency V2DataSearchResultListMetricsMeasureFormat = "currency"
	V2DataSearchResultListMetricsMeasureFormatNumber   V2DataSearchResultListMetricsMeasureFormat = "number"
	V2DataSearchResultListMetricsMeasureFormatPercent  V2DataSearchResultListMetricsMeasureFormat = "percent"
)

type V2DataSearchResultListReportsCatalog string

// List of values that V2DataSearchResultListReportsCatalog can take
const (
	V2DataSearchResultListReportsCatalogReports V2DataSearchResultListReportsCatalog = "reports"
)

// The data type of the column.
type V2DataSearchResultListReportsOutputColumnType string

// List of values that V2DataSearchResultListReportsOutputColumnType can take
const (
	V2DataSearchResultListReportsOutputColumnTypeBigint    V2DataSearchResultListReportsOutputColumnType = "bigint"
	V2DataSearchResultListReportsOutputColumnTypeBoolean   V2DataSearchResultListReportsOutputColumnType = "boolean"
	V2DataSearchResultListReportsOutputColumnTypeDate      V2DataSearchResultListReportsOutputColumnType = "date"
	V2DataSearchResultListReportsOutputColumnTypeDatetime  V2DataSearchResultListReportsOutputColumnType = "datetime"
	V2DataSearchResultListReportsOutputColumnTypeDecimal   V2DataSearchResultListReportsOutputColumnType = "decimal"
	V2DataSearchResultListReportsOutputColumnTypeDouble    V2DataSearchResultListReportsOutputColumnType = "double"
	V2DataSearchResultListReportsOutputColumnTypeInteger   V2DataSearchResultListReportsOutputColumnType = "integer"
	V2DataSearchResultListReportsOutputColumnTypeTimestamp V2DataSearchResultListReportsOutputColumnType = "timestamp"
	V2DataSearchResultListReportsOutputColumnTypeVarchar   V2DataSearchResultListReportsOutputColumnType = "varchar"
)

// The data type of the elements in the array.
type V2DataSearchResultListReportsParametersArrayDetailsElementType string

// List of values that V2DataSearchResultListReportsParametersArrayDetailsElementType can take
const (
	V2DataSearchResultListReportsParametersArrayDetailsElementTypeArray     V2DataSearchResultListReportsParametersArrayDetailsElementType = "array"
	V2DataSearchResultListReportsParametersArrayDetailsElementTypeEnum      V2DataSearchResultListReportsParametersArrayDetailsElementType = "enum"
	V2DataSearchResultListReportsParametersArrayDetailsElementTypeString    V2DataSearchResultListReportsParametersArrayDetailsElementType = "string"
	V2DataSearchResultListReportsParametersArrayDetailsElementTypeTimestamp V2DataSearchResultListReportsParametersArrayDetailsElementType = "timestamp"
)

// The data type of the parameter.
type V2DataSearchResultListReportsParametersType string

// List of values that V2DataSearchResultListReportsParametersType can take
const (
	V2DataSearchResultListReportsParametersTypeArray     V2DataSearchResultListReportsParametersType = "array"
	V2DataSearchResultListReportsParametersTypeEnum      V2DataSearchResultListReportsParametersType = "enum"
	V2DataSearchResultListReportsParametersTypeString    V2DataSearchResultListReportsParametersType = "string"
	V2DataSearchResultListReportsParametersTypeTimestamp V2DataSearchResultListReportsParametersType = "timestamp"
)

type V2DataSearchResultListSchemasCatalog string

// List of values that V2DataSearchResultListSchemasCatalog can take
const (
	V2DataSearchResultListSchemasCatalogSchemas V2DataSearchResultListSchemasCatalog = "schemas"
)

// The data type of the column.
type V2DataSearchResultListSchemasColumnType string

// List of values that V2DataSearchResultListSchemasColumnType can take
const (
	V2DataSearchResultListSchemasColumnTypeBigint    V2DataSearchResultListSchemasColumnType = "bigint"
	V2DataSearchResultListSchemasColumnTypeBoolean   V2DataSearchResultListSchemasColumnType = "boolean"
	V2DataSearchResultListSchemasColumnTypeDate      V2DataSearchResultListSchemasColumnType = "date"
	V2DataSearchResultListSchemasColumnTypeDatetime  V2DataSearchResultListSchemasColumnType = "datetime"
	V2DataSearchResultListSchemasColumnTypeDecimal   V2DataSearchResultListSchemasColumnType = "decimal"
	V2DataSearchResultListSchemasColumnTypeDouble    V2DataSearchResultListSchemasColumnType = "double"
	V2DataSearchResultListSchemasColumnTypeInteger   V2DataSearchResultListSchemasColumnType = "integer"
	V2DataSearchResultListSchemasColumnTypeTimestamp V2DataSearchResultListSchemasColumnType = "timestamp"
	V2DataSearchResultListSchemasColumnTypeVarchar   V2DataSearchResultListSchemasColumnType = "varchar"
)

// The dataset used when submitting an ad-hoc query run.
type V2DataSearchResultListSchemasDataset string

// List of values that V2DataSearchResultListSchemasDataset can take
const (
	V2DataSearchResultListSchemasDatasetAnalytical V2DataSearchResultListSchemasDataset = "analytical"
)

type V2DataSearchResultListDataCatalog string

// List of values that V2DataSearchResultListDataCatalog can take
const (
	V2DataSearchResultListDataCatalogMetrics V2DataSearchResultListDataCatalog = "metrics"
	V2DataSearchResultListDataCatalogReports V2DataSearchResultListDataCatalog = "reports"
	V2DataSearchResultListDataCatalogSchemas V2DataSearchResultListDataCatalog = "schemas"
)

// Domain vocabulary relevant to the request, so that callers interpret the
// returned resources consistently with how Stripe defines these terms.
type V2DataSearchResultListContextBusinessConcept struct {
	// What the term means in the context of this domain.
	Definition string `json:"definition"`
	// The name of the concept, for example "decline reason".
	Term string `json:"term"`
}

// Domain-level intelligence. Present when a known domain matches; omitted
// when the system falls back to direct semantic search.
type V2DataSearchResultListContext struct {
	// Domain vocabulary relevant to the request, so that callers interpret the
	// returned resources consistently with how Stripe defines these terms.
	BusinessConcepts []*V2DataSearchResultListContextBusinessConcept `json:"business_concepts"`
	// The business domain the request was matched to, for example "payments".
	Domain string `json:"domain"`
	// High-level routing guidance for the agent. When `data` is empty, may
	// advise how to refine the request or browse canonical resources.
	Guidance string `json:"guidance"`
}

// How to execute this result.
type V2DataSearchResultListDataExecutionSpec struct {
	// The canonical run endpoint, e.g. "/v2/data/report_runs" or "/v2/data/query_runs".
	ExecutionURL string `json:"execution_url"`
	// The HTTP method for `execution_url`.
	HTTPMethod V2DataSearchResultListDataExecutionSpecHTTPMethod `json:"http_method"`
	// The identifier to submit to `execution_url`, e.g. "report_123" or "table_123".
	ID string `json:"id"`
}

// Dimensions the metric can be filtered and grouped by.
type V2DataSearchResultListMetricsDimension struct {
	// A description of what the dimension represents.
	Description string `json:"description"`
	// The public dimension name accepted by the metric query endpoint.
	Name string `json:"name"`
}

// A pre-defined measure queryable through the metric query endpoint.
type V2DataSearchResultListMetrics struct {
	// Whether the metric query endpoint accepts a `currency` parameter for this
	// metric. False for pure ratios and growth rates.
	AcceptsCurrency bool `json:"accepts_currency"`
	// Dimensions the metric can be filtered and grouped by.
	Dimensions []*V2DataSearchResultListMetricsDimension `json:"dimensions"`
	// How the metric's value is formatted.
	MeasureFormat V2DataSearchResultListMetricsMeasureFormat `json:"measure_format"`
	// The metric namespace. Metrics queried together in one metric query request
	// must share a namespace.
	MetricNamespace string `json:"metric_namespace"`
}

// Shape of the report output when known.
type V2DataSearchResultListReportsOutputColumn struct {
	// What the column represents.
	Description string `json:"description"`
	// The name of the column in the report output.
	Name string `json:"name"`
	// The data type of the column.
	Type V2DataSearchResultListReportsOutputColumnType `json:"type"`
}

// Details about enum elements in the array.
type V2DataSearchResultListReportsParametersArrayDetailsEnumDetails struct {
	// Allowed values of the enum.
	AllowedValues []string `json:"allowed_values"`
}

// For array parameters, details about the array elements.
type V2DataSearchResultListReportsParametersArrayDetails struct {
	// The data type of the elements in the array.
	ElementType V2DataSearchResultListReportsParametersArrayDetailsElementType `json:"element_type"`
	// Details about enum elements in the array.
	EnumDetails *V2DataSearchResultListReportsParametersArrayDetailsEnumDetails `json:"enum_details,omitempty"`
}

// For enum parameters, the list of allowed values.
type V2DataSearchResultListReportsParametersEnumDetails struct {
	// Allowed values of the enum.
	AllowedValues []string `json:"allowed_values"`
}

// Parameter schema, keyed by parameter name, matching the `/v2/data/reports`
// resource.
type V2DataSearchResultListReportsParameters struct {
	// For array parameters, details about the array elements.
	ArrayDetails *V2DataSearchResultListReportsParametersArrayDetails `json:"array_details,omitempty"`
	// Explains the purpose and usage of the parameter.
	Description string `json:"description"`
	// For enum parameters, the list of allowed values.
	EnumDetails *V2DataSearchResultListReportsParametersEnumDetails `json:"enum_details,omitempty"`
	// Indicates whether the parameter must be provided.
	Required bool `json:"required"`
	// The data type of the parameter.
	Type V2DataSearchResultListReportsParametersType `json:"type"`
}

// A Stripe-defined report exposed by `/v2/data/reports`.
type V2DataSearchResultListReports struct {
	// Shape of the report output when known.
	OutputColumns []*V2DataSearchResultListReportsOutputColumn `json:"output_columns"`
	// Parameter schema, keyed by parameter name, matching the `/v2/data/reports`
	// resource.
	Parameters map[string]*V2DataSearchResultListReportsParameters `json:"parameters"`
}

// Columns in other schemas that reference this column as a foreign key.
type V2DataSearchResultListSchemasColumnForeignKeysFrom struct {
	// The name of the referenced column.
	Column string `json:"column"`
	// The identifier of the referenced schema.
	Schema string `json:"schema"`
	// The name of the referenced schema.
	SchemaName string `json:"schema_name"`
}

// Columns in other schemas that this column references as a foreign key.
type V2DataSearchResultListSchemasColumnForeignKeysTo struct {
	// The name of the referenced column.
	Column string `json:"column"`
	// The identifier of the referenced schema.
	Schema string `json:"schema"`
	// The name of the referenced schema.
	SchemaName string `json:"schema_name"`
}

// Compact schema projection.
type V2DataSearchResultListSchemasColumn struct {
	// A description of what the column represents.
	Description string `json:"description"`
	// Columns in other schemas that reference this column as a foreign key.
	ForeignKeysFrom []*V2DataSearchResultListSchemasColumnForeignKeysFrom `json:"foreign_keys_from"`
	// Columns in other schemas that this column references as a foreign key.
	ForeignKeysTo []*V2DataSearchResultListSchemasColumnForeignKeysTo `json:"foreign_keys_to"`
	// Whether the column forms part of the table's primary key.
	IsPrimaryKey bool `json:"is_primary_key"`
	// The name of the column.
	Name string `json:"name"`
	// The data type of the column.
	Type V2DataSearchResultListSchemasColumnType `json:"type"`
}

// Stripe-defined reports that are relevant to this schema.
type V2DataSearchResultListSchemasRelevantReport struct {
	// A description of the `Report`.
	Description string `json:"description"`
	// The unique identifier of the `Report`.
	ID string `json:"id"`
	// The human-readable name of the `Report`.
	Name string `json:"name"`
	// Canonical API path for the `Report`, e.g. "/v2/data/reports/report_123".
	ReferenceURL string `json:"reference_url"`
}

// A queryable table exposed by `/v2/data/schemas`.
type V2DataSearchResultListSchemas struct {
	// Compact schema projection.
	Columns []*V2DataSearchResultListSchemasColumn `json:"columns"`
	// The dataset used when submitting an ad-hoc query run.
	Dataset V2DataSearchResultListSchemasDataset `json:"dataset"`
	// Stripe-defined reports that are relevant to this schema.
	RelevantReports []*V2DataSearchResultListSchemasRelevantReport `json:"relevant_reports"`
}

// Ranked public resources. Always present, including when empty.
type V2DataSearchResultListData struct {
	Catalog V2DataSearchResultListDataCatalog `json:"catalog,omitempty"`
	// What the canonical resource contains.
	Description string `json:"description"`
	// How to execute this result.
	ExecutionSpec *V2DataSearchResultListDataExecutionSpec `json:"execution_spec"`
	// Per-result additional context: constraints, interpretation rules, and gotchas.
	ExtendedDescription string `json:"extended_description,omitempty"`
	// Canonical resource identity.
	ID string `json:"id"`
	// A pre-defined measure queryable through the metric query endpoint.
	Metrics *V2DataSearchResultListMetrics `json:"metrics,omitempty"`
	// Human-readable name of the canonical resource.
	Name string `json:"name"`
	// The canonical resource type this result points at, e.g. "v2.data.schema"
	// or "v2.data.report". May name a resource that is not yet publicly
	// retrievable, in which case `reference_url` is omitted. Distinct from
	// `object`, which identifies this result itself.
	ObjectType string `json:"object_type"`
	// Canonical API path for retrieving the complete resource, e.g.
	// "/v2/data/reports/report_123".
	ReferenceURL string `json:"reference_url"`
	// Present for schemas when known; omitted for resource types without data
	// freshness.
	RefreshedAt time.Time `json:"refreshed_at"`
	// Relative ordering signal from the search service (0-1). Not a
	// calibrated guarantee; callers must not use it as a configurable cutoff.
	RelevanceScore decimal.Decimal `json:"relevance_score"`
	// A Stripe-defined report exposed by `/v2/data/reports`.
	Reports *V2DataSearchResultListReports `json:"reports,omitempty"`
	// A queryable table exposed by `/v2/data/schemas`.
	Schemas *V2DataSearchResultListSchemas `json:"schemas,omitempty"`
}

// A ranked set of public analytical resources matching a natural-language query.
type V2DataSearchResultList struct {
	APIResource
	// Domain-level intelligence. Present when a known domain matches; omitted
	// when the system falls back to direct semantic search.
	Context *V2DataSearchResultListContext `json:"context,omitempty"`
	// Ranked public resources. Always present, including when empty.
	Data []*V2DataSearchResultListData `json:"data"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Token for the next page.
	NextPageURL string `json:"next_page_url,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Token for the previous page.
	PreviousPageURL string `json:"previous_page_url,omitempty"`
}

// UnmarshalJSON handles deserialization of a V2DataSearchResultListData.
// This custom unmarshaling is needed because only the variant named by the
// discriminator is kept; any others the payload carries are cleared.
func (v *V2DataSearchResultListData) UnmarshalJSON(data []byte) error {
	type v2DataSearchResultListData V2DataSearchResultListData
	var u v2DataSearchResultListData
	if err := json.Unmarshal(data, &u); err != nil {
		return err
	}

	*v = V2DataSearchResultListData(u)
	var empty V2DataSearchResultListData
	switch v.Catalog {
	case V2DataSearchResultListDataCatalogMetrics:
		v.Reports = empty.Reports
		v.Schemas = empty.Schemas
	case V2DataSearchResultListDataCatalogReports:
		v.Metrics = empty.Metrics
		v.Schemas = empty.Schemas
	case V2DataSearchResultListDataCatalogSchemas:
		v.Metrics = empty.Metrics
		v.Reports = empty.Reports
	default:
		v.Metrics = empty.Metrics
		v.Reports = empty.Reports
		v.Schemas = empty.Schemas
	}
	return nil
}
