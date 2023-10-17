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

// Type of the account referenced.
type QuoteAutomaticTaxLiabilityType string

// List of values that QuoteAutomaticTaxLiabilityType can take
const (
	QuoteAutomaticTaxLiabilityTypeAccount QuoteAutomaticTaxLiabilityType = "account"
	QuoteAutomaticTaxLiabilityTypeSelf    QuoteAutomaticTaxLiabilityType = "self"
)

// The status of the most recent automated tax calculation for this quote.
type QuoteAutomaticTaxStatus string

// List of values that QuoteAutomaticTaxStatus can take
const (
	QuoteAutomaticTaxStatusComplete               QuoteAutomaticTaxStatus = "complete"
	QuoteAutomaticTaxStatusFailed                 QuoteAutomaticTaxStatus = "failed"
	QuoteAutomaticTaxStatusRequiresLocationInputs QuoteAutomaticTaxStatus = "requires_location_inputs"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay invoices at the end of the subscription cycle or on finalization using the default payment method attached to the subscription or customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically`.
type QuoteCollectionMethod string

// List of values that QuoteCollectionMethod can take
const (
	QuoteCollectionMethodChargeAutomatically QuoteCollectionMethod = "charge_automatically"
	QuoteCollectionMethodSendInvoice         QuoteCollectionMethod = "send_invoice"
)

// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
type QuoteComputedRecurringInterval string

// List of values that QuoteComputedRecurringInterval can take
const (
	QuoteComputedRecurringIntervalDay   QuoteComputedRecurringInterval = "day"
	QuoteComputedRecurringIntervalMonth QuoteComputedRecurringInterval = "month"
	QuoteComputedRecurringIntervalWeek  QuoteComputedRecurringInterval = "week"
	QuoteComputedRecurringIntervalYear  QuoteComputedRecurringInterval = "year"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason string

// List of values that QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason can take
const (
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonCustomerExempt       QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "customer_exempt"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonNotCollecting        QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "not_collecting"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonNotSubjectToTax      QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "not_subject_to_tax"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonNotSupported         QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "not_supported"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonPortionProductExempt QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "portion_product_exempt"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonPortionReducedRated  QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "portion_reduced_rated"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonPortionStandardRated QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "portion_standard_rated"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonProductExempt        QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonProductExemptHoliday QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt_holiday"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonProportionallyRated  QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "proportionally_rated"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonReducedRated         QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "reduced_rated"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonReverseCharge        QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "reverse_charge"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonStandardRated        QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "standard_rated"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonTaxableBasisReduced  QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "taxable_basis_reduced"
	QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonZeroRated            QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason = "zero_rated"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason string

// List of values that QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason can take
const (
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonCustomerExempt       QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "customer_exempt"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonNotCollecting        QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "not_collecting"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonNotSubjectToTax      QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "not_subject_to_tax"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonNotSupported         QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "not_supported"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonPortionProductExempt QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "portion_product_exempt"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonPortionReducedRated  QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "portion_reduced_rated"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonPortionStandardRated QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "portion_standard_rated"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonProductExempt        QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonProductExemptHoliday QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt_holiday"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonProportionallyRated  QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "proportionally_rated"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonReducedRated         QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "reduced_rated"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonReverseCharge        QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "reverse_charge"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonStandardRated        QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "standard_rated"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonTaxableBasisReduced  QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "taxable_basis_reduced"
	QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonZeroRated            QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason = "zero_rated"
)

// Type of the account referenced.
type QuoteInvoiceSettingsIssuerType string

// List of values that QuoteInvoiceSettingsIssuerType can take
const (
	QuoteInvoiceSettingsIssuerTypeAccount QuoteInvoiceSettingsIssuerType = "account"
	QuoteInvoiceSettingsIssuerTypeSelf    QuoteInvoiceSettingsIssuerType = "self"
)

// The status of the quote.
type QuoteStatus string

// List of values that QuoteStatus can take
const (
	QuoteStatusAccepted  QuoteStatus = "accepted"
	QuoteStatusAccepting QuoteStatus = "accepting"
	QuoteStatusCanceled  QuoteStatus = "canceled"
	QuoteStatusDraft     QuoteStatus = "draft"
	QuoteStatusOpen      QuoteStatus = "open"
	QuoteStatusStale     QuoteStatus = "stale"
)

// The reason this quote was marked as canceled.
type QuoteStatusDetailsCanceledReason string

// List of values that QuoteStatusDetailsCanceledReason can take
const (
	QuoteStatusDetailsCanceledReasonCanceled             QuoteStatusDetailsCanceledReason = "canceled"
	QuoteStatusDetailsCanceledReasonQuoteAccepted        QuoteStatusDetailsCanceledReason = "quote_accepted"
	QuoteStatusDetailsCanceledReasonQuoteExpired         QuoteStatusDetailsCanceledReason = "quote_expired"
	QuoteStatusDetailsCanceledReasonQuoteSuperseded      QuoteStatusDetailsCanceledReason = "quote_superseded"
	QuoteStatusDetailsCanceledReasonSubscriptionCanceled QuoteStatusDetailsCanceledReason = "subscription_canceled"
)

// The reason the quote was marked as stale.
type QuoteStatusDetailsStaleLastReasonType string

// List of values that QuoteStatusDetailsStaleLastReasonType can take
const (
	QuoteStatusDetailsStaleLastReasonTypeBillOnAcceptanceInvalid      QuoteStatusDetailsStaleLastReasonType = "bill_on_acceptance_invalid"
	QuoteStatusDetailsStaleLastReasonTypeLineInvalid                  QuoteStatusDetailsStaleLastReasonType = "line_invalid"
	QuoteStatusDetailsStaleLastReasonTypeMarkedStale                  QuoteStatusDetailsStaleLastReasonType = "marked_stale"
	QuoteStatusDetailsStaleLastReasonTypeSubscriptionCanceled         QuoteStatusDetailsStaleLastReasonType = "subscription_canceled"
	QuoteStatusDetailsStaleLastReasonTypeSubscriptionChanged          QuoteStatusDetailsStaleLastReasonType = "subscription_changed"
	QuoteStatusDetailsStaleLastReasonTypeSubscriptionExpired          QuoteStatusDetailsStaleLastReasonType = "subscription_expired"
	QuoteStatusDetailsStaleLastReasonTypeSubscriptionScheduleCanceled QuoteStatusDetailsStaleLastReasonType = "subscription_schedule_canceled"
	QuoteStatusDetailsStaleLastReasonTypeSubscriptionScheduleChanged  QuoteStatusDetailsStaleLastReasonType = "subscription_schedule_changed"
	QuoteStatusDetailsStaleLastReasonTypeSubscriptionScheduleReleased QuoteStatusDetailsStaleLastReasonType = "subscription_schedule_released"
)

// The type of method to specify the `bill_from` time.
type QuoteSubscriptionDataBillOnAcceptanceBillFromType string

// List of values that QuoteSubscriptionDataBillOnAcceptanceBillFromType can take
const (
	QuoteSubscriptionDataBillOnAcceptanceBillFromTypeLineStartsAt         QuoteSubscriptionDataBillOnAcceptanceBillFromType = "line_starts_at"
	QuoteSubscriptionDataBillOnAcceptanceBillFromTypeNow                  QuoteSubscriptionDataBillOnAcceptanceBillFromType = "now"
	QuoteSubscriptionDataBillOnAcceptanceBillFromTypePauseCollectionStart QuoteSubscriptionDataBillOnAcceptanceBillFromType = "pause_collection_start"
	QuoteSubscriptionDataBillOnAcceptanceBillFromTypeQuoteAcceptanceDate  QuoteSubscriptionDataBillOnAcceptanceBillFromType = "quote_acceptance_date"
	QuoteSubscriptionDataBillOnAcceptanceBillFromTypeTimestamp            QuoteSubscriptionDataBillOnAcceptanceBillFromType = "timestamp"
)

// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
type QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationInterval string

// List of values that QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationInterval can take
const (
	QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationIntervalDay   QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationInterval = "day"
	QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationIntervalMonth QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationInterval = "month"
	QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationIntervalWeek  QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationInterval = "week"
	QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationIntervalYear  QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationInterval = "year"
)

// The type of method to specify the `bill_until` time.
type QuoteSubscriptionDataBillOnAcceptanceBillUntilType string

// List of values that QuoteSubscriptionDataBillOnAcceptanceBillUntilType can take
const (
	QuoteSubscriptionDataBillOnAcceptanceBillUntilTypeDuration        QuoteSubscriptionDataBillOnAcceptanceBillUntilType = "duration"
	QuoteSubscriptionDataBillOnAcceptanceBillUntilTypeLineEndsAt      QuoteSubscriptionDataBillOnAcceptanceBillUntilType = "line_ends_at"
	QuoteSubscriptionDataBillOnAcceptanceBillUntilTypeScheduleEnd     QuoteSubscriptionDataBillOnAcceptanceBillUntilType = "schedule_end"
	QuoteSubscriptionDataBillOnAcceptanceBillUntilTypeTimestamp       QuoteSubscriptionDataBillOnAcceptanceBillUntilType = "timestamp"
	QuoteSubscriptionDataBillOnAcceptanceBillUntilTypeUpcomingInvoice QuoteSubscriptionDataBillOnAcceptanceBillUntilType = "upcoming_invoice"
)

// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
type QuoteSubscriptionDataBillingBehavior string

// List of values that QuoteSubscriptionDataBillingBehavior can take
const (
	QuoteSubscriptionDataBillingBehaviorProrateOnNextPhase QuoteSubscriptionDataBillingBehavior = "prorate_on_next_phase"
	QuoteSubscriptionDataBillingBehaviorProrateUpFront     QuoteSubscriptionDataBillingBehavior = "prorate_up_front"
)

// Whether the subscription will always start a new billing period when the quote is accepted.
type QuoteSubscriptionDataBillingCycleAnchor string

// List of values that QuoteSubscriptionDataBillingCycleAnchor can take
const (
	QuoteSubscriptionDataBillingCycleAnchorReset QuoteSubscriptionDataBillingCycleAnchor = "reset"
)

// Behavior of the subscription schedule and underlying subscription when it ends.
type QuoteSubscriptionDataEndBehavior string

// List of values that QuoteSubscriptionDataEndBehavior can take
const (
	QuoteSubscriptionDataEndBehaviorCancel  QuoteSubscriptionDataEndBehavior = "cancel"
	QuoteSubscriptionDataEndBehaviorRelease QuoteSubscriptionDataEndBehavior = "release"
)

// Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) when the quote is accepted.
type QuoteSubscriptionDataProrationBehavior string

// List of values that QuoteSubscriptionDataProrationBehavior can take
const (
	QuoteSubscriptionDataProrationBehaviorAlwaysInvoice    QuoteSubscriptionDataProrationBehavior = "always_invoice"
	QuoteSubscriptionDataProrationBehaviorCreateProrations QuoteSubscriptionDataProrationBehavior = "create_prorations"
	QuoteSubscriptionDataProrationBehaviorNone             QuoteSubscriptionDataProrationBehavior = "none"
)

// Describes whether the quote line is affecting a new schedule or an existing schedule.
type QuoteSubscriptionDataOverrideAppliesToType string

// List of values that QuoteSubscriptionDataOverrideAppliesToType can take
const (
	QuoteSubscriptionDataOverrideAppliesToTypeNewReference         QuoteSubscriptionDataOverrideAppliesToType = "new_reference"
	QuoteSubscriptionDataOverrideAppliesToTypeSubscriptionSchedule QuoteSubscriptionDataOverrideAppliesToType = "subscription_schedule"
)

// The type of method to specify the `bill_from` time.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromType string

// List of values that QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromType can take
const (
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromTypeLineStartsAt         QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromType = "line_starts_at"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromTypeNow                  QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromType = "now"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromTypePauseCollectionStart QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromType = "pause_collection_start"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromTypeQuoteAcceptanceDate  QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromType = "quote_acceptance_date"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromTypeTimestamp            QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromType = "timestamp"
)

// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationInterval string

// List of values that QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationInterval can take
const (
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationIntervalDay   QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationInterval = "day"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationIntervalMonth QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationInterval = "month"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationIntervalWeek  QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationInterval = "week"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationIntervalYear  QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationInterval = "year"
)

// The type of method to specify the `bill_until` time.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilType string

// List of values that QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilType can take
const (
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilTypeDuration        QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilType = "duration"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilTypeLineEndsAt      QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilType = "line_ends_at"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilTypeScheduleEnd     QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilType = "schedule_end"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilTypeTimestamp       QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilType = "timestamp"
	QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilTypeUpcomingInvoice QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilType = "upcoming_invoice"
)

// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
type QuoteSubscriptionDataOverrideBillingBehavior string

// List of values that QuoteSubscriptionDataOverrideBillingBehavior can take
const (
	QuoteSubscriptionDataOverrideBillingBehaviorProrateOnNextPhase QuoteSubscriptionDataOverrideBillingBehavior = "prorate_on_next_phase"
	QuoteSubscriptionDataOverrideBillingBehaviorProrateUpFront     QuoteSubscriptionDataOverrideBillingBehavior = "prorate_up_front"
)

// Behavior of the subscription schedule and underlying subscription when it ends.
type QuoteSubscriptionDataOverrideEndBehavior string

// List of values that QuoteSubscriptionDataOverrideEndBehavior can take
const (
	QuoteSubscriptionDataOverrideEndBehaviorCancel  QuoteSubscriptionDataOverrideEndBehavior = "cancel"
	QuoteSubscriptionDataOverrideEndBehaviorRelease QuoteSubscriptionDataOverrideEndBehavior = "release"
)

// Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) when the quote is accepted.
type QuoteSubscriptionDataOverrideProrationBehavior string

// List of values that QuoteSubscriptionDataOverrideProrationBehavior can take
const (
	QuoteSubscriptionDataOverrideProrationBehaviorAlwaysInvoice    QuoteSubscriptionDataOverrideProrationBehavior = "always_invoice"
	QuoteSubscriptionDataOverrideProrationBehaviorCreateProrations QuoteSubscriptionDataOverrideProrationBehavior = "create_prorations"
	QuoteSubscriptionDataOverrideProrationBehaviorNone             QuoteSubscriptionDataOverrideProrationBehavior = "none"
)

// Describes whether the quote line is affecting a new schedule or an existing schedule.
type QuoteSubscriptionScheduleAppliesToType string

// List of values that QuoteSubscriptionScheduleAppliesToType can take
const (
	QuoteSubscriptionScheduleAppliesToTypeNewReference         QuoteSubscriptionScheduleAppliesToType = "new_reference"
	QuoteSubscriptionScheduleAppliesToTypeSubscriptionSchedule QuoteSubscriptionScheduleAppliesToType = "subscription_schedule"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type QuoteTotalDetailsBreakdownTaxTaxabilityReason string

// List of values that QuoteTotalDetailsBreakdownTaxTaxabilityReason can take
const (
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonCustomerExempt       QuoteTotalDetailsBreakdownTaxTaxabilityReason = "customer_exempt"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonNotCollecting        QuoteTotalDetailsBreakdownTaxTaxabilityReason = "not_collecting"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonNotSubjectToTax      QuoteTotalDetailsBreakdownTaxTaxabilityReason = "not_subject_to_tax"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonNotSupported         QuoteTotalDetailsBreakdownTaxTaxabilityReason = "not_supported"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonPortionProductExempt QuoteTotalDetailsBreakdownTaxTaxabilityReason = "portion_product_exempt"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonPortionReducedRated  QuoteTotalDetailsBreakdownTaxTaxabilityReason = "portion_reduced_rated"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonPortionStandardRated QuoteTotalDetailsBreakdownTaxTaxabilityReason = "portion_standard_rated"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonProductExempt        QuoteTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonProductExemptHoliday QuoteTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt_holiday"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonProportionallyRated  QuoteTotalDetailsBreakdownTaxTaxabilityReason = "proportionally_rated"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonReducedRated         QuoteTotalDetailsBreakdownTaxTaxabilityReason = "reduced_rated"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonReverseCharge        QuoteTotalDetailsBreakdownTaxTaxabilityReason = "reverse_charge"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonStandardRated        QuoteTotalDetailsBreakdownTaxTaxabilityReason = "standard_rated"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonTaxableBasisReduced  QuoteTotalDetailsBreakdownTaxTaxabilityReason = "taxable_basis_reduced"
	QuoteTotalDetailsBreakdownTaxTaxabilityReasonZeroRated            QuoteTotalDetailsBreakdownTaxTaxabilityReason = "zero_rated"
)

// Retrieves the quote with the given ID.
type QuoteParams struct {
	Params `form:"*"`
	// Set to true to allow quote lines to have `starts_at` in the past if collection is paused between `starts_at` and now.
	AllowBackdatedLines *bool `form:"allow_backdated_lines"`
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. There cannot be any line items with recurring prices when using this field.
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. There must be at least 1 line item with a recurring price to use this field.
	ApplicationFeePercent *float64 `form:"application_fee_percent"`
	// Settings for automatic tax lookup for this quote and resulting invoices and subscriptions.
	AutomaticTax *QuoteAutomaticTaxParams `form:"automatic_tax"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay invoices at the end of the subscription cycle or at invoice finalization using the default payment method attached to the subscription or customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically`.
	CollectionMethod *string `form:"collection_method"`
	// The customer for which this quote belongs to. A customer is required before finalizing the quote. Once specified, it cannot be changed.
	Customer *string `form:"customer"`
	// The tax rates that will apply to any line item that does not have `tax_rates` set.
	DefaultTaxRates []*string `form:"default_tax_rates"`
	// A description that will be displayed on the quote PDF. If no value is passed, the default description configured in your [quote template settings](https://dashboard.stripe.com/settings/billing/quote) will be used.
	Description *string `form:"description"`
	// The discounts applied to the quote. You can only set up to one discount.
	Discounts []*QuoteDiscountParams `form:"discounts"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A future timestamp on which the quote will be canceled if in `open` or `draft` status. Measured in seconds since the Unix epoch. If no value is passed, the default expiration date configured in your [quote template settings](https://dashboard.stripe.com/settings/billing/quote) will be used.
	ExpiresAt *int64 `form:"expires_at"`
	// A footer that will be displayed on the quote PDF. If no value is passed, the default footer configured in your [quote template settings](https://dashboard.stripe.com/settings/billing/quote) will be used.
	Footer *string `form:"footer"`
	// Clone an existing quote. The new quote will be created in `status=draft`. When using this parameter, you cannot specify any other parameters except for `expires_at`.
	FromQuote *QuoteFromQuoteParams `form:"from_quote"`
	// A header that will be displayed on the quote PDF. If no value is passed, the default header configured in your [quote template settings](https://dashboard.stripe.com/settings/billing/quote) will be used.
	Header *string `form:"header"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *QuoteInvoiceSettingsParams `form:"invoice_settings"`
	// A list of line items the customer is being quoted for. Each line item includes information about the product, the quantity, and the resulting cost.
	LineItems []*QuoteLineItemParams `form:"line_items"`
	// A list of lines on the quote. These lines describe changes, in the order provided, that will be used to create new subscription schedules or update existing subscription schedules when the quote is accepted.
	Lines []*QuoteLineParams `form:"lines"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The account on behalf of which to charge.
	OnBehalfOf *string `form:"on_behalf_of"`
	// List representing phases of the Quote. Each phase can be customized to have different durations, prices, and coupons.
	Phases []*QuotePhaseParams `form:"phases"`
	// When creating a subscription or subscription schedule, the specified configuration data will be used. There must be at least one line item with a recurring price for a subscription or subscription schedule to be created. A subscription schedule is created if `subscription_data[effective_date]` is present and in the future, otherwise a subscription is created.
	SubscriptionData *QuoteSubscriptionDataParams `form:"subscription_data"`
	// List representing overrides for `subscription_data` configurations for specific groups.
	SubscriptionDataOverrides []*QuoteSubscriptionDataOverrideParams `form:"subscription_data_overrides"`
	// ID of the test clock to attach to the quote.
	TestClock *string `form:"test_clock"`
	// The data with which to automatically create a Transfer for each of the invoices.
	TransferData *QuoteTransferDataParams `form:"transfer_data"`
}

// AddExpand appends a new field to expand.
func (p *QuoteParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *QuoteParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type QuoteAutomaticTaxLiabilityParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// Settings for automatic tax lookup for this quote and resulting invoices and subscriptions.
type QuoteAutomaticTaxParams struct {
	// Controls whether Stripe will automatically compute tax on the resulting invoices or subscriptions as well as the quote itself.
	Enabled *bool `form:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *QuoteAutomaticTaxLiabilityParams `form:"liability"`
}

// Time span for the redeemed discount.
type QuoteDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type QuoteDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *QuoteDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The discounts applied to the quote. You can only set up to one discount.
type QuoteDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuoteDiscountDiscountEndParams `form:"discount_end"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type QuoteInvoiceSettingsIssuerParams struct {
	// The connected account being referenced when `type` is `account`.
	Account *string `form:"account"`
	// Type of the account referenced in the request.
	Type *string `form:"type"`
}

// All invoices will be billed using the specified settings.
type QuoteInvoiceSettingsParams struct {
	// Number of days within which a customer must pay the invoice generated by this quote. This value will be `null` for quotes where `collection_method=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *QuoteInvoiceSettingsIssuerParams `form:"issuer"`
}

// Time span for the redeemed discount.
type QuoteLineItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type QuoteLineItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *QuoteLineItemDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The discounts applied to this line item.
type QuoteLineItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuoteLineItemDiscountDiscountEndParams `form:"discount_end"`
}

// The recurring components of a price such as `interval` and `interval_count`.
type QuoteLineItemPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
type QuoteLineItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the product that this price will belong to.
	Product *string `form:"product"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *QuoteLineItemPriceDataRecurringParams `form:"recurring"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// A list of line items the customer is being quoted for. Each line item includes information about the product, the quantity, and the resulting cost.
type QuoteLineItemParams struct {
	// The discounts applied to this line item.
	Discounts []*QuoteLineItemDiscountParams `form:"discounts"`
	// The ID of an existing line item on the quote.
	ID *string `form:"id"`
	// The ID of the price object. One of `price` or `price_data` is required.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *QuoteLineItemPriceDataParams `form:"price_data"`
	// The quantity of the line item.
	Quantity *int64 `form:"quantity"`
	// The tax rates which apply to the line item. When set, the `default_tax_rates` on the quote do not apply to this line item.
	TaxRates []*string `form:"tax_rates"`
}

// Details to determine how long the discount should be applied for.
type QuoteLineActionAddDiscountDiscountEndParams struct {
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// Details for the `add_discount` type.
type QuoteLineActionAddDiscountParams struct {
	// The coupon code to redeem.
	Coupon *string `form:"coupon"`
	// An ID of an existing discount for a coupon that was already redeemed.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuoteLineActionAddDiscountDiscountEndParams `form:"discount_end"`
	// The index, starting at 0, at which to position the new discount. When not supplied, Stripe defaults to appending the discount to the end of the `discounts` array.
	Index *int64 `form:"index"`
}

// Time span for the redeemed discount.
type QuoteLineActionAddItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type QuoteLineActionAddItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *QuoteLineActionAddItemDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The discounts applied to the item. Subscription item discounts are applied before subscription discounts.
type QuoteLineActionAddItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuoteLineActionAddItemDiscountDiscountEndParams `form:"discount_end"`
}

// Options that configure the trial on the subscription item.
type QuoteLineActionAddItemTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []*string `form:"converts_to"`
	// Determines the type of trial for this item.
	Type *string `form:"type"`
}

// Details for the `add_item` type.
type QuoteLineActionAddItemParams struct {
	// The discounts applied to the item. Subscription item discounts are applied before subscription discounts.
	Discounts []*QuoteLineActionAddItemDiscountParams `form:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The ID of the price object.
	Price *string `form:"price"`
	// Quantity for this item.
	Quantity *int64 `form:"quantity"`
	// The tax rates that apply to this subscription item. When set, the `default_tax_rates` on the subscription do not apply to this `subscription_item`.
	TaxRates []*string `form:"tax_rates"`
	// Options that configure the trial on the subscription item.
	Trial *QuoteLineActionAddItemTrialParams `form:"trial"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *QuoteLineActionAddItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details for the `remove_discount` type.
type QuoteLineActionRemoveDiscountParams struct {
	// The coupon code to remove from the `discounts` array.
	Coupon *string `form:"coupon"`
	// The ID of a discount to remove from the `discounts` array.
	Discount *string `form:"discount"`
}

// Details for the `remove_item` type.
type QuoteLineActionRemoveItemParams struct {
	// ID of a price to remove.
	Price *string `form:"price"`
}

// Details for the `set_discounts` type.
type QuoteLineActionSetDiscountParams struct {
	// The coupon code to replace the `discounts` array with.
	Coupon *string `form:"coupon"`
	// An ID of an existing discount to replace the `discounts` array with.
	Discount *string `form:"discount"`
}

// Time span for the redeemed discount.
type QuoteLineActionSetItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type QuoteLineActionSetItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *QuoteLineActionSetItemDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// If an item with the `price` already exists, passing this will override the `discounts` array on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `discounts`.
type QuoteLineActionSetItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuoteLineActionSetItemDiscountDiscountEndParams `form:"discount_end"`
}

// If an item with the `price` already exists, passing this will override the `trial` configuration on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `trial`.
type QuoteLineActionSetItemTrialParams struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []*string `form:"converts_to"`
	// Determines the type of trial for this item.
	Type *string `form:"type"`
}

