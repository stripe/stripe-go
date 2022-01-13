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

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this subscription at the end of the cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions.
type SubscriptionCollectionMethod string

// List of values that SubscriptionCollectionMethod can take
const (
	SubscriptionCollectionMethodChargeAutomatically SubscriptionCollectionMethod = "charge_automatically"
	SubscriptionCollectionMethodSendInvoice         SubscriptionCollectionMethod = "send_invoice"
)

// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
type SubscriptionPauseCollectionBehavior string

// List of values that SubscriptionPauseCollectionBehavior can take
const (
	SubscriptionPauseCollectionBehaviorKeepAsDraft       SubscriptionPauseCollectionBehavior = "keep_as_draft"
	SubscriptionPauseCollectionBehaviorMarkUncollectible SubscriptionPauseCollectionBehavior = "mark_uncollectible"
	SubscriptionPauseCollectionBehaviorVoid              SubscriptionPauseCollectionBehavior = "void"
)

// Transaction type of the mandate.
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Bank account verification method.
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodInstant       SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
type SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsAmountType string

// List of values that SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsAmountType can take
const (
	SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsAmountTypeFixed   SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsAmountType = "fixed"
	SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsAmountType = "maximum"
)

// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecureAny       SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure = "any"
	SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecureAutomatic SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
)

// The list of payment method types to provide to every invoice created by the subscription. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice's default payment method, the subscription's default payment method, the customer's default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).
type SubscriptionPaymentSettingsPaymentMethodType string

// List of values that SubscriptionPaymentSettingsPaymentMethodType can take
const (
	SubscriptionPaymentSettingsPaymentMethodTypeAchCreditTransfer  SubscriptionPaymentSettingsPaymentMethodType = "ach_credit_transfer"
	SubscriptionPaymentSettingsPaymentMethodTypeAchDebit           SubscriptionPaymentSettingsPaymentMethodType = "ach_debit"
	SubscriptionPaymentSettingsPaymentMethodTypeACSSDebit          SubscriptionPaymentSettingsPaymentMethodType = "acss_debit"
	SubscriptionPaymentSettingsPaymentMethodTypeAUBECSDebit        SubscriptionPaymentSettingsPaymentMethodType = "au_becs_debit"
	SubscriptionPaymentSettingsPaymentMethodTypeBACSDebit          SubscriptionPaymentSettingsPaymentMethodType = "bacs_debit"
	SubscriptionPaymentSettingsPaymentMethodTypeBancontact         SubscriptionPaymentSettingsPaymentMethodType = "bancontact"
	SubscriptionPaymentSettingsPaymentMethodTypeBoleto             SubscriptionPaymentSettingsPaymentMethodType = "boleto"
	SubscriptionPaymentSettingsPaymentMethodTypeCard               SubscriptionPaymentSettingsPaymentMethodType = "card"
	SubscriptionPaymentSettingsPaymentMethodTypeFPX                SubscriptionPaymentSettingsPaymentMethodType = "fpx"
	SubscriptionPaymentSettingsPaymentMethodTypeGiropay            SubscriptionPaymentSettingsPaymentMethodType = "giropay"
	SubscriptionPaymentSettingsPaymentMethodTypeIdeal              SubscriptionPaymentSettingsPaymentMethodType = "ideal"
	SubscriptionPaymentSettingsPaymentMethodTypeSepaCreditTransfer SubscriptionPaymentSettingsPaymentMethodType = "sepa_credit_transfer"
	SubscriptionPaymentSettingsPaymentMethodTypeSepaDebit          SubscriptionPaymentSettingsPaymentMethodType = "sepa_debit"
	SubscriptionPaymentSettingsPaymentMethodTypeSofort             SubscriptionPaymentSettingsPaymentMethodType = "sofort"
	SubscriptionPaymentSettingsPaymentMethodTypeWechatPay          SubscriptionPaymentSettingsPaymentMethodType = "wechat_pay"
)

