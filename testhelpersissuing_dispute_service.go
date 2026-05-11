//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1TestHelpersIssuingDisputeService is used to invoke /v1/issuing/disputes APIs.
type v1TestHelpersIssuingDisputeService struct {
	B   Backend
	Key string
}

// Test helper: populates network_lifecycle.pre_arbitration_submission on a test-mode Visa Issuing Dispute using placeholder file tokens. Only supported for Visa disputes.
func (c v1TestHelpersIssuingDisputeService) SimulateNetworkLifecyclePreArbitrationSubmission(ctx context.Context, id string, params *TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationSubmissionParams) (*IssuingDispute, error) {
	if params == nil {
		params = &TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationSubmissionParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/issuing/disputes/%s/simulate_network_lifecycle_pre_arbitration_submission", id)
	dispute := &IssuingDispute{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, dispute)
	return dispute, err
}
