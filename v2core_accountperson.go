//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Purpose of additional address.
type V2CoreAccountPersonAdditionalAddressPurpose string

// List of values that V2CoreAccountPersonAdditionalAddressPurpose can take
const (
	V2CoreAccountPersonAdditionalAddressPurposeRegistered V2CoreAccountPersonAdditionalAddressPurpose = "registered"
)

// The purpose or type of the additional name.
type V2CoreAccountPersonAdditionalNamePurpose string

// List of values that V2CoreAccountPersonAdditionalNamePurpose can take
const (
	V2CoreAccountPersonAdditionalNamePurposeAlias  V2CoreAccountPersonAdditionalNamePurpose = "alias"
	V2CoreAccountPersonAdditionalNamePurposeMaiden V2CoreAccountPersonAdditionalNamePurpose = "maiden"
)

// The format of the document. Currently supports `files` only.
type V2CoreAccountPersonDocumentsCompanyAuthorizationType string

// List of values that V2CoreAccountPersonDocumentsCompanyAuthorizationType can take
const (
	V2CoreAccountPersonDocumentsCompanyAuthorizationTypeFiles V2CoreAccountPersonDocumentsCompanyAuthorizationType = "files"
)

// The format of the document. Currently supports `files` only.
type V2CoreAccountPersonDocumentsPassportType string

// List of values that V2CoreAccountPersonDocumentsPassportType can take
const (
	V2CoreAccountPersonDocumentsPassportTypeFiles V2CoreAccountPersonDocumentsPassportType = "files"
)

// The format of the verification document. Currently supports `front_back` only.
type V2CoreAccountPersonDocumentsPrimaryVerificationType string

// List of values that V2CoreAccountPersonDocumentsPrimaryVerificationType can take
const (
	V2CoreAccountPersonDocumentsPrimaryVerificationTypeFrontBack V2CoreAccountPersonDocumentsPrimaryVerificationType = "front_back"
)

// The format of the verification document. Currently supports `front_back` only.
type V2CoreAccountPersonDocumentsSecondaryVerificationType string

// List of values that V2CoreAccountPersonDocumentsSecondaryVerificationType can take
const (
	V2CoreAccountPersonDocumentsSecondaryVerificationTypeFrontBack V2CoreAccountPersonDocumentsSecondaryVerificationType = "front_back"
)

// The format of the document. Currently supports `files` only.
type V2CoreAccountPersonDocumentsVisaType string

// List of values that V2CoreAccountPersonDocumentsVisaType can take
const (
	V2CoreAccountPersonDocumentsVisaTypeFiles V2CoreAccountPersonDocumentsVisaType = "files"
)

// The ID number type of an individual.
type V2CoreAccountPersonIDNumberType string

