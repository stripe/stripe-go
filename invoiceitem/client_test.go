package invoiceitem

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
	_ "github.com/stripe/stripe-go/testing"
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
}

func TestInvoiceItemNew(t *testing.T) {
	item, err := New(&stripe.InvoiceItemParams{
		Amount:   123,
		Currency: currency.USD,
		Customer: "cus_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestInvoiceItemUpdate(t *testing.T) {
	item, err := Update("ii_123", &stripe.InvoiceItemParams{
		Desc: "Updated description",
	})
	assert.Nil(t, err)
	assert.NotNil(t, item)
}
