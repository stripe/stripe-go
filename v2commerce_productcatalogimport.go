//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of feed data being imported into the product catalog.
type V2CommerceProductCatalogImportFeedType string

// List of values that V2CommerceProductCatalogImportFeedType can take
const (
	V2CommerceProductCatalogImportFeedTypeInventory V2CommerceProductCatalogImportFeedType = "inventory"
	V2CommerceProductCatalogImportFeedTypePricing   V2CommerceProductCatalogImportFeedType = "pricing"
	V2CommerceProductCatalogImportFeedTypeProduct   V2CommerceProductCatalogImportFeedType = "product"
)

// The current status of this ProductCatalogImport.
type V2CommerceProductCatalogImportStatus string

// List of values that V2CommerceProductCatalogImportStatus can take
const (
	V2CommerceProductCatalogImportStatusAwaitingUpload      V2CommerceProductCatalogImportStatus = "awaiting_upload"
	V2CommerceProductCatalogImportStatusFailed              V2CommerceProductCatalogImportStatus = "failed"
	V2CommerceProductCatalogImportStatusProcessing          V2CommerceProductCatalogImportStatus = "processing"
	V2CommerceProductCatalogImportStatusSucceeded           V2CommerceProductCatalogImportStatus = "succeeded"
	V2CommerceProductCatalogImportStatusSucceededWithErrors V2CommerceProductCatalogImportStatus = "succeeded_with_errors"
)

// The error code for this product catalog processing failure.
type V2CommerceProductCatalogImportStatusDetailsFailedCode string

// List of values that V2CommerceProductCatalogImportStatusDetailsFailedCode can take
const (
	V2CommerceProductCatalogImportStatusDetailsFailedCodeFileNotFound  V2CommerceProductCatalogImportStatusDetailsFailedCode = "file_not_found"
	V2CommerceProductCatalogImportStatusDetailsFailedCodeInternalError V2CommerceProductCatalogImportStatusDetailsFailedCode = "internal_error"
	V2CommerceProductCatalogImportStatusDetailsFailedCodeInvalidFile   V2CommerceProductCatalogImportStatusDetailsFailedCode = "invalid_file"
)

// The error type for this product catalog processing failure.
type V2CommerceProductCatalogImportStatusDetailsFailedType string

// List of values that V2CommerceProductCatalogImportStatusDetailsFailedType can take
const (
	V2CommerceProductCatalogImportStatusDetailsFailedTypeCannotProceed    V2CommerceProductCatalogImportStatusDetailsFailedType = "cannot_proceed"
	V2CommerceProductCatalogImportStatusDetailsFailedTypeTransientFailure V2CommerceProductCatalogImportStatusDetailsFailedType = "transient_failure"
)

// The pre-signed URL information for uploading the catalog file.
type V2CommerceProductCatalogImportStatusDetailsAwaitingUploadUploadURL struct {
	// The timestamp when the upload URL expires.
	ExpiresAt time.Time `json:"expires_at"`
	// The pre-signed URL for uploading the catalog file.
	URL string `json:"url"`
}

// Details when the import is awaiting file upload.
type V2CommerceProductCatalogImportStatusDetailsAwaitingUpload struct {
	// The pre-signed URL information for uploading the catalog file.
	UploadURL *V2CommerceProductCatalogImportStatusDetailsAwaitingUploadUploadURL `json:"upload_url"`
}

// Details when the import didn't complete.
type V2CommerceProductCatalogImportStatusDetailsFailed struct {
	// The error code for this product catalog processing failure.
	Code V2CommerceProductCatalogImportStatusDetailsFailedCode `json:"code"`
	// A message explaining why the import failed.
	FailureMessage string `json:"failure_message"`
	// The error type for this product catalog processing failure.
	Type V2CommerceProductCatalogImportStatusDetailsFailedType `json:"type"`
}