// List of values that V2CoreAccountPersonIDNumberType can take
const (
	V2CoreAccountPersonIDNumberTypeAeEid       V2CoreAccountPersonIDNumberType = "ae_eid"
	V2CoreAccountPersonIDNumberTypeAoNif       V2CoreAccountPersonIDNumberType = "ao_nif"
	V2CoreAccountPersonIDNumberTypeARDni       V2CoreAccountPersonIDNumberType = "ar_dni"
	V2CoreAccountPersonIDNumberTypeAzTin       V2CoreAccountPersonIDNumberType = "az_tin"
	V2CoreAccountPersonIDNumberTypeBdBrc       V2CoreAccountPersonIDNumberType = "bd_brc"
	V2CoreAccountPersonIDNumberTypeBdEtin      V2CoreAccountPersonIDNumberType = "bd_etin"
	V2CoreAccountPersonIDNumberTypeBdNid       V2CoreAccountPersonIDNumberType = "bd_nid"
	V2CoreAccountPersonIDNumberTypeBRCPF       V2CoreAccountPersonIDNumberType = "br_cpf"
	V2CoreAccountPersonIDNumberTypeCrCpf       V2CoreAccountPersonIDNumberType = "cr_cpf"
	V2CoreAccountPersonIDNumberTypeCrDimex     V2CoreAccountPersonIDNumberType = "cr_dimex"
	V2CoreAccountPersonIDNumberTypeCrNite      V2CoreAccountPersonIDNumberType = "cr_nite"
	V2CoreAccountPersonIDNumberTypeDEStn       V2CoreAccountPersonIDNumberType = "de_stn"
	V2CoreAccountPersonIDNumberTypeDORCN       V2CoreAccountPersonIDNumberType = "do_rcn"
	V2CoreAccountPersonIDNumberTypeGtNit       V2CoreAccountPersonIDNumberType = "gt_nit"
	V2CoreAccountPersonIDNumberTypeHkID        V2CoreAccountPersonIDNumberType = "hk_id"
	V2CoreAccountPersonIDNumberTypeKzIIN       V2CoreAccountPersonIDNumberType = "kz_iin"
	V2CoreAccountPersonIDNumberTypeMXRFC       V2CoreAccountPersonIDNumberType = "mx_rfc"
	V2CoreAccountPersonIDNumberTypeMyNric      V2CoreAccountPersonIDNumberType = "my_nric"
	V2CoreAccountPersonIDNumberTypeMzNuit      V2CoreAccountPersonIDNumberType = "mz_nuit"
	V2CoreAccountPersonIDNumberTypeNLBsn       V2CoreAccountPersonIDNumberType = "nl_bsn"
	V2CoreAccountPersonIDNumberTypePeDni       V2CoreAccountPersonIDNumberType = "pe_dni"
	V2CoreAccountPersonIDNumberTypePkCnic      V2CoreAccountPersonIDNumberType = "pk_cnic"
	V2CoreAccountPersonIDNumberTypePkSnic      V2CoreAccountPersonIDNumberType = "pk_snic"
	V2CoreAccountPersonIDNumberTypeSaTin       V2CoreAccountPersonIDNumberType = "sa_tin"
	V2CoreAccountPersonIDNumberTypeSgFin       V2CoreAccountPersonIDNumberType = "sg_fin"
	V2CoreAccountPersonIDNumberTypeSGNRIC      V2CoreAccountPersonIDNumberType = "sg_nric"
	V2CoreAccountPersonIDNumberTypeTHLc        V2CoreAccountPersonIDNumberType = "th_lc"
	V2CoreAccountPersonIDNumberTypeTHPIN       V2CoreAccountPersonIDNumberType = "th_pin"
	V2CoreAccountPersonIDNumberTypeUSItin      V2CoreAccountPersonIDNumberType = "us_itin"
	V2CoreAccountPersonIDNumberTypeUSItinLast4 V2CoreAccountPersonIDNumberType = "us_itin_last_4"
	V2CoreAccountPersonIDNumberTypeUSSSN       V2CoreAccountPersonIDNumberType = "us_ssn"
	V2CoreAccountPersonIDNumberTypeUSSSNLast4  V2CoreAccountPersonIDNumberType = "us_ssn_last_4"
)

// The person's gender (International regulations require either "male" or "female").
type V2CoreAccountPersonLegalGender string

// List of values that V2CoreAccountPersonLegalGender can take
const (
	V2CoreAccountPersonLegalGenderFemale V2CoreAccountPersonLegalGender = "female"
	V2CoreAccountPersonLegalGenderMale   V2CoreAccountPersonLegalGender = "male"
)

// The person's political exposure.
type V2CoreAccountPersonPoliticalExposure string

// List of values that V2CoreAccountPersonPoliticalExposure can take
const (
	V2CoreAccountPersonPoliticalExposureExisting V2CoreAccountPersonPoliticalExposure = "existing"
	V2CoreAccountPersonPoliticalExposureNone     V2CoreAccountPersonPoliticalExposure = "none"
)

// Additional addresses associated with the person.
type V2CoreAccountPersonAdditionalAddress struct {
	// City, district, suburb, town, or village.
	City string `json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code,omitempty"`
	// Purpose of additional address.
	Purpose V2CoreAccountPersonAdditionalAddressPurpose `json:"purpose"`
	// State, county, province, or region.
	State string `json:"state,omitempty"`
	// Town or cho-me.
	Town string `json:"town,omitempty"`
}

