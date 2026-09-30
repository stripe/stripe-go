//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Type of the 3RI Authentication.
type ThreeDSecureAuthenticationChannelThreeRIType string

// List of values that ThreeDSecureAuthenticationChannelThreeRIType can take
const (
	ThreeDSecureAuthenticationChannelThreeRITypeDelayedShipment ThreeDSecureAuthenticationChannelThreeRIType = "delayed_shipment"
	ThreeDSecureAuthenticationChannelThreeRITypeOtherPayment    ThreeDSecureAuthenticationChannelThreeRIType = "other_payment"
	ThreeDSecureAuthenticationChannelThreeRITypeRecurring       ThreeDSecureAuthenticationChannelThreeRIType = "recurring"
	ThreeDSecureAuthenticationChannelThreeRITypeSplitShipment   ThreeDSecureAuthenticationChannelThreeRIType = "split_shipment"
)

// Type of channel you would prefer to use for this 3DS Authentication. Only browser.
type ThreeDSecureAuthenticationChannelType string

// List of values that ThreeDSecureAuthenticationChannelType can take
const (
	ThreeDSecureAuthenticationChannelTypeBrowser ThreeDSecureAuthenticationChannelType = "browser"
	ThreeDSecureAuthenticationChannelTypeThreeRI ThreeDSecureAuthenticationChannelType = "three_r_i"
)

// The 3DS directory server with which this 3DS Authentication was processed.
type ThreeDSecureAuthenticationDirectoryServer string

// List of values that ThreeDSecureAuthenticationDirectoryServer can take
const (
	ThreeDSecureAuthenticationDirectoryServerAmericanExpress ThreeDSecureAuthenticationDirectoryServer = "american_express"
	ThreeDSecureAuthenticationDirectoryServerCartesBancaires ThreeDSecureAuthenticationDirectoryServer = "cartes_bancaires"
	ThreeDSecureAuthenticationDirectoryServerDiscover        ThreeDSecureAuthenticationDirectoryServer = "discover"
	ThreeDSecureAuthenticationDirectoryServerMastercard      ThreeDSecureAuthenticationDirectoryServer = "mastercard"
	ThreeDSecureAuthenticationDirectoryServerVisa            ThreeDSecureAuthenticationDirectoryServer = "visa"
)

// Type of challenge flow you requested for this 3DS Authentication.
type ThreeDSecureAuthenticationFlowPreferenceChallengeType string

// List of values that ThreeDSecureAuthenticationFlowPreferenceChallengeType can take
const (
	ThreeDSecureAuthenticationFlowPreferenceChallengeTypeMandated  ThreeDSecureAuthenticationFlowPreferenceChallengeType = "mandated"
	ThreeDSecureAuthenticationFlowPreferenceChallengeTypePreferred ThreeDSecureAuthenticationFlowPreferenceChallengeType = "preferred"
)

// Type of data share flow you requested for this 3DS Authentication.
type ThreeDSecureAuthenticationFlowPreferenceDataShareType string

// List of values that ThreeDSecureAuthenticationFlowPreferenceDataShareType can take
const (
	ThreeDSecureAuthenticationFlowPreferenceDataShareTypeDsSpecific  ThreeDSecureAuthenticationFlowPreferenceDataShareType = "ds_specific"
	ThreeDSecureAuthenticationFlowPreferenceDataShareTypeEmvStandard ThreeDSecureAuthenticationFlowPreferenceDataShareType = "emv_standard"
)

// Type of frictionless flow you requested for this 3DS Authentication.
type ThreeDSecureAuthenticationFlowPreferenceFrictionlessType string

// List of values that ThreeDSecureAuthenticationFlowPreferenceFrictionlessType can take
const (
	ThreeDSecureAuthenticationFlowPreferenceFrictionlessTypeLowRisk ThreeDSecureAuthenticationFlowPreferenceFrictionlessType = "low_risk"
	ThreeDSecureAuthenticationFlowPreferenceFrictionlessTypeNone    ThreeDSecureAuthenticationFlowPreferenceFrictionlessType = "none"
)

// Type of flow you requested for this 3DS Authentication.
type ThreeDSecureAuthenticationFlowPreferenceType string

// List of values that ThreeDSecureAuthenticationFlowPreferenceType can take
const (
	ThreeDSecureAuthenticationFlowPreferenceTypeChallenge    ThreeDSecureAuthenticationFlowPreferenceType = "challenge"
	ThreeDSecureAuthenticationFlowPreferenceTypeDataShare    ThreeDSecureAuthenticationFlowPreferenceType = "data_share"
	ThreeDSecureAuthenticationFlowPreferenceTypeFrictionless ThreeDSecureAuthenticationFlowPreferenceType = "frictionless"
)

type ThreeDSecureAuthenticationFutureUsageInstallmentExpiryType string

// List of values that ThreeDSecureAuthenticationFutureUsageInstallmentExpiryType can take
const (
	ThreeDSecureAuthenticationFutureUsageInstallmentExpiryTypeDate  ThreeDSecureAuthenticationFutureUsageInstallmentExpiryType = "date"
	ThreeDSecureAuthenticationFutureUsageInstallmentExpiryTypeNever ThreeDSecureAuthenticationFutureUsageInstallmentExpiryType = "never"
)

// The unit of time for `interval_count`.
type ThreeDSecureAuthenticationFutureUsageInstallmentInterval string

// List of values that ThreeDSecureAuthenticationFutureUsageInstallmentInterval can take
const (
	ThreeDSecureAuthenticationFutureUsageInstallmentIntervalDay ThreeDSecureAuthenticationFutureUsageInstallmentInterval = "day"
)

type ThreeDSecureAuthenticationFutureUsageRecurringExpiryType string

// List of values that ThreeDSecureAuthenticationFutureUsageRecurringExpiryType can take
const (
	ThreeDSecureAuthenticationFutureUsageRecurringExpiryTypeDate  ThreeDSecureAuthenticationFutureUsageRecurringExpiryType = "date"
	ThreeDSecureAuthenticationFutureUsageRecurringExpiryTypeNever ThreeDSecureAuthenticationFutureUsageRecurringExpiryType = "never"
)

// The unit of time for `interval_count`.
type ThreeDSecureAuthenticationFutureUsageRecurringInterval string

// List of values that ThreeDSecureAuthenticationFutureUsageRecurringInterval can take
const (
	ThreeDSecureAuthenticationFutureUsageRecurringIntervalDay ThreeDSecureAuthenticationFutureUsageRecurringInterval = "day"
)

// The type of future usage declared for this 3DS Authentication.
type ThreeDSecureAuthenticationFutureUsageType string

// List of values that ThreeDSecureAuthenticationFutureUsageType can take
const (
	ThreeDSecureAuthenticationFutureUsageTypeCardOnFile  ThreeDSecureAuthenticationFutureUsageType = "card_on_file"
	ThreeDSecureAuthenticationFutureUsageTypeInstallment ThreeDSecureAuthenticationFutureUsageType = "installment"
	ThreeDSecureAuthenticationFutureUsageTypeRecurring   ThreeDSecureAuthenticationFutureUsageType = "recurring"
)

// Indicates whether this 3DS Authentication is being performed for a payment or non-payment use case.
type ThreeDSecureAuthenticationMessageCategory string

