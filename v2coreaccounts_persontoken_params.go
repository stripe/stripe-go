//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Additional addresses associated with the person.
type V2CoreAccountsPersonTokenAdditionalAddressParams struct {
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
type V2CoreAccountsPersonTokenAdditionalNameParams struct {
	// The person's full name.
	FullName *string `form:"full_name" json:"full_name,omitempty"`
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The purpose or type of the additional name.
	Purpose *string `form:"purpose" json:"purpose"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Details on the Person's acceptance of the [Stripe Services Agreement]; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountsPersonTokenAdditionalTermsOfServiceAccountParams struct {
	// The boolean value indicating if the terms of service have been accepted.
	ShownAndAccepted *bool `form:"shown_and_accepted" json:"shown_and_accepted,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountsPersonTokenAdditionalTermsOfServiceParams struct {
	// Details on the Person's acceptance of the [Stripe Services Agreement]; IP, date, and User Agent are expanded by Stripe.
	Account *V2CoreAccountsPersonTokenAdditionalTermsOfServiceAccountParams `form:"account" json:"account,omitempty"`
}

// The person's residential address.
type V2CoreAccountsPersonTokenAddressParams struct {
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
type V2CoreAccountsPersonTokenDateOfBirthParams struct {
	// The day of the birth.
	Day *int64 `form:"day" json:"day"`
	// The month of birth.
	Month *int64 `form:"month" json:"month"`
	// The year of birth.
	Year *int64 `form:"year" json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountsPersonTokenDocumentsCompanyAuthorizationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountsPersonTokenDocumentsPassportParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountsPersonTokenDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountsPersonTokenDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountsPersonTokenDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountsPersonTokenDocumentsSecondaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountsPersonTokenDocumentsSecondaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountsPersonTokenDocumentsSecondaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountsPersonTokenDocumentsVisaParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountsPersonTokenDocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountsPersonTokenDocumentsCompanyAuthorizationParams `form:"company_authorization" json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountsPersonTokenDocumentsPassportParams `form:"passport" json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountsPersonTokenDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountsPersonTokenDocumentsSecondaryVerificationParams `form:"secondary_verification" json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountsPersonTokenDocumentsVisaParams `form:"visa" json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the person.
type V2CoreAccountsPersonTokenIDNumberParams struct {
	// The ID number type of an individual.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// The relationship that this person has with the Account's business or legal entity.
type V2CoreAccountsPersonTokenRelationshipParams struct {
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
type V2CoreAccountsPersonTokenScriptAddressesKanaParams struct {
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
type V2CoreAccountsPersonTokenScriptAddressesKanjiParams struct {
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
type V2CoreAccountsPersonTokenScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountsPersonTokenScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountsPersonTokenScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountsPersonTokenScriptNamesKanaParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountsPersonTokenScriptNamesKanjiParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The script names (e.g. non-Latin characters) associated with the person.
type V2CoreAccountsPersonTokenScriptNamesParams struct {
	// Persons name in kana script.
	Kana *V2CoreAccountsPersonTokenScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountsPersonTokenScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Creates a Person Token associated with an Account.
type V2CoreAccountsPersonTokenParams struct {
	Params `form:"*"`
	// The Account the Person is associated with.
	AccountID *string `form:"-" json:"-"` // Included in URL
	// Additional addresses associated with the person.
	AdditionalAddresses []*V2CoreAccountsPersonTokenAdditionalAddressParams `form:"additional_addresses" json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the person.
	AdditionalNames []*V2CoreAccountsPersonTokenAdditionalNameParams `form:"additional_names" json:"additional_names,omitempty"`
	// Attestations of accepted terms of service agreements.
	AdditionalTermsOfService *V2CoreAccountsPersonTokenAdditionalTermsOfServiceParams `form:"additional_terms_of_service" json:"additional_terms_of_service,omitempty"`
	// The person's residential address.
	Address *V2CoreAccountsPersonTokenAddressParams `form:"address" json:"address,omitempty"`
	// The person's date of birth.
	DateOfBirth *V2CoreAccountsPersonTokenDateOfBirthParams `form:"date_of_birth" json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountsPersonTokenDocumentsParams `form:"documents" json:"documents,omitempty"`
	// Email.
	Email *string `form:"email" json:"email,omitempty"`
	// The person's first name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The identification numbers (e.g., SSN) associated with the person.
	IDNumbers []*V2CoreAccountsPersonTokenIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// The person's gender (International regulations require either "male" or "female").
	LegalGender *string `form:"legal_gender" json:"legal_gender,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// The nationalities (countries) this person is associated with.
	Nationalities []*string `form:"nationalities" json:"nationalities,omitempty"`
	// The phone number for this person.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The person's political exposure.
	PoliticalExposure *string `form:"political_exposure" json:"political_exposure,omitempty"`
	// The relationship that this person has with the Account's business or legal entity.
	Relationship *V2CoreAccountsPersonTokenRelationshipParams `form:"relationship" json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the person.
	ScriptAddresses *V2CoreAccountsPersonTokenScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The script names (e.g. non-Latin characters) associated with the person.
	ScriptNames *V2CoreAccountsPersonTokenScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The person's last name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountsPersonTokenParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = &value
}

// Additional addresses associated with the person.
type V2CoreAccountsPersonTokenCreateAdditionalAddressParams struct {
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
type V2CoreAccountsPersonTokenCreateAdditionalNameParams struct {
	// The person's full name.
	FullName *string `form:"full_name" json:"full_name,omitempty"`
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The purpose or type of the additional name.
	Purpose *string `form:"purpose" json:"purpose"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Details on the Person's acceptance of the [Stripe Services Agreement]; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountsPersonTokenCreateAdditionalTermsOfServiceAccountParams struct {
	// The boolean value indicating if the terms of service have been accepted.
	ShownAndAccepted *bool `form:"shown_and_accepted" json:"shown_and_accepted,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountsPersonTokenCreateAdditionalTermsOfServiceParams struct {
	// Details on the Person's acceptance of the [Stripe Services Agreement]; IP, date, and User Agent are expanded by Stripe.
	Account *V2CoreAccountsPersonTokenCreateAdditionalTermsOfServiceAccountParams `form:"account" json:"account,omitempty"`
}

// The person's residential address.
type V2CoreAccountsPersonTokenCreateAddressParams struct {
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
type V2CoreAccountsPersonTokenCreateDateOfBirthParams struct {
	// The day of the birth.
	Day *int64 `form:"day" json:"day"`
	// The month of birth.
	Month *int64 `form:"month" json:"month"`
	// The year of birth.
	Year *int64 `form:"year" json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountsPersonTokenCreateDocumentsCompanyAuthorizationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountsPersonTokenCreateDocumentsPassportParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountsPersonTokenCreateDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountsPersonTokenCreateDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountsPersonTokenCreateDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountsPersonTokenCreateDocumentsSecondaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountsPersonTokenCreateDocumentsSecondaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountsPersonTokenCreateDocumentsSecondaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountsPersonTokenCreateDocumentsVisaParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountsPersonTokenCreateDocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountsPersonTokenCreateDocumentsCompanyAuthorizationParams `form:"company_authorization" json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountsPersonTokenCreateDocumentsPassportParams `form:"passport" json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountsPersonTokenCreateDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountsPersonTokenCreateDocumentsSecondaryVerificationParams `form:"secondary_verification" json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountsPersonTokenCreateDocumentsVisaParams `form:"visa" json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the person.
type V2CoreAccountsPersonTokenCreateIDNumberParams struct {
	// The ID number type of an individual.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// The relationship that this person has with the Account's business or legal entity.
type V2CoreAccountsPersonTokenCreateRelationshipParams struct {
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
type V2CoreAccountsPersonTokenCreateScriptAddressesKanaParams struct {
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
type V2CoreAccountsPersonTokenCreateScriptAddressesKanjiParams struct {
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
type V2CoreAccountsPersonTokenCreateScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountsPersonTokenCreateScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountsPersonTokenCreateScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountsPersonTokenCreateScriptNamesKanaParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountsPersonTokenCreateScriptNamesKanjiParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The script names (e.g. non-Latin characters) associated with the person.
type V2CoreAccountsPersonTokenCreateScriptNamesParams struct {
	// Persons name in kana script.
	Kana *V2CoreAccountsPersonTokenCreateScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountsPersonTokenCreateScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Creates a Person Token associated with an Account.
type V2CoreAccountsPersonTokenCreateParams struct {
	Params `form:"*"`
	// The Account the Person is associated with.
	AccountID *string `form:"-" json:"-"` // Included in URL
	// Additional addresses associated with the person.
	AdditionalAddresses []*V2CoreAccountsPersonTokenCreateAdditionalAddressParams `form:"additional_addresses" json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the person.
	AdditionalNames []*V2CoreAccountsPersonTokenCreateAdditionalNameParams `form:"additional_names" json:"additional_names,omitempty"`
	// Attestations of accepted terms of service agreements.
	AdditionalTermsOfService *V2CoreAccountsPersonTokenCreateAdditionalTermsOfServiceParams `form:"additional_terms_of_service" json:"additional_terms_of_service,omitempty"`
	// The person's residential address.
	Address *V2CoreAccountsPersonTokenCreateAddressParams `form:"address" json:"address,omitempty"`
	// The person's date of birth.
	DateOfBirth *V2CoreAccountsPersonTokenCreateDateOfBirthParams `form:"date_of_birth" json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountsPersonTokenCreateDocumentsParams `form:"documents" json:"documents,omitempty"`
	// Email.
	Email *string `form:"email" json:"email,omitempty"`
	// The person's first name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The identification numbers (e.g., SSN) associated with the person.
	IDNumbers []*V2CoreAccountsPersonTokenCreateIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// The person's gender (International regulations require either "male" or "female").
	LegalGender *string `form:"legal_gender" json:"legal_gender,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// The nationalities (countries) this person is associated with.
	Nationalities []*string `form:"nationalities" json:"nationalities,omitempty"`
	// The phone number for this person.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The person's political exposure.
	PoliticalExposure *string `form:"political_exposure" json:"political_exposure,omitempty"`
	// The relationship that this person has with the Account's business or legal entity.
	Relationship *V2CoreAccountsPersonTokenCreateRelationshipParams `form:"relationship" json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the person.
	ScriptAddresses *V2CoreAccountsPersonTokenCreateScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The script names (e.g. non-Latin characters) associated with the person.
	ScriptNames *V2CoreAccountsPersonTokenCreateScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The person's last name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountsPersonTokenCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = &value
}

// Retrieves a Person Token associated with an Account.
type V2CoreAccountsPersonTokenRetrieveParams struct {
	Params `form:"*"`
	// The Account the Person is associated with.
	AccountID *string `form:"-" json:"-"` // Included in URL
}
