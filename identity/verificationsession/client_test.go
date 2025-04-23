package verificationsession

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestIdentityVerificationSessionCancel(t *testing.T) {
	session, err := Cancel("vs_123", &stripe.IdentityVerificationSessionCancelParams{})
	assert.Nil(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, "identity.verification_session", session.Object)
}

func TestIdentityVerificationSessionRedact(t *testing.T) {
	session, err := Redact("vs_123", &stripe.IdentityVerificationSessionRedactParams{})
	assert.Nil(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, "identity.verification_session", session.Object)
}

func TestIdentityVerificationSessionGet(t *testing.T) {
	session, err := Get("vs_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, "identity.verification_session", session.Object)
}

func TestIdentityVerificationSessionList(t *testing.T) {
	i := List(&stripe.IdentityVerificationSessionListParams{})

	// Verify that we can get at least one session
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IdentityVerificationSession())
	assert.Equal(t, "identity.verification_session", i.IdentityVerificationSession().Object)
	assert.NotNil(t, i.IdentityVerificationSessionList())
}

func TestIdentityVerificationSessionUpdate(t *testing.T) {
	params := &stripe.IdentityVerificationSessionParams{}
	params.AddMetadata("key", "value")
	session, err := Update("vs_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, "identity.verification_session", session.Object)
}
