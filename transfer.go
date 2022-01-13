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
	Params            `form:"*"`
	Amount            *int64  `form:"amount"`
	Currency          *string `form:"currency"`
	Description       *string `form:"description"`
	Destination       *string `form:"destination"`
	SourceTransaction *string `form:"source_transaction"`
	SourceType        *string `form:"source_type"`
	TransferGroup     *string `form:"transfer_group"`
}

// Returns a list of existing transfers sent to connected accounts. The transfers are returned in sorted order, with the most recently created transfers appearing first.
type TransferListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Destination   *string           `form:"destination"`
	TransferGroup *string           `form:"transfer_group"`
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
	Amount             int64                     `json:"amount"`
	AmountReversed     int64                     `json:"amount_reversed"`
	BalanceTransaction *BalanceTransaction       `json:"balance_transaction"`
	Created            int64                     `json:"created"`
	Currency           Currency                  `json:"currency"`
	Description        string                    `json:"description"`
	Destination        *TransferDestination      `json:"destination"`
	DestinationPayment *Charge                   `json:"destination_payment"`
	ID                 string                    `json:"id"`
	Livemode           bool                      `json:"livemode"`
	Metadata           map[string]string         `json:"metadata"`
	Object             string                    `json:"object"`
	Reversals          *ReversalList             `json:"reversals"`
	Reversed           bool                      `json:"reversed"`
	SourceTransaction  *BalanceTransactionSource `json:"source_transaction"`
	SourceType         TransferSourceType        `json:"source_type"`
	TransferGroup      string                    `json:"transfer_group"`
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
