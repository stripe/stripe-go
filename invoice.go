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
	InvoicePaymentSettingsPaymentMethodTypeFPX                InvoicePaymentSettingsPaymentMethodType = "fpx"
	InvoicePaymentSettingsPaymentMethodTypeGiropay            InvoicePaymentSettingsPaymentMethodType = "giropay"
	InvoicePaymentSettingsPaymentMethodTypeIdeal              InvoicePaymentSettingsPaymentMethodType = "ideal"
	InvoicePaymentSettingsPaymentMethodTypeSepaCreditTransfer InvoicePaymentSettingsPaymentMethodType = "sepa_credit_transfer"
	InvoicePaymentSettingsPaymentMethodTypeSepaDebit          InvoicePaymentSettingsPaymentMethodType = "sepa_debit"
	InvoicePaymentSettingsPaymentMethodTypeSofort             InvoicePaymentSettingsPaymentMethodType = "sofort"
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

// You can list all invoices, or list the invoices for a specific customer. The invoices are returned sorted by creation date, with the most recently created invoices appearing first.
type InvoiceListParams struct {
	ListParams       `form:"*"`
	CollectionMethod *string           `form:"collection_method"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	Customer         *string           `form:"customer"`
	DueDate          *int64            `form:"due_date"`
	DueDateRange     *RangeQueryParams `form:"due_date"`
	Status           *string           `form:"status"`
	Subscription     *string           `form:"subscription"`
}

// Settings for automatic tax lookup for this invoice.
type InvoiceAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// A list of up to 4 custom fields to be displayed on the invoice.
type InvoiceCustomFieldParams struct {
	Name  *string `form:"name"`
	Value *string `form:"value"`
}

// The coupons to redeem into discounts for the invoice. If not specified, inherits the discount from the invoice's customer. Pass an empty string to avoid inheriting any discounts.
type InvoiceDiscountParams struct {
	Coupon   *string `form:"coupon"`
	Discount *string `form:"discount"`
}

// Additional fields for Mandate creation
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	TransactionType *string `form:"transaction_type"`
}

// If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitParams struct {
	MandateOptions     *InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                                  `form:"verification_method"`
}

// If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsBancontactParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}

// If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsCardParams struct {
	RequestThreeDSecure *string `form:"request_three_d_secure"`
}

// Payment-method-specific configuration to provide to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsParams struct {
	ACSSDebit  *InvoicePaymentSettingsPaymentMethodOptionsACSSDebitParams  `form:"acss_debit"`
	Bancontact *InvoicePaymentSettingsPaymentMethodOptionsBancontactParams `form:"bancontact"`
	Card       *InvoicePaymentSettingsPaymentMethodOptionsCardParams       `form:"card"`
}

// Configuration settings for the PaymentIntent that is generated when the invoice is finalized.
type InvoicePaymentSettingsParams struct {
	PaymentMethodOptions *InvoicePaymentSettingsPaymentMethodOptionsParams `form:"payment_method_options"`
	PaymentMethodTypes   []*string                                         `form:"payment_method_types"`
}

// If specified, the funds from the invoice will be transferred to the destination and the ID of the resulting transfer will be found on the invoice's charge.
type InvoiceTransferDataParams struct {
	Amount      *int64  `form:"amount"`
	Destination *string `form:"destination"`
}

// This endpoint creates a draft invoice for a given customer. The draft invoice created pulls in all pending invoice items on that customer, including prorations. The invoice remains a draft until you [finalize the invoice, which allows you to [pay](#pay_invoice) or <a href="#send_invoice">send](https://stripe.com/docs/api#finalize_invoice) the invoice to your customers.
type InvoiceParams struct {
	Params               `form:"*"`
	AccountTaxIDs        []*string                     `form:"account_tax_ids"`
	ApplicationFeeAmount *int64                        `form:"application_fee_amount"`
	AutoAdvance          *bool                         `form:"auto_advance"`
	AutomaticTax         *InvoiceAutomaticTaxParams    `form:"automatic_tax"`
	CollectionMethod     *string                       `form:"collection_method"`
	Customer             *string                       `form:"customer"`
	CustomFields         []*InvoiceCustomFieldParams   `form:"custom_fields"`
	DaysUntilDue         *int64                        `form:"days_until_due"`
	DefaultPaymentMethod *string                       `form:"default_payment_method"`
	DefaultSource        *string                       `form:"default_source"`
	DefaultTaxRates      []*string                     `form:"default_tax_rates"`
	Description          *string                       `form:"description"`
	Discounts            []*InvoiceDiscountParams      `form:"discounts"`
	DueDate              *int64                        `form:"due_date"`
	Footer               *string                       `form:"footer"`
	OnBehalfOf           *string                       `form:"on_behalf_of"`
	Paid                 *bool                         `form:"paid"`
	PaymentSettings      *InvoicePaymentSettingsParams `form:"payment_settings"`
	Schedule             *string                       `form:"schedule"`
	StatementDescriptor  *string                       `form:"statement_descriptor"`
	Subscription         *string                       `form:"subscription"`
	TransferData         *InvoiceTransferDataParams    `form:"transfer_data"`
	// These are all for exclusive use by GetNext.
	Coupon                                  *string                               `form:"coupon"`
	CustomerDetails                         *InvoiceUpcomingCustomerDetailsParams `form:"customer_details"`
	InvoiceItems                            []*InvoiceUpcomingInvoiceItemParams   `form:"invoice_items"`
	SubscriptionBillingCycleAnchor          *int64                                `form:"subscription_billing_cycle_anchor"`
	SubscriptionBillingCycleAnchorNow       *bool                                 `form:"-"` // See custom AppendTo
	SubscriptionBillingCycleAnchorUnchanged *bool                                 `form:"-"` // See custom AppendTo
	SubscriptionCancelAt                    *int64                                `form:"subscription_cancel_at"`
	SubscriptionCancelAtPeriodEnd           *bool                                 `form:"subscription_cancel_at_period_end"`
	SubscriptionCancelNow                   *bool                                 `form:"subscription_cancel_now"`
	SubscriptionDefaultTaxRates             []*string                             `form:"subscription_default_tax_rates"`
	SubscriptionItems                       []*SubscriptionItemsParams            `form:"subscription_items"`
	SubscriptionPlan                        *string                               `form:"subscription_plan"`
	SubscriptionProrationBehavior           *string                               `form:"subscription_proration_behavior"`
	SubscriptionProrationDate               *int64                                `form:"subscription_proration_date"`
	SubscriptionQuantity                    *int64                                `form:"subscription_quantity"`
	SubscriptionStartDate                   *int64                                `form:"subscription_start_date"`
	SubscriptionTrialEnd                    *int64                                `form:"subscription_trial_end"`
	SubscriptionTrialEndNow                 *bool                                 `form:"-"` // See custom AppendTo
	SubscriptionTrialFromPlan               *bool                                 `form:"subscription_trial_from_plan"`
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
	Address *AddressParams `form:"address"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// Tax details about the customer.
type InvoiceUpcomingCustomerDetailsTaxParams struct {
	IPAddress *string `form:"ip_address"`
}

// The customer's tax IDs.
type InvoiceUpcomingCustomerDetailsTaxIDParams struct {
	Type  *string `form:"type"`
	Value *string `form:"value"`
}

// Details about the customer you want to invoice or overrides for an existing customer.
type InvoiceUpcomingCustomerDetailsParams struct {
	Address   *AddressParams                                `form:"address"`
	Shipping  *InvoiceUpcomingCustomerDetailsShippingParams `form:"shipping"`
	Tax       *InvoiceUpcomingCustomerDetailsTaxParams      `form:"tax"`
	TaxExempt *string                                       `form:"tax_exempt"`
	TaxIDs    []*InvoiceUpcomingCustomerDetailsTaxIDParams  `form:"tax_ids"`
}

// The period associated with this invoice item.
type InvoiceUpcomingInvoiceItemPeriodParams struct {
	End   *int64 `form:"end"`
	Start *int64 `form:"start"`
}

// List of invoice items to add or update in the upcoming invoice preview.
type InvoiceUpcomingInvoiceItemParams struct {
	Amount            *int64                                  `form:"amount"`
	Currency          *string                                 `form:"currency"`
	Description       *string                                 `form:"description"`
	Discountable      *bool                                   `form:"discountable"`
	Discounts         []*InvoiceItemDiscountParams            `form:"discounts"`
	InvoiceItem       *string                                 `form:"invoiceitem"`
	Metadata          map[string]string                       `form:"metadata"`
	Period            *InvoiceUpcomingInvoiceItemPeriodParams `form:"period"`
	Price             *string                                 `form:"price"`
	PriceData         *InvoiceItemPriceDataParams             `form:"price_data"`
	Quantity          *int64                                  `form:"quantity"`
	Schedule          *string                                 `form:"schedule"`
	TaxRates          []*string                               `form:"tax_rates"`
	UnitAmount        *int64                                  `form:"unit_amount"`
	UnitAmountDecimal *float64                                `form:"unit_amount_decimal,high_precision"`
}

// Stripe automatically creates and then attempts to collect payment on invoices for customers on subscriptions according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to attempt payment on an invoice out of the normal collection schedule or for some other reason, you can do so.
type InvoicePayParams struct {
	Params        `form:"*"`
	Forgive       *bool   `form:"forgive"`
	OffSession    *bool   `form:"off_session"`
	PaidOutOfBand *bool   `form:"paid_out_of_band"`
	PaymentMethod *string `form:"payment_method"`
	Source        *string `form:"source"`
}

// Stripe automatically finalizes drafts before sending and attempting payment on invoices. However, if you'd like to finalize a draft invoice manually, you can do so using this method.
type InvoiceFinalizeParams struct {
	Params      `form:"*"`
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
	Enabled bool                      `json:"enabled"`
	Status  InvoiceAutomaticTaxStatus `json:"status"`
}

// Custom fields displayed on the invoice.
type InvoiceCustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// The customer's tax IDs. Until the invoice is finalized, this field will contain the same tax IDs as `customer.tax_ids`. Once the invoice is finalized, this field will no longer be updated.
type InvoiceCustomerTaxID struct {
	Type  TaxIDType `json:"type"`
	Value string    `json:"value"`
}
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions struct {
	TransactionType InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}

// If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebit struct {
	MandateOptions     *InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions    `json:"mandate_options"`
	VerificationMethod InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsBancontact struct {
	PreferredLanguage string `json:"preferred_language"`
}

// If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsCard struct {
	RequestThreeDSecure InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}