// SubscriptionPaymentBehavior lets you control the behavior of subscription creation in case of
// a failed payment.
type SubscriptionPaymentBehavior string

// List of values that SubscriptionPaymentBehavior can take.
const (
	SubscriptionPaymentBehaviorAllowIncomplete     SubscriptionPaymentBehavior = "allow_incomplete"
	SubscriptionPaymentBehaviorErrorIfIncomplete   SubscriptionPaymentBehavior = "error_if_incomplete"
	SubscriptionPaymentBehaviorPendingIfIncomplete SubscriptionPaymentBehavior = "pending_if_incomplete"
)

// SubscriptionProrationBehavior determines how to handle prorations when billing cycles change.
type SubscriptionProrationBehavior string

// List of values that SubscriptionProrationBehavior can take.
const (
	SubscriptionProrationBehaviorAlwaysInvoice    SubscriptionProrationBehavior = "always_invoice"
	SubscriptionProrationBehaviorCreateProrations SubscriptionProrationBehavior = "create_prorations"
	SubscriptionProrationBehaviorNone             SubscriptionProrationBehavior = "none"
)

// Specifies invoicing frequency. Either `day`, `week`, `month` or `year`.
type SubscriptionPendingInvoiceItemIntervalInterval string

// List of values that SubscriptionPendingInvoiceItemIntervalInterval can take
const (
	SubscriptionPendingInvoiceItemIntervalIntervalDay   SubscriptionPendingInvoiceItemIntervalInterval = "day"
	SubscriptionPendingInvoiceItemIntervalIntervalMonth SubscriptionPendingInvoiceItemIntervalInterval = "month"
	SubscriptionPendingInvoiceItemIntervalIntervalWeek  SubscriptionPendingInvoiceItemIntervalInterval = "week"
	SubscriptionPendingInvoiceItemIntervalIntervalYear  SubscriptionPendingInvoiceItemIntervalInterval = "year"
)

// Possible values are `incomplete`, `incomplete_expired`, `trialing`, `active`, `past_due`, `canceled`, or `unpaid`.
//
// For `collection_method=charge_automatically` a subscription moves into `incomplete` if the initial payment attempt fails. A subscription in this state can only have metadata and default_source updated. Once the first invoice is paid, the subscription moves into an `active` state. If the first invoice is not paid within 23 hours, the subscription transitions to `incomplete_expired`. This is a terminal state, the open invoice will be voided and no further invoices will be generated.
//
// A subscription that is currently in a trial period is `trialing` and moves to `active` when the trial period is over.
//
// If subscription `collection_method=charge_automatically` it becomes `past_due` when payment to renew it fails and `canceled` or `unpaid` (depending on your subscriptions settings) when Stripe has exhausted all payment retry attempts.
//
// If subscription `collection_method=send_invoice` it becomes `past_due` when its invoice is not paid by the due date, and `canceled` or `unpaid` if it is still not paid by an additional deadline after that. Note that when a subscription has a status of `unpaid`, no subsequent invoices will be attempted (invoices will be created, but then immediately automatically closed). After receiving updated payment information from a customer, you may choose to reopen and pay their closed invoices.
type SubscriptionStatus string

// List of values that SubscriptionStatus can take
const (
	SubscriptionStatusActive            SubscriptionStatus = "active"
	SubscriptionStatusAll               SubscriptionStatus = "all"
	SubscriptionStatusCanceled          SubscriptionStatus = "canceled"
	SubscriptionStatusIncomplete        SubscriptionStatus = "incomplete"
	SubscriptionStatusIncompleteExpired SubscriptionStatus = "incomplete_expired"
	SubscriptionStatusPastDue           SubscriptionStatus = "past_due"
	SubscriptionStatusTrialing          SubscriptionStatus = "trialing"
	SubscriptionStatusUnpaid            SubscriptionStatus = "unpaid"
)

