//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The source balance this transfer came from. One of `card`, `fpx`, or `bank_account`.
type TransferSourceType string

// List of values that TransferSourceType can take
const (
	TransferSourceTypeBankAccount TransferSourceType = "bank_account"
	TransferSourceTypeCard        TransferSourceType = "card"
	TransferSourceTypeFPX         TransferSourceType = "fpx"
)

// To send funds from your Stripe account to a connected account, you create a new transfer object. Your [Stripe balance](https://stripe.com/docs/api#balance) must be able to cover the transfer amount, or you'll receive an “Insufficient Funds” error.
type TransferParams struct {
	Params `form:"*"`
	// A positive integer in cents (or local equivalent) representing how much to transfer.
	Amount *int64 `form:"amount"`
	// 3-letter [ISO code for currency](https://stripe.com/docs/payouts).
	Currency *string `form:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// The ID of a connected Stripe account. [See the Connect documentation](https://stripe.com/docs/connect/charges-transfers) for details.
	Destination *string `form:"destination"`
	// You can use this parameter to transfer funds from a charge before they are added to your available balance. A pending balance will transfer immediately but the funds will not become available until the original charge becomes available. [See the Connect documentation](https://stripe.com/docs/connect/charges-transfers#transfer-availability) for details.
	SourceTransaction *string `form:"source_transaction"`
	// The source balance to use for this transfer. One of `bank_account`, `card`, or `fpx`. For most users, this will default to `card`.
	SourceType *string `form:"source_type"`
	// A string that identifies this transaction as part of a group. See the [Connect documentation](https://stripe.com/docs/connect/charges-transfers#transfer-options) for details.
	TransferGroup *string `form:"transfer_group"`
}

// Returns a list of existing transfers sent to connected accounts. The transfers are returned in sorted order, with the most recently created transfers appearing first.
type TransferListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return transfers for the destination specified by this account ID.
	Destination *string `form:"destination"`
	// Only return transfers with the specified transfer group.
	TransferGroup *string `form:"transfer_group"`
}

// A `Transfer` object is created when you move funds between Stripe accounts as
// part of Connect.
//
// Before April 6, 2017, transfers also represented movement of funds from a
// Stripe account to a card or bank account. This behavior has since been split
// out into a [Payout](https://stripe.com/docs/api#payout_object) object, with corresponding payout endpoints. For more
// information, read about the
// [transfer/payout split](https://stripe.com/docs/transfer-payout-split).
//
// Related guide: [Creating Separate Charges and Transfers](https://stripe.com/docs/connect/charges-transfers).
type Transfer struct {
	APIResource
	// Amount in %s to be transferred.
	Amount int64 `json:"amount"`
	// Amount in %s reversed (can be less than the amount attribute on the transfer if a partial reversal was issued).
	AmountReversed int64 `json:"amount_reversed"`
	// Balance transaction that describes the impact of this transfer on your account balance.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	// Time that this record of the transfer was first created.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string               `json:"description"`
	Destination *TransferDestination `json:"destination"`
	// If the destination is a Stripe account, this will be the ID of the payment that the destination account received for the transfer.
	DestinationPayment *Charge `json:"destination_payment"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// A list of reversals that have been applied to the transfer.
	Reversals *ReversalList `json:"reversals"`
	// Whether the transfer has been fully reversed. If the transfer is only partially reversed, this attribute will still be false.
	Reversed bool `json:"reversed"`
	// ID of the charge or payment that was used to fund the transfer. If null, the transfer was funded from the available balance.
	SourceTransaction *BalanceTransactionSource `json:"source_transaction"`
	// The source balance this transfer came from. One of `card`, `fpx`, or `bank_account`.
	SourceType TransferSourceType `json:"source_type"`
	// A string that identifies this transaction as part of a group. See the [Connect documentation](https://stripe.com/docs/connect/charges-transfers#transfer-options) for details.
	TransferGroup string `json:"transfer_group"`
}
type TransferDestination struct {
	ID string `json:"id"`

	Account *Account `json:"-"`
}

// TransferList is a list of Transfers as retrieved from a list endpoint.
type TransferList struct {
	APIResource
	ListMeta
	Data []*Transfer `json:"data"`
}

// UnmarshalJSON handles deserialization of a Transfer.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *Transfer) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type transfer Transfer
	var v transfer
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = Transfer(v)
	return nil
}

// UnmarshalJSON handles deserialization of a TransferDestination.
// This custom unmarshaling is needed because the specific type of
// TransferDestination it refers to is specified in the JSON
func (t *TransferDestination) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type transferDestination TransferDestination
	var v transferDestination
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = TransferDestination(v)
	err := json.Unmarshal(data, &t.Account)
	return err
}
