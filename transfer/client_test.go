package transfer

import (
	"testing"

	. "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/recipient"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	Key = GetTestKey()
}

func TestTransferNew(t *testing.T) {
	recipientParams := &RecipientParams{
		Name: "Recipient Name",
		Type: Individual,
		Bank: &BankAccountParams{
			Country: "US",
			Routing: "110000000",
			Account: "000123456789",
		},
	}

	rec, _ := recipient.New(recipientParams)

	transferParams := &TransferParams{
		Amount:    100,
		Currency:  USD,
		Recipient: rec.Id,
		Desc:      "Transfer Desc",
		Statement: "Transfer",
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

	if target.Date == 0 {
		t.Errorf("Date is not set \n")
	}

	if target.Desc != transferParams.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, transferParams.Desc)
	}

	if target.Recipient.Id != transferParams.Recipient {
		t.Errorf("Recipient %q does not match expected recipient %q\n", target.Recipient.Id, transferParams.Recipient)
	}

	if target.Statement != transferParams.Statement {
		t.Errorf("Statement %q does not match expected statement %q\n", target.Statement, transferParams.Statement)
	}

	if target.Bank == nil {
		t.Errorf("Bank account is not set\n")
	}

	if target.Status != Pending {
		t.Errorf("Unexpected status %q\n", target.Status)
	}

	if target.Type != BankTransfer {
		t.Errorf("Unexpected type %q\n", target.Type)
	}

	recipient.Del(rec.Id)
}

func TestTransferGet(t *testing.T) {
	recipientParams := &RecipientParams{
		Name: "Recipient Name",
		Type: Individual,
		Card: &CardParams{
			Name:   "Test Debit",
			Number: "4000056655665556",
			Month:  "10",
			Year:   "20",
		},
	}

	rec, _ := recipient.New(recipientParams)

	transferParams := &TransferParams{
		Amount:    100,
		Currency:  USD,
		Recipient: rec.Id,
	}

	trans, _ := New(transferParams)

	target, err := Get(trans.Id, nil)

	if err != nil {
		t.Error(err)
	}

	if target.Card == nil {
		t.Errorf("Card is not set\n")
	}

	if target.Type != CardTransfer {
		t.Errorf("Unexpected type %q\n", target.Type)
	}

	recipient.Del(rec.Id)
}

func TestTransferUpdate(t *testing.T) {
	recipientParams := &RecipientParams{
		Name: "Recipient Name",
		Type: Corp,
		Bank: &BankAccountParams{
			Country: "US",
			Routing: "110000000",
			Account: "000123456789",
		},
	}

	rec, _ := recipient.New(recipientParams)

	transferParams := &TransferParams{
		Amount:    100,
		Currency:  USD,
		Recipient: rec.Id,
		Desc:      "Original",
	}

	trans, _ := New(transferParams)

	updated := &TransferParams{
		Desc: "Updated",
	}

	target, err := Update(trans.Id, updated)

	if err != nil {
		t.Error(err)
	}

	if target.Desc != updated.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, updated.Desc)
	}

	recipient.Del(rec.Id)
}

func TestTransferList(t *testing.T) {
	recipientParams := &RecipientParams{
		Name: "Recipient Name",
		Type: Individual,
		Card: &CardParams{
			Name:   "Test Debit",
			Number: "4000056655665556",
			Month:  "10",
			Year:   "20",
		},
	}

	rec, _ := recipient.New(recipientParams)

	transferParams := &TransferParams{
		Amount:    100,
		Currency:  USD,
		Recipient: rec.Id,
	}

	for i := 0; i < 5; i++ {
		New(transferParams)
	}

	i := List(&TransferListParams{Recipient: rec.Id})
	for !i.Stop() {
		target, err := i.Next()

		if err != nil {
			t.Error(err)
		}

		if target == nil {
			t.Error("No nil values expected")
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
	}

	recipient.Del(rec.Id)
}
