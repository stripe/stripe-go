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

// v1AppsSecretService is used to invoke /v1/apps/secrets APIs.
type v1AppsSecretService struct {
	B   Backend
	Key string
}

// Create or replace a secret in the secret store.
func (c v1AppsSecretService) Create(ctx context.Context, params *AppsSecretCreateParams) (*AppsSecret, error) {
	if params == nil {
		params = &AppsSecretCreateParams{}
	}
	params.Context = ctx
	secret := &AppsSecret{}
	err := c.B.Call(http.MethodPost, "/v1/apps/secrets", c.Key, params, secret)
	return secret, err
}

// Deletes a secret from the secret store by name and scope.
func (c v1AppsSecretService) DeleteWhere(ctx context.Context, params *AppsSecretDeleteWhereParams) (*AppsSecret, error) {
	if params == nil {
		params = &AppsSecretDeleteWhereParams{}
	}
	params.Context = ctx
	secret := &AppsSecret{}
	err := c.B.Call(
		http.MethodPost, "/v1/apps/secrets/delete", c.Key, params, secret)
	return secret, err
}

// Finds a secret in the secret store by name and scope.
func (c v1AppsSecretService) Find(ctx context.Context, params *AppsSecretFindParams) (*AppsSecret, error) {
	if params == nil {
		params = &AppsSecretFindParams{}
	}
	params.Context = ctx
	secret := &AppsSecret{}
	err := c.B.Call(
		http.MethodGet, "/v1/apps/secrets/find", c.Key, params, secret)
	return secret, err
}

// List all secrets stored on the given scope.
func (c v1AppsSecretService) List(ctx context.Context, listParams *AppsSecretListParams) Seq2[*AppsSecret, error] {
	if listParams == nil {
		listParams = &AppsSecretListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*AppsSecret, ListContainer, error) {
		list := &AppsSecretList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/apps/secrets", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
