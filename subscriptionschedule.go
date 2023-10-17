//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v76/form"
)

// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
type SubscriptionScheduleBillingBehavior string

// List of values that SubscriptionScheduleBillingBehavior can take
const (
	SubscriptionScheduleBillingBehaviorProrateOnNextPhase SubscriptionScheduleBillingBehavior = "prorate_on_next_phase"
	SubscriptionScheduleBillingBehaviorProrateUpFront     SubscriptionScheduleBillingBehavior = "prorate_up_front"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type SubscriptionScheduleDefaultSettingsBillingCycleAnchor string

// List of values that SubscriptionScheduleDefaultSettingsBillingCycleAnchor can take
const (
	SubscriptionScheduleDefaultSettingsBillingCycleAnchorAutomatic  SubscriptionScheduleDefaultSettingsBillingCycleAnchor = "automatic"
	SubscriptionScheduleDefaultSettingsBillingCycleAnchorPhaseStart SubscriptionScheduleDefaultSettingsBillingCycleAnchor = "phase_start"
)

// Type of the account referenced.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType string

// List of values that SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType can take
const (
	SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerTypeAccount SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType = "account"
	SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerTypeSelf    SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType = "self"
)

// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running.`cancel` will end the subscription schedule and cancel the underlying subscription.
type SubscriptionScheduleEndBehavior string

// List of values that SubscriptionScheduleEndBehavior can take
const (
	SubscriptionScheduleEndBehaviorCancel  SubscriptionScheduleEndBehavior = "cancel"
	SubscriptionScheduleEndBehaviorNone    SubscriptionScheduleEndBehavior = "none"
	SubscriptionScheduleEndBehaviorRelease SubscriptionScheduleEndBehavior = "release"
	SubscriptionScheduleEndBehaviorRenew   SubscriptionScheduleEndBehavior = "renew"
)

// The discount end type.
type SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType string

// List of values that SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType can take
const (
	SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndTypeTimestamp SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType = "timestamp"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type SubscriptionSchedulePhaseBillingCycleAnchor string

// List of values that SubscriptionSchedulePhaseBillingCycleAnchor can take
const (
	SubscriptionSchedulePhaseBillingCycleAnchorAutomatic  SubscriptionSchedulePhaseBillingCycleAnchor = "automatic"
	SubscriptionSchedulePhaseBillingCycleAnchorPhaseStart SubscriptionSchedulePhaseBillingCycleAnchor = "phase_start"
)

// The discount end type.
type SubscriptionSchedulePhaseDiscountDiscountEndType string

// List of values that SubscriptionSchedulePhaseDiscountDiscountEndType can take
const (
	SubscriptionSchedulePhaseDiscountDiscountEndTypeTimestamp SubscriptionSchedulePhaseDiscountDiscountEndType = "timestamp"
)

// Type of the account referenced.
type SubscriptionSchedulePhaseInvoiceSettingsIssuerType string

// List of values that SubscriptionSchedulePhaseInvoiceSettingsIssuerType can take
const (
	SubscriptionSchedulePhaseInvoiceSettingsIssuerTypeAccount SubscriptionSchedulePhaseInvoiceSettingsIssuerType = "account"
	SubscriptionSchedulePhaseInvoiceSettingsIssuerTypeSelf    SubscriptionSchedulePhaseInvoiceSettingsIssuerType = "self"
)

// The discount end type.
type SubscriptionSchedulePhaseItemDiscountDiscountEndType string

// List of values that SubscriptionSchedulePhaseItemDiscountDiscountEndType can take
const (
	SubscriptionSchedulePhaseItemDiscountDiscountEndTypeTimestamp SubscriptionSchedulePhaseItemDiscountDiscountEndType = "timestamp"
)

type SubscriptionSchedulePhaseItemTrialType string

// List of values that SubscriptionSchedulePhaseItemTrialType can take
const (
	SubscriptionSchedulePhaseItemTrialTypeFree SubscriptionSchedulePhaseItemTrialType = "free"
	SubscriptionSchedulePhaseItemTrialTypePaid SubscriptionSchedulePhaseItemTrialType = "paid"
)

// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
type SubscriptionSchedulePhasePauseCollectionBehavior string

// List of values that SubscriptionSchedulePhasePauseCollectionBehavior can take
const (
	SubscriptionSchedulePhasePauseCollectionBehaviorKeepAsDraft       SubscriptionSchedulePhasePauseCollectionBehavior = "keep_as_draft"
	SubscriptionSchedulePhasePauseCollectionBehaviorMarkUncollectible SubscriptionSchedulePhasePauseCollectionBehavior = "mark_uncollectible"
	SubscriptionSchedulePhasePauseCollectionBehaviorVoid              SubscriptionSchedulePhasePauseCollectionBehavior = "void"
)

// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
type SubscriptionSchedulePhaseProrationBehavior string

// List of values that SubscriptionSchedulePhaseProrationBehavior can take
const (
	SubscriptionSchedulePhaseProrationBehaviorAlwaysInvoice    SubscriptionSchedulePhaseProrationBehavior = "always_invoice"
	SubscriptionSchedulePhaseProrationBehaviorCreateProrations SubscriptionSchedulePhaseProrationBehavior = "create_prorations"
	SubscriptionSchedulePhaseProrationBehaviorNone             SubscriptionSchedulePhaseProrationBehavior = "none"
)

// Specify behavior of the trial when crossing schedule phase boundaries
type SubscriptionSchedulePhaseTrialContinuation string

// List of values that SubscriptionSchedulePhaseTrialContinuation can take
const (
	SubscriptionSchedulePhaseTrialContinuationContinue SubscriptionSchedulePhaseTrialContinuation = "continue"
	SubscriptionSchedulePhaseTrialContinuationNone     SubscriptionSchedulePhaseTrialContinuation = "none"
)

// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
type SubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront string

// List of values that SubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront can take
const (
	SubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFrontDefer   SubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront = "defer"
	SubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFrontInclude SubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront = "include"
)

// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period.
type SubscriptionSchedulePrebillingUpdateBehavior string

// List of values that SubscriptionSchedulePrebillingUpdateBehavior can take
const (
	SubscriptionSchedulePrebillingUpdateBehaviorPrebill SubscriptionSchedulePrebillingUpdateBehavior = "prebill"
	SubscriptionSchedulePrebillingUpdateBehaviorReset   SubscriptionSchedulePrebillingUpdateBehavior = "reset"
)

// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
type SubscriptionScheduleStatus string

// List of values that SubscriptionScheduleStatus can take
const (
	SubscriptionScheduleStatusActive     SubscriptionScheduleStatus = "active"
	SubscriptionScheduleStatusCanceled   SubscriptionScheduleStatus = "canceled"
	SubscriptionScheduleStatusCompleted  SubscriptionScheduleStatus = "completed"
	SubscriptionScheduleStatusNotStarted SubscriptionScheduleStatus = "not_started"
	SubscriptionScheduleStatusReleased   SubscriptionScheduleStatus = "released"
)

// Retrieves the list of your subscription schedules.
type SubscriptionScheduleListParams struct {
	ListParams `form:"*"`
	// Only return subscription schedules that were created canceled the given date interval.
	CanceledAt *int64 `form:"canceled_at"`
	// Only return subscription schedules that were created canceled the given date interval.
	CanceledAtRange *RangeQueryParams `form:"canceled_at"`
	// Only return subscription schedules that completed during the given date interval.
	CompletedAt *int64 `form:"completed_at"`
	// Only return subscription schedules that completed during the given date interval.
	CompletedAtRange *RangeQueryParams `form:"completed_at"`
	// Only return subscription schedules that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return subscription schedules that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return subscription schedules for the given customer.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return subscription schedules that were released during the given date interval.
	ReleasedAt *int64 `form:"released_at"`
	// Only return subscription schedules that were released during the given date interval.
	ReleasedAtRange *RangeQueryParams `form:"released_at"`
	// Only return subscription schedules that have not started yet.
	Scheduled *bool `form:"scheduled"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsParams struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `collection_method=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerParams `form:"issuer"`
}

// Object representing the subscription schedule's default settings.
type SubscriptionScheduleDefaultSettingsParams struct {
	Params `form:"*"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent,high_precision"`
	// Default settings for automatic tax computation.
	AutomaticTax *SubscriptionAutomaticTaxParams `form:"automatic_tax"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionBillingThresholdsParams `form:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description *string `form:"description"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleDefaultSettingsInvoiceSettingsParams `form:"invoice_settings"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of"`
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data"`
}

// Time span for the redeemed discount.
type SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The coupons to redeem into discounts for the item.
type SubscriptionSchedulePhaseAddInvoiceItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndParams `form:"discount_end"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
type SubscriptionSchedulePhaseAddInvoiceItemParams struct {
	// The coupons to redeem into discounts for the item.
	Discounts []*SubscriptionSchedulePhaseAddInvoiceItemDiscountParams `form:"discounts"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	// Quantity for this item. Defaults to 1.
	Quantity *int64 `form:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*string `form:"tax_rates"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type SubscriptionSchedulePhaseAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// Automatic tax settings for this phase.
type SubscriptionSchedulePhaseAutomaticTaxParams struct {
	// Enabled automatic tax calculation which will automatically compute tax rates on all invoices generated by the subscription.
	Enabled *bool `form:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *SubscriptionSchedulePhaseAutomaticTaxLiabilityParams `form:"liability"`
}

// Time span for the redeemed discount.
type SubscriptionSchedulePhaseDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionSchedulePhaseDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionSchedulePhaseDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
type SubscriptionSchedulePhaseDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionSchedulePhaseDiscountDiscountEndParams `form:"discount_end"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionSchedulePhaseInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionSchedulePhaseInvoiceSettingsParams struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionSchedulePhaseInvoiceSettingsIssuerParams `form:"issuer"`
}

// Time span for the redeemed discount.
type SubscriptionSchedulePhaseItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionSchedulePhaseItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionSchedulePhaseItemDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The coupons to redeem into discounts for the subscription item.
type SubscriptionSchedulePhaseItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionSchedulePhaseItemDiscountDiscountEndParams `form:"discount_end"`
}

// Options that configure the trial on the subscription item.
type SubscriptionSchedulePhaseItemTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []*string `form:"converts_to"`
	// Determines the type of trial for this item.
	Type *string `form:"type"`
}

// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
type SubscriptionSchedulePhaseItemParams struct {
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. When updating, pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionItemBillingThresholdsParams `form:"billing_thresholds"`
	// The coupons to redeem into discounts for the subscription item.
	Discounts []*SubscriptionSchedulePhaseItemDiscountParams `form:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a configuration item. Metadata on a configuration item will update the underlying subscription item's `metadata` when the phase is entered, adding new keys and replacing existing keys. Individual keys in the subscription item's `metadata` can be unset by posting an empty value to them in the configuration item's `metadata`. To unset all keys in the subscription item's `metadata`, update the subscription item directly or unset every key individually from the configuration item's `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The plan ID to subscribe to. You may specify the same ID in `plan` and `price`.
	Plan *string `form:"plan"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
	PriceData *SubscriptionItemPriceDataParams `form:"price_data"`
	// Quantity for the given price. Can be set only if the price's `usage_type` is `licensed` and not `metered`.
	Quantity *int64 `form:"quantity"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionSchedulePhaseItemTrialParams `form:"trial"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionSchedulePhaseItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// If specified, payment collection for this subscription will be paused.
type SubscriptionSchedulePhasePauseCollectionParams struct {
	// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
	Behavior *string `form:"behavior"`
}

// Defines how the subscription should behave when a trial ends.
type SubscriptionSchedulePhaseTrialSettingsEndBehaviorParams struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront *string `form:"prorate_up_front"`
}

// Settings related to subscription trials.
type SubscriptionSchedulePhaseTrialSettingsParams struct {
	// Defines how the subscription should behave when a trial ends.
	EndBehavior *SubscriptionSchedulePhaseTrialSettingsEndBehaviorParams `form:"end_behavior"`
}

// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase.
type SubscriptionSchedulePhaseParams struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
	AddInvoiceItems []*SubscriptionSchedulePhaseAddInvoiceItemParams `form:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// Automatic tax settings for this phase.
	AutomaticTax *SubscriptionSchedulePhaseAutomaticTaxParams `form:"automatic_tax"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionBillingThresholdsParams `form:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method"`
	// The identifier of the coupon to apply to this phase of the subscription schedule.
	Coupon *string `form:"coupon"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method"`
	// A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will set the Subscription's [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates), which means they will be the Invoice's [`default_tax_rates`](https://stripe.com/docs/api/invoices/create#create_invoice-default_tax_rates) for any Invoices issued by the Subscription during this Phase.
	DefaultTaxRates []*string `form:"default_tax_rates"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description *string `form:"description"`
	// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
	Discounts []*SubscriptionSchedulePhaseDiscountParams `form:"discounts"`
	// The date at which this phase of the subscription schedule ends. If set, `iterations` must not be set.
	EndDate    *int64 `form:"end_date"`
	EndDateNow *bool  `form:"-"` // See custom AppendTo
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionSchedulePhaseInvoiceSettingsParams `form:"invoice_settings"`
	// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
	Items []*SubscriptionSchedulePhaseItemParams `form:"items"`
	// Integer representing the multiplier applied to the price interval. For example, `iterations=2` applied to a price with `interval=month` and `interval_count=3` results in a phase of duration `2 * 3 months = 6 months`. If set, `end_date` must not be set.
	Iterations *int64 `form:"iterations"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered, adding new keys and replacing existing keys in the subscription's `metadata`. Individual keys in the subscription's `metadata` can be unset by posting an empty value to them in the phase's `metadata`. To unset all keys in the subscription's `metadata`, update the subscription directly or unset every key individually from the phase's `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of"`
	// If specified, payment collection for this subscription will be paused.
	PauseCollection *SubscriptionSchedulePhasePauseCollectionParams `form:"pause_collection"`
	// Whether the subscription schedule will create [prorations](https://stripe.com/docs/billing/subscriptions/prorations) when transitioning to this phase. The default value is `create_prorations`. This setting controls prorations when a phase is started asynchronously and it is persisted as a field on the phase. It's different from the request-level [proration_behavior](https://stripe.com/docs/api/subscription_schedules/update#update_subscription_schedule-proration_behavior) parameter which controls what happens if the update request affects the billing configuration of the current phase.
	ProrationBehavior *string `form:"proration_behavior"`
	// The date at which this phase of the subscription schedule starts or `now`. Must be set on the first phase.
	StartDate    *int64 `form:"start_date"`
	StartDateNow *bool  `form:"-"` // See custom AppendTo
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data"`
	// If set to true the entire phase is counted as a trial and the customer will not be charged for any fees.
	Trial *bool `form:"trial"`
	// Specify trial behavior when crossing phase boundaries
	TrialContinuation *string `form:"trial_continuation"`
	// Sets the phase to trialing from the start date to this date. Must be before the phase end date, can not be combined with `trial`
	TrialEnd    *int64 `form:"trial_end"`
	TrialEndNow *bool  `form:"-"` // See custom AppendTo
	// Settings related to subscription trials.
	TrialSettings *SubscriptionSchedulePhaseTrialSettingsParams `form:"trial_settings"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionSchedulePhaseParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// AppendTo implements custom encoding logic for SubscriptionSchedulePhaseParams.
func (p *SubscriptionSchedulePhaseParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.EndDateNow) {
		body.Add(form.FormatKey(append(keyParts, "end_date")), "now")
	}
	if BoolValue(p.TrialEndNow) {
		body.Add(form.FormatKey(append(keyParts, "trial_end")), "now")
	}
	if BoolValue(p.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// If specified, the invoicing for the given billing cycle iterations will be processed now.
type SubscriptionSchedulePrebillingParams struct {
	// This is used to determine the number of billing cycles to prebill.
	Iterations *int64 `form:"iterations"`
	// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period. The default value is `reset`.
	UpdateBehavior *string `form:"update_behavior"`
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
type SubscriptionScheduleParams struct {
	Params `form:"*"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior *string `form:"billing_behavior"`
	// The identifier of the customer to create the subscription schedule for.
	Customer *string `form:"customer"`
	// Object representing the subscription schedule's default settings.
	DefaultSettings *SubscriptionScheduleDefaultSettingsParams `form:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running.`cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior *string `form:"end_behavior"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Migrate an existing subscription to be managed by a subscription schedule. If this parameter is set, a subscription schedule will be created using the subscription's item(s), set to auto-renew using the subscription's interval. When using this parameter, other parameters (such as phase values) cannot be set. To create a subscription schedule with other modifications, we recommend making two separate API calls.
	FromSubscription *string `form:"from_subscription"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase. Note that past phases can be omitted.
	Phases []*SubscriptionSchedulePhaseParams `form:"phases"`
	// If specified, the invoicing for the given billing cycle iterations will be processed now.
	Prebilling *SubscriptionSchedulePrebillingParams `form:"prebilling"`
	// If the update changes the current phase, indicates whether the changes should be prorated. The default value is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior"`
	// When the subscription schedule starts. We recommend using `now` so that it starts the subscription immediately. You can also use a Unix timestamp to backdate the subscription so that it starts on a past date, or set a future date for the subscription to start on.
	StartDate    *int64 `form:"start_date"`
	StartDateNow *bool  `form:"-"` // See custom AppendTo
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// AppendTo implements custom encoding logic for SubscriptionScheduleParams.
func (p *SubscriptionScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// Use the `end` time of a given discount.
type SubscriptionScheduleAmendAmendmentAmendmentEndDiscountEndParams struct {
	// The ID of a specific discount.
	Discount *string `form:"discount"`
}

// Time span for the amendment starting from the `amendment_start`.
type SubscriptionScheduleAmendAmendmentAmendmentEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to identify the end of the time range modified by the proposed change. If not supplied, the amendment is considered a point-in-time operation that only affects the exact timestamp at `amendment_start`, and a restricted set of attributes is supported on the amendment.
type SubscriptionScheduleAmendAmendmentAmendmentEndParams struct {
	// Use the `end` time of a given discount.
	DiscountEnd *SubscriptionScheduleAmendAmendmentAmendmentEndDiscountEndParams `form:"discount_end"`
	// Time span for the amendment starting from the `amendment_start`.
	Duration *SubscriptionScheduleAmendAmendmentAmendmentEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the amendment to end. Must be after the `amendment_start`.
	Timestamp *int64 `form:"timestamp"`
	// Select one of three ways to pass the `amendment_end`.
	Type *string `form:"type"`
}

// Details of another amendment in the same array, immediately after which this amendment should begin.
type SubscriptionScheduleAmendAmendmentAmendmentStartAmendmentEndParams struct {
	// The position of the previous amendment in the `amendments` array after which this amendment should begin. Indexes start from 0 and must be less than the index of the current amendment in the array.
	Index *int64 `form:"index"`
}

// Use the `end` time of a given discount.
type SubscriptionScheduleAmendAmendmentAmendmentStartDiscountEndParams struct {
	// The ID of a specific discount.
	Discount *string `form:"discount"`
}

// Details to identify the earliest timestamp where the proposed change should take effect.
type SubscriptionScheduleAmendAmendmentAmendmentStartParams struct {
	// Details of another amendment in the same array, immediately after which this amendment should begin.
	AmendmentEnd *SubscriptionScheduleAmendAmendmentAmendmentStartAmendmentEndParams `form:"amendment_end"`
	// Use the `end` time of a given discount.
	DiscountEnd *SubscriptionScheduleAmendAmendmentAmendmentStartDiscountEndParams `form:"discount_end"`
	// A precise Unix timestamp for the amendment to start.
	Timestamp *int64 `form:"timestamp"`
	// Select one of three ways to pass the `amendment_start`.
	Type *string `form:"type"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleAmendAmendmentDiscountActionAddDiscountEndParams struct {
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// Details of the discount to add.
type SubscriptionScheduleAmendAmendmentDiscountActionAddParams struct {
	// The coupon code to redeem.
	Coupon *string `form:"coupon"`
	// An ID of an existing discount for a coupon that was already redeemed.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleAmendAmendmentDiscountActionAddDiscountEndParams `form:"discount_end"`
	// The index, starting at 0, at which to position the new discount. When not supplied, Stripe defaults to appending the discount to the end of the `discounts` array.
	Index *int64 `form:"index"`
}

// Details of the discount to remove.
type SubscriptionScheduleAmendAmendmentDiscountActionRemoveParams struct {
	// The coupon code to remove from the `discounts` array.
	Coupon *string `form:"coupon"`
	// The ID of a discount to remove from the `discounts` array.
	Discount *string `form:"discount"`
}

// Details of the discount to replace the existing discounts with.
type SubscriptionScheduleAmendAmendmentDiscountActionSetParams struct {
	// The coupon code to replace the `discounts` array with.
	Coupon *string `form:"coupon"`
	// An ID of an existing discount to replace the `discounts` array with.
	Discount *string `form:"discount"`
}

// Changes to the coupons being redeemed or discounts being applied during the amendment time span.
type SubscriptionScheduleAmendAmendmentDiscountActionParams struct {
	// Details of the discount to add.
	Add *SubscriptionScheduleAmendAmendmentDiscountActionAddParams `form:"add"`
	// Details of the discount to remove.
	Remove *SubscriptionScheduleAmendAmendmentDiscountActionRemoveParams `form:"remove"`
	// Details of the discount to replace the existing discounts with.
	Set *SubscriptionScheduleAmendAmendmentDiscountActionSetParams `form:"set"`
	// Determines the type of discount action.
	Type *string `form:"type"`
}

// Time span for the redeemed discount.
type SubscriptionScheduleAmendAmendmentItemActionAddDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleAmendAmendmentItemActionAddDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionScheduleAmendAmendmentItemActionAddDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The discounts applied to the item. Subscription item discounts are applied before subscription discounts.
type SubscriptionScheduleAmendAmendmentItemActionAddDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleAmendAmendmentItemActionAddDiscountDiscountEndParams `form:"discount_end"`
}

// Options that configure the trial on the subscription item.
type SubscriptionScheduleAmendAmendmentItemActionAddTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []*string `form:"converts_to"`
	// Determines the type of trial for this item.
	Type *string `form:"type"`
}

// Details of the subscription item to add. If an item with the same `price` exists, it will be replaced by this new item. Otherwise, it adds the new item.
type SubscriptionScheduleAmendAmendmentItemActionAddParams struct {
	// The discounts applied to the item. Subscription item discounts are applied before subscription discounts.
	Discounts []*SubscriptionScheduleAmendAmendmentItemActionAddDiscountParams `form:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Quantity for this item.
	Quantity *int64 `form:"quantity"`
	// The tax rates that apply to this subscription item. When set, the `default_tax_rates` on the subscription do not apply to this `subscription_item`.
	TaxRates []*string `form:"tax_rates"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionScheduleAmendAmendmentItemActionAddTrialParams `form:"trial"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleAmendAmendmentItemActionAddParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details of the subscription item to remove.
type SubscriptionScheduleAmendAmendmentItemActionRemoveParams struct {
	// ID of a price to remove.
	Price *string `form:"price"`
}

// Time span for the redeemed discount.
type SubscriptionScheduleAmendAmendmentItemActionSetDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleAmendAmendmentItemActionSetDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionScheduleAmendAmendmentItemActionSetDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// If an item with the `price` already exists, passing this will override the `discounts` array on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `discounts`.
type SubscriptionScheduleAmendAmendmentItemActionSetDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleAmendAmendmentItemActionSetDiscountDiscountEndParams `form:"discount_end"`
}

// If an item with the `price` already exists, passing this will override the `trial` configuration on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `trial`.
type SubscriptionScheduleAmendAmendmentItemActionSetTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []*string `form:"converts_to"`
	// Determines the type of trial for this item.
	Type *string `form:"type"`
}

// Details of the subscription item to replace the existing items with. If an item with the `set[price]` already exists, the `items` array is not cleared. Instead, all of the other `set` properties that are passed in this request will replace the existing values for the configuration item.
type SubscriptionScheduleAmendAmendmentItemActionSetParams struct {
	// If an item with the `price` already exists, passing this will override the `discounts` array on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `discounts`.
	Discounts []*SubscriptionScheduleAmendAmendmentItemActionSetDiscountParams `form:"discounts"`
	// If an item with the `price` already exists, passing this will override the `metadata` on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The ID of the price object.
	Price *string `form:"price"`
	// If an item with the `price` already exists, passing this will override the quantity on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `quantity`.
	Quantity *int64 `form:"quantity"`
	// If an item with the `price` already exists, passing this will override the `tax_rates` array on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `tax_rates`.
	TaxRates []*string `form:"tax_rates"`
	// If an item with the `price` already exists, passing this will override the `trial` configuration on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `trial`.
	Trial *SubscriptionScheduleAmendAmendmentItemActionSetTrialParams `form:"trial"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleAmendAmendmentItemActionSetParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Changes to the subscription items during the amendment time span.
type SubscriptionScheduleAmendAmendmentItemActionParams struct {
	// Details of the subscription item to add. If an item with the same `price` exists, it will be replaced by this new item. Otherwise, it adds the new item.
	Add *SubscriptionScheduleAmendAmendmentItemActionAddParams `form:"add"`
	// Details of the subscription item to remove.
	Remove *SubscriptionScheduleAmendAmendmentItemActionRemoveParams `form:"remove"`
	// Details of the subscription item to replace the existing items with. If an item with the `set[price]` already exists, the `items` array is not cleared. Instead, all of the other `set` properties that are passed in this request will replace the existing values for the configuration item.
	Set *SubscriptionScheduleAmendAmendmentItemActionSetParams `form:"set"`
	// Determines the type of item action.
	Type *string `form:"type"`
}

// Instructions for how to modify phase metadata
type SubscriptionScheduleAmendAmendmentMetadataActionParams struct {
	// Key-value pairs to add to schedule phase metadata. These values will merge with existing schedule phase metadata.
	Add map[string]string `form:"add"`
	// Keys to remove from schedule phase metadata.
	Remove []*string `form:"remove"`
	// Key-value pairs to set as schedule phase metadata. Existing schedule phase metadata will be overwritten.
	Set map[string]string `form:"set"`
	// Select one of three ways to update phase-level `metadata` on subscription schedules.
	Type *string `form:"type"`
}

// Details of the pause_collection behavior to apply to the amendment.
type SubscriptionScheduleAmendAmendmentSetPauseCollectionSetParams struct {
	// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
	Behavior *string `form:"behavior"`
}

// Defines how to pause collection for the underlying subscription throughout the duration of the amendment.
type SubscriptionScheduleAmendAmendmentSetPauseCollectionParams struct {
	// Details of the pause_collection behavior to apply to the amendment.
	Set *SubscriptionScheduleAmendAmendmentSetPauseCollectionSetParams `form:"set"`
	// Determines the type of the pause_collection amendment.
	Type *string `form:"type"`
}

// Defines how the subscription should behave when a trial ends.
type SubscriptionScheduleAmendAmendmentTrialSettingsEndBehaviorParams struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront *string `form:"prorate_up_front"`
}

// Settings related to subscription trials.
type SubscriptionScheduleAmendAmendmentTrialSettingsParams struct {
	// Defines how the subscription should behave when a trial ends.
	EndBehavior *SubscriptionScheduleAmendAmendmentTrialSettingsEndBehaviorParams `form:"end_behavior"`
}

// Changes to apply to the phases of the subscription schedule, in the order provided.
type SubscriptionScheduleAmendAmendmentParams struct {
	// Details to identify the end of the time range modified by the proposed change. If not supplied, the amendment is considered a point-in-time operation that only affects the exact timestamp at `amendment_start`, and a restricted set of attributes is supported on the amendment.
	AmendmentEnd *SubscriptionScheduleAmendAmendmentAmendmentEndParams `form:"amendment_end"`
	// Details to identify the earliest timestamp where the proposed change should take effect.
	AmendmentStart *SubscriptionScheduleAmendAmendmentAmendmentStartParams `form:"amendment_start"`
	// For a point-in-time amendment, this attribute lets you set or update whether the subscription's billing cycle anchor is reset at the `amendment_start` timestamp.
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Changes to the coupons being redeemed or discounts being applied during the amendment time span.
	DiscountActions []*SubscriptionScheduleAmendAmendmentDiscountActionParams `form:"discount_actions"`
	// Changes to the subscription items during the amendment time span.
	ItemActions []*SubscriptionScheduleAmendAmendmentItemActionParams `form:"item_actions"`
	// Instructions for how to modify phase metadata
	MetadataActions []*SubscriptionScheduleAmendAmendmentMetadataActionParams `form:"metadata_actions"`
	// Changes to how Stripe handles prorations during the amendment time span. Affects if and how prorations are created when a future phase starts. In cases where the amendment changes the currently active phase, it is used to determine whether or how to prorate now, at the time of the request. Also supported as a point-in-time operation when `amendment_end` is `null`.
	ProrationBehavior *string `form:"proration_behavior"`
	// Defines how to pause collection for the underlying subscription throughout the duration of the amendment.
	SetPauseCollection *SubscriptionScheduleAmendAmendmentSetPauseCollectionParams `form:"set_pause_collection"`
	// Ends the subscription schedule early as dictated by either the accompanying amendment's start or end.
	SetScheduleEnd *string `form:"set_schedule_end"`
	// Settings related to subscription trials.
	TrialSettings *SubscriptionScheduleAmendAmendmentTrialSettingsParams `form:"trial_settings"`
}

// Start the prebilled period when a specified amendment begins.
type SubscriptionScheduleAmendPrebillingBillFromAmendmentStartParams struct {
	// The position of the amendment in the `amendments` array with which prebilling should begin. Indexes start from 0 and must be less than the total number of supplied amendments.
	Index *int64 `form:"index"`
}

// The beginning of the prebilled time period. The default value is `now`.
type SubscriptionScheduleAmendPrebillingBillFromParams struct {
	// Start the prebilled period when a specified amendment begins.
	AmendmentStart *SubscriptionScheduleAmendPrebillingBillFromAmendmentStartParams `form:"amendment_start"`
	// Start the prebilled period at a precise integer timestamp, starting from the Unix epoch.
	Timestamp *int64 `form:"timestamp"`
	// Select one of several ways to pass the `bill_from` value.
	Type *string `form:"type"`
}

// End the prebilled period when a specified amendment ends.
type SubscriptionScheduleAmendPrebillingBillUntilAmendmentEndParams struct {
	// The position of the amendment in the `amendments` array at which prebilling should end. Indexes start from 0 and must be less than the total number of supplied amendments.
	Index *int64 `form:"index"`
}

// Time span for prebilling, starting from `bill_from`.
type SubscriptionScheduleAmendPrebillingBillUntilDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// The end of the prebilled time period.
type SubscriptionScheduleAmendPrebillingBillUntilParams struct {
	// End the prebilled period when a specified amendment ends.
	AmendmentEnd *SubscriptionScheduleAmendPrebillingBillUntilAmendmentEndParams `form:"amendment_end"`
	// Time span for prebilling, starting from `bill_from`.
	Duration *SubscriptionScheduleAmendPrebillingBillUntilDurationParams `form:"duration"`
	// End the prebilled period at a precise integer timestamp, starting from the Unix epoch.
	Timestamp *int64 `form:"timestamp"`
	// Select one of several ways to pass the `bill_until` value.
	Type *string `form:"type"`
}

// Provide any time periods to bill in advance.
type SubscriptionScheduleAmendPrebillingParams struct {
	// The beginning of the prebilled time period. The default value is `now`.
	BillFrom *SubscriptionScheduleAmendPrebillingBillFromParams `form:"bill_from"`
	// The end of the prebilled time period.
	BillUntil *SubscriptionScheduleAmendPrebillingBillUntilParams `form:"bill_until"`
	// When the prebilling invoice should be created. The default value is `now`.
	InvoiceAt *string `form:"invoice_at"`
	// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period. The default value is `reset`.
	UpdateBehavior *string `form:"update_behavior"`
}

// Changes to apply to the subscription schedule.
type SubscriptionScheduleAmendScheduleSettingsParams struct {
	// Behavior of the subscription schedule and underlying subscription when it ends.
	EndBehavior *string `form:"end_behavior"`
}

// Amends an existing subscription schedule.
type SubscriptionScheduleAmendParams struct {
	Params `form:"*"`
	// Changes to apply to the phases of the subscription schedule, in the order provided.
	Amendments []*SubscriptionScheduleAmendAmendmentParams `form:"amendments"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Provide any time periods to bill in advance.
	Prebilling []*SubscriptionScheduleAmendPrebillingParams `form:"prebilling"`
	// In cases where the amendment changes the currently active phase,
	//  specifies if and how to prorate at the time of the request.
	ProrationBehavior *string `form:"proration_behavior"`
	// Changes to apply to the subscription schedule.
	ScheduleSettings *SubscriptionScheduleAmendScheduleSettingsParams `form:"schedule_settings"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleAmendParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Cancels a subscription schedule and its associated subscription immediately (if the subscription schedule has an active subscription). A subscription schedule can only be canceled if its status is not_started or active.
type SubscriptionScheduleCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// If the subscription schedule is `active`, indicates if a final invoice will be generated that contains any un-invoiced metered usage and new/pending proration invoice items. Defaults to `true`.
	InvoiceNow *bool `form:"invoice_now"`
	// If the subscription schedule is `active`, indicates if the cancellation should be prorated. Defaults to `true`.
	Prorate *bool `form:"prorate"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Releases the subscription schedule immediately, which will stop scheduling of its phases, but leave any existing subscription in place. A schedule can only be released if its status is not_started or active. If the subscription schedule is currently associated with a subscription, releasing it will remove its subscription property and set the subscription's ID to the released_subscription property.
type SubscriptionScheduleReleaseParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Keep any cancellation on the subscription that the schedule has set
	PreserveCancelDate *bool `form:"preserve_cancel_date"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleReleaseParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
type SubscriptionScheduleCurrentPhase struct {
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType `json:"type"`
}

// The subscription schedule's default invoice settings.
type SubscriptionScheduleDefaultSettingsInvoiceSettings struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer `json:"issuer"`
}
type SubscriptionScheduleDefaultSettings struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                   `json:"application_fee_percent"`
	AutomaticTax          *SubscriptionAutomaticTax `json:"automatic_tax"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor SubscriptionScheduleDefaultSettingsBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *SubscriptionBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod *SubscriptionCollectionMethod `json:"collection_method"`
	// ID of the default payment method for the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description"`
	// The subscription schedule's default invoice settings.
	InvoiceSettings *SubscriptionScheduleDefaultSettingsInvoiceSettings `json:"invoice_settings"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *SubscriptionTransferData `json:"transfer_data"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType `json:"type"`
}

// The stackable discounts that will be applied to the item.
type SubscriptionSchedulePhaseAddInvoiceItemDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEnd `json:"discount_end"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
type SubscriptionSchedulePhaseAddInvoiceItem struct {
	// The stackable discounts that will be applied to the item.
	Discounts []*SubscriptionSchedulePhaseAddInvoiceItemDiscount `json:"discounts"`
	// ID of the price used to generate the invoice item.
	Price *Price `json:"price"`
	// The quantity of the invoice item.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*TaxRate `json:"tax_rates"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionSchedulePhaseDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type SubscriptionSchedulePhaseDiscountDiscountEndType `json:"type"`
}

