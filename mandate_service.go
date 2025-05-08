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

// v1MandateService is used to invoke /v1/mandates APIs.
type v1MandateService struct {
	B   Backend
	Key string
}

// Retrieves a Mandate object.
func (c v1MandateService) Retrieve(ctx context.Context, id string, params *MandateRetrieveParams) (*Mandate, error) {
	if params == nil {
		params = &MandateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/mandates/%s", id)
	mandate := &Mandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, mandate)
	return mandate, err
}
