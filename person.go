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
	// A filter on the list of people returned based on whether these people are directors of the account's company.
	Director *bool `form:"director"`
	// A filter on the list of people returned based on whether these people are executives of the account's company.
	Executive *bool `form:"executive"`
	// A filter on the list of people returned based on whether these people are owners of the account's company.
	Owner *bool `form:"owner"`
	// A filter on the list of people returned based on whether these people are the representative of the account's company.
	Representative *bool `form:"representative"`
}

// Returns a list of people associated with the account's legal entity. The people are returned sorted by creation date, with the most recent people appearing first.
type PersonListParams struct {
	ListParams `form:"*"`
	Account    *string `form:"-"` // Included in URL
	// Filters on the list of people returned based on the person's relationship to the account's company.
	Relationship *RelationshipListParams `form:"relationship"`
}

// The person's date of birth.
type DOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type DocumentsCompanyAuthorizationParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// One or more documents showing the person's passport page with photo and personal data.
type DocumentsPassportParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type DocumentsVisaParams struct {
	// One or more document ids returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `account_requirement`.
	Files []*string `form:"files"`
}

// Documents that may be submitted to satisfy various informational requests.
type DocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *DocumentsCompanyAuthorizationParams `form:"company_authorization"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *DocumentsPassportParams `form:"passport"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *DocumentsVisaParams `form:"visa"`
}

// The relationship that this person has with the account's legal entity.
type RelationshipParams struct {
	// Whether the person is a director of the account's legal entity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.
	Director *bool `form:"director"`
	// Whether the person has significant responsibility to control, manage, or direct the organization.
	Executive *bool `form:"executive"`
	// Whether the person is an owner of the account's legal entity.
	Owner *bool `form:"owner"`
	// The percent owned by the person of the account's legal entity.
	PercentOwnership *float64 `form:"percent_ownership"`
	// Whether the person is authorized as the primary representative of the account. This is the person nominated by the business to provide information about themselves, and general information about the account. There can only be one representative at any given time. At the time the account is created, this person should be set to the person responsible for opening the account.
	Representative *bool `form:"representative"`
	// The person's title (e.g., CEO, Support Engineer).
	Title *string `form:"title"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type PersonVerificationDocumentParams struct {
	// The back of an ID returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `identity_document`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back"`
	// The front of an ID returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `identity_document`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front"`
}

// The person's verification status.
type PersonVerificationParams struct {
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	AdditionalDocument *PersonVerificationDocumentParams `form:"additional_document"`
	// An identifying document, either a passport or local ID card.
	Document *PersonVerificationDocumentParams `form:"document"`
}

