//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// How the FundingIntent is confirmed after validation.
type V2MoneyManagementFundingIntentConfirmationMethod string

// List of values that V2MoneyManagementFundingIntentConfirmationMethod can take
const (
	V2MoneyManagementFundingIntentConfirmationMethodAutomatic V2MoneyManagementFundingIntentConfirmationMethod = "automatic"
	V2MoneyManagementFundingIntentConfirmationMethodManual    V2MoneyManagementFundingIntentConfirmationMethod = "manual"
)

// The type of underlying money movement created by the attempt.
type V2MoneyManagementFundingIntentLatestAttemptType string

// List of values that V2MoneyManagementFundingIntentLatestAttemptType can take
const (
	V2MoneyManagementFundingIntentLatestAttemptTypeInboundFunding V2MoneyManagementFundingIntentLatestAttemptType = "inbound_funding"
)

// The reason the attempt failed.
type V2MoneyManagementFundingIntentNextActionHandleFailureReason string

// List of values that V2MoneyManagementFundingIntentNextActionHandleFailureReason can take
const (
	V2MoneyManagementFundingIntentNextActionHandleFailureReasonInsufficientFunds V2MoneyManagementFundingIntentNextActionHandleFailureReason = "insufficient_funds"
	V2MoneyManagementFundingIntentNextActionHandleFailureReasonNotAuthorized     V2MoneyManagementFundingIntentNextActionHandleFailureReason = "not_authorized"
	V2MoneyManagementFundingIntentNextActionHandleFailureReasonStripeRejected    V2MoneyManagementFundingIntentNextActionHandleFailureReason = "stripe_rejected"
)

// A machine-readable reason for the factor.
type V2MoneyManagementFundingIntentNextActionProvideNewCheckReason string

// List of values that V2MoneyManagementFundingIntentNextActionProvideNewCheckReason can take
const (
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonAmountVarianceExceedsThreshold V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "amount_variance_exceeds_threshold"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonAmountVarianceNearingThreshold V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "amount_variance_nearing_threshold"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonBackImageNotProcessed          V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "back_image_not_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonBackLooksLikeFront             V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "back_looks_like_front"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonBackSignatureNotDetected       V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "back_signature_not_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonBadPayeeName                   V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "bad_payee_name"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonDateInvalid                    V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "date_invalid"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonDateInFuture                   V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "date_in_future"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonDateOlderThanOneYear           V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "date_older_than_one_year"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonDateOlderThanSixMonths         V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "date_older_than_six_months"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonDuplicateCheck                 V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "duplicate_check"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonEndorsementMissing             V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "endorsement_missing"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonFoldedOrTornCorners            V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "folded_or_torn_corners"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonFrontImageDimensionMismatch    V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "front_image_dimension_mismatch"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonFrontImageNotProcessed         V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "front_image_not_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonFrontImageTooDark              V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "front_image_too_dark"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonFrontImageTooLight             V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "front_image_too_light"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonFrontImageTooSmall             V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "front_image_too_small"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonFrontSignatureNotDetected      V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "front_signature_not_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonHighRisk                       V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "high_risk"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonImageCannotBeProcessed         V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "image_cannot_be_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonImageOutOfFocus                V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "image_out_of_focus"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonInvalidCheckType               V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "invalid_check_type"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonInvalidRoutingNumber           V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "invalid_routing_number"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonNotAcceptedPleaseContact       V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "not_accepted_please_contact"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonNoRoutingNumberDetected        V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "no_routing_number_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonSafeHarborPhraseMissing        V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "safe_harbor_phrase_missing"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonSavingsBondDetected            V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "savings_bond_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonTravelersCheck                 V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "travelers_check"
	V2MoneyManagementFundingIntentNextActionProvideNewCheckReasonUnableToReadAccountInfo        V2MoneyManagementFundingIntentNextActionProvideNewCheckReason = "unable_to_read_account_info"
)

// A machine-readable reason for the factor.
type V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason string

