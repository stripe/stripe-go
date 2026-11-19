//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"github.com/shopspring/decimal"
	"time"
)

// How Stripe collects payment.
type V2BillingGroupBillingSettingsCollectionMethod string

// List of values that V2BillingGroupBillingSettingsCollectionMethod can take
const (
	V2BillingGroupBillingSettingsCollectionMethodChargeAutomatically V2BillingGroupBillingSettingsCollectionMethod = "charge_automatically"
	V2BillingGroupBillingSettingsCollectionMethodSendInvoice         V2BillingGroupBillingSettingsCollectionMethod = "send_invoice"
)

// The interval unit for the time until due.
type V2BillingGroupBillingSettingsInvoiceTimeUntilDueInterval string

// List of values that V2BillingGroupBillingSettingsInvoiceTimeUntilDueInterval can take
const (
	V2BillingGroupBillingSettingsInvoiceTimeUntilDueIntervalDay   V2BillingGroupBillingSettingsInvoiceTimeUntilDueInterval = "day"
	V2BillingGroupBillingSettingsInvoiceTimeUntilDueIntervalMonth V2BillingGroupBillingSettingsInvoiceTimeUntilDueInterval = "month"
	V2BillingGroupBillingSettingsInvoiceTimeUntilDueIntervalWeek  V2BillingGroupBillingSettingsInvoiceTimeUntilDueInterval = "week"
	V2BillingGroupBillingSettingsInvoiceTimeUntilDueIntervalYear  V2BillingGroupBillingSettingsInvoiceTimeUntilDueInterval = "year"
)

// Transaction type of the mandate.
type V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Verification method.
type V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethodInstant       V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// Preferred language of the Bancontact authorization page that the customer is redirected to.
type V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguageDE V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage = "de"
	V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguageEN V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage = "en"
	V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguageFR V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage = "fr"
	V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguageNL V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage = "nl"
)

// The AmountType for the mandate. One of `fixed` or `maximum`.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountType string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountType can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountTypeFixed   V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountType = "fixed"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountType = "maximum"
)

// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
// for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication).
// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
// Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsCardRequestThreeDSecureAny       V2BillingGroupBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure = "any"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCardRequestThreeDSecureAutomatic V2BillingGroupBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCardRequestThreeDSecureChallenge V2BillingGroupBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure = "challenge"
)

// The desired country code of the bank account information.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryBE V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "BE"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryDE V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "DE"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryES V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "ES"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryFR V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "FR"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryIE V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "IE"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryNL V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "NL"
)

// The bank transfer type that can be used for funding.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeEUBankTransfer V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "eu_bank_transfer"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeGBBankTransfer V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "gb_bank_transfer"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeJPBankTransfer V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "jp_bank_transfer"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeMXBankTransfer V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "mx_bank_transfer"
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeUSBankTransfer V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "us_bank_transfer"
)

// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceFundingType string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceFundingType can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceFundingTypeBankTransfer V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceFundingType = "bank_transfer"
)

// The account subcategories to use to filter for selectable accounts.
type V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategoryChecking V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "checking"
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategorySavings  V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "savings"
)

// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
type V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionBalances      V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "balances"
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionOwnership     V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "ownership"
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionPaymentMethod V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "payment_method"
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionTransactions  V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "transactions"
)

// List of data features that you would like to retrieve upon account creation.
type V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchBalances     V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "balances"
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchOwnership    V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "ownership"
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchTransactions V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "transactions"
)

// Verification method.
type V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod string

// List of values that V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod can take
const (
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethodAutomatic     V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "automatic"
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethodInstant       V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "instant"
	V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethodMicrodeposits V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "microdeposits"
)

// Determines if tax is calculated automatically based on a product tax code or manually based on rules defined by the business.
type V2BillingGroupBillingSettingsTaxType string

// List of values that V2BillingGroupBillingSettingsTaxType can take
const (
	V2BillingGroupBillingSettingsTaxTypeAutomatic V2BillingGroupBillingSettingsTaxType = "automatic"
	V2BillingGroupBillingSettingsTaxTypeManual    V2BillingGroupBillingSettingsTaxType = "manual"
)

// The interval unit for the cadence.
type V2BillingGroupGroupingBehaviorByCadenceType string

