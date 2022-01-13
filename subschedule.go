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

// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type SubscriptionSchedulePhaseBillingCycleAnchor string

// List of values that SubscriptionSchedulePhaseBillingCycleAnchor can take
const (
	SubscriptionSchedulePhaseBillingCycleAnchorAutomatic  SubscriptionSchedulePhaseBillingCycleAnchor = "automatic"
	SubscriptionSchedulePhaseBillingCycleAnchorPhaseStart SubscriptionSchedulePhaseBillingCycleAnchor = "phase_start"
)

// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` and `cancel`.
type SubscriptionScheduleEndBehavior string

// List of values that SubscriptionScheduleEndBehavior can take
const (
	SubscriptionScheduleEndBehaviorCancel  SubscriptionScheduleEndBehavior = "cancel"
	SubscriptionScheduleEndBehaviorNone    SubscriptionScheduleEndBehavior = "none"
	SubscriptionScheduleEndBehaviorRelease SubscriptionScheduleEndBehavior = "release"
	SubscriptionScheduleEndBehaviorRenew   SubscriptionScheduleEndBehavior = "renew"
)

// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
type SubscriptionSchedulePhaseProrationBehavior string

// List of values that SubscriptionSchedulePhaseProrationBehavior can take
const (
	SubscriptionSchedulePhaseProrationBehaviorAlwaysInvoice    SubscriptionSchedulePhaseProrationBehavior = "always_invoice"
	SubscriptionSchedulePhaseProrationBehaviorCreateProrations SubscriptionSchedulePhaseProrationBehavior = "create_prorations"
	SubscriptionSchedulePhaseProrationBehaviorNone             SubscriptionSchedulePhaseProrationBehavior = "none"
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
	ListParams       `form:"*"`
	CanceledAt       int64             `form:"canceled_at"`
	CanceledAtRange  *RangeQueryParams `form:"canceled_at"`
	CompletedAt      int64             `form:"completed_at"`
	CompletedAtRange *RangeQueryParams `form:"completed_at"`
	Created          int64             `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	Customer         string            `form:"customer"`
	ReleasedAt       int64             `form:"released_at"`
	ReleasedAtRange  *RangeQueryParams `form:"released_at"`
	Scheduled        *bool             `form:"scheduled"`
}

// All invoices will be billed using the specified settings.
type SubscriptionScheduleInvoiceSettingsParams struct {
	DaysUntilDue *int64 `form:"days_until_due"`
}

// Object representing the subscription schedule's default settings.
type SubscriptionScheduleDefaultSettingsParams struct {
	Params                `form:"*"`
	ApplicationFeePercent *float64                                   `form:"application_fee_percent,high_precision"`
	AutomaticTax          *SubscriptionAutomaticTaxParams            `form:"automatic_tax"`
	BillingCycleAnchor    *string                                    `form:"billing_cycle_anchor"`
	BillingThresholds     *SubscriptionBillingThresholdsParams       `form:"billing_thresholds"`
	CollectionMethod      *string                                    `form:"collection_method"`
	DefaultPaymentMethod  *string                                    `form:"default_payment_method"`
	InvoiceSettings       *SubscriptionScheduleInvoiceSettingsParams `form:"invoice_settings"`
	TransferData          *SubscriptionTransferDataParams            `form:"transfer_data"`
}

// SubscriptionSchedulePhaseAddInvoiceItemPriceDataRecurringParams is a structure representing the
// parameters to create an inline recurring price for a given invoice item.
type SubscriptionSchedulePhaseAddInvoiceItemPriceDataRecurringParams struct {
	AggregateUsage  *string `form:"aggregate_usage"`
	Interval        *string `form:"interval"`
	IntervalCount   *int64  `form:"interval_count"`
	TrialPeriodDays *int64  `form:"trial_period_days"`
	UsageType       *string `form:"usage_type"`
}

// SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams is a structure representing the parameters
// to create an inline price for a given invoice item.
type SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams struct {
	Currency          *string                                                          `form:"currency"`
	Product           *string                                                          `form:"product"`
	Recurring         *SubscriptionSchedulePhaseAddInvoiceItemPriceDataRecurringParams `form:"recurring"`
	TaxBehavior       *string                                                          `form:"tax_behavior"`
	UnitAmount        *int64                                                           `form:"unit_amount"`
	UnitAmountDecimal *float64                                                         `form:"unit_amount_decimal,high_precision"`
}

// A list of prices and quantities that will generate invoice items appended to the next invoice. You may pass up to 20 items.
type SubscriptionSchedulePhaseAddInvoiceItemParams struct {
	Price     *string                     `form:"price"`
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	Quantity  *int64                      `form:"quantity"`
	TaxRates  []*string                   `form:"tax_rates"`
}

// Automatic tax settings for this phase.
type SubscriptionSchedulePhaseAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// List of configuration items, each with an attached price, to apply during this phase of the subscription schedule.
type SubscriptionSchedulePhaseItemParams struct {
	BillingThresholds *SubscriptionItemBillingThresholdsParams `form:"billing_thresholds"`
	Plan              *string                                  `form:"plan"`
	Price             *string                                  `form:"price"`
	PriceData         *SubscriptionItemPriceDataParams         `form:"price_data"`
	Quantity          *int64                                   `form:"quantity"`
	TaxRates          []*string                                `form:"tax_rates"`
}

// List representing phases of the subscription schedule. Each phase can be customized to have different durations, plans, and coupons. If there are multiple phases, the `end_date` of one phase will always equal the `start_date` of the next phase.
type SubscriptionSchedulePhaseParams struct {
	AddInvoiceItems       []*SubscriptionSchedulePhaseAddInvoiceItemParams `form:"add_invoice_items"`
	ApplicationFeePercent *float64                                         `form:"application_fee_percent"`
	AutomaticTax          *SubscriptionSchedulePhaseAutomaticTaxParams     `form:"automatic_tax"`
	BillingCycleAnchor    *string                                          `form:"billing_cycle_anchor"`
	BillingThresholds     *SubscriptionBillingThresholdsParams             `form:"billing_thresholds"`
	CollectionMethod      *string                                          `form:"collection_method"`
	Coupon                *string                                          `form:"coupon"`
	DefaultPaymentMethod  *string                                          `form:"default_payment_method"`
	DefaultTaxRates       []*string                                        `form:"default_tax_rates"`
	EndDate               *int64                                           `form:"end_date"`
	EndDateNow            *bool                                            `form:"-"` // See custom AppendTo
	InvoiceSettings       *SubscriptionScheduleInvoiceSettingsParams       `form:"invoice_settings"`
	Items                 []*SubscriptionSchedulePhaseItemParams           `form:"items"`
	Iterations            *int64                                           `form:"iterations"`
	ProrationBehavior     *string                                          `form:"proration_behavior"`
	StartDate             *int64                                           `form:"start_date"`
	StartDateNow          *bool                                            `form:"-"` // See custom AppendTo
	TransferData          *SubscriptionTransferDataParams                  `form:"transfer_data"`
	Trial                 *bool                                            `form:"trial"`
	TrialEnd              *int64                                           `form:"trial_end"`
	TrialEndNow           *bool                                            `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for SubscriptionSchedulePhaseParams.
func (s *SubscriptionSchedulePhaseParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(s.EndDateNow) {
		body.Add(form.FormatKey(append(keyParts, "end_date")), "now")
	}
	if BoolValue(s.TrialEndNow) {
		body.Add(form.FormatKey(append(keyParts, "trial_end")), "now")
	}
	if BoolValue(s.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
type SubscriptionScheduleParams struct {
	Params            `form:"*"`
	Customer          *string                                    `form:"customer"`
	DefaultSettings   *SubscriptionScheduleDefaultSettingsParams `form:"default_settings"`
	EndBehavior       *string                                    `form:"end_behavior"`
	FromSubscription  *string                                    `form:"from_subscription"`
	Phases            []*SubscriptionSchedulePhaseParams         `form:"phases"`
	ProrationBehavior *string                                    `form:"proration_behavior"`
	StartDate         *int64                                     `form:"start_date"`
	StartDateNow      *bool                                      `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for SubscriptionScheduleParams.
func (s *SubscriptionScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(s.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// Cancels a subscription schedule and its associated subscription immediately (if the subscription schedule has an active subscription). A subscription schedule can only be canceled if its status is not_started or active.
type SubscriptionScheduleCancelParams struct {
	Params     `form:"*"`
	InvoiceNow *bool `form:"invoice_now"`
	Prorate    *bool `form:"prorate"`
}

// Releases the subscription schedule immediately, which will stop scheduling of its phases, but leave any existing subscription in place. A schedule can only be released if its status is not_started or active. If the subscription schedule is currently associated with a subscription, releasing it will remove its subscription property and set the subscription's ID to the released_subscription property.
type SubscriptionScheduleReleaseParams struct {
	Params             `form:"*"`
	PreserveCancelDate *bool `form:"preserve_cancel_date"`
}

// Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.
type SubscriptionScheduleCurrentPhase struct {
	EndDate   int64 `json:"end_date"`
	StartDate int64 `json:"start_date"`
}

// The subscription schedule's default invoice settings.
type SubscriptionScheduleInvoiceSettings struct {
	DaysUntilDue int64 `json:"days_until_due"`
}
type SubscriptionScheduleDefaultSettings struct {
	ApplicationFeePercent float64                                     `json:"application_fee_percent,string"`
	AutomaticTax          *SubscriptionAutomaticTax                   `json:"automatic_tax"`
	BillingCycleAnchor    SubscriptionSchedulePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	BillingThresholds     *SubscriptionBillingThresholds              `json:"billing_thresholds"`
	CollectionMethod      SubscriptionCollectionMethod                `json:"collection_method"`
	DefaultPaymentMethod  *PaymentMethod                              `json:"default_payment_method"`
	InvoiceSettings       *SubscriptionScheduleInvoiceSettings        `json:"invoice_settings"`
	TransferData          *SubscriptionTransferData                   `json:"transfer_data"`
}

// A list of prices and quantities that will generate invoice items appended to the first invoice for this phase.
type SubscriptionSchedulePhaseAddInvoiceItem struct {
	Price    *Price     `json:"price"`
	Quantity int64      `json:"quantity"`
	TaxRates []*TaxRate `json:"tax_rates"`
}

// Subscription items to configure the subscription to during this phase of the subscription schedule.
type SubscriptionSchedulePhaseItem struct {
	BillingThresholds *SubscriptionItemBillingThresholds `json:"billing_thresholds"`
	Plan              *Plan                              `json:"plan"`
	Price             *Price                             `json:"price"`
	Quantity          int64                              `json:"quantity"`
	TaxRates          []*TaxRate                         `json:"tax_rates"`
}

// Configuration for the subscription schedule's phases.
type SubscriptionSchedulePhase struct {
	AddInvoiceItems       []*SubscriptionSchedulePhaseAddInvoiceItem  `json:"add_invoice_items"`
	ApplicationFeePercent float64                                     `json:"application_fee_percent"`
	AutomaticTax          *SubscriptionAutomaticTax                   `json:"automatic_tax"`
	BillingCycleAnchor    SubscriptionSchedulePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	BillingThresholds     *SubscriptionBillingThresholds              `json:"billing_thresholds"`
	CollectionMethod      SubscriptionCollectionMethod                `json:"collection_method"`
	Coupon                *Coupon                                     `json:"coupon"`
	DefaultPaymentMethod  *PaymentMethod                              `json:"default_payment_method"`
	DefaultTaxRates       []*TaxRate                                  `json:"default_tax_rates"`
	EndDate               int64                                       `json:"end_date"`
	InvoiceSettings       *SubscriptionScheduleInvoiceSettings        `json:"invoice_settings"`
	Items                 []*SubscriptionSchedulePhaseItem            `json:"items"`
	ProrationBehavior     SubscriptionSchedulePhaseProrationBehavior  `json:"proration_behavior"`
	StartDate             int64                                       `json:"start_date"`
	TransferData          *SubscriptionTransferData                   `json:"transfer_data"`
	TrialEnd              int64                                       `json:"trial_end"`
}

// A subscription schedule allows you to create and manage the lifecycle of a subscription by predefining expected changes.
//
// Related guide: [Subscription Schedules](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
type SubscriptionSchedule struct {
	APIResource
	CanceledAt           int64                                `json:"canceled_at"`
	CompletedAt          int64                                `json:"completed_at"`
	Created              int64                                `json:"created"`
	CurrentPhase         *SubscriptionScheduleCurrentPhase    `json:"current_phase"`
	Customer             *Customer                            `json:"customer"`
	DefaultSettings      *SubscriptionScheduleDefaultSettings `json:"default_settings"`
	EndBehavior          SubscriptionScheduleEndBehavior      `json:"end_behavior"`
	ID                   string                               `json:"id"`
	Livemode             bool                                 `json:"livemode"`
	Metadata             map[string]string                    `json:"metadata"`
	Object               string                               `json:"object"`
	Phases               []*SubscriptionSchedulePhase         `json:"phases"`
	ReleasedAt           int64                                `json:"released_at"`
	ReleasedSubscription *Subscription                        `json:"released_subscription"`
	Status               SubscriptionScheduleStatus           `json:"status"`
	Subscription         *Subscription                        `json:"subscription"`
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
