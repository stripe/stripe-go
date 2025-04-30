//
//
// File generated from our OpenAPI spec
//
//

// Package inboundtransfer provides the /v1/treasury/inbound_transfers APIs
package inboundtransfer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/treasury/inbound_transfers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an InboundTransfer.
func New(params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().New(params)
}

// Creates an InboundTransfer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	err := c.B.Call(
		http.MethodPost, "/v1/treasury/inbound_transfers", c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Retrieves the details of an existing InboundTransfer.
func Get(id string, params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().Get(id, params)
}

// Retrieves the details of an existing InboundTransfer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	path := stripe.FormatURLPath("/v1/treasury/inbound_transfers/%s", id)
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Cancels an InboundTransfer.
func Cancel(id string, params *stripe.TreasuryInboundTransferCancelParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().Cancel(id, params)
}

// Cancels an InboundTransfer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.TreasuryInboundTransferCancelParams) (*stripe.TreasuryInboundTransfer, error) {
	path := stripe.FormatURLPath("/v1/treasury/inbound_transfers/%s/cancel", id)
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Returns a list of InboundTransfers sent from the specified FinancialAccount.
func List(params *stripe.TreasuryInboundTransferListParams) *Iter {
	return getC().List(params)
}

// Returns a list of InboundTransfers sent from the specified FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TreasuryInboundTransferListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryInboundTransferList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/inbound_transfers", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for treasury inbound transfers.
type Iter struct {
	*stripe.Iter
}

// TreasuryInboundTransfer returns the treasury inbound transfer which the iterator is currently pointing to.
func (i *Iter) TreasuryInboundTransfer() *stripe.TreasuryInboundTransfer {
	return i.Current().(*stripe.TreasuryInboundTransfer)
}

// TreasuryInboundTransferList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryInboundTransferList() *stripe.TreasuryInboundTransferList {
	return i.List().(*stripe.TreasuryInboundTransferList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
