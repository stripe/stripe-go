//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1PaymentSourceService is used to invoke /v1/customers/{customer}/sources APIs.
type v1PaymentSourceService struct {
	B   Backend
	Key string
}

// When you create a new credit card, you must specify a customer or recipient on which to create it.
//
// If the card's owner has no default card, then the new card will become the default.
// However, if the owner already has a default, then it will not change.
// To change the default, you should [update the customer](https://docs.stripe.com/docs/api#update_customer) to have a new default_source.
func (c v1PaymentSourceService) Create(ctx context.Context, params *PaymentSourceCreateParams) (*PaymentSource, error) {
	if params == nil {
		params = &PaymentSourceCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/sources", StringValue(params.Customer))
	paymentsource := &PaymentSource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// Retrieve a specified source for a given customer.
func (c v1PaymentSourceService) Retrieve(ctx context.Context, id string, params *PaymentSourceRetrieveParams) (*PaymentSource, error) {
	if params == nil {
		params = &PaymentSourceRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	paymentsource := &PaymentSource{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// Update a specified source for a given customer.
func (c v1PaymentSourceService) Update(ctx context.Context, id string, params *PaymentSourceUpdateParams) (*PaymentSource, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	if params.Customer == nil {
		return nil, fmt.Errorf("Invalid source params: customer needs to be set")
	}
	if params == nil {
		params = &PaymentSourceUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	paymentsource := &PaymentSource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// Delete a specified source for a given customer.
func (c v1PaymentSourceService) Delete(ctx context.Context, id string, params *PaymentSourceDeleteParams) (*PaymentSource, error) {
	if params == nil {
		params = &PaymentSourceDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	paymentsource := &PaymentSource{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// Verify verifies a source which is used for bank accounts.
// Verify a specified bank account for a given customer.
func (c v1PaymentSourceService) Verify(ctx context.Context, id string, params *PaymentSourceVerifyParams) (*PaymentSource, error) {
	if params == nil {
		params = &PaymentSourceVerifyParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/sources/%s/verify", StringValue(params.Customer), id)
	paymentsource := &PaymentSource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// List sources for a specified customer.
func (c v1PaymentSourceService) List(ctx context.Context, listParams *PaymentSourceListParams) Seq2[*PaymentSource, error] {
	if listParams == nil {
		listParams = &PaymentSourceListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/sources", StringValue(listParams.Customer))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*PaymentSource, ListContainer, error) {
		list := &PaymentSourceList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
