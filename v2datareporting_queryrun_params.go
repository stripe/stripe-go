//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Optional settings to customize the results of the `QueryRun`.
type V2DataReportingQueryRunResultOptionsParams struct {
	// If set, the generated results file will be compressed into a ZIP folder.
	// This is useful for reducing file size and download time for large results.
	CompressFile *bool `form:"compress_file" json:"compress_file,omitempty"`
}

// Creates a query run to execute ad-hoc SQL and returns a `QueryRun` object to track progress and retrieve results.
type V2DataReportingQueryRunParams struct {
	Params `form:"*"`
	// Optional settings to customize the results of the `QueryRun`.
	ResultOptions *V2DataReportingQueryRunResultOptionsParams `form:"result_options" json:"result_options,omitempty"`
	// The SQL to execute.
	SQL *string `form:"sql" json:"sql,omitempty"`
}

// Optional settings to customize the results of the `QueryRun`.
type V2DataReportingQueryRunCreateResultOptionsParams struct {
	// If set, the generated results file will be compressed into a ZIP folder.
	// This is useful for reducing file size and download time for large results.
	CompressFile *bool `form:"compress_file" json:"compress_file,omitempty"`
}

// Creates a query run to execute ad-hoc SQL and returns a `QueryRun` object to track progress and retrieve results.
type V2DataReportingQueryRunCreateParams struct {
	Params `form:"*"`
	// Optional settings to customize the results of the `QueryRun`.
	ResultOptions *V2DataReportingQueryRunCreateResultOptionsParams `form:"result_options" json:"result_options,omitempty"`
	// The SQL to execute.
	SQL *string `form:"sql" json:"sql"`
}

// Fetches the current state and details of a previously created `QueryRun`. If the `QueryRun`
// has succeeded, the endpoint will provide details for how to retrieve the results.
type V2DataReportingQueryRunRetrieveParams struct {
	Params `form:"*"`
}
