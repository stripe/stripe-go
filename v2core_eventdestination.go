//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The AWS-reported lifecycle state of the partner event source.
type V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus string

// List of values that V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus can take
const (
	V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatusActive  V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus = "active"
	V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatusDeleted V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus = "deleted"
	V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatusPending V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus = "pending"
	V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatusUnknown V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus = "unknown"
)

// The Azure-reported lifecycle state of the partner topic.
type V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatus string

// List of values that V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatus can take
const (
	V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatusActivated      V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatus = "activated"
	V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatusDeleted        V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatus = "deleted"
	V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatusNeverActivated V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatus = "never_activated"
	V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatusUnknown        V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatus = "unknown"
)

// Whether to deliver as snapshot or thin events.
type V2CoreEventDestinationEventPayload string

// List of values that V2CoreEventDestinationEventPayload can take
const (
	V2CoreEventDestinationEventPayloadSnapshot V2CoreEventDestinationEventPayload = "snapshot"
	V2CoreEventDestinationEventPayloadThin     V2CoreEventDestinationEventPayload = "thin"
)

// Whether Stripe currently attempts delivery. Stripe attempts delivery to enabled destinations when their provider configuration is active; disabled destinations do not receive delivery attempts.
type V2CoreEventDestinationStatus string

// List of values that V2CoreEventDestinationStatus can take
const (
	V2CoreEventDestinationStatusDisabled V2CoreEventDestinationStatus = "disabled"
	V2CoreEventDestinationStatusEnabled  V2CoreEventDestinationStatus = "enabled"
)

// Reason event destination has been disabled.
type V2CoreEventDestinationStatusDetailsDisabledReason string

// List of values that V2CoreEventDestinationStatusDetailsDisabledReason can take
const (
	V2CoreEventDestinationStatusDetailsDisabledReasonNoAwsEventSourceExists    V2CoreEventDestinationStatusDetailsDisabledReason = "no_aws_event_source_exists"
	V2CoreEventDestinationStatusDetailsDisabledReasonNoAzurePartnerTopicExists V2CoreEventDestinationStatusDetailsDisabledReason = "no_azure_partner_topic_exists"
	V2CoreEventDestinationStatusDetailsDisabledReasonUser                      V2CoreEventDestinationStatusDetailsDisabledReason = "user"
)

// The delivery transport. Chosen when the destination is created and cannot be changed by update.
type V2CoreEventDestinationType string

// List of values that V2CoreEventDestinationType can take
const (
	V2CoreEventDestinationTypeAmazonEventbridge V2CoreEventDestinationType = "amazon_eventbridge"
	V2CoreEventDestinationTypeAzureEventGrid    V2CoreEventDestinationType = "azure_event_grid"
	V2CoreEventDestinationTypeWebhookEndpoint   V2CoreEventDestinationType = "webhook_endpoint"
)

// Configuration for delivering events through an Amazon EventBridge partner event source.
type V2CoreEventDestinationAmazonEventbridge struct {
	// The AWS account ID that owns the event bus receiving events.
	AwsAccountID string `json:"aws_account_id"`
	// The ARN of the Stripe-created partner event source in your AWS account.
	AwsEventSourceArn string `json:"aws_event_source_arn"`
	// The AWS-reported lifecycle state of the partner event source.
	AwsEventSourceStatus V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus `json:"aws_event_source_status"`
}

// Configuration for delivering events through an Azure Event Grid partner topic.
type V2CoreEventDestinationAzureEventGrid struct {
	// The name of the Stripe-created partner topic that receives events.
	AzurePartnerTopicName string `json:"azure_partner_topic_name"`
	// The Azure-reported lifecycle state of the partner topic.
	AzurePartnerTopicStatus V2CoreEventDestinationAzureEventGridAzurePartnerTopicStatus `json:"azure_partner_topic_status"`
	// The Azure region where the partner topic is located.
	AzureRegion string `json:"azure_region"`
	// The Azure resource group containing the partner topic.
	AzureResourceGroupName string `json:"azure_resource_group_name"`
	// The Azure subscription containing the resource group and partner topic.
	AzureSubscriptionID string `json:"azure_subscription_id"`
}

