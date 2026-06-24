//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves a list of alert history entries for a health alert.
type V2CoreHealthAlertsHistoryListParams struct {
	Params `form:"*"`
	// The ID of the health alert to list history for.
	ID *string `form:"-" json:"-"` // Included in URL
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}
