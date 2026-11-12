//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"github.com/shopspring/decimal"
	"time"
)

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

// The site link route. Set for the connect_site_link_updated action.
type V2IamActivityLogDetailsConnectSiteLinkRoute string

// List of values that V2IamActivityLogDetailsConnectSiteLinkRoute can take
const (
	V2IamActivityLogDetailsConnectSiteLinkRouteAccountManagement       V2IamActivityLogDetailsConnectSiteLinkRoute = "account_management"
	V2IamActivityLogDetailsConnectSiteLinkRouteBalances                V2IamActivityLogDetailsConnectSiteLinkRoute = "balances"
	V2IamActivityLogDetailsConnectSiteLinkRouteDocuments               V2IamActivityLogDetailsConnectSiteLinkRoute = "documents"
	V2IamActivityLogDetailsConnectSiteLinkRouteFinancialAccounts       V2IamActivityLogDetailsConnectSiteLinkRoute = "financial_accounts"
	V2IamActivityLogDetailsConnectSiteLinkRouteIssuingCardTransactions V2IamActivityLogDetailsConnectSiteLinkRoute = "issuing_card_transactions"
	V2IamActivityLogDetailsConnectSiteLinkRouteNotificationBanner      V2IamActivityLogDetailsConnectSiteLinkRoute = "notification_banner"
	V2IamActivityLogDetailsConnectSiteLinkRoutePayments                V2IamActivityLogDetailsConnectSiteLinkRoute = "payments"
	V2IamActivityLogDetailsConnectSiteLinkRoutePayouts                 V2IamActivityLogDetailsConnectSiteLinkRoute = "payouts"
)

// Status of the event destination.
type V2IamActivityLogDetailsEventDestinationNewConfigurationStatus string

// List of values that V2IamActivityLogDetailsEventDestinationNewConfigurationStatus can take
const (
	V2IamActivityLogDetailsEventDestinationNewConfigurationStatusDisabled V2IamActivityLogDetailsEventDestinationNewConfigurationStatus = "disabled"
	V2IamActivityLogDetailsEventDestinationNewConfigurationStatusEnabled  V2IamActivityLogDetailsEventDestinationNewConfigurationStatus = "enabled"
)

// Status of the event destination.
type V2IamActivityLogDetailsEventDestinationOldConfigurationStatus string

// List of values that V2IamActivityLogDetailsEventDestinationOldConfigurationStatus can take
const (
	V2IamActivityLogDetailsEventDestinationOldConfigurationStatusDisabled V2IamActivityLogDetailsEventDestinationOldConfigurationStatus = "disabled"
	V2IamActivityLogDetailsEventDestinationOldConfigurationStatusEnabled  V2IamActivityLogDetailsEventDestinationOldConfigurationStatus = "enabled"
)

// Type of the event destination.
type V2IamActivityLogDetailsEventDestinationType string

// List of values that V2IamActivityLogDetailsEventDestinationType can take
const (
	V2IamActivityLogDetailsEventDestinationTypeAmazonEventbridge V2IamActivityLogDetailsEventDestinationType = "amazon_eventbridge"
	V2IamActivityLogDetailsEventDestinationTypeAzureEventGrid    V2IamActivityLogDetailsEventDestinationType = "azure_event_grid"
	V2IamActivityLogDetailsEventDestinationTypeWebhookEndpoint   V2IamActivityLogDetailsEventDestinationType = "webhook_endpoint"
)

// Risk preference after a preference switch carried by a protection-toggle action.
type V2IamActivityLogDetailsRadarNewRiskPreference string

// List of values that V2IamActivityLogDetailsRadarNewRiskPreference can take
const (
	V2IamActivityLogDetailsRadarNewRiskPreferenceBalanced   V2IamActivityLogDetailsRadarNewRiskPreference = "balanced"
	V2IamActivityLogDetailsRadarNewRiskPreferenceCustom     V2IamActivityLogDetailsRadarNewRiskPreference = "custom"
	V2IamActivityLogDetailsRadarNewRiskPreferenceHighGrowth V2IamActivityLogDetailsRadarNewRiskPreference = "high_growth"
	V2IamActivityLogDetailsRadarNewRiskPreferenceLowRisk    V2IamActivityLogDetailsRadarNewRiskPreference = "low_risk"
)

// Action performed when the rule matches.
type V2IamActivityLogDetailsRadarNewRuleAction string