// Details when the import is processing.
type V2CommerceProductCatalogImportStatusDetailsProcessing struct {
	// The number of records that failed to process so far.
	ErrorCount int64 `json:"error_count,string"`
	// The number of records processed so far.
	SuccessCount int64 `json:"success_count,string"`
}

// Details when the import has succeeded.
type V2CommerceProductCatalogImportStatusDetailsSucceeded struct {
	// The total number of records processed.
	SuccessCount int64 `json:"success_count,string"`
}

// The pre-signed URL information for downloading the error file.
type V2CommerceProductCatalogImportStatusDetailsSucceededWithErrorsErrorFileDownloadURL struct {
	// The timestamp when the download URL expires.
	ExpiresAt time.Time `json:"expires_at"`
	// The pre-signed URL for downloading the error file.
	URL string `json:"url"`
}

// A file containing details about all errors that occurred.
type V2CommerceProductCatalogImportStatusDetailsSucceededWithErrorsErrorFile struct {
	// The MIME type of the error file.
	ContentType string `json:"content_type"`
	// The pre-signed URL information for downloading the error file.
	DownloadURL *V2CommerceProductCatalogImportStatusDetailsSucceededWithErrorsErrorFileDownloadURL `json:"download_url"`
	// The size of the error file in bytes.
	Size int64 `json:"size,string"`
}

// A sample of errors that occurred during processing.
type V2CommerceProductCatalogImportStatusDetailsSucceededWithErrorsSample struct {
	// A description of what went wrong with this record.
	ErrorMessage string `json:"error_message"`
	// The name of the field that caused the error.
	Field string `json:"field"`
	// The identifier of the record that failed to process.
	ID string `json:"id"`
	// The row number in the import file where the error occurred.
	Row int64 `json:"row,string"`
}

// Details when the import completed but some records failed to process.
type V2CommerceProductCatalogImportStatusDetailsSucceededWithErrors struct {
	// The total number of records that failed to process.
	ErrorCount int64 `json:"error_count,string"`
	// A file containing details about all errors that occurred.
	ErrorFile *V2CommerceProductCatalogImportStatusDetailsSucceededWithErrorsErrorFile `json:"error_file"`
	// A sample of errors that occurred during processing.
	Samples []*V2CommerceProductCatalogImportStatusDetailsSucceededWithErrorsSample `json:"samples"`
	// The total number of records processed.
	SuccessCount int64 `json:"success_count,string"`
}

// Details about the current import status.
type V2CommerceProductCatalogImportStatusDetails struct {
	// Details when the import is awaiting file upload.
	AwaitingUpload *V2CommerceProductCatalogImportStatusDetailsAwaitingUpload `json:"awaiting_upload,omitempty"`
	// Details when the import didn't complete.
	Failed *V2CommerceProductCatalogImportStatusDetailsFailed `json:"failed,omitempty"`
	// Details when the import is processing.
	Processing *V2CommerceProductCatalogImportStatusDetailsProcessing `json:"processing,omitempty"`
	// Details when the import has succeeded.
	Succeeded *V2CommerceProductCatalogImportStatusDetailsSucceeded `json:"succeeded,omitempty"`
	// Details when the import completed but some records failed to process.
	SucceededWithErrors *V2CommerceProductCatalogImportStatusDetailsSucceededWithErrors `json:"succeeded_with_errors,omitempty"`
}

// The product catalog import object tracks the long-running background process that handles uploading, processing and validation.
type V2CommerceProductCatalogImport struct {
	APIResource
	// The time this ProductCatalogImport was created.
	Created time.Time `json:"created"`
	// The type of feed data being imported into the product catalog.
	FeedType V2CommerceProductCatalogImportFeedType `json:"feed_type"`
	// The unique identifier for this ProductCatalogImport.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The current status of this ProductCatalogImport.
	Status V2CommerceProductCatalogImportStatus `json:"status"`
	// Details about the current import status.
	StatusDetails *V2CommerceProductCatalogImportStatusDetails `json:"status_details,omitempty"`
}
