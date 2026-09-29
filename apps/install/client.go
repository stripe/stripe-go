//
//
// File generated from our OpenAPI spec
//
//

// Package install provides the /v1/apps/installs APIs
package install

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/apps/installs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an app install. An account installs its own private app with its own key; public and testing installs are made from the Dashboard. An app developer or embedding platform acting on a connected account through Stripe-Account installs or reinstalls its app there. Creating an install for a private app that is already installed at the channel's current version with nothing pending returns the existing install.
func New(params *stripe.AppsInstallParams) (*stripe.AppsInstall, error) {
	return getC().New(params)
}

// Creates an app install. An account installs its own private app with its own key; public and testing installs are made from the Dashboard. An app developer or embedding platform acting on a connected account through Stripe-Account installs or reinstalls its app there. Creating an install for a private app that is already installed at the channel's current version with nothing pending returns the existing install.
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

// Reauthorizes an app install. The installer grants the permissions, content security policy entries, and endpoints that the latest published version of the app requests. An account reauthorizes its own installs on any channel with its own key; app developers and embedding platforms reauthorize installs on connected accounts through Stripe-Account. For private apps, install a new version from the Dashboard to grant its permissions.
func Update(id string, params *stripe.AppsInstallParams) (*stripe.AppsInstall, error) {
	return getC().Update(id, params)
}

// Reauthorizes an app install. The installer grants the permissions, content security policy entries, and endpoints that the latest published version of the app requests. An account reauthorizes its own installs on any channel with its own key; app developers and embedding platforms reauthorize installs on connected accounts through Stripe-Account. For private apps, install a new version from the Dashboard to grant its permissions.
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

// Returns a list of app installs. An app developer or embedding platform filtering by its own app sees the installs across the accounts that installed it; other callers see the installs on their own account. The key selects the environment: a live key lists live installs, a sandbox API key lists the installs on that sandbox, and the key of an app's managed sandbox filtering by app lists that app's installs across every sandbox. For existing accounts that still use legacy test mode, a test mode key lists legacy test mode installs.
func List(params *stripe.AppsInstallListParams) *Iter {
	return getC().List(params)
}

// Returns a list of app installs. An app developer or embedding platform filtering by its own app sees the installs across the accounts that installed it; other callers see the installs on their own account. The key selects the environment: a live key lists live installs, a sandbox API key lists the installs on that sandbox, and the key of an app's managed sandbox filtering by app lists that app's installs across every sandbox. For existing accounts that still use legacy test mode, a test mode key lists legacy test mode installs.
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
