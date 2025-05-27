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

// v1PayoutService is used to invoke /v1/payouts APIs.
type v1PayoutService struct {
	B   Backend
	Key string
}

// To send funds to your own bank account, create a new payout object. Your [Stripe balance](https://docs.stripe.com/api#balance) must cover the payout amount. If it doesn't, you receive an “Insufficient Funds” error.
//
// If your API key is in test mode, money won't actually be sent, though every other action occurs as if you're in live mode.
//
// If you create a manual payout on a Stripe account that uses multiple payment source types, you need to specify the source type balance that the payout draws from. The [balance object](https://docs.stripe.com/api#balance_object) details available and pending amounts by source type.
func (c v1PayoutService) Create(ctx context.Context, params *PayoutCreateParams) (*Payout, error) {
	if params == nil {
		params = &PayoutCreateParams{}
	}
	params.Context = ctx
	payout := &Payout{}
	err := c.B.Call(http.MethodPost, "/v1/payouts", c.Key, params, payout)
	return payout, err
}

// Retrieves the details of an existing payout. Supply the unique payout ID from either a payout creation request or the payout list. Stripe returns the corresponding payout information.
func (c v1PayoutService) Retrieve(ctx context.Context, id string, params *PayoutRetrieveParams) (*Payout, error) {
	if params == nil {
		params = &PayoutRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payouts/%s", id)
	payout := &Payout{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, payout)
	return payout, err
}

// Updates the specified payout by setting the values of the parameters you pass. We don't change parameters that you don't provide. This request only accepts the metadata as arguments.
func (c v1PayoutService) Update(ctx context.Context, id string, params *PayoutUpdateParams) (*Payout, error) {
	if params == nil {
		params = &PayoutUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payouts/%s", id)
	payout := &Payout{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payout)
	return payout, err
}

// Cancels a top-up. Only pending top-ups can be canceled.
func (c v1PayoutService) Cancel(ctx context.Context, id string, params *PayoutCancelParams) (*Payout, error) {
	path := FormatURLPath("/v1/topups/%s/cancel", id)
	topup := &Payout{}
	if params == nil {
		params = &PayoutCancelParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, topup)
	return topup, err
}

// Reverses a payout by debiting the destination bank account. At this time, you can only reverse payouts for connected accounts to US bank accounts. If the payout is manual and in the pending status, use /v1/payouts/:id/cancel instead.
//
// By requesting a reversal through /v1/payouts/:id/reverse, you confirm that the authorized signatory of the selected bank account authorizes the debit on the bank account and that no other authorization is required.
func (c v1PayoutService) Reverse(ctx context.Context, id string, params *PayoutReverseParams) (*Payout, error) {
	if params == nil {
		params = &PayoutReverseParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payouts/%s/reverse", id)
	payout := &Payout{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payout)
	return payout, err
}

// Returns a list of existing payouts sent to third-party bank accounts or payouts that Stripe sent to you. The payouts return in sorted order, with the most recently created payouts appearing first.
func (c v1PayoutService) List(ctx context.Context, listParams *PayoutListParams) Seq2[*Payout, error] {
	if listParams == nil {
		listParams = &PayoutListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Payout, ListContainer, error) {
		list := &PayoutList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payouts", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