// List of values that V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason can take
const (
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonAmountVarianceExceedsThreshold V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "amount_variance_exceeds_threshold"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonAmountVarianceNearingThreshold V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "amount_variance_nearing_threshold"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonBackImageNotProcessed          V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "back_image_not_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonBackLooksLikeFront             V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "back_looks_like_front"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonBackSignatureNotDetected       V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "back_signature_not_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonBadPayeeName                   V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "bad_payee_name"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonDateInvalid                    V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "date_invalid"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonDateInFuture                   V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "date_in_future"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonDateOlderThanOneYear           V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "date_older_than_one_year"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonDateOlderThanSixMonths         V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "date_older_than_six_months"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonDuplicateCheck                 V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "duplicate_check"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonEndorsementMissing             V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "endorsement_missing"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonFoldedOrTornCorners            V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "folded_or_torn_corners"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonFrontImageDimensionMismatch    V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "front_image_dimension_mismatch"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonFrontImageNotProcessed         V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "front_image_not_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonFrontImageTooDark              V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "front_image_too_dark"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonFrontImageTooLight             V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "front_image_too_light"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonFrontImageTooSmall             V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "front_image_too_small"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonFrontSignatureNotDetected      V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "front_signature_not_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonHighRisk                       V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "high_risk"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonImageCannotBeProcessed         V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "image_cannot_be_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonImageOutOfFocus                V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "image_out_of_focus"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonInvalidCheckType               V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "invalid_check_type"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonInvalidRoutingNumber           V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "invalid_routing_number"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonNotAcceptedPleaseContact       V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "not_accepted_please_contact"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonNoRoutingNumberDetected        V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "no_routing_number_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonSafeHarborPhraseMissing        V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "safe_harbor_phrase_missing"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonSavingsBondDetected            V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "savings_bond_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonTravelersCheck                 V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "travelers_check"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReasonUnableToReadAccountInfo        V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason = "unable_to_read_account_info"
)

// A machine-readable reason for the factor.
type V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason string

// List of values that V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason can take
const (
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonAmountVarianceExceedsThreshold V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "amount_variance_exceeds_threshold"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonAmountVarianceNearingThreshold V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "amount_variance_nearing_threshold"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonBackImageNotProcessed          V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "back_image_not_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonBackLooksLikeFront             V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "back_looks_like_front"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonBackSignatureNotDetected       V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "back_signature_not_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonBadPayeeName                   V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "bad_payee_name"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonDateInvalid                    V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "date_invalid"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonDateInFuture                   V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "date_in_future"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonDateOlderThanOneYear           V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "date_older_than_one_year"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonDateOlderThanSixMonths         V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "date_older_than_six_months"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonDuplicateCheck                 V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "duplicate_check"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonEndorsementMissing             V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "endorsement_missing"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonFoldedOrTornCorners            V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "folded_or_torn_corners"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonFrontImageDimensionMismatch    V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "front_image_dimension_mismatch"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonFrontImageNotProcessed         V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "front_image_not_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonFrontImageTooDark              V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "front_image_too_dark"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonFrontImageTooLight             V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "front_image_too_light"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonFrontImageTooSmall             V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "front_image_too_small"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonFrontSignatureNotDetected      V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "front_signature_not_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonHighRisk                       V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "high_risk"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonImageCannotBeProcessed         V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "image_cannot_be_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonImageOutOfFocus                V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "image_out_of_focus"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonInvalidCheckType               V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "invalid_check_type"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonInvalidRoutingNumber           V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "invalid_routing_number"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonNotAcceptedPleaseContact       V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "not_accepted_please_contact"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonNoRoutingNumberDetected        V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "no_routing_number_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonSafeHarborPhraseMissing        V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "safe_harbor_phrase_missing"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonSavingsBondDetected            V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "savings_bond_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonTravelersCheck                 V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "travelers_check"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReasonUnableToReadAccountInfo        V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason = "unable_to_read_account_info"
)

// A machine-readable reason for the factor.
type V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason string

// List of values that V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason can take
const (
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonAmountVarianceExceedsThreshold V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "amount_variance_exceeds_threshold"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonAmountVarianceNearingThreshold V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "amount_variance_nearing_threshold"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonBackImageNotProcessed          V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "back_image_not_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonBackLooksLikeFront             V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "back_looks_like_front"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonBackSignatureNotDetected       V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "back_signature_not_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonBadPayeeName                   V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "bad_payee_name"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonDateInvalid                    V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "date_invalid"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonDateInFuture                   V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "date_in_future"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonDateOlderThanOneYear           V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "date_older_than_one_year"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonDateOlderThanSixMonths         V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "date_older_than_six_months"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonDuplicateCheck                 V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "duplicate_check"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonEndorsementMissing             V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "endorsement_missing"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonFoldedOrTornCorners            V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "folded_or_torn_corners"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonFrontImageDimensionMismatch    V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "front_image_dimension_mismatch"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonFrontImageNotProcessed         V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "front_image_not_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonFrontImageTooDark              V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "front_image_too_dark"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonFrontImageTooLight             V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "front_image_too_light"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonFrontImageTooSmall             V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "front_image_too_small"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonFrontSignatureNotDetected      V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "front_signature_not_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonHighRisk                       V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "high_risk"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonImageCannotBeProcessed         V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "image_cannot_be_processed"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonImageOutOfFocus                V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "image_out_of_focus"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonInvalidCheckType               V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "invalid_check_type"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonInvalidRoutingNumber           V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "invalid_routing_number"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonNotAcceptedPleaseContact       V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "not_accepted_please_contact"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonNoRoutingNumberDetected        V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "no_routing_number_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonSafeHarborPhraseMissing        V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "safe_harbor_phrase_missing"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonSavingsBondDetected            V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "savings_bond_detected"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonTravelersCheck                 V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "travelers_check"
	V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReasonUnableToReadAccountInfo        V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason = "unable_to_read_account_info"
)

