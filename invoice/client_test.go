package invoice

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestInvoiceGet(t *testing.T) {
	invoice, err := Get("in_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}

func TestInvoiceList(t *testing.T) {
	i := List(&stripe.InvoiceListParams{})

	// Verify that we can get at least one invoice
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Invoice())
}

func TestInvoiceListLines(t *testing.T) {
	i := ListLines(&stripe.InvoiceLineListParams{
		ID: "in_123",
	})

	// Verify that we can get at least one invoice
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.InvoiceLine())
}

func TestInvoiceNew(t *testing.T) {
	invoice, err := New(&stripe.InvoiceParams{
		Customer: "cus_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}

func TestInvoicePay(t *testing.T) {
	invoice, err := Pay("in_123", &stripe.InvoicePayParams{
		Source: "src_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}

func TestInvoiceUpdate(t *testing.T) {
	invoice, err := Update("in_123", &stripe.InvoiceParams{
		Closed: true,
	})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}
