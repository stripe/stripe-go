//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of SpendingMandates.
type V2MoneyManagementSpendingMandateListParams struct {
	Params `form:"*"`
	// The ID of the FinancialAccount to filter by.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by bank transfer network. If omitted, returns mandates for all networks.
	Network *string `form:"network" json:"network,omitempty"`
	// Filter by mandate status.
	Statuses []*string `form:"statuses" json:"statuses,omitempty"`
	// The type of SpendingMandate to filter by.
	Type *string `form:"type" json:"type,omitempty"`
}

// ACH authorization details. Required when `network` is `ach`.
type V2MoneyManagementSpendingMandateBankTransferACHParams struct {
	// ACH Company ID identifying the authorized originator. Must contain exactly 10 alphanumeric characters.
	OriginatorCompanyID *string `form:"originator_company_id" json:"originator_company_id,omitempty"`
	// Positive per-debit limit in USD.
	TransactionLimit *Amount `form:"transaction_limit" json:"transaction_limit,omitempty"`
}

// Bank transfer authorization details. Required when `type` is `bank_transfer`.
type V2MoneyManagementSpendingMandateBankTransferParams struct {
	// Display name for the authorized originator, up to 50 characters long.
	AccountHolderName *string `form:"account_holder_name" json:"account_holder_name,omitempty"`
	// ACH authorization details. Required when `network` is `ach`.
	ACH *V2MoneyManagementSpendingMandateBankTransferACHParams `form:"ach" json:"ach,omitempty"`
	// An active US Financial Address belonging to the Financial Account.
	FinancialAddress *string `form:"financial_address" json:"financial_address,omitempty"`
	// The bank transfer network. Only ACH mandates can be created.
	Network *string `form:"network" json:"network,omitempty"`
}

// Creates an active SpendingMandate authorizing an ACH originator to debit a Financial Account.
type V2MoneyManagementSpendingMandateParams struct {
	Params `form:"*"`
	// Bank transfer authorization details. Required when `type` is `bank_transfer`.
	BankTransfer *V2MoneyManagementSpendingMandateBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The Financial Account to authorize debits against.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The mandate type.
	Type *string `form:"type" json:"type,omitempty"`
}

// Cancels an active SpendingMandate.
type V2MoneyManagementSpendingMandateCancelParams struct {
	Params `form:"*"`
}

// ACH authorization details. Required when `network` is `ach`.
type V2MoneyManagementSpendingMandateCreateBankTransferACHParams struct {
	// ACH Company ID identifying the authorized originator. Must contain exactly 10 alphanumeric characters.
	OriginatorCompanyID *string `form:"originator_company_id" json:"originator_company_id"`
	// Positive per-debit limit in USD.
	TransactionLimit *Amount `form:"transaction_limit" json:"transaction_limit,omitempty"`
}

// Bank transfer authorization details. Required when `type` is `bank_transfer`.
type V2MoneyManagementSpendingMandateCreateBankTransferParams struct {
	// Display name for the authorized originator.
	AccountHolderName *string `form:"account_holder_name" json:"account_holder_name"`
	// ACH authorization details. Required when `network` is `ach`.
	ACH *V2MoneyManagementSpendingMandateCreateBankTransferACHParams `form:"ach" json:"ach,omitempty"`
	// An active US Financial Address belonging to the Financial Account.
	FinancialAddress *string `form:"financial_address" json:"financial_address"`
	// The bank transfer network. Only ACH mandates can be created.
	Network *string `form:"network" json:"network"`
}

// Creates an active SpendingMandate authorizing an ACH originator to debit a Financial Account.
type V2MoneyManagementSpendingMandateCreateParams struct {
	Params `form:"*"`
	// Bank transfer authorization details. Required when `type` is `bank_transfer`.
	BankTransfer *V2MoneyManagementSpendingMandateCreateBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The Financial Account to authorize debits against.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// The mandate type.
	Type *string `form:"type" json:"type"`
}

// Retrieves the details of an existing SpendingMandate.
type V2MoneyManagementSpendingMandateRetrieveParams struct {
	Params `form:"*"`
}

// ACH authorization fields to update.
type V2MoneyManagementSpendingMandateUpdateBankTransferACHParams struct {
	// ACH Company ID identifying the authorized originator. Must contain exactly 10 alphanumeric characters.
	OriginatorCompanyID *string `form:"originator_company_id" json:"originator_company_id,omitempty"`
	// Per-debit limit for originator.
	TransactionLimit *Amount `form:"transaction_limit" json:"transaction_limit,omitempty"`
}

// Bank transfer fields to update. Omitted fields remain unchanged.
type V2MoneyManagementSpendingMandateUpdateBankTransferParams struct {
	// Display name for the authorized originator, up to 50 characters long.
	AccountHolderName *string `form:"account_holder_name" json:"account_holder_name,omitempty"`
	// ACH authorization fields to update.
	ACH *V2MoneyManagementSpendingMandateUpdateBankTransferACHParams `form:"ach" json:"ach,omitempty"`
}

// Updates an active ACH SpendingMandate.
type V2MoneyManagementSpendingMandateUpdateParams struct {
	Params `form:"*"`
	// Bank transfer fields to update. Omitted fields remain unchanged.
	BankTransfer *V2MoneyManagementSpendingMandateUpdateBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
}