// Details for the `set_items` type.
type QuoteLineActionSetItemParams struct {
	// If an item with the `price` already exists, passing this will override the `discounts` array on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `discounts`.
	Discounts []*QuoteLineActionSetItemDiscountParams `form:"discounts"`
	// If an item with the `price` already exists, passing this will override the `metadata` on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The ID of the price object.
	Price *string `form:"price"`
	// If an item with the `price` already exists, passing this will override the quantity on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `quantity`.
	Quantity *int64 `form:"quantity"`
	// If an item with the `price` already exists, passing this will override the `tax_rates` array on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `tax_rates`.
	TaxRates []*string `form:"tax_rates"`
	// If an item with the `price` already exists, passing this will override the `trial` configuration on the subscription item that matches that price. Otherwise, the `items` array is cleared and a single new item is added with the supplied `trial`.
	Trial *QuoteLineActionSetItemTrialParams `form:"trial"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *QuoteLineActionSetItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// An array of operations the quote line performs.
type QuoteLineActionParams struct {
	// Details for the `add_discount` type.
	AddDiscount *QuoteLineActionAddDiscountParams `form:"add_discount"`
	// Details for the `add_item` type.
	AddItem *QuoteLineActionAddItemParams `form:"add_item"`
	// Details for the `add_metadata` type: specify a hash of key-value pairs.
	AddMetadata map[string]string `form:"add_metadata"`
	// Details for the `remove_discount` type.
	RemoveDiscount *QuoteLineActionRemoveDiscountParams `form:"remove_discount"`
	// Details for the `remove_item` type.
	RemoveItem *QuoteLineActionRemoveItemParams `form:"remove_item"`
	// Details for the `remove_metadata` type: specify an array of metadata keys.
	RemoveMetadata []*string `form:"remove_metadata"`
	// Details for the `set_discounts` type.
	SetDiscounts []*QuoteLineActionSetDiscountParams `form:"set_discounts"`
	// Details for the `set_items` type.
	SetItems []*QuoteLineActionSetItemParams `form:"set_items"`
	// Details for the `set_metadata` type: specify an array of key-value pairs.
	SetMetadata map[string]string `form:"set_metadata"`
	// The type of action the quote line performs.
	Type *string `form:"type"`
}

// Details to identify the subscription schedule the quote line applies to.
type QuoteLineAppliesToParams struct {
	// A custom string that identifies a new subscription schedule being created upon quote acceptance. All quote lines with the same `new_reference` field will be applied to the creation of a new subscription schedule.
	NewReference *string `form:"new_reference"`
	// The ID of the schedule the line applies to.
	SubscriptionSchedule *string `form:"subscription_schedule"`
	// Describes whether the quote line is affecting a new schedule or an existing schedule.
	Type *string `form:"type"`
}

// Use the `end` time of a given discount.
type QuoteLineEndsAtDiscountEndParams struct {
	// The ID of a specific discount.
	Discount *string `form:"discount"`
}

// Time span for the quote line starting from the `starts_at` date.
type QuoteLineEndsAtDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to identify the end of the time range modified by the proposed change. If not supplied, the quote line is considered a point-in-time operation that only affects the exact timestamp at `starts_at`, and a restricted set of attributes is supported on the quote line.
type QuoteLineEndsAtParams struct {
	// Use the `end` time of a given discount.
	DiscountEnd *QuoteLineEndsAtDiscountEndParams `form:"discount_end"`
	// Time span for the quote line starting from the `starts_at` date.
	Duration *QuoteLineEndsAtDurationParams `form:"duration"`
	// A precise Unix timestamp.
	Timestamp *int64 `form:"timestamp"`
	// Select a way to pass in `ends_at`.
	Type *string `form:"type"`
}

// Details of the pause_collection behavior to apply to the amendment.
type QuoteLineSetPauseCollectionSetParams struct {
	// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
	Behavior *string `form:"behavior"`
}

// Defines how to pause collection for the underlying subscription throughout the duration of the amendment.
type QuoteLineSetPauseCollectionParams struct {
	// Details of the pause_collection behavior to apply to the amendment.
	Set *QuoteLineSetPauseCollectionSetParams `form:"set"`
	// Determines the type of the pause_collection amendment.
	Type *string `form:"type"`
}

// Use the `end` time of a given discount.
type QuoteLineStartsAtDiscountEndParams struct {
	// The ID of a specific discount.
	Discount *string `form:"discount"`
}

// The timestamp the given line ends at.
type QuoteLineStartsAtLineEndsAtParams struct {
	// The ID of a quote line.
	ID *string `form:"id"`
	// The position of the previous quote line in the `lines` array after which this line should begin. Indexes start from 0 and must be less than the index of the current line in the array.
	Index *int64 `form:"index"`
}

// Details to identify the earliest timestamp where the proposed change should take effect.
type QuoteLineStartsAtParams struct {
	// Use the `end` time of a given discount.
	DiscountEnd *QuoteLineStartsAtDiscountEndParams `form:"discount_end"`
	// The timestamp the given line ends at.
	LineEndsAt *QuoteLineStartsAtLineEndsAtParams `form:"line_ends_at"`
	// A precise Unix timestamp.
	Timestamp *int64 `form:"timestamp"`
	// Select a way to pass in `starts_at`.
	Type *string `form:"type"`
}

// Defines how the subscription should behave when a trial ends.
type QuoteLineTrialSettingsEndBehaviorParams struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront *string `form:"prorate_up_front"`
}

// Settings related to subscription trials.
type QuoteLineTrialSettingsParams struct {
	// Defines how the subscription should behave when a trial ends.
	EndBehavior *QuoteLineTrialSettingsEndBehaviorParams `form:"end_behavior"`
}

// A list of lines on the quote. These lines describe changes, in the order provided, that will be used to create new subscription schedules or update existing subscription schedules when the quote is accepted.
type QuoteLineParams struct {
	// An array of operations the quote line performs.
	Actions []*QuoteLineActionParams `form:"actions"`
	// Details to identify the subscription schedule the quote line applies to.
	AppliesTo *QuoteLineAppliesToParams `form:"applies_to"`
	// For a point-in-time operation, this attribute lets you set or update whether the subscription's billing cycle anchor is reset at the `starts_at` timestamp.
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Details to identify the end of the time range modified by the proposed change. If not supplied, the quote line is considered a point-in-time operation that only affects the exact timestamp at `starts_at`, and a restricted set of attributes is supported on the quote line.
	EndsAt *QuoteLineEndsAtParams `form:"ends_at"`
	// The ID of an existing line on the quote.
	ID *string `form:"id"`
	// Changes to how Stripe handles prorations during the quote line's time span. Affects if and how prorations are created when a future phase starts.
	ProrationBehavior *string `form:"proration_behavior"`
	// Defines how to pause collection for the underlying subscription throughout the duration of the amendment.
	SetPauseCollection *QuoteLineSetPauseCollectionParams `form:"set_pause_collection"`
	// Timestamp helper to end the underlying schedule early, based on the acompanying line's start or end date.
	SetScheduleEnd *string `form:"set_schedule_end"`
	// Details to identify the earliest timestamp where the proposed change should take effect.
	StartsAt *QuoteLineStartsAtParams `form:"starts_at"`
	// Settings related to subscription trials.
	TrialSettings *QuoteLineTrialSettingsParams `form:"trial_settings"`
}

// Time span for the redeemed discount.
type QuotePhaseDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type QuotePhaseDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *QuotePhaseDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The coupons to redeem into discounts for the schedule phase. If not specified, inherits the discount from the subscription's customer. Pass an empty string to avoid inheriting any discounts.
type QuotePhaseDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuotePhaseDiscountDiscountEndParams `form:"discount_end"`
}

// All invoices will be billed using the specified settings.
type QuotePhaseInvoiceSettingsParams struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue *int64 `form:"days_until_due"`
}

// Time span for the redeemed discount.
type QuotePhaseLineItemDiscountDiscountEndDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details to determine how long the discount should be applied for.
type QuotePhaseLineItemDiscountDiscountEndParams struct {
	// Time span for the redeemed discount.
	Duration *QuotePhaseLineItemDiscountDiscountEndDurationParams `form:"duration"`
	// A precise Unix timestamp for the discount to end. Must be in the future.
	Timestamp *int64 `form:"timestamp"`
	// The type of calculation made to determine when the discount ends.
	Type *string `form:"type"`
}

// The discounts applied to this line item.
type QuotePhaseLineItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuotePhaseLineItemDiscountDiscountEndParams `form:"discount_end"`
}

// The recurring components of a price such as `interval` and `interval_count`.
type QuotePhaseLineItemPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
type QuotePhaseLineItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of the product that this price will belong to.
	Product *string `form:"product"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *QuotePhaseLineItemPriceDataRecurringParams `form:"recurring"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// A list of line items the customer is being quoted for within this phase. Each line item includes information about the product, the quantity, and the resulting cost.
type QuotePhaseLineItemParams struct {
	// The discounts applied to this line item.
	Discounts []*QuotePhaseLineItemDiscountParams `form:"discounts"`
	// The ID of the price object. One of `price` or `price_data` is required.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *QuotePhaseLineItemPriceDataParams `form:"price_data"`
	// The quantity of the line item.
	Quantity *int64 `form:"quantity"`
	// The tax rates which apply to the line item. When set, the `default_tax_rates` on the quote do not apply to this line item.
	TaxRates []*string `form:"tax_rates"`
}

// Details of a Quote line to start the bill period from.
type QuoteSubscriptionDataBillOnAcceptanceBillFromLineStartsAtParams struct {
	// The ID of a quote line.
	ID *string `form:"id"`
	// The position of the previous quote line in the `lines` array after which this line should begin. Indexes start from 0 and must be less than the index of the current line in the array.
	Index *int64 `form:"index"`
}

// The start of the period to bill from when the Quote is accepted.
type QuoteSubscriptionDataBillOnAcceptanceBillFromParams struct {
	// Details of a Quote line to start the bill period from.
	LineStartsAt *QuoteSubscriptionDataBillOnAcceptanceBillFromLineStartsAtParams `form:"line_starts_at"`
	// A precise Unix timestamp.
	Timestamp *int64 `form:"timestamp"`
	// The type of method to specify the `bill_from` time.
	Type *string `form:"type"`
}

// Details of the duration over which to bill.
type QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details of a Quote line item from which to bill until.
type QuoteSubscriptionDataBillOnAcceptanceBillUntilLineEndsAtParams struct {
	// The ID of a quote line.
	ID *string `form:"id"`
	// The position of the previous quote line in the `lines` array after which this line should begin. Indexes start from 0 and must be less than the index of the current line in the array.
	Index *int64 `form:"index"`
}

// The end of the period to bill until when the Quote is accepted.
type QuoteSubscriptionDataBillOnAcceptanceBillUntilParams struct {
	// Details of the duration over which to bill.
	Duration *QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationParams `form:"duration"`
	// Details of a Quote line item from which to bill until.
	LineEndsAt *QuoteSubscriptionDataBillOnAcceptanceBillUntilLineEndsAtParams `form:"line_ends_at"`
	// A precise Unix timestamp.
	Timestamp *int64 `form:"timestamp"`
	// The type of method to specify the `bill_until` time.
	Type *string `form:"type"`
}

// Describes the period to bill for upon accepting the quote.
type QuoteSubscriptionDataBillOnAcceptanceParams struct {
	// The start of the period to bill from when the Quote is accepted.
	BillFrom *QuoteSubscriptionDataBillOnAcceptanceBillFromParams `form:"bill_from"`
	// The end of the period to bill until when the Quote is accepted.
	BillUntil *QuoteSubscriptionDataBillOnAcceptanceBillUntilParams `form:"bill_until"`
}

// If specified, the invoicing for the given billing cycle iterations will be processed when the quote is accepted. Cannot be used with `effective_date`.
type QuoteSubscriptionDataPrebillingParams struct {
	// This is used to determine the number of billing cycles to prebill.
	Iterations *int64 `form:"iterations"`
}

// When creating a subscription or subscription schedule, the specified configuration data will be used. There must be at least one line item with a recurring price for a subscription or subscription schedule to be created. A subscription schedule is created if `subscription_data[effective_date]` is present and in the future, otherwise a subscription is created.
type QuoteSubscriptionDataParams struct {
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior *string `form:"billing_behavior"`
	// When specified as `reset`, the subscription will always start a new billing period when the quote is accepted.
	BillingCycleAnchor *string `form:"billing_cycle_anchor"`
	// Describes the period to bill for upon accepting the quote.
	BillOnAcceptance *QuoteSubscriptionDataBillOnAcceptanceParams `form:"bill_on_acceptance"`
	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description *string `form:"description"`
	// When creating a new subscription, the date of which the subscription schedule will start after the quote is accepted. When updating a subscription, the date of which the subscription will be updated using a subscription schedule. The special value `current_period_end` can be provided to update a subscription at the end of its current period. The `effective_date` is ignored if it is in the past when the quote is accepted.
	EffectiveDate                 *int64 `form:"effective_date"`
	EffectiveDateCurrentPeriodEnd *bool  `form:"-"` // See custom AppendTo
	// Behavior of the subscription schedule and underlying subscription when it ends.
	EndBehavior *string `form:"end_behavior"`
	// The id of a subscription schedule the quote will update. The quote will inherit the state of the subscription schedule, such as `phases`. Cannot be combined with other parameters.
	FromSchedule *string `form:"from_schedule"`
	// The id of a subscription that the quote will update. By default, the quote will contain the state of the subscription (such as line items, collection method and billing thresholds) unless overridden.
	FromSubscription *string `form:"from_subscription"`
	// If specified, the invoicing for the given billing cycle iterations will be processed when the quote is accepted. Cannot be used with `effective_date`.
	Prebilling *QuoteSubscriptionDataPrebillingParams `form:"prebilling"`
	// Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations). When creating a subscription, valid values are `create_prorations` or `none`.
	//
	// When updating a subscription, valid values are `create_prorations`, `none`, or `always_invoice`.
	//
	// Passing `create_prorations` will cause proration invoice items to be created when applicable. These proration items will only be invoiced immediately under [certain conditions](https://stripe.com/docs/subscriptions/upgrading-downgrading#immediate-payment). In order to always invoice immediately for prorations, pass `always_invoice`.
	//
	// Prorations can be disabled by passing `none`.
	ProrationBehavior *string `form:"proration_behavior"`
	// Integer representing the number of trial period days before the customer is charged for the first time.
	TrialPeriodDays *int64 `form:"trial_period_days"`
}

// AppendTo implements custom encoding logic for QuoteSubscriptionDataParams.
func (p *QuoteSubscriptionDataParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.EffectiveDateCurrentPeriodEnd) {
		body.Add(form.FormatKey(append(keyParts, "effective_date")), "current_period_end")
	}
}

// Whether the override applies to an existing Subscription Schedule or a new Subscription Schedule.
type QuoteSubscriptionDataOverrideAppliesToParams struct {
	// A custom string that identifies a new subscription schedule being created upon quote acceptance. All quote lines with the same `new_reference` field will be applied to the creation of a new subscription schedule.
	NewReference *string `form:"new_reference"`
	// The ID of the schedule the line applies to.
	SubscriptionSchedule *string `form:"subscription_schedule"`
	// Describes whether the quote line is affecting a new schedule or an existing schedule.
	Type *string `form:"type"`
}

// Details of a Quote line to start the bill period from.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromLineStartsAtParams struct {
	// The ID of a quote line.
	ID *string `form:"id"`
	// The position of the previous quote line in the `lines` array after which this line should begin. Indexes start from 0 and must be less than the index of the current line in the array.
	Index *int64 `form:"index"`
}

// The start of the period to bill from when the Quote is accepted.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromParams struct {
	// Details of a Quote line to start the bill period from.
	LineStartsAt *QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromLineStartsAtParams `form:"line_starts_at"`
	// A precise Unix timestamp.
	Timestamp *int64 `form:"timestamp"`
	// The type of method to specify the `bill_from` time.
	Type *string `form:"type"`
}

// Details of the duration over which to bill.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationParams struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount *int64 `form:"interval_count"`
}

// Details of a Quote line item from which to bill until.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilLineEndsAtParams struct {
	// The ID of a quote line.
	ID *string `form:"id"`
	// The position of the previous quote line in the `lines` array after which this line should begin. Indexes start from 0 and must be less than the index of the current line in the array.
	Index *int64 `form:"index"`
}

// The end of the period to bill until when the Quote is accepted.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilParams struct {
	// Details of the duration over which to bill.
	Duration *QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationParams `form:"duration"`
	// Details of a Quote line item from which to bill until.
	LineEndsAt *QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilLineEndsAtParams `form:"line_ends_at"`
	// A precise Unix timestamp.
	Timestamp *int64 `form:"timestamp"`
	// The type of method to specify the `bill_until` time.
	Type *string `form:"type"`
}

// Describes the period to bill for upon accepting the quote.
type QuoteSubscriptionDataOverrideBillOnAcceptanceParams struct {
	// The start of the period to bill from when the Quote is accepted.
	BillFrom *QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromParams `form:"bill_from"`
	// The end of the period to bill until when the Quote is accepted.
	BillUntil *QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilParams `form:"bill_until"`
}

// List representing overrides for `subscription_data` configurations for specific groups.
type QuoteSubscriptionDataOverrideParams struct {
	// Whether the override applies to an existing Subscription Schedule or a new Subscription Schedule.
	AppliesTo *QuoteSubscriptionDataOverrideAppliesToParams `form:"applies_to"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior *string `form:"billing_behavior"`
	// Describes the period to bill for upon accepting the quote.
	BillOnAcceptance *QuoteSubscriptionDataOverrideBillOnAcceptanceParams `form:"bill_on_acceptance"`
	// The customer the Subscription Data override applies to. This is only relevant when `applies_to.type=new_reference`.
	Customer *string `form:"customer"`
	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description *string `form:"description"`
	// Behavior of the subscription schedule and underlying subscription when it ends.
	EndBehavior *string `form:"end_behavior"`
	// Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations). When creating a subscription, valid values are `create_prorations` or `none`.
	//
	// When updating a subscription, valid values are `create_prorations`, `none`, or `always_invoice`.
	//
	// Passing `create_prorations` will cause proration invoice items to be created when applicable. These proration items will only be invoiced immediately under [certain conditions](https://stripe.com/docs/subscriptions/upgrading-downgrading#immediate-payment). In order to always invoice immediately for prorations, pass `always_invoice`.
	//
	// Prorations can be disabled by passing `none`.
	ProrationBehavior *string `form:"proration_behavior"`
}

