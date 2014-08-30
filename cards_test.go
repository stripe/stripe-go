package stripe

import (
	"testing"
)

func TestCardCreate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := c.Customers.Create(customer)

	card := &CardParams{
		Number:   "4242424242424242",
		Month:    "10",
		Year:     "20",
		Customer: cust.Id,
		CVC:      "1234",
	}

	target, err := c.Cards.Create(card)

	if err != nil {
		t.Error(err)
	}

	if target.LastFour != "4242" {
		t.Errorf("Unexpected last four %q for card number %v\n", target.LastFour, card.Number)
	}

	if target.CVCCheck != Pass {
		t.Errorf("CVC check %q does not match expected status\n", target.ZipCheck)
	}

	targetCust, err := c.Customers.Get(cust.Id)
	if err != nil {
		t.Error(err)
	}

	if targetCust.Cards.Count != 2 {
		t.Errorf("Unexpected number of cards %v\n", targetCust.Cards.Count)
	}

	c.Customers.Delete(cust.Id)
}

func TestCardGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	recipient := &RecipientParams{
		Name: "Test Recipient",
		Type: Corp,
		Card: &CardParams{
			Number: "5200828282828210",
			Month:  "06",
			Year:   "20",
		},
	}

	rec, _ := c.Recipients.Create(recipient)

	target, err := c.Cards.Get(rec.DefaultCard.Id, &CardParams{Recipient: rec.Id})

	if err != nil {
		t.Error(err)
	}

	if target.LastFour != "8210" {
		t.Errorf("Unexpected last four %q for card number %v\n", target.LastFour, recipient.Card.Number)
	}

	if target.Brand != MasterCard {
		t.Errorf("Card brand %q does not match expected value\n", target.Brand)
	}

	if target.Funding != DebitFunding {
		t.Errorf("Card funding %q does not match expected value\n", target.Funding)
	}

	c.Recipients.Delete(rec.Id)
}

func TestCardDelete(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := c.Customers.Create(customer)

	err := c.Cards.Delete(cust.DefaultCard.Id, &CardParams{Customer: cust.Id})
	if err != nil {
		t.Error(err)
	}

	c.Customers.Delete(cust.Id)
}

func TestCardUpdate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
			Name:   "Original Name",
		},
	}

	cust, err := c.Customers.Create(customer)
	if err != nil {
		t.Error(err)
	}

	card := &CardParams{
		Customer: cust.Id,
		Name:     "Updated Name",
	}

	target, err := c.Cards.Update(cust.DefaultCard.Id, card)
	if err != nil {
		t.Error(err)
	}

	if target.Name != card.Name {
		t.Errorf("Card name %q does not match expected name %q\n", target.Name, card.Name)
	}

	c.Customers.Delete(cust.Id)
}

func TestCardList(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := c.Customers.Create(customer)

	card := &CardParams{
		Number:   "4242424242424242",
		Month:    "10",
		Year:     "20",
		Customer: cust.Id,
	}

	c.Cards.Create(card)

	target, err := c.Cards.List(&CardListParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != 2 {
		t.Errorf("Number of cards %v does not match expected value\n", len(target.Values))
	}

	c.Customers.Delete(cust.Id)
}
