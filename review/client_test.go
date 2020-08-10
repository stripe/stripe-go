package review

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestReviewApprove(t *testing.T) {
	review, err := Approve("prv_123", &stripe.ReviewApproveParams{})
	assert.Nil(t, err)
	assert.NotNil(t, review)
}

func TestReviewGet(t *testing.T) {
	review, err := Get("prv_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, review)
}

func TestReviewList(t *testing.T) {
	i := List(&stripe.ReviewListParams{})

	// Verify that we can get at least one review
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Review())
	assert.NotNil(t, i.ReviewList())
}
