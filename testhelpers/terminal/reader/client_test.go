package reader

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v75"
	_ "github.com/stripe/stripe-go/v75/testing"
)

func TestTerminalReaderUpdate(t *testing.T) {
	reader, err := PresentPaymentMethod("rdr_123", &stripe.TestHelpersTerminalReaderPresentPaymentMethodParams{
		Type: stripe.String("card_present"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reader)
	assert.Equal(t, "terminal.reader", reader.Object)
}
