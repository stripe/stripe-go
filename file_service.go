//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1FileService is used to invoke /v1/files APIs.
type v1FileService struct {
	B        Backend
	BUploads Backend
	Key      string
}

// To upload a file to Stripe, you need to send a request of type multipart/form-data. Include the file you want to upload in the request, and the parameters for creating a file.
//
// All of Stripe's officially supported Client libraries support sending multipart/form-data.
func (c v1FileService) Create(ctx context.Context, params *FileCreateParams) (*File, error) {
	if params == nil {
		return nil, fmt.Errorf(
			"params cannot be nil, and params.Purpose and params.File must be set")
	}
	bodyBuffer, boundary, err := params.GetBody()
	if err != nil {
		return nil, err
	}
	params.Context = ctx
	file := &File{}
	err = c.BUploads.CallMultipart(http.MethodPost, "/v1/files", c.Key, boundary, bodyBuffer, &params.Params, file)
	return file, err
}

// Retrieves the details of an existing file object. After you supply a unique file ID, Stripe returns the corresponding file object. Learn how to [access file contents](https://docs.stripe.com/docs/file-upload#download-file-contents).
func (c v1FileService) Retrieve(ctx context.Context, id string, params *FileRetrieveParams) (*File, error) {
	if params == nil {
		params = &FileRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/files/%s", id)
	file := &File{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, file)
	return file, err
}

// Returns a list of the files that your account has access to. Stripe sorts and returns the files by their creation dates, placing the most recently created files at the top.
func (c v1FileService) List(ctx context.Context, listParams *FileListParams) Seq2[*File, error] {
	if listParams == nil {
		listParams = &FileListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*File, ListContainer, error) {
		list := &FileList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/files", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
