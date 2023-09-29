//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Describes whether the quote line is affecting a new schedule or an existing schedule.
type QuotePreviewScheduleAppliesToType string

// List of values that QuotePreviewScheduleAppliesToType can take
const (
	QuotePreviewScheduleAppliesToTypeNewReference         QuotePreviewScheduleAppliesToType = "new_reference"
	QuotePreviewScheduleAppliesToTypeSubscriptionSchedule QuotePreviewScheduleAppliesToType = "subscription_schedule"
)

// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
type QuotePreviewScheduleBillingBehavior string

// List of values that QuotePreviewScheduleBillingBehavior can take
const (
	QuotePreviewScheduleBillingBehaviorProrateOnNextPhase QuotePreviewScheduleBillingBehavior = "prorate_on_next_phase"
	QuotePreviewScheduleBillingBehaviorProrateUpFront     QuotePreviewScheduleBillingBehavior = "prorate_up_front"
)

// Type of the account referenced.
type QuotePreviewScheduleDefaultSettingsAutomaticTaxLiabilityType string

// List of values that QuotePreviewScheduleDefaultSettingsAutomaticTaxLiabilityType can take
const (
	QuotePreviewScheduleDefaultSettingsAutomaticTaxLiabilityTypeAccount QuotePreviewScheduleDefaultSettingsAutomaticTaxLiabilityType = "account"
	QuotePreviewScheduleDefaultSettingsAutomaticTaxLiabilityTypeSelf    QuotePreviewScheduleDefaultSettingsAutomaticTaxLiabilityType = "self"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type QuotePreviewScheduleDefaultSettingsBillingCycleAnchor string

// List of values that QuotePreviewScheduleDefaultSettingsBillingCycleAnchor can take
const (
	QuotePreviewScheduleDefaultSettingsBillingCycleAnchorAutomatic  QuotePreviewScheduleDefaultSettingsBillingCycleAnchor = "automatic"
	QuotePreviewScheduleDefaultSettingsBillingCycleAnchorPhaseStart QuotePreviewScheduleDefaultSettingsBillingCycleAnchor = "phase_start"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
type QuotePreviewScheduleDefaultSettingsCollectionMethod string

// List of values that QuotePreviewScheduleDefaultSettingsCollectionMethod can take
const (
	QuotePreviewScheduleDefaultSettingsCollectionMethodChargeAutomatically QuotePreviewScheduleDefaultSettingsCollectionMethod = "charge_automatically"
	QuotePreviewScheduleDefaultSettingsCollectionMethodSendInvoice         QuotePreviewScheduleDefaultSettingsCollectionMethod = "send_invoice"
)

// Type of the account referenced.
type QuotePreviewScheduleDefaultSettingsInvoiceSettingsIssuerType string

// List of values that QuotePreviewScheduleDefaultSettingsInvoiceSettingsIssuerType can take
const (
	QuotePreviewScheduleDefaultSettingsInvoiceSettingsIssuerTypeAccount QuotePreviewScheduleDefaultSettingsInvoiceSettingsIssuerType = "account"
	QuotePreviewScheduleDefaultSettingsInvoiceSettingsIssuerTypeSelf    QuotePreviewScheduleDefaultSettingsInvoiceSettingsIssuerType = "self"
)

// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running.`cancel` will end the subscription schedule and cancel the underlying subscription.
type QuotePreviewScheduleEndBehavior string

// List of values that QuotePreviewScheduleEndBehavior can take
const (
	QuotePreviewScheduleEndBehaviorCancel  QuotePreviewScheduleEndBehavior = "cancel"
	QuotePreviewScheduleEndBehaviorNone    QuotePreviewScheduleEndBehavior = "none"
	QuotePreviewScheduleEndBehaviorRelease QuotePreviewScheduleEndBehavior = "release"
	QuotePreviewScheduleEndBehaviorRenew   QuotePreviewScheduleEndBehavior = "renew"
)

// The discount end type.
type QuotePreviewSchedulePhaseAddInvoiceItemDiscountDiscountEndType string

// List of values that QuotePreviewSchedulePhaseAddInvoiceItemDiscountDiscountEndType can take
const (
	QuotePreviewSchedulePhaseAddInvoiceItemDiscountDiscountEndTypeTimestamp QuotePreviewSchedulePhaseAddInvoiceItemDiscountDiscountEndType = "timestamp"
)

// Type of the account referenced.
type QuotePreviewSchedulePhaseAutomaticTaxLiabilityType string

// List of values that QuotePreviewSchedulePhaseAutomaticTaxLiabilityType can take
const (
	QuotePreviewSchedulePhaseAutomaticTaxLiabilityTypeAccount QuotePreviewSchedulePhaseAutomaticTaxLiabilityType = "account"
	QuotePreviewSchedulePhaseAutomaticTaxLiabilityTypeSelf    QuotePreviewSchedulePhaseAutomaticTaxLiabilityType = "self"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type QuotePreviewSchedulePhaseBillingCycleAnchor string

// List of values that QuotePreviewSchedulePhaseBillingCycleAnchor can take
const (
	QuotePreviewSchedulePhaseBillingCycleAnchorAutomatic  QuotePreviewSchedulePhaseBillingCycleAnchor = "automatic"
	QuotePreviewSchedulePhaseBillingCycleAnchorPhaseStart QuotePreviewSchedulePhaseBillingCycleAnchor = "phase_start"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
type QuotePreviewSchedulePhaseCollectionMethod string

// List of values that QuotePreviewSchedulePhaseCollectionMethod can take
const (
	QuotePreviewSchedulePhaseCollectionMethodChargeAutomatically QuotePreviewSchedulePhaseCollectionMethod = "charge_automatically"
	QuotePreviewSchedulePhaseCollectionMethodSendInvoice         QuotePreviewSchedulePhaseCollectionMethod = "send_invoice"
)

// The discount end type.
type QuotePreviewSchedulePhaseDiscountDiscountEndType string

// List of values that QuotePreviewSchedulePhaseDiscountDiscountEndType can take
const (
	QuotePreviewSchedulePhaseDiscountDiscountEndTypeTimestamp QuotePreviewSchedulePhaseDiscountDiscountEndType = "timestamp"
)

// Type of the account referenced.
type QuotePreviewSchedulePhaseInvoiceSettingsIssuerType string

// List of values that QuotePreviewSchedulePhaseInvoiceSettingsIssuerType can take
const (
	QuotePreviewSchedulePhaseInvoiceSettingsIssuerTypeAccount QuotePreviewSchedulePhaseInvoiceSettingsIssuerType = "account"
	QuotePreviewSchedulePhaseInvoiceSettingsIssuerTypeSelf    QuotePreviewSchedulePhaseInvoiceSettingsIssuerType = "self"
)

// The discount end type.
type QuotePreviewSchedulePhaseItemDiscountDiscountEndType string

// List of values that QuotePreviewSchedulePhaseItemDiscountDiscountEndType can take
const (
	QuotePreviewSchedulePhaseItemDiscountDiscountEndTypeTimestamp QuotePreviewSchedulePhaseItemDiscountDiscountEndType = "timestamp"
)

type QuotePreviewSchedulePhaseItemTrialType string

// List of values that QuotePreviewSchedulePhaseItemTrialType can take
const (
	QuotePreviewSchedulePhaseItemTrialTypeFree QuotePreviewSchedulePhaseItemTrialType = "free"
	QuotePreviewSchedulePhaseItemTrialTypePaid QuotePreviewSchedulePhaseItemTrialType = "paid"
)

// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
type QuotePreviewSchedulePhasePauseCollectionBehavior string

// List of values that QuotePreviewSchedulePhasePauseCollectionBehavior can take
const (
	QuotePreviewSchedulePhasePauseCollectionBehaviorKeepAsDraft       QuotePreviewSchedulePhasePauseCollectionBehavior = "keep_as_draft"
	QuotePreviewSchedulePhasePauseCollectionBehaviorMarkUncollectible QuotePreviewSchedulePhasePauseCollectionBehavior = "mark_uncollectible"
	QuotePreviewSchedulePhasePauseCollectionBehaviorVoid              QuotePreviewSchedulePhasePauseCollectionBehavior = "void"
)

// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
type QuotePreviewSchedulePhaseProrationBehavior string

// List of values that QuotePreviewSchedulePhaseProrationBehavior can take
const (
	QuotePreviewSchedulePhaseProrationBehaviorAlwaysInvoice    QuotePreviewSchedulePhaseProrationBehavior = "always_invoice"
	QuotePreviewSchedulePhaseProrationBehaviorCreateProrations QuotePreviewSchedulePhaseProrationBehavior = "create_prorations"
	QuotePreviewSchedulePhaseProrationBehaviorNone             QuotePreviewSchedulePhaseProrationBehavior = "none"
)

// Specify behavior of the trial when crossing schedule phase boundaries
type QuotePreviewSchedulePhaseTrialContinuation string

// List of values that QuotePreviewSchedulePhaseTrialContinuation can take
const (
	QuotePreviewSchedulePhaseTrialContinuationContinue QuotePreviewSchedulePhaseTrialContinuation = "continue"
	QuotePreviewSchedulePhaseTrialContinuationNone     QuotePreviewSchedulePhaseTrialContinuation = "none"
)

// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
type QuotePreviewSchedulePhaseTrialSettingsEndBehaviorProrateUpFront string

// List of values that QuotePreviewSchedulePhaseTrialSettingsEndBehaviorProrateUpFront can take
const (
	QuotePreviewSchedulePhaseTrialSettingsEndBehaviorProrateUpFrontDefer   QuotePreviewSchedulePhaseTrialSettingsEndBehaviorProrateUpFront = "defer"
	QuotePreviewSchedulePhaseTrialSettingsEndBehaviorProrateUpFrontInclude QuotePreviewSchedulePhaseTrialSettingsEndBehaviorProrateUpFront = "include"
)

// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period.
type QuotePreviewSchedulePrebillingUpdateBehavior string

// List of values that QuotePreviewSchedulePrebillingUpdateBehavior can take
const (
	QuotePreviewSchedulePrebillingUpdateBehaviorPrebill QuotePreviewSchedulePrebillingUpdateBehavior = "prebill"
	QuotePreviewSchedulePrebillingUpdateBehaviorReset   QuotePreviewSchedulePrebillingUpdateBehavior = "reset"
)

// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
type QuotePreviewScheduleStatus string

// List of values that QuotePreviewScheduleStatus can take
const (
	QuotePreviewScheduleStatusActive     QuotePreviewScheduleStatus = "active"
	QuotePreviewScheduleStatusCanceled   QuotePreviewScheduleStatus = "canceled"
	QuotePreviewScheduleStatusCompleted  QuotePreviewScheduleStatus = "completed"
	QuotePreviewScheduleStatusNotStarted QuotePreviewScheduleStatus = "not_started"
	QuotePreviewScheduleStatusReleased   QuotePreviewScheduleStatus = "released"
)

// Preview the schedules that would be generated by accepting the quote
type QuotePreviewScheduleListParams struct {
	ListParams `form:"*"`
	Quote      *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuotePreviewScheduleListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type QuotePreviewScheduleAppliesTo struct {
	// A custom string that identifies a new subscription schedule being created upon quote acceptance. All quote lines with the same `new_reference` field will be applied to the creation of a new subscription schedule.
	NewReference string `json:"new_reference"`
	// The ID of the schedule the line applies to.
	SubscriptionSchedule string `json:"subscription_schedule"`
	// Describes whether the quote line is affecting a new schedule or an existing schedule.
	Type QuotePreviewScheduleAppliesToType `json:"type"`
}

// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
type QuotePreviewScheduleCurrentPhase struct {
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type QuotePreviewScheduleDefaultSettingsAutomaticTaxLiability struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewScheduleDefaultSettingsAutomaticTaxLiabilityType `json:"type"`
}
type QuotePreviewScheduleDefaultSettingsAutomaticTax struct {
	// Whether Stripe automatically computes tax on invoices created during this phase.
	Enabled bool `json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *QuotePreviewScheduleDefaultSettingsAutomaticTaxLiability `json:"liability"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type QuotePreviewScheduleDefaultSettingsBillingThresholds struct {
	// Monetary threshold that triggers the subscription to create an invoice
	AmountGTE int64 `json:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.
	ResetBillingCycleAnchor bool `json:"reset_billing_cycle_anchor"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type QuotePreviewScheduleDefaultSettingsInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewScheduleDefaultSettingsInvoiceSettingsIssuerType `json:"type"`
}

// The subscription schedule's default invoice settings.
type QuotePreviewScheduleDefaultSettingsInvoiceSettings struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *QuotePreviewScheduleDefaultSettingsInvoiceSettingsIssuer `json:"issuer"`
}

// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
type QuotePreviewScheduleDefaultSettingsTransferData struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount is transferred to the destination.
	AmountPercent float64 `json:"amount_percent"`
	// The account where funds from the payment will be transferred to upon payment success.
	Destination *Account `json:"destination"`
}
type QuotePreviewScheduleDefaultSettings struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                                          `json:"application_fee_percent"`
	AutomaticTax          *QuotePreviewScheduleDefaultSettingsAutomaticTax `json:"automatic_tax"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor QuotePreviewScheduleDefaultSettingsBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *QuotePreviewScheduleDefaultSettingsBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod QuotePreviewScheduleDefaultSettingsCollectionMethod `json:"collection_method"`
	// ID of the default payment method for the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description"`
	// The subscription schedule's default invoice settings.
	InvoiceSettings *QuotePreviewScheduleDefaultSettingsInvoiceSettings `json:"invoice_settings"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *QuotePreviewScheduleDefaultSettingsTransferData `json:"transfer_data"`
}

// Details to determine how long the discount should be applied for.
type QuotePreviewSchedulePhaseAddInvoiceItemDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type QuotePreviewSchedulePhaseAddInvoiceItemDiscountDiscountEndType `json:"type"`
}

// The stackable discounts that will be applied to the item.
type QuotePreviewSchedulePhaseAddInvoiceItemDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuotePreviewSchedulePhaseAddInvoiceItemDiscountDiscountEnd `json:"discount_end"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
type QuotePreviewSchedulePhaseAddInvoiceItem struct {
	// The stackable discounts that will be applied to the item.
	Discounts []*QuotePreviewSchedulePhaseAddInvoiceItemDiscount `json:"discounts"`
	// ID of the price used to generate the invoice item.
	Price *Price `json:"price"`
	// The quantity of the invoice item.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*TaxRate `json:"tax_rates"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type QuotePreviewSchedulePhaseAutomaticTaxLiability struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewSchedulePhaseAutomaticTaxLiabilityType `json:"type"`
}
type QuotePreviewSchedulePhaseAutomaticTax struct {
	// Whether Stripe automatically computes tax on invoices created during this phase.
	Enabled bool `json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *QuotePreviewSchedulePhaseAutomaticTaxLiability `json:"liability"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type QuotePreviewSchedulePhaseBillingThresholds struct {
	// Monetary threshold that triggers the subscription to create an invoice
	AmountGTE int64 `json:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.
	ResetBillingCycleAnchor bool `json:"reset_billing_cycle_anchor"`
}

// Details to determine how long the discount should be applied for.
type QuotePreviewSchedulePhaseDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type QuotePreviewSchedulePhaseDiscountDiscountEndType `json:"type"`
}

// The stackable discounts that will be applied to the subscription on this phase. Subscription item discounts are applied before subscription discounts.
type QuotePreviewSchedulePhaseDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuotePreviewSchedulePhaseDiscountDiscountEnd `json:"discount_end"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type QuotePreviewSchedulePhaseInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewSchedulePhaseInvoiceSettingsIssuerType `json:"type"`
}

// The invoice settings applicable during this phase.
type QuotePreviewSchedulePhaseInvoiceSettings struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *QuotePreviewSchedulePhaseInvoiceSettingsIssuer `json:"issuer"`
}

// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
type QuotePreviewSchedulePhaseItemBillingThresholds struct {
	// Usage threshold that triggers the subscription to create an invoice
	UsageGTE int64 `json:"usage_gte"`
}

// Details to determine how long the discount should be applied for.
type QuotePreviewSchedulePhaseItemDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type QuotePreviewSchedulePhaseItemDiscountDiscountEndType `json:"type"`
}

// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
type QuotePreviewSchedulePhaseItemDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuotePreviewSchedulePhaseItemDiscountDiscountEnd `json:"discount_end"`
}

// Options that configure the trial on the subscription item.
type QuotePreviewSchedulePhaseItemTrial struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []string                               `json:"converts_to"`
	Type       QuotePreviewSchedulePhaseItemTrialType `json:"type"`
}

// Subscription items to configure the subscription to during this phase of the subscription schedule.
type QuotePreviewSchedulePhaseItem struct {
	// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
	BillingThresholds *QuotePreviewSchedulePhaseItemBillingThresholds `json:"billing_thresholds"`
	// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*QuotePreviewSchedulePhaseItemDiscount `json:"discounts"`
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
	Trial *QuotePreviewSchedulePhaseItemTrial `json:"trial"`
}

// If specified, payment collection for this subscription will be paused.
type QuotePreviewSchedulePhasePauseCollection struct {
	// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
	Behavior QuotePreviewSchedulePhasePauseCollectionBehavior `json:"behavior"`
}

// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
type QuotePreviewSchedulePhaseTransferData struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount is transferred to the destination.
	AmountPercent float64 `json:"amount_percent"`
	// The account where funds from the payment will be transferred to upon payment success.
	Destination *Account `json:"destination"`
}

