//
//
// File generated from our OpenAPI spec
//
//

// Package alert provides the /v1/billing/alerts APIs
package alert

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/billing/alerts APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a billing alert
func New(params *stripe.BillingAlertParams) (*stripe.BillingAlert, error) {
	return getC().New(params)
}

// Creates a billing alert
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.BillingAlertParams) (*stripe.BillingAlert, error) {
	alert := &stripe.BillingAlert{}
	err := c.B.Call(http.MethodPost, "/v1/billing/alerts", c.Key, params, alert)
	return alert, err
}

// Retrieves a billing alert given an ID
func Get(id string, params *stripe.BillingAlertParams) (*stripe.BillingAlert, error) {
	return getC().Get(id, params)
}

// Retrieves a billing alert given an ID
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.BillingAlertParams) (*stripe.BillingAlert, error) {
	path := stripe.FormatURLPath("/v1/billing/alerts/%s", id)
	alert := &stripe.BillingAlert{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, alert)
	return alert, err
}

// Reactivates this alert, allowing it to trigger again.
func Activate(id string, params *stripe.BillingAlertActivateParams) (*stripe.BillingAlert, error) {
	return getC().Activate(id, params)
}

// Reactivates this alert, allowing it to trigger again.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Activate(id string, params *stripe.BillingAlertActivateParams) (*stripe.BillingAlert, error) {
	path := stripe.FormatURLPath("/v1/billing/alerts/%s/activate", id)
	alert := &stripe.BillingAlert{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, alert)
	return alert, err
}

// Archives this alert, removing it from the list view and APIs. This is non-reversible.
func Archive(id string, params *stripe.BillingAlertArchiveParams) (*stripe.BillingAlert, error) {
	return getC().Archive(id, params)
}

// Archives this alert, removing it from the list view and APIs. This is non-reversible.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Archive(id string, params *stripe.BillingAlertArchiveParams) (*stripe.BillingAlert, error) {
	path := stripe.FormatURLPath("/v1/billing/alerts/%s/archive", id)
	alert := &stripe.BillingAlert{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, alert)
	return alert, err
}

// Deactivates this alert, preventing it from triggering.
func Deactivate(id string, params *stripe.BillingAlertDeactivateParams) (*stripe.BillingAlert, error) {
	return getC().Deactivate(id, params)
}

// Deactivates this alert, preventing it from triggering.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.BillingAlertDeactivateParams) (*stripe.BillingAlert, error) {
	path := stripe.FormatURLPath("/v1/billing/alerts/%s/deactivate", id)
	alert := &stripe.BillingAlert{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, alert)
	return alert, err
}

// Lists billing active and inactive alerts
func List(params *stripe.BillingAlertListParams) *Iter {
	return getC().List(params)
}

// Lists billing active and inactive alerts
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.BillingAlertListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.BillingAlertList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/billing/alerts", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for billing alerts.
type Iter struct {
	*stripe.Iter
}

// BillingAlert returns the billing alert which the iterator is currently pointing to.
func (i *Iter) BillingAlert() *stripe.BillingAlert {
	return i.Current().(*stripe.BillingAlert)
}

// BillingAlertList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BillingAlertList() *stripe.BillingAlertList {
	return i.List().(*stripe.BillingAlertList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
