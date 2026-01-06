//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. Speed of the payout.
type V2MoneyManagementOutboundPaymentDeliveryOptionsSpeed string

// List of values that V2MoneyManagementOutboundPaymentDeliveryOptionsSpeed can take
const (
	V2MoneyManagementOutboundPaymentDeliveryOptionsSpeedInstant         V2MoneyManagementOutboundPaymentDeliveryOptionsSpeed = "instant"
	V2MoneyManagementOutboundPaymentDeliveryOptionsSpeedNextBusinessDay V2MoneyManagementOutboundPaymentDeliveryOptionsSpeed = "next_business_day"
	V2MoneyManagementOutboundPaymentDeliveryOptionsSpeedStandard        V2MoneyManagementOutboundPaymentDeliveryOptionsSpeed = "standard"
)

// Open Enum. Method for bank account.
type V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount string

// List of values that V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount can take
const (
	V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccountAutomatic V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount = "automatic"
	V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccountLocal     V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount = "local"
	V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccountWire      V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount = "wire"
)

// Open Enum. Shipping speed of the paper check.
type V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheckShippingSpeed string

// List of values that V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheckShippingSpeed can take
const (
	V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheckShippingSpeedPriority V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheckShippingSpeed = "priority"
	V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheckShippingSpeedStandard V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheckShippingSpeed = "standard"
)

// Closed Enum. Configuration option to enable or disable notifications to recipients.
// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
type V2MoneyManagementOutboundPaymentRecipientNotificationSetting string

// List of values that V2MoneyManagementOutboundPaymentRecipientNotificationSetting can take
const (
	V2MoneyManagementOutboundPaymentRecipientNotificationSettingConfigured V2MoneyManagementOutboundPaymentRecipientNotificationSetting = "configured"
	V2MoneyManagementOutboundPaymentRecipientNotificationSettingNone       V2MoneyManagementOutboundPaymentRecipientNotificationSetting = "none"
)

// Closed Enum. Current status of the OutboundPayment: `processing`, `failed`, `posted`, `returned`, `canceled`.
// An OutboundPayment is `processing` if it has been created and is processing.
// The status changes to `posted` once the OutboundPayment has been "confirmed" and funds have left the account, or to `failed` or `canceled`.
// If an OutboundPayment fails to arrive at its payout method, its status will change to `returned`.
type V2MoneyManagementOutboundPaymentStatus string

// List of values that V2MoneyManagementOutboundPaymentStatus can take
const (
	V2MoneyManagementOutboundPaymentStatusCanceled   V2MoneyManagementOutboundPaymentStatus = "canceled"
	V2MoneyManagementOutboundPaymentStatusFailed     V2MoneyManagementOutboundPaymentStatus = "failed"
	V2MoneyManagementOutboundPaymentStatusPosted     V2MoneyManagementOutboundPaymentStatus = "posted"
	V2MoneyManagementOutboundPaymentStatusProcessing V2MoneyManagementOutboundPaymentStatus = "processing"
	V2MoneyManagementOutboundPaymentStatusReturned   V2MoneyManagementOutboundPaymentStatus = "returned"
)

// Open Enum. The `failed` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsFailedReason string

// List of values that V2MoneyManagementOutboundPaymentStatusDetailsFailedReason can take
const (
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodDeclined                    V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_declined"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodDoesNotExist                V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_does_not_exist"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodExpired                     V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_expired"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodUnsupported                 V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_unsupported"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonPayoutMethodUsageFrequencyLimitExceeded V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "payout_method_usage_frequency_limit_exceeded"
	V2MoneyManagementOutboundPaymentStatusDetailsFailedReasonUnknownFailure                          V2MoneyManagementOutboundPaymentStatusDetailsFailedReason = "unknown_failure"
)

// Open Enum. The `returned` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason string

// List of values that V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason can take
const (
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodCanceledByCustomer     V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_canceled_by_customer"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodClosed                 V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_closed"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodCurrencyUnsupported    V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_currency_unsupported"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodDoesNotExist           V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_does_not_exist"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodHolderAddressIncorrect V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_holder_address_incorrect"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodHolderDetailsIncorrect V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_holder_details_incorrect"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodHolderNameIncorrect    V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_holder_name_incorrect"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodInvalidAccountNumber   V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_invalid_account_number"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonPayoutMethodRestricted             V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "payout_method_restricted"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonRecalled                           V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "recalled"
	V2MoneyManagementOutboundPaymentStatusDetailsReturnedReasonUnknownFailure                     V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason = "unknown_failure"
)

// Possible values are `pending`, `supported`, and `unsupported`. Initially set to `pending`, it changes to
// `supported` when the recipient bank provides a trace ID, or `unsupported` if the recipient bank doesn't support it.
// Note that this status may not align with the OutboundPayment or OutboundTransfer status and can remain `pending`
// even after the payment or transfer is posted.
type V2MoneyManagementOutboundPaymentTraceIDStatus string

