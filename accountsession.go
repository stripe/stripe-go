//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The list of features enabled in the embedded component.
type AccountSessionComponentsAccountOnboardingFeaturesParams struct{}

// Configuration for the account onboarding embedded component.
type AccountSessionComponentsAccountOnboardingParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsAccountOnboardingFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsDocumentsFeaturesParams struct{}

// Configuration for the documents embedded component.
type AccountSessionComponentsDocumentsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsDocumentsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsPaymentDetailsFeaturesParams struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments *bool `form:"capture_payments"`
	// Whether to allow connected accounts to manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement *bool `form:"destination_on_behalf_of_charge_management"`
	// Whether to allow responding to disputes, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement *bool `form:"dispute_management"`
	// Whether to allow sending refunds. This is `true` by default.
	RefundManagement *bool `form:"refund_management"`
}

// Configuration for the payment details embedded component.
type AccountSessionComponentsPaymentDetailsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsPaymentDetailsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsPaymentsFeaturesParams struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments *bool `form:"capture_payments"`
	// Whether to allow connected accounts to manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement *bool `form:"destination_on_behalf_of_charge_management"`
	// Whether to allow responding to disputes, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement *bool `form:"dispute_management"`
	// Whether to allow sending refunds. This is `true` by default.
	RefundManagement *bool `form:"refund_management"`
}

// Configuration for the payments embedded component.
type AccountSessionComponentsPaymentsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsPaymentsFeaturesParams `form:"features"`
}

// The list of features enabled in the embedded component.
type AccountSessionComponentsPayoutsFeaturesParams struct {
	// Whether to allow payout schedule to be changed. Default `true` when Stripe owns Loss Liability, default `false` otherwise.
	EditPayoutSchedule *bool `form:"edit_payout_schedule"`
	// Whether to allow creation of instant payouts. Default `true` when Stripe owns Loss Liability, default `false` otherwise.
	InstantPayouts *bool `form:"instant_payouts"`
	// Whether to allow creation of standard payouts. Default `true` when Stripe owns Loss Liability, default `false` otherwise.
	StandardPayouts *bool `form:"standard_payouts"`
}

// Configuration for the payouts embedded component.
type AccountSessionComponentsPayoutsParams struct {
	// Whether the embedded component is enabled.
	Enabled *bool `form:"enabled"`
	// The list of features enabled in the embedded component.
	Features *AccountSessionComponentsPayoutsFeaturesParams `form:"features"`
}

// Each key of the dictionary represents an embedded component, and each embedded component maps to its configuration (e.g. whether it has been enabled or not).
type AccountSessionComponentsParams struct {
	// Configuration for the account onboarding embedded component.
	AccountOnboarding *AccountSessionComponentsAccountOnboardingParams `form:"account_onboarding"`
	// Configuration for the documents embedded component.
	Documents *AccountSessionComponentsDocumentsParams `form:"documents"`
	// Configuration for the payment details embedded component.
	PaymentDetails *AccountSessionComponentsPaymentDetailsParams `form:"payment_details"`
	// Configuration for the payments embedded component.
	Payments *AccountSessionComponentsPaymentsParams `form:"payments"`
	// Configuration for the payouts embedded component.
	Payouts *AccountSessionComponentsPayoutsParams `form:"payouts"`
}

