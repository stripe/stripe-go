package file

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
	"bytes"
	"os"
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

const (
	expectedSize int64 = 734
	expectedType       = "pdf"
)

func TestFileGet(t *testing.T) {
	file, err := Get("file_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, file)
}

func TestFileList(t *testing.T) {
	i := List(&stripe.FileListParams{})

	// Verify that we can get at least one file
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.File())
	assert.NotNil(t, i.FileList())
}

type testBackend struct {
	calledMultipart bool
}

func (b *testBackend) Call(method, path, key string, params stripe.ParamsContainer, v stripe.LastResponseSetter) error {
	return nil
}
func (b *testBackend) CallStreaming(method, path, key string, params stripe.ParamsContainer, v stripe.StreamingLastResponseSetter) error {
	return nil
}
func (b *testBackend) CallRaw(method, path, key string, body []byte, params *stripe.Params, v stripe.LastResponseSetter) error {
	return nil
}
func (b *testBackend) CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *stripe.Params, v stripe.LastResponseSetter) error {
	b.calledMultipart = true
	return nil
}
func (b *testBackend) SetMaxNetworkRetries(maxNetworkRetries int64) {}
func TestFileBackend(t *testing.T) {
	orig := stripe.GetBackend(stripe.UploadsBackend)
	b := &testBackend{calledMultipart: false}
	stripe.SetBackend(stripe.UploadsBackend, b)
	fileParams := &stripe.FileParams{}
	New(fileParams)
	assert.Equal(t, b.calledMultipart, true)
	stripe.SetBackend(stripe.UploadsBackend, orig)
}
func TestFileNew(t *testing.T) {
	f, err := os.Open("test_data.pdf")
	if err != nil {
		t.Errorf("Unable to open test file %v\n", err)
	}

	fileParams := &stripe.FileParams{
		Purpose:    stripe.String(string(stripe.FilePurposeDisputeEvidence)),
		FileReader: f,
		Filename:   stripe.String(f.Name()),
		FileLinkData: &stripe.FileFileLinkDataParams{
			Params: stripe.Params{
				Metadata: map[string]string{
					"foo": "bar",
				},
			},
			Create: stripe.Bool(true),
		},
	}

	file, err := New(fileParams)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
