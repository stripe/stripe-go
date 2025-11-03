//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The frequency at which a cadence bills.
type V2BillingCadenceBillingCycleType string

// List of values that V2BillingCadenceBillingCycleType can take
const (
	V2BillingCadenceBillingCycleTypeDay   V2BillingCadenceBillingCycleType = "day"
	V2BillingCadenceBillingCycleTypeMonth V2BillingCadenceBillingCycleType = "month"
	V2BillingCadenceBillingCycleTypeWeek  V2BillingCadenceBillingCycleType = "week"
	V2BillingCadenceBillingCycleTypeYear  V2BillingCadenceBillingCycleType = "year"
)

// A string identifying the type of the payer. Currently the only supported value is `customer`.
type V2BillingCadencePayerType string

// List of values that V2BillingCadencePayerType can take
const (
	V2BillingCadencePayerTypeCustomer V2BillingCadencePayerType = "customer"
)

// Determines if tax will be calculated automatically based on a PTC or manually based on rules defined by the merchant. Defaults to "manual".
type V2BillingCadenceSettingsDataBillCalculationTaxType string

// List of values that V2BillingCadenceSettingsDataBillCalculationTaxType can take
const (
	V2BillingCadenceSettingsDataBillCalculationTaxTypeAutomatic V2BillingCadenceSettingsDataBillCalculationTaxType = "automatic"
	V2BillingCadenceSettingsDataBillCalculationTaxTypeManual    V2BillingCadenceSettingsDataBillCalculationTaxType = "manual"
)

// The interval unit for the time until due.
type V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueInterval string

// List of values that V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueInterval can take
const (
	V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueIntervalDay   V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueInterval = "day"
	V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueIntervalMonth V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueInterval = "month"
	V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueIntervalWeek  V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueInterval = "week"
	V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueIntervalYear  V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueInterval = "year"
)

// Either automatic, or send_invoice. When charging automatically, Stripe will attempt to pay this
// bill at the end of the period using the payment method attached to the payer profile. When sending an invoice,
// Stripe will email your payer profile an invoice with payment instructions.
// Defaults to automatic.
type V2BillingCadenceSettingsDataCollectionCollectionMethod string

// List of values that V2BillingCadenceSettingsDataCollectionCollectionMethod can take
const (
	V2BillingCadenceSettingsDataCollectionCollectionMethodAutomatic   V2BillingCadenceSettingsDataCollectionCollectionMethod = "automatic"
	V2BillingCadenceSettingsDataCollectionCollectionMethodSendInvoice V2BillingCadenceSettingsDataCollectionCollectionMethod = "send_invoice"
)

// Transaction type of the mandate.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Verification method.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitVerificationMethodInstant       V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// Preferred language of the Bancontact authorization page that the customer is redirected to.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguage string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguage can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguageDE V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguage = "de"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguageEN V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguage = "en"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguageFR V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguage = "fr"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguageNL V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguage = "nl"
)

// The AmountType for the mandate. One of `fixed` or `maximum`.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptionsAmountType string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptionsAmountType can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptionsAmountTypeFixed   V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptionsAmountType = "fixed"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptionsAmountType = "maximum"
)

// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
// for authentication based on risk level and [other requirements](https://docs.corp.stripe.com/strong-customer-authentication).
// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
// Read our guide on [manually requesting 3D Secure](https://docs.corp.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardRequestThreeDSecureAny       V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardRequestThreeDSecure = "any"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardRequestThreeDSecureAutomatic V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardRequestThreeDSecureChallenge V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardRequestThreeDSecure = "challenge"
)

// The desired country code of the bank account information.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryBE V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "BE"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryDE V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "DE"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryES V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "ES"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryFR V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "FR"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryIE V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "IE"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryNL V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "NL"
)

