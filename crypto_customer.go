//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The KYC region determined by the customer's address country.
type CryptoCustomerKycRegion string

// List of values that CryptoCustomerKycRegion can take
const (
	CryptoCustomerKycRegionEU CryptoCustomerKycRegion = "eu"
	CryptoCustomerKycRegionUS CryptoCustomerKycRegion = "us"
)

// The KYC tier level (e.g., l0, l1, l2).
type CryptoCustomerKycTierTier string

// List of values that CryptoCustomerKycTierTier can take
const (
	CryptoCustomerKycTierTierL0 CryptoCustomerKycTierTier = "l0"
	CryptoCustomerKycTierTierL1 CryptoCustomerKycTierTier = "l1"
	CryptoCustomerKycTierTierL2 CryptoCustomerKycTierTier = "l2"
)

// List of errors associated with this KYC tier verification.
type CryptoCustomerKycTierVerificationError string

// List of values that CryptoCustomerKycTierVerificationError can take
const (
	CryptoCustomerKycTierVerificationErrorIDDocumentVerificationFailed          CryptoCustomerKycTierVerificationError = "id_document_verification_failed"
	CryptoCustomerKycTierVerificationErrorPhoneVerificationFailed               CryptoCustomerKycTierVerificationError = "phone_verification_failed"
	CryptoCustomerKycTierVerificationErrorUserHasReachedMaxVerificationAttempts CryptoCustomerKycTierVerificationError = "user_has_reached_max_verification_attempts"
)

// The verification status for this KYC tier.
type CryptoCustomerKycTierVerificationStatus string

// List of values that CryptoCustomerKycTierVerificationStatus can take
const (
	CryptoCustomerKycTierVerificationStatusNotAvailable CryptoCustomerKycTierVerificationStatus = "not_available"
	CryptoCustomerKycTierVerificationStatusNotStarted   CryptoCustomerKycTierVerificationStatus = "not_started"
	CryptoCustomerKycTierVerificationStatusPending      CryptoCustomerKycTierVerificationStatus = "pending"
	CryptoCustomerKycTierVerificationStatusRejected     CryptoCustomerKycTierVerificationStatus = "rejected"
	CryptoCustomerKycTierVerificationStatusVerified     CryptoCustomerKycTierVerificationStatus = "verified"
)

// The set of KYC Fields provided for this customers.
type CryptoCustomerProvidedField string

// List of values that CryptoCustomerProvidedField can take
const (
	CryptoCustomerProvidedFieldAddressCity       CryptoCustomerProvidedField = "address_city"
	CryptoCustomerProvidedFieldAddressCountry    CryptoCustomerProvidedField = "address_country"
	CryptoCustomerProvidedFieldAddressLine1      CryptoCustomerProvidedField = "address_line_1"
	CryptoCustomerProvidedFieldAddressLine2      CryptoCustomerProvidedField = "address_line_2"
	CryptoCustomerProvidedFieldAddressPostalCode CryptoCustomerProvidedField = "address_postal_code"
	CryptoCustomerProvidedFieldAddressState      CryptoCustomerProvidedField = "address_state"
	CryptoCustomerProvidedFieldAttestation       CryptoCustomerProvidedField = "attestation"
	CryptoCustomerProvidedFieldBirthCity         CryptoCustomerProvidedField = "birth_city"
	CryptoCustomerProvidedFieldBirthCountry      CryptoCustomerProvidedField = "birth_country"
	CryptoCustomerProvidedFieldDOB               CryptoCustomerProvidedField = "dob"
	CryptoCustomerProvidedFieldFirstName         CryptoCustomerProvidedField = "first_name"
	CryptoCustomerProvidedFieldIDDocument        CryptoCustomerProvidedField = "id_document"
	CryptoCustomerProvidedFieldIDNumber          CryptoCustomerProvidedField = "id_number"
	CryptoCustomerProvidedFieldIDType            CryptoCustomerProvidedField = "id_type"
	CryptoCustomerProvidedFieldIdentifiers       CryptoCustomerProvidedField = "identifiers"
	CryptoCustomerProvidedFieldLastName          CryptoCustomerProvidedField = "last_name"
	CryptoCustomerProvidedFieldNationalities     CryptoCustomerProvidedField = "nationalities"
	CryptoCustomerProvidedFieldSelfie            CryptoCustomerProvidedField = "selfie"
)

