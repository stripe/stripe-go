//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Assessments from Stripe. If set, the value is `fraudulent`.
type ChargeFraudStripeReport string

// List of values that ChargeFraudStripeReport can take
const (
	ChargeFraudStripeReportFraudulent ChargeFraudStripeReport = "fraudulent"
)

// Assessments reported by you. If set, possible values of are `safe` and `fraudulent`.
type ChargeFraudUserReport string

// List of values that ChargeFraudUserReport can take
const (
	ChargeFraudUserReportFraudulent ChargeFraudUserReport = "fraudulent"
	ChargeFraudUserReportSafe       ChargeFraudUserReport = "safe"
)

// For authenticated transactions: how the customer was authenticated by
// the issuing bank.
type ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// Indicates the outcome of 3D Secure authentication.
type ChargePaymentMethodDetailsCardThreeDSecureResult string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResult can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged ChargePaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	ChargePaymentMethodDetailsCardThreeDSecureResultAuthenticated       ChargePaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	ChargePaymentMethodDetailsCardThreeDSecureResultExempted            ChargePaymentMethodDetailsCardThreeDSecureResult = "exempted"
	ChargePaymentMethodDetailsCardThreeDSecureResultFailed              ChargePaymentMethodDetailsCardThreeDSecureResult = "failed"
	ChargePaymentMethodDetailsCardThreeDSecureResultNotSupported        ChargePaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultProcessingError     ChargePaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// Additional information about why 3D Secure succeeded or failed based
// on the `result`.
type ChargePaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResultReason can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           ChargePaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonBypassed            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCanceled            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     ChargePaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported ChargePaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       ChargePaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonRejected            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

// The type of account being debited or credited
type ChargePaymentMethodDetailsCardPresentReceiptAccountType string

// List of values that ChargePaymentMethodDetailsCardPresentReceiptAccountType can take
const (
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeChecking ChargePaymentMethodDetailsCardPresentReceiptAccountType = "checking"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeCredit   ChargePaymentMethodDetailsCardPresentReceiptAccountType = "credit"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypePrepaid  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "prepaid"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeUnknown  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "unknown"
)

// The Klarna payment method used for this transaction.
// Can be one of `pay_later`, `pay_now`, `pay_with_financing`, or `pay_in_installments`
type ChargePaymentMethodDetailsKlarnaPaymentMethodCategory string

// List of values that ChargePaymentMethodDetailsKlarnaPaymentMethodCategory can take
const (
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayLater          ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_later"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayNow            ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_now"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayWithFinancing  ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_with_financing"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayInInstallments ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_in_installments"
)

// The name of the convenience store chain where the payment was completed.
type ChargePaymentMethodDetailsKonbiniStoreChain string

// List of values that ChargePaymentMethodDetailsKonbiniStoreChain can take
const (
	ChargePaymentMethodDetailsKonbiniStoreChainFamilyMart ChargePaymentMethodDetailsKonbiniStoreChain = "familymart"
	ChargePaymentMethodDetailsKonbiniStoreChainLawson     ChargePaymentMethodDetailsKonbiniStoreChain = "lawson"
	ChargePaymentMethodDetailsKonbiniStoreChainMinistop   ChargePaymentMethodDetailsKonbiniStoreChain = "ministop"
	ChargePaymentMethodDetailsKonbiniStoreChainSeicomart  ChargePaymentMethodDetailsKonbiniStoreChain = "seicomart"
)

// The type of transaction-specific details of the payment method used in the payment, one of `ach_credit_transfer`, `ach_debit`, `acss_debit`, `alipay`, `au_becs_debit`, `bancontact`, `card`, `card_present`, `eps`, `giropay`, `ideal`, `klarna`, `multibanco`, `p24`, `sepa_debit`, `sofort`, `stripe_account`, or `wechat`.
// An additional hash is included on `payment_method_details` with a name matching this value.
// It contains information specific to the payment method.
type ChargePaymentMethodDetailsType string

// List of values that ChargePaymentMethodDetailsType can take
const (
	ChargePaymentMethodDetailsTypeAchCreditTransfer ChargePaymentMethodDetailsType = "ach_credit_transfer"
	ChargePaymentMethodDetailsTypeAchDebit          ChargePaymentMethodDetailsType = "ach_debit"
	ChargePaymentMethodDetailsTypeAcssDebit         ChargePaymentMethodDetailsType = "acss_debit"
	ChargePaymentMethodDetailsTypeAlipay            ChargePaymentMethodDetailsType = "alipay"
	ChargePaymentMethodDetailsTypeAUBECSDebit       ChargePaymentMethodDetailsType = "au_becs_debit"
	ChargePaymentMethodDetailsTypeBACSDebit         ChargePaymentMethodDetailsType = "bacs_debit"
	ChargePaymentMethodDetailsTypeBancontact        ChargePaymentMethodDetailsType = "bancontact"
	ChargePaymentMethodDetailsTypeCard              ChargePaymentMethodDetailsType = "card"
	ChargePaymentMethodDetailsTypeCardPresent       ChargePaymentMethodDetailsType = "card_present"
	ChargePaymentMethodDetailsTypeEps               ChargePaymentMethodDetailsType = "eps"
	ChargePaymentMethodDetailsTypeFPX               ChargePaymentMethodDetailsType = "fpx"
	ChargePaymentMethodDetailsTypeGiropay           ChargePaymentMethodDetailsType = "giropay"
	ChargePaymentMethodDetailsTypeGrabpay           ChargePaymentMethodDetailsType = "grabpay"
	ChargePaymentMethodDetailsTypeIdeal             ChargePaymentMethodDetailsType = "ideal"
	ChargePaymentMethodDetailsTypeInteracPresent    ChargePaymentMethodDetailsType = "interac_present"
	ChargePaymentMethodDetailsTypeKlarna            ChargePaymentMethodDetailsType = "klarna"
	ChargePaymentMethodDetailsTypeMultibanco        ChargePaymentMethodDetailsType = "multibanco"
	ChargePaymentMethodDetailsTypeP24               ChargePaymentMethodDetailsType = "p24"
	ChargePaymentMethodDetailsTypeSepaDebit         ChargePaymentMethodDetailsType = "sepa_debit"
	ChargePaymentMethodDetailsTypeSofort            ChargePaymentMethodDetailsType = "sofort"
	ChargePaymentMethodDetailsTypeStripeAccount     ChargePaymentMethodDetailsType = "stripe_account"
	ChargePaymentMethodDetailsTypeWechat            ChargePaymentMethodDetailsType = "wechat"
)

