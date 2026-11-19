//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List Billing Groups.
type V2BillingGroupListParams struct {
	Params `form:"*"`
	// Filter billing groups by Customer ID. Mutually exclusive with test_clock.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The maximum number of results per page. Defaults to 10.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter billing groups by lookup keys. You can specify up to 10 lookup keys.
	LookupKeys []*string `form:"lookup_keys" json:"lookup_keys,omitempty"`
	// Filter billing groups by TestClock ID. Mutually exclusive with customer.
	TestClock *string `form:"test_clock" json:"test_clock,omitempty"`
}

// Controls emails for when the payment is due. Omit to use the applicable default.
type V2BillingGroupBillingSettingsEmailDeliveryPaymentDueParams struct {
	// If true an email for the invoice would be generated and sent out. Omit to leave unchanged. Set to null to restore the applicable default.
	Enabled *bool `form:"enabled" json:"enabled,omitempty"`
	// If true the payment link to hosted invoice page would be included in email and PDF of the invoice. Omit to leave unchanged. Set to null to restore the applicable default.
	IncludePaymentLink *bool `form:"include_payment_link" json:"include_payment_link,omitempty"`
}

// Email delivery setting.
type V2BillingGroupBillingSettingsEmailDeliveryParams struct {
	// Controls emails for when the payment is due. Omit to leave unchanged. Set to null to restore the applicable default.
	PaymentDue *V2BillingGroupBillingSettingsEmailDeliveryPaymentDueParams `form:"payment_due" json:"payment_due,omitempty"`
}

// Settings for invoice presentation.
type V2BillingGroupBillingSettingsInvoicePresentationParams struct {
	// The ID of the invoice rendering template to be used when generating invoices. Omit to use the applicable default.
	InvoiceRenderingTemplate *string `form:"invoice_rendering_template" json:"invoice_rendering_template,omitempty"`
}

// The amount of time until the invoice is overdue for payment.
type V2BillingGroupBillingSettingsInvoiceTimeUntilDueParams struct {
	// The interval unit for the time until due.
	Interval *string `form:"interval" json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30, the invoice is due in 30 days.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Additional fields for Mandate creation.
type V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type" json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitParams struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactParams struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage *string `form:"preferred_language" json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount" json:"amount,string,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType *string `form:"amount_type" json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description *string `form:"description" json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingGroupBillingSettingsPaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network *string `form:"network" json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure" json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information.
	Country *string `form:"country" json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer" json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type *string `form:"type" json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType *string `form:"funding_type" json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsKonbiniParams struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsSEPADebitParams struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []*string `form:"account_subcategories" json:"account_subcategories,omitempty"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters" json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []*string `form:"permissions" json:"permissions,omitempty"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch" json:"prefetch,omitempty"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections" json:"financial_connections"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method"`
}