// List of values that V2MoneyManagementOutboundPaymentTraceIDStatus can take
const (
	V2MoneyManagementOutboundPaymentTraceIDStatusPending     V2MoneyManagementOutboundPaymentTraceIDStatus = "pending"
	V2MoneyManagementOutboundPaymentTraceIDStatusSupported   V2MoneyManagementOutboundPaymentTraceIDStatus = "supported"
	V2MoneyManagementOutboundPaymentTraceIDStatusUnsupported V2MoneyManagementOutboundPaymentTraceIDStatus = "unsupported"
)

// Open Enum. Carrier of the paper check.
type V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckCarrier string

// List of values that V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckCarrier can take
const (
	V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckCarrierFedEx V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckCarrier = "fedex"
	V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckCarrierUSPS  V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckCarrier = "usps"
)

// Open Enum. Tracking status of the paper check.
type V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckTrackingStatus string

// List of values that V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckTrackingStatus can take
const (
	V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckTrackingStatusDelivered V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckTrackingStatus = "delivered"
	V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckTrackingStatusInTransit V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckTrackingStatus = "in_transit"
	V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckTrackingStatusMailed    V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckTrackingStatus = "mailed"
)

// The "presentment amount" for the OutboundPayment.
type V2MoneyManagementOutboundPaymentAmount struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// Delivery options for paper check.
type V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheck struct {
	// Memo printed on the memo field of the check.
	Memo string `json:"memo,omitempty"`
	// Open Enum. Shipping speed of the paper check.
	ShippingSpeed V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheckShippingSpeed `json:"shipping_speed"`
	// Signature for the paper check.
	Signature string `json:"signature"`
}

// Delivery options to be used to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentDeliveryOptions struct {
	// Open Enum. Method for bank account.
	BankAccount V2MoneyManagementOutboundPaymentDeliveryOptionsBankAccount `json:"bank_account,omitempty"`
	// Delivery options for paper check.
	PaperCheck *V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheck `json:"paper_check,omitempty"`
	// Open Enum. Speed of the payout.
	Speed V2MoneyManagementOutboundPaymentDeliveryOptionsSpeed `json:"speed,omitempty"`
}

// The monetary amount debited from the sender, only set on responses.
type V2MoneyManagementOutboundPaymentFromDebited struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// The FinancialAccount that funds were pulled from.
type V2MoneyManagementOutboundPaymentFrom struct {
	// The monetary amount debited from the sender, only set on responses.
	Debited *V2MoneyManagementOutboundPaymentFromDebited `json:"debited"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount string `json:"financial_account"`
}

// Details about the OutboundPayment notification settings for recipient.
type V2MoneyManagementOutboundPaymentRecipientNotification struct {
	// Closed Enum. Configuration option to enable or disable notifications to recipients.
	// Do not send notifications when setting is NONE. Default to account setting when setting is CONFIGURED or not set.
	Setting V2MoneyManagementOutboundPaymentRecipientNotificationSetting `json:"setting"`
}

// The `failed` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsFailed struct {
	// Open Enum. The `failed` status reason.
	Reason V2MoneyManagementOutboundPaymentStatusDetailsFailedReason `json:"reason"`
}

// The `returned` status reason.
type V2MoneyManagementOutboundPaymentStatusDetailsReturned struct {
	// Open Enum. The `returned` status reason.
	Reason V2MoneyManagementOutboundPaymentStatusDetailsReturnedReason `json:"reason"`
}

// Status details for an OutboundPayment in a `failed` or `returned` state.
type V2MoneyManagementOutboundPaymentStatusDetails struct {
	// The `failed` status reason.
	Failed *V2MoneyManagementOutboundPaymentStatusDetailsFailed `json:"failed,omitempty"`
	// The `returned` status reason.
	Returned *V2MoneyManagementOutboundPaymentStatusDetailsReturned `json:"returned,omitempty"`
}

// Hash containing timestamps of when the object transitioned to a particular status.
type V2MoneyManagementOutboundPaymentStatusTransitions struct {
	// Timestamp describing when an OutboundPayment changed status to `canceled`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Timestamp describing when an OutboundPayment changed status to `failed`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	FailedAt time.Time `json:"failed_at,omitempty"`
	// Timestamp describing when an OutboundPayment changed status to `posted`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	PostedAt time.Time `json:"posted_at,omitempty"`
	// Timestamp describing when an OutboundPayment changed status to `returned`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ReturnedAt time.Time `json:"returned_at,omitempty"`
}

// The monetary amount being credited to the destination.
type V2MoneyManagementOutboundPaymentToCredited struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// To which payout method the OutboundPayment was sent.
type V2MoneyManagementOutboundPaymentTo struct {
	// The monetary amount being credited to the destination.
	Credited *V2MoneyManagementOutboundPaymentToCredited `json:"credited"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod string `json:"payout_method"`
	// To which account the OutboundPayment is sent.
	Recipient string `json:"recipient"`
}

