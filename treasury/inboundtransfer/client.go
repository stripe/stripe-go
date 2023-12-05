//
//
// File generated from our OpenAPI spec
//
//

// Package inboundtransfer provides the /treasury/inbound_transfers APIs
package inboundtransfer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /treasury/inbound_transfers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new treasury inbound transfer.
func New(params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().New(params)
}

// New creates a new treasury inbound transfer.
func (c Client) New(params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	var err error
	sr := stripe.NewStripeRequest(
		http.MethodPost,
		"/v1/treasury/inbound_transfers",
		c.Key,
	)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, inboundtransfer)
	return inboundtransfer, err
}

// Get returns the details of a treasury inbound transfer.
func Get(id string, params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().Get(id, params)
}

// Get returns the details of a treasury inbound transfer.
func (c Client) Get(id string, params *stripe.TreasuryInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	path := stripe.FormatURLPath("/v1/treasury/inbound_transfers/%s", id)
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, inboundtransfer)
	return inboundtransfer, err
}

// Cancel is the method for the `POST /v1/treasury/inbound_transfers/{inbound_transfer}/cancel` API.
func Cancel(id string, params *stripe.TreasuryInboundTransferCancelParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/treasury/inbound_transfers/{inbound_transfer}/cancel` API.
func (c Client) Cancel(id string, params *stripe.TreasuryInboundTransferCancelParams) (*stripe.TreasuryInboundTransfer, error) {
	path := stripe.FormatURLPath("/v1/treasury/inbound_transfers/%s/cancel", id)
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, inboundtransfer)
	return inboundtransfer, err
}

// List returns a list of treasury inbound transfers.
func List(params *stripe.TreasuryInboundTransferListParams) *Iter {
	return getC().List(params)
}

// List returns a list of treasury inbound transfers.
func (c Client) List(listParams *stripe.TreasuryInboundTransferListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryInboundTransferList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				"/v1/treasury/inbound_transfers",
				c.Key,
			)
			err := sr.SetRawForm(p, b)
			if err != nil {
				return nil, list, err
			}
			err = c.B.Call(sr, list)

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
