package transfer

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/account"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/currency"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestTransferAllMethods(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Source: &stripe.SourceParams{
			Card: &stripe.CardParams{
				Number: "4000000000000077",
				Month:  "06",
				Year:   "20",
			},
		},
	}

	charge, err := charge.New(chargeParams)
	if err != nil {
		t.Error(err)
	}

	params := &stripe.AccountParams{
		Managed: true,
		Country: "US",
		LegalEntity: &stripe.LegalEntity{
			Type: stripe.Individual,
			DOB: stripe.DOB{
				Day:   1,
				Month: 2,
				Year:  1990,
			},
		},
	}

	acc, err := account.New(params)
	if err != nil {
		t.Error(err)
	}

	transferParams := &stripe.TransferParams{
		Amount:   100,
		Currency: currency.USD,
		Dest:     acc.ID,
		SourceTx: charge.ID,
	}

	target, err := New(transferParams)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != transferParams.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, transferParams.Amount)
	}

	if target.Currency != transferParams.Currency {
		t.Errorf("Curency %q does not match expected currency %q\n", target.Currency, transferParams.Currency)
	}

	if target.Created == 0 {
		t.Errorf("Created date is not set\n")
	}

	transferRetrieved, err := Get(target.ID, nil)
	if err != nil {
		t.Error(err)
	}

	if transferRetrieved.ID != target.ID {
		t.Errorf("ID %q does not match expected ID %q\n", transferRetrieved.ID, target.ID)
	}

	updateParams := &stripe.TransferParams{}
	updateParams.AddMeta("foo", "bar")

	transferUpdated, err := Update(target.ID, updateParams)
	if err != nil {
		t.Error(err)
	}

	if transferUpdated.Meta["foo"] != "bar" {
		t.Error("Transfer metadata not updated")
	}

	multipleTransferParams := &stripe.TransferParams{
		Amount:   100,
		Currency: currency.USD,
		Dest:     acc.ID,
	}

	for i := 0; i < 3; i++ {
		New(multipleTransferParams)
	}

	nbTransfers := 0
	i := List(&stripe.TransferListParams{Dest: acc.ID})
	for i.Next() {
		if i.Transfer() == nil {
			t.Error("No nil values expected")
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
		nbTransfers++
	}
	if err := i.Err(); err != nil {
		t.Error(err)
	}
	if nbTransfers != 4 {
		t.Error("Expected 4 transfers on %q but got %q\n", acc.ID, nbTransfers)
	}

	account.Del(acc.ID)
}
