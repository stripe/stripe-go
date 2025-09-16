//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1FileLinkService is used to invoke /v1/file_links APIs.
type v1FileLinkService struct {
	B   Backend
	Key string
}

// Creates a new file link object.
func (c v1FileLinkService) Create(ctx context.Context, params *FileLinkCreateParams) (*FileLink, error) {
	if params == nil {
		params = &FileLinkCreateParams{}
	}
	params.Context = ctx
	filelink := &FileLink{}
	err := c.B.Call(http.MethodPost, "/v1/file_links", c.Key, params, filelink)
	return filelink, err
}

// Retrieves the file link with the given ID.
func (c v1FileLinkService) Retrieve(ctx context.Context, id string, params *FileLinkRetrieveParams) (*FileLink, error) {
	if params == nil {
		params = &FileLinkRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/file_links/%s", id)
	filelink := &FileLink{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, filelink)
	return filelink, err
}

// Updates an existing file link object. Expired links can no longer be updated.
func (c v1FileLinkService) Update(ctx context.Context, id string, params *FileLinkUpdateParams) (*FileLink, error) {
	if params == nil {
		params = &FileLinkUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/file_links/%s", id)
	filelink := &FileLink{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, filelink)
	return filelink, err
}

// Returns a list of file links.
func (c v1FileLinkService) List(ctx context.Context, listParams *FileLinkListParams) Seq2[*FileLink, error] {
	if listParams == nil {
		listParams = &FileLinkListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*FileLink, ListContainer, error) {
		list := &FileLinkList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/file_links", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
