//
//
// File generated from our OpenAPI spec
//
//

package stripe

type IdentityVerificationSessionLastErrorCode string

const (
	IdentityVerificationSessionLastErrorCodeAbandoned                        IdentityVerificationSessionLastErrorCode = "abandoned"
	IdentityVerificationSessionLastErrorCodeConsentDeclined                  IdentityVerificationSessionLastErrorCode = "consent_declined"
	IdentityVerificationSessionLastErrorCodeCountryNotSupported              IdentityVerificationSessionLastErrorCode = "country_not_supported"
	IdentityVerificationSessionLastErrorCodeDeviceNotSupported               IdentityVerificationSessionLastErrorCode = "device_not_supported"
	IdentityVerificationSessionLastErrorCodeDocumentExpired                  IdentityVerificationSessionLastErrorCode = "document_expired"
	IdentityVerificationSessionLastErrorCodeDocumentTypeNotSupported         IdentityVerificationSessionLastErrorCode = "document_type_not_supported"
	IdentityVerificationSessionLastErrorCodeDocumentUnverifiedOther          IdentityVerificationSessionLastErrorCode = "document_unverified_other"
	IdentityVerificationSessionLastErrorCodeIDNumberInsufficientDocumentData IdentityVerificationSessionLastErrorCode = "id_number_insufficient_document_data"
	IdentityVerificationSessionLastErrorCodeIDNumberMismatch                 IdentityVerificationSessionLastErrorCode = "id_number_mismatch"
	IdentityVerificationSessionLastErrorCodeIDNumberUnverifiedOther          IdentityVerificationSessionLastErrorCode = "id_number_unverified_other"
	IdentityVerificationSessionLastErrorCodeSelfieDocumentMissingPhoto       IdentityVerificationSessionLastErrorCode = "selfie_document_missing_photo"
	IdentityVerificationSessionLastErrorCodeSelfieFaceMismatch               IdentityVerificationSessionLastErrorCode = "selfie_face_mismatch"
	IdentityVerificationSessionLastErrorCodeSelfieManipulated                IdentityVerificationSessionLastErrorCode = "selfie_manipulated"
	IdentityVerificationSessionLastErrorCodeSelfieUnverifiedOther            IdentityVerificationSessionLastErrorCode = "selfie_unverified_other"
	IdentityVerificationSessionLastErrorCodeUnderSupportedAge                IdentityVerificationSessionLastErrorCode = "under_supported_age"
)

type IdentityVerificationSessionOptionsDocumentAllowedType string

const (
	IdentityVerificationSessionOptionsDocumentAllowedTypeDrivingLicense IdentityVerificationSessionOptionsDocumentAllowedType = "driving_license"
	IdentityVerificationSessionOptionsDocumentAllowedTypeIDCard         IdentityVerificationSessionOptionsDocumentAllowedType = "id_card"
	IdentityVerificationSessionOptionsDocumentAllowedTypePassport       IdentityVerificationSessionOptionsDocumentAllowedType = "passport"
)

type IdentityVerificationSessionRedactionStatus string

const (
	IdentityVerificationSessionRedactionStatusProcessing IdentityVerificationSessionRedactionStatus = "processing"
	IdentityVerificationSessionRedactionStatusRedacted   IdentityVerificationSessionRedactionStatus = "redacted"
)

type IdentityVerificationSessionStatus string

const (
	IdentityVerificationSessionStatusCanceled      IdentityVerificationSessionStatus = "canceled"
	IdentityVerificationSessionStatusProcessing    IdentityVerificationSessionStatus = "processing"
	IdentityVerificationSessionStatusRequiresInput IdentityVerificationSessionStatus = "requires_input"
	IdentityVerificationSessionStatusVerified      IdentityVerificationSessionStatus = "verified"
)

type IdentityVerificationSessionType string

const (
	IdentityVerificationSessionTypeDocument IdentityVerificationSessionType = "document"
	IdentityVerificationSessionTypeIDNumber IdentityVerificationSessionType = "id_number"
)

type IdentityVerificationSessionVerifiedOutputsIDNumberType string