// List of values that V2IamActivityLogDetailsRadarNewRuleAction can take
const (
	V2IamActivityLogDetailsRadarNewRuleActionAllow              V2IamActivityLogDetailsRadarNewRuleAction = "allow"
	V2IamActivityLogDetailsRadarNewRuleActionBlock              V2IamActivityLogDetailsRadarNewRuleAction = "block"
	V2IamActivityLogDetailsRadarNewRuleActionBlockFaOutbound    V2IamActivityLogDetailsRadarNewRuleAction = "block_fa_outbound"
	V2IamActivityLogDetailsRadarNewRuleActionManualReview       V2IamActivityLogDetailsRadarNewRuleAction = "manual_review"
	V2IamActivityLogDetailsRadarNewRuleActionPayoutPause        V2IamActivityLogDetailsRadarNewRuleAction = "payout_pause"
	V2IamActivityLogDetailsRadarNewRuleActionRaiseReview        V2IamActivityLogDetailsRadarNewRuleAction = "raise_review"
	V2IamActivityLogDetailsRadarNewRuleActionRequestCredentials V2IamActivityLogDetailsRadarNewRuleAction = "request_credentials"
	V2IamActivityLogDetailsRadarNewRuleActionResolveDispute     V2IamActivityLogDetailsRadarNewRuleAction = "resolve_dispute"
)

// Type of entity evaluated by the rule.
type V2IamActivityLogDetailsRadarNewRuleType string

// List of values that V2IamActivityLogDetailsRadarNewRuleType can take
const (
	V2IamActivityLogDetailsRadarNewRuleTypeAccount V2IamActivityLogDetailsRadarNewRuleType = "account"
	V2IamActivityLogDetailsRadarNewRuleTypePayment V2IamActivityLogDetailsRadarNewRuleType = "payment"
)

// Risk preference before a preference switch carried by a protection-toggle action.
type V2IamActivityLogDetailsRadarOldRiskPreference string

// List of values that V2IamActivityLogDetailsRadarOldRiskPreference can take
const (
	V2IamActivityLogDetailsRadarOldRiskPreferenceBalanced   V2IamActivityLogDetailsRadarOldRiskPreference = "balanced"
	V2IamActivityLogDetailsRadarOldRiskPreferenceCustom     V2IamActivityLogDetailsRadarOldRiskPreference = "custom"
	V2IamActivityLogDetailsRadarOldRiskPreferenceHighGrowth V2IamActivityLogDetailsRadarOldRiskPreference = "high_growth"
	V2IamActivityLogDetailsRadarOldRiskPreferenceLowRisk    V2IamActivityLogDetailsRadarOldRiskPreference = "low_risk"
)

// Action performed when the rule matches.
type V2IamActivityLogDetailsRadarOldRuleAction string

// List of values that V2IamActivityLogDetailsRadarOldRuleAction can take
const (
	V2IamActivityLogDetailsRadarOldRuleActionAllow              V2IamActivityLogDetailsRadarOldRuleAction = "allow"
	V2IamActivityLogDetailsRadarOldRuleActionBlock              V2IamActivityLogDetailsRadarOldRuleAction = "block"
	V2IamActivityLogDetailsRadarOldRuleActionBlockFaOutbound    V2IamActivityLogDetailsRadarOldRuleAction = "block_fa_outbound"
	V2IamActivityLogDetailsRadarOldRuleActionManualReview       V2IamActivityLogDetailsRadarOldRuleAction = "manual_review"
	V2IamActivityLogDetailsRadarOldRuleActionPayoutPause        V2IamActivityLogDetailsRadarOldRuleAction = "payout_pause"
	V2IamActivityLogDetailsRadarOldRuleActionRaiseReview        V2IamActivityLogDetailsRadarOldRuleAction = "raise_review"
	V2IamActivityLogDetailsRadarOldRuleActionRequestCredentials V2IamActivityLogDetailsRadarOldRuleAction = "request_credentials"
	V2IamActivityLogDetailsRadarOldRuleActionResolveDispute     V2IamActivityLogDetailsRadarOldRuleAction = "resolve_dispute"
)

// Type of entity evaluated by the rule.
type V2IamActivityLogDetailsRadarOldRuleType string

// List of values that V2IamActivityLogDetailsRadarOldRuleType can take
const (
	V2IamActivityLogDetailsRadarOldRuleTypeAccount V2IamActivityLogDetailsRadarOldRuleType = "account"
	V2IamActivityLogDetailsRadarOldRuleTypePayment V2IamActivityLogDetailsRadarOldRuleType = "payment"
)

// SSO enforcement level.
type V2IamActivityLogDetailsSsoMandate string

