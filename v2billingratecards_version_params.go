//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List the versions of a Rate Card object. Results are sorted in reverse chronological order (most recent first).
type V2BillingRateCardsVersionListParams struct {
	Params     `form:"*"`
	RateCardID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a specific version of a Rate Card object.
type V2BillingRateCardsVersionParams struct {
	Params     `form:"*"`
	RateCardID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a specific version of a Rate Card object.
type V2BillingRateCardsVersionRetrieveParams struct {
	Params     `form:"*"`
	RateCardID *string `form:"-" json:"-"` // Included in URL
}
