//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The discount end type.
type QuoteLineActionAddDiscountDiscountEndType string

// List of values that QuoteLineActionAddDiscountDiscountEndType can take
const (
	QuoteLineActionAddDiscountDiscountEndTypeLineEndsAt QuoteLineActionAddDiscountDiscountEndType = "line_ends_at"
)

// The discount end type.
type QuoteLineActionAddItemDiscountDiscountEndType string

// List of values that QuoteLineActionAddItemDiscountDiscountEndType can take
const (
	QuoteLineActionAddItemDiscountDiscountEndTypeTimestamp QuoteLineActionAddItemDiscountDiscountEndType = "timestamp"
)

type QuoteLineActionAddItemTrialType string

// List of values that QuoteLineActionAddItemTrialType can take
const (
	QuoteLineActionAddItemTrialTypeFree QuoteLineActionAddItemTrialType = "free"
	QuoteLineActionAddItemTrialTypePaid QuoteLineActionAddItemTrialType = "paid"
)

// The discount end type.
type QuoteLineActionRemoveDiscountDiscountEndType string

// List of values that QuoteLineActionRemoveDiscountDiscountEndType can take
const (
	QuoteLineActionRemoveDiscountDiscountEndTypeTimestamp QuoteLineActionRemoveDiscountDiscountEndType = "timestamp"
)

// The discount end type.
type QuoteLineActionSetDiscountDiscountEndType string

// List of values that QuoteLineActionSetDiscountDiscountEndType can take
const (
	QuoteLineActionSetDiscountDiscountEndTypeTimestamp QuoteLineActionSetDiscountDiscountEndType = "timestamp"
)

// The discount end type.
type QuoteLineActionSetItemDiscountDiscountEndType string

// List of values that QuoteLineActionSetItemDiscountDiscountEndType can take
const (
	QuoteLineActionSetItemDiscountDiscountEndTypeTimestamp QuoteLineActionSetItemDiscountDiscountEndType = "timestamp"
)

type QuoteLineActionSetItemTrialType string

// List of values that QuoteLineActionSetItemTrialType can take
const (
	QuoteLineActionSetItemTrialTypeFree QuoteLineActionSetItemTrialType = "free"
	QuoteLineActionSetItemTrialTypePaid QuoteLineActionSetItemTrialType = "paid"
)

// The type of action the quote line performs.
type QuoteLineActionType string

// List of values that QuoteLineActionType can take
const (
	QuoteLineActionTypeAddDiscount    QuoteLineActionType = "add_discount"
	QuoteLineActionTypeAddItem        QuoteLineActionType = "add_item"
	QuoteLineActionTypeAddMetadata    QuoteLineActionType = "add_metadata"
	QuoteLineActionTypeClearDiscounts QuoteLineActionType = "clear_discounts"
	QuoteLineActionTypeClearMetadata  QuoteLineActionType = "clear_metadata"
	QuoteLineActionTypeRemoveDiscount QuoteLineActionType = "remove_discount"
	QuoteLineActionTypeRemoveItem     QuoteLineActionType = "remove_item"
	QuoteLineActionTypeRemoveMetadata QuoteLineActionType = "remove_metadata"
	QuoteLineActionTypeSetDiscounts   QuoteLineActionType = "set_discounts"
	QuoteLineActionTypeSetItems       QuoteLineActionType = "set_items"
	QuoteLineActionTypeSetMetadata    QuoteLineActionType = "set_metadata"
)

// Describes whether the quote line is affecting a new schedule or an existing schedule.
type QuoteLineAppliesToType string

// List of values that QuoteLineAppliesToType can take
const (
	QuoteLineAppliesToTypeNewReference         QuoteLineAppliesToType = "new_reference"
	QuoteLineAppliesToTypeSubscriptionSchedule QuoteLineAppliesToType = "subscription_schedule"
)

// For a point-in-time operation, this attribute lets you set or update whether the subscription's billing cycle anchor is reset at the `starts_at` timestamp.
type QuoteLineBillingCycleAnchor string

// List of values that QuoteLineBillingCycleAnchor can take
const (
	QuoteLineBillingCycleAnchorAutomatic    QuoteLineBillingCycleAnchor = "automatic"
	QuoteLineBillingCycleAnchorLineStartsAt QuoteLineBillingCycleAnchor = "line_starts_at"
)

// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
type QuoteLineEndsAtDurationInterval string

// List of values that QuoteLineEndsAtDurationInterval can take
const (
	QuoteLineEndsAtDurationIntervalDay   QuoteLineEndsAtDurationInterval = "day"
	QuoteLineEndsAtDurationIntervalMonth QuoteLineEndsAtDurationInterval = "month"
	QuoteLineEndsAtDurationIntervalWeek  QuoteLineEndsAtDurationInterval = "week"
	QuoteLineEndsAtDurationIntervalYear  QuoteLineEndsAtDurationInterval = "year"
)

// Select a way to pass in `ends_at`.
type QuoteLineEndsAtType string

// List of values that QuoteLineEndsAtType can take
const (
	QuoteLineEndsAtTypeDiscountEnd         QuoteLineEndsAtType = "discount_end"
	QuoteLineEndsAtTypeDuration            QuoteLineEndsAtType = "duration"
	QuoteLineEndsAtTypeQuoteAcceptanceDate QuoteLineEndsAtType = "quote_acceptance_date"
	QuoteLineEndsAtTypeScheduleEnd         QuoteLineEndsAtType = "schedule_end"
	QuoteLineEndsAtTypeTimestamp           QuoteLineEndsAtType = "timestamp"
	QuoteLineEndsAtTypeUpcomingInvoice     QuoteLineEndsAtType = "upcoming_invoice"
)

// Changes to how Stripe handles prorations during the quote line's time span. Affects if and how prorations are created when a future phase starts.
type QuoteLineProrationBehavior string

// List of values that QuoteLineProrationBehavior can take
const (
	QuoteLineProrationBehaviorAlwaysInvoice    QuoteLineProrationBehavior = "always_invoice"
	QuoteLineProrationBehaviorCreateProrations QuoteLineProrationBehavior = "create_prorations"
	QuoteLineProrationBehaviorNone             QuoteLineProrationBehavior = "none"
)

// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
type QuoteLineSetPauseCollectionSetBehavior string

// List of values that QuoteLineSetPauseCollectionSetBehavior can take
const (
	QuoteLineSetPauseCollectionSetBehaviorKeepAsDraft       QuoteLineSetPauseCollectionSetBehavior = "keep_as_draft"
	QuoteLineSetPauseCollectionSetBehaviorMarkUncollectible QuoteLineSetPauseCollectionSetBehavior = "mark_uncollectible"
	QuoteLineSetPauseCollectionSetBehaviorVoid              QuoteLineSetPauseCollectionSetBehavior = "void"
)

// Defines the type of the pause_collection behavior for the quote line.
type QuoteLineSetPauseCollectionType string

// List of values that QuoteLineSetPauseCollectionType can take
const (
	QuoteLineSetPauseCollectionTypeRemove QuoteLineSetPauseCollectionType = "remove"
	QuoteLineSetPauseCollectionTypeSet    QuoteLineSetPauseCollectionType = "set"
)

// Timestamp helper to end the underlying schedule early, based on the acompanying line's start or end date.
type QuoteLineSetScheduleEnd string

// List of values that QuoteLineSetScheduleEnd can take
const (
	QuoteLineSetScheduleEndLineEndsAt   QuoteLineSetScheduleEnd = "line_ends_at"
	QuoteLineSetScheduleEndLineStartsAt QuoteLineSetScheduleEnd = "line_starts_at"
)

// Select a way to pass in `starts_at`.
type QuoteLineStartsAtType string

// List of values that QuoteLineStartsAtType can take
const (
	QuoteLineStartsAtTypeDiscountEnd         QuoteLineStartsAtType = "discount_end"
	QuoteLineStartsAtTypeLineEndsAt          QuoteLineStartsAtType = "line_ends_at"
	QuoteLineStartsAtTypeNow                 QuoteLineStartsAtType = "now"
	QuoteLineStartsAtTypeQuoteAcceptanceDate QuoteLineStartsAtType = "quote_acceptance_date"
	QuoteLineStartsAtTypeScheduleEnd         QuoteLineStartsAtType = "schedule_end"
	QuoteLineStartsAtTypeTimestamp           QuoteLineStartsAtType = "timestamp"
	QuoteLineStartsAtTypeUpcomingInvoice     QuoteLineStartsAtType = "upcoming_invoice"
)

// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
type QuoteLineTrialSettingsEndBehaviorProrateUpFront string

// List of values that QuoteLineTrialSettingsEndBehaviorProrateUpFront can take
const (
	QuoteLineTrialSettingsEndBehaviorProrateUpFrontDefer   QuoteLineTrialSettingsEndBehaviorProrateUpFront = "defer"
	QuoteLineTrialSettingsEndBehaviorProrateUpFrontInclude QuoteLineTrialSettingsEndBehaviorProrateUpFront = "include"
)

// Details to determine how long the discount should be applied for.
type QuoteLineActionAddDiscountDiscountEnd struct {
	// The discount end type.
	Type QuoteLineActionAddDiscountDiscountEndType `json:"type"`
}

// Details for the `add_discount` type.
type QuoteLineActionAddDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuoteLineActionAddDiscountDiscountEnd `json:"discount_end"`
	// The index, starting at 0, at which to position the new discount. When not supplied, Stripe defaults to appending the discount to the end of the `discounts` array.
	Index int64 `json:"index"`
}

// Details to determine how long the discount should be applied for.
type QuoteLineActionAddItemDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type QuoteLineActionAddItemDiscountDiscountEndType `json:"type"`
}

// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
type QuoteLineActionAddItemDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuoteLineActionAddItemDiscountDiscountEnd `json:"discount_end"`
}

// Options that configure the trial on the subscription item.
type QuoteLineActionAddItemTrial struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []string                        `json:"converts_to"`
	Type       QuoteLineActionAddItemTrialType `json:"type"`
}

// Details for the `add_item` type.
type QuoteLineActionAddItem struct {
	// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*QuoteLineActionAddItemDiscount `json:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an item. Metadata on this item will update the underlying subscription item's `metadata` when the phase is entered.
	Metadata map[string]string `json:"metadata"`
	// ID of the price to which the customer should be subscribed.
	Price *Price `json:"price"`
	// Quantity of the plan to which the customer should be subscribed.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to this `phase_item`. When set, the `default_tax_rates` on the phase do not apply to this `phase_item`.
	TaxRates []*TaxRate `json:"tax_rates"`
	// Options that configure the trial on the subscription item.
	Trial *QuoteLineActionAddItemTrial `json:"trial"`
}

// Details to determine how long the discount should be applied for.
type QuoteLineActionRemoveDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type QuoteLineActionRemoveDiscountDiscountEndType `json:"type"`
}

// Details for the `remove_discount` type.
type QuoteLineActionRemoveDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuoteLineActionRemoveDiscountDiscountEnd `json:"discount_end"`
}

// Details for the `remove_item` type.
type QuoteLineActionRemoveItem struct {
	// ID of a price to remove.
	Price *Price `json:"price"`
}

// Details to determine how long the discount should be applied for.
type QuoteLineActionSetDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type QuoteLineActionSetDiscountDiscountEndType `json:"type"`
}

// Details for the `set_discounts` type.
type QuoteLineActionSetDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuoteLineActionSetDiscountDiscountEnd `json:"discount_end"`
}

// Details to determine how long the discount should be applied for.
type QuoteLineActionSetItemDiscountDiscountEnd struct {
	// The discount end timestamp.
	Timestamp int64 `json:"timestamp"`
	// The discount end type.
	Type QuoteLineActionSetItemDiscountDiscountEndType `json:"type"`
}

// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
type QuoteLineActionSetItemDiscount struct {
	// ID of the coupon to create a new discount for.
	Coupon *Coupon `json:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *Discount `json:"discount"`
	// Details to determine how long the discount should be applied for.
	DiscountEnd *QuoteLineActionSetItemDiscountDiscountEnd `json:"discount_end"`
}

// Options that configure the trial on the subscription item.
type QuoteLineActionSetItemTrial struct {
	// List of price IDs which, if present on the subscription following a paid trial, constitute opting-in to the paid trial.
	ConvertsTo []string                        `json:"converts_to"`
	Type       QuoteLineActionSetItemTrialType `json:"type"`
}

// Details for the `set_items` type.
type QuoteLineActionSetItem struct {
	// The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.
	Discounts []*QuoteLineActionSetItemDiscount `json:"discounts"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an item. Metadata on this item will update the underlying subscription item's `metadata` when the phase is entered.
	Metadata map[string]string `json:"metadata"`
	// ID of the price to which the customer should be subscribed.
	Price *Price `json:"price"`
	// Quantity of the plan to which the customer should be subscribed.
	Quantity int64 `json:"quantity"`
	// The tax rates which apply to this `phase_item`. When set, the `default_tax_rates` on the phase do not apply to this `phase_item`.
	TaxRates []*TaxRate `json:"tax_rates"`
	// Options that configure the trial on the subscription item.
	Trial *QuoteLineActionSetItemTrial `json:"trial"`
}