// List of values that V2BillingGroupGroupingBehaviorByCadenceType can take
const (
	V2BillingGroupGroupingBehaviorByCadenceTypeDay   V2BillingGroupGroupingBehaviorByCadenceType = "day"
	V2BillingGroupGroupingBehaviorByCadenceTypeMonth V2BillingGroupGroupingBehaviorByCadenceType = "month"
	V2BillingGroupGroupingBehaviorByCadenceTypeWeek  V2BillingGroupGroupingBehaviorByCadenceType = "week"
	V2BillingGroupGroupingBehaviorByCadenceTypeYear  V2BillingGroupGroupingBehaviorByCadenceType = "year"
)

// The grouping strategy.
type V2BillingGroupGroupingBehaviorType string

// List of values that V2BillingGroupGroupingBehaviorType can take
const (
	V2BillingGroupGroupingBehaviorTypeByCadence V2BillingGroupGroupingBehaviorType = "by_cadence"
)

// Max applications type of this discount, ex: indefinite.
type V2BillingGroupInvoiceDiscountRulePercentOffMaximumApplicationsType string

// List of values that V2BillingGroupInvoiceDiscountRulePercentOffMaximumApplicationsType can take
const (
	V2BillingGroupInvoiceDiscountRulePercentOffMaximumApplicationsTypeIndefinite V2BillingGroupInvoiceDiscountRulePercentOffMaximumApplicationsType = "indefinite"
	V2BillingGroupInvoiceDiscountRulePercentOffMaximumApplicationsTypeTimes      V2BillingGroupInvoiceDiscountRulePercentOffMaximumApplicationsType = "times"
)

// The type of the discount.
type V2BillingGroupInvoiceDiscountRuleType string

// List of values that V2BillingGroupInvoiceDiscountRuleType can take
const (
	V2BillingGroupInvoiceDiscountRuleTypePercentOff V2BillingGroupInvoiceDiscountRuleType = "percent_off"
)

// The current status of the billing group.
type V2BillingGroupStatus string

// List of values that V2BillingGroupStatus can take
const (
	V2BillingGroupStatusActive   V2BillingGroupStatus = "active"
	V2BillingGroupStatusCanceled V2BillingGroupStatus = "canceled"
)

// Controls emails for when the payment is due.
type V2BillingGroupBillingSettingsEmailDeliveryPaymentDue struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled bool `json:"enabled"`
	// If true the payment link to hosted invoice page would be included in email and PDF of the invoice.
	IncludePaymentLink bool `json:"include_payment_link"`
}

// Email delivery setting.
type V2BillingGroupBillingSettingsEmailDelivery struct {
	// Controls emails for when the payment is due.
	PaymentDue *V2BillingGroupBillingSettingsEmailDeliveryPaymentDue `json:"payment_due"`
}

// Settings for invoice presentation.
type V2BillingGroupBillingSettingsInvoicePresentation struct {
	// The ID of the invoice rendering template to be used when generating invoices.
	InvoiceRenderingTemplate string `json:"invoice_rendering_template,omitempty"`
}

