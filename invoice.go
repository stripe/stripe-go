//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
)

// The status of the most recent automated tax calculation for this invoice.
type InvoiceAutomaticTaxStatus string

// List of values that InvoiceAutomaticTaxStatus can take
const (
	InvoiceAutomaticTaxStatusComplete               InvoiceAutomaticTaxStatus = "complete"
	InvoiceAutomaticTaxStatusFailed                 InvoiceAutomaticTaxStatus = "failed"
	InvoiceAutomaticTaxStatusRequiresLocationInputs InvoiceAutomaticTaxStatus = "requires_location_inputs"
)

// Indicates the reason why the invoice was created. `subscription_cycle` indicates an invoice created by a subscription advancing into a new period. `subscription_create` indicates an invoice created due to creating a subscription. `subscription_update` indicates an invoice created due to updating a subscription. `subscription` is set for all old invoices to indicate either a change to a subscription or a period advancement. `manual` is set for all invoices unrelated to a subscription (for example: created via the invoice editor). The `upcoming` value is reserved for simulated invoices per the upcoming invoice endpoint. `subscription_threshold` indicates an invoice created due to a billing threshold being reached.
type InvoiceBillingReason string

// List of values that InvoiceBillingReason can take
const (
	InvoiceBillingReasonAutomaticPendingInvoiceItemInvoice InvoiceBillingReason = "automatic_pending_invoice_item_invoice"
	InvoiceBillingReasonManual                             InvoiceBillingReason = "manual"
	InvoiceBillingReasonQuoteAccept                        InvoiceBillingReason = "quote_accept"
	InvoiceBillingReasonSubscription                       InvoiceBillingReason = "subscription"
	InvoiceBillingReasonSubscriptionCreate                 InvoiceBillingReason = "subscription_create"
	InvoiceBillingReasonSubscriptionCycle                  InvoiceBillingReason = "subscription_cycle"
	InvoiceBillingReasonSubscriptionThreshold              InvoiceBillingReason = "subscription_threshold"
	InvoiceBillingReasonSubscriptionUpdate                 InvoiceBillingReason = "subscription_update"
	InvoiceBillingReasonUpcoming                           InvoiceBillingReason = "upcoming"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions.
type InvoiceCollectionMethod string

// List of values that InvoiceCollectionMethod can take
const (
	InvoiceCollectionMethodChargeAutomatically InvoiceCollectionMethod = "charge_automatically"
	InvoiceCollectionMethodSendInvoice         InvoiceCollectionMethod = "send_invoice"
)

// Transaction type of the mandate.
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Bank account verification method.
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodInstant       InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecureAny       InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure = "any"
	InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecureAutomatic InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
)

// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
type InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType string

// List of values that InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType can take
const (
	InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceFundingTypeBankTransfer InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType = "bank_transfer"
)

// The list of permissions to request. The `payment_method` permission must be included.
type InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission string

// List of values that InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission can take
const (
	InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionBalances      InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "balances"
	InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionPaymentMethod InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "payment_method"
	InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionTransactions  InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "transactions"
)

// Bank account verification method.
type InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod string

// List of values that InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod can take
const (
	InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethodAutomatic     InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "automatic"
	InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethodInstant       InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "instant"
	InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethodMicrodeposits InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod = "microdeposits"
)

// The list of payment method types (e.g. card) to provide to the invoice's PaymentIntent. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice's default payment method, the subscription's default payment method, the customer's default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).
type InvoicePaymentSettingsPaymentMethodType string

// List of values that InvoicePaymentSettingsPaymentMethodType can take
const (
	InvoicePaymentSettingsPaymentMethodTypeAchCreditTransfer  InvoicePaymentSettingsPaymentMethodType = "ach_credit_transfer"
	InvoicePaymentSettingsPaymentMethodTypeAchDebit           InvoicePaymentSettingsPaymentMethodType = "ach_debit"
	InvoicePaymentSettingsPaymentMethodTypeACSSDebit          InvoicePaymentSettingsPaymentMethodType = "acss_debit"
	InvoicePaymentSettingsPaymentMethodTypeAUBECSDebit        InvoicePaymentSettingsPaymentMethodType = "au_becs_debit"
	InvoicePaymentSettingsPaymentMethodTypeBACSDebit          InvoicePaymentSettingsPaymentMethodType = "bacs_debit"
	InvoicePaymentSettingsPaymentMethodTypeBancontact         InvoicePaymentSettingsPaymentMethodType = "bancontact"
	InvoicePaymentSettingsPaymentMethodTypeBoleto             InvoicePaymentSettingsPaymentMethodType = "boleto"
	InvoicePaymentSettingsPaymentMethodTypeCard               InvoicePaymentSettingsPaymentMethodType = "card"
	InvoicePaymentSettingsPaymentMethodTypeCustomerBalance    InvoicePaymentSettingsPaymentMethodType = "customer_balance"
	InvoicePaymentSettingsPaymentMethodTypeFPX                InvoicePaymentSettingsPaymentMethodType = "fpx"
	InvoicePaymentSettingsPaymentMethodTypeGiropay            InvoicePaymentSettingsPaymentMethodType = "giropay"
	InvoicePaymentSettingsPaymentMethodTypeGrabpay            InvoicePaymentSettingsPaymentMethodType = "grabpay"
	InvoicePaymentSettingsPaymentMethodTypeIdeal              InvoicePaymentSettingsPaymentMethodType = "ideal"
	InvoicePaymentSettingsPaymentMethodTypeKonbini            InvoicePaymentSettingsPaymentMethodType = "konbini"
	InvoicePaymentSettingsPaymentMethodTypeLink               InvoicePaymentSettingsPaymentMethodType = "link"
	InvoicePaymentSettingsPaymentMethodTypePayNow             InvoicePaymentSettingsPaymentMethodType = "paynow"
	InvoicePaymentSettingsPaymentMethodTypePromptPay          InvoicePaymentSettingsPaymentMethodType = "promptpay"
	InvoicePaymentSettingsPaymentMethodTypeSepaCreditTransfer InvoicePaymentSettingsPaymentMethodType = "sepa_credit_transfer"
	InvoicePaymentSettingsPaymentMethodTypeSepaDebit          InvoicePaymentSettingsPaymentMethodType = "sepa_debit"
	InvoicePaymentSettingsPaymentMethodTypeSofort             InvoicePaymentSettingsPaymentMethodType = "sofort"
	InvoicePaymentSettingsPaymentMethodTypeUSBankAccount      InvoicePaymentSettingsPaymentMethodType = "us_bank_account"
	InvoicePaymentSettingsPaymentMethodTypeWechatPay          InvoicePaymentSettingsPaymentMethodType = "wechat_pay"
)

