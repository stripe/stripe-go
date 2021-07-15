//
//
// File generated from our OpenAPI spec
//
//

// Package webhookendpoint provides the /webhook_endpoints APIs
package webhookendpoint

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /webhook_endpoints APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new webhook endpoint.
func New(params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	return getC().New(params)
}

// New creates a new webhook endpoint.
func (c Client) New(params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	webhookendpoint := &stripe.WebhookEndpoint{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/webhook_endpoints",
		c.Key,
		params,
		webhookendpoint,
	)
	return webhookendpoint, err
}

// Get returns the details of a webhook endpoint.
func Get(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	return getC().Get(id, params)
}

// Get returns the details of a webhook endpoint.
func (c Client) Get(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	path := stripe.FormatURLPath("/v1/webhook_endpoints/%s", id)
	webhookendpoint := &stripe.WebhookEndpoint{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, webhookendpoint)
	return webhookendpoint, err
}

// Update updates a webhook endpoint's properties.
func Update(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	return getC().Update(id, params)
}

// Update updates a webhook endpoint's properties.
func (c Client) Update(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	path := stripe.FormatURLPath("/v1/webhook_endpoints/%s", id)
	webhookendpoint := &stripe.WebhookEndpoint{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, webhookendpoint)
	return webhookendpoint, err
}

// Del removes a webhook endpoint.
func Del(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	return getC().Del(id, params)
}

// Del removes a webhook endpoint.
func (c Client) Del(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	path := stripe.FormatURLPath("/v1/webhook_endpoints/%s", id)
	webhookendpoint := &stripe.WebhookEndpoint{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, webhookendpoint)
	return webhookendpoint, err
}

// List returns a list of webhook endpoints.
func List(params *stripe.WebhookEndpointListParams) *Iter {
	return getC().List(params)
}

// List returns a list of webhook endpoints.
func (c Client) List(listParams *stripe.WebhookEndpointListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.WebhookEndpointList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/webhook_endpoints", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for webhook endpoints.
type Iter struct {
	*stripe.Iter
}

// WebhookEndpoint returns the webhook endpoint which the iterator is currently pointing to.
func (i *Iter) WebhookEndpoint() *stripe.WebhookEndpoint {
	return i.Current().(*stripe.WebhookEndpoint)
}

// WebhookEndpointList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) WebhookEndpointList() *stripe.WebhookEndpointList {
	return i.List().(*stripe.WebhookEndpointList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