// The data with which to automatically create a Transfer for each of the invoices.
type QuoteTransferDataParams struct {
	// The amount that will be transferred automatically when the invoice is paid. If no amount is set, the full amount is transferred. There cannot be any line items with recurring prices when using this field.
	Amount *int64 `form:"amount"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount is transferred to the destination. There must be at least 1 line item with a recurring price to use this field.
	AmountPercent *float64 `form:"amount_percent"`
	// ID of an existing, connected Stripe account.
	Destination *string `form:"destination"`
}

// Clone an existing quote. The new quote will be created in `status=draft`. When using this parameter, you cannot specify any other parameters except for `expires_at`.
type QuoteFromQuoteParams struct {
	// Whether this quote is a revision of the previous quote.
	IsRevision *bool `form:"is_revision"`
	// The `id` of the quote that will be cloned.
	Quote *string `form:"quote"`
}

// Returns a list of your quotes.
type QuoteListParams struct {
	ListParams `form:"*"`
	// The ID of the customer whose quotes will be retrieved.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The subscription which the quote updates.
	FromSubscription *string `form:"from_subscription"`
	// The status of the quote.
	Status *string `form:"status"`
	// Provides a list of quotes that are associated with the specified test clock. The response will not include quotes with test clocks if this and the customer parameter is not set.
	TestClock *string `form:"test_clock"`
}

// AddExpand appends a new field to expand.
func (p *QuoteListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Cancels the quote.
type QuoteCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuoteCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Finalizes the quote.
type QuoteFinalizeQuoteParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A future timestamp on which the quote will be canceled if in `open` or `draft` status. Measured in seconds since the Unix epoch.
	ExpiresAt *int64 `form:"expires_at"`
}

// AddExpand appends a new field to expand.
func (p *QuoteFinalizeQuoteParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Accepts the specified quote.
type QuoteAcceptParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuoteAcceptParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Recompute the upcoming invoice estimate for the quote.
type QuoteReestimateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuoteReestimateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Converts a stale quote to draft.
type QuoteMarkDraftParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuoteMarkDraftParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Converts a draft or open quote to stale.
type QuoteMarkStaleParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Reason the Quote is being marked stale.
	Reason *string `form:"reason"`
}

// AddExpand appends a new field to expand.
func (p *QuoteMarkStaleParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// When retrieving a quote, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type QuoteListLineItemsParams struct {
	ListParams `form:"*"`
	Quote      *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuoteListLineItemsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// When retrieving a quote, there is an includable [computed.upfront.line_items](https://stripe.com/docs/api/quotes/object#quote_object-computed-upfront-line_items) property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of upfront line items.
type QuoteListComputedUpfrontLineItemsParams struct {
	ListParams `form:"*"`
	Quote      *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuoteListComputedUpfrontLineItemsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a paginated list of lines for a quote. These lines describe changes that will be used to create new subscription schedules or update existing subscription schedules when the quote is accepted.
type QuoteListLinesParams struct {
	ListParams `form:"*"`
	Quote      *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuoteListLinesParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Download the PDF for a finalized quote
type QuotePDFParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuotePDFParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Preview the invoice line items that would be generated by accepting the quote.
type QuoteListPreviewInvoiceLinesParams struct {
	ListParams     `form:"*"`
	PreviewInvoice *string `form:"-"` // Included in URL
	Quote          *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuoteListPreviewInvoiceLinesParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
type QuoteAutomaticTaxLiability struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuoteAutomaticTaxLiabilityType `json:"type"`
}
type QuoteAutomaticTax struct {
	// Automatically calculate taxes
	Enabled bool `json:"enabled"`
	// The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.
	Liability *QuoteAutomaticTaxLiability `json:"liability"`
	// The status of the most recent automated tax calculation for this quote.
	Status QuoteAutomaticTaxStatus `json:"status"`
}

// The aggregated discounts.
type QuoteComputedRecurringTotalDetailsBreakdownDiscount struct {
	// The amount discounted.
	Amount int64 `json:"amount"`
	// A discount represents the actual application of a [coupon](https://stripe.com/docs/api#coupons) or [promotion code](https://stripe.com/docs/api#promotion_codes).
	// It contains information about when the discount began, when it will end, and what it is applied to.
	//
	// Related guide: [Applying discounts to subscriptions](https://stripe.com/docs/billing/subscriptions/discounts)
	Discount *Discount `json:"discount"`
}

// The aggregated tax amounts by rate.
type QuoteComputedRecurringTotalDetailsBreakdownTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.
	//
	// Related guide: [Tax rates](https://stripe.com/docs/billing/taxes/tax-rates)
	Rate *TaxRate `json:"rate"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
}
type QuoteComputedRecurringTotalDetailsBreakdown struct {
	// The aggregated discounts.
	Discounts []*QuoteComputedRecurringTotalDetailsBreakdownDiscount `json:"discounts"`
	// The aggregated tax amounts by rate.
	Taxes []*QuoteComputedRecurringTotalDetailsBreakdownTax `json:"taxes"`
}
type QuoteComputedRecurringTotalDetails struct {
	// This is the sum of all the discounts.
	AmountDiscount int64 `json:"amount_discount"`
	// This is the sum of all the shipping amounts.
	AmountShipping int64 `json:"amount_shipping"`
	// This is the sum of all the tax amounts.
	AmountTax int64                                        `json:"amount_tax"`
	Breakdown *QuoteComputedRecurringTotalDetailsBreakdown `json:"breakdown"`
}

