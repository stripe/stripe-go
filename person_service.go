//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1PersonService is used to invoke /v1/accounts/{account}/persons APIs.
type v1PersonService struct {
	B   Backend
	Key string
}

// Creates a new person.
func (c v1PersonService) Create(ctx context.Context, params *PersonCreateParams) (*Person, error) {
	if params == nil {
		params = &PersonCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/accounts/%s/persons", StringValue(params.Account))
	person := &Person{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, person)
	return person, err
}

// Retrieves an existing person.
func (c v1PersonService) Retrieve(ctx context.Context, id string, params *PersonRetrieveParams) (*Person, error) {
	if params == nil {
		params = &PersonRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/persons/%s", StringValue(params.Account), id)
	person := &Person{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, person)
	return person, err
}

// Updates an existing person.
func (c v1PersonService) Update(ctx context.Context, id string, params *PersonUpdateParams) (*Person, error) {
	if params == nil {
		params = &PersonUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/persons/%s", StringValue(params.Account), id)
	person := &Person{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, person)
	return person, err
}

// Deletes an existing person's relationship to the account's legal entity. Any person with a relationship for an account can be deleted through the API, except if the person is the account_opener. If your integration is using the executive parameter, you cannot delete the only verified executive on file.
func (c v1PersonService) Delete(ctx context.Context, id string, params *PersonDeleteParams) (*Person, error) {
	if params == nil {
		params = &PersonDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/persons/%s", StringValue(params.Account), id)
	person := &Person{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, person)
	return person, err
}

// Returns a list of people associated with the account's legal entity. The people are returned sorted by creation date, with the most recent people appearing first.
func (c v1PersonService) List(ctx context.Context, listParams *PersonListParams) Seq2[*Person, error] {
	if listParams == nil {
		listParams = &PersonListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/persons", StringValue(listParams.Account))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Person, ListContainer, error) {
		list := &PersonList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