// Present when the destination was disabled; identifies the cause, time, and provider-side object involved when available.
type V2CoreEventDestinationStatusDetailsDisabled struct {
	// Reason event destination has been disabled.
	Reason V2CoreEventDestinationStatusDetailsDisabledReason `json:"reason"`
}

// Additional lifecycle context for the destination status, when available.
type V2CoreEventDestinationStatusDetails struct {
	// Present when the destination was disabled; identifies the cause, time, and provider-side object involved when available.
	Disabled *V2CoreEventDestinationStatusDetailsDisabled `json:"disabled,omitempty"`
}

// Configuration for delivering events to a webhook endpoint. Live mode requires HTTPS; sandbox mode also supports HTTP.
type V2CoreEventDestinationWebhookEndpoint struct {
	// The secret used to verify Stripe signatures on delivered events. Returned only in the create response when explicitly included; public API clients cannot retrieve it later.
	SigningSecret string `json:"signing_secret,omitempty"`
	// The URL where Stripe sends matching events. Live mode requires HTTPS; sandbox mode also supports HTTP. Returned only when explicitly included.
	URL string `json:"url,omitempty"`
}

// Set up an event destination to receive events from Stripe across multiple destination types, including [webhook endpoints](https://docs.stripe.com/webhooks), [Amazon EventBridge](https://docs.stripe.com/event-destinations/eventbridge), and [Azure Event Grid](https://docs.stripe.com/event-destinations/eventgrid). Event destinations support receiving [thin events](https://docs.stripe.com/api/v2/events) and [snapshot events](https://docs.stripe.com/api/events).
type V2CoreEventDestination struct {
	APIResource
	// Configuration for delivering events through an Amazon EventBridge partner event source.
	AmazonEventbridge *V2CoreEventDestinationAmazonEventbridge `json:"amazon_eventbridge,omitempty"`
	// Configuration for delivering events through an Azure Event Grid partner topic.
	AzureEventGrid *V2CoreEventDestinationAzureEventGrid `json:"azure_event_grid,omitempty"`
	// The time when the destination was created.
	Created time.Time `json:"created"`
	// An optional user-defined description of the destination's purpose.
	Description string `json:"description"`
	// The list of event types enabled for delivery to this destination.
	EnabledEvents []string `json:"enabled_events"`
	// Whether to deliver as snapshot or thin events.
	EventPayload V2CoreEventDestinationEventPayload `json:"event_payload"`
	// Specifies which accounts' events route to this destination.
	// `@self`: Receive events from the account that owns the event destination.
	// `@accounts`: Receive events emitted from other accounts you manage which includes your v1 and v2 accounts.
	// `@organization_members`: Receive events from accounts directly linked to the organization.
	// `@organization_members/@accounts`: Receive events from all accounts connected to any platform accounts in the organization.
	EventsFrom []string `json:"events_from,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// User-defined key/value data for the destination; it has no effect on event matching or delivery.
	Metadata map[string]string `json:"metadata,omitempty"`
	// A user-defined label for identifying the destination in Stripe.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// For snapshot events only, the Stripe API version used to render event objects. You can't change this value after you create the event destination. Thin events are not pinned to an API version.
	SnapshotAPIVersion string `json:"snapshot_api_version,omitempty"`
	// Whether Stripe currently attempts delivery. Stripe attempts delivery to enabled destinations when their provider configuration is active; disabled destinations do not receive delivery attempts.
	Status V2CoreEventDestinationStatus `json:"status"`
	// Additional lifecycle context for the destination status, when available.
	StatusDetails *V2CoreEventDestinationStatusDetails `json:"status_details,omitempty"`
	// The delivery transport. Chosen when the destination is created and cannot be changed by update.
	Type V2CoreEventDestinationType `json:"type"`
	// The time when the destination object was last updated.
	Updated time.Time `json:"updated"`
	// Configuration for delivering events to a webhook endpoint. Live mode requires HTTPS; sandbox mode also supports HTTP.
	WebhookEndpoint *V2CoreEventDestinationWebhookEndpoint `json:"webhook_endpoint,omitempty"`
}
