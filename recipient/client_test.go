package recipient

import (
	"testing"

	. "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	Key = GetTestKey()
}

func TestRecipientCreate(t *testing.T) {
	recipientParams := &RecipientParams{
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

	target, err := Create(recipientParams)

	if err != nil {
		t.Error(err)
	}

	if target.Name != recipientParams.Name {
		t.Errorf("Name %q does not match expected name %q\n", target.Name, recipientParams.Name)
	}

	if target.Type != recipientParams.Type {
		t.Errorf("Type %q does not match expected type %q\n", target.Type, recipientParams.Type)
	}

	if target.Email != recipientParams.Email {
		t.Errorf("Email %q does not match expected email %q\n", target.Email, recipientParams.Email)
	}

	if target.Desc != recipientParams.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, recipientParams.Desc)
	}

	if target.Created == 0 {
		t.Errorf("Created date is not set\n")
	}

	if target.Bank == nil {
		t.Errorf("Bank account is not set\n")
	}

	if target.Bank.Country != recipientParams.Bank.Country {
		t.Errorf("Bank country %q does not match expected country %q\n", target.Bank.Country, recipientParams.Bank.Country)
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

	Delete(target.Id)
}

func TestRecipientGet(t *testing.T) {
	recipientParams := &RecipientParams{
		Name: "Recipient Name",
		Type: Individual,
	}

	rec, _ := Create(recipientParams)

	target, err := Get(rec.Id, nil)

	if err != nil {
		t.Error(err)
	}

	if len(target.Id) == 0 {
		t.Errorf("Recipient not found\n")
	}

	Delete(target.Id)
}

func TestRecipientUpdate(t *testing.T) {
	recipientParams := &RecipientParams{
		Name:  "Original Name",
		Type:  Individual,
		Email: "original@b.com",
		Desc:  "Original Desc",
	}

	original, _ := Create(recipientParams)

	updated := &RecipientParams{
		Name:  "Updated Name",
		Email: "updated@b.com",
		Desc:  "Updated Desc",
	}

	target, err := Update(original.Id, updated)

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

	Delete(target.Id)
}

func TestRecipientList(t *testing.T) {
	recipientParams := &RecipientParams{
		Name: "Recipient Name",
		Type: Individual,
	}

	recipients := make([]string, 5)

	for i := 0; i < 5; i++ {
		rec, _ := Create(recipientParams)
		recipients[i] = rec.Id
	}

	target, err := List(nil)

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != len(recipients) {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	for _, v := range recipients {
		Delete(v)
	}
}
