//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The method for mapping a meter event to a customer.
type BillingMeterCustomerMappingType string

// List of values that BillingMeterCustomerMappingType can take
const (
	BillingMeterCustomerMappingTypeByID BillingMeterCustomerMappingType = "by_id"
)

// Specifies how events are aggregated.
type BillingMeterDefaultAggregationFormula string

// List of values that BillingMeterDefaultAggregationFormula can take
const (
	BillingMeterDefaultAggregationFormulaCount BillingMeterDefaultAggregationFormula = "count"
	BillingMeterDefaultAggregationFormulaSum   BillingMeterDefaultAggregationFormula = "sum"
)

// The time window to pre-aggregate usage events for, if any.
type BillingMeterEventTimeWindow string

// List of values that BillingMeterEventTimeWindow can take
const (
	BillingMeterEventTimeWindowDay  BillingMeterEventTimeWindow = "day"
	BillingMeterEventTimeWindowHour BillingMeterEventTimeWindow = "hour"
)

// The meter's status.
type BillingMeterStatus string

// List of values that BillingMeterStatus can take
const (
	BillingMeterStatusActive   BillingMeterStatus = "active"
	BillingMeterStatusInactive BillingMeterStatus = "inactive"
)

// Retrieve a list of billing meters.
type BillingMeterListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Filter results to only include meters with the given status.
	Status *string `form:"status"`
}

// AddExpand appends a new field to expand.
func (p *BillingMeterListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Fields that specify how to map a meter event to a customer.
type BillingMeterCustomerMappingParams struct {
	// The key in the usage event payload to use for mapping the event to a customer.
	EventPayloadKey *string `form:"event_payload_key"`
	// The method for mapping a meter event to a customer. Must be `by_id`.
	Type *string `form:"type"`
}

// The default settings to aggregate a meter's events with.
type BillingMeterDefaultAggregationParams struct {
	// Specifies how events are aggregated. Allowed values are `count` to count the number of events, `sum` to sum each event's value, or `last` to use the last event's value.
	Formula *string `form:"formula"`
}

// Fields that specify how to calculate a usage event's value.
type BillingMeterValueSettingsParams struct {
	// The key in the usage event payload to use as the value for this meter. For example, if the event payload contains usage on a `bytes_used` field, then set the event_payload_key to "bytes_used".
	EventPayloadKey *string `form:"event_payload_key"`
}

// Creates a billing meter
type BillingMeterParams struct {
	Params `form:"*"`
	// Fields that specify how to map a meter event to a customer.
	CustomerMapping *BillingMeterCustomerMappingParams `form:"customer_mapping"`
	// The default settings to aggregate a meter's events with.
	DefaultAggregation *BillingMeterDefaultAggregationParams `form:"default_aggregation"`
	// The meter's name.
	DisplayName *string `form:"display_name"`
	// The name of the usage event to record usage for. Corresponds with the `event_name` field on usage events.
	EventName *string `form:"event_name"`
	// The time window to pre-aggregate usage events for, if any.
	EventTimeWindow *string `form:"event_time_window"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Fields that specify how to calculate a usage event's value.
	ValueSettings *BillingMeterValueSettingsParams `form:"value_settings"`
}

// AddExpand appends a new field to expand.
func (p *BillingMeterParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Deactivates a billing meter
type BillingMeterDeactivateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingMeterDeactivateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Reactivates a billing meter
type BillingMeterReactivateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingMeterReactivateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type BillingMeterCustomerMapping struct {
	// The key in the usage event payload to use for mapping the event to a customer.
	EventPayloadKey string `json:"event_payload_key"`
	// The method for mapping a meter event to a customer.
	Type BillingMeterCustomerMappingType `json:"type"`
}
type BillingMeterDefaultAggregation struct {
	// Specifies how events are aggregated.
	Formula BillingMeterDefaultAggregationFormula `json:"formula"`
}
type BillingMeterStatusTransitions struct {
	// The time the meter was deactivated, if any. Measured in seconds since Unix epoch.
	DeactivatedAt int64 `json:"deactivated_at"`
}
type BillingMeterValueSettings struct {
	// The key in the usage event payload to use as the value for this meter.
	EventPayloadKey string `json:"event_payload_key"`
}

// A billing meter is a resource that allows you to track usage of a particular event. For example, you might create a billing meter to track the number of API calls made by a particular user. You can then use the billing meter to charge the user for the number of API calls they make.
type BillingMeter struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created            int64                           `json:"created"`
	CustomerMapping    *BillingMeterCustomerMapping    `json:"customer_mapping"`
	DefaultAggregation *BillingMeterDefaultAggregation `json:"default_aggregation"`
	// The meter's name.
	DisplayName string `json:"display_name"`
	// The name of the usage event to record usage for. Corresponds with the `event_name` field on usage events.
	EventName string `json:"event_name"`
	// The time window to pre-aggregate usage events for, if any.
	EventTimeWindow BillingMeterEventTimeWindow `json:"event_time_window"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The meter's status.
	Status            BillingMeterStatus             `json:"status"`
	StatusTransitions *BillingMeterStatusTransitions `json:"status_transitions"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	Updated       int64                      `json:"updated"`
	ValueSettings *BillingMeterValueSettings `json:"value_settings"`
}

// BillingMeterList is a list of Meters as retrieved from a list endpoint.
type BillingMeterList struct {
	APIResource
	ListMeta
	Data []*BillingMeter `json:"data"`
}