// A list of items the customer is being quoted for.
type QuoteLineAction struct {
	// Details for the `add_discount` type.
	AddDiscount *QuoteLineActionAddDiscount `json:"add_discount"`
	// Details for the `add_item` type.
	AddItem *QuoteLineActionAddItem `json:"add_item"`
	// Details for the `add_metadata` type: specify a hash of key-value pairs.
	AddMetadata map[string]string `json:"add_metadata"`
	// Details for the `remove_discount` type.
	RemoveDiscount *QuoteLineActionRemoveDiscount `json:"remove_discount"`
	// Details for the `remove_item` type.
	RemoveItem *QuoteLineActionRemoveItem `json:"remove_item"`
	// Details for the `remove_metadata` type: specify an array of metadata keys.
	RemoveMetadata []string `json:"remove_metadata"`
	// Details for the `set_discounts` type.
	SetDiscounts []*QuoteLineActionSetDiscount `json:"set_discounts"`
	// Details for the `set_items` type.
	SetItems []*QuoteLineActionSetItem `json:"set_items"`
	// Details for the `set_metadata` type: specify an array of key-value pairs.
	SetMetadata map[string]string `json:"set_metadata"`
	// The type of action the quote line performs.
	Type QuoteLineActionType `json:"type"`
}

// Details to identify the subscription schedule the quote line applies to.
type QuoteLineAppliesTo struct {
	// A custom string that identifies a new subscription schedule being created upon quote acceptance. All quote lines with the same `new_reference` field will be applied to the creation of a new subscription schedule.
	NewReference string `json:"new_reference"`
	// The ID of the schedule the line applies to.
	SubscriptionSchedule string `json:"subscription_schedule"`
	// Describes whether the quote line is affecting a new schedule or an existing schedule.
	Type QuoteLineAppliesToType `json:"type"`
}

// Use the `end` time of a given discount.
type QuoteLineEndsAtDiscountEnd struct {
	// The ID of a specific discount.
	Discount string `json:"discount"`
}

// Time span for the quote line starting from the `starts_at` date.
type QuoteLineEndsAtDuration struct {
	// Specifies a type of interval unit. Either `day`, `week`, `month` or `year`.
	Interval QuoteLineEndsAtDurationInterval `json:"interval"`
	// The number of intervals, as an whole number greater than 0. Stripe multiplies this by the interval type to get the overall duration.
	IntervalCount int64 `json:"interval_count"`
}

// Details to identify the end of the time range modified by the proposed change. If not supplied, the quote line is considered a point-in-time operation that only affects the exact timestamp at `starts_at`, and a restricted set of attributes is supported on the quote line.
type QuoteLineEndsAt struct {
	// The timestamp value that will be used to determine when to make changes to the subscription schedule, as computed from the `ends_at` field. For example, if `ends_at[type]=upcoming_invoice`, the upcoming invoice date will be computed at the time the `ends_at` field was specified and saved. This field will not be recomputed upon future requests to update or finalize the quote unless `ends_at` is respecified. This field is guaranteed to be populated after quote acceptance.
	Computed int64 `json:"computed"`
	// Use the `end` time of a given discount.
	DiscountEnd *QuoteLineEndsAtDiscountEnd `json:"discount_end"`
	// Time span for the quote line starting from the `starts_at` date.
	Duration *QuoteLineEndsAtDuration `json:"duration"`
	// A precise Unix timestamp.
	Timestamp int64 `json:"timestamp"`
	// Select a way to pass in `ends_at`.
	Type QuoteLineEndsAtType `json:"type"`
}

