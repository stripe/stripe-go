package loginlink

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestLoginLinkNew(t *testing.T) {
	loginLinkParams := &stripe.LoginLinkParams{
		Account: "acct_EXPRESS",
	}
	target, err := New(loginLinkParams)

	if err != nil {
		t.Error(err)
	}

	if target.URL == nil || len(target.URL) == 0 {
		t.Errorf("Unexpected nil or empty URL on login link\n")
	}
}
