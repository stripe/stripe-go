//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The interval for assessing service. For example, a monthly Rate Card with a rate of $1 for the first 10 "workloads"
// and $2 thereafter means "$1 per workload up to 10 workloads during a month of service." This is similar to but
// distinct from billing interval; the service interval deals with the rate at which the customer accumulates fees,
// while the billing interval in Cadence deals with the rate the customer is billed.
type V2BillingRateCardServiceInterval string

// List of values that V2BillingRateCardServiceInterval can take
const (
	V2BillingRateCardServiceIntervalDay   V2BillingRateCardServiceInterval = "day"
	V2BillingRateCardServiceIntervalMonth V2BillingRateCardServiceInterval = "month"
	V2BillingRateCardServiceIntervalWeek  V2BillingRateCardServiceInterval = "week"
	V2BillingRateCardServiceIntervalYear  V2BillingRateCardServiceInterval = "year"
)

// The Stripe Tax tax behavior - whether the rates are inclusive or exclusive of tax.
type V2BillingRateCardTaxBehavior string

// List of values that V2BillingRateCardTaxBehavior can take
const (
	V2BillingRateCardTaxBehaviorExclusive V2BillingRateCardTaxBehavior = "exclusive"
	V2BillingRateCardTaxBehaviorInclusive V2BillingRateCardTaxBehavior = "inclusive"
)

// A Rate Card represents a versioned set of usage-based prices (rates). Each rate is associated with one Metered Item and
// defines how much to charge for usage of that item. After you've set up a RateCard, you can subscribe customers to it
// by creating a Rate Card Subscription.
type V2BillingRateCard struct {
	APIResource
	// Whether this RateCard is active. Inactive RateCards cannot be used in new activations or have new rates added.
	Active bool `json:"active"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency Currency `json:"currency"`
	// A customer-facing name for the Rate Card.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The ID of this rate card's most recently created version.
	LatestVersion string `json:"latest_version"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The ID of the Rate Card Version that will be used by all subscriptions when no specific version is specified.
	LiveVersion string `json:"live_version"`
	// An internal key you can use to search for a particular RateCard. Maximum length of 200 characters.
	LookupKey string `json:"lookup_key"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The interval for assessing service. For example, a monthly Rate Card with a rate of $1 for the first 10 "workloads"
	// and $2 thereafter means "$1 per workload up to 10 workloads during a month of service." This is similar to but
	// distinct from billing interval; the service interval deals with the rate at which the customer accumulates fees,
	// while the billing interval in Cadence deals with the rate the customer is billed.
	ServiceInterval V2BillingRateCardServiceInterval `json:"service_interval"`
	// The length of the interval for assessing service. For example, set this to 3 and `service_interval` to `"month"` in
	// order to specify quarterly service.
	ServiceIntervalCount int64 `json:"service_interval_count"`
	// The Stripe Tax tax behavior - whether the rates are inclusive or exclusive of tax.
	TaxBehavior V2BillingRateCardTaxBehavior `json:"tax_behavior"`
}