// If specified, payment collection for this subscription will be paused.
type QuoteLineSetPauseCollectionSet struct {
	// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
	Behavior QuoteLineSetPauseCollectionSetBehavior `json:"behavior"`
}

// Details to modify the pause_collection behavior of the subscription schedule.
type QuoteLineSetPauseCollection struct {
	// If specified, payment collection for this subscription will be paused.
	Set *QuoteLineSetPauseCollectionSet `json:"set"`
	// Defines the type of the pause_collection behavior for the quote line.
	Type QuoteLineSetPauseCollectionType `json:"type"`
}

// Use the `end` time of a given discount.
type QuoteLineStartsAtDiscountEnd struct {
	// The ID of a specific discount.
	Discount string `json:"discount"`
}

// The timestamp the given line ends at.
type QuoteLineStartsAtLineEndsAt struct {
	// Unique identifier for the object.
	ID string `json:"id"`
}

// Details to identify the earliest timestamp where the proposed change should take effect.
type QuoteLineStartsAt struct {
	// The timestamp value that will be used to determine when to make changes to the subscription schedule, as computed from the `starts_at` field. For example, if `starts_at[type]=upcoming_invoice`, the upcoming invoice date will be computed at the time the `starts_at` field was specified and saved. This field will not be recomputed upon future requests to update or finalize the quote unless `starts_at` is respecified. This field is guaranteed to be populated after quote acceptance.
	Computed int64 `json:"computed"`
	// Use the `end` time of a given discount.
	DiscountEnd *QuoteLineStartsAtDiscountEnd `json:"discount_end"`
	// The timestamp the given line ends at.
	LineEndsAt *QuoteLineStartsAtLineEndsAt `json:"line_ends_at"`
	// A precise Unix timestamp.
	Timestamp int64 `json:"timestamp"`
	// Select a way to pass in `starts_at`.
	Type QuoteLineStartsAtType `json:"type"`
}

// Defines how the subscription should behave when a trial ends.
type QuoteLineTrialSettingsEndBehavior struct {
	// Configure how an opt-in following a paid trial is billed when using `billing_behavior: prorate_up_front`.
	ProrateUpFront QuoteLineTrialSettingsEndBehaviorProrateUpFront `json:"prorate_up_front"`
}

// Settings related to subscription trials.
type QuoteLineTrialSettings struct {
	// Defines how the subscription should behave when a trial ends.
	EndBehavior *QuoteLineTrialSettingsEndBehavior `json:"end_behavior"`
}

// A quote line defines a set of changes, in the order provided, that will be applied upon quote acceptance.
type QuoteLine struct {
	// A list of items the customer is being quoted for.
	Actions []*QuoteLineAction `json:"actions"`
	// Details to identify the subscription schedule the quote line applies to.
	AppliesTo *QuoteLineAppliesTo `json:"applies_to"`
	// For a point-in-time operation, this attribute lets you set or update whether the subscription's billing cycle anchor is reset at the `starts_at` timestamp.
	BillingCycleAnchor QuoteLineBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Details to identify the end of the time range modified by the proposed change. If not supplied, the quote line is considered a point-in-time operation that only affects the exact timestamp at `starts_at`, and a restricted set of attributes is supported on the quote line.
	EndsAt *QuoteLineEndsAt `json:"ends_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Changes to how Stripe handles prorations during the quote line's time span. Affects if and how prorations are created when a future phase starts.
	ProrationBehavior QuoteLineProrationBehavior `json:"proration_behavior"`
	// Details to modify the pause_collection behavior of the subscription schedule.
	SetPauseCollection *QuoteLineSetPauseCollection `json:"set_pause_collection"`
	// Timestamp helper to end the underlying schedule early, based on the acompanying line's start or end date.
	SetScheduleEnd QuoteLineSetScheduleEnd `json:"set_schedule_end"`
	// Details to identify the earliest timestamp where the proposed change should take effect.
	StartsAt *QuoteLineStartsAt `json:"starts_at"`
	// Settings related to subscription trials.
	TrialSettings *QuoteLineTrialSettings `json:"trial_settings"`
}

// QuoteLineList is a list of QuoteLines as retrieved from a list endpoint.
type QuoteLineList struct {
	APIResource
	ListMeta
	Data []*QuoteLine `json:"data"`
}
