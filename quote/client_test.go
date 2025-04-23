package quote

import (
	"io/ioutil"
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestQuoteGet(t *testing.T) {
	quote, err := Get("qt_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, quote)
}

func TestQuoteList(t *testing.T) {
	i := List(&stripe.QuoteListParams{})

	// Verify that we can get at least one quote
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Quote())
	assert.NotNil(t, i.QuoteList())
}

func TestQuoteListComputedUpfrontLineItems(t *testing.T) {
	i := ListComputedUpfrontLineItems(&stripe.QuoteListComputedUpfrontLineItemsParams{
		Quote: stripe.String("qt_123"),
	})

	// Verify that we can get at least line item
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.LineItem())
	assert.NotNil(t, i.LineItemList())
}

func TestQuoteListLineItems(t *testing.T) {
	i := ListLineItems(&stripe.QuoteListLineItemsParams{
		Quote: stripe.String("qt_123"),
	})

	// Verify that we can get at least one quote
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.LineItem())
	assert.NotNil(t, i.LineItemList())
}

func TestQuoteNew(t *testing.T) {
	quote, err := New(&stripe.QuoteParams{
		CollectionMethod: stripe.String(string(stripe.QuoteCollectionMethodChargeAutomatically)),
		Customer:         stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, quote)
}

func TestQuotePDF(t *testing.T) {
	stream, err := PDF("qt_123", &stripe.QuotePDFParams{})
	assert.Nil(t, err)
	assert.NotNil(t, stream)
	body, err := ioutil.ReadAll(stream.LastResponse.Body)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.Equal(t, []byte("Stripe binary response"), body)
}

func TestQuoteUpdate(t *testing.T) {
	quote, err := Update("qt_123", &stripe.QuoteParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, quote)
}

func TestQuoteFinalizeQuote(t *testing.T) {
	quote, err := FinalizeQuote("qt_123", &stripe.QuoteFinalizeQuoteParams{})
	assert.Nil(t, err)
	assert.NotNil(t, quote)
}
