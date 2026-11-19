//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// How Stripe collects payment.
type V2BillingIntentBillingGroupDataBillingSettingsCollectionMethod string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsCollectionMethod can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsCollectionMethodChargeAutomatically V2BillingIntentBillingGroupDataBillingSettingsCollectionMethod = "charge_automatically"
	V2BillingIntentBillingGroupDataBillingSettingsCollectionMethodSendInvoice         V2BillingIntentBillingGroupDataBillingSettingsCollectionMethod = "send_invoice"
)

// The interval unit for the time until due.
type V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueInterval string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueInterval can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueIntervalDay   V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueInterval = "day"
	V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueIntervalMonth V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueInterval = "month"
	V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueIntervalWeek  V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueInterval = "week"
	V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueIntervalYear  V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueInterval = "year"
)

// Transaction type of the mandate.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Verification method.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethodInstant       V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// Preferred language of the Bancontact authorization page that the customer is redirected to.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguageDE V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage = "de"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguageEN V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage = "en"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguageFR V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage = "fr"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguageNL V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage = "nl"
)

// The AmountType for the mandate. One of `fixed` or `maximum`.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountType string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountType can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountTypeFixed   V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountType = "fixed"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountType = "maximum"
)

// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
// for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication).
// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
// Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardRequestThreeDSecureAny       V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure = "any"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardRequestThreeDSecureAutomatic V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardRequestThreeDSecureChallenge V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure = "challenge"
)

// The desired country code of the bank account information.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryBE V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "BE"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryDE V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "DE"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryES V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "ES"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryFR V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "FR"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryIE V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "IE"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryNL V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "NL"
)

// The bank transfer type that can be used for funding.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeEUBankTransfer V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "eu_bank_transfer"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeGBBankTransfer V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "gb_bank_transfer"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeJPBankTransfer V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "jp_bank_transfer"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeMXBankTransfer V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "mx_bank_transfer"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeUSBankTransfer V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "us_bank_transfer"
)

// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceFundingType string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceFundingType can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceFundingTypeBankTransfer V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceFundingType = "bank_transfer"
)

// The account subcategories to use to filter for selectable accounts.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategoryChecking V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "checking"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategorySavings  V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "savings"
)

// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionBalances      V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "balances"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionOwnership     V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "ownership"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionPaymentMethod V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "payment_method"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionTransactions  V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "transactions"
)

// List of data features that you would like to retrieve upon account creation.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchBalances     V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "balances"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchOwnership    V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "ownership"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchTransactions V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "transactions"
)

// Verification method.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethodAutomatic     V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "automatic"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethodInstant       V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "instant"
	V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethodMicrodeposits V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "microdeposits"
)

// Determines if tax is calculated automatically based on a product tax code or manually based on rules defined by the business.
type V2BillingIntentBillingGroupDataBillingSettingsTaxType string

// List of values that V2BillingIntentBillingGroupDataBillingSettingsTaxType can take
const (
	V2BillingIntentBillingGroupDataBillingSettingsTaxTypeAutomatic V2BillingIntentBillingGroupDataBillingSettingsTaxType = "automatic"
	V2BillingIntentBillingGroupDataBillingSettingsTaxTypeManual    V2BillingIntentBillingGroupDataBillingSettingsTaxType = "manual"
)

// The interval unit for the cadence.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceType string

// List of values that V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceType can take
const (
	V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceTypeDay   V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceType = "day"
	V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceTypeMonth V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceType = "month"
	V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceTypeWeek  V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceType = "week"
	V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceTypeYear  V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceType = "year"
)

// The grouping strategy.
type V2BillingIntentBillingGroupDataGroupingBehaviorType string

// List of values that V2BillingIntentBillingGroupDataGroupingBehaviorType can take
const (
	V2BillingIntentBillingGroupDataGroupingBehaviorTypeByCadence V2BillingIntentBillingGroupDataGroupingBehaviorType = "by_cadence"
)

// Current status of the Billing Intent.
type V2BillingIntentStatus string

// List of values that V2BillingIntentStatus can take
const (
	V2BillingIntentStatusCanceled  V2BillingIntentStatus = "canceled"
	V2BillingIntentStatusCommitted V2BillingIntentStatus = "committed"
	V2BillingIntentStatusDraft     V2BillingIntentStatus = "draft"
	V2BillingIntentStatusReserved  V2BillingIntentStatus = "reserved"
)