// The status of the invoice, one of `draft`, `open`, `paid`, `uncollectible`, or `void`. [Learn more](https://stripe.com/docs/billing/invoices/workflow#workflow-overview)
type InvoiceStatus string

// List of values that InvoiceStatus can take
const (
	InvoiceStatusDeleted       InvoiceStatus = "deleted"
	InvoiceStatusDraft         InvoiceStatus = "draft"
	InvoiceStatusOpen          InvoiceStatus = "open"
	InvoiceStatusPaid          InvoiceStatus = "paid"
	InvoiceStatusUncollectible InvoiceStatus = "uncollectible"
	InvoiceStatusVoid          InvoiceStatus = "void"
)

// Search for invoices you've previously created using Stripe's [Search Query Language](https://stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
type InvoiceSearchParams struct {
	SearchParams `form:"*"`
	// A cursor for pagination across multiple pages of results. Don't include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
	Page *string `form:"page"`
}

// You can list all invoices, or list the invoices for a specific customer. The invoices are returned sorted by creation date, with the most recently created invoices appearing first.
type InvoiceListParams struct {
	ListParams `form:"*"`
	// The collection method of the invoice to retrieve. Either `charge_automatically` or `send_invoice`.
	CollectionMethod *string           `form:"collection_method"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	// Only return invoices for the customer specified by this customer ID.
	Customer     *string           `form:"customer"`
	DueDate      *int64            `form:"due_date"`
	DueDateRange *RangeQueryParams `form:"due_date"`
	// The status of the invoice, one of `draft`, `open`, `paid`, `uncollectible`, or `void`. [Learn more](https://stripe.com/docs/billing/invoices/workflow#workflow-overview)
	Status *string `form:"status"`
	// Only return invoices for the subscription specified by this subscription ID.
	Subscription *string `form:"subscription"`
}

// Settings for automatic tax lookup for this invoice.
type InvoiceAutomaticTaxParams struct {
	// Whether Stripe automatically computes tax on this invoice. Note that incompatible invoice items (invoice items with manually specified [tax rates](https://stripe.com/docs/api/tax_rates), negative amounts, or `tax_behavior=unspecified`) cannot be added to automatic tax invoices.
	Enabled *bool `form:"enabled"`
}

// A list of up to 4 custom fields to be displayed on the invoice.
type InvoiceCustomFieldParams struct {
	// The name of the custom field. This may be up to 30 characters.
	Name *string `form:"name"`
	// The value of the custom field. This may be up to 30 characters.
	Value *string `form:"value"`
}

// The coupons to redeem into discounts for the invoice. If not specified, inherits the discount from the invoice's customer. Pass an empty string to avoid inheriting any discounts.
type InvoiceDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
}

// Additional fields for Mandate creation
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type"`
}

// If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsBancontactParams struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage *string `form:"preferred_language"`
}

// The selected installment plan to use for this invoice.
type InvoicePaymentSettingsPaymentMethodOptionsCardInstallmentsPlanParams struct {
	// For `fixed_count` installment plans, this is the number of installment payments your customer will make to their credit card.
	Count *int64 `form:"count"`
	// For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card.
	// One of `month`.
	Interval *string `form:"interval"`
	// Type of installment plan, one of `fixed_count`.
	Type *string `form:"type"`
}

// Installment configuration for payments attempted on this invoice (Mexico Only).
//
// For more information, see the [installments integration guide](https://stripe.com/docs/payments/installments).
type InvoicePaymentSettingsPaymentMethodOptionsCardInstallmentsParams struct {
	// Setting to true enables installments for this invoice.
	// Setting to false will prevent any selected plan from applying to a payment.
	Enabled *bool `form:"enabled"`
	// The selected installment plan to use for this invoice.
	Plan *InvoicePaymentSettingsPaymentMethodOptionsCardInstallmentsPlanParams `form:"plan"`
}

// If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsCardParams struct {
	// Installment configuration for payments attempted on this invoice (Mexico Only).
	//
	// For more information, see the [installments integration guide](https://stripe.com/docs/payments/installments).
	Installments *InvoicePaymentSettingsPaymentMethodOptionsCardInstallmentsParams `form:"installments"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure"`
}

// Configuration for eu_bank_transfer funding type.
type InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information. Permitted values include: `DE`, `ES`, `FR`, `IE`, or `NL`.
	Country *string `form:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for eu_bank_transfer funding type.
	EUBankTransfer *InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer"`
	// The bank transfer type that can be used for funding. Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, or `mx_bank_transfer`.
	Type *string `form:"type"`
}

// If paying by `customer_balance`, this sub-hash contains details about the Bank transfer payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType *string `form:"funding_type"`
}

// If paying by `konbini`, this sub-hash contains details about the Konbini payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsKonbiniParams struct{}

// Additional fields for Financial Connections Session creation
type InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included. Valid permissions include: `balances`, `ownership`, `payment_method`, and `transactions`.
	Permissions []*string `form:"permissions"`
}

// If paying by `us_bank_account`, this sub-hash contains details about the ACH direct debit payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation
	FinancialConnections *InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// Payment-method-specific configuration to provide to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsParams struct {
	// If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
	ACSSDebit *InvoicePaymentSettingsPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	// If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the invoice's PaymentIntent.
	Bancontact *InvoicePaymentSettingsPaymentMethodOptionsBancontactParams `form:"bancontact"`
	// If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the invoice's PaymentIntent.
	Card *InvoicePaymentSettingsPaymentMethodOptionsCardParams `form:"card"`
	// If paying by `customer_balance`, this sub-hash contains details about the Bank transfer payment method options to pass to the invoice's PaymentIntent.
	CustomerBalance *InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceParams `form:"customer_balance"`
	// If paying by `konbini`, this sub-hash contains details about the Konbini payment method options to pass to the invoice's PaymentIntent.
	Konbini *InvoicePaymentSettingsPaymentMethodOptionsKonbiniParams `form:"konbini"`
	// If paying by `us_bank_account`, this sub-hash contains details about the ACH direct debit payment method options to pass to the invoice's PaymentIntent.
	USBankAccount *InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account"`
}

