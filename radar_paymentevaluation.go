//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Reason given by cardholder for dispute.
type RadarPaymentEvaluationEventDisputeOpenedReason string

// List of values that RadarPaymentEvaluationEventDisputeOpenedReason can take
const (
	RadarPaymentEvaluationEventDisputeOpenedReasonAccountNotAvailable  RadarPaymentEvaluationEventDisputeOpenedReason = "account_not_available"
	RadarPaymentEvaluationEventDisputeOpenedReasonCreditNotProcessed   RadarPaymentEvaluationEventDisputeOpenedReason = "credit_not_processed"
	RadarPaymentEvaluationEventDisputeOpenedReasonCustomerInitiated    RadarPaymentEvaluationEventDisputeOpenedReason = "customer_initiated"
	RadarPaymentEvaluationEventDisputeOpenedReasonDuplicate            RadarPaymentEvaluationEventDisputeOpenedReason = "duplicate"
	RadarPaymentEvaluationEventDisputeOpenedReasonFraudulent           RadarPaymentEvaluationEventDisputeOpenedReason = "fraudulent"
	RadarPaymentEvaluationEventDisputeOpenedReasonGeneral              RadarPaymentEvaluationEventDisputeOpenedReason = "general"
	RadarPaymentEvaluationEventDisputeOpenedReasonNoncompliant         RadarPaymentEvaluationEventDisputeOpenedReason = "noncompliant"
	RadarPaymentEvaluationEventDisputeOpenedReasonProductNotReceived   RadarPaymentEvaluationEventDisputeOpenedReason = "product_not_received"
	RadarPaymentEvaluationEventDisputeOpenedReasonProductUnacceptable  RadarPaymentEvaluationEventDisputeOpenedReason = "product_unacceptable"
	RadarPaymentEvaluationEventDisputeOpenedReasonSubscriptionCanceled RadarPaymentEvaluationEventDisputeOpenedReason = "subscription_canceled"
	RadarPaymentEvaluationEventDisputeOpenedReasonUnrecognized         RadarPaymentEvaluationEventDisputeOpenedReason = "unrecognized"
)

// The type of fraud labeled by the issuer.
type RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudType string

// List of values that RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudType can take
const (
	RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudTypeMadeWithLostCard      RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudType = "made_with_lost_card"
	RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudTypeMadeWithStolenCard    RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudType = "made_with_stolen_card"
	RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudTypeOther                 RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudType = "other"
	RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudTypeUnauthorizedUseOfCard RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudType = "unauthorized_use_of_card"
)

// Indicates the reason for the refund.
type RadarPaymentEvaluationEventRefundedReason string

// List of values that RadarPaymentEvaluationEventRefundedReason can take
const (
	RadarPaymentEvaluationEventRefundedReasonDuplicate           RadarPaymentEvaluationEventRefundedReason = "duplicate"
	RadarPaymentEvaluationEventRefundedReasonFraudulent          RadarPaymentEvaluationEventRefundedReason = "fraudulent"
	RadarPaymentEvaluationEventRefundedReasonOther               RadarPaymentEvaluationEventRefundedReason = "other"
	RadarPaymentEvaluationEventRefundedReasonRequestedByCustomer RadarPaymentEvaluationEventRefundedReason = "requested_by_customer"
)

// Indicates the type of event attached to the payment evaluation.
type RadarPaymentEvaluationEventType string

// List of values that RadarPaymentEvaluationEventType can take
const (
	RadarPaymentEvaluationEventTypeDisputeOpened             RadarPaymentEvaluationEventType = "dispute_opened"
	RadarPaymentEvaluationEventTypeEarlyFraudWarningReceived RadarPaymentEvaluationEventType = "early_fraud_warning_received"
	RadarPaymentEvaluationEventTypeRefunded                  RadarPaymentEvaluationEventType = "refunded"
	RadarPaymentEvaluationEventTypeUserInterventionRaised    RadarPaymentEvaluationEventType = "user_intervention_raised"
	RadarPaymentEvaluationEventTypeUserInterventionResolved  RadarPaymentEvaluationEventType = "user_intervention_resolved"
)

// Type of user intervention raised.
type RadarPaymentEvaluationEventUserInterventionRaisedType string

// List of values that RadarPaymentEvaluationEventUserInterventionRaisedType can take
const (
	RadarPaymentEvaluationEventUserInterventionRaisedType3ds     RadarPaymentEvaluationEventUserInterventionRaisedType = "3ds"
	RadarPaymentEvaluationEventUserInterventionRaisedTypeCaptcha RadarPaymentEvaluationEventUserInterventionRaisedType = "captcha"
	RadarPaymentEvaluationEventUserInterventionRaisedTypeCustom  RadarPaymentEvaluationEventUserInterventionRaisedType = "custom"
)

