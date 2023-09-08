//
//
// File generated from our OpenAPI spec
//
//

// Package carddesign provides the /issuing/card_designs APIs
package carddesign

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/form"
)

// Client is used to invoke /issuing/card_designs APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new issuing card design.
func New(params *stripe.IssuingCardDesignParams) (*stripe.IssuingCardDesign, error) {
	return getC().New(params)
}

// New creates a new issuing card design.
func (c Client) New(params *stripe.IssuingCardDesignParams) (*stripe.IssuingCardDesign, error) {
	carddesign := &stripe.IssuingCardDesign{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/issuing/card_designs",
		c.Key,
		params,
		carddesign,
	)
	return carddesign, err
}

// Get returns the details of an issuing card design.
func Get(id string, params *stripe.IssuingCardDesignParams) (*stripe.IssuingCardDesign, error) {
	return getC().Get(id, params)
}

// Get returns the details of an issuing card design.
func (c Client) Get(id string, params *stripe.IssuingCardDesignParams) (*stripe.IssuingCardDesign, error) {
	path := stripe.FormatURLPath("/v1/issuing/card_designs/%s", id)
	carddesign := &stripe.IssuingCardDesign{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, carddesign)
	return carddesign, err
}

// Update updates an issuing card design's properties.
func Update(id string, params *stripe.IssuingCardDesignParams) (*stripe.IssuingCardDesign, error) {
	return getC().Update(id, params)
}

// Update updates an issuing card design's properties.
func (c Client) Update(id string, params *stripe.IssuingCardDesignParams) (*stripe.IssuingCardDesign, error) {
	path := stripe.FormatURLPath("/v1/issuing/card_designs/%s", id)
	carddesign := &stripe.IssuingCardDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, carddesign)
	return carddesign, err
}

// List returns a list of issuing card designs.
func List(params *stripe.IssuingCardDesignListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuing card designs.
func (c Client) List(listParams *stripe.IssuingCardDesignListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingCardDesignList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/card_designs", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing card designs.
type Iter struct {
	*stripe.Iter
}

// IssuingCardDesign returns the issuing card design which the iterator is currently pointing to.
func (i *Iter) IssuingCardDesign() *stripe.IssuingCardDesign {
	return i.Current().(*stripe.IssuingCardDesign)
}

// IssuingCardDesignList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingCardDesignList() *stripe.IssuingCardDesignList {
	return i.List().(*stripe.IssuingCardDesignList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