// By default, returns a list of subscriptions that have not been canceled. In order to list canceled subscriptions, specify status=canceled.
type SubscriptionListParams struct {
	ListParams              `form:"*"`
	CollectionMethod        *string           `form:"collection_method"`
	Created                 int64             `form:"created"`
	CreatedRange            *RangeQueryParams `form:"created"`
	CurrentPeriodEnd        *int64            `form:"current_period_end"`
	CurrentPeriodEndRange   *RangeQueryParams `form:"current_period_end"`
	CurrentPeriodStart      *int64            `form:"current_period_start"`
	CurrentPeriodStartRange *RangeQueryParams `form:"current_period_start"`
	Customer                string            `form:"customer"`
	Plan                    string            `form:"plan"`
	Price                   string            `form:"price"`
	Status                  string            `form:"status"`
}

// A list of prices and quantities that will generate invoice items appended to the first invoice for this subscription. You may pass up to 20 items.
type SubscriptionAddInvoiceItemParams struct {
	Price     *string                     `form:"price"`
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	Quantity  *int64                      `form:"quantity"`
	TaxRates  []*string                   `form:"tax_rates"`
}

// Automatic tax settings for this subscription.
type SubscriptionAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionBillingThresholdsParams struct {
	AmountGTE               *int64 `form:"amount_gte"`
	ResetBillingCycleAnchor *bool  `form:"reset_billing_cycle_anchor"`
}

// A list of up to 20 subscription items, each with an attached price.
type SubscriptionItemsParams struct {
	Params            `form:"*"`
	BillingThresholds *SubscriptionItemBillingThresholdsParams `form:"billing_thresholds"`
	ClearUsage        *bool                                    `form:"clear_usage"`
	Deleted           *bool                                    `form:"deleted"`
	ID                *string                                  `form:"id"`
	Metadata          map[string]string                        `form:"metadata"`
	Plan              *string                                  `form:"plan"`
	Price             *string                                  `form:"price"`
	PriceData         *SubscriptionItemPriceDataParams         `form:"price_data"`
	Quantity          *int64                                   `form:"quantity"`
	TaxRates          []*string                                `form:"tax_rates"`
}

// Additional fields for Mandate creation
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	TransactionType *string `form:"transaction_type"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitParams struct {
	MandateOptions     *SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                                       `form:"verification_method"`
}

// This sub-hash contains details about the Bancontact payment method options to pass to the invoice's PaymentIntent.
type SubscriptionPaymentSettingsPaymentMethodOptionsBancontactParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsParams struct {
	Amount      *int64  `form:"amount"`
	AmountType  *string `form:"amount_type"`
	Description *string `form:"description"`
}

// This sub-hash contains details about the Card payment method options to pass to the invoice's PaymentIntent.
type SubscriptionPaymentSettingsPaymentMethodOptionsCardParams struct {
	MandateOptions      *SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options"`
	RequestThreeDSecure *string                                                                  `form:"request_three_d_secure"`
}

// Payment-method-specific configuration to provide to invoices created by the subscription.
type SubscriptionPaymentSettingsPaymentMethodOptionsParams struct {
	ACSSDebit  *SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitParams  `form:"acss_debit"`
	Bancontact *SubscriptionPaymentSettingsPaymentMethodOptionsBancontactParams `form:"bancontact"`
	Card       *SubscriptionPaymentSettingsPaymentMethodOptionsCardParams       `form:"card"`
}

// Payment settings to pass to invoices created by the subscription.
type SubscriptionPaymentSettingsParams struct {
	PaymentMethodOptions *SubscriptionPaymentSettingsPaymentMethodOptionsParams `form:"payment_method_options"`
	PaymentMethodTypes   []*string                                              `form:"payment_method_types"`
}

