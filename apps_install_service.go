//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v86/form"
)

// v1AppsInstallService is used to invoke /v1/apps/installs APIs.
type v1AppsInstallService struct {
	B   Backend
	Key string
}

// Creates an app install. An account installs its own private app with its own key; public and testing installs are made from the Dashboard. An app developer or embedding platform acting on a connected account through Stripe-Account installs or reinstalls its app there. Creating an install for a private app that is already installed at the channel's current version with nothing pending returns the existing install.
func (c v1AppsInstallService) Create(ctx context.Context, params *AppsInstallCreateParams) (*AppsInstall, error) {
	if params == nil {
		params = &AppsInstallCreateParams{}
	}
	params.Context = ctx
	install := &AppsInstall{}
	err := c.B.Call(http.MethodPost, "/v1/apps/installs", c.Key, params, install)
	return install, err
}

// Retrieves an app install. The installing account, the app's developer (with the keys of the account that owns the app or of the app's managed sandbox), and the embedding platform that created the install can retrieve it.
func (c v1AppsInstallService) Retrieve(ctx context.Context, id string, params *AppsInstallRetrieveParams) (*AppsInstall, error) {
	if params == nil {
		params = &AppsInstallRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/apps/installs/%s", id)
	install := &AppsInstall{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, install)
	return install, err
}

// Reauthorizes an app install. The installer grants the permissions, content security policy entries, and endpoints that the latest published version of the app requests. An account reauthorizes its own installs on any channel with its own key; app developers and embedding platforms reauthorize installs on connected accounts through Stripe-Account. For private apps, install a new version from the Dashboard to grant its permissions.
func (c v1AppsInstallService) Update(ctx context.Context, id string, params *AppsInstallUpdateParams) (*AppsInstall, error) {
	if params == nil {
		params = &AppsInstallUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/apps/installs/%s", id)
	install := &AppsInstall{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, install)
	return install, err
}

// Uninstalls an app from the account that installed it.
func (c v1AppsInstallService) Uninstall(ctx context.Context, id string, params *AppsInstallUninstallParams) (*AppsInstall, error) {
	if params == nil {
		params = &AppsInstallUninstallParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/apps/installs/%s/uninstall", id)
	install := &AppsInstall{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, install)
	return install, err
}

// Returns a list of app installs. An app developer or embedding platform filtering by its own app sees the installs across the accounts that installed it; other callers see the installs on their own account. The key selects the environment: a live key lists live installs, a sandbox API key lists the installs on that sandbox, and the key of an app's managed sandbox filtering by app lists that app's installs across every sandbox. For existing accounts that still use legacy test mode, a test mode key lists legacy test mode installs.
func (c v1AppsInstallService) List(ctx context.Context, listParams *AppsInstallListParams) *V1List[*AppsInstall] {
	if listParams == nil {
		listParams = &AppsInstallListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*AppsInstall], error) {
		list := &v1Page[*AppsInstall]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/apps/installs", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