// Payment-method-specific configuration to provide to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptions struct {
	ACSSDebit  *InvoicePaymentSettingsPaymentMethodOptionsACSSDebit  `json:"acss_debit"`
	Bancontact *InvoicePaymentSettingsPaymentMethodOptionsBancontact `json:"bancontact"`
	Card       *InvoicePaymentSettingsPaymentMethodOptionsCard       `json:"card"`
}
type InvoicePaymentSettings struct {
	PaymentMethodOptions *InvoicePaymentSettingsPaymentMethodOptions `json:"payment_method_options"`
	PaymentMethodTypes   []InvoicePaymentSettingsPaymentMethodType   `json:"payment_method_types"`
}
type InvoiceStatusTransitions struct {
	FinalizedAt           int64 `json:"finalized_at"`
	MarkedUncollectibleAt int64 `json:"marked_uncollectible_at"`
	PaidAt                int64 `json:"paid_at"`
	VoidedAt              int64 `json:"voided_at"`
}

// Indicates which line items triggered a threshold invoice.
type InvoiceThresholdReasonItemReason struct {
	LineItemIDs []string `json:"line_item_ids"`
	UsageGTE    int64    `json:"usage_gte"`
}
type InvoiceThresholdReason struct {
	AmountGTE   int64                               `json:"amount_gte"`
	ItemReasons []*InvoiceThresholdReasonItemReason `json:"item_reasons"`
}