// Breakdown of the amount for this Billing Intent.
type V2BillingIntentAmountDetails struct {
	// The outstanding amount after discount, tax, and customer balance application.
	AmountDue string `json:"amount_due"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency Currency `json:"currency"`
	// The customer's account balance applied to the amount.
	CustomerBalanceApplied string `json:"customer_balance_applied"`
	// Amount of discount applied.
	Discount string `json:"discount"`
	// Amount of shipping charges.
	Shipping string `json:"shipping"`
	// Subtotal amount before tax and discounts.
	Subtotal string `json:"subtotal"`
	// Amount of tax.
	Tax string `json:"tax"`
	// Total amount for the Billing Intent.
	Total string `json:"total"`
}

// Controls emails for when the payment is due.
type V2BillingIntentBillingGroupDataBillingSettingsEmailDeliveryPaymentDue struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled bool `json:"enabled"`
	// If true the payment link to hosted invoice page would be included in email and PDF of the invoice.
	IncludePaymentLink bool `json:"include_payment_link"`
}

// Email delivery setting.
type V2BillingIntentBillingGroupDataBillingSettingsEmailDelivery struct {
	// Controls emails for when the payment is due.
	PaymentDue *V2BillingIntentBillingGroupDataBillingSettingsEmailDeliveryPaymentDue `json:"payment_due"`
}

// Settings for invoice presentation.
type V2BillingIntentBillingGroupDataBillingSettingsInvoicePresentation struct {
	// The ID of the invoice rendering template to be used when generating invoices.
	InvoiceRenderingTemplate string `json:"invoice_rendering_template,omitempty"`
}

// The amount of time until the invoice is overdue for payment.
type V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDue struct {
	// The interval unit for the time until due.
	Interval V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueInterval `json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30, the invoice is due in 30 days.
	IntervalCount int64 `json:"interval_count"`
}

// Additional fields for Mandate creation.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptions struct {
	// Transaction type of the mandate.
	TransactionType V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebit struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontact struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage `json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptions struct {
	// Amount to be charged for future payments.
	Amount int64 `json:"amount,string,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountType `json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description string `json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCard struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptions `json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network string `json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer struct {
	// The desired country code of the bank account information.
	Country V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry `json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransfer struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer `json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType `json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalance struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceFundingType `json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsKonbini struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsSEPADebit struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory `json:"account_subcategories"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters `json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch `json:"prefetch"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccount struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	// Verification method.
	VerificationMethod V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment Method specific configuration to be stored on the object.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptions struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebit `json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontact `json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCard `json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalance `json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsKonbini `json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsSEPADebit `json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccount `json:"us_bank_account,omitempty"`
}

// Settings for calculating tax.
type V2BillingIntentBillingGroupDataBillingSettingsTax struct {
	// Determines if tax is calculated automatically based on a product tax code or manually based on rules defined by the business.
	Type V2BillingIntentBillingGroupDataBillingSettingsTaxType `json:"type"`
}

// Billing setting overrides for the group. If omitted, the applicable defaults are used.
type V2BillingIntentBillingGroupDataBillingSettings struct {
	// How Stripe collects payment.
	CollectionMethod V2BillingIntentBillingGroupDataBillingSettingsCollectionMethod `json:"collection_method"`
	// Email delivery setting.
	EmailDelivery *V2BillingIntentBillingGroupDataBillingSettingsEmailDelivery `json:"email_delivery"`
	// Settings for invoice presentation.
	InvoicePresentation *V2BillingIntentBillingGroupDataBillingSettingsInvoicePresentation `json:"invoice_presentation,omitempty"`
	// The amount of time until the invoice is overdue for payment.
	InvoiceTimeUntilDue *V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDue `json:"invoice_time_until_due"`
	// The ID of the payment method configuration, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration string `json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration to be stored on the object.
	PaymentMethodOptions *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptions `json:"payment_method_options,omitempty"`
	// Settings for calculating tax.
	Tax *V2BillingIntentBillingGroupDataBillingSettingsTax `json:"tax"`
}

// The time at which the daily grouping interval begins.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceDayTime struct {
	// The hour, from 0 through 23.
	Hour int64 `json:"hour"`
	// The minute, from 0 through 59.
	Minute int64 `json:"minute"`
	// The second, from 0 through 59. Defaults to 0.
	Second int64 `json:"second,omitempty"`
}

// Daily cadence details.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceDay struct {
	// The time at which the daily grouping interval begins.
	Time *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceDayTime `json:"time"`
}

// The time at which the monthly grouping interval begins. Defaults to the time the billing group was created, in UTC.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceMonthTime struct {
	// The hour, from 0 through 23.
	Hour int64 `json:"hour"`
	// The minute, from 0 through 59.
	Minute int64 `json:"minute"`
	// The second, from 0 through 59. Defaults to 0.
	Second int64 `json:"second,omitempty"`
}

// Monthly cadence details.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceMonth struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth int64 `json:"day_of_month,omitempty"`
	// The month of year, from 1 through 12, from which occurrences are calculated.
	MonthOfYear int64 `json:"month_of_year,omitempty"`
	// The time at which the monthly grouping interval begins. Defaults to the time the billing group was created, in UTC.
	Time *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceMonthTime `json:"time"`
}

