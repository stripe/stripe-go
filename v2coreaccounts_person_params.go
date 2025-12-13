//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Returns a list of Persons associated with an Account.
type V2CoreAccountsPersonListParams struct {
	Params `form:"*"`
	// Account the Persons are associated with.
	AccountID *string `form:"-" json:"-"` // Included in URL
	// The upper limit on the number of accounts returned by the List Account request.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Additional addresses associated with the person.
type V2CoreAccountsPersonAdditionalAddressParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// Purpose of additional address.
	Purpose *string `form:"purpose" json:"purpose"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// Additional names (e.g. aliases) associated with the person.
type V2CoreAccountsPersonAdditionalNameParams struct {
	// The person's full name.
	FullName *string `form:"full_name" json:"full_name,omitempty"`
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The purpose or type of the additional name.
	Purpose *string `form:"purpose" json:"purpose"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Stripe terms of service agreement.
type V2CoreAccountsPersonAdditionalTermsOfServiceAccountParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountsPersonAdditionalTermsOfServiceParams struct {
	// Stripe terms of service agreement.
	Account *V2CoreAccountsPersonAdditionalTermsOfServiceAccountParams `form:"account" json:"account,omitempty"`
}

// The person's residential address.
type V2CoreAccountsPersonAddressParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// The person's date of birth.
type V2CoreAccountsPersonDateOfBirthParams struct {
	// The day of the birth.
	Day *int64 `form:"day" json:"day"`
	// The month of birth.
	Month *int64 `form:"month" json:"month"`
	// The year of birth.
	Year *int64 `form:"year" json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountsPersonDocumentsCompanyAuthorizationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountsPersonDocumentsPassportParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountsPersonDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountsPersonDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountsPersonDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountsPersonDocumentsSecondaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountsPersonDocumentsSecondaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountsPersonDocumentsSecondaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountsPersonDocumentsVisaParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountsPersonDocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountsPersonDocumentsCompanyAuthorizationParams `form:"company_authorization" json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountsPersonDocumentsPassportParams `form:"passport" json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountsPersonDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountsPersonDocumentsSecondaryVerificationParams `form:"secondary_verification" json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountsPersonDocumentsVisaParams `form:"visa" json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the person.
type V2CoreAccountsPersonIDNumberParams struct {
	// The ID number type of an individual.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// The relationship that this person has with the Account's business or legal entity.
type V2CoreAccountsPersonRelationshipParams struct {
	// Whether the individual is an authorizer of the Account's legal entity.
	Authorizer *bool `form:"authorizer" json:"authorizer,omitempty"`
	// Indicates whether the person is a director of the associated legal entity.
	Director *bool `form:"director" json:"director,omitempty"`
	// Indicates whether the person is an executive of the associated legal entity.
	Executive *bool `form:"executive" json:"executive,omitempty"`
	// Indicates whether the person is a legal guardian of the associated legal entity.
	LegalGuardian *bool `form:"legal_guardian" json:"legal_guardian,omitempty"`
	// Indicates whether the person is an owner of the associated legal entity.
	Owner *bool `form:"owner" json:"owner,omitempty"`
	// The percentage of ownership the person has in the associated legal entity.
	PercentOwnership *string `form:"percent_ownership" json:"percent_ownership,omitempty"`
	// Indicates whether the person is a representative of the associated legal entity.
	Representative *bool `form:"representative" json:"representative,omitempty"`
	// The title or position the person holds in the associated legal entity.
	Title *string `form:"title" json:"title,omitempty"`
}

// Kana Address.
type V2CoreAccountsPersonScriptAddressesKanaParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountsPersonScriptAddressesKanjiParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// The script addresses (e.g., non-Latin characters) associated with the person.
type V2CoreAccountsPersonScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountsPersonScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountsPersonScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountsPersonScriptNamesKanaParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountsPersonScriptNamesKanjiParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The script names (e.g. non-Latin characters) associated with the person.
type V2CoreAccountsPersonScriptNamesParams struct {
	// Persons name in kana script.
	Kana *V2CoreAccountsPersonScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountsPersonScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Create a Person associated with an Account.
type V2CoreAccountsPersonParams struct {
	Params `form:"*"`
	// Account the Person should be associated with.
	AccountID *string `form:"-" json:"-"` // Included in URL
	// Additional addresses associated with the person.
	AdditionalAddresses []*V2CoreAccountsPersonAdditionalAddressParams `form:"additional_addresses" json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the person.
	AdditionalNames []*V2CoreAccountsPersonAdditionalNameParams `form:"additional_names" json:"additional_names,omitempty"`
	// Attestations of accepted terms of service agreements.
	AdditionalTermsOfService *V2CoreAccountsPersonAdditionalTermsOfServiceParams `form:"additional_terms_of_service" json:"additional_terms_of_service,omitempty"`
	// The primary address associated with the person.
	Address *V2CoreAccountsPersonAddressParams `form:"address" json:"address,omitempty"`
	// The person's date of birth.
	DateOfBirth *V2CoreAccountsPersonDateOfBirthParams `form:"date_of_birth" json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountsPersonDocumentsParams `form:"documents" json:"documents,omitempty"`
	// Email.
	Email *string `form:"email" json:"email,omitempty"`
	// The person's first name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The identification numbers (e.g., SSN) associated with the person.
	IDNumbers []*V2CoreAccountsPersonIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// The person's gender (International regulations require either "male" or "female").
	LegalGender *string `form:"legal_gender" json:"legal_gender,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// The nationalities (countries) this person is associated with.
	Nationalities []*string `form:"nationalities" json:"nationalities,omitempty"`
	// The person token generated by the person token api.
	PersonToken *string `form:"person_token" json:"person_token,omitempty"`
	// The phone number for this person.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The person's political exposure.
	PoliticalExposure *string `form:"political_exposure" json:"political_exposure,omitempty"`
	// The relationship that this person has with the Account's business or legal entity.
	Relationship *V2CoreAccountsPersonRelationshipParams `form:"relationship" json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the person.
	ScriptAddresses *V2CoreAccountsPersonScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The script names (e.g. non-Latin characters) associated with the person.
	ScriptNames *V2CoreAccountsPersonScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The person's last name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountsPersonParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Additional addresses associated with the person.
type V2CoreAccountsPersonCreateAdditionalAddressParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// Purpose of additional address.
	Purpose *string `form:"purpose" json:"purpose"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// Additional names (e.g. aliases) associated with the person.
type V2CoreAccountsPersonCreateAdditionalNameParams struct {
	// The person's full name.
	FullName *string `form:"full_name" json:"full_name,omitempty"`
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The purpose or type of the additional name.
	Purpose *string `form:"purpose" json:"purpose"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Stripe terms of service agreement.
type V2CoreAccountsPersonCreateAdditionalTermsOfServiceAccountParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountsPersonCreateAdditionalTermsOfServiceParams struct {
	// Stripe terms of service agreement.
	Account *V2CoreAccountsPersonCreateAdditionalTermsOfServiceAccountParams `form:"account" json:"account,omitempty"`
}

// The person's residential address.
type V2CoreAccountsPersonCreateAddressParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// The person's date of birth.
type V2CoreAccountsPersonCreateDateOfBirthParams struct {
	// The day of birth.
	Day *int64 `form:"day" json:"day"`
	// The month of birth.
	Month *int64 `form:"month" json:"month"`
	// The year of birth.
	Year *int64 `form:"year" json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountsPersonCreateDocumentsCompanyAuthorizationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountsPersonCreateDocumentsPassportParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountsPersonCreateDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountsPersonCreateDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountsPersonCreateDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountsPersonCreateDocumentsSecondaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountsPersonCreateDocumentsSecondaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountsPersonCreateDocumentsSecondaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountsPersonCreateDocumentsVisaParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountsPersonCreateDocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountsPersonCreateDocumentsCompanyAuthorizationParams `form:"company_authorization" json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountsPersonCreateDocumentsPassportParams `form:"passport" json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountsPersonCreateDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountsPersonCreateDocumentsSecondaryVerificationParams `form:"secondary_verification" json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountsPersonCreateDocumentsVisaParams `form:"visa" json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the person.
type V2CoreAccountsPersonCreateIDNumberParams struct {
	// The ID number type of an individual.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// The relationship that this person has with the Account's business or legal entity.
type V2CoreAccountsPersonCreateRelationshipParams struct {
	// Whether the individual is an authorizer of the Account's legal entity.
	Authorizer *bool `form:"authorizer" json:"authorizer,omitempty"`
	// Indicates whether the person is a director of the associated legal entity.
	Director *bool `form:"director" json:"director,omitempty"`
	// Indicates whether the person is an executive of the associated legal entity.
	Executive *bool `form:"executive" json:"executive,omitempty"`
	// Indicates whether the person is a legal guardian of the associated legal entity.
	LegalGuardian *bool `form:"legal_guardian" json:"legal_guardian,omitempty"`
	// Indicates whether the person is an owner of the associated legal entity.
	Owner *bool `form:"owner" json:"owner,omitempty"`
	// The percentage of ownership the person has in the associated legal entity.
	PercentOwnership *string `form:"percent_ownership" json:"percent_ownership,omitempty"`
	// Indicates whether the person is a representative of the associated legal entity.
	Representative *bool `form:"representative" json:"representative,omitempty"`
	// The title or position the person holds in the associated legal entity.
	Title *string `form:"title" json:"title,omitempty"`
}

// Kana Address.
type V2CoreAccountsPersonCreateScriptAddressesKanaParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountsPersonCreateScriptAddressesKanjiParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// The script addresses (e.g., non-Latin characters) associated with the person.
type V2CoreAccountsPersonCreateScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountsPersonCreateScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountsPersonCreateScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountsPersonCreateScriptNamesKanaParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountsPersonCreateScriptNamesKanjiParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The script names (e.g. non-Latin characters) associated with the person.
type V2CoreAccountsPersonCreateScriptNamesParams struct {
	// Persons name in kana script.
	Kana *V2CoreAccountsPersonCreateScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountsPersonCreateScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Create a Person associated with an Account.
type V2CoreAccountsPersonCreateParams struct {
	Params `form:"*"`
	// Account the Person should be associated with.
	AccountID *string `form:"-" json:"-"` // Included in URL
	// Additional addresses associated with the person.
	AdditionalAddresses []*V2CoreAccountsPersonCreateAdditionalAddressParams `form:"additional_addresses" json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the person.
	AdditionalNames []*V2CoreAccountsPersonCreateAdditionalNameParams `form:"additional_names" json:"additional_names,omitempty"`
	// Attestations of accepted terms of service agreements.
	AdditionalTermsOfService *V2CoreAccountsPersonCreateAdditionalTermsOfServiceParams `form:"additional_terms_of_service" json:"additional_terms_of_service,omitempty"`
	// The person's residential address.
	Address *V2CoreAccountsPersonCreateAddressParams `form:"address" json:"address,omitempty"`
	// The person's date of birth.
	DateOfBirth *V2CoreAccountsPersonCreateDateOfBirthParams `form:"date_of_birth" json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountsPersonCreateDocumentsParams `form:"documents" json:"documents,omitempty"`
	// Email.
	Email *string `form:"email" json:"email,omitempty"`
	// The person's first name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The identification numbers (e.g., SSN) associated with the person.
	IDNumbers []*V2CoreAccountsPersonCreateIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// The person's gender (International regulations require either "male" or "female").
	LegalGender *string `form:"legal_gender" json:"legal_gender,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The nationalities (countries) this person is associated with.
	Nationalities []*string `form:"nationalities" json:"nationalities,omitempty"`
	// The person token generated by the person token api.
	PersonToken *string `form:"person_token" json:"person_token,omitempty"`
	// The phone number for this person.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The person's political exposure.
	PoliticalExposure *string `form:"political_exposure" json:"political_exposure,omitempty"`
	// The relationship that this person has with the Account's business or legal entity.
	Relationship *V2CoreAccountsPersonCreateRelationshipParams `form:"relationship" json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the person.
	ScriptAddresses *V2CoreAccountsPersonCreateScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The script names (e.g. non-Latin characters) associated with the person.
	ScriptNames *V2CoreAccountsPersonCreateScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The person's last name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountsPersonCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Delete a Person associated with an Account.
type V2CoreAccountsPersonDeleteParams struct {
	Params `form:"*"`
	// The Account the Person is associated with.
	AccountID *string `form:"-" json:"-"` // Included in URL
}

// Retrieves a Person associated with an Account.
type V2CoreAccountsPersonRetrieveParams struct {
	Params `form:"*"`
	// The Account the Person is associated with.
	AccountID *string `form:"-" json:"-"` // Included in URL
}

// Additional addresses associated with the person.
type V2CoreAccountsPersonUpdateAdditionalAddressParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// Purpose of additional address.
	Purpose *string `form:"purpose" json:"purpose"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// Additional names (e.g. aliases) associated with the person.
type V2CoreAccountsPersonUpdateAdditionalNameParams struct {
	// The person's full name.
	FullName *string `form:"full_name" json:"full_name,omitempty"`
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The purpose or type of the additional name.
	Purpose *string `form:"purpose" json:"purpose"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Stripe terms of service agreement.
type V2CoreAccountsPersonUpdateAdditionalTermsOfServiceAccountParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountsPersonUpdateAdditionalTermsOfServiceParams struct {
	// Stripe terms of service agreement.
	Account *V2CoreAccountsPersonUpdateAdditionalTermsOfServiceAccountParams `form:"account" json:"account,omitempty"`
}

// The primary address associated with the person.
type V2CoreAccountsPersonUpdateAddressParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// The person's date of birth.
type V2CoreAccountsPersonUpdateDateOfBirthParams struct {
	// The day of the birth.
	Day *int64 `form:"day" json:"day"`
	// The month of birth.
	Month *int64 `form:"month" json:"month"`
	// The year of birth.
	Year *int64 `form:"year" json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountsPersonUpdateDocumentsCompanyAuthorizationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountsPersonUpdateDocumentsPassportParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountsPersonUpdateDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountsPersonUpdateDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountsPersonUpdateDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountsPersonUpdateDocumentsSecondaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountsPersonUpdateDocumentsSecondaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountsPersonUpdateDocumentsSecondaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountsPersonUpdateDocumentsVisaParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountsPersonUpdateDocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountsPersonUpdateDocumentsCompanyAuthorizationParams `form:"company_authorization" json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountsPersonUpdateDocumentsPassportParams `form:"passport" json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountsPersonUpdateDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountsPersonUpdateDocumentsSecondaryVerificationParams `form:"secondary_verification" json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountsPersonUpdateDocumentsVisaParams `form:"visa" json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the person.
type V2CoreAccountsPersonUpdateIDNumberParams struct {
	// The ID number type of an individual.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// The relationship that this person has with the Account's business or legal entity.
type V2CoreAccountsPersonUpdateRelationshipParams struct {
	// Whether the individual is an authorizer of the Account's legal entity.
	Authorizer *bool `form:"authorizer" json:"authorizer,omitempty"`
	// Indicates whether the person is a director of the associated legal entity.
	Director *bool `form:"director" json:"director,omitempty"`
	// Indicates whether the person is an executive of the associated legal entity.
	Executive *bool `form:"executive" json:"executive,omitempty"`
	// Indicates whether the person is a legal guardian of the associated legal entity.
	LegalGuardian *bool `form:"legal_guardian" json:"legal_guardian,omitempty"`
	// Indicates whether the person is an owner of the associated legal entity.
	Owner *bool `form:"owner" json:"owner,omitempty"`
	// The percentage of ownership the person has in the associated legal entity.
	PercentOwnership *string `form:"percent_ownership" json:"percent_ownership,omitempty"`
	// Indicates whether the person is a representative of the associated legal entity.
	Representative *bool `form:"representative" json:"representative,omitempty"`
	// The title or position the person holds in the associated legal entity.
	Title *string `form:"title" json:"title,omitempty"`
}

// Kana Address.
type V2CoreAccountsPersonUpdateScriptAddressesKanaParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountsPersonUpdateScriptAddressesKanjiParams struct {
	// City, district, suburb, town, or village.
	City *string `form:"city" json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town or cho-me.
	Town *string `form:"town" json:"town,omitempty"`
}

// The script addresses (e.g., non-Latin characters) associated with the person.
type V2CoreAccountsPersonUpdateScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountsPersonUpdateScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountsPersonUpdateScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountsPersonUpdateScriptNamesKanaParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountsPersonUpdateScriptNamesKanjiParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The script names (e.g. non-Latin characters) associated with the person.
type V2CoreAccountsPersonUpdateScriptNamesParams struct {
	// Persons name in kana script.
	Kana *V2CoreAccountsPersonUpdateScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountsPersonUpdateScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Updates a Person associated with an Account.
type V2CoreAccountsPersonUpdateParams struct {
	Params `form:"*"`
	// The Account the Person is associated with.
	AccountID *string `form:"-" json:"-"` // Included in URL
	// Additional addresses associated with the person.
	AdditionalAddresses []*V2CoreAccountsPersonUpdateAdditionalAddressParams `form:"additional_addresses" json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the person.
	AdditionalNames []*V2CoreAccountsPersonUpdateAdditionalNameParams `form:"additional_names" json:"additional_names,omitempty"`
	// Attestations of accepted terms of service agreements.
	AdditionalTermsOfService *V2CoreAccountsPersonUpdateAdditionalTermsOfServiceParams `form:"additional_terms_of_service" json:"additional_terms_of_service,omitempty"`
	// The primary address associated with the person.
	Address *V2CoreAccountsPersonUpdateAddressParams `form:"address" json:"address,omitempty"`
	// The person's date of birth.
	DateOfBirth *V2CoreAccountsPersonUpdateDateOfBirthParams `form:"date_of_birth" json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountsPersonUpdateDocumentsParams `form:"documents" json:"documents,omitempty"`
	// Email.
	Email *string `form:"email" json:"email,omitempty"`
	// The person's first name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The identification numbers (e.g., SSN) associated with the person.
	IDNumbers []*V2CoreAccountsPersonUpdateIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// The person's gender (International regulations require either "male" or "female").
	LegalGender *string `form:"legal_gender" json:"legal_gender,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// The nationalities (countries) this person is associated with.
	Nationalities []*string `form:"nationalities" json:"nationalities,omitempty"`
	// The person token generated by the person token api.
	PersonToken *string `form:"person_token" json:"person_token,omitempty"`
	// The phone number for this person.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The person's political exposure.
	PoliticalExposure *string `form:"political_exposure" json:"political_exposure,omitempty"`
	// The relationship that this person has with the Account's business or legal entity.
	Relationship *V2CoreAccountsPersonUpdateRelationshipParams `form:"relationship" json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the person.
	ScriptAddresses *V2CoreAccountsPersonUpdateScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The script names (e.g. non-Latin characters) associated with the person.
	ScriptNames *V2CoreAccountsPersonUpdateScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The person's last name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountsPersonUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
