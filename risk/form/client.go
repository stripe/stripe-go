//
//
// File generated from our OpenAPI spec
//
//

// Package form provides the /v1/risk/forms APIs
package form

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke /v1/risk/forms APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a risk Form by ID.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
func Get(id string, params *stripe.RiskFormParams) (*stripe.RiskForm, error) {
	return getC().Get(id, params)
}

// Retrieves a risk Form by ID.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.RiskFormParams) (*stripe.RiskForm, error) {
	path := stripe.FormatURLPath("/v1/risk/forms/%s", id)
	form := &stripe.RiskForm{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, form)
	return form, err
}

// Updates a risk Form with the connected account's response.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
// | form_closed | 400 | The Form is closed and cannot accept updates. |
// | invalid_field_for_form_type | 400 | Submitted fields contain invalid values. |
// | form_field_already_submitted | 400 | A submitted field has already been submitted. |
// | form_insufficient_fields | 400 | Required fields are missing for this Form type. |
// | resource_missing | 404 | A submitted file does not exist or is not accessible to the requesting account. |
//
// Validation errors include a summary in message. Missing required properties use
// form_insufficient_fields and include missing_fields, an array of public dotted paths.
// Missing required properties take precedence over value violations.
//
// Invalid values use invalid_field_for_form_type and include violations. Each violation
// has a public dotted field path and a reason of missing_value or invalid_value.
// Reviewed-field violations include a corrective message; corrective messages for other
// Form types are omitted. For an array member, target identifies the member's field path.
// For a scalar value, target is explicitly null.
//
// For example, a reviewed-fields submission with an out-of-date timestamp and an omitted
// requested member returns both violations:
//
// json
//
//	{
//	 "error": {
//	 "type": "invalid_request_error",
//	 "code": "invalid_field_for_form_type",
//	 "message": "Submitted fields are invalid for Form type reviewed_fields. See violations for details and corrective actions for these fields: 'reviewed_fields.attestation.attested_at', 'reviewed_fields.attestation.fields'.",
//	 "violations": [
//	 {
//	 "field": "reviewed_fields.attestation.attested_at",
//	 "reason": "invalid_value",
//	 "message": "The attestation timestamp is out of date. Review and attest the requested information again, then provide the new attestation timestamp.",
//	 "target": null
//	 },
//	 {
//	 "field": "reviewed_fields.attestation.fields",
//	 "reason": "missing_value",
//	 "message": "'defaults.profile.product_description' is required by subject.fields_to_review but is missing from reviewed_fields.attestation.fields. Review this information, then include its path in the attestation.",
//	 "target": "defaults.profile.product_description"
//	 }
//	 ]
//	 }
//	}
//
// A v1 request with no extracted field submissions is rejected before the Form is loaded,
// including for reviewed-fields Forms. This early error uses form_insufficient_fields
// and omits missing_fields,
// violations, and param. Partially populated reviewed-fields requests use structured
// missing-property validation.
func Update(id string, params *stripe.RiskFormParams) (*stripe.RiskForm, error) {
	return getC().Update(id, params)
}

// Updates a risk Form with the connected account's response.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
// | form_closed | 400 | The Form is closed and cannot accept updates. |
// | invalid_field_for_form_type | 400 | Submitted fields contain invalid values. |
// | form_field_already_submitted | 400 | A submitted field has already been submitted. |
// | form_insufficient_fields | 400 | Required fields are missing for this Form type. |
// | resource_missing | 404 | A submitted file does not exist or is not accessible to the requesting account. |
//
// Validation errors include a summary in message. Missing required properties use
// form_insufficient_fields and include missing_fields, an array of public dotted paths.
// Missing required properties take precedence over value violations.
//
// Invalid values use invalid_field_for_form_type and include violations. Each violation
// has a public dotted field path and a reason of missing_value or invalid_value.
// Reviewed-field violations include a corrective message; corrective messages for other
// Form types are omitted. For an array member, target identifies the member's field path.
// For a scalar value, target is explicitly null.
//
// For example, a reviewed-fields submission with an out-of-date timestamp and an omitted
// requested member returns both violations:
//
// json
//
//	{
//	 "error": {
//	 "type": "invalid_request_error",
//	 "code": "invalid_field_for_form_type",
//	 "message": "Submitted fields are invalid for Form type reviewed_fields. See violations for details and corrective actions for these fields: 'reviewed_fields.attestation.attested_at', 'reviewed_fields.attestation.fields'.",
//	 "violations": [
//	 {
//	 "field": "reviewed_fields.attestation.attested_at",
//	 "reason": "invalid_value",
//	 "message": "The attestation timestamp is out of date. Review and attest the requested information again, then provide the new attestation timestamp.",
//	 "target": null
//	 },
//	 {
//	 "field": "reviewed_fields.attestation.fields",
//	 "reason": "missing_value",
//	 "message": "'defaults.profile.product_description' is required by subject.fields_to_review but is missing from reviewed_fields.attestation.fields. Review this information, then include its path in the attestation.",
//	 "target": "defaults.profile.product_description"
//	 }
//	 ]
//	 }
//	}
//
// A v1 request with no extracted field submissions is rejected before the Form is loaded,
// including for reviewed-fields Forms. This early error uses form_insufficient_fields
// and omits missing_fields,
// violations, and param. Partially populated reviewed-fields requests use structured
// missing-property validation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.RiskFormParams) (*stripe.RiskForm, error) {
	path := stripe.FormatURLPath("/v1/risk/forms/%s", id)
	form := &stripe.RiskForm{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, form)
	return form, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
