package balance

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	_ "github.com/stripe/stripe-go/v75/testing"
)

func TestBalanceGet(t *testing.T) {
	balance, err := Get(nil)
	assert.Nil(t, err)
	assert.NotNil(t, balance)
}
