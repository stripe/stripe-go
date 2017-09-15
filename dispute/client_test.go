package dispute

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestDisputeClose(t *testing.T) {
	dispute, err := Close("dp_123")
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
}

func TestDisputeUpdate(t *testing.T) {
	dispute, err := Update("dp_123", &stripe.DisputeParams{
		Evidence: &stripe.DisputeEvidenceParams{
			CustomerName: "A Name",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
}
