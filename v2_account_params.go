//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of Accounts. Note that the `include` parameter cannot be passed in on list requests.
type V2AccountListParams struct {
	Params `form:"*"`
	// Filter by the configurations that have been applied to the account. If omitted, returns all Accounts regardless of which configurations they have. Currently only supports `recipient`, to filter by Accounts with the recipient configuration set.
	AppliedConfigurations []*string `form:"applied_configurations" json:"applied_configurations,omitempty"`
	// The limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
type V2AccountConfigurationRecipientDataFeaturesBankAccountsLocalParams struct {
	// Whether or not to request the Feature.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
type V2AccountConfigurationRecipientDataFeaturesBankAccountsWireParams struct {
	// Whether or not to request the Feature.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Features that enable OutboundPayments to a bank account linked to this Account.
type V2AccountConfigurationRecipientDataFeaturesBankAccountsParams struct {
	// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
	Local *V2AccountConfigurationRecipientDataFeaturesBankAccountsLocalParams `form:"local" json:"local,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
	Wire *V2AccountConfigurationRecipientDataFeaturesBankAccountsWireParams `form:"wire" json:"wire,omitempty"`
}

// Feature that enable OutboundPayments to a debit card linked to this Account.
type V2AccountConfigurationRecipientDataFeaturesCardsParams struct {
	// Whether or not to request the Feature.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Features representing the functionality that should be enabled for when this Account is used as a recipient. Features need to be explicitly requested, and the `status` field indicates if the Feature is `active`, `restricted` or `pending`. Once a Feature is `active`, the Account can be used with the product flow that the Feature enables.
type V2AccountConfigurationRecipientDataFeaturesParams struct {
	// Features that enable OutboundPayments to a bank account linked to this Account.
	BankAccounts *V2AccountConfigurationRecipientDataFeaturesBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Feature that enable OutboundPayments to a debit card linked to this Account.
	Cards *V2AccountConfigurationRecipientDataFeaturesCardsParams `form:"cards" json:"cards,omitempty"`
}

// Configuration to enable this Account to be used as a recipient of an OutboundPayment via the OutboundPayments API, or via the dashboard.
type V2AccountConfigurationRecipientDataParams struct {
	// The payout method id to be used as a default outbound destination. This will allow the PayoutMethod to be omitted on OutboundPayments made through API or sending payouts via dashboard. Can also be explicitly set to `null` to clear the existing default outbound destination.
	DefaultOutboundDestination *string `form:"default_outbound_destination" json:"default_outbound_destination,omitempty"`
	// Features representing the functionality that should be enabled for when this Account is used as a recipient. Features need to be explicitly requested, and the `status` field indicates if the Feature is `active`, `restricted` or `pending`. Once a Feature is `active`, the Account can be used with the product flow that the Feature enables.
	Features *V2AccountConfigurationRecipientDataFeaturesParams `form:"features" json:"features,omitempty"`
}

// Configurations applied to this Account in order to allow it to be used in different product flows. Currently only supports the recipient configuration.
type V2AccountConfigurationParams struct {
	// Configuration to enable this Account to be used as a recipient of an OutboundPayment via the OutboundPayments API, or via the dashboard.
	RecipientData *V2AccountConfigurationRecipientDataParams `form:"recipient_data" json:"recipient_data,omitempty"`
}

// The address of the Legal Entity.
type V2AccountLegalEntityDataAddressParams struct {
	// City.
	City *string `form:"city" json:"city,omitempty"`
	// Open Enum. Two-letter country code.
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town.
	Town *string `form:"town" json:"town,omitempty"`
}

// The address of the representative.
type V2AccountLegalEntityDataRepresentativeAddressParams struct {
	// City.
	City *string `form:"city" json:"city,omitempty"`
	// Open Enum. Two-letter country code.
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town.
	Town *string `form:"town" json:"town,omitempty"`
}

// The date of birth of the representative.
type V2AccountLegalEntityDataRepresentativeDOBParams struct {
	// The day of birth of the representative.
	Day *int64 `form:"day" json:"day"`
	// The month of birth of the representative.
	Month *int64 `form:"month" json:"month"`
	// The year of birth of the representative.
	Year *int64 `form:"year" json:"year"`
}

// The representative of the Legal Entity. This is the person nominated by the Legal Entity to provide information about themselves, and general information about the account. For legal entities with `individual` business type, this field is required and should contain the individual's information.
type V2AccountLegalEntityDataRepresentativeParams struct {
	// The address of the representative.
	Address *V2AccountLegalEntityDataRepresentativeAddressParams `form:"address" json:"address,omitempty"`
	// The date of birth of the representative.
	DOB *V2AccountLegalEntityDataRepresentativeDOBParams `form:"dob" json:"dob,omitempty"`
	// The given name of the representative.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The surname of the representative.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Information about the company or individual that this Account represents. Stripe may require Legal Entity information in order to enable Features requested on the Account.
type V2AccountLegalEntityDataParams struct {
	// The address of the Legal Entity.
	Address *V2AccountLegalEntityDataAddressParams `form:"address" json:"address,omitempty"`
	// Closed Enum. The business type of the Legal Entity.
	BusinessType *string `form:"business_type" json:"business_type,omitempty"`
	// Open Enum. Two-letter country code (ISO 3166-1 alpha-2) for where the company or individual is located.
	Country *string `form:"country" json:"country,omitempty"`
	// The legal name of this Legal Entity. Required unless the business type is `individual`.
	Name *string `form:"name" json:"name,omitempty"`
	// The representative of the Legal Entity. This is the person nominated by the Legal Entity to provide information about themselves, and general information about the account. For legal entities with `individual` business type, this field is required and should contain the individual's information.
	Representative *V2AccountLegalEntityDataRepresentativeParams `form:"representative" json:"representative,omitempty"`
}

// Creates an Account. You can optionally provide information about the associated Legal Entity, such as name and business type. The Account can also be configured as a recipient of OutboundPayments by requesting Features on the recipient configuration.
type V2AccountParams struct {
	Params `form:"*"`
	// Configurations applied to this Account in order to allow it to be used in different product flows. Currently only supports the recipient configuration.
	Configuration *V2AccountConfigurationParams `form:"configuration" json:"configuration,omitempty"`
	// The default contact email address for the Account. This field is optional, but must be supplied before the recipient configuration can be populated.
	Email *string `form:"email" json:"email,omitempty"`
	// Closed Enum. Additional fields to include in the response. Currently supports `configuration.recipient_data`, `legal_entity_data`, `requirements`, and `supportable_features.recipient_data`.
	Include []*string `form:"include" json:"include,omitempty"`
	// Information about the company or individual that this Account represents. Stripe may require Legal Entity information in order to enable Features requested on the Account.
	LegalEntityData *V2AccountLegalEntityDataParams `form:"legal_entity_data" json:"legal_entity_data,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// A descriptive name for the Account. This name will be surfaced in the Account directory in the dashboard.
	Name *string `form:"name" json:"name,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2AccountParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Closes and removes access to the Account and its associated resources.
type V2AccountCloseParams struct {
	Params `form:"*"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
type V2AccountCreateConfigurationRecipientDataFeaturesBankAccountsLocalParams struct {
	// Whether or not to request the Feature.
	Requested *bool `form:"requested" json:"requested"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
type V2AccountCreateConfigurationRecipientDataFeaturesBankAccountsWireParams struct {
	// Whether or not to request the Feature.
	Requested *bool `form:"requested" json:"requested"`
}

// Features that enable OutboundPayments to a bank account linked to this Account.
type V2AccountCreateConfigurationRecipientDataFeaturesBankAccountsParams struct {
	// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
	Local *V2AccountCreateConfigurationRecipientDataFeaturesBankAccountsLocalParams `form:"local" json:"local,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
	Wire *V2AccountCreateConfigurationRecipientDataFeaturesBankAccountsWireParams `form:"wire" json:"wire,omitempty"`
}

// Feature that enable OutboundPayments to a debit card linked to this Account.
type V2AccountCreateConfigurationRecipientDataFeaturesCardsParams struct {
	// Whether or not to request the Feature.
	Requested *bool `form:"requested" json:"requested"`
}

// Features representing the functionality that should be enabled for when this Account is used as a recipient. Features need to be explicitly requested, and the `status` field indicates if the Feature is `active`, `restricted` or `pending`. Once a Feature is `active`, the Account can be used with the product flow that the Feature enables.
type V2AccountCreateConfigurationRecipientDataFeaturesParams struct {
	// Features that enable OutboundPayments to a bank account linked to this Account.
	BankAccounts *V2AccountCreateConfigurationRecipientDataFeaturesBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Feature that enable OutboundPayments to a debit card linked to this Account.
	Cards *V2AccountCreateConfigurationRecipientDataFeaturesCardsParams `form:"cards" json:"cards,omitempty"`
}

// Configuration to enable this Account to be used as a recipient of an OutboundPayment via the OutboundPayments API, or via the dashboard.
type V2AccountCreateConfigurationRecipientDataParams struct {
	// Features representing the functionality that should be enabled for when this Account is used as a recipient. Features need to be explicitly requested, and the `status` field indicates if the Feature is `active`, `restricted` or `pending`. Once a Feature is `active`, the Account can be used with the product flow that the Feature enables.
	Features *V2AccountCreateConfigurationRecipientDataFeaturesParams `form:"features" json:"features"`
}

// Configurations applied to this Account in order to allow it to be used in different product flows. Currently only supports the recipient configuration.
type V2AccountCreateConfigurationParams struct {
	// Configuration to enable this Account to be used as a recipient of an OutboundPayment via the OutboundPayments API, or via the dashboard.
	RecipientData *V2AccountCreateConfigurationRecipientDataParams `form:"recipient_data" json:"recipient_data,omitempty"`
}

// The address of the Legal Entity.
type V2AccountCreateLegalEntityDataAddressParams struct {
	// City.
	City *string `form:"city" json:"city,omitempty"`
	// Open Enum. Two-letter country code.
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town.
	Town *string `form:"town" json:"town,omitempty"`
}

// The address of the representative.
type V2AccountCreateLegalEntityDataRepresentativeAddressParams struct {
	// City.
	City *string `form:"city" json:"city,omitempty"`
	// Open Enum. Two-letter country code.
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town.
	Town *string `form:"town" json:"town,omitempty"`
}

// The date of birth of the representative.
type V2AccountCreateLegalEntityDataRepresentativeDOBParams struct {
	// The day of birth of the representative.
	Day *int64 `form:"day" json:"day"`
	// The month of birth of the representative.
	Month *int64 `form:"month" json:"month"`
	// The year of birth of the representative.
	Year *int64 `form:"year" json:"year"`
}

// The representative of the Legal Entity. This is the person nominated by the Legal Entity to provide information about themselves, and general information about the account. For legal entities with `individual` business type, this field is required and should contain the individual's information.
type V2AccountCreateLegalEntityDataRepresentativeParams struct {
	// The address of the representative.
	Address *V2AccountCreateLegalEntityDataRepresentativeAddressParams `form:"address" json:"address,omitempty"`
	// The date of birth of the representative.
	DOB *V2AccountCreateLegalEntityDataRepresentativeDOBParams `form:"dob" json:"dob,omitempty"`
	// The given name of the representative.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The surname of the representative.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Information about the company or individual that this Account represents. Stripe may require Legal Entity information in order to enable Features requested on the Account.
type V2AccountCreateLegalEntityDataParams struct {
	// The address of the Legal Entity.
	Address *V2AccountCreateLegalEntityDataAddressParams `form:"address" json:"address,omitempty"`
	// Closed Enum. The business type of the Legal Entity.
	BusinessType *string `form:"business_type" json:"business_type,omitempty"`
	// Open Enum. Two-letter country code (ISO 3166-1 alpha-2) for where the company or individual is located.
	Country *string `form:"country" json:"country"`
	// The legal name of this Legal Entity. Required unless the business type is `individual`.
	Name *string `form:"name" json:"name,omitempty"`
	// The representative of the Legal Entity. This is the person nominated by the Legal Entity to provide information about themselves, and general information about the account. For legal entities with `individual` business type, this field is required and should contain the individual's information.
	Representative *V2AccountCreateLegalEntityDataRepresentativeParams `form:"representative" json:"representative,omitempty"`
}

// Creates an Account. You can optionally provide information about the associated Legal Entity, such as name and business type. The Account can also be configured as a recipient of OutboundPayments by requesting Features on the recipient configuration.
type V2AccountCreateParams struct {
	Params `form:"*"`
	// Configurations applied to this Account in order to allow it to be used in different product flows. Currently only supports the recipient configuration.
	Configuration *V2AccountCreateConfigurationParams `form:"configuration" json:"configuration,omitempty"`
	// The default contact email address for the Account. This field is optional, but must be supplied before the recipient configuration can be populated.
	Email *string `form:"email" json:"email,omitempty"`
	// Closed Enum. Additional fields to include in the response. Currently supports `configuration.recipient_data`, `legal_entity_data`, `requirements`, and `supportable_features.recipient_data`.
	Include []*string `form:"include" json:"include,omitempty"`
	// Information about the company or individual that this Account represents. Stripe may require Legal Entity information in order to enable Features requested on the Account.
	LegalEntityData *V2AccountCreateLegalEntityDataParams `form:"legal_entity_data" json:"legal_entity_data,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// A descriptive name for the Account. This name will be surfaced in the Account directory in the dashboard.
	Name *string `form:"name" json:"name,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2AccountCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of an Account.
type V2AccountRetrieveParams struct {
	Params `form:"*"`
	// Closed Enum. Additional fields to include in the response. Currently supports `configuration.recipient_data`, `legal_entity_data`, `requirements`, and `supportable_features.recipient_data`.
	Include []*string `form:"include" json:"include,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
type V2AccountUpdateConfigurationRecipientDataFeaturesBankAccountsLocalParams struct {
	// Whether or not to request the Feature.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
type V2AccountUpdateConfigurationRecipientDataFeaturesBankAccountsWireParams struct {
	// Whether or not to request the Feature.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Features that enable OutboundPayments to a bank account linked to this Account.
type V2AccountUpdateConfigurationRecipientDataFeaturesBankAccountsParams struct {
	// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
	Local *V2AccountUpdateConfigurationRecipientDataFeaturesBankAccountsLocalParams `form:"local" json:"local,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
	Wire *V2AccountUpdateConfigurationRecipientDataFeaturesBankAccountsWireParams `form:"wire" json:"wire,omitempty"`
}

// Feature that enable OutboundPayments to a debit card linked to this Account.
type V2AccountUpdateConfigurationRecipientDataFeaturesCardsParams struct {
	// Whether or not to request the Feature.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Features representing the functionality that should be enabled for when this Account is used as a recipient. Features need to be explicitly requested, and the `status` field indicates if the Feature is `active`, `restricted` or `pending`. Once a Feature is `active`, the Account can be used with the product flow that the Feature enables.
type V2AccountUpdateConfigurationRecipientDataFeaturesParams struct {
	// Features that enable OutboundPayments to a bank account linked to this Account.
	BankAccounts *V2AccountUpdateConfigurationRecipientDataFeaturesBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Feature that enable OutboundPayments to a debit card linked to this Account.
	Cards *V2AccountUpdateConfigurationRecipientDataFeaturesCardsParams `form:"cards" json:"cards,omitempty"`
}

// Configuration to enable this Account to be used as a recipient of an OutboundPayment via the OutboundPayments API, or via the dashboard.
type V2AccountUpdateConfigurationRecipientDataParams struct {
	// The payout method id to be used as a default outbound destination. This will allow the PayoutMethod to be omitted on OutboundPayments made through API or sending payouts via dashboard. Can also be explicitly set to `null` to clear the existing default outbound destination.
	DefaultOutboundDestination *string `form:"default_outbound_destination" json:"default_outbound_destination,omitempty"`
	// Features representing the functionality that should be enabled for when this Account is used as a recipient. Features need to be explicitly requested, and the `status` field indicates if the Feature is `active`, `restricted` or `pending`. Once a Feature is `active`, the Account can be used with the product flow that the Feature enables.
	Features *V2AccountUpdateConfigurationRecipientDataFeaturesParams `form:"features" json:"features,omitempty"`
}

// Configurations applied to this Account in order to allow it to be used in different product flows. Currently only supports the recipient configuration.
type V2AccountUpdateConfigurationParams struct {
	// Configuration to enable this Account to be used as a recipient of an OutboundPayment via the OutboundPayments API, or via the dashboard.
	RecipientData *V2AccountUpdateConfigurationRecipientDataParams `form:"recipient_data" json:"recipient_data,omitempty"`
}

// The address of the Legal Entity.
type V2AccountUpdateLegalEntityDataAddressParams struct {
	// City.
	City *string `form:"city" json:"city,omitempty"`
	// Open Enum. Two-letter country code.
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town.
	Town *string `form:"town" json:"town,omitempty"`
}

// The address of the representative.
type V2AccountUpdateLegalEntityDataRepresentativeAddressParams struct {
	// City.
	City *string `form:"city" json:"city,omitempty"`
	// Open Enum. Two-letter country code.
	Country *string `form:"country" json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `form:"line1" json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `form:"line2" json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `form:"postal_code" json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `form:"state" json:"state,omitempty"`
	// Town.
	Town *string `form:"town" json:"town,omitempty"`
}

// The date of birth of the representative.
type V2AccountUpdateLegalEntityDataRepresentativeDOBParams struct {
	// The day of birth of the representative.
	Day *int64 `form:"day" json:"day"`
	// The month of birth of the representative.
	Month *int64 `form:"month" json:"month"`
	// The year of birth of the representative.
	Year *int64 `form:"year" json:"year"`
}

// The representative of the Legal Entity. This is the person nominated by the Legal Entity to provide information about themselves, and general information about the account. For legal entities with `individual` business type, this field is required and should contain the individual's information.
type V2AccountUpdateLegalEntityDataRepresentativeParams struct {
	// The address of the representative.
	Address *V2AccountUpdateLegalEntityDataRepresentativeAddressParams `form:"address" json:"address,omitempty"`
	// The date of birth of the representative.
	DOB *V2AccountUpdateLegalEntityDataRepresentativeDOBParams `form:"dob" json:"dob,omitempty"`
	// The given name of the representative.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The surname of the representative.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Information about the company or individual that this Account represents. Stripe may require Legal Entity information in order to enable Features requested on the Account.
type V2AccountUpdateLegalEntityDataParams struct {
	// The address of the Legal Entity.
	Address *V2AccountUpdateLegalEntityDataAddressParams `form:"address" json:"address,omitempty"`
	// Closed Enum. The business type of the Legal Entity.
	BusinessType *string `form:"business_type" json:"business_type,omitempty"`
	// Open Enum. Two-letter country code (ISO 3166-1 alpha-2) for where the company or individual is located.
	Country *string `form:"country" json:"country,omitempty"`
	// The legal name of this Legal Entity. Required unless the business type is `individual`.
	Name *string `form:"name" json:"name,omitempty"`
	// The representative of the Legal Entity. This is the person nominated by the Legal Entity to provide information about themselves, and general information about the account. For legal entities with `individual` business type, this field is required and should contain the individual's information.
	Representative *V2AccountUpdateLegalEntityDataRepresentativeParams `form:"representative" json:"representative,omitempty"`
}

// Updates the details of an Account. You can also optionally provide or update the details of the associated Legal Entity and recipient configuration.
type V2AccountUpdateParams struct {
	Params `form:"*"`
	// Configurations applied to this Account in order to allow it to be used in different product flows. Currently only supports the recipient configuration.
	Configuration *V2AccountUpdateConfigurationParams `form:"configuration" json:"configuration,omitempty"`
	// The default contact email address for the Account. This field is optional, but must be supplied before the recipient configuration can be populated.
	Email *string `form:"email" json:"email,omitempty"`
	// Closed Enum. Additional fields to include in the response. Currently supports `configuration.recipient_data`, `legal_entity_data`, `requirements`, and `supportable_features.recipient_data`.
	Include []*string `form:"include" json:"include,omitempty"`
	// Information about the company or individual that this Account represents. Stripe may require Legal Entity information in order to enable Features requested on the Account.
	LegalEntityData *V2AccountUpdateLegalEntityDataParams `form:"legal_entity_data" json:"legal_entity_data,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// A descriptive name for the Account. This name will be surfaced in the Account directory in the dashboard.
	Name *string `form:"name" json:"name,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2AccountUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