// Creates a AccountSession object that includes a single-use token that the platform can use on their front-end to grant client-side API access.
type AccountSessionParams struct {
	Params `form:"*"`
	// The identifier of the account to create an Account Session for.
	Account *string `form:"account"`
	// Each key of the dictionary represents an embedded component, and each embedded component maps to its configuration (e.g. whether it has been enabled or not).
	Components *AccountSessionComponentsParams `form:"components"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *AccountSessionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type AccountSessionComponentsAccountOnboardingFeatures struct{}
type AccountSessionComponentsAccountOnboarding struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                               `json:"enabled"`
	Features *AccountSessionComponentsAccountOnboardingFeatures `json:"features"`
}
type AccountSessionComponentsDocumentsFeatures struct{}
type AccountSessionComponentsDocuments struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                       `json:"enabled"`
	Features *AccountSessionComponentsDocumentsFeatures `json:"features"`
}
type AccountSessionComponentsPaymentDetailsFeatures struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments bool `json:"capture_payments"`
	// Whether to allow connected accounts to manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement bool `json:"destination_on_behalf_of_charge_management"`
	// Whether to allow responding to disputes, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement bool `json:"dispute_management"`
	// Whether to allow sending refunds. This is `true` by default.
	RefundManagement bool `json:"refund_management"`
}
type AccountSessionComponentsPaymentDetails struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                            `json:"enabled"`
	Features *AccountSessionComponentsPaymentDetailsFeatures `json:"features"`
}
type AccountSessionComponentsPaymentsFeatures struct {
	// Whether to allow capturing and cancelling payment intents. This is `true` by default.
	CapturePayments bool `json:"capture_payments"`
	// Whether to allow connected accounts to manage destination charges that are created on behalf of them. This is `false` by default.
	DestinationOnBehalfOfChargeManagement bool `json:"destination_on_behalf_of_charge_management"`
	// Whether to allow responding to disputes, including submitting evidence and accepting disputes. This is `true` by default.
	DisputeManagement bool `json:"dispute_management"`
	// Whether to allow sending refunds. This is `true` by default.
	RefundManagement bool `json:"refund_management"`
}
type AccountSessionComponentsPayments struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                      `json:"enabled"`
	Features *AccountSessionComponentsPaymentsFeatures `json:"features"`
}
type AccountSessionComponentsPayoutsFeatures struct {
	// Whether to allow payout schedule to be changed. Default `true` when Stripe owns Loss Liability, default `false` otherwise.
	EditPayoutSchedule bool `json:"edit_payout_schedule"`
	// Whether to allow creation of instant payouts. Default `true` when Stripe owns Loss Liability, default `false` otherwise.
	InstantPayouts bool `json:"instant_payouts"`
	// Whether to allow creation of standard payouts. Default `true` when Stripe owns Loss Liability, default `false` otherwise.
	StandardPayouts bool `json:"standard_payouts"`
}
type AccountSessionComponentsPayouts struct {
	// Whether the embedded component is enabled.
	Enabled  bool                                     `json:"enabled"`
	Features *AccountSessionComponentsPayoutsFeatures `json:"features"`
}
type AccountSessionComponents struct {
	AccountOnboarding *AccountSessionComponentsAccountOnboarding `json:"account_onboarding"`
	Documents         *AccountSessionComponentsDocuments         `json:"documents"`
	PaymentDetails    *AccountSessionComponentsPaymentDetails    `json:"payment_details"`
	Payments          *AccountSessionComponentsPayments          `json:"payments"`
	Payouts           *AccountSessionComponentsPayouts           `json:"payouts"`
}

// An AccountSession allows a Connect platform to grant access to a connected account in Connect embedded components.
//
// We recommend that you create an AccountSession each time you need to display an embedded component
// to your user. Do not save AccountSessions to your database as they expire relatively
// quickly, and cannot be used more than once.
//
// Related guide: [Connect embedded components](https://stripe.com/docs/connect/get-started-connect-embedded-components)
type AccountSession struct {
	APIResource
	// The ID of the account the AccountSession was created for
	Account string `json:"account"`
	// The client secret of this AccountSession. Used on the client to set up secure access to the given `account`.
	//
	// The client secret can be used to provide access to `account` from your frontend. It should not be stored, logged, or exposed to anyone other than the connected account. Make sure that you have TLS enabled on any page that includes the client secret.
	//
	// Refer to our docs to [setup Connect embedded components](https://stripe.com/docs/connect/get-started-connect-embedded-components) and learn about how `client_secret` should be handled.
	ClientSecret string                    `json:"client_secret"`
	Components   *AccountSessionComponents `json:"components"`
	// The timestamp at which this AccountSession will expire.
	ExpiresAt int64 `json:"expires_at"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
