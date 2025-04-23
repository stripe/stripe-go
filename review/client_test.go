package review

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
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
