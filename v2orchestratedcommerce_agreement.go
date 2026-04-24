//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The party that initiated the agreement.
type V2OrchestratedCommerceAgreementInitiatedBy string

// List of values that V2OrchestratedCommerceAgreementInitiatedBy can take
const (
	V2OrchestratedCommerceAgreementInitiatedByOrchestrator V2OrchestratedCommerceAgreementInitiatedBy = "orchestrator"
	V2OrchestratedCommerceAgreementInitiatedBySeller       V2OrchestratedCommerceAgreementInitiatedBy = "seller"
)

// The current status of the agreement.
type V2OrchestratedCommerceAgreementStatus string

// List of values that V2OrchestratedCommerceAgreementStatus can take
const (
	V2OrchestratedCommerceAgreementStatusConfirmed          V2OrchestratedCommerceAgreementStatus = "confirmed"
	V2OrchestratedCommerceAgreementStatusInitiated          V2OrchestratedCommerceAgreementStatus = "initiated"
	V2OrchestratedCommerceAgreementStatusPartiallyConfirmed V2OrchestratedCommerceAgreementStatus = "partially_confirmed"
	V2OrchestratedCommerceAgreementStatusTerminated         V2OrchestratedCommerceAgreementStatus = "terminated"
)

// The party that terminated the agreement, if applicable.
type V2OrchestratedCommerceAgreementTerminatedBy string

// List of values that V2OrchestratedCommerceAgreementTerminatedBy can take
const (
	V2OrchestratedCommerceAgreementTerminatedByOrchestrator V2OrchestratedCommerceAgreementTerminatedBy = "orchestrator"
	V2OrchestratedCommerceAgreementTerminatedBySeller       V2OrchestratedCommerceAgreementTerminatedBy = "seller"
)

// Details about the orchestrator.
type V2OrchestratedCommerceAgreementOrchestratorDetails struct {
	// The name of the orchestrator. This can be the name of the agent or the name of the business.
	Name string `json:"name"`
	// The Network ID of the orchestrator.
	NetworkBusinessProfile string `json:"network_business_profile"`
}

// Details about the seller.
type V2OrchestratedCommerceAgreementSellerDetails struct {
	// The Network ID of the seller.
	NetworkBusinessProfile string `json:"network_business_profile"`
}

// Timestamps of key status transitions for the agreement.
type V2OrchestratedCommerceAgreementStatusTransitions struct {
	// The time at which the orchestrator confirmed the agreement.
	OrchestratorConfirmedAt time.Time `json:"orchestrator_confirmed_at,omitempty"`
	// The time at which the seller confirmed the agreement.
	SellerConfirmedAt time.Time `json:"seller_confirmed_at,omitempty"`
	// The time at which the agreement was terminated.
	TerminatedAt time.Time `json:"terminated_at,omitempty"`
}

// An Orchestrated Commerce Agreement represents a mutual agreement between a seller and an orchestrator/agent on the Stripe network.
type V2OrchestratedCommerceAgreement struct {
	APIResource
	// The time at which the agreement was created.
	Created time.Time `json:"created"`
	// The unique identifier for the agreement.
	ID string `json:"id"`
	// The party that initiated the agreement.
	InitiatedBy V2OrchestratedCommerceAgreementInitiatedBy `json:"initiated_by"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Details about the orchestrator.
	OrchestratorDetails *V2OrchestratedCommerceAgreementOrchestratorDetails `json:"orchestrator_details"`
	// Details about the seller.
	SellerDetails *V2OrchestratedCommerceAgreementSellerDetails `json:"seller_details"`
	// The current status of the agreement.
	Status V2OrchestratedCommerceAgreementStatus `json:"status"`
	// Timestamps of key status transitions for the agreement.
	StatusTransitions *V2OrchestratedCommerceAgreementStatusTransitions `json:"status_transitions"`
	// The party that terminated the agreement, if applicable.
	TerminatedBy V2OrchestratedCommerceAgreementTerminatedBy `json:"terminated_by,omitempty"`
}
