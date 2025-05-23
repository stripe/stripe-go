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

// v1BillingMeterService is used to invoke /v1/billing/meters APIs.
type v1BillingMeterService struct {
	B   Backend
	Key string
}

// Creates a billing meter.
func (c v1BillingMeterService) Create(ctx context.Context, params *BillingMeterCreateParams) (*BillingMeter, error) {
	if params == nil {
		params = &BillingMeterCreateParams{}
	}
	params.Context = ctx
	meter := &BillingMeter{}
	err := c.B.Call(http.MethodPost, "/v1/billing/meters", c.Key, params, meter)
	return meter, err
}

// Retrieves a billing meter given an ID.
func (c v1BillingMeterService) Retrieve(ctx context.Context, id string, params *BillingMeterRetrieveParams) (*BillingMeter, error) {
	if params == nil {
		params = &BillingMeterRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/meters/%s", id)
	meter := &BillingMeter{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, meter)
	return meter, err
}

// Updates a billing meter.
func (c v1BillingMeterService) Update(ctx context.Context, id string, params *BillingMeterUpdateParams) (*BillingMeter, error) {
	if params == nil {
		params = &BillingMeterUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/meters/%s", id)
	meter := &BillingMeter{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, meter)
	return meter, err
}

// When a meter is deactivated, no more meter events will be accepted for this meter. You can't attach a deactivated meter to a price.
func (c v1BillingMeterService) Deactivate(ctx context.Context, id string, params *BillingMeterDeactivateParams) (*BillingMeter, error) {
	if params == nil {
		params = &BillingMeterDeactivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/meters/%s/deactivate", id)
	meter := &BillingMeter{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, meter)
	return meter, err
}

// When a meter is reactivated, events for this meter can be accepted and you can attach the meter to a price.
func (c v1BillingMeterService) Reactivate(ctx context.Context, id string, params *BillingMeterReactivateParams) (*BillingMeter, error) {
	if params == nil {
		params = &BillingMeterReactivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/meters/%s/reactivate", id)
	meter := &BillingMeter{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, meter)
	return meter, err
}

// Retrieve a list of billing meters.
func (c v1BillingMeterService) List(ctx context.Context, listParams *BillingMeterListParams) Seq2[*BillingMeter, error] {
	if listParams == nil {
		listParams = &BillingMeterListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*BillingMeter, ListContainer, error) {
		list := &BillingMeterList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/billing/meters", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
