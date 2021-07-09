//
//
// File generated from our OpenAPI spec
//
//

package stripe

type IdentityVerificationSessionLastErrorCode string

// List of values that IdentityVerificationSessionLastErrorCode can take
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

// List of values that IdentityVerificationSessionOptionsDocumentAllowedType can take
const (
	IdentityVerificationSessionOptionsDocumentAllowedTypeDrivingLicense IdentityVerificationSessionOptionsDocumentAllowedType = "driving_license"
	IdentityVerificationSessionOptionsDocumentAllowedTypeIDCard         IdentityVerificationSessionOptionsDocumentAllowedType = "id_card"
	IdentityVerificationSessionOptionsDocumentAllowedTypePassport       IdentityVerificationSessionOptionsDocumentAllowedType = "passport"
)

type IdentityVerificationSessionRedactionStatus string

// List of values that IdentityVerificationSessionRedactionStatus can take
const (
	IdentityVerificationSessionRedactionStatusProcessing IdentityVerificationSessionRedactionStatus = "processing"
	IdentityVerificationSessionRedactionStatusRedacted   IdentityVerificationSessionRedactionStatus = "redacted"
)

type IdentityVerificationSessionStatus string

// List of values that IdentityVerificationSessionStatus can take
const (
	IdentityVerificationSessionStatusCanceled      IdentityVerificationSessionStatus = "canceled"
	IdentityVerificationSessionStatusProcessing    IdentityVerificationSessionStatus = "processing"
	IdentityVerificationSessionStatusRequiresInput IdentityVerificationSessionStatus = "requires_input"
	IdentityVerificationSessionStatusVerified      IdentityVerificationSessionStatus = "verified"
)

type IdentityVerificationSessionType string

// List of values that IdentityVerificationSessionType can take
const (
	IdentityVerificationSessionTypeDocument IdentityVerificationSessionType = "document"
	IdentityVerificationSessionTypeIDNumber IdentityVerificationSessionType = "id_number"
)

type IdentityVerificationSessionVerifiedOutputsIDNumberType string

// List of values that IdentityVerificationSessionVerifiedOutputsIDNumberType can take
const (
	IdentityVerificationSessionVerifiedOutputsIDNumberTypeBRCPF  IdentityVerificationSessionVerifiedOutputsIDNumberType = "br_cpf"
	IdentityVerificationSessionVerifiedOutputsIDNumberTypeSGNRIC IdentityVerificationSessionVerifiedOutputsIDNumberType = "sg_nric"
	IdentityVerificationSessionVerifiedOutputsIDNumberTypeUSSSN  IdentityVerificationSessionVerifiedOutputsIDNumberType = "us_ssn"
)

// Options that apply to the [document check](https://stripe.com/docs/identity/verification-checks?type=document).
type IdentityVerificationSessionOptionsDocumentParams struct {
	AllowedTypes          []*string `form:"allowed_types"`
	RequireIDNumber       *bool     `form:"require_id_number"`
	RequireLiveCapture    *bool     `form:"require_live_capture"`
	RequireMatchingSelfie *bool     `form:"require_matching_selfie"`
}

// A set of options for the session's verification checks.
type IdentityVerificationSessionOptionsParams struct {
	Document *IdentityVerificationSessionOptionsDocumentParams `form:"document"`
}

// Creates a VerificationSession object.
//
// After the VerificationSession is created, display a verification modal using the session client_secret or send your users to the session's url.
//
// If your API key is in test mode, verification checks won't actually process, though everything else will occur as if in live mode.
//
// Related guide: [Verify your users' identity documents](https://stripe.com/docs/identity/verify-identity-documents).
type IdentityVerificationSessionParams struct {
	Params    `form:"*"`
	Options   *IdentityVerificationSessionOptionsParams `form:"options"`
	ReturnURL *string                                   `form:"return_url"`
	Type      *string                                   `form:"type"`
}

// Returns a list of VerificationSessions
type IdentityVerificationSessionListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Status       *string           `form:"status"`
}

// A VerificationSession object can be canceled when it is in requires_input [status](https://stripe.com/docs/identity/how-sessions-work).
//
// Once canceled, future submission attempts are disabled. This cannot be undone. [Learn more](https://stripe.com/docs/identity/verification-sessions#cancel).
type IdentityVerificationSessionCancelParams struct {
	Params `form:"*"`
}

// Redact a VerificationSession to remove all collected information from Stripe. This will redact
// the VerificationSession and all objects related to it, including VerificationReports, Events,
// request logs, etc.
//
// A VerificationSession object can be redacted when it is in requires_input or verified
// [status](https://stripe.com/docs/identity/how-sessions-work). Redacting a VerificationSession in requires_action
// state will automatically cancel it.
//
// The redaction process may take up to four days. When the redaction process is in progress, the
// VerificationSession's redaction.status field will be set to processing; when the process is
// finished, it will change to redacted and an identity.verification_session.redacted event
// will be emitted.
//
// Redaction is irreversible. Redacted objects are still accessible in the Stripe API, but all the
// fields that contain personal data will be replaced by the string [redacted] or a similar
// placeholder. The metadata field will also be erased. Redacted objects cannot be updated or
// used for any purpose.
//
// [Learn more](https://stripe.com/docs/identity/verification-sessions#redact).
type IdentityVerificationSessionRedactParams struct {
	Params `form:"*"`
}

// If present, this property tells you the last error encountered when processing the verification.
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

// Redaction status of this VerificationSession. If the VerificationSession is not redacted, this field will be null.
type IdentityVerificationSessionRedaction struct {
	Status IdentityVerificationSessionRedactionStatus `json:"status"`
}

// The user's verified date of birth.
type IdentityVerificationSessionVerifiedOutputsDOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// The user's verified data.
type IdentityVerificationSessionVerifiedOutputs struct {
	Address      *Address                                               `json:"address"`
	DOB          *IdentityVerificationSessionVerifiedOutputsDOB         `json:"dob"`
	FirstName    string                                                 `json:"first_name"`
	IDNumber     string                                                 `json:"id_number"`
	IDNumberType IdentityVerificationSessionVerifiedOutputsIDNumberType `json:"id_number_type"`
	LastName     string                                                 `json:"last_name"`
}

// A VerificationSession guides you through the process of collecting and verifying the identities
// of your users. It contains details about the type of verification, such as what [verification
// check](https://stripe.com/docs/identity/verification-checks) to perform. Only create one VerificationSession for
// each verification in your system.
//
// A VerificationSession transitions through [multiple
// statuses](https://stripe.com/docs/identity/how-sessions-work) throughout its lifetime as it progresses through
// the verification flow. The VerificationSession contains the user's verified data after
// verification checks are complete.
//
// Related guide: [The Verification Sessions API](https://stripe.com/docs/identity/verification-sessions)
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

// IdentityVerificationSessionList is a list of VerificationSessions as retrieved from a list endpoint.
type IdentityVerificationSessionList struct {
	APIResource
	ListMeta
	Data []*IdentityVerificationSession `json:"data"`
}
