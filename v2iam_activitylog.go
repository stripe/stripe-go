//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of actor.
type V2IamActivityLogActorType string

// List of values that V2IamActivityLogActorType can take
const (
	V2IamActivityLogActorTypeAPIKey       V2IamActivityLogActorType = "api_key"
	V2IamActivityLogActorTypeStripeAction V2IamActivityLogActorType = "stripe_action"
	V2IamActivityLogActorTypeUser         V2IamActivityLogActorType = "user"
)

// The type of entity.
type V2IamActivityLogDetailsAPIKeyManagedByType string

// List of values that V2IamActivityLogDetailsAPIKeyManagedByType can take
const (
	V2IamActivityLogDetailsAPIKeyManagedByTypeApplication V2IamActivityLogDetailsAPIKeyManagedByType = "application"
)

// Type of the API key.
type V2IamActivityLogDetailsAPIKeyType string

// List of values that V2IamActivityLogDetailsAPIKeyType can take
const (
	V2IamActivityLogDetailsAPIKeyTypePublishableKey V2IamActivityLogDetailsAPIKeyType = "publishable_key"
	V2IamActivityLogDetailsAPIKeyTypeSecretKey      V2IamActivityLogDetailsAPIKeyType = "secret_key"
)

// Type of challenge used for the authentication.
type V2IamActivityLogDetailsAuthenticationChallengeType string

// List of values that V2IamActivityLogDetailsAuthenticationChallengeType can take
const (
	V2IamActivityLogDetailsAuthenticationChallengeTypeExternalAccountCode   V2IamActivityLogDetailsAuthenticationChallengeType = "external_account_code"
	V2IamActivityLogDetailsAuthenticationChallengeTypeOauth                 V2IamActivityLogDetailsAuthenticationChallengeType = "oauth"
	V2IamActivityLogDetailsAuthenticationChallengeTypePreviousAccountNumber V2IamActivityLogDetailsAuthenticationChallengeType = "previous_account_number"
	V2IamActivityLogDetailsAuthenticationChallengeTypeReverseSms            V2IamActivityLogDetailsAuthenticationChallengeType = "reverse_sms"
	V2IamActivityLogDetailsAuthenticationChallengeTypeSms                   V2IamActivityLogDetailsAuthenticationChallengeType = "sms"
	V2IamActivityLogDetailsAuthenticationChallengeTypeStripeIdentity        V2IamActivityLogDetailsAuthenticationChallengeType = "stripe_identity"
	V2IamActivityLogDetailsAuthenticationChallengeTypeTotp                  V2IamActivityLogDetailsAuthenticationChallengeType = "totp"
	V2IamActivityLogDetailsAuthenticationChallengeTypeWebauthn              V2IamActivityLogDetailsAuthenticationChallengeType = "webauthn"
)

// Surface where the authentication occurred.
type V2IamActivityLogDetailsAuthenticationSurface string

// List of values that V2IamActivityLogDetailsAuthenticationSurface can take
const (
	V2IamActivityLogDetailsAuthenticationSurfaceDashboard V2IamActivityLogDetailsAuthenticationSurface = "dashboard"
	V2IamActivityLogDetailsAuthenticationSurfaceExpress   V2IamActivityLogDetailsAuthenticationSurface = "express"
)

// SSO enforcement level.
type V2IamActivityLogDetailsSsoMandate string

// List of values that V2IamActivityLogDetailsSsoMandate can take
const (
	V2IamActivityLogDetailsSsoMandateOff      V2IamActivityLogDetailsSsoMandate = "off"
	V2IamActivityLogDetailsSsoMandateOptional V2IamActivityLogDetailsSsoMandate = "optional"
	V2IamActivityLogDetailsSsoMandateRequired V2IamActivityLogDetailsSsoMandate = "required"
)

// The action group type of the activity log entry.
type V2IamActivityLogDetailsType string