// Payment Method specific configuration to be stored on the object.
type V2BillingGroupBillingSettingsPaymentMethodOptionsParams struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingGroupBillingSettingsPaymentMethodOptionsACSSDebitParams `form:"acss_debit" json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingGroupBillingSettingsPaymentMethodOptionsBancontactParams `form:"bancontact" json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingGroupBillingSettingsPaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingGroupBillingSettingsPaymentMethodOptionsCustomerBalanceParams `form:"customer_balance" json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingGroupBillingSettingsPaymentMethodOptionsKonbiniParams `form:"konbini" json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingGroupBillingSettingsPaymentMethodOptionsSEPADebitParams `form:"sepa_debit" json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingGroupBillingSettingsPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// Settings for calculating tax.
type V2BillingGroupBillingSettingsTaxParams struct {
	// Determines if tax is calculated automatically based on a product tax code or manually based on rules defined by the business. Omit to use the applicable default.
	Type *string `form:"type" json:"type,omitempty"`
}

// Billing setting overrides for the group. If omitted, the applicable defaults are used.
type V2BillingGroupBillingSettingsParams struct {
	// How Stripe collects payment. Set to null to restore the applicable default.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// Email delivery setting. Set to null to restore the applicable default.
	EmailDelivery *V2BillingGroupBillingSettingsEmailDeliveryParams `form:"email_delivery" json:"email_delivery,omitempty"`
	// Settings for invoice presentation. Set to null to restore the applicable defaults.
	InvoicePresentation *V2BillingGroupBillingSettingsInvoicePresentationParams `form:"invoice_presentation" json:"invoice_presentation,omitempty"`
	// The amount of time until the invoice is overdue for payment. Set to null to restore the applicable default.
	InvoiceTimeUntilDue *V2BillingGroupBillingSettingsInvoiceTimeUntilDueParams `form:"invoice_time_until_due" json:"invoice_time_until_due,omitempty"`
	// The ID of the payment method configuration, which controls which payment methods are displayed to your customers. Set to null to restore the applicable default.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration to be stored on the object. Set to null to clear the stored value.
	PaymentMethodOptions *V2BillingGroupBillingSettingsPaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// Tax calculation settings. Set to null to restore the applicable default.
	Tax *V2BillingGroupBillingSettingsTaxParams `form:"tax" json:"tax,omitempty"`
}

// The time at which the daily grouping interval begins.
type V2BillingGroupGroupingBehaviorByCadenceDayTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Daily cadence details.
type V2BillingGroupGroupingBehaviorByCadenceDayParams struct {
	// The time at which the daily grouping interval begins.
	Time *V2BillingGroupGroupingBehaviorByCadenceDayTimeParams `form:"time" json:"time"`
}

// The time at which the monthly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingGroupGroupingBehaviorByCadenceMonthTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Monthly cadence details.
type V2BillingGroupGroupingBehaviorByCadenceMonthParams struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The month of year, from 1 through 12, from which occurrences are calculated. Defaults to the month when the billing group was created.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// The time at which the monthly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingGroupGroupingBehaviorByCadenceMonthTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the weekly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingGroupGroupingBehaviorByCadenceWeekTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Weekly cadence details.
type V2BillingGroupGroupingBehaviorByCadenceWeekParams struct {
	// The ISO-8601 day of week, from 1 for Monday through 7 for Sunday.
	DayOfWeek *int64 `form:"day_of_week" json:"day_of_week"`
	// The time at which the weekly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingGroupGroupingBehaviorByCadenceWeekTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the yearly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingGroupGroupingBehaviorByCadenceYearTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Yearly cadence details.
type V2BillingGroupGroupingBehaviorByCadenceYearParams struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The month of year, from 1 through 12.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year"`
	// The time at which the yearly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingGroupGroupingBehaviorByCadenceYearTimeParams `form:"time" json:"time,omitempty"`
}

// The cadence configuration when type is `by_cadence`.
type V2BillingGroupGroupingBehaviorByCadenceParams struct {
	// Timestamp that anchors the recurring grouping cadence.
	BillingCycleAnchor *time.Time `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Daily cadence details.
	Day *V2BillingGroupGroupingBehaviorByCadenceDayParams `form:"day" json:"day,omitempty"`
	// The number of interval units between groups. Defaults to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
	// Monthly cadence details.
	Month *V2BillingGroupGroupingBehaviorByCadenceMonthParams `form:"month" json:"month,omitempty"`
	// The interval unit for the cadence.
	Type *string `form:"type" json:"type"`
	// Weekly cadence details.
	Week *V2BillingGroupGroupingBehaviorByCadenceWeekParams `form:"week" json:"week,omitempty"`
	// Yearly cadence details.
	Year *V2BillingGroupGroupingBehaviorByCadenceYearParams `form:"year" json:"year,omitempty"`
}

// The strategy used to group member invoices.
type V2BillingGroupGroupingBehaviorParams struct {
	// The cadence configuration when type is `by_cadence`.
	ByCadence *V2BillingGroupGroupingBehaviorByCadenceParams `form:"by_cadence" json:"by_cadence,omitempty"`
	// The grouping strategy.
	Type *string `form:"type" json:"type"`
}

// Create a BillingGroup.
type V2BillingGroupParams struct {
	Params `form:"*"`
	// Billing setting overrides for the group. If omitted, the applicable defaults are used.
	BillingSettings *V2BillingGroupBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// The ID of the Customer object whose invoices are grouped.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The ID of the PaymentMethod object used by default. If unset, the customer's default payment method is used.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// The strategy used to group member invoices.
	GroupingBehavior *V2BillingGroupGroupingBehaviorParams `form:"grouping_behavior" json:"grouping_behavior,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// An internal key you can use to find a particular billing group. Set to null to clear it. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object for storing additional structured information.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingGroupParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Cancel a BillingGroup.
type V2BillingGroupCancelParams struct {
	Params `form:"*"`
}

// Controls emails for when the payment is due. Omit to use the applicable default.
type V2BillingGroupCreateBillingSettingsEmailDeliveryPaymentDueParams struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled *bool `form:"enabled" json:"enabled,omitempty"`
	// If true the payment link to hosted invoice page would be included in email and PDF of the invoice.
	IncludePaymentLink *bool `form:"include_payment_link" json:"include_payment_link,omitempty"`
}

