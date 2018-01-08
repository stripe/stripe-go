package stripe

import "encoding/json"

// TransferSourceType is the list of allowed values for the transfer's source_type field.
// Allowed values are "alipay_account", bank_account", "bitcoin_receiver", "card".
type TransferSourceType string

// TransferDestination describes the destination of a Transfer.
// The Type should indicate which object is fleshed out
// For more details see https://stripe.com/docs/api/go#transfer_object
type TransferDestination struct {
	Account *Account `json:"-"`
	ID      string   `json:"id"`
}

// TransferParams is the set of parameters that can be used when creating or updating a transfer.
// For more details see https://stripe.com/docs/api#create_transfer and https://stripe.com/docs/api#update_transfer.
type TransferParams struct {
	Params        `form:"*"`
	Amount        int64              `form:"amount"`
	Currency      Currency           `form:"currency"`
	Dest          string             `form:"destination"`
	SourceTx      string             `form:"source_transaction"`
	SourceType    TransferSourceType `form:"source_type"`
	TransferGroup string             `form:"transfer_group"`
}

// TransferListParams is the set of parameters that can be used when listing transfers.
// For more details see https://stripe.com/docs/api#list_transfers.
type TransferListParams struct {
	ListParams    `form:"*"`
	Created       int64             `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Currency      Currency          `form:"currency"`
	Dest          string            `form:"destination"`
	TransferGroup string            `form:"transfer_group"`
}

// Transfer is the resource representing a Stripe transfer.
// For more details see https://stripe.com/docs/api#transfers.
type Transfer struct {
	Amount         int64               `json:"amount"`
	AmountReversed int64               `json:"amount_reversed"`
	Created        int64               `json:"created"`
	Currency       Currency            `json:"currency"`
	Dest           TransferDestination `json:"destination"`
	DestPayment    string              `json:"destination_payment"`
	ID             string              `json:"id"`
	Live           bool                `json:"livemode"`
	Meta           map[string]string   `json:"metadata"`
	Reversals      *ReversalList       `json:"reversals"`
	Reversed       bool                `json:"reversed"`
	SourceTx       *TransactionSource  `json:"source_transaction"`
	Statement      string              `json:"statement_descriptor"`
	TransferGroup  string              `json:"transfer_group"`
	Tx             *Transaction        `json:"balance_transaction"`
}

// TransferList is a list of transfers as retrieved from a list endpoint.
type TransferList struct {
	ListMeta
	Values []*Transfer `json:"data"`
}

// UnmarshalJSON handles deserialization of a Transfer.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *Transfer) UnmarshalJSON(data []byte) error {
	type transfer Transfer
	var tb transfer
	err := json.Unmarshal(data, &tb)
	if err == nil {
		*t = Transfer(tb)
	} else {
		// the id is surrounded by "\" characters, so strip them
		t.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// UnmarshalJSON handles deserialization of a TransferDestination.
// This custom unmarshaling is needed because the specific
// type of destination it refers to is specified in the JSON
func (d *TransferDestination) UnmarshalJSON(data []byte) error {
	type dest TransferDestination
	var dd dest
	err := json.Unmarshal(data, &dd)
	if err == nil {
		*d = TransferDestination(dd)

		err = json.Unmarshal(data, &d.Account)

		if err != nil {
			return err
		}
	} else {
		// the id is surrounded by "\" characters, so strip them
		d.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// MarshalJSON handles serialization of a TransferDestination.
// This custom marshaling is needed because we can only send a string
// ID as a destination, even though it can be expanded to a full
// object when retrieving
func (d *TransferDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ID)
}