// Creates a new person.
type PersonParams struct {
	Params  `form:"*"`
	Account *string `form:"-"` // Included in URL
	// The person's address.
	Address *AccountAddressParams `form:"address"`
	// The Kana variation of the person's address (Japan only).
	AddressKana *AccountAddressParams `form:"address_kana"`
	// The Kanji variation of the person's address (Japan only).
	AddressKanji *AccountAddressParams `form:"address_kanji"`
	// The person's date of birth.
	DOB *DOBParams `form:"dob"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *DocumentsParams `form:"documents"`
	// The person's email address.
	Email *string `form:"email"`
	// The person's first name.
	FirstName *string `form:"first_name"`
	// The Kana variation of the person's first name (Japan only).
	FirstNameKana *string `form:"first_name_kana"`
	// The Kanji variation of the person's first name (Japan only).
	FirstNameKanji *string `form:"first_name_kanji"`
	// A list of alternate names or aliases that the person is known by.
	FullNameAliases []*string `form:"full_name_aliases"`
	// The person's gender (International regulations require either "male" or "female").
	Gender *string `form:"gender"`
	// The person's ID number, as appropriate for their country. For example, a social security number in the U.S., social insurance number in Canada, etc. Instead of the number itself, you can also provide a [PII token provided by Stripe.js](https://stripe.com/docs/js/tokens_sources/create_token?type=pii).
	IDNumber *string `form:"id_number"`
	// The person's secondary ID number, as appropriate for their country, will be used for enhanced verification checks. In Thailand, this would be the laser code found on the back of an ID card. Instead of the number itself, you can also provide a [PII token provided by Stripe.js](https://stripe.com/docs/js/tokens_sources/create_token?type=pii).
	IDNumberSecondary *string `form:"id_number_secondary"`
	// The person's last name.
	LastName *string `form:"last_name"`
	// The Kana variation of the person's last name (Japan only).
	LastNameKana *string `form:"last_name_kana"`
	// The Kanji variation of the person's last name (Japan only).
	LastNameKanji *string `form:"last_name_kanji"`
	// The person's maiden name.
	MaidenName *string `form:"maiden_name"`
	// The country where the person is a national. Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)), or "XX" if unavailable.
	Nationality *string `form:"nationality"`
	// A [person token](https://stripe.com/docs/connect/account-tokens), used to securely provide details to the person.
	PersonToken *string `form:"person_token"`
	// The person's phone number.
	Phone *string `form:"phone"`
	// Indicates if the person or any of their representatives, family members, or other closely related persons, declares that they hold or have held an important public job or function, in any jurisdiction.
	PoliticalExposure *string `form:"political_exposure"`
	// The person's registered address.
	RegisteredAddress *AddressParams `form:"registered_address"`
	// The relationship that this person has with the account's legal entity.
	Relationship *RelationshipParams `form:"relationship"`
	// The last four digits of the person's Social Security number (U.S. only).
	SSNLast4 *string `form:"ssn_last_4"`
	// The person's verification status.
	Verification *PersonVerificationParams `form:"verification"`
}
type DOB struct {
	// The day of birth, between 1 and 31.
	Day int64 `json:"day"`
	// The month of birth, between 1 and 12.
	Month int64 `json:"month"`
	// The four-digit year of birth.
	Year int64 `json:"year"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type PersonFutureRequirementsAlternative struct {
	// Fields that can be provided to satisfy all fields in `original_fields_due`.
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	// Fields that are due and can be satisfied by providing all fields in `alternative_fields_due`.
	OriginalFieldsDue []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type PersonFutureRequirementsError struct {
	// The code for the type of error.
	Code string `json:"code"`
	// An informative message that indicates the error type and provides additional details about the error.
	Reason string `json:"reason"`
	// The specific user onboarding requirement field (in the requirements hash) that needs to be resolved.
	Requirement string `json:"requirement"`
}

// Information about the upcoming new requirements for this person, including what information needs to be collected, and by when.
type PersonFutureRequirements struct {
	// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
	Alternatives []*PersonFutureRequirementsAlternative `json:"alternatives"`
	// Fields that need to be collected to keep the person's account enabled. If not collected by the account's `future_requirements[current_deadline]`, these fields will transition to the main `requirements` hash, and may immediately become `past_due`, but the account may also be given a grace period depending on the account's enablement state prior to transition.
	CurrentlyDue []string `json:"currently_due"`
	// Fields that are `currently_due` and need to be collected again because validation or verification failed.
	Errors []*PersonFutureRequirementsError `json:"errors"`
	// Fields that need to be collected assuming all volume thresholds are reached. As they become required, they appear in `currently_due` as well, and the account's `future_requirements[current_deadline]` becomes set.
	EventuallyDue []string `json:"eventually_due"`
	// Fields that weren't collected by the account's `requirements.current_deadline`. These fields need to be collected to enable the person's account. New fields will never appear here; `future_requirements.past_due` will always be a subset of `requirements.past_due`.
	PastDue []string `json:"past_due"`
	// Fields that may become required depending on the results of verification or review. Will be an empty array unless an asynchronous verification is pending. If verification fails, these fields move to `eventually_due` or `currently_due`.
	PendingVerification []string `json:"pending_verification"`
}
type Relationship struct {
	// Whether the person is a director of the account's legal entity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.
	Director bool `json:"director"`
	// Whether the person has significant responsibility to control, manage, or direct the organization.
	Executive bool `json:"executive"`
	// Whether the person is an owner of the account's legal entity.
	Owner bool `json:"owner"`
	// The percent owned by the person of the account's legal entity.
	PercentOwnership float64 `json:"percent_ownership"`
	// Whether the person is authorized as the primary representative of the account. This is the person nominated by the business to provide information about themselves, and general information about the account. There can only be one representative at any given time. At the time the account is created, this person should be set to the person responsible for opening the account.
	Representative bool `json:"representative"`
	// The person's title (e.g., CEO, Support Engineer).
	Title string `json:"title"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type PersonRequirementsAlternative struct {
	// Fields that can be provided to satisfy all fields in `original_fields_due`.
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	// Fields that are due and can be satisfied by providing all fields in `alternative_fields_due`.
	OriginalFieldsDue []string `json:"original_fields_due"`
}

// Information about the requirements for this person, including what information needs to be collected, and by when.
type Requirements struct {
	// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
	Alternatives []*PersonRequirementsAlternative `json:"alternatives"`
	// Fields that need to be collected to keep the person's account enabled. If not collected by the account's `current_deadline`, these fields appear in `past_due` as well, and the account is disabled.
	CurrentlyDue []string `json:"currently_due"`
	// Fields that are `currently_due` and need to be collected again because validation or verification failed.
	Errors []*AccountRequirementsError `json:"errors"`
	// Fields that need to be collected assuming all volume thresholds are reached. As they become required, they appear in `currently_due` as well, and the account's `current_deadline` becomes set.
	EventuallyDue []string `json:"eventually_due"`
	// Fields that weren't collected by the account's `current_deadline`. These fields need to be collected to enable the person's account.
	PastDue []string `json:"past_due"`
	// Fields that may become required depending on the results of verification or review. Will be an empty array unless an asynchronous verification is pending. If verification fails, these fields move to `eventually_due`, `currently_due`, or `past_due`.
	PendingVerification []string `json:"pending_verification"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type PersonVerificationDocument struct {
	// The back of an ID returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `identity_document`.
	Back *File `json:"back"`
	// A user-displayable string describing the verification state of this document. For example, if a document is uploaded and the picture is too fuzzy, this may say "Identity document is too unclear to read".
	Details string `json:"details"`
	// One of `document_corrupt`, `document_country_not_supported`, `document_expired`, `document_failed_copy`, `document_failed_other`, `document_failed_test_mode`, `document_fraudulent`, `document_failed_greyscale`, `document_incomplete`, `document_invalid`, `document_manipulated`, `document_missing_back`, `document_missing_front`, `document_not_readable`, `document_not_uploaded`, `document_photo_mismatch`, `document_too_large`, or `document_type_not_supported`. A machine-readable code specifying the verification state for this document.
	DetailsCode VerificationDocumentDetailsCode `json:"details_code"`
	// The front of an ID returned by a [file upload](https://stripe.com/docs/api#create_file) with a `purpose` value of `identity_document`.
	Front *File `json:"front"`
}
type PersonVerification struct {
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	AdditionalDocument *PersonVerificationDocument `json:"additional_document"`
	// A user-displayable string describing the verification state for the person. For example, this may say "Provided identity information could not be verified".
	Details string `json:"details"`
	// One of `document_address_mismatch`, `document_dob_mismatch`, `document_duplicate_type`, `document_id_number_mismatch`, `document_name_mismatch`, `document_nationality_mismatch`, `failed_keyed_identity`, or `failed_other`. A machine-readable code specifying the verification state for the person.
	DetailsCode PersonVerificationDetailsCode `json:"details_code"`
	Document    *PersonVerificationDocument   `json:"document"`
	// The state of verification for the person. Possible values are `unverified`, `pending`, or `verified`.
	Status IdentityVerificationStatus `json:"status"`
}

// This is an object representing a person associated with a Stripe account.
//
// A platform cannot access a Standard or Express account's persons after the account starts onboarding, such as after generating an account link for the account.
// See the [Standard onboarding](https://stripe.com/docs/connect/standard-accounts) or [Express onboarding documentation](https://stripe.com/docs/connect/express-accounts) for information about platform pre-filling and account onboarding steps.
//
// Related guide: [Handling Identity Verification with the API](https://stripe.com/docs/connect/identity-verification-api#person-information).
type Person struct {
	APIResource
	// The account the person is associated with.
	Account string          `json:"account"`
	Address *AccountAddress `json:"address"`
	// The Kana variation of the person's address (Japan only).
	AddressKana *AccountAddress `json:"address_kana"`
	// The Kanji variation of the person's address (Japan only).
	AddressKanji *AccountAddress `json:"address_kanji"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	Deleted bool  `json:"deleted"`
	DOB     *DOB  `json:"dob"`
	// The person's email address.
	Email string `json:"email"`
	// The person's first name.
	FirstName string `json:"first_name"`
	// The Kana variation of the person's first name (Japan only).
	FirstNameKana string `json:"first_name_kana"`
	// The Kanji variation of the person's first name (Japan only).
	FirstNameKanji string `json:"first_name_kanji"`
	// A list of alternate names or aliases that the person is known by.
	FullNameAliases []string `json:"full_name_aliases"`
	// Information about the upcoming new requirements for this person, including what information needs to be collected, and by when.
	FutureRequirements *PersonFutureRequirements `json:"future_requirements"`
	// The person's gender (International regulations require either "male" or "female").
	Gender string `json:"gender"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Whether the person's `id_number` was provided.
	IDNumberProvided bool `json:"id_number_provided"`
	// Whether the person's `id_number_secondary` was provided.
	IDNumberSecondaryProvided bool `json:"id_number_secondary_provided"`
	// The person's last name.
	LastName string `json:"last_name"`
	// The Kana variation of the person's last name (Japan only).
	LastNameKana string `json:"last_name_kana"`
	// The Kanji variation of the person's last name (Japan only).
	LastNameKanji string `json:"last_name_kanji"`
	// The person's maiden name.
	MaidenName string `json:"maiden_name"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The country where the person is a national.
	Nationality string `json:"nationality"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The person's phone number.
	Phone string `json:"phone"`
	// Indicates if the person or any of their representatives, family members, or other closely related persons, declares that they hold or have held an important public job or function, in any jurisdiction.
	PoliticalExposure PersonPoliticalExposure `json:"political_exposure"`
	RegisteredAddress *Address                `json:"registered_address"`
	Relationship      *Relationship           `json:"relationship"`
	// Information about the requirements for this person, including what information needs to be collected, and by when.
	Requirements *Requirements `json:"requirements"`
	// Whether the last four digits of the person's Social Security number have been provided (U.S. only).
	SSNLast4Provided bool                `json:"ssn_last_4_provided"`
	Verification     *PersonVerification `json:"verification"`
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
