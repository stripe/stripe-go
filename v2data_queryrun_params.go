//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The query to execute.
type V2DataQueryRunQueryParams struct {
	// Ad-hoc SQL to execute.
	SQL *string `form:"sql" json:"sql,omitempty"`
}

// Optional settings that customize the generated result file.
type V2DataQueryRunResultOptionsParams struct {
	// If set, the generated result file is compressed into a ZIP archive before
	// it is stored. This applies only to downloadable file results.
	CompressFile *bool `form:"compress_file" json:"compress_file,omitempty"`
}

// Submits a SQL query for execution against a dataset and returns a `QueryRun` object
// to track progress and retrieve results.
type V2DataQueryRunParams struct {
	Params `form:"*"`
	// The dataset to query.
	Dataset *string `form:"dataset" json:"dataset,omitempty"`
	// The file format for the result.
	Format *string `form:"format" json:"format,omitempty"`
	// Any optional includes (see https://docs.stripe.com/api-includable-response-values).
	Include []*string `form:"include" json:"include,omitempty"`
	// The maximum number of inline `QueryRun` result rows to return. Defaults to 10. Maximum is 1000.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The page token for paginating the inline `QueryRun` result rows.
	Page *string `form:"page" json:"page,omitempty"`
	// The query to execute.
	Query *V2DataQueryRunQueryParams `form:"query" json:"query,omitempty"`
	// Optional settings that customize the generated result file.
	ResultOptions *V2DataQueryRunResultOptionsParams `form:"result_options" json:"result_options,omitempty"`
}

// The query to execute.
type V2DataQueryRunCreateQueryParams struct {
	// Ad-hoc SQL to execute.
	SQL *string `form:"sql" json:"sql,omitempty"`
}

// Optional settings that customize the generated result file.
type V2DataQueryRunCreateResultOptionsParams struct {
	// If set, the generated result file is compressed into a ZIP archive before
	// it is stored. This applies only to downloadable file results.
	CompressFile *bool `form:"compress_file" json:"compress_file,omitempty"`
}

// Submits a SQL query for execution against a dataset and returns a `QueryRun` object
// to track progress and retrieve results.
type V2DataQueryRunCreateParams struct {
	Params `form:"*"`
	// The dataset to query.
	Dataset *string `form:"dataset" json:"dataset"`
	// The file format for the result.
	Format *string `form:"format" json:"format"`
	// The query to execute.
	Query *V2DataQueryRunCreateQueryParams `form:"query" json:"query"`
	// Optional settings that customize the generated result file.
	ResultOptions *V2DataQueryRunCreateResultOptionsParams `form:"result_options" json:"result_options,omitempty"`
}

// Retrieves the status and results of a previously created `QueryRun`.
type V2DataQueryRunRetrieveParams struct {
	Params `form:"*"`
	// Any optional includes (see https://docs.stripe.com/api-includable-response-values).
	Include []*string `form:"include" json:"include,omitempty"`
	// The maximum number of inline `QueryRun` result rows to return. Defaults to 10. Maximum is 1000.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The page token for paginating the inline `QueryRun` result rows.
	Page *string `form:"page" json:"page,omitempty"`
}
