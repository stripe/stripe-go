//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Payload type of events being subscribed to.
type V2CoreEventDestinationEventPayload string

// List of values that V2CoreEventDestinationEventPayload can take
const (
	V2CoreEventDestinationEventPayloadSnapshot V2CoreEventDestinationEventPayload = "snapshot"
	V2CoreEventDestinationEventPayloadThin     V2CoreEventDestinationEventPayload = "thin"
)

// Where events should be routed from.
type V2CoreEventDestinationEventsFrom string

// List of values that V2CoreEventDestinationEventsFrom can take
const (
	V2CoreEventDestinationEventsFromOtherAccounts V2CoreEventDestinationEventsFrom = "other_accounts"
	V2CoreEventDestinationEventsFromSelf          V2CoreEventDestinationEventsFrom = "self"
)

// Status. It can be set to either enabled or disabled.
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
	V2CoreEventDestinationStatusDetailsDisabledReasonNoAwsEventSourceExists V2CoreEventDestinationStatusDetailsDisabledReason = "no_aws_event_source_exists"
	V2CoreEventDestinationStatusDetailsDisabledReasonUser                   V2CoreEventDestinationStatusDetailsDisabledReason = "user"
)

// Event destination type.
type V2CoreEventDestinationType string

// List of values that V2CoreEventDestinationType can take
const (
	V2CoreEventDestinationTypeAmazonEventbridge V2CoreEventDestinationType = "amazon_eventbridge"
	V2CoreEventDestinationTypeWebhookEndpoint   V2CoreEventDestinationType = "webhook_endpoint"
)

// The state of the AWS event source.
type V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus string

// List of values that V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus can take
const (
	V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatusActive  V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus = "active"
	V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatusDeleted V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus = "deleted"
	V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatusPending V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus = "pending"
	V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatusUnknown V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus = "unknown"
)

// Details about why the event destination has been disabled.
type V2CoreEventDestinationStatusDetailsDisabled struct {
	// Reason event destination has been disabled.
	Reason V2CoreEventDestinationStatusDetailsDisabledReason `json:"reason"`
}

// Additional information about event destination status.
type V2CoreEventDestinationStatusDetails struct {
	// Details about why the event destination has been disabled.
	Disabled *V2CoreEventDestinationStatusDetailsDisabled `json:"disabled,omitempty"`
}

// Amazon EventBridge configuration.
type V2CoreEventDestinationAmazonEventbridge struct {
	// The AWS account ID.
	AwsAccountID string `json:"aws_account_id"`
	// The ARN of the AWS event source.
	AwsEventSourceArn string `json:"aws_event_source_arn"`
	// The state of the AWS event source.
	AwsEventSourceStatus V2CoreEventDestinationAmazonEventbridgeAwsEventSourceStatus `json:"aws_event_source_status"`
}

// Webhook endpoint configuration.
type V2CoreEventDestinationWebhookEndpoint struct {
	// The signing secret of the webhook endpoint, only includable on creation.
	SigningSecret string `json:"signing_secret,omitempty"`
	// The URL of the webhook endpoint, includable.
	URL string `json:"url,omitempty"`
}

// Set up an event destination to receive events from Stripe across multiple destination types, including [webhook endpoints](https://docs.stripe.com/webhooks) and [Amazon EventBridge](https://docs.stripe.com/event-destinations/eventbridge). Event destinations support receiving [thin events](https://docs.stripe.com/api/v2/events) and [snapshot events](https://docs.stripe.com/api/events).
type V2CoreEventDestination struct {
	APIResource
	// Amazon EventBridge configuration.
	AmazonEventbridge *V2CoreEventDestinationAmazonEventbridge `json:"amazon_eventbridge,omitempty"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// An optional description of what the event destination is used for.
	Description string `json:"description"`
	// The list of events to enable for this endpoint.
	EnabledEvents []string `json:"enabled_events"`
	// Payload type of events being subscribed to.
	EventPayload V2CoreEventDestinationEventPayload `json:"event_payload"`
	// Where events should be routed from.
	EventsFrom []V2CoreEventDestinationEventsFrom `json:"events_from,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Metadata.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Event destination name.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// If using the snapshot event payload, the API version events are rendered as.
	SnapshotAPIVersion string `json:"snapshot_api_version,omitempty"`
	// Status. It can be set to either enabled or disabled.
	Status V2CoreEventDestinationStatus `json:"status"`
	// Additional information about event destination status.
	StatusDetails *V2CoreEventDestinationStatusDetails `json:"status_details,omitempty"`
	// Event destination type.
	Type V2CoreEventDestinationType `json:"type"`
	// Time at which the object was last updated.
	Updated time.Time `json:"updated"`
	// Webhook endpoint configuration.
	WebhookEndpoint *V2CoreEventDestinationWebhookEndpoint `json:"webhook_endpoint,omitempty"`
}
