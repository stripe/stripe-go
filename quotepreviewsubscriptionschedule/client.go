//
//
// File generated from our OpenAPI spec
//
//

// Package quotepreviewsubscriptionschedule provides the /quotes/{quote}/preview_subscription_schedules APIs
package quotepreviewsubscriptionschedule

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /quotes/{quote}/preview_subscription_schedules APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of quote preview subscription schedules.
func List(params *stripe.QuotePreviewSubscriptionScheduleListParams) *Iter {
	return getC().List(params)
}

// List returns a list of quote preview subscription schedules.
func (c Client) List(listParams *stripe.QuotePreviewSubscriptionScheduleListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/preview_subscription_schedules",
		stripe.StringValue(listParams.Quote),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.QuotePreviewSubscriptionScheduleList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for quote preview subscription schedules.
type Iter struct {
	*stripe.Iter
}

// QuotePreviewSubscriptionSchedule returns the quote preview subscription schedule which the iterator is currently pointing to.
func (i *Iter) QuotePreviewSubscriptionSchedule() *stripe.QuotePreviewSubscriptionSchedule {
	return i.Current().(*stripe.QuotePreviewSubscriptionSchedule)
}

// QuotePreviewSubscriptionScheduleList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) QuotePreviewSubscriptionScheduleList() *stripe.QuotePreviewSubscriptionScheduleList {
	return i.List().(*stripe.QuotePreviewSubscriptionScheduleList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
