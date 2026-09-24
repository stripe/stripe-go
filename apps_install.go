//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The distribution channel associated with the app install.
type AppsInstallChannel string

// List of values that AppsInstallChannel can take
const (
	AppsInstallChannelPrivateLive AppsInstallChannel = "private_live"
	AppsInstallChannelPrivateTest AppsInstallChannel = "private_test"
	AppsInstallChannelPublic      AppsInstallChannel = "public"
	AppsInstallChannelReview      AppsInstallChannel = "review"
	AppsInstallChannelTesting     AppsInstallChannel = "testing"
)

// The status of the app install.
type AppsInstallStatus string

// List of values that AppsInstallStatus can take
const (
	AppsInstallStatusInstallFailed   AppsInstallStatus = "install_failed"
	AppsInstallStatusInstalled       AppsInstallStatus = "installed"
	AppsInstallStatusInstalling      AppsInstallStatus = "installing"
	AppsInstallStatusUninstallFailed AppsInstallStatus = "uninstall_failed"
	AppsInstallStatusUninstalling    AppsInstallStatus = "uninstalling"
)

// Returns a list of app installs. An app developer or embedding platform filtering by its own app sees the installs across the accounts that installed it; other callers see the installs on their own account. The key selects the environment: a live key lists live installs, a sandbox API key lists the installs on that sandbox, and the key of an app's managed sandbox filtering by app lists that app's installs across every sandbox. For existing accounts that still use legacy test mode, a test mode key lists legacy test mode installs.
type AppsInstallListParams struct {
	ListParams `form:"*"`
	// Only return installs made by this account. Only useful to app developers and embedding platforms, whose lists span the accounts that installed their app.
	Account *string `form:"account" json:"account,omitempty"`
	// Only return installs for the app specified by this app ID.
	App *string `form:"app" json:"app,omitempty"`
	// Only return installs whose installer must authorize pending permissions, content security policy entries, or endpoints.
	ApprovalRequired *bool `form:"approval_required" json:"approval_required,omitempty"`
	// Only return installs in the distribution channel specified by this channel name.
	Channel *string `form:"channel" json:"channel,omitempty"`
	// Only return app installs that were created during the given date interval.
	Created *int64 `form:"created" json:"created,omitempty"`
	// Only return installs created by the embedding platform specified by this account ID.
	CreatedBy *string `form:"created_by" json:"created_by,omitempty"`
	// Only return app installs that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created" json:"-"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Only return installs with the given status.
	Status *string `form:"status" json:"status,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *AppsInstallListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates an app install. An account installs its own private app with its own key; public and testing installs are made from the Dashboard. An app developer or embedding platform acting on a connected account through Stripe-Account installs or reinstalls its app there. Creating an install for a private app that is already installed at the channel's current version with nothing pending returns the existing install.
type AppsInstallParams struct {
	Params `form:"*"`
	// The ID of the app to install.
	App *string `form:"app" json:"app,omitempty"`
	// The distribution channel to install from. Defaults to `public`. A private app must be installed on `private_test` or `private_live`, matching the mode of the API key.
	Channel *string `form:"channel" json:"channel,omitempty"`
	// For OAuth apps, the PKCE code challenge used to issue the `auth_code` returned on the install. Must be 43 to 128 characters and contain only letters, numbers, `-`, `.`, `_`, and `~`. Only applies to installs made by the app developer or an embedding platform; ignored when an account installs its own private app.
	CodeChallenge *string `form:"code_challenge" json:"code_challenge,omitempty"`
	// The method used to derive `code_challenge`. Required when `code_challenge` is provided, and must be `S256`.
	CodeChallengeMethod *string `form:"code_challenge_method" json:"code_challenge_method,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *AppsInstallParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Uninstalls an app from the account that installed it.
type AppsInstallUninstallParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *AppsInstallUninstallParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates an app install. An account installs its own private app with its own key; public and testing installs are made from the Dashboard. An app developer or embedding platform acting on a connected account through Stripe-Account installs or reinstalls its app there. Creating an install for a private app that is already installed at the channel's current version with nothing pending returns the existing install.
type AppsInstallCreateParams struct {
	Params `form:"*"`
	// The ID of the app to install.
	App *string `form:"app" json:"app"`
	// The distribution channel to install from. Defaults to `public`. A private app must be installed on `private_test` or `private_live`, matching the mode of the API key.
	Channel *string `form:"channel" json:"channel,omitempty"`
	// For OAuth apps, the PKCE code challenge used to issue the `auth_code` returned on the install. Must be 43 to 128 characters and contain only letters, numbers, `-`, `.`, `_`, and `~`. Only applies to installs made by the app developer or an embedding platform; ignored when an account installs its own private app.
	CodeChallenge *string `form:"code_challenge" json:"code_challenge,omitempty"`
	// The method used to derive `code_challenge`. Required when `code_challenge` is provided, and must be `S256`.
	CodeChallengeMethod *string `form:"code_challenge_method" json:"code_challenge_method,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *AppsInstallCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an app install. The installing account, the app's developer (with the keys of the account that owns the app or of the app's managed sandbox), and the embedding platform that created the install can retrieve it.
type AppsInstallRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *AppsInstallRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Reauthorizes an app install. The installer grants the permissions, content security policy entries, and endpoints that the latest published version of the app requests. An account reauthorizes its own installs on any channel with its own key; app developers and embedding platforms reauthorize installs on connected accounts through Stripe-Account. For private apps, install a new version from the Dashboard to grant its permissions.
type AppsInstallUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *AppsInstallUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type AppsInstallContentSecurityPolicyGranted struct {
	// The URLs that the app can make network requests to.
	ConnectSrc []string `json:"connect_src"`
	// The URLs that the app can load images from.
	ImageSrc []string `json:"image_src"`
}
type AppsInstallContentSecurityPolicyPending struct {
	// The URLs that the app can make network requests to.
	ConnectSrc []string `json:"connect_src"`
	// The URLs that the app can load images from.
	ImageSrc []string `json:"image_src"`
}

// An app install represents a Stripe App that is installed on an account. It reports the permissions,
// content security policy entries, and endpoints that the installing account has authorized, along with any
// that the app's latest version requests but the account has not authorized yet. Use the Install API to
// install, reauthorize, and uninstall apps, and to check the state of existing installs.
type AppsInstall struct {
	APIResource
	// The ID of the account that the app install belongs to.
	Account string `json:"account"`
	// The ID of the app installed.
	App string `json:"app"`
	// Whether the installer must authorize pending permissions, content security policy entries, or endpoints. For private apps, `approval_required` stays `false`. Install a new version from the Dashboard to grant its permissions.
	ApprovalRequired bool `json:"approval_required"`
	// The authorization code for an oauth app install.
	AuthCode string `json:"auth_code"`
	// The distribution channel associated with the app install.
	Channel                      AppsInstallChannel                       `json:"channel"`
	ContentSecurityPolicyGranted *AppsInstallContentSecurityPolicyGranted `json:"content_security_policy_granted"`
	ContentSecurityPolicyPending *AppsInstallContentSecurityPolicyPending `json:"content_security_policy_pending"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The ID of the embedding platform that created the install, if applicable.
	CreatedBy string `json:"created_by"`
	// The endpoint URLs authorized by the installer.
	EndpointsGranted []string `json:"endpoints_granted"`
	// The endpoint URLs requested by the latest app version that the installer has not authorized.
	EndpointsPending []string `json:"endpoints_pending"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The permissions authorized by the installer.
	PermissionsGranted []string `json:"permissions_granted"`
	// The permissions requested by the latest app version that the installer has not authorized.
	PermissionsPending []string `json:"permissions_pending"`
	// The status of the app install.
	Status AppsInstallStatus `json:"status"`
}

// AppsInstallList is a list of Installs as retrieved from a list endpoint.
type AppsInstallList struct {
	APIResource
	ListMeta
	Data []*AppsInstall `json:"data"`
}
