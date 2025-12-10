//
//
// File generated from our OpenAPI spec
//
//

package stripe

// This hash is used to attest that the directors information provided to Stripe is both current and correct; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountTokenIdentityAttestationsDirectorshipDeclarationParams struct {
	// A boolean indicating if the directors information has been attested.
	Attested *bool `form:"attested" json:"attested,omitempty"`
}

// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountTokenIdentityAttestationsOwnershipDeclarationParams struct {
	// A boolean indicating if the beneficial owner information has been attested.
	Attested *bool `form:"attested" json:"attested,omitempty"`
}

// Attestation that all Persons with a specific Relationship value have been provided.
type V2CoreAccountTokenIdentityAttestationsPersonsProvidedParams struct {
	// Whether the company's directors have been provided. Set this Boolean to true after creating all the company's directors with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Directors *bool `form:"directors" json:"directors,omitempty"`
	// Whether the company's executives have been provided. Set this Boolean to true after creating all the company's executives with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Executives *bool `form:"executives" json:"executives,omitempty"`
	// Whether the company's owners have been provided. Set this Boolean to true after creating all the company's owners with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Owners *bool `form:"owners" json:"owners,omitempty"`
	// Reason for why the company is exempt from providing ownership information.
	OwnershipExemptionReason *string `form:"ownership_exemption_reason" json:"ownership_exemption_reason,omitempty"`
}

// This hash is used to attest that the representative is authorized to act as the representative of their legal entity; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountTokenIdentityAttestationsRepresentativeDeclarationParams struct {
	// A boolean indicating if the representative is authorized to act as the representative of their legal entity.
	Attested *bool `form:"attested" json:"attested,omitempty"`
}

// Details on the Account's acceptance of the [Stripe Services Agreement]; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountTokenIdentityAttestationsTermsOfServiceAccountParams struct {
	// The boolean value indicating if the terms of service have been accepted.
	ShownAndAccepted *bool `form:"shown_and_accepted" json:"shown_and_accepted,omitempty"`
}