// The stackable discounts that will be applied to the subscription on this phase. Subscription item discounts are applied before subscription discounts.
type SubscriptionSchedulePhaseDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionSchedulePhaseDiscountDiscountEnd `json:"discount_end"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionSchedulePhaseInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type SubscriptionSchedulePhaseInvoiceSettingsIssuerType `json:"type"`
}

// The invoice settings applicable during this phase.
type SubscriptionSchedulePhaseInvoiceSettings struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionSchedulePhaseInvoiceSettingsIssuer `json:"issuer"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionSchedulePhaseItemDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type SubscriptionSchedulePhaseItemDiscountDiscountEndType `json:"type"`
}

// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
type SubscriptionSchedulePhaseItemDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionSchedulePhaseItemDiscountDiscountEnd `json:"discount_end"`
}

// Options that configure the trial on the subscription item.
type SubscriptionSchedulePhaseItemTrial struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []string                               `json:"converts_to"`
	Type       SubscriptionSchedulePhaseItemTrialType `json:"type"`
}

// Subscription items to configure the subscription to during this phase of the subscription schedule.
type SubscriptionSchedulePhaseItem struct {
	// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
	BillingThresholds *SubscriptionItemBillingThresholds `json:"billing_thresholds"`
	// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*SubscriptionSchedulePhaseItemDiscount `json:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an item. Metadata on this item will update the underlying subscription item's `metadata` when the phase is entered.
	Metadata map[string]string `json:"metadata"`
	// ID of the plan to which the customer should be subscribed.
	Plan *Plan `json:"plan"`
	// ID of the price to which the customer should be subscribed.
	Price *Price `json:"price"`
	// Quantity of the plan to which the customer should be subscribed.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to this `phase_item`. When set, the `default_tax_rates` on the phase do not apply to this `phase_item`.
	TaxRates []*TaxRate `json:"tax_rates"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionSchedulePhaseItemTrial `json:"trial"`
}

// If specified, payment collection for this subscription will be paused.
type SubscriptionSchedulePhasePauseCollection struct {
	// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
	Behavior SubscriptionSchedulePhasePauseCollectionBehavior `json:"behavior"`
}

// Defines how the subscription should behaves when a trial ensd.
type SubscriptionSchedulePhaseTrialSettingsEndBehavior struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront SubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront `json:"prorate_up_front"`
}

// Settings related to any trials on the subscription during this phase.
type SubscriptionSchedulePhaseTrialSettings struct {
	// Defines how the subscription should behaves when a trial ensd.
	EndBehavior *SubscriptionSchedulePhaseTrialSettingsEndBehavior `json:"end_behavior"`
}

// Configuration for the subscription schedule's phases.
type SubscriptionSchedulePhase struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
	AddInvoiceItems []*SubscriptionSchedulePhaseAddInvoiceItem `json:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                   `json:"application_fee_percent"`
	AutomaticTax          *SubscriptionAutomaticTax `json:"automatic_tax"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor SubscriptionSchedulePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *SubscriptionBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod *SubscriptionCollectionMethod `json:"collection_method"`
	// ID of the coupon to use during this phase of the subscription schedule.
	Coupon *Coupon `json:"coupon"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// The default tax rates to apply to the subscription during this phase of the subscription schedule.
	DefaultTaxRates []*TaxRate `json:"default_tax_rates"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description"`
	// The stackable discounts that will be applied to the subscription on this phase. Subscription item discounts are applied before subscription discounts.
	Discounts []*SubscriptionSchedulePhaseDiscount `json:"discounts"`
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The invoice settings applicable during this phase.
	InvoiceSettings *SubscriptionSchedulePhaseInvoiceSettings `json:"invoice_settings"`
	// Subscription items to configure the subscription to during this phase of the subscription schedule.
	Items []*SubscriptionSchedulePhaseItem `json:"items"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered. Updating the underlying subscription's `metadata` directly will not affect the current phase's `metadata`.
	Metadata map[string]string `json:"metadata"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// If specified, payment collection for this subscription will be paused.
	PauseCollection *SubscriptionSchedulePhasePauseCollection `json:"pause_collection"`
	// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
	ProrationBehavior SubscriptionSchedulePhaseProrationBehavior `json:"proration_behavior"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *SubscriptionTransferData `json:"transfer_data"`
	// Specify behavior of the trial when crossing schedule phase boundaries
	TrialContinuation SubscriptionSchedulePhaseTrialContinuation `json:"trial_continuation"`
	// When the trial ends within the phase.
	TrialEnd int64 `json:"trial_end"`
	// Settings related to any trials on the subscription during this phase.
	TrialSettings *SubscriptionSchedulePhaseTrialSettings `json:"trial_settings"`
}

// Time period and invoice for a Subscription billed in advance.
type SubscriptionSchedulePrebilling struct {
	// ID of the prebilling invoice.
	Invoice *Invoice `json:"invoice"`
	// The end of the last period for which the invoice pre-bills.
	PeriodEnd int64 `json:"period_end"`
	// The start of the first period for which the invoice pre-bills.
	PeriodStart int64 `json:"period_start"`
	// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period.
	UpdateBehavior SubscriptionSchedulePrebillingUpdateBehavior `json:"update_behavior"`
}

// A subscription schedule allows you to create and manage the lifecycle of a subscription by predefining expected changes.
//
// Related guide: [Subscription schedules](https://stripe.com/docs/billing/subscriptions/subscription-schedules)
type SubscriptionSchedule struct {
	APIResource
	// ID of the Connect Application that created the schedule.
	Application *Application `json:"application"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior SubscriptionScheduleBillingBehavior `json:"billing_behavior"`
	// Time at which the subscription schedule was canceled. Measured in seconds since the Unix epoch.
	CanceledAt int64 `json:"canceled_at"`
	// Time at which the subscription schedule was completed. Measured in seconds since the Unix epoch.
	CompletedAt int64 `json:"completed_at"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
	CurrentPhase *SubscriptionScheduleCurrentPhase `json:"current_phase"`
	// ID of the customer who owns the subscription schedule.
	Customer        *Customer                            `json:"customer"`
	DefaultSettings *SubscriptionScheduleDefaultSettings `json:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running.`cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior SubscriptionScheduleEndBehavior `json:"end_behavior"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Configuration for the subscription schedule's phases.
	Phases []*SubscriptionSchedulePhase `json:"phases"`
	// Time period and invoice for a Subscription billed in advance.
	Prebilling *SubscriptionSchedulePrebilling `json:"prebilling"`
	// Time at which the subscription schedule was released. Measured in seconds since the Unix epoch.
	ReleasedAt int64 `json:"released_at"`
	// ID of the subscription once managed by the subscription schedule (if it is released).
	ReleasedSubscription *Subscription `json:"released_subscription"`
	// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
	Status SubscriptionScheduleStatus `json:"status"`
	// ID of the subscription managed by the subscription schedule.
	Subscription *Subscription `json:"subscription"`
	// ID of the test clock this subscription schedule belongs to.
	TestClock *TestHelpersTestClock `json:"test_clock"`
}

// SubscriptionScheduleList is a list of SubscriptionSchedules as retrieved from a list endpoint.
type SubscriptionScheduleList struct {
	APIResource
	ListMeta
	Data []*SubscriptionSchedule `json:"data"`
}

// UnmarshalJSON handles deserialization of a SubscriptionSchedule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SubscriptionSchedule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type subscriptionSchedule SubscriptionSchedule
	var v subscriptionSchedule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SubscriptionSchedule(v)
	return nil
}
