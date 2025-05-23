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

// v1RadarEarlyFraudWarningService is used to invoke /v1/radar/early_fraud_warnings APIs.
type v1RadarEarlyFraudWarningService struct {
	B   Backend
	Key string
}

// Retrieves the details of an early fraud warning that has previously been created.
//
// Please refer to the [early fraud warning](https://docs.stripe.com/api#early_fraud_warning_object) object reference for more details.
func (c v1RadarEarlyFraudWarningService) Retrieve(ctx context.Context, id string, params *RadarEarlyFraudWarningRetrieveParams) (*RadarEarlyFraudWarning, error) {
	if params == nil {
		params = &RadarEarlyFraudWarningRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/radar/early_fraud_warnings/%s", id)
	earlyfraudwarning := &RadarEarlyFraudWarning{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, earlyfraudwarning)
	return earlyfraudwarning, err
}

// Returns a list of early fraud warnings.
func (c v1RadarEarlyFraudWarningService) List(ctx context.Context, listParams *RadarEarlyFraudWarningListParams) Seq2[*RadarEarlyFraudWarning, error] {
	if listParams == nil {
		listParams = &RadarEarlyFraudWarningListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*RadarEarlyFraudWarning, ListContainer, error) {
		list := &RadarEarlyFraudWarningList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/radar/early_fraud_warnings", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
