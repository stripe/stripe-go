package invoice

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
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
	assert.NotNil(t, i.InvoiceList())
}

func TestInvoiceListLines(t *testing.T) {
	i := ListLines(&stripe.InvoiceLineListParams{
		ID: stripe.String("in_123"),
	})

	// Verify that we can get at least one invoice
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.InvoiceLine())
	assert.NotNil(t, i.InvoiceLineList())
}

func TestInvoiceNew(t *testing.T) {
	invoice, err := New(&stripe.InvoiceParams{
		CollectionMethod: stripe.String(string(stripe.InvoiceCollectionMethodChargeAutomatically)),
		Customer:         stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}

func TestInvoicePay(t *testing.T) {
	invoice, err := Pay("in_123", &stripe.InvoicePayParams{
		Source: stripe.String("src_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}

func TestInvoiceUpdate(t *testing.T) {
	invoice, err := Update("in_123", &stripe.InvoiceParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}

func TestInvoiceDel(t *testing.T) {
	invoice, err := Del("in_123", &stripe.InvoiceParams{})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}

func TestInvoiceFinalizeInvoice(t *testing.T) {
	invoice, err := FinalizeInvoice("in_123", &stripe.InvoiceFinalizeParams{})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}

func TestInvoiceMarkUncollectible(t *testing.T) {
	invoice, err := MarkUncollectible("in_123", &stripe.InvoiceMarkUncollectibleParams{})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}

func TestInvoiceSendInvoice(t *testing.T) {
	invoice, err := SendInvoice("in_123", &stripe.InvoiceSendParams{})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}

func TestInvoiceVoidInvoice(t *testing.T) {
	invoice, err := VoidInvoice("in_123", &stripe.InvoiceVoidParams{})
	assert.Nil(t, err)
	assert.NotNil(t, invoice)
}
