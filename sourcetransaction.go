//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// SourceTransactionListParams is the set of parameters that can be used when listing SourceTransactions.
type SourceTransactionListParams struct {
	ListParams `form:"*"`
	Source     *string `form:"-"` // Included in URL
}

// SourceTransaction is the resource representing a Stripe source transaction.
type SourceTransaction struct {
	Amount   int64    `json:"amount"`
	Created  int64    `json:"created"`
	Currency Currency `json:"currency"`
	ID       string   `json:"id"`
	Livemode bool     `json:"livemode"`
	Object   string   `json:"object"`
	Source   string   `json:"source"`
	Status   string   `json:"status"`
	Type     string   `json:"type"`

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

// SourceTransactionList is a list object for SourceTransactions.
type SourceTransactionList struct {
	APIResource
	ListMeta
	Data []*SourceTransaction `json:"data"`
}
