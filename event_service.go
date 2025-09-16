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

// v1EventService is used to invoke /v1/events APIs.
type v1EventService struct {
	B   Backend
	Key string
}

// Retrieves the details of an event if it was created in the last 30 days. Supply the unique identifier of the event, which you might have received in a webhook.
func (c v1EventService) Retrieve(ctx context.Context, id string, params *EventRetrieveParams) (*Event, error) {
	if params == nil {
		params = &EventRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/events/%s", id)
	event := &Event{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, event)
	return event, err
}

// List events, going back up to 30 days. Each event data is rendered according to Stripe API version at its creation time, specified in [event object](https://docs.stripe.com/api/events/object) api_version attribute (not according to your current Stripe API version or Stripe-Version header).
func (c v1EventService) List(ctx context.Context, listParams *EventListParams) Seq2[*Event, error] {
	if listParams == nil {
		listParams = &EventListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Event, ListContainer, error) {
		list := &EventList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/events", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
