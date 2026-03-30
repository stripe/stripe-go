//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The current status of the `batch_job`.
type V2CoreBatchJobStatus string

// List of values that V2CoreBatchJobStatus can take
const (
	V2CoreBatchJobStatusBatchFailed      V2CoreBatchJobStatus = "batch_failed"
	V2CoreBatchJobStatusCanceled         V2CoreBatchJobStatus = "canceled"
	V2CoreBatchJobStatusCancelling       V2CoreBatchJobStatus = "cancelling"
	V2CoreBatchJobStatusComplete         V2CoreBatchJobStatus = "complete"
	V2CoreBatchJobStatusInProgress       V2CoreBatchJobStatus = "in_progress"
	V2CoreBatchJobStatusReadyForUpload   V2CoreBatchJobStatus = "ready_for_upload"
	V2CoreBatchJobStatusTimeout          V2CoreBatchJobStatus = "timeout"
	V2CoreBatchJobStatusUploadTimeout    V2CoreBatchJobStatus = "upload_timeout"
	V2CoreBatchJobStatusValidating       V2CoreBatchJobStatus = "validating"
	V2CoreBatchJobStatusValidationFailed V2CoreBatchJobStatus = "validation_failed"
)

// Additional details for the `BATCH_FAILED` status of the `batch_job`.
type V2CoreBatchJobStatusDetailsBatchFailed struct {
	// Details about the `batch_job` failure.
	Error string `json:"error"`
}

// A pre-signed URL that allows secure, time-limited access to download the file.
type V2CoreBatchJobStatusDetailsCanceledOutputFileDownloadURL struct {
	// The time that the URL expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// The URL that can be used for accessing the file.
	URL string `json:"url"`
}

// The output file details. If the `batch_job` is canceled, this is provided only if there is already output at this point.
type V2CoreBatchJobStatusDetailsCanceledOutputFile struct {
	// The content type of the file.
	ContentType string `json:"content_type"`
	// A pre-signed URL that allows secure, time-limited access to download the file.
	DownloadURL *V2CoreBatchJobStatusDetailsCanceledOutputFileDownloadURL `json:"download_url"`
	// The total size of the file in bytes.
	Size int64 `json:"size,string"`
}

// Additional details for the `CANCELED` status of the `batch_job`.
type V2CoreBatchJobStatusDetailsCanceled struct {
	// The total number of records that failed processing.
	FailureCount int64 `json:"failure_count,string"`
	// The output file details. If the `batch_job` is canceled, this is provided only if there is already output at this point.
	OutputFile *V2CoreBatchJobStatusDetailsCanceledOutputFile `json:"output_file"`
	// The total number of records that were successfully processed.
	SuccessCount int64 `json:"success_count,string"`
}

// A pre-signed URL that allows secure, time-limited access to download the file.
type V2CoreBatchJobStatusDetailsCompleteOutputFileDownloadURL struct {
	// The time that the URL expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// The URL that can be used for accessing the file.
	URL string `json:"url"`
}

// The output file details. If the `batch_job` is canceled, this is provided only if there is already output at this point.
type V2CoreBatchJobStatusDetailsCompleteOutputFile struct {
	// The content type of the file.
	ContentType string `json:"content_type"`
	// A pre-signed URL that allows secure, time-limited access to download the file.
	DownloadURL *V2CoreBatchJobStatusDetailsCompleteOutputFileDownloadURL `json:"download_url"`
	// The total size of the file in bytes.
	Size int64 `json:"size,string"`
}

// Additional details for the `COMPLETE` status of the `batch_job`.
type V2CoreBatchJobStatusDetailsComplete struct {
	// The total number of records that failed processing.
	FailureCount int64 `json:"failure_count,string"`
	// The output file details. If the `batch_job` is canceled, this is provided only if there is already output at this point.
	OutputFile *V2CoreBatchJobStatusDetailsCompleteOutputFile `json:"output_file"`
	// The total number of records that were successfully processed.
	SuccessCount int64 `json:"success_count,string"`
}

// Additional details for the `IN_PROGRESS` status of the `batch_job`.
type V2CoreBatchJobStatusDetailsInProgress struct {
	// The number of records that failed processing so far.
	FailureCount int64 `json:"failure_count,string"`
	// The number of records that were successfully processed so far.
	SuccessCount int64 `json:"success_count,string"`
}

// The upload file details.
type V2CoreBatchJobStatusDetailsReadyForUploadUploadURL struct {
	// The time that the URL expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// The URL that can be used for accessing the file.
	URL string `json:"url"`
}

// Additional details for the `READY_FOR_UPLOAD` status of the `batch_job`.
type V2CoreBatchJobStatusDetailsReadyForUpload struct {
	// The upload file details.
	UploadURL *V2CoreBatchJobStatusDetailsReadyForUploadUploadURL `json:"upload_url"`
}

// A pre-signed URL that allows secure, time-limited access to download the file.
type V2CoreBatchJobStatusDetailsTimeoutOutputFileDownloadURL struct {
	// The time that the URL expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// The URL that can be used for accessing the file.
	URL string `json:"url"`
}

