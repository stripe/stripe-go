//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "github.com/stripe/stripe-go/v72/form"

// Possible values for the action parameter on usage record creation.
const (
	UsageRecordActionIncrement string = "increment"
	UsageRecordActionSet       string = "set"
)

// Creates a usage record for a specified subscription item and date, and fills it with a quantity.
//
// Usage records provide quantity information that Stripe uses to track how much a customer is using your service. With usage information and the pricing model set up by the [metered billing](https://stripe.com/docs/billing/subscriptions/metered-billing) plan, Stripe helps you send accurate invoices to your customers.
//
// The default calculation for usage is to add up all the quantity values of the usage records within a billing period. You can change this default behavior with the billing plan's aggregate_usage [parameter](https://stripe.com/docs/api/plans/create#create_plan-aggregate_usage). When there is more than one usage record with the same timestamp, Stripe adds the quantity values together. In most cases, this is the desired resolution, however, you can change this behavior with the action parameter.
//
// The default pricing model for metered billing is [per-unit pricing. For finer granularity, you can configure metered billing to have a <a href="https://stripe.com/docs/billing/subscriptions/tiers">tiered pricing](https://stripe.com/docs/api/plans/object#plan_object-billing_scheme) model.
type UsageRecordParams struct {
	Params           `form:"*"`
	SubscriptionItem *string `form:"-"` // Included in URL
	Action           *string `form:"action"`
	Quantity         *int64  `form:"quantity"`
	Timestamp        *int64  `form:"timestamp"`
	TimestampNow     *bool   `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for UsageRecordParams.
func (u *UsageRecordParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(u.TimestampNow) {
		body.Add(form.FormatKey(append(keyParts, "timestamp")), "now")
	}
}

// Usage records allow you to report customer usage and metrics to Stripe for
// metered billing of subscription prices.
//
// Related guide: [Metered Billing](https://stripe.com/docs/billing/subscriptions/metered-billing).
type UsageRecord struct {
	APIResource
	ID               string `json:"id"`
	Livemode         bool   `json:"livemode"`
	Object           string `json:"object"`
	Quantity         int64  `json:"quantity"`
	SubscriptionItem string `json:"subscription_item"`
	Timestamp        int64  `json:"timestamp"`
}
