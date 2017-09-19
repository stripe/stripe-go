package stripe

import "encoding/json"

// ApplicationFeeParams is the set of parameters that can be used when refunding an application fee.
// For more details see https://stripe.com/docs/api#refund_application_fee.
type ApplicationFeeParams struct {
	Params `form:"*"`
}

// ApplicationFeeListParams is the set of parameters that can be used when listing application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
type ApplicationFeeListParams struct {
	ListParams   `form:"*"`
	Charge       string            `form:"charge"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// ApplicationFee is the resource representing a Stripe application fee.
// For more details see https://stripe.com/docs/api#application_fees.
type ApplicationFee struct {
	Account                *Account                  `json:"account"`
	Amount                 uint64                    `json:"amount"`
	AmountRefunded         uint64                    `json:"amount_refunded"`
	Application            string                    `json:"application"`
	BalanceTransaction     *BalanceTransaction       `json:"balance_transaction"`
	Charge                 *Charge                   `json:"charge"`
	Created                int64                     `json:"created"`
	Currency               Currency                  `json:"currency"`
	ID                     string                    `json:"id"`
	Livemode               bool                      `json:"livemode"`
	OriginatingTransaction *Charge                   `json:"originating_transaction"`
	Refunded               bool                      `json:"refunded"`
	Refunds                *ApplicationFeeRefundList `json:"refunds"`
}

//ApplicationFeeList is a list of application fees as retrieved from a list endpoint.
type ApplicationFeeList struct {
	ListMeta
	Data []*ApplicationFee `json:"data"`
}

// UnmarshalJSON handles deserialization of an ApplicationFee.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *ApplicationFee) UnmarshalJSON(data []byte) error {
	type appfee ApplicationFee
	var ff appfee
	err := json.Unmarshal(data, &ff)
	if err == nil {
		*f = ApplicationFee(ff)
	} else {
		// the id is surrounded by "\" characters, so strip them
		f.ID = string(data[1 : len(data)-1])
	}

	return nil
}