// Email delivery setting.
type V2BillingGroupCreateBillingSettingsEmailDeliveryParams struct {
	// Controls emails for when the payment is due. Omit to use the applicable default.
	PaymentDue *V2BillingGroupCreateBillingSettingsEmailDeliveryPaymentDueParams `form:"payment_due" json:"payment_due,omitempty"`
}

// Settings for invoice presentation.
type V2BillingGroupCreateBillingSettingsInvoicePresentationParams struct {
	// The ID of the invoice rendering template to be used when generating invoices. Omit to use the applicable default.
	InvoiceRenderingTemplate *string `form:"invoice_rendering_template" json:"invoice_rendering_template,omitempty"`
}

// The amount of time until the invoice is overdue for payment.
type V2BillingGroupCreateBillingSettingsInvoiceTimeUntilDueParams struct {
	// The interval unit for the time until due.
	Interval *string `form:"interval" json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30, the invoice is due in 30 days.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Additional fields for Mandate creation.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type" json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsACSSDebitParams struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsBancontactParams struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage *string `form:"preferred_language" json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsCardMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount" json:"amount,string,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType *string `form:"amount_type" json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description *string `form:"description" json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network *string `form:"network" json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure" json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information.
	Country *string `form:"country" json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer" json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type *string `form:"type" json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType *string `form:"funding_type" json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsKonbiniParams struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsSEPADebitParams struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []*string `form:"account_subcategories" json:"account_subcategories,omitempty"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters" json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []*string `form:"permissions" json:"permissions,omitempty"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch" json:"prefetch,omitempty"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections" json:"financial_connections"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method"`
}

// Payment Method specific configuration to be stored on the object.
type V2BillingGroupCreateBillingSettingsPaymentMethodOptionsParams struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsACSSDebitParams `form:"acss_debit" json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsBancontactParams `form:"bancontact" json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsCustomerBalanceParams `form:"customer_balance" json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsKonbiniParams `form:"konbini" json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsSEPADebitParams `form:"sepa_debit" json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// Settings for calculating tax.
type V2BillingGroupCreateBillingSettingsTaxParams struct {
	// Determines if tax is calculated automatically based on a product tax code or manually based on rules defined by the business. Omit to use the applicable default.
	Type *string `form:"type" json:"type,omitempty"`
}

// Billing setting overrides for the group. If omitted, the applicable defaults are used.
type V2BillingGroupCreateBillingSettingsParams struct {
	// How Stripe collects payment.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// Email delivery setting.
	EmailDelivery *V2BillingGroupCreateBillingSettingsEmailDeliveryParams `form:"email_delivery" json:"email_delivery,omitempty"`
	// Settings for invoice presentation.
	InvoicePresentation *V2BillingGroupCreateBillingSettingsInvoicePresentationParams `form:"invoice_presentation" json:"invoice_presentation,omitempty"`
	// The amount of time until the invoice is overdue for payment.
	InvoiceTimeUntilDue *V2BillingGroupCreateBillingSettingsInvoiceTimeUntilDueParams `form:"invoice_time_until_due" json:"invoice_time_until_due,omitempty"`
	// The ID of the payment method configuration, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration to be stored on the object.
	PaymentMethodOptions *V2BillingGroupCreateBillingSettingsPaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// Settings for calculating tax.
	Tax *V2BillingGroupCreateBillingSettingsTaxParams `form:"tax" json:"tax,omitempty"`
}

// The time at which the daily grouping interval begins.
type V2BillingGroupCreateGroupingBehaviorByCadenceDayTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Daily cadence details.
type V2BillingGroupCreateGroupingBehaviorByCadenceDayParams struct {
	// The time at which the daily grouping interval begins.
	Time *V2BillingGroupCreateGroupingBehaviorByCadenceDayTimeParams `form:"time" json:"time"`
}

// The time at which the monthly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingGroupCreateGroupingBehaviorByCadenceMonthTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Monthly cadence details.
type V2BillingGroupCreateGroupingBehaviorByCadenceMonthParams struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The month of year, from 1 through 12, from which occurrences are calculated. Defaults to the month when the billing group was created.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// The time at which the monthly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingGroupCreateGroupingBehaviorByCadenceMonthTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the weekly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingGroupCreateGroupingBehaviorByCadenceWeekTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Weekly cadence details.
type V2BillingGroupCreateGroupingBehaviorByCadenceWeekParams struct {
	// The ISO-8601 day of week, from 1 for Monday through 7 for Sunday.
	DayOfWeek *int64 `form:"day_of_week" json:"day_of_week"`
	// The time at which the weekly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingGroupCreateGroupingBehaviorByCadenceWeekTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the yearly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingGroupCreateGroupingBehaviorByCadenceYearTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Yearly cadence details.
type V2BillingGroupCreateGroupingBehaviorByCadenceYearParams struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The month of year, from 1 through 12.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year"`
	// The time at which the yearly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingGroupCreateGroupingBehaviorByCadenceYearTimeParams `form:"time" json:"time,omitempty"`
}

