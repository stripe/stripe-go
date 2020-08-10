package dispute

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestDisputeClose(t *testing.T) {
	dispute, err := Close("dp_123", &stripe.DisputeParams{})
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
}

func TestDisputeGet(t *testing.T) {
	dispute, err := Get("dp_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
}

func TestDisputeList(t *testing.T) {
	i := List(&stripe.DisputeListParams{})

	// Verify that we can get at least one dispute
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Dispute())
	assert.NotNil(t, i.DisputeList())
}

func TestDisputeUpdate(t *testing.T) {
	dispute, err := Update("dp_123", &stripe.DisputeParams{
		Evidence: &stripe.DisputeEvidenceParams{
			CustomerName: stripe.String("A Name"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
}
