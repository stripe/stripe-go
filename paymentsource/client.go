// Package paymentsource provides the /sources APIs
package paymentsource

import (
	"errors"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
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
	if params == nil {
		return nil, errors.New("params should not be nil")
	}

	if params.Customer == nil {
		return nil, errors.New("Invalid source params: customer needs to be set")
	}

	path := stripe.FormatURLPath("/customers/%s/sources", stripe.StringValue(params.Customer))
	source := &stripe.PaymentSource{}
	err := s.B.Call("POST", path, s.Key, params, source)
	return source, err
}

// Get returns the details of a source.
// For more details see https://stripe.com/docs/api#retrieve_source.
func Get(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	return getC().Get(id, params)
}

func (s Client) Get(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	if params == nil {
		return nil, errors.New("params should not be nil")
	}

	if params.Customer == nil {
		return nil, errors.New("Invalid source params: customer needs to be set")
	}

	path := stripe.FormatURLPath("/customers/%s/sources/%s", stripe.StringValue(params.Customer), id)
	source := &stripe.PaymentSource{}
	err := s.B.Call("GET", path, s.Key, params, source)
	return source, err
}

// Update updates a source's properties.
// For more details see https://stripe.com/docs/api#update_source.
func Update(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	return getC().Update(id, params)
}

func (s Client) Update(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	if params == nil {
		return nil, errors.New("params should not be nil")
	}

	if params.Customer == nil {
		return nil, errors.New("Invalid source params: customer needs to be set")
	}

	path := stripe.FormatURLPath("/customers/%s/sources/%s", stripe.StringValue(params.Customer), id)
	source := &stripe.PaymentSource{}
	err := s.B.Call("POST", path, s.Key, params, source)
	return source, err
}

// Del removes a source.
// For more details see https://stripe.com/docs/api#delete_source.
func Del(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	return getC().Del(id, params)
}

func (s Client) Del(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	if params == nil {
		return nil, errors.New("params should not be nil")
	}

	if params.Customer == nil {
		return nil, errors.New("Invalid source params: customer needs to be set")
	}

	source := &stripe.PaymentSource{}
	path := stripe.FormatURLPath("/customers/%s/sources/%s", stripe.StringValue(params.Customer), id)
	err := s.B.Call("DELETE", path, s.Key, params, source)
	return source, err
}

// List returns a list of sources.
// For more details see https://stripe.com/docs/api#list_sources.
func List(params *stripe.SourceListParams) *Iter {
	return getC().List(params)
}

func (s Client) List(listParams *stripe.SourceListParams) *Iter {
	var outerErr error
	var path string

	if listParams == nil {
		outerErr = errors.New("params should not be nil")
	} else if listParams.Customer == nil {
		outerErr = errors.New("Invalid source params: customer needs to be set")
	} else {
		path = stripe.FormatURLPath("/customers/%s/sources",
			stripe.StringValue(listParams.Customer))
	}

	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.SourceList{}

		if outerErr != nil {
			return nil, list.ListMeta, outerErr
		}

		err := s.B.CallRaw("GET", path, s.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Verify verifies a bank account
// For more details see https://stripe.com/docs/guides/ach-beta
func Verify(id string, params *stripe.SourceVerifyParams) (*stripe.PaymentSource, error) {
	return getC().Verify(id, params)
}

func (s Client) Verify(id string, params *stripe.SourceVerifyParams) (*stripe.PaymentSource, error) {
	if params == nil {
		return nil, errors.New("params should not be nil")
	}

	var path string
	if params.Customer != nil {
		path = stripe.FormatURLPath("/customers/%s/sources/%s/verify",
			stripe.StringValue(params.Customer), id)
	} else if len(params.Values) > 0 {
		path = stripe.FormatURLPath("/sources/%s/verify", id)
	} else {
		return nil, errors.New("Only customer bank accounts or sources can be verified in this manner.")
	}

	source := &stripe.PaymentSource{}
	err := s.B.Call("POST", path, s.Key, params, source)
	return source, err
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
