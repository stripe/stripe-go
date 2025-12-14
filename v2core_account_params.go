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

// Can create commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardParams struct {
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

// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams struct {
	// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
	ChargeCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Stripe as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadParams struct {
	// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Can create commercial issuing charge cards with Stripe as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing prepaid cards with Stripe as BIN sponsor.
type V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardParams struct {
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
	// Can create commercial issuing cards with Stripe as BIN sponsor.
	Lead *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadParams `form:"lead" json:"lead,omitempty"`
	// Can create commercial issuing cards with Stripe as BIN sponsor.
	Stripe *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeParams `form:"stripe" json:"stripe,omitempty"`
}

// Capabilities to request on the CardCreator Configuration.
type V2CoreAccountConfigurationCardCreatorCapabilitiesParams struct {
	// Can create cards for commercial issuing use cases.
	Commercial *V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialParams `form:"commercial" json:"commercial,omitempty"`
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
	// ID of a PaymentMethod attached to the customer account to use as the default for invoices and subscriptions.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// Default invoice settings for the customer account.
	Invoice *V2CoreAccountConfigurationCustomerBillingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// Generates requirements for enabling automatic indirect tax calculation on this customer's invoices or subscriptions. Recommended to request this capability if planning to enable automatic tax calculation on this customer's invoices or subscriptions.
type V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTaxParams struct {
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

// The Customer Configuration allows the Account to be used in inbound payment flows.
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

// Allow the merchant to process ACH debit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesACHDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process ACSS debit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesACSSDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Affirm payments.
type V2CoreAccountConfigurationMerchantCapabilitiesAffirmPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Afterpay/Clearpay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Alma payments.
type V2CoreAccountConfigurationMerchantCapabilitiesAlmaPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Amazon Pay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Australian BECS Direct Debit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesAUBECSDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process BACS Direct Debit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesBACSDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Bancontact payments.
type V2CoreAccountConfigurationMerchantCapabilitiesBancontactPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process BLIK payments.
type V2CoreAccountConfigurationMerchantCapabilitiesBLIKPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Boleto payments.
type V2CoreAccountConfigurationMerchantCapabilitiesBoletoPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to collect card payments.
type V2CoreAccountConfigurationMerchantCapabilitiesCardPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Cartes Bancaires payments.
type V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Cash App payments.
type V2CoreAccountConfigurationMerchantCapabilitiesCashAppPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process EPS payments.
type V2CoreAccountConfigurationMerchantCapabilitiesEPSPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process FPX payments.
type V2CoreAccountConfigurationMerchantCapabilitiesFPXPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process UK bank transfer payments.
type V2CoreAccountConfigurationMerchantCapabilitiesGBBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process GrabPay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process iDEAL payments.
type V2CoreAccountConfigurationMerchantCapabilitiesIDEALPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process JCB card payments.
type V2CoreAccountConfigurationMerchantCapabilitiesJCBPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Japanese bank transfer payments.
type V2CoreAccountConfigurationMerchantCapabilitiesJPBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Kakao Pay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Klarna payments.
type V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Konbini convenience store payments.
type V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Korean card payments.
type V2CoreAccountConfigurationMerchantCapabilitiesKrCardPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Link payments.
type V2CoreAccountConfigurationMerchantCapabilitiesLinkPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process MobilePay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Multibanco payments.
type V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Mexican bank transfer payments.
type V2CoreAccountConfigurationMerchantCapabilitiesMXBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Naver Pay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process OXXO payments.
type V2CoreAccountConfigurationMerchantCapabilitiesOXXOPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Przelewy24 (P24) payments.
type V2CoreAccountConfigurationMerchantCapabilitiesP24PaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Pay by Bank payments.
type V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process PAYCO payments.
type V2CoreAccountConfigurationMerchantCapabilitiesPaycoPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process PayNow payments.
type V2CoreAccountConfigurationMerchantCapabilitiesPayNowPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process PromptPay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesPromptPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Revolut Pay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Samsung Pay payments.
type V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process SEPA bank transfer payments.
type V2CoreAccountConfigurationMerchantCapabilitiesSEPABankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process SEPA Direct Debit payments.
type V2CoreAccountConfigurationMerchantCapabilitiesSEPADebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Swish payments.
type V2CoreAccountConfigurationMerchantCapabilitiesSwishPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process TWINT payments.
type V2CoreAccountConfigurationMerchantCapabilitiesTWINTPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process US bank transfer payments.
type V2CoreAccountConfigurationMerchantCapabilitiesUSBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Zip payments.
type V2CoreAccountConfigurationMerchantCapabilitiesZipPaymentsParams struct {
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
	// Settings for the default [statement descriptor](https://docs.stripe.com/connect/statement-descriptors) text.
	StatementDescriptor *V2CoreAccountConfigurationMerchantStatementDescriptorParams `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Publicly available contact information for sending support issues to.
	Support *V2CoreAccountConfigurationMerchantSupportParams `form:"support" json:"support,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsInstantParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocalParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWireParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities that enable OutboundPayments to a bank account linked to this Account.
type V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsParams struct {
	// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
	Instant *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsInstantParams `form:"instant" json:"instant,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
	Local *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocalParams `form:"local" json:"local,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
	Wire *V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWireParams `form:"wire" json:"wire,omitempty"`
}

// Capabilities that enable OutboundPayments to a card linked to this Account.
type V2CoreAccountConfigurationRecipientCapabilitiesCardsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities that enable OutboundPayments to a crypto wallet linked to this Account.
type V2CoreAccountConfigurationRecipientCapabilitiesCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Enables this Account to receive /v1/transfers into their Stripe Balance (/v1/balance).
type V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams struct {
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

// Can provision a bank-account-like financial address (VBAN) to credit/debit a FinancialAccount.
type V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can provision a crypto wallet like financial address to credit a FinancialAccount.
type V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can provision a financial address to credit/debit a FinancialAccount.
type V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesParams struct {
	// Can provision a bank-account-like financial address (VBAN) to credit/debit a FinancialAccount.
	BankAccounts *V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can provision a crypto wallet like financial address to credit a FinancialAccount.
	CryptoWallets *V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
}

// Can hold storage-type funds on Stripe in EUR.
type V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesEURParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can hold storage-type funds on Stripe in GBP.
type V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesGBPParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can hold storage-type funds on Stripe in USD.
type V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUSDParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can hold storage-type funds on Stripe in USDC.
type V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUsdcParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can hold storage-type funds on Stripe.
type V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesParams struct {
	// Can hold storage-type funds on Stripe in EUR.
	EUR *V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesEURParams `form:"eur" json:"eur,omitempty"`
	// Can hold storage-type funds on Stripe in GBP.
	GBP *V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesGBPParams `form:"gbp" json:"gbp,omitempty"`
	// Can hold storage-type funds on Stripe in USD.
	USD *V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUSDParams `form:"usd" json:"usd,omitempty"`
	// Can hold storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
type V2CoreAccountConfigurationStorerCapabilitiesInboundTransfersBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
type V2CoreAccountConfigurationStorerCapabilitiesInboundTransfersParams struct {
	// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
	BankAccounts *V2CoreAccountConfigurationStorerCapabilitiesInboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
}

// Can send funds from a FinancialAccount to a bank account owned by someone else.
type V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a debit card owned by someone else.
type V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCardsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
type V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
type V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsFinancialAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a destination owned by someone else.
type V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by someone else.
	BankAccounts *V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a debit card owned by someone else.
	Cards *V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCardsParams `form:"cards" json:"cards,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
	CryptoWallets *V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
	FinancialAccounts *V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
}

// Can send funds from a FinancialAccount to a bank account owned by yourself.
type V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
type V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
type V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersFinancialAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a destination owned by yourself.
type V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by yourself.
	BankAccounts *V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
	CryptoWallets *V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
	FinancialAccounts *V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
}

// Capabilities to request on the Storer Configuration.
type V2CoreAccountConfigurationStorerCapabilitiesParams struct {
	// Can provision a financial address to credit/debit a FinancialAccount.
	FinancialAddresses *V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesParams `form:"financial_addresses" json:"financial_addresses,omitempty"`
	// Can hold storage-type funds on Stripe.
	HoldsCurrencies *V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesParams `form:"holds_currencies" json:"holds_currencies,omitempty"`
	// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
	InboundTransfers *V2CoreAccountConfigurationStorerCapabilitiesInboundTransfersParams `form:"inbound_transfers" json:"inbound_transfers,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by someone else.
	OutboundPayments *V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsParams `form:"outbound_payments" json:"outbound_payments,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by yourself.
	OutboundTransfers *V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersParams `form:"outbound_transfers" json:"outbound_transfers,omitempty"`
}

// Details of the regulated activity if the business participates in one.
type V2CoreAccountConfigurationStorerRegulatedActivityParams struct {
	// A detailed description of the regulated activities the business is licensed to conduct.
	Description *string `form:"description" json:"description,omitempty"`
	// The license number or registration number assigned by the business's primary regulator.
	LicenseNumber *string `form:"license_number" json:"license_number,omitempty"`
	// The country of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityCountry *string `form:"primary_regulatory_authority_country" json:"primary_regulatory_authority_country,omitempty"`
	// The name of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityName *string `form:"primary_regulatory_authority_name" json:"primary_regulatory_authority_name,omitempty"`
}

// The Storer Configuration allows the Account to store and move funds using stored-value FinancialAccounts.
type V2CoreAccountConfigurationStorerParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Capabilities to request on the Storer Configuration.
	Capabilities *V2CoreAccountConfigurationStorerCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
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
	RegulatedActivity *V2CoreAccountConfigurationStorerRegulatedActivityParams `form:"regulated_activity" json:"regulated_activity,omitempty"`
	// The source of funds for the business, e.g. profits, income, venture capital, etc.
	SourceOfFunds *string `form:"source_of_funds" json:"source_of_funds,omitempty"`
	// Description of the source of funds for the business' account.
	SourceOfFundsDescription *string `form:"source_of_funds_description" json:"source_of_funds_description,omitempty"`
}

// An Account Configuration which allows the Account to take on a key persona across Stripe products.
type V2CoreAccountConfigurationParams struct {
	// The CardCreator Configuration allows the Account to create and issue cards to users.
	CardCreator *V2CoreAccountConfigurationCardCreatorParams `form:"card_creator" json:"card_creator,omitempty"`
	// The Customer Configuration allows the Account to be used in inbound payment flows.
	Customer *V2CoreAccountConfigurationCustomerParams `form:"customer" json:"customer,omitempty"`
	// Enables the Account to act as a connected account and collect payments facilitated by a Connect platform. You must onboard your platform to Connect before you can add this configuration to your connected accounts. Utilize this configuration when the Account will be the Merchant of Record, like with Direct charges or Destination Charges with on_behalf_of set.
	Merchant *V2CoreAccountConfigurationMerchantParams `form:"merchant" json:"merchant,omitempty"`
	// The Recipient Configuration allows the Account to receive funds. Utilize this configuration if the Account will not be the Merchant of Record, like with Separate Charges & Transfers, or Destination Charges without on_behalf_of set.
	Recipient *V2CoreAccountConfigurationRecipientParams `form:"recipient" json:"recipient,omitempty"`
	// The Storer Configuration allows the Account to store and move funds using stored-value FinancialAccounts.
	Storer *V2CoreAccountConfigurationStorerParams `form:"storer" json:"storer,omitempty"`
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
	// Terms of service acceptances for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
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
	// Terms of service acceptances for Stripe commercial card Global issuing.
	GlobalAccountHolder *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialGlobalAccountHolderParams `form:"global_account_holder" json:"global_account_holder,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadParams `form:"lead" json:"lead,omitempty"`
}

// Details on the Account's acceptance of Issuing-specific terms of service.
type V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorParams struct {
	// Terms of service acceptances to create cards for commercial issuing use cases.
	Commercial *V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialParams `form:"commercial" json:"commercial,omitempty"`
}

// Details on the Account's acceptance of Crypto-storer-specific terms of service.
type V2CoreAccountIdentityAttestationsTermsOfServiceCryptoStorerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Treasury-specific terms of service.
type V2CoreAccountIdentityAttestationsTermsOfServiceStorerParams struct {
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
	// Details on the Account's acceptance of Crypto-storer-specific terms of service.
	CryptoStorer *V2CoreAccountIdentityAttestationsTermsOfServiceCryptoStorerParams `form:"crypto_storer" json:"crypto_storer,omitempty"`
	// Details on the Account's acceptance of Treasury-specific terms of service.
	Storer *V2CoreAccountIdentityAttestationsTermsOfServiceStorerParams `form:"storer" json:"storer,omitempty"`
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

// A non-negative integer representing the amount in the smallest currency unit.
type V2CoreAccountIdentityBusinessDetailsAnnualRevenueAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// The business gross annual revenue for its preceding fiscal year.
type V2CoreAccountIdentityBusinessDetailsAnnualRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *V2CoreAccountIdentityBusinessDetailsAnnualRevenueAmountParams `form:"amount" json:"amount,omitempty"`
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

// One or more documents showing the company's proof of registration with the national business registry.
type V2CoreAccountIdentityBusinessDetailsDocumentsProofOfRegistrationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of ultimate beneficial ownership.
type V2CoreAccountIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
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

// A non-negative integer representing the amount in the smallest currency unit.
type V2CoreAccountIdentityBusinessDetailsMonthlyEstimatedRevenueAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// An estimate of the monthly revenue of the business.
type V2CoreAccountIdentityBusinessDetailsMonthlyEstimatedRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *V2CoreAccountIdentityBusinessDetailsMonthlyEstimatedRevenueAmountParams `form:"amount" json:"amount,omitempty"`
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
	PercentOwnership *string `form:"percent_ownership" json:"percent_ownership,omitempty"`
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
	// The individual's email address.
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
	// The entity type.
	EntityType *string `form:"entity_type" json:"entity_type,omitempty"`
	// Information about the individual represented by the Account. This property is `null` unless `entity_type` is set to `individual`.
	Individual *V2CoreAccountIdentityIndividualParams `form:"individual" json:"individual,omitempty"`
}

// An Account is a representation of a company, individual or other entity that a user interacts with. Accounts contain identifying information about the entity, and configurations that store the features an account has access to. An account can be configured as any or all of the following configurations: Customer, Merchant and/or Recipient.
type V2CoreAccountParams struct {
	Params `form:"*"`
	// The account token generated by the account token api.
	AccountToken *string `form:"account_token" json:"account_token,omitempty"`
	// An Account Configuration which allows the Account to take on a key persona across Stripe products.
	Configuration *V2CoreAccountConfigurationParams `form:"configuration" json:"configuration,omitempty"`
	// The default contact email address for the Account. Required when configuring the account as a merchant or recipient.
	ContactEmail *string `form:"contact_email" json:"contact_email,omitempty"`
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

// Removes access to the Account and its associated resources. Closed Accounts can no longer be operated on, but limited information can still be retrieved through the API in order to be able to track their history.
type V2CoreAccountCloseParams struct {
	Params `form:"*"`
	// Configurations on the Account to be closed. All configurations on the Account must be passed in for this request to succeed.
	AppliedConfigurations []*string `form:"applied_configurations" json:"applied_configurations,omitempty"`
}

// Can create commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardParams struct {
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

// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create commercial issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams struct {
	// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
	ChargeCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create commercial issuing cards with Stripe as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadParams struct {
	// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Can create commercial issuing charge cards with Stripe as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can create commercial issuing prepaid cards with Stripe as BIN sponsor.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardParams struct {
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
	// Can create commercial issuing cards with Stripe as BIN sponsor.
	Lead *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialLeadParams `form:"lead" json:"lead,omitempty"`
	// Can create commercial issuing cards with Stripe as BIN sponsor.
	Stripe *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialStripeParams `form:"stripe" json:"stripe,omitempty"`
}

// Capabilities to request on the CardCreator Configuration.
type V2CoreAccountCreateConfigurationCardCreatorCapabilitiesParams struct {
	// Can create cards for commercial issuing use cases.
	Commercial *V2CoreAccountCreateConfigurationCardCreatorCapabilitiesCommercialParams `form:"commercial" json:"commercial,omitempty"`
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

// Generates requirements for enabling automatic indirect tax calculation on this customer's invoices or subscriptions. Recommended to request this capability if planning to enable automatic tax calculation on this customer's invoices or subscriptions.
type V2CoreAccountCreateConfigurationCustomerCapabilitiesAutomaticIndirectTaxParams struct {
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

// The Customer Configuration allows the Account to be used in inbound payment flows.
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

// Allow the merchant to process ACH debit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesACHDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process ACSS debit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesACSSDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Affirm payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAffirmPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Afterpay/Clearpay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Alma payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAlmaPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Amazon Pay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAmazonPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Australian BECS Direct Debit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process BACS Direct Debit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBACSDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Bancontact payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBancontactPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process BLIK payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBLIKPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Boleto payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesBoletoPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to collect card payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCardPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Cartes Bancaires payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCartesBancairesPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Cash App payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesCashAppPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process EPS payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesEPSPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process FPX payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesFPXPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process UK bank transfer payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesGBBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process GrabPay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesGrabpayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process iDEAL payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesIDEALPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process JCB card payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesJCBPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Japanese bank transfer payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesJPBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Kakao Pay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKakaoPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Klarna payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKlarnaPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Konbini convenience store payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKonbiniPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Korean card payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesKrCardPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Link payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesLinkPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process MobilePay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMobilepayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Multibanco payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMultibancoPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Mexican bank transfer payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesMXBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Naver Pay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesNaverPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process OXXO payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesOXXOPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Przelewy24 (P24) payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesP24PaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Pay by Bank payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPayByBankPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process PAYCO payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPaycoPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process PayNow payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPayNowPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process PromptPay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesPromptPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Revolut Pay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesRevolutPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Samsung Pay payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSamsungPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process SEPA bank transfer payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process SEPA Direct Debit payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSEPADebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Swish payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesSwishPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process TWINT payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesTWINTPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process US bank transfer payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesUSBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Allow the merchant to process Zip payments.
type V2CoreAccountCreateConfigurationMerchantCapabilitiesZipPaymentsParams struct {
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
	// Statement descriptor.
	StatementDescriptor *V2CoreAccountCreateConfigurationMerchantStatementDescriptorParams `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Publicly available contact information for sending support issues to.
	Support *V2CoreAccountCreateConfigurationMerchantSupportParams `form:"support" json:"support,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsInstantParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsLocalParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsWireParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Capabilities that enable OutboundPayments to a bank account linked to this Account.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsParams struct {
	// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
	Instant *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsInstantParams `form:"instant" json:"instant,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
	Local *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsLocalParams `form:"local" json:"local,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
	Wire *V2CoreAccountCreateConfigurationRecipientCapabilitiesBankAccountsWireParams `form:"wire" json:"wire,omitempty"`
}

// Capabilities that enable OutboundPayments to a card linked to this Account.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesCardsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Capabilities that enable OutboundPayments to a crypto wallet linked to this Account.
type V2CoreAccountCreateConfigurationRecipientCapabilitiesCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Enables this Account to receive /v1/transfers into their Stripe Balance (/v1/balance).
type V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams struct {
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
	// Capabilities that enable the recipient to manage their Stripe Balance (/v1/balance).
	StripeBalance *V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceParams `form:"stripe_balance" json:"stripe_balance,omitempty"`
}

// The Recipient Configuration allows the Account to receive funds. Utilize this configuration if the Account will not be the Merchant of Record, like with Separate Charges & Transfers, or Destination Charges without on_behalf_of set.
type V2CoreAccountCreateConfigurationRecipientParams struct {
	// Capabilities to be requested on the Recipient Configuration.
	Capabilities *V2CoreAccountCreateConfigurationRecipientCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
}

// Can provision a bank-account-like financial address (VBAN) to credit/debit a FinancialAccount.
type V2CoreAccountCreateConfigurationStorerCapabilitiesFinancialAddressesBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can provision a crypto wallet like financial address to credit a FinancialAccount.
type V2CoreAccountCreateConfigurationStorerCapabilitiesFinancialAddressesCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can provision a financial address to credit/debit a FinancialAccount.
type V2CoreAccountCreateConfigurationStorerCapabilitiesFinancialAddressesParams struct {
	// Can provision a bank-account-like financial address (VBAN) to credit/debit a FinancialAccount.
	BankAccounts *V2CoreAccountCreateConfigurationStorerCapabilitiesFinancialAddressesBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can provision a crypto wallet like financial address to credit a FinancialAccount.
	CryptoWallets *V2CoreAccountCreateConfigurationStorerCapabilitiesFinancialAddressesCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
}

// Can hold storage-type funds on Stripe in EUR.
type V2CoreAccountCreateConfigurationStorerCapabilitiesHoldsCurrenciesEURParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can hold storage-type funds on Stripe in GBP.
type V2CoreAccountCreateConfigurationStorerCapabilitiesHoldsCurrenciesGBPParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can hold storage-type funds on Stripe in USD.
type V2CoreAccountCreateConfigurationStorerCapabilitiesHoldsCurrenciesUSDParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can hold storage-type funds on Stripe in USDC.
type V2CoreAccountCreateConfigurationStorerCapabilitiesHoldsCurrenciesUsdcParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can hold storage-type funds on Stripe.
type V2CoreAccountCreateConfigurationStorerCapabilitiesHoldsCurrenciesParams struct {
	// Can hold storage-type funds on Stripe in EUR.
	EUR *V2CoreAccountCreateConfigurationStorerCapabilitiesHoldsCurrenciesEURParams `form:"eur" json:"eur,omitempty"`
	// Can hold storage-type funds on Stripe in GBP.
	GBP *V2CoreAccountCreateConfigurationStorerCapabilitiesHoldsCurrenciesGBPParams `form:"gbp" json:"gbp,omitempty"`
	// Can hold storage-type funds on Stripe in USD.
	USD *V2CoreAccountCreateConfigurationStorerCapabilitiesHoldsCurrenciesUSDParams `form:"usd" json:"usd,omitempty"`
	// Can hold storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountCreateConfigurationStorerCapabilitiesHoldsCurrenciesUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
type V2CoreAccountCreateConfigurationStorerCapabilitiesInboundTransfersBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
type V2CoreAccountCreateConfigurationStorerCapabilitiesInboundTransfersParams struct {
	// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
	BankAccounts *V2CoreAccountCreateConfigurationStorerCapabilitiesInboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
}

// Can send funds from a FinancialAccount to a bank account owned by someone else.
type V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundPaymentsBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send funds from a FinancialAccount to a debit card owned by someone else.
type V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundPaymentsCardsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
type V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundPaymentsCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
type V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundPaymentsFinancialAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send funds from a FinancialAccount to a destination owned by someone else.
type V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundPaymentsParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by someone else.
	BankAccounts *V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundPaymentsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a debit card owned by someone else.
	Cards *V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundPaymentsCardsParams `form:"cards" json:"cards,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
	CryptoWallets *V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundPaymentsCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
	FinancialAccounts *V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundPaymentsFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
}

// Can send funds from a FinancialAccount to a bank account owned by yourself.
type V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundTransfersBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
type V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundTransfersCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
type V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundTransfersFinancialAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested"`
}

// Can send funds from a FinancialAccount to a destination owned by yourself.
type V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundTransfersParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by yourself.
	BankAccounts *V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
	CryptoWallets *V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundTransfersCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
	FinancialAccounts *V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundTransfersFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
}

// Capabilities to request on the Storer Configuration.
type V2CoreAccountCreateConfigurationStorerCapabilitiesParams struct {
	// Can provision a financial address to credit/debit a FinancialAccount.
	FinancialAddresses *V2CoreAccountCreateConfigurationStorerCapabilitiesFinancialAddressesParams `form:"financial_addresses" json:"financial_addresses,omitempty"`
	// Can hold storage-type funds on Stripe.
	HoldsCurrencies *V2CoreAccountCreateConfigurationStorerCapabilitiesHoldsCurrenciesParams `form:"holds_currencies" json:"holds_currencies,omitempty"`
	// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
	InboundTransfers *V2CoreAccountCreateConfigurationStorerCapabilitiesInboundTransfersParams `form:"inbound_transfers" json:"inbound_transfers,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by someone else.
	OutboundPayments *V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundPaymentsParams `form:"outbound_payments" json:"outbound_payments,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by yourself.
	OutboundTransfers *V2CoreAccountCreateConfigurationStorerCapabilitiesOutboundTransfersParams `form:"outbound_transfers" json:"outbound_transfers,omitempty"`
}

// Details of the regulated activity if the business participates in one.
type V2CoreAccountCreateConfigurationStorerRegulatedActivityParams struct {
	// A detailed description of the regulated activities the business is licensed to conduct.
	Description *string `form:"description" json:"description,omitempty"`
	// The license number or registration number assigned by the business's primary regulator.
	LicenseNumber *string `form:"license_number" json:"license_number,omitempty"`
	// The country of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityCountry *string `form:"primary_regulatory_authority_country" json:"primary_regulatory_authority_country,omitempty"`
	// The name of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityName *string `form:"primary_regulatory_authority_name" json:"primary_regulatory_authority_name,omitempty"`
}

// The Storer Configuration allows the Account to store and move funds using stored-value FinancialAccounts.
type V2CoreAccountCreateConfigurationStorerParams struct {
	// Capabilities to request on the Storer Configuration.
	Capabilities *V2CoreAccountCreateConfigurationStorerCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
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
	RegulatedActivity *V2CoreAccountCreateConfigurationStorerRegulatedActivityParams `form:"regulated_activity" json:"regulated_activity,omitempty"`
	// The source of funds for the business, e.g. profits, income, venture capital, etc.
	SourceOfFunds *string `form:"source_of_funds" json:"source_of_funds,omitempty"`
	// Description of the source of funds for the business' account.
	SourceOfFundsDescription *string `form:"source_of_funds_description" json:"source_of_funds_description,omitempty"`
}

// An Account Configuration which allows the Account to take on a key persona across Stripe products.
type V2CoreAccountCreateConfigurationParams struct {
	// The CardCreator Configuration allows the Account to create and issue cards to users.
	CardCreator *V2CoreAccountCreateConfigurationCardCreatorParams `form:"card_creator" json:"card_creator,omitempty"`
	// The Customer Configuration allows the Account to be used in inbound payment flows.
	Customer *V2CoreAccountCreateConfigurationCustomerParams `form:"customer" json:"customer,omitempty"`
	// Enables the Account to act as a connected account and collect payments facilitated by a Connect platform. You must onboard your platform to Connect before you can add this configuration to your connected accounts. Utilize this configuration when the Account will be the Merchant of Record, like with Direct charges or Destination Charges with on_behalf_of set.
	Merchant *V2CoreAccountCreateConfigurationMerchantParams `form:"merchant" json:"merchant,omitempty"`
	// The Recipient Configuration allows the Account to receive funds. Utilize this configuration if the Account will not be the Merchant of Record, like with Separate Charges & Transfers, or Destination Charges without on_behalf_of set.
	Recipient *V2CoreAccountCreateConfigurationRecipientParams `form:"recipient" json:"recipient,omitempty"`
	// The Storer Configuration allows the Account to store and move funds using stored-value FinancialAccounts.
	Storer *V2CoreAccountCreateConfigurationStorerParams `form:"storer" json:"storer,omitempty"`
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
	// Terms of service acceptances for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
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
	// Terms of service acceptances for Stripe commercial card Global issuing.
	GlobalAccountHolder *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialGlobalAccountHolderParams `form:"global_account_holder" json:"global_account_holder,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadParams `form:"lead" json:"lead,omitempty"`
}

// Details on the Account's acceptance of Issuing-specific terms of service.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorParams struct {
	// Terms of service acceptances to create cards for commercial issuing use cases.
	Commercial *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCardCreatorCommercialParams `form:"commercial" json:"commercial,omitempty"`
}

// Details on the Account's acceptance of Crypto-storer-specific terms of service.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceCryptoStorerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Treasury-specific terms of service.
type V2CoreAccountCreateIdentityAttestationsTermsOfServiceStorerParams struct {
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
	// Details on the Account's acceptance of Crypto-storer-specific terms of service.
	CryptoStorer *V2CoreAccountCreateIdentityAttestationsTermsOfServiceCryptoStorerParams `form:"crypto_storer" json:"crypto_storer,omitempty"`
	// Details on the Account's acceptance of Treasury-specific terms of service.
	Storer *V2CoreAccountCreateIdentityAttestationsTermsOfServiceStorerParams `form:"storer" json:"storer,omitempty"`
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

// A non-negative integer representing the amount in the smallest currency unit.
type V2CoreAccountCreateIdentityBusinessDetailsAnnualRevenueAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// The business gross annual revenue for its preceding fiscal year.
type V2CoreAccountCreateIdentityBusinessDetailsAnnualRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *V2CoreAccountCreateIdentityBusinessDetailsAnnualRevenueAmountParams `form:"amount" json:"amount,omitempty"`
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

// One or more documents showing the company's proof of registration with the national business registry.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfRegistrationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of ultimate beneficial ownership.
type V2CoreAccountCreateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
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

// A non-negative integer representing the amount in the smallest currency unit.
type V2CoreAccountCreateIdentityBusinessDetailsMonthlyEstimatedRevenueAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// An estimate of the monthly revenue of the business.
type V2CoreAccountCreateIdentityBusinessDetailsMonthlyEstimatedRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *V2CoreAccountCreateIdentityBusinessDetailsMonthlyEstimatedRevenueAmountParams `form:"amount" json:"amount,omitempty"`
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
	PercentOwnership *string `form:"percent_ownership" json:"percent_ownership,omitempty"`
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
	// The individual's email address.
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
	// The entity type.
	EntityType *string `form:"entity_type" json:"entity_type,omitempty"`
	// Information about the person represented by the account.
	Individual *V2CoreAccountCreateIdentityIndividualParams `form:"individual" json:"individual,omitempty"`
}

// An Account is a representation of a company, individual or other entity that a user interacts with. Accounts contain identifying information about the entity, and configurations that store the features an account has access to. An account can be configured as any or all of the following configurations: Customer, Merchant and/or Recipient.
type V2CoreAccountCreateParams struct {
	Params `form:"*"`
	// The account token generated by the account token api.
	AccountToken *string `form:"account_token" json:"account_token,omitempty"`
	// An Account Configuration which allows the Account to take on a key persona across Stripe products.
	Configuration *V2CoreAccountCreateConfigurationParams `form:"configuration" json:"configuration,omitempty"`
	// The default contact email address for the Account. Required when configuring the account as a merchant or recipient.
	ContactEmail *string `form:"contact_email" json:"contact_email,omitempty"`
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

// Can create commercial issuing charge cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing spend cards with Celtic as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardParams struct {
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

// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Cross River Bank as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams struct {
	// Can create commercial issuing charge cards with Cross River Bank as BIN sponsor.
	ChargeCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams `form:"charge_card" json:"charge_card,omitempty"`
	// Can create commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
}

// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing cards with Lead as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadParams struct {
	// Can create commercial issuing prepaid cards with Lead as BIN sponsor.
	PrepaidCard *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams `form:"prepaid_card" json:"prepaid_card,omitempty"`
}

// Can create commercial issuing charge cards with Stripe as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can create commercial issuing prepaid cards with Stripe as BIN sponsor.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardParams struct {
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
	// Can create commercial issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialLeadParams `form:"lead" json:"lead,omitempty"`
	// Can create commercial issuing cards with Stripe as BIN sponsor.
	Stripe *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialStripeParams `form:"stripe" json:"stripe,omitempty"`
}

// Capabilities to request on the CardCreator Configuration.
type V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesParams struct {
	// Can create cards for commercial issuing use cases.
	Commercial *V2CoreAccountUpdateConfigurationCardCreatorCapabilitiesCommercialParams `form:"commercial" json:"commercial,omitempty"`
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
	// ID of a PaymentMethod attached to the customer account to use as the default for invoices and subscriptions.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// Default invoice settings for the customer account.
	Invoice *V2CoreAccountUpdateConfigurationCustomerBillingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// Generates requirements for enabling automatic indirect tax calculation on this customer's invoices or subscriptions. Recommended to request this capability if planning to enable automatic tax calculation on this customer's invoices or subscriptions.
type V2CoreAccountUpdateConfigurationCustomerCapabilitiesAutomaticIndirectTaxParams struct {
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

// Allow the merchant to process ACH debit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesACHDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process ACSS debit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesACSSDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Affirm payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAffirmPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Afterpay/Clearpay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Alma payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAlmaPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Amazon Pay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAmazonPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Australian BECS Direct Debit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesAUBECSDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process BACS Direct Debit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBACSDebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Bancontact payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBancontactPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process BLIK payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBLIKPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Boleto payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesBoletoPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to collect card payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCardPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Cartes Bancaires payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCartesBancairesPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Cash App payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesCashAppPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process EPS payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesEPSPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process FPX payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesFPXPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process UK bank transfer payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesGBBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process GrabPay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesGrabpayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process iDEAL payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesIDEALPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process JCB card payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesJCBPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Japanese bank transfer payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesJPBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Kakao Pay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKakaoPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Klarna payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKlarnaPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Konbini convenience store payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKonbiniPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Korean card payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesKrCardPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Link payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesLinkPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process MobilePay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMobilepayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Multibanco payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMultibancoPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Mexican bank transfer payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesMXBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Naver Pay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesNaverPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process OXXO payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesOXXOPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Przelewy24 (P24) payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesP24PaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Pay by Bank payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayByBankPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process PAYCO payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPaycoPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process PayNow payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPayNowPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process PromptPay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesPromptPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Revolut Pay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesRevolutPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Samsung Pay payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSamsungPayPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process SEPA bank transfer payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPABankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process SEPA Direct Debit payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSEPADebitPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Swish payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesSwishPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process TWINT payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesTWINTPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process US bank transfer payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesUSBankTransferPaymentsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Allow the merchant to process Zip payments.
type V2CoreAccountUpdateConfigurationMerchantCapabilitiesZipPaymentsParams struct {
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
	// Settings for the default [statement descriptor](https://docs.stripe.com/connect/statement-descriptors) text.
	StatementDescriptor *V2CoreAccountUpdateConfigurationMerchantStatementDescriptorParams `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	// Publicly available contact information for sending support issues to.
	Support *V2CoreAccountUpdateConfigurationMerchantSupportParams `form:"support" json:"support,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsInstantParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsLocalParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsWireParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities that enable OutboundPayments to a bank account linked to this Account.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsParams struct {
	// Enables this Account to receive OutboundPayments to linked bank accounts over real time rails.
	Instant *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsInstantParams `form:"instant" json:"instant,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over local networks.
	Local *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsLocalParams `form:"local" json:"local,omitempty"`
	// Enables this Account to receive OutboundPayments to linked bank accounts over wire.
	Wire *V2CoreAccountUpdateConfigurationRecipientCapabilitiesBankAccountsWireParams `form:"wire" json:"wire,omitempty"`
}

// Capability that enable OutboundPayments to a debit card linked to this Account.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesCardsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Capabilities that enable OutboundPayments to a crypto wallet linked to this Account.
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Enables this Account to receive /v1/transfers into their Stripe Balance (/v1/balance).
type V2CoreAccountUpdateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams struct {
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

// Can provision a bank-account-like financial address (VBAN) to credit/debit a FinancialAccount.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesFinancialAddressesBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can provision a crypto wallet like financial address to credit a FinancialAccount.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesFinancialAddressesCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can provision a financial address to credit/debit a FinancialAccount.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesFinancialAddressesParams struct {
	// Can provision a bank-account-like financial address (VBAN) to credit/debit a FinancialAccount.
	BankAccounts *V2CoreAccountUpdateConfigurationStorerCapabilitiesFinancialAddressesBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can provision a crypto wallet like financial address to credit a FinancialAccount.
	CryptoWallets *V2CoreAccountUpdateConfigurationStorerCapabilitiesFinancialAddressesCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
}

// Can hold storage-type funds on Stripe in EUR.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesHoldsCurrenciesEURParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can hold storage-type funds on Stripe in GBP.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesHoldsCurrenciesGBPParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can hold storage-type funds on Stripe in USD.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesHoldsCurrenciesUSDParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can hold storage-type funds on Stripe in USDC.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesHoldsCurrenciesUsdcParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can hold storage-type funds on Stripe.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesHoldsCurrenciesParams struct {
	// Can hold storage-type funds on Stripe in EUR.
	EUR *V2CoreAccountUpdateConfigurationStorerCapabilitiesHoldsCurrenciesEURParams `form:"eur" json:"eur,omitempty"`
	// Can hold storage-type funds on Stripe in GBP.
	GBP *V2CoreAccountUpdateConfigurationStorerCapabilitiesHoldsCurrenciesGBPParams `form:"gbp" json:"gbp,omitempty"`
	// Can hold storage-type funds on Stripe in USD.
	USD *V2CoreAccountUpdateConfigurationStorerCapabilitiesHoldsCurrenciesUSDParams `form:"usd" json:"usd,omitempty"`
	// Can hold storage-type funds on Stripe in USDC.
	Usdc *V2CoreAccountUpdateConfigurationStorerCapabilitiesHoldsCurrenciesUsdcParams `form:"usdc" json:"usdc,omitempty"`
}

// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesInboundTransfersBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesInboundTransfersParams struct {
	// Can pull funds from an external bank account owned by yourself to a FinancialAccount.
	BankAccounts *V2CoreAccountUpdateConfigurationStorerCapabilitiesInboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
}

// Can send funds from a FinancialAccount to a bank account owned by someone else.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundPaymentsBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a debit card owned by someone else.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundPaymentsCardsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundPaymentsCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundPaymentsFinancialAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a destination owned by someone else.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundPaymentsParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by someone else.
	BankAccounts *V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundPaymentsBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a debit card owned by someone else.
	Cards *V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundPaymentsCardsParams `form:"cards" json:"cards,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by someone else.
	CryptoWallets *V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundPaymentsCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by someone else.
	FinancialAccounts *V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundPaymentsFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
}

// Can send funds from a FinancialAccount to a bank account owned by yourself.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundTransfersBankAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundTransfersCryptoWalletsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundTransfersFinancialAccountsParams struct {
	// To request a new Capability for an account, pass true. There can be a delay before the requested Capability becomes active.
	Requested *bool `form:"requested" json:"requested,omitempty"`
}

// Can send funds from a FinancialAccount to a destination owned by yourself.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundTransfersParams struct {
	// Can send funds from a FinancialAccount to a bank account owned by yourself.
	BankAccounts *V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundTransfersBankAccountsParams `form:"bank_accounts" json:"bank_accounts,omitempty"`
	// Can send funds from a FinancialAccount to a crypto wallet owned by yourself.
	CryptoWallets *V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundTransfersCryptoWalletsParams `form:"crypto_wallets" json:"crypto_wallets,omitempty"`
	// Can send funds from a FinancialAccount to another FinancialAccount owned by yourself.
	FinancialAccounts *V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundTransfersFinancialAccountsParams `form:"financial_accounts" json:"financial_accounts,omitempty"`
}

// Capabilities to request on the Storer Configuration.
type V2CoreAccountUpdateConfigurationStorerCapabilitiesParams struct {
	// Can provision a financial address to credit/debit a FinancialAccount.
	FinancialAddresses *V2CoreAccountUpdateConfigurationStorerCapabilitiesFinancialAddressesParams `form:"financial_addresses" json:"financial_addresses,omitempty"`
	// Can hold storage-type funds on Stripe.
	HoldsCurrencies *V2CoreAccountUpdateConfigurationStorerCapabilitiesHoldsCurrenciesParams `form:"holds_currencies" json:"holds_currencies,omitempty"`
	// Can pull funds from an external source, owned by yourself, to a FinancialAccount.
	InboundTransfers *V2CoreAccountUpdateConfigurationStorerCapabilitiesInboundTransfersParams `form:"inbound_transfers" json:"inbound_transfers,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by someone else.
	OutboundPayments *V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundPaymentsParams `form:"outbound_payments" json:"outbound_payments,omitempty"`
	// Can send funds from a FinancialAccount to a destination owned by yourself.
	OutboundTransfers *V2CoreAccountUpdateConfigurationStorerCapabilitiesOutboundTransfersParams `form:"outbound_transfers" json:"outbound_transfers,omitempty"`
}

// Details of the regulated activity if the business participates in one.
type V2CoreAccountUpdateConfigurationStorerRegulatedActivityParams struct {
	// A detailed description of the regulated activities the business is licensed to conduct.
	Description *string `form:"description" json:"description,omitempty"`
	// The license number or registration number assigned by the business's primary regulator.
	LicenseNumber *string `form:"license_number" json:"license_number,omitempty"`
	// The country of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityCountry *string `form:"primary_regulatory_authority_country" json:"primary_regulatory_authority_country,omitempty"`
	// The name of the primary regulatory authority that oversees the business's regulated activities.
	PrimaryRegulatoryAuthorityName *string `form:"primary_regulatory_authority_name" json:"primary_regulatory_authority_name,omitempty"`
}

// The Storer Configuration allows the Account to store and move funds using stored-value FinancialAccounts.
type V2CoreAccountUpdateConfigurationStorerParams struct {
	// Represents the state of the configuration, and can be updated to deactivate or re-apply a configuration.
	Applied *bool `form:"applied" json:"applied,omitempty"`
	// Capabilities to request on the Storer Configuration.
	Capabilities *V2CoreAccountUpdateConfigurationStorerCapabilitiesParams `form:"capabilities" json:"capabilities,omitempty"`
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
	RegulatedActivity *V2CoreAccountUpdateConfigurationStorerRegulatedActivityParams `form:"regulated_activity" json:"regulated_activity,omitempty"`
	// The source of funds for the business, e.g. profits, income, venture capital, etc.
	SourceOfFunds *string `form:"source_of_funds" json:"source_of_funds,omitempty"`
	// Description of the source of funds for the business' account.
	SourceOfFundsDescription *string `form:"source_of_funds_description" json:"source_of_funds_description,omitempty"`
}

// An Account Configuration which allows the Account to take on a key persona across Stripe products.
type V2CoreAccountUpdateConfigurationParams struct {
	// The CardCreator Configuration allows the Account to create and issue cards to users.
	CardCreator *V2CoreAccountUpdateConfigurationCardCreatorParams `form:"card_creator" json:"card_creator,omitempty"`
	// The Customer Configuration allows the Account to be charged.
	Customer *V2CoreAccountUpdateConfigurationCustomerParams `form:"customer" json:"customer,omitempty"`
	// Enables the Account to act as a connected account and collect payments facilitated by a Connect platform. You must onboard your platform to Connect before you can add this configuration to your connected accounts. Utilize this configuration when the Account will be the Merchant of Record, like with Direct charges or Destination Charges with on_behalf_of set.
	Merchant *V2CoreAccountUpdateConfigurationMerchantParams `form:"merchant" json:"merchant,omitempty"`
	// The Recipient Configuration allows the Account to receive funds. Utilize this configuration if the Account will not be the Merchant of Record, like with Separate Charges & Transfers, or Destination Charges without on_behalf_of set.
	Recipient *V2CoreAccountUpdateConfigurationRecipientParams `form:"recipient" json:"recipient,omitempty"`
	// The Storer Configuration allows the Account to store and move funds using stored-value FinancialAccounts.
	Storer *V2CoreAccountUpdateConfigurationStorerParams `form:"storer" json:"storer,omitempty"`
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
	// Terms of service acceptances for commercial issuing spend cards with Cross River Bank as BIN sponsor.
	SpendCard *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankSpendCardParams `form:"spend_card" json:"spend_card,omitempty"`
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
	// Terms of service acceptances for Stripe commercial card Global issuing.
	GlobalAccountHolder *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialGlobalAccountHolderParams `form:"global_account_holder" json:"global_account_holder,omitempty"`
	// Terms of service acceptances for commercial issuing cards with Lead as BIN sponsor.
	Lead *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialLeadParams `form:"lead" json:"lead,omitempty"`
}

// Details on the Account's acceptance of Issuing-specific terms of service.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorParams struct {
	// Terms of service acceptances to create cards for commercial issuing use cases.
	Commercial *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCardCreatorCommercialParams `form:"commercial" json:"commercial,omitempty"`
}

// Details on the Account's acceptance of Crypto-storer-specific terms of service.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCryptoStorerParams struct {
	// The time when the Account's representative accepted the terms of service. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Date *time.Time `form:"date" json:"date,omitempty"`
	// The IP address from which the Account's representative accepted the terms of service.
	IP *string `form:"ip" json:"ip,omitempty"`
	// The user agent of the browser from which the Account's representative accepted the terms of service.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Details on the Account's acceptance of Treasury-specific terms of service.
type V2CoreAccountUpdateIdentityAttestationsTermsOfServiceStorerParams struct {
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
	// Details on the Account's acceptance of Crypto-storer-specific terms of service.
	CryptoStorer *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceCryptoStorerParams `form:"crypto_storer" json:"crypto_storer,omitempty"`
	// Details on the Account's acceptance of Treasury-specific terms of service.
	Storer *V2CoreAccountUpdateIdentityAttestationsTermsOfServiceStorerParams `form:"storer" json:"storer,omitempty"`
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

// A non-negative integer representing the amount in the smallest currency unit.
type V2CoreAccountUpdateIdentityBusinessDetailsAnnualRevenueAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// The business gross annual revenue for its preceding fiscal year.
type V2CoreAccountUpdateIdentityBusinessDetailsAnnualRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *V2CoreAccountUpdateIdentityBusinessDetailsAnnualRevenueAmountParams `form:"amount" json:"amount,omitempty"`
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

// One or more documents showing the company's proof of registration with the national business registry.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfRegistrationParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
	// The format of the document. Currently supports `files` only.
	Type *string `form:"type" json:"type"`
}

// One or more documents that demonstrate proof of ultimate beneficial ownership.
type V2CoreAccountUpdateIdentityBusinessDetailsDocumentsProofOfUltimateBeneficialOwnershipParams struct {
	// One or more document IDs returned by a [file upload](https://docs.stripe.com/api/persons/update#create_file) with a purpose value of `account_requirement`.
	Files []*string `form:"files" json:"files"`
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
	// One or more documents showing the company's proof of registration with the national business registry.
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

// A non-negative integer representing the amount in the smallest currency unit.
type V2CoreAccountUpdateIdentityBusinessDetailsMonthlyEstimatedRevenueAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// An estimate of the monthly revenue of the business.
type V2CoreAccountUpdateIdentityBusinessDetailsMonthlyEstimatedRevenueParams struct {
	// A non-negative integer representing the amount in the smallest currency unit.
	Amount *V2CoreAccountUpdateIdentityBusinessDetailsMonthlyEstimatedRevenueAmountParams `form:"amount" json:"amount,omitempty"`
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
	PercentOwnership *string `form:"percent_ownership" json:"percent_ownership,omitempty"`
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
	// The individual's email address.
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
	// The entity type.
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
	// The default contact email address for the Account. Required when configuring the account as a merchant or recipient.
	ContactEmail *string `form:"contact_email" json:"contact_email,omitempty"`
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
