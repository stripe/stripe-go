//
//
// File generated from our OpenAPI spec
//
//

// Package meterusage provides the /v1/billing/analytics/meter_usage APIs
package meterusage

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/billing/analytics/meter_usage APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Returns aggregated meter usage data for a customer within a specified time interval. The data can be grouped by various dimensions and can include multiple meters if specified.
func Get(params *stripe.BillingAnalyticsMeterUsageParams) (*stripe.BillingAnalyticsMeterUsage, error) {
	return getC().Get(params)
}

// Returns aggregated meter usage data for a customer within a specified time interval. The data can be grouped by various dimensions and can include multiple meters if specified.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(params *stripe.BillingAnalyticsMeterUsageParams) (*stripe.BillingAnalyticsMeterUsage, error) {
	meterusage := &stripe.BillingAnalyticsMeterUsage{}
	err := c.B.Call(
		http.MethodGet, "/v1/billing/analytics/meter_usage", c.Key, params, meterusage)
	return meterusage, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