// List of values that V2IamActivityLogDetailsType can take
const (
	V2IamActivityLogDetailsTypeAccountSecurity V2IamActivityLogDetailsType = "account_security"
	V2IamActivityLogDetailsTypeAPIKey          V2IamActivityLogDetailsType = "api_key"
	V2IamActivityLogDetailsTypeAuthentication  V2IamActivityLogDetailsType = "authentication"
	V2IamActivityLogDetailsTypeIssuing         V2IamActivityLogDetailsType = "issuing"
	V2IamActivityLogDetailsTypePayout          V2IamActivityLogDetailsType = "payout"
	V2IamActivityLogDetailsTypeScim            V2IamActivityLogDetailsType = "scim"
	V2IamActivityLogDetailsTypeSso             V2IamActivityLogDetailsType = "sso"
	V2IamActivityLogDetailsTypeUserAccess      V2IamActivityLogDetailsType = "user_access"
	V2IamActivityLogDetailsTypeUserInvite      V2IamActivityLogDetailsType = "user_invite"
	V2IamActivityLogDetailsTypeUserProfile     V2IamActivityLogDetailsType = "user_profile"
	V2IamActivityLogDetailsTypeUserRoles       V2IamActivityLogDetailsType = "user_roles"
)

// Type of authentication factor.
type V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType string

// List of values that V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType can take
const (
	V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorTypeBackupCode V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType = "backup_code"
	V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorTypeEmailCode  V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType = "email_code"
	V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorTypeOauth      V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType = "oauth"
	V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorTypePasskey    V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType = "passkey"
	V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorTypePassword   V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType = "password"
	V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorTypePhoneCode  V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType = "phone_code"
	V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorTypeSaml       V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType = "saml"
	V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorTypeSms        V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType = "sms"
	V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorTypeTotp       V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType = "totp"
	V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorTypeWebAuthn   V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType = "web_authn"
)

// Type of authentication factor.
type V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType string

// List of values that V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType can take
const (
	V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorTypeBackupCode V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType = "backup_code"
	V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorTypeEmailCode  V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType = "email_code"
	V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorTypeOauth      V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType = "oauth"
	V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorTypePasskey    V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType = "passkey"
	V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorTypePassword   V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType = "password"
	V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorTypePhoneCode  V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType = "phone_code"
	V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorTypeSaml       V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType = "saml"
	V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorTypeSms        V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType = "sms"
	V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorTypeTotp       V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType = "totp"
	V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorTypeWebAuthn   V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType = "web_authn"
)

// Risk level for the user access action.
type V2IamActivityLogDetailsUserAccessRiskLevel string

// List of values that V2IamActivityLogDetailsUserAccessRiskLevel can take
const (
	V2IamActivityLogDetailsUserAccessRiskLevelHigh   V2IamActivityLogDetailsUserAccessRiskLevel = "high"
	V2IamActivityLogDetailsUserAccessRiskLevelLow    V2IamActivityLogDetailsUserAccessRiskLevel = "low"
	V2IamActivityLogDetailsUserAccessRiskLevelMedium V2IamActivityLogDetailsUserAccessRiskLevel = "medium"
)

// Type of risk signal.
type V2IamActivityLogDetailsUserAccessRiskSignalType string

// List of values that V2IamActivityLogDetailsUserAccessRiskSignalType can take
const (
	V2IamActivityLogDetailsUserAccessRiskSignalTypeNovelDevice V2IamActivityLogDetailsUserAccessRiskSignalType = "novel_device"
)

// Surface where the user access action started.
type V2IamActivityLogDetailsUserAccessSurface string

// List of values that V2IamActivityLogDetailsUserAccessSurface can take
const (
	V2IamActivityLogDetailsUserAccessSurfaceDashboard V2IamActivityLogDetailsUserAccessSurface = "dashboard"
	V2IamActivityLogDetailsUserAccessSurfaceExpress   V2IamActivityLogDetailsUserAccessSurface = "express"
)

// Source of the role change.
type V2IamActivityLogDetailsUserRolesSource string

// List of values that V2IamActivityLogDetailsUserRolesSource can take
const (
	V2IamActivityLogDetailsUserRolesSourceDashboard V2IamActivityLogDetailsUserRolesSource = "dashboard"
	V2IamActivityLogDetailsUserRolesSourceScim      V2IamActivityLogDetailsUserRolesSource = "scim"
	V2IamActivityLogDetailsUserRolesSourceSso       V2IamActivityLogDetailsUserRolesSource = "sso"
)

// Type of the object.
type V2IamActivityLogRelatedObjectType string

