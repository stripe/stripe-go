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

// v1AccountSignalsService is used to invoke /v1/accounts/{account_id}/signals APIs.
type v1AccountSignalsService struct {
	B   Backend
	Key string
}

// Retrieves the account's Signal objects
func (c v1AccountSignalsService) Retrieve(ctx context.Context, id string, params *AccountSignalsRetrieveParams) (*AccountSignals, error) {
	if params == nil {
		params = &AccountSignalsRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/accounts/%s/signals", id)
	accountsignals := &AccountSignals{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountsignals)
	return accountsignals, err
}
