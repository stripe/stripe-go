// Package recipient provides the /recipients APIs
package recipient

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
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
	path := stripe.FormatURLPath("/recipients/%s", id)
	recipient := &stripe.Recipient{}
	err := c.B.Call2("GET", path, c.Key, params, recipient)
	return recipient, err
}

// Update updates a recipient's properties.
// For more details see https://stripe.com/docs/api#update_recipient.
func Update(id string, params *stripe.RecipientParams) (*stripe.Recipient, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.RecipientParams) (*stripe.Recipient, error) {
	path := stripe.FormatURLPath("/recipients/%s", id)
	recipient := &stripe.Recipient{}
	err := c.B.Call2("POST", path, c.Key, params, recipient)
	return recipient, err
}

// Del removes a recipient.
// For more details see https://stripe.com/docs/api#delete_recipient.
func Del(id string, params *stripe.RecipientParams) (*stripe.Recipient, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.RecipientParams) (*stripe.Recipient, error) {
	path := stripe.FormatURLPath("/recipients/%s", id)
	recipient := &stripe.Recipient{}
	err := c.B.Call2("DELETE", path, c.Key, params, recipient)
	return recipient, err
}

// List returns a list of recipients.
// For more details see https://stripe.com/docs/api#list_recipients.
func List(params *stripe.RecipientListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.RecipientListParams) *Iter {
	return &Iter{stripe.GetIter2(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.RecipientList{}
		err := c.B.CallRaw("GET", "/recipients", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
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