// Result of the intervention if it has been completed.
type RadarPaymentEvaluationEventUserInterventionResolvedOutcome string

// List of values that RadarPaymentEvaluationEventUserInterventionResolvedOutcome can take
const (
	RadarPaymentEvaluationEventUserInterventionResolvedOutcomeAbandoned RadarPaymentEvaluationEventUserInterventionResolvedOutcome = "abandoned"
	RadarPaymentEvaluationEventUserInterventionResolvedOutcomeFailed    RadarPaymentEvaluationEventUserInterventionResolvedOutcome = "failed"
	RadarPaymentEvaluationEventUserInterventionResolvedOutcomePassed    RadarPaymentEvaluationEventUserInterventionResolvedOutcome = "passed"
)

// Recommended action based on the risk score. Possible values are `block` and `continue`.
type RadarPaymentEvaluationInsightsFraudulentDisputeRecommendedAction string

// List of values that RadarPaymentEvaluationInsightsFraudulentDisputeRecommendedAction can take
const (
	RadarPaymentEvaluationInsightsFraudulentDisputeRecommendedActionBlock    RadarPaymentEvaluationInsightsFraudulentDisputeRecommendedAction = "block"
	RadarPaymentEvaluationInsightsFraudulentDisputeRecommendedActionContinue RadarPaymentEvaluationInsightsFraudulentDisputeRecommendedAction = "continue"
)

// The reason the payment was blocked by the merchant.
type RadarPaymentEvaluationOutcomeMerchantBlockedReason string

// List of values that RadarPaymentEvaluationOutcomeMerchantBlockedReason can take
const (
	RadarPaymentEvaluationOutcomeMerchantBlockedReasonAuthenticationRequired RadarPaymentEvaluationOutcomeMerchantBlockedReason = "authentication_required"
	RadarPaymentEvaluationOutcomeMerchantBlockedReasonBlockedForFraud        RadarPaymentEvaluationOutcomeMerchantBlockedReason = "blocked_for_fraud"
	RadarPaymentEvaluationOutcomeMerchantBlockedReasonInvalidPayment         RadarPaymentEvaluationOutcomeMerchantBlockedReason = "invalid_payment"
	RadarPaymentEvaluationOutcomeMerchantBlockedReasonOther                  RadarPaymentEvaluationOutcomeMerchantBlockedReason = "other"
)

// Result of the address line 1 check.
type RadarPaymentEvaluationOutcomeRejectedCardAddressLine1Check string

// List of values that RadarPaymentEvaluationOutcomeRejectedCardAddressLine1Check can take
const (
	RadarPaymentEvaluationOutcomeRejectedCardAddressLine1CheckFail        RadarPaymentEvaluationOutcomeRejectedCardAddressLine1Check = "fail"
	RadarPaymentEvaluationOutcomeRejectedCardAddressLine1CheckPass        RadarPaymentEvaluationOutcomeRejectedCardAddressLine1Check = "pass"
	RadarPaymentEvaluationOutcomeRejectedCardAddressLine1CheckUnavailable RadarPaymentEvaluationOutcomeRejectedCardAddressLine1Check = "unavailable"
	RadarPaymentEvaluationOutcomeRejectedCardAddressLine1CheckUnchecked   RadarPaymentEvaluationOutcomeRejectedCardAddressLine1Check = "unchecked"
)

// Indicates whether the cardholder provided a postal code and if it matched the cardholder's billing address.
type RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheck string

// List of values that RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheck can take
const (
	RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheckFail        RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheck = "fail"
	RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheckPass        RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheck = "pass"
	RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheckUnavailable RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheck = "unavailable"
	RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheckUnchecked   RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheck = "unchecked"
)

// Result of the CVC check.
type RadarPaymentEvaluationOutcomeRejectedCardCVCCheck string

// List of values that RadarPaymentEvaluationOutcomeRejectedCardCVCCheck can take
const (
	RadarPaymentEvaluationOutcomeRejectedCardCVCCheckFail        RadarPaymentEvaluationOutcomeRejectedCardCVCCheck = "fail"
	RadarPaymentEvaluationOutcomeRejectedCardCVCCheckPass        RadarPaymentEvaluationOutcomeRejectedCardCVCCheck = "pass"
	RadarPaymentEvaluationOutcomeRejectedCardCVCCheckUnavailable RadarPaymentEvaluationOutcomeRejectedCardCVCCheck = "unavailable"
	RadarPaymentEvaluationOutcomeRejectedCardCVCCheckUnchecked   RadarPaymentEvaluationOutcomeRejectedCardCVCCheck = "unchecked"
)

