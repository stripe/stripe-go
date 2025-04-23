package invoiceitem

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestInvoiceItemDel(t *testing.T) {
	item, err := Del("ii_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestInvoiceItemGet(t *testing.T) {
	item, err := Get("ii_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestInvoiceItemList(t *testing.T) {
	i := List(&stripe.InvoiceItemListParams{})

	// Verify that we can get at least one item
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.InvoiceItem())
	assert.NotNil(t, i.InvoiceItemList())
}

func TestInvoiceItemNew(t *testing.T) {
	item, err := New(&stripe.InvoiceItemParams{
		Amount:   stripe.Int64(123),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestInvoiceItemUpdate(t *testing.T) {
	item, err := Update("ii_123", &stripe.InvoiceItemParams{
		Description: stripe.String("Updated description"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, item)
}