// The definitive totals and line items the customer will be charged on a recurring basis. Takes into account the line items with recurring prices and discounts with `duration=forever` coupons only. Defaults to `null` if no inputted line items with recurring prices.
type QuoteComputedRecurring struct {
	// Total before any discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total after discounts and taxes are applied.
	AmountTotal int64 `json:"amount_total"`
	// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
	Interval QuoteComputedRecurringInterval `json:"interval"`
	// The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months.
	IntervalCount int64                               `json:"interval_count"`
	TotalDetails  *QuoteComputedRecurringTotalDetails `json:"total_details"`
}

// The aggregated discounts.
type QuoteComputedUpfrontTotalDetailsBreakdownDiscount struct {
	// The amount discounted.
	Amount int64 `json:"amount"`
	// A discount represents the actual application of a [coupon](https://stripe.com/docs/api#coupons) or [promotion code](https://stripe.com/docs/api#promotion_codes).
	// It contains information about when the discount began, when it will end, and what it is applied to.
	//
	// Related guide: [Applying discounts to subscriptions](https://stripe.com/docs/billing/subscriptions/discounts)
	Discount *Discount `json:"discount"`
}

// The aggregated tax amounts by rate.
type QuoteComputedUpfrontTotalDetailsBreakdownTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.
	//
	// Related guide: [Tax rates](https://stripe.com/docs/billing/taxes/tax-rates)
	Rate *TaxRate `json:"rate"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
}
type QuoteComputedUpfrontTotalDetailsBreakdown struct {
	// The aggregated discounts.
	Discounts []*QuoteComputedUpfrontTotalDetailsBreakdownDiscount `json:"discounts"`
	// The aggregated tax amounts by rate.
	Taxes []*QuoteComputedUpfrontTotalDetailsBreakdownTax `json:"taxes"`
}
type QuoteComputedUpfrontTotalDetails struct {
	// This is the sum of all the discounts.
	AmountDiscount int64 `json:"amount_discount"`
	// This is the sum of all the shipping amounts.
	AmountShipping int64 `json:"amount_shipping"`
	// This is the sum of all the tax amounts.
	AmountTax int64                                      `json:"amount_tax"`
	Breakdown *QuoteComputedUpfrontTotalDetailsBreakdown `json:"breakdown"`
}
type QuoteComputedUpfront struct {
	// Total before any discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total after discounts and taxes are applied.
	AmountTotal int64 `json:"amount_total"`
	// The line items that will appear on the next invoice after this quote is accepted. This does not include pending invoice items that exist on the customer but may still be included in the next invoice.
	LineItems    *LineItemList                     `json:"line_items"`
	TotalDetails *QuoteComputedUpfrontTotalDetails `json:"total_details"`
}
type QuoteComputed struct {
	// The definitive totals and line items the customer will be charged on a recurring basis. Takes into account the line items with recurring prices and discounts with `duration=forever` coupons only. Defaults to `null` if no inputted line items with recurring prices.
	Recurring *QuoteComputedRecurring `json:"recurring"`
	// The time at which the quote's estimated schedules and upcoming invoices were generated.
	UpdatedAt int64                 `json:"updated_at"`
	Upfront   *QuoteComputedUpfront `json:"upfront"`
}

// Details of the quote that was cloned. See the [cloning documentation](https://stripe.com/docs/quotes/clone) for more details.
type QuoteFromQuote struct {
	// Whether this quote is a revision of a different quote.
	IsRevision bool `json:"is_revision"`
	// The quote that was cloned.
	Quote *Quote `json:"quote"`
}

// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
type QuoteInvoiceSettingsIssuer struct {
	// The connected account being referenced when `type` is `account`.
	Account *Account `json:"account"`
	// Type of the account referenced.
	Type QuoteInvoiceSettingsIssuerType `json:"type"`
}

// All invoices will be billed using the specified settings.
type QuoteInvoiceSettings struct {
	// Number of days within which a customer must pay invoices generated by this quote. This value will be `null` for quotes where `collection_method=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
	// The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.
	Issuer *QuoteInvoiceSettingsIssuer `json:"issuer"`
}
type QuoteStatusDetailsCanceled struct {
	// The reason this quote was marked as canceled.
	Reason QuoteStatusDetailsCanceledReason `json:"reason"`
	// Time at which the quote was marked as canceled. Measured in seconds since the Unix epoch.
	TransitionedAt int64 `json:"transitioned_at"`
}
type QuoteStatusDetailsStaleLastReasonSubscriptionChanged struct {
	// The subscription's state before the quote was marked as stale.
	PreviousSubscription *Subscription `json:"previous_subscription"`
}
type QuoteStatusDetailsStaleLastReasonSubscriptionScheduleChanged struct {
	// The subscription schedule's state before the quote was marked as stale.
	PreviousSubscriptionSchedule *SubscriptionSchedule `json:"previous_subscription_schedule"`
}