// Card issuer's reason for the network decline.
type RadarPaymentEvaluationOutcomeRejectedCardReason string

// List of values that RadarPaymentEvaluationOutcomeRejectedCardReason can take
const (
	RadarPaymentEvaluationOutcomeRejectedCardReasonAuthenticationFailed RadarPaymentEvaluationOutcomeRejectedCardReason = "authentication_failed"
	RadarPaymentEvaluationOutcomeRejectedCardReasonDoNotHonor           RadarPaymentEvaluationOutcomeRejectedCardReason = "do_not_honor"
	RadarPaymentEvaluationOutcomeRejectedCardReasonExpired              RadarPaymentEvaluationOutcomeRejectedCardReason = "expired"
	RadarPaymentEvaluationOutcomeRejectedCardReasonIncorrectCVC         RadarPaymentEvaluationOutcomeRejectedCardReason = "incorrect_cvc"
	RadarPaymentEvaluationOutcomeRejectedCardReasonIncorrectNumber      RadarPaymentEvaluationOutcomeRejectedCardReason = "incorrect_number"
	RadarPaymentEvaluationOutcomeRejectedCardReasonIncorrectPostalCode  RadarPaymentEvaluationOutcomeRejectedCardReason = "incorrect_postal_code"
	RadarPaymentEvaluationOutcomeRejectedCardReasonInsufficientFunds    RadarPaymentEvaluationOutcomeRejectedCardReason = "insufficient_funds"
	RadarPaymentEvaluationOutcomeRejectedCardReasonInvalidAccount       RadarPaymentEvaluationOutcomeRejectedCardReason = "invalid_account"
	RadarPaymentEvaluationOutcomeRejectedCardReasonLostCard             RadarPaymentEvaluationOutcomeRejectedCardReason = "lost_card"
	RadarPaymentEvaluationOutcomeRejectedCardReasonOther                RadarPaymentEvaluationOutcomeRejectedCardReason = "other"
	RadarPaymentEvaluationOutcomeRejectedCardReasonProcessingError      RadarPaymentEvaluationOutcomeRejectedCardReason = "processing_error"
	RadarPaymentEvaluationOutcomeRejectedCardReasonReportedStolen       RadarPaymentEvaluationOutcomeRejectedCardReason = "reported_stolen"
	RadarPaymentEvaluationOutcomeRejectedCardReasonTryAgainLater        RadarPaymentEvaluationOutcomeRejectedCardReason = "try_again_later"
)

// Result of the address line 1 check.
type RadarPaymentEvaluationOutcomeSucceededCardAddressLine1Check string

// List of values that RadarPaymentEvaluationOutcomeSucceededCardAddressLine1Check can take
const (
	RadarPaymentEvaluationOutcomeSucceededCardAddressLine1CheckFail        RadarPaymentEvaluationOutcomeSucceededCardAddressLine1Check = "fail"
	RadarPaymentEvaluationOutcomeSucceededCardAddressLine1CheckPass        RadarPaymentEvaluationOutcomeSucceededCardAddressLine1Check = "pass"
	RadarPaymentEvaluationOutcomeSucceededCardAddressLine1CheckUnavailable RadarPaymentEvaluationOutcomeSucceededCardAddressLine1Check = "unavailable"
	RadarPaymentEvaluationOutcomeSucceededCardAddressLine1CheckUnchecked   RadarPaymentEvaluationOutcomeSucceededCardAddressLine1Check = "unchecked"
)

// Indicates whether the cardholder provided a postal code and if it matched the cardholder's billing address.
type RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheck string

// List of values that RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheck can take
const (
	RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheckFail        RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheck = "fail"
	RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheckPass        RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheck = "pass"
	RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheckUnavailable RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheck = "unavailable"
	RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheckUnchecked   RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheck = "unchecked"
)

// Result of the CVC check.
type RadarPaymentEvaluationOutcomeSucceededCardCVCCheck string

