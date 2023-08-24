//
//
// File generated from our OpenAPI spec
//
//

// Package person provides the /accounts/{account}/persons APIs
package person

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/form"
)

// Client is used to invoke /accounts/{account}/persons APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new person.
func New(params *stripe.PersonParams) (*stripe.Person, error) {
	return getC().New(params)
}

// New creates a new person.
func (c Client) New(params *stripe.PersonParams) (*stripe.Person, error) {
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/persons",
		stripe.StringValue(params.Account),
	)
	person := &stripe.Person{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, person)
	return person, err
}

// Get returns the details of a person.
func Get(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	return getC().Get(id, params)
}

// Get returns the details of a person.
func (c Client) Get(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	if params == nil {
		return nil, fmt.Errorf(
			"params cannot be nil, and params.Account must be set",
		)
	}
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/persons/%s",
		stripe.StringValue(params.Account),
		id,
	)
	person := &stripe.Person{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, person)
	return person, err
}

// Update updates a person's properties.
func Update(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	return getC().Update(id, params)
}

// Update updates a person's properties.
func (c Client) Update(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/persons/%s",
		stripe.StringValue(params.Account),
		id,
	)
	person := &stripe.Person{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, person)
	return person, err
}

// Del removes a person.
func Del(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	return getC().Del(id, params)
}

// Del removes a person.
func (c Client) Del(id string, params *stripe.PersonParams) (*stripe.Person, error) {
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/persons/%s",
		stripe.StringValue(params.Account),
		id,
	)
	person := &stripe.Person{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, person)
	return person, err
}

// List returns a list of persons.
func List(params *stripe.PersonListParams) *Iter {
	return getC().List(params)
}

// List returns a list of persons.
func (c Client) List(listParams *stripe.PersonListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/persons",
		stripe.StringValue(listParams.Account),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PersonList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for persons.
type Iter struct {
	*stripe.Iter
}

// Person returns the person which the iterator is currently pointing to.
func (i *Iter) Person() *stripe.Person {
	return i.Current().(*stripe.Person)
}

// PersonList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PersonList() *stripe.PersonList {
	return i.List().(*stripe.PersonList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
