package stripe

import "testing"

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
