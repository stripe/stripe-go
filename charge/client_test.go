package charge

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/bitcoinreceiver"
	"github.com/stripe/stripe-go/currency"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/refund"
	"github.com/stripe/stripe-go/token"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestChargeNew(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Name:   "Stripe Tester",
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Statement: "statement",
		Email:     "a@b.com",
	}

	target, err := New(chargeParams)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != chargeParams.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, chargeParams.Amount)
	}

	if target.Currency != chargeParams.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, chargeParams.Currency)
	}

	if target.Card.Name != chargeParams.Card.Name {
		t.Errorf("Card name %q does not match expected name %q\n", target.Card.Name, chargeParams.Card.Name)
	}

	if target.Statement != chargeParams.Statement {
		t.Errorf("Statement description %q does not match expected description %v\n", target.Statement, chargeParams.Statement)
	}

	if target.Email != chargeParams.Email {
		t.Errorf("Email %q does not match expected email %v\n", target.Email, chargeParams.Email)
	}
}

func TestWithoutIdempotentTwoDifferentCharges(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Name:   "Stripe Tester",
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Statement: "statement",
		Email:     "a@b.com",
	}

	if chargeParams.Params.IdempotencyKey != "" {
		t.Errorf("The default value of a Params.IdempotencyKey was not blank, and it needs to be. (%q).", chargeParams.Params.IdempotencyKey)
	}

	first, err := New(chargeParams)

	if err != nil {
		t.Error(err)
	}

	second, err := New(chargeParams)

	if err != nil {
		t.Error(err)
	}

	if first.ID == second.ID {
		t.Errorf("Created two charges with no Idempotency Key (%s), but they resulted in charges with the same IDs (%q and %q).\n", chargeParams.Params.IdempotencyKey, first.ID, second.ID)
	}
}

func TestChargeNewWithCustomerAndCard(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Customer: cust.ID,
		Card: &stripe.CardParams{
			Token: cust.Cards.Values[0].ID,
		},
		Statement: "statement",
		Email:     "a@b.com",
	}

	target, err := New(chargeParams)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != chargeParams.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, chargeParams.Amount)
	}

	if target.Currency != chargeParams.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, chargeParams.Currency)
	}

	if target.Customer.ID != cust.ID {
		t.Errorf("Customer ID %q doesn't match expected customer ID %q", target.Customer.ID, cust.ID)
	}

	if target.Card.ID != cust.Cards.Values[0].ID {
		t.Errorf("Card ID %q doesn't match expected card ID %q", target.Card.ID, cust.Cards.Values[0].ID)
	}

}

func TestChargeNewWithToken(t *testing.T) {
	tokenParams := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: "4242424242424242",
			Month:  "10",
			Year:   "20",
		},
	}

	tok, _ := token.New(tokenParams)

	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Token: tok.ID,
		},
	}

	target, err := New(chargeParams)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != chargeParams.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, chargeParams.Amount)
	}

	if target.Currency != chargeParams.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, chargeParams.Currency)
	}

	if target.Card.ID != tok.Card.ID {
		t.Errorf("Card Id %q doesn't match card id %q of token %q", target.Card.ID, tok.Card.ID, tok.ID)
	}
}

func TestChargeGet(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1001,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	res, _ := New(chargeParams)

	target, err := Get(res.ID, nil)

	if err != nil {
		t.Error(err)
	}

	if target.ID != res.ID {
		t.Errorf("Charge id %q does not match expected id %q\n", target.ID, res.ID)
	}
}

func TestChargeUpdate(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1002,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Desc: "original description",
	}

	res, _ := New(chargeParams)

	if res.Desc != chargeParams.Desc {
		t.Errorf("Original description %q does not match expected description %q\n", res.Desc, chargeParams.Desc)
	}

	updated := &stripe.ChargeParams{
		Desc: "updated description",
	}

	target, err := Update(res.ID, updated)

	if err != nil {
		t.Error(err)
	}

	if target.Desc != updated.Desc {
		t.Errorf("Updated description %q does not match expected description %q\n", target.Desc, updated.Desc)
	}
}

