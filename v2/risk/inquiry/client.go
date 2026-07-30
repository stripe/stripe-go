//
//
// File generated from our OpenAPI spec
//
//

// Package inquiry provides the inquiry related APIs
package inquiry

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke inquiry related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a risk inquiry by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2RiskInquiryParams) (*stripe.V2RiskInquiry, error) {
	path := stripe.FormatURLPath("/v2/risk/inquiries/%s", id)
	inquiry := &stripe.V2RiskInquiry{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, inquiry)
	return inquiry, err
}

// Submits a response to a risk inquiry.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2RiskInquiryParams) (*stripe.V2RiskInquiry, error) {
	path := stripe.FormatURLPath("/v2/risk/inquiries/%s", id)
	inquiry := &stripe.V2RiskInquiry{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inquiry)
	return inquiry, err
}

// Lists risk inquiries for a connected account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2RiskInquiryListParams) stripe.Seq2[*stripe.V2RiskInquiry, error] {
	if listParams == nil {
		listParams = &stripe.V2RiskInquiryListParams{}
	}
	return stripe.NewV2List("/v2/risk/inquiries", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2RiskInquiry], error) {
		page := &stripe.V2Page[*stripe.V2RiskInquiry]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
