//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Whether the alert was triggered or recovered.
type BillingAlertNotificationAction string

// List of values that BillingAlertNotificationAction can take
const (
	BillingAlertNotificationActionRecovered BillingAlertNotificationAction = "recovered"
	BillingAlertNotificationActionTriggered BillingAlertNotificationAction = "triggered"
)

// The type of billing alert that generated this notification.
type BillingAlertNotificationAlertType string

// List of values that BillingAlertNotificationAlertType can take
const (
	BillingAlertNotificationAlertTypeCreditBalanceThreshold BillingAlertNotificationAlertType = "credit_balance_threshold"
	BillingAlertNotificationAlertTypeSpendThreshold         BillingAlertNotificationAlertType = "spend_threshold"
	BillingAlertNotificationAlertTypeUsageThreshold         BillingAlertNotificationAlertType = "usage_threshold"
)

// Lists sent billing alert triggered and recovered notifications for a billing alert.
type BillingAlertNotificationListParams struct {
	ListParams `form:"*"`
	// The billing alert ID.
	ID *string `form:"-"` // Included in URL
	// Filter results to only include triggered or recovered notifications.
	Action *string `form:"action" json:"action,omitempty"`
	// Filter results to only include notifications for the given billing cadence.
	Cadence *string `form:"cadence" json:"cadence,omitempty"`
	// The customer to list notifications for.
	Customer *string `form:"customer" json:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Filter results to only include notifications for the given meter.
	Meter *string `form:"meter" json:"meter,omitempty"`
	// Filter results according to when the notification was sent.
	NotifiedAt *int64 `form:"notified_at" json:"notified_at,omitempty"`
	// Filter results according to when the notification was sent.
	NotifiedAtRange *RangeQueryParams `form:"notified_at" json:"-"`
	// Filter results to only include notifications for the given subscription.
	Subscription *string `form:"subscription" json:"subscription,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertNotificationListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type BillingAlertNotification struct {
	// Whether the alert was triggered or recovered.
	Action BillingAlertNotificationAction `json:"action"`
	// End of the aggregation period for which this notification was sent. Only present for usage threshold alerts.
	AggregationPeriodEnd int64 `json:"aggregation_period_end"`
	// Start of the aggregation period for which this notification was sent. Only present for usage threshold alerts.
	AggregationPeriodStart int64 `json:"aggregation_period_start"`
	// ID of the billing alert that generated this notification.
	Alert string `json:"alert"`
	// The type of billing alert that generated this notification.
	AlertType BillingAlertNotificationAlertType `json:"alert_type"`
	// The billing cadence associated with this notification. Only present for spend threshold alerts grouped by billing cadence.
	Cadence string `json:"cadence"`
	// Three-letter ISO currency code for the value, in lowercase. Only present for spend and credit balance threshold alerts.
	Currency Currency `json:"currency"`
	// ID of the customer for which the alert notification was sent.
	Customer string `json:"customer"`
	// Custom pricing unit for the threshold value
	CustomPricingUnit string `json:"custom_pricing_unit,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// ID of the billing meter associated with this notification. Only present for usage threshold alerts.
	Meter string `json:"meter"`
	// ID of the event delivered for this notification. Retrievable via the Events API for a limited time; for long-term audit scenarios, capture the full event payload at webhook delivery time.
	NotificationEvent string `json:"notification_event"`
	// Time at which the notification was sent. Measured in seconds since the Unix epoch.
	NotifiedAt int64 `json:"notified_at"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// ID of the subscription associated with this notification. Only present for spend threshold alerts grouped by subscription.
	Subscription string `json:"subscription"`
	// The value that triggered the alert. This may be a decimal string for custom pricing unit alerts. For usage threshold alerts, this is the meter event count. For credit balance and spend threshold alerts, this is the amount in the smallest currency unit.
	Value string `json:"value"`
}

// BillingAlertNotificationList is a list of AlertNotifications as retrieved from a list endpoint.
type BillingAlertNotificationList struct {
	APIResource
	ListMeta
	Data []*BillingAlertNotification `json:"data"`
}
