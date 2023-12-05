//
//
// File generated from our OpenAPI spec
//
//

// Package receivedcredit provides the /treasury/received_credits APIs
package receivedcredit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /treasury/received_credits APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a treasury received credit.
func Get(id string, params *stripe.TreasuryReceivedCreditParams) (*stripe.TreasuryReceivedCredit, error) {
	return getC().Get(id, params)
}

// Get returns the details of a treasury received credit.
func (c Client) Get(id string, params *stripe.TreasuryReceivedCreditParams) (*stripe.TreasuryReceivedCredit, error) {
	path := stripe.FormatURLPath("/v1/treasury/received_credits/%s", id)
	receivedcredit := &stripe.TreasuryReceivedCredit{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, receivedcredit)
	return receivedcredit, err
}

// List returns a list of treasury received credits.
func List(params *stripe.TreasuryReceivedCreditListParams) *Iter {
	return getC().List(params)
}

// List returns a list of treasury received credits.
func (c Client) List(listParams *stripe.TreasuryReceivedCreditListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryReceivedCreditList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				"/v1/treasury/received_credits",
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

// Iter is an iterator for treasury received credits.
type Iter struct {
	*stripe.Iter
}

// TreasuryReceivedCredit returns the treasury received credit which the iterator is currently pointing to.
func (i *Iter) TreasuryReceivedCredit() *stripe.TreasuryReceivedCredit {
	return i.Current().(*stripe.TreasuryReceivedCredit)
}

// TreasuryReceivedCreditList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryReceivedCreditList() *stripe.TreasuryReceivedCreditList {
	return i.List().(*stripe.TreasuryReceivedCreditList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
