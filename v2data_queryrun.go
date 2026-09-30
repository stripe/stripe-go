//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The dataset that was queried.
type V2DataQueryRunDataset string

// List of values that V2DataQueryRunDataset can take
const (
	V2DataQueryRunDatasetAnalytical V2DataQueryRunDataset = "analytical"
)

// The file format of the result. Only applicable when the result is a file.
type V2DataQueryRunFormat string

// List of values that V2DataQueryRunFormat can take
const (
	V2DataQueryRunFormatCsv V2DataQueryRunFormat = "csv"
)

// The data type of the column.
type V2DataQueryRunResultFileColumnType string

// List of values that V2DataQueryRunResultFileColumnType can take
const (
	V2DataQueryRunResultFileColumnTypeBigint    V2DataQueryRunResultFileColumnType = "bigint"
	V2DataQueryRunResultFileColumnTypeBoolean   V2DataQueryRunResultFileColumnType = "boolean"
	V2DataQueryRunResultFileColumnTypeDate      V2DataQueryRunResultFileColumnType = "date"
	V2DataQueryRunResultFileColumnTypeDatetime  V2DataQueryRunResultFileColumnType = "datetime"
	V2DataQueryRunResultFileColumnTypeDecimal   V2DataQueryRunResultFileColumnType = "decimal"
	V2DataQueryRunResultFileColumnTypeDouble    V2DataQueryRunResultFileColumnType = "double"
	V2DataQueryRunResultFileColumnTypeInteger   V2DataQueryRunResultFileColumnType = "integer"
	V2DataQueryRunResultFileColumnTypeTimestamp V2DataQueryRunResultFileColumnType = "timestamp"
	V2DataQueryRunResultFileColumnTypeVarchar   V2DataQueryRunResultFileColumnType = "varchar"
)

// The content type of the file.
type V2DataQueryRunResultFileContentType string

// List of values that V2DataQueryRunResultFileContentType can take
const (
	V2DataQueryRunResultFileContentTypeCsv V2DataQueryRunResultFileContentType = "csv"
)

// The data type of the column.
type V2DataQueryRunResultInlineColumnType string

// List of values that V2DataQueryRunResultInlineColumnType can take
const (
	V2DataQueryRunResultInlineColumnTypeBigint    V2DataQueryRunResultInlineColumnType = "bigint"
	V2DataQueryRunResultInlineColumnTypeBoolean   V2DataQueryRunResultInlineColumnType = "boolean"
	V2DataQueryRunResultInlineColumnTypeDate      V2DataQueryRunResultInlineColumnType = "date"
	V2DataQueryRunResultInlineColumnTypeDatetime  V2DataQueryRunResultInlineColumnType = "datetime"
	V2DataQueryRunResultInlineColumnTypeDecimal   V2DataQueryRunResultInlineColumnType = "decimal"
	V2DataQueryRunResultInlineColumnTypeDouble    V2DataQueryRunResultInlineColumnType = "double"
	V2DataQueryRunResultInlineColumnTypeInteger   V2DataQueryRunResultInlineColumnType = "integer"
	V2DataQueryRunResultInlineColumnTypeTimestamp V2DataQueryRunResultInlineColumnType = "timestamp"
	V2DataQueryRunResultInlineColumnTypeVarchar   V2DataQueryRunResultInlineColumnType = "varchar"
)

// The current status of the `QueryRun`.
type V2DataQueryRunStatus string

// List of values that V2DataQueryRunStatus can take
const (
	V2DataQueryRunStatusCanceled  V2DataQueryRunStatus = "canceled"
	V2DataQueryRunStatusFailed    V2DataQueryRunStatus = "failed"
	V2DataQueryRunStatusRunning   V2DataQueryRunStatus = "running"
	V2DataQueryRunStatusSucceeded V2DataQueryRunStatus = "succeeded"
)

// Error code categorizing the reason the run failed.
type V2DataQueryRunStatusDetailsCode string

// List of values that V2DataQueryRunStatusDetailsCode can take
const (
	V2DataQueryRunStatusDetailsCodeFileSizeAboveLimit V2DataQueryRunStatusDetailsCode = "file_size_above_limit"
	V2DataQueryRunStatusDetailsCodeInternalError      V2DataQueryRunStatusDetailsCode = "internal_error"
	V2DataQueryRunStatusDetailsCodeQueryRunInvalidSQL V2DataQueryRunStatusDetailsCode = "query_run_invalid_sql"
)

// The query that was submitted for execution.
type V2DataQueryRunQuery struct {
	// Ad-hoc SQL to execute.
	SQL string `json:"sql,omitempty"`
}