// List of values that V2IamActivityLogDetailsSsoMandate can take
const (
	V2IamActivityLogDetailsSsoMandateOff      V2IamActivityLogDetailsSsoMandate = "off"
	V2IamActivityLogDetailsSsoMandateOptional V2IamActivityLogDetailsSsoMandate = "optional"
	V2IamActivityLogDetailsSsoMandateRequired V2IamActivityLogDetailsSsoMandate = "required"
)

// The level of Dashboard access granted.
type V2IamActivityLogDetailsStripeAccessAccessType string

// List of values that V2IamActivityLogDetailsStripeAccessAccessType can take
const (
	V2IamActivityLogDetailsStripeAccessAccessTypeDashboardReadOnly  V2IamActivityLogDetailsStripeAccessAccessType = "dashboard_read_only"
	V2IamActivityLogDetailsStripeAccessAccessTypeDashboardReadWrite V2IamActivityLogDetailsStripeAccessAccessType = "dashboard_read_write"
)

// The type of reason for granting access.
type V2IamActivityLogDetailsStripeAccessReasonType string

// List of values that V2IamActivityLogDetailsStripeAccessReasonType can take
const (
	V2IamActivityLogDetailsStripeAccessReasonTypeAccessRequest       V2IamActivityLogDetailsStripeAccessReasonType = "access_request"
	V2IamActivityLogDetailsStripeAccessReasonTypeSupportRelationship V2IamActivityLogDetailsStripeAccessReasonType = "support_relationship"
	V2IamActivityLogDetailsStripeAccessReasonTypeSupportTicket       V2IamActivityLogDetailsStripeAccessReasonType = "support_ticket"
)

// The action group type of the activity log entry.
type V2IamActivityLogDetailsType string