// Configuration settings for the PaymentIntent that is generated when the invoice is finalized.
type InvoicePaymentSettingsParams struct {
	// ID of the mandate to be used for this invoice. It must correspond to the payment method used to pay the invoice, including the invoice's default_payment_method or default_source, if set.
	DefaultMandate *string `form:"default_mandate"`
	// Payment-method-specific configuration to provide to the invoice's PaymentIntent.
	PaymentMethodOptions *InvoicePaymentSettingsPaymentMethodOptionsParams `form:"payment_method_options"`
	// The list of payment method types (e.g. card) to provide to the invoice's PaymentIntent. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice's default payment method, the subscription's default payment method, the customer's default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).
	PaymentMethodTypes []*string `form:"payment_method_types"`
}

// Options for invoice PDF rendering.
type InvoiceRenderingOptionsParams struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs. One of `exclude_tax` or `include_inclusive_tax`. `include_inclusive_tax` will include inclusive tax (and exclude exclusive tax) in invoice PDF amounts. `exclude_tax` will exclude all tax (inclusive and exclusive alike) from invoice PDF amounts.
	AmountTaxDisplay *string `form:"amount_tax_display"`
}

// If specified, the funds from the invoice will be transferred to the destination and the ID of the resulting transfer will be found on the invoice's charge.
type InvoiceTransferDataParams struct {
	// The amount that will be transferred automatically when the invoice is paid. If no amount is set, the full amount is transferred.
	Amount *int64 `form:"amount"`
	// ID of an existing, connected Stripe account.
	Destination *string `form:"destination"`
}

// This endpoint creates a draft invoice for a given customer. The draft invoice created pulls in all pending invoice items on that customer, including prorations. The invoice remains a draft until you [finalize the invoice, which allows you to [pay](#pay_invoice) or <a href="#send_invoice">send](https://stripe.com/docs/api#finalize_invoice) the invoice to your customers.
type InvoiceParams struct {
	Params `form:"*"`
	// The account tax IDs associated with the invoice. Only editable when the invoice is a draft.
	AccountTaxIDs []*string `form:"account_tax_ids"`
	// A fee in cents (or local equivalent) that will be applied to the invoice and transferred to the application owner's Stripe account. The request must be made with an OAuth key or the Stripe-Account header in order to take an application fee. For more information, see the application fees [documentation](https://stripe.com/docs/billing/invoices/connect#collecting-fees).
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// Controls whether Stripe will perform [automatic collection](https://stripe.com/docs/billing/invoices/workflow/#auto_advance) of the invoice.
	AutoAdvance *bool `form:"auto_advance"`
	// Settings for automatic tax lookup for this invoice.
	AutomaticTax *InvoiceAutomaticTaxParams `form:"automatic_tax"`
	// Either `charge_automatically` or `send_invoice`. This field can be updated only on `draft` invoices.
	CollectionMethod *string `form:"collection_method"`
	// The currency to preview this invoice in. Defaults to that of `customer` if not specified.
	Currency *string `form:"currency"`
	// The identifier of the customer whose upcoming invoice you'd like to retrieve.
	Customer *string `form:"customer"`
	// A list of up to 4 custom fields to be displayed on the invoice. If a value for `custom_fields` is specified, the list specified will replace the existing custom field list on this invoice. Pass an empty string to remove previously-defined fields.
	CustomFields []*InvoiceCustomFieldParams `form:"custom_fields"`
	// The number of days from which the invoice is created until it is due. Only valid for invoices where `collection_method=send_invoice`. This field can only be updated on `draft` invoices.
	DaysUntilDue *int64 `form:"days_until_due"`
	// ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription's default payment method, if any, or to the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// ID of the default payment source for the invoice. It must belong to the customer associated with the invoice and be in a chargeable state. If not set, defaults to the subscription's default source, if any, or to the customer's default source.
	DefaultSource *string `form:"default_source"`
	// The tax rates that will apply to any line item that does not have `tax_rates` set. Pass an empty string to remove previously-defined tax rates.
	DefaultTaxRates []*string `form:"default_tax_rates"`
	// An arbitrary string attached to the object. Often useful for displaying to users. Referenced as 'memo' in the Dashboard.
	Description *string `form:"description"`
	// The discounts that will apply to the invoice. Pass an empty string to remove previously-defined discounts.
	Discounts []*InvoiceDiscountParams `form:"discounts"`
	// The date on which payment for this invoice is due. Only valid for invoices where `collection_method=send_invoice`. This field can only be updated on `draft` invoices.
	DueDate *int64 `form:"due_date"`
	// Footer to be displayed on the invoice.
	Footer *string `form:"footer"`
	// The account (if any) for which the funds of the invoice payment are intended. If set, the invoice will be presented with the branding and support information of the specified account. See the [Invoices with Connect](https://stripe.com/docs/billing/invoices/connect) documentation for details.
	OnBehalfOf *string `form:"on_behalf_of"`
	Paid       *bool   `form:"paid"`
	// Configuration settings for the PaymentIntent that is generated when the invoice is finalized.
	PaymentSettings *InvoicePaymentSettingsParams `form:"payment_settings"`
	// How to handle pending invoice items on invoice creation. One of `include`, `exclude`, or `include_and_require`. `include` will include any pending invoice items, and will create an empty draft invoice if no pending invoice items exist. `include_and_require` will include any pending invoice items, if no pending invoice items exist then the request will fail. `exclude` will always create an empty invoice draft regardless if there are pending invoice items or not. Defaults to `include_and_require` if the parameter is omitted.
	PendingInvoiceItemsBehavior *string `form:"pending_invoice_items_behavior"`
	// Options for invoice PDF rendering.
	RenderingOptions *InvoiceRenderingOptionsParams `form:"rendering_options"`
	// The identifier of the unstarted schedule whose upcoming invoice you'd like to retrieve. Cannot be used with subscription or subscription fields.
	Schedule *string `form:"schedule"`
	// Extra information about a charge for the customer's credit card statement. It must contain at least one letter. If not specified and this invoice is part of a subscription, the default `statement_descriptor` will be set to the first subscription item's product's `statement_descriptor`.
	StatementDescriptor *string `form:"statement_descriptor"`
	// The identifier of the subscription for which you'd like to retrieve the upcoming invoice. If not provided, but a `subscription_items` is provided, you will preview creating a subscription with those items. If neither `subscription` nor `subscription_items` is provided, you will retrieve the next upcoming invoice from among the customer's subscriptions.
	Subscription *string `form:"subscription"`
	// If specified, the funds from the invoice will be transferred to the destination and the ID of the resulting transfer will be found on the invoice's charge. This will be unset if you POST an empty value.
	TransferData *InvoiceTransferDataParams `form:"transfer_data"`
	// These are all for exclusive use by GetNext.
	// The code of the coupon to apply. If `subscription` or `subscription_items` is provided, the invoice returned will preview updating or creating a subscription with that coupon. Otherwise, it will preview applying that coupon to the customer for the next upcoming invoice from among the customer's subscriptions. The invoice can be previewed without a coupon by passing this value as an empty string.
	Coupon *string `form:"coupon"`
	// Details about the customer you want to invoice or overrides for an existing customer.
	CustomerDetails *InvoiceUpcomingCustomerDetailsParams `form:"customer_details"`
	// List of invoice items to add or update in the upcoming invoice preview.
	InvoiceItems []*InvoiceUpcomingInvoiceItemParams `form:"invoice_items"`
	// For new subscriptions, a future timestamp to anchor the subscription's [billing cycle](https://stripe.com/docs/subscriptions/billing-cycle). This is used to determine the date of the first full invoice, and, for plans with `month` or `year` intervals, the day of the month for subsequent invoices. For existing subscriptions, the value can only be set to `now` or `unchanged`.
	SubscriptionBillingCycleAnchor          *int64 `form:"subscription_billing_cycle_anchor"`
	SubscriptionBillingCycleAnchorNow       *bool  `form:"-"` // See custom AppendTo
	SubscriptionBillingCycleAnchorUnchanged *bool  `form:"-"` // See custom AppendTo
	// Timestamp indicating when the subscription should be scheduled to cancel. Will prorate if within the current period and prorations have been enabled using `proration_behavior`.
	SubscriptionCancelAt *int64 `form:"subscription_cancel_at"`
	// Boolean indicating whether this subscription should cancel at the end of the current period.
	SubscriptionCancelAtPeriodEnd *bool `form:"subscription_cancel_at_period_end"`
	// This simulates the subscription being canceled or expired immediately.
	SubscriptionCancelNow *bool `form:"subscription_cancel_now"`
	// If provided, the invoice returned will preview updating or creating a subscription with these default tax rates. The default tax rates will apply to any line item that does not have `tax_rates` set.
	SubscriptionDefaultTaxRates []*string `form:"subscription_default_tax_rates"`
	// A list of up to 20 subscription items, each with an attached price.
	SubscriptionItems []*SubscriptionItemsParams `form:"subscription_items"`
	SubscriptionPlan  *string                    `form:"subscription_plan"`
	// Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes.
	SubscriptionProrationBehavior *string `form:"subscription_proration_behavior"`
	// If previewing an update to a subscription, and doing proration, `subscription_proration_date` forces the proration to be calculated as though the update was done at the specified time. The time given must be within the current subscription period, and cannot be before the subscription was on its current plan. If set, `subscription`, and one of `subscription_items`, or `subscription_trial_end` are required. Also, `subscription_proration_behavior` cannot be set to 'none'.
	SubscriptionProrationDate *int64 `form:"subscription_proration_date"`
	SubscriptionQuantity      *int64 `form:"subscription_quantity"`
	// Date a subscription is intended to start (can be future or past)
	SubscriptionStartDate *int64 `form:"subscription_start_date"`
	// If provided, the invoice returned will preview updating or creating a subscription with that trial end. If set, one of `subscription_items` or `subscription` is required.
	SubscriptionTrialEnd    *int64 `form:"subscription_trial_end"`
	SubscriptionTrialEndNow *bool  `form:"-"` // See custom AppendTo
	// Indicates if a plan's `trial_period_days` should be applied to the subscription. Setting `subscription_trial_end` per subscription is preferred, and this defaults to `false`. Setting this flag to `true` together with `subscription_trial_end` is not allowed. See [Using trial periods on subscriptions](https://stripe.com/docs/billing/subscriptions/trials) to learn more.
	SubscriptionTrialFromPlan *bool `form:"subscription_trial_from_plan"`
}

