//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v85/form"
)

// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time. `prorate_up_front` will bill for all phases within the current billing cycle up front.
type SubscriptionScheduleBillingBehavior string

// List of values that SubscriptionScheduleBillingBehavior can take
const (
	SubscriptionScheduleBillingBehaviorProrateOnNextPhase SubscriptionScheduleBillingBehavior = "prorate_on_next_phase"
	SubscriptionScheduleBillingBehaviorProrateUpFront     SubscriptionScheduleBillingBehavior = "prorate_up_front"
)

// Controls how invoices and invoice items display proration amounts and discount amounts.
type SubscriptionScheduleBillingModeFlexibleProrationDiscounts string

// List of values that SubscriptionScheduleBillingModeFlexibleProrationDiscounts can take
const (
	SubscriptionScheduleBillingModeFlexibleProrationDiscountsIncluded SubscriptionScheduleBillingModeFlexibleProrationDiscounts = "included"
	SubscriptionScheduleBillingModeFlexibleProrationDiscountsItemized SubscriptionScheduleBillingModeFlexibleProrationDiscounts = "itemized"
)

// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
type SubscriptionScheduleBillingModeType string

// List of values that SubscriptionScheduleBillingModeType can take
const (
	SubscriptionScheduleBillingModeTypeClassic  SubscriptionScheduleBillingModeType = "classic"
	SubscriptionScheduleBillingModeTypeFlexible SubscriptionScheduleBillingModeType = "flexible"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).
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

// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
type SubscriptionScheduleEndBehavior string

// List of values that SubscriptionScheduleEndBehavior can take
const (
	SubscriptionScheduleEndBehaviorCancel  SubscriptionScheduleEndBehavior = "cancel"
	SubscriptionScheduleEndBehaviorNone    SubscriptionScheduleEndBehavior = "none"
	SubscriptionScheduleEndBehaviorRelease SubscriptionScheduleEndBehavior = "release"
	SubscriptionScheduleEndBehaviorRenew   SubscriptionScheduleEndBehavior = "renew"
)

// The type of error encountered by the price migration.
type SubscriptionScheduleLastPriceMigrationErrorType string

// List of values that SubscriptionScheduleLastPriceMigrationErrorType can take
const (
	SubscriptionScheduleLastPriceMigrationErrorTypePriceUniquenessViolation SubscriptionScheduleLastPriceMigrationErrorType = "price_uniqueness_violation"
)

// The discount end type.
type SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType string

// List of values that SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType can take
const (
	SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndTypeTimestamp SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType = "timestamp"
)

// Select how to calculate the end of the invoice item period.
type SubscriptionSchedulePhaseAddInvoiceItemPeriodEndType string

// List of values that SubscriptionSchedulePhaseAddInvoiceItemPeriodEndType can take
const (
	SubscriptionSchedulePhaseAddInvoiceItemPeriodEndTypeMinItemPeriodEnd SubscriptionSchedulePhaseAddInvoiceItemPeriodEndType = "min_item_period_end"
	SubscriptionSchedulePhaseAddInvoiceItemPeriodEndTypePhaseEnd         SubscriptionSchedulePhaseAddInvoiceItemPeriodEndType = "phase_end"
	SubscriptionSchedulePhaseAddInvoiceItemPeriodEndTypeTimestamp        SubscriptionSchedulePhaseAddInvoiceItemPeriodEndType = "timestamp"
)

// Select how to calculate the start of the invoice item period.
type SubscriptionSchedulePhaseAddInvoiceItemPeriodStartType string

// List of values that SubscriptionSchedulePhaseAddInvoiceItemPeriodStartType can take
const (
	SubscriptionSchedulePhaseAddInvoiceItemPeriodStartTypeMaxItemPeriodStart SubscriptionSchedulePhaseAddInvoiceItemPeriodStartType = "max_item_period_start"
	SubscriptionSchedulePhaseAddInvoiceItemPeriodStartTypePhaseStart         SubscriptionSchedulePhaseAddInvoiceItemPeriodStartType = "phase_start"
	SubscriptionSchedulePhaseAddInvoiceItemPeriodStartTypeTimestamp          SubscriptionSchedulePhaseAddInvoiceItemPeriodStartType = "timestamp"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).
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

// Determines the type of trial for this item.
type SubscriptionSchedulePhaseItemTrialType string

// List of values that SubscriptionSchedulePhaseItemTrialType can take
const (
	SubscriptionSchedulePhaseItemTrialTypeFree SubscriptionSchedulePhaseItemTrialType = "free"
	SubscriptionSchedulePhaseItemTrialTypePaid SubscriptionSchedulePhaseItemTrialType = "paid"
)

// The payment collection behavior for this subscription while paused.
type SubscriptionSchedulePhasePauseCollectionBehavior string

// List of values that SubscriptionSchedulePhasePauseCollectionBehavior can take
const (
	SubscriptionSchedulePhasePauseCollectionBehaviorKeepAsDraft       SubscriptionSchedulePhasePauseCollectionBehavior = "keep_as_draft"
	SubscriptionSchedulePhasePauseCollectionBehaviorMarkUncollectible SubscriptionSchedulePhasePauseCollectionBehavior = "mark_uncollectible"
	SubscriptionSchedulePhasePauseCollectionBehaviorVoid              SubscriptionSchedulePhasePauseCollectionBehavior = "void"
)

// When transitioning phases, controls how prorations are handled (if any). Possible values are `create_prorations`, `none`, and `always_invoice`.
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

// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://docs.stripe.com/billing/subscriptions/subscription-schedules).
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
	CanceledAt *int64 `form:"canceled_at" json:"canceled_at,omitempty"`
	// Only return subscription schedules that were created canceled the given date interval.
	CanceledAtRange *RangeQueryParams `form:"canceled_at" json:"-"`
	// Only return subscription schedules that completed during the given date interval.
	CompletedAt *int64 `form:"completed_at" json:"completed_at,omitempty"`
	// Only return subscription schedules that completed during the given date interval.
	CompletedAtRange *RangeQueryParams `form:"completed_at" json:"-"`
	// Only return subscription schedules that were created during the given date interval.
	Created *int64 `form:"created" json:"created,omitempty"`
	// Only return subscription schedules that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created" json:"-"`
	// Only return subscription schedules for the given customer.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Only return subscription schedules for the given account.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Only return subscription schedules that were released during the given date interval.
	ReleasedAt *int64 `form:"released_at" json:"released_at,omitempty"`
	// Only return subscription schedules that were released during the given date interval.
	ReleasedAtRange *RangeQueryParams `form:"released_at" json:"-"`
	// Only return subscription schedules that have not started yet.
	Scheduled *bool `form:"scheduled" json:"scheduled,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configure behavior for flexible billing mode.
type SubscriptionScheduleBillingModeFlexibleParams struct {
	// Controls how invoices and invoice items display proration amounts and discount amounts.
	ProrationDiscounts *string `form:"proration_discounts" json:"proration_discounts,omitempty"`
}

// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
type SubscriptionScheduleBillingModeParams struct {
	// Configure behavior for flexible billing mode.
	Flexible *SubscriptionScheduleBillingModeFlexibleParams `form:"flexible" json:"flexible,omitempty"`
	// Controls the calculation and orchestration of prorations and invoices for subscriptions. If no value is passed, the default is `flexible`.
	Type *string `form:"type" json:"type"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleDefaultSettingsBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte" json:"amount_gte,omitempty"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor" json:"reset_billing_cycle_anchor,omitempty"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsParams struct {
	// The account tax IDs associated with the subscription schedule. Will be set on invoices generated by the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids" json:"account_tax_ids,omitempty"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `collection_method=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due" json:"days_until_due,omitempty"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer      *SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerParams      `form:"issuer" json:"issuer,omitempty"`
	UnsetFields []SubscriptionScheduleDefaultSettingsInvoiceSettingsParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleDefaultSettingsInvoiceSettingsParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleDefaultSettingsInvoiceSettingsParams.
type SubscriptionScheduleDefaultSettingsInvoiceSettingsParamsUnsetField string

const (
	SubscriptionScheduleDefaultSettingsInvoiceSettingsParamsUnsetFieldAccountTaxIDs SubscriptionScheduleDefaultSettingsInvoiceSettingsParamsUnsetField = "account_tax_ids"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleDefaultSettingsInvoiceSettingsParams) AddUnsetField(field SubscriptionScheduleDefaultSettingsInvoiceSettingsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Object representing the subscription schedule's default settings.
type SubscriptionScheduleDefaultSettingsParams struct {
	Params `form:"*"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent,high_precision"`
	// Default settings for automatic tax computation.
	AutomaticTax *SubscriptionAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleDefaultSettingsBillingThresholdsParams `form:"billing_thresholds" json:"billing_thresholds,omitempty"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description" json:"description,omitempty"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleDefaultSettingsInvoiceSettingsParams `form:"invoice_settings" json:"invoice_settings,omitempty"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams                       `form:"transfer_data" json:"transfer_data,omitempty"`
	UnsetFields  []SubscriptionScheduleDefaultSettingsParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleDefaultSettingsParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleDefaultSettingsParams.
type SubscriptionScheduleDefaultSettingsParamsUnsetField string

const (
	SubscriptionScheduleDefaultSettingsParamsUnsetFieldBillingThresholds SubscriptionScheduleDefaultSettingsParamsUnsetField = "billing_thresholds"
	SubscriptionScheduleDefaultSettingsParamsUnsetFieldDescription       SubscriptionScheduleDefaultSettingsParamsUnsetField = "description"
	SubscriptionScheduleDefaultSettingsParamsUnsetFieldOnBehalfOf        SubscriptionScheduleDefaultSettingsParamsUnsetField = "on_behalf_of"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleDefaultSettingsParams) AddUnsetField(field SubscriptionScheduleDefaultSettingsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Time span for the redeemed discount.
type SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// The coupons to redeem into discounts for the item.
type SubscriptionSchedulePhaseAddInvoiceItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// End of the invoice item period.
type SubscriptionSchedulePhaseAddInvoiceItemPeriodEndParams struct {
	// A precise Unix timestamp for the end of the invoice item period. Must be greater than or equal to `period.start`.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// Select how to calculate the end of the invoice item period.
	Type *string `form:"type" json:"type"`
}

// Start of the invoice item period.
type SubscriptionSchedulePhaseAddInvoiceItemPeriodStartParams struct {
	// A precise Unix timestamp for the start of the invoice item period. Must be less than or equal to `period.end`.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// Select how to calculate the start of the invoice item period.
	Type *string `form:"type" json:"type"`
}

// The period associated with this invoice item. If not set, `period.start.type` defaults to `max_item_period_start` and `period.end.type` defaults to `min_item_period_end`.
type SubscriptionSchedulePhaseAddInvoiceItemPeriodParams struct {
	// End of the invoice item period.
	End *SubscriptionSchedulePhaseAddInvoiceItemPeriodEndParams `form:"end" json:"end"`
	// Start of the invoice item period.
	Start *SubscriptionSchedulePhaseAddInvoiceItemPeriodStartParams `form:"start" json:"start"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
type SubscriptionSchedulePhaseAddInvoiceItemParams struct {
	// The coupons to redeem into discounts for the item.
	Discounts []*SubscriptionSchedulePhaseAddInvoiceItemDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The period associated with this invoice item. If not set, `period.start.type` defaults to `max_item_period_start` and `period.end.type` defaults to `min_item_period_end`.
	Period *SubscriptionSchedulePhaseAddInvoiceItemPeriodParams `form:"period" json:"period,omitempty"`
	// The ID of the price object. One of `price` or `price_data` is required.
	Price *string `form:"price" json:"price,omitempty"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *InvoiceItemPriceDataParams `form:"price_data" json:"price_data,omitempty"`
	// Quantity for this item. Defaults to 1.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates    []*string                                                 `form:"tax_rates" json:"tax_rates,omitempty"`
	UnsetFields []SubscriptionSchedulePhaseAddInvoiceItemParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionSchedulePhaseAddInvoiceItemParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionSchedulePhaseAddInvoiceItemParams.
type SubscriptionSchedulePhaseAddInvoiceItemParamsUnsetField string

const (
	SubscriptionSchedulePhaseAddInvoiceItemParamsUnsetFieldTaxRates SubscriptionSchedulePhaseAddInvoiceItemParamsUnsetField = "tax_rates"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionSchedulePhaseAddInvoiceItemParams) AddUnsetField(field SubscriptionSchedulePhaseAddInvoiceItemParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionSchedulePhaseAddInvoiceItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type SubscriptionSchedulePhaseAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// Automatic tax settings for this phase.
type SubscriptionSchedulePhaseAutomaticTaxParams struct {
	// Enabled automatic tax calculation which will automatically compute tax rates on all invoices generated by the subscription.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *SubscriptionSchedulePhaseAutomaticTaxLiabilityParams `form:"liability" json:"liability,omitempty"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionSchedulePhaseBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte" json:"amount_gte,omitempty"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor" json:"reset_billing_cycle_anchor,omitempty"`
}

// Time span for the redeemed discount.
type SubscriptionSchedulePhaseDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionSchedulePhaseDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionSchedulePhaseDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
type SubscriptionSchedulePhaseDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionSchedulePhaseDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// The number of intervals the phase should last. If set, `end_date` must not be set.
type SubscriptionSchedulePhaseDurationParams struct {
	// Specifies phase duration. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The multiplier applied to the interval.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionSchedulePhaseInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionSchedulePhaseInvoiceSettingsParams struct {
	// The account tax IDs associated with this phase of the subscription schedule. Will be set on invoices generated by this phase of the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids" json:"account_tax_ids,omitempty"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due" json:"days_until_due,omitempty"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer      *SubscriptionSchedulePhaseInvoiceSettingsIssuerParams      `form:"issuer" json:"issuer,omitempty"`
	UnsetFields []SubscriptionSchedulePhaseInvoiceSettingsParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionSchedulePhaseInvoiceSettingsParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionSchedulePhaseInvoiceSettingsParams.
type SubscriptionSchedulePhaseInvoiceSettingsParamsUnsetField string

const (
	SubscriptionSchedulePhaseInvoiceSettingsParamsUnsetFieldAccountTaxIDs SubscriptionSchedulePhaseInvoiceSettingsParamsUnsetField = "account_tax_ids"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionSchedulePhaseInvoiceSettingsParams) AddUnsetField(field SubscriptionSchedulePhaseInvoiceSettingsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionSchedulePhaseItemBillingThresholdsParams struct {
	// Number of units that meets the billing threshold to advance the subscription to a new billing period (e.g., it takes 10 $5 units to meet a $50 [monetary threshold](https://docs.stripe.com/api/subscriptions/update#update_subscription-billing_thresholds-amount_gte))
	UsageGTE *int64 `form:"usage_gte" json:"usage_gte"`
}

// Time span for the redeemed discount.
type SubscriptionSchedulePhaseItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionSchedulePhaseItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionSchedulePhaseItemDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// The coupons to redeem into discounts for the subscription item.
type SubscriptionSchedulePhaseItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionSchedulePhaseItemDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// Options that configure the trial on the subscription item.
type SubscriptionSchedulePhaseItemTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial. Currently only supports at most 1 price ID.
	ConvertsTo []*string `form:"converts_to" json:"converts_to,omitempty"`
	// Determines the type of trial for this item.
	Type *string `form:"type" json:"type"`
}

// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
type SubscriptionSchedulePhaseItemParams struct {
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionSchedulePhaseItemBillingThresholdsParams `form:"billing_thresholds" json:"billing_thresholds,omitempty"`
	// The coupons to redeem into discounts for the subscription item.
	Discounts []*SubscriptionSchedulePhaseItemDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to a configuration item. Metadata on a configuration item will update the underlying subscription item's `metadata` when the phase is entered, adding new keys and replacing existing keys. Individual keys in the subscription item's `metadata` can be unset by posting an empty value to them in the configuration item's `metadata`. To unset all keys in the subscription item's `metadata`, update the subscription item directly or unset every key individually from the configuration item's `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The plan ID to subscribe to. You may specify the same ID in `plan` and `price`.
	Plan *string `form:"plan" json:"plan,omitempty"`
	// The ID of the price object.
	Price *string `form:"price" json:"price,omitempty"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline.
	PriceData *SubscriptionItemPriceDataParams `form:"price_data" json:"price_data,omitempty"`
	// Quantity for the given price. Can be set only if the price's `usage_type` is `licensed` and not `metered`.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
	// A list of [Tax Rate](https://docs.stripe.com/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://docs.stripe.com/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates" json:"tax_rates,omitempty"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionSchedulePhaseItemTrialParams `form:"trial" json:"trial,omitempty"`
	// The ID of the trial offer to apply to the configuration item.
	TrialOffer  *string                                         `form:"trial_offer" json:"trial_offer,omitempty"`
	UnsetFields []SubscriptionSchedulePhaseItemParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionSchedulePhaseItemParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionSchedulePhaseItemParams.
type SubscriptionSchedulePhaseItemParamsUnsetField string

const (
	SubscriptionSchedulePhaseItemParamsUnsetFieldBillingThresholds SubscriptionSchedulePhaseItemParamsUnsetField = "billing_thresholds"
	SubscriptionSchedulePhaseItemParamsUnsetFieldDiscounts         SubscriptionSchedulePhaseItemParamsUnsetField = "discounts"
	SubscriptionSchedulePhaseItemParamsUnsetFieldTaxRates          SubscriptionSchedulePhaseItemParamsUnsetField = "tax_rates"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionSchedulePhaseItemParams) AddUnsetField(field SubscriptionSchedulePhaseItemParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionSchedulePhaseItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://docs.stripe.com/billing/subscriptions/pause-payment).
type SubscriptionSchedulePhasePauseCollectionParams struct {
	// The payment collection behavior for this subscription while paused.
	Behavior *string `form:"behavior" json:"behavior"`
}

// Defines how the subscription should behave when a trial ends.
type SubscriptionSchedulePhaseTrialSettingsEndBehaviorParams struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront *string `form:"prorate_up_front" json:"prorate_up_front,omitempty"`
}

// Settings related to subscription trials.
type SubscriptionSchedulePhaseTrialSettingsParams struct {
	// Defines how the subscription should behave when a trial ends.
	EndBehavior *SubscriptionSchedulePhaseTrialSettingsEndBehaviorParams `form:"end_behavior" json:"end_behavior,omitempty"`
}

// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase.
type SubscriptionSchedulePhaseParams struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
	AddInvoiceItems []*SubscriptionSchedulePhaseAddInvoiceItemParams `form:"add_invoice_items" json:"add_invoice_items,omitempty"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent" json:"application_fee_percent,omitempty"`
	// Automatic tax settings for this phase.
	AutomaticTax *SubscriptionSchedulePhaseAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionSchedulePhaseBillingThresholdsParams `form:"billing_thresholds" json:"billing_thresholds,omitempty"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// A list of [Tax Rate](https://docs.stripe.com/api/tax_rates) ids. These Tax Rates will set the Subscription's [`default_tax_rates`](https://docs.stripe.com/api/subscriptions/create#create_subscription-default_tax_rates), which means they will be the Invoice's [`default_tax_rates`](https://docs.stripe.com/api/invoices/create#create_invoice-default_tax_rates) for any Invoices issued by the Subscription during this Phase.
	DefaultTaxRates []*string `form:"default_tax_rates" json:"default_tax_rates,omitempty"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description" json:"description,omitempty"`
	// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
	Discounts []*SubscriptionSchedulePhaseDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// The number of intervals the phase should last. If set, `end_date` must not be set.
	Duration *SubscriptionSchedulePhaseDurationParams `form:"duration" json:"duration,omitempty"`
	// The date at which this phase of the subscription schedule ends. If set, `duration` must not be set.
	EndDate    *int64 `form:"end_date" json:"end_date,omitempty"`
	EndDateNow *bool  `form:"-"` // See custom AppendTo
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionSchedulePhaseInvoiceSettingsParams `form:"invoice_settings" json:"invoice_settings,omitempty"`
	// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
	Items []*SubscriptionSchedulePhaseItemParams `form:"items" json:"items"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered, adding new keys and replacing existing keys in the subscription's `metadata`. Individual keys in the subscription's `metadata` can be unset by posting an empty value to them in the phase's `metadata`. To unset all keys in the subscription's `metadata`, update the subscription directly or unset every key individually from the phase's `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://docs.stripe.com/billing/subscriptions/pause-payment).
	PauseCollection *SubscriptionSchedulePhasePauseCollectionParams `form:"pause_collection" json:"pause_collection,omitempty"`
	// Controls whether the subscription schedule should create [prorations](https://docs.stripe.com/billing/subscriptions/prorations) when transitioning to this phase if there is a difference in billing configuration. It's different from the request-level [proration_behavior](https://docs.stripe.com/api/subscription_schedules/update#update_subscription_schedule-proration_behavior) parameter which controls what happens if the update request affects the billing configuration (item price, quantity, etc.) of the current phase.
	ProrationBehavior *string `form:"proration_behavior" json:"proration_behavior,omitempty"`
	// The date at which this phase of the subscription schedule starts or `now`. Must be set on the first phase.
	StartDate    *int64 `form:"start_date" json:"start_date,omitempty"`
	StartDateNow *bool  `form:"-"` // See custom AppendTo
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
	// If set to true the entire phase is counted as a trial and the customer will not be charged for any fees.
	Trial *bool `form:"trial" json:"trial,omitempty"`
	// Specify trial behavior when crossing phase boundaries
	TrialContinuation *string `form:"trial_continuation" json:"trial_continuation,omitempty"`
	// Sets the phase to trialing from the start date to this date. Must be before the phase end date, can not be combined with `trial`
	TrialEnd    *int64 `form:"trial_end" json:"trial_end,omitempty"`
	TrialEndNow *bool  `form:"-"` // See custom AppendTo
	// Settings related to subscription trials.
	TrialSettings *SubscriptionSchedulePhaseTrialSettingsParams `form:"trial_settings" json:"trial_settings,omitempty"`
	UnsetFields   []SubscriptionSchedulePhaseParamsUnsetField   `form:"-" json:"-"`
}

// SubscriptionSchedulePhaseParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionSchedulePhaseParams.
type SubscriptionSchedulePhaseParamsUnsetField string

const (
	SubscriptionSchedulePhaseParamsUnsetFieldBillingThresholds SubscriptionSchedulePhaseParamsUnsetField = "billing_thresholds"
	SubscriptionSchedulePhaseParamsUnsetFieldDefaultTaxRates   SubscriptionSchedulePhaseParamsUnsetField = "default_tax_rates"
	SubscriptionSchedulePhaseParamsUnsetFieldDescription       SubscriptionSchedulePhaseParamsUnsetField = "description"
	SubscriptionSchedulePhaseParamsUnsetFieldDiscounts         SubscriptionSchedulePhaseParamsUnsetField = "discounts"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionSchedulePhaseParams) AddUnsetField(field SubscriptionSchedulePhaseParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
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
	Iterations *int64 `form:"iterations" json:"iterations"`
	// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period. The default value is `reset`.
	UpdateBehavior *string `form:"update_behavior" json:"update_behavior,omitempty"`
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
type SubscriptionScheduleParams struct {
	Params `form:"*"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time. `prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior *string `form:"billing_behavior" json:"billing_behavior,omitempty"`
	// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
	BillingMode *SubscriptionScheduleBillingModeParams `form:"billing_mode" json:"billing_mode,omitempty"`
	// The identifier of the customer to create the subscription schedule for.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The identifier of the account to create the subscription schedule for.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// Object representing the subscription schedule's default settings.
	DefaultSettings *SubscriptionScheduleDefaultSettingsParams `form:"default_settings" json:"default_settings,omitempty"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior *string `form:"end_behavior" json:"end_behavior,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Migrate an existing subscription to be managed by a subscription schedule. If this parameter is set, a subscription schedule will be created using the subscription's item(s), set to auto-renew using the subscription's interval. When using this parameter, other parameters (such as phase values) cannot be set. To create a subscription schedule with other modifications, we recommend making two separate API calls.
	FromSubscription *string `form:"from_subscription" json:"from_subscription,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase. Note that past phases can be omitted.
	Phases []*SubscriptionSchedulePhaseParams `form:"phases" json:"phases,omitempty"`
	// If specified, the invoicing for the given billing cycle iterations will be processed now.
	Prebilling *SubscriptionSchedulePrebillingParams `form:"prebilling" json:"prebilling,omitempty"`
	// If the update changes the billing configuration (item price, quantity, etc.) of the current phase, indicates how prorations from this change should be handled. The default value is `create_prorations`.
	ProrationBehavior *string `form:"proration_behavior" json:"proration_behavior,omitempty"`
	// When the subscription schedule starts. We recommend using `now` so that it starts the subscription immediately. You can also use a Unix timestamp to backdate the subscription so that it starts on a past date, or set a future date for the subscription to start on.
	StartDate    *int64                                 `form:"start_date" json:"start_date,omitempty"`
	StartDateNow *bool                                  `form:"-"` // See custom AppendTo
	UnsetFields  []SubscriptionScheduleParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleParams.
type SubscriptionScheduleParamsUnsetField string

const (
	SubscriptionScheduleParamsUnsetFieldMetadata SubscriptionScheduleParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleParams) AddUnsetField(field SubscriptionScheduleParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
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
	Discount *string `form:"discount" json:"discount"`
}

// Time span for the amendment starting from the `amendment_start`.
type SubscriptionScheduleAmendAmendmentAmendmentEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to identify the end of the time range modified by the proposed change. If not supplied, the amendment is considered a point-in-time operation that only affects the exact timestamp at `amendment_start`, and a restricted set of attributes is supported on the amendment.
type SubscriptionScheduleAmendAmendmentAmendmentEndParams struct {
	// Use the `end` time of a given discount.
	DiscountEnd *SubscriptionScheduleAmendAmendmentAmendmentEndDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// Time span for the amendment starting from the `amendment_start`.
	Duration *SubscriptionScheduleAmendAmendmentAmendmentEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the amendment to end. Must be after the `amendment_start`.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// Select one of three ways to pass the `amendment_end`.
	Type *string `form:"type" json:"type"`
}

// Details of another amendment in the same array, immediately after which this amendment should begin.
type SubscriptionScheduleAmendAmendmentAmendmentStartAmendmentEndParams struct {
	// The position of the previous amendment in the `amendments` array after which this amendment should begin. Indexes start from 0 and must be less than the index of the current amendment in the array.
	Index *int64 `form:"index" json:"index"`
}

// Use the `end` time of a given discount.
type SubscriptionScheduleAmendAmendmentAmendmentStartDiscountEndParams struct {
	// The ID of a specific discount.
	Discount *string `form:"discount" json:"discount"`
}

// Details to identify the earliest timestamp where the proposed change should take effect.
type SubscriptionScheduleAmendAmendmentAmendmentStartParams struct {
	// Details of another amendment in the same array, immediately after which this amendment should begin.
	AmendmentEnd *SubscriptionScheduleAmendAmendmentAmendmentStartAmendmentEndParams `form:"amendment_end" json:"amendment_end,omitempty"`
	// Use the `end` time of a given discount.
	DiscountEnd *SubscriptionScheduleAmendAmendmentAmendmentStartDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// A precise Unix timestamp for the amendment to start.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// Select one of three ways to pass the `amendment_start`.
	Type *string `form:"type" json:"type"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleAmendAmendmentDiscountActionAddDiscountEndParams struct {
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// Details of the discount to add.
type SubscriptionScheduleAmendAmendmentDiscountActionAddParams struct {
	// The coupon code to redeem.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// An ID of an existing discount for a coupon that was already redeemed.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleAmendAmendmentDiscountActionAddDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// The index, starting at 0, at which to position the new discount. When not supplied, Stripe defaults to appending the discount to the end of the `discounts` array.
	Index *int64 `form:"index" json:"index,omitempty"`
	// The promotion code to redeem.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// Details of the discount to remove.
type SubscriptionScheduleAmendAmendmentDiscountActionRemoveParams struct {
	// The coupon code to remove from the `discounts` array.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// The ID of a discount to remove from the `discounts` array.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// The ID of a promotion code to remove from the `discounts` array.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// Details of the discount to replace the existing discounts with.
type SubscriptionScheduleAmendAmendmentDiscountActionSetParams struct {
	// The coupon code to replace the `discounts` array with.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// An ID of an existing discount to replace the `discounts` array with.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// An ID of an existing promotion code to replace the `discounts` array with.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// Changes to the coupons being redeemed or discounts being applied during the amendment time span.
type SubscriptionScheduleAmendAmendmentDiscountActionParams struct {
	// Details of the discount to add.
	Add *SubscriptionScheduleAmendAmendmentDiscountActionAddParams `form:"add" json:"add,omitempty"`
	// Details of the discount to remove.
	Remove *SubscriptionScheduleAmendAmendmentDiscountActionRemoveParams `form:"remove" json:"remove,omitempty"`
	// Details of the discount to replace the existing discounts with.
	Set *SubscriptionScheduleAmendAmendmentDiscountActionSetParams `form:"set" json:"set,omitempty"`
	// Determines the type of discount action.
	Type *string `form:"type" json:"type"`
}

// Time span for the redeemed discount.
type SubscriptionScheduleAmendAmendmentItemActionAddDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleAmendAmendmentItemActionAddDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionScheduleAmendAmendmentItemActionAddDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// The discounts applied to the item. Subscription item discounts are applied before subscription discounts.
type SubscriptionScheduleAmendAmendmentItemActionAddDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleAmendAmendmentItemActionAddDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// Options that configure the trial on the subscription item.
type SubscriptionScheduleAmendAmendmentItemActionAddTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial. Currently only supports at most 1 price ID.
	ConvertsTo []*string `form:"converts_to" json:"converts_to,omitempty"`
	// Determines the type of trial for this item.
	Type *string `form:"type" json:"type"`
}

// Details of the subscription item to add. If an item with the same `price` exists, it will be replaced by this new item. Otherwise, it adds the new item.
type SubscriptionScheduleAmendAmendmentItemActionAddParams struct {
	// The discounts applied to the item. Subscription item discounts are applied before subscription discounts.
	Discounts []*SubscriptionScheduleAmendAmendmentItemActionAddDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The ID of the price object.
	Price *string `form:"price" json:"price"`
	// Quantity for this item.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
	// The tax rates that apply to this subscription item. When set, the `default_tax_rates` on the subscription do not apply to this `subscription_item`.
	TaxRates []*string `form:"tax_rates" json:"tax_rates,omitempty"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionScheduleAmendAmendmentItemActionAddTrialParams `form:"trial" json:"trial,omitempty"`
	// The ID of the trial offer to apply to the configuration item.
	TrialOffer *string `form:"trial_offer" json:"trial_offer,omitempty"`
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
	Price *string `form:"price" json:"price"`
}

// Time span for the redeemed discount.
type SubscriptionScheduleAmendAmendmentItemActionSetDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleAmendAmendmentItemActionSetDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionScheduleAmendAmendmentItemActionSetDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// If an item with the `price` already exists, passing this will override the `discounts` array on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `discounts`.
type SubscriptionScheduleAmendAmendmentItemActionSetDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleAmendAmendmentItemActionSetDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// If an item with the `price` already exists, passing this will override the `trial` configuration on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `trial`.
type SubscriptionScheduleAmendAmendmentItemActionSetTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial. Currently only supports at most 1 price ID.
	ConvertsTo []*string `form:"converts_to" json:"converts_to,omitempty"`
	// Determines the type of trial for this item.
	Type *string `form:"type" json:"type"`
}

// Details of the subscription item to replace the existing items with. If an item with the `set[price]` already exists, the `items` array is not cleared. Instead, all of the other `set` properties that are passed in this request will replace the existing values for the configuration item.
type SubscriptionScheduleAmendAmendmentItemActionSetParams struct {
	// If an item with the `price` already exists, passing this will override the `discounts` array on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `discounts`.
	Discounts []*SubscriptionScheduleAmendAmendmentItemActionSetDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// If an item with the `price` already exists, passing this will override the `metadata` on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The ID of the price object.
	Price *string `form:"price" json:"price"`
	// If an item with the `price` already exists, passing this will override the quantity on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `quantity`.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
	// If an item with the `price` already exists, passing this will override the `tax_rates` array on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `tax_rates`.
	TaxRates []*string `form:"tax_rates" json:"tax_rates,omitempty"`
	// If an item with the `price` already exists, passing this will override the `trial` configuration on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `trial`.
	Trial *SubscriptionScheduleAmendAmendmentItemActionSetTrialParams `form:"trial" json:"trial,omitempty"`
	// The ID of the trial offer to apply to the configuration item.
	TrialOffer *string `form:"trial_offer" json:"trial_offer,omitempty"`
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
	Add *SubscriptionScheduleAmendAmendmentItemActionAddParams `form:"add" json:"add,omitempty"`
	// Details of the subscription item to remove.
	Remove *SubscriptionScheduleAmendAmendmentItemActionRemoveParams `form:"remove" json:"remove,omitempty"`
	// Details of the subscription item to replace the existing items with. If an item with the `set[price]` already exists, the `items` array is not cleared. Instead, all of the other `set` properties that are passed in this request will replace the existing values for the configuration item.
	Set *SubscriptionScheduleAmendAmendmentItemActionSetParams `form:"set" json:"set,omitempty"`
	// Determines the type of item action.
	Type *string `form:"type" json:"type"`
}

// Instructions for how to modify phase metadata
type SubscriptionScheduleAmendAmendmentMetadataActionParams struct {
	// Key-value pairs to add to schedule phase metadata. These values will merge with existing schedule phase metadata.
	Add map[string]string `form:"add" json:"add,omitempty"`
	// Keys to remove from schedule phase metadata.
	Remove []*string `form:"remove" json:"remove,omitempty"`
	// Key-value pairs to set as schedule phase metadata. Existing schedule phase metadata will be overwritten.
	Set map[string]string `form:"set" json:"set,omitempty"`
	// Select one of three ways to update phase-level `metadata` on subscription schedules.
	Type        *string                                                            `form:"type" json:"type"`
	UnsetFields []SubscriptionScheduleAmendAmendmentMetadataActionParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleAmendAmendmentMetadataActionParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleAmendAmendmentMetadataActionParams.
type SubscriptionScheduleAmendAmendmentMetadataActionParamsUnsetField string

const (
	SubscriptionScheduleAmendAmendmentMetadataActionParamsUnsetFieldSet SubscriptionScheduleAmendAmendmentMetadataActionParamsUnsetField = "set"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleAmendAmendmentMetadataActionParams) AddUnsetField(field SubscriptionScheduleAmendAmendmentMetadataActionParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Details of the pause_collection behavior to apply to the amendment.
type SubscriptionScheduleAmendAmendmentSetPauseCollectionSetParams struct {
	// The payment collection behavior for this subscription while paused.
	Behavior *string `form:"behavior" json:"behavior"`
}

// Defines how to pause collection for the underlying subscription throughout the duration of the amendment.
type SubscriptionScheduleAmendAmendmentSetPauseCollectionParams struct {
	// Details of the pause_collection behavior to apply to the amendment.
	Set *SubscriptionScheduleAmendAmendmentSetPauseCollectionSetParams `form:"set" json:"set,omitempty"`
	// Determines the type of the pause_collection amendment.
	Type *string `form:"type" json:"type"`
}

// Defines how the subscription should behave when a trial ends.
type SubscriptionScheduleAmendAmendmentTrialSettingsEndBehaviorParams struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront *string `form:"prorate_up_front" json:"prorate_up_front,omitempty"`
}

// Settings related to subscription trials.
type SubscriptionScheduleAmendAmendmentTrialSettingsParams struct {
	// Defines how the subscription should behave when a trial ends.
	EndBehavior *SubscriptionScheduleAmendAmendmentTrialSettingsEndBehaviorParams `form:"end_behavior" json:"end_behavior,omitempty"`
}

// Changes to apply to the phases of the subscription schedule, in the order provided.
type SubscriptionScheduleAmendAmendmentParams struct {
	// Details to identify the end of the time range modified by the proposed change. If not supplied, the amendment is considered a point-in-time operation that only affects the exact timestamp at `amendment_start`, and a restricted set of attributes is supported on the amendment.
	AmendmentEnd *SubscriptionScheduleAmendAmendmentAmendmentEndParams `form:"amendment_end" json:"amendment_end,omitempty"`
	// Details to identify the earliest timestamp where the proposed change should take effect.
	AmendmentStart *SubscriptionScheduleAmendAmendmentAmendmentStartParams `form:"amendment_start" json:"amendment_start"`
	// For point-in-time amendments (having no `amendment_end`), this attribute lets you set or remove whether the subscription's billing cycle anchor is reset at the `amendment_start` timestamp.For time-span based amendments (having both `amendment_start` and `amendment_end`), the only value valid is `automatic`, which removes any previously configured billing cycle anchor resets scheduled to occur during the window of time spanned by the amendment.
	BillingCycleAnchor *string `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Changes to the coupons being redeemed or discounts being applied during the amendment time span.
	DiscountActions []*SubscriptionScheduleAmendAmendmentDiscountActionParams `form:"discount_actions" json:"discount_actions,omitempty"`
	// Changes to the subscription items during the amendment time span.
	ItemActions []*SubscriptionScheduleAmendAmendmentItemActionParams `form:"item_actions" json:"item_actions,omitempty"`
	// Instructions for how to modify phase metadata
	MetadataActions []*SubscriptionScheduleAmendAmendmentMetadataActionParams `form:"metadata_actions" json:"metadata_actions,omitempty"`
	// Changes to how Stripe handles prorations during the amendment time span. Affects if and how prorations are created when a future phase starts. In cases where the amendment changes the currently active phase, it is used to determine whether or how to prorate now, at the time of the request. Also supported as a point-in-time operation when `amendment_end` is `null`.
	ProrationBehavior *string `form:"proration_behavior" json:"proration_behavior,omitempty"`
	// Defines how to pause collection for the underlying subscription throughout the duration of the amendment.
	SetPauseCollection *SubscriptionScheduleAmendAmendmentSetPauseCollectionParams `form:"set_pause_collection" json:"set_pause_collection,omitempty"`
	// Ends the subscription schedule early as dictated by either the accompanying amendment's start or end.
	SetScheduleEnd *string `form:"set_schedule_end" json:"set_schedule_end,omitempty"`
	// Settings related to subscription trials.
	TrialSettings *SubscriptionScheduleAmendAmendmentTrialSettingsParams `form:"trial_settings" json:"trial_settings,omitempty"`
}

// Start the prebilled period when a specified amendment begins.
type SubscriptionScheduleAmendPrebillingBillFromAmendmentStartParams struct {
	// The position of the amendment in the `amendments` array with which prebilling should begin. Indexes start from 0 and must be less than the total number of supplied amendments.
	Index *int64 `form:"index" json:"index"`
}

// The beginning of the prebilled time period. The default value is `now`.
type SubscriptionScheduleAmendPrebillingBillFromParams struct {
	// Start the prebilled period when a specified amendment begins.
	AmendmentStart *SubscriptionScheduleAmendPrebillingBillFromAmendmentStartParams `form:"amendment_start" json:"amendment_start,omitempty"`
	// Start the prebilled period at a precise integer timestamp, starting from the Unix epoch.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// Select one of several ways to pass the `bill_from` value.
	Type *string `form:"type" json:"type"`
}

// End the prebilled period when a specified amendment ends.
type SubscriptionScheduleAmendPrebillingBillUntilAmendmentEndParams struct {
	// The position of the amendment in the `amendments` array at which prebilling should end. Indexes start from 0 and must be less than the total number of supplied amendments.
	Index *int64 `form:"index" json:"index"`
}

// Time span for prebilling, starting from `bill_from`.
type SubscriptionScheduleAmendPrebillingBillUntilDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// The end of the prebilled time period.
type SubscriptionScheduleAmendPrebillingBillUntilParams struct {
	// End the prebilled period when a specified amendment ends.
	AmendmentEnd *SubscriptionScheduleAmendPrebillingBillUntilAmendmentEndParams `form:"amendment_end" json:"amendment_end,omitempty"`
	// Time span for prebilling, starting from `bill_from`.
	Duration *SubscriptionScheduleAmendPrebillingBillUntilDurationParams `form:"duration" json:"duration,omitempty"`
	// End the prebilled period at a precise integer timestamp, starting from the Unix epoch.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// Select one of several ways to pass the `bill_until` value.
	Type *string `form:"type" json:"type"`
}

// Provide any time periods to bill in advance.
type SubscriptionScheduleAmendPrebillingParams struct {
	// The beginning of the prebilled time period. The default value is `now`.
	BillFrom *SubscriptionScheduleAmendPrebillingBillFromParams `form:"bill_from" json:"bill_from,omitempty"`
	// The end of the prebilled time period.
	BillUntil *SubscriptionScheduleAmendPrebillingBillUntilParams `form:"bill_until" json:"bill_until,omitempty"`
	// When the prebilling invoice should be created. The default value is `now`.
	InvoiceAt *string `form:"invoice_at" json:"invoice_at,omitempty"`
	// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period. The default value is `reset`.
	UpdateBehavior *string `form:"update_behavior" json:"update_behavior,omitempty"`
}

// Changes to apply to the subscription schedule.
type SubscriptionScheduleAmendScheduleSettingsParams struct {
	// Behavior of the subscription schedule and underlying subscription when it ends.
	EndBehavior *string `form:"end_behavior" json:"end_behavior,omitempty"`
}

// Amends an existing subscription schedule.
type SubscriptionScheduleAmendParams struct {
	Params `form:"*"`
	// Changes to apply to the phases of the subscription schedule, in the order provided.
	Amendments []*SubscriptionScheduleAmendAmendmentParams `form:"amendments" json:"amendments,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Provide any time periods to bill in advance.
	Prebilling []*SubscriptionScheduleAmendPrebillingParams `form:"prebilling" json:"prebilling,omitempty"`
	// In cases where the amendment changes the currently active phase,
	//  specifies if and how to prorate at the time of the request.
	ProrationBehavior *string `form:"proration_behavior" json:"proration_behavior,omitempty"`
	// Changes to apply to the subscription schedule.
	ScheduleSettings *SubscriptionScheduleAmendScheduleSettingsParams `form:"schedule_settings" json:"schedule_settings,omitempty"`
	UnsetFields      []SubscriptionScheduleAmendParamsUnsetField      `form:"-" json:"-"`
}

// SubscriptionScheduleAmendParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleAmendParams.
type SubscriptionScheduleAmendParamsUnsetField string

const (
	SubscriptionScheduleAmendParamsUnsetFieldPrebilling SubscriptionScheduleAmendParamsUnsetField = "prebilling"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleAmendParams) AddUnsetField(field SubscriptionScheduleAmendParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleAmendParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Cancels a subscription schedule and its associated subscription immediately (if the subscription schedule has an active subscription). A subscription schedule can only be canceled if its status is not_started or active.
type SubscriptionScheduleCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// If the subscription schedule is `active`, indicates if a final invoice will be generated that contains any un-invoiced metered usage and new/pending proration invoice items. Defaults to `true`.
	InvoiceNow *bool `form:"invoice_now" json:"invoice_now,omitempty"`
	// If the subscription schedule is `active`, indicates if the cancellation should be prorated. Defaults to `true`.
	Prorate *bool `form:"prorate" json:"prorate,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Releases the subscription schedule immediately, which will stop scheduling of its phases, but leave any existing subscription in place. A schedule can only be released if its status is not_started or active. If the subscription schedule is currently associated with a subscription, releasing it will remove its subscription property and set the subscription's ID to the released_subscription property.
type SubscriptionScheduleReleaseParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Keep any cancellation on the subscription that the schedule has set
	PreserveCancelDate *bool `form:"preserve_cancel_date" json:"preserve_cancel_date,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleReleaseParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configure behavior for flexible billing mode.
type SubscriptionScheduleCreateBillingModeFlexibleParams struct {
	// Controls how invoices and invoice items display proration amounts and discount amounts.
	ProrationDiscounts *string `form:"proration_discounts" json:"proration_discounts,omitempty"`
}

// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
type SubscriptionScheduleCreateBillingModeParams struct {
	// Configure behavior for flexible billing mode.
	Flexible *SubscriptionScheduleCreateBillingModeFlexibleParams `form:"flexible" json:"flexible,omitempty"`
	// Controls the calculation and orchestration of prorations and invoices for subscriptions. If no value is passed, the default is `flexible`.
	Type *string `form:"type" json:"type"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleCreateDefaultSettingsBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte" json:"amount_gte,omitempty"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor" json:"reset_billing_cycle_anchor,omitempty"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParams struct {
	// The account tax IDs associated with the subscription schedule. Will be set on invoices generated by the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids" json:"account_tax_ids,omitempty"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `collection_method=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due" json:"days_until_due,omitempty"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer      *SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsIssuerParams      `form:"issuer" json:"issuer,omitempty"`
	UnsetFields []SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParams.
type SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParamsUnsetField string

const (
	SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParamsUnsetFieldAccountTaxIDs SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParamsUnsetField = "account_tax_ids"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParams) AddUnsetField(field SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Object representing the subscription schedule's default settings.
type SubscriptionScheduleCreateDefaultSettingsParams struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent,high_precision"`
	// Default settings for automatic tax computation.
	AutomaticTax *SubscriptionAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleCreateDefaultSettingsBillingThresholdsParams `form:"billing_thresholds" json:"billing_thresholds,omitempty"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description" json:"description,omitempty"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleCreateDefaultSettingsInvoiceSettingsParams `form:"invoice_settings" json:"invoice_settings,omitempty"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams                             `form:"transfer_data" json:"transfer_data,omitempty"`
	UnsetFields  []SubscriptionScheduleCreateDefaultSettingsParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleCreateDefaultSettingsParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleCreateDefaultSettingsParams.
type SubscriptionScheduleCreateDefaultSettingsParamsUnsetField string

const (
	SubscriptionScheduleCreateDefaultSettingsParamsUnsetFieldBillingThresholds SubscriptionScheduleCreateDefaultSettingsParamsUnsetField = "billing_thresholds"
	SubscriptionScheduleCreateDefaultSettingsParamsUnsetFieldDescription       SubscriptionScheduleCreateDefaultSettingsParamsUnsetField = "description"
	SubscriptionScheduleCreateDefaultSettingsParamsUnsetFieldOnBehalfOf        SubscriptionScheduleCreateDefaultSettingsParamsUnsetField = "on_behalf_of"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleCreateDefaultSettingsParams) AddUnsetField(field SubscriptionScheduleCreateDefaultSettingsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Time span for the redeemed discount.
type SubscriptionScheduleCreatePhaseAddInvoiceItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleCreatePhaseAddInvoiceItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionScheduleCreatePhaseAddInvoiceItemDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// The coupons to redeem into discounts for the item.
type SubscriptionScheduleCreatePhaseAddInvoiceItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleCreatePhaseAddInvoiceItemDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// End of the invoice item period.
type SubscriptionScheduleCreatePhaseAddInvoiceItemPeriodEndParams struct {
	// A precise Unix timestamp for the end of the invoice item period. Must be greater than or equal to `period.start`.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// Select how to calculate the end of the invoice item period.
	Type *string `form:"type" json:"type"`
}

// Start of the invoice item period.
type SubscriptionScheduleCreatePhaseAddInvoiceItemPeriodStartParams struct {
	// A precise Unix timestamp for the start of the invoice item period. Must be less than or equal to `period.end`.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// Select how to calculate the start of the invoice item period.
	Type *string `form:"type" json:"type"`
}

// The period associated with this invoice item. If not set, `period.start.type` defaults to `max_item_period_start` and `period.end.type` defaults to `min_item_period_end`.
type SubscriptionScheduleCreatePhaseAddInvoiceItemPeriodParams struct {
	// End of the invoice item period.
	End *SubscriptionScheduleCreatePhaseAddInvoiceItemPeriodEndParams `form:"end" json:"end"`
	// Start of the invoice item period.
	Start *SubscriptionScheduleCreatePhaseAddInvoiceItemPeriodStartParams `form:"start" json:"start"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
type SubscriptionScheduleCreatePhaseAddInvoiceItemParams struct {
	// The coupons to redeem into discounts for the item.
	Discounts []*SubscriptionScheduleCreatePhaseAddInvoiceItemDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The period associated with this invoice item. If not set, `period.start.type` defaults to `max_item_period_start` and `period.end.type` defaults to `min_item_period_end`.
	Period *SubscriptionScheduleCreatePhaseAddInvoiceItemPeriodParams `form:"period" json:"period,omitempty"`
	// The ID of the price object. One of `price` or `price_data` is required.
	Price *string `form:"price" json:"price,omitempty"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *InvoiceItemPriceDataParams `form:"price_data" json:"price_data,omitempty"`
	// Quantity for this item. Defaults to 1.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates    []*string                                                       `form:"tax_rates" json:"tax_rates,omitempty"`
	UnsetFields []SubscriptionScheduleCreatePhaseAddInvoiceItemParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleCreatePhaseAddInvoiceItemParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleCreatePhaseAddInvoiceItemParams.
type SubscriptionScheduleCreatePhaseAddInvoiceItemParamsUnsetField string

const (
	SubscriptionScheduleCreatePhaseAddInvoiceItemParamsUnsetFieldTaxRates SubscriptionScheduleCreatePhaseAddInvoiceItemParamsUnsetField = "tax_rates"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleCreatePhaseAddInvoiceItemParams) AddUnsetField(field SubscriptionScheduleCreatePhaseAddInvoiceItemParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleCreatePhaseAddInvoiceItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type SubscriptionScheduleCreatePhaseAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// Automatic tax settings for this phase.
type SubscriptionScheduleCreatePhaseAutomaticTaxParams struct {
	// Enabled automatic tax calculation which will automatically compute tax rates on all invoices generated by the subscription.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *SubscriptionScheduleCreatePhaseAutomaticTaxLiabilityParams `form:"liability" json:"liability,omitempty"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleCreatePhaseBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte" json:"amount_gte,omitempty"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor" json:"reset_billing_cycle_anchor,omitempty"`
}

// Time span for the redeemed discount.
type SubscriptionScheduleCreatePhaseDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleCreatePhaseDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionScheduleCreatePhaseDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
type SubscriptionScheduleCreatePhaseDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleCreatePhaseDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// The number of intervals the phase should last. If set, `end_date` must not be set.
type SubscriptionScheduleCreatePhaseDurationParams struct {
	// Specifies phase duration. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The multiplier applied to the interval.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleCreatePhaseInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleCreatePhaseInvoiceSettingsParams struct {
	// The account tax IDs associated with this phase of the subscription schedule. Will be set on invoices generated by this phase of the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids" json:"account_tax_ids,omitempty"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due" json:"days_until_due,omitempty"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer      *SubscriptionScheduleCreatePhaseInvoiceSettingsIssuerParams      `form:"issuer" json:"issuer,omitempty"`
	UnsetFields []SubscriptionScheduleCreatePhaseInvoiceSettingsParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleCreatePhaseInvoiceSettingsParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleCreatePhaseInvoiceSettingsParams.
type SubscriptionScheduleCreatePhaseInvoiceSettingsParamsUnsetField string

const (
	SubscriptionScheduleCreatePhaseInvoiceSettingsParamsUnsetFieldAccountTaxIDs SubscriptionScheduleCreatePhaseInvoiceSettingsParamsUnsetField = "account_tax_ids"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleCreatePhaseInvoiceSettingsParams) AddUnsetField(field SubscriptionScheduleCreatePhaseInvoiceSettingsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleCreatePhaseItemBillingThresholdsParams struct {
	// Number of units that meets the billing threshold to advance the subscription to a new billing period (e.g., it takes 10 $5 units to meet a $50 [monetary threshold](https://docs.stripe.com/api/subscriptions/update#update_subscription-billing_thresholds-amount_gte))
	UsageGTE *int64 `form:"usage_gte" json:"usage_gte"`
}

// Time span for the redeemed discount.
type SubscriptionScheduleCreatePhaseItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleCreatePhaseItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionScheduleCreatePhaseItemDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// The coupons to redeem into discounts for the subscription item.
type SubscriptionScheduleCreatePhaseItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleCreatePhaseItemDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// Options that configure the trial on the subscription item.
type SubscriptionScheduleCreatePhaseItemTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial. Currently only supports at most 1 price ID.
	ConvertsTo []*string `form:"converts_to" json:"converts_to,omitempty"`
	// Determines the type of trial for this item.
	Type *string `form:"type" json:"type"`
}

// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
type SubscriptionScheduleCreatePhaseItemParams struct {
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleCreatePhaseItemBillingThresholdsParams `form:"billing_thresholds" json:"billing_thresholds,omitempty"`
	// The coupons to redeem into discounts for the subscription item.
	Discounts []*SubscriptionScheduleCreatePhaseItemDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to a configuration item. Metadata on a configuration item will update the underlying subscription item's `metadata` when the phase is entered, adding new keys and replacing existing keys. Individual keys in the subscription item's `metadata` can be unset by posting an empty value to them in the configuration item's `metadata`. To unset all keys in the subscription item's `metadata`, update the subscription item directly or unset every key individually from the configuration item's `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The plan ID to subscribe to. You may specify the same ID in `plan` and `price`.
	Plan *string `form:"plan" json:"plan,omitempty"`
	// The ID of the price object.
	Price *string `form:"price" json:"price,omitempty"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline.
	PriceData *SubscriptionItemPriceDataParams `form:"price_data" json:"price_data,omitempty"`
	// Quantity for the given price. Can be set only if the price's `usage_type` is `licensed` and not `metered`.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
	// A list of [Tax Rate](https://docs.stripe.com/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://docs.stripe.com/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates" json:"tax_rates,omitempty"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionScheduleCreatePhaseItemTrialParams `form:"trial" json:"trial,omitempty"`
	// The ID of the trial offer to apply to the configuration item.
	TrialOffer  *string                                               `form:"trial_offer" json:"trial_offer,omitempty"`
	UnsetFields []SubscriptionScheduleCreatePhaseItemParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleCreatePhaseItemParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleCreatePhaseItemParams.
type SubscriptionScheduleCreatePhaseItemParamsUnsetField string

const (
	SubscriptionScheduleCreatePhaseItemParamsUnsetFieldBillingThresholds SubscriptionScheduleCreatePhaseItemParamsUnsetField = "billing_thresholds"
	SubscriptionScheduleCreatePhaseItemParamsUnsetFieldDiscounts         SubscriptionScheduleCreatePhaseItemParamsUnsetField = "discounts"
	SubscriptionScheduleCreatePhaseItemParamsUnsetFieldTaxRates          SubscriptionScheduleCreatePhaseItemParamsUnsetField = "tax_rates"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleCreatePhaseItemParams) AddUnsetField(field SubscriptionScheduleCreatePhaseItemParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleCreatePhaseItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://docs.stripe.com/billing/subscriptions/pause-payment).
type SubscriptionScheduleCreatePhasePauseCollectionParams struct {
	// The payment collection behavior for this subscription while paused.
	Behavior *string `form:"behavior" json:"behavior"`
}

// Defines how the subscription should behave when a trial ends.
type SubscriptionScheduleCreatePhaseTrialSettingsEndBehaviorParams struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront *string `form:"prorate_up_front" json:"prorate_up_front,omitempty"`
}

// Settings related to subscription trials.
type SubscriptionScheduleCreatePhaseTrialSettingsParams struct {
	// Defines how the subscription should behave when a trial ends.
	EndBehavior *SubscriptionScheduleCreatePhaseTrialSettingsEndBehaviorParams `form:"end_behavior" json:"end_behavior,omitempty"`
}

// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase.
type SubscriptionScheduleCreatePhaseParams struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
	AddInvoiceItems []*SubscriptionScheduleCreatePhaseAddInvoiceItemParams `form:"add_invoice_items" json:"add_invoice_items,omitempty"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent" json:"application_fee_percent,omitempty"`
	// Automatic tax settings for this phase.
	AutomaticTax *SubscriptionScheduleCreatePhaseAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleCreatePhaseBillingThresholdsParams `form:"billing_thresholds" json:"billing_thresholds,omitempty"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// A list of [Tax Rate](https://docs.stripe.com/api/tax_rates) ids. These Tax Rates will set the Subscription's [`default_tax_rates`](https://docs.stripe.com/api/subscriptions/create#create_subscription-default_tax_rates), which means they will be the Invoice's [`default_tax_rates`](https://docs.stripe.com/api/invoices/create#create_invoice-default_tax_rates) for any Invoices issued by the Subscription during this Phase.
	DefaultTaxRates []*string `form:"default_tax_rates" json:"default_tax_rates,omitempty"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description" json:"description,omitempty"`
	// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
	Discounts []*SubscriptionScheduleCreatePhaseDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// The number of intervals the phase should last. If set, `end_date` must not be set.
	Duration *SubscriptionScheduleCreatePhaseDurationParams `form:"duration" json:"duration,omitempty"`
	// The date at which this phase of the subscription schedule ends. If set, `duration` must not be set.
	EndDate *int64 `form:"end_date" json:"end_date,omitempty"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleCreatePhaseInvoiceSettingsParams `form:"invoice_settings" json:"invoice_settings,omitempty"`
	// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
	Items []*SubscriptionScheduleCreatePhaseItemParams `form:"items" json:"items"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered, adding new keys and replacing existing keys in the subscription's `metadata`. Individual keys in the subscription's `metadata` can be unset by posting an empty value to them in the phase's `metadata`. To unset all keys in the subscription's `metadata`, update the subscription directly or unset every key individually from the phase's `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://docs.stripe.com/billing/subscriptions/pause-payment).
	PauseCollection *SubscriptionScheduleCreatePhasePauseCollectionParams `form:"pause_collection" json:"pause_collection,omitempty"`
	// Controls whether the subscription schedule should create [prorations](https://docs.stripe.com/billing/subscriptions/prorations) when transitioning to this phase if there is a difference in billing configuration. It's different from the request-level [proration_behavior](https://docs.stripe.com/api/subscription_schedules/update#update_subscription_schedule-proration_behavior) parameter which controls what happens if the update request affects the billing configuration (item price, quantity, etc.) of the current phase.
	ProrationBehavior *string `form:"proration_behavior" json:"proration_behavior,omitempty"`
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
	// If set to true the entire phase is counted as a trial and the customer will not be charged for any fees.
	Trial *bool `form:"trial" json:"trial,omitempty"`
	// Specify trial behavior when crossing phase boundaries
	TrialContinuation *string `form:"trial_continuation" json:"trial_continuation,omitempty"`
	// Sets the phase to trialing from the start date to this date. Must be before the phase end date, can not be combined with `trial`
	TrialEnd *int64 `form:"trial_end" json:"trial_end,omitempty"`
	// Settings related to subscription trials.
	TrialSettings *SubscriptionScheduleCreatePhaseTrialSettingsParams `form:"trial_settings" json:"trial_settings,omitempty"`
	UnsetFields   []SubscriptionScheduleCreatePhaseParamsUnsetField   `form:"-" json:"-"`
}

// SubscriptionScheduleCreatePhaseParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleCreatePhaseParams.
type SubscriptionScheduleCreatePhaseParamsUnsetField string

const (
	SubscriptionScheduleCreatePhaseParamsUnsetFieldBillingThresholds SubscriptionScheduleCreatePhaseParamsUnsetField = "billing_thresholds"
	SubscriptionScheduleCreatePhaseParamsUnsetFieldDefaultTaxRates   SubscriptionScheduleCreatePhaseParamsUnsetField = "default_tax_rates"
	SubscriptionScheduleCreatePhaseParamsUnsetFieldDescription       SubscriptionScheduleCreatePhaseParamsUnsetField = "description"
	SubscriptionScheduleCreatePhaseParamsUnsetFieldDiscounts         SubscriptionScheduleCreatePhaseParamsUnsetField = "discounts"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleCreatePhaseParams) AddUnsetField(field SubscriptionScheduleCreatePhaseParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleCreatePhaseParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// If specified, the invoicing for the given billing cycle iterations will be processed now.
type SubscriptionScheduleCreatePrebillingParams struct {
	// This is used to determine the number of billing cycles to prebill.
	Iterations *int64 `form:"iterations" json:"iterations"`
	// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period. The default value is `reset`.
	UpdateBehavior *string `form:"update_behavior" json:"update_behavior,omitempty"`
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
type SubscriptionScheduleCreateParams struct {
	Params `form:"*"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time. `prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior *string `form:"billing_behavior" json:"billing_behavior,omitempty"`
	// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
	BillingMode *SubscriptionScheduleCreateBillingModeParams `form:"billing_mode" json:"billing_mode,omitempty"`
	// The identifier of the customer to create the subscription schedule for.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The identifier of the account to create the subscription schedule for.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// Object representing the subscription schedule's default settings.
	DefaultSettings *SubscriptionScheduleCreateDefaultSettingsParams `form:"default_settings" json:"default_settings,omitempty"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior *string `form:"end_behavior" json:"end_behavior,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Migrate an existing subscription to be managed by a subscription schedule. If this parameter is set, a subscription schedule will be created using the subscription's item(s), set to auto-renew using the subscription's interval. When using this parameter, other parameters (such as phase values) cannot be set. To create a subscription schedule with other modifications, we recommend making two separate API calls.
	FromSubscription *string `form:"from_subscription" json:"from_subscription,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase.
	Phases []*SubscriptionScheduleCreatePhaseParams `form:"phases" json:"phases,omitempty"`
	// If specified, the invoicing for the given billing cycle iterations will be processed now.
	Prebilling *SubscriptionScheduleCreatePrebillingParams `form:"prebilling" json:"prebilling,omitempty"`
	// When the subscription schedule starts. We recommend using `now` so that it starts the subscription immediately. You can also use a Unix timestamp to backdate the subscription so that it starts on a past date, or set a future date for the subscription to start on.
	StartDate    *int64                                       `form:"start_date" json:"start_date,omitempty"`
	StartDateNow *bool                                        `form:"-"` // See custom AppendTo
	UnsetFields  []SubscriptionScheduleCreateParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleCreateParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleCreateParams.
type SubscriptionScheduleCreateParamsUnsetField string

const (
	SubscriptionScheduleCreateParamsUnsetFieldMetadata SubscriptionScheduleCreateParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleCreateParams) AddUnsetField(field SubscriptionScheduleCreateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// AppendTo implements custom encoding logic for SubscriptionScheduleCreateParams.
func (p *SubscriptionScheduleCreateParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// Retrieves the details of an existing subscription schedule. You only need to supply the unique subscription schedule identifier that was returned upon subscription schedule creation.
type SubscriptionScheduleRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleUpdateDefaultSettingsBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte" json:"amount_gte,omitempty"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor" json:"reset_billing_cycle_anchor,omitempty"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParams struct {
	// The account tax IDs associated with the subscription schedule. Will be set on invoices generated by the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids" json:"account_tax_ids,omitempty"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `collection_method=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due" json:"days_until_due,omitempty"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer      *SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsIssuerParams      `form:"issuer" json:"issuer,omitempty"`
	UnsetFields []SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParams.
type SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParamsUnsetField string

const (
	SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParamsUnsetFieldAccountTaxIDs SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParamsUnsetField = "account_tax_ids"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParams) AddUnsetField(field SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Object representing the subscription schedule's default settings.
type SubscriptionScheduleUpdateDefaultSettingsParams struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent,high_precision"`
	// Default settings for automatic tax computation.
	AutomaticTax *SubscriptionAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleUpdateDefaultSettingsBillingThresholdsParams `form:"billing_thresholds" json:"billing_thresholds,omitempty"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description" json:"description,omitempty"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleUpdateDefaultSettingsInvoiceSettingsParams `form:"invoice_settings" json:"invoice_settings,omitempty"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams                             `form:"transfer_data" json:"transfer_data,omitempty"`
	UnsetFields  []SubscriptionScheduleUpdateDefaultSettingsParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleUpdateDefaultSettingsParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleUpdateDefaultSettingsParams.
type SubscriptionScheduleUpdateDefaultSettingsParamsUnsetField string

const (
	SubscriptionScheduleUpdateDefaultSettingsParamsUnsetFieldBillingThresholds SubscriptionScheduleUpdateDefaultSettingsParamsUnsetField = "billing_thresholds"
	SubscriptionScheduleUpdateDefaultSettingsParamsUnsetFieldDescription       SubscriptionScheduleUpdateDefaultSettingsParamsUnsetField = "description"
	SubscriptionScheduleUpdateDefaultSettingsParamsUnsetFieldOnBehalfOf        SubscriptionScheduleUpdateDefaultSettingsParamsUnsetField = "on_behalf_of"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleUpdateDefaultSettingsParams) AddUnsetField(field SubscriptionScheduleUpdateDefaultSettingsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Time span for the redeemed discount.
type SubscriptionScheduleUpdatePhaseAddInvoiceItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleUpdatePhaseAddInvoiceItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionScheduleUpdatePhaseAddInvoiceItemDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// The coupons to redeem into discounts for the item.
type SubscriptionScheduleUpdatePhaseAddInvoiceItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleUpdatePhaseAddInvoiceItemDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// End of the invoice item period.
type SubscriptionScheduleUpdatePhaseAddInvoiceItemPeriodEndParams struct {
	// A precise Unix timestamp for the end of the invoice item period. Must be greater than or equal to `period.start`.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// Select how to calculate the end of the invoice item period.
	Type *string `form:"type" json:"type"`
}

// Start of the invoice item period.
type SubscriptionScheduleUpdatePhaseAddInvoiceItemPeriodStartParams struct {
	// A precise Unix timestamp for the start of the invoice item period. Must be less than or equal to `period.end`.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// Select how to calculate the start of the invoice item period.
	Type *string `form:"type" json:"type"`
}

// The period associated with this invoice item. If not set, `period.start.type` defaults to `max_item_period_start` and `period.end.type` defaults to `min_item_period_end`.
type SubscriptionScheduleUpdatePhaseAddInvoiceItemPeriodParams struct {
	// End of the invoice item period.
	End *SubscriptionScheduleUpdatePhaseAddInvoiceItemPeriodEndParams `form:"end" json:"end"`
	// Start of the invoice item period.
	Start *SubscriptionScheduleUpdatePhaseAddInvoiceItemPeriodStartParams `form:"start" json:"start"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
type SubscriptionScheduleUpdatePhaseAddInvoiceItemParams struct {
	// The coupons to redeem into discounts for the item.
	Discounts []*SubscriptionScheduleUpdatePhaseAddInvoiceItemDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The period associated with this invoice item. If not set, `period.start.type` defaults to `max_item_period_start` and `period.end.type` defaults to `min_item_period_end`.
	Period *SubscriptionScheduleUpdatePhaseAddInvoiceItemPeriodParams `form:"period" json:"period,omitempty"`
	// The ID of the price object. One of `price` or `price_data` is required.
	Price *string `form:"price" json:"price,omitempty"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *InvoiceItemPriceDataParams `form:"price_data" json:"price_data,omitempty"`
	// Quantity for this item. Defaults to 1.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates    []*string                                                       `form:"tax_rates" json:"tax_rates,omitempty"`
	UnsetFields []SubscriptionScheduleUpdatePhaseAddInvoiceItemParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleUpdatePhaseAddInvoiceItemParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleUpdatePhaseAddInvoiceItemParams.
type SubscriptionScheduleUpdatePhaseAddInvoiceItemParamsUnsetField string

const (
	SubscriptionScheduleUpdatePhaseAddInvoiceItemParamsUnsetFieldTaxRates SubscriptionScheduleUpdatePhaseAddInvoiceItemParamsUnsetField = "tax_rates"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleUpdatePhaseAddInvoiceItemParams) AddUnsetField(field SubscriptionScheduleUpdatePhaseAddInvoiceItemParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleUpdatePhaseAddInvoiceItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type SubscriptionScheduleUpdatePhaseAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// Automatic tax settings for this phase.
type SubscriptionScheduleUpdatePhaseAutomaticTaxParams struct {
	// Enabled automatic tax calculation which will automatically compute tax rates on all invoices generated by the subscription.
	Enabled *bool `form:"enabled" json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *SubscriptionScheduleUpdatePhaseAutomaticTaxLiabilityParams `form:"liability" json:"liability,omitempty"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleUpdatePhaseBillingThresholdsParams struct {
	// Monetary threshold that triggers the subscription to advance to a new billing period
	AmountGTE *int64 `form:"amount_gte" json:"amount_gte,omitempty"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged.
	ResetBillingCycleAnchor *bool `form:"reset_billing_cycle_anchor" json:"reset_billing_cycle_anchor,omitempty"`
}

// Time span for the redeemed discount.
type SubscriptionScheduleUpdatePhaseDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleUpdatePhaseDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionScheduleUpdatePhaseDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
type SubscriptionScheduleUpdatePhaseDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleUpdatePhaseDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// The number of intervals the phase should last. If set, `end_date` must not be set.
type SubscriptionScheduleUpdatePhaseDurationParams struct {
	// Specifies phase duration. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The multiplier applied to the interval.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionScheduleUpdatePhaseInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// Type of the account referenced in the request.
	Type *string `form:"type" json:"type"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleUpdatePhaseInvoiceSettingsParams struct {
	// The account tax IDs associated with this phase of the subscription schedule. Will be set on invoices generated by this phase of the subscription schedule.
	AccountTaxIDs []*string `form:"account_tax_ids" json:"account_tax_ids,omitempty"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due" json:"days_until_due,omitempty"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer      *SubscriptionScheduleUpdatePhaseInvoiceSettingsIssuerParams      `form:"issuer" json:"issuer,omitempty"`
	UnsetFields []SubscriptionScheduleUpdatePhaseInvoiceSettingsParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleUpdatePhaseInvoiceSettingsParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleUpdatePhaseInvoiceSettingsParams.
type SubscriptionScheduleUpdatePhaseInvoiceSettingsParamsUnsetField string

const (
	SubscriptionScheduleUpdatePhaseInvoiceSettingsParamsUnsetFieldAccountTaxIDs SubscriptionScheduleUpdatePhaseInvoiceSettingsParamsUnsetField = "account_tax_ids"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleUpdatePhaseInvoiceSettingsParams) AddUnsetField(field SubscriptionScheduleUpdatePhaseInvoiceSettingsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
type SubscriptionScheduleUpdatePhaseItemBillingThresholdsParams struct {
	// Number of units that meets the billing threshold to advance the subscription to a new billing period (e.g., it takes 10 $5 units to meet a $50 [monetary threshold](https://docs.stripe.com/api/subscriptions/update#update_subscription-billing_thresholds-amount_gte))
	UsageGTE *int64 `form:"usage_gte" json:"usage_gte"`
}

// Time span for the redeemed discount.
type SubscriptionScheduleUpdatePhaseItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval" json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type SubscriptionScheduleUpdatePhaseItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *SubscriptionScheduleUpdatePhaseItemDiscountDiscountEndDurationParams `form:"duration" json:"duration,omitempty"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type" json:"type"`
}

// The coupons to redeem into discounts for the subscription item.
type SubscriptionScheduleUpdatePhaseItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount" json:"discount,omitempty"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *SubscriptionScheduleUpdatePhaseItemDiscountDiscountEndParams `form:"discount_end" json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
}

// Options that configure the trial on the subscription item.
type SubscriptionScheduleUpdatePhaseItemTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial. Currently only supports at most 1 price ID.
	ConvertsTo []*string `form:"converts_to" json:"converts_to,omitempty"`
	// Determines the type of trial for this item.
	Type *string `form:"type" json:"type"`
}

// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
type SubscriptionScheduleUpdatePhaseItemParams struct {
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleUpdatePhaseItemBillingThresholdsParams `form:"billing_thresholds" json:"billing_thresholds,omitempty"`
	// The coupons to redeem into discounts for the subscription item.
	Discounts []*SubscriptionScheduleUpdatePhaseItemDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to a configuration item. Metadata on a configuration item will update the underlying subscription item's `metadata` when the phase is entered, adding new keys and replacing existing keys. Individual keys in the subscription item's `metadata` can be unset by posting an empty value to them in the configuration item's `metadata`. To unset all keys in the subscription item's `metadata`, update the subscription item directly or unset every key individually from the configuration item's `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The plan ID to subscribe to. You may specify the same ID in `plan` and `price`.
	Plan *string `form:"plan" json:"plan,omitempty"`
	// The ID of the price object.
	Price *string `form:"price" json:"price,omitempty"`
	// Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline.
	PriceData *SubscriptionItemPriceDataParams `form:"price_data" json:"price_data,omitempty"`
	// Quantity for the given price. Can be set only if the price's `usage_type` is `licensed` and not `metered`.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
	// A list of [Tax Rate](https://docs.stripe.com/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://docs.stripe.com/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.
	TaxRates []*string `form:"tax_rates" json:"tax_rates,omitempty"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionScheduleUpdatePhaseItemTrialParams `form:"trial" json:"trial,omitempty"`
	// The ID of the trial offer to apply to the configuration item.
	TrialOffer  *string                                               `form:"trial_offer" json:"trial_offer,omitempty"`
	UnsetFields []SubscriptionScheduleUpdatePhaseItemParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleUpdatePhaseItemParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleUpdatePhaseItemParams.
type SubscriptionScheduleUpdatePhaseItemParamsUnsetField string

const (
	SubscriptionScheduleUpdatePhaseItemParamsUnsetFieldBillingThresholds SubscriptionScheduleUpdatePhaseItemParamsUnsetField = "billing_thresholds"
	SubscriptionScheduleUpdatePhaseItemParamsUnsetFieldDiscounts         SubscriptionScheduleUpdatePhaseItemParamsUnsetField = "discounts"
	SubscriptionScheduleUpdatePhaseItemParamsUnsetFieldTaxRates          SubscriptionScheduleUpdatePhaseItemParamsUnsetField = "tax_rates"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleUpdatePhaseItemParams) AddUnsetField(field SubscriptionScheduleUpdatePhaseItemParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleUpdatePhaseItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://docs.stripe.com/billing/subscriptions/pause-payment).
type SubscriptionScheduleUpdatePhasePauseCollectionParams struct {
	// The payment collection behavior for this subscription while paused.
	Behavior *string `form:"behavior" json:"behavior"`
}

// Defines how the subscription should behave when a trial ends.
type SubscriptionScheduleUpdatePhaseTrialSettingsEndBehaviorParams struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront *string `form:"prorate_up_front" json:"prorate_up_front,omitempty"`
}

// Settings related to subscription trials.
type SubscriptionScheduleUpdatePhaseTrialSettingsParams struct {
	// Defines how the subscription should behave when a trial ends.
	EndBehavior *SubscriptionScheduleUpdatePhaseTrialSettingsEndBehaviorParams `form:"end_behavior" json:"end_behavior,omitempty"`
}

// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase. Note that past phases can be omitted.
type SubscriptionScheduleUpdatePhaseParams struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase. You may pass up to 20 items.
	AddInvoiceItems []*SubscriptionScheduleUpdatePhaseAddInvoiceItemParams `form:"add_invoice_items" json:"add_invoice_items,omitempty"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).
	ApplicationFeePercent *float64 `form:"application_fee_percent" json:"application_fee_percent,omitempty"`
	// Automatic tax settings for this phase.
	AutomaticTax *SubscriptionScheduleUpdatePhaseAutomaticTaxParams `form:"automatic_tax" json:"automatic_tax,omitempty"`
	// Can be set to `phase_start` to set the anchor to the start of the phase or `automatic` to automatically change it if needed. Cannot be set to `phase_start` if this phase specifies a trial. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).
	BillingCycleAnchor *string `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.
	BillingThresholds *SubscriptionScheduleUpdatePhaseBillingThresholdsParams `form:"billing_thresholds" json:"billing_thresholds,omitempty"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically` on creation.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// A list of [Tax Rate](https://docs.stripe.com/api/tax_rates) ids. These Tax Rates will set the Subscription's [`default_tax_rates`](https://docs.stripe.com/api/subscriptions/create#create_subscription-default_tax_rates), which means they will be the Invoice's [`default_tax_rates`](https://docs.stripe.com/api/invoices/create#create_invoice-default_tax_rates) for any Invoices issued by the Subscription during this Phase.
	DefaultTaxRates []*string `form:"default_tax_rates" json:"default_tax_rates,omitempty"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description" json:"description,omitempty"`
	// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
	Discounts []*SubscriptionScheduleUpdatePhaseDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// The number of intervals the phase should last. If set, `end_date` must not be set.
	Duration *SubscriptionScheduleUpdatePhaseDurationParams `form:"duration" json:"duration,omitempty"`
	// The date at which this phase of the subscription schedule ends. If set, `duration` must not be set.
	EndDate    *int64 `form:"end_date" json:"end_date,omitempty"`
	EndDateNow *bool  `form:"-"` // See custom AppendTo
	// All invoices will be billed using the specified settings.
	InvoiceSettings *SubscriptionScheduleUpdatePhaseInvoiceSettingsParams `form:"invoice_settings" json:"invoice_settings,omitempty"`
	// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
	Items []*SubscriptionScheduleUpdatePhaseItemParams `form:"items" json:"items"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered, adding new keys and replacing existing keys in the subscription's `metadata`. Individual keys in the subscription's `metadata` can be unset by posting an empty value to them in the phase's `metadata`. To unset all keys in the subscription's `metadata`, update the subscription directly or unset every key individually from the phase's `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The account on behalf of which to charge, for each of the associated subscription's invoices.
	OnBehalfOf *string `form:"on_behalf_of" json:"on_behalf_of,omitempty"`
	// If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://docs.stripe.com/billing/subscriptions/pause-payment).
	PauseCollection *SubscriptionScheduleUpdatePhasePauseCollectionParams `form:"pause_collection" json:"pause_collection,omitempty"`
	// Controls whether the subscription schedule should create [prorations](https://docs.stripe.com/billing/subscriptions/prorations) when transitioning to this phase if there is a difference in billing configuration. It's different from the request-level [proration_behavior](https://docs.stripe.com/api/subscription_schedules/update#update_subscription_schedule-proration_behavior) parameter which controls what happens if the update request affects the billing configuration (item price, quantity, etc.) of the current phase.
	ProrationBehavior *string `form:"proration_behavior" json:"proration_behavior,omitempty"`
	// The date at which this phase of the subscription schedule starts or `now`. Must be set on the first phase.
	StartDate    *int64 `form:"start_date" json:"start_date,omitempty"`
	StartDateNow *bool  `form:"-"` // See custom AppendTo
	// The data with which to automatically create a Transfer for each of the associated subscription's invoices.
	TransferData *SubscriptionTransferDataParams `form:"transfer_data" json:"transfer_data,omitempty"`
	// If set to true the entire phase is counted as a trial and the customer will not be charged for any fees.
	Trial *bool `form:"trial" json:"trial,omitempty"`
	// Specify trial behavior when crossing phase boundaries
	TrialContinuation *string `form:"trial_continuation" json:"trial_continuation,omitempty"`
	// Sets the phase to trialing from the start date to this date. Must be before the phase end date, can not be combined with `trial`
	TrialEnd    *int64 `form:"trial_end" json:"trial_end,omitempty"`
	TrialEndNow *bool  `form:"-"` // See custom AppendTo
	// Settings related to subscription trials.
	TrialSettings *SubscriptionScheduleUpdatePhaseTrialSettingsParams `form:"trial_settings" json:"trial_settings,omitempty"`
	UnsetFields   []SubscriptionScheduleUpdatePhaseParamsUnsetField   `form:"-" json:"-"`
}

// SubscriptionScheduleUpdatePhaseParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleUpdatePhaseParams.
type SubscriptionScheduleUpdatePhaseParamsUnsetField string

const (
	SubscriptionScheduleUpdatePhaseParamsUnsetFieldBillingThresholds SubscriptionScheduleUpdatePhaseParamsUnsetField = "billing_thresholds"
	SubscriptionScheduleUpdatePhaseParamsUnsetFieldDefaultTaxRates   SubscriptionScheduleUpdatePhaseParamsUnsetField = "default_tax_rates"
	SubscriptionScheduleUpdatePhaseParamsUnsetFieldDescription       SubscriptionScheduleUpdatePhaseParamsUnsetField = "description"
	SubscriptionScheduleUpdatePhaseParamsUnsetFieldDiscounts         SubscriptionScheduleUpdatePhaseParamsUnsetField = "discounts"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleUpdatePhaseParams) AddUnsetField(field SubscriptionScheduleUpdatePhaseParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleUpdatePhaseParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// AppendTo implements custom encoding logic for SubscriptionScheduleUpdatePhaseParams.
func (p *SubscriptionScheduleUpdatePhaseParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.EndDateNow) {
		body.Add(form.FormatKey(append(keyParts, "end_date")), "now")
	}
	if BoolValue(p.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
	if BoolValue(p.TrialEndNow) {
		body.Add(form.FormatKey(append(keyParts, "trial_end")), "now")
	}
}

// If specified, the invoicing for the given billing cycle iterations will be processed now.
type SubscriptionScheduleUpdatePrebillingParams struct {
	// This is used to determine the number of billing cycles to prebill.
	Iterations *int64 `form:"iterations" json:"iterations"`
	// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period. The default value is `reset`.
	UpdateBehavior *string `form:"update_behavior" json:"update_behavior,omitempty"`
}

// Updates an existing subscription schedule.
type SubscriptionScheduleUpdateParams struct {
	Params `form:"*"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time. `prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior *string `form:"billing_behavior" json:"billing_behavior,omitempty"`
	// Object representing the subscription schedule's default settings.
	DefaultSettings *SubscriptionScheduleUpdateDefaultSettingsParams `form:"default_settings" json:"default_settings,omitempty"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior *string `form:"end_behavior" json:"end_behavior,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase. Note that past phases can be omitted.
	Phases []*SubscriptionScheduleUpdatePhaseParams `form:"phases" json:"phases,omitempty"`
	// If specified, the invoicing for the given billing cycle iterations will be processed now.
	Prebilling *SubscriptionScheduleUpdatePrebillingParams `form:"prebilling" json:"prebilling,omitempty"`
	// If the update changes the billing configuration (item price, quantity, etc.) of the current phase, indicates how prorations from this change should be handled. The default value is `create_prorations`.
	ProrationBehavior *string                                      `form:"proration_behavior" json:"proration_behavior,omitempty"`
	UnsetFields       []SubscriptionScheduleUpdateParamsUnsetField `form:"-" json:"-"`
}

// SubscriptionScheduleUpdateParamsUnsetField is the list of fields that can be cleared/unset on SubscriptionScheduleUpdateParams.
type SubscriptionScheduleUpdateParamsUnsetField string

const (
	SubscriptionScheduleUpdateParamsUnsetFieldMetadata SubscriptionScheduleUpdateParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *SubscriptionScheduleUpdateParams) AddUnsetField(field SubscriptionScheduleUpdateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *SubscriptionScheduleUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SubscriptionScheduleUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Configure behavior for flexible billing mode
type SubscriptionScheduleBillingModeFlexible struct {
	// Controls how invoices and invoice items display proration amounts and discount amounts.
	ProrationDiscounts SubscriptionScheduleBillingModeFlexibleProrationDiscounts `json:"proration_discounts,omitempty"`
}

// The billing mode of the subscription.
type SubscriptionScheduleBillingMode struct {
	// Configure behavior for flexible billing mode
	Flexible *SubscriptionScheduleBillingModeFlexible `json:"flexible"`
	// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
	Type SubscriptionScheduleBillingModeType `json:"type"`
	// Details on when the current billing_mode was adopted.
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
type SubscriptionScheduleCurrentPhase struct {
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type SubscriptionScheduleDefaultSettingsBillingThresholds struct {
	// Monetary threshold that triggers the subscription to create an invoice
	AmountGTE int64 `json:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.
	ResetBillingCycleAnchor bool `json:"reset_billing_cycle_anchor"`
}
type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account,omitempty"`
	// Type of the account referenced.
	Type SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType `json:"type"`
}
type SubscriptionScheduleDefaultSettingsInvoiceSettings struct {
	// The account tax IDs associated with the subscription schedule. Will be set on invoices generated by the subscription schedule.
	AccountTaxIDs []*TaxID `json:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64                                                     `json:"days_until_due"`
	Issuer       *SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer `json:"issuer"`
}
type SubscriptionScheduleDefaultSettings struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                   `json:"application_fee_percent"`
	AutomaticTax          *SubscriptionAutomaticTax `json:"automatic_tax,omitempty"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).
	BillingCycleAnchor SubscriptionScheduleDefaultSettingsBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *SubscriptionScheduleDefaultSettingsBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod *SubscriptionCollectionMethod `json:"collection_method"`
	// ID of the default payment method for the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description     string                                              `json:"description"`
	InvoiceSettings *SubscriptionScheduleDefaultSettingsInvoiceSettings `json:"invoice_settings"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *SubscriptionTransferData `json:"transfer_data"`
}

// The involved price pairs in each failed transition.
type SubscriptionScheduleLastPriceMigrationErrorFailedTransition struct {
	// The original price to be migrated.
	SourcePrice string `json:"source_price"`
	// The intended resulting price of the migration.
	TargetPrice string `json:"target_price"`
}

// Details of the most recent price migration that failed for the subscription schedule.
type SubscriptionScheduleLastPriceMigrationError struct {
	// The time at which the price migration encountered an error.
	ErroredAt int64 `json:"errored_at"`
	// The involved price pairs in each failed transition.
	FailedTransitions []*SubscriptionScheduleLastPriceMigrationErrorFailedTransition `json:"failed_transitions"`
	// The type of error encountered by the price migration.
	Type SubscriptionScheduleLastPriceMigrationErrorType `json:"type"`
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
	DiscountEnd *SubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEnd `json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *PromotionCode `json:"promotion_code"`
}
type SubscriptionSchedulePhaseAddInvoiceItemPeriodEnd struct {
	// A precise Unix timestamp for the end of the invoice item period. Must be greater than or equal to `period.start`.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Select how to calculate the end of the invoice item period.
	Type SubscriptionSchedulePhaseAddInvoiceItemPeriodEndType `json:"type"`
}
type SubscriptionSchedulePhaseAddInvoiceItemPeriodStart struct {
	// A precise Unix timestamp for the start of the invoice item period. Must be less than or equal to `period.end`.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Select how to calculate the start of the invoice item period.
	Type SubscriptionSchedulePhaseAddInvoiceItemPeriodStartType `json:"type"`
}
type SubscriptionSchedulePhaseAddInvoiceItemPeriod struct {
	End   *SubscriptionSchedulePhaseAddInvoiceItemPeriodEnd   `json:"end"`
	Start *SubscriptionSchedulePhaseAddInvoiceItemPeriodStart `json:"start"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
type SubscriptionSchedulePhaseAddInvoiceItem struct {
	// The stackable discounts that will be applied to the item.
	Discounts []*SubscriptionSchedulePhaseAddInvoiceItemDiscount `json:"discounts"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string                              `json:"metadata"`
	Period   *SubscriptionSchedulePhaseAddInvoiceItemPeriod `json:"period"`
	// ID of the price used to generate the invoice item.
	Price *Price `json:"price"`
	// The quantity of the invoice item.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*TaxRate `json:"tax_rates,omitempty"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type SubscriptionSchedulePhaseBillingThresholds struct {
	// Monetary threshold that triggers the subscription to create an invoice
	AmountGTE int64 `json:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.
	ResetBillingCycleAnchor bool `json:"reset_billing_cycle_anchor"`
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
	DiscountEnd *SubscriptionSchedulePhaseDiscountDiscountEnd `json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *PromotionCode `json:"promotion_code"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type SubscriptionSchedulePhaseInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account,omitempty"`
	// Type of the account referenced.
	Type SubscriptionSchedulePhaseInvoiceSettingsIssuerType `json:"type"`
}

// The invoice settings applicable during this phase.
type SubscriptionSchedulePhaseInvoiceSettings struct {
	// The account tax IDs associated with this phase of the subscription schedule. Will be set on invoices generated by this phase of the subscription schedule.
	AccountTaxIDs []*TaxID `json:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *SubscriptionSchedulePhaseInvoiceSettingsIssuer `json:"issuer"`
}

// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
type SubscriptionSchedulePhaseItemBillingThresholds struct {
	// Usage threshold that triggers the subscription to create an invoice
	UsageGTE int64 `json:"usage_gte"`
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
	DiscountEnd *SubscriptionSchedulePhaseItemDiscountDiscountEnd `json:"discount_end,omitempty"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *PromotionCode `json:"promotion_code"`
}

// Options that configure the trial on the subscription item.
type SubscriptionSchedulePhaseItemTrial struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []string `json:"converts_to,omitempty"`
	// Determines the type of trial for this item.
	Type SubscriptionSchedulePhaseItemTrialType `json:"type"`
}

// Subscription items to configure the subscription to during this phase of the subscription schedule.
type SubscriptionSchedulePhaseItem struct {
	// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
	BillingThresholds *SubscriptionSchedulePhaseItemBillingThresholds `json:"billing_thresholds"`
	// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*SubscriptionSchedulePhaseItemDiscount `json:"discounts"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an item. Metadata on this item will update the underlying subscription item's `metadata` when the phase is entered.
	Metadata map[string]string `json:"metadata"`
	// ID of the plan to which the customer should be subscribed.
	Plan *Plan `json:"plan"`
	// ID of the price to which the customer should be subscribed.
	Price *Price `json:"price"`
	// Quantity of the plan to which the customer should be subscribed.
	Quantity int64 `json:"quantity,omitempty"`
	// The tax rates which apply to this `phase_item`. When set, the `default_tax_rates` on the phase do not apply to this `phase_item`.
	TaxRates []*TaxRate `json:"tax_rates,omitempty"`
	// Options that configure the trial on the subscription item.
	Trial *SubscriptionSchedulePhaseItemTrial `json:"trial,omitempty"`
	// The ID of the trial offer to apply to the configuration item.
	TrialOffer string `json:"trial_offer,omitempty"`
}

// If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://docs.stripe.com/billing/subscriptions/pause-payment).
type SubscriptionSchedulePhasePauseCollection struct {
	// The payment collection behavior for this subscription while paused.
	Behavior SubscriptionSchedulePhasePauseCollectionBehavior `json:"behavior"`
}

// Defines how the subscription should behave when a trial ends.
type SubscriptionSchedulePhaseTrialSettingsEndBehavior struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront SubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront `json:"prorate_up_front"`
}

// Settings related to any trials on the subscription during this phase.
type SubscriptionSchedulePhaseTrialSettings struct {
	// Defines how the subscription should behave when a trial ends.
	EndBehavior *SubscriptionSchedulePhaseTrialSettingsEndBehavior `json:"end_behavior"`
}

// Configuration for the subscription schedule's phases.
type SubscriptionSchedulePhase struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
	AddInvoiceItems []*SubscriptionSchedulePhaseAddInvoiceItem `json:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                   `json:"application_fee_percent"`
	AutomaticTax          *SubscriptionAutomaticTax `json:"automatic_tax,omitempty"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).
	BillingCycleAnchor SubscriptionSchedulePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *SubscriptionSchedulePhaseBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod *SubscriptionCollectionMethod `json:"collection_method"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// The default tax rates to apply to the subscription during this phase of the subscription schedule.
	DefaultTaxRates []*TaxRate `json:"default_tax_rates,omitempty"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description string `json:"description"`
	// The stackable discounts that will be applied to the subscription on this phase. Subscription item discounts are applied before subscription discounts.
	Discounts []*SubscriptionSchedulePhaseDiscount `json:"discounts"`
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The invoice settings applicable during this phase.
	InvoiceSettings *SubscriptionSchedulePhaseInvoiceSettings `json:"invoice_settings"`
	// Subscription items to configure the subscription to during this phase of the subscription schedule.
	Items []*SubscriptionSchedulePhaseItem `json:"items"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered. Updating the underlying subscription's `metadata` directly will not affect the current phase's `metadata`.
	Metadata map[string]string `json:"metadata"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://docs.stripe.com/billing/subscriptions/pause-payment).
	PauseCollection *SubscriptionSchedulePhasePauseCollection `json:"pause_collection,omitempty"`
	// When transitioning phases, controls how prorations are handled (if any). Possible values are `create_prorations`, `none`, and `always_invoice`.
	ProrationBehavior SubscriptionSchedulePhaseProrationBehavior `json:"proration_behavior"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *SubscriptionTransferData `json:"transfer_data"`
	// Specify behavior of the trial when crossing schedule phase boundaries
	TrialContinuation SubscriptionSchedulePhaseTrialContinuation `json:"trial_continuation,omitempty"`
	// When the trial ends within the phase.
	TrialEnd int64 `json:"trial_end"`
	// Settings related to any trials on the subscription during this phase.
	TrialSettings *SubscriptionSchedulePhaseTrialSettings `json:"trial_settings,omitempty"`
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
	UpdateBehavior SubscriptionSchedulePrebillingUpdateBehavior `json:"update_behavior,omitempty"`
}

// A subscription schedule allows you to create and manage the lifecycle of a subscription by predefining expected changes.
//
// Related guide: [Subscription schedules](https://docs.stripe.com/billing/subscriptions/subscription-schedules)
type SubscriptionSchedule struct {
	APIResource
	// ID of the Connect Application that created the schedule.
	Application *Application `json:"application"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time. `prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior SubscriptionScheduleBillingBehavior `json:"billing_behavior,omitempty"`
	// The billing mode of the subscription.
	BillingMode *SubscriptionScheduleBillingMode `json:"billing_mode"`
	// Time at which the subscription schedule was canceled. Measured in seconds since the Unix epoch.
	CanceledAt int64 `json:"canceled_at"`
	// Time at which the subscription schedule was completed. Measured in seconds since the Unix epoch.
	CompletedAt int64 `json:"completed_at"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
	CurrentPhase *SubscriptionScheduleCurrentPhase `json:"current_phase"`
	// ID of the customer who owns the subscription schedule.
	Customer *Customer `json:"customer"`
	// ID of the account who owns the subscription schedule.
	CustomerAccount string                               `json:"customer_account"`
	DefaultSettings *SubscriptionScheduleDefaultSettings `json:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior SubscriptionScheduleEndBehavior `json:"end_behavior"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Details of the most recent price migration that failed for the subscription schedule.
	LastPriceMigrationError *SubscriptionScheduleLastPriceMigrationError `json:"last_price_migration_error,omitempty"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Configuration for the subscription schedule's phases.
	Phases []*SubscriptionSchedulePhase `json:"phases"`
	// Time period and invoice for a Subscription billed in advance.
	Prebilling *SubscriptionSchedulePrebilling `json:"prebilling,omitempty"`
	// Time at which the subscription schedule was released. Measured in seconds since the Unix epoch.
	ReleasedAt int64 `json:"released_at"`
	// ID of the subscription once managed by the subscription schedule (if it is released).
	ReleasedSubscription *Subscription `json:"released_subscription"`
	// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://docs.stripe.com/billing/subscriptions/subscription-schedules).
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