// List of values that V2IamActivityLogDetailsType can take
const (
	V2IamActivityLogDetailsTypeAccountSecurity  V2IamActivityLogDetailsType = "account_security"
	V2IamActivityLogDetailsTypeAPIKey           V2IamActivityLogDetailsType = "api_key"
	V2IamActivityLogDetailsTypeApp              V2IamActivityLogDetailsType = "app"
	V2IamActivityLogDetailsTypeAuthentication   V2IamActivityLogDetailsType = "authentication"
	V2IamActivityLogDetailsTypeConnect          V2IamActivityLogDetailsType = "connect"
	V2IamActivityLogDetailsTypeDataExport       V2IamActivityLogDetailsType = "data_export"
	V2IamActivityLogDetailsTypeEventDestination V2IamActivityLogDetailsType = "event_destination"
	V2IamActivityLogDetailsTypeIssuing          V2IamActivityLogDetailsType = "issuing"
	V2IamActivityLogDetailsTypePayout           V2IamActivityLogDetailsType = "payout"
	V2IamActivityLogDetailsTypeRadar            V2IamActivityLogDetailsType = "radar"
	V2IamActivityLogDetailsTypeScim             V2IamActivityLogDetailsType = "scim"
	V2IamActivityLogDetailsTypeSso              V2IamActivityLogDetailsType = "sso"
	V2IamActivityLogDetailsTypeStripeAccess     V2IamActivityLogDetailsType = "stripe_access"
	V2IamActivityLogDetailsTypeUserAccess       V2IamActivityLogDetailsType = "user_access"
	V2IamActivityLogDetailsTypeUserInvite       V2IamActivityLogDetailsType = "user_invite"
	V2IamActivityLogDetailsTypeUserProfile      V2IamActivityLogDetailsType = "user_profile"
	V2IamActivityLogDetailsTypeUserRoles        V2IamActivityLogDetailsType = "user_roles"
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
	V2IamActivityLogDetailsUserAccessSurfaceEmbedded  V2IamActivityLogDetailsUserAccessSurface = "embedded"
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
	V2IamActivityLogRelatedObjectTypeAppsInstall            V2IamActivityLogRelatedObjectType = "apps.install"
	V2IamActivityLogRelatedObjectTypeBalanceTransfer        V2IamActivityLogRelatedObjectType = "balance_transfer"
	V2IamActivityLogRelatedObjectTypeBankAccount            V2IamActivityLogRelatedObjectType = "bank_account"
	V2IamActivityLogRelatedObjectTypeBlockchainAddress      V2IamActivityLogRelatedObjectType = "blockchain_address"
	V2IamActivityLogRelatedObjectTypeCard                   V2IamActivityLogRelatedObjectType = "card"
	V2IamActivityLogRelatedObjectTypeV2CoreEventDestination V2IamActivityLogRelatedObjectType = "v2.core.event_destination"
	V2IamActivityLogRelatedObjectTypeIssuingCard            V2IamActivityLogRelatedObjectType = "issuing.card"
	V2IamActivityLogRelatedObjectTypeIssuingCardholder      V2IamActivityLogRelatedObjectType = "issuing.cardholder"
	V2IamActivityLogRelatedObjectTypeIssuingDispute         V2IamActivityLogRelatedObjectType = "issuing.dispute"
	V2IamActivityLogRelatedObjectTypeWebhookEndpoint        V2IamActivityLogRelatedObjectType = "webhook_endpoint"
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
	V2IamActivityLogTypeApplicationAuthorized                   V2IamActivityLogType = "application_authorized"
	V2IamActivityLogTypeApplicationCreated                      V2IamActivityLogType = "application_created"
	V2IamActivityLogTypeApplicationDeauthorized                 V2IamActivityLogType = "application_deauthorized"
	V2IamActivityLogTypeAppInstalled                            V2IamActivityLogType = "app_installed"
	V2IamActivityLogTypeAppUninstalled                          V2IamActivityLogType = "app_uninstalled"
	V2IamActivityLogTypeConnectExpressBrandingUpdated           V2IamActivityLogType = "connect_express_branding_updated"
	V2IamActivityLogTypeConnectOauthDisabled                    V2IamActivityLogType = "connect_oauth_disabled"
	V2IamActivityLogTypeConnectOauthEnabled                     V2IamActivityLogType = "connect_oauth_enabled"
	V2IamActivityLogTypeConnectSettingsUpdated                  V2IamActivityLogType = "connect_settings_updated"
	V2IamActivityLogTypeConnectSiteLinkUpdated                  V2IamActivityLogType = "connect_site_link_updated"
	V2IamActivityLogTypeDataExportRequested                     V2IamActivityLogType = "data_export_requested"
	V2IamActivityLogTypeEventDestinationAdded                   V2IamActivityLogType = "event_destination_added"
	V2IamActivityLogTypeEventDestinationRemoved                 V2IamActivityLogType = "event_destination_removed"
	V2IamActivityLogTypeEventDestinationUpdated                 V2IamActivityLogType = "event_destination_updated"
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
	V2IamActivityLogTypePayoutScheduleUpdated                   V2IamActivityLogType = "payout_schedule_updated"
	V2IamActivityLogTypeRadarAdaptive3dsDisabled                V2IamActivityLogType = "radar_adaptive_3ds_disabled"
	V2IamActivityLogTypeRadarDynamicRiskThresholdsDisabled      V2IamActivityLogType = "radar_dynamic_risk_thresholds_disabled"
	V2IamActivityLogTypeRadarEarlyFraudWarningDisabled          V2IamActivityLogType = "radar_early_fraud_warning_disabled"
	V2IamActivityLogTypeRadarForSetupIntentsDisabled            V2IamActivityLogType = "radar_for_setup_intents_disabled"
	V2IamActivityLogTypeRadarFraudDisputeProtectionDisabled     V2IamActivityLogType = "radar_fraud_dispute_protection_disabled"
	V2IamActivityLogTypeRadarRuleAdded                          V2IamActivityLogType = "radar_rule_added"
	V2IamActivityLogTypeRadarRuleDisabled                       V2IamActivityLogType = "radar_rule_disabled"
	V2IamActivityLogTypeRadarRuleEnabled                        V2IamActivityLogType = "radar_rule_enabled"
	V2IamActivityLogTypeRadarRuleRemoved                        V2IamActivityLogType = "radar_rule_removed"
	V2IamActivityLogTypeRadarRuleUpdated                        V2IamActivityLogType = "radar_rule_updated"
	V2IamActivityLogTypeScimGroupDeleted                        V2IamActivityLogType = "scim_group_deleted"
	V2IamActivityLogTypeScimGroupMemberAdded                    V2IamActivityLogType = "scim_group_member_added"
	V2IamActivityLogTypeScimGroupMemberRemoved                  V2IamActivityLogType = "scim_group_member_removed"
	V2IamActivityLogTypeScimGroupRolesUpdated                   V2IamActivityLogType = "scim_group_roles_updated"
	V2IamActivityLogTypeScimGroupUpdated                        V2IamActivityLogType = "scim_group_updated"
	V2IamActivityLogTypeSsoDomainVerified                       V2IamActivityLogType = "sso_domain_verified"
	V2IamActivityLogTypeSsoSettingsCreated                      V2IamActivityLogType = "sso_settings_created"
	V2IamActivityLogTypeSsoSettingsDeleted                      V2IamActivityLogType = "sso_settings_deleted"
	V2IamActivityLogTypeSsoSettingsUpdated                      V2IamActivityLogType = "sso_settings_updated"
	V2IamActivityLogTypeStripeAccessStarted                     V2IamActivityLogType = "stripe_access_started"
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

// Set when a display name is available for an action performed by Stripe.
type V2IamActivityLogActorStripeAction struct {
	// Public display name captured when the action occurs.
	DisplayName string `json:"display_name"`
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
	// Set when a display name is available for an action performed by Stripe.
	StripeAction *V2IamActivityLogActorStripeAction `json:"stripe_action,omitempty"`
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

// The Connect application.
type V2IamActivityLogDetailsConnectApplication struct {
	// Identifier of the application.
	ID string `json:"id"`
}

// Express Dashboard branding settings after the change. Set for the connect_express_branding_updated action.
type V2IamActivityLogDetailsConnectNewExpressBrandingSettings struct {
	// Border radius for the Express Dashboard.
	BorderRadius string `json:"border_radius,omitempty"`
	// Primary button background color for the Express Dashboard.
	PrimaryButtonBackgroundColor string `json:"primary_button_background_color,omitempty"`
	// Primary color for the Express Dashboard.
	PrimaryColor string `json:"primary_color,omitempty"`
	// Secondary button background color for the Express Dashboard.
	SecondaryButtonBackgroundColor string `json:"secondary_button_background_color,omitempty"`
}

// Connect settings after the change. Set for the connect_settings_updated action.
type V2IamActivityLogDetailsConnectNewSettings struct {
	// The display name shown on the onboarding experience and Dashboard. This is different from the business legal name.
	BusinessName string `json:"business_name,omitempty"`
	// Whether platform branding is copied to new connected accounts. These also appear on invoices, receipts, and other Stripe products.
	CopyPlatformBrandingToNewConnectedAccounts bool `json:"copy_platform_branding_to_new_connected_accounts,omitempty"`
	// Whether the support email is hidden on onboarding surfaces.
	HideOnboardingSupportEmail bool `json:"hide_onboarding_support_email,omitempty"`
	// Whether the support phone number is hidden on onboarding surfaces.
	HideOnboardingSupportPhone bool `json:"hide_onboarding_support_phone,omitempty"`
	// Whether the support URL is hidden on onboarding surfaces.
	HideOnboardingSupportURL bool `json:"hide_onboarding_support_url,omitempty"`
	// Identifier of the icon file (file_* ID).
	Icon string `json:"icon,omitempty"`
	// Identifier of the logo file (file_* ID).
	Logo string `json:"logo,omitempty"`
	// The platform's default payout statement descriptor. It appears on the bank statements of the platform's connected accounts for their payouts.
	PayoutStatementDescriptor string `json:"payout_statement_descriptor,omitempty"`
	// A CSS hex color value representing the primary branding color, applied to the onboarding experience and the Stripe Dashboard.
	PrimaryColor string `json:"primary_color,omitempty"`
	// A CSS hex color value representing the secondary branding color, applied to the onboarding experience and the Stripe Dashboard.
	SecondaryColor string `json:"secondary_color,omitempty"`
}

// Express Dashboard branding settings before the change. Set for the connect_express_branding_updated action.
type V2IamActivityLogDetailsConnectOldExpressBrandingSettings struct {
	// Border radius for the Express Dashboard.
	BorderRadius string `json:"border_radius,omitempty"`
	// Primary button background color for the Express Dashboard.
	PrimaryButtonBackgroundColor string `json:"primary_button_background_color,omitempty"`
	// Primary color for the Express Dashboard.
	PrimaryColor string `json:"primary_color,omitempty"`
	// Secondary button background color for the Express Dashboard.
	SecondaryButtonBackgroundColor string `json:"secondary_button_background_color,omitempty"`
}

// Connect settings before the change. Set for the connect_settings_updated action.
type V2IamActivityLogDetailsConnectOldSettings struct {
	// The display name shown on the onboarding experience and Dashboard. This is different from the business legal name.
	BusinessName string `json:"business_name,omitempty"`
	// Whether platform branding is copied to new connected accounts. These also appear on invoices, receipts, and other Stripe products.
	CopyPlatformBrandingToNewConnectedAccounts bool `json:"copy_platform_branding_to_new_connected_accounts,omitempty"`
	// Whether the support email is hidden on onboarding surfaces.
	HideOnboardingSupportEmail bool `json:"hide_onboarding_support_email,omitempty"`
	// Whether the support phone number is hidden on onboarding surfaces.
	HideOnboardingSupportPhone bool `json:"hide_onboarding_support_phone,omitempty"`
	// Whether the support URL is hidden on onboarding surfaces.
	HideOnboardingSupportURL bool `json:"hide_onboarding_support_url,omitempty"`
	// Identifier of the icon file (file_* ID).
	Icon string `json:"icon,omitempty"`
	// Identifier of the logo file (file_* ID).
	Logo string `json:"logo,omitempty"`
	// The platform's default payout statement descriptor. It appears on the bank statements of the platform's connected accounts for their payouts.
	PayoutStatementDescriptor string `json:"payout_statement_descriptor,omitempty"`
	// A CSS hex color value representing the primary branding color, applied to the onboarding experience and the Stripe Dashboard.
	PrimaryColor string `json:"primary_color,omitempty"`
	// A CSS hex color value representing the secondary branding color, applied to the onboarding experience and the Stripe Dashboard.
	SecondaryColor string `json:"secondary_color,omitempty"`
}

// Details of a Connect action.
type V2IamActivityLogDetailsConnect struct {
	// The Connect application.
	Application *V2IamActivityLogDetailsConnectApplication `json:"application,omitempty"`
	// Express Dashboard branding settings after the change. Set for the connect_express_branding_updated action.
	NewExpressBrandingSettings *V2IamActivityLogDetailsConnectNewExpressBrandingSettings `json:"new_express_branding_settings,omitempty"`
	// Connect settings after the change. Set for the connect_settings_updated action.
	NewSettings *V2IamActivityLogDetailsConnectNewSettings `json:"new_settings,omitempty"`
	// Express Dashboard branding settings before the change. Set for the connect_express_branding_updated action.
	OldExpressBrandingSettings *V2IamActivityLogDetailsConnectOldExpressBrandingSettings `json:"old_express_branding_settings,omitempty"`
	// Connect settings before the change. Set for the connect_settings_updated action.
	OldSettings *V2IamActivityLogDetailsConnectOldSettings `json:"old_settings,omitempty"`
	// The site link route. Set for the connect_site_link_updated action.
	SiteLinkRoute V2IamActivityLogDetailsConnectSiteLinkRoute `json:"site_link_route,omitempty"`
}

// Details of a data export action.
type V2IamActivityLogDetailsDataExport struct {
	// Name of export type.
	Type string `json:"type"`
}

// Amazon EventBridge configuration.
type V2IamActivityLogDetailsEventDestinationNewConfigurationAmazonEventbridge struct {
	// AWS account ID configured for the event destination.
	AwsAccountID string `json:"aws_account_id"`
	// ARN of the AWS event source configured for the event destination.
	AwsEventSourceArn string `json:"aws_event_source_arn"`
	// AWS region configured for the event destination.
	AwsRegion string `json:"aws_region"`
}

// Azure Event Grid configuration.
type V2IamActivityLogDetailsEventDestinationNewConfigurationAzureEventGrid struct {
	// Azure partner topic name configured for the event destination.
	AzurePartnerTopicName string `json:"azure_partner_topic_name"`
	// Azure region configured for the event destination.
	AzureRegion string `json:"azure_region"`
	// Azure resource group name configured for the event destination.
	AzureResourceGroupName string `json:"azure_resource_group_name"`
	// Azure subscription ID configured for the event destination.
	AzureSubscriptionID string `json:"azure_subscription_id"`
}

// Webhook endpoint configuration.
type V2IamActivityLogDetailsEventDestinationNewConfigurationWebhookEndpoint struct {
	// URL configured for the webhook endpoint.
	URL string `json:"url"`
}

// Configuration after the change. Set for added and updated actions, absent for removed actions.
type V2IamActivityLogDetailsEventDestinationNewConfiguration struct {
	// Amazon EventBridge configuration.
	AmazonEventbridge *V2IamActivityLogDetailsEventDestinationNewConfigurationAmazonEventbridge `json:"amazon_eventbridge,omitempty"`
	// Azure Event Grid configuration.
	AzureEventGrid *V2IamActivityLogDetailsEventDestinationNewConfigurationAzureEventGrid `json:"azure_event_grid,omitempty"`
	// Event types subscribed to by the destination. ["*"] subscribes to all V1 webhook events.
	EnabledEvents []string `json:"enabled_events"`
	// Event source selectors: @self, @accounts, @organization_members, or @organization_members/@accounts.
	EventsFrom []string `json:"events_from"`
	// Status of the event destination.
	Status V2IamActivityLogDetailsEventDestinationNewConfigurationStatus `json:"status"`
	// Webhook endpoint configuration.
	WebhookEndpoint *V2IamActivityLogDetailsEventDestinationNewConfigurationWebhookEndpoint `json:"webhook_endpoint,omitempty"`
}

// Amazon EventBridge configuration.
type V2IamActivityLogDetailsEventDestinationOldConfigurationAmazonEventbridge struct {
	// AWS account ID configured for the event destination.
	AwsAccountID string `json:"aws_account_id"`
	// ARN of the AWS event source configured for the event destination.
	AwsEventSourceArn string `json:"aws_event_source_arn"`
	// AWS region configured for the event destination.
	AwsRegion string `json:"aws_region"`
}

// Azure Event Grid configuration.
type V2IamActivityLogDetailsEventDestinationOldConfigurationAzureEventGrid struct {
	// Azure partner topic name configured for the event destination.
	AzurePartnerTopicName string `json:"azure_partner_topic_name"`
	// Azure region configured for the event destination.
	AzureRegion string `json:"azure_region"`
	// Azure resource group name configured for the event destination.
	AzureResourceGroupName string `json:"azure_resource_group_name"`
	// Azure subscription ID configured for the event destination.
	AzureSubscriptionID string `json:"azure_subscription_id"`
}

// Webhook endpoint configuration.
type V2IamActivityLogDetailsEventDestinationOldConfigurationWebhookEndpoint struct {
	// URL configured for the webhook endpoint.
	URL string `json:"url"`
}

// Configuration before the change. Set for updated and removed actions, absent for added actions.
type V2IamActivityLogDetailsEventDestinationOldConfiguration struct {
	// Amazon EventBridge configuration.
	AmazonEventbridge *V2IamActivityLogDetailsEventDestinationOldConfigurationAmazonEventbridge `json:"amazon_eventbridge,omitempty"`
	// Azure Event Grid configuration.
	AzureEventGrid *V2IamActivityLogDetailsEventDestinationOldConfigurationAzureEventGrid `json:"azure_event_grid,omitempty"`
	// Event types subscribed to by the destination. ["*"] subscribes to all V1 webhook events.
	EnabledEvents []string `json:"enabled_events"`
	// Event source selectors: @self, @accounts, @organization_members, or @organization_members/@accounts.
	EventsFrom []string `json:"events_from"`
	// Status of the event destination.
	Status V2IamActivityLogDetailsEventDestinationOldConfigurationStatus `json:"status"`
	// Webhook endpoint configuration.
	WebhookEndpoint *V2IamActivityLogDetailsEventDestinationOldConfigurationWebhookEndpoint `json:"webhook_endpoint,omitempty"`
}

// Details of an event destination action.
type V2IamActivityLogDetailsEventDestination struct {
	// Configuration after the change. Set for added and updated actions, absent for removed actions.
	NewConfiguration *V2IamActivityLogDetailsEventDestinationNewConfiguration `json:"new_configuration,omitempty"`
	// Configuration before the change. Set for updated and removed actions, absent for added actions.
	OldConfiguration *V2IamActivityLogDetailsEventDestinationOldConfiguration `json:"old_configuration,omitempty"`
	// Type of the event destination.
	Type V2IamActivityLogDetailsEventDestinationType `json:"type"`
}

// Rules after the change. Set for add, update, enable, and disable actions; absent for remove actions.
type V2IamActivityLogDetailsRadarNewRule struct {
	// Action performed when the rule matches.
	Action V2IamActivityLogDetailsRadarNewRuleAction `json:"action"`
	// Whether the rule is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Identifier of the rule.
	ID string `json:"id"`
	// Objective of the rule.
	Objective string `json:"objective,omitempty"`
	// Rule expression evaluated to determine whether the rule matches.
	Predicate string `json:"predicate,omitempty"`
	// Percentage of matching traffic to which the rule applies, between 0 and 100, serialized as a decimal string.
	RolloutPercentage decimal.Decimal `json:"rollout_percentage,omitempty"`
	// Type of entity evaluated by the rule.
	Type V2IamActivityLogDetailsRadarNewRuleType `json:"type,omitempty"`
}

// Rules before the change. Set for update, removal, enable, and disable actions; absent for add actions.
type V2IamActivityLogDetailsRadarOldRule struct {
	// Action performed when the rule matches.
	Action V2IamActivityLogDetailsRadarOldRuleAction `json:"action"`
	// Whether the rule is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Identifier of the rule.
	ID string `json:"id"`
	// Objective of the rule.
	Objective string `json:"objective,omitempty"`
	// Rule expression evaluated to determine whether the rule matches.
	Predicate string `json:"predicate,omitempty"`
	// Percentage of matching traffic to which the rule applies, between 0 and 100, serialized as a decimal string.
	RolloutPercentage decimal.Decimal `json:"rollout_percentage,omitempty"`
	// Type of entity evaluated by the rule.
	Type V2IamActivityLogDetailsRadarOldRuleType `json:"type,omitempty"`
}

// Details of a Radar action.
type V2IamActivityLogDetailsRadar struct {
	// Risk preference after a preference switch carried by a protection-toggle action.
	NewRiskPreference V2IamActivityLogDetailsRadarNewRiskPreference `json:"new_risk_preference,omitempty"`
	// Rules after the change. Set for add, update, enable, and disable actions; absent for remove actions.
	NewRules []*V2IamActivityLogDetailsRadarNewRule `json:"new_rules,omitempty"`
	// Risk preference before a preference switch carried by a protection-toggle action.
	OldRiskPreference V2IamActivityLogDetailsRadarOldRiskPreference `json:"old_risk_preference,omitempty"`
	// Rules before the change. Set for update, removal, enable, and disable actions; absent for add actions.
	OldRules []*V2IamActivityLogDetailsRadarOldRule `json:"old_rules,omitempty"`
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

// Set when access is granted through an access request.
type V2IamActivityLogDetailsStripeAccessReasonAccessRequest struct {
	// Identifier of the access request.
	ID string `json:"id"`
}

// Set when access is granted through a support relationship.
type V2IamActivityLogDetailsStripeAccessReasonSupportRelationship struct {
	// Identifier of the support relationship.
	ID string `json:"id"`
	// Name of the support contact, when available.
	Name string `json:"name,omitempty"`
	// Role of the support contact, when available.
	Role string `json:"role,omitempty"`
}

// Set when access is granted through a support ticket.
type V2IamActivityLogDetailsStripeAccessReasonSupportTicket struct {
	// Identifier of the support ticket.
	ID string `json:"id"`
}

// The reason for granting Stripe access.
type V2IamActivityLogDetailsStripeAccessReason struct {
	// Set when access is granted through an access request.
	AccessRequest *V2IamActivityLogDetailsStripeAccessReasonAccessRequest `json:"access_request,omitempty"`
	// Set when access is granted through a support relationship.
	SupportRelationship *V2IamActivityLogDetailsStripeAccessReasonSupportRelationship `json:"support_relationship,omitempty"`
	// Set when access is granted through a support ticket.
	SupportTicket *V2IamActivityLogDetailsStripeAccessReasonSupportTicket `json:"support_ticket,omitempty"`
	// The type of reason for granting access.
	Type V2IamActivityLogDetailsStripeAccessReasonType `json:"type"`
}

// Details of a Stripe access action.
type V2IamActivityLogDetailsStripeAccess struct {
	// The level of Dashboard access granted.
	AccessType V2IamActivityLogDetailsStripeAccessAccessType `json:"access_type"`
	// The reason for granting Stripe access.
	Reason *V2IamActivityLogDetailsStripeAccessReason `json:"reason"`
}

// Dashboard client details for the user access action.
type V2IamActivityLogDetailsUserAccessDashboardClient struct {
	// User-Agent header for the user access action.
	UserAgent string `json:"user_agent"`
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
	// Risk signals for the user access action.
	Signals []*V2IamActivityLogDetailsUserAccessRiskSignal `json:"signals"`
}

// Details of a user access action.
type V2IamActivityLogDetailsUserAccess struct {
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
	// Details of a Connect action.
	Connect *V2IamActivityLogDetailsConnect `json:"connect,omitempty"`
	// Details of a data export action.
	DataExport *V2IamActivityLogDetailsDataExport `json:"data_export,omitempty"`
	// Details of an event destination action.
	EventDestination *V2IamActivityLogDetailsEventDestination `json:"event_destination,omitempty"`
	// Details of a Radar action.
	Radar *V2IamActivityLogDetailsRadar `json:"radar,omitempty"`
	// Details of a SCIM action.
	Scim *V2IamActivityLogDetailsScim `json:"scim,omitempty"`
	// Details of an SSO action.
	Sso *V2IamActivityLogDetailsSso `json:"sso,omitempty"`
	// Details of a Stripe access action.
	StripeAccess *V2IamActivityLogDetailsStripeAccess `json:"stripe_access,omitempty"`
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