// List of values that RadarPaymentEvaluationOutcomeSucceededCardCVCCheck can take
const (
	RadarPaymentEvaluationOutcomeSucceededCardCVCCheckFail        RadarPaymentEvaluationOutcomeSucceededCardCVCCheck = "fail"
	RadarPaymentEvaluationOutcomeSucceededCardCVCCheckPass        RadarPaymentEvaluationOutcomeSucceededCardCVCCheck = "pass"
	RadarPaymentEvaluationOutcomeSucceededCardCVCCheckUnavailable RadarPaymentEvaluationOutcomeSucceededCardCVCCheck = "unavailable"
	RadarPaymentEvaluationOutcomeSucceededCardCVCCheckUnchecked   RadarPaymentEvaluationOutcomeSucceededCardCVCCheck = "unchecked"
)

// Indicates the outcome of the payment evaluation.
type RadarPaymentEvaluationOutcomeType string

// List of values that RadarPaymentEvaluationOutcomeType can take
const (
	RadarPaymentEvaluationOutcomeTypeFailed          RadarPaymentEvaluationOutcomeType = "failed"
	RadarPaymentEvaluationOutcomeTypeMerchantBlocked RadarPaymentEvaluationOutcomeType = "merchant_blocked"
	RadarPaymentEvaluationOutcomeTypeRejected        RadarPaymentEvaluationOutcomeType = "rejected"
	RadarPaymentEvaluationOutcomeTypeSucceeded       RadarPaymentEvaluationOutcomeType = "succeeded"
)

// Describes the presence of the customer during the payment.
type RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresence string

// List of values that RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresence can take
const (
	RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresenceOffSession RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresence = "off_session"
	RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresenceOnSession  RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresence = "on_session"
)

// Describes the type of payment.
type RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType string

// List of values that RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType can take
const (
	RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentTypeOneOff         RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType = "one_off"
	RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentTypeRecurring      RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType = "recurring"
	RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentTypeSetupOneOff    RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType = "setup_one_off"
	RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentTypeSetupRecurring RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType = "setup_recurring"
)

// Describes the type of money movement. Currently only `card` is supported.
type RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsMoneyMovementType string

// List of values that RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsMoneyMovementType can take
const (
	RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsMoneyMovementTypeCard RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsMoneyMovementType = "card"
)

// Details about the Client Device Metadata to associate with the payment evaluation.
type RadarPaymentEvaluationClientDeviceMetadataDetailsParams struct {
	// ID for the Radar Session to associate with the payment evaluation. A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	RadarSession *string `form:"radar_session"`
}

// Details about the customer associated with the payment evaluation.
type RadarPaymentEvaluationCustomerDetailsParams struct {
	// The ID of the customer associated with the payment evaluation.
	Customer *string `form:"customer"`
	// The ID of the Account representing the customer associated with the payment evaluation.
	CustomerAccount *string `form:"customer_account"`
	// The customer's email address.
	Email *string `form:"email"`
	// The customer's full name or business name.
	Name *string `form:"name"`
	// The customer's phone number.
	Phone *string `form:"phone"`
}

// Describes card money movement details for the payment evaluation.
type RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardParams struct {
	// Describes the presence of the customer during the payment.
	CustomerPresence *string `form:"customer_presence"`
	// Describes the type of payment.
	PaymentType *string `form:"payment_type"`
}

// Details about the payment's customer presence and type.
type RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsParams struct {
	// Describes card money movement details for the payment evaluation.
	Card *RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardParams `form:"card"`
	// Describes the type of money movement. Currently only `card` is supported.
	MoneyMovementType *string `form:"money_movement_type"`
}

// Billing information associated with the payment evaluation.
type RadarPaymentEvaluationPaymentDetailsPaymentMethodDetailsBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address"`
	// Email address.
	Email *string `form:"email"`
	// Full name.
	Name *string `form:"name"`
	// Billing phone number (including extension).
	Phone *string `form:"phone"`
}

// Details about the payment method to use for the payment.
type RadarPaymentEvaluationPaymentDetailsPaymentMethodDetailsParams struct {
	// Billing information associated with the payment evaluation.
	BillingDetails *RadarPaymentEvaluationPaymentDetailsPaymentMethodDetailsBillingDetailsParams `form:"billing_details"`
	// ID of the payment method used in this payment evaluation.
	PaymentMethod *string `form:"payment_method"`
}

// Shipping details for the payment evaluation.
type RadarPaymentEvaluationPaymentDetailsShippingDetailsParams struct {
	// Shipping address.
	Address *AddressParams `form:"address"`
	// Shipping name.
	Name *string `form:"name"`
	// Shipping phone number.
	Phone *string `form:"phone"`
}