// The bank transfer type that can be used for funding.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferType string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferType can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferTypeEUBankTransfer V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferType = "eu_bank_transfer"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferTypeGBBankTransfer V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferType = "gb_bank_transfer"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferTypeJPBankTransfer V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferType = "jp_bank_transfer"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferTypeMXBankTransfer V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferType = "mx_bank_transfer"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferTypeUSBankTransfer V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferType = "us_bank_transfer"
)

// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceFundingType string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceFundingType can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceFundingTypeBankTransfer V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceFundingType = "bank_transfer"
)

// The account subcategories to use to filter for selectable accounts.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategoryChecking V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "checking"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategorySavings  V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "savings"
)

// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionBalances      V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "balances"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionOwnership     V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "ownership"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionPaymentMethod V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "payment_method"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionTransactions  V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "transactions"
)

// List of data features that you would like to retrieve upon account creation.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchBalances     V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "balances"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchOwnership    V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "ownership"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchTransactions V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "transactions"
)

// Verification method.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountVerificationMethod string

// List of values that V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountVerificationMethod can take
const (
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountVerificationMethodAutomatic     V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountVerificationMethod = "automatic"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountVerificationMethodInstant       V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountVerificationMethod = "instant"
	V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountVerificationMethodMicrodeposits V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountVerificationMethod = "microdeposits"
)

// The current status of the cadence.
type V2BillingCadenceStatus string

// List of values that V2BillingCadenceStatus can take
const (
	V2BillingCadenceStatusActive   V2BillingCadenceStatus = "active"
	V2BillingCadenceStatusCanceled V2BillingCadenceStatus = "canceled"
)

// The time at which the billing cycle ends.
type V2BillingCadenceBillingCycleDayTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second,omitempty"`
}

// Specific configuration for determining billing dates when type=day.
type V2BillingCadenceBillingCycleDay struct {
	// The time at which the billing cycle ends.
	Time *V2BillingCadenceBillingCycleDayTime `json:"time"`
}

// The time at which the billing cycle ends.
type V2BillingCadenceBillingCycleMonthTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second,omitempty"`
}

// Specific configuration for determining billing dates when type=month.
type V2BillingCadenceBillingCycleMonth struct {
	// The day to anchor the billing on for a type="month" billing cycle from 1-31.
	// If this number is greater than the number of days in the month being billed,
	// this will anchor to the last day of the month.
	DayOfMonth int64 `json:"day_of_month"`
	// The month to anchor the billing on for a type="month" billing cycle from
	// 1-12. Occurrences are calculated from the month anchor.
	MonthOfYear int64 `json:"month_of_year,omitempty"`
	// The time at which the billing cycle ends.
	Time *V2BillingCadenceBillingCycleMonthTime `json:"time"`
}

// The time at which the billing cycle ends.
type V2BillingCadenceBillingCycleWeekTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second,omitempty"`
}

// Specific configuration for determining billing dates when type=week.
type V2BillingCadenceBillingCycleWeek struct {
	// The day of the week to bill the type=week billing cycle on.
	// Numbered from 1-7 for Monday to Sunday respectively, based on the ISO-8601 week day numbering.
	DayOfWeek int64 `json:"day_of_week"`
	// The time at which the billing cycle ends.
	Time *V2BillingCadenceBillingCycleWeekTime `json:"time"`
}

// The time at which the billing cycle ends.
type V2BillingCadenceBillingCycleYearTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second,omitempty"`
}

// Specific configuration for determining billing dates when type=year.
type V2BillingCadenceBillingCycleYear struct {
	// The day to anchor the billing on for a type="month" billing cycle from 1-31.
	// If this number is greater than the number of days in the month being billed,
	// this will anchor to the last day of the month.
	DayOfMonth int64 `json:"day_of_month"`
	// The month to bill on from 1-12. If not provided, this will default to the month the cadence was created.
	MonthOfYear int64 `json:"month_of_year"`
	// The time at which the billing cycle ends.
	Time *V2BillingCadenceBillingCycleYearTime `json:"time"`
}

