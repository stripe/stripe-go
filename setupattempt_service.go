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

// v1SetupAttemptService is used to invoke /v1/setup_attempts APIs.
type v1SetupAttemptService struct {
	B   Backend
	Key string
}

// Returns a list of SetupAttempts that associate with a provided SetupIntent.
func (c v1SetupAttemptService) List(ctx context.Context, listParams *SetupAttemptListParams) Seq2[*SetupAttempt, error] {
	if listParams == nil {
		listParams = &SetupAttemptListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*SetupAttempt, ListContainer, error) {
		list := &SetupAttemptList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/setup_attempts", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
