package stripe

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"path/filepath"
)

// FileUploadParams is the set of parameters that can be used when creating a
// file upload.
// For more details see https://stripe.com/docs/api#create_file_upload.
type FileUploadParams struct {
	Params `form:"*"`

	// FileReader is a reader with the contents of the file that should be uploaded.
	FileReader io.Reader

	// Filename is just the name of the file without path information.
	Filename string

	Purpose FileUploadPurpose
}

// FileUploadListParams is the set of parameters that can be used when listing
// file uploads. For more details see https://stripe.com/docs/api#list_file_uploads.
type FileUploadListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Purpose      FileUploadPurpose `form:"purpose"`
}

// FileUploadPurpose is the purpose of a particular file upload. Allowed values
// are "dispute_evidence" and "identity_document".
type FileUploadPurpose string

// FileUpload is the resource representing a Stripe file upload.
// For more details see https://stripe.com/docs/api#file_uploads.
type FileUpload struct {
	Created int64             `json:"created"`
	ID      string            `json:"id"`
	Purpose FileUploadPurpose `json:"purpose"`
	Size    int64             `json:"size"`
	Type    string            `json:"type"`
	URL     string            `json:"url"`
}

// FileUploadList is a list of file uploads as retrieved from a list endpoint.
type FileUploadList struct {
	ListMeta
	Data []*FileUpload `json:"data"`
}

// AppendDetails adds the file upload details to an io.ReadWriter. It returns
// the boundary string for a multipart/form-data request and an error (if one
// exists).
func (f *FileUploadParams) AppendDetails(body io.ReadWriter) (string, error) {
	writer := multipart.NewWriter(body)
	var err error

	if len(f.Purpose) > 0 {
		err = writer.WriteField("purpose", string(f.Purpose))
		if err != nil {
			return "", err
		}
	}

	// Support both FileReader/Filename and File with
	// the former being the newer preferred version
	if f.FileReader != nil && f.Filename != "" {
		part, err := writer.CreateFormFile("file", filepath.Base(f.Filename))
		if err != nil {
			return "", err
		}

		_, err = io.Copy(part, f.FileReader)
		if err != nil {
			return "", err
		}
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	return writer.Boundary(), nil
}

// UnmarshalJSON handles deserialization of a FileUpload.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *FileUpload) UnmarshalJSON(data []byte) error {
	type file FileUpload
	var ff file
	err := json.Unmarshal(data, &ff)
	if err == nil {
		*f = FileUpload(ff)
	} else {
		// the id is surrounded by "\" characters, so strip them
		f.ID = string(data[1 : len(data)-1])
	}

	return nil
}
