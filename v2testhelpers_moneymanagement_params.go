//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a RecipientVerification with a specified match result for testing purposes in a Sandbox environment.
type V2TestHelpersMoneyManagementRecipientVerificationsParams struct {
	Params `form:"*"`
	// Expected match level of the RecipientVerification to be created: `match`, `close_match`, `no_match`, `unavailable`.
	// For `close_match`, the simulated response appends "close_match" to the provided name in match_result_details.matched_name.
	MatchResult *string `form:"match_result" json:"match_result"`
	// ID of the payout method.
	PayoutMethod *string `form:"payout_method" json:"payout_method"`
	// ID of the recipient account. Required if the recipient distinct from the sender. Leave empty if the recipient and sender are the same entity (i.e. for me-to-me payouts).
	Recipient *string `form:"recipient" json:"recipient,omitempty"`
}
