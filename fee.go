//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Returns a list of application fees you've previously collected. The application fees are returned in sorted order, with the most recent fees appearing first.
type ApplicationFeeListParams struct {
	ListParams   `form:"*"`
	Charge       *string           `form:"charge"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// Retrieves the details of an application fee that your account has collected. The same information is returned when refunding the application fee.
type ApplicationFeeParams struct {
	Params `form:"*"`
}
type ApplicationFee struct {
	APIResource
	Account                *Account            `json:"account"`
	Amount                 int64               `json:"amount"`
	AmountRefunded         int64               `json:"amount_refunded"`
	Application            string              `json:"application"`
	BalanceTransaction     *BalanceTransaction `json:"balance_transaction"`
	Charge                 *Charge             `json:"charge"`
	Created                int64               `json:"created"`
	Currency               Currency            `json:"currency"`
	ID                     string              `json:"id"`
	Livemode               bool                `json:"livemode"`
	Object                 string              `json:"object"`
	OriginatingTransaction *Charge             `json:"originating_transaction"`
	Refunded               bool                `json:"refunded"`
	Refunds                *FeeRefundList      `json:"refunds"`
}

// ApplicationFeeList is a list of ApplicationFees as retrieved from a list endpoint.
type ApplicationFeeList struct {
	APIResource
	ListMeta
	Data []*ApplicationFee `json:"data"`
}

// UnmarshalJSON handles deserialization of an ApplicationFee.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (a *ApplicationFee) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		a.ID = id
		return nil
	}

	type applicationFee ApplicationFee
	var v applicationFee
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*a = ApplicationFee(v)
	return nil
}