// List of values that ThreeDSecureAuthenticationMessageCategory can take
const (
	ThreeDSecureAuthenticationMessageCategoryNonPaymentAuthentication ThreeDSecureAuthenticationMessageCategory = "non_payment_authentication"
	ThreeDSecureAuthenticationMessageCategoryPaymentAuthentication    ThreeDSecureAuthenticationMessageCategory = "payment_authentication"
)

// The outcome of this 3DS Authentication.
type ThreeDSecureAuthenticationOutcome string

// List of values that ThreeDSecureAuthenticationOutcome can take
const (
	ThreeDSecureAuthenticationOutcomeAbandoned           ThreeDSecureAuthenticationOutcome = "abandoned"
	ThreeDSecureAuthenticationOutcomeAttemptAcknowledged ThreeDSecureAuthenticationOutcome = "attempt_acknowledged"
	ThreeDSecureAuthenticationOutcomeAuthenticated       ThreeDSecureAuthenticationOutcome = "authenticated"
	ThreeDSecureAuthenticationOutcomeCanceled            ThreeDSecureAuthenticationOutcome = "canceled"
	ThreeDSecureAuthenticationOutcomeDenied              ThreeDSecureAuthenticationOutcome = "denied"
	ThreeDSecureAuthenticationOutcomeInformational       ThreeDSecureAuthenticationOutcome = "informational"
	ThreeDSecureAuthenticationOutcomeInternalError       ThreeDSecureAuthenticationOutcome = "internal_error"
	ThreeDSecureAuthenticationOutcomeNotSupported        ThreeDSecureAuthenticationOutcome = "not_supported"
	ThreeDSecureAuthenticationOutcomeNotTriggered        ThreeDSecureAuthenticationOutcome = "not_triggered"
	ThreeDSecureAuthenticationOutcomeProcessingError     ThreeDSecureAuthenticationOutcome = "processing_error"
	ThreeDSecureAuthenticationOutcomeRejected            ThreeDSecureAuthenticationOutcome = "rejected"
)

// TransStatus field on the ARes
type ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus string

// List of values that ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus can take
const (
	ThreeDSecureAuthenticationOutcomeDetailsAresTransStatusA ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus = "A"
	ThreeDSecureAuthenticationOutcomeDetailsAresTransStatusC ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus = "C"
	ThreeDSecureAuthenticationOutcomeDetailsAresTransStatusD ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus = "D"
	ThreeDSecureAuthenticationOutcomeDetailsAresTransStatusI ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus = "I"
	ThreeDSecureAuthenticationOutcomeDetailsAresTransStatusN ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus = "N"
	ThreeDSecureAuthenticationOutcomeDetailsAresTransStatusR ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus = "R"
	ThreeDSecureAuthenticationOutcomeDetailsAresTransStatusS ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus = "S"
	ThreeDSecureAuthenticationOutcomeDetailsAresTransStatusU ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus = "U"
	ThreeDSecureAuthenticationOutcomeDetailsAresTransStatusY ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus = "Y"
)

// The 3DS protocol version used for this 3DS Authentication.
type ThreeDSecureAuthenticationOutcomeDetailsProtocolVersion string

// List of values that ThreeDSecureAuthenticationOutcomeDetailsProtocolVersion can take
const (
	ThreeDSecureAuthenticationOutcomeDetailsProtocolVersion210 ThreeDSecureAuthenticationOutcomeDetailsProtocolVersion = "2.1.0"
	ThreeDSecureAuthenticationOutcomeDetailsProtocolVersion220 ThreeDSecureAuthenticationOutcomeDetailsProtocolVersion = "2.2.0"
	ThreeDSecureAuthenticationOutcomeDetailsProtocolVersion231 ThreeDSecureAuthenticationOutcomeDetailsProtocolVersion = "2.3.1"
)

// The indicator provided to the issuer by Stripe in the AReq that indicates whether a challenge is requested for this Authentication. This indicator should match the flow_preference you specified but may be overridden (for compliance reasons for example).
type ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator string

// List of values that ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator can take
const (
	ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator01 ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator = "01"
	ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator02 ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator = "02"
	ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator03 ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator = "03"
	ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator04 ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator = "04"
	ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator05 ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator = "05"
	ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator06 ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator = "06"
)

// TransStatus field on the RReq
type ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus string

// List of values that ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus can take
const (
	ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatusA ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus = "A"
	ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatusC ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus = "C"
	ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatusD ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus = "D"
	ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatusI ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus = "I"
	ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatusN ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus = "N"
	ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatusR ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus = "R"
	ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatusS ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus = "S"
	ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatusU ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus = "U"
	ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatusY ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus = "Y"
)

// The reason for invoking this 3DS Authentication.
type ThreeDSecureAuthenticationReason string

// List of values that ThreeDSecureAuthenticationReason can take
const (
	ThreeDSecureAuthenticationReasonCardholderAuthentication ThreeDSecureAuthenticationReason = "cardholder_authentication"
	ThreeDSecureAuthenticationReasonIssuerRequested          ThreeDSecureAuthenticationReason = "issuer_requested"
	ThreeDSecureAuthenticationReasonLiabilityShift           ThreeDSecureAuthenticationReason = "liability_shift"
	ThreeDSecureAuthenticationReasonProcessingCosts          ThreeDSecureAuthenticationReason = "processing_costs"
	ThreeDSecureAuthenticationReasonRegulatoryCompliance     ThreeDSecureAuthenticationReason = "regulatory_compliance"
)

// Status of this Authentication.
type ThreeDSecureAuthenticationStatus string

// List of values that ThreeDSecureAuthenticationStatus can take
const (
	ThreeDSecureAuthenticationStatusCanceled           ThreeDSecureAuthenticationStatus = "canceled"
	ThreeDSecureAuthenticationStatusError              ThreeDSecureAuthenticationStatus = "error"
	ThreeDSecureAuthenticationStatusFailed             ThreeDSecureAuthenticationStatus = "failed"
	ThreeDSecureAuthenticationStatusRequiresChallenge  ThreeDSecureAuthenticationStatus = "requires_challenge"
	ThreeDSecureAuthenticationStatusRequiresSubmission ThreeDSecureAuthenticationStatus = "requires_submission"
	ThreeDSecureAuthenticationStatusSucceeded          ThreeDSecureAuthenticationStatus = "succeeded"
)

