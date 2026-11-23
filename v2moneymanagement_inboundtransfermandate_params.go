//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Retrieve a list of InboundTransferMandates for the authenticated compartment.
type V2MoneyManagementInboundTransferMandateListParams struct {
	Params `form:"*"`
	// Filter by v2 credential.
	Credential *string `form:"credential" json:"credential,omitempty"`
	// Maximum number of results to return on a single page.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by mandate status.
	Status *string `form:"status" json:"status,omitempty"`
	// Filter by mandate scheme type.
	Type *string `form:"type" json:"type,omitempty"`
}

// Optional Australian BECS-specific parameters.
type V2MoneyManagementInboundTransferMandateAUBECSParams struct {
	// Optional prefix for the generated 18-character lodgement reference. The prefix is
	// normalized to uppercase and must be empty or contain 1-10 letters, digits, or underscores.
	LodgementReferencePrefix *string `form:"lodgement_reference_prefix" json:"lodgement_reference_prefix,omitempty"`
}

// Optional Bacs-specific parameters.
type V2MoneyManagementInboundTransferMandateBACSParams struct {
	// Optional prefix for the generated mandate reference. When nonempty, it must contain at most
	// 10 letters, digits, or underscores. In live mode, Stripe uppercases the prefix.
	// Prefixes beginning with DDIC or STRIPE are reserved regardless of case.
	// In a Sandbox or test mode, set this parameter to one of the following exact,
	// case-sensitive values to simulate a Bacs mandate outcome:
	// - `ACTIVE`: Creates an active mandate immediately.
	// - `PENDING`: Creates a pending mandate that does not automatically activate.
	// - `FAIL`: Creates a canceled mandate with `status_details.canceled.reason` set to
	// `refused_by_network`.
	// - `EXPIRE`: Creates an expired mandate.
	// - `VALIDATION_FAIL`: Returns an `inbound_mandate_cop_failed` error without creating a mandate.
	// This test-only value is exempt from the normal 10-character prefix limit.
	// If omitted, empty, or set to another valid prefix, creates an active mandate.
	// These simulation values do not select test outcomes in live mode.
	ReferencePrefix *string `form:"reference_prefix" json:"reference_prefix,omitempty"`
}

