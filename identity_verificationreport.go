//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type IdentityVerificationReportDocumentErrorCode string

const (
	IdentityVerificationReportDocumentErrorCodeDocumentExpired          IdentityVerificationReportDocumentErrorCode = "document_expired"
	IdentityVerificationReportDocumentErrorCodeDocumentTypeNotSupported IdentityVerificationReportDocumentErrorCode = "document_type_not_supported"
	IdentityVerificationReportDocumentErrorCodeDocumentUnverifiedOther  IdentityVerificationReportDocumentErrorCode = "document_unverified_other"
)

type IdentityVerificationReportDocumentStatus string

const (
	IdentityVerificationReportDocumentStatusUnverified IdentityVerificationReportDocumentStatus = "unverified"
	IdentityVerificationReportDocumentStatusVerified   IdentityVerificationReportDocumentStatus = "verified"
)

type IdentityVerificationReportDocumentType string

const (
	IdentityVerificationReportDocumentTypeDrivingLicense IdentityVerificationReportDocumentType = "driving_license"
	IdentityVerificationReportDocumentTypeIDCard         IdentityVerificationReportDocumentType = "id_card"
	IdentityVerificationReportDocumentTypePassport       IdentityVerificationReportDocumentType = "passport"
)

type IdentityVerificationReportIDNumberErrorCode string

const (
	IdentityVerificationReportIDNumberErrorCodeIDNumberInsufficientDocumentData IdentityVerificationReportIDNumberErrorCode = "id_number_insufficient_document_data"
	IdentityVerificationReportIDNumberErrorCodeIDNumberMismatch                 IdentityVerificationReportIDNumberErrorCode = "id_number_mismatch"
	IdentityVerificationReportIDNumberErrorCodeIDNumberUnverifiedOther          IdentityVerificationReportIDNumberErrorCode = "id_number_unverified_other"
)

type IdentityVerificationReportIDNumberIDNumberType string

const (
	IdentityVerificationReportIDNumberIDNumberTypeBRCPF  IdentityVerificationReportIDNumberIDNumberType = "br_cpf"
	IdentityVerificationReportIDNumberIDNumberTypeSGNRIC IdentityVerificationReportIDNumberIDNumberType = "sg_nric"
	IdentityVerificationReportIDNumberIDNumberTypeUSSSN  IdentityVerificationReportIDNumberIDNumberType = "us_ssn"
)

type IdentityVerificationReportIDNumberStatus string

const (
	IdentityVerificationReportIDNumberStatusUnverified IdentityVerificationReportIDNumberStatus = "unverified"
	IdentityVerificationReportIDNumberStatusVerified   IdentityVerificationReportIDNumberStatus = "verified"
)

type IdentityVerificationReportOptionsDocumentAllowedType string

const (
	IdentityVerificationReportOptionsDocumentAllowedTypeDrivingLicense IdentityVerificationReportOptionsDocumentAllowedType = "driving_license"
	IdentityVerificationReportOptionsDocumentAllowedTypeIDCard         IdentityVerificationReportOptionsDocumentAllowedType = "id_card"
	IdentityVerificationReportOptionsDocumentAllowedTypePassport       IdentityVerificationReportOptionsDocumentAllowedType = "passport"
)

type IdentityVerificationReportSelfieErrorCode string

const (
	IdentityVerificationReportSelfieErrorCodeSelfieDocumentMissingPhoto IdentityVerificationReportSelfieErrorCode = "selfie_document_missing_photo"
	IdentityVerificationReportSelfieErrorCodeSelfieFaceMismatch         IdentityVerificationReportSelfieErrorCode = "selfie_face_mismatch"
	IdentityVerificationReportSelfieErrorCodeSelfieManipulated          IdentityVerificationReportSelfieErrorCode = "selfie_manipulated"
	IdentityVerificationReportSelfieErrorCodeSelfieUnverifiedOther      IdentityVerificationReportSelfieErrorCode = "selfie_unverified_other"
)

type IdentityVerificationReportSelfieStatus string

const (
	IdentityVerificationReportSelfieStatusUnverified IdentityVerificationReportSelfieStatus = "unverified"
	IdentityVerificationReportSelfieStatusVerified   IdentityVerificationReportSelfieStatus = "verified"
)

type IdentityVerificationReportType string

const (
	IdentityVerificationReportTypeDocument IdentityVerificationReportType = "document"
	IdentityVerificationReportTypeIDNumber IdentityVerificationReportType = "id_number"
)

type IdentityVerificationReportParams struct {
	Params `form:"*"`
}
type IdentityVerificationReportListParams struct {
	ListParams          `form:"*"`
	Created             *int64            `form:"created"`
	CreatedRange        *RangeQueryParams `form:"created"`
	Type                *string           `form:"type"`
	VerificationSession *string           `form:"verification_session"`
}
type IdentityVerificationReportDocumentDOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}
type IdentityVerificationReportDocumentError struct {
	Code   IdentityVerificationReportDocumentErrorCode `json:"code"`
	Reason string                                      `json:"reason"`
}
type IdentityVerificationReportDocumentExpirationDate struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}
type IdentityVerificationReportDocumentIssuedDate struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}
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
type IdentityVerificationReportIDNumberDOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}
type IdentityVerificationReportIDNumberError struct {
	Code   IdentityVerificationReportIDNumberErrorCode `json:"code"`
	Reason string                                      `json:"reason"`
}
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
type IdentityVerificationReportSelfieError struct {
	Code   IdentityVerificationReportSelfieErrorCode `json:"code"`
	Reason string                                    `json:"reason"`
}
type IdentityVerificationReportSelfie struct {
	Document string                                 `json:"document"`
	Error    *IdentityVerificationReportSelfieError `json:"error"`
	Selfie   string                                 `json:"selfie"`
	Status   IdentityVerificationReportSelfieStatus `json:"status"`
}
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
