package source

import (
	"errors"
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /sources APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new source.
// For more details see https://stripe.com/docs/api#create_source.
func New(params *stripe.SourceObjectParams) (*stripe.Source, error) {
	return getC().New(params)
}

// New POSTs a new source.
// For more details see https://stripe.com/docs/api#create_source.
func (c Client) New(params *stripe.SourceObjectParams) (*stripe.Source, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		commonParams = &params.Params
		form.AppendTo(body, params)
	}

	p := &stripe.Source{}
	err := c.B.Call("POST", "/sources", c.Key, body, commonParams, p)

	return p, err
}

// Get returns the details of a source
// For more details see https://stripe.com/docs/api#retrieve_source.
func Get(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	return getC().Get(id, params)
}

// Get returns the details of a source
// For more details see https://stripe.com/docs/api#retrieve_source.
func (c Client) Get(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		commonParams = &params.Params
		form.AppendTo(body, params)
	}

	source := &stripe.Source{}
	err := c.B.Call("GET", "/sources/"+id, c.Key, body, commonParams, source)
	return source, err
}

// Update updates a source's properties.
// For more details see	https://stripe.com/docs/api#update_source.
func Update(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		commonParams = &params.Params
		form.AppendTo(body, params)
	}

	source := &stripe.Source{}
	err := c.B.Call("POST", "/sources/"+id, c.Key, body, commonParams, source)

	return source, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}

// Detach detaches the source from its customer object.
func Detach(id string, params *stripe.SourceObjectDetachParams) (*stripe.Source, error) {
	return getC().Detach(id, params)
}

func (c Client) Detach(id string, params *stripe.SourceObjectDetachParams) (*stripe.Source, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	source := &stripe.Source{}
	var err error

	if len(params.Customer) > 0 {
		err = c.B.Call("DELETE", fmt.Sprintf("/customers/%v/sources/%v", params.Customer, id), c.Key, body, commonParams, source)
	} else {
		err = errors.New("Invalid source detach params: Customer needs to be set")
	}

	return source, err
}