// Optional details for online acceptance.
type V2MoneyManagementInboundTransferMandateUserAcceptedDetailsOnlineParams struct {
	// The IP address from which the merchant accepted the mandate. For connected account
	// requests, you must provide a valid IPv4 or IPv6 address. For direct account requests,
	// Stripe derives the address from the request when you don't supply it and rejects the
	// request if it can't obtain one.
	IPAddress *string `form:"ip_address" json:"ip_address,omitempty"`
	// The user agent of the browser from which the merchant accepted the mandate. For direct
	// account requests, derived from the request when not supplied.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Acceptance evidence collected from the account holder. Required when a platform creates a
// mandate for a connected account. Include that account's ID in the [Stripe-Context
// header](https://docs.stripe.com/context), and provide accepted_at, online.ip_address, and online.user_agent. Direct
// account calls can omit details that Stripe derives from request metadata.
type V2MoneyManagementInboundTransferMandateUserAcceptedDetailsParams struct {
	// When the account holder accepted the mandate. Must be in the past.
	// For direct account requests, defaults to the mandate's creation time when not supplied.
	AcceptedAt *time.Time `form:"accepted_at" json:"accepted_at,omitempty"`
	// Optional details for online acceptance.
	Online *V2MoneyManagementInboundTransferMandateUserAcceptedDetailsOnlineParams `form:"online" json:"online,omitempty"`
	// Channel through which acceptance was obtained.
	Type *string `form:"type" json:"type,omitempty"`
}

// Create an InboundTransferMandate for a v2 credential. If a pending or active mandate
// already exists for the same user and credential, return that mandate instead of creating
// another one. The existing mandate's reference and acceptance evidence are preserved.
// Retry with the same idempotency key to replay a successful creation.
// In a Sandbox or test mode, use bacs.reference_prefix to simulate Bacs mandate
// outcomes. See the parameter description for supported values.
type V2MoneyManagementInboundTransferMandateParams struct {
	Params `form:"*"`
	// Optional Australian BECS-specific parameters.
	AUBECS *V2MoneyManagementInboundTransferMandateAUBECSParams `form:"au_becs" json:"au_becs,omitempty"`
	// Optional Bacs-specific parameters.
	BACS *V2MoneyManagementInboundTransferMandateBACSParams `form:"bacs" json:"bacs,omitempty"`
	// The v2 credential (GB Bank Account or equivalent) this mandate is created
	// for. Must belong to the authenticated compartment.
	Credential *string `form:"credential" json:"credential,omitempty"`
	// The mandate scheme type.
	Type *string `form:"type" json:"type,omitempty"`
	// Acceptance evidence collected from the account holder. Required when a platform creates a
	// mandate for a connected account. Include that account's ID in the [Stripe-Context
	// header](https://docs.stripe.com/context), and provide accepted_at, online.ip_address, and online.user_agent. Direct
	// account calls can omit details that Stripe derives from request metadata.
	UserAcceptedDetails *V2MoneyManagementInboundTransferMandateUserAcceptedDetailsParams `form:"user_accepted_details" json:"user_accepted_details,omitempty"`
}

// Cancel a pending or active InboundTransferMandate.
type V2MoneyManagementInboundTransferMandateCancelParams struct {
	Params `form:"*"`
}

// Optional Australian BECS-specific parameters.
type V2MoneyManagementInboundTransferMandateCreateAUBECSParams struct {
	// Optional prefix for the generated 18-character lodgement reference. The prefix is
	// normalized to uppercase and must be empty or contain 1-10 letters, digits, or underscores.
	LodgementReferencePrefix *string `form:"lodgement_reference_prefix" json:"lodgement_reference_prefix,omitempty"`
}

// Optional Bacs-specific parameters.
type V2MoneyManagementInboundTransferMandateCreateBACSParams struct {
	// Optional prefix for the generated mandate reference. When nonempty, it must contain at most
	// 10 letters, digits, or underscores. In live mode, Stripe uppercases the prefix.
	// Prefixes beginning with DDIC or STRIPE are reserved regardless of case.
	// In a Sandbox or test mode, set this parameter to one of the following exact,
	// case-sensitive values to simulate a Bacs mandate outcome:
	// - `ACTIVE`: Creates an active mandate immediately.
	// - `PENDING`: Creates a pending mandate that does not automatically activate.
	// - `FAIL`: Creates a canceled mandate with `status_details.canceled.reason` set to
	// `refused_by_network`.
	// - `EXPIRE`: Creates an expired mandate.
	// - `VALIDATION_FAIL`: Returns an `inbound_mandate_cop_failed` error without creating a mandate.
	// This test-only value is exempt from the normal 10-character prefix limit.
	// If omitted, empty, or set to another valid prefix, creates an active mandate.
	// These simulation values do not select test outcomes in live mode.
	ReferencePrefix *string `form:"reference_prefix" json:"reference_prefix,omitempty"`
}

// Optional details for online acceptance.
type V2MoneyManagementInboundTransferMandateCreateUserAcceptedDetailsOnlineParams struct {
	// The IP address from which the merchant accepted the mandate. For connected account
	// requests, you must provide a valid IPv4 or IPv6 address. For direct account requests,
	// Stripe derives the address from the request when you don't supply it and rejects the
	// request if it can't obtain one.
	IPAddress *string `form:"ip_address" json:"ip_address,omitempty"`
	// The user agent of the browser from which the merchant accepted the mandate. For direct
	// account requests, derived from the request when not supplied.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Acceptance evidence collected from the account holder. Required when a platform creates a
// mandate for a connected account. Include that account's ID in the [Stripe-Context
// header](https://docs.stripe.com/context), and provide accepted_at, online.ip_address, and online.user_agent. Direct
// account calls can omit details that Stripe derives from request metadata.
type V2MoneyManagementInboundTransferMandateCreateUserAcceptedDetailsParams struct {
	// When the account holder accepted the mandate. Must be in the past.
	// For direct account requests, defaults to the mandate's creation time when not supplied.
	AcceptedAt *time.Time `form:"accepted_at" json:"accepted_at,omitempty"`
	// Optional details for online acceptance.
	Online *V2MoneyManagementInboundTransferMandateCreateUserAcceptedDetailsOnlineParams `form:"online" json:"online,omitempty"`
	// Channel through which acceptance was obtained.
	Type *string `form:"type" json:"type,omitempty"`
}

// Create an InboundTransferMandate for a v2 credential. If a pending or active mandate
// already exists for the same user and credential, return that mandate instead of creating
// another one. The existing mandate's reference and acceptance evidence are preserved.
// Retry with the same idempotency key to replay a successful creation.
// In a Sandbox or test mode, use bacs.reference_prefix to simulate Bacs mandate
// outcomes. See the parameter description for supported values.
type V2MoneyManagementInboundTransferMandateCreateParams struct {
	Params `form:"*"`
	// Optional Australian BECS-specific parameters.
	AUBECS *V2MoneyManagementInboundTransferMandateCreateAUBECSParams `form:"au_becs" json:"au_becs,omitempty"`
	// Optional Bacs-specific parameters.
	BACS *V2MoneyManagementInboundTransferMandateCreateBACSParams `form:"bacs" json:"bacs,omitempty"`
	// The v2 credential (GB Bank Account or equivalent) this mandate is created
	// for. Must belong to the authenticated compartment.
	Credential *string `form:"credential" json:"credential"`
	// The mandate scheme type.
	Type *string `form:"type" json:"type"`
	// Acceptance evidence collected from the account holder. Required when a platform creates a
	// mandate for a connected account. Include that account's ID in the [Stripe-Context
	// header](https://docs.stripe.com/context), and provide accepted_at, online.ip_address, and online.user_agent. Direct
	// account calls can omit details that Stripe derives from request metadata.
	UserAcceptedDetails *V2MoneyManagementInboundTransferMandateCreateUserAcceptedDetailsParams `form:"user_accepted_details" json:"user_accepted_details,omitempty"`
}

// Retrieve an InboundTransferMandate by ID.
type V2MoneyManagementInboundTransferMandateRetrieveParams struct {
	Params `form:"*"`
}
