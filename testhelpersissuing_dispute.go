//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Test helper: closes a test-mode Issuing dispute as won or lost.
type TestHelpersIssuingDisputeCloseParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Whether to close the dispute as `won` or `lost`.
	Status *string `form:"status" json:"status"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingDisputeCloseParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Test helper: overrides the grant_deadline and revocable_after timestamps on a test-mode Issuing dispute's provisional credit, allowing tests to simulate timer-driven status transitions without waiting for real regulatory deadlines to pass.
type TestHelpersIssuingDisputeProvisionalCreditParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Override the deadline by which the platform must grant a provisional credit to the consumer.
	GrantDeadline *int64 `form:"grant_deadline" json:"grant_deadline,omitempty"`
	// Override the earliest time after which the platform may revoke the provisional credit.
	RevocableAfter *int64 `form:"revocable_after" json:"revocable_after,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingDisputeProvisionalCreditParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Controls the acquiring merchant's simulated submitted evidence files for the dispute response stage.
type TestHelpersIssuingDisputeSimulateNetworkLifecycleDisputeResponseMerchantEvidenceFilesParams struct {
	// How many simulated merchant evidence file tokens to attach (between 1 and 12).
	NumberToGenerate *int64 `form:"number_to_generate" json:"number_to_generate"`
}

// Test helper: populates network_lifecycle.dispute_response on a test-mode Visa Issuing Dispute using placeholder file tokens. Only supported for Visa disputes.
type TestHelpersIssuingDisputeSimulateNetworkLifecycleDisputeResponseParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Controls the acquiring merchant's simulated submitted evidence files for the dispute response stage.
	MerchantEvidenceFiles *TestHelpersIssuingDisputeSimulateNetworkLifecycleDisputeResponseMerchantEvidenceFilesParams `form:"merchant_evidence_files" json:"merchant_evidence_files"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingDisputeSimulateNetworkLifecycleDisputeResponseParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Controls the acquiring merchant's simulated submitted evidence files for the pre-arbitration response stage.
type TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationResponseMerchantEvidenceFilesParams struct {
	// How many simulated merchant evidence file tokens to attach (between 1 and 12).
	NumberToGenerate *int64 `form:"number_to_generate" json:"number_to_generate"`
}

// Test helper: populates network_lifecycle.pre_arbitration_response on a test-mode Visa Issuing Dispute using placeholder file tokens. Only supported for Visa disputes in the collaboration flow.
type TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationResponseParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Controls the acquiring merchant's simulated submitted evidence files for the pre-arbitration response stage.
	MerchantEvidenceFiles *TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationResponseMerchantEvidenceFilesParams `form:"merchant_evidence_files" json:"merchant_evidence_files"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationResponseParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Controls the acquiring merchant's simulated submitted evidence files for the pre-arbitration submission stage.
type TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationSubmissionMerchantEvidenceFilesParams struct {
	// How many simulated merchant evidence file tokens to attach (between 1 and 12).
	NumberToGenerate *int64 `form:"number_to_generate" json:"number_to_generate"`
}

// Test helper: populates network_lifecycle.pre_arbitration_submission on a test-mode Visa Issuing Dispute using placeholder file tokens. Only supported for Visa disputes.
type TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationSubmissionParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Controls the acquiring merchant's simulated submitted evidence files for the pre-arbitration submission stage.
	MerchantEvidenceFiles *TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationSubmissionMerchantEvidenceFilesParams `form:"merchant_evidence_files" json:"merchant_evidence_files"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingDisputeSimulateNetworkLifecyclePreArbitrationSubmissionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
