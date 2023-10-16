package loginlink

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v75"
	_ "github.com/stripe/stripe-go/v75/testing"
)

func TestLoginLinkNew(t *testing.T) {
	link, err := New(&stripe.LoginLinkParams{
		Account: stripe.String("acct_EXPRESS"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, link)
}