// Details on the Account's acceptance of Treasury-specific terms of service; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountTokenIdentityAttestationsTermsOfServiceStorerParams struct {
	// The boolean value indicating if the terms of service have been accepted.
	ShownAndAccepted *bool `form:"shown_and_accepted" json:"shown_and_accepted,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountTokenIdentityAttestationsTermsOfServiceParams struct {
	// Details on the Account's acceptance of the [Stripe Services Agreement]; IP, date, and User Agent are expanded by Stripe.
	Account *V2CoreAccountTokenIdentityAttestationsTermsOfServiceAccountParams `form:"account" json:"account,omitempty"`
	// Details on the Account's acceptance of Treasury-specific terms of service; IP, date, and User Agent are expanded by Stripe.
	Storer *V2CoreAccountTokenIdentityAttestationsTermsOfServiceStorerParams `form:"storer" json:"storer,omitempty"`
}

// Attestations from the identity's key people, e.g. owners, executives, directors, representatives.
type V2CoreAccountTokenIdentityAttestationsParams struct {
	// This hash is used to attest that the directors information provided to Stripe is both current and correct; IP, date, and User Agent are expanded by Stripe.
	DirectorshipDeclaration *V2CoreAccountTokenIdentityAttestationsDirectorshipDeclarationParams `form:"directorship_declaration" json:"directorship_declaration,omitempty"`
	// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct; IP, date, and User Agent are expanded by Stripe.
	OwnershipDeclaration *V2CoreAccountTokenIdentityAttestationsOwnershipDeclarationParams `form:"ownership_declaration" json:"ownership_declaration,omitempty"`
	// Attestation that all Persons with a specific Relationship value have been provided.
	PersonsProvided *V2CoreAccountTokenIdentityAttestationsPersonsProvidedParams `form:"persons_provided" json:"persons_provided,omitempty"`
	// This hash is used to attest that the representative is authorized to act as the representative of their legal entity; IP, date, and User Agent are expanded by Stripe.
	RepresentativeDeclaration *V2CoreAccountTokenIdentityAttestationsRepresentativeDeclarationParams `form:"representative_declaration" json:"representative_declaration,omitempty"`
	// Attestations of accepted terms of service agreements.
	TermsOfService *V2CoreAccountTokenIdentityAttestationsTermsOfServiceParams `form:"terms_of_service" json:"terms_of_service,omitempty"`
}

// The business registration address of the business entity.
type V2CoreAccountTokenIdentityBusinessDetailsAddressParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// A non-negative integer representing the amount in the smallest currency unit.
type V2CoreAccountTokenIdentityBusinessDetailsAnnualRevenueAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// The business gross annual revenue for its preceding fiscal year.
type V2CoreAccountTokenIdentityBusinessDetailsAnnualRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *V2CoreAccountTokenIdentityBusinessDetailsAnnualRevenueAmountParams `form:"amount" json:"amount,omitempty"`
	// The close-out date of the preceding fiscal year in ISO 8601 format. E.g. 2023-12-31 for the 31st of December, 2023.
	FiscalYearEnd *string `form:"fiscal_year_end" json:"fiscal_year_end,omitempty"`
}

// One or more documents that support the bank account ownership verification requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a check.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsBankAccountOwnershipVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's license to operate.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsCompanyLicenseParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the company's Memorandum of Association.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsCompanyMemorandumOfAssociationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Certain countries only: One or more documents showing the ministerial decree legalizing the company's establishment.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsCompanyMinisterialDecreeParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsCompanyRegistrationVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's tax ID.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsCompanyTaxIDVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document verifying the business.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountTokenIdentityBusinessDetailsDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of address.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsProofOfAddressParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the company's proof of registration with the national business registry.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsProofOfRegistrationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of ultimate beneficial ownership.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// A document verifying the business.
type V2CoreAccountTokenIdentityBusinessDetailsDocumentsParams struct {
	// One or more documents that support the bank account ownership verification requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a check.
	BankAccountOwnershipVerification *V2CoreAccountTokenIdentityBusinessDetailsDocumentsBankAccountOwnershipVerificationParams `form:"bank_account_ownership_verification" json:"bank_account_ownership_verification,omitempty"`
	// One or more documents that demonstrate proof of a company's license to operate.
	CompanyLicense *V2CoreAccountTokenIdentityBusinessDetailsDocumentsCompanyLicenseParams `form:"company_license" json:"company_license,omitempty"`
	// One or more documents showing the company's Memorandum of Association.
	CompanyMemorandumOfAssociation *V2CoreAccountTokenIdentityBusinessDetailsDocumentsCompanyMemorandumOfAssociationParams `form:"company_memorandum_of_association" json:"company_memorandum_of_association,omitempty"`
	// Certain countries only: One or more documents showing the ministerial decree legalizing the company's establishment.
	CompanyMinisterialDecree *V2CoreAccountTokenIdentityBusinessDetailsDocumentsCompanyMinisterialDecreeParams `form:"company_ministerial_decree" json:"company_ministerial_decree,omitempty"`
	// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
	CompanyRegistrationVerification *V2CoreAccountTokenIdentityBusinessDetailsDocumentsCompanyRegistrationVerificationParams `form:"company_registration_verification" json:"company_registration_verification,omitempty"`
	// One or more documents that demonstrate proof of a company's tax ID.
	CompanyTaxIDVerification *V2CoreAccountTokenIdentityBusinessDetailsDocumentsCompanyTaxIDVerificationParams `form:"company_tax_id_verification" json:"company_tax_id_verification,omitempty"`
	// A document verifying the business.
	PrimaryVerification *V2CoreAccountTokenIdentityBusinessDetailsDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// One or more documents that demonstrate proof of address.
	ProofOfAddress *V2CoreAccountTokenIdentityBusinessDetailsDocumentsProofOfAddressParams `form:"proof_of_address" json:"proof_of_address,omitempty"`
	// One or more documents showing the company's proof of registration with the national business registry.
	ProofOfRegistration *V2CoreAccountTokenIdentityBusinessDetailsDocumentsProofOfRegistrationParams `form:"proof_of_registration" json:"proof_of_registration,omitempty"`
	// One or more documents that demonstrate proof of ultimate beneficial ownership.
	ProofOfUltimateBeneficialOwnership *V2CoreAccountTokenIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams `form:"proof_of_ultimate_beneficial_ownership" json:"proof_of_ultimate_beneficial_ownership,omitempty"`
}

// The ID numbers of a business entity.
type V2CoreAccountTokenIdentityBusinessDetailsIDNumberParams struct {
	// The registrar of the ID number (Only valid for DE ID number types).
	Registrar *string `form:"registrar" json:"registrar,omitempty"`
	// Open Enum. The ID number type of a business entity.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// A non-negative integer representing the amount in the smallest currency unit.
type V2CoreAccountTokenIdentityBusinessDetailsMonthlyEstimatedRevenueAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// An estimate of the monthly revenue of the business.
type V2CoreAccountTokenIdentityBusinessDetailsMonthlyEstimatedRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *V2CoreAccountTokenIdentityBusinessDetailsMonthlyEstimatedRevenueAmountParams `form:"amount" json:"amount,omitempty"`
}

// Kana Address.
type V2CoreAccountTokenIdentityBusinessDetailsScriptAddressesKanaParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountTokenIdentityBusinessDetailsScriptAddressesKanjiParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// The business registration address of the business entity in non latin script.
type V2CoreAccountTokenIdentityBusinessDetailsScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountTokenIdentityBusinessDetailsScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountTokenIdentityBusinessDetailsScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Kana name.
type V2CoreAccountTokenIdentityBusinessDetailsScriptNamesKanaParams struct {
	// Registered name of the business.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// Kanji name.
type V2CoreAccountTokenIdentityBusinessDetailsScriptNamesKanjiParams struct {
	// Registered name of the business.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// The business legal name in non latin script.
type V2CoreAccountTokenIdentityBusinessDetailsScriptNamesParams struct {
	// Kana name.
	Kana *V2CoreAccountTokenIdentityBusinessDetailsScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji name.
	Kanji *V2CoreAccountTokenIdentityBusinessDetailsScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Information about the company or business.
type V2CoreAccountTokenIdentityBusinessDetailsParams struct {
	// The business registration address of the business entity.
	Address *V2CoreAccountTokenIdentityBusinessDetailsAddressParams `form:"address" json:"address,omitempty"`
	// The business gross annual revenue for its preceding fiscal year.
	AnnualRevenue *V2CoreAccountTokenIdentityBusinessDetailsAnnualRevenueParams `form:"annual_revenue" json:"annual_revenue,omitempty"`
	// A document verifying the business.
	Documents *V2CoreAccountTokenIdentityBusinessDetailsDocumentsParams `form:"documents" json:"documents,omitempty"`
	// Estimated maximum number of workers currently engaged by the business (including employees, contractors, and vendors).
	EstimatedWorkerCount *int64 `form:"estimated_worker_count" json:"estimated_worker_count,omitempty"`
	// The ID numbers of a business entity.
	IDNumbers []*V2CoreAccountTokenIdentityBusinessDetailsIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// An estimate of the monthly revenue of the business.
	MonthlyEstimatedRevenue *V2CoreAccountTokenIdentityBusinessDetailsMonthlyEstimatedRevenueParams `form:"monthly_estimated_revenue" json:"monthly_estimated_revenue,omitempty"`
	// The phone number of the Business Entity.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The business legal name.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
	// The business registration address of the business entity in non latin script.
	ScriptAddresses *V2CoreAccountTokenIdentityBusinessDetailsScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The business legal name in non latin script.
	ScriptNames *V2CoreAccountTokenIdentityBusinessDetailsScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The category identifying the legal structure of the business.
	Structure *string `form:"structure" json:"structure,omitempty"`
}

// Additional addresses associated with the individual.
type V2CoreAccountTokenIdentityIndividualAdditionalAddressParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// Additional names (e.g. aliases) associated with the individual.
type V2CoreAccountTokenIdentityIndividualAdditionalNameParams struct {
	// The person's full name.
	FullName *string `form:"full_name" json:"full_name,omitempty"`
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The purpose or type of the additional name.
	Purpose *string `form:"purpose" json:"purpose"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The individual's residential address.
