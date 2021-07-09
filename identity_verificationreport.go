//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type IdentityVerificationReportDocumentErrorCode string

// List of values that IdentityVerificationReportDocumentErrorCode can take
const (
	IdentityVerificationReportDocumentErrorCodeDocumentExpired          IdentityVerificationReportDocumentErrorCode = "document_expired"
	IdentityVerificationReportDocumentErrorCodeDocumentTypeNotSupported IdentityVerificationReportDocumentErrorCode = "document_type_not_supported"
	IdentityVerificationReportDocumentErrorCodeDocumentUnverifiedOther  IdentityVerificationReportDocumentErrorCode = "document_unverified_other"
)

type IdentityVerificationReportDocumentStatus string

// List of values that IdentityVerificationReportDocumentStatus can take
const (
	IdentityVerificationReportDocumentStatusUnverified IdentityVerificationReportDocumentStatus = "unverified"
	IdentityVerificationReportDocumentStatusVerified   IdentityVerificationReportDocumentStatus = "verified"
)

type IdentityVerificationReportDocumentType string

// List of values that IdentityVerificationReportDocumentType can take
const (
	IdentityVerificationReportDocumentTypeDrivingLicense IdentityVerificationReportDocumentType = "driving_license"
	IdentityVerificationReportDocumentTypeIDCard         IdentityVerificationReportDocumentType = "id_card"
	IdentityVerificationReportDocumentTypePassport       IdentityVerificationReportDocumentType = "passport"
)

type IdentityVerificationReportIDNumberErrorCode string

// List of values that IdentityVerificationReportIDNumberErrorCode can take
const (
	IdentityVerificationReportIDNumberErrorCodeIDNumberInsufficientDocumentData IdentityVerificationReportIDNumberErrorCode = "id_number_insufficient_document_data"
	IdentityVerificationReportIDNumberErrorCodeIDNumberMismatch                 IdentityVerificationReportIDNumberErrorCode = "id_number_mismatch"
	IdentityVerificationReportIDNumberErrorCodeIDNumberUnverifiedOther          IdentityVerificationReportIDNumberErrorCode = "id_number_unverified_other"
)

type IdentityVerificationReportIDNumberIDNumberType string

// List of values that IdentityVerificationReportIDNumberIDNumberType can take
const (
	IdentityVerificationReportIDNumberIDNumberTypeBRCPF  IdentityVerificationReportIDNumberIDNumberType = "br_cpf"
	IdentityVerificationReportIDNumberIDNumberTypeSGNRIC IdentityVerificationReportIDNumberIDNumberType = "sg_nric"
	IdentityVerificationReportIDNumberIDNumberTypeUSSSN  IdentityVerificationReportIDNumberIDNumberType = "us_ssn"
)

type IdentityVerificationReportIDNumberStatus string

// List of values that IdentityVerificationReportIDNumberStatus can take
const (
	IdentityVerificationReportIDNumberStatusUnverified IdentityVerificationReportIDNumberStatus = "unverified"
	IdentityVerificationReportIDNumberStatusVerified   IdentityVerificationReportIDNumberStatus = "verified"
)

type IdentityVerificationReportOptionsDocumentAllowedType string

// List of values that IdentityVerificationReportOptionsDocumentAllowedType can take
const (
	IdentityVerificationReportOptionsDocumentAllowedTypeDrivingLicense IdentityVerificationReportOptionsDocumentAllowedType = "driving_license"
	IdentityVerificationReportOptionsDocumentAllowedTypeIDCard         IdentityVerificationReportOptionsDocumentAllowedType = "id_card"
	IdentityVerificationReportOptionsDocumentAllowedTypePassport       IdentityVerificationReportOptionsDocumentAllowedType = "passport"
)

type IdentityVerificationReportSelfieErrorCode string

// List of values that IdentityVerificationReportSelfieErrorCode can take
const (
	IdentityVerificationReportSelfieErrorCodeSelfieDocumentMissingPhoto IdentityVerificationReportSelfieErrorCode = "selfie_document_missing_photo"
	IdentityVerificationReportSelfieErrorCodeSelfieFaceMismatch         IdentityVerificationReportSelfieErrorCode = "selfie_face_mismatch"
	IdentityVerificationReportSelfieErrorCodeSelfieManipulated          IdentityVerificationReportSelfieErrorCode = "selfie_manipulated"
	IdentityVerificationReportSelfieErrorCodeSelfieUnverifiedOther      IdentityVerificationReportSelfieErrorCode = "selfie_unverified_other"
)

type IdentityVerificationReportSelfieStatus string

// List of values that IdentityVerificationReportSelfieStatus can take
const (
	IdentityVerificationReportSelfieStatusUnverified IdentityVerificationReportSelfieStatus = "unverified"
	IdentityVerificationReportSelfieStatusVerified   IdentityVerificationReportSelfieStatus = "verified"
)

type IdentityVerificationReportType string

// List of values that IdentityVerificationReportType can take
const (
	IdentityVerificationReportTypeDocument IdentityVerificationReportType = "document"
	IdentityVerificationReportTypeIDNumber IdentityVerificationReportType = "id_number"
)

// Retrieves an existing VerificationReport
type IdentityVerificationReportParams struct {
	Params `form:"*"`
}