func TestChargeCapture(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1004,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		NoCapture: true,
	}

	res, _ := New(chargeParams)

	if res.Captured {
		t.Errorf("The charge should not have been captured\n")
	}

	// full capture
	target, err := Capture(res.ID, nil)

	if err != nil {
		t.Error(err)
	}

	if !target.Captured {
		t.Errorf("Expected to have captured this charge after full capture\n")
	}

	res, err = New(chargeParams)

	// partial capture
	capture := &stripe.CaptureParams{
		Amount: 554,
		Email:  "a@b.com",
	}

	target, err = Capture(res.ID, capture)

	if err != nil {
		t.Error(err)
	}

	if !target.Captured {
		t.Errorf("Expected to have captured this charge after partial capture\n")
	}

	if target.AmountRefunded != chargeParams.Amount-capture.Amount {
		t.Errorf("Refunded amount %v does not match expected amount %v\n", target.AmountRefunded, chargeParams.Amount-capture.Amount)
	}

	if target.Email != capture.Email {
		t.Errorf("Email %q does not match expected email %v\n", target.Email, capture.Email)
	}
}

func TestChargeList(t *testing.T) {
	params := &stripe.ChargeListParams{}
	params.Filters.AddFilter("include[]", "", "total_count")
	params.Filters.AddFilter("limit", "", "5")
	params.Single = true

	i := List(params)
	for i.Next() {
		if i.Charge() == nil {
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

func TestMarkFraudulent(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Name:   "Stripe Tester",
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Statement: "statement",
		Email:     "a@b.com",
	}

	target, _ := New(chargeParams)
	refund.New(&stripe.RefundParams{Charge: target.ID})

	ch, _ := MarkFraudulent(target.ID)

	if ch.FraudDetails.UserReport != ReportFraudulent {
		t.Error("UserReport was not fraudulent for a charge marked as fraudulent")
	}
}

func TestMarkSafe(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Name:   "Stripe Tester",
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Statement: "statement",
		Email:     "a@b.com",
	}

	target, _ := New(chargeParams)

	ch, _ := MarkSafe(target.ID)

	if ch.FraudDetails.UserReport != ReportSafe {
		t.Error("UserReport was not safe for a charge marked as safe: ",
			ch.FraudDetails.UserReport)
	}
}

func TestChargeSourceForCard(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Source: &stripe.SourceParams{
			Card: &stripe.CardParams{
				Name:   "Stripe Tester",
				Number: "378282246310005",
				Month:  "06",
				Year:   "20",
			},
		},
		Statement: "statement",
		Email:     "a@b.com",
	}

	ch, _ := New(chargeParams)

	if ch.Source == nil {
		t.Error("Source is nil for Charge `source` property created by a Card")
	}

	if ch.Source.Type != stripe.PaymentSourceCard {
		t.Error("Source Type for Charge created by Card should be `card`")
	}

	card := ch.Source.Card

	if len(card.ID) == 0 {
		t.Error("Source ID is nil for Charge `source` Card property")
	}

	if card.Display() != "American Express ending in 0005" {
		t.Error("Display value did not match expectation")
	}
}

func TestChargeSourceForBitcoinReceiver(t *testing.T) {
	bitcoinReceiverParams := &stripe.BitcoinReceiverParams{
		Amount:   1000,
		Currency: currency.USD,
		Email:    "do+fill_now@stripe.com",
		Desc:     "some details",
	}

	receiver, _ := bitcoinreceiver.New(bitcoinReceiverParams)

	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Source: &stripe.SourceParams{
			ID: receiver.ID,
		},
		Email: "do+fill_now@stripe.com",
	}

	ch, _ := New(chargeParams)

	if len(ch.ID) == 0 {
		t.Error("ID is nil for Charge")
	}

	if ch.Source == nil {
		t.Error("Source is nil for Charge, should be BitcoinReceiver property")
	}

	if ch.Source.Type != stripe.PaymentSourceBitcoinReceiver {
		t.Error("Source Type for Charge created by BitcoinReceiver should be `bitcoin_receiver`")
	}

	rreceiver := ch.Source.BitcoinReceiver

	if len(rreceiver.ID) == 0 {
		t.Error("Source ID is nil for Charge `source` BitcoinReceiver property")
	}

	if rreceiver.Amount == 0 {
		t.Error("Amount is empty for Charge `source` BitcoinReceiver property")
	}

	if rreceiver.Display() != "Filled bitcoin receiver (1000/1000 usd)" {
		t.Error("Display value did not match expectation")
	}
}
