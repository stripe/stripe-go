//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Owner of the payment method.
type V2ProvisioningPaymentProfilePaymentMethodOwner string

// List of values that V2ProvisioningPaymentProfilePaymentMethodOwner can take
const (
	V2ProvisioningPaymentProfilePaymentMethodOwnerPlatform V2ProvisioningPaymentProfilePaymentMethodOwner = "platform"
)

// Interval over which `max_amount` applies.
type V2ProvisioningPaymentProfileProviderUsageLimitsRecurringInterval string

// List of values that V2ProvisioningPaymentProfileProviderUsageLimitsRecurringInterval can take
const (
	V2ProvisioningPaymentProfileProviderUsageLimitsRecurringIntervalMonth V2ProvisioningPaymentProfileProviderUsageLimitsRecurringInterval = "month"
	V2ProvisioningPaymentProfileProviderUsageLimitsRecurringIntervalWeek  V2ProvisioningPaymentProfileProviderUsageLimitsRecurringInterval = "week"
	V2ProvisioningPaymentProfileProviderUsageLimitsRecurringIntervalYear  V2ProvisioningPaymentProfileProviderUsageLimitsRecurringInterval = "year"
)

// Interval over which `max_amount` applies.
type V2ProvisioningPaymentProfileUsageLimitsRecurringInterval string

// List of values that V2ProvisioningPaymentProfileUsageLimitsRecurringInterval can take
const (
	V2ProvisioningPaymentProfileUsageLimitsRecurringIntervalMonth V2ProvisioningPaymentProfileUsageLimitsRecurringInterval = "month"
	V2ProvisioningPaymentProfileUsageLimitsRecurringIntervalWeek  V2ProvisioningPaymentProfileUsageLimitsRecurringInterval = "week"
	V2ProvisioningPaymentProfileUsageLimitsRecurringIntervalYear  V2ProvisioningPaymentProfileUsageLimitsRecurringInterval = "year"
)

// Usage limit applied to the payment method for this provider.
type V2ProvisioningPaymentProfileProviderUsageLimits struct {
	// Three-letter ISO currency code for `max_amount`.
	Currency Currency `json:"currency"`
	// Maximum amount that can be charged per recurring interval.
	MaxAmount int64 `json:"max_amount,string"`
	// Interval over which `max_amount` applies.
	RecurringInterval V2ProvisioningPaymentProfileProviderUsageLimitsRecurringInterval `json:"recurring_interval"`
}

// Providers the payment method is shared with, and their usage limits.
type V2ProvisioningPaymentProfileProvider struct {
	// Provider the payment method is shared with.
	Provider string `json:"provider"`
	// Usage limit applied to the payment method for this provider.
	UsageLimits *V2ProvisioningPaymentProfileProviderUsageLimits `json:"usage_limits,omitempty"`
}

// Usage limit applied to the payment method.
type V2ProvisioningPaymentProfileUsageLimits struct {
	// Three-letter ISO currency code for `max_amount`.
	Currency Currency `json:"currency"`
	// Maximum amount that can be charged per recurring interval.
	MaxAmount int64 `json:"max_amount,string"`
	// Interval over which `max_amount` applies.
	RecurringInterval V2ProvisioningPaymentProfileUsageLimitsRecurringInterval `json:"recurring_interval"`
}

// A customer's payment method and its usage limits.
type V2ProvisioningPaymentProfile struct {
	APIResource
	// Last 4 digits of the card on the payment method.
	CardLast4 string `json:"card_last4"`
	// Whether the payment method is in live mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Owner of the payment method.
	PaymentMethodOwner V2ProvisioningPaymentProfilePaymentMethodOwner `json:"payment_method_owner,omitempty"`
	// Providers the payment method is shared with, and their usage limits.
	Providers []*V2ProvisioningPaymentProfileProvider `json:"providers"`
	// Deprecated: use providers instead.
	SharedWithProviders []string `json:"shared_with_providers"`
	// Usage limit applied to the payment method.
	UsageLimits *V2ProvisioningPaymentProfileUsageLimits `json:"usage_limits,omitempty"`
}