// The billing cycle is the object that defines future billing cycle dates.
type V2BillingCadenceBillingCycle struct {
	// Specific configuration for determining billing dates when type=day.
	Day *V2BillingCadenceBillingCycleDay `json:"day,omitempty"`
	// The number of intervals (specified in the interval attribute) between cadence billings. For example, type=month and interval_count=3 bills every 3 months.
	IntervalCount int64 `json:"interval_count"`
	// Specific configuration for determining billing dates when type=month.
	Month *V2BillingCadenceBillingCycleMonth `json:"month,omitempty"`
	// The frequency at which a cadence bills.
	Type V2BillingCadenceBillingCycleType `json:"type"`
	// Specific configuration for determining billing dates when type=week.
	Week *V2BillingCadenceBillingCycleWeek `json:"week,omitempty"`
	// Specific configuration for determining billing dates when type=year.
	Year *V2BillingCadenceBillingCycleYear `json:"year,omitempty"`
}

// The payer determines the entity financially responsible for the bill.
type V2BillingCadencePayer struct {
	// The ID of the Billing Profile object which determines how a bill will be paid.
	BillingProfile string `json:"billing_profile"`
	// The ID of the Customer object.
	Customer string `json:"customer,omitempty"`
	// A string identifying the type of the payer. Currently the only supported value is `customer`.
	Type V2BillingCadencePayerType `json:"type"`
}

// Settings that configure bills generation, which includes calculating totals, tax, and presenting invoices.
type V2BillingCadenceSettingsBill struct {
	// The ID of the referenced settings object.
	ID string `json:"id"`
	// Returns the Settings Version when the cadence is pinned to a specific version.
	Version string `json:"version,omitempty"`
}

// Settings that configure and manage the behavior of collecting payments.
type V2BillingCadenceSettingsCollection struct {
	// The ID of the referenced settings object.
	ID string `json:"id"`
	// Returns the Settings Version when the cadence is pinned to a specific version.
	Version string `json:"version,omitempty"`
}

// The settings associated with the cadence.
type V2BillingCadenceSettings struct {
	// Settings that configure bills generation, which includes calculating totals, tax, and presenting invoices.
	Bill *V2BillingCadenceSettingsBill `json:"bill,omitempty"`
	// Settings that configure and manage the behavior of collecting payments.
	Collection *V2BillingCadenceSettingsCollection `json:"collection,omitempty"`
}

// Settings for calculating tax.
type V2BillingCadenceSettingsDataBillCalculationTax struct {
	// Determines if tax will be calculated automatically based on a PTC or manually based on rules defined by the merchant. Defaults to "manual".
	Type V2BillingCadenceSettingsDataBillCalculationTaxType `json:"type"`
}

// Settings related to calculating a bill.
type V2BillingCadenceSettingsDataBillCalculation struct {
	// Settings for calculating tax.
	Tax *V2BillingCadenceSettingsDataBillCalculationTax `json:"tax,omitempty"`
}

