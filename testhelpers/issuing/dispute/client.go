//
//
// File generated from our OpenAPI spec
//
//

// Package dispute provides the /v1/issuing/disputes APIs
package dispute

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/disputes APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Test helper: populates network_lifecycle.pre_arbitration_submission on a test-mode Visa Issuing Dispute using placeholder file tokens. Only supported for Visa disputes.
func SimulateNetworkLifecyclePreArbitrationSubmission(id string, params *stripe.TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationSubmissionParams) (*stripe.IssuingDispute, error) {
	return getC().SimulateNetworkLifecyclePreArbitrationSubmission(id, params)
}

// Test helper: populates network_lifecycle.pre_arbitration_submission on a test-mode Visa Issuing Dispute using placeholder file tokens. Only supported for Visa disputes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SimulateNetworkLifecyclePreArbitrationSubmission(id string, params *stripe.TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationSubmissionParams) (*stripe.IssuingDispute, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/disputes/%s/simulate_network_lifecycle_pre_arbitration_submission", id)
	dispute := &stripe.IssuingDispute{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, dispute)
	return dispute, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
