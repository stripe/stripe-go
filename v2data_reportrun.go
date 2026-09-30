//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The data type of the column.
type V2DataReportRunResultFileColumnType string

// List of values that V2DataReportRunResultFileColumnType can take
const (
	V2DataReportRunResultFileColumnTypeBigint    V2DataReportRunResultFileColumnType = "bigint"
	V2DataReportRunResultFileColumnTypeBoolean   V2DataReportRunResultFileColumnType = "boolean"
	V2DataReportRunResultFileColumnTypeDate      V2DataReportRunResultFileColumnType = "date"
	V2DataReportRunResultFileColumnTypeDatetime  V2DataReportRunResultFileColumnType = "datetime"
	V2DataReportRunResultFileColumnTypeDecimal   V2DataReportRunResultFileColumnType = "decimal"
	V2DataReportRunResultFileColumnTypeDouble    V2DataReportRunResultFileColumnType = "double"
	V2DataReportRunResultFileColumnTypeInteger   V2DataReportRunResultFileColumnType = "integer"
	V2DataReportRunResultFileColumnTypeTimestamp V2DataReportRunResultFileColumnType = "timestamp"
	V2DataReportRunResultFileColumnTypeVarchar   V2DataReportRunResultFileColumnType = "varchar"
)

// The content type of the file.
type V2DataReportRunResultFileContentType string

// List of values that V2DataReportRunResultFileContentType can take
const (
	V2DataReportRunResultFileContentTypeCsv V2DataReportRunResultFileContentType = "csv"
)

// The data type of the column.
type V2DataReportRunResultInlineColumnType string

// List of values that V2DataReportRunResultInlineColumnType can take
const (
	V2DataReportRunResultInlineColumnTypeBigint    V2DataReportRunResultInlineColumnType = "bigint"
	V2DataReportRunResultInlineColumnTypeBoolean   V2DataReportRunResultInlineColumnType = "boolean"
	V2DataReportRunResultInlineColumnTypeDate      V2DataReportRunResultInlineColumnType = "date"
	V2DataReportRunResultInlineColumnTypeDatetime  V2DataReportRunResultInlineColumnType = "datetime"
	V2DataReportRunResultInlineColumnTypeDecimal   V2DataReportRunResultInlineColumnType = "decimal"
	V2DataReportRunResultInlineColumnTypeDouble    V2DataReportRunResultInlineColumnType = "double"
	V2DataReportRunResultInlineColumnTypeInteger   V2DataReportRunResultInlineColumnType = "integer"
	V2DataReportRunResultInlineColumnTypeTimestamp V2DataReportRunResultInlineColumnType = "timestamp"
	V2DataReportRunResultInlineColumnTypeVarchar   V2DataReportRunResultInlineColumnType = "varchar"
)

// The current status of the `ReportRun`.
type V2DataReportRunStatus string

// List of values that V2DataReportRunStatus can take
const (
	V2DataReportRunStatusCanceled  V2DataReportRunStatus = "canceled"
	V2DataReportRunStatusFailed    V2DataReportRunStatus = "failed"
	V2DataReportRunStatusRunning   V2DataReportRunStatus = "running"
	V2DataReportRunStatusSucceeded V2DataReportRunStatus = "succeeded"
)

// Error code categorizing the reason the run failed.
type V2DataReportRunStatusDetailsCode string

// List of values that V2DataReportRunStatusDetailsCode can take
const (
	V2DataReportRunStatusDetailsCodeFileSizeAboveLimit V2DataReportRunStatusDetailsCode = "file_size_above_limit"
	V2DataReportRunStatusDetailsCodeInternalError      V2DataReportRunStatusDetailsCode = "internal_error"
	V2DataReportRunStatusDetailsCodeQueryRunInvalidSQL V2DataReportRunStatusDetailsCode = "query_run_invalid_sql"
)

// The schema of the result data.
type V2DataReportRunResultFileColumn struct {
	// The name of the column.
	Name string `json:"name"`
	// The data type of the column.
	Type V2DataReportRunResultFileColumnType `json:"type"`
}

// A pre-signed URL that allows secure, time-limited access to download the file.
type V2DataReportRunResultFileDownloadURL struct {
	// The time that the URL expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// The URL that can be used for accessing the file.
	URL string `json:"url"`
}

