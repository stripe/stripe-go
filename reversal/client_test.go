package reversal

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestReversalGet(t *testing.T) {
	reversal, err := Get("trr_123", &stripe.ReversalParams{
		Transfer: "tr_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}

func TestReversalList(t *testing.T) {
	i := List(&stripe.ReversalListParams{
		Transfer: "tr_123",
	})

	// Verify that we can get at least one reversal
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Reversal())
}

func TestReversalNew(t *testing.T) {
	reversal, err := New(&stripe.ReversalParams{
		Amount:   123,
		Transfer: "tr_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}

func TestReversalUpdate(t *testing.T) {
	reversal, err := Update("trr_123", &stripe.ReversalParams{
		Params: stripe.Params{
			Meta: map[string]string{
				"foo": "bar",
			},
		},
		Transfer: "tr_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}
