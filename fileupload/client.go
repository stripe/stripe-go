// Package fileupload provides the file upload related APIs
package fileupload

import (
	"bytes"
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
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
	err = c.B.CallMultipart("POST", "/files", c.Key, boundary, body, &params.Params, upload)

	return upload, err
}

// Get returns the details of a file upload.
// For more details see https://stripe.com/docs/api#retrieve_file_upload.
func Get(id string, params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	return getC().Get(id, params)

}

func (c Client) Get(id string, params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	path := stripe.FormatURLPath("/files/%s", id)
	upload := &stripe.FileUpload{}
	err := c.B.Call("GET", path, c.Key, params, upload)
	return upload, err
}

// List returns a list of file uploads.
// For more details see https://stripe.com/docs/api#list_file_uploads.
func List(params *stripe.FileUploadListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.FileUploadListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.FileUploadList{}
		err := c.B.CallRaw("GET", "/files", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of FileUploads.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// FileUpload returns the most recent FileUpload visited by a call to Next.
func (i *Iter) FileUpload() *stripe.FileUpload {
	return i.Current().(*stripe.FileUpload)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.UploadsBackend), stripe.Key}
}
