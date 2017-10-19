package fileupload

//
// This test is unlike other tests: it makes calls to the live Stripe API
// servers. This is because file uploads operate under a different domain from
// the standard api.stripe.com and stripe-mock does not yet support them.
//
// I've nicened this file up a bit for now, and it's not making enough requests
// to cause bad intermittency problems in the test suite, but long term it
// would be nice if we could change these tests to hit a local target so that
// the entire suite can run offline (and more quickly).
//

import (
	"os"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
)

const (
	expectedSize int64 = 734
	expectedType       = "pdf"
)

func init() {
	stripe.Key = "tGN0bIwXnHdwOa85VABjPdSn8nWY7G7I"
}

func TestFileUploadNewThenGet(t *testing.T) {
	t.Skip("File uploads are currently unreliable")

	f, err := os.Open("test_data.pdf")
	if err != nil {
		t.Errorf("Unable to open test file upload file %v\n", err)
	}

	uploadParams := &stripe.FileUploadParams{
		Purpose:    "dispute_evidence",
		FileReader: f,
		Filename:   f.Name(),
	}

	target, err := New(uploadParams)
	assert.NoError(t, err)

	assert.Equal(t, uploadParams.Purpose, target.Purpose)
	assert.Equal(t, expectedSize, target.Size)
	assert.Equal(t, expectedType, target.Type)

	res, err := Get(target.ID, nil)
	assert.NoError(t, err)

	assert.Equal(t, target.ID, res.ID)
}

func TestFileUploadList(t *testing.T) {
	f, err := os.Open("test_data.pdf")
	if err != nil {
		t.Errorf("Unable to open test file upload file %v\n", err)
	}

	uploadParams := &stripe.FileUploadParams{
		Purpose: "dispute_evidence",
		File:    f,
	}

	_, err = New(uploadParams)
	assert.NoError(t, err)

	params := &stripe.FileUploadListParams{}
	params.Filters.AddFilter("limit", "", "5")
	params.Single = true

	i := List(params)
	assert.Nil(t, i.Err())
	for i.Next() {
		assert.NotNil(t, i.FileUpload())
		assert.NotNil(t, i.Meta())
	}
}
