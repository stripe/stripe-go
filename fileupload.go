package stripe

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// FileUploadParams is the set of parameters that can be used when creating a
// file upload.
// For more details see https://stripe.com/docs/api#create_file_upload.
type FileUploadParams struct {
	Purpose FileUploadPurpose
	File    os.File
}

// FileUploadPurpose is the purpose of a particular file upload. Allowed values
// are "dispute_evidence" and "identity_document".
type FileUploadPurpose string

// FileUpload is the resource representing a Stripe file upload.
// For more details see https://stripe.com/docs/api#file_uploads.
type FileUpload struct {
	ID      string            `json:"id"`
	Created int64             `json:"created"`
	Size    int64             `json:"size"`
	Purpose FileUploadPurpose `json:"purpose"`
	URL     string            `json:"url"`
	Mime    string            `json:"mimetype"`
}

func (f *FileUploadParams) AppendDetails(body io.Reader) error {
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(f.Name()))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, f.File)
	if err != nil {
		return nil, err
	}

	err = writer.WriteField("purpose", f.Purpose)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return nil
}

// UnmarshalJSON handles deserialization of a File.
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