// Details about the payment.
type RadarPaymentEvaluationPaymentDetailsParams struct {
	// The intended amount to collect with this payment. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (for example, 100 cents to charge 1.00 USD or 100 to charge 100 Yen, a zero-decimal currency).
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Details about the payment's customer presence and type.
	MoneyMovementDetails *RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsParams `form:"money_movement_details"`
	// Details about the payment method to use for the payment.
	PaymentMethodDetails *RadarPaymentEvaluationPaymentDetailsPaymentMethodDetailsParams `form:"payment_method_details"`
	// Shipping details for the payment evaluation.
	ShippingDetails *RadarPaymentEvaluationPaymentDetailsShippingDetailsParams `form:"shipping_details"`
	// Payment statement descriptor.
	StatementDescriptor *string `form:"statement_descriptor"`
}

// Request a Radar API fraud risk score from Stripe for a payment before sending it for external processor authorization.
type RadarPaymentEvaluationParams struct {
	Params `form:"*"`
	// Details about the Client Device Metadata to associate with the payment evaluation.
	ClientDeviceMetadataDetails *RadarPaymentEvaluationClientDeviceMetadataDetailsParams `form:"client_device_metadata_details"`
	// Details about the customer associated with the payment evaluation.
	CustomerDetails *RadarPaymentEvaluationCustomerDetailsParams `form:"customer_details"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Details about the payment.
	PaymentDetails *RadarPaymentEvaluationPaymentDetailsParams `form:"payment_details"`
}

// AddExpand appends a new field to expand.
func (p *RadarPaymentEvaluationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *RadarPaymentEvaluationParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details about the Client Device Metadata to associate with the payment evaluation.
type RadarPaymentEvaluationCreateClientDeviceMetadataDetailsParams struct {
	// ID for the Radar Session to associate with the payment evaluation. A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	RadarSession *string `form:"radar_session"`
}

// Details about the customer associated with the payment evaluation.
type RadarPaymentEvaluationCreateCustomerDetailsParams struct {
	// The ID of the customer associated with the payment evaluation.
	Customer *string `form:"customer"`
	// The ID of the Account representing the customer associated with the payment evaluation.
	CustomerAccount *string `form:"customer_account"`
	// The customer's email address.
	Email *string `form:"email"`
	// The customer's full name or business name.
	Name *string `form:"name"`
	// The customer's phone number.
	Phone *string `form:"phone"`
}

// Describes card money movement details for the payment evaluation.
type RadarPaymentEvaluationCreatePaymentDetailsMoneyMovementDetailsCardParams struct {
	// Describes the presence of the customer during the payment.
	CustomerPresence *string `form:"customer_presence"`
	// Describes the type of payment.
	PaymentType *string `form:"payment_type"`
}

// Details about the payment's customer presence and type.
type RadarPaymentEvaluationCreatePaymentDetailsMoneyMovementDetailsParams struct {
	// Describes card money movement details for the payment evaluation.
	Card *RadarPaymentEvaluationCreatePaymentDetailsMoneyMovementDetailsCardParams `form:"card"`
	// Describes the type of money movement. Currently only `card` is supported.
	MoneyMovementType *string `form:"money_movement_type"`
}

// Billing information associated with the payment evaluation.
type RadarPaymentEvaluationCreatePaymentDetailsPaymentMethodDetailsBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address"`
	// Email address.
	Email *string `form:"email"`
	// Full name.
	Name *string `form:"name"`
	// Billing phone number (including extension).
	Phone *string `form:"phone"`
}

// Details about the payment method to use for the payment.
type RadarPaymentEvaluationCreatePaymentDetailsPaymentMethodDetailsParams struct {
	// Billing information associated with the payment evaluation.
	BillingDetails *RadarPaymentEvaluationCreatePaymentDetailsPaymentMethodDetailsBillingDetailsParams `form:"billing_details"`
	// ID of the payment method used in this payment evaluation.
	PaymentMethod *string `form:"payment_method"`
}

// Shipping details for the payment evaluation.
type RadarPaymentEvaluationCreatePaymentDetailsShippingDetailsParams struct {
	// Shipping address.
	Address *AddressParams `form:"address"`
	// Shipping name.
	Name *string `form:"name"`
	// Shipping phone number.
	Phone *string `form:"phone"`
}