// List all verification reports.
type IdentityVerificationReportListParams struct {
	ListParams          `form:"*"`
	Created             *int64            `form:"created"`
	CreatedRange        *RangeQueryParams `form:"created"`
	Type                *string           `form:"type"`
	VerificationSession *string           `form:"verification_session"`
}

// Date of birth as it appears in the document.
type IdentityVerificationReportDocumentDOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// Details on the verification error. Present when status is `unverified`.
type IdentityVerificationReportDocumentError struct {
	Code   IdentityVerificationReportDocumentErrorCode `json:"code"`
	Reason string                                      `json:"reason"`
}

// Expiration date of the document.
type IdentityVerificationReportDocumentExpirationDate struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// Issued date of the document.
type IdentityVerificationReportDocumentIssuedDate struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// Result from a document check
type IdentityVerificationReportDocument struct {
	Address        *Address                                          `json:"address"`
	DOB            *IdentityVerificationReportDocumentDOB            `json:"dob"`
	Error          *IdentityVerificationReportDocumentError          `json:"error"`
	ExpirationDate *IdentityVerificationReportDocumentExpirationDate `json:"expiration_date"`
	Files          []string                                          `json:"files"`
	FirstName      string                                            `json:"first_name"`
	IssuedDate     *IdentityVerificationReportDocumentIssuedDate     `json:"issued_date"`
	IssuingCountry string                                            `json:"issuing_country"`
	LastName       string                                            `json:"last_name"`
	Number         string                                            `json:"number"`
	Status         IdentityVerificationReportDocumentStatus          `json:"status"`
	Type           IdentityVerificationReportDocumentType            `json:"type"`
}

// Date of birth.
type IdentityVerificationReportIDNumberDOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// Details on the verification error. Present when status is `unverified`.
type IdentityVerificationReportIDNumberError struct {
	Code   IdentityVerificationReportIDNumberErrorCode `json:"code"`
	Reason string                                      `json:"reason"`
}

// Result from an id_number check
type IdentityVerificationReportIDNumber struct {
	DOB          *IdentityVerificationReportIDNumberDOB         `json:"dob"`
	Error        *IdentityVerificationReportIDNumberError       `json:"error"`
	FirstName    string                                         `json:"first_name"`
	IDNumber     string                                         `json:"id_number"`
	IDNumberType IdentityVerificationReportIDNumberIDNumberType `json:"id_number_type"`
	LastName     string                                         `json:"last_name"`
	Status       IdentityVerificationReportIDNumberStatus       `json:"status"`
}
type IdentityVerificationReportOptionsDocument struct {
	AllowedTypes          []IdentityVerificationReportOptionsDocumentAllowedType `json:"allowed_types"`
	RequireIDNumber       bool                                                   `json:"require_id_number"`
	RequireLiveCapture    bool                                                   `json:"require_live_capture"`
	RequireMatchingSelfie bool                                                   `json:"require_matching_selfie"`
}
type IdentityVerificationReportOptionsIDNumber struct{}
type IdentityVerificationReportOptions struct {
	Document *IdentityVerificationReportOptionsDocument `json:"document"`
	IDNumber *IdentityVerificationReportOptionsIDNumber `json:"id_number"`
}

// Details on the verification error. Present when status is `unverified`.
type IdentityVerificationReportSelfieError struct {
	Code   IdentityVerificationReportSelfieErrorCode `json:"code"`
	Reason string                                    `json:"reason"`
}

// Result from a selfie check
type IdentityVerificationReportSelfie struct {
	Document string                                 `json:"document"`
	Error    *IdentityVerificationReportSelfieError `json:"error"`
	Selfie   string                                 `json:"selfie"`
	Status   IdentityVerificationReportSelfieStatus `json:"status"`
}

// A VerificationReport is the result of an attempt to collect and verify data from a user.
// The collection of verification checks performed is determined from the `type` and `options`
// parameters used. You can find the result of each verification check performed in the
// appropriate sub-resource: `document`, `id_number`, `selfie`.
//
// Each VerificationReport contains a copy of any data collected by the user as well as
// reference IDs which can be used to access collected images through the [FileUpload](https://stripe.com/docs/api/files)
// API. To configure and create VerificationReports, use the
// [VerificationSession](https://stripe.com/docs/api/identity/verification_sessions) API.
//
// Related guides: [Accessing verification results](https://stripe.com/docs/identity/verification-sessions#results).
type IdentityVerificationReport struct {
	APIResource
	Created             int64                               `json:"created"`
	Document            *IdentityVerificationReportDocument `json:"document"`
	ID                  string                              `json:"id"`
	IDNumber            *IdentityVerificationReportIDNumber `json:"id_number"`
	Livemode            bool                                `json:"livemode"`
	Object              string                              `json:"object"`
	Options             *IdentityVerificationReportOptions  `json:"options"`
	Selfie              *IdentityVerificationReportSelfie   `json:"selfie"`
	Type                IdentityVerificationReportType      `json:"type"`
	VerificationSession string                              `json:"verification_session"`
}

// IdentityVerificationReportList is a list of VerificationReports as retrieved from a list endpoint.
type IdentityVerificationReportList struct {
	APIResource
	ListMeta
	Data []*IdentityVerificationReport `json:"data"`
}

// UnmarshalJSON handles deserialization of an IdentityVerificationReport.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IdentityVerificationReport) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type identityVerificationReport IdentityVerificationReport
	var v identityVerificationReport
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IdentityVerificationReport(v)
	return nil
}
