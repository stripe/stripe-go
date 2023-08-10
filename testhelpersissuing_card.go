//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Updates the shipping status of the specified Issuing Card object to delivered.
type TestHelpersIssuingCardDeliverCardParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// Updates the shipping status of the specified Issuing Card object to shipped.
type TestHelpersIssuingCardShipCardParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// Updates the shipping status of the specified Issuing Card object to returned.
type TestHelpersIssuingCardReturnCardParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// Updates the shipping status of the specified Issuing Card object to failure.
type TestHelpersIssuingCardFailCardParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}
