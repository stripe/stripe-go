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
	V2IamActivityLogActorTypeAPIKey V2IamActivityLogActorType = "api_key"
	V2IamActivityLogActorTypeUser   V2IamActivityLogActorType = "user"
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

// The action group type of the activity log entry.
type V2IamActivityLogDetailsType string

// List of values that V2IamActivityLogDetailsType can take
const (
	V2IamActivityLogDetailsTypeAPIKey     V2IamActivityLogDetailsType = "api_key"
	V2IamActivityLogDetailsTypeUserInvite V2IamActivityLogDetailsType = "user_invite"
	V2IamActivityLogDetailsTypeUserRoles  V2IamActivityLogDetailsType = "user_roles"
)

// The type of action that was performed.
type V2IamActivityLogType string

// List of values that V2IamActivityLogType can take
const (
	V2IamActivityLogTypeAPIKeyCreated      V2IamActivityLogType = "api_key_created"
	V2IamActivityLogTypeAPIKeyDeleted      V2IamActivityLogType = "api_key_deleted"
	V2IamActivityLogTypeAPIKeyUpdated      V2IamActivityLogType = "api_key_updated"
	V2IamActivityLogTypeAPIKeyViewed       V2IamActivityLogType = "api_key_viewed"
	V2IamActivityLogTypeUserInviteAccepted V2IamActivityLogType = "user_invite_accepted"
	V2IamActivityLogTypeUserInviteCreated  V2IamActivityLogType = "user_invite_created"
	V2IamActivityLogTypeUserInviteDeleted  V2IamActivityLogType = "user_invite_deleted"
	V2IamActivityLogTypeUserRolesDeleted   V2IamActivityLogType = "user_roles_deleted"
	V2IamActivityLogTypeUserRolesUpdated   V2IamActivityLogType = "user_roles_updated"
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

// Details of a user invite action.
type V2IamActivityLogDetailsUserInvite struct {
	// Email address of the invited user.
	InvitedUserEmail string `json:"invited_user_email"`
	// Roles assigned to the invited user.
	Roles []string `json:"roles"`
}

// Details of a user role change action.
type V2IamActivityLogDetailsUserRoles struct {
	// Roles the user has after the change.
	NewRoles []string `json:"new_roles"`
	// Roles the user had before the change.
	OldRoles []string `json:"old_roles"`
	// Email address of the user whose roles were changed.
	UserEmail string `json:"user_email"`
}

// Action-specific details of the activity log entry.
type V2IamActivityLogDetails struct {
	// Details of an API key action.
	APIKey *V2IamActivityLogDetailsAPIKey `json:"api_key,omitempty"`
	// The action group type of the activity log entry.
	Type V2IamActivityLogDetailsType `json:"type"`
	// Details of a user invite action.
	UserInvite *V2IamActivityLogDetailsUserInvite `json:"user_invite,omitempty"`
	// Details of a user role change action.
	UserRoles *V2IamActivityLogDetailsUserRoles `json:"user_roles,omitempty"`
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
	// The type of action that was performed.
	Type V2IamActivityLogType `json:"type"`
}