// List of values that V2IamActivityLogRelatedObjectType can take
const (
	V2IamActivityLogRelatedObjectTypeBalanceTransfer   V2IamActivityLogRelatedObjectType = "balance_transfer"
	V2IamActivityLogRelatedObjectTypeBankAccount       V2IamActivityLogRelatedObjectType = "bank_account"
	V2IamActivityLogRelatedObjectTypeBlockchainAddress V2IamActivityLogRelatedObjectType = "blockchain_address"
	V2IamActivityLogRelatedObjectTypeCard              V2IamActivityLogRelatedObjectType = "card"
	V2IamActivityLogRelatedObjectTypeIssuingCard       V2IamActivityLogRelatedObjectType = "issuing.card"
	V2IamActivityLogRelatedObjectTypeIssuingCardholder V2IamActivityLogRelatedObjectType = "issuing.cardholder"
	V2IamActivityLogRelatedObjectTypeIssuingDispute    V2IamActivityLogRelatedObjectType = "issuing.dispute"
)

// The type of action that was performed.
type V2IamActivityLogType string

// List of values that V2IamActivityLogType can take
const (
	V2IamActivityLogTypeAnomalyDetectionSettingsUpdated         V2IamActivityLogType = "anomaly_detection_settings_updated"
	V2IamActivityLogTypeAPIKeyCreated                           V2IamActivityLogType = "api_key_created"
	V2IamActivityLogTypeAPIKeyDeleted                           V2IamActivityLogType = "api_key_deleted"
	V2IamActivityLogTypeAPIKeyUpdated                           V2IamActivityLogType = "api_key_updated"
	V2IamActivityLogTypeAPIKeyViewed                            V2IamActivityLogType = "api_key_viewed"
	V2IamActivityLogTypeIssuingActivated                        V2IamActivityLogType = "issuing_activated"
	V2IamActivityLogTypeIssuingBalanceTransferCreated           V2IamActivityLogType = "issuing_balance_transfer_created"
	V2IamActivityLogTypeIssuingCardholderCreated                V2IamActivityLogType = "issuing_cardholder_created"
	V2IamActivityLogTypeIssuingCardholderUpdated                V2IamActivityLogType = "issuing_cardholder_updated"
	V2IamActivityLogTypeIssuingCardCreated                      V2IamActivityLogType = "issuing_card_created"
	V2IamActivityLogTypeIssuingCardSensitiveDetailsViewed       V2IamActivityLogType = "issuing_card_sensitive_details_viewed"
	V2IamActivityLogTypeIssuingCardUpdated                      V2IamActivityLogType = "issuing_card_updated"
	V2IamActivityLogTypeIssuingDisputeCreated                   V2IamActivityLogType = "issuing_dispute_created"
	V2IamActivityLogTypeIssuingDisputeSubmitted                 V2IamActivityLogType = "issuing_dispute_submitted"
	V2IamActivityLogTypeIssuingDisputeUpdated                   V2IamActivityLogType = "issuing_dispute_updated"
	V2IamActivityLogTypeManualPayoutsDisabled                   V2IamActivityLogType = "manual_payouts_disabled"
	V2IamActivityLogTypeManualPayoutsEnabled                    V2IamActivityLogType = "manual_payouts_enabled"
	V2IamActivityLogTypePayoutDestinationAdded                  V2IamActivityLogType = "payout_destination_added"
	V2IamActivityLogTypePayoutDestinationRemoved                V2IamActivityLogType = "payout_destination_removed"
	V2IamActivityLogTypePayoutDestinationUpdated                V2IamActivityLogType = "payout_destination_updated"
	V2IamActivityLogTypePayoutScheduleEditsDisabled             V2IamActivityLogType = "payout_schedule_edits_disabled"
	V2IamActivityLogTypePayoutScheduleEditsEnabled              V2IamActivityLogType = "payout_schedule_edits_enabled"
	V2IamActivityLogTypeScimGroupDeleted                        V2IamActivityLogType = "scim_group_deleted"
	V2IamActivityLogTypeScimGroupMemberAdded                    V2IamActivityLogType = "scim_group_member_added"
	V2IamActivityLogTypeScimGroupMemberRemoved                  V2IamActivityLogType = "scim_group_member_removed"
	V2IamActivityLogTypeScimGroupRolesUpdated                   V2IamActivityLogType = "scim_group_roles_updated"
	V2IamActivityLogTypeScimGroupUpdated                        V2IamActivityLogType = "scim_group_updated"
	V2IamActivityLogTypeSsoDomainVerified                       V2IamActivityLogType = "sso_domain_verified"
	V2IamActivityLogTypeSsoSettingsCreated                      V2IamActivityLogType = "sso_settings_created"
	V2IamActivityLogTypeSsoSettingsDeleted                      V2IamActivityLogType = "sso_settings_deleted"
	V2IamActivityLogTypeSsoSettingsUpdated                      V2IamActivityLogType = "sso_settings_updated"
	V2IamActivityLogTypeTwoStepAuthenticationMandateDisabled    V2IamActivityLogType = "two_step_authentication_mandate_disabled"
	V2IamActivityLogTypeTwoStepAuthenticationMandateEnabled     V2IamActivityLogType = "two_step_authentication_mandate_enabled"
	V2IamActivityLogTypeUserAccessStarted                       V2IamActivityLogType = "user_access_started"
	V2IamActivityLogTypeUserAuthChallengeFailed                 V2IamActivityLogType = "user_auth_challenge_failed"
	V2IamActivityLogTypeUserEmailChanged                        V2IamActivityLogType = "user_email_changed"
	V2IamActivityLogTypeUserEmailVerified                       V2IamActivityLogType = "user_email_verified"
	V2IamActivityLogTypeUserExpressPhoneNumberChanged           V2IamActivityLogType = "user_express_phone_number_changed"
	V2IamActivityLogTypeUserGoogleAccountConnected              V2IamActivityLogType = "user_google_account_connected"
	V2IamActivityLogTypeUserGoogleAccountDisconnected           V2IamActivityLogType = "user_google_account_disconnected"
	V2IamActivityLogTypeUserInviteAccepted                      V2IamActivityLogType = "user_invite_accepted"
	V2IamActivityLogTypeUserInviteCreated                       V2IamActivityLogType = "user_invite_created"
	V2IamActivityLogTypeUserInviteDeleted                       V2IamActivityLogType = "user_invite_deleted"
	V2IamActivityLogTypeUserPasskeyAdded                        V2IamActivityLogType = "user_passkey_added"
	V2IamActivityLogTypeUserPasskeyRemoved                      V2IamActivityLogType = "user_passkey_removed"
	V2IamActivityLogTypeUserPasskeyUpdated                      V2IamActivityLogType = "user_passkey_updated"
	V2IamActivityLogTypeUserPasskeyUpgraded                     V2IamActivityLogType = "user_passkey_upgraded"
	V2IamActivityLogTypeUserPasswordChanged                     V2IamActivityLogType = "user_password_changed"
	V2IamActivityLogTypeUserPasswordInitialized                 V2IamActivityLogType = "user_password_initialized"
	V2IamActivityLogTypeUserPasswordResetFailed                 V2IamActivityLogType = "user_password_reset_failed"
	V2IamActivityLogTypeUserPasswordResetRequested              V2IamActivityLogType = "user_password_reset_requested"
	V2IamActivityLogTypeUserPasswordResetSucceeded              V2IamActivityLogType = "user_password_reset_succeeded"
	V2IamActivityLogTypeUserRolesDeleted                        V2IamActivityLogType = "user_roles_deleted"
	V2IamActivityLogTypeUserRolesUpdated                        V2IamActivityLogType = "user_roles_updated"
	V2IamActivityLogTypeUserTwoStepAuthenticationBackupCodeUsed V2IamActivityLogType = "user_two_step_authentication_backup_code_used"
	V2IamActivityLogTypeUserTwoStepAuthenticationMethodAdded    V2IamActivityLogType = "user_two_step_authentication_method_added"
	V2IamActivityLogTypeUserTwoStepAuthenticationMethodRemoved  V2IamActivityLogType = "user_two_step_authentication_method_removed"
	V2IamActivityLogTypeUserTwoStepAuthenticationMethodReset    V2IamActivityLogType = "user_two_step_authentication_method_reset"
	V2IamActivityLogTypeUserTwoStepAuthenticationMethodUpdated  V2IamActivityLogType = "user_two_step_authentication_method_updated"
	V2IamActivityLogTypeUserTwoStepAuthenticationResetRequested V2IamActivityLogType = "user_two_step_authentication_reset_requested"
)

