package charge

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
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