// Additional names (e.g. aliases) associated with the person.
type V2CoreAccountPersonAdditionalName struct {
	// The individual's full name.
	FullName string `json:"full_name,omitempty"`
	// The individual's first or given name.
	GivenName string `json:"given_name,omitempty"`
	// The purpose or type of the additional name.
	Purpose V2CoreAccountPersonAdditionalNamePurpose `json:"purpose"`
	// The individual's last or family name.
	Surname string `json:"surname,omitempty"`
}

// Stripe terms of service agreement.
type V2CoreAccountPersonAdditionalTermsOfServiceAccount struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date time.Time `json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP string `json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent string `json:"user_agent,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountPersonAdditionalTermsOfService struct {
	// Stripe terms of service agreement.
	Account *V2CoreAccountPersonAdditionalTermsOfServiceAccount `json:"account,omitempty"`
}

// The person's residential address.
type V2CoreAccountPersonAddress struct {
	// City, district, suburb, town, or village.
	City string `json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code,omitempty"`
	// State, county, province, or region.
	State string `json:"state,omitempty"`
	// Town or cho-me.
	Town string `json:"town,omitempty"`
}

// The person's date of birth.
type V2CoreAccountPersonDateOfBirth struct {
	// The day of birth, between 1 and 31.
	Day int64 `json:"day"`
	// The month of birth, between 1 and 12.
	Month int64 `json:"month"`
	// The four-digit year of birth.
	Year int64 `json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountPersonDocumentsCompanyAuthorization struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []string `json:"files"`
	// The format of the document. Currently supports `files` only.
	Type V2CoreAccountPersonDocumentsCompanyAuthorizationType `json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountPersonDocumentsPassport struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []string `json:"files"`
	// The format of the document. Currently supports `files` only.
	Type V2CoreAccountPersonDocumentsPassportType `json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens for the front and back of the verification document.
type V2CoreAccountPersonDocumentsPrimaryVerificationFrontBack struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back string `json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front string `json:"front"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountPersonDocumentsPrimaryVerification struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens for the front and back of the verification document.
	FrontBack *V2CoreAccountPersonDocumentsPrimaryVerificationFrontBack `json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type V2CoreAccountPersonDocumentsPrimaryVerificationType `json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens for the front and back of the verification document.
type V2CoreAccountPersonDocumentsSecondaryVerificationFrontBack struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back string `json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front string `json:"front"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountPersonDocumentsSecondaryVerification struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens for the front and back of the verification document.
	FrontBack *V2CoreAccountPersonDocumentsSecondaryVerificationFrontBack `json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type V2CoreAccountPersonDocumentsSecondaryVerificationType `json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountPersonDocumentsVisa struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []string `json:"files"`
	// The format of the document. Currently supports `files` only.
	Type V2CoreAccountPersonDocumentsVisaType `json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountPersonDocuments struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountPersonDocumentsCompanyAuthorization `json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountPersonDocumentsPassport `json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountPersonDocumentsPrimaryVerification `json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountPersonDocumentsSecondaryVerification `json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountPersonDocumentsVisa `json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the person.
type V2CoreAccountPersonIDNumber struct {
	// The ID number type of an individual.
	Type V2CoreAccountPersonIDNumberType `json:"type"`
}

// The relationship that this person has with the Account's business or legal entity.
type V2CoreAccountPersonRelationship struct {
	// Whether the individual is an authorizer of the Account's legal entity.
	Authorizer bool `json:"authorizer,omitempty"`
	// Whether the individual is a director of the Account's legal entity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.
	Director bool `json:"director,omitempty"`
	// Whether the individual has significant responsibility to control, manage, or direct the organization.
	Executive bool `json:"executive,omitempty"`
	// Whether the individual is the legal guardian of the Account's representative.
	LegalGuardian bool `json:"legal_guardian,omitempty"`
	// Whether the individual is an owner of the Account's legal entity.
	Owner bool `json:"owner,omitempty"`
	// The percent owned by the individual of the Account's legal entity.
	PercentOwnership string `json:"percent_ownership,omitempty"`
	// Whether the individual is authorized as the primary representative of the Account. This is the person nominated by the business to provide information about themselves, and general information about the account. There can only be one representative at any given time. At the time the account is created, this person should be set to the person responsible for opening the account.
	Representative bool `json:"representative,omitempty"`
	// The individual's title (e.g., CEO, Support Engineer).
	Title string `json:"title,omitempty"`
}

// Kana Address.
type V2CoreAccountPersonScriptAddressesKana struct {
	// City, district, suburb, town, or village.
	City string `json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code,omitempty"`
	// State, county, province, or region.
	State string `json:"state,omitempty"`
	// Town or cho-me.
	Town string `json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountPersonScriptAddressesKanji struct {
	// City, district, suburb, town, or village.
	City string `json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code,omitempty"`
	// State, county, province, or region.
	State string `json:"state,omitempty"`
	// Town or cho-me.
	Town string `json:"town,omitempty"`
}