type V2CoreAccountTokenIdentityIndividualAddressParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// The individual's date of birth.
type V2CoreAccountTokenIdentityIndividualDateOfBirthParams struct {
	// The day of the birth.
	Day *int64 `form:"day" json:"day"`
	// The month of birth.
	Month *int64 `form:"month" json:"month"`
	// The year of birth.
	Year *int64 `form:"year" json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountTokenIdentityIndividualDocumentsCompanyAuthorizationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountTokenIdentityIndividualDocumentsPassportParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountTokenIdentityIndividualDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountTokenIdentityIndividualDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountTokenIdentityIndividualDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountTokenIdentityIndividualDocumentsSecondaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountTokenIdentityIndividualDocumentsSecondaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountTokenIdentityIndividualDocumentsSecondaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountTokenIdentityIndividualDocumentsVisaParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountTokenIdentityIndividualDocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountTokenIdentityIndividualDocumentsCompanyAuthorizationParams `form:"company_authorization" json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountTokenIdentityIndividualDocumentsPassportParams `form:"passport" json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountTokenIdentityIndividualDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountTokenIdentityIndividualDocumentsSecondaryVerificationParams `form:"secondary_verification" json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountTokenIdentityIndividualDocumentsVisaParams `form:"visa" json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the individual.
type V2CoreAccountTokenIdentityIndividualIDNumberParams struct {
	// The ID number type of an individual.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// The relationship that this individual has with the account's identity.
type V2CoreAccountTokenIdentityIndividualRelationshipParams struct {
	// Whether the person is a director of the account's identity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.
	Director *bool `form:"director" json:"director,omitempty"`
	// Whether the person has significant responsibility to control, manage, or direct the organization.
	Executive *bool `form:"executive" json:"executive,omitempty"`
	// Whether the person is an owner of the account's identity.
	Owner *bool `form:"owner" json:"owner,omitempty"`
	// The percent owned by the person of the account's legal entity.
	PercentOwnership *string `form:"percent_ownership" json:"percent_ownership,omitempty"`
	// The person's title (e.g., CEO, Support Engineer).
	Title *string `form:"title" json:"title,omitempty"`
}

// Kana Address.
type V2CoreAccountTokenIdentityIndividualScriptAddressesKanaParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountTokenIdentityIndividualScriptAddressesKanjiParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// The script addresses (e.g., non-Latin characters) associated with the individual.
type V2CoreAccountTokenIdentityIndividualScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountTokenIdentityIndividualScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountTokenIdentityIndividualScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountTokenIdentityIndividualScriptNamesKanaParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountTokenIdentityIndividualScriptNamesKanjiParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The individuals primary name in non latin script.
type V2CoreAccountTokenIdentityIndividualScriptNamesParams struct {
	// Persons name in kana script.
	Kana *V2CoreAccountTokenIdentityIndividualScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountTokenIdentityIndividualScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Information about the person represented by the account.
type V2CoreAccountTokenIdentityIndividualParams struct {
	// Additional addresses associated with the individual.
	AdditionalAddresses []*V2CoreAccountTokenIdentityIndividualAdditionalAddressParams `form:"additional_addresses" json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the individual.
	AdditionalNames []*V2CoreAccountTokenIdentityIndividualAdditionalNameParams `form:"additional_names" json:"additional_names,omitempty"`
	// The individual's residential address.
	Address *V2CoreAccountTokenIdentityIndividualAddressParams `form:"address" json:"address,omitempty"`
	// The individual's date of birth.
	DateOfBirth *V2CoreAccountTokenIdentityIndividualDateOfBirthParams `form:"date_of_birth" json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountTokenIdentityIndividualDocumentsParams `form:"documents" json:"documents,omitempty"`
	// The individual's email address.
	Email *string `form:"email" json:"email,omitempty"`
	// The individual's first name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The identification numbers (e.g., SSN) associated with the individual.
	IDNumbers []*V2CoreAccountTokenIdentityIndividualIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// The individual's gender (International regulations require either "male" or "female").
	LegalGender *string `form:"legal_gender" json:"legal_gender,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// The countries where the individual is a national. Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Nationalities []*string `form:"nationalities" json:"nationalities,omitempty"`
	// The individual's phone number.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The individual's political exposure.
	PoliticalExposure *string `form:"political_exposure" json:"political_exposure,omitempty"`
	// The relationship that this individual has with the account's identity.
	Relationship *V2CoreAccountTokenIdentityIndividualRelationshipParams `form:"relationship" json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the individual.
	ScriptAddresses *V2CoreAccountTokenIdentityIndividualScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The individuals primary name in non latin script.
	ScriptNames *V2CoreAccountTokenIdentityIndividualScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The individual's last name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountTokenIdentityIndividualParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Information about the company, individual, and business represented by the Account.
type V2CoreAccountTokenIdentityParams struct {
	// Attestations from the identity's key people, e.g. owners, executives, directors, representatives.
	Attestations *V2CoreAccountTokenIdentityAttestationsParams `form:"attestations" json:"attestations,omitempty"`
	// Information about the company or business.
	BusinessDetails *V2CoreAccountTokenIdentityBusinessDetailsParams `form:"business_details" json:"business_details,omitempty"`
	// The entity type.
	EntityType *string `form:"entity_type" json:"entity_type,omitempty"`
	// Information about the person represented by the account.
	Individual *V2CoreAccountTokenIdentityIndividualParams `form:"individual" json:"individual,omitempty"`
}

// Creates an Account Token.
type V2CoreAccountTokenParams struct {
	Params `form:"*"`
	// The default contact email address for the Account. Required when configuring the account as a merchant or recipient.
	ContactEmail *string `form:"contact_email" json:"contact_email,omitempty"`
	// A descriptive name for the Account. This name will be surfaced in the Stripe Dashboard and on any invoices sent to the Account.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Information about the company, individual, and business represented by the Account.
	Identity *V2CoreAccountTokenIdentityParams `form:"identity" json:"identity,omitempty"`
}

// This hash is used to attest that the directors information provided to Stripe is both current and correct; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountTokenCreateIdentityAttestationsDirectorshipDeclarationParams struct {
	// A boolean indicating if the directors information has been attested.
	Attested *bool `form:"attested" json:"attested,omitempty"`
}

// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountTokenCreateIdentityAttestationsOwnershipDeclarationParams struct {
	// A boolean indicating if the beneficial owner information has been attested.
	Attested *bool `form:"attested" json:"attested,omitempty"`
}

// Attestation that all Persons with a specific Relationship value have been provided.
type V2CoreAccountTokenCreateIdentityAttestationsPersonsProvidedParams struct {
	// Whether the company's directors have been provided. Set this Boolean to true after creating all the company's directors with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Directors *bool `form:"directors" json:"directors,omitempty"`
	// Whether the company's executives have been provided. Set this Boolean to true after creating all the company's executives with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Executives *bool `form:"executives" json:"executives,omitempty"`
	// Whether the company's owners have been provided. Set this Boolean to true after creating all the company's owners with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Owners *bool `form:"owners" json:"owners,omitempty"`
	// Reason for why the company is exempt from providing ownership information.
	OwnershipExemptionReason *string `form:"ownership_exemption_reason" json:"ownership_exemption_reason,omitempty"`
}

// This hash is used to attest that the representative is authorized to act as the representative of their legal entity; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountTokenCreateIdentityAttestationsRepresentativeDeclarationParams struct {
	// A boolean indicating if the representative is authorized to act as the representative of their legal entity.
	Attested *bool `form:"attested" json:"attested,omitempty"`
}

// Details on the Account's acceptance of the [Stripe Services Agreement]; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountTokenCreateIdentityAttestationsTermsOfServiceAccountParams struct {
	// The boolean value indicating if the terms of service have been accepted.
	ShownAndAccepted *bool `form:"shown_and_accepted" json:"shown_and_accepted,omitempty"`
}

// Details on the Account's acceptance of Treasury-specific terms of service; IP, date, and User Agent are expanded by Stripe.
type V2CoreAccountTokenCreateIdentityAttestationsTermsOfServiceStorerParams struct {
	// The boolean value indicating if the terms of service have been accepted.
	ShownAndAccepted *bool `form:"shown_and_accepted" json:"shown_and_accepted,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountTokenCreateIdentityAttestationsTermsOfServiceParams struct {
	// Details on the Account's acceptance of the [Stripe Services Agreement]; IP, date, and User Agent are expanded by Stripe.
	Account *V2CoreAccountTokenCreateIdentityAttestationsTermsOfServiceAccountParams `form:"account" json:"account,omitempty"`
	// Details on the Account's acceptance of Treasury-specific terms of service; IP, date, and User Agent are expanded by Stripe.
	Storer *V2CoreAccountTokenCreateIdentityAttestationsTermsOfServiceStorerParams `form:"storer" json:"storer,omitempty"`
}