// The most recent reason this quote was marked as stale.
type QuoteStatusDetailsStaleLastReason struct {
	// The ID of the line that is invalid if the stale reason type is `line_invalid`.
	LineInvalid string `json:"line_invalid"`
	// The user supplied mark stale reason.
	MarkedStale string `json:"marked_stale"`
	// The ID of the subscription that was canceled.
	SubscriptionCanceled string                                                `json:"subscription_canceled"`
	SubscriptionChanged  *QuoteStatusDetailsStaleLastReasonSubscriptionChanged `json:"subscription_changed"`
	// The ID of the subscription that was expired.
	SubscriptionExpired string `json:"subscription_expired"`
	// The ID of the subscription schedule that was canceled.
	SubscriptionScheduleCanceled string                                                        `json:"subscription_schedule_canceled"`
	SubscriptionScheduleChanged  *QuoteStatusDetailsStaleLastReasonSubscriptionScheduleChanged `json:"subscription_schedule_changed"`
	// The ID of the subscription schedule that was released.
	SubscriptionScheduleReleased string `json:"subscription_schedule_released"`
	// The reason the quote was marked as stale.
	Type QuoteStatusDetailsStaleLastReasonType `json:"type"`
}
type QuoteStatusDetailsStale struct {
	// Time at which the quote expires. Measured in seconds since the Unix epoch.
	ExpiresAt int64 `json:"expires_at"`
	// The most recent reason this quote was marked as stale.
	LastReason *QuoteStatusDetailsStaleLastReason `json:"last_reason"`
	// Time at which the stale reason was updated. Measured in seconds since the Unix epoch.
	LastUpdatedAt int64 `json:"last_updated_at"`
	// Time at which the quote was marked as stale. Measured in seconds since the Unix epoch.
	TransitionedAt int64 `json:"transitioned_at"`
}

