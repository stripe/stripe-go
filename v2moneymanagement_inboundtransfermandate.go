//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The current lifecycle status of the mandate.
type V2MoneyManagementInboundTransferMandateStatus string

// List of values that V2MoneyManagementInboundTransferMandateStatus can take
const (
	V2MoneyManagementInboundTransferMandateStatusActive   V2MoneyManagementInboundTransferMandateStatus = "active"
	V2MoneyManagementInboundTransferMandateStatusCanceled V2MoneyManagementInboundTransferMandateStatus = "canceled"
	V2MoneyManagementInboundTransferMandateStatusExpired  V2MoneyManagementInboundTransferMandateStatus = "expired"
	V2MoneyManagementInboundTransferMandateStatusPending  V2MoneyManagementInboundTransferMandateStatus = "pending"
)

// The reason the mandate was canceled.
type V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReason string

// List of values that V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReason can take
const (
	V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReasonCanceledByNetwork V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReason = "canceled_by_network"
	V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReasonCanceledByUser    V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReason = "canceled_by_user"
	V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReasonRefusedByNetwork  V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReason = "refused_by_network"
	V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReasonRevokedByStripe   V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReason = "revoked_by_stripe"
)

// The mandate scheme type.
type V2MoneyManagementInboundTransferMandateType string

// List of values that V2MoneyManagementInboundTransferMandateType can take
const (
	V2MoneyManagementInboundTransferMandateTypeAUBECS V2MoneyManagementInboundTransferMandateType = "au_becs"
	V2MoneyManagementInboundTransferMandateTypeBACS   V2MoneyManagementInboundTransferMandateType = "bacs"
	V2MoneyManagementInboundTransferMandateTypeNzBECS V2MoneyManagementInboundTransferMandateType = "nz_becs"
	V2MoneyManagementInboundTransferMandateTypeSEPA   V2MoneyManagementInboundTransferMandateType = "sepa"
)

// Channel through which acceptance was obtained.
type V2MoneyManagementInboundTransferMandateUserAcceptedDetailsType string

// List of values that V2MoneyManagementInboundTransferMandateUserAcceptedDetailsType can take
const (
	V2MoneyManagementInboundTransferMandateUserAcceptedDetailsTypeOnline V2MoneyManagementInboundTransferMandateUserAcceptedDetailsType = "online"
)

// Australian BECS-specific details. Present when type is AU_BECS.
type V2MoneyManagementInboundTransferMandateAUBECS struct {
	// The generated AU BECS lodgement reference. It is 18 uppercase alphanumeric or underscore
	// characters and incorporates lodgement_reference_prefix when one was supplied at creation.
	LodgementReference string `json:"lodgement_reference"`
}

// Bacs-specific details. Present when type is BACS.
type V2MoneyManagementInboundTransferMandateBACS struct {
	// The generated Bacs mandate reference. May incorporate the optional
	// reference_prefix supplied at creation time.
	Reference string `json:"reference"`
}

// Present when the mandate is in the CANCELED state.
type V2MoneyManagementInboundTransferMandateStatusDetailsCanceled struct {
	// The reason the mandate was canceled.
	Reason V2MoneyManagementInboundTransferMandateStatusDetailsCanceledReason `json:"reason"`
}

// Additional details about the current status (e.g. cancelation reason).
type V2MoneyManagementInboundTransferMandateStatusDetails struct {
	// Present when the mandate is in the CANCELED state.
	Canceled *V2MoneyManagementInboundTransferMandateStatusDetailsCanceled `json:"canceled,omitempty"`
}

// Timestamps for each state transition.
type V2MoneyManagementInboundTransferMandateStatusTransitions struct {
	// When the mandate became active.
	ActivatedAt time.Time `json:"activated_at,omitempty"`
	// When the mandate was canceled.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// When the mandate expired.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// Optional details for online acceptance.
type V2MoneyManagementInboundTransferMandateUserAcceptedDetailsOnline struct {
	// The IP address from which the merchant accepted the mandate. For direct account requests,
	// derived from the request when not supplied; rejected if obtainable from neither.
	IPAddress string `json:"ip_address,omitempty"`
	// The user agent of the browser from which the merchant accepted the mandate. For direct
	// account requests, derived from the request when not supplied.
	UserAgent string `json:"user_agent,omitempty"`
}

// Evidence of the merchant's acceptance of the mandate.
type V2MoneyManagementInboundTransferMandateUserAcceptedDetails struct {
	// When the merchant accepted the mandate. Must be a past timestamp. For direct account
	// requests, defaults to the mandate's creation time when not supplied.
	AcceptedAt time.Time `json:"accepted_at,omitempty"`
	// Optional details for online acceptance.
	Online *V2MoneyManagementInboundTransferMandateUserAcceptedDetailsOnline `json:"online,omitempty"`
	// Channel through which acceptance was obtained.
	Type V2MoneyManagementInboundTransferMandateUserAcceptedDetailsType `json:"type,omitempty"`
}

// An InboundTransferMandate represents Stripe's authorization to debit a
// merchant's external bank account (v2 credential) on their behalf.
type V2MoneyManagementInboundTransferMandate struct {
	APIResource
	// Australian BECS-specific details. Present when type is AU_BECS.
	AUBECS *V2MoneyManagementInboundTransferMandateAUBECS `json:"au_becs,omitempty"`
	// Bacs-specific details. Present when type is BACS.
	BACS *V2MoneyManagementInboundTransferMandateBACS `json:"bacs,omitempty"`
	// Creation time of the mandate. RFC 3339 UTC, millisecond precision.
	Created time.Time `json:"created"`
	// The v2 credential (e.g. GB Bank Account) this mandate authorizes debits for.
	Credential string `json:"credential"`
	// Unique identifier for the InboundTransferMandate.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The current lifecycle status of the mandate.
	Status V2MoneyManagementInboundTransferMandateStatus `json:"status"`
	// Additional details about the current status (e.g. cancelation reason).
	StatusDetails *V2MoneyManagementInboundTransferMandateStatusDetails `json:"status_details"`
	// Timestamps for each state transition.
	StatusTransitions *V2MoneyManagementInboundTransferMandateStatusTransitions `json:"status_transitions"`
	// The mandate scheme type.
	Type V2MoneyManagementInboundTransferMandateType `json:"type"`
	// Evidence of the merchant's acceptance of the mandate.
	UserAcceptedDetails *V2MoneyManagementInboundTransferMandateUserAcceptedDetails `json:"user_accepted_details"`
}
