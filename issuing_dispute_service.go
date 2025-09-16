//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1IssuingDisputeService is used to invoke /v1/issuing/disputes APIs.
type v1IssuingDisputeService struct {
	B   Backend
	Key string
}

// Creates an Issuing Dispute object. Individual pieces of evidence within the evidence object are optional at this point. Stripe only validates that required evidence is present during submission. Refer to [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence) for more details about evidence requirements.
func (c v1IssuingDisputeService) Create(ctx context.Context, params *IssuingDisputeCreateParams) (*IssuingDispute, error) {
	if params == nil {
		params = &IssuingDisputeCreateParams{}
	}
	params.Context = ctx
	dispute := &IssuingDispute{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/disputes", c.Key, params, dispute)
	return dispute, err
}

// Retrieves an Issuing Dispute object.
func (c v1IssuingDisputeService) Retrieve(ctx context.Context, id string, params *IssuingDisputeRetrieveParams) (*IssuingDispute, error) {
	if params == nil {
		params = &IssuingDisputeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/disputes/%s", id)
	dispute := &IssuingDispute{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, dispute)
	return dispute, err
}

// Updates the specified Issuing Dispute object by setting the values of the parameters passed. Any parameters not provided will be left unchanged. Properties on the evidence object can be unset by passing in an empty string.
func (c v1IssuingDisputeService) Update(ctx context.Context, id string, params *IssuingDisputeUpdateParams) (*IssuingDispute, error) {
	if params == nil {
		params = &IssuingDisputeUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/disputes/%s", id)
	dispute := &IssuingDispute{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, dispute)
	return dispute, err
}

// Submits an Issuing Dispute to the card network. Stripe validates that all evidence fields required for the dispute's reason are present. For more details, see [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence).
func (c v1IssuingDisputeService) Submit(ctx context.Context, id string, params *IssuingDisputeSubmitParams) (*IssuingDispute, error) {
	if params == nil {
		params = &IssuingDisputeSubmitParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/disputes/%s/submit", id)
	dispute := &IssuingDispute{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, dispute)
	return dispute, err
}

// Returns a list of Issuing Dispute objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingDisputeService) List(ctx context.Context, listParams *IssuingDisputeListParams) Seq2[*IssuingDispute, error] {
	if listParams == nil {
		listParams = &IssuingDisputeListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*IssuingDispute, ListContainer, error) {
		list := &IssuingDisputeList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/disputes", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
