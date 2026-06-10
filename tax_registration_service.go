//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stripe/stripe-go/v86/form"
)

// v1TaxRegistrationService is used to invoke /v1/tax/registrations APIs.
type v1TaxRegistrationService struct {
	B   Backend
	Key string
}

// Creates a new Tax Registration object.
func (c v1TaxRegistrationService) Create(ctx context.Context, params *TaxRegistrationCreateParams) (*TaxRegistration, error) {
	if params == nil {
		params = &TaxRegistrationCreateParams{}
	}
	params.Context = ctx
	registration := &TaxRegistration{}
	err := c.B.Call(
		http.MethodPost, "/v1/tax/registrations", c.Key, params, registration)
	return registration, err
}

// Returns a Tax Registration object.
func (c v1TaxRegistrationService) Retrieve(ctx context.Context, id string, params *TaxRegistrationRetrieveParams) (*TaxRegistration, error) {
	if params == nil {
		params = &TaxRegistrationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax/registrations/%s", id)
	registration := &TaxRegistration{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, registration)
	return registration, err
}

// Updates an existing Tax Registration object.
//
// A registration cannot be deleted after it has been created. If you wish to end a registration you may do so by setting expires_at.
func (c v1TaxRegistrationService) Update(ctx context.Context, id string, params *TaxRegistrationUpdateParams) (*TaxRegistration, error) {
	if params == nil {
		params = &TaxRegistrationUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax/registrations/%s", id)
	registration := &TaxRegistration{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, registration)
	return registration, err
}

// Serializes a Registration create request into a batch job JSONL line.
func (c v1TaxRegistrationService) MarshalBatchCreate(params *TaxRegistrationCreateParams) (string, error) {
	itemID, err := newUUID4()
	if err != nil {
		return "", err
	}

	item := struct {
		ID            string            `json:"id"`
		Context       string            `json:"context,omitempty"`
		StripeVersion string            `json:"stripe_version,omitempty"`
		PathParams    map[string]string `json:"path_params,omitempty"`
		Params        interface{}       `json:"params"`
	}{
		ID:            itemID,
		PathParams:    nil,
		StripeVersion: APIVersion,
	}
	if params != nil {
		item.Params = params
		if params.StripeContext != nil {
			item.Context = *params.StripeContext
		}
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Serializes a Registration update request into a batch job JSONL line.
func (c v1TaxRegistrationService) MarshalBatchUpdate(id string, params *TaxRegistrationUpdateParams) (string, error) {
	itemID, err := newUUID4()
	if err != nil {
		return "", err
	}

	item := struct {
		ID            string            `json:"id"`
		Context       string            `json:"context,omitempty"`
		StripeVersion string            `json:"stripe_version,omitempty"`
		PathParams    map[string]string `json:"path_params,omitempty"`
		Params        interface{}       `json:"params"`
	}{
		ID:            itemID,
		PathParams:    map[string]string{"id": id},
		StripeVersion: APIVersion,
	}
	if params != nil {
		item.Params = params
		if params.StripeContext != nil {
			item.Context = *params.StripeContext
		}
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Returns a list of Tax Registration objects.
func (c v1TaxRegistrationService) List(ctx context.Context, listParams *TaxRegistrationListParams) *V1List[*TaxRegistration] {
	if listParams == nil {
		listParams = &TaxRegistrationListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*TaxRegistration], error) {
		list := &v1Page[*TaxRegistration]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/tax/registrations", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
