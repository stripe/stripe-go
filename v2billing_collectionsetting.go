//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Either automatic, or send_invoice. When charging automatically, Stripe will attempt to pay this
// bill at the end of the period using the payment method attached to the payer profile. When sending an invoice,
// Stripe will email your payer profile an invoice with payment instructions.
// Defaults to automatic.
type V2BillingCollectionSettingCollectionMethod string

// List of values that V2BillingCollectionSettingCollectionMethod can take
const (
	V2BillingCollectionSettingCollectionMethodAutomatic   V2BillingCollectionSettingCollectionMethod = "automatic"
	V2BillingCollectionSettingCollectionMethodSendInvoice V2BillingCollectionSettingCollectionMethod = "send_invoice"
)

// Transaction type of the mandate.
type V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Verification method.
type V2BillingCollectionSettingPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     V2BillingCollectionSettingPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	V2BillingCollectionSettingPaymentMethodOptionsACSSDebitVerificationMethodInstant       V2BillingCollectionSettingPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	V2BillingCollectionSettingPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits V2BillingCollectionSettingPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// Preferred language of the Bancontact authorization page that the customer is redirected to.
type V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguage string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguage can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguageDE V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguage = "de"
	V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguageEN V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguage = "en"
	V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguageFR V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguage = "fr"
	V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguageNL V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguage = "nl"
)

// The AmountType for the mandate. One of `fixed` or `maximum`.
type V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsAmountType string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsAmountType can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsAmountTypeFixed   V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsAmountType = "fixed"
	V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsAmountType = "maximum"
)

// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
// for authentication based on risk level and [other requirements](https://docs.corp.stripe.com/strong-customer-authentication).
// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
// Read our guide on [manually requesting 3D Secure](https://docs.corp.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type V2BillingCollectionSettingPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsCardRequestThreeDSecureAny       V2BillingCollectionSettingPaymentMethodOptionsCardRequestThreeDSecure = "any"
	V2BillingCollectionSettingPaymentMethodOptionsCardRequestThreeDSecureAutomatic V2BillingCollectionSettingPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	V2BillingCollectionSettingPaymentMethodOptionsCardRequestThreeDSecureChallenge V2BillingCollectionSettingPaymentMethodOptionsCardRequestThreeDSecure = "challenge"
)

// The desired country code of the bank account information.
type V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryBE V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "BE"
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryDE V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "DE"
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryES V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "ES"
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryFR V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "FR"
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryIE V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "IE"
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryNL V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "NL"
)

// The bank transfer type that can be used for funding.
type V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferType string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferType can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferTypeEUBankTransfer V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferType = "eu_bank_transfer"
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferTypeGBBankTransfer V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferType = "gb_bank_transfer"
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferTypeJPBankTransfer V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferType = "jp_bank_transfer"
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferTypeMXBankTransfer V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferType = "mx_bank_transfer"
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferTypeUSBankTransfer V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferType = "us_bank_transfer"
)

// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
type V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceFundingType string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceFundingType can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceFundingTypeBankTransfer V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceFundingType = "bank_transfer"
)

// The account subcategories to use to filter for selectable accounts.
type V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategoryChecking V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "checking"
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategorySavings  V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "savings"
)

// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
type V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionBalances      V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "balances"
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionOwnership     V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "ownership"
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionPaymentMethod V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "payment_method"
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionTransactions  V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "transactions"
)

// List of data features that you would like to retrieve upon account creation.
type V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchBalances     V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "balances"
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchOwnership    V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "ownership"
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchTransactions V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "transactions"
)

// Verification method.
type V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountVerificationMethod string

// List of values that V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountVerificationMethod can take
const (
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountVerificationMethodAutomatic     V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountVerificationMethod = "automatic"
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountVerificationMethodInstant       V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountVerificationMethod = "instant"
	V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountVerificationMethodMicrodeposits V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountVerificationMethod = "microdeposits"
)

