// Package applepaydomain provides the /apple_pay/domains APIs
package applepaydomain

import (
	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /apple_pay/domains and ApplePayDomain-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new ApplePayDomains.
func New(params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	body := &stripe.RequestValues{}
	body.Add("domain_name", params.DomainName)

	params.AppendTo(body)

	domain := &stripe.ApplePayDomain{}
	err := c.B.Call("POST", "/apple_pay/domains", c.Key, body, &params.Params, domain)
	return domain, err
}

// Get returns the details of an ApplePayDomain.
func Get(id string, params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ApplePayDomainParams) (*stripe.ApplePayDomain, error) {
	var body *stripe.RequestValues
	var commonParams *stripe.Params

	if params != nil {
		body = &stripe.RequestValues{}
		commonParams = &params.Params
		params.AppendTo(body)
	}

	domain := &stripe.ApplePayDomain{}
	err := c.B.Call("GET", "/apple_pay/domains/"+id, c.Key, body, commonParams, domain)

	return domain, err
}

// Del removes an ApplePayDomain.
func Del(id string) (*stripe.ApplePayDomain, error) {
	return getC().Del(id)
}

func (c Client) Del(id string) (*stripe.ApplePayDomain, error) {
	domain := &stripe.ApplePayDomain{}
	err := c.B.Call("DELETE", "/apple_pay/domains/"+id, c.Key, nil, nil, domain)

	return domain, err
}

// List lists available ApplePayDomains.
func List(params *stripe.ApplePayDomainListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.ApplePayDomainListParams) *Iter {
	var body *stripe.RequestValues
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &stripe.RequestValues{}

		params.AppendTo(body)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *stripe.RequestValues) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.ApplePayDomainList{}
		err := c.B.Call("GET", "/apple_pay/domains", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of ApplePayDomain.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// ApplePayDomain returns the most recent ApplePayDomain
// visited by a call to Next.
func (i *Iter) ApplePayDomain() *stripe.ApplePayDomain {
	return i.Current().(*stripe.ApplePayDomain)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
