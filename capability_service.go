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

// v1CapabilityService is used to invoke /v1/accounts/{account}/capabilities APIs.
type v1CapabilityService struct {
	B   Backend
	Key string
}

// Retrieves information about the specified Account Capability.
func (c v1CapabilityService) Retrieve(ctx context.Context, id string, params *CapabilityRetrieveParams) (*Capability, error) {
	if params == nil {
		params = &CapabilityRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/capabilities/%s", StringValue(params.Account), id)
	capability := &Capability{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, capability)
	return capability, err
}

// Updates an existing Account Capability. Request or remove a capability by updating its requested parameter.
func (c v1CapabilityService) Update(ctx context.Context, id string, params *CapabilityUpdateParams) (*Capability, error) {
	if params == nil {
		params = &CapabilityUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/capabilities/%s", StringValue(params.Account), id)
	capability := &Capability{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, capability)
	return capability, err
}

// Returns a list of capabilities associated with the account. The capabilities are returned sorted by creation date, with the most recent capability appearing first.
func (c v1CapabilityService) List(ctx context.Context, listParams *CapabilityListParams) Seq2[*Capability, error] {
	if listParams == nil {
		listParams = &CapabilityListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/capabilities", StringValue(listParams.Account))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Capability, ListContainer, error) {
		list := &CapabilityList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
