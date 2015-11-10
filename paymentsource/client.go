// Package paymentsource provides the /sources APIs
package paymentsource

import (
	"errors"
	"fmt"
	"net/url"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /sources APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new sources for a customer.
// For more details see https://stripe.com/docs/api#create_source.
func New(params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	return getC().New(params)
}

func (s Client) New(params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	body := &url.Values{}
	params.Source.AppendDetails(body, true)
	params.AppendTo(body)

	source := &stripe.PaymentSource{}
	var err error

	if len(params.Customer) > 0 {
		err = s.B.Call("POST", fmt.Sprintf("/customers/%v/sources", params.Customer), s.Key, body, &params.Params, source)
	} else {
		err = errors.New("Invalid source params: customer needs to be set")
	}

	return source, err
}

// Get returns the details of a source.
// For more details see https://stripe.com/docs/api#retrieve_source.
func Get(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	return getC().Get(id, params)
}

func (s Client) Get(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}
		params.AppendTo(body)
	}

	source := &stripe.PaymentSource{}
	var err error

	if len(params.Customer) > 0 {
		err = s.B.Call("GET", fmt.Sprintf("/customers/%v/sources/%v", params.Customer, id), s.Key, body, commonParams, source)
	} else {
		err = errors.New("Invalid source params: customer needs to be set")
	}

	return source, err
}

// Update updates a source's properties.
// For more details see https://stripe.com/docs/api#update_source.
func Update(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	return getC().Update(id, params)
}

func (s Client) Update(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	body := &url.Values{}
	params.Source.AppendDetails(body, false)
	params.AppendTo(body)

	source := &stripe.PaymentSource{}
	var err error

	if len(params.Customer) > 0 {
		err = s.B.Call("POST", fmt.Sprintf("/customers/%v/sources/%v", params.Customer, id), s.Key, body, &params.Params, source)
	} else {
		err = errors.New("Invalid source params: customer needs to be set")
	}

	return source, err
}

// Del removes a source.
// For more details see https://stripe.com/docs/api#delete_source.
func Del(id string, params *stripe.CustomerSourceParams) error {
	return getC().Del(id, params)
}

func (s Client) Del(id string, params *stripe.CustomerSourceParams) error {

	if len(params.Customer) > 0 {
		return s.B.Call("DELETE", fmt.Sprintf("/customers/%v/sources/%v", params.Customer, id), s.Key, nil, &params.Params, nil)
	} else {
		return errors.New("Invalid source params: customer needs to be set")
	}
}

// List returns a list of sources.
// For more details see https://stripe.com/docs/api#list_sources.
func List(params *stripe.SourceListParams) *Iter {
	return getC().List(params)
}

func (s Client) List(params *stripe.SourceListParams) *Iter {
	body := &url.Values{}
	var lp *stripe.ListParams

	params.AppendTo(body)
	lp = &params.ListParams

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.SourceList{}
		var err error

		if len(params.Customer) > 0 {
			err = s.B.Call("GET", fmt.Sprintf("/customers/%v/sources", params.Customer), s.Key, &b, nil, list)
		} else {
			err = errors.New("Invalid source params: customer needs to be set")
		}

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of PaymentSources.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// PaymentSource returns the most recent PaymentSource
// visited by a call to Next.
func (i *Iter) PaymentSource() *stripe.PaymentSource {
	return i.Current().(*stripe.PaymentSource)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
