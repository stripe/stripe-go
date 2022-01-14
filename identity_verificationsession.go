//
//
// File generated from our OpenAPI spec
//
//

package stripe

// A short machine-readable string giving the reason for the verification or user-session failure.
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

// Array of strings of allowed identity document types. If the provided identity document isn't one of the allowed types, the verification check will fail with a document_type_not_allowed error code.
type IdentityVerificationSessionOptionsDocumentAllowedType string

// List of values that IdentityVerificationSessionOptionsDocumentAllowedType can take
const (
	IdentityVerificationSessionOptionsDocumentAllowedTypeDrivingLicense IdentityVerificationSessionOptionsDocumentAllowedType = "driving_license"
	IdentityVerificationSessionOptionsDocumentAllowedTypeIDCard         IdentityVerificationSessionOptionsDocumentAllowedType = "id_card"
	IdentityVerificationSessionOptionsDocumentAllowedTypePassport       IdentityVerificationSessionOptionsDocumentAllowedType = "passport"
)

// Indicates whether this object and its related objects have been redacted or not.
type IdentityVerificationSessionRedactionStatus string

// List of values that IdentityVerificationSessionRedactionStatus can take
const (
	IdentityVerificationSessionRedactionStatusProcessing IdentityVerificationSessionRedactionStatus = "processing"
	IdentityVerificationSessionRedactionStatusRedacted   IdentityVerificationSessionRedactionStatus = "redacted"
)

// Status of this VerificationSession. [Learn more about the lifecycle of sessions](https://stripe.com/docs/identity/how-sessions-work).
type IdentityVerificationSessionStatus string

// List of values that IdentityVerificationSessionStatus can take
const (
	IdentityVerificationSessionStatusCanceled      IdentityVerificationSessionStatus = "canceled"
	IdentityVerificationSessionStatusProcessing    IdentityVerificationSessionStatus = "processing"
	IdentityVerificationSessionStatusRequiresInput IdentityVerificationSessionStatus = "requires_input"
	IdentityVerificationSessionStatusVerified      IdentityVerificationSessionStatus = "verified"
)

// The type of [verification check](https://stripe.com/docs/identity/verification-checks) to be performed.
type IdentityVerificationSessionType string

// List of values that IdentityVerificationSessionType can take
const (
	IdentityVerificationSessionTypeDocument IdentityVerificationSessionType = "document"
	IdentityVerificationSessionTypeIDNumber IdentityVerificationSessionType = "id_number"
)

// The user's verified id number type.
type IdentityVerificationSessionVerifiedOutputsIDNumberType string

// List of values that IdentityVerificationSessionVerifiedOutputsIDNumberType can take
const (
	IdentityVerificationSessionVerifiedOutputsIDNumberTypeBRCPF  IdentityVerificationSessionVerifiedOutputsIDNumberType = "br_cpf"
	IdentityVerificationSessionVerifiedOutputsIDNumberTypeSGNRIC IdentityVerificationSessionVerifiedOutputsIDNumberType = "sg_nric"
	IdentityVerificationSessionVerifiedOutputsIDNumberTypeUSSSN  IdentityVerificationSessionVerifiedOutputsIDNumberType = "us_ssn"
)

// Options that apply to the [document check](https://stripe.com/docs/identity/verification-checks?type=document).
type IdentityVerificationSessionOptionsDocumentParams struct {
	// Array of strings of allowed identity document types. If the provided identity document isn't one of the allowed types, the verification check will fail with a document_type_not_allowed error code.
	AllowedTypes []*string `form:"allowed_types"`
	// Collect an ID number and perform an [ID number check](https://stripe.com/docs/identity/verification-checks?type=id-number) with the document's extracted name and date of birth.
	RequireIDNumber *bool `form:"require_id_number"`
	// Disable image uploads, identity document images have to be captured using the device's camera.
	RequireLiveCapture *bool `form:"require_live_capture"`
	// Capture a face image and perform a [selfie check](https://stripe.com/docs/identity/verification-checks?type=selfie) comparing a photo ID and a picture of your user's face. [Learn more](https://stripe.com/docs/identity/selfie).
	RequireMatchingSelfie *bool `form:"require_matching_selfie"`
}

// A set of options for the session's verification checks.
type IdentityVerificationSessionOptionsParams struct {
	// Options that apply to the [document check](https://stripe.com/docs/identity/verification-checks?type=document).
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
	Params `form:"*"`
	// A set of options for the session's verification checks.
	Options *IdentityVerificationSessionOptionsParams `form:"options"`
	// The URL that the user will be redirected to upon completing the verification flow.
	ReturnURL *string `form:"return_url"`
	// The type of [verification check](https://stripe.com/docs/identity/verification-checks) to be performed.
	Type *string `form:"type"`
}