// Attestations from the identity's key people, e.g. owners, executives, directors, representatives.
type V2CoreAccountTokenCreateIdentityAttestationsParams struct {
	// This hash is used to attest that the directors information provided to Stripe is both current and correct; IP, date, and User Agent are expanded by Stripe.
	DirectorshipDeclaration *V2CoreAccountTokenCreateIdentityAttestationsDirectorshipDeclarationParams `form:"directorship_declaration" json:"directorship_declaration,omitempty"`
	// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct; IP, date, and User Agent are expanded by Stripe.
	OwnershipDeclaration *V2CoreAccountTokenCreateIdentityAttestationsOwnershipDeclarationParams `form:"ownership_declaration" json:"ownership_declaration,omitempty"`
	// Attestation that all Persons with a specific Relationship value have been provided.
	PersonsProvided *V2CoreAccountTokenCreateIdentityAttestationsPersonsProvidedParams `form:"persons_provided" json:"persons_provided,omitempty"`
	// This hash is used to attest that the representative is authorized to act as the representative of their legal entity; IP, date, and User Agent are expanded by Stripe.
	RepresentativeDeclaration *V2CoreAccountTokenCreateIdentityAttestationsRepresentativeDeclarationParams `form:"representative_declaration" json:"representative_declaration,omitempty"`
	// Attestations of accepted terms of service agreements.
	TermsOfService *V2CoreAccountTokenCreateIdentityAttestationsTermsOfServiceParams `form:"terms_of_service" json:"terms_of_service,omitempty"`
}

