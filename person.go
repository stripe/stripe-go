//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Indicates if the person or any of their representatives, family members, or other closely related persons, declares that they hold or have held an important public job or function, in any jurisdiction.
type PersonPoliticalExposure string

// List of values that PersonPoliticalExposure can take
const (
	PersonPoliticalExposureExisting PersonPoliticalExposure = "existing"
	PersonPoliticalExposureNone     PersonPoliticalExposure = "none"
)

// One of `document_corrupt`, `document_country_not_supported`, `document_expired`, `document_failed_copy`, `document_failed_other`, `document_failed_test_mode`, `document_fraudulent`, `document_failed_greyscale`, `document_incomplete`, `document_invalid`, `document_manipulated`, `document_missing_back`, `document_missing_front`, `document_not_readable`, `document_not_uploaded`, `document_photo_mismatch`, `document_too_large`, or `document_type_not_supported`. A machine-readable code specifying the verification state for this document.
type VerificationDocumentDetailsCode string

// List of values that VerificationDocumentDetailsCode can take
const (
	VerificationDocumentDetailsCodeDocumentCorrupt               VerificationDocumentDetailsCode = "document_corrupt"
	VerificationDocumentDetailsCodeDocumentFailedCopy            VerificationDocumentDetailsCode = "document_failed_copy"
	VerificationDocumentDetailsCodeDocumentFailedGreyscale       VerificationDocumentDetailsCode = "document_failed_greyscale"
	VerificationDocumentDetailsCodeDocumentFailedOther           VerificationDocumentDetailsCode = "document_failed_other"
	VerificationDocumentDetailsCodeDocumentFailedTestMode        VerificationDocumentDetailsCode = "document_failed_test_mode"
	VerificationDocumentDetailsCodeDocumentFraudulent            VerificationDocumentDetailsCode = "document_fraudulent"
	VerificationDocumentDetailsCodeDocumentIDTypeNotSupported    VerificationDocumentDetailsCode = "document_id_type_not_supported"
	VerificationDocumentDetailsCodeDocumentIDCountryNotSupported VerificationDocumentDetailsCode = "document_id_country_not_supported"
	VerificationDocumentDetailsCodeDocumentManipulated           VerificationDocumentDetailsCode = "document_manipulated"
	VerificationDocumentDetailsCodeDocumentMissingBack           VerificationDocumentDetailsCode = "document_missing_back"
	VerificationDocumentDetailsCodeDocumentMissingFront          VerificationDocumentDetailsCode = "document_missing_front"
	VerificationDocumentDetailsCodeDocumentNotReadable           VerificationDocumentDetailsCode = "document_not_readable"
	VerificationDocumentDetailsCodeDocumentNotUploaded           VerificationDocumentDetailsCode = "document_not_uploaded"
	VerificationDocumentDetailsCodeDocumentTooLarge              VerificationDocumentDetailsCode = "document_too_large"
)

// One of `document_address_mismatch`, `document_dob_mismatch`, `document_duplicate_type`, `document_id_number_mismatch`, `document_name_mismatch`, `document_nationality_mismatch`, `failed_keyed_identity`, or `failed_other`. A machine-readable code specifying the verification state for the person.
type PersonVerificationDetailsCode string

// List of values that PersonVerificationDetailsCode can take
const (
	PersonVerificationDetailsCodeFailedKeyedIdentity PersonVerificationDetailsCode = "failed_keyed_identity"
	PersonVerificationDetailsCodeFailedOther         PersonVerificationDetailsCode = "failed_other"
	PersonVerificationDetailsCodeScanNameMismatch    PersonVerificationDetailsCode = "scan_name_mismatch"
)

// The state of verification for the person. Possible values are `unverified`, `pending`, or `verified`.
type IdentityVerificationStatus string

// List of values that IdentityVerificationStatus can take
const (
	IdentityVerificationStatusPending    IdentityVerificationStatus = "pending"
	IdentityVerificationStatusUnverified IdentityVerificationStatus = "unverified"
	IdentityVerificationStatusVerified   IdentityVerificationStatus = "verified"
)

