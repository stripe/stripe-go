package stripe

import (
	"encoding/json"
	"os"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestFile_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v File
		err := json.Unmarshal([]byte(`"file_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "file_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := File{ID: "file_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "file_123", v.ID)
	}
}

func TestFileParams_GetBody(t *testing.T) {
	f, err := os.Open("file/test_data.pdf")
	if err != nil {
		t.Errorf("Unable to open test file %v\n", err)
	}

	p := &FileParams{
		FileReader: f,
		Filename:   String(f.Name()),
		FileLinkData: &FileFileLinkDataParams{
			Create: Bool(true),
		},
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