// The amount of time until the invoice is overdue for payment.
type V2BillingGroupBillingSettingsInvoiceTimeUntilDue struct {
	// The interval unit for the time until due.
	Interval V2BillingGroupBillingSettingsInvoiceTimeUntilDueInterval `json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30, the invoice is due in 30 days.
	IntervalCount int64 `json:"interval_count"`
}

// Additional fields for Mandate creation.
type V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptions struct {
	// Transaction type of the mandate.
	TransactionType V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebit struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingGroupBillingSettingsPaymentMethodOptionsBancontact struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactPreferredLanguage `json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptions struct {
	// Amount to be charged for future payments.
	Amount int64 `json:"amount,string,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptionsAmountType `json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description string `json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCard struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptions `json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network string `json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure V2BillingGroupBillingSettingsPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer struct {
	// The desired country code of the bank account information.
	Country V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry `json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransfer struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer `json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferType `json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalance struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceFundingType `json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsKonbini struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsSEPADebit struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory `json:"account_subcategories"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters `json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch `json:"prefetch"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccount struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	// Verification method.
	VerificationMethod V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment Method specific configuration to be stored on the object.
type V2BillingGroupBillingSettingsPaymentMethodOptions struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebit `json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingGroupBillingSettingsPaymentMethodOptionsBancontact `json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingGroupBillingSettingsPaymentMethodOptionsCard `json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalance `json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingGroupBillingSettingsPaymentMethodOptionsKonbini `json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingGroupBillingSettingsPaymentMethodOptionsSEPADebit `json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccount `json:"us_bank_account,omitempty"`
}

// Settings for calculating tax.
type V2BillingGroupBillingSettingsTax struct {
	// Determines if tax is calculated automatically based on a product tax code or manually based on rules defined by the business.
	Type V2BillingGroupBillingSettingsTaxType `json:"type"`
}

// The effective billing settings for the group. Returned only when `billing_settings` is included in the request.
type V2BillingGroupBillingSettings struct {
	// How Stripe collects payment.
	CollectionMethod V2BillingGroupBillingSettingsCollectionMethod `json:"collection_method"`
	// Email delivery setting.
	EmailDelivery *V2BillingGroupBillingSettingsEmailDelivery `json:"email_delivery"`
	// Settings for invoice presentation.
	InvoicePresentation *V2BillingGroupBillingSettingsInvoicePresentation `json:"invoice_presentation,omitempty"`
	// The amount of time until the invoice is overdue for payment.
	InvoiceTimeUntilDue *V2BillingGroupBillingSettingsInvoiceTimeUntilDue `json:"invoice_time_until_due"`
	// The ID of the payment method configuration, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration string `json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration to be stored on the object.
	PaymentMethodOptions *V2BillingGroupBillingSettingsPaymentMethodOptions `json:"payment_method_options,omitempty"`
	// Settings for calculating tax.
	Tax *V2BillingGroupBillingSettingsTax `json:"tax"`
}

// The time at which the daily grouping interval begins.
type V2BillingGroupGroupingBehaviorByCadenceDayTime struct {
	// The hour, from 0 through 23.
	Hour int64 `json:"hour"`
	// The minute, from 0 through 59.
	Minute int64 `json:"minute"`
	// The second, from 0 through 59. Defaults to 0.
	Second int64 `json:"second,omitempty"`
}

// Daily cadence details.
type V2BillingGroupGroupingBehaviorByCadenceDay struct {
	// The time at which the daily grouping interval begins.
	Time *V2BillingGroupGroupingBehaviorByCadenceDayTime `json:"time"`
}

// The time at which the monthly grouping interval begins. Defaults to the time the billing group was created, in UTC.
type V2BillingGroupGroupingBehaviorByCadenceMonthTime struct {
	// The hour, from 0 through 23.
	Hour int64 `json:"hour"`
	// The minute, from 0 through 59.
	Minute int64 `json:"minute"`
	// The second, from 0 through 59. Defaults to 0.
	Second int64 `json:"second,omitempty"`
}

// Monthly cadence details.
type V2BillingGroupGroupingBehaviorByCadenceMonth struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth int64 `json:"day_of_month,omitempty"`
	// The month of year, from 1 through 12, from which occurrences are calculated.
	MonthOfYear int64 `json:"month_of_year,omitempty"`
	// The time at which the monthly grouping interval begins. Defaults to the time the billing group was created, in UTC.
	Time *V2BillingGroupGroupingBehaviorByCadenceMonthTime `json:"time"`
}

// The time at which the weekly grouping interval begins.
type V2BillingGroupGroupingBehaviorByCadenceWeekTime struct {
	// The hour, from 0 through 23.
	Hour int64 `json:"hour"`
	// The minute, from 0 through 59.
	Minute int64 `json:"minute"`
	// The second, from 0 through 59. Defaults to 0.
	Second int64 `json:"second,omitempty"`
}

// Weekly cadence details.
type V2BillingGroupGroupingBehaviorByCadenceWeek struct {
	// The ISO-8601 day of week, from 1 for Monday through 7 for Sunday.
	DayOfWeek int64 `json:"day_of_week,omitempty"`
	// The time at which the weekly grouping interval begins.
	Time *V2BillingGroupGroupingBehaviorByCadenceWeekTime `json:"time"`
}

// The time at which the yearly grouping interval begins. Defaults to the time the billing group was created, in UTC.
type V2BillingGroupGroupingBehaviorByCadenceYearTime struct {
	// The hour, from 0 through 23.
	Hour int64 `json:"hour"`
	// The minute, from 0 through 59.
	Minute int64 `json:"minute"`
	// The second, from 0 through 59. Defaults to 0.
	Second int64 `json:"second,omitempty"`
}

// Yearly cadence details.
type V2BillingGroupGroupingBehaviorByCadenceYear struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth int64 `json:"day_of_month,omitempty"`
	// The month of year, from 1 through 12.
	MonthOfYear int64 `json:"month_of_year,omitempty"`
	// The time at which the yearly grouping interval begins. Defaults to the time the billing group was created, in UTC.
	Time *V2BillingGroupGroupingBehaviorByCadenceYearTime `json:"time"`
}