// Returns a list of 3D Secure Authentications.
type ThreeDSecureAuthenticationListParams struct {
	ListParams `form:"*"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp or a dictionary with a number of different query options.
	Created *int64 `form:"created" json:"created,omitempty"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp or a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created" json:"-"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Only return 3D Secure Authentications for specified status.
	Status *string `form:"status" json:"status,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *ThreeDSecureAuthenticationListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Contains additional details about the acquirer for this 3DS Authentication.
//
// Refer to the [Pass acquirer details and directory server section of the standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#pass-acquirer-details-and-directory-server) for more information.
type ThreeDSecureAuthenticationAcquirerDetailsParams struct {
	// The Acquirer BIN (specific to the directory_server).
	AcquirerBin *string `form:"acquirer_bin" json:"acquirer_bin"`
	// The two-letter country code of the acquirer ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	AcquirerCountry *string `form:"acquirer_country" json:"acquirer_country"`
	// The Merchant ID (or Card Acceptor ID) that your acquirer assigned you (specific to the directory_server).
	AcquirerMerchantID *string `form:"acquirer_merchant_id" json:"acquirer_merchant_id"`
	// The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) as defined by each payment system or directory server.
	MCC *string `form:"mcc" json:"mcc,omitempty"`
	// The merchant name assigned by the acquirer or payment system. Same name used in the authorization message as defined in [ISO 8583](https://en.wikipedia.org/wiki/ISO_8583).
	MerchantName *string `form:"merchant_name" json:"merchant_name,omitempty"`
	// Requestor ID if you're enrolled in the card network's 3DS program. Otherwise, you can omit this field because Stripe assigns a Requestor ID with the card networks.
	RequestorID *string `form:"requestor_id" json:"requestor_id,omitempty"`
}

// Contains additional details about the browser details you collected.
type ThreeDSecureAuthenticationChannelBrowserParams struct {
	// The HTTP accept headers from the cardholder's browser. Collected server-side.
	AcceptHeader *string `form:"accept_header" json:"accept_header"`
	// The color depth of the cardholder's screen.
	//
	// Returned from the `screen.colorDepth` property.
	ColorDepth *int64 `form:"color_depth" json:"color_depth,omitempty"`
	// Unique and immutable identifier linked to a device that is consistent across 3DS transactions for the specific user device. For example: hardware device ID or a platform-calculated device fingerprint.
	DeviceID *string `form:"device_id" json:"device_id,omitempty"`
	// The IP address of the browser. Included in the HTTP request to your server before you create the 3DS Authentication.
	//
	// Collected server-side.
	IPAddress *string `form:"ip_address" json:"ip_address"`
	// The cardholder browser's ability to execute Java. Returned from the navigator.javaEnabled property.
	JavaEnabled *bool `form:"java_enabled" json:"java_enabled,omitempty"`
	// The cardholder browser's ability to execute JavaScript.
	JavascriptEnabled *bool `form:"javascript_enabled" json:"javascript_enabled"`
	// An IETF BCP 47 language tag representing the browser language. Typically returned from the `navigator.language` property, but might also be returned from `navigator.languages` or `navigator.browserLanguage`.
	//
	//  In some cases, this value might be an array. To cast it to a string or null value, you can use the `getBrowserLanguage()` [example function](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#pass-client-side-collected-channel-information).
	Language *string `form:"language" json:"language"`
	// The total height of the cardholder's screen in pixels.
	//
	// Returned from the `screen.height` property.
	ScreenHeight *int64 `form:"screen_height" json:"screen_height,omitempty"`
	// The total width of the cardholder's screen in pixels.
	//
	// Returned from the `screen.width` property.
	ScreenWidth *int64 `form:"screen_width" json:"screen_width,omitempty"`
	// The time difference between UTC time and the local time of the cardholder's browser, in minutes.
	//
	// Returned by `new Date().getTimezoneOffset()`
	TimezoneOffset *int64 `form:"timezone_offset" json:"timezone_offset,omitempty"`
	// The browser user agent. You can retrieve this value on the client side using the `navigator.userAgent` property, or in the HTTP request to your server before you create the 3DS Authentication.
	UserAgent *string `form:"user_agent" json:"user_agent"`
}

// Contains additional details about the 3DS Requestor Initiated (3RI) channel.
type ThreeDSecureAuthenticationChannelThreeRIParams struct {
	// ID of a prior `Authentication`. For example, the first recurring transaction that was authenticated by the cardholder.
	PreviousAuthentication *string `form:"previous_authentication" json:"previous_authentication"`
	// It provides additional information to the ACS to determine the best approach for handling a 3RI request.
	Type *string `form:"type" json:"type"`
}

// Contains additional details on the channel used for this 3DS Authentication.
type ThreeDSecureAuthenticationChannelParams struct {
	// Contains additional details about the browser details you collected.
	Browser *ThreeDSecureAuthenticationChannelBrowserParams `form:"browser" json:"browser,omitempty"`
	// Contains additional details about the 3DS Requestor Initiated (3RI) channel.
	ThreeRI *ThreeDSecureAuthenticationChannelThreeRIParams `form:"three_r_i" json:"three_r_i,omitempty"`
	// Type of channel you would prefer to use for this 3DS Authentication.
	Type *string `form:"type" json:"type"`
}

// Contains additional details about your challenge flow preference for this 3DS Authentication.
type ThreeDSecureAuthenticationFlowPreferenceChallengeParams struct {
	// Type of challenge flow you requested for this 3DS Authentication.
	Type *string `form:"type" json:"type"`
}

// Contains additional details about your data share only flow preference for this 3DS Authentication.
type ThreeDSecureAuthenticationFlowPreferenceDataShareParams struct {
	// Type of data share only flow you requested for this 3DS Authentication.
	Type *string `form:"type" json:"type"`
}

// Contains additional details about your frictionless flow preference for this 3DS Authentication.
type ThreeDSecureAuthenticationFlowPreferenceFrictionlessParams struct {
	// Type of frictionless flow you requested for this 3DS Authentication.
	Type *string `form:"type" json:"type"`
}

// Contains additional details on your flow preference for this 3DS Authentication.
//
// Refer to the [Specify a flow preference section of the standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#specify-a-flow-preference) for more information.
type ThreeDSecureAuthenticationFlowPreferenceParams struct {
	// Contains additional details about your challenge flow preference for this 3DS Authentication.
	Challenge *ThreeDSecureAuthenticationFlowPreferenceChallengeParams `form:"challenge" json:"challenge,omitempty"`
	// Contains additional details about your data share only flow preference for this 3DS Authentication.
	DataShare *ThreeDSecureAuthenticationFlowPreferenceDataShareParams `form:"data_share" json:"data_share,omitempty"`
	// Contains additional details about your frictionless flow preference for this 3DS Authentication.
	Frictionless *ThreeDSecureAuthenticationFlowPreferenceFrictionlessParams `form:"frictionless" json:"frictionless,omitempty"`
	// Type of flow you requested for this 3DS Authentication.
	Type *string `form:"type" json:"type"`
}

// Information about the expiry of the future usage of this authentication.
type ThreeDSecureAuthenticationFutureUsageInstallmentExpiryParams struct {
	// The date before which the last authorization related to this authentication will occur.
	Date *string `form:"date" json:"date,omitempty"`
	// The type of expiry for the future use of this authentication.
	Type *string `form:"type" json:"type"`
}

// Parameters related to an installment payment.
type ThreeDSecureAuthenticationFutureUsageInstallmentParams struct {
	// A non-negative integer representing the future authorizations' amount in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount *int64 `form:"amount" json:"amount"`
	// Information about the expiry of the future usage of this authentication.
	Expiry *ThreeDSecureAuthenticationFutureUsageInstallmentExpiryParams `form:"expiry" json:"expiry"`
	// The unit of time for `interval_count`.
	Interval *string `form:"interval" json:"interval,omitempty"`
	// The minimum number of time intervals between authorizations. Must be greater than 0, and defaults to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
	// The maximum number of installments. Must be greater than 1.
	Number *int64 `form:"number" json:"number"`
}

// Information about the expiry of the future usage of this authentication.
type ThreeDSecureAuthenticationFutureUsageRecurringExpiryParams struct {
	// The date before which the last authorization related to this authentication will occur.
	Date *string `form:"date" json:"date,omitempty"`
	// The type of expiry for the future use of this authentication.
	Type *string `form:"type" json:"type"`
}

// Parameters related to a recurring payment.
type ThreeDSecureAuthenticationFutureUsageRecurringParams struct {
	// A non-negative integer representing the future authorizations' amount in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount *int64 `form:"amount" json:"amount"`
	// Information about the expiry of the future usage of this authentication.
	Expiry *ThreeDSecureAuthenticationFutureUsageRecurringExpiryParams `form:"expiry" json:"expiry"`
	// The unit of time for `interval_count`.
	Interval *string `form:"interval" json:"interval,omitempty"`
	// The minimum number of time intervals between authorizations. Must be greater than 0, and defaults to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
}

// Contains information about future usage of this 3DS Authentication
type ThreeDSecureAuthenticationFutureUsageParams struct {
	// Parameters related to an installment payment.
	Installment *ThreeDSecureAuthenticationFutureUsageInstallmentParams `form:"installment" json:"installment,omitempty"`
	// Parameters related to a recurring payment.
	Recurring *ThreeDSecureAuthenticationFutureUsageRecurringParams `form:"recurring" json:"recurring,omitempty"`
	// The type of future usage declared for this 3DS Authentication
	Type *string `form:"type" json:"type"`
}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type ThreeDSecureAuthenticationPaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
}
type ThreeDSecureAuthenticationPaymentMethodDataCardParams struct {
	CVC      *string `form:"cvc" json:"cvc,omitempty"`
	ExpMonth *int64  `form:"exp_month" json:"exp_month,omitempty"`
	ExpYear  *int64  `form:"exp_year" json:"exp_year,omitempty"`
	Number   *string `form:"number" json:"number,omitempty"`
	Token    *string `form:"token" json:"token,omitempty"`
}

// Hash used to generate the PaymentMethod to be used for this Authentication. This is mutually exclusive with the `payment_method` parameter.
type ThreeDSecureAuthenticationPaymentMethodDataParams struct {
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *ThreeDSecureAuthenticationPaymentMethodDataBillingDetailsParams `form:"billing_details" json:"billing_details,omitempty"`
	Card           *ThreeDSecureAuthenticationPaymentMethodDataCardParams           `form:"card" json:"card"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type" json:"type"`
}