// Defines how the subscription should behaves when a trial ensd.
type QuotePreviewSchedulePhaseTrialSettingsEndBehavior struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront QuotePreviewSchedulePhaseTrialSettingsEndBehaviorProrateUpFront `json:"prorate_up_front"`
}

// Settings related to any trials on the subscription during this phase.
type QuotePreviewSchedulePhaseTrialSettings struct {
	// Defines how the subscription should behaves when a trial ensd.
	EndBehavior *QuotePreviewSchedulePhaseTrialSettingsEndBehavior `json:"end_behavior"`
}

// Configuration for the subscription schedule's phases.
type QuotePreviewSchedulePhase struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
	AddInvoiceItems []*QuotePreviewSchedulePhaseAddInvoiceItem `json:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                                `json:"application_fee_percent"`
	AutomaticTax          *QuotePreviewSchedulePhaseAutomaticTax `json:"automatic_tax"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor QuotePreviewSchedulePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *QuotePreviewSchedulePhaseBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod QuotePreviewSchedulePhaseCollectionMethod `json:"collection_method"`
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
	Discounts []*QuotePreviewSchedulePhaseDiscount `json:"discounts"`
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The invoice settings applicable during this phase.
	InvoiceSettings *QuotePreviewSchedulePhaseInvoiceSettings `json:"invoice_settings"`
	// Subscription items to configure the subscription to during this phase of the subscription schedule.
	Items []*QuotePreviewSchedulePhaseItem `json:"items"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered. Updating the underlying subscription's `metadata` directly will not affect the current phase's `metadata`.
	Metadata map[string]string `json:"metadata"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// If specified, payment collection for this subscription will be paused.
	PauseCollection *QuotePreviewSchedulePhasePauseCollection `json:"pause_collection"`
	// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
	ProrationBehavior QuotePreviewSchedulePhaseProrationBehavior `json:"proration_behavior"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *QuotePreviewSchedulePhaseTransferData `json:"transfer_data"`
	// Specify behavior of the trial when crossing schedule phase boundaries
	TrialContinuation QuotePreviewSchedulePhaseTrialContinuation `json:"trial_continuation"`
	// When the trial ends within the phase.
	TrialEnd int64 `json:"trial_end"`
	// Settings related to any trials on the subscription during this phase.
	TrialSettings *QuotePreviewSchedulePhaseTrialSettings `json:"trial_settings"`
}

// Time period and invoice for a Subscription billed in advance.
type QuotePreviewSchedulePrebilling struct {
	// ID of the prebilling invoice.
	Invoice *Invoice `json:"invoice"`
	// The end of the last period for which the invoice pre-bills.
	PeriodEnd int64 `json:"period_end"`
	// The start of the first period for which the invoice pre-bills.
	PeriodStart int64 `json:"period_start"`
	// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period.
	UpdateBehavior QuotePreviewSchedulePrebillingUpdateBehavior `json:"update_behavior"`
}
type QuotePreviewSchedule struct {
	// ID of the Connect Application that created the schedule.
	Application *Application                   `json:"application"`
	AppliesTo   *QuotePreviewScheduleAppliesTo `json:"applies_to"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior QuotePreviewScheduleBillingBehavior `json:"billing_behavior"`
	// Time at which the subscription schedule was canceled. Measured in seconds since the Unix epoch.
	CanceledAt int64 `json:"canceled_at"`
	// Time at which the subscription schedule was completed. Measured in seconds since the Unix epoch.
	CompletedAt int64 `json:"completed_at"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
	CurrentPhase *QuotePreviewScheduleCurrentPhase `json:"current_phase"`
	// ID of the customer who owns the subscription schedule.
	Customer        *Customer                            `json:"customer"`
	DefaultSettings *QuotePreviewScheduleDefaultSettings `json:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running.`cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior QuotePreviewScheduleEndBehavior `json:"end_behavior"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Configuration for the subscription schedule's phases.
	Phases []*QuotePreviewSchedulePhase `json:"phases"`
	// Time period and invoice for a Subscription billed in advance.
	Prebilling *QuotePreviewSchedulePrebilling `json:"prebilling"`
	// Time at which the subscription schedule was released. Measured in seconds since the Unix epoch.
	ReleasedAt int64 `json:"released_at"`
	// ID of the subscription once managed by the subscription schedule (if it is released).
	ReleasedSubscription string `json:"released_subscription"`
	// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
	Status QuotePreviewScheduleStatus `json:"status"`
	// ID of the subscription managed by the subscription schedule.
	Subscription *Subscription `json:"subscription"`
	// ID of the test clock this subscription schedule belongs to.
	TestClock *TestHelpersTestClock `json:"test_clock"`
}

// QuotePreviewScheduleList is a list of QuotePreviewSchedules as retrieved from a list endpoint.
type QuotePreviewScheduleList struct {
	APIResource
	ListMeta
	Data []*QuotePreviewSchedule `json:"data"`
}
