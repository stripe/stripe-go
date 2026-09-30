//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of Stripe-defined reports that the caller can create a `ReportRun` for.
type V2DataReportListParams struct {
	Params `form:"*"`
	// Any optional includes (see https://docs.stripe.com/api-includable-response-values).
	Include []*string `form:"include" json:"include,omitempty"`
	// The maximum number of results per page. Defaults to 10. Maximum is 100.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// If supplied, only return reports with this exact, case-sensitive name.
	Name *string `form:"name" json:"name,omitempty"`
}

// Retrieves metadata about a specific `Report`, including its name, description, and
// the parameters it accepts. It's useful for understanding the capabilities and
// requirements of a particular `Report` before requesting a `ReportRun`.
type V2DataReportParams struct {
	Params `form:"*"`
	// Any optional includes (see https://docs.stripe.com/api-includable-response-values).
	Include []*string `form:"include" json:"include,omitempty"`
}

// Retrieves metadata about a specific `Report`, including its name, description, and
// the parameters it accepts. It's useful for understanding the capabilities and
// requirements of a particular `Report` before requesting a `ReportRun`.
type V2DataReportRetrieveParams struct {
	Params `form:"*"`
	// Any optional includes (see https://docs.stripe.com/api-includable-response-values).
	Include []*string `form:"include" json:"include,omitempty"`
}