// Set when the actor is an API key.
type V2IamActivityLogActorAPIKey struct {
	// Unique identifier of the API key.
	ID string `json:"id"`
}

// Set when the actor is a user.
type V2IamActivityLogActorUser struct {
	// Email address of the user.
	Email string `json:"email"`
}

// The actor that performed the action.
type V2IamActivityLogActor struct {
	// Set when the actor is an API key.
	APIKey *V2IamActivityLogActorAPIKey `json:"api_key,omitempty"`
	// The type of actor.
	Type V2IamActivityLogActorType `json:"type"`
	// Set when the actor is a user.
	User *V2IamActivityLogActorUser `json:"user,omitempty"`
}

// Anomaly detection settings after the change.
type V2IamActivityLogDetailsAccountSecurityNewAnomalySettings struct {
	// Whether dormant API key protection is enabled.
	DormantAPIKeyProtectionEnabled bool `json:"dormant_api_key_protection_enabled,omitempty"`
	// Whether money movement anomaly detection is enabled.
	MoneyMovementAnomalyDetectionEnabled bool `json:"money_movement_anomaly_detection_enabled,omitempty"`
	// Whether request-level anomaly detection is enabled.
	RequestLevelAnomalyDetectionEnabled bool `json:"request_level_anomaly_detection_enabled,omitempty"`
}

