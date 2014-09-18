package stripe

import "encoding/json"

// FeeParams is the set of parameters that can be used when refunding an application fee.
// For more details see https://stripe.com/docs/api#refund_application_fee.
type FeeParams struct {
	Params
	Amount uint64
}

// FeeListParams is the set of parameters that can be used when listing application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
type FeeListParams struct {
	ListParams
	Created int64
	Charge  string
}

// Fee is the resource representing a Stripe application fee.
// For more details see https://stripe.com/docs/api#application_fees.
type Fee struct {
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

func (f *Fee) UnmarshalJSON(data []byte) error {
	type appfee Fee
	var ff appfee
	err := json.Unmarshal(data, &ff)
	if err == nil {
		*f = Fee(ff)
	} else {
		// the id is surrounded by escaped \, so ignore those
		f.Id = string(data[1 : len(data)-1])
	}

	return nil
}