// The type of action required.
type V2MoneyManagementFundingIntentNextActionType string

// List of values that V2MoneyManagementFundingIntentNextActionType can take
const (
	V2MoneyManagementFundingIntentNextActionTypeHandleFailure       V2MoneyManagementFundingIntentNextActionType = "handle_failure"
	V2MoneyManagementFundingIntentNextActionTypeProvideNewCheck     V2MoneyManagementFundingIntentNextActionType = "provide_new_check"
	V2MoneyManagementFundingIntentNextActionTypeProvideNewImages    V2MoneyManagementFundingIntentNextActionType = "provide_new_images"
	V2MoneyManagementFundingIntentNextActionTypeRequiresCheckImages V2MoneyManagementFundingIntentNextActionType = "requires_check_images"
)

// A machine-readable reason for the factor.
type V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason string

// List of values that V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason can take
const (
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonAmountVarianceExceedsThreshold V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "amount_variance_exceeds_threshold"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonAmountVarianceNearingThreshold V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "amount_variance_nearing_threshold"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonBackImageNotProcessed          V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "back_image_not_processed"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonBackLooksLikeFront             V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "back_looks_like_front"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonBackSignatureNotDetected       V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "back_signature_not_detected"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonBadPayeeName                   V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "bad_payee_name"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonDateInvalid                    V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "date_invalid"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonDateInFuture                   V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "date_in_future"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonDateOlderThanOneYear           V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "date_older_than_one_year"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonDateOlderThanSixMonths         V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "date_older_than_six_months"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonDuplicateCheck                 V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "duplicate_check"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonEndorsementMissing             V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "endorsement_missing"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonFoldedOrTornCorners            V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "folded_or_torn_corners"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonFrontImageDimensionMismatch    V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "front_image_dimension_mismatch"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonFrontImageNotProcessed         V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "front_image_not_processed"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonFrontImageTooDark              V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "front_image_too_dark"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonFrontImageTooLight             V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "front_image_too_light"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonFrontImageTooSmall             V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "front_image_too_small"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonFrontSignatureNotDetected      V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "front_signature_not_detected"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonHighRisk                       V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "high_risk"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonImageCannotBeProcessed         V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "image_cannot_be_processed"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonImageOutOfFocus                V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "image_out_of_focus"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonInvalidCheckType               V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "invalid_check_type"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonInvalidRoutingNumber           V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "invalid_routing_number"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonNotAcceptedPleaseContact       V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "not_accepted_please_contact"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonNoRoutingNumberDetected        V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "no_routing_number_detected"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonSafeHarborPhraseMissing        V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "safe_harbor_phrase_missing"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonSavingsBondDetected            V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "savings_bond_detected"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonTravelersCheck                 V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "travelers_check"
	V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReasonUnableToReadAccountInfo        V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason = "unable_to_read_account_info"
)

// The type of confirmation required.
type V2MoneyManagementFundingIntentNextConfirmationType string

// List of values that V2MoneyManagementFundingIntentNextConfirmationType can take
const (
	V2MoneyManagementFundingIntentNextConfirmationTypeRequiresReview V2MoneyManagementFundingIntentNextConfirmationType = "requires_review"
)

// The current status of the FundingIntent.
type V2MoneyManagementFundingIntentStatus string

// List of values that V2MoneyManagementFundingIntentStatus can take
const (
	V2MoneyManagementFundingIntentStatusCanceled             V2MoneyManagementFundingIntentStatus = "canceled"
	V2MoneyManagementFundingIntentStatusProcessing           V2MoneyManagementFundingIntentStatus = "processing"
	V2MoneyManagementFundingIntentStatusRequiresAction       V2MoneyManagementFundingIntentStatus = "requires_action"
	V2MoneyManagementFundingIntentStatusRequiresConfirmation V2MoneyManagementFundingIntentStatus = "requires_confirmation"
	V2MoneyManagementFundingIntentStatusSucceeded            V2MoneyManagementFundingIntentStatus = "succeeded"
	V2MoneyManagementFundingIntentStatusValidating           V2MoneyManagementFundingIntentStatus = "validating"
)

