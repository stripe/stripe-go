//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// WebhookEndpointParams is the set of parameters that can be used when creating a webhook endpoint.
// For more details see https://stripe.com/docs/api#create_webhook_endpoint.
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

// WebhookEndpointListParams is the set of parameters that can be used when listing webhook endpoints.
// For more detail see https://stripe.com/docs/api#list_webhook_endpoints.
type WebhookEndpointListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// WebhookEndpoint is the resource representing a Stripe webhook endpoint.
// For more details see https://stripe.com/docs/api#webhook_endpoints.
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

// WebhookEndpointList is a list of webhook endpoints as retrieved from a list endpoint.
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
