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
type V2BillingCollectionSettingVersionCollectionMethod string

// List of values that V2BillingCollectionSettingVersionCollectionMethod can take
const (
	V2BillingCollectionSettingVersionCollectionMethodAutomatic   V2BillingCollectionSettingVersionCollectionMethod = "automatic"
	V2BillingCollectionSettingVersionCollectionMethodSendInvoice V2BillingCollectionSettingVersionCollectionMethod = "send_invoice"
)

// Transaction type of the mandate.
type V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Verification method.
type V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitVerificationMethodInstant       V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// Preferred language of the Bancontact authorization page that the customer is redirected to.
type V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguage string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguage can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguageDE V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguage = "de"
	V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguageEN V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguage = "en"
	V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguageFR V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguage = "fr"
	V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguageNL V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguage = "nl"
)

// The AmountType for the mandate. One of `fixed` or `maximum`.
type V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptionsAmountType string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptionsAmountType can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptionsAmountTypeFixed   V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptionsAmountType = "fixed"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptionsAmountType = "maximum"
)

// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
// for authentication based on risk level and [other requirements](https://docs.corp.stripe.com/strong-customer-authentication).
// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
// Read our guide on [manually requesting 3D Secure](https://docs.corp.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type V2BillingCollectionSettingVersionPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsCardRequestThreeDSecureAny       V2BillingCollectionSettingVersionPaymentMethodOptionsCardRequestThreeDSecure = "any"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCardRequestThreeDSecureAutomatic V2BillingCollectionSettingVersionPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCardRequestThreeDSecureChallenge V2BillingCollectionSettingVersionPaymentMethodOptionsCardRequestThreeDSecure = "challenge"
)

// The desired country code of the bank account information.
type V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryBE V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "BE"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryDE V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "DE"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryES V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "ES"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryFR V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "FR"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryIE V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "IE"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryNL V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "NL"
)

// The bank transfer type that can be used for funding.
type V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferType string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferType can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferTypeEUBankTransfer V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferType = "eu_bank_transfer"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferTypeGBBankTransfer V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferType = "gb_bank_transfer"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferTypeJPBankTransfer V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferType = "jp_bank_transfer"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferTypeMXBankTransfer V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferType = "mx_bank_transfer"
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferTypeUSBankTransfer V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferType = "us_bank_transfer"
)

// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
type V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceFundingType string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceFundingType can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceFundingTypeBankTransfer V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceFundingType = "bank_transfer"
)

// The account subcategories to use to filter for selectable accounts.
type V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategoryChecking V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "checking"
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategorySavings  V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "savings"
)

// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
type V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionBalances      V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "balances"
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionOwnership     V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "ownership"
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionPaymentMethod V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "payment_method"
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionTransactions  V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "transactions"
)

// List of data features that you would like to retrieve upon account creation.
type V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchBalances     V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "balances"
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchOwnership    V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "ownership"
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchTransactions V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "transactions"
)

// Verification method.
type V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountVerificationMethod string

// List of values that V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountVerificationMethod can take
const (
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountVerificationMethodAutomatic     V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountVerificationMethod = "automatic"
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountVerificationMethodInstant       V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountVerificationMethod = "instant"
	V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountVerificationMethodMicrodeposits V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountVerificationMethod = "microdeposits"
)

// Controls emails for when the payment is due. For example after the invoice is finalized and transitions to Open state.
type V2BillingCollectionSettingVersionEmailDeliveryPaymentDue struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled bool `json:"enabled"`
	// If true the payment link to hosted invoice page would be included in email and PDF of the invoice.
	IncludePaymentLink bool `json:"include_payment_link"`
}

// Email delivery settings.
type V2BillingCollectionSettingVersionEmailDelivery struct {
	// Controls emails for when the payment is due. For example after the invoice is finalized and transitions to Open state.
	PaymentDue *V2BillingCollectionSettingVersionEmailDeliveryPaymentDue `json:"payment_due,omitempty"`
}

// Additional fields for Mandate creation.
type V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitMandateOptions struct {
	// Transaction type of the mandate.
	TransactionType V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebit struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingCollectionSettingVersionPaymentMethodOptionsBancontact struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage V2BillingCollectionSettingVersionPaymentMethodOptionsBancontactPreferredLanguage `json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptions struct {
	// Amount to be charged for future payments.
	Amount int64 `json:"amount,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptionsAmountType `json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description string `json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingCollectionSettingVersionPaymentMethodOptionsCard struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptions `json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network string `json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.corp.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.corp.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure V2BillingCollectionSettingVersionPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer struct {
	// The desired country code of the bank account information.
	Country V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry `json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransfer struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer `json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransferType `json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalance struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalanceFundingType `json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingCollectionSettingVersionPaymentMethodOptionsKonbini struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingCollectionSettingVersionPaymentMethodOptionsSEPADebit struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory `json:"account_subcategories"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters `json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch `json:"prefetch"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccount struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	// Verification method.
	VerificationMethod V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment Method specific configuration stored on the object.
type V2BillingCollectionSettingVersionPaymentMethodOptions struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingCollectionSettingVersionPaymentMethodOptionsACSSDebit `json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingCollectionSettingVersionPaymentMethodOptionsBancontact `json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingCollectionSettingVersionPaymentMethodOptionsCard `json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingCollectionSettingVersionPaymentMethodOptionsCustomerBalance `json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingCollectionSettingVersionPaymentMethodOptionsKonbini `json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingCollectionSettingVersionPaymentMethodOptionsSEPADebit `json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingCollectionSettingVersionPaymentMethodOptionsUSBankAccount `json:"us_bank_account,omitempty"`
}
type V2BillingCollectionSettingVersion struct {
	APIResource
	// Either automatic, or send_invoice. When charging automatically, Stripe will attempt to pay this
	// bill at the end of the period using the payment method attached to the payer profile. When sending an invoice,
	// Stripe will email your payer profile an invoice with payment instructions.
	// Defaults to automatic.
	CollectionMethod V2BillingCollectionSettingVersionCollectionMethod `json:"collection_method,omitempty"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Email delivery settings.
	EmailDelivery *V2BillingCollectionSettingVersionEmailDelivery `json:"email_delivery,omitempty"`
	// The ID of the CollectionSettingVersion object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the PaymentMethodConfiguration object, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration string `json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration stored on the object.
	PaymentMethodOptions *V2BillingCollectionSettingVersionPaymentMethodOptions `json:"payment_method_options,omitempty"`
}
