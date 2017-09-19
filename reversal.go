package stripe

import "encoding/json"

// ReversalParams is the set of parameters that can be used when reversing a transfer.
type ReversalParams struct {
	Params               `form:"*"`
	Amount               uint64 `form:"amount"`
	RefundApplicationFee bool   `form:"refund_application_fee"`
	Transfer             string `form:"-"` // Included in URL
}

// ReversalListParams is the set of parameters that can be used when listing reversals.
type ReversalListParams struct {
	ListParams `form:"*"`
	Transfer   string `form:"-"` // Included in URL
}

// Reversal represents a transfer reversal.
type Reversal struct {
	Amount             uint64              `json:"amount"`
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	Created            int64               `json:"created"`
	Currency           Currency            `json:"currency"`
	ID                 string              `json:"id"`
	Metadata           map[string]string   `json:"metadata"`
	Transfer           string              `json:"transfer"`
}

// ReversalList is a list of object for reversals.
type ReversalList struct {
	ListMeta
	Data []*Reversal `json:"data"`
}

// UnmarshalJSON handles deserialization of a Reversal.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *Reversal) UnmarshalJSON(data []byte) error {
	type reversal Reversal
	var rr reversal
	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Reversal(rr)
	} else {
		// the id is surrounded by "\" characters, so strip them
		r.ID = string(data[1 : len(data)-1])
	}

	return nil
}
