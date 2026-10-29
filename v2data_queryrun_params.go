//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The query to execute.
type V2DataQueryRunQueryParams struct {
	// Ad-hoc SQL to execute. The maximum length is 100,000 characters.
	SQL *string `form:"sql" json:"sql,omitempty"`
}

// Optional settings that customize the generated result file.
type V2DataQueryRunResultOptionsParams struct {
	// The compression applied to the generated result file. Defaults to `none`. This applies
	// only to downloadable file results.
	Compression *string `form:"compression" json:"compression,omitempty"`
}

// Submits a SQL query for execution against a dataset and returns a `QueryRun` object
// to track progress and retrieve results.
type V2DataQueryRunParams struct {
	Params `form:"*"`
	// The dataset to query.
	Dataset *string `form:"dataset" json:"dataset,omitempty"`
	// The file format for the result.
	Format *string `form:"format" json:"format,omitempty"`
	// Any optional includes (see [include-dependent response values](https://docs.stripe.com/api-includable-response-values)).
	Include []*string `form:"include" json:"include,omitempty"`
	// The maximum number of inline `QueryRun` result rows to return. Defaults to 100. Maximum is 1000.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The page token for paginating the inline `QueryRun` result rows.
	Page *string `form:"page" json:"page,omitempty"`
	// Named bind parameters for the query. Keys are parameter names referenced in
	// the SQL (without any prefix). Omit this field for a non-parameterized query.
	// Values are strings, integers, or booleans. Fractional numbers, timestamps,
	// and dates should be passed as strings.
	Parameters map[string]any `form:"parameters" json:"parameters,omitempty"`
	// The query to execute.
	Query *V2DataQueryRunQueryParams `form:"query" json:"query,omitempty"`
	// Optional settings that customize the generated result file.
	ResultOptions *V2DataQueryRunResultOptionsParams `form:"result_options" json:"result_options,omitempty"`
}

// The query to execute.
type V2DataQueryRunCreateQueryParams struct {
	// Ad-hoc SQL to execute. The maximum length is 100,000 characters.
	SQL *string `form:"sql" json:"sql,omitempty"`
}

// Optional settings that customize the generated result file.
type V2DataQueryRunCreateResultOptionsParams struct {
	// The compression applied to the generated result file. Defaults to `none`. This applies
	// only to downloadable file results.
	Compression *string `form:"compression" json:"compression,omitempty"`
}

// Submits a SQL query for execution against a dataset and returns a `QueryRun` object
// to track progress and retrieve results.
type V2DataQueryRunCreateParams struct {
	Params `form:"*"`
	// The dataset to query.
	Dataset *string `form:"dataset" json:"dataset"`
	// The file format for the result.
	Format *string `form:"format" json:"format"`
	// Named bind parameters for the query. Keys are parameter names referenced in
	// the SQL (without any prefix). Omit this field for a non-parameterized query.
	// Values are strings, integers, or booleans. Fractional numbers, timestamps,
	// and dates should be passed as strings.
	Parameters map[string]any `form:"parameters" json:"parameters,omitempty"`
	// The query to execute.
	Query *V2DataQueryRunCreateQueryParams `form:"query" json:"query"`
	// Optional settings that customize the generated result file.
	ResultOptions *V2DataQueryRunCreateResultOptionsParams `form:"result_options" json:"result_options,omitempty"`
}

// Retrieves the status and results of a previously created `QueryRun`. Returns a `404` error
// if the `QueryRun` doesn't exist or its retention period has ended.
type V2DataQueryRunRetrieveParams struct {
	Params `form:"*"`
	// Any optional includes (see [include-dependent response values](https://docs.stripe.com/api-includable-response-values)).
	Include []*string `form:"include" json:"include,omitempty"`
	// The maximum number of inline `QueryRun` result rows to return. Defaults to 100. Maximum is 1000.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The page token for paginating the inline `QueryRun` result rows.
	Page *string `form:"page" json:"page,omitempty"`
}