// This endpoint creates a 3DS Authentication. Refer to the [Create a 3DS Authentication object section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#create-a-3ds-authentication-object) for more information.
//
// You can pass the submit parameter to automatically submit the 3DS Authentication object when you create it. Refer to the [Submit at creation section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-at-creation) for more information.
type ThreeDSecureAuthenticationParams struct {
	Params `form:"*"`
	// Contains additional details about the acquirer for this 3DS Authentication.
	//
	// Refer to the [Pass acquirer details and directory server section of the standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#pass-acquirer-details-and-directory-server) for more information.
	AcquirerDetails *ThreeDSecureAuthenticationAcquirerDetailsParams `form:"acquirer_details" json:"acquirer_details,omitempty"`
	// A non-negative integer representing the amount in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). You can't include this parameter if `message_category` is `non_payment_authentication`
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// Contains additional details on the channel used for this 3DS Authentication.
	Channel *ThreeDSecureAuthenticationChannelParams `form:"channel" json:"channel,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The 3DS directory server with which this 3DS Authentication was processed.
	DirectoryServer *string `form:"directory_server" json:"directory_server,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Contains additional details on your flow preference for this 3DS Authentication.
	//
	// Refer to the [Specify a flow preference section of the standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#specify-a-flow-preference) for more information.
	FlowPreference *ThreeDSecureAuthenticationFlowPreferenceParams `form:"flow_preference" json:"flow_preference,omitempty"`
	// Contains information about future usage of this 3DS Authentication
	FutureUsage *ThreeDSecureAuthenticationFutureUsageParams `form:"future_usage" json:"future_usage,omitempty"`
	// Indicates whether this 3DS Authentication is being performed for a payment or non-payment use case.
	MessageCategory *string `form:"message_category" json:"message_category,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// ID of the payment method (a PaymentMethod object) to attach to this 3DS Authentication.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// Hash used to generate the PaymentMethod to be used for this Authentication. This is mutually exclusive with the `payment_method` parameter.
	PaymentMethodData *ThreeDSecureAuthenticationPaymentMethodDataParams `form:"payment_method_data" json:"payment_method_data,omitempty"`
	// The reason for invoking standalone 3DS. This is tailored specifically for cases when you want Stripe to help determine the standalone 3DS flow to fit your use case instead of needing to select a specific 3DS flow.
	//
	// This parameter is exclusive with `flow_preference`. You can either use `reason` for controlling 3DS according to your business requirements, or use `flow_preference` for having fine-grained control over your 3DS flow preference.
	Reason *string `form:"reason" json:"reason,omitempty"`
	// The shipping address requested by the cardholder. You should try to include as complete address information as possible.
	ShippingAddress *AddressParams `form:"shipping_address" json:"shipping_address,omitempty"`
	// Set to `always` to skip the fingerprinting step and submit this Authentication immediately or `if_fingerprinting_not_supported` to submit this Authentication only if fingerprinting is not available. This parameter defaults to `never`.
	//
	// Refer to the [Submit at creation section of the standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-at-creation) for more information.
	Submit      *string                                      `form:"submit" json:"submit,omitempty"`
	UnsetFields []ThreeDSecureAuthenticationParamsUnsetField `form:"-" json:"-"`
}

// ThreeDSecureAuthenticationParamsUnsetField is the list of fields that can be cleared/unset on ThreeDSecureAuthenticationParams.
type ThreeDSecureAuthenticationParamsUnsetField string

const (
	ThreeDSecureAuthenticationParamsUnsetFieldMetadata ThreeDSecureAuthenticationParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *ThreeDSecureAuthenticationParams) AddUnsetField(field ThreeDSecureAuthenticationParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *ThreeDSecureAuthenticationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *ThreeDSecureAuthenticationParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// This endpoint cancels a 3DS Authentication. You can cancel a 3DS Authentication object when it's in a non-final status:
// requires_submission or requires_challenge.
type ThreeDSecureAuthenticationCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata    map[string]string                                  `form:"metadata" json:"metadata,omitempty"`
	UnsetFields []ThreeDSecureAuthenticationCancelParamsUnsetField `form:"-" json:"-"`
}

// ThreeDSecureAuthenticationCancelParamsUnsetField is the list of fields that can be cleared/unset on ThreeDSecureAuthenticationCancelParams.
type ThreeDSecureAuthenticationCancelParamsUnsetField string

const (
	ThreeDSecureAuthenticationCancelParamsUnsetFieldMetadata ThreeDSecureAuthenticationCancelParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *ThreeDSecureAuthenticationCancelParams) AddUnsetField(field ThreeDSecureAuthenticationCancelParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *ThreeDSecureAuthenticationCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *ThreeDSecureAuthenticationCancelParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// This endpoint submits a 3DS Authentication. You can submit a 3DS Authentication object when it has status requires_submission. Refer to the [Submit the 3DS Authentication object section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-the-3ds-authentication-object) for more information.
type ThreeDSecureAuthenticationSubmitParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The fingerprinting result of the issuer fingerprinting step.
	//
	// Refer to the [Issuer fingerprinting section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#issuer-fingerprinting) for more information.
	FingerprintingResult *string `form:"fingerprinting_result" json:"fingerprinting_result,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata    map[string]string                                  `form:"metadata" json:"metadata,omitempty"`
	UnsetFields []ThreeDSecureAuthenticationSubmitParamsUnsetField `form:"-" json:"-"`
}

