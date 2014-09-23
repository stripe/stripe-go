package client

import (
	"testing"

	. "github.com/stripe/stripe-go"
)

func TestErrors(t *testing.T) {
	c := &API{}
	c.Init("bad_key", nil)

	_, err := c.Account.Get()

	if err == nil {
		t.Errorf("Expected an error")
	}

	stripeErr := err.(*Error)

	if stripeErr.Type != InvalidRequest {
		t.Errorf("Type %v does not match expected type\n", stripeErr.Type)
	}

	if stripeErr.HTTPStatusCode != 401 {
		t.Errorf("HTTPStatusCode %q does not match expected value of \"401\"", stripeErr.HTTPStatusCode)
	}
}
