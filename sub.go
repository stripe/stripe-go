package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/form"
)

// SubscriptionStatus is the list of allowed values for the subscription's status.
// Allowed values are "trialing", "active", "past_due", "canceled", "unpaid", "all".
type SubscriptionStatus string

// SubscriptionBilling is the type of billing method for this subscription's invoices.
// Currently supported values are "send_invoice" and "charge_automatically".
type SubscriptionBilling string

// SubscriptionParams is the set of parameters that can be used when creating or updating a subscription.
// For more details see https://stripe.com/docs/api#create_subscription and https://stripe.com/docs/api#update_subscription.
type SubscriptionParams struct {
	Params                      `form:"*"`
	ApplicationFeePercent       float64                    `form:"application_fee_percent"`
	ApplicationFeePercentZero   bool                       `form:"application_fee_percent,zero"`
	Billing                     SubscriptionBilling        `form:"billing"`
	BillingCycleAnchor          int64                      `form:"billing_cycle_anchor"`
	BillingCycleAnchorNow       bool                       `form:"-"` // See custom AppendTo
	BillingCycleAnchorUnchanged bool                       `form:"-"` // See custom AppendTo
	Card                        *CardParams                `form:"card"`
	Coupon                      string                     `form:"coupon"`
	CouponEmpty                 bool                       `form:"coupon,empty"`
	Customer                    string                     `form:"customer"`
	DaysUntilDue                uint64                     `form:"days_until_due"`
	Items                       []*SubscriptionItemsParams `form:"items,indexed"`
	NoProrate                   bool                       `form:"prorate,invert"`
	OnBehalfOf                  string                     `form:"on_behalf_of"`
	Plan                        string                     `form:"plan"`
	ProrationDate               int64                      `form:"proration_date"`
	Quantity                    uint64                     `form:"quantity"`
	QuantityZero                bool                       `form:"quantity,zero"`
	Source                      string                     `form:"source"`
	TaxPercent                  float64                    `form:"tax_percent"`
	TaxPercentZero              bool                       `form:"tax_percent,zero"`
	TrialEnd                    int64                      `form:"trial_end"`
	TrialEndNow                 bool                       `form:"-"` // See custom AppendTo
	TrialPeriodDays             int64                      `form:"trial_period_days"`

	// Used for Cancel

	AtPeriodEnd bool `form:"at_period_end"`
}

// AppendTo implements custom encoding logic for SubscriptionParams so that the special
// "now" value for billing_cycle_anchor and trial_end can be implemented
// (they're otherwise timestamps rather than strings).
func (p *SubscriptionParams) AppendTo(body *form.Values, keyParts []string) {
	if p.BillingCycleAnchorNow {
		body.Add(form.FormatKey(append(keyParts, "billing_cycle_anchor")), "now")
	}

	if p.BillingCycleAnchorUnchanged {
		body.Add(form.FormatKey(append(keyParts, "billing_cycle_anchor")), "unchanged")
	}

	if p.TrialEndNow {
		body.Add(form.FormatKey(append(keyParts, "trial_end")), "now")
	}
}

// SubscriptionItemsParams is the set of parameters that can be used when creating or updating a subscription item on a subscription
// For more details see https://stripe.com/docs/api#create_subscription and https://stripe.com/docs/api#update_subscription.
type SubscriptionItemsParams struct {
	Params       `form:"*"`
	Deleted      bool   `form:"deleted"`
	ID           string `form:"id"`
	Plan         string `form:"plan"`
	Quantity     uint64 `form:"quantity"`
	QuantityZero bool   `form:"quantity,zero"`
}

// SubscriptionListParams is the set of parameters that can be used when listing active subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
type SubscriptionListParams struct {
	ListParams   `form:"*"`
	Billing      SubscriptionBilling `form:"billing"`
	Created      int64               `form:"created"`
	CreatedRange *RangeQueryParams   `form:"created"`
	Customer     string              `form:"customer"`
	Plan         string              `form:"plan"`
	Status       SubscriptionStatus  `form:"status"`
}

// Subscription is the resource representing a Stripe subscription.
// For more details see https://stripe.com/docs/api#subscriptions.
type Subscription struct {
	ApplicationFeePercent float64               `json:"application_fee_percent"`
	Billing               SubscriptionBilling   `json:"billing"`
	BillingCycleAnchor    int64                 `json:"billing_cycle_anchor"`
	CanceledAt            int64                 `json:"canceled_at"`
	Created               int64                 `json:"created"`
	CurrentPeriodEnd      int64                 `json:"current_period_end"`
	CurrentPeriodStart    int64                 `json:"current_period_start"`
	Customer              *Customer             `json:"customer"`
	DaysUntilDue          uint64                `json:"days_until_due"`
	Discount              *Discount             `json:"discount"`
	CancelAtPeriodEnd     bool                  `json:"cancel_at_period_end"`
	EndedAt               int64                 `json:"ended_at"`
	ID                    string                `json:"id"`
	Items                 *SubscriptionItemList `json:"items"`
	Metadata              map[string]string     `json:"metadata"`
	Plan                  *Plan                 `json:"plan"`
	Quantity              uint64                `json:"quantity"`
	Start                 int64                 `json:"start"`
	Status                SubscriptionStatus    `json:"status"`
	TaxPercent            float64               `json:"tax_percent"`
	TrialEnd              int64                 `json:"trial_end"`
	TrialStart            int64                 `json:"trial_start"`
}

// SubscriptionList is a list object for subscriptions.
type SubscriptionList struct {
	ListMeta
	Data []*Subscription `json:"data"`
}

// UnmarshalJSON handles deserialization of a Subscription.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *Subscription) UnmarshalJSON(data []byte) error {
	type sub Subscription
	var ss sub
	err := json.Unmarshal(data, &ss)
	if err == nil {
		*s = Subscription(ss)
	} else {
		// the id is surrounded by "\" characters, so strip them
		s.ID = string(data[1 : len(data)-1])
	}

	return nil
}