// Specifies an interval for how often to bill for any pending invoice items. It is analogous to calling [Create an invoice](https://stripe.com/docs/api#create_invoice) for the given subscription at the specified interval.
type SubscriptionPendingInvoiceItemIntervalParams struct {
	Interval      *string `form:"interval"`
	IntervalCount *int64  `form:"interval_count"`
}

// If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.
type SubscriptionTransferDataParams struct {
	AmountPercent *float64 `form:"amount_percent"`
	Destination   *string  `form:"destination"`
}

// Creates a new subscription on an existing customer. Each customer can have up to 500 active or scheduled subscriptions.
type SubscriptionParams struct {
	Params                      `form:"*"`
	AddInvoiceItems             []*SubscriptionAddInvoiceItemParams           `form:"add_invoice_items"`
	ApplicationFeePercent       *float64                                      `form:"application_fee_percent"`
	AutomaticTax                *SubscriptionAutomaticTaxParams               `form:"automatic_tax"`
	BackdateStartDate           *int64                                        `form:"backdate_start_date"`
	BillingCycleAnchor          *int64                                        `form:"billing_cycle_anchor"`
	BillingCycleAnchorNow       *bool                                         `form:"-"` // See custom AppendTo
	BillingCycleAnchorUnchanged *bool                                         `form:"-"` // See custom AppendTo
	BillingThresholds           *SubscriptionBillingThresholdsParams          `form:"billing_thresholds"`
	CancelAt                    *int64                                        `form:"cancel_at"`
	CancelAtPeriodEnd           *bool                                         `form:"cancel_at_period_end"`
	Card                        *CardParams                                   `form:"card"`
	CollectionMethod            *string                                       `form:"collection_method"`
	Coupon                      *string                                       `form:"coupon"`
	Customer                    *string                                       `form:"customer"`
	DaysUntilDue                *int64                                        `form:"days_until_due"`
	DefaultPaymentMethod        *string                                       `form:"default_payment_method"`
	DefaultSource               *string                                       `form:"default_source"`
	DefaultTaxRates             []*string                                     `form:"default_tax_rates"`
	Items                       []*SubscriptionItemsParams                    `form:"items"`
	OffSession                  *bool                                         `form:"off_session"`
	OnBehalfOf                  *string                                       `form:"on_behalf_of"`
	PauseCollection             *SubscriptionPauseCollectionParams            `form:"pause_collection"`
	PaymentBehavior             *string                                       `form:"payment_behavior"`
	PaymentSettings             *SubscriptionPaymentSettingsParams            `form:"payment_settings"`
	PendingInvoiceItemInterval  *SubscriptionPendingInvoiceItemIntervalParams `form:"pending_invoice_item_interval"`
	Plan                        *string                                       `form:"plan"`
	PromotionCode               *string                                       `form:"promotion_code"`
	ProrationBehavior           *string                                       `form:"proration_behavior"`
	ProrationDate               *int64                                        `form:"proration_date"`
	Quantity                    *int64                                        `form:"quantity"`
	TransferData                *SubscriptionTransferDataParams               `form:"transfer_data"`
	TrialEnd                    *int64                                        `form:"trial_end"`
	TrialEndNow                 *bool                                         `form:"-"` // See custom AppendTo
	TrialFromPlan               *bool                                         `form:"trial_from_plan"`
	TrialPeriodDays             *int64                                        `form:"trial_period_days"`
}

// AppendTo implements custom encoding logic for SubscriptionParams.
func (s *SubscriptionParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(s.BillingCycleAnchorNow) {
		body.Add(form.FormatKey(append(keyParts, "billing_cycle_anchor")), "now")
	}
	if BoolValue(s.BillingCycleAnchorUnchanged) {
		body.Add(form.FormatKey(append(keyParts, "billing_cycle_anchor")), "unchanged")
	}
	if BoolValue(s.TrialEndNow) {
		body.Add(form.FormatKey(append(keyParts, "trial_end")), "now")
	}
}

