//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Returns a list of your webhook endpoints.
type WebhookEndpointListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// A webhook endpoint must have a url and a list of enabled_events. You may optionally specify the Boolean connect parameter. If set to true, then a Connect webhook endpoint that notifies the specified url about events from all connected accounts is created; otherwise an account webhook endpoint that notifies the specified url only about events from your account is created. You can also create webhook endpoints in the [webhooks settings](https://dashboard.stripe.com/account/webhooks) section of the Dashboard.
type WebhookEndpointParams struct {
	Params        `form:"*"`
	Connect       *bool     `form:"connect"`
	Description   *string   `form:"description"`
	Disabled      *bool     `form:"disabled"`
	EnabledEvents []*string `form:"enabled_events"`
	URL           *string   `form:"url"`

	// This parameter is only available on creation.
	// We recommend setting the API version that the library is pinned to. See apiversion in stripe.go
	APIVersion *string `form:"api_version"`
}

// You can configure [webhook endpoints](https://stripe.com/docs/webhooks/) via the API to be
// notified about events that happen in your Stripe account or connected
// accounts.
//
// Most users configure webhooks from [the dashboard](https://dashboard.stripe.com/webhooks), which provides a user interface for registering and testing your webhook endpoints.
//
// Related guide: [Setting up Webhooks](https://stripe.com/docs/webhooks/configure).
type WebhookEndpoint struct {
	APIResource
	APIVersion    string            `json:"api_version"`
	Application   string            `json:"application"`
	Connect       bool              `json:"connect"`
	Created       int64             `json:"created"`
	Deleted       bool              `json:"deleted"`
	Description   string            `json:"description"`
	EnabledEvents []string          `json:"enabled_events"`
	ID            string            `json:"id"`
	Livemode      bool              `json:"livemode"`
	Metadata      map[string]string `json:"metadata"`
	Object        string            `json:"object"`
	Secret        string            `json:"secret"`
	Status        string            `json:"status"`
	URL           string            `json:"url"`
}

// WebhookEndpointList is a list of WebhookEndpoints as retrieved from a list endpoint.
type WebhookEndpointList struct {
	APIResource
	ListMeta
	Data []*WebhookEndpoint `json:"data"`
}

// UnmarshalJSON handles deserialization of a WebhookEndpoint.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (w *WebhookEndpoint) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		w.ID = id
		return nil
	}

	type webhookEndpoint WebhookEndpoint
	var v webhookEndpoint
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*w = WebhookEndpoint(v)
	return nil
}
