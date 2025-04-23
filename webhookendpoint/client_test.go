package webhookendpoint

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestWebhookEndpointDel(t *testing.T) {
	endpoint, err := Del("we_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, endpoint)
}

func TestWebhookEndpointGet(t *testing.T) {
	endpoint, err := Get("we_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, endpoint)
}

func TestWebhookEndpointList(t *testing.T) {
	i := List(&stripe.WebhookEndpointListParams{})

	// Verify that we can get at least one webhook endpoint
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.WebhookEndpoint())
	assert.NotNil(t, i.WebhookEndpointList())
}

func TestWebhookEndpointNew(t *testing.T) {
	endpoint, err := New(&stripe.WebhookEndpointParams{
		EnabledEvents: stripe.StringSlice([]string{
			"charge.succeeded",
		}),
		URL: stripe.String("https://stripe.com"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, endpoint)
}

func TestWebhookEndpointUpdate(t *testing.T) {
	endpoint, err := Update("we_123", &stripe.WebhookEndpointParams{
		EnabledEvents: stripe.StringSlice([]string{
			"charge.succeeded",
		}),
	})
	assert.Nil(t, err)
	assert.NotNil(t, endpoint)
}