// The time at which the weekly grouping interval begins.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceWeekTime struct {
	// The hour, from 0 through 23.
	Hour int64 `json:"hour"`
	// The minute, from 0 through 59.
	Minute int64 `json:"minute"`
	// The second, from 0 through 59. Defaults to 0.
	Second int64 `json:"second,omitempty"`
}

// Weekly cadence details.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceWeek struct {
	// The ISO-8601 day of week, from 1 for Monday through 7 for Sunday.
	DayOfWeek int64 `json:"day_of_week,omitempty"`
	// The time at which the weekly grouping interval begins.
	Time *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceWeekTime `json:"time"`
}

// The time at which the yearly grouping interval begins. Defaults to the time the billing group was created, in UTC.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceYearTime struct {
	// The hour, from 0 through 23.
	Hour int64 `json:"hour"`
	// The minute, from 0 through 59.
	Minute int64 `json:"minute"`
	// The second, from 0 through 59. Defaults to 0.
	Second int64 `json:"second,omitempty"`
}

// Yearly cadence details.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceYear struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth int64 `json:"day_of_month,omitempty"`
	// The month of year, from 1 through 12.
	MonthOfYear int64 `json:"month_of_year,omitempty"`
	// The time at which the yearly grouping interval begins. Defaults to the time the billing group was created, in UTC.
	Time *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceYearTime `json:"time"`
}

// The cadence configuration when type is `by_cadence`.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadence struct {
	// Timestamp that anchors the recurring grouping cadence.
	BillingCycleAnchor time.Time `json:"billing_cycle_anchor,omitempty"`
	// Daily cadence details.
	Day *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceDay `json:"day,omitempty"`
	// The number of interval units between groups.
	IntervalCount int64 `json:"interval_count"`
	// Monthly cadence details.
	Month *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceMonth `json:"month,omitempty"`
	// The interval unit for the cadence.
	Type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceType `json:"type"`
	// Weekly cadence details.
	Week *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceWeek `json:"week,omitempty"`
	// Yearly cadence details.
	Year *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceYear `json:"year,omitempty"`
}

// The strategy used to group member invoices.
type V2BillingIntentBillingGroupDataGroupingBehavior struct {
	// The cadence configuration when type is `by_cadence`.
	ByCadence *V2BillingIntentBillingGroupDataGroupingBehaviorByCadence `json:"by_cadence,omitempty"`
	// The grouping strategy.
	Type V2BillingIntentBillingGroupDataGroupingBehaviorType `json:"type"`
}

// Data for creating a new Billing Group.
type V2BillingIntentBillingGroupData struct {
	// Billing setting overrides for the group. If omitted, the applicable defaults are used.
	BillingSettings *V2BillingIntentBillingGroupDataBillingSettings `json:"billing_settings,omitempty"`
	// The ID of the Customer object whose invoices are grouped.
	Customer string `json:"customer"`
	// The ID of the PaymentMethod object used by default. If unset, the customer's default payment method is used.
	DefaultPaymentMethod string `json:"default_payment_method,omitempty"`
	// The strategy used to group member invoices.
	GroupingBehavior *V2BillingIntentBillingGroupDataGroupingBehavior `json:"grouping_behavior,omitempty"`
	// An internal key you can use to find a particular billing group. Maximum length of 200 characters.
	LookupKey string `json:"lookup_key,omitempty"`
}

// Timestamps for status transitions of the Billing Intent.
type V2BillingIntentStatusTransitions struct {
	// Time at which the Billing Intent was canceled.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Time at which the Billing Intent was committed.
	CommittedAt time.Time `json:"committed_at,omitempty"`
	// Time at which the Billing Intent was drafted.
	DraftedAt time.Time `json:"drafted_at,omitempty"`
	// Time at which the Billing Intent expires.
	ExpiresAt time.Time `json:"expires_at"`
	// Time at which the Billing Intent was reserved.
	ReservedAt time.Time `json:"reserved_at,omitempty"`
}

// A Billing Intent represents a proposed change to a customer's billing configuration, such as subscribing to a new service,
// modifying an existing subscription, or canceling service. Intents follow a draft-reserve-commit workflow where they can be
// previewed before committing, allowing you to see the billing impact before changes take effect.
type V2BillingIntent struct {
	APIResource
	// Breakdown of the amount for this Billing Intent.
	AmountDetails *V2BillingIntentAmountDetails `json:"amount_details"`
	// ID of an existing Billing Group to use.
	BillingGroup string `json:"billing_group,omitempty"`
	// Data for creating a new Billing Group.
	BillingGroupData *V2BillingIntentBillingGroupData `json:"billing_group_data,omitempty"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Current status of the Billing Intent.
	Status V2BillingIntentStatus `json:"status"`
	// Timestamps for status transitions of the Billing Intent.
	StatusTransitions *V2BillingIntentStatusTransitions `json:"status_transitions"`
}
