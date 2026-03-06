//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
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

// Amends an existing subscription schedule.
func (c v1SubscriptionScheduleService) Amend(ctx context.Context, id string, params *SubscriptionScheduleAmendParams) (*SubscriptionSchedule, error) {
	if params == nil {
		params = &SubscriptionScheduleAmendParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/subscription_schedules/%s/amend", id)
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

// Serializes a SubscriptionSchedule create request into a batch job JSONL line.
func (c v1SubscriptionScheduleService) SerializeBatchCreate(ctx context.Context, params *SubscriptionScheduleCreateParams) (string, error) {
	// Generate a UUID v4 using crypto/rand
	var uuidBytes [16]byte
	if _, err := rand.Read(uuidBytes[:]); err != nil {
		return "", err
	}
	uuidBytes[6] = (uuidBytes[6] & 0x0f) | 0x40 // version 4
	uuidBytes[8] = (uuidBytes[8] & 0x3f) | 0x80 // variant bits
	itemID := fmt.Sprintf("%x-%x-%x-%x-%x", uuidBytes[0:4], uuidBytes[4:6], uuidBytes[6:8], uuidBytes[8:10], uuidBytes[10:16])

	item := BatchItemParams{
		ID:            itemID,
		PathParams:    nil,
		Params:        params,
		StripeVersion: APIVersion, // Default to global API version
	}
	if params.StripeContext != nil {
		item.Context = *params.StripeContext
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Retrieves the list of your subscription schedules.
func (c v1SubscriptionScheduleService) List(ctx context.Context, listParams *SubscriptionScheduleListParams) Seq2[*SubscriptionSchedule, error] {
	if listParams == nil {
		listParams = &SubscriptionScheduleListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*SubscriptionSchedule], error) {
		list := &v1Page[*SubscriptionSchedule]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/subscription_schedules", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
