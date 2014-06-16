package stripe

import (
	"testing"
)

func TestTokenCreate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	token := &TokenParams{
		Card: &CardParams{
			Number: "4242424242424242",
			Month:  "10",
			Year:   "20",
		},
	}

	target, err := c.Tokens.Create(token)

	if err != nil {
		t.Error(err)
	}

	if target.Created == 0 {
		t.Errorf("Created date is not set\n")
	}

	if target.Type != CardToken {
		t.Errorf("Type %v does not match expected value\n", target.Type)
	}

	if target.Card == nil {
		t.Errorf("Card is not set\n")
	}

	if target.Card.LastFour != "4242" {
		t.Errorf("Unexpected last four %q for card number %v\n", target.Card.LastFour, token.Card.Number)
	}
}

func TestTokenGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	token := &TokenParams{
		Bank: &BankAccountParams{
			Country: "US",
			Routing: "110000000",
			Account: "000123456789",
		},
	}

	tok, _ := c.Tokens.Create(token)

	target, err := c.Tokens.Get(tok.Id)

	if err != nil {
		t.Error(err)
	}

	if target.Type != BankToken {
		t.Errorf("Type %v does not match expected value\n", target.Type)
	}

	if target.Bank == nil {
		t.Errorf("Bank account is not set\n")
	}
}
