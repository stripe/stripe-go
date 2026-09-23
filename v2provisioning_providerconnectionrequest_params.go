//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a new provider connection.
type V2ProvisioningProviderConnectionRequestParams struct {
	Params `form:"*"`
	// PKCE code challenge: BASE64URL(SHA256(code_verifier)). Optional; when present the OAuth
	// callback must supply the matching code_verifier. Not a secret (it is a hash of the verifier).
	CodeChallenge *string `form:"code_challenge" json:"code_challenge,omitempty"`
	// PKCE code challenge method. Only "S256" is supported.
	CodeChallengeMethod *string `form:"code_challenge_method" json:"code_challenge_method,omitempty"`
	// Provider-specific configuration payload for the connection.
	Configuration map[string]any `form:"configuration" json:"configuration,omitempty"`
	// Project this provider connection is created for. Used to infer the catalog partition for provider
	// calls. Optional; when absent the provider connection defaults to the prod catalog.
	Project *string `form:"project" json:"project,omitempty"`
	// Identifier of the provider to connect to.
	Provider *string `form:"provider" json:"provider,omitempty"`
	// Deprecated identifier of the provider to connect to; use `provider` instead.
	ProviderName *string `form:"provider_name" json:"provider_name,omitempty"`
}

// Submits additional information requested by the provider for a provider connection.
type V2ProvisioningProviderConnectionRequestSubmitInformationParams struct {
	Params `form:"*"`
	// Secret used to confirm the request when submitting on behalf of a resource without
	// an authenticated session.
	ConfirmationSecret *string `form:"confirmation_secret" json:"confirmation_secret,omitempty"`
	// Information requested by the provider, matching the connection's needs_information_schema.
	Information map[string]any `form:"information" json:"information"`
}

// Creates a new provider connection.
type V2ProvisioningProviderConnectionRequestCreateParams struct {
	Params `form:"*"`
	// PKCE code challenge: BASE64URL(SHA256(code_verifier)). Optional; when present the OAuth
	// callback must supply the matching code_verifier. Not a secret (it is a hash of the verifier).
	CodeChallenge *string `form:"code_challenge" json:"code_challenge,omitempty"`
	// PKCE code challenge method. Only "S256" is supported.
	CodeChallengeMethod *string `form:"code_challenge_method" json:"code_challenge_method,omitempty"`
	// Provider-specific configuration payload for the connection.
	Configuration map[string]any `form:"configuration" json:"configuration"`
	// Project this provider connection is created for. Used to infer the catalog partition for provider
	// calls. Optional; when absent the provider connection defaults to the prod catalog.
	Project *string `form:"project" json:"project,omitempty"`
	// Identifier of the provider to connect to.
	Provider *string `form:"provider" json:"provider,omitempty"`
	// Deprecated identifier of the provider to connect to; use `provider` instead.
	ProviderName *string `form:"provider_name" json:"provider_name,omitempty"`
}

// Retrieves a provider connection.
type V2ProvisioningProviderConnectionRequestRetrieveParams struct {
	Params `form:"*"`
}
