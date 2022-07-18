//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Reason for cancellation of this SetupIntent, one of `abandoned`, `requested_by_customer`, or `duplicate`.
type SetupIntentCancellationReason string

// List of values that SetupIntentCancellationReason can take
const (
	SetupIntentCancellationReasonAbandoned           SetupIntentCancellationReason = "abandoned"
	SetupIntentCancellationReasonDuplicate           SetupIntentCancellationReason = "duplicate"
	SetupIntentCancellationReasonFailedInvoice       SetupIntentCancellationReason = "failed_invoice"
	SetupIntentCancellationReasonFraudulent          SetupIntentCancellationReason = "fraudulent"
	SetupIntentCancellationReasonRequestedByCustomer SetupIntentCancellationReason = "requested_by_customer"
)

// Indicates the directions of money movement for which this payment method is intended to be used.
//
// Include `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.
type SetupIntentFlowDirection string

// List of values that SetupIntentFlowDirection can take
const (
	SetupIntentFlowDirectionInbound  SetupIntentFlowDirection = "inbound"
	SetupIntentFlowDirectionOutbound SetupIntentFlowDirection = "outbound"
)

// Type of the next action to perform, one of `redirect_to_url`, `use_stripe_sdk`, `alipay_handle_redirect`, `oxxo_display_details`, or `verify_with_microdeposits`.
type SetupIntentNextActionType string

// List of values that SetupIntentNextActionType can take
const (
	SetupIntentNextActionTypeRedirectToURL SetupIntentNextActionType = "redirect_to_url"
)

// The type of the microdeposit sent to the customer. Used to distinguish between different verification methods.
type SetupIntentNextActionVerifyWithMicrodepositsMicrodepositType string

// List of values that SetupIntentNextActionVerifyWithMicrodepositsMicrodepositType can take
const (
	SetupIntentNextActionVerifyWithMicrodepositsMicrodepositTypeAmounts        SetupIntentNextActionVerifyWithMicrodepositsMicrodepositType = "amounts"
	SetupIntentNextActionVerifyWithMicrodepositsMicrodepositTypeDescriptorCode SetupIntentNextActionVerifyWithMicrodepositsMicrodepositType = "descriptor_code"
)

// Currency supported by the bank account
type SetupIntentPaymentMethodOptionsACSSDebitCurrency string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitCurrency can take
const (
	SetupIntentPaymentMethodOptionsACSSDebitCurrencyCAD SetupIntentPaymentMethodOptionsACSSDebitCurrency = "cad"
	SetupIntentPaymentMethodOptionsACSSDebitCurrencyUSD SetupIntentPaymentMethodOptionsACSSDebitCurrency = "usd"
)

// List of Stripe products where this mandate can be selected automatically.
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor can take
const (
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultForInvoice      SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor = "invoice"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultForSubscription SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor = "subscription"
)

// Payment schedule for the mandate.
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule can take
const (
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleCombined SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "combined"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleInterval SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "interval"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleSporadic SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "sporadic"
)

// Transaction type of the mandate.
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Bank account verification method.
type SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	SetupIntentPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	SetupIntentPaymentMethodOptionsACSSDebitVerificationMethodInstant       SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	SetupIntentPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// Frequency interval of each recurring payment.
type SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionInterval string

// List of values that SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionInterval can take
const (
	SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionIntervalDay   SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionInterval = "day"
	SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionIntervalMonth SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionInterval = "month"
	SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionIntervalWeek  SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionInterval = "week"
	SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionIntervalYear  SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionInterval = "year"
)

// Type of the mandate.
type SetupIntentPaymentMethodOptionsBLIKMandateOptionsType string

// List of values that SetupIntentPaymentMethodOptionsBLIKMandateOptionsType can take
const (
	SetupIntentPaymentMethodOptionsBLIKMandateOptionsTypeOffSession SetupIntentPaymentMethodOptionsBLIKMandateOptionsType = "off_session"
	SetupIntentPaymentMethodOptionsBLIKMandateOptionsTypeOnSession  SetupIntentPaymentMethodOptionsBLIKMandateOptionsType = "on_session"
)

// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
type SetupIntentPaymentMethodOptionsCardMandateOptionsAmountType string

// List of values that SetupIntentPaymentMethodOptionsCardMandateOptionsAmountType can take
const (
	SetupIntentPaymentMethodOptionsCardMandateOptionsAmountTypeFixed   SetupIntentPaymentMethodOptionsCardMandateOptionsAmountType = "fixed"
	SetupIntentPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum SetupIntentPaymentMethodOptionsCardMandateOptionsAmountType = "maximum"
)

// Specifies payment frequency. One of `day`, `week`, `month`, `year`, or `sporadic`.
type SetupIntentPaymentMethodOptionsCardMandateOptionsInterval string

// List of values that SetupIntentPaymentMethodOptionsCardMandateOptionsInterval can take
const (
	SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalDay      SetupIntentPaymentMethodOptionsCardMandateOptionsInterval = "day"
	SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalMonth    SetupIntentPaymentMethodOptionsCardMandateOptionsInterval = "month"
	SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalSporadic SetupIntentPaymentMethodOptionsCardMandateOptionsInterval = "sporadic"
	SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalWeek     SetupIntentPaymentMethodOptionsCardMandateOptionsInterval = "week"
	SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalYear     SetupIntentPaymentMethodOptionsCardMandateOptionsInterval = "year"
)

// Specifies the type of mandates supported. Possible values are `india`.
type SetupIntentPaymentMethodOptionsCardMandateOptionsSupportedType string

// List of values that SetupIntentPaymentMethodOptionsCardMandateOptionsSupportedType can take
const (
	SetupIntentPaymentMethodOptionsCardMandateOptionsSupportedTypeIndia SetupIntentPaymentMethodOptionsCardMandateOptionsSupportedType = "india"
)

