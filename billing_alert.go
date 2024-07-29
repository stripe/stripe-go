//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Defines the type of the alert.
type BillingAlertAlertType string

// List of values that BillingAlertAlertType can take
const (
	BillingAlertAlertTypeUsageThreshold BillingAlertAlertType = "usage_threshold"
)

// Status of the alert. This can be active, inactive or archived.
type BillingAlertStatus string

// List of values that BillingAlertStatus can take
const (
	BillingAlertStatusActive   BillingAlertStatus = "active"
	BillingAlertStatusArchived BillingAlertStatus = "archived"
	BillingAlertStatusInactive BillingAlertStatus = "inactive"
)

// Defines how the alert will behave.
type BillingAlertUsageThresholdConfigRecurrence string

// List of values that BillingAlertUsageThresholdConfigRecurrence can take
const (
	BillingAlertUsageThresholdConfigRecurrenceOneTime BillingAlertUsageThresholdConfigRecurrence = "one_time"
)

// Limits the scope of the alert to a specific [customer](https://stripe.com/docs/api/customers).
type BillingAlertFilter struct {
	// Limit the scope of the alert to this customer ID
	Customer *Customer `json:"customer"`
}

// Encapsulates configuration of the alert to monitor usage on a specific [Billing Meter](https://stripe.com/docs/api/billing/meter).
type BillingAlertUsageThresholdConfig struct {
	// The value at which this alert will trigger.
	GTE int64 `json:"gte"`
	// The [Billing Meter](https://stripe.com/api/billing/meter) ID whose usage is monitored.
	Meter *BillingMeter `json:"meter"`
	// Defines how the alert will behave.
	Recurrence BillingAlertUsageThresholdConfigRecurrence `json:"recurrence"`
}

// A billing alert is a resource that notifies you when a certain usage threshold on a meter is crossed. For example, you might create a billing alert to notify you when a certain user made 100 API requests.
type BillingAlert struct {
	// Defines the type of the alert.
	AlertType BillingAlertAlertType `json:"alert_type"`
	// Limits the scope of the alert to a specific [customer](https://stripe.com/docs/api/customers).
	Filter *BillingAlertFilter `json:"filter"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Status of the alert. This can be active, inactive or archived.
	Status BillingAlertStatus `json:"status"`
	// Title of the alert.
	Title string `json:"title"`
	// Encapsulates configuration of the alert to monitor usage on a specific [Billing Meter](https://stripe.com/docs/api/billing/meter).
	UsageThresholdConfig *BillingAlertUsageThresholdConfig `json:"usage_threshold_config"`
}
