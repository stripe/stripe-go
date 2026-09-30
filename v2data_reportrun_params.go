//
//
// File generated from our OpenAPI spec
//
//

package stripe

// A reference to the `Report` to run, by ID or name.
type V2DataReportRunReportParams struct {
	// The unique identifier of the `Report`.
	ID *string `form:"id" json:"id,omitempty"`
	// The human-readable name of the `Report`.
	Name *string `form:"name" json:"name,omitempty"`
}

// Optional settings that customize the generated result file.
type V2DataReportRunResultOptionsParams struct {
	// If set, the generated result file is compressed into a ZIP archive before
	// it is stored. This applies only to downloadable file results.
	CompressFile *bool `form:"compress_file" json:"compress_file,omitempty"`
}

// Initiates the generation of a `ReportRun` based on the specified `Report` and
// caller-provided parameters. Returns a `ReportRun` object which can be used to track
// the progress and retrieve the results of the report.
type V2DataReportRunParams struct {
	Params `form:"*"`
	// The file format for the result.
	Format *string `form:"format" json:"format,omitempty"`
	// Any optional includes (see https://docs.stripe.com/api-includable-response-values).
	Include []*string `form:"include" json:"include,omitempty"`
	// The maximum number of inline `ReportRun` result rows to return. Defaults to 10. Maximum is 1000.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The page token for paginating the inline `ReportRun` result rows.
	Page *string `form:"page" json:"page,omitempty"`
	// A map of parameter names to values, specifying how the report should be customized.
	// The accepted parameters depend on the specific `Report` being run.
	Parameters map[string]any `form:"parameters" json:"parameters,omitempty"`
	// A reference to the `Report` to run, by ID or name.
	Report *V2DataReportRunReportParams `form:"report" json:"report,omitempty"`
	// Optional settings that customize the generated result file.
	ResultOptions *V2DataReportRunResultOptionsParams `form:"result_options" json:"result_options,omitempty"`
}

// A reference to the `Report` to run, by ID or name.
type V2DataReportRunCreateReportParams struct {
	// The unique identifier of the `Report`.
	ID *string `form:"id" json:"id,omitempty"`
	// The human-readable name of the `Report`.
	Name *string `form:"name" json:"name,omitempty"`
}

// Optional settings that customize the generated result file.
type V2DataReportRunCreateResultOptionsParams struct {
	// If set, the generated result file is compressed into a ZIP archive before
	// it is stored. This applies only to downloadable file results.
	CompressFile *bool `form:"compress_file" json:"compress_file,omitempty"`
}

// Initiates the generation of a `ReportRun` based on the specified `Report` and
// caller-provided parameters. Returns a `ReportRun` object which can be used to track
// the progress and retrieve the results of the report.
type V2DataReportRunCreateParams struct {
	Params `form:"*"`
	// The file format for the result.
	Format *string `form:"format" json:"format"`
	// A map of parameter names to values, specifying how the report should be customized.
	// The accepted parameters depend on the specific `Report` being run.
	Parameters map[string]any `form:"parameters" json:"parameters"`
	// A reference to the `Report` to run, by ID or name.
	Report *V2DataReportRunCreateReportParams `form:"report" json:"report"`
	// Optional settings that customize the generated result file.
	ResultOptions *V2DataReportRunCreateResultOptionsParams `form:"result_options" json:"result_options,omitempty"`
}

// Fetches the current state and details of a previously created `ReportRun`. If the
// `ReportRun` has succeeded, the endpoint provides details for how to retrieve the results.
type V2DataReportRunRetrieveParams struct {
	Params `form:"*"`
	// Any optional includes (see https://docs.stripe.com/api-includable-response-values).
	Include []*string `form:"include" json:"include,omitempty"`
	// The maximum number of inline `ReportRun` result rows to return. Defaults to 10. Maximum is 1000.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The page token for paginating the inline `ReportRun` result rows.
	Page *string `form:"page" json:"page,omitempty"`
}