// File result with a download URL. This is the default result type.
type V2DataReportRunResultFile struct {
	// The schema of the result data.
	Columns []*V2DataReportRunResultFileColumn `json:"columns"`
	// The content type of the file.
	ContentType V2DataReportRunResultFileContentType `json:"content_type"`
	// A pre-signed URL that allows secure, time-limited access to download the file.
	DownloadURL *V2DataReportRunResultFileDownloadURL `json:"download_url"`
	// The total size of the file in bytes.
	Size int64 `json:"size,string"`
}

// The schema of the result data.
type V2DataReportRunResultInlineColumn struct {
	// The name of the column.
	Name string `json:"name"`
	// The data type of the column.
	Type V2DataReportRunResultInlineColumnType `json:"type"`
}

// The result rows, each represented as a map of column name to value.
type V2DataReportRunResultInlineRow struct {
	// The column data in this row, keyed by column name.
	Data map[string]any `json:"data"`
}

// Inline result with data returned directly. Only present when requested via
// `include[0]=result.inline`.
type V2DataReportRunResultInline struct {
	// The schema of the result data.
	Columns []*V2DataReportRunResultInlineColumn `json:"columns"`
	// Token for the next page of rows.
	NextPageURL string `json:"next_page_url,omitempty"`
	// Token for the previous page of rows.
	PreviousPageURL string `json:"previous_page_url,omitempty"`
	// The result rows, each represented as a map of column name to value.
	Rows []*V2DataReportRunResultInlineRow `json:"rows"`
}

// The result of the `ReportRun`, populated when it has completed.
type V2DataReportRunResult struct {
	// The total number of columns in the result.
	ColCount int64 `json:"col_count,string,omitempty"`
	// File result with a download URL. This is the default result type.
	File *V2DataReportRunResultFile `json:"file,omitempty"`
	// Inline result with data returned directly. Only present when requested via
	// `include[0]=result.inline`.
	Inline *V2DataReportRunResultInline `json:"inline,omitempty"`
	// The total number of data rows in the result, excluding any header row.
	RowCount int64 `json:"row_count,string,omitempty"`
}

// Settings applied to the generated result file.
type V2DataReportRunResultOptions struct {
	// If set, the generated result file is compressed into a ZIP archive before
	// it is stored. This applies only to downloadable file results.
	CompressFile bool `json:"compress_file,omitempty"`
}

// Additional details about the current state of the `ReportRun`.
type V2DataReportRunStatusDetails struct {
	// Time at which the run was canceled. Populated when the run is in the `canceled` state.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Error code categorizing the reason the run failed.
	Code V2DataReportRunStatusDetailsCode `json:"code,omitempty"`
	// Error message with additional details about the failure.
	Message string `json:"message,omitempty"`
}

// The `ReportRun` resource represents an instance of a `Report` generated with specific
// parameter values. Once the object is created, Stripe begins processing the report. When
// the report has finished running, it provides a reference to the results.
type V2DataReportRun struct {
	APIResource
	// Time at which the `ReportRun` was created.
	Created time.Time `json:"created"`
	// The unique identifier of the `ReportRun`.
	ID string `json:"id"`
	// Whether the `ReportRun` was executed in live mode.
	Livemode bool `json:"livemode"`
	// The human-readable name of the `Report` which was run.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The parameters used to customize the generation of the report.
	Parameters map[string]any `json:"parameters"`
	// Time at which the data used by this report was last refreshed.
	RefreshedAt time.Time `json:"refreshed_at,omitempty"`
	// The unique identifier of the `Report` which was run.
	Report string `json:"report"`
	// The result of the `ReportRun`, populated when it has completed.
	Result *V2DataReportRunResult `json:"result,omitempty"`
	// Settings applied to the generated result file.
	ResultOptions *V2DataReportRunResultOptions `json:"result_options,omitempty"`
	// The fully-resolved SQL that was executed. Only present when requested via
	// `include[0]=sql`.
	SQL string `json:"sql,omitempty"`
	// The current status of the `ReportRun`.
	Status V2DataReportRunStatus `json:"status"`
	// Additional details about the current state of the `ReportRun`.
	StatusDetails *V2DataReportRunStatusDetails `json:"status_details,omitempty"`
}