// Details on when and why a quote has been marked as stale or canceled.
type QuoteStatusDetails struct {
	Canceled *QuoteStatusDetailsCanceled `json:"canceled"`
	Stale    *QuoteStatusDetailsStale    `json:"stale"`
}
type QuoteStatusTransitions struct {
	// The time that the quote was accepted. Measured in seconds since Unix epoch.
	AcceptedAt int64 `json:"accepted_at"`
	// The time that the quote was canceled. Measured in seconds since Unix epoch.
	CanceledAt int64 `json:"canceled_at"`
	// The time that the quote was finalized. Measured in seconds since Unix epoch.
	FinalizedAt int64 `json:"finalized_at"`
}

// The timestamp the given line starts at.
type QuoteSubscriptionDataBillOnAcceptanceBillFromLineStartsAt struct {
	// Unique identifier for the object.
	ID string `json:"id"`
}

// The start of the period to bill from when the Quote is accepted.
type QuoteSubscriptionDataBillOnAcceptanceBillFrom struct {
	// The materialized time.
	Computed int64 `json:"computed"`
	// The timestamp the given line starts at.
	LineStartsAt *QuoteSubscriptionDataBillOnAcceptanceBillFromLineStartsAt `json:"line_starts_at"`
	// A precise Unix timestamp.
	Timestamp int64 `json:"timestamp"`
	// The type of method to specify the `bill_from` time.
	Type QuoteSubscriptionDataBillOnAcceptanceBillFromType `json:"type"`
}