// The reason the FundingIntent was canceled.
type V2MoneyManagementFundingIntentStatusDetailsCanceledReason string

// List of values that V2MoneyManagementFundingIntentStatusDetailsCanceledReason can take
const (
	V2MoneyManagementFundingIntentStatusDetailsCanceledReasonCanceledByUser V2MoneyManagementFundingIntentStatusDetailsCanceledReason = "canceled_by_user"
)

// The type of FundingIntent.
type V2MoneyManagementFundingIntentType string

// List of values that V2MoneyManagementFundingIntentType can take
const (
	V2MoneyManagementFundingIntentTypeCheck V2MoneyManagementFundingIntentType = "check"
)

// Details for a check FundingIntent.
type V2MoneyManagementFundingIntentCheck struct {
	// The ID of a File containing the back image of the check.
	BackImage string `json:"back_image,omitempty"`
	// The ID of a File containing the front image of the check.
	FrontImage string `json:"front_image,omitempty"`
}

// Details about the latest underlying money movement attempt.
type V2MoneyManagementFundingIntentLatestAttempt struct {
	// The ID of the InboundFunding created by the attempt.
	InboundFunding string `json:"inbound_funding"`
	// The type of underlying money movement created by the attempt.
	Type V2MoneyManagementFundingIntentLatestAttemptType `json:"type"`
}

// Details for handling a failed attempt.
type V2MoneyManagementFundingIntentNextActionHandleFailure struct {
	// The reason the attempt failed.
	Reason V2MoneyManagementFundingIntentNextActionHandleFailureReason `json:"reason"`
	// Whether you can retry the FundingIntent.
	Retryable bool `json:"retryable"`
}

// The reason that you must provide a different check.
type V2MoneyManagementFundingIntentNextActionProvideNewCheck struct {
	// A human-readable description of the factor.
	Description string `json:"description"`
	// A machine-readable reason for the factor.
	Reason V2MoneyManagementFundingIntentNextActionProvideNewCheckReason `json:"reason"`
}

// Issues detected in the check amount.
type V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssue struct {
	// A human-readable description of the factor.
	Description string `json:"description"`
	// A machine-readable reason for the factor.
	Reason V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssueReason `json:"reason"`
}

// Issues detected in the back image of the check.
type V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImage struct {
	// A human-readable description of the factor.
	Description string `json:"description"`
	// A machine-readable reason for the factor.
	Reason V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImageReason `json:"reason"`
}

// Issues detected in the front image of the check.
type V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImage struct {
	// A human-readable description of the factor.
	Description string `json:"description"`
	// A machine-readable reason for the factor.
	Reason V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImageReason `json:"reason"`
}

// Issues that require corrected check images or an updated amount.
type V2MoneyManagementFundingIntentNextActionProvideNewImages struct {
	// Issues detected in the check amount.
	AmountIssues []*V2MoneyManagementFundingIntentNextActionProvideNewImagesAmountIssue `json:"amount_issues"`
	// Issues detected in the back image of the check.
	BackImage []*V2MoneyManagementFundingIntentNextActionProvideNewImagesBackImage `json:"back_image"`
	// Issues detected in the front image of the check.
	FrontImage []*V2MoneyManagementFundingIntentNextActionProvideNewImagesFrontImage `json:"front_image"`
}

// The action required to continue processing the FundingIntent.
type V2MoneyManagementFundingIntentNextAction struct {
	// Details for handling a failed attempt.
	HandleFailure *V2MoneyManagementFundingIntentNextActionHandleFailure `json:"handle_failure,omitempty"`
	// The reason that you must provide a different check.
	ProvideNewCheck *V2MoneyManagementFundingIntentNextActionProvideNewCheck `json:"provide_new_check,omitempty"`
	// Issues that require corrected check images or an updated amount.
	ProvideNewImages *V2MoneyManagementFundingIntentNextActionProvideNewImages `json:"provide_new_images,omitempty"`
	// The type of action required.
	Type V2MoneyManagementFundingIntentNextActionType `json:"type"`
}

// Factors that require review.
type V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactor struct {
	// A human-readable description of the factor.
	Description string `json:"description"`
	// A machine-readable reason for the factor.
	Reason V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactorReason `json:"reason"`
}

// Factors to review before confirming the FundingIntent.
type V2MoneyManagementFundingIntentNextConfirmationRequiresReview struct {
	// Factors that require review.
	Factors []*V2MoneyManagementFundingIntentNextConfirmationRequiresReviewFactor `json:"factors"`
}