// ThreeDSecureAuthenticationSubmitParamsUnsetField is the list of fields that can be cleared/unset on ThreeDSecureAuthenticationSubmitParams.
type ThreeDSecureAuthenticationSubmitParamsUnsetField string

const (
	ThreeDSecureAuthenticationSubmitParamsUnsetFieldMetadata ThreeDSecureAuthenticationSubmitParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *ThreeDSecureAuthenticationSubmitParams) AddUnsetField(field ThreeDSecureAuthenticationSubmitParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *ThreeDSecureAuthenticationSubmitParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *ThreeDSecureAuthenticationSubmitParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Contains additional details about the acquirer for this 3DS Authentication.
//
// Refer to the [Pass acquirer details and directory server section of the standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#pass-acquirer-details-and-directory-server) for more information.
type ThreeDSecureAuthenticationCreateAcquirerDetailsParams struct {
	// The Acquirer BIN (specific to the directory_server).
	AcquirerBin *string `form:"acquirer_bin" json:"acquirer_bin"`
	// The two-letter country code of the acquirer ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	AcquirerCountry *string `form:"acquirer_country" json:"acquirer_country"`
	// The Merchant ID (or Card Acceptor ID) that your acquirer assigned you (specific to the directory_server).
	AcquirerMerchantID *string `form:"acquirer_merchant_id" json:"acquirer_merchant_id"`
	// The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) as defined by each payment system or directory server.
	MCC *string `form:"mcc" json:"mcc,omitempty"`
	// The merchant name assigned by the acquirer or payment system. Same name used in the authorization message as defined in [ISO 8583](https://en.wikipedia.org/wiki/ISO_8583).
	MerchantName *string `form:"merchant_name" json:"merchant_name,omitempty"`
	// Requestor ID if you're enrolled in the card network's 3DS program. Otherwise, you can omit this field because Stripe assigns a Requestor ID with the card networks.
	RequestorID *string `form:"requestor_id" json:"requestor_id,omitempty"`
}

// Contains additional details about the browser details you collected.
type ThreeDSecureAuthenticationCreateChannelBrowserParams struct {
	// The HTTP accept headers from the cardholder's browser. Collected server-side.
	AcceptHeader *string `form:"accept_header" json:"accept_header"`
	// The color depth of the cardholder's screen.
	//
	// Returned from the `screen.colorDepth` property.
	ColorDepth *int64 `form:"color_depth" json:"color_depth,omitempty"`
	// Unique and immutable identifier linked to a device that is consistent across 3DS transactions for the specific user device. For example: hardware device ID or a platform-calculated device fingerprint.
	DeviceID *string `form:"device_id" json:"device_id,omitempty"`
	// The IP address of the browser. Included in the HTTP request to your server before you create the 3DS Authentication.
	//
	// Collected server-side.
	IPAddress *string `form:"ip_address" json:"ip_address"`
	// The cardholder browser's ability to execute Java. Returned from the navigator.javaEnabled property.
	JavaEnabled *bool `form:"java_enabled" json:"java_enabled,omitempty"`
	// The cardholder browser's ability to execute JavaScript.
	JavascriptEnabled *bool `form:"javascript_enabled" json:"javascript_enabled"`
	// An IETF BCP 47 language tag representing the browser language. Typically returned from the `navigator.language` property, but might also be returned from `navigator.languages` or `navigator.browserLanguage`.
	//
	//  In some cases, this value might be an array. To cast it to a string or null value, you can use the `getBrowserLanguage()` [example function](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#pass-client-side-collected-channel-information).
	Language *string `form:"language" json:"language"`
	// The total height of the cardholder's screen in pixels.
	//
	// Returned from the `screen.height` property.
	ScreenHeight *int64 `form:"screen_height" json:"screen_height,omitempty"`
	// The total width of the cardholder's screen in pixels.
	//
	// Returned from the `screen.width` property.
	ScreenWidth *int64 `form:"screen_width" json:"screen_width,omitempty"`
	// The time difference between UTC time and the local time of the cardholder's browser, in minutes.
	//
	// Returned by `new Date().getTimezoneOffset()`
	TimezoneOffset *int64 `form:"timezone_offset" json:"timezone_offset,omitempty"`
	// The browser user agent. You can retrieve this value on the client side using the `navigator.userAgent` property, or in the HTTP request to your server before you create the 3DS Authentication.
	UserAgent *string `form:"user_agent" json:"user_agent"`
}

// Contains additional details about the 3DS Requestor Initiated (3RI) channel.
type ThreeDSecureAuthenticationCreateChannelThreeRIParams struct {
	// ID of a prior `Authentication`. For example, the first recurring transaction that was authenticated by the cardholder.
	PreviousAuthentication *string `form:"previous_authentication" json:"previous_authentication"`
	// It provides additional information to the ACS to determine the best approach for handling a 3RI request.
	Type *string `form:"type" json:"type"`
}

// Contains additional details on the channel used for this 3DS Authentication.
type ThreeDSecureAuthenticationCreateChannelParams struct {
	// Contains additional details about the browser details you collected.
	Browser *ThreeDSecureAuthenticationCreateChannelBrowserParams `form:"browser" json:"browser,omitempty"`
	// Contains additional details about the 3DS Requestor Initiated (3RI) channel.
	ThreeRI *ThreeDSecureAuthenticationCreateChannelThreeRIParams `form:"three_r_i" json:"three_r_i,omitempty"`
	// Type of channel you would prefer to use for this 3DS Authentication.
	Type *string `form:"type" json:"type"`
}

// Contains additional details about your challenge flow preference for this 3DS Authentication.
type ThreeDSecureAuthenticationCreateFlowPreferenceChallengeParams struct {
	// Type of challenge flow you requested for this 3DS Authentication.
	Type *string `form:"type" json:"type"`
}

// Contains additional details about your data share only flow preference for this 3DS Authentication.
type ThreeDSecureAuthenticationCreateFlowPreferenceDataShareParams struct {
	// Type of data share only flow you requested for this 3DS Authentication.
	Type *string `form:"type" json:"type"`
}

// Contains additional details about your frictionless flow preference for this 3DS Authentication.
type ThreeDSecureAuthenticationCreateFlowPreferenceFrictionlessParams struct {
	// Type of frictionless flow you requested for this 3DS Authentication.
	Type *string `form:"type" json:"type"`
}

// Contains additional details on your flow preference for this 3DS Authentication.
//
// Refer to the [Specify a flow preference section of the standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#specify-a-flow-preference) for more information.
type ThreeDSecureAuthenticationCreateFlowPreferenceParams struct {
	// Contains additional details about your challenge flow preference for this 3DS Authentication.
	Challenge *ThreeDSecureAuthenticationCreateFlowPreferenceChallengeParams `form:"challenge" json:"challenge,omitempty"`
	// Contains additional details about your data share only flow preference for this 3DS Authentication.
	DataShare *ThreeDSecureAuthenticationCreateFlowPreferenceDataShareParams `form:"data_share" json:"data_share,omitempty"`
	// Contains additional details about your frictionless flow preference for this 3DS Authentication.
	Frictionless *ThreeDSecureAuthenticationCreateFlowPreferenceFrictionlessParams `form:"frictionless" json:"frictionless,omitempty"`
	// Type of flow you requested for this 3DS Authentication.
	Type *string `form:"type" json:"type"`
}

// Information about the expiry of the future usage of this authentication.
type ThreeDSecureAuthenticationCreateFutureUsageInstallmentExpiryParams struct {
	// The date before which the last authorization related to this authentication will occur.
	Date *string `form:"date" json:"date,omitempty"`
	// The type of expiry for the future use of this authentication.
	Type *string `form:"type" json:"type"`
}

// Parameters related to an installment payment.
type ThreeDSecureAuthenticationCreateFutureUsageInstallmentParams struct {
	// A non-negative integer representing the future authorizations' amount in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount *int64 `form:"amount" json:"amount"`
	// Information about the expiry of the future usage of this authentication.
	Expiry *ThreeDSecureAuthenticationCreateFutureUsageInstallmentExpiryParams `form:"expiry" json:"expiry"`
	// The unit of time for `interval_count`.
	Interval *string `form:"interval" json:"interval,omitempty"`
	// The minimum number of time intervals between authorizations. Must be greater than 0, and defaults to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
	// The maximum number of installments. Must be greater than 1.
	Number *int64 `form:"number" json:"number"`
}

// Information about the expiry of the future usage of this authentication.
type ThreeDSecureAuthenticationCreateFutureUsageRecurringExpiryParams struct {
	// The date before which the last authorization related to this authentication will occur.
	Date *string `form:"date" json:"date,omitempty"`
	// The type of expiry for the future use of this authentication.
	Type *string `form:"type" json:"type"`
}

// Parameters related to a recurring payment.
type ThreeDSecureAuthenticationCreateFutureUsageRecurringParams struct {
	// A non-negative integer representing the future authorizations' amount in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount *int64 `form:"amount" json:"amount"`
	// Information about the expiry of the future usage of this authentication.
	Expiry *ThreeDSecureAuthenticationCreateFutureUsageRecurringExpiryParams `form:"expiry" json:"expiry"`
	// The unit of time for `interval_count`.
	Interval *string `form:"interval" json:"interval,omitempty"`
	// The minimum number of time intervals between authorizations. Must be greater than 0, and defaults to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
}

// Contains information about future usage of this 3DS Authentication
type ThreeDSecureAuthenticationCreateFutureUsageParams struct {
	// Parameters related to an installment payment.
	Installment *ThreeDSecureAuthenticationCreateFutureUsageInstallmentParams `form:"installment" json:"installment,omitempty"`
	// Parameters related to a recurring payment.
	Recurring *ThreeDSecureAuthenticationCreateFutureUsageRecurringParams `form:"recurring" json:"recurring,omitempty"`
	// The type of future usage declared for this 3DS Authentication
	Type *string `form:"type" json:"type"`
}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type ThreeDSecureAuthenticationCreatePaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
}
type ThreeDSecureAuthenticationCreatePaymentMethodDataCardParams struct {
	CVC      *string `form:"cvc" json:"cvc,omitempty"`
	ExpMonth *int64  `form:"exp_month" json:"exp_month,omitempty"`
	ExpYear  *int64  `form:"exp_year" json:"exp_year,omitempty"`
	Number   *string `form:"number" json:"number,omitempty"`
	Token    *string `form:"token" json:"token,omitempty"`
}

