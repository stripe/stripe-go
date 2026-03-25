//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of grouping used for this alert notification
type BillingAlertTriggeredGroupByType string

// List of values that BillingAlertTriggeredGroupByType can take
const (
	BillingAlertTriggeredGroupByTypeBillingCadence          BillingAlertTriggeredGroupByType = "billing_cadence"
	BillingAlertTriggeredGroupByTypePricingPlanSubscription BillingAlertTriggeredGroupByType = "pricing_plan_subscription"
)

// The aggregation period for which this alert triggered
type BillingAlertTriggeredAggregationPeriod struct {
	// End time of the aggregation period
	EndsAt int64 `json:"ends_at"`
	// Start time of the aggregation period
	StartsAt int64 `json:"starts_at"`
}

// Populated specifically for spend alerts to notify merchants which group_by entity has the exceeded spend
type BillingAlertTriggeredGroupBy struct {
	// The billing cadence ID, populated when type is `billing_cadence`
	BillingCadence string `json:"billing_cadence"`
	// The pricing plan subscription ID, populated when type is `pricing_plan_subscription`
	PricingPlanSubscription string `json:"pricing_plan_subscription"`
	// The type of grouping used for this alert notification
	Type BillingAlertTriggeredGroupByType `json:"type"`
}
type BillingAlertTriggered struct {
	// The aggregation period for which this alert triggered
	AggregationPeriod *BillingAlertTriggeredAggregationPeriod `json:"aggregation_period,omitempty"`
	// A billing alert is a resource that notifies you when a certain usage threshold on a meter is crossed. For example, you might create a billing alert to notify you when a certain user made 100 API requests.
	Alert *BillingAlert `json:"alert"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Currency for the threshold value
	Currency Currency `json:"currency,omitempty"`
	// ID of customer for which the alert triggered
	Customer string `json:"customer"`
	// Custom pricing unit for the threshold value
	CustomPricingUnit string `json:"custom_pricing_unit,omitempty"`
	// External customer ID for the customer for which the alert triggered
	ExternalCustomerID string `json:"external_customer_id,omitempty"`
	// Populated specifically for spend alerts to notify merchants which group_by entity has the exceeded spend
	GroupBy *BillingAlertTriggeredGroupBy `json:"group_by,omitempty"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Timestamp at which the threshold was crossed
	TriggeredAt int64 `json:"triggered_at,omitempty"`
	// The value triggering the alert
	Value float64 `json:"value,string"`
}