// The business registration address of the business entity.
type V2CoreAccountTokenCreateIdentityBusinessDetailsAddressParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// A non-negative integer representing the amount in the smallest currency unit.
type V2CoreAccountTokenCreateIdentityBusinessDetailsAnnualRevenueAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// The business gross annual revenue for its preceding fiscal year.
type V2CoreAccountTokenCreateIdentityBusinessDetailsAnnualRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *V2CoreAccountTokenCreateIdentityBusinessDetailsAnnualRevenueAmountParams `form:"amount" json:"amount,omitempty"`
	// The close-out date of the preceding fiscal year in ISO 8601 format. E.g. 2023-12-31 for the 31st of December, 2023.
	FiscalYearEnd *string `form:"fiscal_year_end" json:"fiscal_year_end,omitempty"`
}

// One or more documents that support the bank account ownership verification requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a check.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsBankAccountOwnershipVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's license to operate.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsCompanyLicenseParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the company's Memorandum of Association.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsCompanyMemorandumOfAssociationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Certain countries only: One or more documents showing the ministerial decree legalizing the company's establishment.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsCompanyMinisterialDecreeParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsCompanyRegistrationVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's tax ID.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsCompanyTaxIDVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document verifying the business.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of address.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsProofOfAddressParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the company's proof of registration with the national business registry.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsProofOfRegistrationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of ultimate beneficial ownership.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// A document verifying the business.
type V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsParams struct {
	// One or more documents that support the bank account ownership verification requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a check.
	BankAccountOwnershipVerification *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsBankAccountOwnershipVerificationParams `form:"bank_account_ownership_verification" json:"bank_account_ownership_verification,omitempty"`
	// One or more documents that demonstrate proof of a company's license to operate.
	CompanyLicense *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsCompanyLicenseParams `form:"company_license" json:"company_license,omitempty"`
	// One or more documents showing the company's Memorandum of Association.
	CompanyMemorandumOfAssociation *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsCompanyMemorandumOfAssociationParams `form:"company_memorandum_of_association" json:"company_memorandum_of_association,omitempty"`
	// Certain countries only: One or more documents showing the ministerial decree legalizing the company's establishment.
	CompanyMinisterialDecree *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsCompanyMinisterialDecreeParams `form:"company_ministerial_decree" json:"company_ministerial_decree,omitempty"`
	// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
	CompanyRegistrationVerification *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsCompanyRegistrationVerificationParams `form:"company_registration_verification" json:"company_registration_verification,omitempty"`
	// One or more documents that demonstrate proof of a company's tax ID.
	CompanyTaxIDVerification *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsCompanyTaxIDVerificationParams `form:"company_tax_id_verification" json:"company_tax_id_verification,omitempty"`
	// A document verifying the business.
	PrimaryVerification *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// One or more documents that demonstrate proof of address.
	ProofOfAddress *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsProofOfAddressParams `form:"proof_of_address" json:"proof_of_address,omitempty"`
	// One or more documents showing the company's proof of registration with the national business registry.
	ProofOfRegistration *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsProofOfRegistrationParams `form:"proof_of_registration" json:"proof_of_registration,omitempty"`
	// One or more documents that demonstrate proof of ultimate beneficial ownership.
	ProofOfUltimateBeneficialOwnership *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams `form:"proof_of_ultimate_beneficial_ownership" json:"proof_of_ultimate_beneficial_ownership,omitempty"`
}

