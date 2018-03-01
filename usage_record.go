package stripe

const (
	// UsageRecordParamsActionIncrement indicates that if two usage records conflict
	// (i.E. are reported a the same timestamp), their Quantity will be summed
	UsageRecordParamsActionIncrement string = "increment"

	// UsageRecordParamsActionSet indicates that if two usage records conflict
	// (i.E. are reported a the same timestamp), the Quantity of the most recent
	// usage record will overwrite any other quantity.
	UsageRecordParamsActionSet string = "set"
)

// UsageRecord represents a usage record.
// See https://stripe.com/docs/api#usage_records
type UsageRecord struct {
	ID               string `json:"id"`
	Live             bool   `json:"livemode"`
	Quantity         uint64 `json:"quantity"`
	SubscriptionItem string `json:"subscription_item"`
	Timestamp        uint64 `json:"timestamp"`
}

// UsageRecordParams create a usage record for a specified subscription item
// and date, and fills it with a quantity.
type UsageRecordParams struct {
	Params           `form:"*"`
	Action           string `form:"action"`
	Quantity         uint64 `form:"quantity"`
	QuantityZero     bool   `form:"quantity,zero"`
	SubscriptionItem string `form:"-"` // passed in the URL
	Timestamp        uint64 `form:"timestamp"`
}