// The script addresses (e.g., non-Latin characters) associated with the person.
type V2CoreAccountPersonScriptAddresses struct {
	// Kana Address.
	Kana *V2CoreAccountPersonScriptAddressesKana `json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountPersonScriptAddressesKanji `json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountPersonScriptNamesKana struct {
	// The person's first or given name.
	GivenName string `json:"given_name,omitempty"`
	// The person's last or family name.
	Surname string `json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountPersonScriptNamesKanji struct {
	// The person's first or given name.
	GivenName string `json:"given_name,omitempty"`
	// The person's last or family name.
	Surname string `json:"surname,omitempty"`
}

// The script names (e.g. non-Latin characters) associated with the person.
type V2CoreAccountPersonScriptNames struct {
	// Persons name in kana script.
	Kana *V2CoreAccountPersonScriptNamesKana `json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountPersonScriptNamesKanji `json:"kanji,omitempty"`
}

// Person retrieval response schema.
type V2CoreAccountPerson struct {
	APIResource
	// The account ID which the individual belongs to.
	Account string `json:"account"`
	// Additional addresses associated with the person.
	AdditionalAddresses []*V2CoreAccountPersonAdditionalAddress `json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the person.
	AdditionalNames []*V2CoreAccountPersonAdditionalName `json:"additional_names,omitempty"`
	// Attestations of accepted terms of service agreements.
	AdditionalTermsOfService *V2CoreAccountPersonAdditionalTermsOfService `json:"additional_terms_of_service,omitempty"`
	// The person's residential address.
	Address *V2CoreAccountPersonAddress `json:"address,omitempty"`
	// Time at which the object was created. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// The person's date of birth.
	DateOfBirth *V2CoreAccountPersonDateOfBirth `json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountPersonDocuments `json:"documents,omitempty"`
	// The person's email address.
	Email string `json:"email,omitempty"`
	// The person's first name.
	GivenName string `json:"given_name,omitempty"`
	// Unique identifier for the Person.
	ID string `json:"id"`
	// The identification numbers (e.g., SSN) associated with the person.
	IDNumbers []*V2CoreAccountPersonIDNumber `json:"id_numbers,omitempty"`
	// The person's gender (International regulations require either "male" or "female").
	LegalGender V2CoreAccountPersonLegalGender `json:"legal_gender,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The countries where the person is a national. Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Nationalities []string `json:"nationalities,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The person's phone number.
	Phone string `json:"phone,omitempty"`
	// The person's political exposure.
	PoliticalExposure V2CoreAccountPersonPoliticalExposure `json:"political_exposure,omitempty"`
	// The relationship that this person has with the Account's business or legal entity.
	Relationship *V2CoreAccountPersonRelationship `json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the person.
	ScriptAddresses *V2CoreAccountPersonScriptAddresses `json:"script_addresses,omitempty"`
	// The script names (e.g. non-Latin characters) associated with the person.
	ScriptNames *V2CoreAccountPersonScriptNames `json:"script_names,omitempty"`
	// The person's last name.
	Surname string `json:"surname,omitempty"`
	// Time at which the object was last updated. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Updated time.Time `json:"updated"`
}
