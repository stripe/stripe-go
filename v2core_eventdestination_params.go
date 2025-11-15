//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists all event destinations.
type V2CoreEventDestinationListParams struct {
	Params `form:"*"`
	// Additional fields to include in the response. Currently supports `webhook_endpoint.url`.
	Include []*string `form:"include" json:"include,omitempty"`
	// The page size.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Amazon EventBridge configuration.
type V2CoreEventDestinationAmazonEventbridgeParams struct {
	// The AWS account ID.
	AwsAccountID *string `form:"aws_account_id" json:"aws_account_id"`
	// The region of the AWS event source.
	AwsRegion *string `form:"aws_region" json:"aws_region"`
}

// Webhook endpoint configuration.
type V2CoreEventDestinationWebhookEndpointParams struct {
	// The URL of the webhook endpoint.
	URL *string `form:"url" json:"url"`
}

// Create a new event destination.
type V2CoreEventDestinationParams struct {
	Params `form:"*"`
	// Amazon EventBridge configuration.
	AmazonEventbridge *V2CoreEventDestinationAmazonEventbridgeParams `form:"amazon_eventbridge" json:"amazon_eventbridge,omitempty"`
	// An optional description of what the event destination is used for.
	Description *string `form:"description" json:"description,omitempty"`
	// The list of events to enable for this endpoint.
	EnabledEvents []*string `form:"enabled_events" json:"enabled_events,omitempty"`
	// Payload type of events being subscribed to.
	EventPayload *string `form:"event_payload" json:"event_payload,omitempty"`
	// Where events should be routed from.
	EventsFrom []*string `form:"events_from" json:"events_from,omitempty"`
	// Additional fields to include in the response. Currently supports `webhook_endpoint.url`.
	Include []*string `form:"include" json:"include,omitempty"`
	// Metadata.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Event destination name.
	Name *string `form:"name" json:"name,omitempty"`
	// If using the snapshot event payload, the API version events are rendered as.
	SnapshotAPIVersion *string `form:"snapshot_api_version" json:"snapshot_api_version,omitempty"`
	// Event destination type.
	Type *string `form:"type" json:"type,omitempty"`
	// Webhook endpoint configuration.
	WebhookEndpoint *V2CoreEventDestinationWebhookEndpointParams `form:"webhook_endpoint" json:"webhook_endpoint,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreEventDestinationParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Disable an event destination.
type V2CoreEventDestinationDisableParams struct {
	Params `form:"*"`
}

// Enable an event destination.
type V2CoreEventDestinationEnableParams struct {
	Params `form:"*"`
}

// Send a `ping` event to an event destination.
type V2CoreEventDestinationPingParams struct {
	Params `form:"*"`
}

// Amazon EventBridge configuration.
type V2CoreEventDestinationCreateAmazonEventbridgeParams struct {
	// The AWS account ID.
	AwsAccountID *string `form:"aws_account_id" json:"aws_account_id"`
	// The region of the AWS event source.
	AwsRegion *string `form:"aws_region" json:"aws_region"`
}

// Webhook endpoint configuration.
type V2CoreEventDestinationCreateWebhookEndpointParams struct {
	// The URL of the webhook endpoint.
	URL *string `form:"url" json:"url"`
}

// Create a new event destination.
type V2CoreEventDestinationCreateParams struct {
	Params `form:"*"`
	// Amazon EventBridge configuration.
	AmazonEventbridge *V2CoreEventDestinationCreateAmazonEventbridgeParams `form:"amazon_eventbridge" json:"amazon_eventbridge,omitempty"`
	// An optional description of what the event destination is used for.
	Description *string `form:"description" json:"description,omitempty"`
	// The list of events to enable for this endpoint.
	EnabledEvents []*string `form:"enabled_events" json:"enabled_events"`
	// Payload type of events being subscribed to.
	EventPayload *string `form:"event_payload" json:"event_payload"`
	// Where events should be routed from.
	EventsFrom []*string `form:"events_from" json:"events_from,omitempty"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Metadata.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Event destination name.
	Name *string `form:"name" json:"name"`
	// If using the snapshot event payload, the API version events are rendered as.
	SnapshotAPIVersion *string `form:"snapshot_api_version" json:"snapshot_api_version,omitempty"`
	// Event destination type.
	Type *string `form:"type" json:"type"`
	// Webhook endpoint configuration.
	WebhookEndpoint *V2CoreEventDestinationCreateWebhookEndpointParams `form:"webhook_endpoint" json:"webhook_endpoint,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreEventDestinationCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Delete an event destination.
type V2CoreEventDestinationDeleteParams struct {
	Params `form:"*"`
}

// Retrieves the details of an event destination.
type V2CoreEventDestinationRetrieveParams struct {
	Params `form:"*"`
	// Additional fields to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// Webhook endpoint configuration.
type V2CoreEventDestinationUpdateWebhookEndpointParams struct {
	// The URL of the webhook endpoint.
	URL *string `form:"url" json:"url"`
}

// Update the details of an event destination.
type V2CoreEventDestinationUpdateParams struct {
	Params `form:"*"`
	// An optional description of what the event destination is used for.
	Description *string `form:"description" json:"description,omitempty"`
	// The list of events to enable for this endpoint.
	EnabledEvents []*string `form:"enabled_events" json:"enabled_events,omitempty"`
	// Additional fields to include in the response. Currently supports `webhook_endpoint.url`.
	Include []*string `form:"include" json:"include,omitempty"`
	// Metadata.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Event destination name.
	Name *string `form:"name" json:"name,omitempty"`
	// Webhook endpoint configuration.
	WebhookEndpoint *V2CoreEventDestinationUpdateWebhookEndpointParams `form:"webhook_endpoint" json:"webhook_endpoint,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreEventDestinationUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