// If specified, payment collection for this subscription will be paused.
type SubscriptionPauseCollectionParams struct {
	Behavior  *string `form:"behavior"`
	ResumesAt *int64  `form:"resumes_at"`
}

// Cancels a customer's subscription immediately. The customer will not be charged again for the subscription.
//
// Note, however, that any pending invoice items that you've created will still be charged for at the end of the period, unless manually [deleted](https://stripe.com/docs/api#delete_invoiceitem). If you've set the subscription to cancel at the end of the period, any pending prorations will also be left in place and collected at the end of the period. But if the subscription is set to cancel immediately, pending prorations will be removed.
//
// By default, upon subscription cancellation, Stripe will stop automatic collection of all finalized invoices for the customer. This is intended to prevent unexpected payment attempts after the customer has canceled a subscription. However, you can resume automatic collection of the invoices manually after subscription cancellation to have us proceed. Or, you could check for unpaid invoices before allowing the customer to cancel the subscription at all.
type SubscriptionCancelParams struct {
	Params     `form:"*"`
	InvoiceNow *bool `form:"invoice_now"`
	Prorate    *bool `form:"prorate"`
}
type SubscriptionAutomaticTax struct {
	Enabled bool `json:"enabled"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type SubscriptionBillingThresholds struct {
	AmountGTE               int64 `json:"amount_gte"`
	ResetBillingCycleAnchor bool  `json:"reset_billing_cycle_anchor"`
}

// If specified, payment collection for this subscription will be paused.
type SubscriptionPauseCollection struct {
	Behavior  SubscriptionPauseCollectionBehavior `json:"behavior"`
	ResumesAt int64                               `json:"resumes_at"`
}
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions struct {
	TransactionType SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to invoices created by the subscription.
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebit struct {
	MandateOptions     *SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions    `json:"mandate_options"`
	VerificationMethod SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// This sub-hash contains details about the Bancontact payment method options to pass to invoices created by the subscription.
type SubscriptionPaymentSettingsPaymentMethodOptionsBancontact struct {
	PreferredLanguage string `json:"preferred_language"`
}
type SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptions struct {
	Amount      int64                                                                       `json:"amount"`
	AmountType  SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsAmountType `json:"amount_type"`
	Description string                                                                      `json:"description"`
}

// This sub-hash contains details about the Card payment method options to pass to invoices created by the subscription.
type SubscriptionPaymentSettingsPaymentMethodOptionsCard struct {
	MandateOptions      *SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptions     `json:"mandate_options"`
	RequestThreeDSecure SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}

// Payment-method-specific configuration to provide to invoices created by the subscription.
type SubscriptionPaymentSettingsPaymentMethodOptions struct {
	ACSSDebit  *SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebit  `json:"acss_debit"`
	Bancontact *SubscriptionPaymentSettingsPaymentMethodOptionsBancontact `json:"bancontact"`
	Card       *SubscriptionPaymentSettingsPaymentMethodOptionsCard       `json:"card"`
}

// Payment settings passed on to invoices created by the subscription.
type SubscriptionPaymentSettings struct {
	PaymentMethodOptions *SubscriptionPaymentSettingsPaymentMethodOptions `json:"payment_method_options"`
	PaymentMethodTypes   []SubscriptionPaymentSettingsPaymentMethodType   `json:"payment_method_types"`
}

// Specifies an interval for how often to bill for any pending invoice items. It is analogous to calling [Create an invoice](https://stripe.com/docs/api#create_invoice) for the given subscription at the specified interval.
type SubscriptionPendingInvoiceItemInterval struct {
	Interval      SubscriptionPendingInvoiceItemIntervalInterval `json:"interval"`
	IntervalCount int64                                          `json:"interval_count"`
}

// If specified, [pending updates](https://stripe.com/docs/billing/subscriptions/pending-updates) that will be applied to the subscription once the `latest_invoice` has been paid.
type SubscriptionPendingUpdate struct {
	BillingCycleAnchor int64               `json:"billing_cycle_anchor"`
	ExpiresAt          int64               `json:"expires_at"`
	SubscriptionItems  []*SubscriptionItem `json:"subscription_items"`
	TrialEnd           int64               `json:"trial_end"`
	TrialFromPlan      bool                `json:"trial_from_plan"`
}

// The account (if any) the subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
type SubscriptionTransferData struct {
	AmountPercent float64  `json:"amount_percent"`
	Destination   *Account `json:"destination"`
}

// Subscriptions allow you to charge a customer on a recurring basis.
//
// Related guide: [Creating Subscriptions](https://stripe.com/docs/billing/subscriptions/creating).
type Subscription struct {
	APIResource
	ApplicationFeePercent         float64                                `json:"application_fee_percent"`
	AutomaticTax                  *SubscriptionAutomaticTax              `json:"automatic_tax"`
	BillingCycleAnchor            int64                                  `json:"billing_cycle_anchor"`
	BillingThresholds             *SubscriptionBillingThresholds         `json:"billing_thresholds"`
	CancelAt                      int64                                  `json:"cancel_at"`
	CancelAtPeriodEnd             bool                                   `json:"cancel_at_period_end"`
	CanceledAt                    int64                                  `json:"canceled_at"`
	CollectionMethod              SubscriptionCollectionMethod           `json:"collection_method"`
	Created                       int64                                  `json:"created"`
	CurrentPeriodEnd              int64                                  `json:"current_period_end"`
	CurrentPeriodStart            int64                                  `json:"current_period_start"`
	Customer                      *Customer                              `json:"customer"`
	DaysUntilDue                  int64                                  `json:"days_until_due"`
	DefaultPaymentMethod          *PaymentMethod                         `json:"default_payment_method"`
	DefaultSource                 *PaymentSource                         `json:"default_source"`
	DefaultTaxRates               []*TaxRate                             `json:"default_tax_rates"`
	Discount                      *Discount                              `json:"discount"`
	EndedAt                       int64                                  `json:"ended_at"`
	ID                            string                                 `json:"id"`
	Items                         *SubscriptionItemList                  `json:"items"`
	LatestInvoice                 *Invoice                               `json:"latest_invoice"`
	Livemode                      bool                                   `json:"livemode"`
	Metadata                      map[string]string                      `json:"metadata"`
	NextPendingInvoiceItemInvoice int64                                  `json:"next_pending_invoice_item_invoice"`
	Object                        string                                 `json:"object"`
	OnBehalfOf                    *Account                               `json:"on_behalf_of"`
	PauseCollection               SubscriptionPauseCollection            `json:"pause_collection"`
	PaymentSettings               *SubscriptionPaymentSettings           `json:"payment_settings"`
	PendingInvoiceItemInterval    SubscriptionPendingInvoiceItemInterval `json:"pending_invoice_item_interval"`
	PendingSetupIntent            *SetupIntent                           `json:"pending_setup_intent"`
	PendingUpdate                 *SubscriptionPendingUpdate             `json:"pending_update"`
	Plan                          *Plan                                  `json:"plan"`
	Quantity                      int64                                  `json:"quantity"`
	Schedule                      *SubscriptionSchedule                  `json:"schedule"`
	StartDate                     int64                                  `json:"start_date"`
	Status                        SubscriptionStatus                     `json:"status"`
	TransferData                  *SubscriptionTransferData              `json:"transfer_data"`
	TrialEnd                      int64                                  `json:"trial_end"`
	TrialStart                    int64                                  `json:"trial_start"`
}

// SubscriptionList is a list of Subscriptions as retrieved from a list endpoint.
type SubscriptionList struct {
	APIResource
	ListMeta
	Data []*Subscription `json:"data"`
}

// UnmarshalJSON handles deserialization of a Subscription.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *Subscription) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type subscription Subscription
	var v subscription
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = Subscription(v)
	return nil
}