// The cadence configuration when type is `by_cadence`.
type V2BillingGroupCreateGroupingBehaviorByCadenceParams struct {
	// Timestamp that anchors the recurring grouping cadence.
	BillingCycleAnchor *time.Time `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Daily cadence details.
	Day *V2BillingGroupCreateGroupingBehaviorByCadenceDayParams `form:"day" json:"day,omitempty"`
	// The number of interval units between groups. Defaults to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
	// Monthly cadence details.
	Month *V2BillingGroupCreateGroupingBehaviorByCadenceMonthParams `form:"month" json:"month,omitempty"`
	// The interval unit for the cadence.
	Type *string `form:"type" json:"type"`
	// Weekly cadence details.
	Week *V2BillingGroupCreateGroupingBehaviorByCadenceWeekParams `form:"week" json:"week,omitempty"`
	// Yearly cadence details.
	Year *V2BillingGroupCreateGroupingBehaviorByCadenceYearParams `form:"year" json:"year,omitempty"`
}

// The strategy used to group member invoices.
type V2BillingGroupCreateGroupingBehaviorParams struct {
	// The cadence configuration when type is `by_cadence`.
	ByCadence *V2BillingGroupCreateGroupingBehaviorByCadenceParams `form:"by_cadence" json:"by_cadence,omitempty"`
	// The grouping strategy.
	Type *string `form:"type" json:"type"`
}

// Create a BillingGroup.
type V2BillingGroupCreateParams struct {
	Params `form:"*"`
	// Billing setting overrides for the group. If omitted, the applicable defaults are used.
	BillingSettings *V2BillingGroupCreateBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// The ID of the Customer object whose invoices are grouped.
	Customer *string `form:"customer" json:"customer"`
	// The ID of the PaymentMethod object used by default. If unset, the customer's default payment method is used.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// The strategy used to group member invoices.
	GroupingBehavior *V2BillingGroupCreateGroupingBehaviorParams `form:"grouping_behavior" json:"grouping_behavior,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// An internal key you can use to find a particular billing group. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object for storing additional structured information.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingGroupCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a BillingGroup.
type V2BillingGroupRetrieveParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// Controls emails for when the payment is due. Omit to leave unchanged. Set to null to restore the applicable default.
type V2BillingGroupUpdateBillingSettingsEmailDeliveryPaymentDueParams struct {
	// If true an email for the invoice would be generated and sent out. Omit to leave unchanged. Set to null to restore the applicable default.
	Enabled *bool `form:"enabled" json:"enabled,omitempty"`
	// If true the payment link to hosted invoice page would be included in email and PDF of the invoice. Omit to leave unchanged. Set to null to restore the applicable default.
	IncludePaymentLink *bool `form:"include_payment_link" json:"include_payment_link,omitempty"`
}

// Email delivery setting. Set to null to restore the applicable default.
type V2BillingGroupUpdateBillingSettingsEmailDeliveryParams struct {
	// Controls emails for when the payment is due. Omit to leave unchanged. Set to null to restore the applicable default.
	PaymentDue *V2BillingGroupUpdateBillingSettingsEmailDeliveryPaymentDueParams `form:"payment_due" json:"payment_due,omitempty"`
}

// Settings for invoice presentation. Set to null to restore the applicable defaults.
type V2BillingGroupUpdateBillingSettingsInvoicePresentationParams struct {
	// The ID of the invoice rendering template to be used when generating invoices. Set to null to use no template.
	InvoiceRenderingTemplate *string `form:"invoice_rendering_template" json:"invoice_rendering_template,omitempty"`
}

// The amount of time until the invoice is overdue for payment. Set to null to restore the applicable default.
type V2BillingGroupUpdateBillingSettingsInvoiceTimeUntilDueParams struct {
	// The interval unit for the time until due.
	Interval *string `form:"interval" json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30, the invoice is due in 30 days.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Additional fields for Mandate creation.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type" json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsACSSDebitParams struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsBancontactParams struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage *string `form:"preferred_language" json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsCardMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount" json:"amount,string,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType *string `form:"amount_type" json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description *string `form:"description" json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network *string `form:"network" json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure" json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information.
	Country *string `form:"country" json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer" json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type *string `form:"type" json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType *string `form:"funding_type" json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsKonbiniParams struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsSEPADebitParams struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []*string `form:"account_subcategories" json:"account_subcategories,omitempty"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters" json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []*string `form:"permissions" json:"permissions,omitempty"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch" json:"prefetch,omitempty"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections" json:"financial_connections"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method"`
}

// Payment Method specific configuration to be stored on the object. Set to null to clear the stored value.
type V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsParams struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsACSSDebitParams `form:"acss_debit" json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsBancontactParams `form:"bancontact" json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsCustomerBalanceParams `form:"customer_balance" json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsKonbiniParams `form:"konbini" json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsSEPADebitParams `form:"sepa_debit" json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// Tax calculation settings. Set to null to restore the applicable default.
type V2BillingGroupUpdateBillingSettingsTaxParams struct {
	// Determines if tax is calculated automatically based on a product tax code or manually based on rules defined by the business.
	Type *string `form:"type" json:"type"`
}

// Billing setting overrides to update. Set to null to restore all applicable defaults.
type V2BillingGroupUpdateBillingSettingsParams struct {
	// How Stripe collects payment. Set to null to restore the applicable default.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// Email delivery setting. Set to null to restore the applicable default.
	EmailDelivery *V2BillingGroupUpdateBillingSettingsEmailDeliveryParams `form:"email_delivery" json:"email_delivery,omitempty"`
	// Settings for invoice presentation. Set to null to restore the applicable defaults.
	InvoicePresentation *V2BillingGroupUpdateBillingSettingsInvoicePresentationParams `form:"invoice_presentation" json:"invoice_presentation,omitempty"`
	// The amount of time until the invoice is overdue for payment. Set to null to restore the applicable default.
	InvoiceTimeUntilDue *V2BillingGroupUpdateBillingSettingsInvoiceTimeUntilDueParams `form:"invoice_time_until_due" json:"invoice_time_until_due,omitempty"`
	// The ID of the payment method configuration, which controls which payment methods are displayed to your customers. Set to null to restore the applicable default.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration to be stored on the object. Set to null to clear the stored value.
	PaymentMethodOptions *V2BillingGroupUpdateBillingSettingsPaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// Tax calculation settings. Set to null to restore the applicable default.
	Tax *V2BillingGroupUpdateBillingSettingsTaxParams `form:"tax" json:"tax,omitempty"`
}

// Update a BillingGroup.
type V2BillingGroupUpdateParams struct {
	Params `form:"*"`
	// Billing setting overrides to update. Set to null to restore all applicable defaults.
	BillingSettings *V2BillingGroupUpdateBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// The ID of the default PaymentMethod object. Set to null to use the customer's default payment method.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// An internal key you can use to find a particular billing group. Set to null to clear it. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs to add, update, or remove.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingGroupUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