// The amount of time until the invoice will be overdue for payment.
type V2BillingCadenceSettingsDataBillInvoiceTimeUntilDue struct {
	// The interval unit for the time until due.
	Interval V2BillingCadenceSettingsDataBillInvoiceTimeUntilDueInterval `json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30,
	// the invoice will be due in 30 days.
	IntervalCount int64 `json:"interval_count"`
}

// Settings related to invoice behavior.
type V2BillingCadenceSettingsDataBillInvoice struct {
	// The amount of time until the invoice will be overdue for payment.
	TimeUntilDue *V2BillingCadenceSettingsDataBillInvoiceTimeUntilDue `json:"time_until_due,omitempty"`
}

// Expanded bill settings data with actual configuration values.
type V2BillingCadenceSettingsDataBill struct {
	// Settings related to calculating a bill.
	Calculation *V2BillingCadenceSettingsDataBillCalculation `json:"calculation"`
	// Settings related to invoice behavior.
	Invoice *V2BillingCadenceSettingsDataBillInvoice `json:"invoice"`
	// The ID of the invoice rendering template to be used when generating invoices.
	InvoiceRenderingTemplate string `json:"invoice_rendering_template,omitempty"`
}

// Controls emails for when the payment is due. For example after the invoice is finalized and transitions to Open state.
type V2BillingCadenceSettingsDataCollectionEmailDeliveryPaymentDue struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled bool `json:"enabled"`
	// If true the payment link to hosted invoice page would be included in email and PDF of the invoice.
	IncludePaymentLink bool `json:"include_payment_link"`
}

// Email delivery settings.
type V2BillingCadenceSettingsDataCollectionEmailDelivery struct {
	// Controls emails for when the payment is due. For example after the invoice is finalized and transitions to Open state.
	PaymentDue *V2BillingCadenceSettingsDataCollectionEmailDeliveryPaymentDue `json:"payment_due,omitempty"`
}

// Additional fields for Mandate creation.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitMandateOptions struct {
	// Transaction type of the mandate.
	TransactionType V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebit struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontact struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontactPreferredLanguage `json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptions struct {
	// Amount to be charged for future payments.
	Amount int64 `json:"amount,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptionsAmountType `json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description string `json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCard struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptions `json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network string `json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.corp.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.corp.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer struct {
	// The desired country code of the bank account information.
	Country V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry `json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransfer struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer `json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransferType `json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalance struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalanceFundingType `json:"funding_type,omitempty"`
}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory `json:"account_subcategories"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters `json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch `json:"prefetch"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccount struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	// Verification method.
	VerificationMethod V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment Method specific configuration stored on the object.
type V2BillingCadenceSettingsDataCollectionPaymentMethodOptions struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsACSSDebit `json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsBancontact `json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCard `json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCustomerBalance `json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini map[string]any `json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit map[string]any `json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsUSBankAccount `json:"us_bank_account,omitempty"`
}

// Expanded collection settings data with actual configuration values.
type V2BillingCadenceSettingsDataCollection struct {
	// Either automatic, or send_invoice. When charging automatically, Stripe will attempt to pay this
	// bill at the end of the period using the payment method attached to the payer profile. When sending an invoice,
	// Stripe will email your payer profile an invoice with payment instructions.
	// Defaults to automatic.
	CollectionMethod V2BillingCadenceSettingsDataCollectionCollectionMethod `json:"collection_method"`
	// Email delivery settings.
	EmailDelivery *V2BillingCadenceSettingsDataCollectionEmailDelivery `json:"email_delivery"`
	// The ID of the PaymentMethodConfiguration object, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration string `json:"payment_method_configuration"`
	// Payment Method specific configuration stored on the object.
	PaymentMethodOptions *V2BillingCadenceSettingsDataCollectionPaymentMethodOptions `json:"payment_method_options"`
}

// Settings data that contains expanded billing settings configuration with actual values.
type V2BillingCadenceSettingsData struct {
	// Expanded bill settings data with actual configuration values.
	Bill *V2BillingCadenceSettingsDataBill `json:"bill"`
	// Expanded collection settings data with actual configuration values.
	Collection *V2BillingCadenceSettingsDataCollection `json:"collection"`
}
type V2BillingCadence struct {
	APIResource
	// The billing cycle is the object that defines future billing cycle dates.
	BillingCycle *V2BillingCadenceBillingCycle `json:"billing_cycle"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// A lookup key used to retrieve cadences dynamically from a static string. Maximum length of 200 characters.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The date that the billing cadence will next bill. Null if the cadence is not active.
	NextBillingDate time.Time `json:"next_billing_date,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The payer determines the entity financially responsible for the bill.
	Payer *V2BillingCadencePayer `json:"payer"`
	// The settings associated with the cadence.
	Settings *V2BillingCadenceSettings `json:"settings,omitempty"`
	// Settings data that contains expanded billing settings configuration with actual values.
	SettingsData *V2BillingCadenceSettingsData `json:"settings_data,omitempty"`
	// The current status of the cadence.
	Status V2BillingCadenceStatus `json:"status"`
	// The ID of the Test Clock.
	TestClock string `json:"test_clock,omitempty"`
}
