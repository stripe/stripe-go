//
//
// File generated from our OpenAPI spec
//
//

// Package setupintent provides the /setup_intents APIs
package setupintent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /setup_intents APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new setup intent.
func New(params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	return getC().New(params)
}

// New creates a new setup intent.
func (c Client) New(params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	setupintent := &stripe.SetupIntent{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/setup_intents",
		c.Key,
		params,
		setupintent,
	)
	return setupintent, err
}

// Get returns the details of a setup intent.
func Get(id string, params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	return getC().Get(id, params)
}

// Get returns the details of a setup intent.
func (c Client) Get(id string, params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	path := stripe.FormatURLPath("/v1/setup_intents/%s", id)
	setupintent := &stripe.SetupIntent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, setupintent)
	return setupintent, err
}

// Update updates a setup intent's properties.
func Update(id string, params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	return getC().Update(id, params)
}

// Update updates a setup intent's properties.
func (c Client) Update(id string, params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	path := stripe.FormatURLPath("/v1/setup_intents/%s", id)
	setupintent := &stripe.SetupIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, setupintent)
	return setupintent, err
}

// Cancel is the method for the `POST /v1/setup_intents/{intent}/cancel` API.
func Cancel(id string, params *stripe.SetupIntentCancelParams) (*stripe.SetupIntent, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/setup_intents/{intent}/cancel` API.
func (c Client) Cancel(id string, params *stripe.SetupIntentCancelParams) (*stripe.SetupIntent, error) {
	path := stripe.FormatURLPath("/v1/setup_intents/%s/cancel", id)
	setupintent := &stripe.SetupIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, setupintent)
	return setupintent, err
}

// Confirm is the method for the `POST /v1/setup_intents/{intent}/confirm` API.
func Confirm(id string, params *stripe.SetupIntentConfirmParams) (*stripe.SetupIntent, error) {
	return getC().Confirm(id, params)
}

// Confirm is the method for the `POST /v1/setup_intents/{intent}/confirm` API.
func (c Client) Confirm(id string, params *stripe.SetupIntentConfirmParams) (*stripe.SetupIntent, error) {
	path := stripe.FormatURLPath("/v1/setup_intents/%s/confirm", id)
	setupintent := &stripe.SetupIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, setupintent)
	return setupintent, err
}

// List returns a list of setup intents.
func List(params *stripe.SetupIntentListParams) *Iter {
	return getC().List(params)
}

// List returns a list of setup intents.
func (c Client) List(listParams *stripe.SetupIntentListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SetupIntentList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/setup_intents", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for setup intents.
type Iter struct {
	*stripe.Iter
}

// SetupIntent returns the setup intent which the iterator is currently pointing to.
func (i *Iter) SetupIntent() *stripe.SetupIntent {
	return i.Current().(*stripe.SetupIntent)
}

// SetupIntentList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SetupIntentList() *stripe.SetupIntentList {
	return i.List().(*stripe.SetupIntentList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