// Time span for the quote line starting from the `starts_at` date.
type QuoteSubscriptionDataBillOnAcceptanceBillUntilDuration struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval QuoteSubscriptionDataBillOnAcceptanceBillUntilDurationInterval `json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount int64 `json:"interval_count"`
}

// The timestamp the given line ends at.
type QuoteSubscriptionDataBillOnAcceptanceBillUntilLineEndsAt struct {
	// Unique identifier for the object.
	ID string `json:"id"`
}

// The end of the period to bill until when the Quote is accepted.
type QuoteSubscriptionDataBillOnAcceptanceBillUntil struct {
	// The materialized time.
	Computed int64 `json:"computed"`
	// Time span for the quote line starting from the `starts_at` date.
	Duration *QuoteSubscriptionDataBillOnAcceptanceBillUntilDuration `json:"duration"`
	// The timestamp the given line ends at.
	LineEndsAt *QuoteSubscriptionDataBillOnAcceptanceBillUntilLineEndsAt `json:"line_ends_at"`
	// A precise Unix timestamp.
	Timestamp int64 `json:"timestamp"`
	// The type of method to specify the `bill_until` time.
	Type QuoteSubscriptionDataBillOnAcceptanceBillUntilType `json:"type"`
}

// Describes the period to bill for upon accepting the quote.
type QuoteSubscriptionDataBillOnAcceptance struct {
	// The start of the period to bill from when the Quote is accepted.
	BillFrom *QuoteSubscriptionDataBillOnAcceptanceBillFrom `json:"bill_from"`
	// The end of the period to bill until when the Quote is accepted.
	BillUntil *QuoteSubscriptionDataBillOnAcceptanceBillUntil `json:"bill_until"`
}

// If specified, the invoicing for the given billing cycle iterations will be processed when the quote is accepted. Cannot be used with `effective_date`.
type QuoteSubscriptionDataPrebilling struct {
	Iterations int64 `json:"iterations"`
}
type QuoteSubscriptionData struct {
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior QuoteSubscriptionDataBillingBehavior `json:"billing_behavior"`
	// Whether the subscription will always start a new billing period when the quote is accepted.
	BillingCycleAnchor QuoteSubscriptionDataBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Describes the period to bill for upon accepting the quote.
	BillOnAcceptance *QuoteSubscriptionDataBillOnAcceptance `json:"bill_on_acceptance"`
	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description"`
	// When creating a new subscription, the date of which the subscription schedule will start after the quote is accepted. This date is ignored if it is in the past when the quote is accepted. Measured in seconds since the Unix epoch.
	EffectiveDate int64 `json:"effective_date"`
	// Behavior of the subscription schedule and underlying subscription when it ends.
	EndBehavior QuoteSubscriptionDataEndBehavior `json:"end_behavior"`
	// The id of the subscription schedule that will be updated when the quote is accepted.
	FromSchedule *SubscriptionSchedule `json:"from_schedule"`
	// The id of the subscription that will be updated when the quote is accepted.
	FromSubscription *Subscription `json:"from_subscription"`
	// If specified, the invoicing for the given billing cycle iterations will be processed when the quote is accepted. Cannot be used with `effective_date`.
	Prebilling *QuoteSubscriptionDataPrebilling `json:"prebilling"`
	// Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) when the quote is accepted.
	ProrationBehavior QuoteSubscriptionDataProrationBehavior `json:"proration_behavior"`
	// Integer representing the number of trial period days before the customer is charged for the first time.
	TrialPeriodDays int64 `json:"trial_period_days"`
}
type QuoteSubscriptionDataOverrideAppliesTo struct {
	// A custom string that identifies a new subscription schedule being created upon quote acceptance. All quote lines with the same `new_reference` field will be applied to the creation of a new subscription schedule.
	NewReference string `json:"new_reference"`
	// The ID of the schedule the line applies to.
	SubscriptionSchedule string `json:"subscription_schedule"`
	// Describes whether the quote line is affecting a new schedule or an existing schedule.
	Type QuoteSubscriptionDataOverrideAppliesToType `json:"type"`
}

// The timestamp the given line starts at.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromLineStartsAt struct {
	// Unique identifier for the object.
	ID string `json:"id"`
}

// The start of the period to bill from when the Quote is accepted.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillFrom struct {
	// The materialized time.
	Computed int64 `json:"computed"`
	// The timestamp the given line starts at.
	LineStartsAt *QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromLineStartsAt `json:"line_starts_at"`
	// A precise Unix timestamp.
	Timestamp int64 `json:"timestamp"`
	// The type of method to specify the `bill_from` time.
	Type QuoteSubscriptionDataOverrideBillOnAcceptanceBillFromType `json:"type"`
}

// Time span for the quote line starting from the `starts_at` date.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDuration struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDurationInterval `json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount int64 `json:"interval_count"`
}

// The timestamp the given line ends at.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilLineEndsAt struct {
	// Unique identifier for the object.
	ID string `json:"id"`
}

// The end of the period to bill until when the Quote is accepted.
type QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntil struct {
	// The materialized time.
	Computed int64 `json:"computed"`
	// Time span for the quote line starting from the `starts_at` date.
	Duration *QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilDuration `json:"duration"`
	// The timestamp the given line ends at.
	LineEndsAt *QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilLineEndsAt `json:"line_ends_at"`
	// A precise Unix timestamp.
	Timestamp int64 `json:"timestamp"`
	// The type of method to specify the `bill_until` time.
	Type QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntilType `json:"type"`
}

