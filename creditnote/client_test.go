package creditnote

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
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
}

func TestCreditNoteNew(t *testing.T) {
	params := &stripe.CreditNoteParams{
		Amount:  stripe.Int64(100),
		Invoice: stripe.String("in_123"),
		Reason:  stripe.String(string(stripe.CreditNoteReasonDuplicate)),
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

func TestCreditNoteVoidCreditNote(t *testing.T) {
	params := &stripe.CreditNoteVoidParams{}
	cn, err := VoidCreditNote("cn_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, cn)
}
