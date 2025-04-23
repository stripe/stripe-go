package dispute

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestIssuingDisputeGet(t *testing.T) {
	dispute, err := Get("idp_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
	assert.Equal(t, "issuing.dispute", dispute.Object)
}

func TestIssuingDisputeList(t *testing.T) {
	params := &stripe.IssuingDisputeListParams{
		Status:      stripe.String(string(stripe.IssuingDisputeStatusWon)),
		Transaction: stripe.String("ipi_123"),
	}
	i := List(params)

	// Verify that we can get at least one dispute
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IssuingDispute())
	assert.Equal(t, "issuing.dispute", i.IssuingDispute().Object)
	assert.NotNil(t, i.IssuingDisputeList())
}

func TestIssuingDisputeNew(t *testing.T) {
	params := &stripe.IssuingDisputeParams{
		Evidence: &stripe.IssuingDisputeEvidenceParams{
			Canceled: &stripe.IssuingDisputeEvidenceCanceledParams{
				AdditionalDocumentation:    stripe.String("file_123"),
				CanceledAt:                 stripe.Int64(1577836800),
				CancellationPolicyProvided: stripe.Bool(true),
				CancellationReason:         stripe.String("reason for cancellation"),
				ExpectedAt:                 stripe.Int64(1577836800),
				Explanation:                stripe.String("explanation"),
				ProductDescription:         stripe.String("product description"),
				ProductType:                stripe.String(string(stripe.IssuingDisputeEvidenceCanceledProductTypeMerchandise)),
				ReturnStatus:               stripe.String(string(stripe.IssuingDisputeEvidenceCanceledReturnStatusMerchantRejected)),
				ReturnedAt:                 stripe.Int64(1577836800),
			},
			Reason: stripe.String(string(stripe.IssuingDisputeEvidenceReasonCanceled)),
		},
		Transaction: stripe.String("ipi_123"),
	}
	dispute, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
	assert.Equal(t, "issuing.dispute", dispute.Object)
}

func TestIssuingDisputeSubmit(t *testing.T) {
	params := &stripe.IssuingDisputeSubmitParams{}
	dispute, err := Submit("idp_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
	assert.Equal(t, "issuing.dispute", dispute.Object)
}

func TestIssuingDisputeUpdate(t *testing.T) {
	params := &stripe.IssuingDisputeParams{}
	dispute, err := Update("idp_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, dispute)
	assert.Equal(t, "issuing.dispute", dispute.Object)
}
