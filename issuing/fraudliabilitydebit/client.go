//
//
// File generated from our OpenAPI spec
//
//

// Package fraudliabilitydebit provides the /issuing/fraud_liability_debits APIs
package fraudliabilitydebit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/form"
)

// Client is used to invoke /issuing/fraud_liability_debits APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an Issuing FraudLiabilityDebit object.
func Get(id string, params *stripe.IssuingFraudLiabilityDebitParams) (*stripe.IssuingFraudLiabilityDebit, error) {
	return getC().Get(id, params)
}

// Retrieves an Issuing FraudLiabilityDebit object.
func (c Client) Get(id string, params *stripe.IssuingFraudLiabilityDebitParams) (*stripe.IssuingFraudLiabilityDebit, error) {
	path := stripe.FormatURLPath("/v1/issuing/fraud_liability_debits/%s", id)
	fraudliabilitydebit := &stripe.IssuingFraudLiabilityDebit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, fraudliabilitydebit)
	return fraudliabilitydebit, err
}

// Returns a list of Issuing FraudLiabilityDebit objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingFraudLiabilityDebitListParams) *Iter {
	return getC().List(params)
}

// Returns a list of Issuing FraudLiabilityDebit objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c Client) List(listParams *stripe.IssuingFraudLiabilityDebitListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingFraudLiabilityDebitList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/fraud_liability_debits", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing fraud liability debits.
type Iter struct {
	*stripe.Iter
}

// IssuingFraudLiabilityDebit returns the issuing fraud liability debit which the iterator is currently pointing to.
func (i *Iter) IssuingFraudLiabilityDebit() *stripe.IssuingFraudLiabilityDebit {
	return i.Current().(*stripe.IssuingFraudLiabilityDebit)
}

// IssuingFraudLiabilityDebitList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingFraudLiabilityDebitList() *stripe.IssuingFraudLiabilityDebitList {
	return i.List().(*stripe.IssuingFraudLiabilityDebitList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
