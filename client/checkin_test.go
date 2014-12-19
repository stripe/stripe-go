package client

import (
	"fmt"
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
	"github.com/stripe/stripe-go/plan"
)

const testKey = "tGN0bIwXnHdwOa85VABjPdSn8nWY7G7I"

func TestCheckinIdempotentDefaults(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   100,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Name:   "Go Bindings Cardholder",
			Number: "4242424242424242",
			Month:  "12",
			Year:   "24",
		},
	}

	if chargeParams.Params.IdempotencyKey != "" {
		t.Errorf("The default value of a Params.IdempotencyKey was not the empty string (%q).\n", chargeParams.Params.IdempotencyKey)
	}

	chargeParams.Params.GenerateIdempotencyKey()

	if chargeParams.Params.IdempotencyKey == "" {
		t.Error("Callind GenerateIdempotencyKey on Params generated an empty idempotency key.")
	}
}

func TestCheckinIdempotentSetter(t *testing.T) {
	foo := &stripe.Params{}
	if foo.IdempotencyKey != "" {
		t.Errorf("The default value of a Params.IdempotencyKey was not the empty string (%q).\n", foo.IdempotencyKey)
	}
	err := foo.SetIdempotencyKey("hello")
	if err != nil {
		t.Error(err)
	}
	err = foo.SetIdempotencyKey("  ")
	if err == nil {
		t.Errorf("Expected a blank idempotency key to fail, even if it had spaces.\n")
	}
	err = foo.SetIdempotencyKey(" h ")
	if err != nil {
		t.Error(err)
	}
	if foo.IdempotencyKey != "h" {
		t.Errorf("Expected the idempotency key submitted as ' h ' to be reduced to 'h'.\n")
	}
	err = foo.SetIdempotencyKey("")
	if err != nil {
		t.Error(err)
	}
	if foo.IdempotencyKey != "" {
		t.Errorf("After clearing the idempotency key, it should be blank.\n")
	}
}

func TestCheckinConnectivity(t *testing.T) {
	c := &API{}
	c.Init(testKey, nil)

	target, err := c.Account.Get()

	if err != nil {
		t.Error(err)
	}

	if target.ID != "cuD9Rwx8pgmRZRpVe02lsuR9cwp2Bzf7" {
		t.Errorf("Invalid account id: %q\n", target.ID)
	}

	if target.Name != "Stripe Test" {
		t.Errorf("Invalid account name: %q\n", target.Name)
	}

	if target.Email != "test+bindings@stripe.com" {
		t.Errorf("Invalid account email: %q\n", target.Email)
	}
}

func TestCheckinError(t *testing.T) {
	c := &API{}
	c.Init("bad_key", nil)

	_, err := c.Account.Get()

	if err == nil {
		t.Errorf("Expected an error")
	}

	stripeErr := err.(*stripe.Error)

	if stripeErr.Type != stripe.InvalidRequest {
		t.Errorf("Type %v does not match expected type\n", stripeErr.Type)
	}
}

func TestCheckinPost(t *testing.T) {
	c := &API{}
	c.Init(testKey, nil)

	charge := &stripe.ChargeParams{
		Amount:   100,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Name:   "Go Bindings Cardholder",
			Number: "4242424242424242",
			Month:  "12",
			Year:   "24",
		},
	}

	target, err := c.Charges.New(charge)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != charge.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, charge.Amount)
	}

	if target.Currency != charge.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, charge.Currency)
	}

	if target.Card.Name != charge.Card.Name {
		t.Errorf("Card name %q does not match expected name %q\n", target.Card.Name, charge.Card.Name)
	}
}

func TestCheckinDel(t *testing.T) {
	c := &API{}
	c.Init(testKey, nil)

	plan := &stripe.PlanParams{
		ID:       "go_binding",
		Name:     "Go Test Plan",
		Amount:   100,
		Currency: currency.USD,
		Interval: plan.Month,
	}

	_, err := c.Plans.New(plan)

	if err != nil {
		t.Error(err)
	}

	err = c.Plans.Del(plan.ID)

	if err != nil {
		t.Error(err)
	}
}

func TestCheckinList(t *testing.T) {
	const runs = 4
	c := &API{}
	c.Init(testKey, nil)

	for i := 0; i < runs; i++ {
		plan := &stripe.PlanParams{
			ID:       fmt.Sprintf("go_binding_%v", i),
			Name:     fmt.Sprintf("Go Test Plan %v", i),
			Amount:   100,
			Currency: currency.USD,
			Interval: plan.Month,
		}

		c.Plans.New(plan)
	}

	params := &stripe.PlanListParams{}
	params.Filters.AddFilter("limit", "", "2")
	params.Single = true

	i := c.Plans.List(params)
	for i.Next() {
		target := i.Plan()

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}

		if target.Amount != 100 {
			t.Errorf("Amount %v does not match expected value\n", target.Amount)
		}
	}
	if err := i.Err(); err != nil {
		t.Error(err)
	}

	for i := 0; i < runs; i++ {
		c.Plans.Del(fmt.Sprintf("go_binding_%v", i))
	}
}
