// Package fileupload provides the file upload related APIs
package fileupload

import (
	"bytes"
	"fmt"
	"io"
	"net/url"

	stripe "github.com/stripe/stripe-go"
)

const (
	uploadsURL = "https://uploads.stripe.com/v1"
)

// Client is used to invoke file upload related APIs
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
	body := &bytes.Buffer{}
	if params != nil {
		params.AppendDetails(body)
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "multipart/form-data"

	upload := &stripe.FileUpload{}
	err := c.B.AbstractCall("POST", uploadsURL+"/files", c.Key, body, headers, upload)

	return upload, err
}

// Get returns the details of a file upload.
// For more details see https://stripe.com/docs/api#retrieve_file_upload.
func Get(id string, params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	var form *url.Values
	path := fmt.Sprintf("%v/files/%v", uploadsURL, id)

	if params != nil {
		form = &url.Values{}
		params.AppendTo(form)
	}
	if form != nil && len(*form) > 0 {
		data := form.Encode()
		path += "?" + data
	}

	var headers map[string]string
	var body io.Reader
	upload := &stripe.FileUpload{}
	err := c.B.AbstractCall("GET", path, c.Key, body, headers, upload)

	return upload, err
}

func getC() Client {
	return Client{stripe.GetBackend(), stripe.Key}
}
