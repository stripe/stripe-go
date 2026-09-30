//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The data type of the elements in the array.
type V2DataReportParametersArrayDetailsElementType string

// List of values that V2DataReportParametersArrayDetailsElementType can take
const (
	V2DataReportParametersArrayDetailsElementTypeArray     V2DataReportParametersArrayDetailsElementType = "array"
	V2DataReportParametersArrayDetailsElementTypeEnum      V2DataReportParametersArrayDetailsElementType = "enum"
	V2DataReportParametersArrayDetailsElementTypeString    V2DataReportParametersArrayDetailsElementType = "string"
	V2DataReportParametersArrayDetailsElementTypeTimestamp V2DataReportParametersArrayDetailsElementType = "timestamp"
)

// The data type of the parameter.
type V2DataReportParametersType string

// List of values that V2DataReportParametersType can take
const (
	V2DataReportParametersTypeArray     V2DataReportParametersType = "array"
	V2DataReportParametersTypeEnum      V2DataReportParametersType = "enum"
	V2DataReportParametersTypeString    V2DataReportParametersType = "string"
	V2DataReportParametersTypeTimestamp V2DataReportParametersType = "timestamp"
)

// Details about enum elements in the array.
type V2DataReportParametersArrayDetailsEnumDetails struct {
	// Allowed values of the enum.
	AllowedValues []string `json:"allowed_values"`
}

// For array parameters, provides details about the array elements.
type V2DataReportParametersArrayDetails struct {
	// The data type of the elements in the array.
	ElementType V2DataReportParametersArrayDetailsElementType `json:"element_type"`
	// Details about enum elements in the array.
	EnumDetails *V2DataReportParametersArrayDetailsEnumDetails `json:"enum_details,omitempty"`
}

// For enum parameters, provides the list of allowed values.
type V2DataReportParametersEnumDetails struct {
	// Allowed values of the enum.
	AllowedValues []string `json:"allowed_values"`
}

// Specification of the parameters that the `Report` accepts, keyed by parameter name.
type V2DataReportParameters struct {
	// For array parameters, provides details about the array elements.
	ArrayDetails *V2DataReportParametersArrayDetails `json:"array_details,omitempty"`
	// Explains the purpose and usage of the parameter.
	Description string `json:"description"`
	// For enum parameters, provides the list of allowed values.
	EnumDetails *V2DataReportParametersEnumDetails `json:"enum_details,omitempty"`
	// Indicates whether the parameter must be provided.
	Required bool `json:"required"`
	// The data type of the parameter.
	Type V2DataReportParametersType `json:"type"`
}

// The `Report` resource represents a Stripe-defined, parameterized report that provides
// insights into various aspects of your Stripe integration.
type V2DataReport struct {
	APIResource
	// Representative SQL generated using common parameter values, or an explanatory message when
	// the report's SQL cannot be exposed. Only present when requested via `include[0]=default_sql`.
	DefaultSQL string `json:"default_sql,omitempty"`
	// A human-readable description of what this report contains.
	Description string `json:"description"`
	// The unique identifier of the `Report`.
	ID string `json:"id"`
	// Whether this `Report` is available in live mode.
	Livemode bool `json:"livemode"`
	// The human-readable name of the `Report`.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Specification of the parameters that the `Report` accepts, keyed by parameter name.
	Parameters map[string]*V2DataReportParameters `json:"parameters,omitempty"`
}
