//
//
// File generated from our OpenAPI spec
//
//

// Package calculation provides the /v1/tax/calculations APIs
package calculation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/tax/calculations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Calculates tax based on the input and returns a Tax Calculation object.
func New(params *stripe.TaxCalculationParams) (*stripe.TaxCalculation, error) {
	return getC().New(params)
}

// Calculates tax based on the input and returns a Tax Calculation object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TaxCalculationParams) (*stripe.TaxCalculation, error) {
	calculation := &stripe.TaxCalculation{}
	err := c.B.Call(
		http.MethodPost, "/v1/tax/calculations", c.Key, params, calculation)
	return calculation, err
}

// Retrieves a Tax Calculation object, if the calculation hasn't expired.
func Get(id string, params *stripe.TaxCalculationParams) (*stripe.TaxCalculation, error) {
	return getC().Get(id, params)
}

// Retrieves a Tax Calculation object, if the calculation hasn't expired.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TaxCalculationParams) (*stripe.TaxCalculation, error) {
	path := stripe.FormatURLPath("/v1/tax/calculations/%s", id)
	calculation := &stripe.TaxCalculation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, calculation)
	return calculation, err
}

// Retrieves the line items of a tax calculation as a collection, if the calculation hasn't expired.
func ListLineItems(params *stripe.TaxCalculationListLineItemsParams) *LineItemIter {
	return getC().ListLineItems(params)
}

// Retrieves the line items of a tax calculation as a collection, if the calculation hasn't expired.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLineItems(listParams *stripe.TaxCalculationListLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath(
		"/v1/tax/calculations/%s/line_items", stripe.StringValue(
			listParams.Calculation))
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxCalculationLineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

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