// AppendTo implements custom encoding logic for InvoiceParams.
func (i *InvoiceParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(i.SubscriptionBillingCycleAnchorNow) {
		body.Add(form.FormatKey(append(keyParts, "subscription_billing_cycle_anchor")), "now")
	}
	if BoolValue(i.SubscriptionBillingCycleAnchorUnchanged) {
		body.Add(form.FormatKey(append(keyParts, "subscription_billing_cycle_anchor")), "unchanged")
	}
	if BoolValue(i.SubscriptionTrialEndNow) {
		body.Add(form.FormatKey(append(keyParts, "subscription_trial_end")), "now")
	}
}

type InvoiceUpcomingAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// The customer's shipping information. Appears on invoices emailed to this customer.
type InvoiceUpcomingCustomerDetailsShippingParams struct {
	// Customer shipping address.
	Address *AddressParams `form:"address"`
	// Customer name.
	Name *string `form:"name"`
	// Customer phone (including extension).
	Phone *string `form:"phone"`
}

// Tax details about the customer.
type InvoiceUpcomingCustomerDetailsTaxParams struct {
	// A recent IP address of the customer used for tax reporting and tax location inference. Stripe recommends updating the IP address when a new PaymentMethod is attached or the address field on the customer is updated. We recommend against updating this field more frequently since it could result in unexpected tax location/reporting outcomes.
	IPAddress *string `form:"ip_address"`
}

// The customer's tax IDs.
type InvoiceUpcomingCustomerDetailsTaxIDParams struct {
	// Type of the tax ID, one of `ae_trn`, `au_abn`, `au_arn`, `bg_uic`, `br_cnpj`, `br_cpf`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `ch_vat`, `cl_tin`, `es_cif`, `eu_oss_vat`, `eu_vat`, `gb_vat`, `ge_vat`, `hk_br`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `jp_cn`, `jp_rn`, `kr_brn`, `li_uid`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `no_vat`, `nz_gst`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `th_vat`, `tw_vat`, `ua_vat`, `us_ein`, or `za_vat`
	Type *string `form:"type"`
	// Value of the tax ID.
	Value *string `form:"value"`
}

// Details about the customer you want to invoice or overrides for an existing customer.
type InvoiceUpcomingCustomerDetailsParams struct {
	// The customer's address.
	Address *AddressParams `form:"address"`
	// The customer's shipping information. Appears on invoices emailed to this customer.
	Shipping *InvoiceUpcomingCustomerDetailsShippingParams `form:"shipping"`
	// Tax details about the customer.
	Tax *InvoiceUpcomingCustomerDetailsTaxParams `form:"tax"`
	// The customer's tax exemption. One of `none`, `exempt`, or `reverse`.
	TaxExempt *string `form:"tax_exempt"`
	// The customer's tax IDs.
	TaxIDs []*InvoiceUpcomingCustomerDetailsTaxIDParams `form:"tax_ids"`
}

