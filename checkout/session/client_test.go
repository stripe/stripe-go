package session

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestCheckoutSessionGet(t *testing.T) {
	session, err := Get("cs_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, session)
}

func TestCheckoutSessionNew(t *testing.T) {
	params := &stripe.CheckoutSessionParams{
		CancelURL:         stripe.String("https://stripe.com/cancel"),
		ClientReferenceID: stripe.String("1234"),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Quantity: stripe.Int64(2),
			},
		},
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			Description: stripe.String("description"),
			Metadata: map[string]string{
				"attr1": "val1",
				"attr2": "val2",
			},
			Shipping: &stripe.ShippingDetailsParams{
				Address: &stripe.AddressParams{
					Line1: stripe.String("line1"),
					City:  stripe.String("city"),
				},
				Carrier: stripe.String("carrier"),
				Name:    stripe.String("name"),
			},
		},
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			Metadata: map[string]string{
				"attr1": "val1",
				"attr2": "val2",
			},
		},
		SuccessURL: stripe.String("https://stripe.com/success"),
	}
	params.AddExpand("line_items")
	session, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, session.LineItems.Data[0].Object, "item")
}

func TestCheckoutSessionList(t *testing.T) {
	i := List(&stripe.CheckoutSessionListParams{})

	// Verify that we can get at least one session.
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.CheckoutSession())
	assert.NotNil(t, i.CheckoutSessionList())
}

func TestCheckoutSessionListLineItems(t *testing.T) {
	params := &stripe.CheckoutSessionListLineItemsParams{Session: stripe.String("cs_123")}
	i := ListLineItems(params)

	// Verify that we can get at least one line item.
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.LineItem())
	assert.NotNil(t, i.LineItemList())
}
