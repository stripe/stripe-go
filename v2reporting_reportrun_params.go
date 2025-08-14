//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Parameter with an array data type.
type V2ReportingReportRunReportParametersArrayValueParams struct {
	// The list of string values in the array.
	Items []*string `form:"items" json:"items"`
}

// A map of parameter names to values, specifying how the report should be customized.
// The accepted parameters depend on the specific `Report` being run.
type V2ReportingReportRunReportParametersParams struct {
	// Parameter with an array data type.
	ArrayValue *V2ReportingReportRunReportParametersArrayValueParams `form:"array_value" json:"array_value,omitempty"`
	// Parameter with a string data type.
	StringValue *string `form:"string_value" json:"string_value,omitempty"`
	// Parameter with a timestamp data type.
	TimestampValue *time.Time `form:"timestamp_value" json:"timestamp_value,omitempty"`
}

// Optional settings to customize the results of the `ReportRun`.
type V2ReportingReportRunResultOptionsParams struct {
	// If set, the generated report file will be compressed into a ZIP folder.
	// This is useful for reducing file size and download time for large reports.
	CompressFile *bool `form:"compress_file" json:"compress_file,omitempty"`
}

// Initiates the generation of a `ReportRun` based on the specified report template
// and user-provided parameters. It's the starting point for obtaining report data,
// and returns a `ReportRun` object which can be used to track the progress and retrieve
// the results of the report.
type V2ReportingReportRunParams struct {
	Params `form:"*"`
	// The unique identifier of the `Report` being requested.
	Report *string `form:"report" json:"report,omitempty"`
	// A map of parameter names to values, specifying how the report should be customized.
	// The accepted parameters depend on the specific `Report` being run.
	ReportParameters map[string]*V2ReportingReportRunReportParametersParams `form:"report_parameters" json:"report_parameters,omitempty"`
	// Optional settings to customize the results of the `ReportRun`.
	ResultOptions *V2ReportingReportRunResultOptionsParams `form:"result_options" json:"result_options,omitempty"`
}

// Parameter with an array data type.
type V2ReportingReportRunCreateReportParametersArrayValueParams struct {
	// The list of string values in the array.
	Items []*string `form:"items" json:"items"`
}

// A map of parameter names to values, specifying how the report should be customized.
// The accepted parameters depend on the specific `Report` being run.
type V2ReportingReportRunCreateReportParametersParams struct {
	// Parameter with an array data type.
	ArrayValue *V2ReportingReportRunCreateReportParametersArrayValueParams `form:"array_value" json:"array_value,omitempty"`
	// Parameter with a string data type.
	StringValue *string `form:"string_value" json:"string_value,omitempty"`
	// Parameter with a timestamp data type.
	TimestampValue *time.Time `form:"timestamp_value" json:"timestamp_value,omitempty"`
}

// Optional settings to customize the results of the `ReportRun`.
type V2ReportingReportRunCreateResultOptionsParams struct {
	// If set, the generated report file will be compressed into a ZIP folder.
	// This is useful for reducing file size and download time for large reports.
	CompressFile *bool `form:"compress_file" json:"compress_file,omitempty"`
}

// Initiates the generation of a `ReportRun` based on the specified report template
// and user-provided parameters. It's the starting point for obtaining report data,
// and returns a `ReportRun` object which can be used to track the progress and retrieve
// the results of the report.
type V2ReportingReportRunCreateParams struct {
	Params `form:"*"`
	// The unique identifier of the `Report` being requested.
	Report *string `form:"report" json:"report"`
	// A map of parameter names to values, specifying how the report should be customized.
	// The accepted parameters depend on the specific `Report` being run.
	ReportParameters map[string]*V2ReportingReportRunCreateReportParametersParams `form:"report_parameters" json:"report_parameters"`
	// Optional settings to customize the results of the `ReportRun`.
	ResultOptions *V2ReportingReportRunCreateResultOptionsParams `form:"result_options" json:"result_options,omitempty"`
}

// Fetches the current state and details of a previously created `ReportRun`. If the `ReportRun`
// has succeeded, the endpoint will provide details for how to retrieve the results.
type V2ReportingReportRunRetrieveParams struct {
	Params `form:"*"`
}
