//
//
// File generated from our OpenAPI spec
//
//

// Package alertnotification provides the /v1/billing/alerts/{id}/notifications APIs
package alertnotification

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/billing/alerts/{id}/notifications APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Lists sent billing alert triggered and recovered notifications for a billing alert.
func List(params *stripe.BillingAlertNotificationListParams) *Iter {
	return getC().List(params)
}

// Lists sent billing alert triggered and recovered notifications for a billing alert.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.BillingAlertNotificationListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/billing/alerts/%s/notifications", stripe.StringValue(listParams.ID))
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.BillingAlertNotificationList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for billing alert notifications.
type Iter struct {
	*stripe.Iter
}

// BillingAlertNotification returns the billing alert notification which the iterator is currently pointing to.
func (i *Iter) BillingAlertNotification() *stripe.BillingAlertNotification {
	return i.Current().(*stripe.BillingAlertNotification)
}

// BillingAlertNotificationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BillingAlertNotificationList() *stripe.BillingAlertNotificationList {
	return i.List().(*stripe.BillingAlertNotificationList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
