// Package subschedule provides the /subscription_schedules APIs
package subschedule

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to invoke /subscription_schedules APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Cancel removes a subscription schedule.
func Cancel(id string, params *stripe.SubscriptionScheduleCancelParams) (*stripe.SubscriptionSchedule, error) {
	return getC().Cancel(id, params)
}

// Cancel removes a subscription schedule.
func (c Client) Cancel(id string, params *stripe.SubscriptionScheduleCancelParams) (*stripe.SubscriptionSchedule, error) {
	path := stripe.FormatURLPath("/v1/subscription_schedules/%s/cancel", id)
	sched := &stripe.SubscriptionSchedule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, sched)
	return sched, err
}

// Get returns the details of a subscription schedule.
func Get(id string, params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	return getC().Get(id, params)
}

// Get returns the details of a subscription schedule.
func (c Client) Get(id string, params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	path := stripe.FormatURLPath("/v1/subscription_schedules/%s", id)
	sched := &stripe.SubscriptionSchedule{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, sched)
	return sched, err
}

// List returns a list of subscriptions.
func List(params *stripe.SubscriptionScheduleListParams) *Iter {
	return getC().List(params)
}

// List returns a list of subscriptions.
func (c Client) List(listParams *stripe.SubscriptionScheduleListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.SubscriptionScheduleList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/subscription_schedules", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// New creates a new subscription schedule.
func New(params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	return getC().New(params)
}

// New creates a new subscription schedule.
func (c Client) New(params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	sched := &stripe.SubscriptionSchedule{}
	err := c.B.Call(http.MethodPost, "/v1/subscription_schedules", c.Key, params, sched)
	return sched, err
}

// Release releases a subscription schedule's properties.
func Release(id string, params *stripe.SubscriptionScheduleReleaseParams) (*stripe.SubscriptionSchedule, error) {
	return getC().Release(id, params)
}

// Release releases a subscription schedule's properties.
func (c Client) Release(id string, params *stripe.SubscriptionScheduleReleaseParams) (*stripe.SubscriptionSchedule, error) {
	path := stripe.FormatURLPath("/v1/subscription_schedules/%s/release", id)
	sched := &stripe.SubscriptionSchedule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, sched)

	return sched, err
}

// Update updates a subscription schedule's properties.
func Update(id string, params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	return getC().Update(id, params)
}

// Update updates a subscription schedule's properties.
func (c Client) Update(id string, params *stripe.SubscriptionScheduleParams) (*stripe.SubscriptionSchedule, error) {
	path := stripe.FormatURLPath("/v1/subscription_schedules/%s", id)
	sched := &stripe.SubscriptionSchedule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, sched)

	return sched, err
}

// Iter is an iterator for subscription schedules.
type Iter struct {
	*stripe.Iter
}

// SubscriptionSchedule returns the subscription schedule which the iterator is currently pointing to.
func (i *Iter) SubscriptionSchedule() *stripe.SubscriptionSchedule {
	return i.Current().(*stripe.SubscriptionSchedule)
}

// SubscriptionScheduleList returns the current list object which the iterator
// is currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SubscriptionScheduleList() *stripe.SubscriptionScheduleList {
	return i.List().(*stripe.SubscriptionScheduleList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
