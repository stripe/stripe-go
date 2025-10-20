//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all CollectionSetting objects.
type V2BillingCollectionSettingListParams struct {
	Params `form:"*"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Only return the settings with these lookup_keys, if any exist.
	// You can specify up to 10 lookup_keys.
	LookupKeys []*string `form:"lookup_keys,flat_array" json:"lookup_keys,omitempty"`
}

// Controls emails for when the payment is due. For example after the invoice is finilized and transition to Open state.
type V2BillingCollectionSettingEmailDeliveryPaymentDueParams struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled *bool `form:"enabled" json:"enabled"`
	// If true the payment link to hosted invocie page would be included in email and PDF of the invoice.
	IncludePaymentLink *bool `form:"include_payment_link" json:"include_payment_link"`
}

// Email delivery setting.
type V2BillingCollectionSettingEmailDeliveryParams struct {
	// Controls emails for when the payment is due. For example after the invoice is finilized and transition to Open state.
	PaymentDue *V2BillingCollectionSettingEmailDeliveryPaymentDueParams `form:"payment_due" json:"payment_due,omitempty"`
}

// Additional fields for Mandate creation.
type V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type" json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsACSSDebitParams struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingCollectionSettingPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingCollectionSettingPaymentMethodOptionsBancontactParams struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage *string `form:"preferred_language" json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType *string `form:"amount_type" json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description *string `form:"description" json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network *string `form:"network" json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.corp.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.corp.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure" json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information.
	Country *string `form:"country" json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer" json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type *string `form:"type" json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType *string `form:"funding_type" json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsKonbiniParams struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsSEPADebitParams struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []*string `form:"account_subcategories,flat_array" json:"account_subcategories,omitempty"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters" json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []*string `form:"permissions,flat_array" json:"permissions,omitempty"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch,flat_array" json:"prefetch,omitempty"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections" json:"financial_connections"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method"`
}

// Payment Method specific configuration to be stored on the object.
type V2BillingCollectionSettingPaymentMethodOptionsParams struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingCollectionSettingPaymentMethodOptionsACSSDebitParams `form:"acss_debit" json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingCollectionSettingPaymentMethodOptionsBancontactParams `form:"bancontact" json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingCollectionSettingPaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingCollectionSettingPaymentMethodOptionsCustomerBalanceParams `form:"customer_balance" json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingCollectionSettingPaymentMethodOptionsKonbiniParams `form:"konbini" json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingCollectionSettingPaymentMethodOptionsSEPADebitParams `form:"sepa_debit" json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingCollectionSettingPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// Create a CollectionSetting object.
type V2BillingCollectionSettingParams struct {
	Params `form:"*"`
	// Either automatic, or send_invoice. When charging automatically, Stripe will attempt to pay this
	// bill at the end of the period using the payment method attached to the payer profile. When sending an invoice,
	// Stripe will email your payer profile an invoice with payment instructions.
	// Defaults to automatic.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// An optional customer-facing display name for the CollectionSetting object.
	// To remove the display name, set it to an empty string in the request.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Email delivery settings.
	EmailDelivery *V2BillingCollectionSettingEmailDeliveryParams `form:"email_delivery" json:"email_delivery,omitempty"`
	// Optionally change the live version of the CollectionSetting. Billing Cadences and other objects that refer to this
	// CollectionSetting will use this version when no overrides are set. Providing `live_version = "latest"` will set the
	// CollectionSetting's `live_version` to its latest version.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// A lookup key used to retrieve settings dynamically from a static string.
	// This may be up to 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// The ID of the PaymentMethodConfiguration object, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration to be stored on the object.
	PaymentMethodOptions *V2BillingCollectionSettingPaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
}

// Controls emails for when the payment is due. For example after the invoice is finilized and transition to Open state.
type V2BillingCollectionSettingCreateEmailDeliveryPaymentDueParams struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled *bool `form:"enabled" json:"enabled"`
	// If true the payment link to hosted invocie page would be included in email and PDF of the invoice.
	IncludePaymentLink *bool `form:"include_payment_link" json:"include_payment_link"`
}

// Email delivery setting.
type V2BillingCollectionSettingCreateEmailDeliveryParams struct {
	// Controls emails for when the payment is due. For example after the invoice is finilized and transition to Open state.
	PaymentDue *V2BillingCollectionSettingCreateEmailDeliveryPaymentDueParams `form:"payment_due" json:"payment_due,omitempty"`
}

// Additional fields for Mandate creation.
type V2BillingCollectionSettingCreatePaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type" json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingCollectionSettingCreatePaymentMethodOptionsACSSDebitParams struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingCollectionSettingCreatePaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingCollectionSettingCreatePaymentMethodOptionsBancontactParams struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage *string `form:"preferred_language" json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingCollectionSettingCreatePaymentMethodOptionsCardMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType *string `form:"amount_type" json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description *string `form:"description" json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingCollectionSettingCreatePaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingCollectionSettingCreatePaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network *string `form:"network" json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.corp.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.corp.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure" json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingCollectionSettingCreatePaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information.
	Country *string `form:"country" json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingCollectionSettingCreatePaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingCollectionSettingCreatePaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer" json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type *string `form:"type" json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingCollectionSettingCreatePaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingCollectionSettingCreatePaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType *string `form:"funding_type" json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingCollectionSettingCreatePaymentMethodOptionsKonbiniParams struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingCollectionSettingCreatePaymentMethodOptionsSEPADebitParams struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingCollectionSettingCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []*string `form:"account_subcategories,flat_array" json:"account_subcategories,omitempty"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingCollectionSettingCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingCollectionSettingCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters" json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []*string `form:"permissions,flat_array" json:"permissions,omitempty"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch,flat_array" json:"prefetch,omitempty"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingCollectionSettingCreatePaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingCollectionSettingCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections" json:"financial_connections"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method"`
}

