//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type ReserveTransaction struct {
	Amount      int64    `json:"amount"`
	Currency    Currency `json:"currency"`
	Description string   `json:"description"`
	ID          string   `json:"id"`
	Object      string   `json:"object"`
}

// UnmarshalJSON handles deserialization of a ReserveTransaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *ReserveTransaction) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		r.ID = id
		return nil
	}

	type reserveTransaction ReserveTransaction
	var v reserveTransaction
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*r = ReserveTransaction(v)
	return nil
}
