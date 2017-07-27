package loginlink

import (
	stripe "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

/*
 To run this test you need an Express account connected to your
 platform. This would break if you run the test with your own API
 key so we're commenting it out.
func TestLoginLinkNew(t *testing.T) {
	loginLinkParams := &stripe.LoginLinkParams{
		Account: "acct_EXPRESS",
	}
	target, err := New(loginLinkParams)

	if err != nil {
		t.Error(err)
	}

	if len(target.Url) == 0 {
		t.Errorf("Unexpected nil or empty Url on login link\n")
	}
}
*/