// The information to review before confirming the FundingIntent.
type V2MoneyManagementFundingIntentNextConfirmation struct {
	// Factors to review before confirming the FundingIntent.
	RequiresReview *V2MoneyManagementFundingIntentNextConfirmationRequiresReview `json:"requires_review,omitempty"`
	// The type of confirmation required.
	Type V2MoneyManagementFundingIntentNextConfirmationType `json:"type"`
}

// Controls when the funds become available after processing.
type V2MoneyManagementFundingIntentSettlementTiming struct {
	// The number of days to delay the availability of funds after processing completes.
	DelayDaysOverride int64 `json:"delay_days_override,omitempty"`
}

// Details present when the FundingIntent is canceled.
type V2MoneyManagementFundingIntentStatusDetailsCanceled struct {
	// The reason the FundingIntent was canceled.
	Reason V2MoneyManagementFundingIntentStatusDetailsCanceledReason `json:"reason"`
}

// Additional details about the current status.
type V2MoneyManagementFundingIntentStatusDetails struct {
	// Details present when the FundingIntent is canceled.
	Canceled *V2MoneyManagementFundingIntentStatusDetailsCanceled `json:"canceled,omitempty"`
}

// Timestamps for the FundingIntent's status transitions.
type V2MoneyManagementFundingIntentStatusTransitions struct {
	// Time at which the FundingIntent entered the `canceled` status.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Time at which the FundingIntent entered the `processing` status.
	ProcessingAt time.Time `json:"processing_at,omitempty"`
	// Time at which the FundingIntent entered the `requires_action` status.
	RequiresActionAt time.Time `json:"requires_action_at,omitempty"`
	// Time at which the FundingIntent entered the `requires_confirmation` status.
	RequiresConfirmationAt time.Time `json:"requires_confirmation_at,omitempty"`
	// Time at which the FundingIntent entered the `succeeded` status.
	SucceededAt time.Time `json:"succeeded_at,omitempty"`
	// Time at which the FundingIntent entered the `validating` status.
	ValidatingAt time.Time `json:"validating_at,omitempty"`
}

// The destination for the funds.
type V2MoneyManagementFundingIntentTo struct {
	// The currency of the destination FinancialAccount balance.
	Currency Currency `json:"currency"`
	// The ID of the FinancialAccount that receives the funds.
	FinancialAccount string `json:"financial_account"`
}

// A FundingIntent represents the lifecycle of adding funds to a FinancialAccount.
type V2MoneyManagementFundingIntent struct {
	APIResource
	// The amount to add to the destination.
	Amount Amount `json:"amount"`
	// Whether you can cancel the FundingIntent.
	Cancelable bool `json:"cancelable"`
	// Details for a check FundingIntent.
	Check *V2MoneyManagementFundingIntentCheck `json:"check,omitempty"`
	// How the FundingIntent is confirmed after validation.
	ConfirmationMethod V2MoneyManagementFundingIntentConfirmationMethod `json:"confirmation_method"`
	// Time at which the FundingIntent was created.
	Created time.Time `json:"created"`
	// An arbitrary string attached to the FundingIntent.
	Description string `json:"description,omitempty"`
	// Unique identifier for the FundingIntent.
	ID string `json:"id"`
	// Details about the latest underlying money movement attempt.
	LatestAttempt *V2MoneyManagementFundingIntentLatestAttempt `json:"latest_attempt,omitempty"`
	// Whether the FundingIntent exists in live mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs attached to the FundingIntent.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The action required to continue processing the FundingIntent.
	NextAction *V2MoneyManagementFundingIntentNextAction `json:"next_action,omitempty"`
	// The information to review before confirming the FundingIntent.
	NextConfirmation *V2MoneyManagementFundingIntentNextConfirmation `json:"next_confirmation,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Controls when the funds become available after processing.
	SettlementTiming *V2MoneyManagementFundingIntentSettlementTiming `json:"settlement_timing,omitempty"`
	// The current status of the FundingIntent.
	Status V2MoneyManagementFundingIntentStatus `json:"status"`
	// Additional details about the current status.
	StatusDetails *V2MoneyManagementFundingIntentStatusDetails `json:"status_details,omitempty"`
	// Timestamps for the FundingIntent's status transitions.
	StatusTransitions *V2MoneyManagementFundingIntentStatusTransitions `json:"status_transitions"`
	// The destination for the funds.
	To *V2MoneyManagementFundingIntentTo `json:"to"`
	// The type of FundingIntent.
	Type V2MoneyManagementFundingIntentType `json:"type"`
}
