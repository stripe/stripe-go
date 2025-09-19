package stripe_test

import (
	"context"
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v82"
	. "github.com/stripe/stripe-go/v82/testing"
)

func TestList_HasLastResponse(t *testing.T) {
	sc := stripe.NewClient(TestAPIKey)
	i := sc.V1Disputes.List(context.TODO(), &stripe.DisputeListParams{
		PaymentIntent: stripe.String("pi_123"),
	})
	i(func(dispute *stripe.Dispute, err error) bool {
		assert.Nil(t, err)
		assert.NotNil(t, dispute)
		assert.NotNil(t, dispute.LastResponse)
		rawJSON := dispute.LastResponse.RawJSON
		lastResponseDispute := &stripe.Dispute{}
		if err := json.Unmarshal(rawJSON, &lastResponseDispute); err != nil {
			assert.NoError(t, err)
		}
		assert.Equal(t, dispute.ID, lastResponseDispute.ID)
		return true
	})
}