const (
	IdentityVerificationSessionVerifiedOutputsIDNumberTypeBRCPF  IdentityVerificationSessionVerifiedOutputsIDNumberType = "br_cpf"
	IdentityVerificationSessionVerifiedOutputsIDNumberTypeSGNRIC IdentityVerificationSessionVerifiedOutputsIDNumberType = "sg_nric"
	IdentityVerificationSessionVerifiedOutputsIDNumberTypeUSSSN  IdentityVerificationSessionVerifiedOutputsIDNumberType = "us_ssn"
)

type IdentityVerificationSessionOptionsDocumentParams struct {
	AllowedTypes          []*string `form:"allowed_types"`
	RequireIDNumber       *bool     `form:"require_id_number"`
	RequireLiveCapture    *bool     `form:"require_live_capture"`
	RequireMatchingSelfie *bool     `form:"require_matching_selfie"`
}
type IdentityVerificationSessionOptionsParams struct {
	Document *IdentityVerificationSessionOptionsDocumentParams `form:"document"`
}
type IdentityVerificationSessionParams struct {
	Params    `form:"*"`
	Options   *IdentityVerificationSessionOptionsParams `form:"options"`
	ReturnURL *string                                   `form:"return_url"`
	Type      *string                                   `form:"type"`
}
type IdentityVerificationSessionListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Status       *string           `form:"status"`
}
type IdentityVerificationSessionCancelParams struct {
	Params `form:"*"`
}
type IdentityVerificationSessionRedactParams struct {
	Params `form:"*"`
}
type IdentityVerificationSessionLastError struct {
	Code   IdentityVerificationSessionLastErrorCode `json:"code"`
	Reason string                                   `json:"reason"`
}
type IdentityVerificationSessionOptionsDocument struct {
	AllowedTypes          []IdentityVerificationSessionOptionsDocumentAllowedType `json:"allowed_types"`
	RequireIDNumber       bool                                                    `json:"require_id_number"`
	RequireLiveCapture    bool                                                    `json:"require_live_capture"`
	RequireMatchingSelfie bool                                                    `json:"require_matching_selfie"`
}
type IdentityVerificationSessionOptionsIDNumber struct{}
type IdentityVerificationSessionOptions struct {
	Document *IdentityVerificationSessionOptionsDocument `json:"document"`
	IDNumber *IdentityVerificationSessionOptionsIDNumber `json:"id_number"`
}
type IdentityVerificationSessionRedaction struct {
	Status IdentityVerificationSessionRedactionStatus `json:"status"`
}
type IdentityVerificationSessionVerifiedOutputsDOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}
type IdentityVerificationSessionVerifiedOutputs struct {
	Address      *Address                                               `json:"address"`
	DOB          *IdentityVerificationSessionVerifiedOutputsDOB         `json:"dob"`
	FirstName    string                                                 `json:"first_name"`
	IDNumber     string                                                 `json:"id_number"`
	IDNumberType IdentityVerificationSessionVerifiedOutputsIDNumberType `json:"id_number_type"`
	LastName     string                                                 `json:"last_name"`
}
type IdentityVerificationSession struct {
	APIResource
	ClientSecret           string                                      `json:"client_secret"`
	Created                int64                                       `json:"created"`
	ID                     string                                      `json:"id"`
	LastError              *IdentityVerificationSessionLastError       `json:"last_error"`
	LastVerificationReport *IdentityVerificationReport                 `json:"last_verification_report"`
	Livemode               bool                                        `json:"livemode"`
	Metadata               map[string]string                           `json:"metadata"`
	Object                 string                                      `json:"object"`
	Options                *IdentityVerificationSessionOptions         `json:"options"`
	Redaction              *IdentityVerificationSessionRedaction       `json:"redaction"`
	Status                 IdentityVerificationSessionStatus           `json:"status"`
	Type                   IdentityVerificationSessionType             `json:"type"`
	URL                    string                                      `json:"url"`
	VerifiedOutputs        *IdentityVerificationSessionVerifiedOutputs `json:"verified_outputs"`
}
type IdentityVerificationSessionList struct {
	APIResource
	ListMeta
	Data []*IdentityVerificationSession `json:"data"`
}
