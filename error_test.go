package stripe

import "testing"

func TestErrors(t *testing.T) {
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
