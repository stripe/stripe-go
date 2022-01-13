//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Transaction type: `adjustment`, `applied_to_invoice`, `credit_note`, `initial`, `invoice_too_large`, `invoice_too_small`, `unspent_receiver_credit`, or `unapplied_from_invoice`. See the [Customer Balance page](https://stripe.com/docs/billing/customer/balance#types) to learn more about transaction types.
type CustomerBalanceTransactionType string

// List of values that CustomerBalanceTransactionType can take
const (
	CustomerBalanceTransactionTypeAdjustment            CustomerBalanceTransactionType = "adjustment"
	CustomerBalanceTransactionTypeAppliedToInvoice      CustomerBalanceTransactionType = "applied_to_invoice"
	CustomerBalanceTransactionTypeCreditNote            CustomerBalanceTransactionType = "credit_note"
	CustomerBalanceTransactionTypeInitial               CustomerBalanceTransactionType = "initial"
	CustomerBalanceTransactionTypeInvoiceTooLarge       CustomerBalanceTransactionType = "invoice_too_large"
	CustomerBalanceTransactionTypeInvoiceTooSmall       CustomerBalanceTransactionType = "invoice_too_small"
	CustomerBalanceTransactionTypeMigration             CustomerBalanceTransactionType = "migration"
	CustomerBalanceTransactionTypeUnappliedFromInvoice  CustomerBalanceTransactionType = "unapplied_from_invoice"
	CustomerBalanceTransactionTypeUnspentReceiverCredit CustomerBalanceTransactionType = "unspent_receiver_credit"
)

// Retrieves a specific customer balance transaction that updated the customer's [balances](https://stripe.com/docs/billing/customer/balance).
type CustomerBalanceTransactionParams struct {
	Params      `form:"*"`
	Customer    *string `form:"-"` // Included in URL
	Amount      *int64  `form:"amount"`
	Currency    *string `form:"currency"`
	Description *string `form:"description"`
}

// Returns a list of transactions that updated the customer's [balances](https://stripe.com/docs/billing/customer/balance).
type CustomerBalanceTransactionListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
}

// Each customer has a [`balance`](https://stripe.com/docs/api/customers/object#customer_object-balance) value,
// which denotes a debit or credit that's automatically applied to their next invoice upon finalization.
// You may modify the value directly by using the [update customer API](https://stripe.com/docs/api/customers/update),
// or by creating a Customer Balance Transaction, which increments or decrements the customer's `balance` by the specified `amount`.
//
// Related guide: [Customer Balance](https://stripe.com/docs/billing/customer/balance) to learn more.
type CustomerBalanceTransaction struct {
	APIResource
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

// CustomerBalanceTransactionList is a list of CustomerBalanceTransactions as retrieved from a list endpoint.
type CustomerBalanceTransactionList struct {
	APIResource
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

	type customerBalanceTransaction CustomerBalanceTransaction
	var v customerBalanceTransaction
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = CustomerBalanceTransaction(v)
	return nil
}