// Account holder type: individual or company.
type ChargePaymentMethodDetailsUSBankAccountAccountHolderType string

// List of values that ChargePaymentMethodDetailsUSBankAccountAccountHolderType can take
const (
	ChargePaymentMethodDetailsUSBankAccountAccountHolderTypeCompany    ChargePaymentMethodDetailsUSBankAccountAccountHolderType = "company"
	ChargePaymentMethodDetailsUSBankAccountAccountHolderTypeIndividual ChargePaymentMethodDetailsUSBankAccountAccountHolderType = "individual"
)

// Account type: checkings or savings. Defaults to checking if omitted.
type ChargePaymentMethodDetailsUSBankAccountAccountType string

// List of values that ChargePaymentMethodDetailsUSBankAccountAccountType can take
const (
	ChargePaymentMethodDetailsUSBankAccountAccountTypeChecking ChargePaymentMethodDetailsUSBankAccountAccountType = "checking"
	ChargePaymentMethodDetailsUSBankAccountAccountTypeSavings  ChargePaymentMethodDetailsUSBankAccountAccountType = "savings"
)

// The status of the payment is either `succeeded`, `pending`, or `failed`.
type ChargeStatus string

// List of values that ChargeStatus can take
const (
	ChargeStatusFailed    ChargeStatus = "failed"
	ChargeStatusPending   ChargeStatus = "pending"
	ChargeStatusSucceeded ChargeStatus = "succeeded"
)

// Search for charges you've previously created using Stripe's [Search Query Language](https://stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
type ChargeSearchParams struct {
	SearchParams `form:"*"`
	// A cursor for pagination across multiple pages of results. Don't include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
	Page *string `form:"page"`
}

// Returns a list of charges you've previously created. The charges are returned in sorted order, with the most recent charges appearing first.
type ChargeListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return charges for the customer specified by this customer ID.
	Customer *string `form:"customer"`
	// Only return charges that were created by the PaymentIntent specified by this PaymentIntent ID.
	PaymentIntent *string `form:"payment_intent"`
	// Only return charges for this transfer group.
	TransferGroup *string `form:"transfer_group"`
}
type DestinationParams struct {
	// ID of an existing, connected Stripe account.
	Account *string `form:"account"`
	// The amount to transfer to the destination account without creating an `Application Fee` object. Cannot be combined with the `application_fee` parameter. Must be less than or equal to the charge amount.
	Amount *int64 `form:"amount"`
}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type ChargeRadarOptionsParams struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://stripe.com/docs/connect/destination-charges) for details.
type ChargeTransferDataParams struct {
	// The amount transferred to the destination account, if specified. By default, the entire charge amount is transferred to the destination account.
	Amount *int64 `form:"amount"`
	// This parameter can only be used on Charge creation.
	// ID of an existing, connected Stripe account.
	Destination *string `form:"destination"`
}
type ChargeLevel3LineItemsParams struct {
	DiscountAmount     *int64  `form:"discount_amount"`
	ProductCode        *string `form:"product_code"`
	ProductDescription *string `form:"product_description"`
	Quantity           *int64  `form:"quantity"`
	TaxAmount          *int64  `form:"tax_amount"`
	UnitCost           *int64  `form:"unit_cost"`
}
type ChargeLevel3Params struct {
	CustomerReference  *string                        `form:"customer_reference"`
	LineItems          []*ChargeLevel3LineItemsParams `form:"line_items"`
	MerchantReference  *string                        `form:"merchant_reference"`
	ShippingAddressZip *string                        `form:"shipping_address_zip"`
	ShippingAmount     *int64                         `form:"shipping_amount"`
	ShippingFromZip    *string                        `form:"shipping_from_zip"`
}

// To charge a credit card or other payment source, you create a Charge object. If your API key is in test mode, the supplied payment source (e.g., card) won't actually be charged, although everything else will occur as if in live mode. (Stripe assumes that the charge would have completed successfully).
type ChargeParams struct {
	Params `form:"*"`
	// Amount intended to be collected by this payment. A positive integer representing how much to charge in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://stripe.com/docs/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount         *int64 `form:"amount"`
	ApplicationFee *int64 `form:"application_fee"`
	// A fee in cents (or local equivalent) that will be applied to the charge and transferred to the application owner's Stripe account. The request must be made with an OAuth key or the `Stripe-Account` header in order to take an application fee. For more information, see the application fees [documentation](https://stripe.com/docs/connect/direct-charges#collecting-fees).
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// Whether to immediately capture the charge. Defaults to `true`. When `false`, the charge issues an authorization (or pre-authorization), and will need to be [captured](https://stripe.com/docs/api#capture_charge) later. Uncaptured charges expire after a set number of days (7 by default). For more information, see the [authorizing charges and settling later](https://stripe.com/docs/charges/placing-a-hold) documentation.
	Capture *bool `form:"capture"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of an existing customer that will be associated with this request. This field may only be updated if there is no existing associated customer with this charge.
	Customer *string `form:"customer"`
	// An arbitrary string which you can attach to a charge object. It is displayed when in the web interface alongside the charge. Note that if you use Stripe to send automatic email receipts to your customers, your receipt emails will include the `description` of the charge(s) that they are describing.
	Description  *string            `form:"description"`
	Destination  *DestinationParams `form:"destination"`
	ExchangeRate *float64           `form:"exchange_rate"`
	// A set of key-value pairs you can attach to a charge giving information about its riskiness. If you believe a charge is fraudulent, include a `user_report` key with a value of `fraudulent`. If you believe a charge is safe, include a `user_report` key with a value of `safe`. Stripe will use the information you send to improve our fraud detection algorithms.
	FraudDetails *FraudDetailsParams `form:"fraud_details"`
	Level3       *ChargeLevel3Params `form:"level3"`
	// The Stripe account ID for which these funds are intended. Automatically set if you use the `destination` parameter. For details, see [Creating Separate Charges and Transfers](https://stripe.com/docs/connect/charges-transfers#on-behalf-of).
	OnBehalfOf *string `form:"on_behalf_of"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *ChargeRadarOptionsParams `form:"radar_options"`
	// This is the email address that the receipt for this charge will be sent to. If this field is updated, then a new email receipt will be sent to the updated address.
	ReceiptEmail *string `form:"receipt_email"`
	// Shipping information for the charge. Helps prevent fraud on charges for physical goods.
	Shipping *ShippingDetailsParams `form:"shipping"`
	Source   *SourceParams          `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	// For card charges, use `statement_descriptor_suffix` instead. Otherwise, you can use this value as the complete description of a charge on your customers' statements. Must contain at least one letter, maximum 22 characters.
	StatementDescriptor *string `form:"statement_descriptor"`
	// Provides information about the charge that customers see on their statements. Concatenated with the prefix (shortened descriptor) or statement descriptor that's set on the account to form the complete statement descriptor. Maximum 22 characters for the concatenated descriptor.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix"`
	// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://stripe.com/docs/connect/destination-charges) for details.
	TransferData *ChargeTransferDataParams `form:"transfer_data"`
	// A string that identifies this transaction as part of a group. `transfer_group` may only be provided if it has not been set. See the [Connect documentation](https://stripe.com/docs/connect/charges-transfers#transfer-options) for details.
	TransferGroup *string `form:"transfer_group"`
}

