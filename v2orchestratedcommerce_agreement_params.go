//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List Agreements for the profile associated with the authenticated merchant.
type V2OrchestratedCommerceAgreementListParams struct {
	Params `form:"*"`
	// The limit for the number of results per page.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter list to Orchestrated Commerce Agreements with a specific counterparty.
	NetworkBusinessProfile *string `form:"network_business_profile" json:"network_business_profile,omitempty"`
}

// Create a new Agreement.
type V2OrchestratedCommerceAgreementParams struct {
	Params `form:"*"`
	// The Network ID of the orchestrator to create an agreement with.
	Orchestrator *string `form:"orchestrator" json:"orchestrator,omitempty"`
}

// Confirm an Agreement.
type V2OrchestratedCommerceAgreementConfirmParams struct {
	Params `form:"*"`
}

// Terminate an Agreement.
type V2OrchestratedCommerceAgreementTerminateParams struct {
	Params `form:"*"`
}

// Create a new Agreement.
type V2OrchestratedCommerceAgreementCreateParams struct {
	Params `form:"*"`
	// The Network ID of the orchestrator to create an agreement with.
	Orchestrator *string `form:"orchestrator" json:"orchestrator"`
}

// Retrieve an Agreement by ID.
type V2OrchestratedCommerceAgreementRetrieveParams struct {
	Params `form:"*"`
}
