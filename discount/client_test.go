package discount

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestDiscountDel(t *testing.T) {
	discount, err := Del("cus_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, discount)
}

func TestDiscountDelSub(t *testing.T) {
	discount, err := DelSubscription("sub_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, discount)
}
