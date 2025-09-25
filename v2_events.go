package stripe

// V2Events: The beginning of the section generated from our OpenAPI spec
import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// Open Enum.
type V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode string

// List of values that V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode can take
const (
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeArchivedMeter                   V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "archived_meter"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeMeterEventCustomerNotFound      V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "meter_event_customer_not_found"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeMeterEventDimensionCountTooHigh V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "meter_event_dimension_count_too_high"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeMeterEventInvalidValue          V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "meter_event_invalid_value"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeMeterEventNoCustomerDefined     V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "meter_event_no_customer_defined"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeMissingDimensionPayloadKeys     V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "missing_dimension_payload_keys"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeNoMeter                         V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "no_meter"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeTimestampInFuture               V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "timestamp_in_future"
	V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCodeTimestampTooFarInPast           V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode = "timestamp_too_far_in_past"
)

// Open Enum.
type V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode string

// List of values that V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode can take
const (
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeArchivedMeter                   V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "archived_meter"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeMeterEventCustomerNotFound      V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "meter_event_customer_not_found"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeMeterEventDimensionCountTooHigh V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "meter_event_dimension_count_too_high"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeMeterEventInvalidValue          V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "meter_event_invalid_value"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeMeterEventNoCustomerDefined     V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "meter_event_no_customer_defined"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeMissingDimensionPayloadKeys     V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "missing_dimension_payload_keys"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeNoMeter                         V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "no_meter"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeTimestampInFuture               V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "timestamp_in_future"
	V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCodeTimestampTooFarInPast           V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode = "timestamp_too_far_in_past"
)

// V2Event is the interface implemented by V2 Events. To get the underlying Event,
// use a type switch or type assertion to one of the concrete event types.
type V2Event interface {
	getBaseEvent() *V2BaseEvent
}

// V2RawEvent is the raw event type for V2 events. It is used to unmarshal the
// event data into a generic structure, and can also be used a default event
// type when the event type is not known.
type V2RawEvent struct {
	V2BaseEvent
	Data          *json.RawMessage `json:"data"`
	RelatedObject *RelatedObject   `json:"related_object"`
}

// Used for everything internal to the EventNotifications
type eventNotificationParams struct {
	Params `form:"*"`
}

// V1BillingMeterErrorReportTriggeredEvent is the Go struct for the "v1.billing.meter.error_report_triggered" event.
// Occurs when a Meter has invalid async usage events.
type V1BillingMeterErrorReportTriggeredEvent struct {
	V2BaseEvent
	Data               V1BillingMeterErrorReportTriggeredEventData `json:"data"`
	RelatedObject      RelatedObject                               `json:"related_object"`
	fetchRelatedObject func() (*BillingMeter, error)
}

// FetchRelatedObject fetches the BillingMeter related to the event.
func (e *V1BillingMeterErrorReportTriggeredEvent) FetchRelatedObject(ctx context.Context) (*BillingMeter, error) {
	return e.fetchRelatedObject()
}

// V1BillingMeterErrorReportTriggeredEventNotification is the webhook payload you'll get when handling an event with type "v1.billing.meter.error_report_triggered"
// Occurs when a Meter has invalid async usage events.
type V1BillingMeterErrorReportTriggeredEventNotification struct {
	V2EventNotification
	RelatedObject RelatedObject `json:"related_object"`
}

// GetEventNotification ensures we conform to `EventNotificationContainer`.
func (en *V1BillingMeterErrorReportTriggeredEventNotification) GetEventNotification() *V2EventNotification {
	return &en.V2EventNotification
}

// FetchEvent retrieves the V1BillingMeterErrorReportTriggeredEvent that created this Notification
func (en *V1BillingMeterErrorReportTriggeredEventNotification) FetchEvent(ctx context.Context) (*V1BillingMeterErrorReportTriggeredEvent, error) {
	evt, err := en.V2EventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V1BillingMeterErrorReportTriggeredEvent), nil
}

// FetchRelatedObject fetches the BillingMeter related to the event.
func (en *V1BillingMeterErrorReportTriggeredEventNotification) FetchRelatedObject(ctx context.Context) (*BillingMeter, error) {
	relatedObj := &BillingMeter{}
	err := en.client.backend.Call(
		http.MethodGet, en.RelatedObject.URL, en.client.key, &eventNotificationParams{Params: Params{Context: ctx, StripeContext: en.Context}}, relatedObj)
	return relatedObj, err
}

