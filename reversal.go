//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// When you create a new reversal, you must specify a transfer to create it on.
//
// When reversing transfers, you can optionally reverse part of the transfer. You can do so as many times as you wish until the entire transfer has been reversed.
//
// Once entirely reversed, a transfer can't be reversed again. This method will return an error when called on an already-reversed transfer, or when trying to reverse more money than is left on a transfer.
type ReversalParams struct {
	Params               `form:"*"`
	Transfer             *string `form:"-"` // Included in URL
	Amount               *int64  `form:"amount"`
	Description          *string `form:"description"`
	RefundApplicationFee *bool   `form:"refund_application_fee"`
}

// You can see a list of the reversals belonging to a specific transfer. Note that the 10 most recent reversals are always available by default on the transfer object. If you need more than those 10, you can use this API method and the limit and starting_after parameters to page through additional reversals.
type ReversalListParams struct {
	ListParams `form:"*"`
	Transfer   *string `form:"-"` // Included in URL
}

// [Stripe Connect](https://stripe.com/docs/connect) platforms can reverse transfers made to a
// connected account, either entirely or partially, and can also specify whether
// to refund any related application fees. Transfer reversals add to the
// platform's balance and subtract from the destination account's balance.
//
// Reversing a transfer that was made for a [destination
// charge](https://stripe.com/docs/connect/destination-charges) is allowed only up to the amount of
// the charge. It is possible to reverse a
// [transfer_group](https://stripe.com/docs/connect/charges-transfers#transfer-options)
// transfer only if the destination account has enough balance to cover the
// reversal.
//
// Related guide: [Reversing Transfers](https://stripe.com/docs/connect/charges-transfers#reversing-transfers).
type Reversal struct {
	APIResource
	Amount                   int64               `json:"amount"`
	BalanceTransaction       *BalanceTransaction `json:"balance_transaction"`
	Created                  int64               `json:"created"`
	Currency                 Currency            `json:"currency"`
	Description              string              `json:"description"`
	DestinationPaymentRefund *Refund             `json:"destination_payment_refund"`
	ID                       string              `json:"id"`
	Metadata                 map[string]string   `json:"metadata"`
	Object                   string              `json:"object"`
	SourceRefund             *Refund             `json:"source_refund"`
	Transfer                 string              `json:"transfer"`
}

// ReversalList is a list of Reversals as retrieved from a list endpoint.
type ReversalList struct {
	APIResource
	ListMeta
	Data []*Reversal `json:"data"`
}

// UnmarshalJSON handles deserialization of a Reversal.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *Reversal) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		r.ID = id
		return nil
	}

	type reversal Reversal
	var v reversal
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*r = Reversal(v)
	return nil
}
