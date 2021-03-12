//
//
// File generated from our OpenAPI spec
//
//

package stripe

// UsageRecordSummaryListParams is the set of parameters that can be used when listing charges.
type UsageRecordSummaryListParams struct {
	ListParams       `form:"*"`
	SubscriptionItem *string `form:"-"` // Included in URL
}

// UsageRecordSummary represents a usage record summary.
// See https://stripe.com/docs/api#usage_records
type UsageRecordSummary struct {
	ID               string  `json:"id"`
	Invoice          string  `json:"invoice"`
	Livemode         bool    `json:"livemode"`
	Object           string  `json:"object"`
	Period           *Period `json:"period"`
	SubscriptionItem string  `json:"subscription_item"`
	TotalUsage       int64   `json:"total_usage"`
}

// UsageRecordSummaryList is a list of usage record summaries as retrieved from a list endpoint.
type UsageRecordSummaryList struct {
	APIResource
	ListMeta
	Data []*UsageRecordSummary `json:"data"`
}