// Selected network to process this SetupIntent on. Depends on the available networks of the card attached to the setup intent. Can be only set confirm-time.
type SetupIntentPaymentMethodOptionsCardNetwork string

// List of values that SetupIntentPaymentMethodOptionsCardNetwork can take
const (
	SetupIntentPaymentMethodOptionsCardNetworkAmex            SetupIntentPaymentMethodOptionsCardNetwork = "amex"
	SetupIntentPaymentMethodOptionsCardNetworkCartesBancaires SetupIntentPaymentMethodOptionsCardNetwork = "cartes_bancaires"
	SetupIntentPaymentMethodOptionsCardNetworkDiners          SetupIntentPaymentMethodOptionsCardNetwork = "diners"
	SetupIntentPaymentMethodOptionsCardNetworkDiscover        SetupIntentPaymentMethodOptionsCardNetwork = "discover"
	SetupIntentPaymentMethodOptionsCardNetworkInterac         SetupIntentPaymentMethodOptionsCardNetwork = "interac"
	SetupIntentPaymentMethodOptionsCardNetworkJCB             SetupIntentPaymentMethodOptionsCardNetwork = "jcb"
	SetupIntentPaymentMethodOptionsCardNetworkMastercard      SetupIntentPaymentMethodOptionsCardNetwork = "mastercard"
	SetupIntentPaymentMethodOptionsCardNetworkUnionpay        SetupIntentPaymentMethodOptionsCardNetwork = "unionpay"
	SetupIntentPaymentMethodOptionsCardNetworkUnknown         SetupIntentPaymentMethodOptionsCardNetwork = "unknown"
	SetupIntentPaymentMethodOptionsCardNetworkVisa            SetupIntentPaymentMethodOptionsCardNetwork = "visa"
)

// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Permitted values include: `automatic` or `any`. If not provided, defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type SetupIntentPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that SetupIntentPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAny           SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "any"
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAutomatic     SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureChallengeOnly SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "challenge_only"
)

// The list of permissions to request. The `payment_method` permission must be included.
type SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission string

// List of values that SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission can take
const (
	SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionBalances      SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "balances"
	SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionOwnership     SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "ownership"
	SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionPaymentMethod SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "payment_method"
	SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermissionTransactions  SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission = "transactions"
)

// Bank account verification method.
type SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod string

// List of values that SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod can take
const (
	SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethodAutomatic     SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod = "automatic"
	SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethodInstant       SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod = "instant"
	SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethodMicrodeposits SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod = "microdeposits"
)

// [Status](https://stripe.com/docs/payments/intents#intent-statuses) of this SetupIntent, one of `requires_payment_method`, `requires_confirmation`, `requires_action`, `processing`, `canceled`, or `succeeded`.
type SetupIntentStatus string

// List of values that SetupIntentStatus can take
const (
	SetupIntentStatusCanceled              SetupIntentStatus = "canceled"
	SetupIntentStatusProcessing            SetupIntentStatus = "processing"
	SetupIntentStatusRequiresAction        SetupIntentStatus = "requires_action"
	SetupIntentStatusRequiresConfirmation  SetupIntentStatus = "requires_confirmation"
	SetupIntentStatusRequiresPaymentMethod SetupIntentStatus = "requires_payment_method"
	SetupIntentStatusSucceeded             SetupIntentStatus = "succeeded"
)

// Indicates how the payment method is intended to be used in the future.
//
// Use `on_session` if you intend to only reuse the payment method when the customer is in your checkout flow. Use `off_session` if your customer may or may not be in your checkout flow. If not provided, this value defaults to `off_session`.
type SetupIntentUsage string

// List of values that SetupIntentUsage can take
const (
	SetupIntentUsageOffSession SetupIntentUsage = "off_session"
	SetupIntentUsageOnSession  SetupIntentUsage = "on_session"
)

// If this is a Mandate accepted offline, this hash contains details about the offline acceptance.
type SetupIntentMandateDataCustomerAcceptanceOfflineParams struct{}

// If this is a Mandate accepted online, this hash contains details about the online acceptance.
type SetupIntentMandateDataCustomerAcceptanceOnlineParams struct {
	// The IP address from which the Mandate was accepted by the customer.
	IPAddress *string `form:"ip_address"`
	// The user agent of the browser from which the Mandate was accepted by the customer.
	UserAgent *string `form:"user_agent"`
}

// This hash contains details about the customer acceptance of the Mandate.
type SetupIntentMandateDataCustomerAcceptanceParams struct {
	// The time at which the customer accepted the Mandate.
	AcceptedAt int64 `form:"accepted_at"`
	// If this is a Mandate accepted offline, this hash contains details about the offline acceptance.
	Offline *SetupIntentMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	// If this is a Mandate accepted online, this hash contains details about the online acceptance.
	Online *SetupIntentMandateDataCustomerAcceptanceOnlineParams `form:"online"`
	// The type of customer acceptance information included with the Mandate.
	Type MandateCustomerAcceptanceType `form:"type"`
}

// This hash contains details about the Mandate to create. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
type SetupIntentMandateDataParams struct {
	// This hash contains details about the customer acceptance of the Mandate.
	CustomerAcceptance *SetupIntentMandateDataCustomerAcceptanceParams `form:"customer_acceptance"`
}

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type SetupIntentPaymentMethodDataACSSDebitParams struct {
	// Customer's bank account number.
	AccountNumber *string `form:"account_number"`
	// Institution number of the customer's bank.
	InstitutionNumber *string `form:"institution_number"`
	// Transit number of the customer's bank.
	TransitNumber *string `form:"transit_number"`
}

// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
type SetupIntentPaymentMethodDataAffirmParams struct{}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type SetupIntentPaymentMethodDataAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type SetupIntentPaymentMethodDataAlipayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type SetupIntentPaymentMethodDataAUBECSDebitParams struct {
	// The account number for the bank account.
	AccountNumber *string `form:"account_number"`
	// Bank-State-Branch number of the bank account.
	BSBNumber *string `form:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type SetupIntentPaymentMethodDataBACSDebitParams struct {
	// Account number of the bank account that the funds will be debited from.
	AccountNumber *string `form:"account_number"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode *string `form:"sort_code"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type SetupIntentPaymentMethodDataBancontactParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type SetupIntentPaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address"`
	// Email address.
	Email *string `form:"email"`
	// Full name.
	Name *string `form:"name"`
	// Billing phone number (including extension).
	Phone *string `form:"phone"`
}

// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
type SetupIntentPaymentMethodDataBLIKParams struct{}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type SetupIntentPaymentMethodDataBoletoParams struct {
	// The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)
	TaxID *string `form:"tax_id"`
}

// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
type SetupIntentPaymentMethodDataCustomerBalanceParams struct{}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type SetupIntentPaymentMethodDataEPSParams struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type SetupIntentPaymentMethodDataFPXParams struct {
	// Account holder type for FPX transaction
	AccountHolderType *string `form:"account_holder_type"`
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type SetupIntentPaymentMethodDataGiropayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type SetupIntentPaymentMethodDataGrabpayParams struct{}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type SetupIntentPaymentMethodDataIdealParams struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type SetupIntentPaymentMethodDataInteracPresentParams struct{}

// Customer's date of birth
type SetupIntentPaymentMethodDataKlarnaDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
type SetupIntentPaymentMethodDataKlarnaParams struct {
	// Customer's date of birth
	DOB *SetupIntentPaymentMethodDataKlarnaDOBParams `form:"dob"`
}

// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
type SetupIntentPaymentMethodDataKonbiniParams struct{}

// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
type SetupIntentPaymentMethodDataLinkParams struct{}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type SetupIntentPaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type SetupIntentPaymentMethodDataP24Params struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
type SetupIntentPaymentMethodDataPayNowParams struct{}

// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
type SetupIntentPaymentMethodDataPromptPayParams struct{}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type SetupIntentPaymentMethodDataRadarOptionsParams struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type SetupIntentPaymentMethodDataSepaDebitParams struct {
	// IBAN of the bank account.
	Iban *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type SetupIntentPaymentMethodDataSofortParams struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country *string `form:"country"`
}

// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
type SetupIntentPaymentMethodDataUSBankAccountParams struct {
	// Account holder type: individual or company.
	AccountHolderType *string `form:"account_holder_type"`
	// Account number of the bank account.
	AccountNumber *string `form:"account_number"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType *string `form:"account_type"`
	// The ID of a Financial Connections Account to use as a payment method.
	FinancialConnectionsAccount *string `form:"financial_connections_account"`
	// Routing number of the bank account.
	RoutingNumber *string `form:"routing_number"`
}

// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
type SetupIntentPaymentMethodDataWechatPayParams struct{}

// When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method)
// value in the SetupIntent.
type SetupIntentPaymentMethodDataParams struct {
	// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
	ACSSDebit *SetupIntentPaymentMethodDataACSSDebitParams `form:"acss_debit"`
	// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
	Affirm *SetupIntentPaymentMethodDataAffirmParams `form:"affirm"`
	// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
	AfterpayClearpay *SetupIntentPaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay"`
	// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
	Alipay *SetupIntentPaymentMethodDataAlipayParams `form:"alipay"`
	// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
	AUBECSDebit *SetupIntentPaymentMethodDataAUBECSDebitParams `form:"au_becs_debit"`
	// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
	BACSDebit *SetupIntentPaymentMethodDataBACSDebitParams `form:"bacs_debit"`
	// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
	Bancontact *SetupIntentPaymentMethodDataBancontactParams `form:"bancontact"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *SetupIntentPaymentMethodDataBillingDetailsParams `form:"billing_details"`
	// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
	BLIK *SetupIntentPaymentMethodDataBLIKParams `form:"blik"`
	// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
	Boleto *SetupIntentPaymentMethodDataBoletoParams `form:"boleto"`
	// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
	CustomerBalance *SetupIntentPaymentMethodDataCustomerBalanceParams `form:"customer_balance"`
	// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
	EPS *SetupIntentPaymentMethodDataEPSParams `form:"eps"`
	// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
	FPX *SetupIntentPaymentMethodDataFPXParams `form:"fpx"`
	// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
	Giropay *SetupIntentPaymentMethodDataGiropayParams `form:"giropay"`
	// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
	Grabpay *SetupIntentPaymentMethodDataGrabpayParams `form:"grabpay"`
	// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
	Ideal *SetupIntentPaymentMethodDataIdealParams `form:"ideal"`
	// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
	InteracPresent *SetupIntentPaymentMethodDataInteracPresentParams `form:"interac_present"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
	Klarna *SetupIntentPaymentMethodDataKlarnaParams `form:"klarna"`
	// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
	Konbini *SetupIntentPaymentMethodDataKonbiniParams `form:"konbini"`
	// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
	Link *SetupIntentPaymentMethodDataLinkParams `form:"link"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
	OXXO *SetupIntentPaymentMethodDataOXXOParams `form:"oxxo"`
	// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
	P24 *SetupIntentPaymentMethodDataP24Params `form:"p24"`
	// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
	PayNow *SetupIntentPaymentMethodDataPayNowParams `form:"paynow"`
	// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
	PromptPay *SetupIntentPaymentMethodDataPromptPayParams `form:"promptpay"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *SetupIntentPaymentMethodDataRadarOptionsParams `form:"radar_options"`
	// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
	SepaDebit *SetupIntentPaymentMethodDataSepaDebitParams `form:"sepa_debit"`
	// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
	Sofort *SetupIntentPaymentMethodDataSofortParams `form:"sofort"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type"`
	// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
	USBankAccount *SetupIntentPaymentMethodDataUSBankAccountParams `form:"us_bank_account"`
	// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
	WechatPay *SetupIntentPaymentMethodDataWechatPayParams `form:"wechat_pay"`
}

// Additional fields for Mandate creation
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// A URL for custom mandate text to render during confirmation step.
	// The URL will be rendered with additional GET parameters `payment_intent` and `payment_intent_client_secret` when confirming a Payment Intent,
	// or `setup_intent` and `setup_intent_client_secret` when confirming a Setup Intent.
	CustomMandateURL *string `form:"custom_mandate_url"`
	// List of Stripe products where this mandate can be selected automatically.
	DefaultFor []*string `form:"default_for"`
	// Description of the mandate interval. Only required if 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription *string `form:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule *string `form:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type"`
}

