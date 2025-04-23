package source

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestSourceGet(t *testing.T) {
	source, err := Get("src_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceNew(t *testing.T) {
	source, err := New(&stripe.SourceParams{
		Type:     stripe.String("ach_credit_transfer"),
		Amount:   stripe.Int64(1000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Flow:     stripe.String(string(stripe.SourceFlowReceiver)),
		Mandate: &stripe.SourceMandateParams{
			Amount: stripe.Int64(1000),
			Acceptance: &stripe.SourceMandateAcceptanceParams{
				Date: stripe.Int64(1528573382),
				IP:   stripe.String("127.0.0.1"),
				Online: &stripe.SourceMandateAcceptanceOnlineParams{
					Date:      stripe.Int64(1528573382),
					IP:        stripe.String("127.0.0.1"),
					UserAgent: stripe.String("User-Agent"),
				},
				UserAgent: stripe.String("User-Agent"),
				Status:    stripe.String("accepted"),
			},
			Currency:           stripe.String(string(stripe.CurrencyUSD)),
			Interval:           stripe.String("one_time"),
			NotificationMethod: stripe.String("none"),
		},
		Owner: &stripe.SourceOwnerParams{
			Email: stripe.String("jenny.rosen@example.com"),
		},
		Receiver: &stripe.SourceReceiverParams{
			RefundAttributesMethod: stripe.String("manual"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceUpdate(t *testing.T) {
	source, err := Update("src_123", &stripe.SourceParams{
		Owner: &stripe.SourceOwnerParams{
			Email: stripe.String("jenny.rosen@example.com"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceDetach(t *testing.T) {
	source, err := Detach("src_123", &stripe.SourceDetachParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceSharing(t *testing.T) {
	params := &stripe.SourceParams{
		Type:           stripe.String("card"),
		Customer:       stripe.String("cus_123"),
		OriginalSource: stripe.String("src_123"),
		Usage:          stripe.String(string(stripe.SourceUsageReusable)),
	}
	params.SetStripeAccount("acct_123")
	source, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, source)
}
