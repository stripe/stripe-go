package stripe

import (
	"testing"
)

func TestRecipientCreate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	recipient := &RecipientParams{
		Name:  "Recipient Name",
		Type:  Individual,
		TaxId: "000000000",
		Email: "a@b.com",
		Desc:  "Recipient Desc",
		Bank: &BankAccountParams{
			Country: "US",
			Routing: "110000000",
			Account: "000123456789",
		},
		Card: &CardParams{
			Name:   "Test Debit",
			Number: "4000056655665556",
			Month:  "10",
			Year:   "20",
		},
	}

	target, err := c.Recipients.Create(recipient)

	if err != nil {
		t.Error(err)
	}

	if target.Name != recipient.Name {
		t.Errorf("Name %q does not match expected name %q\n", target.Name, recipient.Name)
	}

	if target.Type != recipient.Type {
		t.Errorf("Type %q does not match expected type %q\n", target.Type, recipient.Type)
	}

	if target.Email != recipient.Email {
		t.Errorf("Email %q does not match expected email %q\n", target.Email, recipient.Email)
	}

	if target.Desc != recipient.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, recipient.Desc)
	}

	if target.Created == 0 {
		t.Errorf("Created date is not set\n")
	}

	if target.Bank == nil {
		t.Errorf("Bank account is not set\n")
	}

	if target.Bank.Country != recipient.Bank.Country {
		t.Errorf("Bank country %q does not match expected country %q\n", target.Bank.Country, recipient.Bank.Country)
	}

	if target.Bank.Currency != USD {
		t.Errorf("Bank currency %q does not match expected value\n", target.Bank.Currency)
	}

	if target.Bank.LastFour != "6789" {
		t.Errorf("Bank last four %q does not match expected value\n", target.Bank.LastFour)
	}

	if len(target.Bank.Name) == 0 {
		t.Errorf("Bank name is not set\n")
	}

	if target.Cards == nil || target.Cards.Count != 1 {
		t.Errorf("Recipient cards not set\n")
	}

	if len(target.DefaultCard.Id) == 0 {
		t.Errorf("Recipient default card is not set\n")
	}

	c.Recipients.Delete(target.Id)
}

func TestRecipientGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	recipient := &RecipientParams{
		Name: "Recipient Name",
		Type: Individual,
	}

	rec, _ := c.Recipients.Create(recipient)

	target, err := c.Recipients.Get(rec.Id)

	if err != nil {
		t.Error(err)
	}

	if len(target.Id) == 0 {
		t.Errorf("Recipient not found\n")
	}

	c.Recipients.Delete(target.Id)
}

func TestRecipientUpdate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	recipient := &RecipientParams{
		Name:  "Original Name",
		Type:  Individual,
		Email: "original@b.com",
		Desc:  "Original Desc",
	}

	original, _ := c.Recipients.Create(recipient)

	updated := &RecipientParams{
		Name:  "Updated Name",
		Email: "updated@b.com",
		Desc:  "Updated Desc",
	}

	target, err := c.Recipients.Update(original.Id, updated)

	if err != nil {
		t.Error(err)
	}

	if target.Name != updated.Name {
		t.Errorf("Name %q does not match expected name %q\n", target.Name, updated.Name)
	}

	if target.Email != updated.Email {
		t.Errorf("Email %q does not match expected email %q\n", target.Email, updated.Email)
	}

	if target.Desc != updated.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, updated.Desc)
	}

	c.Recipients.Delete(target.Id)
}

func TestRecipientList(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	recipient := &RecipientParams{
		Name: "Recipient Name",
		Type: Individual,
	}

	recipients := make([]string, 5)

	for i := 0; i < 5; i++ {
		rec, _ := c.Recipients.Create(recipient)
		recipients[i] = rec.Id
	}

	target, err := c.Recipients.List(nil)

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != len(recipients) {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	for _, v := range recipients {
		c.Recipients.Delete(v)
	}
}