// The ID numbers of a business entity.
type V2CoreAccountTokenCreateIdentityBusinessDetailsIDNumberParams struct {
	// The registrar of the ID number (Only valid for DE ID number types).
	Registrar *string `form:"registrar" json:"registrar,omitempty"`
	// Open Enum. The ID number type of a business entity.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// A non-negative integer representing the amount in the smallest currency unit.
type V2CoreAccountTokenCreateIdentityBusinessDetailsMonthlyEstimatedRevenueAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// An estimate of the monthly revenue of the business.
type V2CoreAccountTokenCreateIdentityBusinessDetailsMonthlyEstimatedRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *V2CoreAccountTokenCreateIdentityBusinessDetailsMonthlyEstimatedRevenueAmountParams `form:"amount" json:"amount,omitempty"`
}

// Kana Address.
type V2CoreAccountTokenCreateIdentityBusinessDetailsScriptAddressesKanaParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountTokenCreateIdentityBusinessDetailsScriptAddressesKanjiParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// The business registration address of the business entity in non latin script.
type V2CoreAccountTokenCreateIdentityBusinessDetailsScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountTokenCreateIdentityBusinessDetailsScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountTokenCreateIdentityBusinessDetailsScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Kana name.
type V2CoreAccountTokenCreateIdentityBusinessDetailsScriptNamesKanaParams struct {
	// Registered name of the business.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// Kanji name.
type V2CoreAccountTokenCreateIdentityBusinessDetailsScriptNamesKanjiParams struct {
	// Registered name of the business.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// The business legal name in non latin script.
type V2CoreAccountTokenCreateIdentityBusinessDetailsScriptNamesParams struct {
	// Kana name.
	Kana *V2CoreAccountTokenCreateIdentityBusinessDetailsScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji name.
	Kanji *V2CoreAccountTokenCreateIdentityBusinessDetailsScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Information about the company or business.
type V2CoreAccountTokenCreateIdentityBusinessDetailsParams struct {
	// The business registration address of the business entity.
	Address *V2CoreAccountTokenCreateIdentityBusinessDetailsAddressParams `form:"address" json:"address,omitempty"`
	// The business gross annual revenue for its preceding fiscal year.
	AnnualRevenue *V2CoreAccountTokenCreateIdentityBusinessDetailsAnnualRevenueParams `form:"annual_revenue" json:"annual_revenue,omitempty"`
	// A document verifying the business.
	Documents *V2CoreAccountTokenCreateIdentityBusinessDetailsDocumentsParams `form:"documents" json:"documents,omitempty"`
	// Estimated maximum number of workers currently engaged by the business (including employees, contractors, and vendors).
	EstimatedWorkerCount *int64 `form:"estimated_worker_count" json:"estimated_worker_count,omitempty"`
	// The ID numbers of a business entity.
	IDNumbers []*V2CoreAccountTokenCreateIdentityBusinessDetailsIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// An estimate of the monthly revenue of the business.
	MonthlyEstimatedRevenue *V2CoreAccountTokenCreateIdentityBusinessDetailsMonthlyEstimatedRevenueParams `form:"monthly_estimated_revenue" json:"monthly_estimated_revenue,omitempty"`
	// The phone number of the Business Entity.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The business legal name.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
	// The business registration address of the business entity in non latin script.
	ScriptAddresses *V2CoreAccountTokenCreateIdentityBusinessDetailsScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The business legal name in non latin script.
	ScriptNames *V2CoreAccountTokenCreateIdentityBusinessDetailsScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The category identifying the legal structure of the business.
	Structure *string `form:"structure" json:"structure,omitempty"`
}

// Additional addresses associated with the individual.
type V2CoreAccountTokenCreateIdentityIndividualAdditionalAddressParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// Additional names (e.g. aliases) associated with the individual.
type V2CoreAccountTokenCreateIdentityIndividualAdditionalNameParams struct {
	// The person's full name.
	FullName *string `form:"full_name" json:"full_name,omitempty"`
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The purpose or type of the additional name.
	Purpose *string `form:"purpose" json:"purpose"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The individual's residential address.
type V2CoreAccountTokenCreateIdentityIndividualAddressParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// The individual's date of birth.
type V2CoreAccountTokenCreateIdentityIndividualDateOfBirthParams struct {
	// The day of the birth.
	Day *int64 `form:"day" json:"day"`
	// The month of birth.
	Month *int64 `form:"month" json:"month"`
	// The year of birth.
	Year *int64 `form:"year" json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountTokenCreateIdentityIndividualDocumentsCompanyAuthorizationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountTokenCreateIdentityIndividualDocumentsPassportParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountTokenCreateIdentityIndividualDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountTokenCreateIdentityIndividualDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountTokenCreateIdentityIndividualDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountTokenCreateIdentityIndividualDocumentsSecondaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountTokenCreateIdentityIndividualDocumentsSecondaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountTokenCreateIdentityIndividualDocumentsSecondaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountTokenCreateIdentityIndividualDocumentsVisaParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountTokenCreateIdentityIndividualDocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountTokenCreateIdentityIndividualDocumentsCompanyAuthorizationParams `form:"company_authorization" json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountTokenCreateIdentityIndividualDocumentsPassportParams `form:"passport" json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountTokenCreateIdentityIndividualDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountTokenCreateIdentityIndividualDocumentsSecondaryVerificationParams `form:"secondary_verification" json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountTokenCreateIdentityIndividualDocumentsVisaParams `form:"visa" json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the individual.
type V2CoreAccountTokenCreateIdentityIndividualIDNumberParams struct {
	// The ID number type of an individual.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// The relationship that this individual has with the account's identity.
type V2CoreAccountTokenCreateIdentityIndividualRelationshipParams struct {
	// Whether the person is a director of the account's identity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.
	Director *bool `form:"director" json:"director,omitempty"`
	// Whether the person has significant responsibility to control, manage, or direct the organization.
	Executive *bool `form:"executive" json:"executive,omitempty"`
	// Whether the person is an owner of the account's identity.
	Owner *bool `form:"owner" json:"owner,omitempty"`
	// The percent owned by the person of the account's legal entity.
	PercentOwnership *string `form:"percent_ownership" json:"percent_ownership,omitempty"`
	// The person's title (e.g., CEO, Support Engineer).
	Title *string `form:"title" json:"title,omitempty"`
}

// Kana Address.
type V2CoreAccountTokenCreateIdentityIndividualScriptAddressesKanaParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountTokenCreateIdentityIndividualScriptAddressesKanjiParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// The script addresses (e.g., non-Latin characters) associated with the individual.
type V2CoreAccountTokenCreateIdentityIndividualScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountTokenCreateIdentityIndividualScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountTokenCreateIdentityIndividualScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountTokenCreateIdentityIndividualScriptNamesKanaParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountTokenCreateIdentityIndividualScriptNamesKanjiParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The individuals primary name in non latin script.
type V2CoreAccountTokenCreateIdentityIndividualScriptNamesParams struct {
	// Persons name in kana script.
	Kana *V2CoreAccountTokenCreateIdentityIndividualScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountTokenCreateIdentityIndividualScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Information about the person represented by the account.
type V2CoreAccountTokenCreateIdentityIndividualParams struct {
	// Additional addresses associated with the individual.
	AdditionalAddresses []*V2CoreAccountTokenCreateIdentityIndividualAdditionalAddressParams `form:"additional_addresses" json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the individual.
	AdditionalNames []*V2CoreAccountTokenCreateIdentityIndividualAdditionalNameParams `form:"additional_names" json:"additional_names,omitempty"`
	// The individual's residential address.
	Address *V2CoreAccountTokenCreateIdentityIndividualAddressParams `form:"address" json:"address,omitempty"`
	// The individual's date of birth.
	DateOfBirth *V2CoreAccountTokenCreateIdentityIndividualDateOfBirthParams `form:"date_of_birth" json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountTokenCreateIdentityIndividualDocumentsParams `form:"documents" json:"documents,omitempty"`
	// The individual's email address.
	Email *string `form:"email" json:"email,omitempty"`
	// The individual's first name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The identification numbers (e.g., SSN) associated with the individual.
	IDNumbers []*V2CoreAccountTokenCreateIdentityIndividualIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// The individual's gender (International regulations require either "male" or "female").
	LegalGender *string `form:"legal_gender" json:"legal_gender,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// The countries where the individual is a national. Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Nationalities []*string `form:"nationalities" json:"nationalities,omitempty"`
	// The individual's phone number.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The individual's political exposure.
	PoliticalExposure *string `form:"political_exposure" json:"political_exposure,omitempty"`
	// The relationship that this individual has with the account's identity.
	Relationship *V2CoreAccountTokenCreateIdentityIndividualRelationshipParams `form:"relationship" json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the individual.
	ScriptAddresses *V2CoreAccountTokenCreateIdentityIndividualScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The individuals primary name in non latin script.
	ScriptNames *V2CoreAccountTokenCreateIdentityIndividualScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The individual's last name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountTokenCreateIdentityIndividualParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Information about the company, individual, and business represented by the Account.
type V2CoreAccountTokenCreateIdentityParams struct {
	// Attestations from the identity's key people, e.g. owners, executives, directors, representatives.
	Attestations *V2CoreAccountTokenCreateIdentityAttestationsParams `form:"attestations" json:"attestations,omitempty"`
	// Information about the company or business.
	BusinessDetails *V2CoreAccountTokenCreateIdentityBusinessDetailsParams `form:"business_details" json:"business_details,omitempty"`
	// The entity type.
	EntityType *string `form:"entity_type" json:"entity_type,omitempty"`
	// Information about the person represented by the account.
	Individual *V2CoreAccountTokenCreateIdentityIndividualParams `form:"individual" json:"individual,omitempty"`
}

// Creates an Account Token.
type V2CoreAccountTokenCreateParams struct {
	Params `form:"*"`
	// The default contact email address for the Account. Required when configuring the account as a merchant or recipient.
	ContactEmail *string `form:"contact_email" json:"contact_email,omitempty"`
	// A descriptive name for the Account. This name will be surfaced in the Stripe Dashboard and on any invoices sent to the Account.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Information about the company, individual, and business represented by the Account.
	Identity *V2CoreAccountTokenCreateIdentityParams `form:"identity" json:"identity,omitempty"`
}

// Retrieves an Account Token.
type V2CoreAccountTokenRetrieveParams struct {
	Params `form:"*"`
}