// Returns a list of VerificationSessions
type IdentityVerificationSessionListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return VerificationSessions with this status. [Learn more about the lifecycle of sessions](https://stripe.com/docs/identity/how-sessions-work).
	Status *string `form:"status"`
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
	// A short machine-readable string giving the reason for the verification or user-session failure.
	Code IdentityVerificationSessionLastErrorCode `json:"code"`
	// A message that explains the reason for verification or user-session failure.
	Reason string `json:"reason"`
}
type IdentityVerificationSessionOptionsDocument struct {
	// Array of strings of allowed identity document types. If the provided identity document isn't one of the allowed types, the verification check will fail with a document_type_not_allowed error code.
	AllowedTypes []IdentityVerificationSessionOptionsDocumentAllowedType `json:"allowed_types"`
	// Collect an ID number and perform an [ID number check](https://stripe.com/docs/identity/verification-checks?type=id-number) with the document's extracted name and date of birth.
	RequireIDNumber bool `json:"require_id_number"`
	// Disable image uploads, identity document images have to be captured using the device's camera.
	RequireLiveCapture bool `json:"require_live_capture"`
	// Capture a face image and perform a [selfie check](https://stripe.com/docs/identity/verification-checks?type=selfie) comparing a photo ID and a picture of your user's face. [Learn more](https://stripe.com/docs/identity/selfie).
	RequireMatchingSelfie bool `json:"require_matching_selfie"`
}
type IdentityVerificationSessionOptionsIDNumber struct{}
type IdentityVerificationSessionOptions struct {
	Document *IdentityVerificationSessionOptionsDocument `json:"document"`
	IDNumber *IdentityVerificationSessionOptionsIDNumber `json:"id_number"`
}

// Redaction status of this VerificationSession. If the VerificationSession is not redacted, this field will be null.
type IdentityVerificationSessionRedaction struct {
	// Indicates whether this object and its related objects have been redacted or not.
	Status IdentityVerificationSessionRedactionStatus `json:"status"`
}

// The user's verified date of birth.
type IdentityVerificationSessionVerifiedOutputsDOB struct {
	// Numerical day between 1 and 31.
	Day int64 `json:"day"`
	// Numerical month between 1 and 12.
	Month int64 `json:"month"`
	// The four-digit year.
	Year int64 `json:"year"`
}

// The user's verified data.
type IdentityVerificationSessionVerifiedOutputs struct {
	// The user's verified address.
	Address *Address `json:"address"`
	// The user's verified date of birth.
	DOB *IdentityVerificationSessionVerifiedOutputsDOB `json:"dob"`
	// The user's verified first name.
	FirstName string `json:"first_name"`
	// The user's verified id number.
	IDNumber string `json:"id_number"`
	// The user's verified id number type.
	IDNumberType IdentityVerificationSessionVerifiedOutputsIDNumberType `json:"id_number_type"`
	// The user's verified last name.
	LastName string `json:"last_name"`
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
	// The short-lived client secret used by Stripe.js to [show a verification modal](https://stripe.com/docs/js/identity/modal) inside your app. This client secret expires after 24 hours and can only be used once. Don't store it, log it, embed it in a URL, or expose it to anyone other than the user. Make sure that you have TLS enabled on any page that includes the client secret. Refer to our docs on [passing the client secret to the frontend](https://stripe.com/docs/identity/verification-sessions#client-secret) to learn more.
	ClientSecret string `json:"client_secret"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If present, this property tells you the last error encountered when processing the verification.
	LastError *IdentityVerificationSessionLastError `json:"last_error"`
	// ID of the most recent VerificationReport. [Learn more about accessing detailed verification results.](https://stripe.com/docs/identity/verification-sessions#results)
	LastVerificationReport *IdentityVerificationReport `json:"last_verification_report"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object  string                              `json:"object"`
	Options *IdentityVerificationSessionOptions `json:"options"`
	// Redaction status of this VerificationSession. If the VerificationSession is not redacted, this field will be null.
	Redaction *IdentityVerificationSessionRedaction `json:"redaction"`
	// Status of this VerificationSession. [Learn more about the lifecycle of sessions](https://stripe.com/docs/identity/how-sessions-work).
	Status IdentityVerificationSessionStatus `json:"status"`
	// The type of [verification check](https://stripe.com/docs/identity/verification-checks) to be performed.
	Type IdentityVerificationSessionType `json:"type"`
	// The short-lived URL that you use to redirect a user to Stripe to submit their identity information. This URL expires after 48 hours and can only be used once. Don't store it, log it, send it in emails or expose it to anyone other than the user. Refer to our docs on [verifying identity documents](https://stripe.com/docs/identity/verify-identity-documents?platform=web&type=redirect) to learn how to redirect users to Stripe.
	URL string `json:"url"`
	// The user's verified data.
	VerifiedOutputs *IdentityVerificationSessionVerifiedOutputs `json:"verified_outputs"`
}

// IdentityVerificationSessionList is a list of VerificationSessions as retrieved from a list endpoint.
type IdentityVerificationSessionList struct {
	APIResource
	ListMeta
	Data []*IdentityVerificationSession `json:"data"`
}
