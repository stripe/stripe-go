package filelink

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestFileLinkGet(t *testing.T) {
	fileLink, err := Get("link_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, fileLink)
}

func TestFileLinkList(t *testing.T) {
	i := List(&stripe.FileLinkListParams{})

	// Verify that we can get at least one fileLink
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.FileLink())
	assert.NotNil(t, i.FileLinkList())
}

func TestFileLinkNew(t *testing.T) {
	fileLink, err := New(&stripe.FileLinkParams{
		File: stripe.String("file_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, fileLink)
}

func TestFileLinkUpdate(t *testing.T) {
	params := &stripe.FileLinkParams{}
	params.AddMetadata("key", "value")
	fileLink, err := Update("link_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, fileLink)
}