// A unique identifier that can be used to track this OutboundPayment with recipient bank. Banks might call this a “reference number” or something similar.
type V2MoneyManagementOutboundPaymentTraceID struct {
	// Possible values are `pending`, `supported`, and `unsupported`. Initially set to `pending`, it changes to
	// `supported` when the recipient bank provides a trace ID, or `unsupported` if the recipient bank doesn't support it.
	// Note that this status may not align with the OutboundPayment or OutboundTransfer status and can remain `pending`
	// even after the payment or transfer is posted.
	Status V2MoneyManagementOutboundPaymentTraceIDStatus `json:"status"`
	// The trace ID value if `trace_id.status` is `supported`, otherwise empty.
	Value string `json:"value,omitempty"`
}

// Mailing address of the paper check.
type V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckMailingAddress struct {
	// City, district, suburb, town, or village.
	City string `json:"city,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 string `json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 string `json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode string `json:"postal_code,omitempty"`
	// State, county, province, or region.
	State string `json:"state,omitempty"`
	// Town or district.
	Town string `json:"town,omitempty"`
}

// Paper check tracking details.
type V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheck struct {
	// Open Enum. Carrier of the paper check.
	Carrier V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckCarrier `json:"carrier"`
	// Check number.
	CheckNumber string `json:"check_number"`
	// Postal code of the latest tracking update.
	CurrentPostalCode string `json:"current_postal_code"`
	// Mailing address of the paper check.
	MailingAddress *V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckMailingAddress `json:"mailing_address"`
	// Tracking number for the check.
	TrackingNumber string `json:"tracking_number"`
	// Open Enum. Tracking status of the paper check.
	TrackingStatus V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckTrackingStatus `json:"tracking_status"`
	// When the tracking details were last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

// Information to track this OutboundPayment with the recipient bank.
type V2MoneyManagementOutboundPaymentTrackingDetails struct {
	// Paper check tracking details.
	PaperCheck *V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheck `json:"paper_check,omitempty"`
}

// OutboundPayment represents a single money movement from one FinancialAccount you own to a payout method someone else owns.
type V2MoneyManagementOutboundPayment struct {
	APIResource
	// The "presentment amount" for the OutboundPayment.
	Amount *V2MoneyManagementOutboundPaymentAmount `json:"amount"`
	// Returns true if the OutboundPayment can be canceled, and false otherwise.
	Cancelable bool `json:"cancelable"`
	// Time at which the OutboundPayment was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Delivery options to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentDeliveryOptions `json:"delivery_options,omitempty"`
	// An arbitrary string attached to the OutboundPayment. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// The date when funds are expected to arrive in the payout method. This field is not set if the payout method is in a `failed`, `canceled`, or `returned` state.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ExpectedArrivalDate time.Time `json:"expected_arrival_date,omitempty"`
	// The FinancialAccount that funds were pulled from.
	From *V2MoneyManagementOutboundPaymentFrom `json:"from"`
	// Unique identifier for the OutboundPayment.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The quote for this OutboundPayment. Only required for countries with regulatory mandates to display fee estimates before OutboundPayment creation.
	OutboundPaymentQuote string `json:"outbound_payment_quote,omitempty"`
	// A link to the Stripe-hosted receipt for this OutboundPayment. The receipt link remains active for 60 days from the OutboundPayment creation date. After this period, the link will expire and the receipt url value will be null.
	ReceiptURL string `json:"receipt_url,omitempty"`
	// Details about the OutboundPayment notification settings for recipient.
	RecipientNotification *V2MoneyManagementOutboundPaymentRecipientNotification `json:"recipient_notification"`
	// The recipient verification id for this OutboundPayment. Only required for countries with regulatory mandates to verify recipient names before OutboundPayment creation.
	RecipientVerification string `json:"recipient_verification,omitempty"`
	// The description that appears on the receiving end for an OutboundPayment (for example, bank statement for external bank transfer). It will default to `STRIPE` if not set on the account settings.
	StatementDescriptor string `json:"statement_descriptor"`
	// Closed Enum. Current status of the OutboundPayment: `processing`, `failed`, `posted`, `returned`, `canceled`.
	// An OutboundPayment is `processing` if it has been created and is processing.
	// The status changes to `posted` once the OutboundPayment has been "confirmed" and funds have left the account, or to `failed` or `canceled`.
	// If an OutboundPayment fails to arrive at its payout method, its status will change to `returned`.
	Status V2MoneyManagementOutboundPaymentStatus `json:"status"`
	// Status details for an OutboundPayment in a `failed` or `returned` state.
	StatusDetails *V2MoneyManagementOutboundPaymentStatusDetails `json:"status_details,omitempty"`
	// Hash containing timestamps of when the object transitioned to a particular status.
	StatusTransitions *V2MoneyManagementOutboundPaymentStatusTransitions `json:"status_transitions,omitempty"`
	// To which payout method the OutboundPayment was sent.
	To *V2MoneyManagementOutboundPaymentTo `json:"to"`
	// A unique identifier that can be used to track this OutboundPayment with recipient bank. Banks might call this a “reference number” or something similar.
	TraceID *V2MoneyManagementOutboundPaymentTraceID `json:"trace_id"`
	// Information to track this OutboundPayment with the recipient bank.
	TrackingDetails *V2MoneyManagementOutboundPaymentTrackingDetails `json:"tracking_details,omitempty"`
}
