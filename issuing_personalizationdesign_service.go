//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v83/form"
)

// v1IssuingPersonalizationDesignService is used to invoke /v1/issuing/personalization_designs APIs.
type v1IssuingPersonalizationDesignService struct {
	B   Backend
	Key string
}

// Creates a personalization design object.
func (c v1IssuingPersonalizationDesignService) Create(ctx context.Context, params *IssuingPersonalizationDesignCreateParams) (*IssuingPersonalizationDesign, error) {
	if params == nil {
		params = &IssuingPersonalizationDesignCreateParams{}
	}
	params.Context = ctx
	personalizationdesign := &IssuingPersonalizationDesign{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/personalization_designs", c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Retrieves a personalization design object.
func (c v1IssuingPersonalizationDesignService) Retrieve(ctx context.Context, id string, params *IssuingPersonalizationDesignRetrieveParams) (*IssuingPersonalizationDesign, error) {
	if params == nil {
		params = &IssuingPersonalizationDesignRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/personalization_designs/%s", id)
	personalizationdesign := &IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Updates a card personalization object.
func (c v1IssuingPersonalizationDesignService) Update(ctx context.Context, id string, params *IssuingPersonalizationDesignUpdateParams) (*IssuingPersonalizationDesign, error) {
	if params == nil {
		params = &IssuingPersonalizationDesignUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/personalization_designs/%s", id)
	personalizationdesign := &IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Returns a list of personalization design objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingPersonalizationDesignService) List(ctx context.Context, listParams *IssuingPersonalizationDesignListParams) *V1List[*IssuingPersonalizationDesign] {
	if listParams == nil {
		listParams = &IssuingPersonalizationDesignListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*V1Page[*IssuingPersonalizationDesign], error) {
		list := &V1Page[*IssuingPersonalizationDesign]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/personalization_designs", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
