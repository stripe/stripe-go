//
//
// File generated from our OpenAPI spec
//
//

// Package applicationfee provides the /v1/application_fees APIs
package applicationfee

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/application_fees APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an application fee that your account has collected. The same information is returned when refunding the application fee.
func Get(id string, params *stripe.ApplicationFeeParams) (*stripe.ApplicationFee, error) {
	return getC().Get(id, params)
}

// Retrieves the details of an application fee that your account has collected. The same information is returned when refunding the application fee.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ApplicationFeeParams) (*stripe.ApplicationFee, error) {
	path := stripe.FormatURLPath("/v1/application_fees/%s", id)
	applicationfee := &stripe.ApplicationFee{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, applicationfee)
	return applicationfee, err
}

// Returns a list of application fees you've previously collected. The application fees are returned in sorted order, with the most recent fees appearing first.
func List(params *stripe.ApplicationFeeListParams) *Iter {
	return getC().List(params)
}

// Returns a list of application fees you've previously collected. The application fees are returned in sorted order, with the most recent fees appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ApplicationFeeListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ApplicationFeeList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/application_fees", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for application fees.
type Iter struct {
	*stripe.Iter
}

// ApplicationFee returns the application fee which the iterator is currently pointing to.
func (i *Iter) ApplicationFee() *stripe.ApplicationFee {
	return i.Current().(*stripe.ApplicationFee)
}

// ApplicationFeeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ApplicationFeeList() *stripe.ApplicationFeeList {
	return i.List().(*stripe.ApplicationFeeList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
