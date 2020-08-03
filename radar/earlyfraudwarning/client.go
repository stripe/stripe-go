// Package earlyfraudwarning provides API functions related to early fraud
// warnings.
//
// For more details, see: https://stripe.com/docs/api/early_fraud_warnings?lang=go
package earlyfraudwarning

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to interact with the /radar/early_fraud_warnings API.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an early fraud warning.
func Get(id string, params *stripe.RadarEarlyFraudWarningParams) (*stripe.RadarEarlyFraudWarning, error) {
	return getC().Get(id, params)
}

// Get returns the details of an early fraud warning.
func (c Client) Get(id string, params *stripe.RadarEarlyFraudWarningParams) (*stripe.RadarEarlyFraudWarning, error) {
	path := stripe.FormatURLPath("/v1/radar/early_fraud_warnings/%s", id)
	ifr := &stripe.RadarEarlyFraudWarning{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, ifr)
	return ifr, err
}

// List returns a list of early fraud warnings.
func List(params *stripe.RadarEarlyFraudWarningListParams) *Iter {
	return getC().List(params)
}

// List returns a list of early fraud warnings.
func (c Client) List(listParams *stripe.RadarEarlyFraudWarningListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.RadarEarlyFraudWarningList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/radar/early_fraud_warnings", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for early fraud warnings.
type Iter struct {
	*stripe.Iter
}

// RadarEarlyFraudWarning returns the early fraud warning which the iterator is currently pointing to.
func (i *Iter) RadarEarlyFraudWarning() *stripe.RadarEarlyFraudWarning {
	return i.Current().(*stripe.RadarEarlyFraudWarning)
}

// RadarEarlyFraudWarningList returns the current list object which the
// iterator is currently using. List objects will change as new API calls are
// made to continue pagination.
func (i *Iter) RadarEarlyFraudWarningList() *stripe.RadarEarlyFraudWarningList {
	return i.List().(*stripe.RadarEarlyFraudWarningList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
