package stripe

import (
	"fmt"
	"testing"
)

const testKey = "tGN0bIwXnHdwOa85VABjPdSn8nWY7G7I"

func TestCheckinConnectivity(t *testing.T) {
	c := &Client{}
	c.Init(testKey, nil, nil)

	target, err := c.Account.Get()

	if err != nil {
		t.Error(err)
	}

	if target.Id != "cuD9Rwx8pgmRZRpVe02lsuR9cwp2Bzf7" {
		t.Errorf("Invalid account id: %q\n", target.Id)
	}

	if target.Name != "Stripe Test" {
		t.Errorf("Invalid account name: %q\n", target.Name)
	}

	if target.Email != "test+bindings@stripe.com" {
		t.Errorf("Invalid account email: %q\n", target.Email)
	}
}

func TestCheckinError(t *testing.T) {
	c := &Client{}
	c.Init("bad_key", nil, nil)

	_, err := c.Account.Get()

	if err == nil {
		t.Errorf("Expected an error")
	}

	stripeErr := err.(*Error)

	if stripeErr.Type != InvalidRequest {
		t.Errorf("Type %v does not match expected type\n", stripeErr.Type)
	}
}

func TestCheckinPost(t *testing.T) {
	c := &Client{}
	c.Init(testKey, nil, nil)

	charge := &ChargeParams{
		Amount:   100,
		Currency: USD,
		Card: &CardParams{
			Name:   "Go Bindings Cardholder",
			Number: "4242424242424242",
			Month:  "12",
			Year:   "24",
		},
	}

	target, err := c.Charges.Create(charge)

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

func TestCheckinDelete(t *testing.T) {
	c := &Client{}
	c.Init(testKey, nil, nil)

	plan := &PlanParams{
		Id:       "go_binding",
		Name:     "Go Test Plan",
		Amount:   100,
		Currency: USD,
		Interval: Month,
	}

	_, err := c.Plans.Create(plan)

	if err != nil {
		t.Error(err)
	}

	err = c.Plans.Delete(plan.Id)

	if err != nil {
		t.Error(err)
	}
}

func TestCheckinList(t *testing.T) {
	const runs = 4
	c := &Client{}
	c.Init(testKey, nil, nil)

	for i := 0; i < runs; i++ {
		plan := &PlanParams{
			Id:       fmt.Sprintf("go_binding_%v", i),
			Name:     fmt.Sprintf("Go Test Plan %v", i),
			Amount:   100,
			Currency: USD,
			Interval: Month,
		}

		c.Plans.Create(plan)
	}

	params := &PlanListParams{}
	params.Filters.AddFilter("limit", "", "2")
	target, err := c.Plans.List(params)

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != runs/2 {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	params.Filters.AddFilter("starting_after", "", target.Values[len(target.Values)-1].Id)
	target, err = c.Plans.List(params)

	if len(target.Values) != runs/2 {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	for i := 0; i < runs; i++ {
		c.Plans.Delete(fmt.Sprintf("go_binding_%v", i))
	}
}
