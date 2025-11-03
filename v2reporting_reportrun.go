//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The content type of the file.
type V2ReportingReportRunResultFileContentType string

// List of values that V2ReportingReportRunResultFileContentType can take
const (
	V2ReportingReportRunResultFileContentTypeCsv V2ReportingReportRunResultFileContentType = "csv"
	V2ReportingReportRunResultFileContentTypeZip V2ReportingReportRunResultFileContentType = "zip"
)

// The type of the `ReportRun` result.
type V2ReportingReportRunResultType string

// List of values that V2ReportingReportRunResultType can take
const (
	V2ReportingReportRunResultTypeFile V2ReportingReportRunResultType = "file"
)

// The current status of the `ReportRun`.
type V2ReportingReportRunStatus string

// List of values that V2ReportingReportRunStatus can take
const (
	V2ReportingReportRunStatusFailed    V2ReportingReportRunStatus = "failed"
	V2ReportingReportRunStatusRunning   V2ReportingReportRunStatus = "running"
	V2ReportingReportRunStatusSucceeded V2ReportingReportRunStatus = "succeeded"
)

// Error code categorizing the reason the `ReportRun` failed.
type V2ReportingReportRunStatusDetailsErrorCode string

// List of values that V2ReportingReportRunStatusDetailsErrorCode can take
const (
	V2ReportingReportRunStatusDetailsErrorCodeFileSizeAboveLimit V2ReportingReportRunStatusDetailsErrorCode = "file_size_above_limit"
	V2ReportingReportRunStatusDetailsErrorCodeInternalError      V2ReportingReportRunStatusDetailsErrorCode = "internal_error"
)

// A pre-signed URL that allows secure, time-limited access to download the file.
type V2ReportingReportRunResultFileDownloadURL struct {
	// The time that the URL expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// The URL that can be used for accessing the file.
	URL string `json:"url"`
}

// Contains metadata about the file produced by the `ReportRun`, including
// its content type, size, and a URL to download its contents.
type V2ReportingReportRunResultFile struct {
	// The content type of the file.
	ContentType V2ReportingReportRunResultFileContentType `json:"content_type"`
	// A pre-signed URL that allows secure, time-limited access to download the file.
	DownloadURL *V2ReportingReportRunResultFileDownloadURL `json:"download_url"`
	// The total size of the file in bytes.
	Size int64 `json:"size"`
}

// Details how to retrieve the results of a successfully completed `ReportRun`.
type V2ReportingReportRunResult struct {
	// Contains metadata about the file produced by the `ReportRun`, including
	// its content type, size, and a URL to download its contents.
	File *V2ReportingReportRunResultFile `json:"file"`
	// The type of the `ReportRun` result.
	Type V2ReportingReportRunResultType `json:"type"`
}

// The options specified for customizing the output file of the `ReportRun`.
type V2ReportingReportRunResultOptions struct {
	// If set, the generated report file will be compressed into a ZIP folder.
	// This is useful for reducing file size and download time for large reports.
	CompressFile bool `json:"compress_file,omitempty"`
}

// Additional details about the current state of the `ReportRun`. The field is currently only populated when a `ReportRun`
// is in the `failed` state, providing more information about why the report failed to generate successfully.
type V2ReportingReportRunStatusDetails struct {
	// Error code categorizing the reason the `ReportRun` failed.
	ErrorCode V2ReportingReportRunStatusDetailsErrorCode `json:"error_code,omitempty"`
	// Error message with additional details about the failure.
	ErrorMessage string `json:"error_message,omitempty"`
}

// The `ReportRun` object represents an instance of a `Report` generated with specific
// run parameters. Once the object is created, Stripe begins processing the report. When
// the report has finished running, it will give you a reference to the results.
type V2ReportingReportRun struct {
	APIResource
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// The unique identifier of the `ReportRun` object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The unique identifier of the `Report` object which was run.
	Report string `json:"report"`
	// The human-readable name of the `Report` which was run.
	ReportName string `json:"report_name"`
	// The parameters used to customize the generation of the report.
	ReportParameters map[string]any `json:"report_parameters"`
	// Details how to retrieve the results of a successfully completed `ReportRun`.
	Result *V2ReportingReportRunResult `json:"result,omitempty"`
	// The options specified for customizing the output file of the `ReportRun`.
	ResultOptions *V2ReportingReportRunResultOptions `json:"result_options,omitempty"`
	// The current status of the `ReportRun`.
	Status V2ReportingReportRunStatus `json:"status"`
	// Additional details about the current state of the `ReportRun`. The field is currently only populated when a `ReportRun`
	// is in the `failed` state, providing more information about why the report failed to generate successfully.
	StatusDetails map[string]*V2ReportingReportRunStatusDetails `json:"status_details"`
}
