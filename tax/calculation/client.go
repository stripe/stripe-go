//
//
// File generated from our OpenAPI spec
//
//

// Package calculation provides the /tax/calculations APIs
package calculation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/form"
)

// Client is used to invoke /tax/calculations APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new tax calculation.
func New(params *stripe.TaxCalculationParams) (*stripe.TaxCalculation, error) {
	return getC().New(params)
}

// New creates a new tax calculation.
func (c Client) New(params *stripe.TaxCalculationParams) (*stripe.TaxCalculation, error) {
	calculation := &stripe.TaxCalculation{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/tax/calculations",
		c.Key,
		params,
		calculation,
	)
	return calculation, err
}

// ListLineItems is the method for the `GET /v1/tax/calculations/{calculation}/line_items` API.
func ListLineItems(params *stripe.TaxCalculationListLineItemsParams) *LineItemIter {
	return getC().ListLineItems(params)
}

// ListLineItems is the method for the `GET /v1/tax/calculations/{calculation}/line_items` API.
func (c Client) ListLineItems(listParams *stripe.TaxCalculationListLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/tax/calculations/%s/line_items",
		stripe.StringValue(listParams.Calculation),
	)
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxCalculationLineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// LineItemIter is an iterator for tax calculation line items.
type LineItemIter struct {
	*stripe.Iter
}

// TaxCalculationLineItem returns the tax calculation line item which the iterator is currently pointing to.
func (i *LineItemIter) TaxCalculationLineItem() *stripe.TaxCalculationLineItem {
	return i.Current().(*stripe.TaxCalculationLineItem)
}

// TaxCalculationLineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) TaxCalculationLineItemList() *stripe.TaxCalculationLineItemList {
	return i.List().(*stripe.TaxCalculationLineItemList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
