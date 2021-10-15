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

// UsageRecordParams create a usage record for a specified subscription item
// and date, and fills it with a quantity.
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

// UsageRecord represents a usage record.
// See https://stripe.com/docs/api#usage_records
type UsageRecord struct {
	APIResource
	ID               string `json:"id"`
	Livemode         bool   `json:"livemode"`
	Object           string `json:"object"`
	Quantity         int64  `json:"quantity"`
	SubscriptionItem string `json:"subscription_item"`
	Timestamp        int64  `json:"timestamp"`
}