// If this is a `acss_debit` SetupIntent, this sub-hash contains details about the ACSS Debit payment method options.
type SetupIntentPaymentMethodOptionsACSSDebitParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
type SetupIntentPaymentMethodOptionsBLIKParams struct {
	// The 6-digit BLIK code that a customer has generated using their banking application. Can only be set on confirmation.
	Code *string `form:"code"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type SetupIntentPaymentMethodOptionsCardMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount"`
	// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
	AmountType *string `form:"amount_type"`
	// Currency in which future payments will be charged. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// A description of the mandate or subscription that is meant to be displayed to the customer.
	Description *string `form:"description"`
	// End date of the mandate or subscription. If not provided, the mandate will be active until canceled. If provided, end date should be after start date.
	EndDate *int64 `form:"end_date"`
	// Specifies payment frequency. One of `day`, `week`, `month`, `year`, or `sporadic`.
	Interval *string `form:"interval"`
	// The number of intervals between payments. For example, `interval=month` and `interval_count=3` indicates one payment every three months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks). This parameter is optional when `interval=sporadic`.
	IntervalCount *int64 `form:"interval_count"`
	// Unique identifier for the mandate or subscription.
	Reference *string `form:"reference"`
	// Start date of the mandate or subscription. Start date should not be lesser than yesterday.
	StartDate *int64 `form:"start_date"`
	// Specifies the type of mandates supported. Possible values are `india`.
	SupportedTypes []*string `form:"supported_types"`
}

// Configuration for any card setup attempted on this SetupIntent.
type SetupIntentPaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *SetupIntentPaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options"`
	// When specified, this parameter signals that a card has been collected
	// as MOTO (Mail Order Telephone Order) and thus out of scope for SCA. This
	// parameter can only be provided during confirmation.
	MOTO *bool `form:"moto"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Permitted values include: `automatic` or `any`. If not provided, defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure"`
}

// If this is a `link` PaymentMethod, this sub-hash contains details about the Link payment method options.
type SetupIntentPaymentMethodOptionsLinkParams struct {
	// Token used for persistent Link logins.
	PersistentToken *string `form:"persistent_token"`
}

// Additional fields for Mandate creation
type SetupIntentPaymentMethodOptionsSepaDebitMandateOptionsParams struct{}

// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
type SetupIntentPaymentMethodOptionsSepaDebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentPaymentMethodOptionsSepaDebitMandateOptionsParams `form:"mandate_options"`
}

// Additional fields for Financial Connections Session creation
type SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included. Valid permissions include: `balances`, `ownership`, `payment_method`, and `transactions`.
	Permissions []*string `form:"permissions"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL *string `form:"return_url"`
}

// Additional fields for network related functions
type SetupIntentPaymentMethodOptionsUSBankAccountNetworksParams struct {
	// Triggers validations to run across the selected networks
	Requested []*string `form:"requested"`
}

// If this is a `us_bank_account` SetupIntent, this sub-hash contains details about the US bank account payment method options.
type SetupIntentPaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation
	FinancialConnections *SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections"`
	// Additional fields for network related functions
	Networks *SetupIntentPaymentMethodOptionsUSBankAccountNetworksParams `form:"networks"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// Payment-method-specific configuration for this SetupIntent.
type SetupIntentPaymentMethodOptionsParams struct {
	// If this is a `acss_debit` SetupIntent, this sub-hash contains details about the ACSS Debit payment method options.
	ACSSDebit *SetupIntentPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
	BLIK *SetupIntentPaymentMethodOptionsBLIKParams `form:"blik"`
	// Configuration for any card setup attempted on this SetupIntent.
	Card *SetupIntentPaymentMethodOptionsCardParams `form:"card"`
	// If this is a `link` PaymentMethod, this sub-hash contains details about the Link payment method options.
	Link *SetupIntentPaymentMethodOptionsLinkParams `form:"link"`
	// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
	SepaDebit *SetupIntentPaymentMethodOptionsSepaDebitParams `form:"sepa_debit"`
	// If this is a `us_bank_account` SetupIntent, this sub-hash contains details about the US bank account payment method options.
	USBankAccount *SetupIntentPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account"`
}

// If this hash is populated, this SetupIntent will generate a single_use Mandate on success.
type SetupIntentSingleUseParams struct {
	// Amount the customer is granting permission to collect later. A positive integer representing how much to charge in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://stripe.com/docs/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
}

