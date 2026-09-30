//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists all event destinations.
type V2CoreEventDestinationListParams struct {
	Params `form:"*"`
	// Include the normally redacted `webhook_endpoint.url` in each returned destination.
	Include []*string `form:"include" json:"include,omitempty"`
	// The page size.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// AWS account and region where Stripe creates the EventBridge partner event source.
type V2CoreEventDestinationAmazonEventbridgeParams struct {
	// Your AWS account where Stripe creates the partner event source.
	AwsAccountID *string `form:"aws_account_id" json:"aws_account_id"`
	// The AWS region where Stripe creates the partner event source.
	AwsRegion *string `form:"aws_region" json:"aws_region"`
}

// Azure subscription, resource group, and region where Stripe creates the partner topic.
type V2CoreEventDestinationAzureEventGridParams struct {
	// The Azure region where Stripe creates the partner topic.
	AzureRegion *string `form:"azure_region" json:"azure_region"`
	// The Azure resource group where Stripe creates the partner topic.
	AzureResourceGroupName *string `form:"azure_resource_group_name" json:"azure_resource_group_name"`
	// The Azure subscription where Stripe creates the partner topic.
	AzureSubscriptionID *string `form:"azure_subscription_id" json:"azure_subscription_id"`
}

// Delivery target for the webhook endpoint. Live mode requires HTTPS; sandbox mode also supports HTTP.
type V2CoreEventDestinationWebhookEndpointParams struct {
	// The URL where Stripe sends matching events. Live mode requires HTTPS; sandbox mode also supports HTTP.
	URL *string `form:"url" json:"url"`
}

// Create a new event destination.
type V2CoreEventDestinationParams struct {
	Params `form:"*"`
	// AWS account and region where Stripe creates the EventBridge partner event source.
	AmazonEventbridge *V2CoreEventDestinationAmazonEventbridgeParams `form:"amazon_eventbridge" json:"amazon_eventbridge,omitempty"`
	// Azure subscription, resource group, and region where Stripe creates the partner topic.
	AzureEventGrid *V2CoreEventDestinationAzureEventGridParams `form:"azure_event_grid" json:"azure_event_grid,omitempty"`
	// An optional user-defined description of the destination's purpose; it does not control routing.
	Description *string `form:"description" json:"description,omitempty"`
	// The list of event types enabled for delivery to this destination. Event scopes are configured when the destination is created.
	EnabledEvents []*string `form:"enabled_events" json:"enabled_events,omitempty"`
	// Whether to deliver as snapshot or thin events.
	EventPayload *string `form:"event_payload" json:"event_payload,omitempty"`
	// The account or organization scopes that can supply events. Use this with `enabled_events` to define the subscription.
	// `@self`: Receive events from the account that owns the event destination.
	// `@accounts`: Receive events emitted from other accounts you manage, including your v1 and v2 accounts.
	// `@organization_members`: Receive events from accounts directly linked to the organization.
	// `@organization_members/@accounts`: Receive events from all accounts connected to any platform accounts in the organization.
	EventsFrom []*string `form:"events_from" json:"events_from,omitempty"`
	// Include normally redacted webhook fields in the create response. Public API clients must include `webhook_endpoint.signing_secret` to receive the signing secret.
	Include []*string `form:"include" json:"include,omitempty"`
	// User-defined key/value data for the destination.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// A user-defined label for identifying the destination; it does not control routing.
	Name *string `form:"name" json:"name,omitempty"`
	// For snapshot events only, the Stripe API version used to render event objects; do not provide this for thin events.
	SnapshotAPIVersion *string `form:"snapshot_api_version" json:"snapshot_api_version,omitempty"`
	// The delivery transport. Chosen when the destination is created and cannot be changed by update.
	Type *string `form:"type" json:"type,omitempty"`
	// New delivery target for the webhook endpoint. Live mode requires HTTPS; sandbox mode also supports HTTP.
	WebhookEndpoint *V2CoreEventDestinationWebhookEndpointParams `form:"webhook_endpoint" json:"webhook_endpoint,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreEventDestinationParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
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

// AWS account and region where Stripe creates the EventBridge partner event source.
type V2CoreEventDestinationCreateAmazonEventbridgeParams struct {
	// Your AWS account where Stripe creates the partner event source.
	AwsAccountID *string `form:"aws_account_id" json:"aws_account_id"`
	// The AWS region where Stripe creates the partner event source.
	AwsRegion *string `form:"aws_region" json:"aws_region"`
}

// Azure subscription, resource group, and region where Stripe creates the partner topic.
type V2CoreEventDestinationCreateAzureEventGridParams struct {
	// The Azure region where Stripe creates the partner topic.
	AzureRegion *string `form:"azure_region" json:"azure_region"`
	// The Azure resource group where Stripe creates the partner topic.
	AzureResourceGroupName *string `form:"azure_resource_group_name" json:"azure_resource_group_name"`
	// The Azure subscription where Stripe creates the partner topic.
	AzureSubscriptionID *string `form:"azure_subscription_id" json:"azure_subscription_id"`
}

// Delivery target for the webhook endpoint. Live mode requires HTTPS; sandbox mode also supports HTTP.
type V2CoreEventDestinationCreateWebhookEndpointParams struct {
	// The URL where Stripe sends matching events. Live mode requires HTTPS; sandbox mode also supports HTTP.
	URL *string `form:"url" json:"url"`
}

// Create a new event destination.
type V2CoreEventDestinationCreateParams struct {
	Params `form:"*"`
	// AWS account and region where Stripe creates the EventBridge partner event source.
	AmazonEventbridge *V2CoreEventDestinationCreateAmazonEventbridgeParams `form:"amazon_eventbridge" json:"amazon_eventbridge,omitempty"`
	// Azure subscription, resource group, and region where Stripe creates the partner topic.
	AzureEventGrid *V2CoreEventDestinationCreateAzureEventGridParams `form:"azure_event_grid" json:"azure_event_grid,omitempty"`
	// An optional user-defined description of the destination's purpose.
	Description *string `form:"description" json:"description,omitempty"`
	// The list of event types enabled for delivery to this destination.
	EnabledEvents []*string `form:"enabled_events" json:"enabled_events"`
	// Whether to deliver as snapshot or thin events.
	EventPayload *string `form:"event_payload" json:"event_payload"`
	// The account or organization scopes that can supply events. Use this with `enabled_events` to define the subscription.
	// `@self`: Receive events from the account that owns the event destination.
	// `@accounts`: Receive events emitted from other accounts you manage, including your v1 and v2 accounts.
	// `@organization_members`: Receive events from accounts directly linked to the organization.
	// `@organization_members/@accounts`: Receive events from all accounts connected to any platform accounts in the organization.
	EventsFrom []*string `form:"events_from" json:"events_from,omitempty"`
	// Include normally redacted webhook fields in the create response. Public API clients must include `webhook_endpoint.signing_secret` to receive the signing secret.
	Include []*string `form:"include" json:"include,omitempty"`
	// User-defined key/value data for the destination.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// A user-defined label for identifying the destination.
	Name *string `form:"name" json:"name"`
	// For snapshot events only, the Stripe API version used to render event objects; do not provide this for thin events.
	SnapshotAPIVersion *string `form:"snapshot_api_version" json:"snapshot_api_version,omitempty"`
	// The delivery transport. Chosen when the destination is created and cannot be changed by update.
	Type *string `form:"type" json:"type"`
	// Delivery target for the webhook endpoint. Live mode requires HTTPS; sandbox mode also supports HTTP.
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

// New delivery target for the webhook endpoint. Live mode requires HTTPS; sandbox mode also supports HTTP.
type V2CoreEventDestinationUpdateWebhookEndpointParams struct {
	// The URL where Stripe sends matching events. Live mode requires HTTPS; sandbox mode also supports HTTP.
	URL *string `form:"url" json:"url"`
}

// Update the details of an event destination.
type V2CoreEventDestinationUpdateParams struct {
	Params `form:"*"`
	// An optional user-defined description of the destination's purpose; it does not control routing.
	Description *string `form:"description" json:"description,omitempty"`
	// The list of event types enabled for delivery to this destination. Event scopes are configured when the destination is created.
	EnabledEvents []*string `form:"enabled_events" json:"enabled_events,omitempty"`
	// Include the normally redacted `webhook_endpoint.url` in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Metadata.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// A user-defined label for identifying the destination; it does not control routing.
	Name *string `form:"name" json:"name,omitempty"`
	// New delivery target for the webhook endpoint. Live mode requires HTTPS; sandbox mode also supports HTTP.
	WebhookEndpoint *V2CoreEventDestinationUpdateWebhookEndpointParams `form:"webhook_endpoint" json:"webhook_endpoint,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2CoreEventDestinationUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