// Anomaly detection settings before the change.
type V2IamActivityLogDetailsAccountSecurityOldAnomalySettings struct {
	// Whether dormant API key protection is enabled.
	DormantAPIKeyProtectionEnabled bool `json:"dormant_api_key_protection_enabled,omitempty"`
	// Whether money movement anomaly detection is enabled.
	MoneyMovementAnomalyDetectionEnabled bool `json:"money_movement_anomaly_detection_enabled,omitempty"`
	// Whether request-level anomaly detection is enabled.
	RequestLevelAnomalyDetectionEnabled bool `json:"request_level_anomaly_detection_enabled,omitempty"`
}

// Details of an account security action.
type V2IamActivityLogDetailsAccountSecurity struct {
	// Anomaly detection settings after the change.
	NewAnomalySettings *V2IamActivityLogDetailsAccountSecurityNewAnomalySettings `json:"new_anomaly_settings,omitempty"`
	// Anomaly detection settings before the change.
	OldAnomalySettings *V2IamActivityLogDetailsAccountSecurityOldAnomalySettings `json:"old_anomaly_settings,omitempty"`
}

// An application.
type V2IamActivityLogDetailsAPIKeyManagedByApplication struct {
	// Identifier of the application.
	ID string `json:"id"`
}

// Information about the entity managing this API key.
type V2IamActivityLogDetailsAPIKeyManagedBy struct {
	// An application.
	Application *V2IamActivityLogDetailsAPIKeyManagedByApplication `json:"application,omitempty"`
	// The type of entity.
	Type V2IamActivityLogDetailsAPIKeyManagedByType `json:"type"`
}

// Details of an API key action.
type V2IamActivityLogDetailsAPIKey struct {
	// Timestamp when the API key was created.
	Created time.Time `json:"created"`
	// Timestamp when the API key expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Unique identifier of the API key.
	ID string `json:"id"`
	// List of IP addresses allowed to use this API key.
	IPAllowlist []string `json:"ip_allowlist"`
	// Information about the entity managing this API key.
	ManagedBy *V2IamActivityLogDetailsAPIKeyManagedBy `json:"managed_by,omitempty"`
	// Name of the API key.
	Name string `json:"name,omitempty"`
	// Unique identifier of the new API key, set when this key was rotated.
	NewKey string `json:"new_key,omitempty"`
	// Note or description for the API key.
	Note string `json:"note,omitempty"`
	// Type of the API key.
	Type V2IamActivityLogDetailsAPIKeyType `json:"type"`
}

