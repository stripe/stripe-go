package stripe

import (
	"testing"
)

func TestTransferCreate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	recipient := &RecipientParams{
		Name: "Recipient Name",
		Type: Individual,
		Bank: &BankAccountParams{
			Country: "US",
			Routing: "110000000",
			Account: "000123456789",
		},
	}

	rec, _ := c.Recipients.Create(recipient)

	transfer := &TransferParams{
		Amount:    100,
		Currency:  USD,
		Recipient: rec.Id,
		Desc:      "Transfer Desc",
		Statement: "Transfer",
	}

	target, err := c.Transfers.Create(transfer)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != transfer.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, transfer.Amount)
	}

	if target.Currency != transfer.Currency {
		t.Errorf("Curency %q does not match expected currency %q\n", target.Currency, transfer.Currency)
	}

	if target.Created == 0 {
		t.Errorf("Created date is not set\n")
	}

	if target.Date == 0 {
		t.Errorf("Date is not set \n")
	}

	if target.Desc != transfer.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, target.Desc)
	}

	if target.Recipient.Id != transfer.Recipient {
		t.Errorf("Recipient %q does not match expected recipient %q\n", target.Recipient.Id, target.Recipient)
	}

	if target.Statement != transfer.Statement {
		t.Errorf("Statement %q does not match expected statement %q\n", target.Statement, target.Statement)
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

	c.Recipients.Delete(rec.Id)
}

func TestTransferGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	recipient := &RecipientParams{
		Name: "Recipient Name",
		Type: Individual,
		Card: &CardParams{
			Name:   "Test Debit",
			Number: "4000056655665556",
			Month:  "10",
			Year:   "20",
		},
	}

	rec, _ := c.Recipients.Create(recipient)

	transfer := &TransferParams{
		Amount:    100,
		Currency:  USD,
		Recipient: rec.Id,
	}

	trans, _ := c.Transfers.Create(transfer)

	target, err := c.Transfers.Get(trans.Id)

	if err != nil {
		t.Error(err)
	}

	if target.Card == nil {
		t.Errorf("Card is not set\n")
	}

	if target.Type != CardTransfer {
		t.Errorf("Unexpected type %q\n", target.Type)
	}

	c.Recipients.Delete(rec.Id)
}

func TestTransferUpdate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	recipient := &RecipientParams{
		Name: "Recipient Name",
		Type: Corp,
		Bank: &BankAccountParams{
			Country: "US",
			Routing: "110000000",
			Account: "000123456789",
		},
	}

	rec, _ := c.Recipients.Create(recipient)

	transfer := &TransferParams{
		Amount:    100,
		Currency:  USD,
		Recipient: rec.Id,
		Desc:      "Original",
	}

	trans, _ := c.Transfers.Create(transfer)

	updated := &TransferParams{
		Desc: "Updated",
	}

	target, err := c.Transfers.Update(trans.Id, updated)

	if err != nil {
		t.Error(err)
	}

	if target.Desc != updated.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, updated.Desc)
	}

	c.Recipients.Delete(rec.Id)
}

func TestTransferList(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	recipient := &RecipientParams{
		Name: "Recipient Name",
		Type: Individual,
		Card: &CardParams{
			Name:   "Test Debit",
			Number: "4000056655665556",
			Month:  "10",
			Year:   "20",
		},
	}

	rec, _ := c.Recipients.Create(recipient)

	transfer := &TransferParams{
		Amount:    100,
		Currency:  USD,
		Recipient: rec.Id,
	}

	for i := 0; i < 5; i++ {
		c.Transfers.Create(transfer)
	}

	target, err := c.Transfers.List(&TransferListParams{Recipient: rec.Id})

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != 5 {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	c.Recipients.Delete(rec.Id)
}
