//
//
// File generated from our OpenAPI spec
//
//

package stripe

// A code indicating the reason for the error.
type PrivacyRedactionJobValidationErrorCode string

// List of values that PrivacyRedactionJobValidationErrorCode can take
const (
	PrivacyRedactionJobValidationErrorCodeInvalidCascadingSource PrivacyRedactionJobValidationErrorCode = "invalid_cascading_source"
	PrivacyRedactionJobValidationErrorCodeInvalidFilePurpose     PrivacyRedactionJobValidationErrorCode = "invalid_file_purpose"
	PrivacyRedactionJobValidationErrorCodeInvalidState           PrivacyRedactionJobValidationErrorCode = "invalid_state"
	PrivacyRedactionJobValidationErrorCodeLockedByOtherJob       PrivacyRedactionJobValidationErrorCode = "locked_by_other_job"
	PrivacyRedactionJobValidationErrorCodeTooManyObjects         PrivacyRedactionJobValidationErrorCode = "too_many_objects"
)

// Returns a list of validation errors for the specified redaction job.
type PrivacyRedactionJobValidationErrorListParams struct {
	ListParams `form:"*"`
	// RedactionJob object identifier
	Job *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PrivacyRedactionJobValidationErrorListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// If the error is related to a specific object, this field includes the object's identifier and object type.
type PrivacyRedactionJobValidationErrorErroringObject struct {
	// Unique identifier for the object.
	ID string `json:"id"`
	// Erroring object type
	ObjectType string `json:"object_type"`
}

// The Redaction Job validation error object contains information about
// errors that affect the ability to redact a specific object in a
// redaction job.
type PrivacyRedactionJobValidationError struct {
	// A code indicating the reason for the error.
	Code PrivacyRedactionJobValidationErrorCode `json:"code"`
	// If the error is related to a specific object, this field includes the object's identifier and object type.
	ErroringObject *PrivacyRedactionJobValidationErrorErroringObject `json:"erroring_object"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A human-readable message providing more details about the error.
	Message string `json:"message"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// PrivacyRedactionJobValidationErrorList is a list of RedactionJobValidationErrors as retrieved from a list endpoint.
type PrivacyRedactionJobValidationErrorList struct {
	APIResource
	ListMeta
	Data []*PrivacyRedactionJobValidationError `json:"data"`
}
