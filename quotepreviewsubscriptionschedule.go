//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Describes whether the quote line is affecting a new schedule or an existing schedule.
type QuotePreviewSubscriptionScheduleAppliesToType string

// List of values that QuotePreviewSubscriptionScheduleAppliesToType can take
const (
	QuotePreviewSubscriptionScheduleAppliesToTypeNewReference         QuotePreviewSubscriptionScheduleAppliesToType = "new_reference"
	QuotePreviewSubscriptionScheduleAppliesToTypeSubscriptionSchedule QuotePreviewSubscriptionScheduleAppliesToType = "subscription_schedule"
)

// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
type QuotePreviewSubscriptionScheduleBillingBehavior string

// List of values that QuotePreviewSubscriptionScheduleBillingBehavior can take
const (
	QuotePreviewSubscriptionScheduleBillingBehaviorProrateOnNextPhase QuotePreviewSubscriptionScheduleBillingBehavior = "prorate_on_next_phase"
	QuotePreviewSubscriptionScheduleBillingBehaviorProrateUpFront     QuotePreviewSubscriptionScheduleBillingBehavior = "prorate_up_front"
)

// Type of the account referenced.
type QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiabilityType string

// List of values that QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiabilityType can take
const (
	QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiabilityTypeAccount QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiabilityType = "account"
	QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiabilityTypeSelf    QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiabilityType = "self"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type QuotePreviewSubscriptionScheduleDefaultSettingsBillingCycleAnchor string

// List of values that QuotePreviewSubscriptionScheduleDefaultSettingsBillingCycleAnchor can take
const (
	QuotePreviewSubscriptionScheduleDefaultSettingsBillingCycleAnchorAutomatic  QuotePreviewSubscriptionScheduleDefaultSettingsBillingCycleAnchor = "automatic"
	QuotePreviewSubscriptionScheduleDefaultSettingsBillingCycleAnchorPhaseStart QuotePreviewSubscriptionScheduleDefaultSettingsBillingCycleAnchor = "phase_start"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
type QuotePreviewSubscriptionScheduleDefaultSettingsCollectionMethod string

// List of values that QuotePreviewSubscriptionScheduleDefaultSettingsCollectionMethod can take
const (
	QuotePreviewSubscriptionScheduleDefaultSettingsCollectionMethodChargeAutomatically QuotePreviewSubscriptionScheduleDefaultSettingsCollectionMethod = "charge_automatically"
	QuotePreviewSubscriptionScheduleDefaultSettingsCollectionMethodSendInvoice         QuotePreviewSubscriptionScheduleDefaultSettingsCollectionMethod = "send_invoice"
)

// Type of the account referenced.
type QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType string

// List of values that QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType can take
const (
	QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerTypeAccount QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType = "account"
	QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerTypeSelf    QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType = "self"
)

// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running.`cancel` will end the subscription schedule and cancel the underlying subscription.
type QuotePreviewSubscriptionScheduleEndBehavior string

// List of values that QuotePreviewSubscriptionScheduleEndBehavior can take
const (
	QuotePreviewSubscriptionScheduleEndBehaviorCancel  QuotePreviewSubscriptionScheduleEndBehavior = "cancel"
	QuotePreviewSubscriptionScheduleEndBehaviorNone    QuotePreviewSubscriptionScheduleEndBehavior = "none"
	QuotePreviewSubscriptionScheduleEndBehaviorRelease QuotePreviewSubscriptionScheduleEndBehavior = "release"
	QuotePreviewSubscriptionScheduleEndBehaviorRenew   QuotePreviewSubscriptionScheduleEndBehavior = "renew"
)

// The discount end type.
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType string

// List of values that QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType can take
const (
	QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndTypeTimestamp QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType = "timestamp"
)

// Type of the account referenced.
type QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiabilityType string

// List of values that QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiabilityType can take
const (
	QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiabilityTypeAccount QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiabilityType = "account"
	QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiabilityTypeSelf    QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiabilityType = "self"
)

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type QuotePreviewSubscriptionSchedulePhaseBillingCycleAnchor string

// List of values that QuotePreviewSubscriptionSchedulePhaseBillingCycleAnchor can take
const (
	QuotePreviewSubscriptionSchedulePhaseBillingCycleAnchorAutomatic  QuotePreviewSubscriptionSchedulePhaseBillingCycleAnchor = "automatic"
	QuotePreviewSubscriptionSchedulePhaseBillingCycleAnchorPhaseStart QuotePreviewSubscriptionSchedulePhaseBillingCycleAnchor = "phase_start"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
type QuotePreviewSubscriptionSchedulePhaseCollectionMethod string

// List of values that QuotePreviewSubscriptionSchedulePhaseCollectionMethod can take
const (
	QuotePreviewSubscriptionSchedulePhaseCollectionMethodChargeAutomatically QuotePreviewSubscriptionSchedulePhaseCollectionMethod = "charge_automatically"
	QuotePreviewSubscriptionSchedulePhaseCollectionMethodSendInvoice         QuotePreviewSubscriptionSchedulePhaseCollectionMethod = "send_invoice"
)

// The discount end type.
type QuotePreviewSubscriptionSchedulePhaseDiscountDiscountEndType string

// List of values that QuotePreviewSubscriptionSchedulePhaseDiscountDiscountEndType can take
const (
	QuotePreviewSubscriptionSchedulePhaseDiscountDiscountEndTypeTimestamp QuotePreviewSubscriptionSchedulePhaseDiscountDiscountEndType = "timestamp"
)

// Type of the account referenced.
type QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuerType string

// List of values that QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuerType can take
const (
	QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuerTypeAccount QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuerType = "account"
	QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuerTypeSelf    QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuerType = "self"
)

// The discount end type.
type QuotePreviewSubscriptionSchedulePhaseItemDiscountDiscountEndType string

// List of values that QuotePreviewSubscriptionSchedulePhaseItemDiscountDiscountEndType can take
const (
	QuotePreviewSubscriptionSchedulePhaseItemDiscountDiscountEndTypeTimestamp QuotePreviewSubscriptionSchedulePhaseItemDiscountDiscountEndType = "timestamp"
)

type QuotePreviewSubscriptionSchedulePhaseItemTrialType string

// List of values that QuotePreviewSubscriptionSchedulePhaseItemTrialType can take
const (
	QuotePreviewSubscriptionSchedulePhaseItemTrialTypeFree QuotePreviewSubscriptionSchedulePhaseItemTrialType = "free"
	QuotePreviewSubscriptionSchedulePhaseItemTrialTypePaid QuotePreviewSubscriptionSchedulePhaseItemTrialType = "paid"
)

// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
type QuotePreviewSubscriptionSchedulePhasePauseCollectionBehavior string

// List of values that QuotePreviewSubscriptionSchedulePhasePauseCollectionBehavior can take
const (
	QuotePreviewSubscriptionSchedulePhasePauseCollectionBehaviorKeepAsDraft       QuotePreviewSubscriptionSchedulePhasePauseCollectionBehavior = "keep_as_draft"
	QuotePreviewSubscriptionSchedulePhasePauseCollectionBehaviorMarkUncollectible QuotePreviewSubscriptionSchedulePhasePauseCollectionBehavior = "mark_uncollectible"
	QuotePreviewSubscriptionSchedulePhasePauseCollectionBehaviorVoid              QuotePreviewSubscriptionSchedulePhasePauseCollectionBehavior = "void"
)

// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
type QuotePreviewSubscriptionSchedulePhaseProrationBehavior string

// List of values that QuotePreviewSubscriptionSchedulePhaseProrationBehavior can take
const (
	QuotePreviewSubscriptionSchedulePhaseProrationBehaviorAlwaysInvoice    QuotePreviewSubscriptionSchedulePhaseProrationBehavior = "always_invoice"
	QuotePreviewSubscriptionSchedulePhaseProrationBehaviorCreateProrations QuotePreviewSubscriptionSchedulePhaseProrationBehavior = "create_prorations"
	QuotePreviewSubscriptionSchedulePhaseProrationBehaviorNone             QuotePreviewSubscriptionSchedulePhaseProrationBehavior = "none"
)

// Specify behavior of the trial when crossing schedule phase boundaries
type QuotePreviewSubscriptionSchedulePhaseTrialContinuation string

// List of values that QuotePreviewSubscriptionSchedulePhaseTrialContinuation can take
const (
	QuotePreviewSubscriptionSchedulePhaseTrialContinuationContinue QuotePreviewSubscriptionSchedulePhaseTrialContinuation = "continue"
	QuotePreviewSubscriptionSchedulePhaseTrialContinuationNone     QuotePreviewSubscriptionSchedulePhaseTrialContinuation = "none"
)

// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
type QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront string

// List of values that QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront can take
const (
	QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFrontDefer   QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront = "defer"
	QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFrontInclude QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront = "include"
)

// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period.
type QuotePreviewSubscriptionSchedulePrebillingUpdateBehavior string

// List of values that QuotePreviewSubscriptionSchedulePrebillingUpdateBehavior can take
const (
	QuotePreviewSubscriptionSchedulePrebillingUpdateBehaviorPrebill QuotePreviewSubscriptionSchedulePrebillingUpdateBehavior = "prebill"
	QuotePreviewSubscriptionSchedulePrebillingUpdateBehaviorReset   QuotePreviewSubscriptionSchedulePrebillingUpdateBehavior = "reset"
)

// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
type QuotePreviewSubscriptionScheduleStatus string

// List of values that QuotePreviewSubscriptionScheduleStatus can take
const (
	QuotePreviewSubscriptionScheduleStatusActive     QuotePreviewSubscriptionScheduleStatus = "active"
	QuotePreviewSubscriptionScheduleStatusCanceled   QuotePreviewSubscriptionScheduleStatus = "canceled"
	QuotePreviewSubscriptionScheduleStatusCompleted  QuotePreviewSubscriptionScheduleStatus = "completed"
	QuotePreviewSubscriptionScheduleStatusNotStarted QuotePreviewSubscriptionScheduleStatus = "not_started"
	QuotePreviewSubscriptionScheduleStatusReleased   QuotePreviewSubscriptionScheduleStatus = "released"
)

// Preview the schedules that would be generated by accepting the quote
type QuotePreviewSubscriptionScheduleListParams struct {
	ListParams `form:"*"`
	Quote      *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuotePreviewSubscriptionScheduleListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type QuotePreviewSubscriptionScheduleAppliesTo struct {
	// A custom string that identifies a new subscription schedule being created upon quote acceptance. All quote lines with the same `new_reference` field will be applied to the creation of a new subscription schedule.
	NewReference string `json:"new_reference"`
	// The ID of the schedule the line applies to.
	SubscriptionSchedule string `json:"subscription_schedule"`
	// Describes whether the quote line is affecting a new schedule or an existing schedule.
	Type QuotePreviewSubscriptionScheduleAppliesToType `json:"type"`
}

// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
type QuotePreviewSubscriptionScheduleCurrentPhase struct {
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiability struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiabilityType `json:"type"`
}
type QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTax struct {
	// Whether Stripe automatically computes tax on invoices created during this phase.
	Enabled bool `json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiability `json:"liability"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type QuotePreviewSubscriptionScheduleDefaultSettingsBillingThresholds struct {
	// Monetary threshold that triggers the subscription to create an invoice
	AmountGTE int64 `json:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.
	ResetBillingCycleAnchor bool `json:"reset_billing_cycle_anchor"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType `json:"type"`
}

// The subscription schedule's default invoice settings.
type QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettings struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer `json:"issuer"`
}

// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
type QuotePreviewSubscriptionScheduleDefaultSettingsTransferData struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount is transferred to the destination.
	AmountPercent float64 `json:"amount_percent"`
	// The account where funds from the payment will be transferred to upon payment success.
	Destination *Account `json:"destination"`
}
type QuotePreviewSubscriptionScheduleDefaultSettings struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                                                      `json:"application_fee_percent"`
	AutomaticTax          *QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTax `json:"automatic_tax"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor QuotePreviewSubscriptionScheduleDefaultSettingsBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *QuotePreviewSubscriptionScheduleDefaultSettingsBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod QuotePreviewSubscriptionScheduleDefaultSettingsCollectionMethod `json:"collection_method"`
	// ID of the default payment method for the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description"`
	// The subscription schedule's default invoice settings.
	InvoiceSettings *QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettings `json:"invoice_settings"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *QuotePreviewSubscriptionScheduleDefaultSettingsTransferData `json:"transfer_data"`
}

// Details to determine how long the discount should be applied for.
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType `json:"type"`
}

// The stackable discounts that will be applied to the item.
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEnd `json:"discount_end"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItem struct {
	// The stackable discounts that will be applied to the item.
	Discounts []*QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscount `json:"discounts"`
	// ID of the price used to generate the invoice item.
	Price *Price `json:"price"`
	// The quantity of the invoice item.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates []*TaxRate `json:"tax_rates"`
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiability struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiabilityType `json:"type"`
}
type QuotePreviewSubscriptionSchedulePhaseAutomaticTax struct {
	// Whether Stripe automatically computes tax on invoices created during this phase.
	Enabled bool `json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiability `json:"liability"`
}

// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
type QuotePreviewSubscriptionSchedulePhaseBillingThresholds struct {
	// Monetary threshold that triggers the subscription to create an invoice
	AmountGTE int64 `json:"amount_gte"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.
	ResetBillingCycleAnchor bool `json:"reset_billing_cycle_anchor"`
}

// Details to determine how long the discount should be applied for.
type QuotePreviewSubscriptionSchedulePhaseDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type QuotePreviewSubscriptionSchedulePhaseDiscountDiscountEndType `json:"type"`
}

// The stackable discounts that will be applied to the subscription on this phase. Subscription item discounts are applied before subscription discounts.
type QuotePreviewSubscriptionSchedulePhaseDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuotePreviewSubscriptionSchedulePhaseDiscountDiscountEnd `json:"discount_end"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuerType `json:"type"`
}

// The invoice settings applicable during this phase.
type QuotePreviewSubscriptionSchedulePhaseInvoiceSettings struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuer `json:"issuer"`
}

// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
type QuotePreviewSubscriptionSchedulePhaseItemBillingThresholds struct {
	// Usage threshold that triggers the subscription to create an invoice
	UsageGTE int64 `json:"usage_gte"`
}

// Details to determine how long the discount should be applied for.
type QuotePreviewSubscriptionSchedulePhaseItemDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type QuotePreviewSubscriptionSchedulePhaseItemDiscountDiscountEndType `json:"type"`
}

// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
type QuotePreviewSubscriptionSchedulePhaseItemDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuotePreviewSubscriptionSchedulePhaseItemDiscountDiscountEnd `json:"discount_end"`
}

// Options that configure the trial on the subscription item.
type QuotePreviewSubscriptionSchedulePhaseItemTrial struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []string                                           `json:"converts_to"`
	Type       QuotePreviewSubscriptionSchedulePhaseItemTrialType `json:"type"`
}

// Subscription items to configure the subscription to during this phase of the subscription schedule.
type QuotePreviewSubscriptionSchedulePhaseItem struct {
	// Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period
	BillingThresholds *QuotePreviewSubscriptionSchedulePhaseItemBillingThresholds `json:"billing_thresholds"`
	// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*QuotePreviewSubscriptionSchedulePhaseItemDiscount `json:"discounts"`
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
	Trial *QuotePreviewSubscriptionSchedulePhaseItemTrial `json:"trial"`
}

// If specified, payment collection for this subscription will be paused.
type QuotePreviewSubscriptionSchedulePhasePauseCollection struct {
	// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
	Behavior QuotePreviewSubscriptionSchedulePhasePauseCollectionBehavior `json:"behavior"`
}

// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
type QuotePreviewSubscriptionSchedulePhaseTransferData struct {
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount is transferred to the destination.
	AmountPercent float64 `json:"amount_percent"`
	// The account where funds from the payment will be transferred to upon payment success.
	Destination *Account `json:"destination"`
}

// Defines how the subscription should behaves when a trial ensd.
type QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehavior struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront `json:"prorate_up_front"`
}

// Settings related to any trials on the subscription during this phase.
type QuotePreviewSubscriptionSchedulePhaseTrialSettings struct {
	// Defines how the subscription should behaves when a trial ensd.
	EndBehavior *QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehavior `json:"end_behavior"`
}

// Configuration for the subscription schedule's phases.
type QuotePreviewSubscriptionSchedulePhase struct {
	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
	AddInvoiceItems []*QuotePreviewSubscriptionSchedulePhaseAddInvoiceItem `json:"add_invoice_items"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.
	ApplicationFeePercent float64                                            `json:"application_fee_percent"`
	AutomaticTax          *QuotePreviewSubscriptionSchedulePhaseAutomaticTax `json:"automatic_tax"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor QuotePreviewSubscriptionSchedulePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period
	BillingThresholds *QuotePreviewSubscriptionSchedulePhaseBillingThresholds `json:"billing_thresholds"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod QuotePreviewSubscriptionSchedulePhaseCollectionMethod `json:"collection_method"`
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
	Discounts []*QuotePreviewSubscriptionSchedulePhaseDiscount `json:"discounts"`
	// The end of this phase of the subscription schedule.
	EndDate int64 `json:"end_date"`
	// The invoice settings applicable during this phase.
	InvoiceSettings *QuotePreviewSubscriptionSchedulePhaseInvoiceSettings `json:"invoice_settings"`
	// Subscription items to configure the subscription to during this phase of the subscription schedule.
	Items []*QuotePreviewSubscriptionSchedulePhaseItem `json:"items"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered. Updating the underlying subscription's `metadata` directly will not affect the current phase's `metadata`.
	Metadata map[string]string `json:"metadata"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// If specified, payment collection for this subscription will be paused.
	PauseCollection *QuotePreviewSubscriptionSchedulePhasePauseCollection `json:"pause_collection"`
	// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
	ProrationBehavior QuotePreviewSubscriptionSchedulePhaseProrationBehavior `json:"proration_behavior"`
	// The start of this phase of the subscription schedule.
	StartDate int64 `json:"start_date"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *QuotePreviewSubscriptionSchedulePhaseTransferData `json:"transfer_data"`
	// Specify behavior of the trial when crossing schedule phase boundaries
	TrialContinuation QuotePreviewSubscriptionSchedulePhaseTrialContinuation `json:"trial_continuation"`
	// When the trial ends within the phase.
	TrialEnd int64 `json:"trial_end"`
	// Settings related to any trials on the subscription during this phase.
	TrialSettings *QuotePreviewSubscriptionSchedulePhaseTrialSettings `json:"trial_settings"`
}

// Time period and invoice for a Subscription billed in advance.
type QuotePreviewSubscriptionSchedulePrebilling struct {
	// ID of the prebilling invoice.
	Invoice *Invoice `json:"invoice"`
	// The end of the last period for which the invoice pre-bills.
	PeriodEnd int64 `json:"period_end"`
	// The start of the first period for which the invoice pre-bills.
	PeriodStart int64 `json:"period_start"`
	// Whether to cancel or preserve `prebilling` if the subscription is updated during the prebilled period.
	UpdateBehavior QuotePreviewSubscriptionSchedulePrebillingUpdateBehavior `json:"update_behavior"`
}
type QuotePreviewSubscriptionSchedule struct {
	// ID of the Connect Application that created the schedule.
	Application *Application                               `json:"application"`
	AppliesTo   *QuotePreviewSubscriptionScheduleAppliesTo `json:"applies_to"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior QuotePreviewSubscriptionScheduleBillingBehavior `json:"billing_behavior"`
	// Time at which the subscription schedule was canceled. Measured in seconds since the Unix epoch.
	CanceledAt int64 `json:"canceled_at"`
	// Time at which the subscription schedule was completed. Measured in seconds since the Unix epoch.
	CompletedAt int64 `json:"completed_at"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
	CurrentPhase *QuotePreviewSubscriptionScheduleCurrentPhase `json:"current_phase"`
	// ID of the customer who owns the subscription schedule.
	Customer        *Customer                                        `json:"customer"`
	DefaultSettings *QuotePreviewSubscriptionScheduleDefaultSettings `json:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running.`cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior QuotePreviewSubscriptionScheduleEndBehavior `json:"end_behavior"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Configuration for the subscription schedule's phases.
	Phases []*QuotePreviewSubscriptionSchedulePhase `json:"phases"`
	// Time period and invoice for a Subscription billed in advance.
	Prebilling *QuotePreviewSubscriptionSchedulePrebilling `json:"prebilling"`
	// Time at which the subscription schedule was released. Measured in seconds since the Unix epoch.
	ReleasedAt int64 `json:"released_at"`
	// ID of the subscription once managed by the subscription schedule (if it is released).
	ReleasedSubscription string `json:"released_subscription"`
	// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
	Status QuotePreviewSubscriptionScheduleStatus `json:"status"`
	// ID of the subscription managed by the subscription schedule.
	Subscription *Subscription `json:"subscription"`
	// ID of the test clock this subscription schedule belongs to.
	TestClock *TestHelpersTestClock `json:"test_clock"`
}

// QuotePreviewSubscriptionScheduleList is a list of QuotePreviewSubscriptionSchedules as retrieved from a list endpoint.
type QuotePreviewSubscriptionScheduleList struct {
	APIResource
	ListMeta
	Data []*QuotePreviewSubscriptionSchedule `json:"data"`
}
