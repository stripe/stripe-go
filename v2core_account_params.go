//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Returns a list of Accounts.
type V2CoreAccountListParams struct {
	Params `form:"*"`
	// Filter only accounts that have all of the configurations specified. If omitted, returns all accounts regardless of which configurations they have.
	AppliedConfigurations []*string `form:"applied_configurations" json:"applied_configurations,omitempty"`
	// Filter by whether the account is closed. If omitted, returns only Accounts that are not closed.
	Closed *bool `form:"closed" json:"closed,omitempty"`
	// The upper limit on the number of accounts returned by the List Account request.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Celtic as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticParams struct {
	// Can create commercial issuing charge cards with Celtic as BIN sponsor.
	ChargeCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing spend cards with Celtic as BIN sponsor.
	SpendCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams struct {
	// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
	ChargeCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
	// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Fifth Third as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialFifthThirdParams struct {
	// Can create commercial issuing charge cards with Fifth Third as BIN sponsor.
	ChargeCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Stripe as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadParams struct {
	// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Stripe as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing prepaid cards with Stripe as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Stripe as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeParams struct {
	// Can create commercial issuing charge cards with Stripe as BIN sponsor.
	ChargeCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing prepaid cards with Stripe as BIN sponsor.
	PrepaidCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Can create cards for commercial issuing use cases.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialParams struct {
	// Can create commercial issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Can create commercial issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Can create commercial issuing cards with Fifth Third as BIN sponsor.
	FifthThird *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialFifthThirdParams `form:"fifth_third" json:"fifth_third,omitempty"`
	// Can create commercial issuing cards with Stripe as BIN sponsor.
	Lead *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadParams `form:"lead" json:"lead,omitempty"`
	// Can create commercial issuing cards with Stripe as BIN sponsor.
	Stripe *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeParams `form:"stripe" json:"stripe,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create consumer issuing cards with Celtic as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCelticParams struct {
	// Can create consumer issuing revolving credit cards with Celtic as BIN sponsor.
	RevolvingCreditCard *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardParams `form:"revolving_credit_card" json:"revolving_credit_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create consumer issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankParams struct {
	// Can create consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create consumer issuing cards with Lead as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadParams struct {
	// Can create consumer issuing debit cards with Lead as BIN sponsor.
	DebitCard *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardParams `form:"debit_card" json:"debit_card,omitempty"`
	// Can create consumer issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Can create cards for consumer issuing use cases.
type V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerParams struct {
	// Can create consumer issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Can create consumer issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Can create consumer issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadParams `form:"lead" json:"lead,omitempty"`
}

// Capabilities to request on the CardCreator Configuration.
type V2CoreAccountConfigurationCardCreatorCapabilitiesParams struct {
	// Can create cards for commercial issuing use cases.
	Commercial *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialParams `form:"commercial" json:"commercial,omitempty"`
	// Can create cards for consumer issuing use cases.
	Consumer *V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerParams `form:"consumer" json:"consumer,omitempty"`
}

// The CardCreator Configuration allows the Account to create and issue cards to users.
type V2CoreAccountConfigurationCardCreatorParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Capabilities to request on the CardCreator Configuration.
	Capabilities *V2CoreAccountConfigurationCardCreatorCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
}

// Automatic indirect tax settings to be used when automatic tax calculation is enabled on the customer's invoices, subscriptions, checkout sessions, or payment links. Surfaces if automatic tax calculation is possible given the current customer location information.
type V2CoreAccountConfigurationCustomerAutomaticIndirectTaxParams struct {
	// Describes the customer's tax exemption status, which is `none`, `exempt`, or `reverse`. When set to reverse, invoice and receipt PDFs include the following text: “Reverse charge”.
	Exempt *string `form:"exempt" json:"exempt,omitempty"`
	// A recent IP address of the customer used for tax reporting and tax location inference.
	IPAddress *string `form:"ip_address" json:"ip_address,omitempty"`
	// The data source used to identify the customer's tax location - defaults to `identity_address`. Will only be used for automatic tax calculation on the customer's Invoices and Subscriptions. This behavior is now deprecated for new users.
	LocationSource *string `form:"location_source" json:"location_source,omitempty"`
	// A per-request flag that indicates when Stripe should validate the customer tax location - defaults to `auto`.
	ValidateLocation *string `form:"validate_location" json:"validate_location,omitempty"`
}

// The list of up to 4 default custom fields to be displayed on invoices for this customer.
type V2CoreAccountConfigurationCustomerBillingInvoiceCustomFieldParams struct {
	// The name of the custom field. This may be up to 40 characters.
	Name *string `form:"name" json:"name"`
	// The value of the custom field. This may be up to 140 characters. When updating, pass an empty string to remove previously-defined values.
	Value *string `form:"value" json:"value"`
}

// Default invoice PDF rendering options.
type V2CoreAccountConfigurationCustomerBillingInvoiceRenderingParams struct {
	// Indicates whether displayed line item prices and amounts on invoice PDFs include inclusive tax amounts. Must be either `include_inclusive_tax` or `exclude_tax`.
	AmountTaxDisplay *string `form:"amount_tax_display" json:"amount_tax_display,omitempty"`
	// ID of the invoice rendering template to use for future invoices.
	Template *string `form:"template" json:"template,omitempty"`
}

// Default invoice settings for the customer account.
type V2CoreAccountConfigurationCustomerBillingInvoiceParams struct {
	// The list of up to 4 default custom fields to be displayed on invoices for this customer.
	CustomFields []*V2CoreAccountConfigurationCustomerBillingInvoiceCustomFieldParams `form:"custom_fields" json:"custom_fields,omitempty"`
	// Default invoice footer.
	Footer *string `form:"footer" json:"footer,omitempty"`
	// Sequence number to use on the customer account's next invoice. Defaults to 1.
	NextSequence *int64 `form:"next_sequence" json:"next_sequence,omitempty"`
	// Prefix used to generate unique invoice numbers. Must be 3-12 uppercase letters or numbers.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
	// Default invoice PDF rendering options.
	Rendering *V2CoreAccountConfigurationCustomerBillingInvoiceRenderingParams `form:"rendering" json:"rendering,omitempty"`
}

// Billing settings - default settings used for this customer in Billing flows such as Invoices and Subscriptions.
type V2CoreAccountConfigurationCustomerBillingParams struct {
	// The ID of a `PaymentMethod` attached to this Account's `customer` configuration, used as the default payment method for invoices and subscriptions.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// Default invoice settings for the customer account.
	Invoice *V2CoreAccountConfigurationCustomerBillingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Generates requirements for enabling automatic indirect tax calculation on this customer's invoices or subscriptions. Recommended to request this capability if planning to enable automatic tax calculation on this customer's invoices or subscriptions.
type V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTaxParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities that have been requested on the Customer Configuration.
type V2CoreAccountConfigurationCustomerCapabilitiesParams struct {
	// Generates requirements for enabling automatic indirect tax calculation on this customer's invoices or subscriptions. Recommended to request this capability if planning to enable automatic tax calculation on this customer's invoices or subscriptions.
	AutomaticIndirectTax *V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTaxParams `form:"automatic_indirect_tax" json:"automatic_indirect_tax,omitempty"`
}

// The customer's shipping information. Appears on invoices emailed to this customer.
type V2CoreAccountConfigurationCustomerShippingParams struct {
	// Customer shipping address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Customer name.
	Name *string `form:"name" json:"name,omitempty"`
	// Customer phone (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// The Customer Configuration allows the Account to be used in inbound payment flows (i.e. customer-facing payment and billing flows).
type V2CoreAccountConfigurationCustomerParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Automatic indirect tax settings to be used when automatic tax calculation is enabled on the customer's invoices, subscriptions, checkout sessions, or payment links. Surfaces if automatic tax calculation is possible given the current customer location information.
	AutomaticIndirectTax *V2CoreAccountConfigurationCustomerAutomaticIndirectTaxParams `form:"automatic_indirect_tax" json:"automatic_indirect_tax,omitempty"`
	// Billing settings - default settings used for this customer in Billing flows such as Invoices and Subscriptions.
	Billing *V2CoreAccountConfigurationCustomerBillingParams `form:"billing" json:"billing,omitempty"`
	// Capabilities that have been requested on the Customer Configuration.
	Capabilities *V2CoreAccountConfigurationCustomerCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// The customer's shipping information. Appears on invoices emailed to this customer.
	Shipping *V2CoreAccountConfigurationCustomerShippingParams `form:"shipping" json:"shipping,omitempty"`
	// ID of the test clock to attach to the customer. Can only be set on testmode Accounts, and when the Customer Configuration is first set on an Account.
	TestClock *string `form:"test_clock" json:"test_clock,omitempty"`
}

// Settings used for Bacs debit payments.
type V2CoreAccountConfigurationMerchantBACSDebitPaymentsParams struct {
	// Display name for Bacs Direct Debit payments.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
}

// Settings used to apply the merchant's branding to email receipts, invoices, Checkout, and other products.
type V2CoreAccountConfigurationMerchantBrandingParams struct {
	// ID of a [file upload](https://docs.stripe.com/api/persons/update#create_file): An icon for the merchant. Must be square and at least 128px x 128px.
	Icon *string `form:"icon" json:"icon,omitempty"`
	// ID of a [file upload](https://docs.stripe.com/api/persons/update#create_file): A logo for the merchant that will be used in Checkout instead of the icon and without the merchant's name next to it if provided. Must be at least 128px x 128px.
	Logo *string `form:"logo" json:"logo,omitempty"`
	// A CSS hex color value representing the primary branding color for the merchant.
	PrimaryColor *string `form:"primary_color" json:"primary_color,omitempty"`
	// A CSS hex color value representing the secondary branding color for the merchant.
	SecondaryColor *string `form:"secondary_color" json:"secondary_color,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process ACH debit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesACHDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process ACSS debit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesACSSDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Affirm payments.
type V2CoreAccountConfigurationMerchantCapabilitiesAffirmPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Afterpay/Clearpay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Alma payments.
type V2CoreAccountConfigurationMerchantCapabilitiesAlmaPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Amazon Pay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Australian BECS Direct Debit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesAUBECSDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process BACS Direct Debit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesBACSDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Bancontact payments.
type V2CoreAccountConfigurationMerchantCapabilitiesBancontactPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process BLIK payments.
type V2CoreAccountConfigurationMerchantCapabilitiesBLIKPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Boleto payments.
type V2CoreAccountConfigurationMerchantCapabilitiesBoletoPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesCardPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesCardPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesCardPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to collect card payments.
type V2CoreAccountConfigurationMerchantCapabilitiesCardPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesCardPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Cartes Bancaires payments.
type V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Cash App payments.
type V2CoreAccountConfigurationMerchantCapabilitiesCashAppPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesEPSPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesEPSPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesEPSPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process EPS payments.
type V2CoreAccountConfigurationMerchantCapabilitiesEPSPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesEPSPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesFPXPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesFPXPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesFPXPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process FPX payments.
type V2CoreAccountConfigurationMerchantCapabilitiesFPXPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesFPXPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process UK bank transfer payments.
type V2CoreAccountConfigurationMerchantCapabilitiesGBBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process GrabPay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process iDEAL payments.
type V2CoreAccountConfigurationMerchantCapabilitiesIDEALPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesJCBPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesJCBPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesJCBPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process JCB card payments.
type V2CoreAccountConfigurationMerchantCapabilitiesJCBPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesJCBPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Japanese bank transfer payments.
type V2CoreAccountConfigurationMerchantCapabilitiesJPBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Kakao Pay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Klarna payments.
type V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Konbini convenience store payments.
type V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Korean card payments.
type V2CoreAccountConfigurationMerchantCapabilitiesKrCardPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesLinkPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesLinkPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesLinkPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Link payments.
type V2CoreAccountConfigurationMerchantCapabilitiesLinkPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesLinkPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process MobilePay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Multibanco payments.
type V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Mexican bank transfer payments.
type V2CoreAccountConfigurationMerchantCapabilitiesMXBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Naver Pay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process OXXO payments.
type V2CoreAccountConfigurationMerchantCapabilitiesOXXOPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesP24PaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesP24PaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesP24PaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Przelewy24 (P24) payments.
type V2CoreAccountConfigurationMerchantCapabilitiesP24PaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesP24PaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Pay by Bank payments.
type V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process PAYCO payments.
type V2CoreAccountConfigurationMerchantCapabilitiesPaycoPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process PayNow payments.
type V2CoreAccountConfigurationMerchantCapabilitiesPayNowPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process PromptPay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesPromptPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Revolut Pay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Samsung Pay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process SEPA bank transfer payments.
type V2CoreAccountConfigurationMerchantCapabilitiesSEPABankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process SEPA Direct Debit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesSEPADebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Sunbit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesSunbitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesSwishPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesSwishPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesSwishPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Swish payments.
type V2CoreAccountConfigurationMerchantCapabilitiesSwishPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesSwishPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process TWINT payments.
type V2CoreAccountConfigurationMerchantCapabilitiesTWINTPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process US bank transfer payments.
type V2CoreAccountConfigurationMerchantCapabilitiesUSBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMerchantCapabilitiesZipPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMerchantCapabilitiesZipPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMerchantCapabilitiesZipPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Zip payments.
type V2CoreAccountConfigurationMerchantCapabilitiesZipPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMerchantCapabilitiesZipPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities to request on the Merchant Configuration.
type V2CoreAccountConfigurationMerchantCapabilitiesParams struct {
	// Allow the merchant to process ACH debit payments.
	ACHDebitPayments *V2CoreAccountConfigurationMerchantCapabilitiesACHDebitPaymentsParams `form:"ach_debit_payments" json:"ach_debit_payments,omitempty"`
	// Allow the merchant to process ACSS debit payments.
	ACSSDebitPayments *V2CoreAccountConfigurationMerchantCapabilitiesACSSDebitPaymentsParams `form:"acss_debit_payments" json:"acss_debit_payments,omitempty"`
	// Allow the merchant to process Affirm payments.
	AffirmPayments *V2CoreAccountConfigurationMerchantCapabilitiesAffirmPaymentsParams `form:"affirm_payments" json:"affirm_payments,omitempty"`
	// Allow the merchant to process Afterpay/Clearpay payments.
	AfterpayClearpayPayments *V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsParams `form:"afterpay_clearpay_payments" json:"afterpay_clearpay_payments,omitempty"`
	// Allow the merchant to process Alma payments.
	AlmaPayments *V2CoreAccountConfigurationMerchantCapabilitiesAlmaPaymentsParams `form:"alma_payments" json:"alma_payments,omitempty"`
	// Allow the merchant to process Amazon Pay payments.
	AmazonPayPayments *V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPaymentsParams `form:"amazon_pay_payments" json:"amazon_pay_payments,omitempty"`
	// Allow the merchant to process Australian BECS Direct Debit payments.
	AUBECSDebitPayments *V2CoreAccountConfigurationMerchantCapabilitiesAUBECSDebitPaymentsParams `form:"au_becs_debit_payments" json:"au_becs_debit_payments,omitempty"`
	// Allow the merchant to process BACS Direct Debit payments.
	BACSDebitPayments *V2CoreAccountConfigurationMerchantCapabilitiesBACSDebitPaymentsParams `form:"bacs_debit_payments" json:"bacs_debit_payments,omitempty"`
	// Allow the merchant to process Bancontact payments.
	BancontactPayments *V2CoreAccountConfigurationMerchantCapabilitiesBancontactPaymentsParams `form:"bancontact_payments" json:"bancontact_payments,omitempty"`
	// Allow the merchant to process BLIK payments.
	BLIKPayments *V2CoreAccountConfigurationMerchantCapabilitiesBLIKPaymentsParams `form:"blik_payments" json:"blik_payments,omitempty"`
	// Allow the merchant to process Boleto payments.
	BoletoPayments *V2CoreAccountConfigurationMerchantCapabilitiesBoletoPaymentsParams `form:"boleto_payments" json:"boleto_payments,omitempty"`
	// Allow the merchant to collect card payments.
	CardPayments *V2CoreAccountConfigurationMerchantCapabilitiesCardPaymentsParams `form:"card_payments" json:"card_payments,omitempty"`
	// Allow the merchant to process Cartes Bancaires payments.
	CartesBancairesPayments *V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPaymentsParams `form:"cartes_bancaires_payments" json:"cartes_bancaires_payments,omitempty"`
	// Allow the merchant to process Cash App payments.
	CashAppPayments *V2CoreAccountConfigurationMerchantCapabilitiesCashAppPaymentsParams `form:"cashapp_payments" json:"cashapp_payments,omitempty"`
	// Allow the merchant to process EPS payments.
	EPSPayments *V2CoreAccountConfigurationMerchantCapabilitiesEPSPaymentsParams `form:"eps_payments" json:"eps_payments,omitempty"`
	// Allow the merchant to process FPX payments.
	FPXPayments *V2CoreAccountConfigurationMerchantCapabilitiesFPXPaymentsParams `form:"fpx_payments" json:"fpx_payments,omitempty"`
	// Allow the merchant to process UK bank transfer payments.
	GBBankTransferPayments *V2CoreAccountConfigurationMerchantCapabilitiesGBBankTransferPaymentsParams `form:"gb_bank_transfer_payments" json:"gb_bank_transfer_payments,omitempty"`
	// Allow the merchant to process GrabPay payments.
	GrabpayPayments *V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPaymentsParams `form:"grabpay_payments" json:"grabpay_payments,omitempty"`
	// Allow the merchant to process iDEAL payments.
	IDEALPayments *V2CoreAccountConfigurationMerchantCapabilitiesIDEALPaymentsParams `form:"ideal_payments" json:"ideal_payments,omitempty"`
	// Allow the merchant to process JCB card payments.
	JCBPayments *V2CoreAccountConfigurationMerchantCapabilitiesJCBPaymentsParams `form:"jcb_payments" json:"jcb_payments,omitempty"`
	// Allow the merchant to process Japanese bank transfer payments.
	JPBankTransferPayments *V2CoreAccountConfigurationMerchantCapabilitiesJPBankTransferPaymentsParams `form:"jp_bank_transfer_payments" json:"jp_bank_transfer_payments,omitempty"`
	// Allow the merchant to process Kakao Pay payments.
	KakaoPayPayments *V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPaymentsParams `form:"kakao_pay_payments" json:"kakao_pay_payments,omitempty"`
	// Allow the merchant to process Klarna payments.
	KlarnaPayments *V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPaymentsParams `form:"klarna_payments" json:"klarna_payments,omitempty"`
	// Allow the merchant to process Konbini convenience store payments.
	KonbiniPayments *V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPaymentsParams `form:"konbini_payments" json:"konbini_payments,omitempty"`
	// Allow the merchant to process Korean card payments.
	KrCardPayments *V2CoreAccountConfigurationMerchantCapabilitiesKrCardPaymentsParams `form:"kr_card_payments" json:"kr_card_payments,omitempty"`
	// Allow the merchant to process Link payments.
	LinkPayments *V2CoreAccountConfigurationMerchantCapabilitiesLinkPaymentsParams `form:"link_payments" json:"link_payments,omitempty"`
	// Allow the merchant to process MobilePay payments.
	MobilepayPayments *V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPaymentsParams `form:"mobilepay_payments" json:"mobilepay_payments,omitempty"`
	// Allow the merchant to process Multibanco payments.
	MultibancoPayments *V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPaymentsParams `form:"multibanco_payments" json:"multibanco_payments,omitempty"`
	// Allow the merchant to process Mexican bank transfer payments.
	MXBankTransferPayments *V2CoreAccountConfigurationMerchantCapabilitiesMXBankTransferPaymentsParams `form:"mx_bank_transfer_payments" json:"mx_bank_transfer_payments,omitempty"`
	// Allow the merchant to process Naver Pay payments.
	NaverPayPayments *V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPaymentsParams `form:"naver_pay_payments" json:"naver_pay_payments,omitempty"`
	// Allow the merchant to process OXXO payments.
	OXXOPayments *V2CoreAccountConfigurationMerchantCapabilitiesOXXOPaymentsParams `form:"oxxo_payments" json:"oxxo_payments,omitempty"`
	// Allow the merchant to process Przelewy24 (P24) payments.
	P24Payments *V2CoreAccountConfigurationMerchantCapabilitiesP24PaymentsParams `form:"p24_payments" json:"p24_payments,omitempty"`
	// Allow the merchant to process Pay by Bank payments.
	PayByBankPayments *V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPaymentsParams `form:"pay_by_bank_payments" json:"pay_by_bank_payments,omitempty"`
	// Allow the merchant to process PAYCO payments.
	PaycoPayments *V2CoreAccountConfigurationMerchantCapabilitiesPaycoPaymentsParams `form:"payco_payments" json:"payco_payments,omitempty"`
	// Allow the merchant to process PayNow payments.
	PayNowPayments *V2CoreAccountConfigurationMerchantCapabilitiesPayNowPaymentsParams `form:"paynow_payments" json:"paynow_payments,omitempty"`
	// Allow the merchant to process PromptPay payments.
	PromptPayPayments *V2CoreAccountConfigurationMerchantCapabilitiesPromptPayPaymentsParams `form:"promptpay_payments" json:"promptpay_payments,omitempty"`
	// Allow the merchant to process Revolut Pay payments.
	RevolutPayPayments *V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPaymentsParams `form:"revolut_pay_payments" json:"revolut_pay_payments,omitempty"`
	// Allow the merchant to process Samsung Pay payments.
	SamsungPayPayments *V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPaymentsParams `form:"samsung_pay_payments" json:"samsung_pay_payments,omitempty"`
	// Allow the merchant to process SEPA bank transfer payments.
	SEPABankTransferPayments *V2CoreAccountConfigurationMerchantCapabilitiesSEPABankTransferPaymentsParams `form:"sepa_bank_transfer_payments" json:"sepa_bank_transfer_payments,omitempty"`
	// Allow the merchant to process SEPA Direct Debit payments.
	SEPADebitPayments *V2CoreAccountConfigurationMerchantCapabilitiesSEPADebitPaymentsParams `form:"sepa_debit_payments" json:"sepa_debit_payments,omitempty"`
	// Allow the merchant to process Sunbit payments.
	SunbitPayments *V2CoreAccountConfigurationMerchantCapabilitiesSunbitPaymentsParams `form:"sunbit_payments" json:"sunbit_payments,omitempty"`
	// Allow the merchant to process Swish payments.
	SwishPayments *V2CoreAccountConfigurationMerchantCapabilitiesSwishPaymentsParams `form:"swish_payments" json:"swish_payments,omitempty"`
	// Allow the merchant to process TWINT payments.
	TWINTPayments *V2CoreAccountConfigurationMerchantCapabilitiesTWINTPaymentsParams `form:"twint_payments" json:"twint_payments,omitempty"`
	// Allow the merchant to process US bank transfer payments.
	USBankTransferPayments *V2CoreAccountConfigurationMerchantCapabilitiesUSBankTransferPaymentsParams `form:"us_bank_transfer_payments" json:"us_bank_transfer_payments,omitempty"`
	// Allow the merchant to process Zip payments.
	ZipPayments *V2CoreAccountConfigurationMerchantCapabilitiesZipPaymentsParams `form:"zip_payments" json:"zip_payments,omitempty"`
}

// Automatically declines certain charge types regardless of whether the card issuer accepted or declined the charge.
type V2CoreAccountConfigurationMerchantCardPaymentsDeclineOnParams struct {
	// Whether Stripe automatically declines charges with an incorrect ZIP or postal code. This setting only applies when a ZIP or postal code is provided and they fail bank verification.
	AVSFailure *bool `form:"avs_failure" json:"avs_failure,omitempty"`
	// Whether Stripe automatically declines charges with an incorrect CVC. This setting only applies when a CVC is provided and it fails bank verification.
	CVCFailure *bool `form:"cvc_failure" json:"cvc_failure,omitempty"`
}

// Card payments settings.
type V2CoreAccountConfigurationMerchantCardPaymentsParams struct {
	// Automatically declines certain charge types regardless of whether the card issuer accepted or declined the charge.
	DeclineOn *V2CoreAccountConfigurationMerchantCardPaymentsDeclineOnParams `form:"decline_on" json:"decline_on,omitempty"`
}

// Support hours for Konbini payments.
type V2CoreAccountConfigurationMerchantKonbiniPaymentsSupportHoursParams struct {
	// Support hours end time (JST time of day) for in `HH:MM` format.
	EndTime *string `form:"end_time" json:"end_time,omitempty"`
	// Support hours start time (JST time of day) for in `HH:MM` format.
	StartTime *string `form:"start_time" json:"start_time,omitempty"`
}

// Support for Konbini payments.
type V2CoreAccountConfigurationMerchantKonbiniPaymentsSupportParams struct {
	// Support email address for Konbini payments.
	Email *string `form:"email" json:"email,omitempty"`
	// Support hours for Konbini payments.
	Hours *V2CoreAccountConfigurationMerchantKonbiniPaymentsSupportHoursParams `form:"hours" json:"hours,omitempty"`
	// Support phone number for Konbini payments.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Settings specific to Konbini payments on the account.
type V2CoreAccountConfigurationMerchantKonbiniPaymentsParams struct {
	// Support for Konbini payments.
	Support *V2CoreAccountConfigurationMerchantKonbiniPaymentsSupportParams `form:"support" json:"support,omitempty"`
}

// The Kana variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
type V2CoreAccountConfigurationMerchantScriptStatementDescriptorKanaParams struct {
	// The default text that appears on statements for non-card charges outside of Japan. For card charges, if you don't set a statement_descriptor_prefix, this text is also used as the statement descriptor prefix. In that case, if concatenating the statement descriptor suffix causes the combined statement descriptor to exceed 22 characters, we truncate the statement_descriptor text to limit the full descriptor to 22 characters. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Descriptor *string `form:"descriptor" json:"descriptor,omitempty"`
	// Default text that appears on statements for card charges outside of Japan, prefixing any dynamic statement_descriptor_suffix specified on the charge. To maximize space for the dynamic part of the descriptor, keep this text short. If you don't specify this value, statement_descriptor is used as the prefix. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
}

// The Kanji variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
type V2CoreAccountConfigurationMerchantScriptStatementDescriptorKanjiParams struct {
	// The default text that appears on statements for non-card charges outside of Japan. For card charges, if you don't set a statement_descriptor_prefix, this text is also used as the statement descriptor prefix. In that case, if concatenating the statement descriptor suffix causes the combined statement descriptor to exceed 22 characters, we truncate the statement_descriptor text to limit the full descriptor to 22 characters. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Descriptor *string `form:"descriptor" json:"descriptor,omitempty"`
	// Default text that appears on statements for card charges outside of Japan, prefixing any dynamic statement_descriptor_suffix specified on the charge. To maximize space for the dynamic part of the descriptor, keep this text short. If you don't specify this value, statement_descriptor is used as the prefix. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
}

// Settings for the default text that appears on statements for language variations.
type V2CoreAccountConfigurationMerchantScriptStatementDescriptorParams struct {
	// The Kana variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
	Kana *V2CoreAccountConfigurationMerchantScriptStatementDescriptorKanaParams `form:"kana" json:"kana,omitempty"`
	// The Kanji variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
	Kanji *V2CoreAccountConfigurationMerchantScriptStatementDescriptorKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Settings for Smart Disputes auto_respond.
type V2CoreAccountConfigurationMerchantSmartDisputesAutoRespondParams struct {
	// The preference for automatic dispute responses.
	Preference *string `form:"preference" json:"preference,omitempty"`
}

// Settings used for Smart Disputes.
type V2CoreAccountConfigurationMerchantSmartDisputesParams struct {
	// Settings for Smart Disputes auto_respond.
	AutoRespond *V2CoreAccountConfigurationMerchantSmartDisputesAutoRespondParams `form:"auto_respond" json:"auto_respond,omitempty"`
}

// Statement descriptor.
type V2CoreAccountConfigurationMerchantStatementDescriptorParams struct {
	// The default text that appears on statements for non-card charges outside of Japan. For card charges, if you don't set a statement_descriptor_prefix, this text is also used as the statement descriptor prefix. In that case, if concatenating the statement descriptor suffix causes the combined statement descriptor to exceed 22 characters, we truncate the statement_descriptor text to limit the full descriptor to 22 characters. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Descriptor *string `form:"descriptor" json:"descriptor,omitempty"`
	// Default text that appears on statements for card charges outside of Japan, prefixing any dynamic statement_descriptor_suffix specified on the charge. To maximize space for the dynamic part of the descriptor, keep this text short. If you don't specify this value, statement_descriptor is used as the prefix. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
}

// A publicly available mailing address for sending support issues to.
type V2CoreAccountConfigurationMerchantSupportAddressParams struct {
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

// Publicly available contact information for sending support issues to.
type V2CoreAccountConfigurationMerchantSupportParams struct {
	// A publicly available mailing address for sending support issues to.
	Address *V2CoreAccountConfigurationMerchantSupportAddressParams `form:"address" json:"address,omitempty"`
	// A publicly available email address for sending support issues to.
	Email *string `form:"email" json:"email,omitempty"`
	// A publicly available phone number to call with support issues.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// A publicly available website for handling support issues.
	URL *string `form:"url" json:"url,omitempty"`
}

// Enables the Account to act as a connected account and collect payments facilitated by a Connect platform. You must onboard your platform to Connect before you can add this configuration to your connected accounts. Utilize this configuration when the Account will be the Merchant of Record, like with Direct charges or Destination Charges with on_behalf_of set.
type V2CoreAccountConfigurationMerchantParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Settings for Bacs Direct Debit payments.
	BACSDebitPayments *V2CoreAccountConfigurationMerchantBACSDebitPaymentsParams `form:"bacs_debit_payments" json:"bacs_debit_payments,omitempty"`
	// Settings used to apply the merchant's branding to email receipts, invoices, Checkout, and other products.
	Branding *V2CoreAccountConfigurationMerchantBrandingParams `form:"branding" json:"branding,omitempty"`
	// Capabilities to request on the Merchant Configuration.
	Capabilities *V2CoreAccountConfigurationMerchantCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// Card payments settings.
	CardPayments *V2CoreAccountConfigurationMerchantCardPaymentsParams `form:"card_payments" json:"card_payments,omitempty"`
	// Settings specific to Konbini payments on the account.
	KonbiniPayments *V2CoreAccountConfigurationMerchantKonbiniPaymentsParams `form:"konbini_payments" json:"konbini_payments,omitempty"`
	// The Merchant Category Code (MCC) for the Merchant Configuration. MCCs classify businesses based on the goods or services they provide.
	MCC *string `form:"mcc" json:"mcc,omitempty"`
	// Settings for the default text that appears on statements for language variations.
	ScriptStatementDescriptor *V2CoreAccountConfigurationMerchantScriptStatementDescriptorParams `form:"script_statement_descriptor" json:"script_statement_descriptor,omitempty"`
	// Settings for Smart Disputes automatic response feature.
	SmartDisputes *V2CoreAccountConfigurationMerchantSmartDisputesParams `form:"smart_disputes" json:"smart_disputes,omitempty"`
	// Settings for the default [statement descriptor](https://docs.stripe.com/connect/statement-descriptors) text.
	StatementDescriptor *V2CoreAccountConfigurationMerchantStatementDescriptorParams `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Publicly available contact information for sending support issues to.
	Support *V2CoreAccountConfigurationMerchantSupportParams `form:"support" json:"support,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in AUD.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in CAD.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in EUR.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in GBP.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in USD.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in USDC.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can receive business storage-type funds on Stripe.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundParams struct {
	// Can receive business storage-type funds on Stripe in AUD.
	AUD *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDParams `form:"aud" json:"aud,omitempty"`
	// Can receive business storage-type funds on Stripe in CAD.
	CAD *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADParams `form:"cad" json:"cad,omitempty"`
	// Can receive business storage-type funds on Stripe in EUR.
	EUR *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURParams `form:"eur" json:"eur,omitempty"`
	// Can receive business storage-type funds on Stripe in GBP.
	GBP *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPParams `form:"gbp" json:"gbp,omitempty"`
	// Can receive business storage-type funds on Stripe in USD.
	USD *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can receive business storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in AUD.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in CAD.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in EUR.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in GBP.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in USD.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in USDC.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send business storage-type funds on Stripe.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundParams struct {
	// Can send business storage-type funds on Stripe in AUD.
	AUD *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDParams `form:"aud" json:"aud,omitempty"`
	// Can send business storage-type funds on Stripe in CAD.
	CAD *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADParams `form:"cad" json:"cad,omitempty"`
	// Can send business storage-type funds on Stripe in EUR.
	EUR *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURParams `form:"eur" json:"eur,omitempty"`
	// Can send business storage-type funds on Stripe in GBP.
	GBP *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPParams `form:"gbp" json:"gbp,omitempty"`
	// Can send business storage-type funds on Stripe in USD.
	USD *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can send business storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Can send or receive business storage-type funds on Stripe.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageParams struct {
	// Can receive business storage-type funds on Stripe.
	Inbound *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundParams `form:"inbound" json:"inbound,omitempty"`
	// Can send business storage-type funds on Stripe.
	Outbound *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundParams `form:"outbound" json:"outbound,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive consumer storage-type funds on Stripe in USD.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive consumer storage-type funds on Stripe in USDC.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can receive consumer storage-type funds on Stripe.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundParams struct {
	// Can receive consumer storage-type funds on Stripe in USD.
	USD *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can receive consumer storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send consumer storage-type funds on Stripe in USD.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send consumer storage-type funds on Stripe in USDC.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send consumer storage-type funds on Stripe.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundParams struct {
	// Can send consumer storage-type funds on Stripe in USD.
	USD *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can send consumer storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Can send or receive consumer storage-type funds on Stripe.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageParams struct {
	// Can receive consumer storage-type funds on Stripe.
	Inbound *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageInboundParams `form:"inbound" json:"inbound,omitempty"`
	// Can send consumer storage-type funds on Stripe.
	Outbound *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundParams `form:"outbound" json:"outbound,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesInboundTransfersParams struct {
	// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
	BankAccounts *V2CoreAccountConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a bank account owned by someone else.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a debit card owned by someone else.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to someone else via paper check.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a destination owned by someone else.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by someone else.
	BankAccounts *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a debit card owned by someone else.
	Cards *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsParams `form:"cards" json:"cards,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
	CryptoWallets *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
	FinancialAccounts *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
	// Can send funds from a FinancialAccount to someone else via paper check.
	PaperChecks *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksParams `form:"paper_checks" json:"paper_checks,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a bank account owned by yourself.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a destination owned by yourself.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by yourself.
	BankAccounts *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
	CryptoWallets *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
	FinancialAccounts *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive funds on a bank-account-like financial address (VBAN) to credit a FinancialAccount.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive funds on a crypto wallet like financial address to credit a FinancialAccount.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can receive funds into a FinancialAccount.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsParams struct {
	// Can receive funds on a bank-account-like financial address (VBAN) to credit a FinancialAccount.
	BankAccounts *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can receive funds on a crypto wallet like financial address to credit a FinancialAccount.
	CryptoWallets *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive debits to a FinancialAccount from a bank account.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can receive debits to a FinancialAccount.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedDebitsParams struct {
	// Can receive debits to a FinancialAccount from a bank account.
	BankAccounts *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
}

// Capabilities to request on the Money Manager Configuration.
type V2CoreAccountConfigurationMoneyManagerCapabilitiesParams struct {
	// Can send or receive business storage-type funds on Stripe.
	BusinessStorage *V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageParams `form:"business_storage" json:"business_storage,omitempty"`
	// Can send or receive consumer storage-type funds on Stripe.
	ConsumerStorage *V2CoreAccountConfigurationMoneyManagerCapabilitiesConsumerStorageParams `form:"consumer_storage" json:"consumer_storage,omitempty"`
	// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
	InboundTransfers *V2CoreAccountConfigurationMoneyManagerCapabilitiesInboundTransfersParams `form:"inbound_transfers" json:"inbound_transfers,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by someone else.
	OutboundPayments *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsParams `form:"outbound_payments" json:"outbound_payments,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by yourself.
	OutboundTransfers *V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersParams `form:"outbound_transfers" json:"outbound_transfers,omitempty"`
	// Can receive funds on a financial address to credit a FinancialAccount.
	ReceivedCredits *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsParams `form:"received_credits" json:"received_credits,omitempty"`
	// Can receive debits to a FinancialAccount.
	ReceivedDebits *V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedDebitsParams `form:"received_debits" json:"received_debits,omitempty"`
}

// Details of the regulated activity if the business participates in one.
type V2CoreAccountConfigurationMoneyManagerRegulatedActivityParams struct {
	// A detailed description of the regulated activities the business is licensed to conduct.
	Description *string `form:"description" json:"description,omitempty"`
	// The license number or registration number assigned by the business's primary regulator.
	LicenseNumber *string `form:"license_number" json:"license_number,omitempty"`
	// The country of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityCountry *string `form:"primary_regulatory_authority_country" json:"primary_regulatory_authority_country,omitempty"`
	// The name of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityName *string `form:"primary_regulatory_authority_name" json:"primary_regulatory_authority_name,omitempty"`
}

// The Money Manager Configuration allows the Account to store and move funds using FinancialAccounts.
type V2CoreAccountConfigurationMoneyManagerParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Capabilities to request on the Money Manager Configuration.
	Capabilities *V2CoreAccountConfigurationMoneyManagerCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// List of high-risk activities the business is involved in.
	HighRiskActivities []*string `form:"high_risk_activities" json:"high_risk_activities,omitempty"`
	// Description of the high-risk activities the business offers.
	HighRiskActivitiesDescription *string `form:"high_risk_activities_description" json:"high_risk_activities_description,omitempty"`
	// Description of the money services offered by the business.
	MoneyServicesDescription *string `form:"money_services_description" json:"money_services_description,omitempty"`
	// Indicates whether the business operates in any prohibited countries.
	OperatesInProhibitedCountries *bool `form:"operates_in_prohibited_countries" json:"operates_in_prohibited_countries,omitempty"`
	// Indicates whether the business participates in any regulated activity.
	ParticipatesInRegulatedActivity *bool `form:"participates_in_regulated_activity" json:"participates_in_regulated_activity,omitempty"`
	// Primary purpose of the stored funds.
	PurposeOfFunds *string `form:"purpose_of_funds" json:"purpose_of_funds,omitempty"`
	// Description of the purpose of the stored funds.
	PurposeOfFundsDescription *string `form:"purpose_of_funds_description" json:"purpose_of_funds_description,omitempty"`
	// Details of the regulated activity if the business participates in one.
	RegulatedActivity *V2CoreAccountConfigurationMoneyManagerRegulatedActivityParams `form:"regulated_activity" json:"regulated_activity,omitempty"`
	// The source of funds for the business, e.g. profits, income, venture capital, etc.
	SourceOfFunds *string `form:"source_of_funds" json:"source_of_funds,omitempty"`
	// Description of the source of funds for the business' account.
	SourceOfFundsDescription *string `form:"source_of_funds_description" json:"source_of_funds_description,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsACHProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsACHProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsACHProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over ACH rails.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsACHParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsACHProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over BECS rails.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsBECSParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsEftProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsEftProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsEftProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over EFT rails.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsEftParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsEftProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over Fedwire or CHIPS.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFedwireParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over FPS rails.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFPSParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsInstantParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocalParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsNppProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsNppProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsNppProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over NPP (real time) rails.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsNppParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsNppProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over RTP rails.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsRTPParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA credit rails.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPACreditParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA instant (real time) rails.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPAInstantParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over SWIFT.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSwiftParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWireProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWireProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWireProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWireParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWireProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities that enable OutboundPayments to a bank account linked to this Account.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsParams struct {
	// Enables this Account to receive OutboundPayments to linked bank accounts over ACH rails.
	ACH *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsACHParams `form:"ach" json:"ach,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over BECS rails.
	BECS *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsBECSParams `form:"becs" json:"becs,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over EFT rails.
	Eft *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsEftParams `form:"eft" json:"eft,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over Fedwire or CHIPS.
	Fedwire *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFedwireParams `form:"fedwire" json:"fedwire,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over FPS rails.
	FPS *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsFPSParams `form:"fps" json:"fps,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
	Instant *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsInstantParams `form:"instant" json:"instant,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
	Local *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocalParams `form:"local" json:"local,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over NPP (real time) rails.
	Npp *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsNppParams `form:"npp" json:"npp,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over RTP rails.
	RTP *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsRTPParams `form:"rtp" json:"rtp,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA credit rails.
	SEPACredit *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPACreditParams `form:"sepa_credit" json:"sepa_credit,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA instant (real time) rails.
	SEPAInstant *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSEPAInstantParams `form:"sepa_instant" json:"sepa_instant,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over SWIFT.
	Swift *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsSwiftParams `form:"swift" json:"swift,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
	Wire *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWireParams `form:"wire" json:"wire,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesCardsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesCardsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesCardsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Capabilities that enable OutboundPayments to a card linked to this Account.
type V2CoreAccountConfigurationRecipientCapabilitiesCardsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesCardsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Capabilities that enable OutboundPayments to a crypto wallet linked to this Account.
type V2CoreAccountConfigurationRecipientCapabilitiesCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesPaperChecksProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesPaperChecksProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesPaperChecksProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Capabilities that enable OutboundPayments via paper check.
type V2CoreAccountConfigurationRecipientCapabilitiesPaperChecksParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesPaperChecksProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive /v1/transfers into their Stripe Balance (/v1/balance).
type V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities that enable the recipient to manage their Stripe Balance (/v1/balance).
type V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceParams struct {
	// Enables this Account to receive /v1/transfers into their Stripe Balance (/v1/balance).
	StripeTransfers *V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams `form:"stripe_transfers" json:"stripe_transfers,omitempty"`
}

// Capabilities to be requested on the Recipient Configuration.
type V2CoreAccountConfigurationRecipientCapabilitiesParams struct {
	// Capabilities that enable OutboundPayments to a bank account linked to this Account.
	BankAccounts *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Capability that enable OutboundPayments to a debit card linked to this Account.
	Cards *V2CoreAccountConfigurationRecipientCapabilitiesCardsParams `form:"cards" json:"cards,omitempty"`
	// Capabilities that enable OutboundPayments to a crypto wallet linked to this Account.
	CryptoWallets *V2CoreAccountConfigurationRecipientCapabilitiesCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Capabilities that enable OutboundPayments via paper check.
	PaperChecks *V2CoreAccountConfigurationRecipientCapabilitiesPaperChecksParams `form:"paper_checks" json:"paper_checks,omitempty"`
	// Capabilities that enable the recipient to manage their Stripe Balance (/v1/balance).
	StripeBalance *V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceParams `form:"stripe_balance" json:"stripe_balance,omitempty"`
}

// The Recipient Configuration allows the Account to receive funds. Utilize this configuration if the Account will not be the Merchant of Record, like with Separate Charges & Transfers, or Destination Charges without on_behalf_of set.
type V2CoreAccountConfigurationRecipientParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Capabilities to be requested on the Recipient Configuration.
	Capabilities *V2CoreAccountConfigurationRecipientCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// The payout method id to be used as a default outbound destination. This will allow the PayoutMethod to be omitted on OutboundPayments made through API or sending payouts via dashboard. Can also be explicitly set to `null` to clear the existing default outbound destination. For further details about creating an Outbound Destination, see [Collect recipient's payment details](https://docs.stripe.com/global-payouts-private-preview/quickstart?dashboard-or-api=api#collect-bank-account-details).
	DefaultOutboundDestination *string `form:"default_outbound_destination" json:"default_outbound_destination,omitempty"`
}

// An Account Configuration which allows the Account to take on a key persona across Stripe products.
type V2CoreAccountConfigurationParams struct {
	// The CardCreator Configuration allows the Account to create and issue cards to users.
	CardCreator *V2CoreAccountConfigurationCardCreatorParams `form:"card_creator" json:"card_creator,omitempty"`
	// The Customer Configuration allows the Account to be used in inbound payment flows (i.e. customer-facing payment and billing flows).
	Customer *V2CoreAccountConfigurationCustomerParams `form:"customer" json:"customer,omitempty"`
	// Enables the Account to act as a connected account and collect payments facilitated by a Connect platform. You must onboard your platform to Connect before you can add this configuration to your connected accounts. Utilize this configuration when the Account will be the Merchant of Record, like with Direct charges or Destination Charges with on_behalf_of set.
	Merchant *V2CoreAccountConfigurationMerchantParams `form:"merchant" json:"merchant,omitempty"`
	// The Money Manager Configuration allows the Account to store and move funds using FinancialAccounts.
	MoneyManager *V2CoreAccountConfigurationMoneyManagerParams `form:"money_manager" json:"money_manager,omitempty"`
	// The Recipient Configuration allows the Account to receive funds. Utilize this configuration if the Account will not be the Merchant of Record, like with Separate Charges & Transfers, or Destination Charges without on_behalf_of set.
	Recipient *V2CoreAccountConfigurationRecipientParams `form:"recipient" json:"recipient,omitempty"`
}

// Account profile information.
type V2CoreAccountDefaultsProfileParams struct {
	// The business's publicly-available website.
	BusinessURL *string `form:"business_url" json:"business_url,omitempty"`
	// The name which is used by the business.
	DoingBusinessAs *string `form:"doing_business_as" json:"doing_business_as,omitempty"`
	// Internal-only description of the product sold or service provided by the business. It's used by Stripe for risk and underwriting purposes.
	ProductDescription *string `form:"product_description" json:"product_description,omitempty"`
}

// Default responsibilities held by either Stripe or the platform.
type V2CoreAccountDefaultsResponsibilitiesParams struct {
	// A value indicating the party responsible for collecting fees from this account.
	FeesCollector *string `form:"fees_collector" json:"fees_collector"`
	// A value indicating who is responsible for losses when this Account can't pay back negative balances from payments.
	LossesCollector *string `form:"losses_collector" json:"losses_collector"`
	// A value indicating the party responsible for collecting requirements on this account.
	RequirementsCollector *string `form:"requirements_collector" json:"requirements_collector,omitempty"`
}

// Default values to be used on Account Configurations.
type V2CoreAccountDefaultsParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The Account's preferred locales (languages), ordered by preference.
	Locales []*string `form:"locales" json:"locales,omitempty"`
	// Account profile information.
	Profile *V2CoreAccountDefaultsProfileParams `form:"profile" json:"profile,omitempty"`
	// Default responsibilities held by either Stripe or the platform.
	Responsibilities *V2CoreAccountDefaultsResponsibilitiesParams `form:"responsibilities" json:"responsibilities,omitempty"`
	// The Account's local timezone. A list of possible time zone values is maintained at the [IANA Time Zone Database](https://www.iana.org/time-zones).
	Timezone *string `form:"timezone" json:"timezone,omitempty"`
}

// This hash is used to attest that the directors information provided to Stripe is both current and correct.
type V2CoreAccountIdentityAttestationsDirectorshipDeclarationParams struct {
	// The time marking when the director attestation was made. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the director attestation was made.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the director attestation was made.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
type V2CoreAccountIdentityAttestationsOwnershipDeclarationParams struct {
	// The time marking when the beneficial owner attestation was made. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the beneficial owner attestation was made.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the beneficial owner attestation was made.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Attestation that all Persons with a specific Relationship value have been provided.
type V2CoreAccountIdentityAttestationsPersonsProvidedParams struct {
	// Whether the company's directors have been provided. Set this Boolean to true after creating all the company's directors with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Directors *bool `form:"directors" json:"directors,omitempty"`
	// Whether the company's executives have been provided. Set this Boolean to true after creating all the company's executives with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Executives *bool `form:"executives" json:"executives,omitempty"`
	// Whether the company's owners have been provided. Set this Boolean to true after creating all the company's owners with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Owners *bool `form:"owners" json:"owners,omitempty"`
	// Reason for why the company is exempt from providing ownership information.
	OwnershipExemptionReason *string `form:"ownership_exemption_reason" json:"ownership_exemption_reason,omitempty"`
}

// This hash is used to attest that the representative is authorized to act as the representative of their legal entity.
type V2CoreAccountIdentityAttestationsRepresentativeDeclarationParams struct {
	// The time marking when the representative attestation was made. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the representative attestation was made.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the representative attestation was made.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of the [Stripe Services Agreement](https://docs.stripe.com/connect/updating-accounts#tos-acceptance).
type V2CoreAccountIdentityAttestationsTermsOfServiceAccountParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for Stripe commercial card issuing.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardParams struct {
	// Bank terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Platform terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
	Platform *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardParams struct {
	// Bank terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
	FinancingDisclosures *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
	Platform *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Celtic as BIN sponsor.
	ApplePay *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing charge cards with Celtic as BIN sponsor.
	ChargeCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Terms of service acceptances for commercial issuing spend cards with Celtic as BIN sponsor.
	SpendCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardParams struct {
	// Bank terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	FinancingDisclosures *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	Platform *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardParams struct {
	// Bank terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Platform terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	Platform *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardParams struct {
	// Bank terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	FinancingDisclosures *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Cross River Bank as BIN sponsor.
	ApplePay *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	ChargeCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Terms of service acceptances for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
	// Terms of service acceptances for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Fifth Third as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardParams struct {
	// Bank terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
	FinancingDisclosures *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
	Platform *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Fifth Third as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Fifth Third as BIN sponsor.
	ApplePay *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing charge cards with Fifth Third as BIN sponsor.
	ChargeCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
}

// Terms of service acceptances for Stripe commercial card Global issuing.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialGlobalAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardParams struct {
	// Bank terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Platform terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
	Platform *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Lead as BIN sponsor.
	ApplePay *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Terms of service acceptances to create cards for commercial issuing use cases.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialParams struct {
	// Terms of service acceptances for Stripe commercial card issuing.
	AccountHolder *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialAccountHolderParams `form:"account_holder" json:"account_holder,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Fifth Third as BIN sponsor.
	FifthThird *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdParams `form:"fifth_third" json:"fifth_third,omitempty"`
	// Terms of service acceptances for Stripe commercial card Global issuing.
	GlobalAccountHolder *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialGlobalAccountHolderParams `form:"global_account_holder" json:"global_account_holder,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadParams `form:"lead" json:"lead,omitempty"`
}

// Terms of service acceptances for Stripe consumer card issuing.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing Apple Pay cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardParams struct {
	// Bank terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
	FinancingDisclosures *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
	Platform *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for consumer issuing cards with Celtic as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticParams struct {
	// Terms of service acceptances for consumer issuing Apple Pay cards with Celtic as BIN sponsor.
	ApplePay *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for consumer issuing revolving credit cards with Celtic as BIN sponsor.
	RevolvingCreditCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardParams `form:"revolving_credit_card" json:"revolving_credit_card,omitempty"`
}

// Bank terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardParams struct {
	// Bank terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	FinancingDisclosures *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	Platform *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for consumer issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankParams struct {
	// Terms of service acceptances for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Terms of service acceptances for Stripe consumer card Global issuing.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerGlobalAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardParams struct {
	// Bank terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
	FinancingDisclosures *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
	Platform *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardParams struct {
	// Bank terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
	BankTerms *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
	FinancingDisclosures *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
	Platform *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for consumer issuing cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadParams struct {
	// Terms of service acceptances for consumer issuing Apple Pay cards with Lead as BIN sponsor.
	ApplePay *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for consumer issuing debit cards with Lead as BIN sponsor.
	DebitCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardParams `form:"debit_card" json:"debit_card,omitempty"`
	// Terms of service acceptances for consumer issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Terms of service acceptances to create cards for consumer issuing use cases.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerParams struct {
	// Terms of service acceptances for Stripe consumer card issuing.
	AccountHolder *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerAccountHolderParams `form:"account_holder" json:"account_holder,omitempty"`
	// Terms of service acceptances for consumer issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Terms of service acceptances for consumer issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Terms of service acceptances for Stripe consumer card Global issuing.
	GlobalAccountHolder *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerGlobalAccountHolderParams `form:"global_account_holder" json:"global_account_holder,omitempty"`
	// Terms of service acceptances for consumer issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadParams `form:"lead" json:"lead,omitempty"`
}

// Details on the Account's acceptance of Issuing-specific terms of service.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorParams struct {
	// Terms of service acceptances to create cards for commercial issuing use cases.
	Commercial *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialParams `form:"commercial" json:"commercial,omitempty"`
	// Terms of service acceptances to create cards for consumer issuing use cases.
	Consumer *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerParams `form:"consumer" json:"consumer,omitempty"`
}

// Details on the Account's acceptance of Consumer-specific terms of service.
type V2CoreAccountIdentityAttestationsTermsOfServiceConsumerMoneyManagerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Consumer-privacy-disclosures-specific terms of service.
type V2CoreAccountIdentityAttestationsTermsOfServiceConsumerPrivacyDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Crypto-specific terms of service.
type V2CoreAccountIdentityAttestationsTermsOfServiceCryptoMoneyManagerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Treasury-specific terms of service.
type V2CoreAccountIdentityAttestationsTermsOfServiceMoneyManagerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountIdentityAttestationsTermsOfServiceParams struct {
	// Details on the Account's acceptance of the [Stripe Services Agreement](https://docs.stripe.com/connect/updating-accounts#tos-acceptance).
	Account *V2CoreAccountIdentityAttestationsTermsOfServiceAccountParams `form:"account" json:"account,omitempty"`
	// Details on the Account's acceptance of Issuing-specific terms of service.
	CardCreator *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorParams `form:"card_creator" json:"card_creator,omitempty"`
	// Details on the Account's acceptance of Consumer-specific terms of service.
	ConsumerMoneyManager *V2CoreAccountIdentityAttestationsTermsOfServiceConsumerMoneyManagerParams `form:"consumer_money_manager" json:"consumer_money_manager,omitempty"`
	// Details on the Account's acceptance of Consumer-privacy-disclosures-specific terms of service.
	ConsumerPrivacyDisclosures *V2CoreAccountIdentityAttestationsTermsOfServiceConsumerPrivacyDisclosuresParams `form:"consumer_privacy_disclosures" json:"consumer_privacy_disclosures,omitempty"`
	// Details on the Account's acceptance of Crypto-specific terms of service.
	CryptoMoneyManager *V2CoreAccountIdentityAttestationsTermsOfServiceCryptoMoneyManagerParams `form:"crypto_money_manager" json:"crypto_money_manager,omitempty"`
	// Details on the Account's acceptance of Treasury-specific terms of service.
	MoneyManager *V2CoreAccountIdentityAttestationsTermsOfServiceMoneyManagerParams `form:"money_manager" json:"money_manager,omitempty"`
}

// Attestations from the identity's key people, e.g. owners, executives, directors, representatives.
type V2CoreAccountIdentityAttestationsParams struct {
	// This hash is used to attest that the directors information provided to Stripe is both current and correct.
	DirectorshipDeclaration *V2CoreAccountIdentityAttestationsDirectorshipDeclarationParams `form:"directorship_declaration" json:"directorship_declaration,omitempty"`
	// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
	OwnershipDeclaration *V2CoreAccountIdentityAttestationsOwnershipDeclarationParams `form:"ownership_declaration" json:"ownership_declaration,omitempty"`
	// Attestation that all Persons with a specific Relationship value have been provided.
	PersonsProvided *V2CoreAccountIdentityAttestationsPersonsProvidedParams `form:"persons_provided" json:"persons_provided,omitempty"`
	// This hash is used to attest that the representative is authorized to act as the representative of their legal entity.
	RepresentativeDeclaration *V2CoreAccountIdentityAttestationsRepresentativeDeclarationParams `form:"representative_declaration" json:"representative_declaration,omitempty"`
	// Attestations of accepted terms of service agreements.
	TermsOfService *V2CoreAccountIdentityAttestationsTermsOfServiceParams `form:"terms_of_service" json:"terms_of_service,omitempty"`
}

// The business registration address of the business entity.
type V2CoreAccountIdentityBusinessDetailsAddressParams struct {
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

// The business gross annual revenue for its preceding fiscal year.
type V2CoreAccountIdentityBusinessDetailsAnnualRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// The close-out date of the preceding fiscal year in ISO 8601 format. E.g. 2023-12-31 for the 31st of December, 2023.
	FiscalYearEnd *string `form:"fiscal_year_end" json:"fiscal_year_end,omitempty"`
}

// One or more documents that support the bank account ownership verification requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a check.
type V2CoreAccountIdentityBusinessDetailsDocumentsBankAccountOwnershipVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's license to operate.
type V2CoreAccountIdentityBusinessDetailsDocumentsCompanyLicenseParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the company's Memorandum of Association.
type V2CoreAccountIdentityBusinessDetailsDocumentsCompanyMemorandumOfAssociationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Certain countries only: One or more documents showing the ministerial decree legalizing the company's establishment.
type V2CoreAccountIdentityBusinessDetailsDocumentsCompanyMinisterialDecreeParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
type V2CoreAccountIdentityBusinessDetailsDocumentsCompanyRegistrationVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's tax ID.
type V2CoreAccountIdentityBusinessDetailsDocumentsCompanyTaxIDVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountIdentityBusinessDetailsDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document verifying the business.
type V2CoreAccountIdentityBusinessDetailsDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountIdentityBusinessDetailsDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of address.
type V2CoreAccountIdentityBusinessDetailsDocumentsProofOfAddressParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Person that is signing the document.
type V2CoreAccountIdentityBusinessDetailsDocumentsProofOfRegistrationSignerParams struct {
	// Person signing the document.
	Person *string `form:"person" json:"person"`
}

// One or more documents showing the company's proof of registration with the national business registry.
type V2CoreAccountIdentityBusinessDetailsDocumentsProofOfRegistrationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// Person that is signing the document.
	Signer *V2CoreAccountIdentityBusinessDetailsDocumentsProofOfRegistrationSignerParams `form:"signer" json:"signer,omitempty"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Person that is signing the document.
type V2CoreAccountIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipSignerParams struct {
	// Person signing the document.
	Person *string `form:"person" json:"person"`
}

// One or more documents that demonstrate proof of ultimate beneficial ownership.
type V2CoreAccountIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// Person that is signing the document.
	Signer *V2CoreAccountIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipSignerParams `form:"signer" json:"signer,omitempty"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// A document verifying the business.
type V2CoreAccountIdentityBusinessDetailsDocumentsParams struct {
	// One or more documents that support the bank account ownership verification requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a check.
	BankAccountOwnershipVerification *V2CoreAccountIdentityBusinessDetailsDocumentsBankAccountOwnershipVerificationParams `form:"bank_account_ownership_verification" json:"bank_account_ownership_verification,omitempty"`
	// One or more documents that demonstrate proof of a company's license to operate.
	CompanyLicense *V2CoreAccountIdentityBusinessDetailsDocumentsCompanyLicenseParams `form:"company_license" json:"company_license,omitempty"`
	// One or more documents showing the company's Memorandum of Association.
	CompanyMemorandumOfAssociation *V2CoreAccountIdentityBusinessDetailsDocumentsCompanyMemorandumOfAssociationParams `form:"company_memorandum_of_association" json:"company_memorandum_of_association,omitempty"`
	// Certain countries only: One or more documents showing the ministerial decree legalizing the company's establishment.
	CompanyMinisterialDecree *V2CoreAccountIdentityBusinessDetailsDocumentsCompanyMinisterialDecreeParams `form:"company_ministerial_decree" json:"company_ministerial_decree,omitempty"`
	// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
	CompanyRegistrationVerification *V2CoreAccountIdentityBusinessDetailsDocumentsCompanyRegistrationVerificationParams `form:"company_registration_verification" json:"company_registration_verification,omitempty"`
	// One or more documents that demonstrate proof of a company's tax ID.
	CompanyTaxIDVerification *V2CoreAccountIdentityBusinessDetailsDocumentsCompanyTaxIDVerificationParams `form:"company_tax_id_verification" json:"company_tax_id_verification,omitempty"`
	// A document verifying the business.
	PrimaryVerification *V2CoreAccountIdentityBusinessDetailsDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// One or more documents that demonstrate proof of address.
	ProofOfAddress *V2CoreAccountIdentityBusinessDetailsDocumentsProofOfAddressParams `form:"proof_of_address" json:"proof_of_address,omitempty"`
	// One or more documents showing the company's proof of registration with the national business registry.
	ProofOfRegistration *V2CoreAccountIdentityBusinessDetailsDocumentsProofOfRegistrationParams `form:"proof_of_registration" json:"proof_of_registration,omitempty"`
	// One or more documents that demonstrate proof of ultimate beneficial ownership.
	ProofOfUltimateBeneficialOwnership *V2CoreAccountIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams `form:"proof_of_ultimate_beneficial_ownership" json:"proof_of_ultimate_beneficial_ownership,omitempty"`
}

// The ID numbers of a business entity.
type V2CoreAccountIdentityBusinessDetailsIDNumberParams struct {
	// The registrar of the ID number (Only valid for DE ID number types).
	Registrar *string `form:"registrar" json:"registrar,omitempty"`
	// Open Enum. The ID number type of a business entity.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// An estimate of the monthly revenue of the business.
type V2CoreAccountIdentityBusinessDetailsMonthlyEstimatedRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
}

// When the business was incorporated or registered.
type V2CoreAccountIdentityBusinessDetailsRegistrationDateParams struct {
	// The day of registration, between 1 and 31.
	Day *int64 `form:"day" json:"day"`
	// The month of registration, between 1 and 12.
	Month *int64 `form:"month" json:"month"`
	// The four-digit year of registration.
	Year *int64 `form:"year" json:"year"`
}

// Kana Address.
type V2CoreAccountIdentityBusinessDetailsScriptAddressesKanaParams struct {
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
type V2CoreAccountIdentityBusinessDetailsScriptAddressesKanjiParams struct {
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
type V2CoreAccountIdentityBusinessDetailsScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountIdentityBusinessDetailsScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountIdentityBusinessDetailsScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Kana name.
type V2CoreAccountIdentityBusinessDetailsScriptNamesKanaParams struct {
	// Registered name of the business.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// Kanji name.
type V2CoreAccountIdentityBusinessDetailsScriptNamesKanjiParams struct {
	// Registered name of the business.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// The business legal name in non latin script.
type V2CoreAccountIdentityBusinessDetailsScriptNamesParams struct {
	// Kana name.
	Kana *V2CoreAccountIdentityBusinessDetailsScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji name.
	Kanji *V2CoreAccountIdentityBusinessDetailsScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Information about the company or business.
type V2CoreAccountIdentityBusinessDetailsParams struct {
	// The business registration address of the business entity.
	Address *V2CoreAccountIdentityBusinessDetailsAddressParams `form:"address" json:"address,omitempty"`
	// The business gross annual revenue for its preceding fiscal year.
	AnnualRevenue *V2CoreAccountIdentityBusinessDetailsAnnualRevenueParams `form:"annual_revenue" json:"annual_revenue,omitempty"`
	// A detailed description of the business's compliance and anti-money laundering controls and practices.
	ComplianceScreeningDescription *string `form:"compliance_screening_description" json:"compliance_screening_description,omitempty"`
	// A document verifying the business.
	Documents *V2CoreAccountIdentityBusinessDetailsDocumentsParams `form:"documents" json:"documents,omitempty"`
	// Estimated maximum number of workers currently engaged by the business (including employees, contractors, and vendors).
	EstimatedWorkerCount *int64 `form:"estimated_worker_count" json:"estimated_worker_count,omitempty"`
	// The ID numbers of a business entity.
	IDNumbers []*V2CoreAccountIdentityBusinessDetailsIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// An estimate of the monthly revenue of the business.
	MonthlyEstimatedRevenue *V2CoreAccountIdentityBusinessDetailsMonthlyEstimatedRevenueParams `form:"monthly_estimated_revenue" json:"monthly_estimated_revenue,omitempty"`
	// The phone number of the Business Entity.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The business legal name.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
	// When the business was incorporated or registered.
	RegistrationDate *V2CoreAccountIdentityBusinessDetailsRegistrationDateParams `form:"registration_date" json:"registration_date,omitempty"`
	// The business registration address of the business entity in non latin script.
	ScriptAddresses *V2CoreAccountIdentityBusinessDetailsScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The business legal name in non latin script.
	ScriptNames *V2CoreAccountIdentityBusinessDetailsScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The category identifying the legal structure of the business.
	Structure *string `form:"structure" json:"structure,omitempty"`
}

// Additional addresses associated with the individual.
type V2CoreAccountIdentityIndividualAdditionalAddressParams struct {
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
type V2CoreAccountIdentityIndividualAdditionalNameParams struct {
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
type V2CoreAccountIdentityIndividualAddressParams struct {
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
type V2CoreAccountIdentityIndividualDateOfBirthParams struct {
	// The day of the birth.
	Day *int64 `form:"day" json:"day"`
	// The month of birth.
	Month *int64 `form:"month" json:"month"`
	// The year of birth.
	Year *int64 `form:"year" json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountIdentityIndividualDocumentsCompanyAuthorizationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountIdentityIndividualDocumentsPassportParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountIdentityIndividualDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountIdentityIndividualDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountIdentityIndividualDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountIdentityIndividualDocumentsSecondaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountIdentityIndividualDocumentsSecondaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountIdentityIndividualDocumentsSecondaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountIdentityIndividualDocumentsVisaParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountIdentityIndividualDocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountIdentityIndividualDocumentsCompanyAuthorizationParams `form:"company_authorization" json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountIdentityIndividualDocumentsPassportParams `form:"passport" json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountIdentityIndividualDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountIdentityIndividualDocumentsSecondaryVerificationParams `form:"secondary_verification" json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountIdentityIndividualDocumentsVisaParams `form:"visa" json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the individual.
type V2CoreAccountIdentityIndividualIDNumberParams struct {
	// The ID number type of an individual.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// The relationship that this individual has with the account's identity.
type V2CoreAccountIdentityIndividualRelationshipParams struct {
	// Whether the person is a director of the account's identity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.
	Director *bool `form:"director" json:"director,omitempty"`
	// Whether the person has significant responsibility to control, manage, or direct the organization.
	Executive *bool `form:"executive" json:"executive,omitempty"`
	// Whether the person is an owner of the account's identity.
	Owner *bool `form:"owner" json:"owner,omitempty"`
	// The percent owned by the person of the account's legal entity.
	PercentOwnership *float64 `form:"percent_ownership,high_precision" json:"percent_ownership,string,omitempty"`
	// The person's title (e.g., CEO, Support Engineer).
	Title *string `form:"title" json:"title,omitempty"`
}

// Kana Address.
type V2CoreAccountIdentityIndividualScriptAddressesKanaParams struct {
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
type V2CoreAccountIdentityIndividualScriptAddressesKanjiParams struct {
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
type V2CoreAccountIdentityIndividualScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountIdentityIndividualScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountIdentityIndividualScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountIdentityIndividualScriptNamesKanaParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountIdentityIndividualScriptNamesKanjiParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The individuals primary name in non latin script.
type V2CoreAccountIdentityIndividualScriptNamesParams struct {
	// Persons name in kana script.
	Kana *V2CoreAccountIdentityIndividualScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountIdentityIndividualScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Information about the person represented by the account.
type V2CoreAccountIdentityIndividualParams struct {
	// Additional addresses associated with the individual.
	AdditionalAddresses []*V2CoreAccountIdentityIndividualAdditionalAddressParams `form:"additional_addresses" json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the individual.
	AdditionalNames []*V2CoreAccountIdentityIndividualAdditionalNameParams `form:"additional_names" json:"additional_names,omitempty"`
	// The individual's residential address.
	Address *V2CoreAccountIdentityIndividualAddressParams `form:"address" json:"address,omitempty"`
	// The individual's date of birth.
	DateOfBirth *V2CoreAccountIdentityIndividualDateOfBirthParams `form:"date_of_birth" json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountIdentityIndividualDocumentsParams `form:"documents" json:"documents,omitempty"`
	// The individual's email address. You can only set this field when the Account is configured as a `merchant` or `recipient`. Use `contact_email` as the primary contact email for this Account.
	Email *string `form:"email" json:"email,omitempty"`
	// The individual's first name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The identification numbers (e.g., SSN) associated with the individual.
	IDNumbers []*V2CoreAccountIdentityIndividualIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
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
	Relationship *V2CoreAccountIdentityIndividualRelationshipParams `form:"relationship" json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the individual.
	ScriptAddresses *V2CoreAccountIdentityIndividualScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The individuals primary name in non latin script.
	ScriptNames *V2CoreAccountIdentityIndividualScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The individual's last name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountIdentityIndividualParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Information about the company, individual, and business represented by the Account.
type V2CoreAccountIdentityParams struct {
	// Attestations from the identity's key people, e.g. owners, executives, directors, representatives.
	Attestations *V2CoreAccountIdentityAttestationsParams `form:"attestations" json:"attestations,omitempty"`
	// Information about the company or business.
	BusinessDetails *V2CoreAccountIdentityBusinessDetailsParams `form:"business_details" json:"business_details,omitempty"`
	// The country in which the account holder resides, or in which the business is legally established. This should be an [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code.
	Country *string `form:"country" json:"country,omitempty"`
	// The entity type represented by the Account. Ensure this field is accurate before adding configurations that rely on identity information, as it determines which identity fields apply and how the Account is validated.
	EntityType *string `form:"entity_type" json:"entity_type,omitempty"`
	// Information about the individual represented by the Account. This property is `null` unless `entity_type` is set to `individual`.
	Individual *V2CoreAccountIdentityIndividualParams `form:"individual" json:"individual,omitempty"`
}

// Create an Account that represents a company, individual, or other entity that your business interacts with. Accounts contain identifying information about the entity, and configurations that store the features an account has access to. An account can be configured as any or all of the following configurations: Customer, Merchant and/or Recipient.
type V2CoreAccountParams struct {
	Params `form:"*"`
	// The account token generated by the account token api.
	AccountToken *string `form:"account_token" json:"account_token,omitempty"`
	// An Account Configuration which allows the Account to take on a key persona across Stripe products.
	Configuration *V2CoreAccountConfigurationParams `form:"configuration" json:"configuration,omitempty"`
	// The primary contact email address for the Account.
	ContactEmail *string `form:"contact_email" json:"contact_email,omitempty"`
	// The default contact phone for the Account.
	ContactPhone *string `form:"contact_phone" json:"contact_phone,omitempty"`
	// A value indicating the Stripe dashboard this Account has access to. This will depend on which configurations are enabled for this account.
	Dashboard *string `form:"dashboard" json:"dashboard,omitempty"`
	// Default values to be used on Account Configurations.
	Defaults *V2CoreAccountDefaultsParams `form:"defaults" json:"defaults,omitempty"`
	// A descriptive name for the Account. This name will be surfaced in the Stripe Dashboard and on any invoices sent to the Account.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Information about the company, individual, and business represented by the Account.
	Identity *V2CoreAccountIdentityParams `form:"identity" json:"identity,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Terms of service acceptances for consumer issuing Apple Pay cards with Lead as BIN sponsor.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Removes access to the Account and its associated resources. Closed Accounts can no longer be operated on, but limited information can still be retrieved through the API in order to be able to track their history.
type V2CoreAccountCloseParams struct {
	Params `form:"*"`
	// Configurations on the Account to be closed. All configurations on the Account must be passed in for this request to succeed.
	AppliedConfigurations []*string `form:"applied_configurations" json:"applied_configurations,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create commercial issuing cards with Celtic as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticParams struct {
	// Can create commercial issuing charge cards with Celtic as BIN sponsor.
	ChargeCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing spend cards with Celtic as BIN sponsor.
	SpendCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create commercial issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams struct {
	// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
	ChargeCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
	// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create commercial issuing cards with Fifth Third as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialFifthThirdParams struct {
	// Can create commercial issuing charge cards with Fifth Third as BIN sponsor.
	ChargeCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create commercial issuing cards with Stripe as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadParams struct {
	// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Stripe as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing prepaid cards with Stripe as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create commercial issuing cards with Stripe as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripeParams struct {
	// Can create commercial issuing charge cards with Stripe as BIN sponsor.
	ChargeCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing prepaid cards with Stripe as BIN sponsor.
	PrepaidCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Can create cards for commercial issuing use cases.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialParams struct {
	// Can create commercial issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Can create commercial issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Can create commercial issuing cards with Fifth Third as BIN sponsor.
	FifthThird *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialFifthThirdParams `form:"fifth_third" json:"fifth_third,omitempty"`
	// Can create commercial issuing cards with Stripe as BIN sponsor.
	Lead *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadParams `form:"lead" json:"lead,omitempty"`
	// Can create commercial issuing cards with Stripe as BIN sponsor.
	Stripe *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripeParams `form:"stripe" json:"stripe,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create consumer issuing cards with Celtic as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCelticParams struct {
	// Can create consumer issuing charge cards with Celtic as BIN sponsor.
	RevolvingCreditCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardParams `form:"revolving_credit_card" json:"revolving_credit_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create consumer issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankParams struct {
	// Can create consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create consumer issuing cards with Lead as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadParams struct {
	// Can create consumer issuing debit cards with Lead as BIN sponsor.
	DebitCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardParams `form:"debit_card" json:"debit_card,omitempty"`
	// Can create consumer issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Can create cards for consumer issuing use cases.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerParams struct {
	// Can create consumer issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Can create consumer issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Can create consumer issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerLeadParams `form:"lead" json:"lead,omitempty"`
}

// Capabilities to request on the CardCreator Configuration.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesParams struct {
	// Can create cards for commercial issuing use cases.
	Commercial *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialParams `form:"commercial" json:"commercial,omitempty"`
	// Can create cards for consumer issuing use cases.
	Consumer *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesConsumerParams `form:"consumer" json:"consumer,omitempty"`
}

// The CardCreator Configuration allows the Account to create and issue cards to users.
type V2CoreAccountCreateConfigurationCardCreatorParams struct {
	// Capabilities to request on the CardCreator Configuration.
	Capabilities *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
}

// Automatic indirect tax settings to be used when automatic tax calculation is enabled on the customer's invoices, subscriptions, checkout sessions, or payment links. Surfaces if automatic tax calculation is possible given the current customer location information.
type V2CoreAccountCreateConfigurationCustomerAutomaticIndirectTaxParams struct {
	// Describes the customer's tax exemption status, which is `none`, `exempt`, or `reverse`. When set to reverse, invoice and receipt PDFs include the following text: “Reverse charge”.
	Exempt *string `form:"exempt" json:"exempt,omitempty"`
	// A recent IP address of the customer used for tax reporting and tax location inference.
	IPAddress *string `form:"ip_address" json:"ip_address,omitempty"`
	// The data source used to identify the customer's tax location - defaults to `identity_address`. Will only be used for automatic tax calculation on the customer's Invoices and Subscriptions. This behavior is now deprecated for new users.
	LocationSource *string `form:"location_source" json:"location_source,omitempty"`
}

// The list of up to 4 default custom fields to be displayed on invoices for this customer.
type V2CoreAccountCreateConfigurationCustomerBillingInvoiceCustomFieldParams struct {
	// The name of the custom field. This may be up to 40 characters.
	Name *string `form:"name" json:"name"`
	// The value of the custom field. This may be up to 140 characters. When updating, pass an empty string to remove previously-defined values.
	Value *string `form:"value" json:"value"`
}

// Default invoice PDF rendering options.
type V2CoreAccountCreateConfigurationCustomerBillingInvoiceRenderingParams struct {
	// Indicates whether displayed line item prices and amounts on invoice PDFs include inclusive tax amounts. Must be either `include_inclusive_tax` or `exclude_tax`.
	AmountTaxDisplay *string `form:"amount_tax_display" json:"amount_tax_display,omitempty"`
	// ID of the invoice rendering template to use for future invoices.
	Template *string `form:"template" json:"template,omitempty"`
}

// Default invoice settings for the customer account.
type V2CoreAccountCreateConfigurationCustomerBillingInvoiceParams struct {
	// The list of up to 4 default custom fields to be displayed on invoices for this customer.
	CustomFields []*V2CoreAccountCreateConfigurationCustomerBillingInvoiceCustomFieldParams `form:"custom_fields" json:"custom_fields,omitempty"`
	// Default invoice footer.
	Footer *string `form:"footer" json:"footer,omitempty"`
	// Sequence number to use on the customer account's next invoice. Defaults to 1.
	NextSequence *int64 `form:"next_sequence" json:"next_sequence,omitempty"`
	// Prefix used to generate unique invoice numbers. Must be 3-12 uppercase letters or numbers.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
	// Default invoice PDF rendering options.
	Rendering *V2CoreAccountCreateConfigurationCustomerBillingInvoiceRenderingParams `form:"rendering" json:"rendering,omitempty"`
}

// Billing settings - default settings used for this customer in Billing flows such as Invoices and Subscriptions.
type V2CoreAccountCreateConfigurationCustomerBillingParams struct {
	// Default invoice settings for the customer account.
	Invoice *V2CoreAccountCreateConfigurationCustomerBillingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Generates requirements for enabling automatic indirect tax calculation on this customer's invoices or subscriptions. Recommended to request this capability if planning to enable automatic tax calculation on this customer's invoices or subscriptions.
type V2CoreAccountCreateConfigurationCustomerCapabilitiesAutomaticIndirectTaxParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Capabilities that have been requested on the Customer Configuration.
type V2CoreAccountCreateConfigurationCustomerCapabilitiesParams struct {
	// Generates requirements for enabling automatic indirect tax calculation on this customer's invoices or subscriptions. Recommended to request this capability if planning to enable automatic tax calculation on this customer's invoices or subscriptions.
	AutomaticIndirectTax *V2CoreAccountCreateConfigurationCustomerCapabilitiesAutomaticIndirectTaxParams `form:"automatic_indirect_tax" json:"automatic_indirect_tax,omitempty"`
}

// The customer's shipping information. Appears on invoices emailed to this customer.
type V2CoreAccountCreateConfigurationCustomerShippingParams struct {
	// Customer shipping address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Customer name.
	Name *string `form:"name" json:"name,omitempty"`
	// Customer phone (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// The Customer Configuration allows the Account to be used in inbound payment flows (i.e. customer-facing payment and billing flows).
type V2CoreAccountCreateConfigurationCustomerParams struct {
	// Automatic indirect tax settings to be used when automatic tax calculation is enabled on the customer's invoices, subscriptions, checkout sessions, or payment links. Surfaces if automatic tax calculation is possible given the current customer location information.
	AutomaticIndirectTax *V2CoreAccountCreateConfigurationCustomerAutomaticIndirectTaxParams `form:"automatic_indirect_tax" json:"automatic_indirect_tax,omitempty"`
	// Billing settings - default settings used for this customer in Billing flows such as Invoices and Subscriptions.
	Billing *V2CoreAccountCreateConfigurationCustomerBillingParams `form:"billing" json:"billing,omitempty"`
	// Capabilities that have been requested on the Customer Configuration.
	Capabilities *V2CoreAccountCreateConfigurationCustomerCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// The customer's shipping information. Appears on invoices emailed to this customer.
	Shipping *V2CoreAccountCreateConfigurationCustomerShippingParams `form:"shipping" json:"shipping,omitempty"`
	// ID of the test clock to attach to the customer. Can only be set on testmode Accounts, and when the Customer Configuration is first set on an Account.
	TestClock *string `form:"test_clock" json:"test_clock,omitempty"`
}

// Settings used for Bacs debit payments.
type V2CoreAccountCreateConfigurationMerchantBACSDebitPaymentsParams struct {
	// Display name for Bacs Direct Debit payments.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
}

// Settings used to apply the merchant's branding to email receipts, invoices, Checkout, and other products.
type V2CoreAccountCreateConfigurationMerchantBrandingParams struct {
	// ID of a [file upload](https://docs.stripe.com/api/persons/update#create_file): An icon for the merchant. Must be square and at least 128px x 128px.
	Icon *string `form:"icon" json:"icon,omitempty"`
	// ID of a [file upload](https://docs.stripe.com/api/persons/update#create_file): A logo for the merchant that will be used in Checkout instead of the icon and without the merchant's name next to it if provided. Must be at least 128px x 128px.
	Logo *string `form:"logo" json:"logo,omitempty"`
	// A CSS hex color value representing the primary branding color for the merchant.
	PrimaryColor *string `form:"primary_color" json:"primary_color,omitempty"`
	// A CSS hex color value representing the secondary branding color for the merchant.
	SecondaryColor *string `form:"secondary_color" json:"secondary_color,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process ACH debit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesACHDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process ACSS debit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesACSSDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Affirm payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAffirmPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Afterpay/Clearpay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Alma payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAlmaPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Amazon Pay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAmazonPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Australian BECS Direct Debit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process BACS Direct Debit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBACSDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Bancontact payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBancontactPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process BLIK payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBLIKPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Boleto payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBoletoPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCardPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCardPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesCardPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to collect card payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCardPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesCardPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Cartes Bancaires payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCartesBancairesPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Cash App payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCashAppPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesEPSPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesEPSPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesEPSPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process EPS payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesEPSPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesEPSPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesFPXPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesFPXPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesFPXPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process FPX payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesFPXPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesFPXPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process UK bank transfer payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesGBBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process GrabPay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesGrabpayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process iDEAL payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesIDEALPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesJCBPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesJCBPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesJCBPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process JCB card payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesJCBPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesJCBPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Japanese bank transfer payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesJPBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Kakao Pay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKakaoPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Klarna payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKlarnaPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Konbini convenience store payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKonbiniPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Korean card payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKrCardPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesLinkPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesLinkPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesLinkPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Link payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesLinkPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesLinkPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process MobilePay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMobilepayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Multibanco payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMultibancoPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Mexican bank transfer payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMXBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Naver Pay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesNaverPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process OXXO payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesOXXOPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesP24PaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesP24PaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesP24PaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Przelewy24 (P24) payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesP24PaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesP24PaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Pay by Bank payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPayByBankPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process PAYCO payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPaycoPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process PayNow payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPayNowPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process PromptPay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPromptPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Revolut Pay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesRevolutPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Samsung Pay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSamsungPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process SEPA bank transfer payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process SEPA Direct Debit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPADebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Sunbit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSunbitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSwishPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSwishPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesSwishPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Swish payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSwishPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesSwishPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process TWINT payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesTWINTPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process US bank transfer payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesUSBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesZipPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMerchantCapabilitiesZipPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMerchantCapabilitiesZipPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Zip payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesZipPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMerchantCapabilitiesZipPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Capabilities to request on the Merchant Configuration.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesParams struct {
	// Allow the merchant to process ACH debit payments.
	ACHDebitPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesACHDebitPaymentsParams `form:"ach_debit_payments" json:"ach_debit_payments,omitempty"`
	// Allow the merchant to process ACSS debit payments.
	ACSSDebitPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesACSSDebitPaymentsParams `form:"acss_debit_payments" json:"acss_debit_payments,omitempty"`
	// Allow the merchant to process Affirm payments.
	AffirmPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesAffirmPaymentsParams `form:"affirm_payments" json:"affirm_payments,omitempty"`
	// Allow the merchant to process Afterpay/Clearpay payments.
	AfterpayClearpayPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsParams `form:"afterpay_clearpay_payments" json:"afterpay_clearpay_payments,omitempty"`
	// Allow the merchant to process Alma payments.
	AlmaPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesAlmaPaymentsParams `form:"alma_payments" json:"alma_payments,omitempty"`
	// Allow the merchant to process Amazon Pay payments.
	AmazonPayPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesAmazonPayPaymentsParams `form:"amazon_pay_payments" json:"amazon_pay_payments,omitempty"`
	// Allow the merchant to process Australian BECS Direct Debit payments.
	AUBECSDebitPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsParams `form:"au_becs_debit_payments" json:"au_becs_debit_payments,omitempty"`
	// Allow the merchant to process BACS Direct Debit payments.
	BACSDebitPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesBACSDebitPaymentsParams `form:"bacs_debit_payments" json:"bacs_debit_payments,omitempty"`
	// Allow the merchant to process Bancontact payments.
	BancontactPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesBancontactPaymentsParams `form:"bancontact_payments" json:"bancontact_payments,omitempty"`
	// Allow the merchant to process BLIK payments.
	BLIKPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesBLIKPaymentsParams `form:"blik_payments" json:"blik_payments,omitempty"`
	// Allow the merchant to process Boleto payments.
	BoletoPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesBoletoPaymentsParams `form:"boleto_payments" json:"boleto_payments,omitempty"`
	// Allow the merchant to collect card payments.
	CardPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesCardPaymentsParams `form:"card_payments" json:"card_payments,omitempty"`
	// Allow the merchant to process Cartes Bancaires payments.
	CartesBancairesPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesCartesBancairesPaymentsParams `form:"cartes_bancaires_payments" json:"cartes_bancaires_payments,omitempty"`
	// Allow the merchant to process Cash App payments.
	CashAppPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesCashAppPaymentsParams `form:"cashapp_payments" json:"cashapp_payments,omitempty"`
	// Allow the merchant to process EPS payments.
	EPSPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesEPSPaymentsParams `form:"eps_payments" json:"eps_payments,omitempty"`
	// Allow the merchant to process FPX payments.
	FPXPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesFPXPaymentsParams `form:"fpx_payments" json:"fpx_payments,omitempty"`
	// Allow the merchant to process UK bank transfer payments.
	GBBankTransferPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesGBBankTransferPaymentsParams `form:"gb_bank_transfer_payments" json:"gb_bank_transfer_payments,omitempty"`
	// Allow the merchant to process GrabPay payments.
	GrabpayPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesGrabpayPaymentsParams `form:"grabpay_payments" json:"grabpay_payments,omitempty"`
	// Allow the merchant to process iDEAL payments.
	IDEALPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesIDEALPaymentsParams `form:"ideal_payments" json:"ideal_payments,omitempty"`
	// Allow the merchant to process JCB card payments.
	JCBPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesJCBPaymentsParams `form:"jcb_payments" json:"jcb_payments,omitempty"`
	// Allow the merchant to process Japanese bank transfer payments.
	JPBankTransferPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesJPBankTransferPaymentsParams `form:"jp_bank_transfer_payments" json:"jp_bank_transfer_payments,omitempty"`
	// Allow the merchant to process Kakao Pay payments.
	KakaoPayPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesKakaoPayPaymentsParams `form:"kakao_pay_payments" json:"kakao_pay_payments,omitempty"`
	// Allow the merchant to process Klarna payments.
	KlarnaPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesKlarnaPaymentsParams `form:"klarna_payments" json:"klarna_payments,omitempty"`
	// Allow the merchant to process Konbini convenience store payments.
	KonbiniPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesKonbiniPaymentsParams `form:"konbini_payments" json:"konbini_payments,omitempty"`
	// Allow the merchant to process Korean card payments.
	KrCardPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesKrCardPaymentsParams `form:"kr_card_payments" json:"kr_card_payments,omitempty"`
	// Allow the merchant to process Link payments.
	LinkPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesLinkPaymentsParams `form:"link_payments" json:"link_payments,omitempty"`
	// Allow the merchant to process MobilePay payments.
	MobilepayPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesMobilepayPaymentsParams `form:"mobilepay_payments" json:"mobilepay_payments,omitempty"`
	// Allow the merchant to process Multibanco payments.
	MultibancoPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesMultibancoPaymentsParams `form:"multibanco_payments" json:"multibanco_payments,omitempty"`
	// Allow the merchant to process Mexican bank transfer payments.
	MXBankTransferPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesMXBankTransferPaymentsParams `form:"mx_bank_transfer_payments" json:"mx_bank_transfer_payments,omitempty"`
	// Allow the merchant to process Naver Pay payments.
	NaverPayPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesNaverPayPaymentsParams `form:"naver_pay_payments" json:"naver_pay_payments,omitempty"`
	// Allow the merchant to process OXXO payments.
	OXXOPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesOXXOPaymentsParams `form:"oxxo_payments" json:"oxxo_payments,omitempty"`
	// Allow the merchant to process Przelewy24 (P24) payments.
	P24Payments *V2CoreAccountCreateConfigurationMerchantCapabilitiesP24PaymentsParams `form:"p24_payments" json:"p24_payments,omitempty"`
	// Allow the merchant to process Pay by Bank payments.
	PayByBankPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesPayByBankPaymentsParams `form:"pay_by_bank_payments" json:"pay_by_bank_payments,omitempty"`
	// Allow the merchant to process PAYCO payments.
	PaycoPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesPaycoPaymentsParams `form:"payco_payments" json:"payco_payments,omitempty"`
	// Allow the merchant to process PayNow payments.
	PayNowPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesPayNowPaymentsParams `form:"paynow_payments" json:"paynow_payments,omitempty"`
	// Allow the merchant to process PromptPay payments.
	PromptPayPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesPromptPayPaymentsParams `form:"promptpay_payments" json:"promptpay_payments,omitempty"`
	// Allow the merchant to process Revolut Pay payments.
	RevolutPayPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesRevolutPayPaymentsParams `form:"revolut_pay_payments" json:"revolut_pay_payments,omitempty"`
	// Allow the merchant to process Samsung Pay payments.
	SamsungPayPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesSamsungPayPaymentsParams `form:"samsung_pay_payments" json:"samsung_pay_payments,omitempty"`
	// Allow the merchant to process SEPA bank transfer payments.
	SEPABankTransferPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsParams `form:"sepa_bank_transfer_payments" json:"sepa_bank_transfer_payments,omitempty"`
	// Allow the merchant to process SEPA Direct Debit payments.
	SEPADebitPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPADebitPaymentsParams `form:"sepa_debit_payments" json:"sepa_debit_payments,omitempty"`
	// Allow the merchant to process Sunbit payments.
	SunbitPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesSunbitPaymentsParams `form:"sunbit_payments" json:"sunbit_payments,omitempty"`
	// Allow the merchant to process Swish payments.
	SwishPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesSwishPaymentsParams `form:"swish_payments" json:"swish_payments,omitempty"`
	// Allow the merchant to process TWINT payments.
	TWINTPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesTWINTPaymentsParams `form:"twint_payments" json:"twint_payments,omitempty"`
	// Allow the merchant to process US bank transfer payments.
	USBankTransferPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesUSBankTransferPaymentsParams `form:"us_bank_transfer_payments" json:"us_bank_transfer_payments,omitempty"`
	// Allow the merchant to process Zip payments.
	ZipPayments *V2CoreAccountCreateConfigurationMerchantCapabilitiesZipPaymentsParams `form:"zip_payments" json:"zip_payments,omitempty"`
}

// Automatically declines certain charge types regardless of whether the card issuer accepted or declined the charge.
type V2CoreAccountCreateConfigurationMerchantCardPaymentsDeclineOnParams struct {
	// Whether Stripe automatically declines charges with an incorrect ZIP or postal code. This setting only applies when a ZIP or postal code is provided and they fail bank verification.
	AVSFailure *bool `form:"avs_failure" json:"avs_failure,omitempty"`
	// Whether Stripe automatically declines charges with an incorrect CVC. This setting only applies when a CVC is provided and it fails bank verification.
	CVCFailure *bool `form:"cvc_failure" json:"cvc_failure,omitempty"`
}

// Card payments settings.
type V2CoreAccountCreateConfigurationMerchantCardPaymentsParams struct {
	// Automatically declines certain charge types regardless of whether the card issuer accepted or declined the charge.
	DeclineOn *V2CoreAccountCreateConfigurationMerchantCardPaymentsDeclineOnParams `form:"decline_on" json:"decline_on,omitempty"`
}

// Support hours for Konbini payments.
type V2CoreAccountCreateConfigurationMerchantKonbiniPaymentsSupportHoursParams struct {
	// Support hours end time (JST time of day) for in `HH:MM` format.
	EndTime *string `form:"end_time" json:"end_time,omitempty"`
	// Support hours start time (JST time of day) for in `HH:MM` format.
	StartTime *string `form:"start_time" json:"start_time,omitempty"`
}

// Support for Konbini payments.
type V2CoreAccountCreateConfigurationMerchantKonbiniPaymentsSupportParams struct {
	// Support email address for Konbini payments.
	Email *string `form:"email" json:"email,omitempty"`
	// Support hours for Konbini payments.
	Hours *V2CoreAccountCreateConfigurationMerchantKonbiniPaymentsSupportHoursParams `form:"hours" json:"hours,omitempty"`
	// Support phone number for Konbini payments.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Settings specific to Konbini payments on the account.
type V2CoreAccountCreateConfigurationMerchantKonbiniPaymentsParams struct {
	// Support for Konbini payments.
	Support *V2CoreAccountCreateConfigurationMerchantKonbiniPaymentsSupportParams `form:"support" json:"support,omitempty"`
}

// The Kana variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
type V2CoreAccountCreateConfigurationMerchantScriptStatementDescriptorKanaParams struct {
	// The default text that appears on statements for non-card charges outside of Japan. For card charges, if you don't set a statement_descriptor_prefix, this text is also used as the statement descriptor prefix. In that case, if concatenating the statement descriptor suffix causes the combined statement descriptor to exceed 22 characters, we truncate the statement_descriptor text to limit the full descriptor to 22 characters. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Descriptor *string `form:"descriptor" json:"descriptor,omitempty"`
	// Default text that appears on statements for card charges outside of Japan, prefixing any dynamic statement_descriptor_suffix specified on the charge. To maximize space for the dynamic part of the descriptor, keep this text short. If you don't specify this value, statement_descriptor is used as the prefix. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
}

// The Kanji variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
type V2CoreAccountCreateConfigurationMerchantScriptStatementDescriptorKanjiParams struct {
	// The default text that appears on statements for non-card charges outside of Japan. For card charges, if you don't set a statement_descriptor_prefix, this text is also used as the statement descriptor prefix. In that case, if concatenating the statement descriptor suffix causes the combined statement descriptor to exceed 22 characters, we truncate the statement_descriptor text to limit the full descriptor to 22 characters. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Descriptor *string `form:"descriptor" json:"descriptor,omitempty"`
	// Default text that appears on statements for card charges outside of Japan, prefixing any dynamic statement_descriptor_suffix specified on the charge. To maximize space for the dynamic part of the descriptor, keep this text short. If you don't specify this value, statement_descriptor is used as the prefix. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
}

// Settings for the default text that appears on statements for language variations.
type V2CoreAccountCreateConfigurationMerchantScriptStatementDescriptorParams struct {
	// The Kana variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
	Kana *V2CoreAccountCreateConfigurationMerchantScriptStatementDescriptorKanaParams `form:"kana" json:"kana,omitempty"`
	// The Kanji variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
	Kanji *V2CoreAccountCreateConfigurationMerchantScriptStatementDescriptorKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Settings for Smart Disputes auto_respond.
type V2CoreAccountCreateConfigurationMerchantSmartDisputesAutoRespondParams struct {
	// The preference for Smart Disputes auto-respond.
	Preference *string `form:"preference" json:"preference,omitempty"`
}

// Settings used for Smart Disputes.
type V2CoreAccountCreateConfigurationMerchantSmartDisputesParams struct {
	// Settings for Smart Disputes auto_respond.
	AutoRespond *V2CoreAccountCreateConfigurationMerchantSmartDisputesAutoRespondParams `form:"auto_respond" json:"auto_respond,omitempty"`
}

// Statement descriptor.
type V2CoreAccountCreateConfigurationMerchantStatementDescriptorParams struct {
	// The default text that appears on statements for non-card charges outside of Japan. For card charges, if you don't set a statement_descriptor_prefix, this text is also used as the statement descriptor prefix. In that case, if concatenating the statement descriptor suffix causes the combined statement descriptor to exceed 22 characters, we truncate the statement_descriptor text to limit the full descriptor to 22 characters. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Descriptor *string `form:"descriptor" json:"descriptor,omitempty"`
	// Default text that appears on statements for card charges outside of Japan, prefixing any dynamic statement_descriptor_suffix specified on the charge. To maximize space for the dynamic part of the descriptor, keep this text short. If you don't specify this value, statement_descriptor is used as the prefix. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
}

// A publicly available mailing address for sending support issues to.
type V2CoreAccountCreateConfigurationMerchantSupportAddressParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// Publicly available contact information for sending support issues to.
type V2CoreAccountCreateConfigurationMerchantSupportParams struct {
	// A publicly available mailing address for sending support issues to.
	Address *V2CoreAccountCreateConfigurationMerchantSupportAddressParams `form:"address" json:"address,omitempty"`
	// A publicly available email address for sending support issues to.
	Email *string `form:"email" json:"email,omitempty"`
	// A publicly available phone number to call with support issues.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// A publicly available website for handling support issues.
	URL *string `form:"url" json:"url,omitempty"`
}

// Enables the Account to act as a connected account and collect payments facilitated by a Connect platform. You must onboard your platform to Connect before you can add this configuration to your connected accounts. Utilize this configuration when the Account will be the Merchant of Record, like with Direct charges or Destination Charges with on_behalf_of set.
type V2CoreAccountCreateConfigurationMerchantParams struct {
	// Settings used for Bacs debit payments.
	BACSDebitPayments *V2CoreAccountCreateConfigurationMerchantBACSDebitPaymentsParams `form:"bacs_debit_payments" json:"bacs_debit_payments,omitempty"`
	// Settings used to apply the merchant's branding to email receipts, invoices, Checkout, and other products.
	Branding *V2CoreAccountCreateConfigurationMerchantBrandingParams `form:"branding" json:"branding,omitempty"`
	// Capabilities to request on the Merchant Configuration.
	Capabilities *V2CoreAccountCreateConfigurationMerchantCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// Card payments settings.
	CardPayments *V2CoreAccountCreateConfigurationMerchantCardPaymentsParams `form:"card_payments" json:"card_payments,omitempty"`
	// Settings specific to Konbini payments on the account.
	KonbiniPayments *V2CoreAccountCreateConfigurationMerchantKonbiniPaymentsParams `form:"konbini_payments" json:"konbini_payments,omitempty"`
	// The Merchant Category Code (MCC) for the Merchant Configuration. MCCs classify businesses based on the goods or services they provide.
	MCC *string `form:"mcc" json:"mcc,omitempty"`
	// Settings for the default text that appears on statements for language variations.
	ScriptStatementDescriptor *V2CoreAccountCreateConfigurationMerchantScriptStatementDescriptorParams `form:"script_statement_descriptor" json:"script_statement_descriptor,omitempty"`
	// Settings used for Smart Disputes.
	SmartDisputes *V2CoreAccountCreateConfigurationMerchantSmartDisputesParams `form:"smart_disputes" json:"smart_disputes,omitempty"`
	// Statement descriptor.
	StatementDescriptor *V2CoreAccountCreateConfigurationMerchantStatementDescriptorParams `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Publicly available contact information for sending support issues to.
	Support *V2CoreAccountCreateConfigurationMerchantSupportParams `form:"support" json:"support,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in AUD.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in CAD.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in EUR.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in GBP.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in USD.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in USDC.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can receive business storage-type funds on Stripe.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundParams struct {
	// Can receive business storage-type funds on Stripe in AUD.
	AUD *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDParams `form:"aud" json:"aud,omitempty"`
	// Can receive business storage-type funds on Stripe in CAD.
	CAD *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADParams `form:"cad" json:"cad,omitempty"`
	// Can receive business storage-type funds on Stripe in EUR.
	EUR *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURParams `form:"eur" json:"eur,omitempty"`
	// Can receive business storage-type funds on Stripe in GBP.
	GBP *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPParams `form:"gbp" json:"gbp,omitempty"`
	// Can receive business storage-type funds on Stripe in USD.
	USD *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can receive business storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in AUD.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in CAD.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in EUR.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in GBP.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in USD.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in USDC.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send business storage-type funds on Stripe.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundParams struct {
	// Can send business storage-type funds on Stripe in AUD.
	AUD *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDParams `form:"aud" json:"aud,omitempty"`
	// Can send business storage-type funds on Stripe in CAD.
	CAD *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADParams `form:"cad" json:"cad,omitempty"`
	// Can send business storage-type funds on Stripe in EUR.
	EUR *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURParams `form:"eur" json:"eur,omitempty"`
	// Can send business storage-type funds on Stripe in GBP.
	GBP *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPParams `form:"gbp" json:"gbp,omitempty"`
	// Can send business storage-type funds on Stripe in USD.
	USD *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can send business storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Can send or receive business storage-type funds on Stripe.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageParams struct {
	// Can receive business storage-type funds on Stripe.
	Inbound *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundParams `form:"inbound" json:"inbound,omitempty"`
	// Can send business storage-type funds on Stripe.
	Outbound *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundParams `form:"outbound" json:"outbound,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive consumer storage-type funds on Stripe in USD.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive consumer storage-type funds on Stripe in USDC.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can receive consumer storage-type funds on Stripe.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundParams struct {
	// Can receive consumer storage-type funds on Stripe in USD.
	USD *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can receive consumer storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send consumer storage-type funds on Stripe in USD.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send consumer storage-type funds on Stripe in USDC.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send consumer storage-type funds on Stripe.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundParams struct {
	// Can send consumer storage-type funds on Stripe in USD.
	USD *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can send consumer storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Can send or receive consumer storage-type funds on Stripe.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageParams struct {
	// Can receive consumer storage-type funds on Stripe.
	Inbound *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundParams `form:"inbound" json:"inbound,omitempty"`
	// Can send consumer storage-type funds on Stripe.
	Outbound *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundParams `form:"outbound" json:"outbound,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesInboundTransfersParams struct {
	// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
	BankAccounts *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a bank account owned by someone else.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a debit card owned by someone else.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to someone else via paper check.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send funds from a FinancialAccount to a destination owned by someone else.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by someone else.
	BankAccounts *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a debit card owned by someone else.
	Cards *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsParams `form:"cards" json:"cards,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
	CryptoWallets *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
	FinancialAccounts *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
	// Can send funds from a FinancialAccount to someone else via paper check.
	PaperChecks *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksParams `form:"paper_checks" json:"paper_checks,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a bank account owned by yourself.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send funds from a FinancialAccount to a destination owned by yourself.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by yourself.
	BankAccounts *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
	CryptoWallets *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
	FinancialAccounts *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive funds on a bank-account-like financial address (VBAN) to credit a FinancialAccount.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive funds on a crypto wallet like financial address to credit a FinancialAccount.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can receive funds into a FinancialAccount.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsParams struct {
	// Can receive funds on a bank-account-like financial address (VBAN) to credit a FinancialAccount.
	BankAccounts *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can receive funds on a crypto wallet like financial address to credit a FinancialAccount.
	CryptoWallets *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive debits to a FinancialAccount from a bank account.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can receive debits to a FinancialAccount.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedDebitsParams struct {
	// Can receive debits to a FinancialAccount from a bank account.
	BankAccounts *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
}

// Capabilities to request on the Money Manager Configuration.
type V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesParams struct {
	// Can send or receive business storage-type funds on Stripe.
	BusinessStorage *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesBusinessStorageParams `form:"business_storage" json:"business_storage,omitempty"`
	// Can send or receive consumer storage-type funds on Stripe.
	ConsumerStorage *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesConsumerStorageParams `form:"consumer_storage" json:"consumer_storage,omitempty"`
	// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
	InboundTransfers *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesInboundTransfersParams `form:"inbound_transfers" json:"inbound_transfers,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by someone else.
	OutboundPayments *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundPaymentsParams `form:"outbound_payments" json:"outbound_payments,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by yourself.
	OutboundTransfers *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesOutboundTransfersParams `form:"outbound_transfers" json:"outbound_transfers,omitempty"`
	// Can receive funds into a FinancialAccount.
	ReceivedCredits *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedCreditsParams `form:"received_credits" json:"received_credits,omitempty"`
	// Can receive debits to a FinancialAccount.
	ReceivedDebits *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesReceivedDebitsParams `form:"received_debits" json:"received_debits,omitempty"`
}

// Details of the regulated activity if the business participates in one.
type V2CoreAccountCreateConfigurationMoneyManagerRegulatedActivityParams struct {
	// A detailed description of the regulated activities the business is licensed to conduct.
	Description *string `form:"description" json:"description,omitempty"`
	// The license number or registration number assigned by the business's primary regulator.
	LicenseNumber *string `form:"license_number" json:"license_number,omitempty"`
	// The country of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityCountry *string `form:"primary_regulatory_authority_country" json:"primary_regulatory_authority_country,omitempty"`
	// The name of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityName *string `form:"primary_regulatory_authority_name" json:"primary_regulatory_authority_name,omitempty"`
}

// The Money Manager Configuration allows the Account to store and move funds using FinancialAccounts.
type V2CoreAccountCreateConfigurationMoneyManagerParams struct {
	// Capabilities to request on the Money Manager Configuration.
	Capabilities *V2CoreAccountCreateConfigurationMoneyManagerCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// List of high-risk activities the business is involved in.
	HighRiskActivities []*string `form:"high_risk_activities" json:"high_risk_activities,omitempty"`
	// Description of the high-risk activities the business offers.
	HighRiskActivitiesDescription *string `form:"high_risk_activities_description" json:"high_risk_activities_description,omitempty"`
	// Description of the money services offered by the business.
	MoneyServicesDescription *string `form:"money_services_description" json:"money_services_description,omitempty"`
	// Indicates whether the business operates in any prohibited countries.
	OperatesInProhibitedCountries *bool `form:"operates_in_prohibited_countries" json:"operates_in_prohibited_countries,omitempty"`
	// Indicates whether the business participates in any regulated activity.
	ParticipatesInRegulatedActivity *bool `form:"participates_in_regulated_activity" json:"participates_in_regulated_activity,omitempty"`
	// Primary purpose of the stored funds.
	PurposeOfFunds *string `form:"purpose_of_funds" json:"purpose_of_funds,omitempty"`
	// Description of the purpose of the stored funds.
	PurposeOfFundsDescription *string `form:"purpose_of_funds_description" json:"purpose_of_funds_description,omitempty"`
	// Details of the regulated activity if the business participates in one.
	RegulatedActivity *V2CoreAccountCreateConfigurationMoneyManagerRegulatedActivityParams `form:"regulated_activity" json:"regulated_activity,omitempty"`
	// The source of funds for the business, e.g. profits, income, venture capital, etc.
	SourceOfFunds *string `form:"source_of_funds" json:"source_of_funds,omitempty"`
	// Description of the source of funds for the business' account.
	SourceOfFundsDescription *string `form:"source_of_funds_description" json:"source_of_funds_description,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsACHProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsACHProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsACHProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over ACH rails.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsACHParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsACHProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over BECS rails.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsBECSParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsEftProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsEftProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsEftProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over EFT rails.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsEftParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsEftProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over Fedwire or CHIPS.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFedwireParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over FPS rails.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFPSParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsInstantParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsLocalParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsNppProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsNppProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsNppProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over NPP (real time) rails.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsNppParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsNppProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over RTP rails.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsRTPParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA credit rails.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPACreditParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA instant (real time) rails.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over SWIFT.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSwiftParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsWireProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsWireProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsWireProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsWireParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsWireProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Capabilities that enable OutboundPayments to a bank account linked to this Account.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsParams struct {
	// Enables this Account to receive OutboundPayments to linked bank accounts over ACH rails.
	ACH *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsACHParams `form:"ach" json:"ach,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over BECS rails.
	BECS *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsBECSParams `form:"becs" json:"becs,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over EFT rails.
	Eft *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsEftParams `form:"eft" json:"eft,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over Fedwire or CHIPS.
	Fedwire *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFedwireParams `form:"fedwire" json:"fedwire,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over FPS rails.
	FPS *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsFPSParams `form:"fps" json:"fps,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
	Instant *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsInstantParams `form:"instant" json:"instant,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
	Local *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsLocalParams `form:"local" json:"local,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over NPP (real time) rails.
	Npp *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsNppParams `form:"npp" json:"npp,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over RTP rails.
	RTP *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsRTPParams `form:"rtp" json:"rtp,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA credit rails.
	SEPACredit *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPACreditParams `form:"sepa_credit" json:"sepa_credit,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA instant (real time) rails.
	SEPAInstant *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantParams `form:"sepa_instant" json:"sepa_instant,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over SWIFT.
	Swift *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsSwiftParams `form:"swift" json:"swift,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
	Wire *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsWireParams `form:"wire" json:"wire,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesCardsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesCardsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesCardsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Capabilities that enable OutboundPayments to a card linked to this Account.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesCardsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesCardsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Capabilities that enable OutboundPayments to a crypto wallet linked to this Account.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesPaperChecksProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesPaperChecksProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesPaperChecksProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Capabilities that enable OutboundPayments via paper check.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesPaperChecksParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesPaperChecksProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive /v1/transfers into their Stripe Balance (/v1/balance).
type V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Capabilities that enable the recipient to manage their Stripe Balance (/v1/balance).
type V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceParams struct {
	// Enables this Account to receive /v1/transfers into their Stripe Balance (/v1/balance).
	StripeTransfers *V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams `form:"stripe_transfers" json:"stripe_transfers,omitempty"`
}

// Capabilities to be requested on the Recipient Configuration.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesParams struct {
	// Capabilities that enable OutboundPayments to a bank account linked to this Account.
	BankAccounts *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Capabilities that enable OutboundPayments to a card linked to this Account.
	Cards *V2CoreAccountCreateConfigurationRecipientCapabilitiesCardsParams `form:"cards" json:"cards,omitempty"`
	// Capabilities that enable OutboundPayments to a crypto wallet linked to this Account.
	CryptoWallets *V2CoreAccountCreateConfigurationRecipientCapabilitiesCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Capabilities that enable OutboundPayments via paper check.
	PaperChecks *V2CoreAccountCreateConfigurationRecipientCapabilitiesPaperChecksParams `form:"paper_checks" json:"paper_checks,omitempty"`
	// Capabilities that enable the recipient to manage their Stripe Balance (/v1/balance).
	StripeBalance *V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceParams `form:"stripe_balance" json:"stripe_balance,omitempty"`
}

// The Recipient Configuration allows the Account to receive funds. Utilize this configuration if the Account will not be the Merchant of Record, like with Separate Charges & Transfers, or Destination Charges without on_behalf_of set.
type V2CoreAccountCreateConfigurationRecipientParams struct {
	// Capabilities to be requested on the Recipient Configuration.
	Capabilities *V2CoreAccountCreateConfigurationRecipientCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
}

// An Account Configuration which allows the Account to take on a key persona across Stripe products.
type V2CoreAccountCreateConfigurationParams struct {
	// The CardCreator Configuration allows the Account to create and issue cards to users.
	CardCreator *V2CoreAccountCreateConfigurationCardCreatorParams `form:"card_creator" json:"card_creator,omitempty"`
	// The Customer Configuration allows the Account to be used in inbound payment flows (i.e. customer-facing payment and billing flows).
	Customer *V2CoreAccountCreateConfigurationCustomerParams `form:"customer" json:"customer,omitempty"`
	// Enables the Account to act as a connected account and collect payments facilitated by a Connect platform. You must onboard your platform to Connect before you can add this configuration to your connected accounts. Utilize this configuration when the Account will be the Merchant of Record, like with Direct charges or Destination Charges with on_behalf_of set.
	Merchant *V2CoreAccountCreateConfigurationMerchantParams `form:"merchant" json:"merchant,omitempty"`
	// The Money Manager Configuration allows the Account to store and move funds using FinancialAccounts.
	MoneyManager *V2CoreAccountCreateConfigurationMoneyManagerParams `form:"money_manager" json:"money_manager,omitempty"`
	// The Recipient Configuration allows the Account to receive funds. Utilize this configuration if the Account will not be the Merchant of Record, like with Separate Charges & Transfers, or Destination Charges without on_behalf_of set.
	Recipient *V2CoreAccountCreateConfigurationRecipientParams `form:"recipient" json:"recipient,omitempty"`
}

// Account profile information.
type V2CoreAccountCreateDefaultsProfileParams struct {
	// The business's publicly-available website.
	BusinessURL *string `form:"business_url" json:"business_url,omitempty"`
	// The name which is used by the business.
	DoingBusinessAs *string `form:"doing_business_as" json:"doing_business_as,omitempty"`
	// Internal-only description of the product sold or service provided by the business. It's used by Stripe for risk and underwriting purposes.
	ProductDescription *string `form:"product_description" json:"product_description,omitempty"`
}

// Default responsibilities held by either Stripe or the platform.
type V2CoreAccountCreateDefaultsResponsibilitiesParams struct {
	// A value indicating the party responsible for collecting fees from this account.
	FeesCollector *string `form:"fees_collector" json:"fees_collector"`
	// A value indicating who is responsible for losses when this Account can't pay back negative balances from payments.
	LossesCollector *string `form:"losses_collector" json:"losses_collector"`
	// A value indicating the party responsible for collecting requirements on this account.
	RequirementsCollector *string `form:"requirements_collector" json:"requirements_collector,omitempty"`
}

// Default values to be used on Account Configurations.
type V2CoreAccountCreateDefaultsParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The Account's preferred locales (languages), ordered by preference.
	Locales []*string `form:"locales" json:"locales,omitempty"`
	// Account profile information.
	Profile *V2CoreAccountCreateDefaultsProfileParams `form:"profile" json:"profile,omitempty"`
	// Default responsibilities held by either Stripe or the platform.
	Responsibilities *V2CoreAccountCreateDefaultsResponsibilitiesParams `form:"responsibilities" json:"responsibilities,omitempty"`
	// The Account's local timezone. A list of possible time zone values is maintained at the [IANA Time Zone Database](https://www.iana.org/time-zones).
	Timezone *string `form:"timezone" json:"timezone,omitempty"`
}

// This hash is used to attest that the directors information provided to Stripe is both current and correct.
type V2CoreAccountCreateIdentityAttestationsDirectorshipDeclarationParams struct {
	// The time marking when the director attestation was made. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the director attestation was made.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the director attestation was made.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
type V2CoreAccountCreateIdentityAttestationsOwnershipDeclarationParams struct {
	// The time marking when the beneficial owner attestation was made. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the beneficial owner attestation was made.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the beneficial owner attestation was made.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Attestation that all Persons with a specific Relationship value have been provided.
type V2CoreAccountCreateIdentityAttestationsPersonsProvidedParams struct {
	// Whether the company's directors have been provided. Set this Boolean to true after creating all the company's directors with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Directors *bool `form:"directors" json:"directors,omitempty"`
	// Whether the company's executives have been provided. Set this Boolean to true after creating all the company's executives with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Executives *bool `form:"executives" json:"executives,omitempty"`
	// Whether the company's owners have been provided. Set this Boolean to true after creating all the company's owners with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Owners *bool `form:"owners" json:"owners,omitempty"`
	// Reason for why the company is exempt from providing ownership information.
	OwnershipExemptionReason *string `form:"ownership_exemption_reason" json:"ownership_exemption_reason,omitempty"`
}

// This hash is used to attest that the representative is authorized to act as the representative of their legal entity.
type V2CoreAccountCreateIdentityAttestationsRepresentativeDeclarationParams struct {
	// The time marking when the representative attestation was made. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the representative attestation was made.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the representative attestation was made.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of the [Stripe Services Agreement](https://docs.stripe.com/connect/updating-accounts#tos-acceptance).
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceAccountParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for Stripe commercial card issuing.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardParams struct {
	// Bank terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Platform terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
	Platform *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardParams struct {
	// Bank terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
	FinancingDisclosures *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
	Platform *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Celtic as BIN sponsor.
	ApplePay *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing charge cards with Celtic as BIN sponsor.
	ChargeCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Terms of service acceptances for commercial issuing spend cards with Celtic as BIN sponsor.
	SpendCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardParams struct {
	// Bank terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	FinancingDisclosures *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	Platform *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardParams struct {
	// Bank terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Platform terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	Platform *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardParams struct {
	// Bank terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	FinancingDisclosures *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Cross River Bank as BIN sponsor.
	ApplePay *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	ChargeCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Terms of service acceptances for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
	// Terms of service acceptances for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Fifth Third as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardParams struct {
	// Bank terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
	FinancingDisclosures *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
	Platform *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Fifth Third as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Fifth Third as BIN sponsor.
	ApplePay *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing charge cards with Fifth Third as BIN sponsor.
	ChargeCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
}

// Terms of service acceptances for Stripe commercial card Global issuing.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialGlobalAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardParams struct {
	// Bank terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Platform terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
	Platform *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Lead as BIN sponsor.
	ApplePay *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Terms of service acceptances to create cards for commercial issuing use cases.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialParams struct {
	// Terms of service acceptances for Stripe commercial card issuing.
	AccountHolder *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialAccountHolderParams `form:"account_holder" json:"account_holder,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Fifth Third as BIN sponsor.
	FifthThird *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdParams `form:"fifth_third" json:"fifth_third,omitempty"`
	// Terms of service acceptances for Stripe commercial card Global issuing.
	GlobalAccountHolder *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialGlobalAccountHolderParams `form:"global_account_holder" json:"global_account_holder,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadParams `form:"lead" json:"lead,omitempty"`
}

// Terms of service acceptances for Stripe consumer card issuing.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing Apple Pay cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardParams struct {
	// Bank terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
	FinancingDisclosures *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
	Platform *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for consumer issuing cards with Celtic as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticParams struct {
	// Terms of service acceptances for consumer issuing Apple Pay cards with Celtic as BIN sponsor.
	ApplePay *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for consumer issuing charge cards with Celtic as BIN sponsor.
	RevolvingCreditCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardParams `form:"revolving_credit_card" json:"revolving_credit_card,omitempty"`
}

// Bank terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardParams struct {
	// Bank terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	FinancingDisclosures *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	Platform *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for consumer issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankParams struct {
	// Terms of service acceptances for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Terms of service acceptances for Stripe consumer card Global issuing.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerGlobalAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardParams struct {
	// Bank terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
	FinancingDisclosures *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
	Platform *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardParams struct {
	// Bank terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
	BankTerms *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
	FinancingDisclosures *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
	Platform *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for consumer issuing cards with Lead as BIN sponsor.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadParams struct {
	// Terms of service acceptances for consumer issuing debit cards with Lead as BIN sponsor.
	DebitCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardParams `form:"debit_card" json:"debit_card,omitempty"`
	// Terms of service acceptances for consumer issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Terms of service acceptances to create cards for consumer issuing use cases.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerParams struct {
	// Terms of service acceptances for Stripe consumer card issuing.
	AccountHolder *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerAccountHolderParams `form:"account_holder" json:"account_holder,omitempty"`
	// Terms of service acceptances for consumer issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Terms of service acceptances for consumer issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Terms of service acceptances for Stripe consumer card Global issuing.
	GlobalAccountHolder *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerGlobalAccountHolderParams `form:"global_account_holder" json:"global_account_holder,omitempty"`
	// Terms of service acceptances for consumer issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadParams `form:"lead" json:"lead,omitempty"`
}

// Details on the Account's acceptance of Issuing-specific terms of service.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorParams struct {
	// Terms of service acceptances to create cards for commercial issuing use cases.
	Commercial *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialParams `form:"commercial" json:"commercial,omitempty"`
	// Terms of service acceptances to create cards for consumer issuing use cases.
	Consumer *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorConsumerParams `form:"consumer" json:"consumer,omitempty"`
}

// Details on the Account's acceptance of Consumer-specific terms of service.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceConsumerMoneyManagerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Consumer-privacy-disclosures-specific terms of service.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceConsumerPrivacyDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Crypto-specific terms of service.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCryptoMoneyManagerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Treasury-specific terms of service.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceMoneyManagerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceParams struct {
	// Details on the Account's acceptance of the [Stripe Services Agreement](https://docs.stripe.com/connect/updating-accounts#tos-acceptance).
	Account *V2CoreAccountCreateIdentityAttestationsTermsOfServiceAccountParams `form:"account" json:"account,omitempty"`
	// Details on the Account's acceptance of Issuing-specific terms of service.
	CardCreator *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorParams `form:"card_creator" json:"card_creator,omitempty"`
	// Details on the Account's acceptance of Consumer-specific terms of service.
	ConsumerMoneyManager *V2CoreAccountCreateIdentityAttestationsTermsOfServiceConsumerMoneyManagerParams `form:"consumer_money_manager" json:"consumer_money_manager,omitempty"`
	// Details on the Account's acceptance of Consumer-privacy-disclosures-specific terms of service.
	ConsumerPrivacyDisclosures *V2CoreAccountCreateIdentityAttestationsTermsOfServiceConsumerPrivacyDisclosuresParams `form:"consumer_privacy_disclosures" json:"consumer_privacy_disclosures,omitempty"`
	// Details on the Account's acceptance of Crypto-specific terms of service.
	CryptoMoneyManager *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCryptoMoneyManagerParams `form:"crypto_money_manager" json:"crypto_money_manager,omitempty"`
	// Details on the Account's acceptance of Treasury-specific terms of service.
	MoneyManager *V2CoreAccountCreateIdentityAttestationsTermsOfServiceMoneyManagerParams `form:"money_manager" json:"money_manager,omitempty"`
}

// Attestations from the identity's key people, e.g. owners, executives, directors, representatives.
type V2CoreAccountCreateIdentityAttestationsParams struct {
	// This hash is used to attest that the directors information provided to Stripe is both current and correct.
	DirectorshipDeclaration *V2CoreAccountCreateIdentityAttestationsDirectorshipDeclarationParams `form:"directorship_declaration" json:"directorship_declaration,omitempty"`
	// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
	OwnershipDeclaration *V2CoreAccountCreateIdentityAttestationsOwnershipDeclarationParams `form:"ownership_declaration" json:"ownership_declaration,omitempty"`
	// Attestation that all Persons with a specific Relationship value have been provided.
	PersonsProvided *V2CoreAccountCreateIdentityAttestationsPersonsProvidedParams `form:"persons_provided" json:"persons_provided,omitempty"`
	// This hash is used to attest that the representative is authorized to act as the representative of their legal entity.
	RepresentativeDeclaration *V2CoreAccountCreateIdentityAttestationsRepresentativeDeclarationParams `form:"representative_declaration" json:"representative_declaration,omitempty"`
	// Attestations of accepted terms of service agreements.
	TermsOfService *V2CoreAccountCreateIdentityAttestationsTermsOfServiceParams `form:"terms_of_service" json:"terms_of_service,omitempty"`
}

// The business registration address of the business entity.
type V2CoreAccountCreateIdentityBusinessDetailsAddressParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// The business gross annual revenue for its preceding fiscal year.
type V2CoreAccountCreateIdentityBusinessDetailsAnnualRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// The close-out date of the preceding fiscal year in ISO 8601 format. E.g. 2023-12-31 for the 31st of December, 2023.
	FiscalYearEnd *string `form:"fiscal_year_end" json:"fiscal_year_end,omitempty"`
}

// One or more documents that support the bank account ownership verification requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a check.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsBankAccountOwnershipVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's license to operate.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsCompanyLicenseParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the company's Memorandum of Association.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsCompanyMemorandumOfAssociationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Certain countries only: One or more documents showing the ministerial decree legalizing the company's establishment.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsCompanyMinisterialDecreeParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsCompanyRegistrationVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's tax ID.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsCompanyTaxIDVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front"`
}

// A document verifying the business.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountCreateIdentityBusinessDetailsDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of address.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfAddressParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Person that is signing the document.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfRegistrationSignerParams struct {
	// Person signing the document.
	Person *string `form:"person" json:"person"`
}

// One or more documents showing the company's proof of registration with the national business registry.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfRegistrationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// Person that is signing the document.
	Signer *V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfRegistrationSignerParams `form:"signer" json:"signer,omitempty"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Person that is signing the document.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipSignerParams struct {
	// Person signing the document.
	Person *string `form:"person" json:"person"`
}

// One or more documents that demonstrate proof of ultimate beneficial ownership.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// Person that is signing the document.
	Signer *V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipSignerParams `form:"signer" json:"signer,omitempty"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// A document verifying the business.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsParams struct {
	// One or more documents that support the bank account ownership verification requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a check.
	BankAccountOwnershipVerification *V2CoreAccountCreateIdentityBusinessDetailsDocumentsBankAccountOwnershipVerificationParams `form:"bank_account_ownership_verification" json:"bank_account_ownership_verification,omitempty"`
	// One or more documents that demonstrate proof of a company's license to operate.
	CompanyLicense *V2CoreAccountCreateIdentityBusinessDetailsDocumentsCompanyLicenseParams `form:"company_license" json:"company_license,omitempty"`
	// One or more documents showing the company's Memorandum of Association.
	CompanyMemorandumOfAssociation *V2CoreAccountCreateIdentityBusinessDetailsDocumentsCompanyMemorandumOfAssociationParams `form:"company_memorandum_of_association" json:"company_memorandum_of_association,omitempty"`
	// Certain countries only: One or more documents showing the ministerial decree legalizing the company's establishment.
	CompanyMinisterialDecree *V2CoreAccountCreateIdentityBusinessDetailsDocumentsCompanyMinisterialDecreeParams `form:"company_ministerial_decree" json:"company_ministerial_decree,omitempty"`
	// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
	CompanyRegistrationVerification *V2CoreAccountCreateIdentityBusinessDetailsDocumentsCompanyRegistrationVerificationParams `form:"company_registration_verification" json:"company_registration_verification,omitempty"`
	// One or more documents that demonstrate proof of a company's tax ID.
	CompanyTaxIDVerification *V2CoreAccountCreateIdentityBusinessDetailsDocumentsCompanyTaxIDVerificationParams `form:"company_tax_id_verification" json:"company_tax_id_verification,omitempty"`
	// A document verifying the business.
	PrimaryVerification *V2CoreAccountCreateIdentityBusinessDetailsDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// One or more documents that demonstrate proof of address.
	ProofOfAddress *V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfAddressParams `form:"proof_of_address" json:"proof_of_address,omitempty"`
	// One or more documents showing the company's proof of registration with the national business registry.
	ProofOfRegistration *V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfRegistrationParams `form:"proof_of_registration" json:"proof_of_registration,omitempty"`
	// One or more documents that demonstrate proof of ultimate beneficial ownership.
	ProofOfUltimateBeneficialOwnership *V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams `form:"proof_of_ultimate_beneficial_ownership" json:"proof_of_ultimate_beneficial_ownership,omitempty"`
}

// The ID numbers of a business entity.
type V2CoreAccountCreateIdentityBusinessDetailsIDNumberParams struct {
	// The registrar of the ID number (Only valid for DE ID number types).
	Registrar *string `form:"registrar" json:"registrar,omitempty"`
	// Open Enum. The ID number type of a business entity.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// An estimate of the monthly revenue of the business.
type V2CoreAccountCreateIdentityBusinessDetailsMonthlyEstimatedRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
}

// When the business was incorporated or registered.
type V2CoreAccountCreateIdentityBusinessDetailsRegistrationDateParams struct {
	// The day of registration, between 1 and 31.
	Day *int64 `form:"day" json:"day"`
	// The month of registration, between 1 and 12.
	Month *int64 `form:"month" json:"month"`
	// The four-digit year of registration.
	Year *int64 `form:"year" json:"year"`
}

// Kana Address.
type V2CoreAccountCreateIdentityBusinessDetailsScriptAddressesKanaParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountCreateIdentityBusinessDetailsScriptAddressesKanjiParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// The business registration address of the business entity in non latin script.
type V2CoreAccountCreateIdentityBusinessDetailsScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountCreateIdentityBusinessDetailsScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountCreateIdentityBusinessDetailsScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Kana name.
type V2CoreAccountCreateIdentityBusinessDetailsScriptNamesKanaParams struct {
	// Registered name of the business.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// Kanji name.
type V2CoreAccountCreateIdentityBusinessDetailsScriptNamesKanjiParams struct {
	// Registered name of the business.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// The business legal name in non latin script.
type V2CoreAccountCreateIdentityBusinessDetailsScriptNamesParams struct {
	// Kana name.
	Kana *V2CoreAccountCreateIdentityBusinessDetailsScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji name.
	Kanji *V2CoreAccountCreateIdentityBusinessDetailsScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Information about the company or business.
type V2CoreAccountCreateIdentityBusinessDetailsParams struct {
	// The business registration address of the business entity.
	Address *V2CoreAccountCreateIdentityBusinessDetailsAddressParams `form:"address" json:"address,omitempty"`
	// The business gross annual revenue for its preceding fiscal year.
	AnnualRevenue *V2CoreAccountCreateIdentityBusinessDetailsAnnualRevenueParams `form:"annual_revenue" json:"annual_revenue,omitempty"`
	// A detailed description of the business's compliance and anti-money laundering controls and practices.
	ComplianceScreeningDescription *string `form:"compliance_screening_description" json:"compliance_screening_description,omitempty"`
	// A document verifying the business.
	Documents *V2CoreAccountCreateIdentityBusinessDetailsDocumentsParams `form:"documents" json:"documents,omitempty"`
	// Estimated maximum number of workers currently engaged by the business (including employees, contractors, and vendors).
	EstimatedWorkerCount *int64 `form:"estimated_worker_count" json:"estimated_worker_count,omitempty"`
	// The ID numbers of a business entity.
	IDNumbers []*V2CoreAccountCreateIdentityBusinessDetailsIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// An estimate of the monthly revenue of the business.
	MonthlyEstimatedRevenue *V2CoreAccountCreateIdentityBusinessDetailsMonthlyEstimatedRevenueParams `form:"monthly_estimated_revenue" json:"monthly_estimated_revenue,omitempty"`
	// The phone number of the Business Entity.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The business legal name.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
	// When the business was incorporated or registered.
	RegistrationDate *V2CoreAccountCreateIdentityBusinessDetailsRegistrationDateParams `form:"registration_date" json:"registration_date,omitempty"`
	// The business registration address of the business entity in non latin script.
	ScriptAddresses *V2CoreAccountCreateIdentityBusinessDetailsScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The business legal name in non latin script.
	ScriptNames *V2CoreAccountCreateIdentityBusinessDetailsScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The category identifying the legal structure of the business.
	Structure *string `form:"structure" json:"structure,omitempty"`
}

// Additional addresses associated with the individual.
type V2CoreAccountCreateIdentityIndividualAdditionalAddressParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// Additional names (e.g. aliases) associated with the individual.
type V2CoreAccountCreateIdentityIndividualAdditionalNameParams struct {
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
type V2CoreAccountCreateIdentityIndividualAddressParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// The individual's date of birth.
type V2CoreAccountCreateIdentityIndividualDateOfBirthParams struct {
	// The day of birth.
	Day *int64 `form:"day" json:"day"`
	// The month of birth.
	Month *int64 `form:"month" json:"month"`
	// The year of birth.
	Year *int64 `form:"year" json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountCreateIdentityIndividualDocumentsCompanyAuthorizationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountCreateIdentityIndividualDocumentsPassportParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountCreateIdentityIndividualDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountCreateIdentityIndividualDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountCreateIdentityIndividualDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountCreateIdentityIndividualDocumentsSecondaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountCreateIdentityIndividualDocumentsSecondaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountCreateIdentityIndividualDocumentsSecondaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountCreateIdentityIndividualDocumentsVisaParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountCreateIdentityIndividualDocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountCreateIdentityIndividualDocumentsCompanyAuthorizationParams `form:"company_authorization" json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountCreateIdentityIndividualDocumentsPassportParams `form:"passport" json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountCreateIdentityIndividualDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountCreateIdentityIndividualDocumentsSecondaryVerificationParams `form:"secondary_verification" json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountCreateIdentityIndividualDocumentsVisaParams `form:"visa" json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the individual.
type V2CoreAccountCreateIdentityIndividualIDNumberParams struct {
	// The ID number type of an individual.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// The relationship that this individual has with the account's identity.
type V2CoreAccountCreateIdentityIndividualRelationshipParams struct {
	// Whether the person is a director of the account's identity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.
	Director *bool `form:"director" json:"director,omitempty"`
	// Whether the person has significant responsibility to control, manage, or direct the organization.
	Executive *bool `form:"executive" json:"executive,omitempty"`
	// Whether the person is an owner of the account's identity.
	Owner *bool `form:"owner" json:"owner,omitempty"`
	// The percent owned by the person of the account's legal entity.
	PercentOwnership *float64 `form:"percent_ownership,high_precision" json:"percent_ownership,string,omitempty"`
	// The person's title (e.g., CEO, Support Engineer).
	Title *string `form:"title" json:"title,omitempty"`
}

// Kana Address.
type V2CoreAccountCreateIdentityIndividualScriptAddressesKanaParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// Kanji Address.
type V2CoreAccountCreateIdentityIndividualScriptAddressesKanjiParams struct {
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
	// Town or district.
	Town *string `form:"town" json:"town,omitempty"`
}

// The script addresses (e.g., non-Latin characters) associated with the individual.
type V2CoreAccountCreateIdentityIndividualScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountCreateIdentityIndividualScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountCreateIdentityIndividualScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountCreateIdentityIndividualScriptNamesKanaParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountCreateIdentityIndividualScriptNamesKanjiParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The individuals primary name in non latin script.
type V2CoreAccountCreateIdentityIndividualScriptNamesParams struct {
	// Persons name in kana script.
	Kana *V2CoreAccountCreateIdentityIndividualScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountCreateIdentityIndividualScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Information about the person represented by the account.
type V2CoreAccountCreateIdentityIndividualParams struct {
	// Additional addresses associated with the individual.
	AdditionalAddresses []*V2CoreAccountCreateIdentityIndividualAdditionalAddressParams `form:"additional_addresses" json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the individual.
	AdditionalNames []*V2CoreAccountCreateIdentityIndividualAdditionalNameParams `form:"additional_names" json:"additional_names,omitempty"`
	// The individual's residential address.
	Address *V2CoreAccountCreateIdentityIndividualAddressParams `form:"address" json:"address,omitempty"`
	// The individual's date of birth.
	DateOfBirth *V2CoreAccountCreateIdentityIndividualDateOfBirthParams `form:"date_of_birth" json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountCreateIdentityIndividualDocumentsParams `form:"documents" json:"documents,omitempty"`
	// The individual's email address. You can only set this field when the Account is configured as a `merchant` or `recipient`. Use `contact_email` as the primary contact email for this Account.
	Email *string `form:"email" json:"email,omitempty"`
	// The individual's first name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The identification numbers (e.g., SSN) associated with the individual.
	IDNumbers []*V2CoreAccountCreateIdentityIndividualIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// The individual's gender (International regulations require either "male" or "female").
	LegalGender *string `form:"legal_gender" json:"legal_gender,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The countries where the individual is a national. Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Nationalities []*string `form:"nationalities" json:"nationalities,omitempty"`
	// The individual's phone number.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The individual's political exposure.
	PoliticalExposure *string `form:"political_exposure" json:"political_exposure,omitempty"`
	// The relationship that this individual has with the account's identity.
	Relationship *V2CoreAccountCreateIdentityIndividualRelationshipParams `form:"relationship" json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the individual.
	ScriptAddresses *V2CoreAccountCreateIdentityIndividualScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The individuals primary name in non latin script.
	ScriptNames *V2CoreAccountCreateIdentityIndividualScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The individual's last name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountCreateIdentityIndividualParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Information about the company, individual, and business represented by the Account.
type V2CoreAccountCreateIdentityParams struct {
	// Attestations from the identity's key people, e.g. owners, executives, directors, representatives.
	Attestations *V2CoreAccountCreateIdentityAttestationsParams `form:"attestations" json:"attestations,omitempty"`
	// Information about the company or business.
	BusinessDetails *V2CoreAccountCreateIdentityBusinessDetailsParams `form:"business_details" json:"business_details,omitempty"`
	// The country in which the account holder resides, or in which the business is legally established. This should be an [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code.
	Country *string `form:"country" json:"country,omitempty"`
	// The entity type represented by the Account. Ensure this field is accurate before adding configurations that rely on identity information, as it determines which identity fields apply and how the Account is validated.
	EntityType *string `form:"entity_type" json:"entity_type,omitempty"`
	// Information about the person represented by the account.
	Individual *V2CoreAccountCreateIdentityIndividualParams `form:"individual" json:"individual,omitempty"`
}

// Create an Account that represents a company, individual, or other entity that your business interacts with. Accounts contain identifying information about the entity, and configurations that store the features an account has access to. An account can be configured as any or all of the following configurations: Customer, Merchant and/or Recipient.
type V2CoreAccountCreateParams struct {
	Params `form:"*"`
	// The account token generated by the account token api.
	AccountToken *string `form:"account_token" json:"account_token,omitempty"`
	// An Account Configuration which allows the Account to take on a key persona across Stripe products.
	Configuration *V2CoreAccountCreateConfigurationParams `form:"configuration" json:"configuration,omitempty"`
	// The primary contact email address for the Account.
	ContactEmail *string `form:"contact_email" json:"contact_email,omitempty"`
	// The default contact phone for the Account.
	ContactPhone *string `form:"contact_phone" json:"contact_phone,omitempty"`
	// A value indicating the Stripe dashboard this Account has access to. This will depend on which configurations are enabled for this account.
	Dashboard *string `form:"dashboard" json:"dashboard,omitempty"`
	// Default values to be used on Account Configurations.
	Defaults *V2CoreAccountCreateDefaultsParams `form:"defaults" json:"defaults,omitempty"`
	// A descriptive name for the Account. This name will be surfaced in the Stripe Dashboard and on any invoices sent to the Account.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Information about the company, individual, and business represented by the Account.
	Identity *V2CoreAccountCreateIdentityParams `form:"identity" json:"identity,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of an Account.
type V2CoreAccountRetrieveParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticParams struct {
	// Can create commercial issuing charge cards with Celtic as BIN sponsor.
	ChargeCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing spend cards with Celtic as BIN sponsor.
	SpendCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams struct {
	// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
	ChargeCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
	// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Fifth Third as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialFifthThirdParams struct {
	// Can create commercial issuing charge cards with Fifth Third as BIN sponsor.
	ChargeCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Lead as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadParams struct {
	// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing charge cards with Stripe as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create commercial issuing prepaid cards with Stripe as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Stripe as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripeParams struct {
	// Can create commercial issuing charge cards with Stripe as BIN sponsor.
	ChargeCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing prepaid cards with Stripe as BIN sponsor.
	PrepaidCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Can create cards for commercial issuing use cases.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialParams struct {
	// Can create commercial issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Can create commercial issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Can create commercial issuing cards with Fifth Third as BIN sponsor.
	FifthThird *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialFifthThirdParams `form:"fifth_third" json:"fifth_third,omitempty"`
	// Can create commercial issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadParams `form:"lead" json:"lead,omitempty"`
	// Can create commercial issuing cards with Stripe as BIN sponsor.
	Stripe *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripeParams `form:"stripe" json:"stripe,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing revolving credit cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create consumer issuing cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCelticParams struct {
	// Can create consumer issuing revolving credit cards with Celtic as BIN sponsor.
	RevolvingCreditCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardParams `form:"revolving_credit_card" json:"revolving_credit_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create consumer issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankParams struct {
	// Can create consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can create consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create consumer issuing cards with Lead as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadParams struct {
	// Can create consumer issuing debit cards with Lead as BIN sponsor.
	DebitCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardParams `form:"debit_card" json:"debit_card,omitempty"`
	// Can create consumer issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Can create cards for consumer issuing use cases.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerParams struct {
	// Can create consumer issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Can create consumer issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Can create consumer issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerLeadParams `form:"lead" json:"lead,omitempty"`
}

// Capabilities to request on the CardCreator Configuration.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesParams struct {
	// Can create cards for commercial issuing use cases.
	Commercial *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialParams `form:"commercial" json:"commercial,omitempty"`
	// Can create cards for consumer issuing use cases.
	Consumer *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesConsumerParams `form:"consumer" json:"consumer,omitempty"`
}

// The CardCreator Configuration allows the Account to create and issue cards to users.
type V2CoreAccountUpdateConfigurationCardCreatorParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Capabilities to request on the CardCreator Configuration.
	Capabilities *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
}

// Automatic indirect tax settings to be used when automatic tax calculation is enabled on the customer's invoices, subscriptions, checkout sessions, or payment links. Surfaces if automatic tax calculation is possible given the current customer location information.
type V2CoreAccountUpdateConfigurationCustomerAutomaticIndirectTaxParams struct {
	// The customer account's tax exemption status: `none`, `exempt`, or `reverse`. When `reverse`, invoice and receipt PDFs include "Reverse charge".
	Exempt *string `form:"exempt" json:"exempt,omitempty"`
	// A recent IP address of the customer used for tax reporting and tax location inference.
	IPAddress *string `form:"ip_address" json:"ip_address,omitempty"`
	// Data source used to identify the customer account's tax location. Defaults to `identity_address`. Used for automatic indirect tax calculation.
	LocationSource *string `form:"location_source" json:"location_source,omitempty"`
	// A per-request flag that indicates when Stripe should validate the customer tax location - defaults to `auto`.
	ValidateLocation *string `form:"validate_location" json:"validate_location,omitempty"`
}

// The list of up to 4 default custom fields to be displayed on invoices for this customer.
type V2CoreAccountUpdateConfigurationCustomerBillingInvoiceCustomFieldParams struct {
	// The name of the custom field. This may be up to 40 characters.
	Name *string `form:"name" json:"name"`
	// The value of the custom field. This may be up to 140 characters. When updating, pass an empty string to remove previously-defined values.
	Value *string `form:"value" json:"value"`
}

// Default invoice PDF rendering options.
type V2CoreAccountUpdateConfigurationCustomerBillingInvoiceRenderingParams struct {
	// Indicates whether displayed line item prices and amounts on invoice PDFs include inclusive tax amounts. Must be either `include_inclusive_tax` or `exclude_tax`.
	AmountTaxDisplay *string `form:"amount_tax_display" json:"amount_tax_display,omitempty"`
	// ID of the invoice rendering template to use for future invoices.
	Template *string `form:"template" json:"template,omitempty"`
}

// Default invoice settings for the customer account.
type V2CoreAccountUpdateConfigurationCustomerBillingInvoiceParams struct {
	// The list of up to 4 default custom fields to be displayed on invoices for this customer.
	CustomFields []*V2CoreAccountUpdateConfigurationCustomerBillingInvoiceCustomFieldParams `form:"custom_fields" json:"custom_fields,omitempty"`
	// Default invoice footer.
	Footer *string `form:"footer" json:"footer,omitempty"`
	// Sequence number to use on the customer account's next invoice. Defaults to 1.
	NextSequence *int64 `form:"next_sequence" json:"next_sequence,omitempty"`
	// Prefix used to generate unique invoice numbers. Must be 3-12 uppercase letters or numbers.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
	// Default invoice PDF rendering options.
	Rendering *V2CoreAccountUpdateConfigurationCustomerBillingInvoiceRenderingParams `form:"rendering" json:"rendering,omitempty"`
}

// Billing settings - default settings used for this customer in Billing flows such as Invoices and Subscriptions.
type V2CoreAccountUpdateConfigurationCustomerBillingParams struct {
	// The ID of a `PaymentMethod` attached to this Account's `customer` configuration, used as the default payment method for invoices and subscriptions.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// Default invoice settings for the customer account.
	Invoice *V2CoreAccountUpdateConfigurationCustomerBillingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Generates requirements for enabling automatic indirect tax calculation on this customer's invoices or subscriptions. Recommended to request this capability if planning to enable automatic tax calculation on this customer's invoices or subscriptions.
type V2CoreAccountUpdateConfigurationCustomerCapabilitiesAutomaticIndirectTaxParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationCustomerCapabilitiesAutomaticIndirectTaxProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities that have been requested on the Customer Configuration.
type V2CoreAccountUpdateConfigurationCustomerCapabilitiesParams struct {
	// Generates requirements for enabling automatic indirect tax calculation on this customer's invoices or subscriptions. Recommended to request this capability if planning to enable automatic tax calculation on this customer's invoices or subscriptions.
	AutomaticIndirectTax *V2CoreAccountUpdateConfigurationCustomerCapabilitiesAutomaticIndirectTaxParams `form:"automatic_indirect_tax" json:"automatic_indirect_tax,omitempty"`
}

// The customer's shipping information. Appears on invoices emailed to this customer.
type V2CoreAccountUpdateConfigurationCustomerShippingParams struct {
	// Customer shipping address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Customer name.
	Name *string `form:"name" json:"name,omitempty"`
	// Customer phone (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// The Customer Configuration allows the Account to be charged.
type V2CoreAccountUpdateConfigurationCustomerParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Automatic indirect tax settings to be used when automatic tax calculation is enabled on the customer's invoices, subscriptions, checkout sessions, or payment links. Surfaces if automatic tax calculation is possible given the current customer location information.
	AutomaticIndirectTax *V2CoreAccountUpdateConfigurationCustomerAutomaticIndirectTaxParams `form:"automatic_indirect_tax" json:"automatic_indirect_tax,omitempty"`
	// Billing settings - default settings used for this customer in Billing flows such as Invoices and Subscriptions.
	Billing *V2CoreAccountUpdateConfigurationCustomerBillingParams `form:"billing" json:"billing,omitempty"`
	// Capabilities that have been requested on the Customer Configuration.
	Capabilities *V2CoreAccountUpdateConfigurationCustomerCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// The customer's shipping information. Appears on invoices emailed to this customer.
	Shipping *V2CoreAccountUpdateConfigurationCustomerShippingParams `form:"shipping" json:"shipping,omitempty"`
	// ID of the test clock to attach to the customer. Can only be set on testmode Accounts, and when the Customer Configuration is first set on an Account.
	TestClock *string `form:"test_clock" json:"test_clock,omitempty"`
}

// Settings for Bacs Direct Debit payments.
type V2CoreAccountUpdateConfigurationMerchantBACSDebitPaymentsParams struct {
	// Display name for Bacs Direct Debit payments.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
}

// Settings used to apply the merchant's branding to email receipts, invoices, Checkout, and other products.
type V2CoreAccountUpdateConfigurationMerchantBrandingParams struct {
	// ID of a [file upload](https://docs.stripe.com/api/persons/update#create_file): An icon for the merchant. Must be square and at least 128px x 128px.
	Icon *string `form:"icon" json:"icon,omitempty"`
	// ID of a [file upload](https://docs.stripe.com/api/persons/update#create_file): A logo for the merchant that will be used in Checkout instead of the icon and without the merchant's name next to it if provided. Must be at least 128px x 128px.
	Logo *string `form:"logo" json:"logo,omitempty"`
	// A CSS hex color value representing the primary branding color for the merchant.
	PrimaryColor *string `form:"primary_color" json:"primary_color,omitempty"`
	// A CSS hex color value representing the secondary branding color for the merchant.
	SecondaryColor *string `form:"secondary_color" json:"secondary_color,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process ACH debit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesACHDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesACHDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process ACSS debit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesACSSDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesACSSDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Affirm payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAffirmPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAffirmPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Afterpay/Clearpay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Alma payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAlmaPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAlmaPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Amazon Pay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAmazonPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAmazonPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Australian BECS Direct Debit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process BACS Direct Debit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBACSDebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBACSDebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Bancontact payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBancontactPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBancontactPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process BLIK payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBLIKPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBLIKPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Boleto payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBoletoPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBoletoPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCardPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCardPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesCardPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to collect card payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCardPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesCardPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Cartes Bancaires payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCartesBancairesPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesCartesBancairesPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Cash App payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCashAppPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesCashAppPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesEPSPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesEPSPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesEPSPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process EPS payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesEPSPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesEPSPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesFPXPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesFPXPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesFPXPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process FPX payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesFPXPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesFPXPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process UK bank transfer payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesGBBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesGBBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process GrabPay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesGrabpayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesGrabpayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process iDEAL payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesIDEALPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesIDEALPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesJCBPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesJCBPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesJCBPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process JCB card payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesJCBPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesJCBPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Japanese bank transfer payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesJPBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesJPBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Kakao Pay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKakaoPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKakaoPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Klarna payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKlarnaPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKlarnaPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Konbini convenience store payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKonbiniPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKonbiniPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Korean card payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKrCardPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKrCardPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesLinkPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesLinkPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesLinkPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Link payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesLinkPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesLinkPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process MobilePay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMobilepayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesMobilepayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Multibanco payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMultibancoPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesMultibancoPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Mexican bank transfer payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMXBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesMXBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Naver Pay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesNaverPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesNaverPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process OXXO payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesOXXOPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesOXXOPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesP24PaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesP24PaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesP24PaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Przelewy24 (P24) payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesP24PaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesP24PaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Pay by Bank payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayByBankPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayByBankPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process PAYCO payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPaycoPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPaycoPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process PayNow payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayNowPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayNowPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process PromptPay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPromptPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPromptPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Revolut Pay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesRevolutPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesRevolutPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Samsung Pay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSamsungPayPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSamsungPayPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process SEPA bank transfer payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process SEPA Direct Debit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPADebitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPADebitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Sunbit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSunbitPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSunbitPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSwishPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSwishPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSwishPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Swish payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSwishPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSwishPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process TWINT payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesTWINTPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesTWINTPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process US bank transfer payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesUSBankTransferPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesUSBankTransferPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesZipPaymentsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesZipPaymentsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMerchantCapabilitiesZipPaymentsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Allow the merchant to process Zip payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesZipPaymentsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMerchantCapabilitiesZipPaymentsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities to request on the Merchant Configuration.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesParams struct {
	// Allow the merchant to process ACH debit payments.
	ACHDebitPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesACHDebitPaymentsParams `form:"ach_debit_payments" json:"ach_debit_payments,omitempty"`
	// Allow the merchant to process ACSS debit payments.
	ACSSDebitPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesACSSDebitPaymentsParams `form:"acss_debit_payments" json:"acss_debit_payments,omitempty"`
	// Allow the merchant to process Affirm payments.
	AffirmPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAffirmPaymentsParams `form:"affirm_payments" json:"affirm_payments,omitempty"`
	// Allow the merchant to process Afterpay/Clearpay payments.
	AfterpayClearpayPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsParams `form:"afterpay_clearpay_payments" json:"afterpay_clearpay_payments,omitempty"`
	// Allow the merchant to process Alma payments.
	AlmaPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAlmaPaymentsParams `form:"alma_payments" json:"alma_payments,omitempty"`
	// Allow the merchant to process Amazon Pay payments.
	AmazonPayPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAmazonPayPaymentsParams `form:"amazon_pay_payments" json:"amazon_pay_payments,omitempty"`
	// Allow the merchant to process Australian BECS Direct Debit payments.
	AUBECSDebitPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsParams `form:"au_becs_debit_payments" json:"au_becs_debit_payments,omitempty"`
	// Allow the merchant to process BACS Direct Debit payments.
	BACSDebitPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBACSDebitPaymentsParams `form:"bacs_debit_payments" json:"bacs_debit_payments,omitempty"`
	// Allow the merchant to process Bancontact payments.
	BancontactPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBancontactPaymentsParams `form:"bancontact_payments" json:"bancontact_payments,omitempty"`
	// Allow the merchant to process BLIK payments.
	BLIKPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBLIKPaymentsParams `form:"blik_payments" json:"blik_payments,omitempty"`
	// Allow the merchant to process Boleto payments.
	BoletoPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesBoletoPaymentsParams `form:"boleto_payments" json:"boleto_payments,omitempty"`
	// Allow the merchant to collect card payments.
	CardPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesCardPaymentsParams `form:"card_payments" json:"card_payments,omitempty"`
	// Allow the merchant to process Cartes Bancaires payments.
	CartesBancairesPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesCartesBancairesPaymentsParams `form:"cartes_bancaires_payments" json:"cartes_bancaires_payments,omitempty"`
	// Allow the merchant to process Cash App payments.
	CashAppPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesCashAppPaymentsParams `form:"cashapp_payments" json:"cashapp_payments,omitempty"`
	// Allow the merchant to process EPS payments.
	EPSPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesEPSPaymentsParams `form:"eps_payments" json:"eps_payments,omitempty"`
	// Allow the merchant to process FPX payments.
	FPXPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesFPXPaymentsParams `form:"fpx_payments" json:"fpx_payments,omitempty"`
	// Allow the merchant to process UK bank transfer payments.
	GBBankTransferPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesGBBankTransferPaymentsParams `form:"gb_bank_transfer_payments" json:"gb_bank_transfer_payments,omitempty"`
	// Allow the merchant to process GrabPay payments.
	GrabpayPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesGrabpayPaymentsParams `form:"grabpay_payments" json:"grabpay_payments,omitempty"`
	// Allow the merchant to process iDEAL payments.
	IDEALPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesIDEALPaymentsParams `form:"ideal_payments" json:"ideal_payments,omitempty"`
	// Allow the merchant to process JCB card payments.
	JCBPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesJCBPaymentsParams `form:"jcb_payments" json:"jcb_payments,omitempty"`
	// Allow the merchant to process Japanese bank transfer payments.
	JPBankTransferPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesJPBankTransferPaymentsParams `form:"jp_bank_transfer_payments" json:"jp_bank_transfer_payments,omitempty"`
	// Allow the merchant to process Kakao Pay payments.
	KakaoPayPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKakaoPayPaymentsParams `form:"kakao_pay_payments" json:"kakao_pay_payments,omitempty"`
	// Allow the merchant to process Klarna payments.
	KlarnaPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKlarnaPaymentsParams `form:"klarna_payments" json:"klarna_payments,omitempty"`
	// Allow the merchant to process Konbini convenience store payments.
	KonbiniPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKonbiniPaymentsParams `form:"konbini_payments" json:"konbini_payments,omitempty"`
	// Allow the merchant to process Korean card payments.
	KrCardPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesKrCardPaymentsParams `form:"kr_card_payments" json:"kr_card_payments,omitempty"`
	// Allow the merchant to process Link payments.
	LinkPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesLinkPaymentsParams `form:"link_payments" json:"link_payments,omitempty"`
	// Allow the merchant to process MobilePay payments.
	MobilepayPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesMobilepayPaymentsParams `form:"mobilepay_payments" json:"mobilepay_payments,omitempty"`
	// Allow the merchant to process Multibanco payments.
	MultibancoPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesMultibancoPaymentsParams `form:"multibanco_payments" json:"multibanco_payments,omitempty"`
	// Allow the merchant to process Mexican bank transfer payments.
	MXBankTransferPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesMXBankTransferPaymentsParams `form:"mx_bank_transfer_payments" json:"mx_bank_transfer_payments,omitempty"`
	// Allow the merchant to process Naver Pay payments.
	NaverPayPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesNaverPayPaymentsParams `form:"naver_pay_payments" json:"naver_pay_payments,omitempty"`
	// Allow the merchant to process OXXO payments.
	OXXOPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesOXXOPaymentsParams `form:"oxxo_payments" json:"oxxo_payments,omitempty"`
	// Allow the merchant to process Przelewy24 (P24) payments.
	P24Payments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesP24PaymentsParams `form:"p24_payments" json:"p24_payments,omitempty"`
	// Allow the merchant to process Pay by Bank payments.
	PayByBankPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayByBankPaymentsParams `form:"pay_by_bank_payments" json:"pay_by_bank_payments,omitempty"`
	// Allow the merchant to process PAYCO payments.
	PaycoPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPaycoPaymentsParams `form:"payco_payments" json:"payco_payments,omitempty"`
	// Allow the merchant to process PayNow payments.
	PayNowPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayNowPaymentsParams `form:"paynow_payments" json:"paynow_payments,omitempty"`
	// Allow the merchant to process PromptPay payments.
	PromptPayPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesPromptPayPaymentsParams `form:"promptpay_payments" json:"promptpay_payments,omitempty"`
	// Allow the merchant to process Revolut Pay payments.
	RevolutPayPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesRevolutPayPaymentsParams `form:"revolut_pay_payments" json:"revolut_pay_payments,omitempty"`
	// Allow the merchant to process Samsung Pay payments.
	SamsungPayPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSamsungPayPaymentsParams `form:"samsung_pay_payments" json:"samsung_pay_payments,omitempty"`
	// Allow the merchant to process SEPA bank transfer payments.
	SEPABankTransferPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsParams `form:"sepa_bank_transfer_payments" json:"sepa_bank_transfer_payments,omitempty"`
	// Allow the merchant to process SEPA Direct Debit payments.
	SEPADebitPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPADebitPaymentsParams `form:"sepa_debit_payments" json:"sepa_debit_payments,omitempty"`
	// Allow the merchant to process Sunbit payments.
	SunbitPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSunbitPaymentsParams `form:"sunbit_payments" json:"sunbit_payments,omitempty"`
	// Allow the merchant to process Swish payments.
	SwishPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesSwishPaymentsParams `form:"swish_payments" json:"swish_payments,omitempty"`
	// Allow the merchant to process TWINT payments.
	TWINTPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesTWINTPaymentsParams `form:"twint_payments" json:"twint_payments,omitempty"`
	// Allow the merchant to process US bank transfer payments.
	USBankTransferPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesUSBankTransferPaymentsParams `form:"us_bank_transfer_payments" json:"us_bank_transfer_payments,omitempty"`
	// Allow the merchant to process Zip payments.
	ZipPayments *V2CoreAccountUpdateConfigurationMerchantCapabilitiesZipPaymentsParams `form:"zip_payments" json:"zip_payments,omitempty"`
}

// Automatically declines certain charge types regardless of whether the card issuer accepted or declined the charge.
type V2CoreAccountUpdateConfigurationMerchantCardPaymentsDeclineOnParams struct {
	// Whether Stripe automatically declines charges with an incorrect ZIP or postal code. This setting only applies when a ZIP or postal code is provided and they fail bank verification.
	AVSFailure *bool `form:"avs_failure" json:"avs_failure,omitempty"`
	// Whether Stripe automatically declines charges with an incorrect CVC. This setting only applies when a CVC is provided and it fails bank verification.
	CVCFailure *bool `form:"cvc_failure" json:"cvc_failure,omitempty"`
}

// Card payments settings.
type V2CoreAccountUpdateConfigurationMerchantCardPaymentsParams struct {
	// Automatically declines certain charge types regardless of whether the card issuer accepted or declined the charge.
	DeclineOn *V2CoreAccountUpdateConfigurationMerchantCardPaymentsDeclineOnParams `form:"decline_on" json:"decline_on,omitempty"`
}

// Support hours for Konbini payments.
type V2CoreAccountUpdateConfigurationMerchantKonbiniPaymentsSupportHoursParams struct {
	// Support hours end time (JST time of day) for in `HH:MM` format.
	EndTime *string `form:"end_time" json:"end_time,omitempty"`
	// Support hours start time (JST time of day) for in `HH:MM` format.
	StartTime *string `form:"start_time" json:"start_time,omitempty"`
}

// Support for Konbini payments.
type V2CoreAccountUpdateConfigurationMerchantKonbiniPaymentsSupportParams struct {
	// Support email address for Konbini payments.
	Email *string `form:"email" json:"email,omitempty"`
	// Support hours for Konbini payments.
	Hours *V2CoreAccountUpdateConfigurationMerchantKonbiniPaymentsSupportHoursParams `form:"hours" json:"hours,omitempty"`
	// Support phone number for Konbini payments.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Settings specific to Konbini payments on the account.
type V2CoreAccountUpdateConfigurationMerchantKonbiniPaymentsParams struct {
	// Support for Konbini payments.
	Support *V2CoreAccountUpdateConfigurationMerchantKonbiniPaymentsSupportParams `form:"support" json:"support,omitempty"`
}

// The Kana variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
type V2CoreAccountUpdateConfigurationMerchantScriptStatementDescriptorKanaParams struct {
	// The default text that appears on statements for non-card charges outside of Japan. For card charges, if you don't set a statement_descriptor_prefix, this text is also used as the statement descriptor prefix. In that case, if concatenating the statement descriptor suffix causes the combined statement descriptor to exceed 22 characters, we truncate the statement_descriptor text to limit the full descriptor to 22 characters. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Descriptor *string `form:"descriptor" json:"descriptor,omitempty"`
	// Default text that appears on statements for card charges outside of Japan, prefixing any dynamic statement_descriptor_suffix specified on the charge. To maximize space for the dynamic part of the descriptor, keep this text short. If you don't specify this value, statement_descriptor is used as the prefix. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
}

// The Kanji variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
type V2CoreAccountUpdateConfigurationMerchantScriptStatementDescriptorKanjiParams struct {
	// The default text that appears on statements for non-card charges outside of Japan. For card charges, if you don't set a statement_descriptor_prefix, this text is also used as the statement descriptor prefix. In that case, if concatenating the statement descriptor suffix causes the combined statement descriptor to exceed 22 characters, we truncate the statement_descriptor text to limit the full descriptor to 22 characters. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Descriptor *string `form:"descriptor" json:"descriptor,omitempty"`
	// Default text that appears on statements for card charges outside of Japan, prefixing any dynamic statement_descriptor_suffix specified on the charge. To maximize space for the dynamic part of the descriptor, keep this text short. If you don't specify this value, statement_descriptor is used as the prefix. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
}

// Settings for the default text that appears on statements for language variations.
type V2CoreAccountUpdateConfigurationMerchantScriptStatementDescriptorParams struct {
	// The Kana variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
	Kana *V2CoreAccountUpdateConfigurationMerchantScriptStatementDescriptorKanaParams `form:"kana" json:"kana,omitempty"`
	// The Kanji variation of statement_descriptor used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).
	Kanji *V2CoreAccountUpdateConfigurationMerchantScriptStatementDescriptorKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Settings for Smart Disputes auto_respond.
type V2CoreAccountUpdateConfigurationMerchantSmartDisputesAutoRespondParams struct {
	// The preference for automatic dispute responses.
	Preference *string `form:"preference" json:"preference,omitempty"`
}

// Settings for Smart Disputes automatic response feature.
type V2CoreAccountUpdateConfigurationMerchantSmartDisputesParams struct {
	// Settings for Smart Disputes auto_respond.
	AutoRespond *V2CoreAccountUpdateConfigurationMerchantSmartDisputesAutoRespondParams `form:"auto_respond" json:"auto_respond,omitempty"`
}

// Settings for the default [statement descriptor](https://docs.stripe.com/connect/statement-descriptors) text.
type V2CoreAccountUpdateConfigurationMerchantStatementDescriptorParams struct {
	// The default text that appears on statements for non-card charges outside of Japan. For card charges, if you don't set a statement_descriptor_prefix, this text is also used as the statement descriptor prefix. In that case, if concatenating the statement descriptor suffix causes the combined statement descriptor to exceed 22 characters, we truncate the statement_descriptor text to limit the full descriptor to 22 characters. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Descriptor *string `form:"descriptor" json:"descriptor,omitempty"`
	// Default text that appears on statements for card charges outside of Japan, prefixing any dynamic statement_descriptor_suffix specified on the charge. To maximize space for the dynamic part of the descriptor, keep this text short. If you don't specify this value, statement_descriptor is used as the prefix. For more information about statement descriptors and their requirements, see the Merchant Configuration settings documentation.
	Prefix *string `form:"prefix" json:"prefix,omitempty"`
}

// A publicly available mailing address for sending support issues to.
type V2CoreAccountUpdateConfigurationMerchantSupportAddressParams struct {
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

// Publicly available contact information for sending support issues to.
type V2CoreAccountUpdateConfigurationMerchantSupportParams struct {
	// A publicly available mailing address for sending support issues to.
	Address *V2CoreAccountUpdateConfigurationMerchantSupportAddressParams `form:"address" json:"address,omitempty"`
	// A publicly available email address for sending support issues to.
	Email *string `form:"email" json:"email,omitempty"`
	// A publicly available phone number to call with support issues.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// A publicly available website for handling support issues.
	URL *string `form:"url" json:"url,omitempty"`
}

// Enables the Account to act as a connected account and collect payments facilitated by a Connect platform. You must onboard your platform to Connect before you can add this configuration to your connected accounts. Utilize this configuration when the Account will be the Merchant of Record, like with Direct charges or Destination Charges with on_behalf_of set.
type V2CoreAccountUpdateConfigurationMerchantParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Settings for Bacs Direct Debit payments.
	BACSDebitPayments *V2CoreAccountUpdateConfigurationMerchantBACSDebitPaymentsParams `form:"bacs_debit_payments" json:"bacs_debit_payments,omitempty"`
	// Settings used to apply the merchant's branding to email receipts, invoices, Checkout, and other products.
	Branding *V2CoreAccountUpdateConfigurationMerchantBrandingParams `form:"branding" json:"branding,omitempty"`
	// Capabilities to request on the Merchant Configuration.
	Capabilities *V2CoreAccountUpdateConfigurationMerchantCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// Card payments settings.
	CardPayments *V2CoreAccountUpdateConfigurationMerchantCardPaymentsParams `form:"card_payments" json:"card_payments,omitempty"`
	// Settings specific to Konbini payments on the account.
	KonbiniPayments *V2CoreAccountUpdateConfigurationMerchantKonbiniPaymentsParams `form:"konbini_payments" json:"konbini_payments,omitempty"`
	// The Merchant Category Code (MCC) for the merchant. MCCs classify businesses based on the goods or services they provide.
	MCC *string `form:"mcc" json:"mcc,omitempty"`
	// Settings for the default text that appears on statements for language variations.
	ScriptStatementDescriptor *V2CoreAccountUpdateConfigurationMerchantScriptStatementDescriptorParams `form:"script_statement_descriptor" json:"script_statement_descriptor,omitempty"`
	// Settings for Smart Disputes automatic response feature.
	SmartDisputes *V2CoreAccountUpdateConfigurationMerchantSmartDisputesParams `form:"smart_disputes" json:"smart_disputes,omitempty"`
	// Settings for the default [statement descriptor](https://docs.stripe.com/connect/statement-descriptors) text.
	StatementDescriptor *V2CoreAccountUpdateConfigurationMerchantStatementDescriptorParams `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Publicly available contact information for sending support issues to.
	Support *V2CoreAccountUpdateConfigurationMerchantSupportParams `form:"support" json:"support,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in AUD.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in CAD.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in EUR.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in GBP.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in USD.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive business storage-type funds on Stripe in USDC.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can receive business storage-type funds on Stripe.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundParams struct {
	// Can receive business storage-type funds on Stripe in AUD.
	AUD *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundAUDParams `form:"aud" json:"aud,omitempty"`
	// Can receive business storage-type funds on Stripe in CAD.
	CAD *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundCADParams `form:"cad" json:"cad,omitempty"`
	// Can receive business storage-type funds on Stripe in EUR.
	EUR *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundEURParams `form:"eur" json:"eur,omitempty"`
	// Can receive business storage-type funds on Stripe in GBP.
	GBP *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundGBPParams `form:"gbp" json:"gbp,omitempty"`
	// Can receive business storage-type funds on Stripe in USD.
	USD *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can receive business storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in AUD.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in CAD.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in EUR.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in GBP.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in USD.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send business storage-type funds on Stripe in USDC.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send business storage-type funds on Stripe.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundParams struct {
	// Can send business storage-type funds on Stripe in AUD.
	AUD *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundAUDParams `form:"aud" json:"aud,omitempty"`
	// Can send business storage-type funds on Stripe in CAD.
	CAD *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundCADParams `form:"cad" json:"cad,omitempty"`
	// Can send business storage-type funds on Stripe in EUR.
	EUR *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundEURParams `form:"eur" json:"eur,omitempty"`
	// Can send business storage-type funds on Stripe in GBP.
	GBP *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundGBPParams `form:"gbp" json:"gbp,omitempty"`
	// Can send business storage-type funds on Stripe in USD.
	USD *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can send business storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Can send or receive business storage-type funds on Stripe.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageParams struct {
	// Can receive business storage-type funds on Stripe.
	Inbound *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageInboundParams `form:"inbound" json:"inbound,omitempty"`
	// Can send business storage-type funds on Stripe.
	Outbound *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundParams `form:"outbound" json:"outbound,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive consumer storage-type funds on Stripe in USD.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive consumer storage-type funds on Stripe in USDC.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can receive consumer storage-type funds on Stripe.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundParams struct {
	// Can receive consumer storage-type funds on Stripe in USD.
	USD *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can receive consumer storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send consumer storage-type funds on Stripe in USD.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send consumer storage-type funds on Stripe in USDC.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send consumer storage-type funds on Stripe.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundParams struct {
	// Can send consumer storage-type funds on Stripe in USD.
	USD *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUSDParams `form:"usd" json:"usd,omitempty"`
	// Can send consumer storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Can send or receive consumer storage-type funds on Stripe.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageParams struct {
	// Can receive consumer storage-type funds on Stripe.
	Inbound *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageInboundParams `form:"inbound" json:"inbound,omitempty"`
	// Can send consumer storage-type funds on Stripe.
	Outbound *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageOutboundParams `form:"outbound" json:"outbound,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesInboundTransfersParams struct {
	// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
	BankAccounts *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesInboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a bank account owned by someone else.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a debit card owned by someone else.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to someone else via paper check.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a destination owned by someone else.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by someone else.
	BankAccounts *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a debit card owned by someone else.
	Cards *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCardsParams `form:"cards" json:"cards,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
	CryptoWallets *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
	FinancialAccounts *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
	// Can send funds from a FinancialAccount to someone else via paper check.
	PaperChecks *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsPaperChecksParams `form:"paper_checks" json:"paper_checks,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a bank account owned by yourself.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a destination owned by yourself.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by yourself.
	BankAccounts *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
	CryptoWallets *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
	FinancialAccounts *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive funds on a bank-account-like financial address (VBAN) to credit a FinancialAccount.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive funds on a crypto wallet like financial address to credit a FinancialAccount.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can receive funds on a financial address to credit a FinancialAccount.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsParams struct {
	// Can receive funds on a bank-account-like financial address (VBAN) to credit a FinancialAccount.
	BankAccounts *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can receive funds on a crypto wallet like financial address to credit a FinancialAccount.
	CryptoWallets *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Can receive debits to a FinancialAccount from a bank account.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can receive debits to a FinancialAccount.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedDebitsParams struct {
	// Can receive debits to a FinancialAccount from a bank account.
	BankAccounts *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedDebitsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
}

// Capabilities to request on the Money Manager Configuration.
type V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesParams struct {
	// Can send or receive business storage-type funds on Stripe.
	BusinessStorage *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesBusinessStorageParams `form:"business_storage" json:"business_storage,omitempty"`
	// Can send or receive consumer storage-type funds on Stripe.
	ConsumerStorage *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesConsumerStorageParams `form:"consumer_storage" json:"consumer_storage,omitempty"`
	// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
	InboundTransfers *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesInboundTransfersParams `form:"inbound_transfers" json:"inbound_transfers,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by someone else.
	OutboundPayments *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundPaymentsParams `form:"outbound_payments" json:"outbound_payments,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by yourself.
	OutboundTransfers *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesOutboundTransfersParams `form:"outbound_transfers" json:"outbound_transfers,omitempty"`
	// Can receive funds on a financial address to credit a FinancialAccount.
	ReceivedCredits *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedCreditsParams `form:"received_credits" json:"received_credits,omitempty"`
	// Can receive debits to a FinancialAccount.
	ReceivedDebits *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesReceivedDebitsParams `form:"received_debits" json:"received_debits,omitempty"`
}

// Details of the regulated activity if the business participates in one.
type V2CoreAccountUpdateConfigurationMoneyManagerRegulatedActivityParams struct {
	// A detailed description of the regulated activities the business is licensed to conduct.
	Description *string `form:"description" json:"description,omitempty"`
	// The license number or registration number assigned by the business's primary regulator.
	LicenseNumber *string `form:"license_number" json:"license_number,omitempty"`
	// The country of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityCountry *string `form:"primary_regulatory_authority_country" json:"primary_regulatory_authority_country,omitempty"`
	// The name of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityName *string `form:"primary_regulatory_authority_name" json:"primary_regulatory_authority_name,omitempty"`
}

// The Money Manager Configuration allows the Account to store and move funds using FinancialAccounts.
type V2CoreAccountUpdateConfigurationMoneyManagerParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Capabilities to request on the Money Manager Configuration.
	Capabilities *V2CoreAccountUpdateConfigurationMoneyManagerCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// List of high-risk activities the business is involved in.
	HighRiskActivities []*string `form:"high_risk_activities" json:"high_risk_activities,omitempty"`
	// Description of the high-risk activities the business offers.
	HighRiskActivitiesDescription *string `form:"high_risk_activities_description" json:"high_risk_activities_description,omitempty"`
	// Description of the money services offered by the business.
	MoneyServicesDescription *string `form:"money_services_description" json:"money_services_description,omitempty"`
	// Does the business operate in any prohibited countries.
	OperatesInProhibitedCountries *bool `form:"operates_in_prohibited_countries" json:"operates_in_prohibited_countries,omitempty"`
	// Indicates whether the business participates in any regulated activity.
	ParticipatesInRegulatedActivity *bool `form:"participates_in_regulated_activity" json:"participates_in_regulated_activity,omitempty"`
	// Primary purpose of the stored funds.
	PurposeOfFunds *string `form:"purpose_of_funds" json:"purpose_of_funds,omitempty"`
	// Description of the purpose of the stored funds.
	PurposeOfFundsDescription *string `form:"purpose_of_funds_description" json:"purpose_of_funds_description,omitempty"`
	// Details of the regulated activity if the business participates in one.
	RegulatedActivity *V2CoreAccountUpdateConfigurationMoneyManagerRegulatedActivityParams `form:"regulated_activity" json:"regulated_activity,omitempty"`
	// The source of funds for the business, e.g. profits, income, venture capital, etc.
	SourceOfFunds *string `form:"source_of_funds" json:"source_of_funds,omitempty"`
	// Description of the source of funds for the business' account.
	SourceOfFundsDescription *string `form:"source_of_funds_description" json:"source_of_funds_description,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsACHProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsACHProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsACHProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over ACH rails.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsACHParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsACHProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over BECS rails.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsBECSParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsBECSProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsEftProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsEftProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsEftProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over EFT rails.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsEftParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsEftProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over Fedwire or CHIPS.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFedwireParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFedwireProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over FPS rails.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFPSParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFPSProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsInstantParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsInstantProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsLocalParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsLocalProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsNppProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsNppProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsNppProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over NPP (real time) rails.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsNppParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsNppProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over RTP rails.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsRTPParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsRTPProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA credit rails.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPACreditParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPACreditProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA instant (real time) rails.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over SWIFT.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSwiftParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSwiftProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsWireProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsWireProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsWireProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsWireParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsWireProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities that enable OutboundPayments to a bank account linked to this Account.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsParams struct {
	// Enables this Account to receive OutboundPayments to linked bank accounts over ACH rails.
	ACH *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsACHParams `form:"ach" json:"ach,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over BECS rails.
	BECS *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsBECSParams `form:"becs" json:"becs,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over EFT rails.
	Eft *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsEftParams `form:"eft" json:"eft,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over Fedwire or CHIPS.
	Fedwire *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFedwireParams `form:"fedwire" json:"fedwire,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over FPS rails.
	FPS *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsFPSParams `form:"fps" json:"fps,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
	Instant *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsInstantParams `form:"instant" json:"instant,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
	Local *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsLocalParams `form:"local" json:"local,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over NPP (real time) rails.
	Npp *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsNppParams `form:"npp" json:"npp,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over RTP rails.
	RTP *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsRTPParams `form:"rtp" json:"rtp,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA credit rails.
	SEPACredit *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPACreditParams `form:"sepa_credit" json:"sepa_credit,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over SEPA instant (real time) rails.
	SEPAInstant *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSEPAInstantParams `form:"sepa_instant" json:"sepa_instant,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over SWIFT.
	Swift *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsSwiftParams `form:"swift" json:"swift,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
	Wire *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsWireParams `form:"wire" json:"wire,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesCardsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesCardsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesCardsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Capability that enable OutboundPayments to a debit card linked to this Account.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesCardsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesCardsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesCryptoWalletsProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesCryptoWalletsProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesCryptoWalletsProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Capabilities that enable OutboundPayments to a crypto wallet linked to this Account.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesCryptoWalletsParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesCryptoWalletsProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesPaperChecksProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesPaperChecksProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesPaperChecksProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Capabilities that enable OutboundPayments via paper check.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesPaperChecksParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesPaperChecksProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Parameter to request psp_migration protection.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsPspMigrationParams struct {
	// To request a protection, pass true.
	Requested *bool `form:"requested" json:"requested"`
}

// Protection types to request for this capability (e.g. "psp_migration").
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsParams struct {
	// Parameter to request psp_migration protection.
	PspMigration *V2CoreAccountUpdateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsPspMigrationParams `form:"psp_migration" json:"psp_migration"`
}

// Enables this Account to receive /v1/transfers into their Stripe Balance (/v1/balance).
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams struct {
	// Protection types to request for this capability (e.g. "psp_migration").
	Protections *V2CoreAccountUpdateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersProtectionsParams `form:"protections" json:"protections,omitempty"`
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities that enable the recipient to manage their Stripe Balance (/v1/balance).
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesStripeBalanceParams struct {
	// Enables this Account to receive /v1/transfers into their Stripe Balance (/v1/balance).
	StripeTransfers *V2CoreAccountUpdateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams `form:"stripe_transfers" json:"stripe_transfers,omitempty"`
}

// Capabilities to request on the Recipient Configuration.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesParams struct {
	// Capabilities that enable OutboundPayments to a bank account linked to this Account.
	BankAccounts *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Capability that enable OutboundPayments to a debit card linked to this Account.
	Cards *V2CoreAccountUpdateConfigurationRecipientCapabilitiesCardsParams `form:"cards" json:"cards,omitempty"`
	// Capabilities that enable OutboundPayments to a crypto wallet linked to this Account.
	CryptoWallets *V2CoreAccountUpdateConfigurationRecipientCapabilitiesCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Capabilities that enable OutboundPayments via paper check.
	PaperChecks *V2CoreAccountUpdateConfigurationRecipientCapabilitiesPaperChecksParams `form:"paper_checks" json:"paper_checks,omitempty"`
	// Capabilities that enable the recipient to manage their Stripe Balance (/v1/balance).
	StripeBalance *V2CoreAccountUpdateConfigurationRecipientCapabilitiesStripeBalanceParams `form:"stripe_balance" json:"stripe_balance,omitempty"`
}

// The Recipient Configuration allows the Account to receive funds. Utilize this configuration if the Account will not be the Merchant of Record, like with Separate Charges & Transfers, or Destination Charges without on_behalf_of set.
type V2CoreAccountUpdateConfigurationRecipientParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Capabilities to request on the Recipient Configuration.
	Capabilities *V2CoreAccountUpdateConfigurationRecipientCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
	// The payout method id to be used as a default outbound destination. This will allow the PayoutMethod to be omitted on OutboundPayments made through API or sending payouts via dashboard. Can also be explicitly set to `null` to clear the existing default outbound destination. For further details about creating an Outbound Destination, see [Collect recipient's payment details](https://docs.stripe.com/global-payouts-private-preview/quickstart?dashboard-or-api=api#collect-bank-account-details).
	DefaultOutboundDestination *string `form:"default_outbound_destination" json:"default_outbound_destination,omitempty"`
}

// An Account Configuration which allows the Account to take on a key persona across Stripe products.
type V2CoreAccountUpdateConfigurationParams struct {
	// The CardCreator Configuration allows the Account to create and issue cards to users.
	CardCreator *V2CoreAccountUpdateConfigurationCardCreatorParams `form:"card_creator" json:"card_creator,omitempty"`
	// The Customer Configuration allows the Account to be charged.
	Customer *V2CoreAccountUpdateConfigurationCustomerParams `form:"customer" json:"customer,omitempty"`
	// Enables the Account to act as a connected account and collect payments facilitated by a Connect platform. You must onboard your platform to Connect before you can add this configuration to your connected accounts. Utilize this configuration when the Account will be the Merchant of Record, like with Direct charges or Destination Charges with on_behalf_of set.
	Merchant *V2CoreAccountUpdateConfigurationMerchantParams `form:"merchant" json:"merchant,omitempty"`
	// The Money Manager Configuration allows the Account to store and move funds using FinancialAccounts.
	MoneyManager *V2CoreAccountUpdateConfigurationMoneyManagerParams `form:"money_manager" json:"money_manager,omitempty"`
	// The Recipient Configuration allows the Account to receive funds. Utilize this configuration if the Account will not be the Merchant of Record, like with Separate Charges & Transfers, or Destination Charges without on_behalf_of set.
	Recipient *V2CoreAccountUpdateConfigurationRecipientParams `form:"recipient" json:"recipient,omitempty"`
}

// Account profile information.
type V2CoreAccountUpdateDefaultsProfileParams struct {
	// The business's publicly-available website.
	BusinessURL *string `form:"business_url" json:"business_url,omitempty"`
	// The name which is used by the business.
	DoingBusinessAs *string `form:"doing_business_as" json:"doing_business_as,omitempty"`
	// Internal-only description of the product sold or service provided by the business. It's used by Stripe for risk and underwriting purposes.
	ProductDescription *string `form:"product_description" json:"product_description,omitempty"`
}

// Default responsibilities held by either Stripe or the platform.
type V2CoreAccountUpdateDefaultsResponsibilitiesParams struct {
	// A value indicating the party responsible for collecting fees from this account.
	FeesCollector *string `form:"fees_collector" json:"fees_collector"`
	// A value indicating who is responsible for losses when this Account can't pay back negative balances from payments.
	LossesCollector *string `form:"losses_collector" json:"losses_collector"`
	// A value indicating the party responsible for collecting requirements on this account.
	RequirementsCollector *string `form:"requirements_collector" json:"requirements_collector,omitempty"`
}

// Default values to be used on Account Configurations.
type V2CoreAccountUpdateDefaultsParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The Account's preferred locales (languages), ordered by preference.
	Locales []*string `form:"locales" json:"locales,omitempty"`
	// Account profile information.
	Profile *V2CoreAccountUpdateDefaultsProfileParams `form:"profile" json:"profile,omitempty"`
	// Default responsibilities held by either Stripe or the platform.
	Responsibilities *V2CoreAccountUpdateDefaultsResponsibilitiesParams `form:"responsibilities" json:"responsibilities,omitempty"`
	// The Account's local timezone. A list of possible time zone values is maintained at the [IANA Time Zone Database](https://www.iana.org/time-zones).
	Timezone *string `form:"timezone" json:"timezone,omitempty"`
}

// This hash is used to attest that the directors information provided to Stripe is both current and correct.
type V2CoreAccountUpdateIdentityAttestationsDirectorshipDeclarationParams struct {
	// The time marking when the director attestation was made. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the director attestation was made.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the director attestation was made.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
type V2CoreAccountUpdateIdentityAttestationsOwnershipDeclarationParams struct {
	// The time marking when the beneficial owner attestation was made. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the beneficial owner attestation was made.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the beneficial owner attestation was made.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Attestation that all Persons with a specific Relationship value have been provided.
type V2CoreAccountUpdateIdentityAttestationsPersonsProvidedParams struct {
	// Whether the company's directors have been provided. Set this Boolean to true after creating all the company's directors with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Directors *bool `form:"directors" json:"directors,omitempty"`
	// Whether the company's executives have been provided. Set this Boolean to true after creating all the company's executives with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Executives *bool `form:"executives" json:"executives,omitempty"`
	// Whether the company's owners have been provided. Set this Boolean to true after creating all the company's owners with the [Persons API](https://docs.stripe.com/api/v2/core/accounts/createperson).
	Owners *bool `form:"owners" json:"owners,omitempty"`
	// Reason for why the company is exempt from providing ownership information.
	OwnershipExemptionReason *string `form:"ownership_exemption_reason" json:"ownership_exemption_reason,omitempty"`
}

// This hash is used to attest that the representative is authorized to act as the representative of their legal entity.
type V2CoreAccountUpdateIdentityAttestationsRepresentativeDeclarationParams struct {
	// The time marking when the representative attestation was made. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the representative attestation was made.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the representative attestation was made.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of the [Stripe Services Agreement](https://docs.stripe.com/connect/updating-accounts#tos-acceptance).
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceAccountParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for Stripe commercial card issuing.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardParams struct {
	// Bank terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Platform terms of service acceptance for commercial issuing charge cards with Celtic as BIN sponsor.
	Platform *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardParams struct {
	// Bank terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
	FinancingDisclosures *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for commercial issuing spend cards with Celtic as BIN sponsor.
	Platform *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Celtic as BIN sponsor.
	ApplePay *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing charge cards with Celtic as BIN sponsor.
	ChargeCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Terms of service acceptances for commercial issuing spend cards with Celtic as BIN sponsor.
	SpendCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardParams struct {
	// Bank terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	FinancingDisclosures *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	Platform *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardParams struct {
	// Bank terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Platform terms of service acceptance for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	Platform *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardParams struct {
	// Bank terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	FinancingDisclosures *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Cross River Bank as BIN sponsor.
	ApplePay *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing charge cards with Cross River Bank as BIN sponsor.
	ChargeCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Terms of service acceptances for commercial issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
	// Terms of service acceptances for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Fifth Third as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing charge cards with Fifth Third as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardParams struct {
	// Bank terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Platform terms of service acceptance for commercial issuing charge cards with Fifth Third as BIN sponsor.
	Platform *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Fifth Third as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Fifth Third as BIN sponsor.
	ApplePay *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing charge cards with Fifth Third as BIN sponsor.
	ChargeCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
}

// Terms of service acceptances for Stripe commercial card Global issuing.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialGlobalAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing Apple Pay cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardParams struct {
	// Bank terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Platform terms of service acceptance for commercial issuing prepaid cards with Lead as BIN sponsor.
	Platform *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for commercial issuing cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadParams struct {
	// Terms of service acceptances for commercial issuing Apple Pay cards with Lead as BIN sponsor.
	ApplePay *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for commercial issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Terms of service acceptances to create cards for commercial issuing use cases.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialParams struct {
	// Terms of service acceptances for Stripe commercial card issuing.
	AccountHolder *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialAccountHolderParams `form:"account_holder" json:"account_holder,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Fifth Third as BIN sponsor.
	FifthThird *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialFifthThirdParams `form:"fifth_third" json:"fifth_third,omitempty"`
	// Terms of service acceptances for Stripe commercial card Global issuing.
	GlobalAccountHolder *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialGlobalAccountHolderParams `form:"global_account_holder" json:"global_account_holder,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadParams `form:"lead" json:"lead,omitempty"`
}

// Terms of service acceptances for Stripe consumer card issuing.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing Apple Pay cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing revolving credit cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardParams struct {
	// Bank terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
	FinancingDisclosures *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing spend cards with Celtic as BIN sponsor.
	Platform *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for consumer issuing cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticParams struct {
	// Terms of service acceptances for consumer issuing Apple Pay cards with Celtic as BIN sponsor.
	ApplePay *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for consumer issuing revolving credit cards with Celtic as BIN sponsor.
	RevolvingCreditCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticRevolvingCreditCardParams `form:"revolving_credit_card" json:"revolving_credit_card,omitempty"`
}

// Bank terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardParams struct {
	// Bank terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	FinancingDisclosures *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	Platform *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for consumer issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankParams struct {
	// Terms of service acceptances for consumer issuing prepaid cards with Cross River Bank as BIN sponsor.
	PrepaidCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Terms of service acceptances for Stripe consumer card Global issuing.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerGlobalAccountHolderParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing Apple Pay cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadApplePayParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Bank terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing debit cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardParams struct {
	// Bank terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
	FinancingDisclosures *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing debit cards with Lead as BIN sponsor.
	Platform *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Bank terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardBankTermsParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardFinancingDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Platform terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardPlatformParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Terms of service acceptances for consumer issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardParams struct {
	// Bank terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
	BankTerms *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardBankTermsParams `form:"bank_terms" json:"bank_terms,omitempty"`
	// Financial disclosures terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
	FinancingDisclosures *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardFinancingDisclosuresParams `form:"financing_disclosures" json:"financing_disclosures,omitempty"`
	// Platform terms of service acceptance for consumer issuing prepaid cards with Lead as BIN sponsor.
	Platform *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardPlatformParams `form:"platform" json:"platform,omitempty"`
}

// Terms of service acceptances for consumer issuing cards with Lead as BIN sponsor.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadParams struct {
	// Terms of service acceptances for consumer issuing Apple Pay cards with Lead as BIN sponsor.
	ApplePay *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadApplePayParams `form:"apple_pay" json:"apple_pay,omitempty"`
	// Terms of service acceptances for consumer issuing debit cards with Lead as BIN sponsor.
	DebitCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadDebitCardParams `form:"debit_card" json:"debit_card,omitempty"`
	// Terms of service acceptances for consumer issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Terms of service acceptances to create cards for consumer issuing use cases.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerParams struct {
	// Terms of service acceptances for Stripe consumer card issuing.
	AccountHolder *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerAccountHolderParams `form:"account_holder" json:"account_holder,omitempty"`
	// Terms of service acceptances for consumer issuing cards with Celtic as BIN sponsor.
	Celtic *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCelticParams `form:"celtic" json:"celtic,omitempty"`
	// Terms of service acceptances for consumer issuing cards with Cross River Bank as BIN sponsor.
	CrossRiverBank *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerCrossRiverBankParams `form:"cross_river_bank" json:"cross_river_bank,omitempty"`
	// Terms of service acceptances for Stripe consumer card Global issuing.
	GlobalAccountHolder *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerGlobalAccountHolderParams `form:"global_account_holder" json:"global_account_holder,omitempty"`
	// Terms of service acceptances for consumer issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadParams `form:"lead" json:"lead,omitempty"`
}

// Details on the Account's acceptance of Issuing-specific terms of service.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorParams struct {
	// Terms of service acceptances to create cards for commercial issuing use cases.
	Commercial *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialParams `form:"commercial" json:"commercial,omitempty"`
	// Terms of service acceptances to create cards for consumer issuing use cases.
	Consumer *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorConsumerParams `form:"consumer" json:"consumer,omitempty"`
}

// Details on the Account's acceptance of Consumer-specific terms of service.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceConsumerMoneyManagerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Consumer-privacy-disclosures-specific terms of service.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceConsumerPrivacyDisclosuresParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Crypto-specific terms of service.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCryptoMoneyManagerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Treasury-specific terms of service.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceMoneyManagerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Attestations of accepted terms of service agreements.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceParams struct {
	// Details on the Account's acceptance of the [Stripe Services Agreement](https://docs.stripe.com/connect/updating-accounts#tos-acceptance).
	Account *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceAccountParams `form:"account" json:"account,omitempty"`
	// Details on the Account's acceptance of Issuing-specific terms of service.
	CardCreator *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorParams `form:"card_creator" json:"card_creator,omitempty"`
	// Details on the Account's acceptance of Consumer-specific terms of service.
	ConsumerMoneyManager *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceConsumerMoneyManagerParams `form:"consumer_money_manager" json:"consumer_money_manager,omitempty"`
	// Details on the Account's acceptance of Consumer-privacy-disclosures-specific terms of service.
	ConsumerPrivacyDisclosures *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceConsumerPrivacyDisclosuresParams `form:"consumer_privacy_disclosures" json:"consumer_privacy_disclosures,omitempty"`
	// Details on the Account's acceptance of Crypto-specific terms of service.
	CryptoMoneyManager *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCryptoMoneyManagerParams `form:"crypto_money_manager" json:"crypto_money_manager,omitempty"`
	// Details on the Account's acceptance of Treasury-specific terms of service.
	MoneyManager *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceMoneyManagerParams `form:"money_manager" json:"money_manager,omitempty"`
}

// Attestations from the identity's key people, e.g. owners, executives, directors, representatives.
type V2CoreAccountUpdateIdentityAttestationsParams struct {
	// This hash is used to attest that the directors information provided to Stripe is both current and correct.
	DirectorshipDeclaration *V2CoreAccountUpdateIdentityAttestationsDirectorshipDeclarationParams `form:"directorship_declaration" json:"directorship_declaration,omitempty"`
	// This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.
	OwnershipDeclaration *V2CoreAccountUpdateIdentityAttestationsOwnershipDeclarationParams `form:"ownership_declaration" json:"ownership_declaration,omitempty"`
	// Attestation that all Persons with a specific Relationship value have been provided.
	PersonsProvided *V2CoreAccountUpdateIdentityAttestationsPersonsProvidedParams `form:"persons_provided" json:"persons_provided,omitempty"`
	// This hash is used to attest that the representative is authorized to act as the representative of their legal entity.
	RepresentativeDeclaration *V2CoreAccountUpdateIdentityAttestationsRepresentativeDeclarationParams `form:"representative_declaration" json:"representative_declaration,omitempty"`
	// Attestations of accepted terms of service agreements.
	TermsOfService *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceParams `form:"terms_of_service" json:"terms_of_service,omitempty"`
}

// The business registration address of the business entity.
type V2CoreAccountUpdateIdentityBusinessDetailsAddressParams struct {
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

// The business gross annual revenue for its preceding fiscal year.
type V2CoreAccountUpdateIdentityBusinessDetailsAnnualRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// The close-out date of the preceding fiscal year in ISO 8601 format. E.g. 2023-12-31 for the 31st of December, 2023.
	FiscalYearEnd *string `form:"fiscal_year_end" json:"fiscal_year_end,omitempty"`
}

// One or more documents that support the bank account ownership verification requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a check.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsBankAccountOwnershipVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's license to operate.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsCompanyLicenseParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the company's Memorandum of Association.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsCompanyMemorandumOfAssociationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Certain countries only: One or more documents showing the ministerial decree legalizing the company's establishment.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsCompanyMinisterialDecreeParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsCompanyRegistrationVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of a company's tax ID.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsCompanyTaxIDVerificationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document verifying the business.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of address.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfAddressParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Person that is signing the document.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfRegistrationSignerParams struct {
	// Person signing the document.
	Person *string `form:"person" json:"person"`
}

// One or more documents that demonstrate proof of ultimate beneficial ownership.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfRegistrationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// Person that is signing the document.
	Signer *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfRegistrationSignerParams `form:"signer" json:"signer,omitempty"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Person that is signing the document.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipSignerParams struct {
	// Person signing the document.
	Person *string `form:"person" json:"person"`
}

// One or more documents that demonstrate proof of ultimate beneficial ownership.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// Person that is signing the document.
	Signer *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipSignerParams `form:"signer" json:"signer,omitempty"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// A document verifying the business.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsParams struct {
	// One or more documents that support the bank account ownership verification requirement. Must be a document associated with the account's primary active bank account that displays the last 4 digits of the account number, either a statement or a check.
	BankAccountOwnershipVerification *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsBankAccountOwnershipVerificationParams `form:"bank_account_ownership_verification" json:"bank_account_ownership_verification,omitempty"`
	// One or more documents that demonstrate proof of a company's license to operate.
	CompanyLicense *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsCompanyLicenseParams `form:"company_license" json:"company_license,omitempty"`
	// One or more documents showing the company's Memorandum of Association.
	CompanyMemorandumOfAssociation *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsCompanyMemorandumOfAssociationParams `form:"company_memorandum_of_association" json:"company_memorandum_of_association,omitempty"`
	// Certain countries only: One or more documents showing the ministerial decree legalizing the company's establishment.
	CompanyMinisterialDecree *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsCompanyMinisterialDecreeParams `form:"company_ministerial_decree" json:"company_ministerial_decree,omitempty"`
	// One or more documents that demonstrate proof of a company's registration with the appropriate local authorities.
	CompanyRegistrationVerification *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsCompanyRegistrationVerificationParams `form:"company_registration_verification" json:"company_registration_verification,omitempty"`
	// One or more documents that demonstrate proof of a company's tax ID.
	CompanyTaxIDVerification *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsCompanyTaxIDVerificationParams `form:"company_tax_id_verification" json:"company_tax_id_verification,omitempty"`
	// A document verifying the business.
	PrimaryVerification *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// One or more documents that demonstrate proof of address.
	ProofOfAddress *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfAddressParams `form:"proof_of_address" json:"proof_of_address,omitempty"`
	// One or more documents that demonstrate proof of ultimate beneficial ownership.
	ProofOfRegistration *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfRegistrationParams `form:"proof_of_registration" json:"proof_of_registration,omitempty"`
	// One or more documents that demonstrate proof of ultimate beneficial ownership.
	ProofOfUltimateBeneficialOwnership *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams `form:"proof_of_ultimate_beneficial_ownership" json:"proof_of_ultimate_beneficial_ownership,omitempty"`
}

// The ID numbers of a business entity.
type V2CoreAccountUpdateIdentityBusinessDetailsIDNumberParams struct {
	// The registrar of the ID number (Only valid for DE ID number types).
	Registrar *string `form:"registrar" json:"registrar,omitempty"`
	// Open Enum. The ID number type of a business entity.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// An estimate of the monthly revenue of the business.
type V2CoreAccountUpdateIdentityBusinessDetailsMonthlyEstimatedRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
}

// When the business was incorporated or registered.
type V2CoreAccountUpdateIdentityBusinessDetailsRegistrationDateParams struct {
	// The day of registration, between 1 and 31.
	Day *int64 `form:"day" json:"day"`
	// The month of registration, between 1 and 12.
	Month *int64 `form:"month" json:"month"`
	// The four-digit year of registration.
	Year *int64 `form:"year" json:"year"`
}

// Kana Address.
type V2CoreAccountUpdateIdentityBusinessDetailsScriptAddressesKanaParams struct {
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
type V2CoreAccountUpdateIdentityBusinessDetailsScriptAddressesKanjiParams struct {
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
type V2CoreAccountUpdateIdentityBusinessDetailsScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountUpdateIdentityBusinessDetailsScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountUpdateIdentityBusinessDetailsScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Kana name.
type V2CoreAccountUpdateIdentityBusinessDetailsScriptNamesKanaParams struct {
	// Registered name of the business.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// Kanji name.
type V2CoreAccountUpdateIdentityBusinessDetailsScriptNamesKanjiParams struct {
	// Registered name of the business.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// The business legal name in non latin script.
type V2CoreAccountUpdateIdentityBusinessDetailsScriptNamesParams struct {
	// Kana name.
	Kana *V2CoreAccountUpdateIdentityBusinessDetailsScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji name.
	Kanji *V2CoreAccountUpdateIdentityBusinessDetailsScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Information about the company or business.
type V2CoreAccountUpdateIdentityBusinessDetailsParams struct {
	// The business registration address of the business entity.
	Address *V2CoreAccountUpdateIdentityBusinessDetailsAddressParams `form:"address" json:"address,omitempty"`
	// The business gross annual revenue for its preceding fiscal year.
	AnnualRevenue *V2CoreAccountUpdateIdentityBusinessDetailsAnnualRevenueParams `form:"annual_revenue" json:"annual_revenue,omitempty"`
	// A detailed description of the business's compliance and anti-money laundering controls and practices.
	ComplianceScreeningDescription *string `form:"compliance_screening_description" json:"compliance_screening_description,omitempty"`
	// A document verifying the business.
	Documents *V2CoreAccountUpdateIdentityBusinessDetailsDocumentsParams `form:"documents" json:"documents,omitempty"`
	// Estimated maximum number of workers currently engaged by the business (including employees, contractors, and vendors).
	EstimatedWorkerCount *int64 `form:"estimated_worker_count" json:"estimated_worker_count,omitempty"`
	// The ID numbers of a business entity.
	IDNumbers []*V2CoreAccountUpdateIdentityBusinessDetailsIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
	// An estimate of the monthly revenue of the business.
	MonthlyEstimatedRevenue *V2CoreAccountUpdateIdentityBusinessDetailsMonthlyEstimatedRevenueParams `form:"monthly_estimated_revenue" json:"monthly_estimated_revenue,omitempty"`
	// The phone number of the Business Entity.
	Phone *string `form:"phone" json:"phone,omitempty"`
	// The business legal name.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
	// When the business was incorporated or registered.
	RegistrationDate *V2CoreAccountUpdateIdentityBusinessDetailsRegistrationDateParams `form:"registration_date" json:"registration_date,omitempty"`
	// The business registration address of the business entity in non latin script.
	ScriptAddresses *V2CoreAccountUpdateIdentityBusinessDetailsScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The business legal name in non latin script.
	ScriptNames *V2CoreAccountUpdateIdentityBusinessDetailsScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The category identifying the legal structure of the business.
	Structure *string `form:"structure" json:"structure,omitempty"`
}

// Additional addresses associated with the individual.
type V2CoreAccountUpdateIdentityIndividualAdditionalAddressParams struct {
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
type V2CoreAccountUpdateIdentityIndividualAdditionalNameParams struct {
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
type V2CoreAccountUpdateIdentityIndividualAddressParams struct {
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
type V2CoreAccountUpdateIdentityIndividualDateOfBirthParams struct {
	// The day of the birth.
	Day *int64 `form:"day" json:"day"`
	// The month of birth.
	Month *int64 `form:"month" json:"month"`
	// The year of birth.
	Year *int64 `form:"year" json:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type V2CoreAccountUpdateIdentityIndividualDocumentsCompanyAuthorizationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's passport page with photo and personal data.
type V2CoreAccountUpdateIdentityIndividualDocumentsPassportParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountUpdateIdentityIndividualDocumentsPrimaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// An identifying document showing the person's name, either a passport or local ID card.
type V2CoreAccountUpdateIdentityIndividualDocumentsPrimaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountUpdateIdentityIndividualDocumentsPrimaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
type V2CoreAccountUpdateIdentityIndividualDocumentsSecondaryVerificationFrontBackParams struct {
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the back of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Back *string `form:"back" json:"back,omitempty"`
	// A [file upload](https://docs.stripe.com/api/persons/update#create_file) token representing the front of the verification document. The purpose of the uploaded file should be 'identity_document'. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.
	Front *string `form:"front" json:"front,omitempty"`
}

// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
type V2CoreAccountUpdateIdentityIndividualDocumentsSecondaryVerificationParams struct {
	// The [file upload](https://docs.stripe.com/api/persons/update#create_file) tokens referring to each side of the document.
	FrontBack *V2CoreAccountUpdateIdentityIndividualDocumentsSecondaryVerificationFrontBackParams `form:"front_back" json:"front_back"`
	// The format of the verification document. Currently supports `front_back` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type V2CoreAccountUpdateIdentityIndividualDocumentsVisaParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// Documents that may be submitted to satisfy various informational requests.
type V2CoreAccountUpdateIdentityIndividualDocumentsParams struct {
	// One or more documents that demonstrate proof that this person is authorized to represent the company.
	CompanyAuthorization *V2CoreAccountUpdateIdentityIndividualDocumentsCompanyAuthorizationParams `form:"company_authorization" json:"company_authorization,omitempty"`
	// One or more documents showing the person's passport page with photo and personal data.
	Passport *V2CoreAccountUpdateIdentityIndividualDocumentsPassportParams `form:"passport" json:"passport,omitempty"`
	// An identifying document showing the person's name, either a passport or local ID card.
	PrimaryVerification *V2CoreAccountUpdateIdentityIndividualDocumentsPrimaryVerificationParams `form:"primary_verification" json:"primary_verification,omitempty"`
	// A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.
	SecondaryVerification *V2CoreAccountUpdateIdentityIndividualDocumentsSecondaryVerificationParams `form:"secondary_verification" json:"secondary_verification,omitempty"`
	// One or more documents showing the person's visa required for living in the country where they are residing.
	Visa *V2CoreAccountUpdateIdentityIndividualDocumentsVisaParams `form:"visa" json:"visa,omitempty"`
}

// The identification numbers (e.g., SSN) associated with the individual.
type V2CoreAccountUpdateIdentityIndividualIDNumberParams struct {
	// The ID number type of an individual.
	Type *string `form:"type" json:"type"`
	// The value of the ID number.
	Value *string `form:"value" json:"value"`
}

// The relationship that this individual has with the account's identity.
type V2CoreAccountUpdateIdentityIndividualRelationshipParams struct {
	// Whether the person is a director of the account's identity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.
	Director *bool `form:"director" json:"director,omitempty"`
	// Whether the person has significant responsibility to control, manage, or direct the organization.
	Executive *bool `form:"executive" json:"executive,omitempty"`
	// Whether the person is an owner of the account's identity.
	Owner *bool `form:"owner" json:"owner,omitempty"`
	// The percent owned by the person of the account's legal entity.
	PercentOwnership *float64 `form:"percent_ownership,high_precision" json:"percent_ownership,string,omitempty"`
	// The person's title (e.g., CEO, Support Engineer).
	Title *string `form:"title" json:"title,omitempty"`
}

// Kana Address.
type V2CoreAccountUpdateIdentityIndividualScriptAddressesKanaParams struct {
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
type V2CoreAccountUpdateIdentityIndividualScriptAddressesKanjiParams struct {
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
type V2CoreAccountUpdateIdentityIndividualScriptAddressesParams struct {
	// Kana Address.
	Kana *V2CoreAccountUpdateIdentityIndividualScriptAddressesKanaParams `form:"kana" json:"kana,omitempty"`
	// Kanji Address.
	Kanji *V2CoreAccountUpdateIdentityIndividualScriptAddressesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Persons name in kana script.
type V2CoreAccountUpdateIdentityIndividualScriptNamesKanaParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// Persons name in kanji script.
type V2CoreAccountUpdateIdentityIndividualScriptNamesKanjiParams struct {
	// The person's first or given name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The person's last or family name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// The individuals primary name in non latin script.
type V2CoreAccountUpdateIdentityIndividualScriptNamesParams struct {
	// Persons name in kana script.
	Kana *V2CoreAccountUpdateIdentityIndividualScriptNamesKanaParams `form:"kana" json:"kana,omitempty"`
	// Persons name in kanji script.
	Kanji *V2CoreAccountUpdateIdentityIndividualScriptNamesKanjiParams `form:"kanji" json:"kanji,omitempty"`
}

// Information about the individual represented by the Account. This property is `null` unless `entity_type` is set to `individual`.
type V2CoreAccountUpdateIdentityIndividualParams struct {
	// Additional addresses associated with the individual.
	AdditionalAddresses []*V2CoreAccountUpdateIdentityIndividualAdditionalAddressParams `form:"additional_addresses" json:"additional_addresses,omitempty"`
	// Additional names (e.g. aliases) associated with the individual.
	AdditionalNames []*V2CoreAccountUpdateIdentityIndividualAdditionalNameParams `form:"additional_names" json:"additional_names,omitempty"`
	// The individual's residential address.
	Address *V2CoreAccountUpdateIdentityIndividualAddressParams `form:"address" json:"address,omitempty"`
	// The individual's date of birth.
	DateOfBirth *V2CoreAccountUpdateIdentityIndividualDateOfBirthParams `form:"date_of_birth" json:"date_of_birth,omitempty"`
	// Documents that may be submitted to satisfy various informational requests.
	Documents *V2CoreAccountUpdateIdentityIndividualDocumentsParams `form:"documents" json:"documents,omitempty"`
	// The individual's email address. You can only set this field when the Account is configured as a `merchant` or `recipient`. Use `contact_email` as the primary contact email for this Account.
	Email *string `form:"email" json:"email,omitempty"`
	// The individual's first name.
	GivenName *string `form:"given_name" json:"given_name,omitempty"`
	// The identification numbers (e.g., SSN) associated with the individual.
	IDNumbers []*V2CoreAccountUpdateIdentityIndividualIDNumberParams `form:"id_numbers" json:"id_numbers,omitempty"`
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
	Relationship *V2CoreAccountUpdateIdentityIndividualRelationshipParams `form:"relationship" json:"relationship,omitempty"`
	// The script addresses (e.g., non-Latin characters) associated with the individual.
	ScriptAddresses *V2CoreAccountUpdateIdentityIndividualScriptAddressesParams `form:"script_addresses" json:"script_addresses,omitempty"`
	// The individuals primary name in non latin script.
	ScriptNames *V2CoreAccountUpdateIdentityIndividualScriptNamesParams `form:"script_names" json:"script_names,omitempty"`
	// The individual's last name.
	Surname *string `form:"surname" json:"surname,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountUpdateIdentityIndividualParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Information about the company, individual, and business represented by the Account.
type V2CoreAccountUpdateIdentityParams struct {
	// Attestations from the identity's key people, e.g. owners, executives, directors, representatives.
	Attestations *V2CoreAccountUpdateIdentityAttestationsParams `form:"attestations" json:"attestations,omitempty"`
	// Information about the company or business.
	BusinessDetails *V2CoreAccountUpdateIdentityBusinessDetailsParams `form:"business_details" json:"business_details,omitempty"`
	// The country in which the account holder resides, or in which the business is legally established. This should be an [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code.
	Country *string `form:"country" json:"country,omitempty"`
	// The entity type represented by the Account. Ensure this field is accurate before adding configurations that rely on identity information, as it determines which identity fields apply and how the Account is validated.
	EntityType *string `form:"entity_type" json:"entity_type,omitempty"`
	// Information about the individual represented by the Account. This property is `null` unless `entity_type` is set to `individual`.
	Individual *V2CoreAccountUpdateIdentityIndividualParams `form:"individual" json:"individual,omitempty"`
}

// Updates the details of an Account.
type V2CoreAccountUpdateParams struct {
	Params `form:"*"`
	// The account token generated by the account token api.
	AccountToken *string `form:"account_token" json:"account_token,omitempty"`
	// An Account Configuration which allows the Account to take on a key persona across Stripe products.
	Configuration *V2CoreAccountUpdateConfigurationParams `form:"configuration" json:"configuration,omitempty"`
	// The primary contact email address for the Account.
	ContactEmail *string `form:"contact_email" json:"contact_email,omitempty"`
	// The default contact phone for the Account.
	ContactPhone *string `form:"contact_phone" json:"contact_phone,omitempty"`
	// A value indicating the Stripe dashboard this Account has access to. This will depend on which configurations are enabled for this account.
	Dashboard *string `form:"dashboard" json:"dashboard,omitempty"`
	// Default values to be used on Account Configurations.
	Defaults *V2CoreAccountUpdateDefaultsParams `form:"defaults" json:"defaults,omitempty"`
	// A descriptive name for the Account. This name will be surfaced in the Stripe Dashboard and on any invoices sent to the Account.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Information about the company, individual, and business represented by the Account.
	Identity *V2CoreAccountUpdateIdentityParams `form:"identity" json:"identity,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreAccountUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
