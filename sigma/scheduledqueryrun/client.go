//
//
// File generated from our OpenAPI spec
//
//

// Package scheduledqueryrun provides the /v1/sigma/scheduled_query_runs APIs
// For more details, see: https://stripe.com/docs/api#scheduled_queries
package scheduledqueryrun

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/sigma/scheduled_query_runs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an scheduled query run.
func Get(id string, params *stripe.SigmaScheduledQueryRunParams) (*stripe.SigmaScheduledQueryRun, error) {
	return getC().Get(id, params)
}

// Retrieves the details of an scheduled query run.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.SigmaScheduledQueryRunParams) (*stripe.SigmaScheduledQueryRun, error) {
	path := stripe.FormatURLPath("/v1/sigma/scheduled_query_runs/%s", id)
	scheduledqueryrun := &stripe.SigmaScheduledQueryRun{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, scheduledqueryrun)
	return scheduledqueryrun, err
}

// Returns a list of scheduled query runs.
func List(params *stripe.SigmaScheduledQueryRunListParams) *Iter {
	return getC().List(params)
}

// Returns a list of scheduled query runs.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.SigmaScheduledQueryRunListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SigmaScheduledQueryRunList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/sigma/scheduled_query_runs", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for sigma scheduled query runs.
type Iter struct {
	*stripe.Iter
}

// SigmaScheduledQueryRun returns the sigma scheduled query run which the iterator is currently pointing to.
func (i *Iter) SigmaScheduledQueryRun() *stripe.SigmaScheduledQueryRun {
	return i.Current().(*stripe.SigmaScheduledQueryRun)
}

// SigmaScheduledQueryRunList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SigmaScheduledQueryRunList() *stripe.SigmaScheduledQueryRunList {
	return i.List().(*stripe.SigmaScheduledQueryRunList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
