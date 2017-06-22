package recipient

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestRecipientDel(t *testing.T) {
	recipient, err := Del("rp_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, recipient)
}

func TestRecipientGet(t *testing.T) {
	recipient, err := Get("rp_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, recipient)
}

func TestRecipientList(t *testing.T) {
	i := List(&stripe.RecipientListParams{})

	// Verify that we can get at least one recipient
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Recipient())
}

func TestRecipientUpdate(t *testing.T) {
	recipient, err := Update("rp_123", &stripe.RecipientParams{
		Name: "Updated Name",
	})
	assert.Nil(t, err)
	assert.NotNil(t, recipient)
}
