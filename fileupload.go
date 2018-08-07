package stripe

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"path/filepath"
)

// FileUploadPurpose is the purpose of a particular file upload.
type FileUploadPurpose string

// List of values that FileUploadPurpose can take.
const (
	FileUploadPurposeDisputeEvidence  FileUploadPurpose = "dispute_evidence"
	FileUploadPurposeIdentityDocument FileUploadPurpose = "identity_document"
)

// FileUploadParams is the set of parameters that can be used when creating a
// file upload.
// For more details see https://stripe.com/docs/api#create_file_upload.
type FileUploadParams struct {
	Params `form:"*"`

	// FileReader is a reader with the contents of the file that should be uploaded.
	FileReader io.Reader

	// Filename is just the name of the file without path information.
	Filename *string

	Purpose *string
}

// FileUploadListParams is the set of parameters that can be used when listing
// file uploads. For more details see https://stripe.com/docs/api#list_file_uploads.
type FileUploadListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Purpose      *string           `form:"purpose"`
}

// FileUpload is the resource representing a Stripe file upload.
// For more details see https://stripe.com/docs/api#file_uploads.
type FileUpload struct {
	Created  int64             `json:"created"`
	ID       string            `json:"id"`
	Filename string            `json:"filename"`
	Links    *FileLinkList     `json:"links"`
	Purpose  FileUploadPurpose `json:"purpose"`
	Size     int64             `json:"size"`
	Type     string            `json:"type"`
	URL      string            `json:"url"`
}

// FileUploadList is a list of file uploads as retrieved from a list endpoint.
type FileUploadList struct {
	ListMeta
	Data []*FileUpload `json:"data"`
}

// GetBody gets an appropriate multipart form payload to use in a request body
// to create a new file.
func (f *FileUploadParams) GetBody() (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if f.Purpose != nil {
		err := writer.WriteField("purpose", StringValue(f.Purpose))
		if err != nil {
			return nil, "", err
		}
	}

	if f.FileReader != nil && f.Filename != nil {
		part, err := writer.CreateFormFile("file", filepath.Base(StringValue(f.Filename)))
		if err != nil {
			return nil, "", err
		}

		_, err = io.Copy(part, f.FileReader)
		if err != nil {
			return nil, "", err
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, "", err
	}

	return body, writer.Boundary(), nil
}

// UnmarshalJSON handles deserialization of a FileUpload.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *FileUpload) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		f.ID = id
		return nil
	}

	type fileUpload FileUpload
	var v fileUpload
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*f = FileUpload(v)
	return nil
}
