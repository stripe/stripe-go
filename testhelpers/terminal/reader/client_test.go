package reader

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestTerminalReaderUpdate(t *testing.T) {
	reader, err := PresentPaymentMethod("rdr_123", &stripe.TestHelpersTerminalReaderPresentPaymentMethodParams{
		Type: stripe.String("card_present"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reader)
	assert.Equal(t, "terminal.reader", reader.Object)
}