// Details of an authentication action.
type V2IamActivityLogDetailsAuthentication struct {
	// Backup email address involved in the authentication.
	BackupEmail string `json:"backup_email,omitempty"`
	// Type of challenge used for the authentication.
	ChallengeType V2IamActivityLogDetailsAuthenticationChallengeType `json:"challenge_type,omitempty"`
	// Surface where the authentication occurred.
	Surface V2IamActivityLogDetailsAuthenticationSurface `json:"surface,omitempty"`
	// Target email address involved in the authentication.
	TargetEmail string `json:"target_email,omitempty"`
}

// Details of a SCIM action.
type V2IamActivityLogDetailsScim struct {
	// Name of the SCIM group.
	GroupName string `json:"group_name"`
	// Group roles after the change; only set for the group roles-updated action (scim_group_roles_updated).
	NewRoles []string `json:"new_roles"`
	// Group roles before the change; only set for the group roles-updated action (scim_group_roles_updated).
	OldRoles []string `json:"old_roles"`
	// The context the roles were assigned in.
	RoleAssignedContext string `json:"role_assigned_context,omitempty"`
	// Email address of the affected member.
	UserEmail string `json:"user_email,omitempty"`
}

// Details of an SSO action.
type V2IamActivityLogDetailsSso struct {
	// SSO enforcement level.
	Mandate V2IamActivityLogDetailsSsoMandate `json:"mandate,omitempty"`
}

// Primary authentication factor.
type V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactor struct {
	// SSO provider for the authentication factor.
	SsoProvider string `json:"sso_provider,omitempty"`
	// Type of authentication factor.
	Type V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactorType `json:"type"`
}

// Secondary authentication factors.
type V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactor struct {
	// SSO provider for the authentication factor.
	SsoProvider string `json:"sso_provider,omitempty"`
	// Type of authentication factor.
	Type V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactorType `json:"type"`
}

// Authentication details for the user access action.
type V2IamActivityLogDetailsUserAccessAuthentication struct {
	// Primary authentication factor.
	PrimaryFactor *V2IamActivityLogDetailsUserAccessAuthenticationPrimaryFactor `json:"primary_factor"`
	// Secondary authentication factors.
	SecondaryFactors []*V2IamActivityLogDetailsUserAccessAuthenticationSecondaryFactor `json:"secondary_factors"`
}

// Dashboard client details for the user access action.
type V2IamActivityLogDetailsUserAccessDashboardClient struct {
	// Browser used for the user access action.
	Browser string `json:"browser"`
	// Browser version used for the user access action.
	BrowserVersion string `json:"browser_version"`
	// Device type used for the user access action.
	DeviceType string `json:"device_type"`
	// Operating system used for the user access action.
	Os string `json:"os"`
}

// Network details for the user access action.
type V2IamActivityLogDetailsUserAccessNetwork struct {
	// City for the user access action.
	City string `json:"city"`
	// Country for the user access action.
	Country string `json:"country"`
	// IP address for the user access action.
	IPAddress string `json:"ip_address"`
	// Region for the user access action.
	Region string `json:"region"`
}

// The user access action used a novel device.
type V2IamActivityLogDetailsUserAccessRiskSignalNovelDevice struct{}

// Risk signals for the user access action.
type V2IamActivityLogDetailsUserAccessRiskSignal struct {
	// The user access action used a novel device.
	NovelDevice *V2IamActivityLogDetailsUserAccessRiskSignalNovelDevice `json:"novel_device,omitempty"`
	// Type of risk signal.
	Type V2IamActivityLogDetailsUserAccessRiskSignalType `json:"type"`
}

// Risk details for the user access action.
type V2IamActivityLogDetailsUserAccessRisk struct {
	// Risk level for the user access action.
	Level V2IamActivityLogDetailsUserAccessRiskLevel `json:"level"`
	// Risk signals for the user access action.
	Signals []*V2IamActivityLogDetailsUserAccessRiskSignal `json:"signals"`
}

// Details of a user access action.
type V2IamActivityLogDetailsUserAccess struct {
	// Authentication details for the user access action.
	Authentication *V2IamActivityLogDetailsUserAccessAuthentication `json:"authentication"`
	// Dashboard client details for the user access action.
	DashboardClient *V2IamActivityLogDetailsUserAccessDashboardClient `json:"dashboard_client,omitempty"`
	// Timestamp when the user access expires.
	ExpiresAt time.Time `json:"expires_at"`
	// Network details for the user access action.
	Network *V2IamActivityLogDetailsUserAccessNetwork `json:"network"`
	// Risk details for the user access action.
	Risk *V2IamActivityLogDetailsUserAccessRisk `json:"risk"`
	// Roles associated with the user access action.
	Roles []string `json:"roles"`
	// Session fingerprint for the user access action.
	SessionFingerprint string `json:"session_fingerprint"`
	// Surface where the user access action started.
	Surface V2IamActivityLogDetailsUserAccessSurface `json:"surface"`
}