// Payment Method specific configuration to be stored on the object.
type V2BillingCollectionSettingCreatePaymentMethodOptionsParams struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingCollectionSettingCreatePaymentMethodOptionsACSSDebitParams `form:"acss_debit" json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingCollectionSettingCreatePaymentMethodOptionsBancontactParams `form:"bancontact" json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingCollectionSettingCreatePaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingCollectionSettingCreatePaymentMethodOptionsCustomerBalanceParams `form:"customer_balance" json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingCollectionSettingCreatePaymentMethodOptionsKonbiniParams `form:"konbini" json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingCollectionSettingCreatePaymentMethodOptionsSEPADebitParams `form:"sepa_debit" json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingCollectionSettingCreatePaymentMethodOptionsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// Create a CollectionSetting object.
type V2BillingCollectionSettingCreateParams struct {
	Params `form:"*"`
	// Either automatic, or send_invoice. When charging automatically, Stripe will attempt to pay this
	// bill at the end of the period using the payment method attached to the payer profile. When sending an invoice,
	// Stripe will email your payer profile an invoice with payment instructions.
	// Defaults to automatic.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// An optional customer-facing display name for the CollectionSetting object.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Email delivery setting.
	EmailDelivery *V2BillingCollectionSettingCreateEmailDeliveryParams `form:"email_delivery" json:"email_delivery,omitempty"`
	// A lookup key used to retrieve settings dynamically from a static string.
	// This may be up to 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// The ID of the PaymentMethodConfiguration object, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration to be stored on the object.
	PaymentMethodOptions *V2BillingCollectionSettingCreatePaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
}

// Retrieve a CollectionSetting by ID.
type V2BillingCollectionSettingRetrieveParams struct {
	Params `form:"*"`
}

// Controls emails for when the payment is due. For example after the invoice is finilized and transition to Open state.
type V2BillingCollectionSettingUpdateEmailDeliveryPaymentDueParams struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled *bool `form:"enabled" json:"enabled"`
	// If true the payment link to hosted invocie page would be included in email and PDF of the invoice.
	IncludePaymentLink *bool `form:"include_payment_link" json:"include_payment_link"`
}

// Email delivery settings.
type V2BillingCollectionSettingUpdateEmailDeliveryParams struct {
	// Controls emails for when the payment is due. For example after the invoice is finilized and transition to Open state.
	PaymentDue *V2BillingCollectionSettingUpdateEmailDeliveryPaymentDueParams `form:"payment_due" json:"payment_due,omitempty"`
}

// Additional fields for Mandate creation.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type" json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsACSSDebitParams struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingCollectionSettingUpdatePaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsBancontactParams struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage *string `form:"preferred_language" json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsCardMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType *string `form:"amount_type" json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description *string `form:"description" json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingCollectionSettingUpdatePaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network *string `form:"network" json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.corp.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.corp.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure" json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information.
	Country *string `form:"country" json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingCollectionSettingUpdatePaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer" json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type *string `form:"type" json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingCollectionSettingUpdatePaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType *string `form:"funding_type" json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsKonbiniParams struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsSEPADebitParams struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []*string `form:"account_subcategories,flat_array" json:"account_subcategories,omitempty"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingCollectionSettingUpdatePaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters" json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []*string `form:"permissions,flat_array" json:"permissions,omitempty"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch,flat_array" json:"prefetch,omitempty"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingCollectionSettingUpdatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections" json:"financial_connections"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method"`
}

// Payment Method specific configuration to be stored on the object.
type V2BillingCollectionSettingUpdatePaymentMethodOptionsParams struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingCollectionSettingUpdatePaymentMethodOptionsACSSDebitParams `form:"acss_debit" json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingCollectionSettingUpdatePaymentMethodOptionsBancontactParams `form:"bancontact" json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingCollectionSettingUpdatePaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingCollectionSettingUpdatePaymentMethodOptionsCustomerBalanceParams `form:"customer_balance" json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingCollectionSettingUpdatePaymentMethodOptionsKonbiniParams `form:"konbini" json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingCollectionSettingUpdatePaymentMethodOptionsSEPADebitParams `form:"sepa_debit" json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingCollectionSettingUpdatePaymentMethodOptionsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// Update fields on an existing CollectionSetting.
type V2BillingCollectionSettingUpdateParams struct {
	Params `form:"*"`
	// Either automatic, or send_invoice. When charging automatically, Stripe will attempt to pay this
	// bill at the end of the period using the payment method attached to the payer profile. When sending an invoice,
	// Stripe will email your payer profile an invoice with payment instructions.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// An optional customer-facing display name for the CollectionSetting object.
	// To remove the display name, set it to an empty string in the request.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Email delivery settings.
	EmailDelivery *V2BillingCollectionSettingUpdateEmailDeliveryParams `form:"email_delivery" json:"email_delivery,omitempty"`
	// Optionally change the live version of the CollectionSetting. Billing Cadences and other objects that refer to this
	// CollectionSetting will use this version when no overrides are set. Providing `live_version = "latest"` will set the
	// CollectionSetting's `live_version` to its latest version.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// A lookup key used to retrieve settings dynamically from a static string.
	// This may be up to 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// The ID of the PaymentMethodConfiguration object, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration to be stored on the object.
	PaymentMethodOptions *V2BillingCollectionSettingUpdatePaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
}
