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

// v1WebhookEndpointService is used to invoke /v1/webhook_endpoints APIs.
type v1WebhookEndpointService struct {
	B   Backend
	Key string
}

// A webhook endpoint must have a url and a list of enabled_events. You may optionally specify the Boolean connect parameter. If set to true, then a Connect webhook endpoint that notifies the specified url about events from all connected accounts is created; otherwise an account webhook endpoint that notifies the specified url only about events from your account is created. You can also create webhook endpoints in the [webhooks settings](https://dashboard.stripe.com/account/webhooks) section of the Dashboard.
func (c v1WebhookEndpointService) Create(ctx context.Context, params *WebhookEndpointCreateParams) (*WebhookEndpoint, error) {
	if params == nil {
		params = &WebhookEndpointCreateParams{}
	}
	params.Context = ctx
	webhookendpoint := &WebhookEndpoint{}
	err := c.B.Call(
		http.MethodPost, "/v1/webhook_endpoints", c.Key, params, webhookendpoint)
	return webhookendpoint, err
}

// Retrieves the webhook endpoint with the given ID.
func (c v1WebhookEndpointService) Retrieve(ctx context.Context, id string, params *WebhookEndpointRetrieveParams) (*WebhookEndpoint, error) {
	if params == nil {
		params = &WebhookEndpointRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/webhook_endpoints/%s", id)
	webhookendpoint := &WebhookEndpoint{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, webhookendpoint)
	return webhookendpoint, err
}

// Updates the webhook endpoint. You may edit the url, the list of enabled_events, and the status of your endpoint.
func (c v1WebhookEndpointService) Update(ctx context.Context, id string, params *WebhookEndpointUpdateParams) (*WebhookEndpoint, error) {
	if params == nil {
		params = &WebhookEndpointUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/webhook_endpoints/%s", id)
	webhookendpoint := &WebhookEndpoint{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, webhookendpoint)
	return webhookendpoint, err
}

// You can also delete webhook endpoints via the [webhook endpoint management](https://dashboard.stripe.com/account/webhooks) page of the Stripe dashboard.
func (c v1WebhookEndpointService) Delete(ctx context.Context, id string, params *WebhookEndpointDeleteParams) (*WebhookEndpoint, error) {
	if params == nil {
		params = &WebhookEndpointDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/webhook_endpoints/%s", id)
	webhookendpoint := &WebhookEndpoint{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, webhookendpoint)
	return webhookendpoint, err
}

// Returns a list of your webhook endpoints.
func (c v1WebhookEndpointService) List(ctx context.Context, listParams *WebhookEndpointListParams) Seq2[*WebhookEndpoint, error] {
	if listParams == nil {
		listParams = &WebhookEndpointListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*WebhookEndpoint, ListContainer, error) {
		list := &WebhookEndpointList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/webhook_endpoints", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