// Filters on the list of people returned based on the person's relationship to the account's company.
type RelationshipListParams struct {
	Director       *bool `form:"director"`
	Executive      *bool `form:"executive"`
	Owner          *bool `form:"owner"`
	Representative *bool `form:"representative"`
}

// Returns a list of people associated with the account's legal entity. The people are returned sorted by creation date, with the most recent people appearing first.
type PersonListParams struct {
	ListParams   `form:"*"`
	Account      *string                 `form:"-"` // Included in URL
	Relationship *RelationshipListParams `form:"relationship"`
}

// The person's date of birth.
type DOBParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type DocumentsCompanyAuthorizationParams struct {
	Files []*string `form:"files"`
}

// One or more documents showing the person's passport page with photo and personal data.
type DocumentsPassportParams struct {
	Files []*string `form:"files"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type DocumentsVisaParams struct {
	Files []*string `form:"files"`
}

// Documents that may be submitted to satisfy various informational requests.
type DocumentsParams struct {
	CompanyAuthorization *DocumentsCompanyAuthorizationParams `form:"company_authorization"`
	Passport             *DocumentsPassportParams             `form:"passport"`
	Visa                 *DocumentsVisaParams                 `form:"visa"`
}

// The relationship that this person has with the account's legal entity.
type RelationshipParams struct {
	Director         *bool    `form:"director"`
	Executive        *bool    `form:"executive"`
	Owner            *bool    `form:"owner"`
	PercentOwnership *float64 `form:"percent_ownership"`
	Representative   *bool    `form:"representative"`
	Title            *string  `form:"title"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type PersonVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// The person's verification status.
type PersonVerificationParams struct {
	AdditionalDocument *PersonVerificationDocumentParams `form:"additional_document"`
	Document           *PersonVerificationDocumentParams `form:"document"`
}

// Creates a new person.
type PersonParams struct {
	Params            `form:"*"`
	Account           *string                   `form:"-"` // Included in URL
	Address           *AccountAddressParams     `form:"address"`
	AddressKana       *AccountAddressParams     `form:"address_kana"`
	AddressKanji      *AccountAddressParams     `form:"address_kanji"`
	DOB               *DOBParams                `form:"dob"`
	Documents         *DocumentsParams          `form:"documents"`
	Email             *string                   `form:"email"`
	FirstName         *string                   `form:"first_name"`
	FirstNameKana     *string                   `form:"first_name_kana"`
	FirstNameKanji    *string                   `form:"first_name_kanji"`
	FullNameAliases   []*string                 `form:"full_name_aliases"`
	Gender            *string                   `form:"gender"`
	IDNumber          *string                   `form:"id_number"`
	LastName          *string                   `form:"last_name"`
	LastNameKana      *string                   `form:"last_name_kana"`
	LastNameKanji     *string                   `form:"last_name_kanji"`
	MaidenName        *string                   `form:"maiden_name"`
	Nationality       *string                   `form:"nationality"`
	PersonToken       *string                   `form:"person_token"`
	Phone             *string                   `form:"phone"`
	PoliticalExposure *string                   `form:"political_exposure"`
	Relationship      *RelationshipParams       `form:"relationship"`
	SSNLast4          *string                   `form:"ssn_last_4"`
	Verification      *PersonVerificationParams `form:"verification"`
}
type DOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type PersonFutureRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type PersonFutureRequirementsError struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Requirement string `json:"requirement"`
}

