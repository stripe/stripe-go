//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Controls whether this SetupIntent will accept redirect-based payment methods.
//
// Redirect-based payment methods may require your customer to be redirected to a payment method's app or site for authentication or additional steps. To [confirm](https://stripe.com/docs/api/setup_intents/confirm) this SetupIntent, you may be required to provide a `return_url` to redirect customers back to your site after they authenticate or complete the setup.
type SetupIntentAutomaticPaymentMethodsAllowRedirects string

// List of values that SetupIntentAutomaticPaymentMethodsAllowRedirects can take
const (
	SetupIntentAutomaticPaymentMethodsAllowRedirectsAlways SetupIntentAutomaticPaymentMethodsAllowRedirects = "always"
	SetupIntentAutomaticPaymentMethodsAllowRedirectsNever  SetupIntentAutomaticPaymentMethodsAllowRedirects = "never"
)

// Reason for cancellation of this SetupIntent, one of `abandoned`, `requested_by_customer`, or `duplicate`.
type SetupIntentCancellationReason string

// List of values that SetupIntentCancellationReason can take
const (
	SetupIntentCancellationReasonAbandoned           SetupIntentCancellationReason = "abandoned"
	SetupIntentCancellationReasonDuplicate           SetupIntentCancellationReason = "duplicate"
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

// Type of the next action to perform. Refer to the other child attributes under `next_action` for available values. Examples include: `redirect_to_url`, `use_stripe_sdk`, `alipay_handle_redirect`, `oxxo_display_details`, or `verify_with_microdeposits`.
type SetupIntentNextActionType string

// List of values that SetupIntentNextActionType can take
const (
	SetupIntentNextActionTypeRedirectToURL           SetupIntentNextActionType = "redirect_to_url"
	SetupIntentNextActionTypeUseStripeSDK            SetupIntentNextActionType = "use_stripe_sdk"
	SetupIntentNextActionTypeAlipayHandleRedirect    SetupIntentNextActionType = "alipay_handle_redirect"
	SetupIntentNextActionTypeOXXODisplayDetails      SetupIntentNextActionType = "oxxo_display_details"
	SetupIntentNextActionTypeVerifyWithMicrodeposits SetupIntentNextActionType = "verify_with_microdeposits"
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
	SetupIntentPaymentMethodOptionsCardNetworkEFTPOSAU        SetupIntentPaymentMethodOptionsCardNetwork = "eftpos_au"
	SetupIntentPaymentMethodOptionsCardNetworkGirocard        SetupIntentPaymentMethodOptionsCardNetwork = "girocard"
	SetupIntentPaymentMethodOptionsCardNetworkInterac         SetupIntentPaymentMethodOptionsCardNetwork = "interac"
	SetupIntentPaymentMethodOptionsCardNetworkJCB             SetupIntentPaymentMethodOptionsCardNetwork = "jcb"
	SetupIntentPaymentMethodOptionsCardNetworkLink            SetupIntentPaymentMethodOptionsCardNetwork = "link"
	SetupIntentPaymentMethodOptionsCardNetworkMastercard      SetupIntentPaymentMethodOptionsCardNetwork = "mastercard"
	SetupIntentPaymentMethodOptionsCardNetworkUnionpay        SetupIntentPaymentMethodOptionsCardNetwork = "unionpay"
	SetupIntentPaymentMethodOptionsCardNetworkUnknown         SetupIntentPaymentMethodOptionsCardNetwork = "unknown"
	SetupIntentPaymentMethodOptionsCardNetworkVisa            SetupIntentPaymentMethodOptionsCardNetwork = "visa"
)

// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type SetupIntentPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that SetupIntentPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAny       SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "any"
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAutomatic SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureChallenge SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "challenge"
)

// The account subcategories to use to filter for possible accounts to link. Valid subcategories are `checking` and `savings`.
type SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory string

// List of values that SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory can take
const (
	SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategoryChecking SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "checking"
	SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategorySavings  SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory = "savings"
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

// Data features requested to be retrieved upon account creation.
type SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch string

// List of values that SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch can take
const (
	SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchBalances     SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "balances"
	SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchOwnership    SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "ownership"
	SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetchTransactions SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch = "transactions"
)

// Mandate collection method
type SetupIntentPaymentMethodOptionsUSBankAccountMandateOptionsCollectionMethod string

// List of values that SetupIntentPaymentMethodOptionsUSBankAccountMandateOptionsCollectionMethod can take
const (
	SetupIntentPaymentMethodOptionsUSBankAccountMandateOptionsCollectionMethodPaper SetupIntentPaymentMethodOptionsUSBankAccountMandateOptionsCollectionMethod = "paper"
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
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return SetupIntents that associate with the specified payment method.
	PaymentMethod *string `form:"payment_method"`
}

// AddExpand appends a new field to expand.
func (p *SetupIntentListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// When you enable this parameter, this SetupIntent accepts payment methods that you enable in the Dashboard and that are compatible with its other parameters.
type SetupIntentAutomaticPaymentMethodsParams struct {
	// Controls whether this SetupIntent will accept redirect-based payment methods.
	//
	// Redirect-based payment methods may require your customer to be redirected to a payment method's app or site for authentication or additional steps. To [confirm](https://stripe.com/docs/api/setup_intents/confirm) this SetupIntent, you may be required to provide a `return_url` to redirect customers back to your site after they authenticate or complete the setup.
	AllowRedirects *string `form:"allow_redirects"`
	// Whether this feature is enabled.
	Enabled *bool `form:"enabled"`
}

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
	AcceptedAt *int64 `form:"accepted_at"`
	// If this is a Mandate accepted offline, this hash contains details about the offline acceptance.
	Offline *SetupIntentMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	// If this is a Mandate accepted online, this hash contains details about the online acceptance.
	Online *SetupIntentMandateDataCustomerAcceptanceOnlineParams `form:"online"`
	// The type of customer acceptance information included with the Mandate. One of `online` or `offline`.
	Type MandateCustomerAcceptanceType `form:"type"`
}

// This hash contains details about the mandate to create. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
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

// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
type SetupIntentPaymentMethodDataAlmaParams struct{}

// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
type SetupIntentPaymentMethodDataAmazonPayParams struct{}

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

// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
type SetupIntentPaymentMethodDataBillieParams struct{}

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
	// Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.
	TaxID *string `form:"tax_id"`
}

// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
type SetupIntentPaymentMethodDataBLIKParams struct{}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type SetupIntentPaymentMethodDataBoletoParams struct {
	// The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)
	TaxID *string `form:"tax_id"`
}

// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
type SetupIntentPaymentMethodDataCashAppParams struct{}

// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
type SetupIntentPaymentMethodDataCryptoParams struct{}

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
type SetupIntentPaymentMethodDataIDEALParams struct {
	// The customer's bank. Only use this parameter for existing customers. Don't use it for new customers.
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type SetupIntentPaymentMethodDataInteracPresentParams struct{}

// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
type SetupIntentPaymentMethodDataKakaoPayParams struct{}

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

// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
type SetupIntentPaymentMethodDataKrCardParams struct{}

// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
type SetupIntentPaymentMethodDataLinkParams struct{}

// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
type SetupIntentPaymentMethodDataMobilepayParams struct{}

// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
type SetupIntentPaymentMethodDataMultibancoParams struct{}

// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
type SetupIntentPaymentMethodDataNaverPayParams struct {
	// Whether to use Naver Pay points or a card to fund this transaction. If not provided, this defaults to `card`.
	Funding *string `form:"funding"`
}

// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
type SetupIntentPaymentMethodDataNzBankAccountParams struct {
	// The name on the bank account. Only required if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod's billing details.
	AccountHolderName *string `form:"account_holder_name"`
	// The account number for the bank account.
	AccountNumber *string `form:"account_number"`
	// The numeric code for the bank account's bank.
	BankCode *string `form:"bank_code"`
	// The numeric code for the bank account's bank branch.
	BranchCode *string `form:"branch_code"`
	Reference  *string `form:"reference"`
	// The suffix of the bank account number.
	Suffix *string `form:"suffix"`
}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type SetupIntentPaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type SetupIntentPaymentMethodDataP24Params struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
type SetupIntentPaymentMethodDataPayByBankParams struct{}

// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
type SetupIntentPaymentMethodDataPaycoParams struct{}

// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
type SetupIntentPaymentMethodDataPayNowParams struct{}

// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
type SetupIntentPaymentMethodDataPaypalParams struct{}

// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
type SetupIntentPaymentMethodDataPixParams struct{}

// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
type SetupIntentPaymentMethodDataPromptPayParams struct{}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type SetupIntentPaymentMethodDataRadarOptionsParams struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
type SetupIntentPaymentMethodDataRevolutPayParams struct{}

// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
type SetupIntentPaymentMethodDataSamsungPayParams struct{}

// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
type SetupIntentPaymentMethodDataSatispayParams struct{}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type SetupIntentPaymentMethodDataSEPADebitParams struct {
	// IBAN of the bank account.
	IBAN *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type SetupIntentPaymentMethodDataSofortParams struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country *string `form:"country"`
}

// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
type SetupIntentPaymentMethodDataSwishParams struct{}

// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
type SetupIntentPaymentMethodDataTWINTParams struct{}

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
type SetupIntentPaymentMethodDataWeChatPayParams struct{}

// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
type SetupIntentPaymentMethodDataZipParams struct{}

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
	// This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow. The field defaults to `unspecified`.
	AllowRedisplay *string `form:"allow_redisplay"`
	// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
	Alma *SetupIntentPaymentMethodDataAlmaParams `form:"alma"`
	// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
	AmazonPay *SetupIntentPaymentMethodDataAmazonPayParams `form:"amazon_pay"`
	// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
	AUBECSDebit *SetupIntentPaymentMethodDataAUBECSDebitParams `form:"au_becs_debit"`
	// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
	BACSDebit *SetupIntentPaymentMethodDataBACSDebitParams `form:"bacs_debit"`
	// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
	Bancontact *SetupIntentPaymentMethodDataBancontactParams `form:"bancontact"`
	// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
	Billie *SetupIntentPaymentMethodDataBillieParams `form:"billie"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *SetupIntentPaymentMethodDataBillingDetailsParams `form:"billing_details"`
	// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
	BLIK *SetupIntentPaymentMethodDataBLIKParams `form:"blik"`
	// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
	Boleto *SetupIntentPaymentMethodDataBoletoParams `form:"boleto"`
	// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
	CashApp *SetupIntentPaymentMethodDataCashAppParams `form:"cashapp"`
	// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
	Crypto *SetupIntentPaymentMethodDataCryptoParams `form:"crypto"`
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
	IDEAL *SetupIntentPaymentMethodDataIDEALParams `form:"ideal"`
	// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
	InteracPresent *SetupIntentPaymentMethodDataInteracPresentParams `form:"interac_present"`
	// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
	KakaoPay *SetupIntentPaymentMethodDataKakaoPayParams `form:"kakao_pay"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
	Klarna *SetupIntentPaymentMethodDataKlarnaParams `form:"klarna"`
	// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
	Konbini *SetupIntentPaymentMethodDataKonbiniParams `form:"konbini"`
	// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
	KrCard *SetupIntentPaymentMethodDataKrCardParams `form:"kr_card"`
	// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
	Link *SetupIntentPaymentMethodDataLinkParams `form:"link"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
	Mobilepay *SetupIntentPaymentMethodDataMobilepayParams `form:"mobilepay"`
	// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
	Multibanco *SetupIntentPaymentMethodDataMultibancoParams `form:"multibanco"`
	// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
	NaverPay *SetupIntentPaymentMethodDataNaverPayParams `form:"naver_pay"`
	// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
	NzBankAccount *SetupIntentPaymentMethodDataNzBankAccountParams `form:"nz_bank_account"`
	// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
	OXXO *SetupIntentPaymentMethodDataOXXOParams `form:"oxxo"`
	// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
	P24 *SetupIntentPaymentMethodDataP24Params `form:"p24"`
	// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
	PayByBank *SetupIntentPaymentMethodDataPayByBankParams `form:"pay_by_bank"`
	// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
	Payco *SetupIntentPaymentMethodDataPaycoParams `form:"payco"`
	// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
	PayNow *SetupIntentPaymentMethodDataPayNowParams `form:"paynow"`
	// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
	Paypal *SetupIntentPaymentMethodDataPaypalParams `form:"paypal"`
	// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
	Pix *SetupIntentPaymentMethodDataPixParams `form:"pix"`
	// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
	PromptPay *SetupIntentPaymentMethodDataPromptPayParams `form:"promptpay"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *SetupIntentPaymentMethodDataRadarOptionsParams `form:"radar_options"`
	// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
	RevolutPay *SetupIntentPaymentMethodDataRevolutPayParams `form:"revolut_pay"`
	// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
	SamsungPay *SetupIntentPaymentMethodDataSamsungPayParams `form:"samsung_pay"`
	// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
	Satispay *SetupIntentPaymentMethodDataSatispayParams `form:"satispay"`
	// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
	SEPADebit *SetupIntentPaymentMethodDataSEPADebitParams `form:"sepa_debit"`
	// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
	Sofort *SetupIntentPaymentMethodDataSofortParams `form:"sofort"`
	// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
	Swish *SetupIntentPaymentMethodDataSwishParams `form:"swish"`
	// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
	TWINT *SetupIntentPaymentMethodDataTWINTParams `form:"twint"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type"`
	// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
	USBankAccount *SetupIntentPaymentMethodDataUSBankAccountParams `form:"us_bank_account"`
	// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
	WeChatPay *SetupIntentPaymentMethodDataWeChatPayParams `form:"wechat_pay"`
	// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
	Zip *SetupIntentPaymentMethodDataZipParams `form:"zip"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SetupIntentPaymentMethodDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
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
	// Bank account verification method.
	VerificationMethod *string `form:"verification_method"`
}

// If this is a `amazon_pay` SetupIntent, this sub-hash contains details about the AmazonPay payment method options.
type SetupIntentPaymentMethodOptionsAmazonPayParams struct{}

// Additional fields for Mandate creation
type SetupIntentPaymentMethodOptionsBACSDebitMandateOptionsParams struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'DDIC' or 'STRIPE'.
	ReferencePrefix *string `form:"reference_prefix"`
}

// If this is a `bacs_debit` SetupIntent, this sub-hash contains details about the Bacs Debit payment method options.
type SetupIntentPaymentMethodOptionsBACSDebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentPaymentMethodOptionsBACSDebitMandateOptionsParams `form:"mandate_options"`
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

// Cartes Bancaires-specific 3DS fields.
type SetupIntentPaymentMethodOptionsCardThreeDSecureNetworkOptionsCartesBancairesParams struct {
	// The cryptogram calculation algorithm used by the card Issuer's ACS
	// to calculate the Authentication cryptogram. Also known as `cavvAlgorithm`.
	// messageExtension: CB-AVALGO
	CbAvalgo *string `form:"cb_avalgo"`
	// The exemption indicator returned from Cartes Bancaires in the ARes.
	// message extension: CB-EXEMPTION; string (4 characters)
	// This is a 3 byte bitmap (low significant byte first and most significant
	// bit first) that has been Base64 encoded
	CbExemption *string `form:"cb_exemption"`
	// The risk score returned from Cartes Bancaires in the ARes.
	// message extension: CB-SCORE; numeric value 0-99
	CbScore *int64 `form:"cb_score"`
}

// Network specific 3DS fields. Network specific arguments require an
// explicit card brand choice. The parameter `payment_method_options.card.network“
// must be populated accordingly
type SetupIntentPaymentMethodOptionsCardThreeDSecureNetworkOptionsParams struct {
	// Cartes Bancaires-specific 3DS fields.
	CartesBancaires *SetupIntentPaymentMethodOptionsCardThreeDSecureNetworkOptionsCartesBancairesParams `form:"cartes_bancaires"`
}

// If 3D Secure authentication was performed with a third-party provider,
// the authentication details to use for this setup.
type SetupIntentPaymentMethodOptionsCardThreeDSecureParams struct {
	// The `transStatus` returned from the card Issuer's ACS in the ARes.
	AresTransStatus *string `form:"ares_trans_status"`
	// The cryptogram, also known as the "authentication value" (AAV, CAVV or
	// AEVV). This value is 20 bytes, base64-encoded into a 28-character string.
	// (Most 3D Secure providers will return the base64-encoded version, which
	// is what you should specify here.)
	Cryptogram *string `form:"cryptogram"`
	// The Electronic Commerce Indicator (ECI) is returned by your 3D Secure
	// provider and indicates what degree of authentication was performed.
	ElectronicCommerceIndicator *string `form:"electronic_commerce_indicator"`
	// Network specific 3DS fields. Network specific arguments require an
	// explicit card brand choice. The parameter `payment_method_options.card.network``
	// must be populated accordingly
	NetworkOptions *SetupIntentPaymentMethodOptionsCardThreeDSecureNetworkOptionsParams `form:"network_options"`
	// The challenge indicator (`threeDSRequestorChallengeInd`) which was requested in the
	// AReq sent to the card Issuer's ACS. A string containing 2 digits from 01-99.
	RequestorChallengeIndicator *string `form:"requestor_challenge_indicator"`
	// For 3D Secure 1, the XID. For 3D Secure 2, the Directory Server
	// Transaction ID (dsTransID).
	TransactionID *string `form:"transaction_id"`
	// The version of 3D Secure that was performed.
	Version *string `form:"version"`
}

// Configuration for any card setup attempted on this SetupIntent.
type SetupIntentPaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *SetupIntentPaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options"`
	// When specified, this parameter signals that a card has been collected
	// as MOTO (Mail Order Telephone Order) and thus out of scope for SCA. This
	// parameter can only be provided during confirmation.
	MOTO *bool `form:"moto"`
	// Selected network to process this SetupIntent on. Depends on the available networks of the card attached to the SetupIntent. Can be only set confirm-time.
	Network *string `form:"network"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure"`
	// If 3D Secure authentication was performed with a third-party provider,
	// the authentication details to use for this setup.
	ThreeDSecure *SetupIntentPaymentMethodOptionsCardThreeDSecureParams `form:"three_d_secure"`
}

