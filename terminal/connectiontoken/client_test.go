package connectiontoken

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v86"
	_ "github.com/stripe/stripe-go/v86/testing"
)

func TestTerminalConnectionTokenNew(t *testing.T) {
	connectiontoken, err := New(&stripe.TerminalConnectionTokenParams{})
	assert.Nil(t, err)
	assert.NotNil(t, connectiontoken)
}
