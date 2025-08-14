//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Data type of the elements in the array.
type V2ReportingReportParametersArrayDetailsElementType string

// List of values that V2ReportingReportParametersArrayDetailsElementType can take
const (
	V2ReportingReportParametersArrayDetailsElementTypeEnum V2ReportingReportParametersArrayDetailsElementType = "enum"
)

// The data type of the parameter.
type V2ReportingReportParametersType string

// List of values that V2ReportingReportParametersType can take
const (
	V2ReportingReportParametersTypeArray     V2ReportingReportParametersType = "array"
	V2ReportingReportParametersTypeEnum      V2ReportingReportParametersType = "enum"
	V2ReportingReportParametersTypeString    V2ReportingReportParametersType = "string"
	V2ReportingReportParametersTypeTimestamp V2ReportingReportParametersType = "timestamp"
)

// Details about enum elements in the array.
type V2ReportingReportParametersArrayDetailsEnumDetails struct {
	// Allowed values of the enum.
	AllowedValues []string `json:"allowed_values"`
}

// For array parameters, provides details about the array elements.
type V2ReportingReportParametersArrayDetails struct {
	// Data type of the elements in the array.
	ElementType V2ReportingReportParametersArrayDetailsElementType `json:"element_type"`
	// Details about enum elements in the array.
	EnumDetails *V2ReportingReportParametersArrayDetailsEnumDetails `json:"enum_details"`
}

// For enum parameters, provides the list of allowed values.
type V2ReportingReportParametersEnumDetails struct {
	// Allowed values of the enum.
	AllowedValues []string `json:"allowed_values"`
}

// For timestamp parameters, specifies the allowed date range.
type V2ReportingReportParametersTimestampDetails struct {
	// Maximum permitted timestamp which can be requested.
	Max time.Time `json:"max"`
	// Minimum permitted timestamp which can be requested.
	Min time.Time `json:"min"`
}

// Specification of the parameters that the `Report` accepts. It details each parameter's
// name, description, whether it is required, and any validations performed.
type V2ReportingReportParameters struct {
	// For array parameters, provides details about the array elements.
	ArrayDetails *V2ReportingReportParametersArrayDetails `json:"array_details"`
	// Explains the purpose and usage of the parameter.
	Description string `json:"description"`
	// For enum parameters, provides the list of allowed values.
	EnumDetails *V2ReportingReportParametersEnumDetails `json:"enum_details"`
	// Indicates whether the parameter must be provided.
	Required bool `json:"required"`
	// For timestamp parameters, specifies the allowed date range.
	TimestampDetails *V2ReportingReportParametersTimestampDetails `json:"timestamp_details"`
	// The data type of the parameter.
	Type V2ReportingReportParametersType `json:"type"`
}

// The Report resource represents a customizable report template that provides insights into various aspects of your Stripe integration.
type V2ReportingReport struct {
	APIResource
	// The unique identifier of the `Report` object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The human-readable name of the `Report`.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Specification of the parameters that the `Report` accepts. It details each parameter's
	// name, description, whether it is required, and any validations performed.
	Parameters map[string]*V2ReportingReportParameters `json:"parameters"`
}