// Describes the period to bill for upon accepting the quote.
type QuoteSubscriptionDataOverrideBillOnAcceptance struct {
	// The start of the period to bill from when the Quote is accepted.
	BillFrom *QuoteSubscriptionDataOverrideBillOnAcceptanceBillFrom `json:"bill_from"`
	// The end of the period to bill until when the Quote is accepted.
	BillUntil *QuoteSubscriptionDataOverrideBillOnAcceptanceBillUntil `json:"bill_until"`
}
type QuoteSubscriptionDataOverride struct {
	AppliesTo *QuoteSubscriptionDataOverrideAppliesTo `json:"applies_to"`
	// Configures when the subscription schedule generates prorations for phase transitions. Possible values are `prorate_on_next_phase` or `prorate_up_front` with the default being `prorate_on_next_phase`. `prorate_on_next_phase` will apply phase changes and generate prorations at transition time.`prorate_up_front` will bill for all phases within the current billing cycle up front.
	BillingBehavior QuoteSubscriptionDataOverrideBillingBehavior `json:"billing_behavior"`
	// Describes the period to bill for upon accepting the quote.
	BillOnAcceptance *QuoteSubscriptionDataOverrideBillOnAcceptance `json:"bill_on_acceptance"`
	// The customer which this quote belongs to. A customer is required before finalizing the quote. Once specified, it cannot be changed.
	Customer string `json:"customer"`
	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description"`
	// Behavior of the subscription schedule and underlying subscription when it ends.
	EndBehavior QuoteSubscriptionDataOverrideEndBehavior `json:"end_behavior"`
	// Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) when the quote is accepted.
	ProrationBehavior QuoteSubscriptionDataOverrideProrationBehavior `json:"proration_behavior"`
}
type QuoteSubscriptionScheduleAppliesTo struct {
	// A custom string that identifies a new subscription schedule being created upon quote acceptance. All quote lines with the same `new_reference` field will be applied to the creation of a new subscription schedule.
	NewReference string `json:"new_reference"`
	// The ID of the schedule the line applies to.
	SubscriptionSchedule string `json:"subscription_schedule"`
	// Describes whether the quote line is affecting a new schedule or an existing schedule.
	Type QuoteSubscriptionScheduleAppliesToType `json:"type"`
}

// The subscription schedule that was created or updated from this quote.
type QuoteSubscriptionSchedule struct {
	AppliesTo *QuoteSubscriptionScheduleAppliesTo `json:"applies_to"`
	// The subscription schedule that was created or updated from this quote.
	SubscriptionSchedule string `json:"subscription_schedule"`
}

// The aggregated discounts.
type QuoteTotalDetailsBreakdownDiscount struct {
	// The amount discounted.
	Amount int64 `json:"amount"`
	// A discount represents the actual application of a [coupon](https://stripe.com/docs/api#coupons) or [promotion code](https://stripe.com/docs/api#promotion_codes).
	// It contains information about when the discount began, when it will end, and what it is applied to.
	//
	// Related guide: [Applying discounts to subscriptions](https://stripe.com/docs/billing/subscriptions/discounts)
	Discount *Discount `json:"discount"`
}

// The aggregated tax amounts by rate.
type QuoteTotalDetailsBreakdownTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.
	//
	// Related guide: [Tax rates](https://stripe.com/docs/billing/taxes/tax-rates)
	Rate *TaxRate `json:"rate"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason QuoteTotalDetailsBreakdownTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
}
type QuoteTotalDetailsBreakdown struct {
	// The aggregated discounts.
	Discounts []*QuoteTotalDetailsBreakdownDiscount `json:"discounts"`
	// The aggregated tax amounts by rate.
	Taxes []*QuoteTotalDetailsBreakdownTax `json:"taxes"`
}
type QuoteTotalDetails struct {
	// This is the sum of all the discounts.
	AmountDiscount int64 `json:"amount_discount"`
	// This is the sum of all the shipping amounts.
	AmountShipping int64 `json:"amount_shipping"`
	// This is the sum of all the tax amounts.
	AmountTax int64                       `json:"amount_tax"`
	Breakdown *QuoteTotalDetailsBreakdown `json:"breakdown"`
}

// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the invoices.
type QuoteTransferData struct {
	// The amount in cents (or local equivalent) that will be transferred to the destination account when the invoice is paid. By default, the entire amount is transferred to the destination.
	Amount int64 `json:"amount"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount will be transferred to the destination.
	AmountPercent float64 `json:"amount_percent"`
	// The account where funds from the payment will be transferred to upon payment success.
	Destination *Account `json:"destination"`
}

// A Quote is a way to model prices that you'd like to provide to a customer.
// Once accepted, it will automatically create an invoice, subscription or subscription schedule.
type Quote struct {
	APIResource
	// Allow quote lines to have `starts_at` in the past if collection is paused between `starts_at` and now.
	AllowBackdatedLines bool `json:"allow_backdated_lines"`
	// Total before any discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total after discounts and taxes are applied.
	AmountTotal int64 `json:"amount_total"`
	// ID of the Connect Application that created the quote.
	Application *Application `json:"application"`
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. Only applicable if there are no line items with recurring prices on the quote.
	ApplicationFeeAmount int64 `json:"application_fee_amount"`
	// A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. Only applicable if there are line items with recurring prices on the quote.
	ApplicationFeePercent float64            `json:"application_fee_percent"`
	AutomaticTax          *QuoteAutomaticTax `json:"automatic_tax"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay invoices at the end of the subscription cycle or on finalization using the default payment method attached to the subscription or customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically`.
	CollectionMethod QuoteCollectionMethod `json:"collection_method"`
	Computed         *QuoteComputed        `json:"computed"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The customer which this quote belongs to. A customer is required before finalizing the quote. Once specified, it cannot be changed.
	Customer *Customer `json:"customer"`
	// The tax rates applied to this quote.
	DefaultTaxRates []*TaxRate `json:"default_tax_rates"`
	// A description that will be displayed on the quote PDF.
	Description string `json:"description"`
	// The discounts applied to this quote.
	Discounts []*Discount `json:"discounts"`
	// The date on which the quote will be canceled if in `open` or `draft` status. Measured in seconds since the Unix epoch.
	ExpiresAt int64 `json:"expires_at"`
	// A footer that will be displayed on the quote PDF.
	Footer string `json:"footer"`
	// Details of the quote that was cloned. See the [cloning documentation](https://stripe.com/docs/quotes/clone) for more details.
	FromQuote *QuoteFromQuote `json:"from_quote"`
	// A header that will be displayed on the quote PDF.
	Header string `json:"header"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The invoice that was created from this quote.
	Invoice *Invoice `json:"invoice"`
	// All invoices will be billed using the specified settings.
	InvoiceSettings *QuoteInvoiceSettings `json:"invoice_settings"`
	// A list of items the customer is being quoted for.
	LineItems *LineItemList `json:"line_items"`
	// A list of lines on the quote. These lines describe changes, in the order provided, that will be used to create new subscription schedules or update existing subscription schedules when the quote is accepted.
	Lines []string `json:"lines"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// A unique number that identifies this particular quote. This number is assigned once the quote is [finalized](https://stripe.com/docs/quotes/overview#finalize).
	Number string `json:"number"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account on behalf of which to charge. See the [Connect documentation](https://support.stripe.com/questions/sending-invoices-on-behalf-of-connected-accounts) for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// The status of the quote.
	Status QuoteStatus `json:"status"`
	// Details on when and why a quote has been marked as stale or canceled.
	StatusDetails     *QuoteStatusDetails     `json:"status_details"`
	StatusTransitions *QuoteStatusTransitions `json:"status_transitions"`
	// The subscription that was created or updated from this quote.
	Subscription              *Subscription                    `json:"subscription"`
	SubscriptionData          *QuoteSubscriptionData           `json:"subscription_data"`
	SubscriptionDataOverrides []*QuoteSubscriptionDataOverride `json:"subscription_data_overrides"`
	// The subscription schedule that was created or updated from this quote.
	SubscriptionSchedule *SubscriptionSchedule `json:"subscription_schedule"`
	// The subscription schedules that were created or updated from this quote.
	SubscriptionSchedules []*QuoteSubscriptionSchedule `json:"subscription_schedules"`
	// ID of the test clock this quote belongs to.
	TestClock    *TestHelpersTestClock `json:"test_clock"`
	TotalDetails *QuoteTotalDetails    `json:"total_details"`
	// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the invoices.
	TransferData *QuoteTransferData `json:"transfer_data"`
}

// QuoteList is a list of Quotes as retrieved from a list endpoint.
type QuoteList struct {
	APIResource
	ListMeta
	Data []*Quote `json:"data"`
}

// UnmarshalJSON handles deserialization of a Quote.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (q *Quote) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		q.ID = id
		return nil
	}

	type quote Quote
	var v quote
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*q = Quote(v)
	return nil
}