// Creates a SetupIntent object.
//
// After the SetupIntent is created, attach a payment method and [confirm](https://stripe.com/docs/api/setup_intents/confirm)
// to collect any required permissions to charge the payment method later.
type SetupIntentParams struct {
	Params `form:"*"`
	// If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.
	//
	// It can only be used for this Stripe Account's own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.
	AttachToSelf *bool `form:"attach_to_self"`
	// The client secret of the SetupIntent. Required if a publishable key is used to retrieve the SetupIntent.
	ClientSecret *string `form:"client_secret"`
	// Set to `true` to attempt to confirm this SetupIntent immediately. This parameter defaults to `false`. If the payment method attached is a card, a return_url may be provided in case additional authentication is required.
	Confirm *bool `form:"confirm"`
	// ID of the Customer this SetupIntent belongs to, if one exists.
	//
	// If present, the SetupIntent's payment method will be attached to the Customer on successful setup. Payment methods attached to other Customers cannot be used with this SetupIntent.
	Customer *string `form:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Indicates the directions of money movement for which this payment method is intended to be used.
	//
	// Include `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.
	FlowDirections []*string `form:"flow_directions"`
	// This hash contains details about the Mandate to create. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
	MandateData *SetupIntentMandateDataParams `form:"mandate_data"`
	// The Stripe account ID for which this SetupIntent is created.
	OnBehalfOf *string `form:"on_behalf_of"`
	// ID of the payment method (a PaymentMethod, Card, or saved Source object) to attach to this SetupIntent.
	PaymentMethod *string `form:"payment_method"`
	// When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method)
	// value in the SetupIntent.
	PaymentMethodData *SetupIntentPaymentMethodDataParams `form:"payment_method_data"`
	// Payment-method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	// The list of payment method types (e.g. card) that this SetupIntent is allowed to set up. If this is not provided, defaults to ["card"].
	PaymentMethodTypes []*string `form:"payment_method_types"`
	// The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method's app or site. If you'd prefer to redirect to a mobile application, you can alternatively supply an application URI scheme. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
	ReturnURL *string `form:"return_url"`
	// If this hash is populated, this SetupIntent will generate a single_use Mandate on success.
	SingleUse *SetupIntentSingleUseParams `form:"single_use"`
	// Indicates how the payment method is intended to be used in the future. If not provided, this value defaults to `off_session`.
	Usage *string `form:"usage"`
}

// Returns a list of SetupIntents.
type SetupIntentListParams struct {
	ListParams `form:"*"`
	// If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.
	//
	// It can only be used for this Stripe Account's own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.
	AttachToSelf *bool `form:"attach_to_self"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	Created *int64 `form:"created"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return SetupIntents for the customer specified by this customer ID.
	Customer *string `form:"customer"`
	// Only return SetupIntents associated with the specified payment method.
	PaymentMethod *string `form:"payment_method"`
}

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type SetupIntentConfirmPaymentMethodDataACSSDebitParams struct {
	// Customer's bank account number.
	AccountNumber *string `form:"account_number"`
	// Institution number of the customer's bank.
	InstitutionNumber *string `form:"institution_number"`
	// Transit number of the customer's bank.
	TransitNumber *string `form:"transit_number"`
}

// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
type SetupIntentConfirmPaymentMethodDataAffirmParams struct{}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type SetupIntentConfirmPaymentMethodDataAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type SetupIntentConfirmPaymentMethodDataAlipayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type SetupIntentConfirmPaymentMethodDataAUBECSDebitParams struct {
	// The account number for the bank account.
	AccountNumber *string `form:"account_number"`
	// Bank-State-Branch number of the bank account.
	BSBNumber *string `form:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type SetupIntentConfirmPaymentMethodDataBACSDebitParams struct {
	// Account number of the bank account that the funds will be debited from.
	AccountNumber *string `form:"account_number"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode *string `form:"sort_code"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type SetupIntentConfirmPaymentMethodDataBancontactParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type SetupIntentConfirmPaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address"`
	// Email address.
	Email *string `form:"email"`
	// Full name.
	Name *string `form:"name"`
	// Billing phone number (including extension).
	Phone *string `form:"phone"`
}

// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
type SetupIntentConfirmPaymentMethodDataBLIKParams struct{}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type SetupIntentConfirmPaymentMethodDataBoletoParams struct {
	// The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)
	TaxID *string `form:"tax_id"`
}

// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
type SetupIntentConfirmPaymentMethodDataCustomerBalanceParams struct{}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type SetupIntentConfirmPaymentMethodDataEPSParams struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type SetupIntentConfirmPaymentMethodDataFPXParams struct {
	// Account holder type for FPX transaction
	AccountHolderType *string `form:"account_holder_type"`
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type SetupIntentConfirmPaymentMethodDataGiropayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type SetupIntentConfirmPaymentMethodDataGrabpayParams struct{}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type SetupIntentConfirmPaymentMethodDataIdealParams struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type SetupIntentConfirmPaymentMethodDataInteracPresentParams struct{}

// Customer's date of birth
type SetupIntentConfirmPaymentMethodDataKlarnaDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
type SetupIntentConfirmPaymentMethodDataKlarnaParams struct {
	// Customer's date of birth
	DOB *SetupIntentConfirmPaymentMethodDataKlarnaDOBParams `form:"dob"`
}

// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
type SetupIntentConfirmPaymentMethodDataKonbiniParams struct{}

// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
type SetupIntentConfirmPaymentMethodDataLinkParams struct{}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type SetupIntentConfirmPaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type SetupIntentConfirmPaymentMethodDataP24Params struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
type SetupIntentConfirmPaymentMethodDataPayNowParams struct{}

// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
type SetupIntentConfirmPaymentMethodDataPromptPayParams struct{}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type SetupIntentConfirmPaymentMethodDataRadarOptionsParams struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type SetupIntentConfirmPaymentMethodDataSepaDebitParams struct {
	// IBAN of the bank account.
	Iban *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type SetupIntentConfirmPaymentMethodDataSofortParams struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country *string `form:"country"`
}

// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
type SetupIntentConfirmPaymentMethodDataUSBankAccountParams struct {
	// Account holder type: individual or company.
	AccountHolderType *string `form:"account_holder_type"`
	// Account number of the bank account.
	AccountNumber *string `form:"account_number"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType *string `form:"account_type"`
	// The ID of a Financial Connections Account to use as a payment method.
	FinancialConnectionsAccount *string `form:"financial_connections_account"`
	// Routing number of the bank account.
	RoutingNumber *string `form:"routing_number"`
}

// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
type SetupIntentConfirmPaymentMethodDataWechatPayParams struct{}

// When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method)
// value in the SetupIntent.
type SetupIntentConfirmPaymentMethodDataParams struct {
	// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
	ACSSDebit *SetupIntentConfirmPaymentMethodDataACSSDebitParams `form:"acss_debit"`
	// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
	Affirm *SetupIntentConfirmPaymentMethodDataAffirmParams `form:"affirm"`
	// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
	AfterpayClearpay *SetupIntentConfirmPaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay"`
	// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
	Alipay *SetupIntentConfirmPaymentMethodDataAlipayParams `form:"alipay"`
	// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
	AUBECSDebit *SetupIntentConfirmPaymentMethodDataAUBECSDebitParams `form:"au_becs_debit"`
	// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
	BACSDebit *SetupIntentConfirmPaymentMethodDataBACSDebitParams `form:"bacs_debit"`
	// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
	Bancontact *SetupIntentConfirmPaymentMethodDataBancontactParams `form:"bancontact"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *SetupIntentConfirmPaymentMethodDataBillingDetailsParams `form:"billing_details"`
	// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
	BLIK *SetupIntentConfirmPaymentMethodDataBLIKParams `form:"blik"`
	// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
	Boleto *SetupIntentConfirmPaymentMethodDataBoletoParams `form:"boleto"`
	// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
	CustomerBalance *SetupIntentConfirmPaymentMethodDataCustomerBalanceParams `form:"customer_balance"`
	// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
	EPS *SetupIntentConfirmPaymentMethodDataEPSParams `form:"eps"`
	// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
	FPX *SetupIntentConfirmPaymentMethodDataFPXParams `form:"fpx"`
	// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
	Giropay *SetupIntentConfirmPaymentMethodDataGiropayParams `form:"giropay"`
	// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
	Grabpay *SetupIntentConfirmPaymentMethodDataGrabpayParams `form:"grabpay"`
	// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
	Ideal *SetupIntentConfirmPaymentMethodDataIdealParams `form:"ideal"`
	// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
	InteracPresent *SetupIntentConfirmPaymentMethodDataInteracPresentParams `form:"interac_present"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
	Klarna *SetupIntentConfirmPaymentMethodDataKlarnaParams `form:"klarna"`
	// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
	Konbini *SetupIntentConfirmPaymentMethodDataKonbiniParams `form:"konbini"`
	// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
	Link *SetupIntentConfirmPaymentMethodDataLinkParams `form:"link"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
	OXXO *SetupIntentConfirmPaymentMethodDataOXXOParams `form:"oxxo"`
	// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
	P24 *SetupIntentConfirmPaymentMethodDataP24Params `form:"p24"`
	// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
	PayNow *SetupIntentConfirmPaymentMethodDataPayNowParams `form:"paynow"`
	// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
	PromptPay *SetupIntentConfirmPaymentMethodDataPromptPayParams `form:"promptpay"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *SetupIntentConfirmPaymentMethodDataRadarOptionsParams `form:"radar_options"`
	// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
	SepaDebit *SetupIntentConfirmPaymentMethodDataSepaDebitParams `form:"sepa_debit"`
	// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
	Sofort *SetupIntentConfirmPaymentMethodDataSofortParams `form:"sofort"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type"`
	// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
	USBankAccount *SetupIntentConfirmPaymentMethodDataUSBankAccountParams `form:"us_bank_account"`
	// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
	WechatPay *SetupIntentConfirmPaymentMethodDataWechatPayParams `form:"wechat_pay"`
}

// Confirm that your customer intends to set up the current or
// provided payment method. For example, you would confirm a SetupIntent
// when a customer hits the “Save” button on a payment method management
// page on your website.
//
// If the selected payment method does not require any additional
// steps from the customer, the SetupIntent will transition to the
// succeeded status.
//
// Otherwise, it will transition to the requires_action status and
// suggest additional actions via next_action. If setup fails,
// the SetupIntent will transition to the
// requires_payment_method status.
type SetupIntentConfirmParams struct {
	Params `form:"*"`
	// This hash contains details about the Mandate to create
	MandateData *SetupIntentMandateDataParams `form:"mandate_data"`
	// ID of the payment method (a PaymentMethod, Card, or saved Source object) to attach to this SetupIntent.
	PaymentMethod *string `form:"payment_method"`
	// When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method)
	// value in the SetupIntent.
	PaymentMethodData *SetupIntentConfirmPaymentMethodDataParams `form:"payment_method_data"`
	// Payment-method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	// The URL to redirect your customer back to after they authenticate on the payment method's app or site.
	// If you'd prefer to redirect to a mobile application, you can alternatively supply an application URI scheme.
	// This parameter is only used for cards and other redirect-based payment methods.
	ReturnURL *string `form:"return_url"`
}

// A SetupIntent object can be canceled when it is in one of these statuses: requires_payment_method, requires_confirmation, or requires_action.
//
// Once canceled, setup is abandoned and any operations on the SetupIntent will fail with an error.
type SetupIntentCancelParams struct {
	Params `form:"*"`
	// Reason for canceling this SetupIntent. Possible values are `abandoned`, `requested_by_customer`, or `duplicate`
	CancellationReason *string `form:"cancellation_reason"`
}

// Verifies microdeposits on a SetupIntent object.
type SetupIntentVerifyMicrodepositsParams struct {
	Params `form:"*"`
	// Two positive integers, in *cents*, equal to the values of the microdeposits sent to the bank account.
	Amounts []*int64 `form:"amounts"`
	// A six-character code starting with SM present in the microdeposit sent to the bank account.
	DescriptorCode *string `form:"descriptor_code"`
}
type SetupIntentNextActionRedirectToURL struct {
	// If the customer does not exit their browser while authenticating, they will be redirected to this specified URL after completion.
	ReturnURL string `json:"return_url"`
	// The URL you must redirect your customer to in order to authenticate.
	URL string `json:"url"`
}

// When confirming a SetupIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js.
type SetupIntentNextActionUseStripeSDK struct{}
type SetupIntentNextActionVerifyWithMicrodeposits struct {
	// The timestamp when the microdeposits are expected to land.
	ArrivalDate int64 `json:"arrival_date"`
	// The URL for the hosted verification page, which allows customers to verify their bank account.
	HostedVerificationURL string `json:"hosted_verification_url"`
	// The type of the microdeposit sent to the customer. Used to distinguish between different verification methods.
	MicrodepositType SetupIntentNextActionVerifyWithMicrodepositsMicrodepositType `json:"microdeposit_type"`
}

// If present, this property tells you what actions you need to take in order for your customer to continue payment setup.
type SetupIntentNextAction struct {
	RedirectToURL *SetupIntentNextActionRedirectToURL `json:"redirect_to_url"`
	// Type of the next action to perform, one of `redirect_to_url`, `use_stripe_sdk`, `alipay_handle_redirect`, `oxxo_display_details`, or `verify_with_microdeposits`.
	Type SetupIntentNextActionType `json:"type"`
	// When confirming a SetupIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js.
	UseStripeSDK            *SetupIntentNextActionUseStripeSDK            `json:"use_stripe_sdk"`
	VerifyWithMicrodeposits *SetupIntentNextActionVerifyWithMicrodeposits `json:"verify_with_microdeposits"`
}
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptions struct {
	// A URL for custom mandate text
	CustomMandateURL string `json:"custom_mandate_url"`
	// List of Stripe products where this mandate can be selected automatically.
	DefaultFor []SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor `json:"default_for"`
	// Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription string `json:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule `json:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}
type SetupIntentPaymentMethodOptionsACSSDebit struct {
	// See SetupIntentPaymentMethodOptionsACSSDebitCurrency for allowed values
	Currency       string                                                  `json:"currency"`
	MandateOptions *SetupIntentPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options"`
	// Bank account verification method.
	VerificationMethod SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}
type SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSession struct {
	// Amount of each recurring payment.
	Amount int64 `json:"amount"`
	// Currency of each recurring payment.
	Currency Currency `json:"currency"`
	// Frequency interval of each recurring payment.
	Interval SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSessionInterval `json:"interval"`
	// Frequency indicator of each recurring payment.
	IntervalCount int64 `json:"interval_count"`
}
type SetupIntentPaymentMethodOptionsBLIKMandateOptions struct {
	// Date at which the mandate expires.
	ExpiresAfter int64                                                        `json:"expires_after"`
	OffSession   *SetupIntentPaymentMethodOptionsBLIKMandateOptionsOffSession `json:"off_session"`
	// Type of the mandate.
	Type SetupIntentPaymentMethodOptionsBLIKMandateOptionsType `json:"type"`
}
type SetupIntentPaymentMethodOptionsBLIK struct {
	MandateOptions *SetupIntentPaymentMethodOptionsBLIKMandateOptions `json:"mandate_options"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type SetupIntentPaymentMethodOptionsCardMandateOptions struct {
	// Amount to be charged for future payments.
	Amount int64 `json:"amount"`
	// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
	AmountType SetupIntentPaymentMethodOptionsCardMandateOptionsAmountType `json:"amount_type"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A description of the mandate or subscription that is meant to be displayed to the customer.
	Description string `json:"description"`
	// End date of the mandate or subscription. If not provided, the mandate will be active until canceled. If provided, end date should be after start date.
	EndDate int64 `json:"end_date"`
	// Specifies payment frequency. One of `day`, `week`, `month`, `year`, or `sporadic`.
	Interval SetupIntentPaymentMethodOptionsCardMandateOptionsInterval `json:"interval"`
	// The number of intervals between payments. For example, `interval=month` and `interval_count=3` indicates one payment every three months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks). This parameter is optional when `interval=sporadic`.
	IntervalCount int64 `json:"interval_count"`
	// Unique identifier for the mandate or subscription.
	Reference string `json:"reference"`
	// Start date of the mandate or subscription. Start date should not be lesser than yesterday.
	StartDate int64 `json:"start_date"`
	// Specifies the type of mandates supported. Possible values are `india`.
	SupportedTypes []SetupIntentPaymentMethodOptionsCardMandateOptionsSupportedType `json:"supported_types"`
}
type SetupIntentPaymentMethodOptionsCard struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *SetupIntentPaymentMethodOptionsCardMandateOptions `json:"mandate_options"`
	// Selected network to process this SetupIntent on. Depends on the available networks of the card attached to the setup intent. Can be only set confirm-time.
	Network SetupIntentPaymentMethodOptionsCardNetwork `json:"network"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Permitted values include: `automatic` or `any`. If not provided, defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure SetupIntentPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}
type SetupIntentPaymentMethodOptionsLink struct {
	// Token used for persistent Link logins.
	PersistentToken string `json:"persistent_token"`
}
type SetupIntentPaymentMethodOptionsSepaDebitMandateOptions struct{}
type SetupIntentPaymentMethodOptionsSepaDebit struct {
	MandateOptions *SetupIntentPaymentMethodOptionsSepaDebitMandateOptions `json:"mandate_options"`
}
type SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	// The list of permissions to request. The `payment_method` permission must be included.
	Permissions []SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL string `json:"return_url"`
}
type SetupIntentPaymentMethodOptionsUSBankAccount struct {
	FinancialConnections *SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	// Bank account verification method.
	VerificationMethod SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment-method-specific configuration for this SetupIntent.
type SetupIntentPaymentMethodOptions struct {
	ACSSDebit     *SetupIntentPaymentMethodOptionsACSSDebit     `json:"acss_debit"`
	BLIK          *SetupIntentPaymentMethodOptionsBLIK          `json:"blik"`
	Card          *SetupIntentPaymentMethodOptionsCard          `json:"card"`
	Link          *SetupIntentPaymentMethodOptionsLink          `json:"link"`
	SepaDebit     *SetupIntentPaymentMethodOptionsSepaDebit     `json:"sepa_debit"`
	USBankAccount *SetupIntentPaymentMethodOptionsUSBankAccount `json:"us_bank_account"`
}

// A SetupIntent guides you through the process of setting up and saving a customer's payment credentials for future payments.
// For example, you could use a SetupIntent to set up and save your customer's card without immediately collecting a payment.
// Later, you can use [PaymentIntents](https://stripe.com/docs/api#payment_intents) to drive the payment flow.
//
// Create a SetupIntent as soon as you're ready to collect your customer's payment credentials.
// Do not maintain long-lived, unconfirmed SetupIntents as they may no longer be valid.
// The SetupIntent then transitions through multiple [statuses](https://stripe.com/docs/payments/intents#intent-statuses) as it guides
// you through the setup process.
//
// Successful SetupIntents result in payment credentials that are optimized for future payments.
// For example, cardholders in [certain regions](https://stripe.com/guides/strong-customer-authentication) may need to be run through
// [Strong Customer Authentication](https://stripe.com/docs/strong-customer-authentication) at the time of payment method collection
// in order to streamline later [off-session payments](https://stripe.com/docs/payments/setup-intents).
// If the SetupIntent is used with a [Customer](https://stripe.com/docs/api#setup_intent_object-customer), upon success,
// it will automatically attach the resulting payment method to that Customer.
// We recommend using SetupIntents or [setup_future_usage](https://stripe.com/docs/api#payment_intent_object-setup_future_usage) on
// PaymentIntents to save payment methods in order to prevent saving invalid or unoptimized payment methods.
//
// By using SetupIntents, you ensure that your customers experience the minimum set of required friction,
// even as regulations change over time.
//
// Related guide: [Setup Intents API](https://stripe.com/docs/payments/setup-intents).
type SetupIntent struct {
	APIResource
	// ID of the Connect application that created the SetupIntent.
	Application *Application `json:"application"`
	// If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.
	//
	// It can only be used for this Stripe Account's own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.
	AttachToSelf bool `json:"attach_to_self"`
	// Reason for cancellation of this SetupIntent, one of `abandoned`, `requested_by_customer`, or `duplicate`.
	CancellationReason SetupIntentCancellationReason `json:"cancellation_reason"`
	// The client secret of this SetupIntent. Used for client-side retrieval using a publishable key.
	//
	// The client secret can be used to complete payment setup from your frontend. It should not be stored, logged, or exposed to anyone other than the customer. Make sure that you have TLS enabled on any page that includes the client secret.
	ClientSecret string `json:"client_secret"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// ID of the Customer this SetupIntent belongs to, if one exists.
	//
	// If present, the SetupIntent's payment method will be attached to the Customer on successful setup. Payment methods attached to other Customers cannot be used with this SetupIntent.
	Customer *Customer `json:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Indicates the directions of money movement for which this payment method is intended to be used.
	//
	// Include `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.
	FlowDirections []SetupIntentFlowDirection `json:"flow_directions"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The error encountered in the previous SetupIntent confirmation.
	LastSetupError *Error `json:"last_setup_error"`
	// The most recent SetupAttempt for this SetupIntent.
	LatestAttempt *SetupAttempt `json:"latest_attempt"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// ID of the multi use Mandate generated by the SetupIntent.
	Mandate *Mandate `json:"mandate"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// If present, this property tells you what actions you need to take in order for your customer to continue payment setup.
	NextAction *SetupIntentNextAction `json:"next_action"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account (if any) for which the setup is intended.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// ID of the payment method used with this SetupIntent.
	PaymentMethod *PaymentMethod `json:"payment_method"`
	// Payment-method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentPaymentMethodOptions `json:"payment_method_options"`
	// The list of payment method types (e.g. card) that this SetupIntent is allowed to set up.
	PaymentMethodTypes []string `json:"payment_method_types"`
	// ID of the single_use Mandate generated by the SetupIntent.
	SingleUseMandate *Mandate `json:"single_use_mandate"`
	// [Status](https://stripe.com/docs/payments/intents#intent-statuses) of this SetupIntent, one of `requires_payment_method`, `requires_confirmation`, `requires_action`, `processing`, `canceled`, or `succeeded`.
	Status SetupIntentStatus `json:"status"`
	// Indicates how the payment method is intended to be used in the future.
	//
	// Use `on_session` if you intend to only reuse the payment method when the customer is in your checkout flow. Use `off_session` if your customer may or may not be in your checkout flow. If not provided, this value defaults to `off_session`.
	Usage SetupIntentUsage `json:"usage"`
}

// SetupIntentList is a list of SetupIntents as retrieved from a list endpoint.
type SetupIntentList struct {
	APIResource
	ListMeta
	Data []*SetupIntent `json:"data"`
}

// UnmarshalJSON handles deserialization of a SetupIntent.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SetupIntent) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type setupIntent SetupIntent
	var v setupIntent
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SetupIntent(v)
	return nil
}
