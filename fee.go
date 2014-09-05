package stripe

import "encoding/json"

// AppFeeParams is the set of parameters that can be used when refunding an application fee.
// For more details see https://stripe.com/docs/api#refund_application_fee.
type AppFeeParams struct {
	Params
	Amount uint64
}

// AppFeeListParams is the set of parameters that can be used when listing application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
type AppFeeListParams struct {
	ListParams
	Created int64
	Charge  string
}

// AppFee is the resource representing a Stripe application fee.
// For more details see https://stripe.com/docs/api#application_fees.
type AppFee struct {
	Id             string         `json:"id"`
	Live           bool           `json:"livemode"`
	Account        *Account       `json:"account"`
	Amount         uint64         `json:"amount"`
	App            string         `json:"application"`
	Tx             *Transaction   `json:"balance_transaction"`
	Charge         *Charge        `json:"charge"`
	Created        int64          `json:"created"`
	Currency       Currency       `json:"currency"`
	Refunded       bool           `json:"refunded"`
	Refunds        *FeeRefundList `json:"refunds"`
	AmountRefunded uint64         `json:"amount_refunded"`
}

// AppFeeList is a list object for application fees.
type AppFeeList struct {
	ListResponse
	Values []*AppFee `json:"data"`
}

func (f *AppFee) UnmarshalJSON(data []byte) error {
	type appfee AppFee
	var ff appfee
	err := json.Unmarshal(data, &ff)
	if err == nil {
		*f = AppFee(ff)
	} else {
		// the id is surrounded by escaped \, so ignore those
		f.Id = string(data[1 : len(data)-1])
	}

	return nil
}
