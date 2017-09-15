package stripe

import "encoding/json"

// BitcoinTransactionListParams is the set of parameters that can be used when listing BitcoinTransactions.
type BitcoinTransactionListParams struct {
	ListParams `form:"*"`
	Customer   string `form:"customer"`
	Receiver   string `form:"-"` // Sent in with the URL
}

// BitcoinTransactionList is a list object for BitcoinTransactions.
// It is a child object of BitcoinRecievers
// For more details see https://stripe.com/docs/api/#retrieve_bitcoin_receiver
type BitcoinTransactionList struct {
	ListMeta
	Values []*BitcoinTransaction `json:"data"`
}

// BitcoinTransaction is the resource representing a Stripe bitcoin transaction.
// For more details see https://stripe.com/docs/api/#bitcoin_receivers
type BitcoinTransaction struct {
	Amount        uint64   `json:"amount"`
	BitcoinAmount uint64   `json:"bitcoin_amount"`
	Created       int64    `json:"created"`
	Currency      Currency `json:"currency"`
	Customer      string   `json:"customer"`
	ID            string   `json:"id"`
	Receiver      string   `json:"receiver"`
}

// UnmarshalJSON handles deserialization of a BitcoinTransaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (bt *BitcoinTransaction) UnmarshalJSON(data []byte) error {
	type bitcoinTransaction BitcoinTransaction
	var t bitcoinTransaction
	err := json.Unmarshal(data, &t)
	if err == nil {
		*bt = BitcoinTransaction(t)
	} else {
		// the id is surrounded by "\" characters, so strip them
		bt.ID = string(data[1 : len(data)-1])
	}

	return nil
}
