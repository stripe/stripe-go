//
//
// File generated from our OpenAPI spec
//
//

// Package personalizationdesign provides the /issuing/personalization_designs APIs
package personalizationdesign

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /issuing/personalization_designs APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new issuing personalization design.
func New(params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().New(params)
}

// New creates a new issuing personalization design.
func (c Client) New(params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/issuing/personalization_designs",
		c.Key,
		params,
		personalizationdesign,
	)
	return personalizationdesign, err
}

// Get returns the details of an issuing personalization design.
func Get(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().Get(id, params)
}

// Get returns the details of an issuing personalization design.
func (c Client) Get(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	path := stripe.FormatURLPath("/v1/issuing/personalization_designs/%s", id)
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Update updates an issuing personalization design's properties.
func Update(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	return getC().Update(id, params)
}

// Update updates an issuing personalization design's properties.
func (c Client) Update(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	path := stripe.FormatURLPath("/v1/issuing/personalization_designs/%s", id)
	personalizationdesign := &stripe.IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// List returns a list of issuing personalization designs.
func List(params *stripe.IssuingPersonalizationDesignListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuing personalization designs.
func (c Client) List(listParams *stripe.IssuingPersonalizationDesignListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingPersonalizationDesignList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/personalization_designs", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing personalization designs.
type Iter struct {
	*stripe.Iter
}

// IssuingPersonalizationDesign returns the issuing personalization design which the iterator is currently pointing to.
func (i *Iter) IssuingPersonalizationDesign() *stripe.IssuingPersonalizationDesign {
	return i.Current().(*stripe.IssuingPersonalizationDesign)
}

// IssuingPersonalizationDesignList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingPersonalizationDesignList() *stripe.IssuingPersonalizationDesignList {
	return i.List().(*stripe.IssuingPersonalizationDesignList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