// List of errors associated with the verification.
type CryptoCustomerVerificationError string

// List of values that CryptoCustomerVerificationError can take
const (
	CryptoCustomerVerificationErrorIDDocumentVerificationFailed          CryptoCustomerVerificationError = "id_document_verification_failed"
	CryptoCustomerVerificationErrorPhoneVerificationFailed               CryptoCustomerVerificationError = "phone_verification_failed"
	CryptoCustomerVerificationErrorUserHasReachedMaxVerificationAttempts CryptoCustomerVerificationError = "user_has_reached_max_verification_attempts"
)

// Type of verification.
type CryptoCustomerVerificationName string

// List of values that CryptoCustomerVerificationName can take
const (
	CryptoCustomerVerificationNameIDDocumentVerified CryptoCustomerVerificationName = "id_document_verified"
	CryptoCustomerVerificationNameKycVerified        CryptoCustomerVerificationName = "kyc_verified"
	CryptoCustomerVerificationNamePhoneVerified      CryptoCustomerVerificationName = "phone_verified"
)

// Outcome of the verification.
type CryptoCustomerVerificationStatus string

// List of values that CryptoCustomerVerificationStatus can take
const (
	CryptoCustomerVerificationStatusNotAvailable CryptoCustomerVerificationStatus = "not_available"
	CryptoCustomerVerificationStatusNotStarted   CryptoCustomerVerificationStatus = "not_started"
	CryptoCustomerVerificationStatusPending      CryptoCustomerVerificationStatus = "pending"
	CryptoCustomerVerificationStatusRejected     CryptoCustomerVerificationStatus = "rejected"
	CryptoCustomerVerificationStatusVerified     CryptoCustomerVerificationStatus = "verified"
)

// Retrieves the details of a Crypto Customer.
type CryptoCustomerParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoCustomerParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a Crypto Customer.
type CryptoCustomerRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoCustomerRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// List of KYC tiers and their verification status.
type CryptoCustomerKycTier struct {
	// The KYC tier level (e.g., l0, l1, l2).
	Tier CryptoCustomerKycTierTier `json:"tier"`
	// List of errors associated with this KYC tier verification.
	VerificationErrors []CryptoCustomerKycTierVerificationError `json:"verification_errors"`
	// The verification status for this KYC tier.
	VerificationStatus CryptoCustomerKycTierVerificationStatus `json:"verification_status"`
}

// List of verifications and their outcome.
type CryptoCustomerVerification struct {
	// List of errors associated with the verification.
	Errors []CryptoCustomerVerificationError `json:"errors"`
	// Type of verification.
	Name CryptoCustomerVerificationName `json:"name"`
	// Outcome of the verification.
	Status CryptoCustomerVerificationStatus `json:"status"`
}

// This object represents a crypto onramp customer. Use it to get their kyc status and payment methods.
type CryptoCustomer struct {
	APIResource
	// Unique identifier for the object.
	ID string `json:"id"`
	// The KYC region determined by the customer's address country.
	KycRegion CryptoCustomerKycRegion `json:"kyc_region"`
	// List of KYC tiers and their verification status.
	KycTiers []*CryptoCustomerKycTier `json:"kyc_tiers"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The set of KYC Fields provided for this customers.
	ProvidedFields []CryptoCustomerProvidedField `json:"provided_fields"`
	// List of verifications and their outcome.
	Verifications []*CryptoCustomerVerification `json:"verifications"`
}
