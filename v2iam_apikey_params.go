//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all API keys of an account.
type V2IamAPIKeyListParams struct {
	Params `form:"*"`
	// Whether to include expired API keys in the response.
	IncludeExpired *bool `form:"include_expired" json:"include_expired,omitempty"`
	// Limit of items to return per page.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// PEM-formatted public key.
type V2IamAPIKeyPublicKeyPemKeyParams struct {
	// The encryption algorithm used with this key (e.g., RSA).
	Algorithm *string `form:"algorithm" json:"algorithm"`
	// The PEM-encoded public key data. Newlines are required between header/footer and body, and optional within the body.
	Data *string `form:"data" json:"data"`
}

// Public key for encrypting the API key secret.
// This must a PEM-formatted RSA key suitable for encryption, >= 2048 bits.
// A public key is required when creating secret keys.
// Publishable keys are not encrypted and a public key should not be included.
type V2IamAPIKeyPublicKeyParams struct {
	// Caller's identifier of the public key. This is used for tracking purposes only, and will be echoed in the response if provided.
	ID *string `form:"id" json:"id,omitempty"`
	// PEM-formatted public key.
	PemKey *V2IamAPIKeyPublicKeyPemKeyParams `form:"pem_key" json:"pem_key,omitempty"`
}

// Create an API key.
type V2IamAPIKeyParams struct {
	Params `form:"*"`
	// Name to set for the API key. If blank, the field is left unchanged.
	Name *string `form:"name" json:"name,omitempty"`
	// Note or description to set for the API key. If blank, the field is left unchanged.
	Note *string `form:"note" json:"note,omitempty"`
	// Public key for encrypting the API key secret.
	// This must a PEM-formatted RSA key suitable for encryption, >= 2048 bits.
	// A public key is required when creating secret keys.
	// Publishable keys are not encrypted and a public key should not be included.
	PublicKey *V2IamAPIKeyPublicKeyParams `form:"public_key" json:"public_key,omitempty"`
	// Type of the API key to create (secret or publishable).
	Type *string `form:"type" json:"type,omitempty"`
}

// Expire an API key.
type V2IamAPIKeyExpireParams struct {
	Params `form:"*"`
}

// PEM-formatted public key.
type V2IamAPIKeyRotatePublicKeyPemKeyParams struct {
	// The encryption algorithm used with this key (e.g., RSA).
	Algorithm *string `form:"algorithm" json:"algorithm"`
	// The PEM-encoded public key data. Newlines are required between header/footer and body, and optional within the body.
	Data *string `form:"data" json:"data"`
}

// Public key for encrypting the new API key secret.
// This must a PEM-formatted RSA key suitable for encryption, >= 2048 bits.
// A public key is required when rotating secret keys.
// Publishable keys are not encrypted and a public key should not be included.
type V2IamAPIKeyRotatePublicKeyParams struct {
	// Caller's identifier of the public key. This is used for tracking purposes only, and will be echoed in the response if provided.
	ID *string `form:"id" json:"id,omitempty"`
	// PEM-formatted public key.
	PemKey *V2IamAPIKeyRotatePublicKeyPemKeyParams `form:"pem_key" json:"pem_key,omitempty"`
}

// Rotate an API key.
type V2IamAPIKeyRotateParams struct {
	Params `form:"*"`
	// Duration in minutes before the current key expires, with a maximum of 7 days (10080 minutes).
	// If not provided, the current key expires immediately.
	ExpireCurrentKeyInMinutes *int64 `form:"expire_current_key_in_minutes" json:"expire_current_key_in_minutes,omitempty"`
	// Public key for encrypting the new API key secret.
	// This must a PEM-formatted RSA key suitable for encryption, >= 2048 bits.
	// A public key is required when rotating secret keys.
	// Publishable keys are not encrypted and a public key should not be included.
	PublicKey *V2IamAPIKeyRotatePublicKeyParams `form:"public_key" json:"public_key,omitempty"`
}

// PEM-formatted public key.
type V2IamAPIKeyCreatePublicKeyPemKeyParams struct {
	// The encryption algorithm used with this key (e.g., RSA).
	Algorithm *string `form:"algorithm" json:"algorithm"`
	// The PEM-encoded public key data. Newlines are required between header/footer and body, and optional within the body.
	Data *string `form:"data" json:"data"`
}

// Public key for encrypting the API key secret.
// This must a PEM-formatted RSA key suitable for encryption, >= 2048 bits.
// A public key is required when creating secret keys.
// Publishable keys are not encrypted and a public key should not be included.
type V2IamAPIKeyCreatePublicKeyParams struct {
	// Caller's identifier of the public key. This is used for tracking purposes only, and will be echoed in the response if provided.
	ID *string `form:"id" json:"id,omitempty"`
	// PEM-formatted public key.
	PemKey *V2IamAPIKeyCreatePublicKeyPemKeyParams `form:"pem_key" json:"pem_key,omitempty"`
}

// Create an API key.
type V2IamAPIKeyCreateParams struct {
	Params `form:"*"`
	// Name for the API key.
	Name *string `form:"name" json:"name,omitempty"`
	// Note or description for the API key.
	Note *string `form:"note" json:"note,omitempty"`
	// Public key for encrypting the API key secret.
	// This must a PEM-formatted RSA key suitable for encryption, >= 2048 bits.
	// A public key is required when creating secret keys.
	// Publishable keys are not encrypted and a public key should not be included.
	PublicKey *V2IamAPIKeyCreatePublicKeyParams `form:"public_key" json:"public_key,omitempty"`
	// Type of the API key to create (secret or publishable).
	Type *string `form:"type" json:"type"`
}

// Retrieve an API key.
type V2IamAPIKeyRetrieveParams struct {
	Params `form:"*"`
}

// Update an API key.
type V2IamAPIKeyUpdateParams struct {
	Params `form:"*"`
	// Name to set for the API key. If blank, the field is left unchanged.
	Name *string `form:"name" json:"name,omitempty"`
	// Note or description to set for the API key. If blank, the field is left unchanged.
	Note *string `form:"note" json:"note,omitempty"`
}
