package creditnote

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestCreditNoteGet(t *testing.T) {
	cn, err := Get("cn_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, cn)
}

func TestCreditNoteList(t *testing.T) {
	params := &stripe.CreditNoteListParams{
		Invoice: stripe.String("in_123"),
	}
	i := List(params)

	// Verify that we can get at least one credit note
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.CreditNote())
	assert.NotNil(t, i.CreditNoteList())
}

func TestCreditNoteListLines(t *testing.T) {
	i := ListLines(&stripe.CreditNoteLineItemListParams{
		ID: stripe.String("cn_123"),
	})

	// Verify that we can get at least one invoice
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.CreditNoteLineItem())
	assert.NotNil(t, i.CreditNoteLineItemList())
}

func TestCreditNoteListPreviewLines(t *testing.T) {
	params := &stripe.CreditNoteLineItemListPreviewParams{
		Invoice: stripe.String("in_123"),
		Lines: []*stripe.CreditNoteLineParams{
			{
				Type:            stripe.String(string(stripe.CreditNoteLineItemTypeInvoiceLineItem)),
				Amount:          stripe.Int64(100),
				InvoiceLineItem: stripe.String("ili_123"),
				TaxRates: stripe.StringSlice([]string{
					"txr_123",
				}),
			},
		},
	}
	i := ListPreviewLines(params)

	// Verify that we can get at least one invoice
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.CreditNoteLineItem())
}

func TestCreditNoteNew(t *testing.T) {
	params := &stripe.CreditNoteParams{
		Amount:  stripe.Int64(100),
		Invoice: stripe.String("in_123"),
		Reason:  stripe.String(string(stripe.CreditNoteReasonDuplicate)),
		Lines: []*stripe.CreditNoteLineParams{
			{
				Type:            stripe.String(string(stripe.CreditNoteLineItemTypeInvoiceLineItem)),
				Amount:          stripe.Int64(100),
				InvoiceLineItem: stripe.String("ili_123"),
				TaxRates: stripe.StringSlice([]string{
					"txr_123",
				}),
			},
		},
	}
	cn, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, cn)
}

func TestCreditNoteUpdate(t *testing.T) {
	params := &stripe.CreditNoteParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	}
	cn, err := Update("cn_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, cn)
}

func TestCreditNotePreview(t *testing.T) {
	params := &stripe.CreditNotePreviewParams{
		Amount:  stripe.Int64(100),
		Invoice: stripe.String("in_123"),
		Lines: []*stripe.CreditNoteLineParams{
			{
				Type:            stripe.String(string(stripe.CreditNoteLineItemTypeInvoiceLineItem)),
				Amount:          stripe.Int64(100),
				InvoiceLineItem: stripe.String("ili_123"),
				TaxRates: stripe.StringSlice([]string{
					"txr_123",
				}),
			},
		},
	}
	cn, err := Preview(params)
	assert.Nil(t, err)
	assert.NotNil(t, cn)
}

func TestCreditNoteVoidCreditNote(t *testing.T) {
	params := &stripe.CreditNoteVoidParams{}
	cn, err := VoidCreditNote("cn_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, cn)
}
