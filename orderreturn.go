package stripe

import "encoding/json"

type OrderReturn struct {
	Amount   int64       `json:"amount"`
	Created  int64       `json:"created"`
	Currency Currency    `json:"currency"`
	ID       string      `json:"id"`
	Items    []OrderItem `json:"items"`
	Order    Order       `json:"order"`
	Live     bool        `json:"livemode"`
	Refund   *Refund     `json:"refund"`
}

// OrderReturnList is a list of returns as retrieved from a list endpoint.
type OrderReturnList struct {
	ListMeta
	Values []*OrderReturn `json:"data"`
}

// OrderReturnListParams is the set of parameters that can be used when listing
// returns. For more details, see: https://stripe.com/docs/api#list_order_returns.
type OrderReturnListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Order        string            `form:"order"`
}

// UnmarshalJSON handles deserialization of an OrderReturn.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (ret *OrderReturn) UnmarshalJSON(data []byte) error {
	type orderReturn OrderReturn
	var rr orderReturn
	err := json.Unmarshal(data, &rr)
	if err == nil {
		*ret = OrderReturn(rr)
	} else {
		// the id is surrounded by "\" characters, so strip them
		ret.ID = string(data[1 : len(data)-1])
	}

	return nil
}
