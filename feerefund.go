package stripe

import (
	"encoding/json"
)

// ApplicationFeeRefundParams is the set of parameters that can be used when refunding an application fee.
// For more details see https://stripe.com/docs/api#fee_refund.
type ApplicationFeeRefundParams struct {
	Params         `form:"*"`
	Amount         uint64 `form:"amount"`
	ApplicationFee string `form:"-"` // Included in the URL
}

// ApplicationFeeRefundListParams is the set of parameters that can be used when listing application fee refunds.
// For more details see https://stripe.com/docs/api#list_fee_refunds.
type ApplicationFeeRefundListParams struct {
	ListParams     `form:"*"`
	ApplicationFee string `form:"-"` // Included in the URL
}

// ApplicationFeeRefund is the resource representing a Stripe application fee refund.
// For more details see https://stripe.com/docs/api#fee_refunds.
type ApplicationFeeRefund struct {
	Amount             uint64              `json:"amount"`
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	Created            int64               `json:"created"`
	Currency           Currency            `json:"currency"`
	Fee                string              `json:"fee"`
	ID                 string              `json:"id"`
	Metadata           map[string]string   `json:"metadata"`
}

// ApplicationFeeRefundList is a list object for application fee refunds.
type ApplicationFeeRefundList struct {
	ListMeta
	Data []*ApplicationFeeRefund `json:"data"`
}

// UnmarshalJSON handles deserialization of a ApplicationFeeRefund.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *ApplicationFeeRefund) UnmarshalJSON(data []byte) error {
	type feerefund ApplicationFeeRefund
	var ff feerefund
	err := json.Unmarshal(data, &ff)
	if err == nil {
		*f = ApplicationFeeRefund(ff)
	} else {
		// the id is surrounded by "\" characters, so strip them
		f.ID = string(data[1 : len(data)-1])
	}

	return nil
}
