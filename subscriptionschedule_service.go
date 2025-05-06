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

// v1SubscriptionScheduleService is used to invoke /v1/subscription_schedules APIs.
type v1SubscriptionScheduleService struct {
	B   Backend
	Key string
}

// Creates a new subscription schedule object. Each customer can have up to 500 active or scheduled subscriptions.
func (c v1SubscriptionScheduleService) Create(ctx context.Context, params *SubscriptionScheduleCreateParams) (*SubscriptionSchedule, error) {
	if params == nil {
		params = &SubscriptionScheduleCreateParams{}
	}
	params.Context = ctx
	subscriptionschedule := &SubscriptionSchedule{}
	err := c.B.Call(
		http.MethodPost, "/v1/subscription_schedules", c.Key, params, subscriptionschedule)
	return subscriptionschedule, err
}

// Retrieves the details of an existing subscription schedule. You only need to supply the unique subscription schedule identifier that was returned upon subscription schedule creation.
func (c v1SubscriptionScheduleService) Retrieve(ctx context.Context, id string, params *SubscriptionScheduleRetrieveParams) (*SubscriptionSchedule, error) {
	if params == nil {
		params = &SubscriptionScheduleRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/subscription_schedules/%s", id)
	subscriptionschedule := &SubscriptionSchedule{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, subscriptionschedule)
	return subscriptionschedule, err
}

// Updates an existing subscription schedule.
func (c v1SubscriptionScheduleService) Update(ctx context.Context, id string, params *SubscriptionScheduleUpdateParams) (*SubscriptionSchedule, error) {
	if params == nil {
		params = &SubscriptionScheduleUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/subscription_schedules/%s", id)
	subscriptionschedule := &SubscriptionSchedule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscriptionschedule)
	return subscriptionschedule, err
}

// Cancels a subscription schedule and its associated subscription immediately (if the subscription schedule has an active subscription). A subscription schedule can only be canceled if its status is not_started or active.
func (c v1SubscriptionScheduleService) Cancel(ctx context.Context, id string, params *SubscriptionScheduleCancelParams) (*SubscriptionSchedule, error) {
	if params == nil {
		params = &SubscriptionScheduleCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/subscription_schedules/%s/cancel", id)
	subscriptionschedule := &SubscriptionSchedule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscriptionschedule)
	return subscriptionschedule, err
}

// Releases the subscription schedule immediately, which will stop scheduling of its phases, but leave any existing subscription in place. A schedule can only be released if its status is not_started or active. If the subscription schedule is currently associated with a subscription, releasing it will remove its subscription property and set the subscription's ID to the released_subscription property.
func (c v1SubscriptionScheduleService) Release(ctx context.Context, id string, params *SubscriptionScheduleReleaseParams) (*SubscriptionSchedule, error) {
	if params == nil {
		params = &SubscriptionScheduleReleaseParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/subscription_schedules/%s/release", id)
	subscriptionschedule := &SubscriptionSchedule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, subscriptionschedule)
	return subscriptionschedule, err
}

// Retrieves the list of your subscription schedules.
func (c v1SubscriptionScheduleService) List(ctx context.Context, listParams *SubscriptionScheduleListParams) Seq2[*SubscriptionSchedule, error] {
	if listParams == nil {
		listParams = &SubscriptionScheduleListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*SubscriptionSchedule, ListContainer, error) {
		list := &SubscriptionScheduleList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/subscription_schedules", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
