//
//
// File generated from our OpenAPI spec
//
//

// Package quotepreviewschedule provides the /quotes/{quote}/preview_subscription_schedules APIs
package quotepreviewschedule

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/form"
)

// Client is used to invoke /quotes/{quote}/preview_subscription_schedules APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of quote preview schedules.
func List(params *stripe.QuotePreviewScheduleListParams) *Iter {
	return getC().List(params)
}

// List returns a list of quote preview schedules.
func (c Client) List(listParams *stripe.QuotePreviewScheduleListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/quotes/%s/preview_subscription_schedules",
		stripe.StringValue(listParams.Quote),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.QuotePreviewScheduleList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for quote preview schedules.
type Iter struct {
	*stripe.Iter
}

// QuotePreviewSchedule returns the quote preview schedule which the iterator is currently pointing to.
func (i *Iter) QuotePreviewSchedule() *stripe.QuotePreviewSchedule {
	return i.Current().(*stripe.QuotePreviewSchedule)
}

// QuotePreviewScheduleList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) QuotePreviewScheduleList() *stripe.QuotePreviewScheduleList {
	return i.List().(*stripe.QuotePreviewScheduleList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