// Details about the payment.
type RadarPaymentEvaluationCreatePaymentDetailsParams struct {
	// The intended amount to collect with this payment. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (for example, 100 cents to charge 1.00 USD or 100 to charge 100 Yen, a zero-decimal currency).
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Details about the payment's customer presence and type.
	MoneyMovementDetails *RadarPaymentEvaluationCreatePaymentDetailsMoneyMovementDetailsParams `form:"money_movement_details"`
	// Details about the payment method to use for the payment.
	PaymentMethodDetails *RadarPaymentEvaluationCreatePaymentDetailsPaymentMethodDetailsParams `form:"payment_method_details"`
	// Shipping details for the payment evaluation.
	ShippingDetails *RadarPaymentEvaluationCreatePaymentDetailsShippingDetailsParams `form:"shipping_details"`
	// Payment statement descriptor.
	StatementDescriptor *string `form:"statement_descriptor"`
}

// Request a Radar API fraud risk score from Stripe for a payment before sending it for external processor authorization.
type RadarPaymentEvaluationCreateParams struct {
	Params `form:"*"`
	// Details about the Client Device Metadata to associate with the payment evaluation.
	ClientDeviceMetadataDetails *RadarPaymentEvaluationCreateClientDeviceMetadataDetailsParams `form:"client_device_metadata_details"`
	// Details about the customer associated with the payment evaluation.
	CustomerDetails *RadarPaymentEvaluationCreateCustomerDetailsParams `form:"customer_details"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Details about the payment.
	PaymentDetails *RadarPaymentEvaluationCreatePaymentDetailsParams `form:"payment_details"`
}

// AddExpand appends a new field to expand.
func (p *RadarPaymentEvaluationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *RadarPaymentEvaluationCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Client device metadata attached to this payment evaluation.
type RadarPaymentEvaluationClientDeviceMetadataDetails struct {
	// ID for the Radar Session associated with the payment evaluation. A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	RadarSession string `json:"radar_session"`
}

// Customer details attached to this payment evaluation.
type RadarPaymentEvaluationCustomerDetails struct {
	// The ID of the customer associated with the payment evaluation.
	Customer string `json:"customer"`
	// The ID of the Account representing the customer associated with the payment evaluation.
	CustomerAccount string `json:"customer_account"`
	// The customer's email address.
	Email string `json:"email"`
	// The customer's full name or business name.
	Name string `json:"name"`
	// The customer's phone number.
	Phone string `json:"phone"`
}

// Dispute opened event details attached to this payment evaluation.
type RadarPaymentEvaluationEventDisputeOpened struct {
	// Amount to dispute for this payment. A positive integer representing how much to charge in [the smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (for example, 100 cents to charge 1.00 USD or 100 to charge 100 Yen, a zero-decimal currency).
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Reason given by cardholder for dispute.
	Reason RadarPaymentEvaluationEventDisputeOpenedReason `json:"reason"`
}

// Early Fraud Warning Received event details attached to this payment evaluation.
type RadarPaymentEvaluationEventEarlyFraudWarningReceived struct {
	// The type of fraud labeled by the issuer.
	FraudType RadarPaymentEvaluationEventEarlyFraudWarningReceivedFraudType `json:"fraud_type"`
}

// Refunded Event details attached to this payment evaluation.
type RadarPaymentEvaluationEventRefunded struct {
	// Amount refunded for this payment. A positive integer representing how much to charge in [the smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (for example, 100 cents to charge 1.00 USD or 100 to charge 100 Yen, a zero-decimal currency).
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Indicates the reason for the refund.
	Reason RadarPaymentEvaluationEventRefundedReason `json:"reason"`
}

// User intervention raised custom event details attached to this payment evaluation
type RadarPaymentEvaluationEventUserInterventionRaisedCustom struct {
	// Custom type of user intervention raised. The string must use a snake case description for the type of intervention performed.
	Type string `json:"type"`
}

// User intervention raised event details attached to this payment evaluation
type RadarPaymentEvaluationEventUserInterventionRaised struct {
	// User intervention raised custom event details attached to this payment evaluation
	Custom *RadarPaymentEvaluationEventUserInterventionRaisedCustom `json:"custom"`
	// Unique identifier for the user intervention event.
	Key string `json:"key"`
	// Type of user intervention raised.
	Type RadarPaymentEvaluationEventUserInterventionRaisedType `json:"type"`
}

// User Intervention Resolved Event details attached to this payment evaluation
type RadarPaymentEvaluationEventUserInterventionResolved struct {
	// Unique ID of this intervention. Use this to provide the result.
	Key string `json:"key"`
	// Result of the intervention if it has been completed.
	Outcome RadarPaymentEvaluationEventUserInterventionResolvedOutcome `json:"outcome"`
}

// Event information associated with the payment evaluation, such as refunds, dispute, early fraud warnings, or user interventions.
type RadarPaymentEvaluationEvent struct {
	// Dispute opened event details attached to this payment evaluation.
	DisputeOpened *RadarPaymentEvaluationEventDisputeOpened `json:"dispute_opened"`
	// Early Fraud Warning Received event details attached to this payment evaluation.
	EarlyFraudWarningReceived *RadarPaymentEvaluationEventEarlyFraudWarningReceived `json:"early_fraud_warning_received"`
	// Timestamp when the event occurred.
	OccurredAt int64 `json:"occurred_at"`
	// Refunded Event details attached to this payment evaluation.
	Refunded *RadarPaymentEvaluationEventRefunded `json:"refunded"`
	// Indicates the type of event attached to the payment evaluation.
	Type RadarPaymentEvaluationEventType `json:"type"`
	// User intervention raised event details attached to this payment evaluation
	UserInterventionRaised *RadarPaymentEvaluationEventUserInterventionRaised `json:"user_intervention_raised"`
	// User Intervention Resolved Event details attached to this payment evaluation
	UserInterventionResolved *RadarPaymentEvaluationEventUserInterventionResolved `json:"user_intervention_resolved"`
}

// Scores, insights and recommended action for one scorer for this PaymentEvaluation.
type RadarPaymentEvaluationInsightsFraudulentDispute struct {
	// Recommended action based on the risk score. Possible values are `block` and `continue`.
	RecommendedAction RadarPaymentEvaluationInsightsFraudulentDisputeRecommendedAction `json:"recommended_action"`
	// Stripe Radar's evaluation of the risk level of the payment. Possible values for evaluated payments are between 0 and 100, with higher scores indicating higher risk.
	RiskScore int64 `json:"risk_score"`
}

// Collection of scores and insights for this payment evaluation.
type RadarPaymentEvaluationInsights struct {
	// The timestamp when the evaluation was performed.
	EvaluatedAt int64 `json:"evaluated_at"`
	// Scores, insights and recommended action for one scorer for this PaymentEvaluation.
	FraudulentDispute *RadarPaymentEvaluationInsightsFraudulentDispute `json:"fraudulent_dispute"`
}

// Details of a merchant_blocked outcome attached to this payment evaluation.
type RadarPaymentEvaluationOutcomeMerchantBlocked struct {
	// The reason the payment was blocked by the merchant.
	Reason RadarPaymentEvaluationOutcomeMerchantBlockedReason `json:"reason"`
}

// Details of an rejected card outcome attached to this payment evaluation.
type RadarPaymentEvaluationOutcomeRejectedCard struct {
	// Result of the address line 1 check.
	AddressLine1Check RadarPaymentEvaluationOutcomeRejectedCardAddressLine1Check `json:"address_line1_check"`
	// Indicates whether the cardholder provided a postal code and if it matched the cardholder's billing address.
	AddressPostalCodeCheck RadarPaymentEvaluationOutcomeRejectedCardAddressPostalCodeCheck `json:"address_postal_code_check"`
	// Result of the CVC check.
	CVCCheck RadarPaymentEvaluationOutcomeRejectedCardCVCCheck `json:"cvc_check"`
	// Card issuer's reason for the network decline.
	Reason RadarPaymentEvaluationOutcomeRejectedCardReason `json:"reason"`
}

// Details of an rejected outcome attached to this payment evaluation.
type RadarPaymentEvaluationOutcomeRejected struct {
	// Details of an rejected card outcome attached to this payment evaluation.
	Card *RadarPaymentEvaluationOutcomeRejectedCard `json:"card"`
}

// Details of an succeeded card outcome attached to this payment evaluation.
type RadarPaymentEvaluationOutcomeSucceededCard struct {
	// Result of the address line 1 check.
	AddressLine1Check RadarPaymentEvaluationOutcomeSucceededCardAddressLine1Check `json:"address_line1_check"`
	// Indicates whether the cardholder provided a postal code and if it matched the cardholder's billing address.
	AddressPostalCodeCheck RadarPaymentEvaluationOutcomeSucceededCardAddressPostalCodeCheck `json:"address_postal_code_check"`
	// Result of the CVC check.
	CVCCheck RadarPaymentEvaluationOutcomeSucceededCardCVCCheck `json:"cvc_check"`
}

// Details of a succeeded outcome attached to this payment evaluation.
type RadarPaymentEvaluationOutcomeSucceeded struct {
	// Details of an succeeded card outcome attached to this payment evaluation.
	Card *RadarPaymentEvaluationOutcomeSucceededCard `json:"card"`
}

// Indicates the final outcome for the payment evaluation.
type RadarPaymentEvaluationOutcome struct {
	// Details of a merchant_blocked outcome attached to this payment evaluation.
	MerchantBlocked *RadarPaymentEvaluationOutcomeMerchantBlocked `json:"merchant_blocked"`
	// The PaymentIntent ID associated with the payment evaluation.
	PaymentIntentID string `json:"payment_intent_id"`
	// Details of an rejected outcome attached to this payment evaluation.
	Rejected *RadarPaymentEvaluationOutcomeRejected `json:"rejected"`
	// Details of a succeeded outcome attached to this payment evaluation.
	Succeeded *RadarPaymentEvaluationOutcomeSucceeded `json:"succeeded"`
	// Indicates the outcome of the payment evaluation.
	Type RadarPaymentEvaluationOutcomeType `json:"type"`
}

// Describes card money movement details for the payment evaluation.
type RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCard struct {
	// Describes the presence of the customer during the payment.
	CustomerPresence RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresence `json:"customer_presence"`
	// Describes the type of payment.
	PaymentType RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType `json:"payment_type"`
}

// Details about the payment's customer presence and type.
type RadarPaymentEvaluationPaymentDetailsMoneyMovementDetails struct {
	// Describes card money movement details for the payment evaluation.
	Card *RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsCard `json:"card"`
	// Describes the type of money movement. Currently only `card` is supported.
	MoneyMovementType RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsMoneyMovementType `json:"money_movement_type"`
}

// Billing information associated with the payment evaluation.
type RadarPaymentEvaluationPaymentDetailsPaymentMethodDetailsBillingDetails struct {
	// Address data.
	Address *Address `json:"address"`
	// Email address.
	Email string `json:"email"`
	// Full name.
	Name string `json:"name"`
	// Billing phone number (including extension).
	Phone string `json:"phone"`
}

// Details about the payment method used for the payment.
type RadarPaymentEvaluationPaymentDetailsPaymentMethodDetails struct {
	// Billing information associated with the payment evaluation.
	BillingDetails *RadarPaymentEvaluationPaymentDetailsPaymentMethodDetailsBillingDetails `json:"billing_details"`
	// The payment method used in this payment evaluation.
	PaymentMethod *PaymentMethod `json:"payment_method"`
}

// Shipping details for the payment evaluation.
type RadarPaymentEvaluationPaymentDetailsShippingDetails struct {
	// Address data.
	Address *Address `json:"address"`
	// Shipping name.
	Name string `json:"name"`
	// Shipping phone number.
	Phone string `json:"phone"`
}

// Payment details attached to this payment evaluation.
type RadarPaymentEvaluationPaymentDetails struct {
	// Amount intended to be collected by this payment. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://docs.stripe.com/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Details about the payment's customer presence and type.
	MoneyMovementDetails *RadarPaymentEvaluationPaymentDetailsMoneyMovementDetails `json:"money_movement_details"`
	// Details about the payment method used for the payment.
	PaymentMethodDetails *RadarPaymentEvaluationPaymentDetailsPaymentMethodDetails `json:"payment_method_details"`
	// Shipping details for the payment evaluation.
	ShippingDetails *RadarPaymentEvaluationPaymentDetailsShippingDetails `json:"shipping_details"`
	// Payment statement descriptor.
	StatementDescriptor string `json:"statement_descriptor"`
}

// Payment Evaluations represent the risk lifecycle of an externally processed payment. It includes the Radar risk score from Stripe, payment outcome taken by the merchant or processor, and any post transaction events, such as refunds or disputes. See the [Radar API guide](https://docs.stripe.com/radar/multiprocessor) for integration steps.
type RadarPaymentEvaluation struct {
	APIResource
	// Client device metadata attached to this payment evaluation.
	ClientDeviceMetadataDetails *RadarPaymentEvaluationClientDeviceMetadataDetails `json:"client_device_metadata_details"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	CreatedAt int64 `json:"created_at"`
	// Customer details attached to this payment evaluation.
	CustomerDetails *RadarPaymentEvaluationCustomerDetails `json:"customer_details"`
	// Event information associated with the payment evaluation, such as refunds, dispute, early fraud warnings, or user interventions.
	Events []*RadarPaymentEvaluationEvent `json:"events"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Collection of scores and insights for this payment evaluation.
	Insights *RadarPaymentEvaluationInsights `json:"insights"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Indicates the final outcome for the payment evaluation.
	Outcome *RadarPaymentEvaluationOutcome `json:"outcome"`
	// Payment details attached to this payment evaluation.
	PaymentDetails *RadarPaymentEvaluationPaymentDetails `json:"payment_details"`
}
