//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The `Provider` resource represents a third-party provider available in the
// provisioning catalog.
type V2ProvisioningProvider struct {
	APIResource
	// Capabilities supported by the provider.
	Capabilities []string `json:"capabilities"`
	// Categories the provider belongs to.
	Categories []string `json:"categories"`
	// Schema describing the configuration accepted by this provider.
	ConfigurationSchema map[string]any `json:"configuration_schema"`
	// Time at which the provider was created.
	Created time.Time `json:"created"`
	// Deep-link purposes supported by the provider.
	DeepLinkPurposes []string `json:"deep_link_purposes"`
	// Description of the provider.
	Description string `json:"description"`
	// proto3 scalar defaults apply: if unset, this value is `false`.
	Development bool `json:"development"`
	// Unique identifier for the provider.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// URL of additional context about the provider intended for LLM consumption.
	LlmContext string `json:"llm_context,omitempty"`
	// Human-readable name of the provider.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// URL of the provider's privacy policy.
	PrivacyPolicyURL string `json:"privacy_policy_url,omitempty"`
	// URL of the provider's terms of service.
	TOSURL string `json:"tos_url,omitempty"`
	// URL of the provider's website.
	WebsiteURL string `json:"website_url,omitempty"`
}
