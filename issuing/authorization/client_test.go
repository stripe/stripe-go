package authorization

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestIssuingAuthorizationApprove(t *testing.T) {
	authorization, err := Approve("iauth_123", &stripe.IssuingAuthorizationApproveParams{})
	assert.Nil(t, err)
	assert.NotNil(t, authorization)
	assert.Equal(t, "issuing.authorization", authorization.Object)
}

func TestIssuingAuthorizationDecline(t *testing.T) {
	authorization, err := Decline("iauth_123", &stripe.IssuingAuthorizationDeclineParams{})
	assert.Nil(t, err)
	assert.NotNil(t, authorization)
	assert.Equal(t, "issuing.authorization", authorization.Object)
}

func TestIssuingAuthorizationGet(t *testing.T) {
	authorization, err := Get("iauth_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, authorization)
	assert.Equal(t, "issuing.authorization", authorization.Object)
}

func TestIssuingAuthorizationList(t *testing.T) {
	i := List(&stripe.IssuingAuthorizationListParams{})

	// Verify that we can get at least one authorization
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IssuingAuthorization())
	assert.Equal(t, "issuing.authorization", i.IssuingAuthorization().Object)
}

func TestIssuingAuthorizationUpdate(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	params.AddMetadata("key", "value")
	authorization, err := Update("iauth_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, authorization)
	assert.Equal(t, "issuing.authorization", authorization.Object)
}
