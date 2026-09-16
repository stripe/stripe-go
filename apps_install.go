//
//
// File generated from our OpenAPI spec
//
//

package stripe

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

type AppsInstallAuthorizedContentSecurityPolicy struct {
	ConnectSrc []string `json:"connect_src"`
	ImageSrc   []string `json:"image_src"`
	Purpose    string   `json:"purpose"`
}

// The content security policy entries authorized by the installer.
type AppsInstallContentSecurityPolicyGranted struct {
	ConnectSrc []string `json:"connect_src"`
	ImageSrc   []string `json:"image_src"`
}
type AppsInstallContentSecurityPolicyPending struct {
	ConnectSrc []string `json:"connect_src"`
	ImageSrc   []string `json:"image_src"`
}

// An object representing an app installation.
type AppsInstall struct {
	// The ID of the account that the app install belongs to.
	Account string `json:"account"`
	// The ID of the app installed.
	App string `json:"app"`
	// Whether the installer must authorize pending permissions, content security policy entries, or endpoints.
	ApprovalRequired bool `json:"approval_required"`
	// The authorization code for an oauth app install.
	AuthCode                        string                                      `json:"auth_code"`
	AuthorizedContentSecurityPolicy *AppsInstallAuthorizedContentSecurityPolicy `json:"authorized_content_security_policy"`
	// The endpoint URLs authorized by the installer.
	AuthorizedEndpoints []string `json:"authorized_endpoints"`
	// The permissions authorized by the installer.
	AuthorizedPermissions []string `json:"authorized_permissions"`
	// The distribution channel associated with the app install.
	Channel string `json:"channel"`
	// The content security policy entries authorized by the installer.
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
	State string `json:"state"`
	// The status of the app install.
	Status AppsInstallStatus `json:"status"`
}
