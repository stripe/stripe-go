package stripe

import "encoding/json"

// CustomerBalanceTransactionType is the list of allowed values for the customer's balance
// transaction type.
type CustomerBalanceTransactionType string

// List of values that CustomerBalanceTransactionDuration can take.
const (
	CustomerBalanceTransactionTypeAdjustment            CustomerBalanceTransactionType = "adjustment"
	CustomerBalanceTransactionTypeAppliedToInvoice      CustomerBalanceTransactionType = "applied_to_invoice"
	CustomerBalanceTransactionTypeCreditNote            CustomerBalanceTransactionType = "credit_note"
	CustomerBalanceTransactionTypeInitial               CustomerBalanceTransactionType = "initial"
	CustomerBalanceTransactionTypeInvoiceTooLarge       CustomerBalanceTransactionType = "invoice_too_large"
	CustomerBalanceTransactionTypeInvoiceTooSmall       CustomerBalanceTransactionType = "invoice_too_small"
	CustomerBalanceTransactionTypeUnspentReceiverCredit CustomerBalanceTransactionType = "unspent_receiver_credit"
)

// CustomerBalanceTransactionParams is the set of parameters that can be used when creating or
// updating a customer balance transactions.
// For more details see https://stripe.com/docs/api/customers/create_customer_balance_transaction
type CustomerBalanceTransactionParams struct {
	Params      `form:"*"`
	Amount      *int64  `form:"amount"`
	Customer    *string `form:"-"`
	Currency    *string `form:"currency"`
	Description *string `form:"description"`
}

// CustomerBalanceTransactionListParams is the set of parameters that can be used when listing
// customer balance transactions.
// For more detail see https://stripe.com/docs/api/customers/customer_balance_transactions
type CustomerBalanceTransactionListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"`
}

// CustomerBalanceTransaction is the resource representing a customer balance transaction.
// For more details see https://stripe.com/docs/api/customers/customer_balance_transaction_object
type CustomerBalanceTransaction struct {
	Amount        int64                          `json:"amount"`
	Created       int64                          `json:"created"`
	CreditNote    *CreditNote                    `json:"credit_note"`
	Currency      Currency                       `json:"currency"`
	Customer      *Customer                      `json:"customer"`
	Description   string                         `json:"description"`
	EndingBalance int64                          `json:"ending_balance"`
	ID            string                         `json:"id"`
	Invoice       *Invoice                       `json:"invoice"`
	Livemode      bool                           `json:"livemode"`
	Metadata      map[string]string              `json:"metadata"`
	Object        string                         `json:"object"`
	Type          CustomerBalanceTransactionType `json:"type"`
}

// CustomerBalanceTransactionList is a list of customer balance transactions as retrieved from a
// list endpoint.
type CustomerBalanceTransactionList struct {
	ListMeta
	Data []*CustomerBalanceTransaction `json:"data"`
}

// UnmarshalJSON handles deserialization of a CustomerBalanceTransaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *CustomerBalanceTransaction) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type transaction CustomerBalanceTransaction
	var v transaction
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = CustomerBalanceTransaction(v)
	return nil
}
