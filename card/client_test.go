package card

import (
	"testing"

	. "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/recipient"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	Key = GetTestKey()
}

func TestCardNew(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	cardParams := &CardParams{
		Number:   "4242424242424242",
		Month:    "10",
		Year:     "20",
		Customer: cust.Id,
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

	targetCust, err := customer.Get(cust.Id, nil)

	if err != nil {
		t.Error(err)
	}

	if targetCust.Cards.Count != 2 {
		t.Errorf("Unexpected number of cards %v\n", targetCust.Cards.Count)
	}

	customer.Del(cust.Id)
}

func TestCardGet(t *testing.T) {
	recipientParams := &RecipientParams{
		Name: "Test Recipient",
		Type: Corp,
		Card: &CardParams{
			Number: "5200828282828210",
			Month:  "06",
			Year:   "20",
		},
	}

	rec, _ := recipient.New(recipientParams)

	target, err := Get(rec.DefaultCard.Id, &CardParams{Recipient: rec.Id})

	if err != nil {
		t.Error(err)
	}

	if target.LastFour != "8210" {
		t.Errorf("Unexpected last four %q for card number %v\n", target.LastFour, recipientParams.Card.Number)
	}

	if target.Brand != MasterCard {
		t.Errorf("Card brand %q does not match expected value\n", target.Brand)
	}

	if target.Funding != DebitFunding {
		t.Errorf("Card funding %q does not match expected value\n", target.Funding)
	}

	recipient.Del(rec.Id)
}

func TestCardDel(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	err := Del(cust.DefaultCard.Id, &CardParams{Customer: cust.Id})
	if err != nil {
		t.Error(err)
	}

	customer.Del(cust.Id)
}

func TestCardUpdate(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
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

	cardParams := &CardParams{
		Customer: cust.Id,
		Name:     "Updated Name",
	}

	target, err := Update(cust.DefaultCard.Id, cardParams)

	if err != nil {
		t.Error(err)
	}

	if target.Name != cardParams.Name {
		t.Errorf("Card name %q does not match expected name %q\n", target.Name, cardParams.Name)
	}

	customer.Del(cust.Id)
}

func TestCardList(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	card := &CardParams{
		Number:   "4242424242424242",
		Month:    "10",
		Year:     "20",
		Customer: cust.Id,
	}

	New(card)

	i := List(&CardListParams{Customer: cust.Id})
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

	customer.Del(cust.Id)
}
