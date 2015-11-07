package bitcointransaction

import (
	"testing"

	stripe "github.com/seenickcode/stripe-go"
	"github.com/seenickcode/stripe-go/bitcoinreceiver"
	"github.com/seenickcode/stripe-go/currency"
	. "github.com/seenickcode/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestBitcoinTransactionList(t *testing.T) {
	bitcoinReceiverParams := &stripe.BitcoinReceiverParams{
		Amount:   1000,
		Currency: currency.USD,
		Email:    "do+fill_now@stripe.com",
		Desc:     "some details",
	}

	r, _ := bitcoinreceiver.New(bitcoinReceiverParams)

	params := &stripe.BitcoinTransactionListParams{
		Receiver: r.ID,
	}
	params.Filters.AddFilter("limit", "", "5")
	params.Single = true

	i := List(params)
	for i.Next() {
		if i.BitcoinTransaction() == nil {
			t.Error("No nil values expected")
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
	}
	if err := i.Err(); err != nil {
		t.Error(err)
	}
}
