package checkoutsession

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestCheckoutSessionNew(t *testing.T) {
	session, err := New(&stripe.CheckoutSessionParams{
		AllowedSourceTypes: []*string{
			stripe.String("card"),
		},
		CancelURL:         stripe.String("https://stripe.com/cancel"),
		ClientReferenceID: stripe.String("1234"),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Amount:      stripe.Int64(1234),
				Currency:    stripe.String(string(stripe.CurrencyUSD)),
				Description: stripe.String("description"),
				Images: []*string{
					stripe.String("https://stripe.com/image1"),
				},
				Name:     stripe.String("name"),
				Quantity: stripe.Int64(2),
			},
		},
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			Description: stripe.String("description"),
			Shipping: &stripe.ShippingDetailsParams{
				Address: &stripe.AddressParams{
					Line1: stripe.String("line1"),
					City:  stripe.String("city"),
				},
				Carrier: stripe.String("carrier"),
				Name:    stripe.String("name"),
			},
		},
		SuccessURL: stripe.String("https://stripe.com/success"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, session)
}
