//
//
// File generated from our OpenAPI spec
//
//

// Package setupattempt provides the /setup_attempts APIs
// For more details, see: https://stripe.com/docs/api/?lang=go#setup_attempts
package setupattempt

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /setup_attempts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of setup attempts.
func List(params *stripe.SetupAttemptListParams) *Iter {
	return getC().List(params)
}

// List returns a list of setup attempts.
func (c Client) List(listParams *stripe.SetupAttemptListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SetupAttemptList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				"/v1/setup_attempts",
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

// Iter is an iterator for setup attempts.
type Iter struct {
	*stripe.Iter
}

// SetupAttempt returns the setup attempt which the iterator is currently pointing to.
func (i *Iter) SetupAttempt() *stripe.SetupAttempt {
	return i.Current().(*stripe.SetupAttempt)
}

// SetupAttemptList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SetupAttemptList() *stripe.SetupAttemptList {
	return i.List().(*stripe.SetupAttemptList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