// V1BillingMeterNoMeterFoundEvent is the Go struct for the "v1.billing.meter.no_meter_found" event.
// Occurs when a Meter's id is missing or invalid in async usage events.
type V1BillingMeterNoMeterFoundEvent struct {
	V2BaseEvent
	Data V1BillingMeterNoMeterFoundEventData `json:"data"`
}

// V1BillingMeterNoMeterFoundEventNotification is the webhook payload you'll get when handling an event with type "v1.billing.meter.no_meter_found"
// Occurs when a Meter's id is missing or invalid in async usage events.
type V1BillingMeterNoMeterFoundEventNotification struct{ V2EventNotification }

// GetEventNotification ensures we conform to `EventNotificationContainer`.
func (en *V1BillingMeterNoMeterFoundEventNotification) GetEventNotification() *V2EventNotification {
	return &en.V2EventNotification
}

// FetchEvent retrieves the V1BillingMeterNoMeterFoundEvent that created this Notification
func (en *V1BillingMeterNoMeterFoundEventNotification) FetchEvent(ctx context.Context) (*V1BillingMeterNoMeterFoundEvent, error) {
	evt, err := en.V2EventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V1BillingMeterNoMeterFoundEvent), nil
}

// V2CoreEventDestinationPingEvent is the Go struct for the "v2.core.event_destination.ping" event.
// A ping event used to test the connection to an EventDestination.
type V2CoreEventDestinationPingEvent struct {
	V2BaseEvent
	RelatedObject      RelatedObject `json:"related_object"`
	fetchRelatedObject func() (*V2EventDestination, error)
}

// FetchRelatedObject fetches the V2EventDestination related to the event.
func (e *V2CoreEventDestinationPingEvent) FetchRelatedObject(ctx context.Context) (*V2EventDestination, error) {
	return e.fetchRelatedObject()
}

// V2CoreEventDestinationPingEventNotification is the webhook payload you'll get when handling an event with type "v2.core.event_destination.ping"
// A ping event used to test the connection to an EventDestination.
type V2CoreEventDestinationPingEventNotification struct {
	V2EventNotification
	RelatedObject RelatedObject `json:"related_object"`
}

// GetEventNotification ensures we conform to `EventNotificationContainer`.
func (en *V2CoreEventDestinationPingEventNotification) GetEventNotification() *V2EventNotification {
	return &en.V2EventNotification
}

// FetchEvent retrieves the V2CoreEventDestinationPingEvent that created this Notification
func (en *V2CoreEventDestinationPingEventNotification) FetchEvent(ctx context.Context) (*V2CoreEventDestinationPingEvent, error) {
	evt, err := en.V2EventNotification.fetchEvent(ctx)
	if err != nil {
		return nil, err
	}
	return evt.(*V2CoreEventDestinationPingEvent), nil
}

// FetchRelatedObject fetches the V2EventDestination related to the event.
func (en *V2CoreEventDestinationPingEventNotification) FetchRelatedObject(ctx context.Context) (*V2EventDestination, error) {
	relatedObj := &V2EventDestination{}
	err := en.client.backend.Call(
		http.MethodGet, en.RelatedObject.URL, en.client.key, &eventNotificationParams{Params: Params{Context: ctx, StripeContext: en.Context}}, relatedObj)
	return relatedObj, err
}

// The request causes the error.
type V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeSampleErrorRequest struct {
	// The request idempotency key.
	Identifier string `json:"identifier"`
}

// A list of sample errors of this type.
type V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeSampleError struct {
	// The error message.
	ErrorMessage string `json:"error_message"`
	// The request causes the error.
	Request *V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeSampleErrorRequest `json:"request"`
}

// The error details.
type V1BillingMeterErrorReportTriggeredEventDataReasonErrorType struct {
	// Open Enum.
	Code V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeCode `json:"code"`
	// The number of errors of this type.
	ErrorCount int64 `json:"error_count"`
	// A list of sample errors of this type.
	SampleErrors []*V1BillingMeterErrorReportTriggeredEventDataReasonErrorTypeSampleError `json:"sample_errors"`
}

// This contains information about why meter error happens.
type V1BillingMeterErrorReportTriggeredEventDataReason struct {
	// The total error count within this window.
	ErrorCount int64 `json:"error_count"`
	// The error details.
	ErrorTypes []*V1BillingMeterErrorReportTriggeredEventDataReasonErrorType `json:"error_types"`
}

// Occurs when a Meter has invalid async usage events.
type V1BillingMeterErrorReportTriggeredEventData struct {
	// Extra field included in the event's `data` when fetched from /v2/events.
	DeveloperMessageSummary string `json:"developer_message_summary"`
	// This contains information about why meter error happens.
	Reason *V1BillingMeterErrorReportTriggeredEventDataReason `json:"reason"`
	// The end of the window that is encapsulated by this summary.
	ValidationEnd time.Time `json:"validation_end"`
	// The start of the window that is encapsulated by this summary.
	ValidationStart time.Time `json:"validation_start"`
}

