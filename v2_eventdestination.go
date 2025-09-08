//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Payload type of events being subscribed to.
type V2EventDestinationEventPayload string

// List of values that V2EventDestinationEventPayload can take
const (
	V2EventDestinationEventPayloadSnapshot V2EventDestinationEventPayload = "snapshot"
	V2EventDestinationEventPayloadThin     V2EventDestinationEventPayload = "thin"
)

// Where events should be routed from.
type V2EventDestinationEventsFrom string

// List of values that V2EventDestinationEventsFrom can take
const (
	V2EventDestinationEventsFromOtherAccounts V2EventDestinationEventsFrom = "other_accounts"
	V2EventDestinationEventsFromSelf          V2EventDestinationEventsFrom = "self"
)

// Status. It can be set to either enabled or disabled.
type V2EventDestinationStatus string

// List of values that V2EventDestinationStatus can take
const (
	V2EventDestinationStatusDisabled V2EventDestinationStatus = "disabled"
	V2EventDestinationStatusEnabled  V2EventDestinationStatus = "enabled"
)

// Reason event destination has been disabled.
type V2EventDestinationStatusDetailsDisabledReason string

// List of values that V2EventDestinationStatusDetailsDisabledReason can take
const (
	V2EventDestinationStatusDetailsDisabledReasonNoAwsEventSourceExists V2EventDestinationStatusDetailsDisabledReason = "no_aws_event_source_exists"
	V2EventDestinationStatusDetailsDisabledReasonUser                   V2EventDestinationStatusDetailsDisabledReason = "user"
)

// Event destination type.
type V2EventDestinationType string

// List of values that V2EventDestinationType can take
const (
	V2EventDestinationTypeAmazonEventbridge V2EventDestinationType = "amazon_eventbridge"
	V2EventDestinationTypeWebhookEndpoint   V2EventDestinationType = "webhook_endpoint"
)

// The state of the AWS event source.
type V2EventDestinationAmazonEventbridgeAwsEventSourceStatus string

// List of values that V2EventDestinationAmazonEventbridgeAwsEventSourceStatus can take
const (
	V2EventDestinationAmazonEventbridgeAwsEventSourceStatusActive  V2EventDestinationAmazonEventbridgeAwsEventSourceStatus = "active"
	V2EventDestinationAmazonEventbridgeAwsEventSourceStatusDeleted V2EventDestinationAmazonEventbridgeAwsEventSourceStatus = "deleted"
	V2EventDestinationAmazonEventbridgeAwsEventSourceStatusPending V2EventDestinationAmazonEventbridgeAwsEventSourceStatus = "pending"
	V2EventDestinationAmazonEventbridgeAwsEventSourceStatusUnknown V2EventDestinationAmazonEventbridgeAwsEventSourceStatus = "unknown"
)

// Details about why the event destination has been disabled.
type V2EventDestinationStatusDetailsDisabled struct {
	// Reason event destination has been disabled.
	Reason V2EventDestinationStatusDetailsDisabledReason `json:"reason"`
}

// Additional information about event destination status.
type V2EventDestinationStatusDetails struct {
	// Details about why the event destination has been disabled.
	Disabled *V2EventDestinationStatusDetailsDisabled `json:"disabled"`
}

// Amazon EventBridge configuration.
type V2EventDestinationAmazonEventbridge struct {
	// The AWS account ID.
	AwsAccountID string `json:"aws_account_id"`
	// The ARN of the AWS event source.
	AwsEventSourceArn string `json:"aws_event_source_arn"`
	// The state of the AWS event source.
	AwsEventSourceStatus V2EventDestinationAmazonEventbridgeAwsEventSourceStatus `json:"aws_event_source_status"`
}

// Webhook endpoint configuration.
type V2EventDestinationWebhookEndpoint struct {
	// The signing secret of the webhook endpoint, only includable on creation.
	SigningSecret string `json:"signing_secret"`
	// The URL of the webhook endpoint, includable.
	URL string `json:"url"`
}

// Set up an event destination to receive events from Stripe across multiple destination types, including [webhook endpoints](https://docs.stripe.com/webhooks) and [Amazon EventBridge](https://docs.stripe.com/event-destinations/eventbridge). Event destinations support receiving [thin events](https://docs.stripe.com/api/v2/events) and [snapshot events](https://docs.stripe.com/api/events).
type V2EventDestination struct {
	APIResource
	// Amazon EventBridge configuration.
	AmazonEventbridge *V2EventDestinationAmazonEventbridge `json:"amazon_eventbridge"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// An optional description of what the event destination is used for.
	Description string `json:"description"`
	// The list of events to enable for this endpoint.
	EnabledEvents []string `json:"enabled_events"`
	// Payload type of events being subscribed to.
	EventPayload V2EventDestinationEventPayload `json:"event_payload"`
	// Where events should be routed from.
	EventsFrom []V2EventDestinationEventsFrom `json:"events_from"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Metadata.
	Metadata map[string]string `json:"metadata"`
	// Event destination name.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// If using the snapshot event payload, the API version events are rendered as.
	SnapshotAPIVersion string `json:"snapshot_api_version"`
	// Status. It can be set to either enabled or disabled.
	Status V2EventDestinationStatus `json:"status"`
	// Additional information about event destination status.
	StatusDetails *V2EventDestinationStatusDetails `json:"status_details"`
	// Event destination type.
	Type V2EventDestinationType `json:"type"`
	// Time at which the object was last updated.
	Updated time.Time `json:"updated"`
	// Webhook endpoint configuration.
	WebhookEndpoint *V2EventDestinationWebhookEndpoint `json:"webhook_endpoint"`
}
