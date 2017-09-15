package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/form"
)

// SubStatus is the list of allowed values for the subscription's status.
// Allowed values are "trialing", "active", "past_due", "canceled", "unpaid", "all".
type SubStatus string

// SubBilling is the type of billing method for this subscription's invoices.
// Currently supported values are "send_invoice" and "charge_automatically".
type SubBilling string

// SubParams is the set of parameters that can be used when creating or updating a subscription.
// For more details see https://stripe.com/docs/api#create_subscription and https://stripe.com/docs/api#update_subscription.
type SubParams struct {
	Params                `form:"*"`
	Billing               SubBilling        `form:"billing"`
	BillingCycleAnchor    int64             `form:"billing_cycle_anchor"`
	BillingCycleAnchorNow bool              `form:"-"` // See custom AppendTo
	Card                  *CardParams       `form:"card"`
	Coupon                string            `form:"coupon"`
	CouponEmpty           bool              `form:"coupon,empty"`
	Customer              string            `form:"customer"`
	DaysUntilDue          uint64            `form:"days_until_due"`
	FeePercent            float64           `form:"application_fee_percent"`
	FeePercentZero        bool              `form:"application_fee_percent,zero"`
	Items                 []*SubItemsParams `form:"items,indexed"`
	NoProrate             bool              `form:"prorate,invert"`
	OnBehalfOf            string            `form:"on_behalf_of"`
	Plan                  string            `form:"plan"`
	ProrationDate         int64             `form:"proration_date"`
	Quantity              uint64            `form:"quantity"`
	QuantityZero          bool              `form:"quantity,zero"`
	TaxPercent            float64           `form:"tax_percent"`
	TaxPercentZero        bool              `form:"tax_percent,zero"`
	Token                 string            `form:"card"`
	TrialEnd              int64             `form:"trial_end"`
	TrialEndNow           bool              `form:"-"` // See custom AppendTo
	TrialPeriod           int64             `form:"trial_period_days"`

	// Used for Cancel

	EndCancel bool `form:"at_period_end"`
}

// AppendTo implements custom encoding logic for SubParams so that the special
// "now" value for billing_cycle_anchor and trial_end can be implemented
// (they're otherwise timestamps rather than strings).
func (p *SubParams) AppendTo(body *form.Values, keyParts []string) {
	if p.BillingCycleAnchorNow {
		body.Add(form.FormatKey(append(keyParts, "billing_cycle_anchor")), "now")
	}

	if p.TrialEndNow {
		body.Add(form.FormatKey(append(keyParts, "trial_end")), "now")
	}
}

// SubItemsParams is the set of parameters that can be used when creating or updating a subscription item on a subscription
// For more details see https://stripe.com/docs/api#create_subscription and https://stripe.com/docs/api#update_subscription.
type SubItemsParams struct {
	Params       `form:"*"`
	Deleted      bool   `form:"deleted"`
	ID           string `form:"id"`
	Plan         string `form:"plan"`
	Quantity     uint64 `form:"quantity"`
	QuantityZero bool   `form:"quantity,zero"`
}

// SubListParams is the set of parameters that can be used when listing active subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
type SubListParams struct {
	ListParams   `form:"*"`
	Billing      SubBilling        `form:"billing"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Customer     string            `form:"customer"`
	Plan         string            `form:"plan"`
	Status       SubStatus         `form:"status"`
}

// Sub is the resource representing a Stripe subscription.
// For more details see https://stripe.com/docs/api#subscriptions.
type Sub struct {
	Billing      SubBilling        `json:"billing"`
	Canceled     int64             `json:"canceled_at"`
	Created      int64             `json:"created"`
	Customer     *Customer         `json:"customer"`
	DaysUntilDue uint64            `json:"days_until_due"`
	Discount     *Discount         `json:"discount"`
	EndCancel    bool              `json:"cancel_at_period_end"`
	Ended        int64             `json:"ended_at"`
	FeePercent   float64           `json:"application_fee_percent"`
	ID           string            `json:"id"`
	Items        *SubItemList      `json:"items"`
	Meta         map[string]string `json:"metadata"`
	PeriodEnd    int64             `json:"current_period_end"`
	PeriodStart  int64             `json:"current_period_start"`
	Plan         *Plan             `json:"plan"`
	Quantity     uint64            `json:"quantity"`
	Start        int64             `json:"start"`
	Status       SubStatus         `json:"status"`
	TaxPercent   float64           `json:"tax_percent"`
	TrialEnd     int64             `json:"trial_end"`
	TrialStart   int64             `json:"trial_start"`
}

// SubList is a list object for subscriptions.
type SubList struct {
	ListMeta
	Values []*Sub `json:"data"`
}

// UnmarshalJSON handles deserialization of a Sub.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *Sub) UnmarshalJSON(data []byte) error {
	type sub Sub
	var ss sub
	err := json.Unmarshal(data, &ss)
	if err == nil {
		*s = Sub(ss)
	} else {
		// the id is surrounded by "\" characters, so strip them
		s.ID = string(data[1 : len(data)-1])
	}

	return nil
}
