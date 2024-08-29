//
//
// File generated from our OpenAPI spec
//
//

package stripe

type BillingMeterErrorReportReasonErrorTypeSampleErrorAPIRequest struct {
	// Unique identifier for the object.
	ID string `json:"id"`
	// idempotency_key of the request
	IdempotencyKey string `json:"idempotency_key"`
}
type BillingMeterErrorReportReasonErrorTypeSampleError struct {
	APIRequest *BillingMeterErrorReportReasonErrorTypeSampleErrorAPIRequest `json:"api_request"`
	// message of the error
	ErrorMessage string `json:"error_message"`
}

// More information about errors
type BillingMeterErrorReportReasonErrorType struct {
	SampleErrors []*BillingMeterErrorReportReasonErrorTypeSampleError `json:"sample_errors"`
}
type BillingMeterErrorReportReason struct {
	// The number of errors generated
	ErrorCount int64 `json:"error_count"`
	// More information about errors
	ErrorTypes []*BillingMeterErrorReportReasonErrorType `json:"error_types"`
}

// The related objects about the error
type BillingMeterErrorReportRelatedObject struct {
	// Unique identifier for the object.
	ID string `json:"id"`
	// The type of meter error related object. Should be 'meter'
	Object string `json:"object"`
	// The url of the meter object
	URL string `json:"url"`
}
type BillingMeterErrorReport struct {
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string                         `json:"object"`
	Reason *BillingMeterErrorReportReason `json:"reason"`
	// The related objects about the error
	RelatedObject *BillingMeterErrorReportRelatedObject `json:"related_object"`
	// Summary of invalid events
	Summary string `json:"summary"`
	// Time when validation ended. Measured in seconds since the Unix epoch
	ValidationEnd int64 `json:"validation_end"`
	// Time when validation started. Measured in seconds since the Unix epoch
	ValidationStart int64 `json:"validation_start"`
}