// The cadence configuration when type is `by_cadence`.
type V2BillingGroupGroupingBehaviorByCadence struct {
	// Timestamp that anchors the recurring grouping cadence.
	BillingCycleAnchor time.Time `json:"billing_cycle_anchor,omitempty"`
	// Daily cadence details.
	Day *V2BillingGroupGroupingBehaviorByCadenceDay `json:"day,omitempty"`
	// The number of interval units between groups.
	IntervalCount int64 `json:"interval_count"`
	// Monthly cadence details.
	Month *V2BillingGroupGroupingBehaviorByCadenceMonth `json:"month,omitempty"`
	// The interval unit for the cadence.
	Type V2BillingGroupGroupingBehaviorByCadenceType `json:"type"`
	// Weekly cadence details.
	Week *V2BillingGroupGroupingBehaviorByCadenceWeek `json:"week,omitempty"`
	// Yearly cadence details.
	Year *V2BillingGroupGroupingBehaviorByCadenceYear `json:"year,omitempty"`
}

// The strategy used to group member invoices.
type V2BillingGroupGroupingBehavior struct {
	// The cadence configuration when type is `by_cadence`.
	ByCadence *V2BillingGroupGroupingBehaviorByCadence `json:"by_cadence,omitempty"`
	// The grouping strategy.
	Type V2BillingGroupGroupingBehaviorType `json:"type"`
}

// The maximum applications configuration for this discount.
type V2BillingGroupInvoiceDiscountRulePercentOffMaximumApplications struct {
	// The number of times this discount can be applied when the type is `times`.
	Times int64 `json:"times,omitempty"`
	// Max applications type of this discount, ex: indefinite.
	Type V2BillingGroupInvoiceDiscountRulePercentOffMaximumApplicationsType `json:"type"`
}

// Details if the discount is a percentage off.
type V2BillingGroupInvoiceDiscountRulePercentOff struct {
	// The maximum applications configuration for this discount.
	MaximumApplications *V2BillingGroupInvoiceDiscountRulePercentOffMaximumApplications `json:"maximum_applications"`
	// Percent that is taken off the amount. For example, a percent_off of 50.0 reduces a 100 USD amount to 50 USD.
	PercentOff decimal.Decimal `json:"percent_off"`
}

// The discount rules applied to all invoices for the group's underlying cadence. Returned only when `invoice_discount_rules` is included in the request.
type V2BillingGroupInvoiceDiscountRule struct {
	// Unique identifier for the object.
	ID string `json:"id"`
	// Details if the discount is a percentage off.
	PercentOff *V2BillingGroupInvoiceDiscountRulePercentOff `json:"percent_off,omitempty"`
	// The type of the discount.
	Type V2BillingGroupInvoiceDiscountRuleType `json:"type"`
}

// Timestamps for the billing group's status transitions.
type V2BillingGroupStatusTransitions struct {
	// Timestamp of when the billing group became active.
	ActivatedAt time.Time `json:"activated_at,omitempty"`
	// Timestamp of when the billing group was canceled.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
}

// A Billing Group combines member invoices into a consolidated invoice according to shared grouping behavior.
type V2BillingGroup struct {
	APIResource
	// The effective billing settings for the group. Returned only when `billing_settings` is included in the request.
	BillingSettings *V2BillingGroupBillingSettings `json:"billing_settings,omitempty"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// The ID of the Customer object whose invoices are grouped.
	Customer string `json:"customer"`
	// The ID of the PaymentMethod object used by default. If unset, the customer's default payment method is used.
	DefaultPaymentMethod string `json:"default_payment_method,omitempty"`
	// The strategy used to group member invoices.
	GroupingBehavior *V2BillingGroupGroupingBehavior `json:"grouping_behavior"`
	// The ID of the BillingGroup object.
	ID string `json:"id"`
	// The discount rules applied to all invoices for the group's underlying cadence. Returned only when `invoice_discount_rules` is included in the request.
	InvoiceDiscountRules []*V2BillingGroupInvoiceDiscountRule `json:"invoice_discount_rules,omitempty"`
	// Whether the billing group exists in live mode.
	Livemode bool `json:"livemode"`
	// An internal key you can use to find a particular billing group. Maximum length of 200 characters.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object for storing additional structured information.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The date when the billing group will next create a consolidated invoice.
	NextBillingDate time.Time `json:"next_billing_date"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The current status of the billing group.
	Status V2BillingGroupStatus `json:"status"`
	// Timestamps for the billing group's status transitions.
	StatusTransitions *V2BillingGroupStatusTransitions `json:"status_transitions"`
	// The ID of the TestClock object associated with the billing group.
	TestClock string `json:"test_clock,omitempty"`
}
