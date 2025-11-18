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

// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time. `prorate_up_front` will bill for all phases within the current billing cycle up front.
type QuotePreviewSubscriptionScheduleBillingBehavior string

// List of values that QuotePreviewSubscriptionScheduleBillingBehavior can take
const (
	QuotePreviewSubscriptionScheduleBillingBehaviorProrateOnNextPhase QuotePreviewSubscriptionScheduleBillingBehavior = "prorate_on_next_phase"
	QuotePreviewSubscriptionScheduleBillingBehaviorProrateUpFront     QuotePreviewSubscriptionScheduleBillingBehavior = "prorate_up_front"
)

// Controls how invoices and invoice items display proration amounts and discount amounts.
type QuotePreviewSubscriptionScheduleBillingModeFlexibleProrationDiscounts string

// List of values that QuotePreviewSubscriptionScheduleBillingModeFlexibleProrationDiscounts can take
const (
	QuotePreviewSubscriptionScheduleBillingModeFlexibleProrationDiscountsIncluded QuotePreviewSubscriptionScheduleBillingModeFlexibleProrationDiscounts = "included"
	QuotePreviewSubscriptionScheduleBillingModeFlexibleProrationDiscountsItemized QuotePreviewSubscriptionScheduleBillingModeFlexibleProrationDiscounts = "itemized"
)

// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
type QuotePreviewSubscriptionScheduleBillingModeType string

// List of values that QuotePreviewSubscriptionScheduleBillingModeType can take
const (
	QuotePreviewSubscriptionScheduleBillingModeTypeClassic  QuotePreviewSubscriptionScheduleBillingModeType = "classic"
	QuotePreviewSubscriptionScheduleBillingModeTypeFlexible QuotePreviewSubscriptionScheduleBillingModeType = "flexible"
)

// If Stripe disabled automatic tax, this enum describes why.
type QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxDisabledReason string

