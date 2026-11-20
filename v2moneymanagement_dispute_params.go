//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves a list of Disputes.
type V2MoneyManagementDisputeListParams struct {
	Params `form:"*"`
	// Filter results by the ID of the FinancialAccount associated with the dispute.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Maximum number of results to return, between 1 and 100. Defaults to 10.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter results to Disputes with any of the provided statuses.
	Statuses []*string `form:"statuses" json:"statuses,omitempty"`
	// Filter results by the ID of the disputed money movement.
	Target *string `form:"target" json:"target,omitempty"`
}

// The money movement being disputed.
type V2MoneyManagementDisputeTargetParams struct {
	// If applicable, the ID of the ReceivedDebit being disputed.
	ReceivedDebit *string `form:"received_debit" json:"received_debit,omitempty"`
	// The type of money movement being disputed.
	Type *string `form:"type" json:"type"`
}

// Creates a Dispute for the specified money movement.
type V2MoneyManagementDisputeParams struct {
	Params `form:"*"`
	// Additional information explaining the dispute. Must contain non-whitespace text and be no longer than 5000 characters.
	Explanation *string `form:"explanation" json:"explanation,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The FinancialAccount holder's reason for disputing the money movement.
	Reason *string `form:"reason" json:"reason,omitempty"`
	// The money movement being disputed.
	Target *V2MoneyManagementDisputeTargetParams `form:"target" json:"target,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementDisputeParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The money movement being disputed.
type V2MoneyManagementDisputeCreateTargetParams struct {
	// If applicable, the ID of the ReceivedDebit being disputed.
	ReceivedDebit *string `form:"received_debit" json:"received_debit,omitempty"`
	// The type of money movement being disputed.
	Type *string `form:"type" json:"type"`
}

// Creates a Dispute for the specified money movement.
type V2MoneyManagementDisputeCreateParams struct {
	Params `form:"*"`
	// Additional information explaining the dispute. Must contain non-whitespace text and be no longer than 5000 characters.
	Explanation *string `form:"explanation" json:"explanation,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The FinancialAccount holder's reason for disputing the money movement.
	Reason *string `form:"reason" json:"reason"`
	// The money movement being disputed.
	Target *V2MoneyManagementDisputeCreateTargetParams `form:"target" json:"target"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2MoneyManagementDisputeCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves a Dispute.
type V2MoneyManagementDisputeRetrieveParams struct {
	Params `form:"*"`
}
