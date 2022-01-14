//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// List source transactions for a given source.
type SourceTransactionListParams struct {
	ListParams `form:"*"`
	Source     *string `form:"-"` // Included in URL
}

// Some payment methods have no required amount that a customer must send.
// Customers can be instructed to send any amount, and it can be made up of
// multiple transactions. As such, sources can have multiple associated
// transactions.
type SourceTransaction struct {
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the amount your customer has pushed to the receiver.
	Amount int64 `json:"amount"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The ID of the source this transaction is attached to.
	Source string `json:"source"`
	// The status of the transaction, one of `succeeded`, `pending`, or `failed`.
	Status string `json:"status"`
	// The type of source this transaction is attached to.
	Type string `json:"type"`

	// See custom UnmarshalJSON
	TypeData map[string]interface{}

	// Deprecated
	CustomerData string `json:"customer_data"`
}

// UnmarshalJSON handles deserialization of a SourceTransaction. This custom
// unmarshaling is needed to extract the type specific data (accessible under
// `TypeData`) but stored in JSON under a hash named after the `type` of the source transaction.
func (t *SourceTransaction) UnmarshalJSON(data []byte) error {
	type sourceTransaction SourceTransaction
	var v sourceTransaction
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*t = SourceTransaction(v)

	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	if d, ok := raw[t.Type]; ok {
		if m, ok := d.(map[string]interface{}); ok {
			t.TypeData = m
		}
	}

	return nil
}

// SourceTransactionList is a list of SourceTransactions as retrieved from a list endpoint.
type SourceTransactionList struct {
	APIResource
	ListMeta
	Data []*SourceTransaction `json:"data"`
}
