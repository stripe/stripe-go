// Package subschedulerevision provides the /subscription_schedules/revisions APIs
package subschedulerevision

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /subscription_schedule/revisions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a subscription schedule revision
func Get(id string, params *stripe.SubscriptionScheduleRevisionParams) (*stripe.SubscriptionScheduleRevision, error) {
	return getC().Get(id, params)
}

// Get returns the details of a subscription schedule revision
func (c Client) Get(id string, params *stripe.SubscriptionScheduleRevisionParams) (*stripe.SubscriptionScheduleRevision, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.Schedule == nil {
		return nil, fmt.Errorf("params.Schedule must be set")
	}

	path := stripe.FormatURLPath("/v1/subscription_schedules/%s/revisions/%s",
		stripe.StringValue(params.Schedule), id)
	rev := &stripe.SubscriptionScheduleRevision{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, rev)
	return rev, err
}

// List returns a list of subscription schedule revisions.
func List(params *stripe.SubscriptionScheduleRevisionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of subscription schedule revisions.
func (c Client) List(listParams *stripe.SubscriptionScheduleRevisionListParams) *Iter {
	path := stripe.FormatURLPath("/v1/subscription_schedules/%s/revisions",
		stripe.StringValue(listParams.Schedule))

	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.SubscriptionScheduleRevisionList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for subscription schedule revisions.
type Iter struct {
	*stripe.Iter
}

// SubscriptionScheduleRevision returns the subscription schedule revision which the iterator is
// currently pointing to.
func (i *Iter) SubscriptionScheduleRevision() *stripe.SubscriptionScheduleRevision {
	return i.Current().(*stripe.SubscriptionScheduleRevision)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
