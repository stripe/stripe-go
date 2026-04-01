//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Error code categorizing the reason the `QueryRun` failed.
type V2DataReportingQueryRunStatusDetailsErrorCode string

// List of values that V2DataReportingQueryRunStatusDetailsErrorCode can take
const (
	V2DataReportingQueryRunStatusDetailsErrorCodeFileSizeAboveLimit V2DataReportingQueryRunStatusDetailsErrorCode = "file_size_above_limit"
	V2DataReportingQueryRunStatusDetailsErrorCodeInternalError      V2DataReportingQueryRunStatusDetailsErrorCode = "internal_error"
)

// The type of the `ReportRun` or `QueryRun` result.
type V2DataReportingQueryRunResultType string

// List of values that V2DataReportingQueryRunResultType can take
const (
	V2DataReportingQueryRunResultTypeFile V2DataReportingQueryRunResultType = "file"
)

// The content type of the file.
type V2DataReportingQueryRunResultFileContentType string

// List of values that V2DataReportingQueryRunResultFileContentType can take
const (
	V2DataReportingQueryRunResultFileContentTypeCsv V2DataReportingQueryRunResultFileContentType = "csv"
	V2DataReportingQueryRunResultFileContentTypeZip V2DataReportingQueryRunResultFileContentType = "zip"
)

// The current status of the `QueryRun`.
type V2DataReportingQueryRunStatus string

// List of values that V2DataReportingQueryRunStatus can take
const (
	V2DataReportingQueryRunStatusFailed    V2DataReportingQueryRunStatus = "failed"
	V2DataReportingQueryRunStatusRunning   V2DataReportingQueryRunStatus = "running"
	V2DataReportingQueryRunStatusSucceeded V2DataReportingQueryRunStatus = "succeeded"
)

// Additional details about the current state of the `QueryRun`. Populated when the `QueryRun`
// is in the `failed` state, providing more information about why the query failed.
type V2DataReportingQueryRunStatusDetails struct {
	// Error code categorizing the reason the `QueryRun` failed.
	ErrorCode V2DataReportingQueryRunStatusDetailsErrorCode `json:"error_code,omitempty"`
	// Error message with additional details about the failure.
	ErrorMessage string `json:"error_message,omitempty"`
}

// A pre-signed URL that allows secure, time-limited access to download the file.
type V2DataReportingQueryRunResultFileDownloadURL struct {
	// The time that the URL expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// The URL that can be used for accessing the file.
	URL string `json:"url"`
}

// Contains metadata about the file produced by the `ReportRun` or `QueryRun`, including
// its content type, size, and a URL to download its contents.
type V2DataReportingQueryRunResultFile struct {
	// The content type of the file.
	ContentType V2DataReportingQueryRunResultFileContentType `json:"content_type"`
	// A pre-signed URL that allows secure, time-limited access to download the file.
	DownloadURL *V2DataReportingQueryRunResultFileDownloadURL `json:"download_url"`
	// The total size of the file in bytes.
	Size int64 `json:"size,string"`
}

// Details how to retrieve the results of a successfully completed `QueryRun`.
type V2DataReportingQueryRunResult struct {
	// Contains metadata about the file produced by the `ReportRun` or `QueryRun`, including
	// its content type, size, and a URL to download its contents.
	File *V2DataReportingQueryRunResultFile `json:"file,omitempty"`
	// The type of the `ReportRun` or `QueryRun` result.
	Type V2DataReportingQueryRunResultType `json:"type"`
}

// The options specified for customizing the output of the `QueryRun`.
type V2DataReportingQueryRunResultOptions struct {
	// If set, the generated results file will be compressed into a ZIP folder.
	// This is useful for reducing file size and download time for large results.
	CompressFile bool `json:"compress_file,omitempty"`
}

// The `QueryRun` object represents an ad-hoc SQL execution. Once created, Stripe processes the query. When
// the query has finished running, the object provides a reference to the results.
type V2DataReportingQueryRun struct {
	APIResource
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// The unique identifier of the `QueryRun` object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Details how to retrieve the results of a successfully completed `QueryRun`.
	Result *V2DataReportingQueryRunResult `json:"result,omitempty"`
	// The options specified for customizing the output of the `QueryRun`.
	ResultOptions *V2DataReportingQueryRunResultOptions `json:"result_options,omitempty"`
	// The SQL that was executed.
	SQL string `json:"sql"`
	// The current status of the `QueryRun`.
	Status V2DataReportingQueryRunStatus `json:"status"`
	// Additional details about the current state of the `QueryRun`. Populated when the `QueryRun`
	// is in the `failed` state, providing more information about why the query failed.
	StatusDetails map[string]*V2DataReportingQueryRunStatusDetails `json:"status_details"`
}
