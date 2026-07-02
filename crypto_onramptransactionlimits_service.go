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

// v1CryptoOnrampTransactionLimitsService is used to invoke /v1/crypto/onramp_transaction_limits APIs.
type v1CryptoOnrampTransactionLimitsService struct {
	B   Backend
	Key string
}

// Retrieves the remaining onramp limit for a crypto customer.
func (c v1CryptoOnrampTransactionLimitsService) Retrieve(ctx context.Context, params *CryptoOnrampTransactionLimitsRetrieveParams) (*CryptoOnrampTransactionLimits, error) {
	if params == nil {
		params = &CryptoOnrampTransactionLimitsRetrieveParams{}
	}
	params.Context = ctx
	onramptransactionlimits := &CryptoOnrampTransactionLimits{}
	err := c.B.Call(
		http.MethodGet, "/v1/crypto/onramp_transaction_limits", c.Key, params, onramptransactionlimits)
	return onramptransactionlimits, err
}