// SetSource adds valid sources to a ChargeParams object,
// returning an error for unsupported sources.
func (p *ChargeParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// A set of key-value pairs you can attach to a charge giving information about its riskiness. If you believe a charge is fraudulent, include a `user_report` key with a value of `fraudulent`. If you believe a charge is safe, include a `user_report` key with a value of `safe`. Stripe will use the information you send to improve our fraud detection algorithms.
type FraudDetailsParams struct {
	// Either `safe` or `fraudulent`.
	UserReport *string `form:"user_report"`
}

// Capture the payment of an existing, uncaptured, charge. This is the second half of the two-step payment flow, where first you [created a charge](https://stripe.com/docs/api#create_charge) with the capture option set to false.
//
// Uncaptured payments expire a set number of days after they are created ([7 by default](https://stripe.com/docs/charges/placing-a-hold)). If they are not captured by that point in time, they will be marked as refunded and will no longer be capturable.
type CaptureParams struct {
	Params `form:"*"`
	// The amount to capture, which must be less than or equal to the original amount. Any additional amount will be automatically refunded.
	Amount *int64 `form:"amount"`
	// An application fee to add on to this charge.
	ApplicationFee *int64 `form:"application_fee"`
	// An application fee amount to add on to this charge, which must be less than or equal to the original amount.
	ApplicationFeeAmount *int64   `form:"application_fee_amount"`
	ExchangeRate         *float64 `form:"exchange_rate"`
	// The email address to send this charge's receipt to. This will override the previously-specified email address for this charge, if one was set. Receipts will not be sent in test mode.
	ReceiptEmail *string `form:"receipt_email"`
	// For card charges, use `statement_descriptor_suffix` instead. Otherwise, you can use this value as the complete description of a charge on your customers' statements. Must contain at least one letter, maximum 22 characters.
	StatementDescriptor *string `form:"statement_descriptor"`
	// Provides information about the charge that customers see on their statements. Concatenated with the prefix (shortened descriptor) or statement descriptor that's set on the account to form the complete statement descriptor. Maximum 22 characters for the concatenated descriptor.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix"`
	// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://stripe.com/docs/connect/destination-charges) for details.
	TransferData *ChargeTransferDataParams `form:"transfer_data"`
	// A string that identifies this transaction as part of a group. `transfer_group` may only be provided if it has not been set. See the [Connect documentation](https://stripe.com/docs/connect/charges-transfers#transfer-options) for details.
	TransferGroup *string `form:"transfer_group"`
}

// Information on fraud assessments for the charge.
type FraudDetails struct {
	// Assessments from Stripe. If set, the value is `fraudulent`.
	StripeReport ChargeFraudStripeReport `json:"stripe_report"`
	// Assessments reported by you. If set, possible values of are `safe` and `fraudulent`.
	UserReport ChargeFraudUserReport `json:"user_report"`
}
type ChargeLevel3LineItem struct {
	DiscountAmount     int64  `json:"discount_amount"`
	ProductCode        string `json:"product_code"`
	ProductDescription string `json:"product_description"`
	Quantity           int64  `json:"quantity"`
	TaxAmount          int64  `json:"tax_amount"`
	UnitCost           int64  `json:"unit_cost"`
}
type ChargeLevel3 struct {
	CustomerReference  string                  `json:"customer_reference"`
	LineItems          []*ChargeLevel3LineItem `json:"line_items"`
	MerchantReference  string                  `json:"merchant_reference"`
	ShippingAddressZip string                  `json:"shipping_address_zip"`
	ShippingAmount     int64                   `json:"shipping_amount"`
	ShippingFromZip    string                  `json:"shipping_from_zip"`
}

// The ID of the Radar rule that matched the payment, if applicable.
type ChargeOutcomeRule struct {
	// The action taken on the payment.
	Action string `json:"action"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The predicate to evaluate the payment against.
	Predicate string `json:"predicate"`
}

// Details about whether the payment was accepted, and why. See [understanding declines](https://stripe.com/docs/declines) for details.
type ChargeOutcome struct {
	// Possible values are `approved_by_network`, `declined_by_network`, `not_sent_to_network`, and `reversed_after_approval`. The value `reversed_after_approval` indicates the payment was [blocked by Stripe](https://stripe.com/docs/declines#blocked-payments) after bank authorization, and may temporarily appear as "pending" on a cardholder's statement.
	NetworkStatus string `json:"network_status"`
	// An enumerated value providing a more detailed explanation of the outcome's `type`. Charges blocked by Radar's default block rule have the value `highest_risk_level`. Charges placed in review by Radar's default review rule have the value `elevated_risk_level`. Charges authorized, blocked, or placed in review by custom rules have the value `rule`. See [understanding declines](https://stripe.com/docs/declines) for more details.
	Reason string `json:"reason"`
	// Stripe Radar's evaluation of the riskiness of the payment. Possible values for evaluated payments are `normal`, `elevated`, `highest`. For non-card payments, and card-based payments predating the public assignment of risk levels, this field will have the value `not_assessed`. In the event of an error in the evaluation, this field will have the value `unknown`. This field is only available with Radar.
	RiskLevel string `json:"risk_level"`
	// Stripe Radar's evaluation of the riskiness of the payment. Possible values for evaluated payments are between 0 and 100. For non-card payments, card-based payments predating the public assignment of risk scores, or in the event of an error during evaluation, this field will not be present. This field is only available with Radar for Fraud Teams.
	RiskScore int64 `json:"risk_score"`
	// The ID of the Radar rule that matched the payment, if applicable.
	Rule *ChargeOutcomeRule `json:"rule"`
	// A human-readable description of the outcome type and reason, designed for you (the recipient of the payment), not your customer.
	SellerMessage string `json:"seller_message"`
	// Possible values are `authorized`, `manual_review`, `issuer_declined`, `blocked`, and `invalid`. See [understanding declines](https://stripe.com/docs/declines) and [Radar reviews](https://stripe.com/docs/radar/reviews) for details.
	Type string `json:"type"`
}

// UnmarshalJSON handles deserialization of a ChargeOutcomeRule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *ChargeOutcomeRule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type chargeOutcomeRule ChargeOutcomeRule
	var v chargeOutcomeRule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = ChargeOutcomeRule(v)
	return nil
}

type ChargePaymentMethodDetailsAchCreditTransfer struct {
	// Account number to transfer funds to.
	AccountNumber string `json:"account_number"`
	// Name of the bank associated with the routing number.
	BankName string `json:"bank_name"`
	// Routing transit number for the bank account to transfer funds to.
	RoutingNumber string `json:"routing_number"`
	// SWIFT code of the bank associated with the routing number.
	SwiftCode string `json:"swift_code"`
}
type ChargePaymentMethodDetailsAchDebit struct {
	AccountHolderType BankAccountAccountHolderType `json:"account_holder_type"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// Routing transit number of the bank account.
	RoutingNumber string `json:"routing_number"`
}
type ChargePaymentMethodDetailsAcssDebit struct {
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Institution number of the bank account
	InstitutionNumber string `json:"institution_number"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate string `json:"mandate"`
	// Transit number of the bank account.
	TransitNumber string `json:"transit_number"`
}
type ChargePaymentMethodDetailsAffirm struct{}
type ChargePaymentMethodDetailsAfterpayClearpay struct {
	// Order identifier shown to the merchant in Afterpay's online portal.
	Reference string `json:"reference"`
}
type ChargePaymentMethodDetailsAlipay struct {
	// Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.
	BuyerID string `json:"buyer_id"`
	// Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Transaction ID of this particular Alipay transaction.
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsAUBECSDebit struct {
	// Bank-State-Branch number of the bank account.
	BSBNumber string `json:"bsb_number"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate string `json:"mandate"`
}
type ChargePaymentMethodDetailsBACSDebit struct {
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate string `json:"mandate"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode string `json:"sort_code"`
}
type ChargePaymentMethodDetailsBancontact struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	Bic string `json:"bic"`
	// The ID of the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSepaDebit *PaymentMethod `json:"generated_sepa_debit"`
	// The mandate for the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSepaDebitMandate *Mandate `json:"generated_sepa_debit_mandate"`
	// Last four characters of the IBAN.
	IbanLast4 string `json:"iban_last4"`
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	// Can be one of `en`, `de`, `fr`, or `nl`
	PreferredLanguage string `json:"preferred_language"`
	// Owner's verified full name. Values are verified or provided by Bancontact directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsBLIK struct{}
type ChargePaymentMethodDetailsBoleto struct {
	// The tax ID of the customer (CPF for individuals consumers or CNPJ for businesses consumers)
	TaxID string `json:"tax_id"`
}

// Check results by Card networks on Card address and CVC at time of payment.
type ChargePaymentMethodDetailsCardChecks struct {
	AddressLine1Check      CardVerification `json:"address_line1_check"`
	AddressPostalCodeCheck CardVerification `json:"address_postal_code_check"`
	CVCCheck               CardVerification `json:"cvc_check"`
}

// Installment details for this payment (Mexico only).
//
// For more information, see the [installments integration guide](https://stripe.com/docs/payments/installments).
type ChargePaymentMethodDetailsCardInstallments struct {
	// Installment plan selected for the payment.
	Plan *PaymentIntentPaymentMethodOptionsCardInstallmentsPlan `json:"plan"`
}

// Populated if this transaction used 3D Secure authentication.
type ChargePaymentMethodDetailsCardThreeDSecure struct {
	// For authenticated transactions: how the customer was authenticated by
	// the issuing bank.
	AuthenticationFlow ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	// Indicates the outcome of 3D Secure authentication.
	Result ChargePaymentMethodDetailsCardThreeDSecureResult `json:"result"`
	// Additional information about why 3D Secure succeeded or failed based
	// on the `result`.
	ResultReason ChargePaymentMethodDetailsCardThreeDSecureResultReason `json:"result_reason"`
	// The version of 3D Secure that was used.
	Version string `json:"version"`
}
type ChargePaymentMethodDetailsCardWalletAmexExpressCheckout struct{}
type ChargePaymentMethodDetailsCardWalletApplePay struct{}
type ChargePaymentMethodDetailsCardWalletGooglePay struct{}
type ChargePaymentMethodDetailsCardWalletMasterpass struct {
	// Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	BillingAddress *Address `json:"billing_address"`
	// Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Email string `json:"email"`
	// Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Name string `json:"name"`
	// Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	ShippingAddress *Address `json:"shipping_address"`
}
type ChargePaymentMethodDetailsCardWalletSamsungPay struct{}
type ChargePaymentMethodDetailsCardWalletVisaCheckout struct {
	// Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	BillingAddress *Address `json:"billing_address"`
	// Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Email string `json:"email"`
	// Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Name string `json:"name"`
	// Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	ShippingAddress *Address `json:"shipping_address"`
}

// If this Card is part of a card wallet, this contains the details of the card wallet.
type ChargePaymentMethodDetailsCardWallet struct {
	AmexExpressCheckout *ChargePaymentMethodDetailsCardWalletAmexExpressCheckout `json:"amex_express_checkout"`
	ApplePay            *ChargePaymentMethodDetailsCardWalletApplePay            `json:"apple_pay"`
	// (For tokenized numbers only.) The last four digits of the device account number.
	DynamicLast4 string                                          `json:"dynamic_last4"`
	GooglePay    *ChargePaymentMethodDetailsCardWalletGooglePay  `json:"google_pay"`
	Masterpass   *ChargePaymentMethodDetailsCardWalletMasterpass `json:"masterpass"`
	SamsungPay   *ChargePaymentMethodDetailsCardWalletSamsungPay `json:"samsung_pay"`
	// The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, or `visa_checkout`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
	Type         PaymentMethodCardWalletType                       `json:"type"`
	VisaCheckout *ChargePaymentMethodDetailsCardWalletVisaCheckout `json:"visa_checkout"`
}
type ChargePaymentMethodDetailsCard struct {
	// Card brand. Can be `amex`, `diners`, `discover`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Brand PaymentMethodCardBrand `json:"brand"`
	// Check results by Card networks on Card address and CVC at time of payment.
	Checks *ChargePaymentMethodDetailsCardChecks `json:"checks"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// Two-digit number representing the card's expiration month.
	ExpMonth uint64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear uint64 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *Starting May 1, 2021, card fingerprint in India for Connect will change to allow two fingerprints for the same card --- one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding CardFunding `json:"funding"`
	// Installment details for this payment (Mexico only).
	//
	// For more information, see the [installments integration guide](https://stripe.com/docs/payments/installments).
	Installments *ChargePaymentMethodDetailsCardInstallments `json:"installments"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment or created by it.
	Mandate string `json:"mandate"`
	// True if this payment was marked as MOTO and out of scope for SCA.
	MOTO bool `json:"moto"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `interac`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network PaymentMethodCardNetwork `json:"network"`
	// Populated if this transaction used 3D Secure authentication.
	ThreeDSecure *ChargePaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
	// If this Card is part of a card wallet, this contains the details of the card wallet.
	Wallet *ChargePaymentMethodDetailsCardWallet `json:"wallet"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}

// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
type ChargePaymentMethodDetailsCardPresentReceipt struct {
	// The type of account being debited or credited
	AccountType ChargePaymentMethodDetailsCardPresentReceiptAccountType `json:"account_type"`
	// EMV tag 9F26, cryptogram generated by the integrated circuit chip.
	ApplicationCryptogram string `json:"application_cryptogram"`
	// Mnenomic of the Application Identifier.
	ApplicationPreferredName string `json:"application_preferred_name"`
	// Identifier for this transaction.
	AuthorizationCode string `json:"authorization_code"`
	// EMV tag 8A. A code returned by the card issuer.
	AuthorizationResponseCode string `json:"authorization_response_code"`
	// How the cardholder verified ownership of the card.
	CardholderVerificationMethod string `json:"cardholder_verification_method"`
	// EMV tag 84. Similar to the application identifier stored on the integrated circuit chip.
	DedicatedFileName string `json:"dedicated_file_name"`
	// The outcome of a series of EMV functions performed by the card reader.
	TerminalVerificationResults string `json:"terminal_verification_results"`
	// An indication of various EMV functions performed during the transaction.
	TransactionStatusInformation string `json:"transaction_status_information"`
}
type ChargePaymentMethodDetailsCardPresent struct {
	// The authorized amount
	AmountAuthorized int64 `json:"amount_authorized"`
	// Card brand. Can be `amex`, `diners`, `discover`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Brand PaymentMethodCardBrand `json:"brand"`
	// When using manual capture, a future timestamp after which the charge will be automatically refunded if uncaptured.
	CaptureBefore int64 `json:"capture_before"`
	// The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.
	CardholderName string `json:"cardholder_name"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// Authorization response cryptogram.
	EmvAuthData string `json:"emv_auth_data"`
	// Two-digit number representing the card's expiration month.
	ExpMonth uint64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear uint64 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *Starting May 1, 2021, card fingerprint in India for Connect will change to allow two fingerprints for the same card --- one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding CardFunding `json:"funding"`
	// ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod.
	GeneratedCard string `json:"generated_card"`
	// Whether this [PaymentIntent](https://stripe.com/docs/api/payment_intents) is eligible for incremental authorizations. Request support using [request_incremental_authorization_support](https://stripe.com/docs/api/payment_intents/create#create_payment_intent-payment_method_options-card_present-request_incremental_authorization_support).
	IncrementalAuthorizationSupported bool `json:"incremental_authorization_supported"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `interac`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network PaymentMethodCardNetwork `json:"network"`
	// Defines whether the authorized amount can be over-captured or not
	OvercaptureSupported bool `json:"overcapture_supported"`
	// How card details were read in this transaction.
	ReadMethod string `json:"read_method"`
	// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
	Receipt *ChargePaymentMethodDetailsCardPresentReceipt `json:"receipt"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}
type ChargePaymentMethodDetailsCustomerBalance struct{}
type ChargePaymentMethodDetailsEps struct {
	// The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.
	Bank string `json:"bank"`
	// Owner's verified full name. Values are verified or provided by EPS directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	// EPS rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsFPX struct {
	// Account holder type, if provided. Can be one of `individual` or `company`.
	AccountHolderType PaymentMethodFPXAccountHolderType `json:"account_holder_type"`
	// The customer's bank. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, or `pb_enterprise`.
	Bank string `json:"bank"`
	// Unique transaction id generated by FPX for every request from the merchant
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsGiropay struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	Bic string `json:"bic"`
	// Owner's verified full name. Values are verified or provided by Giropay directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	// Giropay rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsGrabpay struct {
	// Unique transaction id generated by GrabPay
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsIdeal struct {
	// The customer's bank. Can be one of `abn_amro`, `asn_bank`, `bunq`, `handelsbanken`, `ing`, `knab`, `moneyou`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, or `van_lanschot`.
	Bank string `json:"bank"`
	// The Bank Identifier Code of the customer's bank.
	Bic string `json:"bic"`
	// The ID of the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSepaDebit *PaymentMethod `json:"generated_sepa_debit"`
	// The mandate for the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSepaDebitMandate *Mandate `json:"generated_sepa_debit_mandate"`
	// Last four characters of the IBAN.
	IbanLast4 string `json:"iban_last4"`
	// Owner's verified full name. Values are verified or provided by iDEAL directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}

// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
type ChargePaymentMethodDetailsInteracPresentReceipt struct {
	// The type of account being debited or credited
	AccountType string `json:"account_type"`
	// EMV tag 9F26, cryptogram generated by the integrated circuit chip.
	ApplicationCryptogram string `json:"application_cryptogram"`
	// Mnenomic of the Application Identifier.
	ApplicationPreferredName string `json:"application_preferred_name"`
	// Identifier for this transaction.
	AuthorizationCode string `json:"authorization_code"`
	// EMV tag 8A. A code returned by the card issuer.
	AuthorizationResponseCode string `json:"authorization_response_code"`
	// How the cardholder verified ownership of the card.
	CardholderVerificationMethod string `json:"cardholder_verification_method"`
	// EMV tag 84. Similar to the application identifier stored on the integrated circuit chip.
	DedicatedFileName string `json:"dedicated_file_name"`
	// The outcome of a series of EMV functions performed by the card reader.
	TerminalVerificationResults string `json:"terminal_verification_results"`
	// An indication of various EMV functions performed during the transaction.
	TransactionStatusInformation string `json:"transaction_status_information"`
}
type ChargePaymentMethodDetailsInteracPresent struct {
	// Card brand. Can be `interac`, `mastercard` or `visa`.
	Brand string `json:"brand"`
	// The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.
	CardholderName string `json:"cardholder_name"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// Authorization response cryptogram.
	EmvAuthData string `json:"emv_auth_data"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *Starting May 1, 2021, card fingerprint in India for Connect will change to allow two fingerprints for the same card --- one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding string `json:"funding"`
	// ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod.
	GeneratedCard string `json:"generated_card"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `interac`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network string `json:"network"`
	// EMV tag 5F2D. Preferred languages specified by the integrated circuit chip.
	PreferredLocales []string `json:"preferred_locales"`
	// How card details were read in this transaction.
	ReadMethod string `json:"read_method"`
	// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
	Receipt *ChargePaymentMethodDetailsInteracPresentReceipt `json:"receipt"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}
type ChargePaymentMethodDetailsKlarna struct {
	// The Klarna payment method used for this transaction.
	// Can be one of `pay_later`, `pay_now`, `pay_with_financing`, or `pay_in_installments`
	PaymentMethodCategory ChargePaymentMethodDetailsKlarnaPaymentMethodCategory `json:"payment_method_category"`
	// Preferred language of the Klarna authorization page that the customer is redirected to.
	// Can be one of `de-AT`, `en-AT`, `nl-BE`, `fr-BE`, `en-BE`, `de-DE`, `en-DE`, `da-DK`, `en-DK`, `es-ES`, `en-ES`, `fi-FI`, `sv-FI`, `en-FI`, `en-GB`, `en-IE`, `it-IT`, `en-IT`, `nl-NL`, `en-NL`, `nb-NO`, `en-NO`, `sv-SE`, `en-SE`, `en-US`, `es-US`, `fr-FR`, `en-FR`, `en-AU`, `en-NZ`, `en-CA`, or `fr-CA`
	PreferredLocale string `json:"preferred_locale"`
}

// If the payment succeeded, this contains the details of the convenience store where the payment was completed.
type ChargePaymentMethodDetailsKonbiniStore struct {
	// The name of the convenience store chain where the payment was completed.
	Chain ChargePaymentMethodDetailsKonbiniStoreChain `json:"chain"`
}
type ChargePaymentMethodDetailsKonbini struct {
	// If the payment succeeded, this contains the details of the convenience store where the payment was completed.
	Store *ChargePaymentMethodDetailsKonbiniStore `json:"store"`
}
type ChargePaymentMethodDetailsLink struct{}
type ChargePaymentMethodDetailsMultibanco struct {
	// Entity number associated with this Multibanco payment.
	Entity string `json:"entity"`
	// Reference number associated with this Multibanco payment.
	Reference string `json:"reference"`
}
type ChargePaymentMethodDetailsOXXO struct {
	// OXXO reference number
	Number string `json:"number"`
}
type ChargePaymentMethodDetailsP24 struct {
	// The customer's bank. Can be one of `ing`, `citi_handlowy`, `tmobile_usbugi_bankowe`, `plus_bank`, `etransfer_pocztowy24`, `banki_spbdzielcze`, `bank_nowy_bfg_sa`, `getin_bank`, `blik`, `noble_pay`, `ideabank`, `envelobank`, `santander_przelew24`, `nest_przelew`, `mbank_mtransfer`, `inteligo`, `pbac_z_ipko`, `bnp_paribas`, `credit_agricole`, `toyota_bank`, `bank_pekao_sa`, `volkswagen_bank`, `bank_millennium`, `alior_bank`, or `boz`.
	Bank string `json:"bank"`
	// Unique reference for this Przelewy24 payment.
	Reference string `json:"reference"`
	// Owner's verified full name. Values are verified or provided by Przelewy24 directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	// Przelewy24 rarely provides this information so the attribute is usually empty.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsPayNow struct {
	// Reference number associated with this PayNow payment
	Reference string `json:"reference"`
}
type ChargePaymentMethodDetailsPromptPay struct {
	// Bill reference generated by PromptPay
	Reference string `json:"reference"`
}
type ChargePaymentMethodDetailsSepaCreditTransfer struct {
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	Bic string `json:"bic"`
	// IBAN of the bank account to transfer funds to.
	Iban string `json:"iban"`
}
type ChargePaymentMethodDetailsSepaDebit struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Branch code of bank associated with the bank account.
	BranchCode string `json:"branch_code"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four characters of the IBAN.
	Last4 string `json:"last4"`
	// ID of the mandate used to make this payment.
	Mandate *Mandate `json:"mandate"`
}
type ChargePaymentMethodDetailsSofort struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	Bic string `json:"bic"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
	// The ID of the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSepaDebit *PaymentMethod `json:"generated_sepa_debit"`
	// The mandate for the SEPA Direct Debit PaymentMethod which was generated by this Charge.
	GeneratedSepaDebitMandate *Mandate `json:"generated_sepa_debit_mandate"`
	// Last four characters of the IBAN.
	IbanLast4 string `json:"iban_last4"`
	// Preferred language of the SOFORT authorization page that the customer is redirected to.
	// Can be one of `de`, `en`, `es`, `fr`, `it`, `nl`, or `pl`
	PreferredLanguage string `json:"preferred_language"`
	// Owner's verified full name. Values are verified or provided by SOFORT directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsStripeAccount struct{}
type ChargePaymentMethodDetailsUSBankAccount struct {
	// Account holder type: individual or company.
	AccountHolderType ChargePaymentMethodDetailsUSBankAccountAccountHolderType `json:"account_holder_type"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType ChargePaymentMethodDetailsUSBankAccountAccountType `json:"account_type"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// Routing number of the bank account.
	RoutingNumber string `json:"routing_number"`
}
type ChargePaymentMethodDetailsWechat struct{}
type ChargePaymentMethodDetailsWechatPay struct {
	// Uniquely identifies this particular WeChat Pay account. You can use this attribute to check whether two WeChat accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Transaction ID of this particular WeChat Pay transaction.
	TransactionID string `json:"transaction_id"`
}

// Details about the payment method at the time of the transaction.
type ChargePaymentMethodDetails struct {
	AchCreditTransfer  *ChargePaymentMethodDetailsAchCreditTransfer  `json:"ach_credit_transfer"`
	AchDebit           *ChargePaymentMethodDetailsAchDebit           `json:"ach_debit"`
	AcssDebit          *ChargePaymentMethodDetailsAcssDebit          `json:"acss_debit"`
	Affirm             *ChargePaymentMethodDetailsAffirm             `json:"affirm"`
	AfterpayClearpay   *ChargePaymentMethodDetailsAfterpayClearpay   `json:"afterpay_clearpay"`
	Alipay             *ChargePaymentMethodDetailsAlipay             `json:"alipay"`
	AUBECSDebit        *ChargePaymentMethodDetailsAUBECSDebit        `json:"au_becs_debit"`
	BACSDebit          *ChargePaymentMethodDetailsBACSDebit          `json:"bacs_debit"`
	Bancontact         *ChargePaymentMethodDetailsBancontact         `json:"bancontact"`
	BLIK               *ChargePaymentMethodDetailsBLIK               `json:"blik"`
	Boleto             *ChargePaymentMethodDetailsBoleto             `json:"boleto"`
	Card               *ChargePaymentMethodDetailsCard               `json:"card"`
	CardPresent        *ChargePaymentMethodDetailsCardPresent        `json:"card_present"`
	CustomerBalance    *ChargePaymentMethodDetailsCustomerBalance    `json:"customer_balance"`
	Eps                *ChargePaymentMethodDetailsEps                `json:"eps"`
	FPX                *ChargePaymentMethodDetailsFPX                `json:"fpx"`
	Giropay            *ChargePaymentMethodDetailsGiropay            `json:"giropay"`
	Grabpay            *ChargePaymentMethodDetailsGrabpay            `json:"grabpay"`
	Ideal              *ChargePaymentMethodDetailsIdeal              `json:"ideal"`
	InteracPresent     *ChargePaymentMethodDetailsInteracPresent     `json:"interac_present"`
	Klarna             *ChargePaymentMethodDetailsKlarna             `json:"klarna"`
	Konbini            *ChargePaymentMethodDetailsKonbini            `json:"konbini"`
	Link               *ChargePaymentMethodDetailsLink               `json:"link"`
	Multibanco         *ChargePaymentMethodDetailsMultibanco         `json:"multibanco"`
	OXXO               *ChargePaymentMethodDetailsOXXO               `json:"oxxo"`
	P24                *ChargePaymentMethodDetailsP24                `json:"p24"`
	PayNow             *ChargePaymentMethodDetailsPayNow             `json:"paynow"`
	PromptPay          *ChargePaymentMethodDetailsPromptPay          `json:"promptpay"`
	SepaCreditTransfer *ChargePaymentMethodDetailsSepaCreditTransfer `json:"sepa_credit_transfer"`
	SepaDebit          *ChargePaymentMethodDetailsSepaDebit          `json:"sepa_debit"`
	Sofort             *ChargePaymentMethodDetailsSofort             `json:"sofort"`
	StripeAccount      *ChargePaymentMethodDetailsStripeAccount      `json:"stripe_account"`
	// The type of transaction-specific details of the payment method used in the payment, one of `ach_credit_transfer`, `ach_debit`, `acss_debit`, `alipay`, `au_becs_debit`, `bancontact`, `card`, `card_present`, `eps`, `giropay`, `ideal`, `klarna`, `multibanco`, `p24`, `sepa_debit`, `sofort`, `stripe_account`, or `wechat`.
	// An additional hash is included on `payment_method_details` with a name matching this value.
	// It contains information specific to the payment method.
	Type          ChargePaymentMethodDetailsType           `json:"type"`
	USBankAccount *ChargePaymentMethodDetailsUSBankAccount `json:"us_bank_account"`
	Wechat        *ChargePaymentMethodDetailsWechat        `json:"wechat"`
	WechatPay     *ChargePaymentMethodDetailsWechatPay     `json:"wechat_pay"`
}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type ChargeRadarOptions struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session string `json:"session"`
}

// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://stripe.com/docs/connect/destination-charges) for details.
type ChargeTransferData struct {
	// The amount transferred to the destination account, if specified. By default, the entire charge amount is transferred to the destination account.
	Amount int64 `json:"amount"`
	// ID of an existing, connected Stripe account to transfer funds to if `transfer_data` was specified in the charge request.
	Destination *Account `json:"destination"`
}

// To charge a credit or a debit card, you create a `Charge` object. You can
// retrieve and refund individual charges as well as list all charges. Charges
// are identified by a unique, random ID.
//
// Related guide: [Accept a payment with the Charges API](https://stripe.com/docs/payments/accept-a-payment-charges).
type Charge struct {
	APIResource
	// Amount intended to be collected by this payment. A positive integer representing how much to charge in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://stripe.com/docs/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount int64 `json:"amount"`
	// Amount in %s captured (can be less than the amount attribute on the charge if a partial capture was made).
	AmountCaptured int64 `json:"amount_captured"`
	// Amount in %s refunded (can be less than the amount attribute on the charge if a partial refund was issued).
	AmountRefunded int64 `json:"amount_refunded"`
	// ID of the Connect application that created the charge.
	Application *Application `json:"application"`
	// The application fee (if any) for the charge. [See the Connect documentation](https://stripe.com/docs/connect/direct-charges#collecting-fees) for details.
	ApplicationFee *ApplicationFee `json:"application_fee"`
	// The amount of the application fee (if any) requested for the charge. [See the Connect documentation](https://stripe.com/docs/connect/direct-charges#collecting-fees) for details.
	ApplicationFeeAmount int64 `json:"application_fee_amount"`
	// Authorization code on the charge.
	AuthorizationCode string `json:"authorization_code"`
	// ID of the balance transaction that describes the impact of this charge on your account balance (not including refunds or disputes).
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	BillingDetails     *BillingDetails     `json:"billing_details"`
	// The full statement descriptor that is passed to card networks, and that is displayed on your customers' credit card and bank statements. Allows you to see what the statement descriptor looks like after the static and dynamic portions are combined.
	CalculatedStatementDescriptor string `json:"calculated_statement_descriptor"`
	// If the charge was created without capturing, this Boolean represents whether it is still uncaptured or has since been captured.
	Captured bool `json:"captured"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// ID of the customer this charge is for if one exists.
	Customer *Customer `json:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// ID of an existing, connected Stripe account to transfer funds to if `transfer_data` was specified in the charge request.
	Destination *Account `json:"destination"`
	// Details about the dispute if the charge has been disputed.
	Dispute *Dispute `json:"dispute"`
	// Whether the charge has been disputed.
	Disputed bool `json:"disputed"`
	// ID of the balance transaction that describes the reversal of the balance on your account due to payment failure.
	FailureBalanceTransaction *BalanceTransaction `json:"failure_balance_transaction"`
	// Error code explaining reason for charge failure if available (see [the errors section](https://stripe.com/docs/api#errors) for a list of codes).
	FailureCode string `json:"failure_code"`
	// Message to user further explaining reason for charge failure if available.
	FailureMessage string `json:"failure_message"`
	// Information on fraud assessments for the charge.
	FraudDetails *FraudDetails `json:"fraud_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// ID of the invoice this charge is for if one exists.
	Invoice *Invoice     `json:"invoice"`
	Level3  ChargeLevel3 `json:"level3"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account (if any) the charge was made on behalf of without triggering an automatic transfer. See the [Connect documentation](https://stripe.com/docs/connect/charges-transfers) for details.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// Deprecated
	Order *Order `json:"order"`
	// Details about whether the payment was accepted, and why. See [understanding declines](https://stripe.com/docs/declines) for details.
	Outcome *ChargeOutcome `json:"outcome"`
	// `true` if the charge succeeded, or was successfully authorized for later capture.
	Paid bool `json:"paid"`
	// ID of the PaymentIntent associated with this charge, if one exists.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
	// ID of the payment method used in this charge.
	PaymentMethod string `json:"payment_method"`
	// Details about the payment method at the time of the transaction.
	PaymentMethodDetails *ChargePaymentMethodDetails `json:"payment_method_details"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *ChargeRadarOptions `json:"radar_options"`
	// This is the email address that the receipt for this charge was sent to.
	ReceiptEmail string `json:"receipt_email"`
	// This is the transaction number that appears on email receipts sent for this charge. This attribute will be `null` until a receipt has been sent.
	ReceiptNumber string `json:"receipt_number"`
	// This is the URL to view the receipt for this charge. The receipt is kept up-to-date to the latest state of the charge, including any refunds. If the charge is for an Invoice, the receipt will be stylized as an Invoice receipt.
	ReceiptURL string `json:"receipt_url"`
	// Whether the charge has been fully refunded. If the charge is only partially refunded, this attribute will still be false.
	Refunded bool `json:"refunded"`
	// A list of refunds that have been applied to the charge.
	Refunds *RefundList `json:"refunds"`
	// ID of the review associated with this charge if one exists.
	Review *Review `json:"review"`
	// Shipping information for the charge.
	Shipping *ShippingDetails `json:"shipping"`
	// This is a legacy field that will be removed in the future. It contains the Source, Card, or BankAccount object used for the charge. For details about the payment method used for this charge, refer to `payment_method` or `payment_method_details` instead.
	Source *PaymentSource `json:"source"`
	// The transfer ID which created this charge. Only present if the charge came from another Stripe account. [See the Connect documentation](https://stripe.com/docs/connect/destination-charges) for details.
	SourceTransfer *Transfer `json:"source_transfer"`
	// For card charges, use `statement_descriptor_suffix` instead. Otherwise, you can use this value as the complete description of a charge on your customers' statements. Must contain at least one letter, maximum 22 characters.
	StatementDescriptor string `json:"statement_descriptor"`
	// Provides information about the charge that customers see on their statements. Concatenated with the prefix (shortened descriptor) or statement descriptor that's set on the account to form the complete statement descriptor. Maximum 22 characters for the concatenated descriptor.
	StatementDescriptorSuffix string `json:"statement_descriptor_suffix"`
	// The status of the payment is either `succeeded`, `pending`, or `failed`.
	Status ChargeStatus `json:"status"`
	// ID of the transfer to the `destination` account (only applicable if the charge was created using the `destination` parameter).
	Transfer *Transfer `json:"transfer"`
	// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://stripe.com/docs/connect/destination-charges) for details.
	TransferData *ChargeTransferData `json:"transfer_data"`
	// A string that identifies this transaction as part of a group. See the [Connect documentation](https://stripe.com/docs/connect/charges-transfers#transfer-options) for details.
	TransferGroup string `json:"transfer_group"`
}

// ChargeList is a list of Charges as retrieved from a list endpoint.
type ChargeList struct {
	APIResource
	ListMeta
	Data []*Charge `json:"data"`
}

// ChargeSearchResult is a list of Charge search results as retrieved from a search endpoint.
type ChargeSearchResult struct {
	APIResource
	SearchMeta
	Data []*Charge `json:"data"`
}

// UnmarshalJSON handles deserialization of a Charge.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Charge) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type charge Charge
	var v charge
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Charge(v)
	return nil
}