// The request causes the error.
type V1BillingMeterNoMeterFoundEventDataReasonErrorTypeSampleErrorRequest struct {
	// The request idempotency key.
	Identifier string `json:"identifier"`
}

// A list of sample errors of this type.
type V1BillingMeterNoMeterFoundEventDataReasonErrorTypeSampleError struct {
	// The error message.
	ErrorMessage string `json:"error_message"`
	// The request causes the error.
	Request *V1BillingMeterNoMeterFoundEventDataReasonErrorTypeSampleErrorRequest `json:"request"`
}

// The error details.
type V1BillingMeterNoMeterFoundEventDataReasonErrorType struct {
	// Open Enum.
	Code V1BillingMeterNoMeterFoundEventDataReasonErrorTypeCode `json:"code"`
	// The number of errors of this type.
	ErrorCount int64 `json:"error_count"`
	// A list of sample errors of this type.
	SampleErrors []*V1BillingMeterNoMeterFoundEventDataReasonErrorTypeSampleError `json:"sample_errors"`
}

// This contains information about why meter error happens.
type V1BillingMeterNoMeterFoundEventDataReason struct {
	// The total error count within this window.
	ErrorCount int64 `json:"error_count"`
	// The error details.
	ErrorTypes []*V1BillingMeterNoMeterFoundEventDataReasonErrorType `json:"error_types"`
}

// Occurs when a Meter's id is missing or invalid in async usage events.
type V1BillingMeterNoMeterFoundEventData struct {
	// Extra field included in the event's `data` when fetched from /v2/events.
	DeveloperMessageSummary string `json:"developer_message_summary"`
	// This contains information about why meter error happens.
	Reason *V1BillingMeterNoMeterFoundEventDataReason `json:"reason"`
	// The end of the window that is encapsulated by this summary.
	ValidationEnd time.Time `json:"validation_end"`
	// The start of the window that is encapsulated by this summary.
	ValidationStart time.Time `json:"validation_start"`
}

// ConvertRawEvent converts a raw event to a concrete event type.
// If the event type is not known, it returns the raw event.
func ConvertRawEvent(event *V2RawEvent, backend Backend, key string) (V2Event, error) {
	switch event.Type {
	case "v1.billing.meter.error_report_triggered":
		result := &V1BillingMeterErrorReportTriggeredEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*BillingMeter, error) {
			v := &BillingMeter{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v1.billing.meter.no_meter_found":
		result := &V1BillingMeterNoMeterFoundEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		if err := json.Unmarshal(*event.Data, &result.Data); err != nil {
			return nil, err
		}
		return result, nil
	case "v2.core.event_destination.ping":
		result := &V2CoreEventDestinationPingEvent{}
		result.V2BaseEvent = event.V2BaseEvent
		result.RelatedObject = *event.RelatedObject
		result.fetchRelatedObject = func() (*V2EventDestination, error) {
			v := &V2EventDestination{}
			err := backend.Call(http.MethodGet, event.RelatedObject.URL, key, nil, v)
			return v, err
		}
		return result, nil
	default:
		return event, nil
	}
}

// V2Events: The end of the section generated from our OpenAPI spec

// EventNotificationFromJSON is a helper for constructing an Event Notification. Doesn't perform signature validation,
// so you should use [Client.ParseEventNotification] instead for initial handling.
// This is useful in unit tests and working with EventNotifications that you've already validated the authenticity of.
func EventNotificationFromJSON(payload []byte, client Client) (EventNotificationContainer, error) {
	// we can pull the type out to base our matching on
	var result = &struct {
		Type string `json:"type"`
	}{}
	if err := json.Unmarshal(payload, result); err != nil {
		return nil, err
	}

	// V2EventNotificationTypes: The beginning of the section generated from our OpenAPI spec
	switch result.Type {
	case "v1.billing.meter.error_report_triggered":
		evt := V1BillingMeterErrorReportTriggeredEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v1.billing.meter.no_meter_found":
		evt := V1BillingMeterNoMeterFoundEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	case "v2.core.event_destination.ping":
		evt := V2CoreEventDestinationPingEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	default:
		evt := V2UnknownEventNotification{}
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, err
		}
		evt.client = client
		return &evt, nil
	}
	// V2EventNotificationTypes: The end of the section generated from our OpenAPI spec
}