// Hash used to generate the PaymentMethod to be used for this Authentication. This is mutually exclusive with the `payment_method` parameter.
type ThreeDSecureAuthenticationCreatePaymentMethodDataParams struct {
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *ThreeDSecureAuthenticationCreatePaymentMethodDataBillingDetailsParams `form:"billing_details" json:"billing_details,omitempty"`
	Card           *ThreeDSecureAuthenticationCreatePaymentMethodDataCardParams           `form:"card" json:"card"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type" json:"type"`
}

// This endpoint creates a 3DS Authentication. Refer to the [Create a 3DS Authentication object section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#create-a-3ds-authentication-object) for more information.
//
// You can pass the submit parameter to automatically submit the 3DS Authentication object when you create it. Refer to the [Submit at creation section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-at-creation) for more information.
type ThreeDSecureAuthenticationCreateParams struct {
	Params `form:"*"`
	// Contains additional details about the acquirer for this 3DS Authentication.
	//
	// Refer to the [Pass acquirer details and directory server section of the standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#pass-acquirer-details-and-directory-server) for more information.
	AcquirerDetails *ThreeDSecureAuthenticationCreateAcquirerDetailsParams `form:"acquirer_details" json:"acquirer_details,omitempty"`
	// A non-negative integer representing the amount in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). You can't include this parameter if `message_category` is `non_payment_authentication`
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// Contains additional details on the channel used for this 3DS Authentication.
	Channel *ThreeDSecureAuthenticationCreateChannelParams `form:"channel" json:"channel"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The 3DS directory server with which this 3DS Authentication was processed.
	DirectoryServer *string `form:"directory_server" json:"directory_server,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Contains additional details on your flow preference for this 3DS Authentication.
	//
	// Refer to the [Specify a flow preference section of the standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#specify-a-flow-preference) for more information.
	FlowPreference *ThreeDSecureAuthenticationCreateFlowPreferenceParams `form:"flow_preference" json:"flow_preference,omitempty"`
	// Contains information about future usage of this 3DS Authentication
	FutureUsage *ThreeDSecureAuthenticationCreateFutureUsageParams `form:"future_usage" json:"future_usage,omitempty"`
	// Indicates whether this 3DS Authentication is being performed for a payment or non-payment use case.
	MessageCategory *string `form:"message_category" json:"message_category"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// ID of the payment method (a PaymentMethod object) to attach to this 3DS Authentication.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// Hash used to generate the PaymentMethod to be used for this Authentication. This is mutually exclusive with the `payment_method` parameter.
	PaymentMethodData *ThreeDSecureAuthenticationCreatePaymentMethodDataParams `form:"payment_method_data" json:"payment_method_data,omitempty"`
	// The reason for invoking standalone 3DS. This is tailored specifically for cases when you want Stripe to help determine the standalone 3DS flow to fit your use case instead of needing to select a specific 3DS flow.
	//
	// This parameter is exclusive with `flow_preference`. You can either use `reason` for controlling 3DS according to your business requirements, or use `flow_preference` for having fine-grained control over your 3DS flow preference.
	Reason *string `form:"reason" json:"reason,omitempty"`
	// The shipping address requested by the cardholder. You should try to include as complete address information as possible.
	ShippingAddress *AddressParams `form:"shipping_address" json:"shipping_address,omitempty"`
	// Set to `always` to skip the fingerprinting step and submit this Authentication immediately or `if_fingerprinting_not_supported` to submit this Authentication only if fingerprinting is not available. This parameter defaults to `never`.
	//
	// Refer to the [Submit at creation section of the standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-at-creation) for more information.
	Submit      *string                                            `form:"submit" json:"submit,omitempty"`
	UnsetFields []ThreeDSecureAuthenticationCreateParamsUnsetField `form:"-" json:"-"`
}

// ThreeDSecureAuthenticationCreateParamsUnsetField is the list of fields that can be cleared/unset on ThreeDSecureAuthenticationCreateParams.
type ThreeDSecureAuthenticationCreateParamsUnsetField string

const (
	ThreeDSecureAuthenticationCreateParamsUnsetFieldMetadata ThreeDSecureAuthenticationCreateParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *ThreeDSecureAuthenticationCreateParams) AddUnsetField(field ThreeDSecureAuthenticationCreateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *ThreeDSecureAuthenticationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *ThreeDSecureAuthenticationCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// This endpoint retrieves a 3DS Authentication.
type ThreeDSecureAuthenticationRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *ThreeDSecureAuthenticationRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Contains additional details about the acquirer for a 3DS Authentication.
type ThreeDSecureAuthenticationAcquirerDetails struct {
	// The Acquirer BIN (specific to the directory_server).
	AcquirerBin string `json:"acquirer_bin,omitempty"`
	// The two-letter country code of the acquirer ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	AcquirerCountry string `json:"acquirer_country,omitempty"`
	// The Merchant ID (or Card Acceptor ID) that your acquirer assigned you (specific to the directory_server).
	AcquirerMerchantID string `json:"acquirer_merchant_id,omitempty"`
	// The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) as defined by each payment system or directory server.
	MCC string `json:"mcc,omitempty"`
	// The merchant name assigned by the acquirer or payment system. Same name used in the authorization message as defined in [ISO 8583](https://en.wikipedia.org/wiki/ISO_8583).
	MerchantName string `json:"merchant_name,omitempty"`
	// Requestor ID if you're enrolled in the card network's 3DS program. Otherwise, you can omit this field because Stripe assigns a Requestor ID with the card networks.
	RequestorID string `json:"requestor_id,omitempty"`
}

// Contains details on the browser for a standalone 3DS Authentication.
type ThreeDSecureAuthenticationChannelBrowser struct {
	// The HTTP accept headers from the cardholder's browser.
	AcceptHeader string `json:"accept_header"`
	// The color depth of the cardholder's screen.
	ColorDepth int64 `json:"color_depth,omitempty"`
	// The IP address of the browser.
	IPAddress string `json:"ip_address"`
	// The cardholder browser's ability to execute Java.
	JavaEnabled bool `json:"java_enabled,omitempty"`
	// The cardholder browser's ability to execute JavaScript.
	JavascriptEnabled bool `json:"javascript_enabled"`
	// An IETF BCP 47 language tag representing the browser language.
	Language string `json:"language"`
	// The total height of the cardholder's screen in pixels.
	ScreenHeight int64 `json:"screen_height,omitempty"`
	// The total width of the cardholder's screen in pixels.
	ScreenWidth int64 `json:"screen_width,omitempty"`
	// The time difference between UTC time and the local time of the cardholder's browser, in minutes.
	TimezoneOffset int64 `json:"timezone_offset,omitempty"`
	// The browser user agent.
	UserAgent string `json:"user_agent"`
}

// Contains details for a 3RI standalone 3DS Authentication.
type ThreeDSecureAuthenticationChannelThreeRI struct {
	// ID of the previous initial authenticated 3DS Authentication object.
	PreviousAuthentication string `json:"previous_authentication"`
	// Type of the 3RI Authentication.
	Type ThreeDSecureAuthenticationChannelThreeRIType `json:"type"`
}

// Contains details on the channel used (browser, 3RI) for a standalone 3DS Authentication.
type ThreeDSecureAuthenticationChannel struct {
	// Contains details on the browser for a standalone 3DS Authentication.
	Browser *ThreeDSecureAuthenticationChannelBrowser `json:"browser,omitempty"`
	// Contains details for a 3RI standalone 3DS Authentication.
	ThreeRI *ThreeDSecureAuthenticationChannelThreeRI `json:"three_r_i,omitempty"`
	// Type of channel you would prefer to use for this 3DS Authentication. Only browser.
	Type ThreeDSecureAuthenticationChannelType `json:"type"`
}
type ThreeDSecureAuthenticationFlowPreferenceChallenge struct {
	// Type of challenge flow you requested for this 3DS Authentication.
	Type ThreeDSecureAuthenticationFlowPreferenceChallengeType `json:"type"`
}
type ThreeDSecureAuthenticationFlowPreferenceDataShare struct {
	// Type of data share flow you requested for this 3DS Authentication.
	Type ThreeDSecureAuthenticationFlowPreferenceDataShareType `json:"type"`
}
type ThreeDSecureAuthenticationFlowPreferenceFrictionless struct {
	// Type of frictionless flow you requested for this 3DS Authentication.
	Type ThreeDSecureAuthenticationFlowPreferenceFrictionlessType `json:"type"`
}

// Contains details of the flow preference used for a standalone 3DS Authentication.
type ThreeDSecureAuthenticationFlowPreference struct {
	Challenge    *ThreeDSecureAuthenticationFlowPreferenceChallenge    `json:"challenge,omitempty"`
	DataShare    *ThreeDSecureAuthenticationFlowPreferenceDataShare    `json:"data_share,omitempty"`
	Frictionless *ThreeDSecureAuthenticationFlowPreferenceFrictionless `json:"frictionless,omitempty"`
	// Type of flow you requested for this 3DS Authentication.
	Type ThreeDSecureAuthenticationFlowPreferenceType `json:"type"`
}

// Information about recurring payment expiry
type ThreeDSecureAuthenticationFutureUsageInstallmentExpiry struct {
	Date string                                                     `json:"date,omitempty"`
	Type ThreeDSecureAuthenticationFutureUsageInstallmentExpiryType `json:"type"`
}

// Details about installment payments
type ThreeDSecureAuthenticationFutureUsageInstallment struct {
	// A non-negative integer representing the amount in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount int64 `json:"amount,omitempty"`
	// Information about recurring payment expiry
	Expiry *ThreeDSecureAuthenticationFutureUsageInstallmentExpiry `json:"expiry"`
	// The unit of time for `interval_count`.
	Interval ThreeDSecureAuthenticationFutureUsageInstallmentInterval `json:"interval"`
	// The minimum number of time intervals between authorizations.
	IntervalCount int64 `json:"interval_count"`
	// The maximum number of installments.
	Number int64 `json:"number"`
}

// Information about recurring payment expiry
type ThreeDSecureAuthenticationFutureUsageRecurringExpiry struct {
	Date string                                                   `json:"date,omitempty"`
	Type ThreeDSecureAuthenticationFutureUsageRecurringExpiryType `json:"type"`
}

// Details about recurring payments
type ThreeDSecureAuthenticationFutureUsageRecurring struct {
	// A non-negative integer representing the amount in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount int64 `json:"amount,omitempty"`
	// Information about recurring payment expiry
	Expiry *ThreeDSecureAuthenticationFutureUsageRecurringExpiry `json:"expiry"`
	// The unit of time for `interval_count`.
	Interval ThreeDSecureAuthenticationFutureUsageRecurringInterval `json:"interval"`
	// The minimum number of time intervals between authorizations.
	IntervalCount int64 `json:"interval_count"`
}

// Contains information about the future authorisations related to this authentication
type ThreeDSecureAuthenticationFutureUsage struct {
	// Details about installment payments
	Installment *ThreeDSecureAuthenticationFutureUsageInstallment `json:"installment,omitempty"`
	// Details about recurring payments
	Recurring *ThreeDSecureAuthenticationFutureUsageRecurring `json:"recurring,omitempty"`
	// The type of future usage declared for this 3DS Authentication.
	Type ThreeDSecureAuthenticationFutureUsageType `json:"type"`
}

// Contains details for Cartes Bancaires specific fields in the authentication outcomes.
type ThreeDSecureAuthenticationOutcomeDetailsNetworkDetailsCartesBancaires struct {
	// The cryptogram calculation algorithm used by the card Issuer's ACS to calculate the Authentication cryptogram. Also known as cavvAlgorithm. ARes/RReq messageExtension: `CB-AVALGO`
	Avalgo string `json:"avalgo"`
	// The exemption indicator returned from Cartes Bancaires in the ARes. This is a 3 byte bitmap (lowest significant byte first and most significant bit first) that has been Base64 encoded. String (4 characters). ARes message extension: `CB-EXEMPTION`
	CbExemption string `json:"cb_exemption"`
	// The risk score returned from Cartes Bancaires in the ARes. Numeric value 0-99. ARes/RReq message extension: `CB-SCORE`
	CbScore string `json:"cb_score"`
}

// Contains details specific to the individual network.
type ThreeDSecureAuthenticationOutcomeDetailsNetworkDetails struct {
	// Contains details for Cartes Bancaires specific fields in the authentication outcomes.
	CartesBancaires *ThreeDSecureAuthenticationOutcomeDetailsNetworkDetailsCartesBancaires `json:"cartes_bancaires,omitempty"`
}

// Contains details on the result for a standalone 3DS Authentication.
type ThreeDSecureAuthenticationOutcomeDetails struct {
	// Universally unique transaction identifier assigned by the issuer to identify the transaction.
	AcsTransactionID string `json:"acs_transaction_id,omitempty"`
	// The Authentication Response Message (ARes) is the issuer's response to the AReq message.
	Ares string `json:"ares,omitempty"`
	// TransStatus field on the ARes
	AresTransStatus ThreeDSecureAuthenticationOutcomeDetailsAresTransStatus `json:"ares_trans_status,omitempty"`
	// A 28-character Base64 string proving that 3DS was completed. Store this value securely, and don't reuse it for multiple authorizations.
	Cryptogram string `json:"cryptogram,omitempty"`
	// The 3DS2 Directory Server Transaction ID.
	DsTransactionID string `json:"ds_transaction_id,omitempty"`
	// Electronic Commerce Indicator provided by the issuer to indicate the result of this 3DS Authentication.
	Eci string `json:"eci,omitempty"`
	// Contains details specific to the individual network.
	NetworkDetails *ThreeDSecureAuthenticationOutcomeDetailsNetworkDetails `json:"network_details,omitempty"`
	// The 3DS protocol version used for this 3DS Authentication.
	ProtocolVersion ThreeDSecureAuthenticationOutcomeDetailsProtocolVersion `json:"protocol_version"`
	// The indicator provided to the issuer by Stripe in the AReq that indicates whether a challenge is requested for this Authentication. This indicator should match the flow_preference you specified but may be overridden (for compliance reasons for example).
	RequestorChallengeIndicator ThreeDSecureAuthenticationOutcomeDetailsRequestorChallengeIndicator `json:"requestor_challenge_indicator,omitempty"`
	// The Results Request Message (RReq) communicates the results of the authentication or verification.
	Rreq string `json:"rreq,omitempty"`
	// TransStatus field on the RReq
	RreqTransStatus ThreeDSecureAuthenticationOutcomeDetailsRreqTransStatus `json:"rreq_trans_status,omitempty"`
	// Universally unique transaction identifier assigned by Stripe to identify the transaction.
	ThreeDsServerTransactionID string `json:"three_ds_server_transaction_id"`
}

// The Standalone 3DS API allows you to run EMV 3D Secure (3DS) authentication using Stripe while authorizing the payment with any PSP.
//
// Related guide: [Standalone 3DS](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure)
type ThreeDSecureAuthentication struct {
	APIResource
	// Contains additional details about the acquirer for a 3DS Authentication.
	AcquirerDetails *ThreeDSecureAuthenticationAcquirerDetails `json:"acquirer_details,omitempty"`
	// The amount for this 3DS Authentication.
	Amount int64 `json:"amount,omitempty"`
	// The URL for presenting a challenge to your cardholder, present if status is requires_challenge.
	ChallengeURL string `json:"challenge_url,omitempty"`
	// Contains details on the channel used (browser, 3RI) for a standalone 3DS Authentication.
	Channel *ThreeDSecureAuthenticationChannel `json:"channel"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// The 3DS directory server with which this 3DS Authentication was processed.
	DirectoryServer ThreeDSecureAuthenticationDirectoryServer `json:"directory_server"`
	// The URL for performing issuer fingerprinting, present if fingerprinting is supported for the given payment method.
	FingerprintingURL string `json:"fingerprinting_url,omitempty"`
	// Contains details of the flow preference used for a standalone 3DS Authentication.
	FlowPreference *ThreeDSecureAuthenticationFlowPreference `json:"flow_preference,omitempty"`
	// Contains information about the future authorisations related to this authentication
	FutureUsage *ThreeDSecureAuthenticationFutureUsage `json:"future_usage,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Indicates whether this 3DS Authentication is being performed for a payment or non-payment use case.
	MessageCategory ThreeDSecureAuthenticationMessageCategory `json:"message_category"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The outcome of this 3DS Authentication.
	Outcome ThreeDSecureAuthenticationOutcome `json:"outcome,omitempty"`
	// Contains details on the result for a standalone 3DS Authentication.
	OutcomeDetails *ThreeDSecureAuthenticationOutcomeDetails `json:"outcome_details,omitempty"`
	// ID of the payment method (a PaymentMethod object) to attach to this 3DS Authentication.
	PaymentMethod *PaymentMethod `json:"payment_method"`
	// The reason for invoking this 3DS Authentication.
	Reason ThreeDSecureAuthenticationReason `json:"reason,omitempty"`
	// Contains details about the shipping address for a 3DS Authentication.
	ShippingAddress *Address `json:"shipping_address,omitempty"`
	// Status of this Authentication.
	Status ThreeDSecureAuthenticationStatus `json:"status"`
}

// ThreeDSecureAuthenticationList is a list of Authentications as retrieved from a list endpoint.
type ThreeDSecureAuthenticationList struct {
	APIResource
	ListMeta
	Data []*ThreeDSecureAuthentication `json:"data"`
}
