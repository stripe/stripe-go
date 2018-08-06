// Package fileupload provides the file upload related APIs
package fileupload

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke file upload APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new file upload.
func New(params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	return getC().New(params)
}

// New creates a new file upload.
func (c Client) New(params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil, and params.Purpose and params.File must be set")
	}

	bodyBuffer, boundary, err := params.GetBody()
	if err != nil {
		return nil, err
	}

	upload := &stripe.FileUpload{}
	err = c.B.CallMultipart(http.MethodPost, "/files", c.Key, boundary, bodyBuffer, &params.Params, upload)

	return upload, err
}

// Get returns the details of a file upload.
func Get(id string, params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	return getC().Get(id, params)

}

// Get returns the details of a file upload.
func (c Client) Get(id string, params *stripe.FileUploadParams) (*stripe.FileUpload, error) {
	path := stripe.FormatURLPath("/files/%s", id)
	upload := &stripe.FileUpload{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, upload)
	return upload, err
}

// List returns a list of file uploads.
func List(params *stripe.FileUploadListParams) *Iter {
	return getC().List(params)
}

// List returns a list of file uploads.
func (c Client) List(listParams *stripe.FileUploadListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.FileUploadList{}
		err := c.B.CallRaw(http.MethodGet, "/files", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for file uploads.
type Iter struct {
	*stripe.Iter
}

// FileUpload returns the file upload which the iterator is currently pointing to.
func (i *Iter) FileUpload() *stripe.FileUpload {
	return i.Current().(*stripe.FileUpload)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.UploadsBackend), stripe.Key}
}
