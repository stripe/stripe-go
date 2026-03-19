//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1CardService is used to invoke card related APIs.
type v1CardService struct {
	B   Backend
	Key string
}

// Create creates a new card
func (c v1CardService) Create(ctx context.Context, params *CardCreateParams) (*Card, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}

	var path string
	if (params.Account != nil && params.Customer != nil) || (params.Account == nil && params.Customer == nil) {
		return nil, fmt.Errorf("Invalid bank account params: exactly one of Account or Customer need to be set")
	} else if params.Account != nil {
		path = FormatURLPath("/v1/accounts/%s/external_accounts", StringValue(params.Account))
	} else if params.Customer != nil {
		path = FormatURLPath("/v1/customers/%s/sources", StringValue(params.Customer))
	}
	params.Context = ctx
	body := &form.Values{}

	// Note that we call this special append method instead of the standard one
	// from the form package. We should not use form's because doing so will
	// include some parameters that are undesirable here.
	if err := params.AppendToAsCardSourceOrExternalAccount(body, nil); err != nil {
		return nil, err
	}

	// Because card creation uses the custom append above, we have to
	// make an explicit call using a form and CallRaw instead of the standard
	// Call (which takes a set of parameters).
	card := &Card{}
	err := c.B.CallRaw(http.MethodPost, path, c.Key, []byte(body.Encode()), &params.Params, card)
	return card, err
}

// Get returns the details of a card.
func (c v1CardService) Retrieve(ctx context.Context, id string, params *CardRetrieveParams) (*Card, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	var path string
	if (params.Account != nil && params.Customer != nil) || (params.Account == nil && params.Customer == nil) {
		return nil, fmt.Errorf("Invalid card params: exactly one of Account or Customer need to be set")
	} else if params.Account != nil {
		path = FormatURLPath("/v1/accounts/%s/external_accounts/%s", StringValue(params.Account), id)
	} else if params.Customer != nil {
		path = FormatURLPath("/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	}
	params.Context = ctx
	card := &Card{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, card)
	return card, err
}

// Update a specified source for a given customer.
func (c v1CardService) Update(ctx context.Context, id string, params *CardUpdateParams) (*Card, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	var path string
	if (params.Account != nil && params.Customer != nil) || (params.Account == nil && params.Customer == nil) {
		return nil, fmt.Errorf("Invalid card params: exactly one of Account or Customer need to be set")
	} else if params.Account != nil {
		path = FormatURLPath("/v1/accounts/%s/external_accounts/%s", StringValue(params.Account), id)
	} else if params.Customer != nil {
		path = FormatURLPath("/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	}
	params.Context = ctx
	card := &Card{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Delete a specified source for a given customer.
func (c v1CardService) Delete(ctx context.Context, id string, params *CardDeleteParams) (*Card, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	var path string
	if (params.Account != nil && params.Customer != nil) || (params.Account == nil && params.Customer == nil) {
		return nil, fmt.Errorf("Invalid card params: exactly one of Account or Customer need to be set")
	} else if params.Account != nil {
		path = FormatURLPath("/v1/accounts/%s/external_accounts/%s", StringValue(params.Account), id)
	} else if params.Customer != nil {
		path = FormatURLPath("/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	}
	params.Context = ctx
	card := &Card{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, card)
	return card, err
}
func (c v1CardService) List(ctx context.Context, listParams *CardListParams) *V1List[*Card] {
	var path string
	var outerErr error

	// There's no cards list URL, so we use one sources or external
	// accounts. An override on CardListParam's `AppendTo` will add the
	// filter `object=card` to make sure that only cards come
	// back with the response.
	if listParams == nil {
		outerErr = fmt.Errorf("params should not be nil")
	} else if (listParams.Account != nil && listParams.Customer != nil) || (listParams.Account == nil && listParams.Customer == nil) {
		outerErr = fmt.Errorf("Invalid card params: exactly one of Account or Customer need to be set")
	} else if listParams.Account != nil {
		path = FormatURLPath("/v1/accounts/%s/external_accounts",
			StringValue(listParams.Account))
	} else if listParams.Customer != nil {
		path = FormatURLPath("/v1/customers/%s/sources",
			StringValue(listParams.Customer))
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*Card], error) {
		list := &v1Page[*Card]{}

		if outerErr != nil {
			return nil, outerErr
		}

		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