// The output file details. If the `batch_job` is canceled, this is provided only if there is already output at this point.
type V2CoreBatchJobStatusDetailsTimeoutOutputFile struct {
	// The content type of the file.
	ContentType string `json:"content_type"`
	// A pre-signed URL that allows secure, time-limited access to download the file.
	DownloadURL *V2CoreBatchJobStatusDetailsTimeoutOutputFileDownloadURL `json:"download_url"`
	// The total size of the file in bytes.
	Size int64 `json:"size,string"`
}

// Additional details for the `TIMEOUT` status of the `batch_job`.
type V2CoreBatchJobStatusDetailsTimeout struct {
	// The total number of records that failed processing.
	FailureCount int64 `json:"failure_count,string"`
	// The output file details. If the `batch_job` is canceled, this is provided only if there is already output at this point.
	OutputFile *V2CoreBatchJobStatusDetailsTimeoutOutputFile `json:"output_file"`
	// The total number of records that were successfully processed.
	SuccessCount int64 `json:"success_count,string"`
}

// Additional details for the `VALIDATING` status of the `batch_job`.
type V2CoreBatchJobStatusDetailsValidating struct {
	// The number of records that were validated. Note that there is no failure counter here;
	// once we have any validation failures we give up.
	ValidatedCount int64 `json:"validated_count,string"`
}

// A pre-signed URL that allows secure, time-limited access to download the file.
type V2CoreBatchJobStatusDetailsValidationFailedOutputFileDownloadURL struct {
	// The time that the URL expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// The URL that can be used for accessing the file.
	URL string `json:"url"`
}

// The output file details. If the `batch_job` is canceled, this is provided only if there is already output at this point.
type V2CoreBatchJobStatusDetailsValidationFailedOutputFile struct {
	// The content type of the file.
	ContentType string `json:"content_type"`
	// A pre-signed URL that allows secure, time-limited access to download the file.
	DownloadURL *V2CoreBatchJobStatusDetailsValidationFailedOutputFileDownloadURL `json:"download_url"`
	// The total size of the file in bytes.
	Size int64 `json:"size,string"`
}

// Additional details for the `VALIDATION_FAILED` status of the `batch_job`.
type V2CoreBatchJobStatusDetailsValidationFailed struct {
	// The total number of records that failed processing.
	FailureCount int64 `json:"failure_count,string"`
	// The output file details. If the `batch_job` is canceled, this is provided only if there is already output at this point.
	OutputFile *V2CoreBatchJobStatusDetailsValidationFailedOutputFile `json:"output_file"`
	// The total number of records that were successfully processed.
	SuccessCount int64 `json:"success_count,string"`
}

// Additional details about the current state of the `batch_job`.
type V2CoreBatchJobStatusDetails struct {
	// Additional details for the `BATCH_FAILED` status of the `batch_job`.
	BatchFailed *V2CoreBatchJobStatusDetailsBatchFailed `json:"batch_failed,omitempty"`
	// Additional details for the `CANCELED` status of the `batch_job`.
	Canceled *V2CoreBatchJobStatusDetailsCanceled `json:"canceled,omitempty"`
	// Additional details for the `COMPLETE` status of the `batch_job`.
	Complete *V2CoreBatchJobStatusDetailsComplete `json:"complete,omitempty"`
	// Additional details for the `IN_PROGRESS` status of the `batch_job`.
	InProgress *V2CoreBatchJobStatusDetailsInProgress `json:"in_progress,omitempty"`
	// Additional details for the `READY_FOR_UPLOAD` status of the `batch_job`.
	ReadyForUpload *V2CoreBatchJobStatusDetailsReadyForUpload `json:"ready_for_upload,omitempty"`
	// Additional details for the `TIMEOUT` status of the `batch_job`.
	Timeout *V2CoreBatchJobStatusDetailsTimeout `json:"timeout,omitempty"`
	// Additional details for the `VALIDATING` status of the `batch_job`.
	Validating *V2CoreBatchJobStatusDetailsValidating `json:"validating,omitempty"`
	// Additional details for the `VALIDATION_FAILED` status of the `batch_job`.
	ValidationFailed *V2CoreBatchJobStatusDetailsValidationFailed `json:"validation_failed,omitempty"`
}

// A batch job allows you to perform an API operation on a large set of records asynchronously.
type V2CoreBatchJob struct {
	APIResource
	// Timestamp at which the `batch_job` was created.
	Created time.Time `json:"created"`
	// Unique identifier for the `batch_job`.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The maximum requests per second defined for the `batch_job`.
	MaximumRps int64 `json:"maximum_rps"`
	// The metadata of the `batch_job`.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Whether validation runs before executing the `batch_job`.
	SkipValidation bool `json:"skip_validation"`
	// The current status of the `batch_job`.
	Status V2CoreBatchJobStatus `json:"status"`
	// Additional details about the current state of the `batch_job`.
	StatusDetails *V2CoreBatchJobStatusDetails `json:"status_details,omitempty"`
}