// Information about the upcoming new requirements for this person, including what information needs to be collected, and by when.
type PersonFutureRequirements struct {
	Alternatives        []*PersonFutureRequirementsAlternative `json:"alternatives"`
	CurrentlyDue        []string                               `json:"currently_due"`
	Errors              []*PersonFutureRequirementsError       `json:"errors"`
	EventuallyDue       []string                               `json:"eventually_due"`
	PastDue             []string                               `json:"past_due"`
	PendingVerification []string                               `json:"pending_verification"`
}
type Relationship struct {
	Director         bool    `json:"director"`
	Executive        bool    `json:"executive"`
	Owner            bool    `json:"owner"`
	PercentOwnership float64 `json:"percent_ownership"`
	Representative   bool    `json:"representative"`
	Title            string  `json:"title"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type PersonRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Information about the requirements for this person, including what information needs to be collected, and by when.
type Requirements struct {
	Alternatives        []*PersonRequirementsAlternative `json:"alternatives"`
	CurrentlyDue        []string                         `json:"currently_due"`
	Errors              []*AccountRequirementsError      `json:"errors"`
	EventuallyDue       []string                         `json:"eventually_due"`
	PastDue             []string                         `json:"past_due"`
	PendingVerification []string                         `json:"pending_verification"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type PersonVerificationDocument struct {
	Back        *File                           `json:"back"`
	Details     string                          `json:"details"`
	DetailsCode VerificationDocumentDetailsCode `json:"details_code"`
	Front       *File                           `json:"front"`
}
type PersonVerification struct {
	AdditionalDocument *PersonVerificationDocument   `json:"additional_document"`
	Details            string                        `json:"details"`
	DetailsCode        PersonVerificationDetailsCode `json:"details_code"`
	Document           *PersonVerificationDocument   `json:"document"`
	Status             IdentityVerificationStatus    `json:"status"`
}

// This is an object representing a person associated with a Stripe account.
//
// A platform cannot access a Standard or Express account's persons after the account starts onboarding, such as after generating an account link for the account.
// See the [Standard onboarding](https://stripe.com/docs/connect/standard-accounts) or [Express onboarding documentation](https://stripe.com/docs/connect/express-accounts) for information about platform pre-filling and account onboarding steps.
//
// Related guide: [Handling Identity Verification with the API](https://stripe.com/docs/connect/identity-verification-api#person-information).
type Person struct {
	APIResource
	Account            string                    `json:"account"`
	Address            *AccountAddress           `json:"address"`
	AddressKana        *AccountAddress           `json:"address_kana"`
	AddressKanji       *AccountAddress           `json:"address_kanji"`
	Created            int64                     `json:"created"`
	Deleted            bool                      `json:"deleted"`
	DOB                *DOB                      `json:"dob"`
	Email              string                    `json:"email"`
	FirstName          string                    `json:"first_name"`
	FirstNameKana      string                    `json:"first_name_kana"`
	FirstNameKanji     string                    `json:"first_name_kanji"`
	FullNameAliases    []string                  `json:"full_name_aliases"`
	FutureRequirements *PersonFutureRequirements `json:"future_requirements"`
	Gender             string                    `json:"gender"`
	ID                 string                    `json:"id"`
	IDNumberProvided   bool                      `json:"id_number_provided"`
	LastName           string                    `json:"last_name"`
	LastNameKana       string                    `json:"last_name_kana"`
	LastNameKanji      string                    `json:"last_name_kanji"`
	MaidenName         string                    `json:"maiden_name"`
	Metadata           map[string]string         `json:"metadata"`
	Nationality        string                    `json:"nationality"`
	Object             string                    `json:"object"`
	Phone              string                    `json:"phone"`
	PoliticalExposure  PersonPoliticalExposure   `json:"political_exposure"`
	Relationship       *Relationship             `json:"relationship"`
	Requirements       *Requirements             `json:"requirements"`
	SSNLast4Provided   bool                      `json:"ssn_last_4_provided"`
	Verification       *PersonVerification       `json:"verification"`
}

// PersonList is a list of Persons as retrieved from a list endpoint.
type PersonList struct {
	APIResource
	ListMeta
	Data []*Person `json:"data"`
}

// UnmarshalJSON handles deserialization of a Person.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *Person) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type person Person
	var v person
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = Person(v)
	return nil
}
