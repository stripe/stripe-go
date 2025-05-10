//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1TaxAssociationService is used to invoke association related APIs.
type v1TaxAssociationService struct {
	B   Backend
	Key string
}

// Finds a tax association object by PaymentIntent id.
func (c v1TaxAssociationService) Find(ctx context.Context, params *TaxAssociationFindParams) (*TaxAssociation, error) {
	if params == nil {
		params = &TaxAssociationFindParams{}
	}
	params.Context = ctx
	association := &TaxAssociation{}
	err := c.B.Call(
		http.MethodGet, "/v1/tax/associations/find", c.Key, params, association)
	return association, err
}