// Details of a user invite action.
type V2IamActivityLogDetailsUserInvite struct {
	// Email address of the invited user.
	InvitedUserEmail string `json:"invited_user_email"`
	// Roles assigned to the invited user.
	Roles []string `json:"roles"`
}

// Details of a user profile action.
type V2IamActivityLogDetailsUserProfile struct {
	// Email address after the change.
	NewEmail string `json:"new_email,omitempty"`
	// Redacted phone number after the change.
	NewRedactedPhoneNumber string `json:"new_redacted_phone_number,omitempty"`
	// Email address before the change.
	OldEmail string `json:"old_email,omitempty"`
	// Redacted phone number before the change.
	OldRedactedPhoneNumber string `json:"old_redacted_phone_number,omitempty"`
}

// Details of a user role change action.
type V2IamActivityLogDetailsUserRoles struct {
	// Roles the user has after the change.
	NewRoles []string `json:"new_roles"`
	// Roles the user had before the change.
	OldRoles []string `json:"old_roles"`
	// Source of the role change.
	Source V2IamActivityLogDetailsUserRolesSource `json:"source"`
	// Email address of the user whose roles were changed.
	UserEmail string `json:"user_email"`
}

// Action-specific details of the activity log entry.
type V2IamActivityLogDetails struct {
	// Details of an account security action.
	AccountSecurity *V2IamActivityLogDetailsAccountSecurity `json:"account_security,omitempty"`
	// Details of an API key action.
	APIKey *V2IamActivityLogDetailsAPIKey `json:"api_key,omitempty"`
	// Details of an authentication action.
	Authentication *V2IamActivityLogDetailsAuthentication `json:"authentication,omitempty"`
	// Details of a SCIM action.
	Scim *V2IamActivityLogDetailsScim `json:"scim,omitempty"`
	// Details of an SSO action.
	Sso *V2IamActivityLogDetailsSso `json:"sso,omitempty"`
	// The action group type of the activity log entry.
	Type V2IamActivityLogDetailsType `json:"type"`
	// Details of a user access action.
	UserAccess *V2IamActivityLogDetailsUserAccess `json:"user_access,omitempty"`
	// Details of a user invite action.
	UserInvite *V2IamActivityLogDetailsUserInvite `json:"user_invite,omitempty"`
	// Details of a user profile action.
	UserProfile *V2IamActivityLogDetailsUserProfile `json:"user_profile,omitempty"`
	// Details of a user role change action.
	UserRoles *V2IamActivityLogDetailsUserRoles `json:"user_roles,omitempty"`
}

// The object related to the activity log entry.
type V2IamActivityLogRelatedObject struct {
	// Unique identifier of the object.
	ID string `json:"id"`
	// Type of the object.
	Type V2IamActivityLogRelatedObjectType `json:"type"`
}

// The API request that instigated the action.
type V2IamActivityLogRequest struct {
	// ID of the API request.
	ID string `json:"id"`
}

// An activity log records a single action performed on an account.
type V2IamActivityLog struct {
	APIResource
	// The actor that performed the action.
	Actor *V2IamActivityLogActor `json:"actor"`
	// The account on which the action was performed.
	Context string `json:"context"`
	// Timestamp when the activity log entry was created.
	Created time.Time `json:"created"`
	// Action-specific details of the activity log entry.
	Details *V2IamActivityLogDetails `json:"details"`
	// Unique identifier of the activity log entry.
	ID string `json:"id"`
	// Whether the action was performed in live mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The object related to the activity log entry.
	RelatedObject *V2IamActivityLogRelatedObject `json:"related_object,omitempty"`
	// The API request that instigated the action.
	Request *V2IamActivityLogRequest `json:"request,omitempty"`
	// The type of action that was performed.
	Type V2IamActivityLogType `json:"type"`
}
