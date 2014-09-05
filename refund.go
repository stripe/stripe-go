package stripe

import "encoding/json"

// RefundParams is the set of parameters that can be used when refunding a charge.
// For more details see https://stripe.com/docs/api#refund.
type RefundParams struct {
	Params
	Charge string
	Amount uint64
	Fee    bool
}

// RefundListParams is the set of parameters that can be used when listing refunds.
// For more details see https://stripe.com/docs/api#list_refunds.
type RefundListParams struct {
	ListParams
	Charge string
}

// Refund is the resource representing a Stripe refund.
// For more details see https://stripe.com/docs/api#refunds.
type Refund struct {
	Id       string            `json:"id"`
	Amount   uint64            `json:"amount"`
	Created  int64             `json:"created"`
	Currency Currency          `json:"currency"`
	Tx       *Transaction      `json:"balance_transaction"`
	Charge   string            `json:"charge"`
	Meta     map[string]string `json:"metadata"`
}

// Refundist is a list object for refunds.
type RefundList struct {
	ListResponse
	Values []*Refund `json:"data"`
}

func (r *Refund) UnmarshalJSON(data []byte) error {
	type refund Refund
	var rr refund
	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Refund(rr)
	} else {
		// the id is surrounded by escaped \, so ignore those
		r.Id = string(data[1 : len(data)-1])
	}

	return nil
}
