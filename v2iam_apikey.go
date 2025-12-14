//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of entity.
type V2IamAPIKeyManagedByType string

// List of values that V2IamAPIKeyManagedByType can take
const (
	V2IamAPIKeyManagedByTypeApplication V2IamAPIKeyManagedByType = "application"
)

// Current status of the API key (e.g., active, expired).
type V2IamAPIKeyStatus string

// List of values that V2IamAPIKeyStatus can take
const (
	V2IamAPIKeyStatusActive  V2IamAPIKeyStatus = "active"
	V2IamAPIKeyStatusExpired V2IamAPIKeyStatus = "expired"
)

// Type of the API key.
type V2IamAPIKeyType string

// List of values that V2IamAPIKeyType can take
const (
	V2IamAPIKeyTypePublishableKey V2IamAPIKeyType = "publishable_key"
	V2IamAPIKeyTypeSecretKey      V2IamAPIKeyType = "secret_key"
)

// An application.
type V2IamAPIKeyManagedByApplication struct {
	// Identifier of the application.
	ID string `json:"id"`
}

// Account that manages this API key (for keys managed by platforms).
type V2IamAPIKeyManagedBy struct {
	// An application.
	Application *V2IamAPIKeyManagedByApplication `json:"application,omitempty"`
	// The type of entity.
	Type V2IamAPIKeyManagedByType `json:"type"`
}

// Token set for a publishable key.
type V2IamAPIKeyPublishableKey struct {
	// The plaintext token for the API key.
	Token string `json:"token"`
}

// The encrypted secret for the API key. Only included when a key is first created.
type V2IamAPIKeySecretKeyEncryptedSecret struct {
	// The encrypted secret data in base64 format.
	Ciphertext string `json:"ciphertext"`
	// The format of the encrypted secret (e.g., jwe_compact).
	Format string `json:"format"`
	// The caller's identifier of the public key provided.
	RecipientKeyID string `json:"recipient_key_id,omitempty"`
}

// Token set for a secret key.
type V2IamAPIKeySecretKey struct {
	// The encrypted secret for the API key. Only included when a key is first created.
	EncryptedSecret *V2IamAPIKeySecretKeyEncryptedSecret `json:"encrypted_secret,omitempty"`
	// Redacted version of the secret token for display purposes.
	SecretTokenRedacted string `json:"secret_token_redacted,omitempty"`
	// The plaintext token for the API key. Only included for testmode keys.
	Token string `json:"token,omitempty"`
}

// An API key.
type V2IamAPIKey struct {
	APIResource
	// Timestamp when the API key was created.
	Created time.Time `json:"created"`
	// Timestamp when the API key expires.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Unique identifier of the API key.
	ID string `json:"id"`
	// List of IP addresses allowed to use this API key. Addresses use IPv4 protocol, and may be a CIDR range (e.g., [100.10.38.255, 100.10.38.0/24]).
	IPAllowlist []string `json:"ip_allowlist"`
	// Timestamp when the API key was last used.
	LastUsed time.Time `json:"last_used,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Account that manages this API key (for keys managed by platforms).
	ManagedBy *V2IamAPIKeyManagedBy `json:"managed_by,omitempty"`
	// Name of the API key.
	Name string `json:"name,omitempty"`
	// Note or description for the API key.
	Note string `json:"note,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Token set for a publishable key.
	PublishableKey *V2IamAPIKeyPublishableKey `json:"publishable_key,omitempty"`
	// Token set for a secret key.
	SecretKey *V2IamAPIKeySecretKey `json:"secret_key,omitempty"`
	// Current status of the API key (e.g., active, expired).
	Status V2IamAPIKeyStatus `json:"status"`
	// Type of the API key.
	Type V2IamAPIKeyType `json:"type"`
}
