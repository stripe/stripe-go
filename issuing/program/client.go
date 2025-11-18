//
//
// File generated from our OpenAPI spec
//
//

// Package program provides the /v1/issuing/programs APIs
package program

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/form"
)

// Client is used to invoke /v1/issuing/programs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Program object.
func New(params *stripe.IssuingProgramParams) (*stripe.IssuingProgram, error) {
	return getC().New(params)
}

// Create a Program object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.IssuingProgramParams) (*stripe.IssuingProgram, error) {
	program := &stripe.IssuingProgram{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/programs", c.Key, params, program)
	return program, err
}

// Retrieves the program specified by the given id.
func Get(id string, params *stripe.IssuingProgramParams) (*stripe.IssuingProgram, error) {
	return getC().Get(id, params)
}

// Retrieves the program specified by the given id.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.IssuingProgramParams) (*stripe.IssuingProgram, error) {
	path := stripe.FormatURLPath("/v1/issuing/programs/%s", id)
	program := &stripe.IssuingProgram{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, program)
	return program, err
}

// Updates a Program object.
func Update(id string, params *stripe.IssuingProgramParams) (*stripe.IssuingProgram, error) {
	return getC().Update(id, params)
}

// Updates a Program object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingProgramParams) (*stripe.IssuingProgram, error) {
	path := stripe.FormatURLPath("/v1/issuing/programs/%s", id)
	program := &stripe.IssuingProgram{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, program)
	return program, err
}

// List all of the programs the given Issuing user has access to.
func List(params *stripe.IssuingProgramListParams) *Iter {
	return getC().List(params)
}

// List all of the programs the given Issuing user has access to.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IssuingProgramListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingProgramList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/programs", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing programs.
type Iter struct {
	*stripe.Iter
}

// IssuingProgram returns the issuing program which the iterator is currently pointing to.
func (i *Iter) IssuingProgram() *stripe.IssuingProgram {
	return i.Current().(*stripe.IssuingProgram)
}

// IssuingProgramList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingProgramList() *stripe.IssuingProgramList {
	return i.List().(*stripe.IssuingProgramList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
