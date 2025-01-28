//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

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
	ValidationEnd time.Time `json:"validation_end"`
	// Time when validation started. Measured in seconds since the Unix epoch
	ValidationStart time.Time `json:"validation_start"`
}

// UnmarshalJSON handles deserialization of a BillingMeterErrorReport.
// This custom unmarshaling is needed to handle the time fields correctly.
func (b *BillingMeterErrorReport) UnmarshalJSON(data []byte) error {
	type billingMeterErrorReport BillingMeterErrorReport
	v := struct {
		ValidationEnd   int64 `json:"validation_end"`
		ValidationStart int64 `json:"validation_start"`
		*billingMeterErrorReport
	}{
		billingMeterErrorReport: (*billingMeterErrorReport)(b),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	b.ValidationEnd = time.Unix(v.ValidationEnd, 0)
	b.ValidationStart = time.Unix(v.ValidationStart, 0)
	return nil
}

// MarshalJSON handles serialization of a BillingMeterErrorReport.
// This custom marshaling is needed to handle the time fields correctly.
func (b BillingMeterErrorReport) MarshalJSON() ([]byte, error) {
	type billingMeterErrorReport BillingMeterErrorReport
	v := struct {
		ValidationEnd   int64 `json:"validation_end"`
		ValidationStart int64 `json:"validation_start"`
		billingMeterErrorReport
	}{
		billingMeterErrorReport: (billingMeterErrorReport)(b),
		ValidationEnd:           b.ValidationEnd.Unix(),
		ValidationStart:         b.ValidationStart.Unix(),
	}
	bb, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bb, err
}
