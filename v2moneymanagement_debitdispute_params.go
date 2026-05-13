//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves a list of DebitDisputes.
type V2MoneyManagementDebitDisputeListParams struct {
	Params `form:"*"`
	// Filter by a Financial Account.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by status.
	Status *string `form:"status" json:"status,omitempty"`
}

// Parameters for bank transfer disputes.
type V2MoneyManagementDebitDisputeBankTransferParams struct {
	// The reason for the dispute.
	Reason *string `form:"reason" json:"reason,omitempty"`
}

// Creates a new DebitDispute for a ReceivedDebit.
type V2MoneyManagementDebitDisputeParams struct {
	Params `form:"*"`
	// Parameters for bank transfer disputes.
	BankTransfer *V2MoneyManagementDebitDisputeBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The ID of the ReceivedDebit to dispute.
	ReceivedDebit *string `form:"received_debit" json:"received_debit,omitempty"`
}

// Parameters for bank transfer disputes.
type V2MoneyManagementDebitDisputeCreateBankTransferParams struct {
	// The reason for the dispute.
	Reason *string `form:"reason" json:"reason,omitempty"`
}

// Creates a new DebitDispute for a ReceivedDebit.
type V2MoneyManagementDebitDisputeCreateParams struct {
	Params `form:"*"`
	// Parameters for bank transfer disputes.
	BankTransfer *V2MoneyManagementDebitDisputeCreateBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The ID of the ReceivedDebit to dispute.
	ReceivedDebit *string `form:"received_debit" json:"received_debit"`
}

// Retrieves the details of an existing DebitDispute.
type V2MoneyManagementDebitDisputeRetrieveParams struct {
	Params `form:"*"`
}
