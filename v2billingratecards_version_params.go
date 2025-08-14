//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List the versions of a RateCard object. Results are sorted in reverse chronological order (most recent first).
type V2BillingRateCardsVersionListParams struct {
	Params `form:"*"`
	// The ID of the RateCard to retrieve versions for.
	RateCardID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a specific version of a RateCard object.
type V2BillingRateCardsVersionParams struct {
	Params `form:"*"`
	// The ID of the RateCard object.
	RateCardID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a specific version of a RateCard object.
type V2BillingRateCardsVersionRetrieveParams struct {
	Params `form:"*"`
	// The ID of the RateCard object.
	RateCardID *string `form:"-" json:"-"` // Included in URL
}