// The period associated with this invoice item. When set to different values, the period will be rendered on the invoice.
type InvoiceUpcomingInvoiceItemPeriodParams struct {
	// The end of the period, which must be greater than or equal to the start.
	End *int64 `form:"end"`
	// The start of the period.
	Start *int64 `form:"start"`
}

// List of invoice items to add or update in the upcoming invoice preview.
type InvoiceUpcomingInvoiceItemParams struct {
	// The integer amount in cents (or local equivalent) of previewed invoice item.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies). Only applicable to new invoice items.
	Currency *string `form:"currency"`
	// An arbitrary string which you can attach to the invoice item. The description is displayed in the invoice for easy tracking.
	Description *string `form:"description"`
	// Explicitly controls whether discounts apply to this invoice item. Defaults to true, except for negative invoice items.
	Discountable *bool `form:"discountable"`
	// The coupons to redeem into discounts for the invoice item in the preview.
	Discounts []*InvoiceItemDiscountParams `form:"discounts"`
	// The ID of the invoice item to update in preview. If not specified, a new invoice item will be added to the preview of the upcoming invoice.
	InvoiceItem *string `form:"invoiceitem"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The period associated with this invoice item. When set to different values, the period will be rendered on the invoice.
	Period *InvoiceUpcomingInvoiceItemPeriodParams `form:"period"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	// Non-negative integer. The quantity of units for the invoice item.
	Quantity *int64  `form:"quantity"`
	Schedule *string `form:"schedule"`
	// The tax rates that apply to the item. When set, any `default_tax_rates` do not apply to this item.
	TaxRates []*string `form:"tax_rates"`
	// The integer unit amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. This unit_amount will be multiplied by the quantity to get the full amount. If you want to apply a credit to the customer's account, pass a negative unit_amount.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Stripe automatically creates and then attempts to collect payment on invoices for customers on subscriptions according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to attempt payment on an invoice out of the normal collection schedule or for some other reason, you can do so.
type InvoicePayParams struct {
	Params `form:"*"`
	// In cases where the source used to pay the invoice has insufficient funds, passing `forgive=true` controls whether a charge should be attempted for the full amount available on the source, up to the amount to fully pay the invoice. This effectively forgives the difference between the amount available on the source and the amount due.
	//
	// Passing `forgive=false` will fail the charge if the source hasn't been pre-funded with the right amount. An example for this case is with ACH Credit Transfers and wires: if the amount wired is less than the amount due by a small amount, you might want to forgive the difference. Defaults to `false`.
	Forgive *bool `form:"forgive"`
	// ID of the mandate to be used for this invoice. It must correspond to the payment method used to pay the invoice, including the payment_method param or the invoice's default_payment_method or default_source, if set.
	Mandate *string `form:"mandate"`
	// Indicates if a customer is on or off-session while an invoice payment is attempted. Defaults to `true` (off-session).
	OffSession *bool `form:"off_session"`
	// Boolean representing whether an invoice is paid outside of Stripe. This will result in no charge being made. Defaults to `false`.
	PaidOutOfBand *bool `form:"paid_out_of_band"`
	// A PaymentMethod to be charged. The PaymentMethod must be the ID of a PaymentMethod belonging to the customer associated with the invoice being paid.
	PaymentMethod *string `form:"payment_method"`
	// A payment source to be charged. The source must be the ID of a source belonging to the customer associated with the invoice being paid.
	Source *string `form:"source"`
}

// Stripe automatically finalizes drafts before sending and attempting payment on invoices. However, if you'd like to finalize a draft invoice manually, you can do so using this method.
type InvoiceFinalizeParams struct {
	Params `form:"*"`
	// Controls whether Stripe will perform [automatic collection](https://stripe.com/docs/invoicing/automatic-charging) of the invoice. When `false`, the invoice's state will not automatically advance without an explicit action.
	AutoAdvance *bool `form:"auto_advance"`
}

// Stripe will automatically send invoices to customers according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to manually send an invoice to your customer out of the normal schedule, you can do so. When sending invoices that have already been paid, there will be no reference to the payment in the email.
//
// Requests made in test-mode result in no emails being sent, despite sending an invoice.sent event.
type InvoiceSendParams struct {
	Params `form:"*"`
}

// Marking an invoice as uncollectible is useful for keeping track of bad debts that can be written off for accounting purposes.
type InvoiceMarkUncollectibleParams struct {
	Params `form:"*"`
}

// Mark a finalized invoice as void. This cannot be undone. Voiding an invoice is similar to [deletion](https://stripe.com/docs/api#delete_invoice), however it only applies to finalized invoices and maintains a papertrail where the invoice can still be found.
type InvoiceVoidParams struct {
	Params `form:"*"`
}

// When retrieving an invoice, you'll get a lines property containing the total count of line items and the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type InvoiceLineListParams struct {
	ListParams `form:"*"`
	// ID is the invoice ID to list invoice lines for.
	ID           *string `form:"-"` // Included in URL
	Customer     *string `form:"customer"`
	Subscription *string `form:"subscription"`
}
type InvoiceAutomaticTax struct {
	// Whether Stripe automatically computes tax on this invoice. Note that incompatible invoice items (invoice items with manually specified [tax rates](https://stripe.com/docs/api/tax_rates), negative amounts, or `tax_behavior=unspecified`) cannot be added to automatic tax invoices.
	Enabled bool `json:"enabled"`
	// The status of the most recent automated tax calculation for this invoice.
	Status InvoiceAutomaticTaxStatus `json:"status"`
}

// Custom fields displayed on the invoice.
type InvoiceCustomField struct {
	// The name of the custom field.
	Name string `json:"name"`
	// The value of the custom field.
	Value string `json:"value"`
}

// The customer's tax IDs. Until the invoice is finalized, this field will contain the same tax IDs as `customer.tax_ids`. Once the invoice is finalized, this field will no longer be updated.
type InvoiceCustomerTaxID struct {
	// The type of the tax ID, one of `eu_vat`, `br_cnpj`, `br_cpf`, `eu_oss_vat`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, or `unknown`
	Type TaxIDType `json:"type"`
	// The value of the tax ID.
	Value string `json:"value"`
}
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions struct {
	// Transaction type of the mandate.
	TransactionType InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}

