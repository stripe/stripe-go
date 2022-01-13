//
//
// File generated from our OpenAPI spec
//
//

package stripe

// For the specified subscription item, returns a list of summary objects. Each object in the list provides usage information that's been summarized from multiple usage records and over a subscription billing period (e.g., 15 usage records in the month of September).
//
// The list is sorted in reverse-chronological order (newest first). The first list item represents the most current usage period that hasn't ended yet. Since new usage records can still be added, the returned summary information for the subscription item's ID should be seen as unstable until the subscription billing period ends.
type UsageRecordSummaryListParams struct {
	ListParams       `form:"*"`
	SubscriptionItem *string `form:"-"` // Included in URL
}
type UsageRecordSummary struct {
	ID               string  `json:"id"`
	Invoice          string  `json:"invoice"`
	Livemode         bool    `json:"livemode"`
	Object           string  `json:"object"`
	Period           *Period `json:"period"`
	SubscriptionItem string  `json:"subscription_item"`
	TotalUsage       int64   `json:"total_usage"`
}

// UsageRecordSummaryList is a list of UsageRecordSummaries as retrieved from a list endpoint.
type UsageRecordSummaryList struct {
	APIResource
	ListMeta
	Data []*UsageRecordSummary `json:"data"`
}
