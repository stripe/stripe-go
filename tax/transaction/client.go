//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the /tax/transactions APIs
package transaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /tax/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a tax transaction.
func Get(id string, params *stripe.TaxTransactionParams) (*stripe.TaxTransaction, error) {
	return getC().Get(id, params)
}

// Get returns the details of a tax transaction.
func (c Client) Get(id string, params *stripe.TaxTransactionParams) (*stripe.TaxTransaction, error) {
	path := stripe.FormatURLPath("/v1/tax/transactions/%s", id)
	transaction := &stripe.TaxTransaction{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, transaction)
	return transaction, err
}

// CreateFromCalculation is the method for the `POST /v1/tax/transactions/create_from_calculation` API.
func CreateFromCalculation(params *stripe.TaxTransactionCreateFromCalculationParams) (*stripe.TaxTransaction, error) {
	return getC().CreateFromCalculation(params)
}

// CreateFromCalculation is the method for the `POST /v1/tax/transactions/create_from_calculation` API.
func (c Client) CreateFromCalculation(params *stripe.TaxTransactionCreateFromCalculationParams) (*stripe.TaxTransaction, error) {
	transaction := &stripe.TaxTransaction{}
	var err error
	sr := stripe.NewStripeRequest(
		http.MethodPost,
		"/v1/tax/transactions/create_from_calculation",
		c.Key,
	)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, transaction)
	return transaction, err
}

// CreateReversal is the method for the `POST /v1/tax/transactions/create_reversal` API.
func CreateReversal(params *stripe.TaxTransactionCreateReversalParams) (*stripe.TaxTransaction, error) {
	return getC().CreateReversal(params)
}

// CreateReversal is the method for the `POST /v1/tax/transactions/create_reversal` API.
func (c Client) CreateReversal(params *stripe.TaxTransactionCreateReversalParams) (*stripe.TaxTransaction, error) {
	transaction := &stripe.TaxTransaction{}
	var err error
	sr := stripe.NewStripeRequest(
		http.MethodPost,
		"/v1/tax/transactions/create_reversal",
		c.Key,
	)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, transaction)
	return transaction, err
}

// ListLineItems is the method for the `GET /v1/tax/transactions/{transaction}/line_items` API.
func ListLineItems(params *stripe.TaxTransactionListLineItemsParams) *LineItemIter {
	return getC().ListLineItems(params)
}

// ListLineItems is the method for the `GET /v1/tax/transactions/{transaction}/line_items` API.
func (c Client) ListLineItems(listParams *stripe.TaxTransactionListLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/tax/transactions/%s/line_items",
		stripe.StringValue(listParams.Transaction),
	)
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxTransactionLineItemList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				path,
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

// LineItemIter is an iterator for tax transaction line items.
type LineItemIter struct {
	*stripe.Iter
}

// TaxTransactionLineItem returns the tax transaction line item which the iterator is currently pointing to.
func (i *LineItemIter) TaxTransactionLineItem() *stripe.TaxTransactionLineItem {
	return i.Current().(*stripe.TaxTransactionLineItem)
}

// TaxTransactionLineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) TaxTransactionLineItemList() *stripe.TaxTransactionLineItemList {
	return i.List().(*stripe.TaxTransactionLineItemList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