// If this is a `card_present` PaymentMethod, this sub-hash contains details about the card-present payment method options.
type SetupIntentPaymentMethodOptionsCardPresentParams struct{}

// On-demand details if setting up a payment method for on-demand payments.
type SetupIntentPaymentMethodOptionsKlarnaOnDemandParams struct {
	// Your average amount value. You can use a value across your customer base, or segment based on customer type, country, etc.
	AverageAmount *int64 `form:"average_amount"`
	// The maximum value you may charge a customer per purchase. You can use a value across your customer base, or segment based on customer type, country, etc.
	MaximumAmount *int64 `form:"maximum_amount"`
	// The lowest or minimum value you may charge a customer per purchase. You can use a value across your customer base, or segment based on customer type, country, etc.
	MinimumAmount *int64 `form:"minimum_amount"`
	// Interval at which the customer is making purchases
	PurchaseInterval *string `form:"purchase_interval"`
	// The number of `purchase_interval` between charges
	PurchaseIntervalCount *int64 `form:"purchase_interval_count"`
}

// Describes the upcoming charge for this subscription.
type SetupIntentPaymentMethodOptionsKlarnaSubscriptionNextBillingParams struct {
	// The amount of the next charge for the subscription.
	Amount *int64 `form:"amount"`
	// The date of the next charge for the subscription in YYYY-MM-DD format.
	Date *string `form:"date"`
}

// Subscription details if setting up or charging a subscription
type SetupIntentPaymentMethodOptionsKlarnaSubscriptionParams struct {
	// Unit of time between subscription charges.
	Interval *string `form:"interval"`
	// The number of intervals (specified in the `interval` attribute) between subscription charges. For example, `interval=month` and `interval_count=3` charges every 3 months.
	IntervalCount *int64 `form:"interval_count"`
	// Name for subscription.
	Name *string `form:"name"`
	// Describes the upcoming charge for this subscription.
	NextBilling *SetupIntentPaymentMethodOptionsKlarnaSubscriptionNextBillingParams `form:"next_billing"`
	// A non-customer-facing reference to correlate subscription charges in the Klarna app. Use a value that persists across subscription charges.
	Reference *string `form:"reference"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method options.
type SetupIntentPaymentMethodOptionsKlarnaParams struct {
	// The currency of the SetupIntent. Three letter ISO currency code.
	Currency *string `form:"currency"`
	// On-demand details if setting up a payment method for on-demand payments.
	OnDemand *SetupIntentPaymentMethodOptionsKlarnaOnDemandParams `form:"on_demand"`
	// Preferred language of the Klarna authorization page that the customer is redirected to
	PreferredLocale *string `form:"preferred_locale"`
	// Subscription details if setting up or charging a subscription
	Subscriptions []*SetupIntentPaymentMethodOptionsKlarnaSubscriptionParams `form:"subscriptions"`
}

// If this is a `link` PaymentMethod, this sub-hash contains details about the Link payment method options.
type SetupIntentPaymentMethodOptionsLinkParams struct {
	// [Deprecated] This is a legacy parameter that no longer has any function.
	// Deprecated:
	PersistentToken *string `form:"persistent_token"`
}

// If this is a `paypal` PaymentMethod, this sub-hash contains details about the PayPal payment method options.
type SetupIntentPaymentMethodOptionsPaypalParams struct {
	// The PayPal Billing Agreement ID (BAID). This is an ID generated by PayPal which represents the mandate between the merchant and the customer.
	BillingAgreementID *string `form:"billing_agreement_id"`
}

// Additional fields for Mandate creation
type SetupIntentPaymentMethodOptionsSEPADebitMandateOptionsParams struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'STRIPE'.
	ReferencePrefix *string `form:"reference_prefix"`
}

// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
type SetupIntentPaymentMethodOptionsSEPADebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentPaymentMethodOptionsSEPADebitMandateOptionsParams `form:"mandate_options"`
}

// Provide filters for the linked accounts that the customer can select for the payment method.
type SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts. Valid subcategories are `checking` and `savings`.
	AccountSubcategories []*string `form:"account_subcategories"`
}

// Additional fields for Financial Connections Session creation
type SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included. Valid permissions include: `balances`, `ownership`, `payment_method`, and `transactions`.
	Permissions []*string `form:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL *string `form:"return_url"`
}

// Additional fields for Mandate creation
type SetupIntentPaymentMethodOptionsUSBankAccountMandateOptionsParams struct {
	// The method used to collect offline mandate customer acceptance.
	CollectionMethod *string `form:"collection_method"`
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
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentPaymentMethodOptionsUSBankAccountMandateOptionsParams `form:"mandate_options"`
	// Additional fields for network related functions
	Networks *SetupIntentPaymentMethodOptionsUSBankAccountNetworksParams `form:"networks"`
	// Bank account verification method.
	VerificationMethod *string `form:"verification_method"`
}

// Payment method-specific configuration for this SetupIntent.
type SetupIntentPaymentMethodOptionsParams struct {
	// If this is a `acss_debit` SetupIntent, this sub-hash contains details about the ACSS Debit payment method options.
	ACSSDebit *SetupIntentPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	// If this is a `amazon_pay` SetupIntent, this sub-hash contains details about the AmazonPay payment method options.
	AmazonPay *SetupIntentPaymentMethodOptionsAmazonPayParams `form:"amazon_pay"`
	// If this is a `bacs_debit` SetupIntent, this sub-hash contains details about the Bacs Debit payment method options.
	BACSDebit *SetupIntentPaymentMethodOptionsBACSDebitParams `form:"bacs_debit"`
	// Configuration for any card setup attempted on this SetupIntent.
	Card *SetupIntentPaymentMethodOptionsCardParams `form:"card"`
	// If this is a `card_present` PaymentMethod, this sub-hash contains details about the card-present payment method options.
	CardPresent *SetupIntentPaymentMethodOptionsCardPresentParams `form:"card_present"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method options.
	Klarna *SetupIntentPaymentMethodOptionsKlarnaParams `form:"klarna"`
	// If this is a `link` PaymentMethod, this sub-hash contains details about the Link payment method options.
	Link *SetupIntentPaymentMethodOptionsLinkParams `form:"link"`
	// If this is a `paypal` PaymentMethod, this sub-hash contains details about the PayPal payment method options.
	Paypal *SetupIntentPaymentMethodOptionsPaypalParams `form:"paypal"`
	// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
	SEPADebit *SetupIntentPaymentMethodOptionsSEPADebitParams `form:"sepa_debit"`
	// If this is a `us_bank_account` SetupIntent, this sub-hash contains details about the US bank account payment method options.
	USBankAccount *SetupIntentPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account"`
}

// If you populate this hash, this SetupIntent generates a `single_use` mandate after successful completion.
//
// Single-use mandates are only valid for the following payment methods: `acss_debit`, `alipay`, `au_becs_debit`, `bacs_debit`, `bancontact`, `boleto`, `ideal`, `link`, `sepa_debit`, and `us_bank_account`.
type SetupIntentSingleUseParams struct {
	// Amount the customer is granting permission to collect later. A positive integer representing how much to charge in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://stripe.com/docs/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
}

