package card

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/recipient"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestCardNew(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	cardParams := &stripe.CardParams{
		Number:   "4242424242424242",
		Month:    "10",
		Year:     "20",
		Customer: cust.ID,
		CVC:      "1234",
	}

	target, err := New(cardParams)

	if err != nil {
		t.Error(err)
	}

	if target.LastFour != "4242" {
		t.Errorf("Unexpected last four %q for card number %v\n", target.LastFour, cardParams.Number)
	}

	if target.CVCCheck != Pass {
		t.Errorf("CVC check %q does not match expected status\n", target.ZipCheck)
	}

	targetCust, err := customer.Get(cust.ID, nil)

	if err != nil {
		t.Error(err)
	}

	if targetCust.Cards.Count != 2 {
		t.Errorf("Unexpected number of cards %v\n", targetCust.Cards.Count)
	}

	customer.Del(cust.ID)
}

func TestCardGet(t *testing.T) {
	recipientParams := &stripe.RecipientParams{
		Name: "Test Recipient",
		Type: recipient.Corp,
		Card: &stripe.CardParams{
			Number: "5200828282828210",
			Month:  "06",
			Year:   "20",
		},
	}

	rec, _ := recipient.New(recipientParams)

	target, err := Get(rec.DefaultCard.ID, &stripe.CardParams{Recipient: rec.ID})

	if err != nil {
		t.Error(err)
	}

	if target.LastFour != "8210" {
		t.Errorf("Unexpected last four %q for card number %v\n", target.LastFour, recipientParams.Card.Number)
	}

	if target.Brand != MasterCard {
		t.Errorf("Card brand %q does not match expected value\n", target.Brand)
	}

	if target.Funding != Debit {
		t.Errorf("Card funding %q does not match expected value\n", target.Funding)
	}

	recipient.Del(rec.ID)
}

func TestCardDel(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	err := Del(cust.DefaultCard.ID, &stripe.CardParams{Customer: cust.ID})
	if err != nil {
		t.Error(err)
	}

	customer.Del(cust.ID)
}

func TestCardUpdate(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
			Name:   "Original Name",
		},
	}

	cust, err := customer.New(customerParams)

	if err != nil {
		t.Error(err)
	}

	cardParams := &stripe.CardParams{
		Customer: cust.ID,
		Name:     "Updated Name",
	}

	target, err := Update(cust.DefaultCard.ID, cardParams)

	if err != nil {
		t.Error(err)
	}

	if target.Name != cardParams.Name {
		t.Errorf("Card name %q does not match expected name %q\n", target.Name, cardParams.Name)
	}

	customer.Del(cust.ID)
}

func TestCardList(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	card := &stripe.CardParams{
		Number:   "4242424242424242",
		Month:    "10",
		Year:     "20",
		Customer: cust.ID,
	}

	New(card)

	i := List(&stripe.CardListParams{Customer: cust.ID})
	for i.Next() {
		if i.Card() == nil {
			t.Error("No nil values expected")
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
	}
	if err := i.Err(); err != nil {
		t.Error(err)
	}

	customer.Del(cust.ID)
}
