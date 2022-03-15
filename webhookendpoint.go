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
	Params `form:"*"`
	// Whether this endpoint should receive events from connected accounts (`true`), or from your account (`false`). Defaults to `false`.
	Connect *bool `form:"connect"`
	// An optional description of what the webhook is used for.
	Description *string `form:"description"`
	// Disable the webhook endpoint if set to true.
	Disabled *bool `form:"disabled"`
	// The list of events to enable for this endpoint. You may specify `['*']` to enable all events, except those that require explicit selection.
	EnabledEvents []*string `form:"enabled_events"`
	// The URL of the webhook endpoint.
	URL *string `form:"url"`

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
	// The API version events are rendered as for this webhook endpoint.
	APIVersion string `json:"api_version"`
	// The ID of the associated Connect application.
	Application string `json:"application"`
	Connect     bool   `json:"connect"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	Deleted bool  `json:"deleted"`
	// An optional description of what the webhook is used for.
	Description string `json:"description"`
	// The list of events to enable for this endpoint. `['*']` indicates that all events are enabled, except those that require explicit selection.
	EnabledEvents []string `json:"enabled_events"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The endpoint's secret, used to generate [webhook signatures](https://stripe.com/docs/webhooks/signatures). Only returned at creation.
	Secret string `json:"secret"`
	// The status of the webhook. It can be `enabled` or `disabled`.
	Status string `json:"status"`
	// The URL of the webhook endpoint.
	URL string `json:"url"`
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
