package stripe

import (
	"encoding/json"
)

// BitcoinReceiverListParams is the set of parameters that can be used when listing BitcoinReceivers.
// For more details see https://stripe.com/docs/api/#list_bitcoin_receivers.
type BitcoinReceiverListParams struct {
	ListParams `form:"*"`
	NotActive  bool `form:"active,invert"`
	NotFilled  bool `form:"filled,invert"`
	Uncaptured bool `form:"uncaptured_funds"`
}

// BitcoinReceiverParams is the set of parameters that can be used when creating a BitcoinReceiver.
// For more details see https://stripe.com/docs/api/#create_bitcoin_receiver.
type BitcoinReceiverParams struct {
	Params   `form:"*"`
	Amount   uint64   `form:"amount"`
	Currency Currency `form:"currency"`
	Desc     string   `form:"description"`
	Email    string   `form:"email"`
}

// BitcoinReceiverUpdateParams is the set of parameters that can be used when
// updating a BitcoinReceiver. For more details see
// https://stripe.com/docs/api/#update_bitcoin_receiver.
type BitcoinReceiverUpdateParams struct {
	Params     `form:"*"`
	Desc       string `form:"description"`
	Email      string `form:"email"`
	RefundAddr string `form:"refund_address"`
}

// BitcoinReceiver is the resource representing a Stripe bitcoin receiver.
// For more details see https://stripe.com/docs/api/#bitcoin_receivers
type BitcoinReceiver struct {
	Active                bool                    `json:"active"`
	Amount                uint64                  `json:"amount"`
	AmountReceived        uint64                  `json:"amount_received"`
	BitcoinAmount         uint64                  `json:"bitcoin_amount"`
	BitcoinAmountReceived uint64                  `json:"bitcoin_amount_received"`
	BitcoinUri            string                  `json:"bitcoin_uri"`
	Created               int64                   `json:"created"`
	Currency              Currency                `json:"currency"`
	Customer              string                  `json:"customer"`
	Desc                  string                  `json:"description"`
	Email                 string                  `json:"email"`
	Filled                bool                    `json:"filled"`
	ID                    string                  `json:"id"`
	InboundAddress        string                  `json:"inbound_address"`
	Meta                  map[string]string       `json:"metadata"`
	Payment               string                  `json:"payment"`
	RefundAddress         string                  `json:"refund_address"`
	RejectTransactions    bool                    `json:"reject_transactions"`
	Transactions          *BitcoinTransactionList `json:"transactions"`
}

// BitcoinReceiverList is a list of bitcoin receivers as retrieved from a list endpoint.
type BitcoinReceiverList struct {
	ListMeta
	Values []*BitcoinReceiver `json:"data"`
}

// UnmarshalJSON handles deserialization of a BitcoinReceiver.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (br *BitcoinReceiver) UnmarshalJSON(data []byte) error {
	type bitcoinReceiver BitcoinReceiver
	var r bitcoinReceiver
	err := json.Unmarshal(data, &r)
	if err == nil {
		*br = BitcoinReceiver(r)
	} else {
		// the id is surrounded by "\" characters, so strip them
		br.ID = string(data[1 : len(data)-1])
	}

	return nil
}
