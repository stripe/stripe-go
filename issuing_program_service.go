//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1IssuingProgramService is used to invoke /v1/issuing/programs APIs.
type v1IssuingProgramService struct {
	B   Backend
	Key string
}

// Create a Program object.
func (c v1IssuingProgramService) Create(ctx context.Context, params *IssuingProgramCreateParams) (*IssuingProgram, error) {
	if params == nil {
		params = &IssuingProgramCreateParams{}
	}
	params.Context = ctx
	program := &IssuingProgram{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/programs", c.Key, params, program)
	return program, err
}

// Retrieves the program specified by the given id.
func (c v1IssuingProgramService) Retrieve(ctx context.Context, id string, params *IssuingProgramRetrieveParams) (*IssuingProgram, error) {
	if params == nil {
		params = &IssuingProgramRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/programs/%s", id)
	program := &IssuingProgram{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, program)
	return program, err
}

// Updates a Program object.
func (c v1IssuingProgramService) Update(ctx context.Context, id string, params *IssuingProgramUpdateParams) (*IssuingProgram, error) {
	if params == nil {
		params = &IssuingProgramUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/programs/%s", id)
	program := &IssuingProgram{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, program)
	return program, err
}

// List all of the programs the given Issuing user has access to.
func (c v1IssuingProgramService) List(ctx context.Context, listParams *IssuingProgramListParams) Seq2[*IssuingProgram, error] {
	if listParams == nil {
		listParams = &IssuingProgramListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*IssuingProgram], error) {
		list := &v1Page[*IssuingProgram]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/programs", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
