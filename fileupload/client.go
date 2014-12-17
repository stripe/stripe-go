// Package fileupload provides the file upload related APIs
package fileupload

import (
	"bytes"
	"fmt"
	"net/url"

	stripe "github.com/stripe/stripe-go"
)

const (
	DisputeEvidenceFile stripe.FileUploadPurpose = "dispute_evidence"
	IdentityDocFile     stripe.FileUploadPurpose = "identity_document"
)

// Client is used to invoke file upload APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new file uploads.
// For more details see https://stripe.com/docs/api#create_file_upload.
func New(params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil, and params.Purpose and params.File must be set")
	}

	body := &bytes.Buffer{}
	boundary, err := params.AppendDetails(body)
	if err != nil {
		return nil, err
	}

	upload := &stripe.FileUpload{}
	err = c.B.CallMultipart("POST", "/files", c.Key, boundary, body, upload)

	return upload, err
}

// Get returns the details of a file upload.
// For more details see https://stripe.com/docs/api#retrieve_file_upload.
func Get(id string, params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	return getC().Get(id, params)

}

func (c Client) Get(id string, params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	upload := &stripe.FileUpload{}
	err := c.B.Call("GET", "/files/"+id, c.Key, body, upload)

	return upload, err

}

func getC() Client {
	return Client{stripe.GetBackend(stripe.UploadsBackend), stripe.Key}
}