// The aggregate amounts calculated per discount across all line items.
type InvoiceDiscountAmount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// The aggregate amounts calculated per tax rate for all line items.
type InvoiceTaxAmount struct {
	Amount    int64    `json:"amount"`
	Inclusive bool     `json:"inclusive"`
	TaxRate   *TaxRate `json:"tax_rate"`
}

// The account (if any) the payment will be attributed to for tax reporting, and where funds from the payment will be transferred to for the invoice.
type InvoiceTransferData struct {
	Amount      int64    `json:"amount"`
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
	AccountCountry               string                   `json:"account_country"`
	AccountName                  string                   `json:"account_name"`
	AccountTaxIDs                []*TaxID                 `json:"account_tax_ids"`
	AmountDue                    int64                    `json:"amount_due"`
	AmountPaid                   int64                    `json:"amount_paid"`
	AmountRemaining              int64                    `json:"amount_remaining"`
	ApplicationFeeAmount         int64                    `json:"application_fee_amount"`
	AttemptCount                 int64                    `json:"attempt_count"`
	Attempted                    bool                     `json:"attempted"`
	AutoAdvance                  bool                     `json:"auto_advance"`
	AutomaticTax                 *InvoiceAutomaticTax     `json:"automatic_tax"`
	BillingReason                InvoiceBillingReason     `json:"billing_reason"`
	Charge                       *Charge                  `json:"charge"`
	CollectionMethod             *InvoiceCollectionMethod `json:"collection_method"`
	Created                      int64                    `json:"created"`
	Currency                     Currency                 `json:"currency"`
	Customer                     *Customer                `json:"customer"`
	CustomerAddress              *Address                 `json:"customer_address"`
	CustomerEmail                string                   `json:"customer_email"`
	CustomerName                 *string                  `json:"customer_name"`
	CustomerPhone                *string                  `json:"customer_phone"`
	CustomerShipping             *CustomerShippingDetails `json:"customer_shipping"`
	CustomerTaxExempt            CustomerTaxExempt        `json:"customer_tax_exempt"`
	CustomerTaxIDs               []*InvoiceCustomerTaxID  `json:"customer_tax_ids"`
	CustomFields                 []*InvoiceCustomField    `json:"custom_fields"`
	DefaultPaymentMethod         *PaymentMethod           `json:"default_payment_method"`
	DefaultSource                *PaymentSource           `json:"default_source"`
	DefaultTaxRates              []*TaxRate               `json:"default_tax_rates"`
	Deleted                      bool                     `json:"deleted"`
	Description                  string                   `json:"description"`
	Discount                     *Discount                `json:"discount"`
	Discounts                    []*Discount              `json:"discounts"`
	DueDate                      int64                    `json:"due_date"`
	EndingBalance                int64                    `json:"ending_balance"`
	Footer                       string                   `json:"footer"`
	HostedInvoiceURL             string                   `json:"hosted_invoice_url"`
	ID                           string                   `json:"id"`
	InvoicePDF                   string                   `json:"invoice_pdf"`
	LastFinalizationError        *Error                   `json:"last_finalization_error"`
	Lines                        *InvoiceLineList         `json:"lines"`
	Livemode                     bool                     `json:"livemode"`
	Metadata                     map[string]string        `json:"metadata"`
	NextPaymentAttempt           int64                    `json:"next_payment_attempt"`
	Number                       string                   `json:"number"`
	Object                       string                   `json:"object"`
	OnBehalfOf                   *Account                 `json:"on_behalf_of"`
	Paid                         bool                     `json:"paid"`
	PaymentIntent                *PaymentIntent           `json:"payment_intent"`
	PaymentSettings              *InvoicePaymentSettings  `json:"payment_settings"`
	PeriodEnd                    int64                    `json:"period_end"`
	PeriodStart                  int64                    `json:"period_start"`
	PostPaymentCreditNotesAmount int64                    `json:"post_payment_credit_notes_amount"`
	PrePaymentCreditNotesAmount  int64                    `json:"pre_payment_credit_notes_amount"`
	Quote                        *Quote                   `json:"quote"`
	ReceiptNumber                string                   `json:"receipt_number"`
	StartingBalance              int64                    `json:"starting_balance"`
	StatementDescriptor          string                   `json:"statement_descriptor"`
	Status                       InvoiceStatus            `json:"status"`
	StatusTransitions            InvoiceStatusTransitions `json:"status_transitions"`
	Subscription                 *Subscription            `json:"subscription"`
	SubscriptionProrationDate    int64                    `json:"subscription_proration_date"`
	Subtotal                     int64                    `json:"subtotal"`
	Tax                          int64                    `json:"tax"`
	ThreasholdReason             *InvoiceThresholdReason  `json:"threshold_reason"`
	Total                        int64                    `json:"total"`
	TotalDiscountAmounts         []*InvoiceDiscountAmount `json:"total_discount_amounts"`
	TotalTaxAmounts              []*InvoiceTaxAmount      `json:"total_tax_amounts"`
	TransferData                 *InvoiceTransferData     `json:"transfer_data"`
	WebhooksDeliveredAt          int64                    `json:"webhooks_delivered_at"`
}

// InvoiceList is a list of Invoices as retrieved from a list endpoint.
type InvoiceList struct {
	APIResource
	ListMeta
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
