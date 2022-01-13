//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Refunds an application fee that has previously been collected but not yet refunded.
// Funds will be refunded to the Stripe account from which the fee was originally collected.
//
// You can optionally refund only part of an application fee.
// You can do so multiple times, until the entire fee has been refunded.
//
// Once entirely refunded, an application fee can't be refunded again.
// This method will raise an error when called on an already-refunded application fee,
// or when trying to refund more money than is left on an application fee.
type FeeRefundParams struct {
	Params         `form:"*"`
	ApplicationFee *string `form:"-"` // Included in URL
	Amount         *int64  `form:"amount"`
}

// You can see a list of the refunds belonging to a specific application fee. Note that the 10 most recent refunds are always available by default on the application fee object. If you need more than those 10, you can use this API method and the limit and starting_after parameters to page through additional refunds.
type FeeRefundListParams struct {
	ListParams     `form:"*"`
	ApplicationFee *string `form:"-"` // Included in URL
}

// `Application Fee Refund` objects allow you to refund an application fee that
// has previously been created but not yet refunded. Funds will be refunded to
// the Stripe account from which the fee was originally collected.
//
// Related guide: [Refunding Application Fees](https://stripe.com/docs/connect/destination-charges#refunding-app-fee).
type FeeRefund struct {
	APIResource
	Amount             int64               `json:"amount"`
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	Created            int64               `json:"created"`
	Currency           Currency            `json:"currency"`
	Fee                *ApplicationFee     `json:"fee"`
	ID                 string              `json:"id"`
	Metadata           map[string]string   `json:"metadata"`
	Object             string              `json:"object"`
}

// FeeRefundList is a list of FeeRefunds as retrieved from a list endpoint.
type FeeRefundList struct {
	APIResource
	ListMeta
	Data []*FeeRefund `json:"data"`
}

// UnmarshalJSON handles deserialization of a FeeRefund.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *FeeRefund) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		f.ID = id
		return nil
	}

	type feeRefund FeeRefund
	var v feeRefund
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*f = FeeRefund(v)
	return nil
}
