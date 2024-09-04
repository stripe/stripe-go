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

// Lists billing active and inactive alerts
type BillingAlertListParams struct {
	ListParams `form:"*"`
	// Filter results to only include this type of alert.
	AlertType *string `form:"alert_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Filter results to only include alerts with the given meter.
	Meter *string `form:"meter"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Filters to limit the scope of an alert.
type BillingAlertFilterParams struct {
	// Limit the scope to this alert only to this customer.
	Customer *string `form:"customer"`
	// Limit the scope of this rated usage alert to this subscription.
	Subscription *string `form:"subscription"`
	// Limit the scope of this rated usage alert to this subscription item.
	SubscriptionItem *string `form:"subscription_item"`
}

// The configuration of the usage threshold.
type BillingAlertUsageThresholdConfigParams struct {
	// Defines at which value the alert will fire.
	GTE *int64 `form:"gte"`
	// The [Billing Meter](https://stripe.com/api/billing/meter) ID whose usage is monitored.
	Meter *string `form:"meter"`
	// Whether the alert should only fire only once, or once per billing cycle.
	Recurrence *string `form:"recurrence"`
}

// Creates a billing alert
type BillingAlertParams struct {
	Params `form:"*"`
	// The type of alert to create.
	AlertType *string `form:"alert_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Filters to limit the scope of an alert.
	Filter *BillingAlertFilterParams `form:"filter"`
	// The title of the alert.
	Title *string `form:"title"`
	// The configuration of the usage threshold.
	UsageThresholdConfig *BillingAlertUsageThresholdConfigParams `form:"usage_threshold_config"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Reactivates this alert, allowing it to trigger again.
type BillingAlertActivateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertActivateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Archives this alert, removing it from the list view and APIs. This is non-reversible.
type BillingAlertArchiveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertArchiveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Deactivates this alert, preventing it from triggering.
type BillingAlertDeactivateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertDeactivateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

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
	APIResource
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

// BillingAlertList is a list of Alerts as retrieved from a list endpoint.
type BillingAlertList struct {
	APIResource
	ListMeta
	Data []*BillingAlert `json:"data"`
}