// List of values that QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxDisabledReason can take
const (
	QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxDisabledReasonRequiresLocationInputs QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxDisabledReason = "requires_location_inputs"
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

// Configures how the subscription schedule handles billing for phase transitions. Possible values are `phase_start` (default) or `billing_period_start`. `phase_start` bills based on the current state of the subscription, ignoring changes scheduled in future phases. `billing_period_start` bills predictively for upcoming phase transitions within the current billing cycle, including pricing changes and service period adjustments that will occur before the next invoice.
type QuotePreviewSubscriptionScheduleDefaultSettingsPhaseEffectiveAt string

// List of values that QuotePreviewSubscriptionScheduleDefaultSettingsPhaseEffectiveAt can take
const (
	QuotePreviewSubscriptionScheduleDefaultSettingsPhaseEffectiveAtBillingPeriodStart QuotePreviewSubscriptionScheduleDefaultSettingsPhaseEffectiveAt = "billing_period_start"
	QuotePreviewSubscriptionScheduleDefaultSettingsPhaseEffectiveAtPhaseStart         QuotePreviewSubscriptionScheduleDefaultSettingsPhaseEffectiveAt = "phase_start"
)

// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
type QuotePreviewSubscriptionScheduleEndBehavior string

// List of values that QuotePreviewSubscriptionScheduleEndBehavior can take
const (
	QuotePreviewSubscriptionScheduleEndBehaviorCancel  QuotePreviewSubscriptionScheduleEndBehavior = "cancel"
	QuotePreviewSubscriptionScheduleEndBehaviorNone    QuotePreviewSubscriptionScheduleEndBehavior = "none"
	QuotePreviewSubscriptionScheduleEndBehaviorRelease QuotePreviewSubscriptionScheduleEndBehavior = "release"
	QuotePreviewSubscriptionScheduleEndBehaviorRenew   QuotePreviewSubscriptionScheduleEndBehavior = "renew"
)

// The type of error encountered by the price migration.
type QuotePreviewSubscriptionScheduleLastPriceMigrationErrorType string

// List of values that QuotePreviewSubscriptionScheduleLastPriceMigrationErrorType can take
const (
	QuotePreviewSubscriptionScheduleLastPriceMigrationErrorTypePriceUniquenessViolation QuotePreviewSubscriptionScheduleLastPriceMigrationErrorType = "price_uniqueness_violation"
)

// The discount end type.
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType string

// List of values that QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType can take
const (
	QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndTypeTimestamp QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscountDiscountEndType = "timestamp"
)

// Select how to calculate the end of the invoice item period.
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEndType string

// List of values that QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEndType can take
const (
	QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEndTypeMinItemPeriodEnd QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEndType = "min_item_period_end"
	QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEndTypePhaseEnd         QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEndType = "phase_end"
	QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEndTypeTimestamp        QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEndType = "timestamp"
)

// Select how to calculate the start of the invoice item period.
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStartType string

// List of values that QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStartType can take
const (
	QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStartTypeMaxItemPeriodStart QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStartType = "max_item_period_start"
	QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStartTypePhaseStart         QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStartType = "phase_start"
	QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStartTypeTimestamp          QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStartType = "timestamp"
)

// If Stripe disabled automatic tax, this enum describes why.
type QuotePreviewSubscriptionSchedulePhaseAutomaticTaxDisabledReason string

// List of values that QuotePreviewSubscriptionSchedulePhaseAutomaticTaxDisabledReason can take
const (
	QuotePreviewSubscriptionSchedulePhaseAutomaticTaxDisabledReasonRequiresLocationInputs QuotePreviewSubscriptionSchedulePhaseAutomaticTaxDisabledReason = "requires_location_inputs"
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

// Determines the type of trial for this item.
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

// When transitioning phases, controls how prorations are handled (if any). Possible values are `create_prorations`, `none`, and `always_invoice`.
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

// Configures how the subscription schedule handles billing for phase transitions. Possible values are `phase_start` (default) or `billing_period_start`. `phase_start` bills based on the current state of the subscription, ignoring changes scheduled in future phases. `billing_period_start` bills predictively for upcoming phase transitions within the current billing cycle, including pricing changes and service period adjustments that will occur before the next invoice.
type QuotePreviewSubscriptionSchedulePhaseEffectiveAt string

// List of values that QuotePreviewSubscriptionSchedulePhaseEffectiveAt can take
const (
	QuotePreviewSubscriptionSchedulePhaseEffectiveAtBillingPeriodStart QuotePreviewSubscriptionSchedulePhaseEffectiveAt = "billing_period_start"
	QuotePreviewSubscriptionSchedulePhaseEffectiveAtPhaseStart         QuotePreviewSubscriptionSchedulePhaseEffectiveAt = "phase_start"
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

// Controls which subscription items the billing schedule applies to.
type QuotePreviewSubscriptionScheduleBillingScheduleAppliesToType string

// List of values that QuotePreviewSubscriptionScheduleBillingScheduleAppliesToType can take
const (
	QuotePreviewSubscriptionScheduleBillingScheduleAppliesToTypePrice QuotePreviewSubscriptionScheduleBillingScheduleAppliesToType = "price"
)

// Specifies billing duration. Possible values are `day`, `week`, `month`, or `year`.
type QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeInterval string

// List of values that QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeInterval can take
const (
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeIntervalDay   QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeInterval = "day"
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeIntervalMonth QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeInterval = "month"
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeIntervalWeek  QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeInterval = "week"
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeIntervalYear  QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeInterval = "year"
)

// Describes how the billing schedule determines the start date. Possible values are `timestamp`, `relative`, `amendment_start`, `now`, `quote_acceptance_date`, `line_starts_at`, or `pause_collection_start`.
type QuotePreviewSubscriptionScheduleBillingScheduleBillFromType string

// List of values that QuotePreviewSubscriptionScheduleBillingScheduleBillFromType can take
const (
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromTypeAmendmentStart       QuotePreviewSubscriptionScheduleBillingScheduleBillFromType = "amendment_start"
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromTypeLineStartsAt         QuotePreviewSubscriptionScheduleBillingScheduleBillFromType = "line_starts_at"
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromTypeNow                  QuotePreviewSubscriptionScheduleBillingScheduleBillFromType = "now"
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromTypePauseCollectionStart QuotePreviewSubscriptionScheduleBillingScheduleBillFromType = "pause_collection_start"
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromTypeQuoteAcceptanceDate  QuotePreviewSubscriptionScheduleBillingScheduleBillFromType = "quote_acceptance_date"
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromTypeRelative             QuotePreviewSubscriptionScheduleBillingScheduleBillFromType = "relative"
	QuotePreviewSubscriptionScheduleBillingScheduleBillFromTypeTimestamp            QuotePreviewSubscriptionScheduleBillingScheduleBillFromType = "timestamp"
)

// Specifies billing duration. Either `day`, `week`, `month` or `year`.
type QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationInterval string

// List of values that QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationInterval can take
const (
	QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationIntervalDay   QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationInterval = "day"
	QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationIntervalMonth QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationInterval = "month"
	QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationIntervalWeek  QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationInterval = "week"
	QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationIntervalYear  QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationInterval = "year"
)

// Describes how the billing schedule will determine the end date. Either `duration` or `timestamp`.
type QuotePreviewSubscriptionScheduleBillingScheduleBillUntilType string

// List of values that QuotePreviewSubscriptionScheduleBillingScheduleBillUntilType can take
const (
	QuotePreviewSubscriptionScheduleBillingScheduleBillUntilTypeAmendmentEnd    QuotePreviewSubscriptionScheduleBillingScheduleBillUntilType = "amendment_end"
	QuotePreviewSubscriptionScheduleBillingScheduleBillUntilTypeDuration        QuotePreviewSubscriptionScheduleBillingScheduleBillUntilType = "duration"
	QuotePreviewSubscriptionScheduleBillingScheduleBillUntilTypeLineEndsAt      QuotePreviewSubscriptionScheduleBillingScheduleBillUntilType = "line_ends_at"
	QuotePreviewSubscriptionScheduleBillingScheduleBillUntilTypeScheduleEnd     QuotePreviewSubscriptionScheduleBillingScheduleBillUntilType = "schedule_end"
	QuotePreviewSubscriptionScheduleBillingScheduleBillUntilTypeTimestamp       QuotePreviewSubscriptionScheduleBillingScheduleBillUntilType = "timestamp"
	QuotePreviewSubscriptionScheduleBillingScheduleBillUntilTypeUpcomingInvoice QuotePreviewSubscriptionScheduleBillingScheduleBillUntilType = "upcoming_invoice"
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

// Configure behavior for flexible billing mode
type QuotePreviewSubscriptionScheduleBillingModeFlexible struct {
	// Controls how invoices and invoice items display proration amounts and discount amounts.
	ProrationDiscounts QuotePreviewSubscriptionScheduleBillingModeFlexibleProrationDiscounts `json:"proration_discounts"`
}

// The billing mode of the subscription.
type QuotePreviewSubscriptionScheduleBillingMode struct {
	// Configure behavior for flexible billing mode
	Flexible *QuotePreviewSubscriptionScheduleBillingModeFlexible `json:"flexible"`
	// Controls how prorations and invoices for subscriptions are calculated and orchestrated.
	Type QuotePreviewSubscriptionScheduleBillingModeType `json:"type"`
	// Details on when the current billing_mode was adopted.
	UpdatedAt int64 `json:"updated_at"`
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
	// If Stripe disabled automatic tax, this enum describes why.
	DisabledReason QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxDisabledReason `json:"disabled_reason"`
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
type QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerType `json:"type"`
}
type QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettings struct {
	// The account tax IDs associated with the subscription schedule. Will be set on invoices generated by the subscription schedule.
	AccountTaxIDs []*TaxID `json:"account_tax_ids"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64                                                                 `json:"days_until_due"`
	Issuer       *QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer `json:"issuer"`
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
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description     string                                                          `json:"description"`
	InvoiceSettings *QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettings `json:"invoice_settings"`
	// The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// Configures how the subscription schedule handles billing for phase transitions. Possible values are `phase_start` (default) or `billing_period_start`. `phase_start` bills based on the current state of the subscription, ignoring changes scheduled in future phases. `billing_period_start` bills predictively for upcoming phase transitions within the current billing cycle, including pricing changes and service period adjustments that will occur before the next invoice.
	PhaseEffectiveAt QuotePreviewSubscriptionScheduleDefaultSettingsPhaseEffectiveAt `json:"phase_effective_at"`
	// The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.
	TransferData *QuotePreviewSubscriptionScheduleDefaultSettingsTransferData `json:"transfer_data"`
}

// The involved price pairs in each failed transition.
type QuotePreviewSubscriptionScheduleLastPriceMigrationErrorFailedTransition struct {
	// The original price to be migrated.
	SourcePrice string `json:"source_price"`
	// The intended resulting price of the migration.
	TargetPrice string `json:"target_price"`
}

// Details of the most recent price migration that failed for the subscription schedule.
type QuotePreviewSubscriptionScheduleLastPriceMigrationError struct {
	// The time at which the price migration encountered an error.
	ErroredAt int64 `json:"errored_at"`
	// The involved price pairs in each failed transition.
	FailedTransitions []*QuotePreviewSubscriptionScheduleLastPriceMigrationErrorFailedTransition `json:"failed_transitions"`
	// The type of error encountered by the price migration.
	Type QuotePreviewSubscriptionScheduleLastPriceMigrationErrorType `json:"type"`
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
	// ID of the promotion code to create a new discount for.
	PromotionCode *PromotionCode `json:"promotion_code"`
}
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEnd struct {
	// A precise Unix timestamp for the end of the invoice item period. Must be greater than or equal to `period.start`.
	Timestamp int64 `json:"timestamp"`
	// Select how to calculate the end of the invoice item period.
	Type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEndType `json:"type"`
}
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStart struct {
	// A precise Unix timestamp for the start of the invoice item period. Must be less than or equal to `period.end`.
	Timestamp int64 `json:"timestamp"`
	// Select how to calculate the start of the invoice item period.
	Type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStartType `json:"type"`
}
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriod struct {
	End   *QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodEnd   `json:"end"`
	Start *QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriodStart `json:"start"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
type QuotePreviewSubscriptionSchedulePhaseAddInvoiceItem struct {
	// The stackable discounts that will be applied to the item.
	Discounts []*QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemDiscount `json:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string                                          `json:"metadata"`
	Period   *QuotePreviewSubscriptionSchedulePhaseAddInvoiceItemPeriod `json:"period"`
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
	// If Stripe disabled automatic tax, this enum describes why.
	DisabledReason QuotePreviewSubscriptionSchedulePhaseAutomaticTaxDisabledReason `json:"disabled_reason"`
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
	// ID of the promotion code to create a new discount for.
	PromotionCode *PromotionCode `json:"promotion_code"`
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
	// The account tax IDs associated with this phase of the subscription schedule. Will be set on invoices generated by this phase of the subscription schedule.
	AccountTaxIDs []*TaxID `json:"account_tax_ids"`
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
	// ID of the promotion code to create a new discount for.
	PromotionCode *PromotionCode `json:"promotion_code"`
}

// Options that configure the trial on the subscription item.
type QuotePreviewSubscriptionSchedulePhaseItemTrial struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []string `json:"converts_to"`
	// Determines the type of trial for this item.
	Type QuotePreviewSubscriptionSchedulePhaseItemTrialType `json:"type"`
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

// If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://stripe.com/docs/billing/subscriptions/pause-payment).
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

// Defines how the subscription should behave when a trial ends.
type QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehavior struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront QuotePreviewSubscriptionSchedulePhaseTrialSettingsEndBehaviorProrateUpFront `json:"prorate_up_front"`
}

// Settings related to any trials on the subscription during this phase.
type QuotePreviewSubscriptionSchedulePhaseTrialSettings struct {
	// Defines how the subscription should behave when a trial ends.
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
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method"`
	// The default tax rates to apply to the subscription during this phase of the subscription schedule.
	DefaultTaxRates []*TaxRate `json:"default_tax_rates"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description string `json:"description"`
	// The stackable discounts that will be applied to the subscription on this phase. Subscription item discounts are applied before subscription discounts.
	Discounts []*QuotePreviewSubscriptionSchedulePhaseDiscount `json:"discounts"`
	// Configures how the subscription schedule handles billing for phase transitions. Possible values are `phase_start` (default) or `billing_period_start`. `phase_start` bills based on the current state of the subscription, ignoring changes scheduled in future phases. `billing_period_start` bills predictively for upcoming phase transitions within the current billing cycle, including pricing changes and service period adjustments that will occur before the next invoice.
	EffectiveAt QuotePreviewSubscriptionSchedulePhaseEffectiveAt `json:"effective_at"`
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
	// If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://stripe.com/docs/billing/subscriptions/pause-payment).
	PauseCollection *QuotePreviewSubscriptionSchedulePhasePauseCollection `json:"pause_collection"`
	// When transitioning phases, controls how prorations are handled (if any). Possible values are `create_prorations`, `none`, and `always_invoice`.
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

// Specifies which subscription items the billing schedule applies to.
type QuotePreviewSubscriptionScheduleBillingScheduleAppliesTo struct {
	// The billing schedule will apply to the subscription item with the given price ID.
	Price *Price `json:"price"`
	// Controls which subscription items the billing schedule applies to.
	Type QuotePreviewSubscriptionScheduleBillingScheduleAppliesToType `json:"type"`
}

// Use an index to specify the position of an amendment to start prebilling with.
type QuotePreviewSubscriptionScheduleBillingScheduleBillFromAmendmentStart struct {
	// Use an index to specify the position of an amendment to start prebilling with.
	Index int64 `json:"index"`
}

// Lets you bill the period starting from a particular Quote line.
type QuotePreviewSubscriptionScheduleBillingScheduleBillFromLineStartsAt struct {
	// Unique identifier for the object.
	ID string `json:"id"`
}

// Timestamp is calculated from the request time.
type QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelative struct {
	// Specifies billing duration. Possible values are `day`, `week`, `month`, or `year`.
	Interval QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelativeInterval `json:"interval"`
	// The multiplier applied to the interval.
	IntervalCount int64 `json:"interval_count"`
}

// Specifies the start of the billing period.
type QuotePreviewSubscriptionScheduleBillingScheduleBillFrom struct {
	// Use an index to specify the position of an amendment to start prebilling with.
	AmendmentStart *QuotePreviewSubscriptionScheduleBillingScheduleBillFromAmendmentStart `json:"amendment_start"`
	// The time the billing schedule applies from.
	ComputedTimestamp int64 `json:"computed_timestamp"`
	// Lets you bill the period starting from a particular Quote line.
	LineStartsAt *QuotePreviewSubscriptionScheduleBillingScheduleBillFromLineStartsAt `json:"line_starts_at"`
	// Timestamp is calculated from the request time.
	Relative *QuotePreviewSubscriptionScheduleBillingScheduleBillFromRelative `json:"relative"`
	// Use a precise Unix timestamp for prebilling to start. Must be earlier than `bill_until`.
	Timestamp int64 `json:"timestamp"`
	// Describes how the billing schedule determines the start date. Possible values are `timestamp`, `relative`, `amendment_start`, `now`, `quote_acceptance_date`, `line_starts_at`, or `pause_collection_start`.
	Type QuotePreviewSubscriptionScheduleBillingScheduleBillFromType `json:"type"`
}

// Use an index to specify the position of an amendment to end prebilling with.
type QuotePreviewSubscriptionScheduleBillingScheduleBillUntilAmendmentEnd struct {
	// Use an index to specify the position of an amendment to end prebilling with.
	Index int64 `json:"index"`
}

// Specifies the billing period.
type QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDuration struct {
	// Specifies billing duration. Either `day`, `week`, `month` or `year`.
	Interval QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDurationInterval `json:"interval"`
	// The multiplier applied to the interval.
	IntervalCount int64 `json:"interval_count"`
}

// Lets you bill the period ending at a particular Quote line.
type QuotePreviewSubscriptionScheduleBillingScheduleBillUntilLineEndsAt struct {
	// Unique identifier for the object.
	ID string `json:"id"`
}

// Specifies the end of billing period.
type QuotePreviewSubscriptionScheduleBillingScheduleBillUntil struct {
	// Use an index to specify the position of an amendment to end prebilling with.
	AmendmentEnd *QuotePreviewSubscriptionScheduleBillingScheduleBillUntilAmendmentEnd `json:"amendment_end"`
	// The timestamp the billing schedule will apply until.
	ComputedTimestamp int64 `json:"computed_timestamp"`
	// Specifies the billing period.
	Duration *QuotePreviewSubscriptionScheduleBillingScheduleBillUntilDuration `json:"duration"`
	// Lets you bill the period ending at a particular Quote line.
	LineEndsAt *QuotePreviewSubscriptionScheduleBillingScheduleBillUntilLineEndsAt `json:"line_ends_at"`
	// If specified, the billing schedule will apply until the specified timestamp.
	Timestamp int64 `json:"timestamp"`
	// Describes how the billing schedule will determine the end date. Either `duration` or `timestamp`.
	Type QuotePreviewSubscriptionScheduleBillingScheduleBillUntilType `json:"type"`
}

// Billing schedules for this subscription schedule.
type QuotePreviewSubscriptionScheduleBillingSchedule struct {
	// Specifies which subscription items the billing schedule applies to.
	AppliesTo []*QuotePreviewSubscriptionScheduleBillingScheduleAppliesTo `json:"applies_to"`
	// Specifies the start of the billing period.
	BillFrom *QuotePreviewSubscriptionScheduleBillingScheduleBillFrom `json:"bill_from"`
	// Specifies the end of billing period.
	BillUntil *QuotePreviewSubscriptionScheduleBillingScheduleBillUntil `json:"bill_until"`
	// Unique identifier for the billing schedule.
	Key string `json:"key"`
}
type QuotePreviewSubscriptionSchedule struct {
	// ID of the Connect Application that created the schedule.
	Application *Application                               `json:"application"`
	AppliesTo   *QuotePreviewSubscriptionScheduleAppliesTo `json:"applies_to"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time. `prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior QuotePreviewSubscriptionScheduleBillingBehavior `json:"billing_behavior"`
	// The billing mode of the subscription.
	BillingMode *QuotePreviewSubscriptionScheduleBillingMode `json:"billing_mode"`
	// Billing schedules for this subscription schedule.
	BillingSchedules []*QuotePreviewSubscriptionScheduleBillingSchedule `json:"billing_schedules"`
	// Time at which the subscription schedule was canceled. Measured in seconds since the Unix epoch.
	CanceledAt int64 `json:"canceled_at"`
	// Time at which the subscription schedule was completed. Measured in seconds since the Unix epoch.
	CompletedAt int64 `json:"completed_at"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
	CurrentPhase *QuotePreviewSubscriptionScheduleCurrentPhase `json:"current_phase"`
	// ID of the customer who owns the subscription schedule.
	Customer *Customer `json:"customer"`
	// ID of the account who owns the subscription schedule.
	CustomerAccount string                                           `json:"customer_account"`
	DefaultSettings *QuotePreviewSubscriptionScheduleDefaultSettings `json:"default_settings"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior QuotePreviewSubscriptionScheduleEndBehavior `json:"end_behavior"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Details of the most recent price migration that failed for the subscription schedule.
	LastPriceMigrationError *QuotePreviewSubscriptionScheduleLastPriceMigrationError `json:"last_price_migration_error"`
	// The most recent invoice this subscription schedule has generated.
	LatestInvoice *Invoice `json:"latest_invoice"`
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