// The schema of the result data.
type V2DataQueryRunResultFileColumn struct {
	// The name of the column.
	Name string `json:"name"`
	// The data type of the column.
	Type V2DataQueryRunResultFileColumnType `json:"type"`
}

// A pre-signed URL that allows secure, time-limited access to download the file.
type V2DataQueryRunResultFileDownloadURL struct {
	// The time that the URL expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// The URL that can be used for accessing the file.
	URL string `json:"url"`
}

// File result with a download URL. This is the default result type.
type V2DataQueryRunResultFile struct {
	// The schema of the result data.
	Columns []*V2DataQueryRunResultFileColumn `json:"columns"`
	// The content type of the file.
	ContentType V2DataQueryRunResultFileContentType `json:"content_type"`
	// A pre-signed URL that allows secure, time-limited access to download the file.
	DownloadURL *V2DataQueryRunResultFileDownloadURL `json:"download_url"`
	// The total size of the file in bytes.
	Size int64 `json:"size,string"`
}

// The schema of the result data.
type V2DataQueryRunResultInlineColumn struct {
	// The name of the column.
	Name string `json:"name"`
	// The data type of the column.
	Type V2DataQueryRunResultInlineColumnType `json:"type"`
}

// The result rows, each represented as a map of column name to value.
type V2DataQueryRunResultInlineRow struct {
	// The column data in this row, keyed by column name.
	Data map[string]any `json:"data"`
}

// Inline result with data returned directly. Only present when requested via
// `include[0]=result.inline`.
type V2DataQueryRunResultInline struct {
	// The schema of the result data.
	Columns []*V2DataQueryRunResultInlineColumn `json:"columns"`
	// Token for the next page of rows.
	NextPageURL string `json:"next_page_url,omitempty"`
	// Token for the previous page of rows.
	PreviousPageURL string `json:"previous_page_url,omitempty"`
	// The result rows, each represented as a map of column name to value.
	Rows []*V2DataQueryRunResultInlineRow `json:"rows"`
}

// The result of the `QueryRun`, populated when it has completed.
type V2DataQueryRunResult struct {
	// The total number of columns in the result.
	ColCount int64 `json:"col_count,string,omitempty"`
	// File result with a download URL. This is the default result type.
	File *V2DataQueryRunResultFile `json:"file,omitempty"`
	// Inline result with data returned directly. Only present when requested via
	// `include[0]=result.inline`.
	Inline *V2DataQueryRunResultInline `json:"inline,omitempty"`
	// The total number of data rows in the result, excluding any header row.
	RowCount int64 `json:"row_count,string,omitempty"`
}

// Settings applied to the generated result file.
type V2DataQueryRunResultOptions struct {
	// If set, the generated result file is compressed into a ZIP archive before
	// it is stored. This applies only to downloadable file results.
	CompressFile bool `json:"compress_file,omitempty"`
}

// Additional details about the current state of the `QueryRun`.
type V2DataQueryRunStatusDetails struct {
	// Time at which the run was canceled. Populated when the run is in the `canceled` state.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Error code categorizing the reason the run failed.
	Code V2DataQueryRunStatusDetailsCode `json:"code,omitempty"`
	// Error message with additional details about the failure.
	Message string `json:"message,omitempty"`
}

// The `QueryRun` resource represents an execution of ad-hoc SQL against a dataset. Once
// created, Stripe processes the query. When the query has finished running, the object
// provides a reference to the results.
type V2DataQueryRun struct {
	APIResource
	// Time at which the `QueryRun` was created.
	Created time.Time `json:"created"`
	// The dataset that was queried.
	Dataset V2DataQueryRunDataset `json:"dataset"`
	// The file format of the result. Only applicable when the result is a file.
	Format V2DataQueryRunFormat `json:"format"`
	// The unique identifier of the `QueryRun`.
	ID string `json:"id"`
	// Whether the `QueryRun` was executed in live mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The query that was submitted for execution.
	Query *V2DataQueryRunQuery `json:"query"`
	// Time at which the data used by this query was last refreshed.
	RefreshedAt time.Time `json:"refreshed_at,omitempty"`
	// The result of the `QueryRun`, populated when it has completed.
	Result *V2DataQueryRunResult `json:"result,omitempty"`
	// Settings applied to the generated result file.
	ResultOptions *V2DataQueryRunResultOptions `json:"result_options,omitempty"`
	// The current status of the `QueryRun`.
	Status V2DataQueryRunStatus `json:"status"`
	// Additional details about the current state of the `QueryRun`.
	StatusDetails *V2DataQueryRunStatusDetails `json:"status_details,omitempty"`
}
