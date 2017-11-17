package source

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/VividCortex/stripe-go"
	"github.com/VividCortex/stripe-go/currency"
	_ "github.com/VividCortex/stripe-go/testing"
)

func TestSourceGet(t *testing.T) {
	source, err := Get("src_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceNew(t *testing.T) {
	source, err := New(&stripe.SourceObjectParams{
		Type:     "bitcoin",
		Amount:   1000,
		Currency: currency.USD,
		Owner: &stripe.SourceOwnerParams{
			Email: "jenny.rosen@example.com",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceUpdate(t *testing.T) {
	source, err := Update("src_123", &stripe.SourceObjectParams{
		Owner: &stripe.SourceOwnerParams{
			Email: "jenny.rosen@example.com",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceDetach(t *testing.T) {
	source, err := Detach("src_123", &stripe.SourceObjectDetachParams{
		Customer: "cus_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}
