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

// Returns a list of Tax Registration objects.
func (c v1TaxRegistrationService) List(ctx context.Context, listParams *TaxRegistrationListParams) Seq2[*TaxRegistration, error] {
	if listParams == nil {
		listParams = &TaxRegistrationListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TaxRegistration, ListContainer, error) {
		list := &TaxRegistrationList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/tax/registrations", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
