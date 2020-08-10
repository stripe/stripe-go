package person

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestPersonDel(t *testing.T) {
	person, err := Del("person_123", &stripe.PersonParams{
		Account: stripe.String("acct_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, person)
}

func TestPersonGet(t *testing.T) {
	person, err := Get("person_123", &stripe.PersonParams{
		Account: stripe.String("acct_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, person)
}

func TestPersonList(t *testing.T) {
	i := List(&stripe.PersonListParams{
		Account: stripe.String("acct_123"),
		Relationship: &stripe.RelationshipListParams{
			Owner: stripe.Bool(true),
		},
	})

	// Verify that we can get at least one person
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Person())
	assert.NotNil(t, i.PersonList())
}

func TestPersonNew(t *testing.T) {
	person, err := New(&stripe.PersonParams{
		Account:   stripe.String("acct_123"),
		FirstName: stripe.String("John"),
		Relationship: &stripe.RelationshipParams{
			Owner: stripe.Bool(true),
		},
		Verification: &stripe.PersonVerificationParams{
			Document: &stripe.PersonVerificationDocumentParams{
				Back:  stripe.String("file_123"),
				Front: stripe.String("file_234"),
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, person)
}

func TestPersonUpdate(t *testing.T) {
	person, err := Update("person_123", &stripe.PersonParams{
		Account:   stripe.String("acct_123"),
		FirstName: stripe.String("John"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, person)
}
