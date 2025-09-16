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

// v1BillingAlertService is used to invoke /v1/billing/alerts APIs.
type v1BillingAlertService struct {
	B   Backend
	Key string
}

// Creates a billing alert
func (c v1BillingAlertService) Create(ctx context.Context, params *BillingAlertCreateParams) (*BillingAlert, error) {
	if params == nil {
		params = &BillingAlertCreateParams{}
	}
	params.Context = ctx
	alert := &BillingAlert{}
	err := c.B.Call(http.MethodPost, "/v1/billing/alerts", c.Key, params, alert)
	return alert, err
}

// Retrieves a billing alert given an ID
func (c v1BillingAlertService) Retrieve(ctx context.Context, id string, params *BillingAlertRetrieveParams) (*BillingAlert, error) {
	if params == nil {
		params = &BillingAlertRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/alerts/%s", id)
	alert := &BillingAlert{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, alert)
	return alert, err
}

// Reactivates this alert, allowing it to trigger again.
func (c v1BillingAlertService) Activate(ctx context.Context, id string, params *BillingAlertActivateParams) (*BillingAlert, error) {
	if params == nil {
		params = &BillingAlertActivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/alerts/%s/activate", id)
	alert := &BillingAlert{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, alert)
	return alert, err
}

// Archives this alert, removing it from the list view and APIs. This is non-reversible.
func (c v1BillingAlertService) Archive(ctx context.Context, id string, params *BillingAlertArchiveParams) (*BillingAlert, error) {
	if params == nil {
		params = &BillingAlertArchiveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/alerts/%s/archive", id)
	alert := &BillingAlert{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, alert)
	return alert, err
}

// Deactivates this alert, preventing it from triggering.
func (c v1BillingAlertService) Deactivate(ctx context.Context, id string, params *BillingAlertDeactivateParams) (*BillingAlert, error) {
	if params == nil {
		params = &BillingAlertDeactivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/alerts/%s/deactivate", id)
	alert := &BillingAlert{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, alert)
	return alert, err
}

// Lists billing active and inactive alerts
func (c v1BillingAlertService) List(ctx context.Context, listParams *BillingAlertListParams) Seq2[*BillingAlert, error] {
	if listParams == nil {
		listParams = &BillingAlertListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*BillingAlert, ListContainer, error) {
		list := &BillingAlertList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/billing/alerts", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