// Creates a SetupIntent object.
//
// After you create the SetupIntent, attach a payment method and [confirm](https://docs.stripe.com/docs/api/setup_intents/confirm)
// it to collect any required permissions to charge the payment method later.
type SetupIntentParams struct {
	Params `form:"*"`
	// If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.
	//
	// It can only be used for this Stripe Account's own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.
	AttachToSelf *bool `form:"attach_to_self"`
	// When you enable this parameter, this SetupIntent accepts payment methods that you enable in the Dashboard and that are compatible with its other parameters.
	AutomaticPaymentMethods *SetupIntentAutomaticPaymentMethodsParams `form:"automatic_payment_methods"`
	// The client secret of the SetupIntent. We require this string if you use a publishable key to retrieve the SetupIntent.
	ClientSecret *string `form:"client_secret"`
	// Set to `true` to attempt to confirm this SetupIntent immediately. This parameter defaults to `false`. If a card is the attached payment method, you can provide a `return_url` in case further authentication is necessary.
	Confirm *bool `form:"confirm"`
	// ID of the ConfirmationToken used to confirm this SetupIntent.
	//
	// If the provided ConfirmationToken contains properties that are also being provided in this request, such as `payment_method`, then the values in this request will take precedence.
	ConfirmationToken *string `form:"confirmation_token"`
	// ID of the Customer this SetupIntent belongs to, if one exists.
	//
	// If present, the SetupIntent's payment method will be attached to the Customer on successful setup. Payment methods attached to other Customers cannot be used with this SetupIntent.
	Customer *string `form:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Indicates the directions of money movement for which this payment method is intended to be used.
	//
	// Include `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.
	FlowDirections []*string `form:"flow_directions"`
	// This hash contains details about the mandate to create. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
	MandateData *SetupIntentMandateDataParams `form:"mandate_data"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The Stripe account ID created for this SetupIntent.
	OnBehalfOf *string `form:"on_behalf_of"`
	// ID of the payment method (a PaymentMethod, Card, or saved Source object) to attach to this SetupIntent. To unset this field to null, pass in an empty string.
	PaymentMethod *string `form:"payment_method"`
	// The ID of the [payment method configuration](https://stripe.com/docs/api/payment_method_configurations) to use with this SetupIntent.
	PaymentMethodConfiguration *string `form:"payment_method_configuration"`
	// When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method)
	// value in the SetupIntent.
	PaymentMethodData *SetupIntentPaymentMethodDataParams `form:"payment_method_data"`
	// Payment method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	// The list of payment method types (for example, card) that this SetupIntent can set up. If you don't provide this, Stripe will dynamically show relevant payment methods from your [payment method settings](https://dashboard.stripe.com/settings/payment_methods). A list of valid payment method types can be found [here](https://docs.stripe.com/api/payment_methods/object#payment_method_object-type).
	PaymentMethodTypes []*string `form:"payment_method_types"`
	// The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method's app or site. To redirect to a mobile application, you can alternatively supply an application URI scheme. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
	ReturnURL *string `form:"return_url"`
	// If you populate this hash, this SetupIntent generates a `single_use` mandate after successful completion.
	//
	// Single-use mandates are only valid for the following payment methods: `acss_debit`, `alipay`, `au_becs_debit`, `bacs_debit`, `bancontact`, `boleto`, `ideal`, `link`, `sepa_debit`, and `us_bank_account`.
	SingleUse *SetupIntentSingleUseParams `form:"single_use"`
	// Indicates how the payment method is intended to be used in the future. If not provided, this value defaults to `off_session`.
	Usage *string `form:"usage"`
	// Set to `true` when confirming server-side and using Stripe.js, iOS, or Android client-side SDKs to handle the next actions.
	UseStripeSDK *bool `form:"use_stripe_sdk"`
}

// AddExpand appends a new field to expand.
func (p *SetupIntentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SetupIntentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// You can cancel a SetupIntent object when it's in one of these statuses: requires_payment_method, requires_confirmation, or requires_action.
//
// After you cancel it, setup is abandoned and any operations on the SetupIntent fail with an error. You can't cancel the SetupIntent for a Checkout Session. [Expire the Checkout Session](https://docs.stripe.com/docs/api/checkout/sessions/expire) instead.
type SetupIntentCancelParams struct {
	Params `form:"*"`
	// Reason for canceling this SetupIntent. Possible values are: `abandoned`, `requested_by_customer`, or `duplicate`
	CancellationReason *string `form:"cancellation_reason"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *SetupIntentCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
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

// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
type SetupIntentConfirmPaymentMethodDataAlmaParams struct{}

// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
type SetupIntentConfirmPaymentMethodDataAmazonPayParams struct{}

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

// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
type SetupIntentConfirmPaymentMethodDataBillieParams struct{}

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
	// Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.
	TaxID *string `form:"tax_id"`
}

// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
type SetupIntentConfirmPaymentMethodDataBLIKParams struct{}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type SetupIntentConfirmPaymentMethodDataBoletoParams struct {
	// The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)
	TaxID *string `form:"tax_id"`
}

// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
type SetupIntentConfirmPaymentMethodDataCashAppParams struct{}

// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
type SetupIntentConfirmPaymentMethodDataCryptoParams struct{}

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
type SetupIntentConfirmPaymentMethodDataIDEALParams struct {
	// The customer's bank. Only use this parameter for existing customers. Don't use it for new customers.
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type SetupIntentConfirmPaymentMethodDataInteracPresentParams struct{}

// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
type SetupIntentConfirmPaymentMethodDataKakaoPayParams struct{}

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

// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
type SetupIntentConfirmPaymentMethodDataKrCardParams struct{}

// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
type SetupIntentConfirmPaymentMethodDataLinkParams struct{}

// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
type SetupIntentConfirmPaymentMethodDataMobilepayParams struct{}

// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
type SetupIntentConfirmPaymentMethodDataMultibancoParams struct{}

// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
type SetupIntentConfirmPaymentMethodDataNaverPayParams struct {
	// Whether to use Naver Pay points or a card to fund this transaction. If not provided, this defaults to `card`.
	Funding *string `form:"funding"`
}

// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
type SetupIntentConfirmPaymentMethodDataNzBankAccountParams struct {
	// The name on the bank account. Only required if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod's billing details.
	AccountHolderName *string `form:"account_holder_name"`
	// The account number for the bank account.
	AccountNumber *string `form:"account_number"`
	// The numeric code for the bank account's bank.
	BankCode *string `form:"bank_code"`
	// The numeric code for the bank account's bank branch.
	BranchCode *string `form:"branch_code"`
	Reference  *string `form:"reference"`
	// The suffix of the bank account number.
	Suffix *string `form:"suffix"`
}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type SetupIntentConfirmPaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type SetupIntentConfirmPaymentMethodDataP24Params struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
type SetupIntentConfirmPaymentMethodDataPayByBankParams struct{}

// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
type SetupIntentConfirmPaymentMethodDataPaycoParams struct{}

// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
type SetupIntentConfirmPaymentMethodDataPayNowParams struct{}

// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
type SetupIntentConfirmPaymentMethodDataPaypalParams struct{}

// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
type SetupIntentConfirmPaymentMethodDataPixParams struct{}

// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
type SetupIntentConfirmPaymentMethodDataPromptPayParams struct{}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type SetupIntentConfirmPaymentMethodDataRadarOptionsParams struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
type SetupIntentConfirmPaymentMethodDataRevolutPayParams struct{}

// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
type SetupIntentConfirmPaymentMethodDataSamsungPayParams struct{}

// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
type SetupIntentConfirmPaymentMethodDataSatispayParams struct{}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type SetupIntentConfirmPaymentMethodDataSEPADebitParams struct {
	// IBAN of the bank account.
	IBAN *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type SetupIntentConfirmPaymentMethodDataSofortParams struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country *string `form:"country"`
}

// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
type SetupIntentConfirmPaymentMethodDataSwishParams struct{}

// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
type SetupIntentConfirmPaymentMethodDataTWINTParams struct{}

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
type SetupIntentConfirmPaymentMethodDataWeChatPayParams struct{}

// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
type SetupIntentConfirmPaymentMethodDataZipParams struct{}

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
	// This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow. The field defaults to `unspecified`.
	AllowRedisplay *string `form:"allow_redisplay"`
	// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
	Alma *SetupIntentConfirmPaymentMethodDataAlmaParams `form:"alma"`
	// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
	AmazonPay *SetupIntentConfirmPaymentMethodDataAmazonPayParams `form:"amazon_pay"`
	// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
	AUBECSDebit *SetupIntentConfirmPaymentMethodDataAUBECSDebitParams `form:"au_becs_debit"`
	// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
	BACSDebit *SetupIntentConfirmPaymentMethodDataBACSDebitParams `form:"bacs_debit"`
	// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
	Bancontact *SetupIntentConfirmPaymentMethodDataBancontactParams `form:"bancontact"`
	// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
	Billie *SetupIntentConfirmPaymentMethodDataBillieParams `form:"billie"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *SetupIntentConfirmPaymentMethodDataBillingDetailsParams `form:"billing_details"`
	// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
	BLIK *SetupIntentConfirmPaymentMethodDataBLIKParams `form:"blik"`
	// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
	Boleto *SetupIntentConfirmPaymentMethodDataBoletoParams `form:"boleto"`
	// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
	CashApp *SetupIntentConfirmPaymentMethodDataCashAppParams `form:"cashapp"`
	// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
	Crypto *SetupIntentConfirmPaymentMethodDataCryptoParams `form:"crypto"`
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
	IDEAL *SetupIntentConfirmPaymentMethodDataIDEALParams `form:"ideal"`
	// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
	InteracPresent *SetupIntentConfirmPaymentMethodDataInteracPresentParams `form:"interac_present"`
	// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
	KakaoPay *SetupIntentConfirmPaymentMethodDataKakaoPayParams `form:"kakao_pay"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
	Klarna *SetupIntentConfirmPaymentMethodDataKlarnaParams `form:"klarna"`
	// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
	Konbini *SetupIntentConfirmPaymentMethodDataKonbiniParams `form:"konbini"`
	// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
	KrCard *SetupIntentConfirmPaymentMethodDataKrCardParams `form:"kr_card"`
	// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
	Link *SetupIntentConfirmPaymentMethodDataLinkParams `form:"link"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
	Mobilepay *SetupIntentConfirmPaymentMethodDataMobilepayParams `form:"mobilepay"`
	// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
	Multibanco *SetupIntentConfirmPaymentMethodDataMultibancoParams `form:"multibanco"`
	// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
	NaverPay *SetupIntentConfirmPaymentMethodDataNaverPayParams `form:"naver_pay"`
	// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
	NzBankAccount *SetupIntentConfirmPaymentMethodDataNzBankAccountParams `form:"nz_bank_account"`
	// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
	OXXO *SetupIntentConfirmPaymentMethodDataOXXOParams `form:"oxxo"`
	// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
	P24 *SetupIntentConfirmPaymentMethodDataP24Params `form:"p24"`
	// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
	PayByBank *SetupIntentConfirmPaymentMethodDataPayByBankParams `form:"pay_by_bank"`
	// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
	Payco *SetupIntentConfirmPaymentMethodDataPaycoParams `form:"payco"`
	// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
	PayNow *SetupIntentConfirmPaymentMethodDataPayNowParams `form:"paynow"`
	// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
	Paypal *SetupIntentConfirmPaymentMethodDataPaypalParams `form:"paypal"`
	// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
	Pix *SetupIntentConfirmPaymentMethodDataPixParams `form:"pix"`
	// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
	PromptPay *SetupIntentConfirmPaymentMethodDataPromptPayParams `form:"promptpay"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *SetupIntentConfirmPaymentMethodDataRadarOptionsParams `form:"radar_options"`
	// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
	RevolutPay *SetupIntentConfirmPaymentMethodDataRevolutPayParams `form:"revolut_pay"`
	// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
	SamsungPay *SetupIntentConfirmPaymentMethodDataSamsungPayParams `form:"samsung_pay"`
	// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
	Satispay *SetupIntentConfirmPaymentMethodDataSatispayParams `form:"satispay"`
	// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
	SEPADebit *SetupIntentConfirmPaymentMethodDataSEPADebitParams `form:"sepa_debit"`
	// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
	Sofort *SetupIntentConfirmPaymentMethodDataSofortParams `form:"sofort"`
	// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
	Swish *SetupIntentConfirmPaymentMethodDataSwishParams `form:"swish"`
	// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
	TWINT *SetupIntentConfirmPaymentMethodDataTWINTParams `form:"twint"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type"`
	// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
	USBankAccount *SetupIntentConfirmPaymentMethodDataUSBankAccountParams `form:"us_bank_account"`
	// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
	WeChatPay *SetupIntentConfirmPaymentMethodDataWeChatPayParams `form:"wechat_pay"`
	// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
	Zip *SetupIntentConfirmPaymentMethodDataZipParams `form:"zip"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SetupIntentConfirmPaymentMethodDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
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
// requires_payment_method status or the canceled status if the
// confirmation limit is reached.
type SetupIntentConfirmParams struct {
	Params `form:"*"`
	// ID of the ConfirmationToken used to confirm this SetupIntent.
	//
	// If the provided ConfirmationToken contains properties that are also being provided in this request, such as `payment_method`, then the values in this request will take precedence.
	ConfirmationToken *string `form:"confirmation_token"`
	// Specifies which fields in the response should be expanded.
	Expand      []*string                     `form:"expand"`
	MandateData *SetupIntentMandateDataParams `form:"mandate_data"`
	// ID of the payment method (a PaymentMethod, Card, or saved Source object) to attach to this SetupIntent.
	PaymentMethod *string `form:"payment_method"`
	// When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method)
	// value in the SetupIntent.
	PaymentMethodData *SetupIntentConfirmPaymentMethodDataParams `form:"payment_method_data"`
	// Payment method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	// The URL to redirect your customer back to after they authenticate on the payment method's app or site.
	// If you'd prefer to redirect to a mobile application, you can alternatively supply an application URI scheme.
	// This parameter is only used for cards and other redirect-based payment methods.
	ReturnURL *string `form:"return_url"`
	// Set to `true` when confirming server-side and using Stripe.js, iOS, or Android client-side SDKs to handle the next actions.
	UseStripeSDK *bool `form:"use_stripe_sdk"`
}

// AddExpand appends a new field to expand.
func (p *SetupIntentConfirmParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Verifies microdeposits on a SetupIntent object.
type SetupIntentVerifyMicrodepositsParams struct {
	Params `form:"*"`
	// Two positive integers, in *cents*, equal to the values of the microdeposits sent to the bank account.
	Amounts []*int64 `form:"amounts"`
	// A six-character code starting with SM present in the microdeposit sent to the bank account.
	DescriptorCode *string `form:"descriptor_code"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *SetupIntentVerifyMicrodepositsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// When you enable this parameter, this SetupIntent accepts payment methods that you enable in the Dashboard and that are compatible with its other parameters.
type SetupIntentCreateAutomaticPaymentMethodsParams struct {
	// Controls whether this SetupIntent will accept redirect-based payment methods.
	//
	// Redirect-based payment methods may require your customer to be redirected to a payment method's app or site for authentication or additional steps. To [confirm](https://stripe.com/docs/api/setup_intents/confirm) this SetupIntent, you may be required to provide a `return_url` to redirect customers back to your site after they authenticate or complete the setup.
	AllowRedirects *string `form:"allow_redirects"`
	// Whether this feature is enabled.
	Enabled *bool `form:"enabled"`
}

// If this is a Mandate accepted offline, this hash contains details about the offline acceptance.
type SetupIntentCreateMandateDataCustomerAcceptanceOfflineParams struct{}

// If this is a Mandate accepted online, this hash contains details about the online acceptance.
type SetupIntentCreateMandateDataCustomerAcceptanceOnlineParams struct {
	// The IP address from which the Mandate was accepted by the customer.
	IPAddress *string `form:"ip_address"`
	// The user agent of the browser from which the Mandate was accepted by the customer.
	UserAgent *string `form:"user_agent"`
}

// This hash contains details about the customer acceptance of the Mandate.
type SetupIntentCreateMandateDataCustomerAcceptanceParams struct {
	// The time at which the customer accepted the Mandate.
	AcceptedAt *int64 `form:"accepted_at"`
	// If this is a Mandate accepted offline, this hash contains details about the offline acceptance.
	Offline *SetupIntentCreateMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	// If this is a Mandate accepted online, this hash contains details about the online acceptance.
	Online *SetupIntentCreateMandateDataCustomerAcceptanceOnlineParams `form:"online"`
	// The type of customer acceptance information included with the Mandate. One of `online` or `offline`.
	Type *string `form:"type"`
}

// This hash contains details about the mandate to create. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
type SetupIntentCreateMandateDataParams struct {
	// This hash contains details about the customer acceptance of the Mandate.
	CustomerAcceptance *SetupIntentCreateMandateDataCustomerAcceptanceParams `form:"customer_acceptance"`
}

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type SetupIntentCreatePaymentMethodDataACSSDebitParams struct {
	// Customer's bank account number.
	AccountNumber *string `form:"account_number"`
	// Institution number of the customer's bank.
	InstitutionNumber *string `form:"institution_number"`
	// Transit number of the customer's bank.
	TransitNumber *string `form:"transit_number"`
}

// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
type SetupIntentCreatePaymentMethodDataAffirmParams struct{}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type SetupIntentCreatePaymentMethodDataAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type SetupIntentCreatePaymentMethodDataAlipayParams struct{}

// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
type SetupIntentCreatePaymentMethodDataAlmaParams struct{}

// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
type SetupIntentCreatePaymentMethodDataAmazonPayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type SetupIntentCreatePaymentMethodDataAUBECSDebitParams struct {
	// The account number for the bank account.
	AccountNumber *string `form:"account_number"`
	// Bank-State-Branch number of the bank account.
	BSBNumber *string `form:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type SetupIntentCreatePaymentMethodDataBACSDebitParams struct {
	// Account number of the bank account that the funds will be debited from.
	AccountNumber *string `form:"account_number"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode *string `form:"sort_code"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type SetupIntentCreatePaymentMethodDataBancontactParams struct{}

// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
type SetupIntentCreatePaymentMethodDataBillieParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type SetupIntentCreatePaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address"`
	// Email address.
	Email *string `form:"email"`
	// Full name.
	Name *string `form:"name"`
	// Billing phone number (including extension).
	Phone *string `form:"phone"`
	// Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.
	TaxID *string `form:"tax_id"`
}

// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
type SetupIntentCreatePaymentMethodDataBLIKParams struct{}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type SetupIntentCreatePaymentMethodDataBoletoParams struct {
	// The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)
	TaxID *string `form:"tax_id"`
}

// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
type SetupIntentCreatePaymentMethodDataCashAppParams struct{}

// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
type SetupIntentCreatePaymentMethodDataCryptoParams struct{}

// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
type SetupIntentCreatePaymentMethodDataCustomerBalanceParams struct{}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type SetupIntentCreatePaymentMethodDataEPSParams struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type SetupIntentCreatePaymentMethodDataFPXParams struct {
	// Account holder type for FPX transaction
	AccountHolderType *string `form:"account_holder_type"`
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type SetupIntentCreatePaymentMethodDataGiropayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type SetupIntentCreatePaymentMethodDataGrabpayParams struct{}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type SetupIntentCreatePaymentMethodDataIDEALParams struct {
	// The customer's bank. Only use this parameter for existing customers. Don't use it for new customers.
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type SetupIntentCreatePaymentMethodDataInteracPresentParams struct{}

// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
type SetupIntentCreatePaymentMethodDataKakaoPayParams struct{}

// Customer's date of birth
type SetupIntentCreatePaymentMethodDataKlarnaDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
type SetupIntentCreatePaymentMethodDataKlarnaParams struct {
	// Customer's date of birth
	DOB *SetupIntentCreatePaymentMethodDataKlarnaDOBParams `form:"dob"`
}

// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
type SetupIntentCreatePaymentMethodDataKonbiniParams struct{}

// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
type SetupIntentCreatePaymentMethodDataKrCardParams struct{}

// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
type SetupIntentCreatePaymentMethodDataLinkParams struct{}

// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
type SetupIntentCreatePaymentMethodDataMobilepayParams struct{}

// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
type SetupIntentCreatePaymentMethodDataMultibancoParams struct{}

// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
type SetupIntentCreatePaymentMethodDataNaverPayParams struct {
	// Whether to use Naver Pay points or a card to fund this transaction. If not provided, this defaults to `card`.
	Funding *string `form:"funding"`
}

// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
type SetupIntentCreatePaymentMethodDataNzBankAccountParams struct {
	// The name on the bank account. Only required if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod's billing details.
	AccountHolderName *string `form:"account_holder_name"`
	// The account number for the bank account.
	AccountNumber *string `form:"account_number"`
	// The numeric code for the bank account's bank.
	BankCode *string `form:"bank_code"`
	// The numeric code for the bank account's bank branch.
	BranchCode *string `form:"branch_code"`
	Reference  *string `form:"reference"`
	// The suffix of the bank account number.
	Suffix *string `form:"suffix"`
}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type SetupIntentCreatePaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type SetupIntentCreatePaymentMethodDataP24Params struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
type SetupIntentCreatePaymentMethodDataPayByBankParams struct{}

// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
type SetupIntentCreatePaymentMethodDataPaycoParams struct{}

// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
type SetupIntentCreatePaymentMethodDataPayNowParams struct{}

// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
type SetupIntentCreatePaymentMethodDataPaypalParams struct{}

// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
type SetupIntentCreatePaymentMethodDataPixParams struct{}

// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
type SetupIntentCreatePaymentMethodDataPromptPayParams struct{}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type SetupIntentCreatePaymentMethodDataRadarOptionsParams struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
type SetupIntentCreatePaymentMethodDataRevolutPayParams struct{}

// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
type SetupIntentCreatePaymentMethodDataSamsungPayParams struct{}

// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
type SetupIntentCreatePaymentMethodDataSatispayParams struct{}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type SetupIntentCreatePaymentMethodDataSEPADebitParams struct {
	// IBAN of the bank account.
	IBAN *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type SetupIntentCreatePaymentMethodDataSofortParams struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country *string `form:"country"`
}

// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
type SetupIntentCreatePaymentMethodDataSwishParams struct{}

// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
type SetupIntentCreatePaymentMethodDataTWINTParams struct{}

// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
type SetupIntentCreatePaymentMethodDataUSBankAccountParams struct {
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
type SetupIntentCreatePaymentMethodDataWeChatPayParams struct{}

// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
type SetupIntentCreatePaymentMethodDataZipParams struct{}

// When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method)
// value in the SetupIntent.
type SetupIntentCreatePaymentMethodDataParams struct {
	// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
	ACSSDebit *SetupIntentCreatePaymentMethodDataACSSDebitParams `form:"acss_debit"`
	// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
	Affirm *SetupIntentCreatePaymentMethodDataAffirmParams `form:"affirm"`
	// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
	AfterpayClearpay *SetupIntentCreatePaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay"`
	// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
	Alipay *SetupIntentCreatePaymentMethodDataAlipayParams `form:"alipay"`
	// This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow. The field defaults to `unspecified`.
	AllowRedisplay *string `form:"allow_redisplay"`
	// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
	Alma *SetupIntentCreatePaymentMethodDataAlmaParams `form:"alma"`
	// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
	AmazonPay *SetupIntentCreatePaymentMethodDataAmazonPayParams `form:"amazon_pay"`
	// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
	AUBECSDebit *SetupIntentCreatePaymentMethodDataAUBECSDebitParams `form:"au_becs_debit"`
	// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
	BACSDebit *SetupIntentCreatePaymentMethodDataBACSDebitParams `form:"bacs_debit"`
	// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
	Bancontact *SetupIntentCreatePaymentMethodDataBancontactParams `form:"bancontact"`
	// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
	Billie *SetupIntentCreatePaymentMethodDataBillieParams `form:"billie"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *SetupIntentCreatePaymentMethodDataBillingDetailsParams `form:"billing_details"`
	// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
	BLIK *SetupIntentCreatePaymentMethodDataBLIKParams `form:"blik"`
	// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
	Boleto *SetupIntentCreatePaymentMethodDataBoletoParams `form:"boleto"`
	// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
	CashApp *SetupIntentCreatePaymentMethodDataCashAppParams `form:"cashapp"`
	// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
	Crypto *SetupIntentCreatePaymentMethodDataCryptoParams `form:"crypto"`
	// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
	CustomerBalance *SetupIntentCreatePaymentMethodDataCustomerBalanceParams `form:"customer_balance"`
	// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
	EPS *SetupIntentCreatePaymentMethodDataEPSParams `form:"eps"`
	// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
	FPX *SetupIntentCreatePaymentMethodDataFPXParams `form:"fpx"`
	// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
	Giropay *SetupIntentCreatePaymentMethodDataGiropayParams `form:"giropay"`
	// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
	Grabpay *SetupIntentCreatePaymentMethodDataGrabpayParams `form:"grabpay"`
	// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
	IDEAL *SetupIntentCreatePaymentMethodDataIDEALParams `form:"ideal"`
	// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
	InteracPresent *SetupIntentCreatePaymentMethodDataInteracPresentParams `form:"interac_present"`
	// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
	KakaoPay *SetupIntentCreatePaymentMethodDataKakaoPayParams `form:"kakao_pay"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
	Klarna *SetupIntentCreatePaymentMethodDataKlarnaParams `form:"klarna"`
	// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
	Konbini *SetupIntentCreatePaymentMethodDataKonbiniParams `form:"konbini"`
	// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
	KrCard *SetupIntentCreatePaymentMethodDataKrCardParams `form:"kr_card"`
	// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
	Link *SetupIntentCreatePaymentMethodDataLinkParams `form:"link"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
	Mobilepay *SetupIntentCreatePaymentMethodDataMobilepayParams `form:"mobilepay"`
	// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
	Multibanco *SetupIntentCreatePaymentMethodDataMultibancoParams `form:"multibanco"`
	// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
	NaverPay *SetupIntentCreatePaymentMethodDataNaverPayParams `form:"naver_pay"`
	// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
	NzBankAccount *SetupIntentCreatePaymentMethodDataNzBankAccountParams `form:"nz_bank_account"`
	// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
	OXXO *SetupIntentCreatePaymentMethodDataOXXOParams `form:"oxxo"`
	// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
	P24 *SetupIntentCreatePaymentMethodDataP24Params `form:"p24"`
	// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
	PayByBank *SetupIntentCreatePaymentMethodDataPayByBankParams `form:"pay_by_bank"`
	// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
	Payco *SetupIntentCreatePaymentMethodDataPaycoParams `form:"payco"`
	// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
	PayNow *SetupIntentCreatePaymentMethodDataPayNowParams `form:"paynow"`
	// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
	Paypal *SetupIntentCreatePaymentMethodDataPaypalParams `form:"paypal"`
	// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
	Pix *SetupIntentCreatePaymentMethodDataPixParams `form:"pix"`
	// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
	PromptPay *SetupIntentCreatePaymentMethodDataPromptPayParams `form:"promptpay"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *SetupIntentCreatePaymentMethodDataRadarOptionsParams `form:"radar_options"`
	// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
	RevolutPay *SetupIntentCreatePaymentMethodDataRevolutPayParams `form:"revolut_pay"`
	// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
	SamsungPay *SetupIntentCreatePaymentMethodDataSamsungPayParams `form:"samsung_pay"`
	// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
	Satispay *SetupIntentCreatePaymentMethodDataSatispayParams `form:"satispay"`
	// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
	SEPADebit *SetupIntentCreatePaymentMethodDataSEPADebitParams `form:"sepa_debit"`
	// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
	Sofort *SetupIntentCreatePaymentMethodDataSofortParams `form:"sofort"`
	// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
	Swish *SetupIntentCreatePaymentMethodDataSwishParams `form:"swish"`
	// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
	TWINT *SetupIntentCreatePaymentMethodDataTWINTParams `form:"twint"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type"`
	// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
	USBankAccount *SetupIntentCreatePaymentMethodDataUSBankAccountParams `form:"us_bank_account"`
	// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
	WeChatPay *SetupIntentCreatePaymentMethodDataWeChatPayParams `form:"wechat_pay"`
	// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
	Zip *SetupIntentCreatePaymentMethodDataZipParams `form:"zip"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SetupIntentCreatePaymentMethodDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Additional fields for Mandate creation
type SetupIntentCreatePaymentMethodOptionsACSSDebitMandateOptionsParams struct {
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
type SetupIntentCreatePaymentMethodOptionsACSSDebitParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentCreatePaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	// Bank account verification method.
	VerificationMethod *string `form:"verification_method"`
}

// If this is a `amazon_pay` SetupIntent, this sub-hash contains details about the AmazonPay payment method options.
type SetupIntentCreatePaymentMethodOptionsAmazonPayParams struct{}

// Additional fields for Mandate creation
type SetupIntentCreatePaymentMethodOptionsBACSDebitMandateOptionsParams struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'DDIC' or 'STRIPE'.
	ReferencePrefix *string `form:"reference_prefix"`
}

// If this is a `bacs_debit` SetupIntent, this sub-hash contains details about the Bacs Debit payment method options.
type SetupIntentCreatePaymentMethodOptionsBACSDebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentCreatePaymentMethodOptionsBACSDebitMandateOptionsParams `form:"mandate_options"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type SetupIntentCreatePaymentMethodOptionsCardMandateOptionsParams struct {
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

// Cartes Bancaires-specific 3DS fields.
type SetupIntentCreatePaymentMethodOptionsCardThreeDSecureNetworkOptionsCartesBancairesParams struct {
	// The cryptogram calculation algorithm used by the card Issuer's ACS
	// to calculate the Authentication cryptogram. Also known as `cavvAlgorithm`.
	// messageExtension: CB-AVALGO
	CbAvalgo *string `form:"cb_avalgo"`
	// The exemption indicator returned from Cartes Bancaires in the ARes.
	// message extension: CB-EXEMPTION; string (4 characters)
	// This is a 3 byte bitmap (low significant byte first and most significant
	// bit first) that has been Base64 encoded
	CbExemption *string `form:"cb_exemption"`
	// The risk score returned from Cartes Bancaires in the ARes.
	// message extension: CB-SCORE; numeric value 0-99
	CbScore *int64 `form:"cb_score"`
}

// Network specific 3DS fields. Network specific arguments require an
// explicit card brand choice. The parameter `payment_method_options.card.network“
// must be populated accordingly
type SetupIntentCreatePaymentMethodOptionsCardThreeDSecureNetworkOptionsParams struct {
	// Cartes Bancaires-specific 3DS fields.
	CartesBancaires *SetupIntentCreatePaymentMethodOptionsCardThreeDSecureNetworkOptionsCartesBancairesParams `form:"cartes_bancaires"`
}

// If 3D Secure authentication was performed with a third-party provider,
// the authentication details to use for this setup.
type SetupIntentCreatePaymentMethodOptionsCardThreeDSecureParams struct {
	// The `transStatus` returned from the card Issuer's ACS in the ARes.
	AresTransStatus *string `form:"ares_trans_status"`
	// The cryptogram, also known as the "authentication value" (AAV, CAVV or
	// AEVV). This value is 20 bytes, base64-encoded into a 28-character string.
	// (Most 3D Secure providers will return the base64-encoded version, which
	// is what you should specify here.)
	Cryptogram *string `form:"cryptogram"`
	// The Electronic Commerce Indicator (ECI) is returned by your 3D Secure
	// provider and indicates what degree of authentication was performed.
	ElectronicCommerceIndicator *string `form:"electronic_commerce_indicator"`
	// Network specific 3DS fields. Network specific arguments require an
	// explicit card brand choice. The parameter `payment_method_options.card.network``
	// must be populated accordingly
	NetworkOptions *SetupIntentCreatePaymentMethodOptionsCardThreeDSecureNetworkOptionsParams `form:"network_options"`
	// The challenge indicator (`threeDSRequestorChallengeInd`) which was requested in the
	// AReq sent to the card Issuer's ACS. A string containing 2 digits from 01-99.
	RequestorChallengeIndicator *string `form:"requestor_challenge_indicator"`
	// For 3D Secure 1, the XID. For 3D Secure 2, the Directory Server
	// Transaction ID (dsTransID).
	TransactionID *string `form:"transaction_id"`
	// The version of 3D Secure that was performed.
	Version *string `form:"version"`
}

// Configuration for any card setup attempted on this SetupIntent.
type SetupIntentCreatePaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *SetupIntentCreatePaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options"`
	// When specified, this parameter signals that a card has been collected
	// as MOTO (Mail Order Telephone Order) and thus out of scope for SCA. This
	// parameter can only be provided during confirmation.
	MOTO *bool `form:"moto"`
	// Selected network to process this SetupIntent on. Depends on the available networks of the card attached to the SetupIntent. Can be only set confirm-time.
	Network *string `form:"network"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure"`
	// If 3D Secure authentication was performed with a third-party provider,
	// the authentication details to use for this setup.
	ThreeDSecure *SetupIntentCreatePaymentMethodOptionsCardThreeDSecureParams `form:"three_d_secure"`
}

// If this is a `card_present` PaymentMethod, this sub-hash contains details about the card-present payment method options.
type SetupIntentCreatePaymentMethodOptionsCardPresentParams struct{}

// On-demand details if setting up a payment method for on-demand payments.
type SetupIntentCreatePaymentMethodOptionsKlarnaOnDemandParams struct {
	// Your average amount value. You can use a value across your customer base, or segment based on customer type, country, etc.
	AverageAmount *int64 `form:"average_amount"`
	// The maximum value you may charge a customer per purchase. You can use a value across your customer base, or segment based on customer type, country, etc.
	MaximumAmount *int64 `form:"maximum_amount"`
	// The lowest or minimum value you may charge a customer per purchase. You can use a value across your customer base, or segment based on customer type, country, etc.
	MinimumAmount *int64 `form:"minimum_amount"`
	// Interval at which the customer is making purchases
	PurchaseInterval *string `form:"purchase_interval"`
	// The number of `purchase_interval` between charges
	PurchaseIntervalCount *int64 `form:"purchase_interval_count"`
}

// Describes the upcoming charge for this subscription.
type SetupIntentCreatePaymentMethodOptionsKlarnaSubscriptionNextBillingParams struct {
	// The amount of the next charge for the subscription.
	Amount *int64 `form:"amount"`
	// The date of the next charge for the subscription in YYYY-MM-DD format.
	Date *string `form:"date"`
}

// Subscription details if setting up or charging a subscription
type SetupIntentCreatePaymentMethodOptionsKlarnaSubscriptionParams struct {
	// Unit of time between subscription charges.
	Interval *string `form:"interval"`
	// The number of intervals (specified in the `interval` attribute) between subscription charges. For example, `interval=month` and `interval_count=3` charges every 3 months.
	IntervalCount *int64 `form:"interval_count"`
	// Name for subscription.
	Name *string `form:"name"`
	// Describes the upcoming charge for this subscription.
	NextBilling *SetupIntentCreatePaymentMethodOptionsKlarnaSubscriptionNextBillingParams `form:"next_billing"`
	// A non-customer-facing reference to correlate subscription charges in the Klarna app. Use a value that persists across subscription charges.
	Reference *string `form:"reference"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method options.
type SetupIntentCreatePaymentMethodOptionsKlarnaParams struct {
	// The currency of the SetupIntent. Three letter ISO currency code.
	Currency *string `form:"currency"`
	// On-demand details if setting up a payment method for on-demand payments.
	OnDemand *SetupIntentCreatePaymentMethodOptionsKlarnaOnDemandParams `form:"on_demand"`
	// Preferred language of the Klarna authorization page that the customer is redirected to
	PreferredLocale *string `form:"preferred_locale"`
	// Subscription details if setting up or charging a subscription
	Subscriptions []*SetupIntentCreatePaymentMethodOptionsKlarnaSubscriptionParams `form:"subscriptions"`
}

// If this is a `link` PaymentMethod, this sub-hash contains details about the Link payment method options.
type SetupIntentCreatePaymentMethodOptionsLinkParams struct {
	// [Deprecated] This is a legacy parameter that no longer has any function.
	// Deprecated:
	PersistentToken *string `form:"persistent_token"`
}

// If this is a `paypal` PaymentMethod, this sub-hash contains details about the PayPal payment method options.
type SetupIntentCreatePaymentMethodOptionsPaypalParams struct {
	// The PayPal Billing Agreement ID (BAID). This is an ID generated by PayPal which represents the mandate between the merchant and the customer.
	BillingAgreementID *string `form:"billing_agreement_id"`
}

// Additional fields for Mandate creation
type SetupIntentCreatePaymentMethodOptionsSEPADebitMandateOptionsParams struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'STRIPE'.
	ReferencePrefix *string `form:"reference_prefix"`
}

// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
type SetupIntentCreatePaymentMethodOptionsSEPADebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentCreatePaymentMethodOptionsSEPADebitMandateOptionsParams `form:"mandate_options"`
}

// Provide filters for the linked accounts that the customer can select for the payment method.
type SetupIntentCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts. Valid subcategories are `checking` and `savings`.
	AccountSubcategories []*string `form:"account_subcategories"`
}

// Additional fields for Financial Connections Session creation
type SetupIntentCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *SetupIntentCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included. Valid permissions include: `balances`, `ownership`, `payment_method`, and `transactions`.
	Permissions []*string `form:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL *string `form:"return_url"`
}

// Additional fields for Mandate creation
type SetupIntentCreatePaymentMethodOptionsUSBankAccountMandateOptionsParams struct {
	// The method used to collect offline mandate customer acceptance.
	CollectionMethod *string `form:"collection_method"`
}

// Additional fields for network related functions
type SetupIntentCreatePaymentMethodOptionsUSBankAccountNetworksParams struct {
	// Triggers validations to run across the selected networks
	Requested []*string `form:"requested"`
}

// If this is a `us_bank_account` SetupIntent, this sub-hash contains details about the US bank account payment method options.
type SetupIntentCreatePaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation
	FinancialConnections *SetupIntentCreatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections"`
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentCreatePaymentMethodOptionsUSBankAccountMandateOptionsParams `form:"mandate_options"`
	// Additional fields for network related functions
	Networks *SetupIntentCreatePaymentMethodOptionsUSBankAccountNetworksParams `form:"networks"`
	// Bank account verification method.
	VerificationMethod *string `form:"verification_method"`
}

// Payment method-specific configuration for this SetupIntent.
type SetupIntentCreatePaymentMethodOptionsParams struct {
	// If this is a `acss_debit` SetupIntent, this sub-hash contains details about the ACSS Debit payment method options.
	ACSSDebit *SetupIntentCreatePaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	// If this is a `amazon_pay` SetupIntent, this sub-hash contains details about the AmazonPay payment method options.
	AmazonPay *SetupIntentCreatePaymentMethodOptionsAmazonPayParams `form:"amazon_pay"`
	// If this is a `bacs_debit` SetupIntent, this sub-hash contains details about the Bacs Debit payment method options.
	BACSDebit *SetupIntentCreatePaymentMethodOptionsBACSDebitParams `form:"bacs_debit"`
	// Configuration for any card setup attempted on this SetupIntent.
	Card *SetupIntentCreatePaymentMethodOptionsCardParams `form:"card"`
	// If this is a `card_present` PaymentMethod, this sub-hash contains details about the card-present payment method options.
	CardPresent *SetupIntentCreatePaymentMethodOptionsCardPresentParams `form:"card_present"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method options.
	Klarna *SetupIntentCreatePaymentMethodOptionsKlarnaParams `form:"klarna"`
	// If this is a `link` PaymentMethod, this sub-hash contains details about the Link payment method options.
	Link *SetupIntentCreatePaymentMethodOptionsLinkParams `form:"link"`
	// If this is a `paypal` PaymentMethod, this sub-hash contains details about the PayPal payment method options.
	Paypal *SetupIntentCreatePaymentMethodOptionsPaypalParams `form:"paypal"`
	// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
	SEPADebit *SetupIntentCreatePaymentMethodOptionsSEPADebitParams `form:"sepa_debit"`
	// If this is a `us_bank_account` SetupIntent, this sub-hash contains details about the US bank account payment method options.
	USBankAccount *SetupIntentCreatePaymentMethodOptionsUSBankAccountParams `form:"us_bank_account"`
}

// If you populate this hash, this SetupIntent generates a `single_use` mandate after successful completion.
//
// Single-use mandates are only valid for the following payment methods: `acss_debit`, `alipay`, `au_becs_debit`, `bacs_debit`, `bancontact`, `boleto`, `ideal`, `link`, `sepa_debit`, and `us_bank_account`.
type SetupIntentCreateSingleUseParams struct {
	// Amount the customer is granting permission to collect later. A positive integer representing how much to charge in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://stripe.com/docs/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
}

// Creates a SetupIntent object.
//
// After you create the SetupIntent, attach a payment method and [confirm](https://docs.stripe.com/docs/api/setup_intents/confirm)
// it to collect any required permissions to charge the payment method later.
type SetupIntentCreateParams struct {
	Params `form:"*"`
	// If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.
	//
	// It can only be used for this Stripe Account's own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.
	AttachToSelf *bool `form:"attach_to_self"`
	// When you enable this parameter, this SetupIntent accepts payment methods that you enable in the Dashboard and that are compatible with its other parameters.
	AutomaticPaymentMethods *SetupIntentCreateAutomaticPaymentMethodsParams `form:"automatic_payment_methods"`
	// Set to `true` to attempt to confirm this SetupIntent immediately. This parameter defaults to `false`. If a card is the attached payment method, you can provide a `return_url` in case further authentication is necessary.
	Confirm *bool `form:"confirm"`
	// ID of the ConfirmationToken used to confirm this SetupIntent.
	//
	// If the provided ConfirmationToken contains properties that are also being provided in this request, such as `payment_method`, then the values in this request will take precedence.
	ConfirmationToken *string `form:"confirmation_token"`
	// ID of the Customer this SetupIntent belongs to, if one exists.
	//
	// If present, the SetupIntent's payment method will be attached to the Customer on successful setup. Payment methods attached to other Customers cannot be used with this SetupIntent.
	Customer *string `form:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Indicates the directions of money movement for which this payment method is intended to be used.
	//
	// Include `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.
	FlowDirections []*string `form:"flow_directions"`
	// This hash contains details about the mandate to create. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
	MandateData *SetupIntentCreateMandateDataParams `form:"mandate_data"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The Stripe account ID created for this SetupIntent.
	OnBehalfOf *string `form:"on_behalf_of"`
	// ID of the payment method (a PaymentMethod, Card, or saved Source object) to attach to this SetupIntent.
	PaymentMethod *string `form:"payment_method"`
	// The ID of the [payment method configuration](https://stripe.com/docs/api/payment_method_configurations) to use with this SetupIntent.
	PaymentMethodConfiguration *string `form:"payment_method_configuration"`
	// When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method)
	// value in the SetupIntent.
	PaymentMethodData *SetupIntentCreatePaymentMethodDataParams `form:"payment_method_data"`
	// Payment method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentCreatePaymentMethodOptionsParams `form:"payment_method_options"`
	// The list of payment method types (for example, card) that this SetupIntent can use. If you don't provide this, Stripe will dynamically show relevant payment methods from your [payment method settings](https://dashboard.stripe.com/settings/payment_methods). A list of valid payment method types can be found [here](https://docs.stripe.com/api/payment_methods/object#payment_method_object-type).
	PaymentMethodTypes []*string `form:"payment_method_types"`
	// The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method's app or site. To redirect to a mobile application, you can alternatively supply an application URI scheme. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
	ReturnURL *string `form:"return_url"`
	// If you populate this hash, this SetupIntent generates a `single_use` mandate after successful completion.
	//
	// Single-use mandates are only valid for the following payment methods: `acss_debit`, `alipay`, `au_becs_debit`, `bacs_debit`, `bancontact`, `boleto`, `ideal`, `link`, `sepa_debit`, and `us_bank_account`.
	SingleUse *SetupIntentCreateSingleUseParams `form:"single_use"`
	// Indicates how the payment method is intended to be used in the future. If not provided, this value defaults to `off_session`.
	Usage *string `form:"usage"`
	// Set to `true` when confirming server-side and using Stripe.js, iOS, or Android client-side SDKs to handle the next actions.
	UseStripeSDK *bool `form:"use_stripe_sdk"`
}

// AddExpand appends a new field to expand.
func (p *SetupIntentCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SetupIntentCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of a SetupIntent that has previously been created.
//
// Client-side retrieval using a publishable key is allowed when the client_secret is provided in the query string.
//
// When retrieved with a publishable key, only a subset of properties will be returned. Please refer to the [SetupIntent](https://docs.stripe.com/api#setup_intent_object) object reference for more details.
type SetupIntentRetrieveParams struct {
	Params `form:"*"`
	// The client secret of the SetupIntent. We require this string if you use a publishable key to retrieve the SetupIntent.
	ClientSecret *string `form:"client_secret"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *SetupIntentRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type SetupIntentUpdatePaymentMethodDataACSSDebitParams struct {
	// Customer's bank account number.
	AccountNumber *string `form:"account_number"`
	// Institution number of the customer's bank.
	InstitutionNumber *string `form:"institution_number"`
	// Transit number of the customer's bank.
	TransitNumber *string `form:"transit_number"`
}

// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
type SetupIntentUpdatePaymentMethodDataAffirmParams struct{}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type SetupIntentUpdatePaymentMethodDataAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type SetupIntentUpdatePaymentMethodDataAlipayParams struct{}

// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
type SetupIntentUpdatePaymentMethodDataAlmaParams struct{}

// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
type SetupIntentUpdatePaymentMethodDataAmazonPayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type SetupIntentUpdatePaymentMethodDataAUBECSDebitParams struct {
	// The account number for the bank account.
	AccountNumber *string `form:"account_number"`
	// Bank-State-Branch number of the bank account.
	BSBNumber *string `form:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type SetupIntentUpdatePaymentMethodDataBACSDebitParams struct {
	// Account number of the bank account that the funds will be debited from.
	AccountNumber *string `form:"account_number"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode *string `form:"sort_code"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type SetupIntentUpdatePaymentMethodDataBancontactParams struct{}

// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
type SetupIntentUpdatePaymentMethodDataBillieParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type SetupIntentUpdatePaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address"`
	// Email address.
	Email *string `form:"email"`
	// Full name.
	Name *string `form:"name"`
	// Billing phone number (including extension).
	Phone *string `form:"phone"`
	// Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.
	TaxID *string `form:"tax_id"`
}

// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
type SetupIntentUpdatePaymentMethodDataBLIKParams struct{}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type SetupIntentUpdatePaymentMethodDataBoletoParams struct {
	// The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)
	TaxID *string `form:"tax_id"`
}

// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
type SetupIntentUpdatePaymentMethodDataCashAppParams struct{}

// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
type SetupIntentUpdatePaymentMethodDataCryptoParams struct{}

// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
type SetupIntentUpdatePaymentMethodDataCustomerBalanceParams struct{}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type SetupIntentUpdatePaymentMethodDataEPSParams struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type SetupIntentUpdatePaymentMethodDataFPXParams struct {
	// Account holder type for FPX transaction
	AccountHolderType *string `form:"account_holder_type"`
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type SetupIntentUpdatePaymentMethodDataGiropayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type SetupIntentUpdatePaymentMethodDataGrabpayParams struct{}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type SetupIntentUpdatePaymentMethodDataIDEALParams struct {
	// The customer's bank. Only use this parameter for existing customers. Don't use it for new customers.
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type SetupIntentUpdatePaymentMethodDataInteracPresentParams struct{}

// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
type SetupIntentUpdatePaymentMethodDataKakaoPayParams struct{}

// Customer's date of birth
type SetupIntentUpdatePaymentMethodDataKlarnaDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
type SetupIntentUpdatePaymentMethodDataKlarnaParams struct {
	// Customer's date of birth
	DOB *SetupIntentUpdatePaymentMethodDataKlarnaDOBParams `form:"dob"`
}

// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
type SetupIntentUpdatePaymentMethodDataKonbiniParams struct{}

// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
type SetupIntentUpdatePaymentMethodDataKrCardParams struct{}

// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
type SetupIntentUpdatePaymentMethodDataLinkParams struct{}

// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
type SetupIntentUpdatePaymentMethodDataMobilepayParams struct{}

// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
type SetupIntentUpdatePaymentMethodDataMultibancoParams struct{}

// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
type SetupIntentUpdatePaymentMethodDataNaverPayParams struct {
	// Whether to use Naver Pay points or a card to fund this transaction. If not provided, this defaults to `card`.
	Funding *string `form:"funding"`
}

// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
type SetupIntentUpdatePaymentMethodDataNzBankAccountParams struct {
	// The name on the bank account. Only required if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod's billing details.
	AccountHolderName *string `form:"account_holder_name"`
	// The account number for the bank account.
	AccountNumber *string `form:"account_number"`
	// The numeric code for the bank account's bank.
	BankCode *string `form:"bank_code"`
	// The numeric code for the bank account's bank branch.
	BranchCode *string `form:"branch_code"`
	Reference  *string `form:"reference"`
	// The suffix of the bank account number.
	Suffix *string `form:"suffix"`
}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type SetupIntentUpdatePaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type SetupIntentUpdatePaymentMethodDataP24Params struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
type SetupIntentUpdatePaymentMethodDataPayByBankParams struct{}

// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
type SetupIntentUpdatePaymentMethodDataPaycoParams struct{}

// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
type SetupIntentUpdatePaymentMethodDataPayNowParams struct{}

// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
type SetupIntentUpdatePaymentMethodDataPaypalParams struct{}

// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
type SetupIntentUpdatePaymentMethodDataPixParams struct{}

// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
type SetupIntentUpdatePaymentMethodDataPromptPayParams struct{}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type SetupIntentUpdatePaymentMethodDataRadarOptionsParams struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
type SetupIntentUpdatePaymentMethodDataRevolutPayParams struct{}

// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
type SetupIntentUpdatePaymentMethodDataSamsungPayParams struct{}

// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
type SetupIntentUpdatePaymentMethodDataSatispayParams struct{}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type SetupIntentUpdatePaymentMethodDataSEPADebitParams struct {
	// IBAN of the bank account.
	IBAN *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type SetupIntentUpdatePaymentMethodDataSofortParams struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country *string `form:"country"`
}

// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
type SetupIntentUpdatePaymentMethodDataSwishParams struct{}

// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
type SetupIntentUpdatePaymentMethodDataTWINTParams struct{}

// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
type SetupIntentUpdatePaymentMethodDataUSBankAccountParams struct {
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
type SetupIntentUpdatePaymentMethodDataWeChatPayParams struct{}

// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
type SetupIntentUpdatePaymentMethodDataZipParams struct{}

// When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method)
// value in the SetupIntent.
type SetupIntentUpdatePaymentMethodDataParams struct {
	// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
	ACSSDebit *SetupIntentUpdatePaymentMethodDataACSSDebitParams `form:"acss_debit"`
	// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
	Affirm *SetupIntentUpdatePaymentMethodDataAffirmParams `form:"affirm"`
	// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
	AfterpayClearpay *SetupIntentUpdatePaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay"`
	// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
	Alipay *SetupIntentUpdatePaymentMethodDataAlipayParams `form:"alipay"`
	// This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow. The field defaults to `unspecified`.
	AllowRedisplay *string `form:"allow_redisplay"`
	// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
	Alma *SetupIntentUpdatePaymentMethodDataAlmaParams `form:"alma"`
	// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
	AmazonPay *SetupIntentUpdatePaymentMethodDataAmazonPayParams `form:"amazon_pay"`
	// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
	AUBECSDebit *SetupIntentUpdatePaymentMethodDataAUBECSDebitParams `form:"au_becs_debit"`
	// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
	BACSDebit *SetupIntentUpdatePaymentMethodDataBACSDebitParams `form:"bacs_debit"`
	// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
	Bancontact *SetupIntentUpdatePaymentMethodDataBancontactParams `form:"bancontact"`
	// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
	Billie *SetupIntentUpdatePaymentMethodDataBillieParams `form:"billie"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *SetupIntentUpdatePaymentMethodDataBillingDetailsParams `form:"billing_details"`
	// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
	BLIK *SetupIntentUpdatePaymentMethodDataBLIKParams `form:"blik"`
	// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
	Boleto *SetupIntentUpdatePaymentMethodDataBoletoParams `form:"boleto"`
	// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
	CashApp *SetupIntentUpdatePaymentMethodDataCashAppParams `form:"cashapp"`
	// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
	Crypto *SetupIntentUpdatePaymentMethodDataCryptoParams `form:"crypto"`
	// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
	CustomerBalance *SetupIntentUpdatePaymentMethodDataCustomerBalanceParams `form:"customer_balance"`
	// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
	EPS *SetupIntentUpdatePaymentMethodDataEPSParams `form:"eps"`
	// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
	FPX *SetupIntentUpdatePaymentMethodDataFPXParams `form:"fpx"`
	// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
	Giropay *SetupIntentUpdatePaymentMethodDataGiropayParams `form:"giropay"`
	// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
	Grabpay *SetupIntentUpdatePaymentMethodDataGrabpayParams `form:"grabpay"`
	// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
	IDEAL *SetupIntentUpdatePaymentMethodDataIDEALParams `form:"ideal"`
	// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
	InteracPresent *SetupIntentUpdatePaymentMethodDataInteracPresentParams `form:"interac_present"`
	// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
	KakaoPay *SetupIntentUpdatePaymentMethodDataKakaoPayParams `form:"kakao_pay"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
	Klarna *SetupIntentUpdatePaymentMethodDataKlarnaParams `form:"klarna"`
	// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
	Konbini *SetupIntentUpdatePaymentMethodDataKonbiniParams `form:"konbini"`
	// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
	KrCard *SetupIntentUpdatePaymentMethodDataKrCardParams `form:"kr_card"`
	// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
	Link *SetupIntentUpdatePaymentMethodDataLinkParams `form:"link"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
	Mobilepay *SetupIntentUpdatePaymentMethodDataMobilepayParams `form:"mobilepay"`
	// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
	Multibanco *SetupIntentUpdatePaymentMethodDataMultibancoParams `form:"multibanco"`
	// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
	NaverPay *SetupIntentUpdatePaymentMethodDataNaverPayParams `form:"naver_pay"`
	// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
	NzBankAccount *SetupIntentUpdatePaymentMethodDataNzBankAccountParams `form:"nz_bank_account"`
	// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
	OXXO *SetupIntentUpdatePaymentMethodDataOXXOParams `form:"oxxo"`
	// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
	P24 *SetupIntentUpdatePaymentMethodDataP24Params `form:"p24"`
	// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
	PayByBank *SetupIntentUpdatePaymentMethodDataPayByBankParams `form:"pay_by_bank"`
	// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
	Payco *SetupIntentUpdatePaymentMethodDataPaycoParams `form:"payco"`
	// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
	PayNow *SetupIntentUpdatePaymentMethodDataPayNowParams `form:"paynow"`
	// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
	Paypal *SetupIntentUpdatePaymentMethodDataPaypalParams `form:"paypal"`
	// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
	Pix *SetupIntentUpdatePaymentMethodDataPixParams `form:"pix"`
	// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
	PromptPay *SetupIntentUpdatePaymentMethodDataPromptPayParams `form:"promptpay"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *SetupIntentUpdatePaymentMethodDataRadarOptionsParams `form:"radar_options"`
	// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
	RevolutPay *SetupIntentUpdatePaymentMethodDataRevolutPayParams `form:"revolut_pay"`
	// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
	SamsungPay *SetupIntentUpdatePaymentMethodDataSamsungPayParams `form:"samsung_pay"`
	// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
	Satispay *SetupIntentUpdatePaymentMethodDataSatispayParams `form:"satispay"`
	// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
	SEPADebit *SetupIntentUpdatePaymentMethodDataSEPADebitParams `form:"sepa_debit"`
	// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
	Sofort *SetupIntentUpdatePaymentMethodDataSofortParams `form:"sofort"`
	// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
	Swish *SetupIntentUpdatePaymentMethodDataSwishParams `form:"swish"`
	// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
	TWINT *SetupIntentUpdatePaymentMethodDataTWINTParams `form:"twint"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type"`
	// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
	USBankAccount *SetupIntentUpdatePaymentMethodDataUSBankAccountParams `form:"us_bank_account"`
	// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
	WeChatPay *SetupIntentUpdatePaymentMethodDataWeChatPayParams `form:"wechat_pay"`
	// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
	Zip *SetupIntentUpdatePaymentMethodDataZipParams `form:"zip"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SetupIntentUpdatePaymentMethodDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Additional fields for Mandate creation
type SetupIntentUpdatePaymentMethodOptionsACSSDebitMandateOptionsParams struct {
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
type SetupIntentUpdatePaymentMethodOptionsACSSDebitParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentUpdatePaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	// Bank account verification method.
	VerificationMethod *string `form:"verification_method"`
}

// If this is a `amazon_pay` SetupIntent, this sub-hash contains details about the AmazonPay payment method options.
type SetupIntentUpdatePaymentMethodOptionsAmazonPayParams struct{}

// Additional fields for Mandate creation
type SetupIntentUpdatePaymentMethodOptionsBACSDebitMandateOptionsParams struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'DDIC' or 'STRIPE'.
	ReferencePrefix *string `form:"reference_prefix"`
}

// If this is a `bacs_debit` SetupIntent, this sub-hash contains details about the Bacs Debit payment method options.
type SetupIntentUpdatePaymentMethodOptionsBACSDebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentUpdatePaymentMethodOptionsBACSDebitMandateOptionsParams `form:"mandate_options"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type SetupIntentUpdatePaymentMethodOptionsCardMandateOptionsParams struct {
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

// Cartes Bancaires-specific 3DS fields.
type SetupIntentUpdatePaymentMethodOptionsCardThreeDSecureNetworkOptionsCartesBancairesParams struct {
	// The cryptogram calculation algorithm used by the card Issuer's ACS
	// to calculate the Authentication cryptogram. Also known as `cavvAlgorithm`.
	// messageExtension: CB-AVALGO
	CbAvalgo *string `form:"cb_avalgo"`
	// The exemption indicator returned from Cartes Bancaires in the ARes.
	// message extension: CB-EXEMPTION; string (4 characters)
	// This is a 3 byte bitmap (low significant byte first and most significant
	// bit first) that has been Base64 encoded
	CbExemption *string `form:"cb_exemption"`
	// The risk score returned from Cartes Bancaires in the ARes.
	// message extension: CB-SCORE; numeric value 0-99
	CbScore *int64 `form:"cb_score"`
}

// Network specific 3DS fields. Network specific arguments require an
// explicit card brand choice. The parameter `payment_method_options.card.network“
// must be populated accordingly
type SetupIntentUpdatePaymentMethodOptionsCardThreeDSecureNetworkOptionsParams struct {
	// Cartes Bancaires-specific 3DS fields.
	CartesBancaires *SetupIntentUpdatePaymentMethodOptionsCardThreeDSecureNetworkOptionsCartesBancairesParams `form:"cartes_bancaires"`
}

// If 3D Secure authentication was performed with a third-party provider,
// the authentication details to use for this setup.
type SetupIntentUpdatePaymentMethodOptionsCardThreeDSecureParams struct {
	// The `transStatus` returned from the card Issuer's ACS in the ARes.
	AresTransStatus *string `form:"ares_trans_status"`
	// The cryptogram, also known as the "authentication value" (AAV, CAVV or
	// AEVV). This value is 20 bytes, base64-encoded into a 28-character string.
	// (Most 3D Secure providers will return the base64-encoded version, which
	// is what you should specify here.)
	Cryptogram *string `form:"cryptogram"`
	// The Electronic Commerce Indicator (ECI) is returned by your 3D Secure
	// provider and indicates what degree of authentication was performed.
	ElectronicCommerceIndicator *string `form:"electronic_commerce_indicator"`
	// Network specific 3DS fields. Network specific arguments require an
	// explicit card brand choice. The parameter `payment_method_options.card.network``
	// must be populated accordingly
	NetworkOptions *SetupIntentUpdatePaymentMethodOptionsCardThreeDSecureNetworkOptionsParams `form:"network_options"`
	// The challenge indicator (`threeDSRequestorChallengeInd`) which was requested in the
	// AReq sent to the card Issuer's ACS. A string containing 2 digits from 01-99.
	RequestorChallengeIndicator *string `form:"requestor_challenge_indicator"`
	// For 3D Secure 1, the XID. For 3D Secure 2, the Directory Server
	// Transaction ID (dsTransID).
	TransactionID *string `form:"transaction_id"`
	// The version of 3D Secure that was performed.
	Version *string `form:"version"`
}

// Configuration for any card setup attempted on this SetupIntent.
type SetupIntentUpdatePaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *SetupIntentUpdatePaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options"`
	// When specified, this parameter signals that a card has been collected
	// as MOTO (Mail Order Telephone Order) and thus out of scope for SCA. This
	// parameter can only be provided during confirmation.
	MOTO *bool `form:"moto"`
	// Selected network to process this SetupIntent on. Depends on the available networks of the card attached to the SetupIntent. Can be only set confirm-time.
	Network *string `form:"network"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure"`
	// If 3D Secure authentication was performed with a third-party provider,
	// the authentication details to use for this setup.
	ThreeDSecure *SetupIntentUpdatePaymentMethodOptionsCardThreeDSecureParams `form:"three_d_secure"`
}

// If this is a `card_present` PaymentMethod, this sub-hash contains details about the card-present payment method options.
type SetupIntentUpdatePaymentMethodOptionsCardPresentParams struct{}

// On-demand details if setting up a payment method for on-demand payments.
type SetupIntentUpdatePaymentMethodOptionsKlarnaOnDemandParams struct {
	// Your average amount value. You can use a value across your customer base, or segment based on customer type, country, etc.
	AverageAmount *int64 `form:"average_amount"`
	// The maximum value you may charge a customer per purchase. You can use a value across your customer base, or segment based on customer type, country, etc.
	MaximumAmount *int64 `form:"maximum_amount"`
	// The lowest or minimum value you may charge a customer per purchase. You can use a value across your customer base, or segment based on customer type, country, etc.
	MinimumAmount *int64 `form:"minimum_amount"`
	// Interval at which the customer is making purchases
	PurchaseInterval *string `form:"purchase_interval"`
	// The number of `purchase_interval` between charges
	PurchaseIntervalCount *int64 `form:"purchase_interval_count"`
}

// Describes the upcoming charge for this subscription.
type SetupIntentUpdatePaymentMethodOptionsKlarnaSubscriptionNextBillingParams struct {
	// The amount of the next charge for the subscription.
	Amount *int64 `form:"amount"`
	// The date of the next charge for the subscription in YYYY-MM-DD format.
	Date *string `form:"date"`
}

// Subscription details if setting up or charging a subscription
type SetupIntentUpdatePaymentMethodOptionsKlarnaSubscriptionParams struct {
	// Unit of time between subscription charges.
	Interval *string `form:"interval"`
	// The number of intervals (specified in the `interval` attribute) between subscription charges. For example, `interval=month` and `interval_count=3` charges every 3 months.
	IntervalCount *int64 `form:"interval_count"`
	// Name for subscription.
	Name *string `form:"name"`
	// Describes the upcoming charge for this subscription.
	NextBilling *SetupIntentUpdatePaymentMethodOptionsKlarnaSubscriptionNextBillingParams `form:"next_billing"`
	// A non-customer-facing reference to correlate subscription charges in the Klarna app. Use a value that persists across subscription charges.
	Reference *string `form:"reference"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method options.
type SetupIntentUpdatePaymentMethodOptionsKlarnaParams struct {
	// The currency of the SetupIntent. Three letter ISO currency code.
	Currency *string `form:"currency"`
	// On-demand details if setting up a payment method for on-demand payments.
	OnDemand *SetupIntentUpdatePaymentMethodOptionsKlarnaOnDemandParams `form:"on_demand"`
	// Preferred language of the Klarna authorization page that the customer is redirected to
	PreferredLocale *string `form:"preferred_locale"`
	// Subscription details if setting up or charging a subscription
	Subscriptions []*SetupIntentUpdatePaymentMethodOptionsKlarnaSubscriptionParams `form:"subscriptions"`
}

// If this is a `link` PaymentMethod, this sub-hash contains details about the Link payment method options.
type SetupIntentUpdatePaymentMethodOptionsLinkParams struct {
	// [Deprecated] This is a legacy parameter that no longer has any function.
	// Deprecated:
	PersistentToken *string `form:"persistent_token"`
}

// If this is a `paypal` PaymentMethod, this sub-hash contains details about the PayPal payment method options.
type SetupIntentUpdatePaymentMethodOptionsPaypalParams struct {
	// The PayPal Billing Agreement ID (BAID). This is an ID generated by PayPal which represents the mandate between the merchant and the customer.
	BillingAgreementID *string `form:"billing_agreement_id"`
}

// Additional fields for Mandate creation
type SetupIntentUpdatePaymentMethodOptionsSEPADebitMandateOptionsParams struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'STRIPE'.
	ReferencePrefix *string `form:"reference_prefix"`
}

// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
type SetupIntentUpdatePaymentMethodOptionsSEPADebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentUpdatePaymentMethodOptionsSEPADebitMandateOptionsParams `form:"mandate_options"`
}

// Provide filters for the linked accounts that the customer can select for the payment method.
type SetupIntentUpdatePaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts. Valid subcategories are `checking` and `savings`.
	AccountSubcategories []*string `form:"account_subcategories"`
}

// Additional fields for Financial Connections Session creation
type SetupIntentUpdatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *SetupIntentUpdatePaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included. Valid permissions include: `balances`, `ownership`, `payment_method`, and `transactions`.
	Permissions []*string `form:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL *string `form:"return_url"`
}

// Additional fields for Mandate creation
type SetupIntentUpdatePaymentMethodOptionsUSBankAccountMandateOptionsParams struct {
	// The method used to collect offline mandate customer acceptance.
	CollectionMethod *string `form:"collection_method"`
}

// Additional fields for network related functions
type SetupIntentUpdatePaymentMethodOptionsUSBankAccountNetworksParams struct {
	// Triggers validations to run across the selected networks
	Requested []*string `form:"requested"`
}

// If this is a `us_bank_account` SetupIntent, this sub-hash contains details about the US bank account payment method options.
type SetupIntentUpdatePaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation
	FinancialConnections *SetupIntentUpdatePaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections"`
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentUpdatePaymentMethodOptionsUSBankAccountMandateOptionsParams `form:"mandate_options"`
	// Additional fields for network related functions
	Networks *SetupIntentUpdatePaymentMethodOptionsUSBankAccountNetworksParams `form:"networks"`
	// Bank account verification method.
	VerificationMethod *string `form:"verification_method"`
}

// Payment method-specific configuration for this SetupIntent.
type SetupIntentUpdatePaymentMethodOptionsParams struct {
	// If this is a `acss_debit` SetupIntent, this sub-hash contains details about the ACSS Debit payment method options.
	ACSSDebit *SetupIntentUpdatePaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	// If this is a `amazon_pay` SetupIntent, this sub-hash contains details about the AmazonPay payment method options.
	AmazonPay *SetupIntentUpdatePaymentMethodOptionsAmazonPayParams `form:"amazon_pay"`
	// If this is a `bacs_debit` SetupIntent, this sub-hash contains details about the Bacs Debit payment method options.
	BACSDebit *SetupIntentUpdatePaymentMethodOptionsBACSDebitParams `form:"bacs_debit"`
	// Configuration for any card setup attempted on this SetupIntent.
	Card *SetupIntentUpdatePaymentMethodOptionsCardParams `form:"card"`
	// If this is a `card_present` PaymentMethod, this sub-hash contains details about the card-present payment method options.
	CardPresent *SetupIntentUpdatePaymentMethodOptionsCardPresentParams `form:"card_present"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method options.
	Klarna *SetupIntentUpdatePaymentMethodOptionsKlarnaParams `form:"klarna"`
	// If this is a `link` PaymentMethod, this sub-hash contains details about the Link payment method options.
	Link *SetupIntentUpdatePaymentMethodOptionsLinkParams `form:"link"`
	// If this is a `paypal` PaymentMethod, this sub-hash contains details about the PayPal payment method options.
	Paypal *SetupIntentUpdatePaymentMethodOptionsPaypalParams `form:"paypal"`
	// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
	SEPADebit *SetupIntentUpdatePaymentMethodOptionsSEPADebitParams `form:"sepa_debit"`
	// If this is a `us_bank_account` SetupIntent, this sub-hash contains details about the US bank account payment method options.
	USBankAccount *SetupIntentUpdatePaymentMethodOptionsUSBankAccountParams `form:"us_bank_account"`
}

// Updates a SetupIntent object.
type SetupIntentUpdateParams struct {
	Params `form:"*"`
	// If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.
	//
	// It can only be used for this Stripe Account's own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.
	AttachToSelf *bool `form:"attach_to_self"`
	// ID of the Customer this SetupIntent belongs to, if one exists.
	//
	// If present, the SetupIntent's payment method will be attached to the Customer on successful setup. Payment methods attached to other Customers cannot be used with this SetupIntent.
	Customer *string `form:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Indicates the directions of money movement for which this payment method is intended to be used.
	//
	// Include `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.
	FlowDirections []*string `form:"flow_directions"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// ID of the payment method (a PaymentMethod, Card, or saved Source object) to attach to this SetupIntent. To unset this field to null, pass in an empty string.
	PaymentMethod *string `form:"payment_method"`
	// The ID of the [payment method configuration](https://stripe.com/docs/api/payment_method_configurations) to use with this SetupIntent.
	PaymentMethodConfiguration *string `form:"payment_method_configuration"`
	// When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method)
	// value in the SetupIntent.
	PaymentMethodData *SetupIntentUpdatePaymentMethodDataParams `form:"payment_method_data"`
	// Payment method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentUpdatePaymentMethodOptionsParams `form:"payment_method_options"`
	// The list of payment method types (for example, card) that this SetupIntent can set up. If you don't provide this, Stripe will dynamically show relevant payment methods from your [payment method settings](https://dashboard.stripe.com/settings/payment_methods). A list of valid payment method types can be found [here](https://docs.stripe.com/api/payment_methods/object#payment_method_object-type).
	PaymentMethodTypes []*string `form:"payment_method_types"`
}

// AddExpand appends a new field to expand.
func (p *SetupIntentUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *SetupIntentUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Settings for dynamic payment methods compatible with this Setup Intent
type SetupIntentAutomaticPaymentMethods struct {
	// Controls whether this SetupIntent will accept redirect-based payment methods.
	//
	// Redirect-based payment methods may require your customer to be redirected to a payment method's app or site for authentication or additional steps. To [confirm](https://stripe.com/docs/api/setup_intents/confirm) this SetupIntent, you may be required to provide a `return_url` to redirect customers back to your site after they authenticate or complete the setup.
	AllowRedirects SetupIntentAutomaticPaymentMethodsAllowRedirects `json:"allow_redirects"`
	// Automatically calculates compatible payment methods
	Enabled bool `json:"enabled"`
}
type SetupIntentNextActionCashAppHandleRedirectOrDisplayQRCodeQRCode struct {
	// The date (unix timestamp) when the QR code expires.
	ExpiresAt int64 `json:"expires_at"`
	// The image_url_png string used to render QR code
	ImageURLPNG string `json:"image_url_png"`
	// The image_url_svg string used to render QR code
	ImageURLSVG string `json:"image_url_svg"`
}
type SetupIntentNextActionCashAppHandleRedirectOrDisplayQRCode struct {
	// The URL to the hosted Cash App Pay instructions page, which allows customers to view the QR code, and supports QR code refreshing on expiration.
	HostedInstructionsURL string `json:"hosted_instructions_url"`
	// The url for mobile redirect based auth
	MobileAuthURL string                                                           `json:"mobile_auth_url"`
	QRCode        *SetupIntentNextActionCashAppHandleRedirectOrDisplayQRCodeQRCode `json:"qr_code"`
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
	CashAppHandleRedirectOrDisplayQRCode *SetupIntentNextActionCashAppHandleRedirectOrDisplayQRCode `json:"cashapp_handle_redirect_or_display_qr_code"`
	RedirectToURL                        *SetupIntentNextActionRedirectToURL                        `json:"redirect_to_url"`
	// Type of the next action to perform. Refer to the other child attributes under `next_action` for available values. Examples include: `redirect_to_url`, `use_stripe_sdk`, `alipay_handle_redirect`, `oxxo_display_details`, or `verify_with_microdeposits`.
	Type SetupIntentNextActionType `json:"type"`
	// When confirming a SetupIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js.
	UseStripeSDK            *SetupIntentNextActionUseStripeSDK            `json:"use_stripe_sdk"`
	VerifyWithMicrodeposits *SetupIntentNextActionVerifyWithMicrodeposits `json:"verify_with_microdeposits"`
}

// Information about the [payment method configuration](https://stripe.com/docs/api/payment_method_configurations) used for this Setup Intent.
type SetupIntentPaymentMethodConfigurationDetails struct {
	// ID of the payment method configuration used.
	ID string `json:"id"`
	// ID of the parent payment method configuration used.
	Parent string `json:"parent"`
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
	// Currency supported by the bank account
	Currency       SetupIntentPaymentMethodOptionsACSSDebitCurrency        `json:"currency"`
	MandateOptions *SetupIntentPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options"`
	// Bank account verification method.
	VerificationMethod SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}
type SetupIntentPaymentMethodOptionsAmazonPay struct{}
type SetupIntentPaymentMethodOptionsBACSDebitMandateOptions struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'DDIC' or 'STRIPE'.
	ReferencePrefix string `json:"reference_prefix"`
}
type SetupIntentPaymentMethodOptionsBACSDebit struct {
	MandateOptions *SetupIntentPaymentMethodOptionsBACSDebitMandateOptions `json:"mandate_options"`
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
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure SetupIntentPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}
type SetupIntentPaymentMethodOptionsCardPresent struct{}
type SetupIntentPaymentMethodOptionsKlarna struct {
	// The currency of the setup intent. Three letter ISO currency code.
	Currency Currency `json:"currency"`
	// Preferred locale of the Klarna checkout page that the customer is redirected to.
	PreferredLocale string `json:"preferred_locale"`
}
type SetupIntentPaymentMethodOptionsLink struct {
	// [Deprecated] This is a legacy parameter that no longer has any function.
	// Deprecated:
	PersistentToken string `json:"persistent_token"`
}
type SetupIntentPaymentMethodOptionsPaypal struct {
	// The PayPal Billing Agreement ID (BAID). This is an ID generated by PayPal which represents the mandate between the merchant and the customer.
	BillingAgreementID string `json:"billing_agreement_id"`
}
type SetupIntentPaymentMethodOptionsSEPADebitMandateOptions struct {
	// Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'STRIPE'.
	ReferencePrefix string `json:"reference_prefix"`
}
type SetupIntentPaymentMethodOptionsSEPADebit struct {
	MandateOptions *SetupIntentPaymentMethodOptionsSEPADebitMandateOptions `json:"mandate_options"`
}
type SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters struct {
	// The account subcategories to use to filter for possible accounts to link. Valid subcategories are `checking` and `savings`.
	AccountSubcategories []SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersAccountSubcategory `json:"account_subcategories"`
}
type SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnections struct {
	Filters *SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsFilters `json:"filters"`
	// The list of permissions to request. The `payment_method` permission must be included.
	Permissions []SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPermission `json:"permissions"`
	// Data features requested to be retrieved upon account creation.
	Prefetch []SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnectionsPrefetch `json:"prefetch"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL string `json:"return_url"`
}
type SetupIntentPaymentMethodOptionsUSBankAccountMandateOptions struct {
	// Mandate collection method
	CollectionMethod SetupIntentPaymentMethodOptionsUSBankAccountMandateOptionsCollectionMethod `json:"collection_method"`
}
type SetupIntentPaymentMethodOptionsUSBankAccount struct {
	FinancialConnections *SetupIntentPaymentMethodOptionsUSBankAccountFinancialConnections `json:"financial_connections"`
	MandateOptions       *SetupIntentPaymentMethodOptionsUSBankAccountMandateOptions       `json:"mandate_options"`
	// Bank account verification method.
	VerificationMethod SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment method-specific configuration for this SetupIntent.
type SetupIntentPaymentMethodOptions struct {
	ACSSDebit     *SetupIntentPaymentMethodOptionsACSSDebit     `json:"acss_debit"`
	AmazonPay     *SetupIntentPaymentMethodOptionsAmazonPay     `json:"amazon_pay"`
	BACSDebit     *SetupIntentPaymentMethodOptionsBACSDebit     `json:"bacs_debit"`
	Card          *SetupIntentPaymentMethodOptionsCard          `json:"card"`
	CardPresent   *SetupIntentPaymentMethodOptionsCardPresent   `json:"card_present"`
	Klarna        *SetupIntentPaymentMethodOptionsKlarna        `json:"klarna"`
	Link          *SetupIntentPaymentMethodOptionsLink          `json:"link"`
	Paypal        *SetupIntentPaymentMethodOptionsPaypal        `json:"paypal"`
	SEPADebit     *SetupIntentPaymentMethodOptionsSEPADebit     `json:"sepa_debit"`
	USBankAccount *SetupIntentPaymentMethodOptionsUSBankAccount `json:"us_bank_account"`
}

// A SetupIntent guides you through the process of setting up and saving a customer's payment credentials for future payments.
// For example, you can use a SetupIntent to set up and save your customer's card without immediately collecting a payment.
// Later, you can use [PaymentIntents](https://stripe.com/docs/api#payment_intents) to drive the payment flow.
//
// Create a SetupIntent when you're ready to collect your customer's payment credentials.
// Don't maintain long-lived, unconfirmed SetupIntents because they might not be valid.
// The SetupIntent transitions through multiple [statuses](https://docs.stripe.com/payments/intents#intent-statuses) as it guides
// you through the setup process.
//
// Successful SetupIntents result in payment credentials that are optimized for future payments.
// For example, cardholders in [certain regions](https://stripe.com/guides/strong-customer-authentication) might need to be run through
// [Strong Customer Authentication](https://docs.stripe.com/strong-customer-authentication) during payment method collection
// to streamline later [off-session payments](https://docs.stripe.com/payments/setup-intents).
// If you use the SetupIntent with a [Customer](https://stripe.com/docs/api#setup_intent_object-customer),
// it automatically attaches the resulting payment method to that Customer after successful setup.
// We recommend using SetupIntents or [setup_future_usage](https://stripe.com/docs/api#payment_intent_object-setup_future_usage) on
// PaymentIntents to save payment methods to prevent saving invalid or unoptimized payment methods.
//
// By using SetupIntents, you can reduce friction for your customers, even as regulations change over time.
//
// Related guide: [Setup Intents API](https://docs.stripe.com/payments/setup-intents)
type SetupIntent struct {
	APIResource
	// ID of the Connect application that created the SetupIntent.
	Application *Application `json:"application"`
	// If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.
	//
	// It can only be used for this Stripe Account's own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.
	AttachToSelf bool `json:"attach_to_self"`
	// Settings for dynamic payment methods compatible with this Setup Intent
	AutomaticPaymentMethods *SetupIntentAutomaticPaymentMethods `json:"automatic_payment_methods"`
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
	// ID of the payment method used with this SetupIntent. If the payment method is `card_present` and isn't a digital wallet, then the [generated_card](https://docs.stripe.com/api/setup_attempts/object#setup_attempt_object-payment_method_details-card_present-generated_card) associated with the `latest_attempt` is attached to the Customer instead.
	PaymentMethod *PaymentMethod `json:"payment_method"`
	// Information about the [payment method configuration](https://stripe.com/docs/api/payment_method_configurations) used for this Setup Intent.
	PaymentMethodConfigurationDetails *SetupIntentPaymentMethodConfigurationDetails `json:"payment_method_configuration_details"`
	// Payment method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentPaymentMethodOptions `json:"payment_method_options"`
	// The list of payment method types (e.g. card) that this SetupIntent is allowed to set up. A list of valid payment method types can be found [here](https://docs.stripe.com/api/payment_methods/object#payment_method_object-type).
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
