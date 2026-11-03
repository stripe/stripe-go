//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves a list of InboundTransfers.
type V2MoneyManagementInboundTransferListParams struct {
	Params `form:"*"`
	// Hash of options for filtering on creation time.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// The verified bank account credential to debit for this InboundTransfer.
type V2MoneyManagementInboundTransferFromParams struct {
	// An optional currency field used to specify which currency is debited from the Payment Method.
	// Since many Payment Methods support only one currency, this field is optional.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// ID of the verified bank account credential to debit for this InboundTransfer (for example, `usba_`), not a Payments API `pm_` PaymentMethod.
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
}

// The FinancialAccount and currency to credit with the transferred funds.
type V2MoneyManagementInboundTransferToParams struct {
	// The currency to credit to the FinancialAccount.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that receives the funds.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Creates an InboundTransfer to pull funds from a bank account into a financial account. Per-transfer and rolling weekly limits vary by account and payment rail. See [transfer limits](https://docs.stripe.com/treasury/add-funds#transfer-limits).
type V2MoneyManagementInboundTransferParams struct {
	Params `form:"*"`
	// The amount and currency to credit to the FinancialAccount.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// An optional, freeform description field intended to store metadata.
	Description *string `form:"description" json:"description,omitempty"`
	// The verified bank account credential to debit for this InboundTransfer.
	From *V2MoneyManagementInboundTransferFromParams `form:"from" json:"from,omitempty"`
	// The FinancialAccount and currency to credit with the transferred funds.
	To *V2MoneyManagementInboundTransferToParams `form:"to" json:"to,omitempty"`
}

// The verified bank account credential to debit for this InboundTransfer.
type V2MoneyManagementInboundTransferCreateFromParams struct {
	// An optional currency field used to specify which currency is debited from the Payment Method.
	// Since many Payment Methods support only one currency, this field is optional.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// ID of the verified bank account credential to debit for this InboundTransfer (for example, `usba_`), not a Payments API `pm_` PaymentMethod.
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
}

// The FinancialAccount and currency to credit with the transferred funds.
type V2MoneyManagementInboundTransferCreateToParams struct {
	// The currency to credit to the FinancialAccount.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that receives the funds.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Creates an InboundTransfer to pull funds from a bank account into a financial account. Per-transfer and rolling weekly limits vary by account and payment rail. See [transfer limits](https://docs.stripe.com/treasury/add-funds#transfer-limits).
type V2MoneyManagementInboundTransferCreateParams struct {
	Params `form:"*"`
	// The amount and currency to credit to the FinancialAccount.
	Amount *Amount `form:"amount" json:"amount"`
	// An optional, freeform description field intended to store metadata.
	Description *string `form:"description" json:"description,omitempty"`
	// The verified bank account credential to debit for this InboundTransfer.
	From *V2MoneyManagementInboundTransferCreateFromParams `form:"from" json:"from"`
	// The FinancialAccount and currency to credit with the transferred funds.
	To *V2MoneyManagementInboundTransferCreateToParams `form:"to" json:"to"`
}

// Retrieve an InboundTransfer by ID.
type V2MoneyManagementInboundTransferRetrieveParams struct {
	Params `form:"*"`
}
