package source

import (
	"errors"

	stripe "github.com/stripe/stripe-go"
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
	p := &stripe.Source{}
	err := c.B.Call2("POST", "/sources", c.Key, params, p)
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
	path := stripe.FormatURLPath("/sources/%s", id)
	source := &stripe.Source{}
	err := c.B.Call2("GET", path, c.Key, params, source)
	return source, err
}

// Update updates a source's properties.
// For more details see	https://stripe.com/docs/api#update_source.
func Update(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	path := stripe.FormatURLPath("/sources/%s", id)
	source := &stripe.Source{}
	err := c.B.Call2("POST", path, c.Key, params, source)
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
	if params.Customer == nil {
		return nil, errors.New("Invalid source detach params: Customer needs to be set")
	}

	path := stripe.FormatURLPath("/customers/%s/sources/%s",
		stripe.StringValue(params.Customer), id)
	source := &stripe.Source{}
	err := c.B.Call2("DELETE", path, c.Key, params, source)
	return source, err
}
