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
	Params   `form:"*"`
	Transfer *string `form:"-"` // Included in URL
	// A positive integer in cents (or local equivalent) representing how much of this transfer to reverse. Can only reverse up to the unreversed amount remaining of the transfer. Partial transfer reversals are only allowed for transfers to Stripe Accounts. Defaults to the entire transfer amount.
	Amount *int64 `form:"amount"`
	// An arbitrary string which you can attach to a reversal object. It is displayed alongside the reversal in the Dashboard. This will be unset if you POST an empty value.
	Description *string `form:"description"`
	// Boolean indicating whether the application fee should be refunded when reversing this transfer. If a full transfer reversal is given, the full application fee will be refunded. Otherwise, the application fee will be refunded with an amount proportional to the amount of the transfer reversed.
	RefundApplicationFee *bool `form:"refund_application_fee"`
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
	// Amount, in %s.
	Amount int64 `json:"amount"`
	// Balance transaction that describes the impact on your account balance.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency    Currency `json:"currency"`
	Description string   `json:"description"`
	// Linked payment refund for the transfer reversal.
	DestinationPaymentRefund *Refund `json:"destination_payment_refund"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// ID of the refund responsible for the transfer reversal.
	SourceRefund *Refund `json:"source_refund"`
	// ID of the transfer that was reversed.
	Transfer string `json:"transfer"`
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
