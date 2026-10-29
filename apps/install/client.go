//
//
// File generated from our OpenAPI spec
//
//

// Package install provides the /v1/apps/installs APIs
package install

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
	"github.com/stripe/stripe-go/v87/form"
)

// Client is used to invoke /v1/apps/installs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an app install. An account installs its own private app with its own key; public and testing installs are made from the Dashboard. An app developer acting on a connected account through Stripe-Account installs or reinstalls its app there, and an embedding platform can do the same once the app's developer approves its request to embed the app. For a private app, creating an install installs the newest completed upload; when that version is already installed with nothing pending, the existing install is returned.
func New(params *stripe.AppsInstallParams) (*stripe.AppsInstall, error) {
	return getC().New(params)
}

// Creates an app install. An account installs its own private app with its own key; public and testing installs are made from the Dashboard. An app developer acting on a connected account through Stripe-Account installs or reinstalls its app there, and an embedding platform can do the same once the app's developer approves its request to embed the app. For a private app, creating an install installs the newest completed upload; when that version is already installed with nothing pending, the existing install is returned.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.AppsInstallParams) (*stripe.AppsInstall, error) {
	install := &stripe.AppsInstall{}
	err := c.B.Call(http.MethodPost, "/v1/apps/installs", c.Key, params, install)
	return install, err
}

// Retrieves an app install. The installing account, the app's developer (with the keys of the account that owns the app or of the app's managed sandbox), and the embedding platform that created the install can retrieve it.
func Get(id string, params *stripe.AppsInstallParams) (*stripe.AppsInstall, error) {
	return getC().Get(id, params)
}

// Retrieves an app install. The installing account, the app's developer (with the keys of the account that owns the app or of the app's managed sandbox), and the embedding platform that created the install can retrieve it.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.AppsInstallParams) (*stripe.AppsInstall, error) {
	path := stripe.FormatURLPath("/v1/apps/installs/%s", id)
	install := &stripe.AppsInstall{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, install)
	return install, err
}

// Reauthorizes an app install. The installer grants the permissions, content security policy entries, and endpoints that the version being installed requests. An account reauthorizes its own installs on any channel with its own key, which grants all of that access, so only give app_install_write to keys that may approve an app's access. App developers and embedding platforms reauthorize installs on connected accounts through Stripe-Account. An app developer can't grant new access. An embedding platform can grant new access only once the app's developer approves its request to embed the app. For private apps, the version being installed is the newest completed upload.
func Update(id string, params *stripe.AppsInstallParams) (*stripe.AppsInstall, error) {
	return getC().Update(id, params)
}

// Reauthorizes an app install. The installer grants the permissions, content security policy entries, and endpoints that the version being installed requests. An account reauthorizes its own installs on any channel with its own key, which grants all of that access, so only give app_install_write to keys that may approve an app's access. App developers and embedding platforms reauthorize installs on connected accounts through Stripe-Account. An app developer can't grant new access. An embedding platform can grant new access only once the app's developer approves its request to embed the app. For private apps, the version being installed is the newest completed upload.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.AppsInstallParams) (*stripe.AppsInstall, error) {
	path := stripe.FormatURLPath("/v1/apps/installs/%s", id)
	install := &stripe.AppsInstall{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, install)
	return install, err
}

// Uninstalls an app from the account that installed it.
func Uninstall(id string, params *stripe.AppsInstallUninstallParams) (*stripe.AppsInstall, error) {
	return getC().Uninstall(id, params)
}

// Uninstalls an app from the account that installed it.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Uninstall(id string, params *stripe.AppsInstallUninstallParams) (*stripe.AppsInstall, error) {
	path := stripe.FormatURLPath("/v1/apps/installs/%s/uninstall", id)
	install := &stripe.AppsInstall{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, install)
	return install, err
}

// Returns a list of app installs. An app developer filtering by its own app with its own key sees that app's installs across the accounts that installed it. An app developer acting on a connected account through Stripe-Account and filtering by its app sees that account's installs of the app, and an embedding platform acting on a connected account sees only the installs it created there. Other callers see the installs on their own account. A live key lists live installs and a test key lists test installs; the key of an app's managed sandbox filtering by app lists that app's installs across every sandbox.
func List(params *stripe.AppsInstallListParams) *Iter {
	return getC().List(params)
}

// Returns a list of app installs. An app developer filtering by its own app with its own key sees that app's installs across the accounts that installed it. An app developer acting on a connected account through Stripe-Account and filtering by its app sees that account's installs of the app, and an embedding platform acting on a connected account sees only the installs it created there. Other callers see the installs on their own account. A live key lists live installs and a test key lists test installs; the key of an app's managed sandbox filtering by app lists that app's installs across every sandbox.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.AppsInstallListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.AppsInstallList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/apps/installs", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for apps installs.
type Iter struct {
	*stripe.Iter
}

// AppsInstall returns the apps install which the iterator is currently pointing to.
func (i *Iter) AppsInstall() *stripe.AppsInstall {
	return i.Current().(*stripe.AppsInstall)
}

// AppsInstallList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) AppsInstallList() *stripe.AppsInstallList {
	return i.List().(*stripe.AppsInstallList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
