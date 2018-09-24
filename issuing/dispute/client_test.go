package dispute

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestIssuingDisputeGet(t *testing.T) {
	dispute, err := Get("idp_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
	assert.Equal(t, "issuing.dispute", dispute.Object)
}

func TestIssuingDisputeList(t *testing.T) {
	i := List(&stripe.IssuingDisputeListParams{})

	// Verify that we can get at least one dispute
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IssuingDispute())
	assert.Equal(t, "issuing.dispute", i.IssuingDispute().Object)
}

func TestIssuingDisputeNew(t *testing.T) {
	dispute, err := New(&stripe.IssuingDisputeParams{
		Evidence: &stripe.IssuingDisputeEvidenceParams{
			Fraudulent: &stripe.IssuingDisputeEvidenceFraudulentParams{
				DisputeExplanation: stripe.String("My explanation"),
				UncategorizedFile:  stripe.String("file_123"),
			},
		},
		Reason:              stripe.String(string(stripe.IssuingDisputeReasonFraudulent)),
		DisputedTransaction: stripe.String("ipi_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
	assert.Equal(t, "issuing.dispute", dispute.Object)
}

func TestIssuingDisputeUpdate(t *testing.T) {
	params := &stripe.IssuingDisputeParams{}
	params.AddMetadata("key", "value")
	dispute, err := Update("idp_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
	assert.Equal(t, "issuing.dispute", dispute.Object)
}