// If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebit struct {
	MandateOptions *InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options"`
	// Bank account verification method.
	VerificationMethod InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsBancontact struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage string `json:"preferred_language"`
}
type InvoicePaymentSettingsPaymentMethodOptionsCardInstallments struct {
	// Whether Installments are enabled for this Invoice.
	Enabled bool `json:"enabled"`
}

// If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsCard struct {
	Installments *InvoicePaymentSettingsPaymentMethodOptionsCardInstallments `json:"installments"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}
type InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer struct {
	// The desired country code of the bank account information. Permitted values include: `DE`, `ES`, `FR`, `IE`, or `NL`.
	Country string `json:"country"`
}
type InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransfer struct {
	EUBankTransfer *InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer `json:"eu_bank_transfer"`
	// The bank transfer type that can be used for funding. Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, or `mx_bank_transfer`.
	Type string `json:"type"`
}

// If paying by `customer_balance`, this sub-hash contains details about the Bank transfer payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsCustomerBalance struct {
	BankTransfer *InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType `json:"funding_type"`
}

// If paying by `konbini`, this sub-hash contains details about the Konbini payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsKonbini struct{}
type InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	// The list of permissions to request. The `payment_method` permission must be included.
	Permissions []InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
}

// If paying by `us_bank_account`, this sub-hash contains details about the ACH direct debit payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsUSBankAccount struct {
	FinancialConnections *InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	// Bank account verification method.
	VerificationMethod InvoicePaymentSettingsPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment-method-specific configuration to provide to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptions struct {
	// If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
	ACSSDebit *InvoicePaymentSettingsPaymentMethodOptionsACSSDebit `json:"acss_debit"`
	// If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the invoice's PaymentIntent.
	Bancontact *InvoicePaymentSettingsPaymentMethodOptionsBancontact `json:"bancontact"`
	// If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the invoice's PaymentIntent.
	Card *InvoicePaymentSettingsPaymentMethodOptionsCard `json:"card"`
	// If paying by `customer_balance`, this sub-hash contains details about the Bank transfer payment method options to pass to the invoice's PaymentIntent.
	CustomerBalance *InvoicePaymentSettingsPaymentMethodOptionsCustomerBalance `json:"customer_balance"`
	// If paying by `konbini`, this sub-hash contains details about the Konbini payment method options to pass to the invoice's PaymentIntent.
	Konbini *InvoicePaymentSettingsPaymentMethodOptionsKonbini `json:"konbini"`
	// If paying by `us_bank_account`, this sub-hash contains details about the ACH direct debit payment method options to pass to the invoice's PaymentIntent.
	USBankAccount *InvoicePaymentSettingsPaymentMethodOptionsUSBankAccount `json:"us_bank_account"`
}
type InvoicePaymentSettings struct {
	// ID of the mandate to be used for this invoice. It must correspond to the payment method used to pay the invoice, including the invoice's default_payment_method or default_source, if set.
	DefaultMandate string `json:"default_mandate"`
	// Payment-method-specific configuration to provide to the invoice's PaymentIntent.
	PaymentMethodOptions *InvoicePaymentSettingsPaymentMethodOptions `json:"payment_method_options"`
	// The list of payment method types (e.g. card) to provide to the invoice's PaymentIntent. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice's default payment method, the subscription's default payment method, the customer's default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).
	PaymentMethodTypes []InvoicePaymentSettingsPaymentMethodType `json:"payment_method_types"`
}

// Options for invoice PDF rendering.
type InvoiceRenderingOptions struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.
	AmountTaxDisplay string `json:"amount_tax_display"`
}
type InvoiceStatusTransitions struct {
	// The time that the invoice draft was finalized.
	FinalizedAt int64 `json:"finalized_at"`
	// The time that the invoice was marked uncollectible.
	MarkedUncollectibleAt int64 `json:"marked_uncollectible_at"`
	// The time that the invoice was paid.
	PaidAt int64 `json:"paid_at"`
	// The time that the invoice was voided.
	VoidedAt int64 `json:"voided_at"`
}

// Indicates which line items triggered a threshold invoice.
type InvoiceThresholdReasonItemReason struct {
	// The IDs of the line items that triggered the threshold invoice.
	LineItemIDs []string `json:"line_item_ids"`
	// The quantity threshold boundary that applied to the given line item.
	UsageGTE int64 `json:"usage_gte"`
}
type InvoiceThresholdReason struct {
	// The total invoice amount threshold boundary if it triggered the threshold invoice.
	AmountGTE int64 `json:"amount_gte"`
	// Indicates which line items triggered a threshold invoice.
	ItemReasons []*InvoiceThresholdReasonItemReason `json:"item_reasons"`
}

// The aggregate amounts calculated per discount across all line items.
type InvoiceDiscountAmount struct {
	// The amount, in %s, of the discount.
	Amount int64 `json:"amount"`
	// The discount that was applied to get this discount amount.
	Discount *Discount `json:"discount"`
}

// The aggregate amounts calculated per tax rate for all line items.
type InvoiceTaxAmount struct {
	// The amount, in %s, of the tax.
	Amount int64 `json:"amount"`
	// Whether this tax amount is inclusive or exclusive.
	Inclusive bool `json:"inclusive"`
	// The tax rate that was applied to get this tax amount.
	TaxRate *TaxRate `json:"tax_rate"`
}

// The account (if any) the payment will be attributed to for tax reporting, and where funds from the payment will be transferred to for the invoice.
type InvoiceTransferData struct {
	// The amount in %s that will be transferred to the destination account when the invoice is paid. By default, the entire amount is transferred to the destination.
	Amount int64 `json:"amount"`
	// The account where funds from the payment will be transferred to upon payment success.
	Destination *Account `json:"destination"`
}

// Invoices are statements of amounts owed by a customer, and are either
// generated one-off, or generated periodically from a subscription.
//
// They contain [invoice items](https://stripe.com/docs/api#invoiceitems), and proration adjustments
// that may be caused by subscription upgrades/downgrades (if necessary).
//
// If your invoice is configured to be billed through automatic charges,
// Stripe automatically finalizes your invoice and attempts payment. Note
// that finalizing the invoice,
// [when automatic](https://stripe.com/docs/billing/invoices/workflow/#auto_advance), does
// not happen immediately as the invoice is created. Stripe waits
// until one hour after the last webhook was successfully sent (or the last
// webhook timed out after failing). If you (and the platforms you may have
// connected to) have no webhooks configured, Stripe waits one hour after
// creation to finalize the invoice.
//
// If your invoice is configured to be billed by sending an email, then based on your
// [email settings](https://dashboard.stripe.com/account/billing/automatic),
// Stripe will email the invoice to your customer and await payment. These
// emails can contain a link to a hosted page to pay the invoice.
//
// Stripe applies any customer credit on the account before determining the
// amount due for the invoice (i.e., the amount that will be actually
// charged). If the amount due for the invoice is less than Stripe's [minimum allowed charge
// per currency](https://stripe.com/docs/currencies#minimum-and-maximum-charge-amounts), the
// invoice is automatically marked paid, and we add the amount due to the
// customer's credit balance which is applied to the next invoice.
//
// More details on the customer's credit balance are
// [here](https://stripe.com/docs/billing/customer/balance).
//
// Related guide: [Send Invoices to Customers](https://stripe.com/docs/billing/invoices/sending).
type Invoice struct {
	APIResource
	// The country of the business associated with this invoice, most often the business creating the invoice.
	AccountCountry string `json:"account_country"`
	// The public name of the business associated with this invoice, most often the business creating the invoice.
	AccountName string `json:"account_name"`
	// The account tax IDs associated with the invoice. Only editable when the invoice is a draft.
	AccountTaxIDs []*TaxID `json:"account_tax_ids"`
	// Final amount due at this time for this invoice. If the invoice's total is smaller than the minimum charge amount, for example, or if there is account credit that can be applied to the invoice, the `amount_due` may be 0. If there is a positive `starting_balance` for the invoice (the customer owes money), the `amount_due` will also take that into account. The charge that gets generated for the invoice will be for the amount specified in `amount_due`.
	AmountDue int64 `json:"amount_due"`
	// The amount, in %s, that was paid.
	AmountPaid int64 `json:"amount_paid"`
	// The difference between amount_due and amount_paid, in %s.
	AmountRemaining int64 `json:"amount_remaining"`
	// ID of the Connect Application that created the invoice.
	Application *Application `json:"application"`
	// The fee in %s that will be applied to the invoice and transferred to the application owner's Stripe account when the invoice is paid.
	ApplicationFeeAmount int64 `json:"application_fee_amount"`
	// Number of payment attempts made for this invoice, from the perspective of the payment retry schedule. Any payment attempt counts as the first attempt, and subsequently only automatic retries increment the attempt count. In other words, manual payment attempts after the first attempt do not affect the retry schedule.
	AttemptCount int64 `json:"attempt_count"`
	// Whether an attempt has been made to pay the invoice. An invoice is not attempted until 1 hour after the `invoice.created` webhook, for example, so you might not want to display that invoice as unpaid to your users.
	Attempted bool `json:"attempted"`
	// Controls whether Stripe will perform [automatic collection](https://stripe.com/docs/billing/invoices/workflow/#auto_advance) of the invoice. When `false`, the invoice's state will not automatically advance without an explicit action.
	AutoAdvance  bool                 `json:"auto_advance"`
	AutomaticTax *InvoiceAutomaticTax `json:"automatic_tax"`
	// Indicates the reason why the invoice was created. `subscription_cycle` indicates an invoice created by a subscription advancing into a new period. `subscription_create` indicates an invoice created due to creating a subscription. `subscription_update` indicates an invoice created due to updating a subscription. `subscription` is set for all old invoices to indicate either a change to a subscription or a period advancement. `manual` is set for all invoices unrelated to a subscription (for example: created via the invoice editor). The `upcoming` value is reserved for simulated invoices per the upcoming invoice endpoint. `subscription_threshold` indicates an invoice created due to a billing threshold being reached.
	BillingReason InvoiceBillingReason `json:"billing_reason"`
	// ID of the latest charge generated for this invoice, if any.
	Charge *Charge `json:"charge"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions.
	CollectionMethod *InvoiceCollectionMethod `json:"collection_method"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The ID of the customer who will be billed.
	Customer *Customer `json:"customer"`
	// The customer's address. Until the invoice is finalized, this field will equal `customer.address`. Once the invoice is finalized, this field will no longer be updated.
	CustomerAddress *Address `json:"customer_address"`
	// The customer's email. Until the invoice is finalized, this field will equal `customer.email`. Once the invoice is finalized, this field will no longer be updated.
	CustomerEmail string `json:"customer_email"`
	// The customer's name. Until the invoice is finalized, this field will equal `customer.name`. Once the invoice is finalized, this field will no longer be updated.
	CustomerName *string `json:"customer_name"`
	// The customer's phone number. Until the invoice is finalized, this field will equal `customer.phone`. Once the invoice is finalized, this field will no longer be updated.
	CustomerPhone *string `json:"customer_phone"`
	// The customer's shipping information. Until the invoice is finalized, this field will equal `customer.shipping`. Once the invoice is finalized, this field will no longer be updated.
	CustomerShipping *CustomerShippingDetails `json:"customer_shipping"`
	// The customer's tax exempt status. Until the invoice is finalized, this field will equal `customer.tax_exempt`. Once the invoice is finalized, this field will no longer be updated.
	CustomerTaxExempt CustomerTaxExempt `json:"customer_tax_exempt"`
	// The customer's tax IDs. Until the invoice is finalized, this field will contain the same tax IDs as `customer.tax_ids`. Once the invoice is finalized, this field will no longer be updated.
	CustomerTaxIDs []*InvoiceCustomerTaxID `json:"customer_tax_ids"`
	// Custom fields displayed on the invoice.
	CustomFields []*InvoiceCustomField `json:"custom_fields"`
	// ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription's default payment method, if any, or to the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// ID of the default payment source for the invoice. It must belong to the customer associated with the invoice and be in a chargeable state. If not set, defaults to the subscription's default source, if any, or to the customer's default source.
	DefaultSource *PaymentSource `json:"default_source"`
	// The tax rates applied to this invoice, if any.
	DefaultTaxRates []*TaxRate `json:"default_tax_rates"`
	Deleted         bool       `json:"deleted"`
	// An arbitrary string attached to the object. Often useful for displaying to users. Referenced as 'memo' in the Dashboard.
	Description string `json:"description"`
	// Describes the current discount applied to this invoice, if there is one. Not populated if there are multiple discounts.
	Discount *Discount `json:"discount"`
	// The discounts applied to the invoice. Line item discounts are applied before invoice discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*Discount `json:"discounts"`
	// The date on which payment for this invoice is due. This value will be `null` for invoices where `collection_method=charge_automatically`.
	DueDate int64 `json:"due_date"`
	// Ending customer balance after the invoice is finalized. Invoices are finalized approximately an hour after successful webhook delivery or when payment collection is attempted for the invoice. If the invoice has not been finalized yet, this will be null.
	EndingBalance int64 `json:"ending_balance"`
	// Footer displayed on the invoice.
	Footer string `json:"footer"`
	// The URL for the hosted invoice page, which allows customers to view and pay an invoice. If the invoice has not been finalized yet, this will be null.
	HostedInvoiceURL string `json:"hosted_invoice_url"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The link to download the PDF for the invoice. If the invoice has not been finalized yet, this will be null.
	InvoicePDF string `json:"invoice_pdf"`
	// The error encountered during the previous attempt to finalize the invoice. This field is cleared when the invoice is successfully finalized.
	LastFinalizationError *Error `json:"last_finalization_error"`
	// The individual line items that make up the invoice. `lines` is sorted as follows: invoice items in reverse chronological order, followed by the subscription, if any.
	Lines *InvoiceLineList `json:"lines"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The time at which payment will next be attempted. This value will be `null` for invoices where `collection_method=send_invoice`.
	NextPaymentAttempt int64 `json:"next_payment_attempt"`
	// A unique, identifying string that appears on emails sent to the customer for this invoice. This starts with the customer's unique invoice_prefix if it is specified.
	Number string `json:"number"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account (if any) for which the funds of the invoice payment are intended. If set, the invoice will be presented with the branding and support information of the specified account. See the [Invoices with Connect](https://stripe.com/docs/billing/invoices/connect) documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// Whether payment was successfully collected for this invoice. An invoice can be paid (most commonly) with a charge or with credit from the customer's account balance.
	Paid bool `json:"paid"`
	// Returns true if the invoice was manually marked paid, returns false if the invoice hasn't been paid yet or was paid on Stripe.
	PaidOutOfBand bool `json:"paid_out_of_band"`
	// The PaymentIntent associated with this invoice. The PaymentIntent is generated when the invoice is finalized, and can then be used to pay the invoice. Note that voiding an invoice will cancel the PaymentIntent.
	PaymentIntent   *PaymentIntent          `json:"payment_intent"`
	PaymentSettings *InvoicePaymentSettings `json:"payment_settings"`
	// End of the usage period during which invoice items were added to this invoice.
	PeriodEnd int64 `json:"period_end"`
	// Start of the usage period during which invoice items were added to this invoice.
	PeriodStart int64 `json:"period_start"`
	// Total amount of all post-payment credit notes issued for this invoice.
	PostPaymentCreditNotesAmount int64 `json:"post_payment_credit_notes_amount"`
	// Total amount of all pre-payment credit notes issued for this invoice.
	PrePaymentCreditNotesAmount int64 `json:"pre_payment_credit_notes_amount"`
	// The quote this invoice was generated from.
	Quote *Quote `json:"quote"`
	// This is the transaction number that appears on email receipts sent for this invoice.
	ReceiptNumber string `json:"receipt_number"`
	// Options for invoice PDF rendering.
	RenderingOptions *InvoiceRenderingOptions `json:"rendering_options"`
	// Starting customer balance before the invoice is finalized. If the invoice has not been finalized yet, this will be the current customer balance.
	StartingBalance int64 `json:"starting_balance"`
	// Extra information about an invoice for the customer's credit card statement.
	StatementDescriptor string `json:"statement_descriptor"`
	// The status of the invoice, one of `draft`, `open`, `paid`, `uncollectible`, or `void`. [Learn more](https://stripe.com/docs/billing/invoices/workflow#workflow-overview)
	Status            InvoiceStatus            `json:"status"`
	StatusTransitions InvoiceStatusTransitions `json:"status_transitions"`
	// The subscription that this invoice was prepared for, if any.
	Subscription *Subscription `json:"subscription"`
	// Only set for upcoming invoices that preview prorations. The time used to calculate prorations.
	SubscriptionProrationDate int64 `json:"subscription_proration_date"`
	// Total of all subscriptions, invoice items, and prorations on the invoice before any invoice level discount or exclusive tax is applied. Item discounts are already incorporated
	Subtotal int64 `json:"subtotal"`
	// The integer amount in %s representing the subtotal of the invoice before any invoice level discount or tax is applied. Item discounts are already incorporated
	SubtotalExcludingTax int64 `json:"subtotal_excluding_tax"`
	// The amount of tax on this invoice. This is the sum of all the tax amounts on this invoice.
	Tax int64 `json:"tax"`
	// ID of the test clock this invoice belongs to.
	TestClock        *TestHelpersTestClock   `json:"test_clock"`
	ThreasholdReason *InvoiceThresholdReason `json:"threshold_reason"`
	// Total after discounts and taxes.
	Total int64 `json:"total"`
	// The aggregate amounts calculated per discount across all line items.
	TotalDiscountAmounts []*InvoiceDiscountAmount `json:"total_discount_amounts"`
	// The integer amount in %s representing the total amount of the invoice including all discounts but excluding all tax.
	TotalExcludingTax int64 `json:"total_excluding_tax"`
	// The aggregate amounts calculated per tax rate for all line items.
	TotalTaxAmounts []*InvoiceTaxAmount `json:"total_tax_amounts"`
	// The account (if any) the payment will be attributed to for tax reporting, and where funds from the payment will be transferred to for the invoice.
	TransferData *InvoiceTransferData `json:"transfer_data"`
	// Invoices are automatically paid or sent 1 hour after webhooks are delivered, or until all webhook delivery attempts have [been exhausted](https://stripe.com/docs/billing/webhooks#understand). This field tracks the time when webhooks for this invoice were successfully delivered. If the invoice had no webhooks to deliver, this will be set while the invoice is being created.
	WebhooksDeliveredAt int64 `json:"webhooks_delivered_at"`
}

// InvoiceList is a list of Invoices as retrieved from a list endpoint.
type InvoiceList struct {
	APIResource
	ListMeta
	Data []*Invoice `json:"data"`
}

// InvoiceSearchResult is a list of Invoice search results as retrieved from a search endpoint.
type InvoiceSearchResult struct {
	APIResource
	SearchMeta
	Data []*Invoice `json:"data"`
}

// UnmarshalJSON handles deserialization of an Invoice.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *Invoice) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type invoice Invoice
	var v invoice
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = Invoice(v)
	return nil
}
