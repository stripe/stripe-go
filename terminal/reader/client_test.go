package reader

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/channelmeter/stripe-go"
	_ "github.com/channelmeter/stripe-go/testing"
)

func TestTerminalReaderDel(t *testing.T) {
	reader, err := Del("loc_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, reader)
	assert.Equal(t, "terminal.reader", reader.Object)
}

func TestTerminalReaderGet(t *testing.T) {
	reader, err := Get("rdr_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, reader)
	assert.Equal(t, "terminal.reader", reader.Object)
}

func TestTerminalReaderList(t *testing.T) {
	i := List(&stripe.TerminalReaderListParams{})

	// Verify that we can get at least one reader
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.TerminalReader())
	assert.Equal(t, "terminal.reader", i.TerminalReader().Object)
}

func TestTerminalReaderNew(t *testing.T) {
	reader, err := New(&stripe.TerminalReaderParams{
		Label:            stripe.String("name"),
		RegistrationCode: stripe.String("a-b-c"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reader)
	assert.Equal(t, "terminal.reader", reader.Object)
}

func TestTerminalReaderUpdate(t *testing.T) {
	reader, err := Update("rdr_123", &stripe.TerminalReaderParams{
		Label: stripe.String("new name"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reader)
	assert.Equal(t, "terminal.reader", reader.Object)
}
