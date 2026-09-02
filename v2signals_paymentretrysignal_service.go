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

// v2SignalsPaymentRetrySignalService is used to invoke paymentretrysignal related APIs.
type v2SignalsPaymentRetrySignalService struct {
	B   Backend
	Key string
}

// Retrieves a payment retry signal by ID.
func (c v2SignalsPaymentRetrySignalService) Retrieve(ctx context.Context, id string, params *V2SignalsPaymentRetrySignalRetrieveParams) (*V2SignalsPaymentRetrySignal, error) {
	if params == nil {
		params = &V2SignalsPaymentRetrySignalRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/signals/payment_retry_signals/%s", id)
	paymentretrysignal := &V2SignalsPaymentRetrySignal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentretrysignal)
	return paymentretrysignal, err
}
