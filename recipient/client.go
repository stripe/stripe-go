// Package recipient provides the /recipients APIs
package recipient

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	Individual stripe.RecipientType = "individual"
	Corp       stripe.RecipientType = "corporation"
)

// Client is used to invoke /recipients APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Since API version 2017-04-06, new recipients can't be created anymore.
// For that reason, there isn't a New() method for the Recipient resource.

// Get returns the details of a recipient.
// For more details see https://stripe.com/docs/api#retrieve_recipient.
func Get(id string, params *stripe.RecipientParams) (*stripe.Recipient, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.RecipientParams) (*stripe.Recipient, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	recipient := &stripe.Recipient{}
	err := c.B.Call("GET", "/recipients/"+id, c.Key, body, commonParams, recipient)

	return recipient, err
}

// Update updates a recipient's properties.
// For more details see https://stripe.com/docs/api#update_recipient.
func Update(id string, params *stripe.RecipientParams) (*stripe.Recipient, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.RecipientParams) (*stripe.Recipient, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	recipient := &stripe.Recipient{}
	err := c.B.Call("POST", "/recipients/"+id, c.Key, body, commonParams, recipient)

	return recipient, err
}

// Del removes a recipient.
// For more details see https://stripe.com/docs/api#delete_recipient.
func Del(id string, params *stripe.RecipientParams) (*stripe.Recipient, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.RecipientParams) (*stripe.Recipient, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	recipient := &stripe.Recipient{}
	err := c.B.Call("DELETE", "/recipients/"+id, c.Key, body, commonParams, recipient)

	return recipient, err
}

// List returns a list of recipients.
// For more details see https://stripe.com/docs/api#list_recipients.
func List(params *stripe.RecipientListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.RecipientListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.RecipientList{}
		err := c.B.Call("GET", "/recipients", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Recipients.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Recipient returns the most recent Recipient
// visited by a call to Next.
func (i *Iter) Recipient() *stripe.Recipient {
	return i.Current().(*stripe.Recipient)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
