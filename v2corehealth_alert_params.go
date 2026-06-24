//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Retrieves a list of health alerts.
type V2CoreHealthAlertListParams struct {
	Params `form:"*"`
	// Filter for alerts created at the specified timestamp.
	Created *time.Time `form:"created" json:"created,omitempty"`
	// Filter for alerts created after the specified timestamp.
	CreatedGt *time.Time `form:"created_gt" json:"created_gt,omitempty"`
	// Filter for alerts created on or after the specified timestamp.
	CreatedGTE *time.Time `form:"created_gte" json:"created_gte,omitempty"`
	// Filter for alerts created before the specified timestamp.
	CreatedLT *time.Time `form:"created_lt" json:"created_lt,omitempty"`
	// Filter for alerts created on or before the specified timestamp.
	CreatedLte *time.Time `form:"created_lte" json:"created_lte,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by alert severity.
	Severity *string `form:"severity" json:"severity,omitempty"`
	// Filter by alert status.
	Status *string `form:"status" json:"status,omitempty"`
	// Filter by alert types.
	Types []*string `form:"types" json:"types,omitempty"`
}

// Retrieves a health alert by ID.
type V2CoreHealthAlertParams struct {
	Params `form:"*"`
}

// Retrieves a health alert by ID.
type V2CoreHealthAlertRetrieveParams struct {
	Params `form:"*"`
}
