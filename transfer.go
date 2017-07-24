package stripe

import "encoding/json"

// TransferSourceType is the list of allowed values for the transfer's source_type field.
// Allowed values are "alipay_account", bank_account", "bitcoin_receiver", "card".
type TransferSourceType string

// TransferDestination describes the destination of a Transfer.
// The Type should indicate which object is fleshed out
// For more details see https://stripe.com/docs/api/go#transfer_object
type TransferDestination struct {
	ID      string   `json:"id"`
	Account *Account `json:"-"`
}

// TransferParams is the set of parameters that can be used when creating or updating a transfer.
// For more details see https://stripe.com/docs/api#create_transfer and https://stripe.com/docs/api#update_transfer.
type TransferParams struct {
	Params
	Amount        int64
	Currency      Currency
	Dest          string
	Meta          map[string]string
	SourceTx      string
	SourceType    TransferSourceType
	TransferGroup string
}

// TransferListParams is the set of parameters that can be used when listing transfers.
// For more details see https://stripe.com/docs/api#list_transfers.
type TransferListParams struct {
	ListParams
	Amount        int64
	Created       int64
	CreatedRange  *RangeQueryParams
	Currency      Currency
	Dest          string
	TransferGroup string
}

// Transfer is the resource representing a Stripe transfer.
// For more details see https://stripe.com/docs/api#transfers.
type Transfer struct {
	ID             string              `json:"id"`
	Live           bool                `json:"livemode"`
	Amount         int64               `json:"amount"`
	AmountReversed int64               `json:"amount_reversed"`
	Currency       Currency            `json:"currency"`
	Created        int64               `json:"created"`
	Dest           TransferDestination `json:"destination"`
	Tx             *Transaction        `json:"balance_transaction"`
	Meta           map[string]string   `json:"metadata"`
	Statement      string              `json:"statement_descriptor"`
	Reversals      *ReversalList       `json:"reversals"`
	Reversed       bool                `json:"reversed"`
	SourceTx       *TransactionSource  `json:"source_transaction"`
	DestPayment    string              `json:"destination_payment"`
	TransferGroup  string              `json:"transfer_group"`
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

		json.Unmarshal(data, &d.Account)
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
