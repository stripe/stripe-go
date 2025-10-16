//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves a requested session
type DelegatedCheckoutRequestedSessionParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Confirms a requested session
type DelegatedCheckoutRequestedSessionConfirmParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The PaymentMethod to use with the requested session.
	PaymentMethod *string `form:"payment_method"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionConfirmParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Expires a requested session
type DelegatedCheckoutRequestedSessionExpireParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionExpireParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a requested session
type DelegatedCheckoutRequestedSessionRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates a requested session
type DelegatedCheckoutRequestedSessionUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a requested session
type DelegatedCheckoutRequestedSessionCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *DelegatedCheckoutRequestedSessionCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type DelegatedCheckoutRequestedSessionFulfillmentDetails struct{}

// A requested session is a session that has been requested by a customer.
type DelegatedCheckoutRequestedSession struct {
	APIResource
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The customer for this requested session.
	Customer           string                                               `json:"customer"`
	FulfillmentDetails *DelegatedCheckoutRequestedSessionFulfillmentDetails `json:"fulfillment_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
