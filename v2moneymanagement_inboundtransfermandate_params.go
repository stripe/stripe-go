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
	// Optional prefix for the generated mandate reference (max 10 chars).
	ReferencePrefix *string `form:"reference_prefix" json:"reference_prefix,omitempty"`
}

// Optional details for online acceptance.
type V2MoneyManagementInboundTransferMandateUserAcceptedDetailsOnlineParams struct {
	// The IP address from which the merchant accepted the mandate. For direct account requests,
	// derived from the request when not supplied; rejected if obtainable from neither.
	IPAddress *string `form:"ip_address" json:"ip_address,omitempty"`
	// The user agent of the browser from which the merchant accepted the mandate. For direct
	// account requests, derived from the request when not supplied.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Optional acceptance evidence collected from the merchant. Direct account calls can omit
// details that are derived from request metadata. Platform calls creating a mandate for a
// connected account must provide accepted_at, online.ip_address, and online.user_agent.
type V2MoneyManagementInboundTransferMandateUserAcceptedDetailsParams struct {
	// When the merchant accepted the mandate. Must be a past timestamp. For direct account
	// requests, defaults to the mandate's creation time when not supplied.
	AcceptedAt *time.Time `form:"accepted_at" json:"accepted_at,omitempty"`
	// Optional details for online acceptance.
	Online *V2MoneyManagementInboundTransferMandateUserAcceptedDetailsOnlineParams `form:"online" json:"online,omitempty"`
	// Channel through which acceptance was obtained.
	Type *string `form:"type" json:"type,omitempty"`
}

// Create an InboundTransferMandate for a v2 credential. If a pending or
// active mandate already exists for the same user and credential, that
// mandate is returned instead of creating a new one.
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
	// Optional acceptance evidence collected from the merchant. Direct account calls can omit
	// details that are derived from request metadata. Platform calls creating a mandate for a
	// connected account must provide accepted_at, online.ip_address, and online.user_agent.
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
	// Optional prefix for the generated mandate reference (max 10 chars).
	ReferencePrefix *string `form:"reference_prefix" json:"reference_prefix,omitempty"`
}

// Optional details for online acceptance.
type V2MoneyManagementInboundTransferMandateCreateUserAcceptedDetailsOnlineParams struct {
	// The IP address from which the merchant accepted the mandate. For direct account requests,
	// derived from the request when not supplied; rejected if obtainable from neither.
	IPAddress *string `form:"ip_address" json:"ip_address,omitempty"`
	// The user agent of the browser from which the merchant accepted the mandate. For direct
	// account requests, derived from the request when not supplied.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Optional acceptance evidence collected from the merchant. Direct account calls can omit
// details that are derived from request metadata. Platform calls creating a mandate for a
// connected account must provide accepted_at, online.ip_address, and online.user_agent.
type V2MoneyManagementInboundTransferMandateCreateUserAcceptedDetailsParams struct {
	// When the merchant accepted the mandate. Must be a past timestamp. For direct account
	// requests, defaults to the mandate's creation time when not supplied.
	AcceptedAt *time.Time `form:"accepted_at" json:"accepted_at,omitempty"`
	// Optional details for online acceptance.
	Online *V2MoneyManagementInboundTransferMandateCreateUserAcceptedDetailsOnlineParams `form:"online" json:"online,omitempty"`
	// Channel through which acceptance was obtained.
	Type *string `form:"type" json:"type,omitempty"`
}

// Create an InboundTransferMandate for a v2 credential. If a pending or
// active mandate already exists for the same user and credential, that
// mandate is returned instead of creating a new one.
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
	// Optional acceptance evidence collected from the merchant. Direct account calls can omit
	// details that are derived from request metadata. Platform calls creating a mandate for a
	// connected account must provide accepted_at, online.ip_address, and online.user_agent.
	UserAcceptedDetails *V2MoneyManagementInboundTransferMandateCreateUserAcceptedDetailsParams `form:"user_accepted_details" json:"user_accepted_details,omitempty"`
}

// Retrieve an InboundTransferMandate by ID.
type V2MoneyManagementInboundTransferMandateRetrieveParams struct {
	Params `form:"*"`
}
