package stripe

import (
	"encoding/json"
	"os"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestFileUpload_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v FileUpload
		err := json.Unmarshal([]byte(`"file_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "file_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := FileUpload{ID: "file_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "file_123", v.ID)
	}
}

func TestFileUploadParams_GetBody(t *testing.T) {
	f, err := os.Open("fileupload/test_data.pdf")
	if err != nil {
		t.Errorf("Unable to open test file upload file %v\n", err)
	}

	p := &FileUploadParams{
		FileReader: f,
		Filename:   String(f.Name()),
	}

	buffer, boundary, err := p.GetBody()
	assert.NoError(t, err)

	assert.NotEqual(t, 0, buffer.Len())

	// Copied from the check performed by `multipart.Writer.SetBoundary`. A
	// very basic check that the string we got back indeed looks like a
	// boundary.
	//
	// rfc2046#section-5.1.1
	if len(boundary) < 1 || len(boundary) > 70 {
		t.Errorf("invalid boundary length")
	}
}
