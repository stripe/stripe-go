package authorization

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
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
	assert.NotNil(t, i.IssuingAuthorizationList())
}

func TestIssuingAuthorizationUpdate(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	params.AddMetadata("key", "value")
	authorization, err := Update("iauth_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, authorization)
	assert.Equal(t, "issuing.authorization", authorization.Object)
}