// Controls emails for when the payment is due. For example after the invoice is finilized and transition to Open state.
type V2BillingCollectionSettingEmailDeliveryPaymentDue struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled bool `json:"enabled"`
	// If true the payment link to hosted invocie page would be included in email and PDF of the invoice.
	IncludePaymentLink bool `json:"include_payment_link"`
}

// Email delivery settings.
type V2BillingCollectionSettingEmailDelivery struct {
	// Controls emails for when the payment is due. For example after the invoice is finilized and transition to Open state.
	PaymentDue *V2BillingCollectionSettingEmailDeliveryPaymentDue `json:"payment_due"`
}

// Additional fields for Mandate creation.
type V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptions struct {
	// Transaction type of the mandate.
	TransactionType V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsACSSDebit struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options"`
	// Verification method.
	VerificationMethod V2BillingCollectionSettingPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingCollectionSettingPaymentMethodOptionsBancontact struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage V2BillingCollectionSettingPaymentMethodOptionsBancontactPreferredLanguage `json:"preferred_language"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptions struct {
	// Amount to be charged for future payments.
	Amount int64 `json:"amount"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsAmountType `json:"amount_type"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description string `json:"description"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsCard struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptions `json:"mandate_options"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network string `json:"network"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.corp.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.corp.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure V2BillingCollectionSettingPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer struct {
	// The desired country code of the bank account information.
	Country V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry `json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransfer struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer `json:"eu_bank_transfer"`
	// The bank transfer type that can be used for funding.
	Type V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferType `json:"type"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsCustomerBalance struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceFundingType `json:"funding_type"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsKonbini struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsSEPADebit struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory `json:"account_subcategories"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters `json:"filters"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch `json:"prefetch"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsUSBankAccount struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	// Verification method.
	VerificationMethod V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment Method specific configuration stored on the object.
type V2BillingCollectionSettingPaymentMethodOptions struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingCollectionSettingPaymentMethodOptionsACSSDebit `json:"acss_debit"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingCollectionSettingPaymentMethodOptionsBancontact `json:"bancontact"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingCollectionSettingPaymentMethodOptionsCard `json:"card"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingCollectionSettingPaymentMethodOptionsCustomerBalance `json:"customer_balance"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingCollectionSettingPaymentMethodOptionsKonbini `json:"konbini"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingCollectionSettingPaymentMethodOptionsSEPADebit `json:"sepa_debit"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingCollectionSettingPaymentMethodOptionsUSBankAccount `json:"us_bank_account"`
}

// Settings that configure and manage the behavior of collecting payments.
type V2BillingCollectionSetting struct {
	APIResource
	// Either automatic, or send_invoice. When charging automatically, Stripe will attempt to pay this
	// bill at the end of the period using the payment method attached to the payer profile. When sending an invoice,
	// Stripe will email your payer profile an invoice with payment instructions.
	// Defaults to automatic.
	CollectionMethod V2BillingCollectionSettingCollectionMethod `json:"collection_method"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// An optional field for adding a display name for the CollectionSetting object.
	DisplayName string `json:"display_name"`
	// Email delivery settings.
	EmailDelivery *V2BillingCollectionSettingEmailDelivery `json:"email_delivery"`
	// The ID of the CollectionSetting.
	ID string `json:"id"`
	// The latest version of the current settings object. This will be
	// Updated every time an attribute of the settings is updated.
	LatestVersion string `json:"latest_version"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The current live version of the settings object. This can be different from
	// latest_version if settings are updated without setting live_version='latest'.
	LiveVersion string `json:"live_version"`
	// A lookup key used to retrieve settings dynamically from a static string.
	// This may be up to 200 characters.
	LookupKey string `json:"lookup_key"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the PaymentMethodConfiguration object, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration string `json:"payment_method_configuration"`
	// Payment Method specific configuration stored on the object.
	PaymentMethodOptions *V2BillingCollectionSettingPaymentMethodOptions `json:"payment_method_options"`
}
