//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The status of the most recent automated tax calculation for this session.
type CheckoutSessionAutomaticTaxStatus string

// List of values that CheckoutSessionAutomaticTaxStatus can take
const (
	CheckoutSessionAutomaticTaxStatusComplete               CheckoutSessionAutomaticTaxStatus = "complete"
	CheckoutSessionAutomaticTaxStatusFailed                 CheckoutSessionAutomaticTaxStatus = "failed"
	CheckoutSessionAutomaticTaxStatusRequiresLocationInputs CheckoutSessionAutomaticTaxStatus = "requires_location_inputs"
)

// Describes whether Checkout should collect the customer's billing address.
type CheckoutSessionBillingAddressCollection string

// List of values that CheckoutSessionBillingAddressCollection can take
const (
	CheckoutSessionBillingAddressCollectionAuto     CheckoutSessionBillingAddressCollection = "auto"
	CheckoutSessionBillingAddressCollectionRequired CheckoutSessionBillingAddressCollection = "required"
)

// If `opt_in`, the customer consents to receiving promotional communications
// from the merchant about this Checkout Session.
type CheckoutSessionConsentPromotions string

// List of values that CheckoutSessionConsentPromotions can take
const (
	CheckoutSessionConsentPromotionsOptIn  CheckoutSessionConsentPromotions = "opt_in"
	CheckoutSessionConsentPromotionsOptOut CheckoutSessionConsentPromotions = "opt_out"
)

// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
// Session will determine whether to display an option to opt into promotional communication
// from the merchant depending on the customer's locale. Only available to US merchants.
type CheckoutSessionConsentCollectionPromotions string

// List of values that CheckoutSessionConsentCollectionPromotions can take
const (
	CheckoutSessionConsentCollectionPromotionsAuto CheckoutSessionConsentCollectionPromotions = "auto"
)

// The customer's tax exempt status at time of checkout.
type CheckoutSessionCustomerDetailsTaxExempt string

// List of values that CheckoutSessionCustomerDetailsTaxExempt can take
const (
	CheckoutSessionCustomerDetailsTaxExemptExempt  CheckoutSessionCustomerDetailsTaxExempt = "exempt"
	CheckoutSessionCustomerDetailsTaxExemptNone    CheckoutSessionCustomerDetailsTaxExempt = "none"
	CheckoutSessionCustomerDetailsTaxExemptReverse CheckoutSessionCustomerDetailsTaxExempt = "reverse"
)

// The type of the tax ID, one of `eu_vat`, `br_cnpj`, `br_cpf`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, or `unknown`
type CheckoutSessionCustomerDetailsTaxIDsType string

// List of values that CheckoutSessionCustomerDetailsTaxIDsType can take
const (
	CheckoutSessionCustomerDetailsTaxIDsTypeAETRN    CheckoutSessionCustomerDetailsTaxIDsType = "ae_trn"
	CheckoutSessionCustomerDetailsTaxIDsTypeAUABN    CheckoutSessionCustomerDetailsTaxIDsType = "au_abn"
	CheckoutSessionCustomerDetailsTaxIDsTypeAUARN    CheckoutSessionCustomerDetailsTaxIDsType = "au_arn"
	CheckoutSessionCustomerDetailsTaxIDsTypeBRCNPJ   CheckoutSessionCustomerDetailsTaxIDsType = "br_cnpj"
	CheckoutSessionCustomerDetailsTaxIDsTypeBRCPF    CheckoutSessionCustomerDetailsTaxIDsType = "br_cpf"
	CheckoutSessionCustomerDetailsTaxIDsTypeCABN     CheckoutSessionCustomerDetailsTaxIDsType = "ca_bn"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAGSTHST CheckoutSessionCustomerDetailsTaxIDsType = "ca_gst_hst"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTBC  CheckoutSessionCustomerDetailsTaxIDsType = "ca_pst_bc"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTMB  CheckoutSessionCustomerDetailsTaxIDsType = "ca_pst_mb"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTSK  CheckoutSessionCustomerDetailsTaxIDsType = "ca_pst_sk"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAQST    CheckoutSessionCustomerDetailsTaxIDsType = "ca_qst"
	CheckoutSessionCustomerDetailsTaxIDsTypeCHVAT    CheckoutSessionCustomerDetailsTaxIDsType = "ch_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeCLTIN    CheckoutSessionCustomerDetailsTaxIDsType = "cl_tin"
	CheckoutSessionCustomerDetailsTaxIDsTypeESCIF    CheckoutSessionCustomerDetailsTaxIDsType = "es_cif"
	CheckoutSessionCustomerDetailsTaxIDsTypeEUVAT    CheckoutSessionCustomerDetailsTaxIDsType = "eu_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeGBVAT    CheckoutSessionCustomerDetailsTaxIDsType = "gb_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeGEVAT    CheckoutSessionCustomerDetailsTaxIDsType = "ge_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeHKBR     CheckoutSessionCustomerDetailsTaxIDsType = "hk_br"
	CheckoutSessionCustomerDetailsTaxIDsTypeIDNPWP   CheckoutSessionCustomerDetailsTaxIDsType = "id_npwp"
	CheckoutSessionCustomerDetailsTaxIDsTypeILVAT    CheckoutSessionCustomerDetailsTaxIDsType = "il_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeINGST    CheckoutSessionCustomerDetailsTaxIDsType = "in_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeJPCN     CheckoutSessionCustomerDetailsTaxIDsType = "jp_cn"
	CheckoutSessionCustomerDetailsTaxIDsTypeJPRN     CheckoutSessionCustomerDetailsTaxIDsType = "jp_rn"
	CheckoutSessionCustomerDetailsTaxIDsTypeKRBRN    CheckoutSessionCustomerDetailsTaxIDsType = "kr_brn"
	CheckoutSessionCustomerDetailsTaxIDsTypeLIUID    CheckoutSessionCustomerDetailsTaxIDsType = "li_uid"
	CheckoutSessionCustomerDetailsTaxIDsTypeMXRFC    CheckoutSessionCustomerDetailsTaxIDsType = "mx_rfc"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYFRP    CheckoutSessionCustomerDetailsTaxIDsType = "my_frp"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYITN    CheckoutSessionCustomerDetailsTaxIDsType = "my_itn"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYSST    CheckoutSessionCustomerDetailsTaxIDsType = "my_sst"
	CheckoutSessionCustomerDetailsTaxIDsTypeNOVAT    CheckoutSessionCustomerDetailsTaxIDsType = "no_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeNZGST    CheckoutSessionCustomerDetailsTaxIDsType = "nz_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeRUINN    CheckoutSessionCustomerDetailsTaxIDsType = "ru_inn"
	CheckoutSessionCustomerDetailsTaxIDsTypeRUKPP    CheckoutSessionCustomerDetailsTaxIDsType = "ru_kpp"
	CheckoutSessionCustomerDetailsTaxIDsTypeSAVAT    CheckoutSessionCustomerDetailsTaxIDsType = "sa_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeSGGST    CheckoutSessionCustomerDetailsTaxIDsType = "sg_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeSGUEN    CheckoutSessionCustomerDetailsTaxIDsType = "sg_uen"
	CheckoutSessionCustomerDetailsTaxIDsTypeTHVAT    CheckoutSessionCustomerDetailsTaxIDsType = "th_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeTWVAT    CheckoutSessionCustomerDetailsTaxIDsType = "tw_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeUAVAT    CheckoutSessionCustomerDetailsTaxIDsType = "ua_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeUnknown  CheckoutSessionCustomerDetailsTaxIDsType = "unknown"
	CheckoutSessionCustomerDetailsTaxIDsTypeUSEIN    CheckoutSessionCustomerDetailsTaxIDsType = "us_ein"
	CheckoutSessionCustomerDetailsTaxIDsTypeZAVAT    CheckoutSessionCustomerDetailsTaxIDsType = "za_vat"
)

// The mode of the Checkout Session.
type CheckoutSessionMode string

// List of values that CheckoutSessionMode can take
const (
	CheckoutSessionModePayment      CheckoutSessionMode = "payment"
	CheckoutSessionModeSetup        CheckoutSessionMode = "setup"
	CheckoutSessionModeSubscription CheckoutSessionMode = "subscription"
)

// List of Stripe products where this mandate can be selected automatically. Returned when the Session is in `setup` mode.
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultForInvoice      CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor = "invoice"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultForSubscription CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor = "subscription"
)

// Payment schedule for the mandate.
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleCombined CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "combined"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleInterval CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "interval"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleSporadic CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "sporadic"
)

// Transaction type of the mandate.
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Bank account verification method.
type CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethodInstant       CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// The payment status of the Checkout Session, one of `paid`, `unpaid`, or `no_payment_required`.
// You can use this value to decide when to fulfill your customer's order.
type CheckoutSessionPaymentStatus string

// List of values that CheckoutSessionPaymentStatus can take
const (
	CheckoutSessionPaymentStatusNoPaymentRequired CheckoutSessionPaymentStatus = "no_payment_required"
	CheckoutSessionPaymentStatusPaid              CheckoutSessionPaymentStatus = "paid"
	CheckoutSessionPaymentStatusUnpaid            CheckoutSessionPaymentStatus = "unpaid"
)

// The status of the Checkout Session, one of `open`, `complete`, or `expired`.
type CheckoutSessionStatus string

// List of values that CheckoutSessionStatus can take
const (
	CheckoutSessionStatusComplete CheckoutSessionStatus = "complete"
	CheckoutSessionStatusExpired  CheckoutSessionStatus = "expired"
	CheckoutSessionStatusOpen     CheckoutSessionStatus = "open"
)

// Describes the type of transaction being performed by Checkout in order to customize
// relevant text on the page, such as the submit button. `submit_type` can only be
// specified on Checkout Sessions in `payment` mode, but not Checkout Sessions
// in `subscription` or `setup` mode.
type CheckoutSessionSubmitType string

// List of values that CheckoutSessionSubmitType can take
const (
	CheckoutSessionSubmitTypeAuto   CheckoutSessionSubmitType = "auto"
	CheckoutSessionSubmitTypeBook   CheckoutSessionSubmitType = "book"
	CheckoutSessionSubmitTypeDonate CheckoutSessionSubmitType = "donate"
	CheckoutSessionSubmitTypePay    CheckoutSessionSubmitType = "pay"
)

// Returns a list of Checkout Sessions.
type CheckoutSessionListParams struct {
	ListParams    `form:"*"`
	PaymentIntent *string `form:"payment_intent"`
	Subscription  *string `form:"subscription"`
}

// Configure a Checkout Session that can be used to recover an expired session.
type CheckoutSessionAfterExpirationRecoveryParams struct {
	AllowPromotionCodes *bool `form:"allow_promotion_codes"`
	Enabled             *bool `form:"enabled"`
}

// Configure actions after a Checkout Session has expired.
type CheckoutSessionAfterExpirationParams struct {
	Recovery *CheckoutSessionAfterExpirationRecoveryParams `form:"recovery"`
}

// Settings for automatic tax lookup for this session and resulting payments, invoices, and subscriptions.
type CheckoutSessionAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// Configure fields for the Checkout Session to gather active consent from customers.
type CheckoutSessionConsentCollectionParams struct {
	Promotions *string `form:"promotions"`
}

// Controls what fields on Customer can be updated by the Checkout Session. Can only be provided when `customer` is provided.
type CheckoutSessionCustomerUpdateParams struct {
	Address  *string `form:"address"`
	Name     *string `form:"name"`
	Shipping *string `form:"shipping"`
}

// The coupon or promotion code to apply to this Session. Currently, only up to one may be specified.
type CheckoutSessionDiscountParams struct {
	Coupon        *string `form:"coupon"`
	PromotionCode *string `form:"promotion_code"`
}

// When set, provides configuration for this item's quantity to be adjusted by the customer during Checkout.
type CheckoutSessionLineItemAdjustableQuantityParams struct {
	Enabled *bool  `form:"enabled"`
	Maximum *int64 `form:"maximum"`
	Minimum *int64 `form:"minimum"`
}

// Data used to generate a new product object inline. One of `product` or `product_data` is required.
type CheckoutSessionLineItemPriceDataProductDataParams struct {
	Description *string           `form:"description"`
	Images      []*string         `form:"images"`
	Metadata    map[string]string `form:"metadata"`
	Name        *string           `form:"name"`
	TaxCode     *string           `form:"tax_code"`
}

// The recurring components of a price such as `interval` and `usage_type`.
type CheckoutSessionLineItemPriceDataRecurringParams struct {
	AggregateUsage  *string `form:"aggregate_usage"`
	Interval        *string `form:"interval"`
	IntervalCount   *int64  `form:"interval_count"`
	TrialPeriodDays *int64  `form:"trial_period_days"`
	UsageType       *string `form:"usage_type"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
type CheckoutSessionLineItemPriceDataParams struct {
	Currency          *string                                            `form:"currency"`
	Product           *string                                            `form:"product"`
	ProductData       *CheckoutSessionLineItemPriceDataProductDataParams `form:"product_data"`
	Recurring         *CheckoutSessionLineItemPriceDataRecurringParams   `form:"recurring"`
	TaxBehavior       *string                                            `form:"tax_behavior"`
	UnitAmount        *int64                                             `form:"unit_amount"`
	UnitAmountDecimal *float64                                           `form:"unit_amount_decimal,high_precision"`
}

// A list of items the customer is purchasing. Use this parameter to pass one-time or recurring [Prices](https://stripe.com/docs/api/prices).
//
// For `payment` mode, there is a maximum of 100 line items, however it is recommended to consolidate line items if there are more than a few dozen.
//
// For `subscription` mode, there is a maximum of 20 line items with recurring Prices and 20 line items with one-time Prices. Line items with one-time Prices in will be on the initial invoice only.
type CheckoutSessionLineItemParams struct {
	AdjustableQuantity *CheckoutSessionLineItemAdjustableQuantityParams `form:"adjustable_quantity"`
	Amount             *int64                                           `form:"amount"`
	Currency           *string                                          `form:"currency"`
	Description        *string                                          `form:"description"`
	DynamicTaxRates    []*string                                        `form:"dynamic_tax_rates"`
	Images             []*string                                        `form:"images"`
	Name               *string                                          `form:"name"`
	Price              *string                                          `form:"price"`
	PriceData          *CheckoutSessionLineItemPriceDataParams          `form:"price_data"`
	Quantity           *int64                                           `form:"quantity"`
	TaxRates           []*string                                        `form:"tax_rates"`
}

// The parameters used to automatically create a Transfer when the payment succeeds.
// For more information, see the PaymentIntents [use case for connected accounts](https://stripe.com/docs/payments/connected-accounts).
type CheckoutSessionPaymentIntentDataTransferDataParams struct {
	Amount      *int64  `form:"amount"`
	Destination *string `form:"destination"`
}

// A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.
type CheckoutSessionPaymentIntentDataParams struct {
	Params                    `form:"*"`
	ApplicationFeeAmount      *int64                                              `form:"application_fee_amount"`
	CaptureMethod             *string                                             `form:"capture_method"`
	Description               *string                                             `form:"description"`
	Metadata                  map[string]string                                   `form:"metadata"`
	OnBehalfOf                *string                                             `form:"on_behalf_of"`
	ReceiptEmail              *string                                             `form:"receipt_email"`
	SetupFutureUsage          *string                                             `form:"setup_future_usage"`
	Shipping                  *ShippingDetailsParams                              `form:"shipping"`
	StatementDescriptor       *string                                             `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                                             `form:"statement_descriptor_suffix"`
	TransferData              *CheckoutSessionPaymentIntentDataTransferDataParams `form:"transfer_data"`
	TransferGroup             *string                                             `form:"transfer_group"`
}

// Additional fields for Mandate creation
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	CustomMandateURL    *string   `form:"custom_mandate_url"`
	DefaultFor          []*string `form:"default_for"`
	IntervalDescription *string   `form:"interval_description"`
	PaymentSchedule     *string   `form:"payment_schedule"`
	TransactionType     *string   `form:"transaction_type"`
}

// contains details about the ACSS Debit payment method options.
type CheckoutSessionPaymentMethodOptionsACSSDebitParams struct {
	Currency           *string                                                           `form:"currency"`
	MandateOptions     *CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                           `form:"verification_method"`
}

// contains details about the Boleto payment method options.
type CheckoutSessionPaymentMethodOptionsBoletoParams struct {
	ExpiresAfterDays *int64 `form:"expires_after_days"`
}

// contains details about the OXXO payment method options.
type CheckoutSessionPaymentMethodOptionsOXXOParams struct {
	ExpiresAfterDays *int64 `form:"expires_after_days"`
}

// contains details about the Wechat Pay payment method options.
type CheckoutSessionPaymentMethodOptionsWechatPayParams struct {
	AppID  *string `form:"app_id"`
	Client *string `form:"client"`
}

// Payment-method-specific configuration.
type CheckoutSessionPaymentMethodOptionsParams struct {
	ACSSDebit *CheckoutSessionPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	Boleto    *CheckoutSessionPaymentMethodOptionsBoletoParams    `form:"boleto"`
	OXXO      *CheckoutSessionPaymentMethodOptionsOXXOParams      `form:"oxxo"`
	WechatPay *CheckoutSessionPaymentMethodOptionsWechatPayParams `form:"wechat_pay"`
}

// Controls phone number collection settings for the session.
//
// We recommend that you review your privacy policy and check with your legal contacts
// before using this feature. Learn more about [collecting phone numbers with Checkout](https://stripe.com/docs/payments/checkout/phone-numbers).
type CheckoutSessionPhoneNumberCollectionParams struct {
	Enabled *bool `form:"enabled"`
}

// A subset of parameters to be passed to SetupIntent creation for Checkout Sessions in `setup` mode.
type CheckoutSessionSetupIntentDataParams struct {
	Params      `form:"*"`
	Description *string `form:"description"`
	OnBehalfOf  *string `form:"on_behalf_of"`
}

// When set, provides configuration for Checkout to collect a shipping address from a customer.
type CheckoutSessionShippingAddressCollectionParams struct {
	AllowedCountries []*string `form:"allowed_countries"`
}

// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
type CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMaximumParams struct {
	Unit  *string `form:"unit"`
	Value *int64  `form:"value"`
}

// The lower bound of the estimated range. If empty, represents no lower bound.
type CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMinimumParams struct {
	Unit  *string `form:"unit"`
	Value *int64  `form:"value"`
}

// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
type CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateParams struct {
	Maximum *CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMaximumParams `form:"maximum"`
	Minimum *CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMinimumParams `form:"minimum"`
}

// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
type CheckoutSessionShippingOptionShippingRateDataFixedAmountParams struct {
	Amount   *int64  `form:"amount"`
	Currency *string `form:"currency"`
}

// Parameters to be passed to Shipping Rate creation for this shipping option
type CheckoutSessionShippingOptionShippingRateDataParams struct {
	DeliveryEstimate *CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateParams `form:"delivery_estimate"`
	DisplayName      *string                                                              `form:"display_name"`
	FixedAmount      *CheckoutSessionShippingOptionShippingRateDataFixedAmountParams      `form:"fixed_amount"`
	Metadata         map[string]string                                                    `form:"metadata"`
	TaxBehavior      *string                                                              `form:"tax_behavior"`
	TaxCode          *string                                                              `form:"tax_code"`
	Type             *string                                                              `form:"type"`
}

// The shipping rate options to apply to this Session.
type CheckoutSessionShippingOptionParams struct {
	ShippingRate     *string                                              `form:"shipping_rate"`
	ShippingRateData *CheckoutSessionShippingOptionShippingRateDataParams `form:"shipping_rate_data"`
}

// A list of items, each with an attached plan, that the customer is subscribing to. Prefer using `line_items`.
type CheckoutSessionSubscriptionDataItemsParams struct {
	Plan     *string   `form:"plan"`
	Quantity *int64    `form:"quantity"`
	TaxRates []*string `form:"tax_rates"`
}

// If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.
type CheckoutSessionSubscriptionDataTransferDataParams struct {
	AmountPercent *float64 `form:"amount_percent"`
	Destination   *string  `form:"destination"`
}

// A subset of parameters to be passed to subscription creation for Checkout Sessions in `subscription` mode.
type CheckoutSessionSubscriptionDataParams struct {
	Params                `form:"*"`
	ApplicationFeePercent *float64                                           `form:"application_fee_percent"`
	Coupon                *string                                            `form:"coupon"`
	DefaultTaxRates       []*string                                          `form:"default_tax_rates"`
	Items                 []*CheckoutSessionSubscriptionDataItemsParams      `form:"items"`
	Metadata              map[string]string                                  `form:"metadata"`
	TransferData          *CheckoutSessionSubscriptionDataTransferDataParams `form:"transfer_data"`
	TrialEnd              *int64                                             `form:"trial_end"`
	TrialFromPlan         *bool                                              `form:"trial_from_plan"`
	TrialPeriodDays       *int64                                             `form:"trial_period_days"`
}

// Controls tax ID collection settings for the session.
type CheckoutSessionTaxIDCollectionParams struct {
	Enabled *bool `form:"enabled"`
}

// Creates a Session object.
type CheckoutSessionParams struct {
	Params                    `form:"*"`
	AfterExpiration           *CheckoutSessionAfterExpirationParams           `form:"after_expiration"`
	AllowPromotionCodes       *bool                                           `form:"allow_promotion_codes"`
	AutomaticTax              *CheckoutSessionAutomaticTaxParams              `form:"automatic_tax"`
	BillingAddressCollection  *string                                         `form:"billing_address_collection"`
	CancelURL                 *string                                         `form:"cancel_url"`
	ClientReferenceID         *string                                         `form:"client_reference_id"`
	ConsentCollection         *CheckoutSessionConsentCollectionParams         `form:"consent_collection"`
	Customer                  *string                                         `form:"customer"`
	CustomerEmail             *string                                         `form:"customer_email"`
	CustomerUpdate            *CheckoutSessionCustomerUpdateParams            `form:"customer_update"`
	Discounts                 []*CheckoutSessionDiscountParams                `form:"discounts"`
	ExpiresAt                 *int64                                          `form:"expires_at"`
	LineItems                 []*CheckoutSessionLineItemParams                `form:"line_items"`
	Locale                    *string                                         `form:"locale"`
	Mode                      *string                                         `form:"mode"`
	PaymentIntentData         *CheckoutSessionPaymentIntentDataParams         `form:"payment_intent_data"`
	PaymentMethodOptions      *CheckoutSessionPaymentMethodOptionsParams      `form:"payment_method_options"`
	PaymentMethodTypes        []*string                                       `form:"payment_method_types"`
	PhoneNumberCollection     *CheckoutSessionPhoneNumberCollectionParams     `form:"phone_number_collection"`
	SetupIntentData           *CheckoutSessionSetupIntentDataParams           `form:"setup_intent_data"`
	ShippingAddressCollection *CheckoutSessionShippingAddressCollectionParams `form:"shipping_address_collection"`
	ShippingOptions           []*CheckoutSessionShippingOptionParams          `form:"shipping_options"`
	ShippingRates             []*string                                       `form:"shipping_rates"`
	SubmitType                *string                                         `form:"submit_type"`
	SubscriptionData          *CheckoutSessionSubscriptionDataParams          `form:"subscription_data"`
	SuccessURL                *string                                         `form:"success_url"`
	TaxIDCollection           *CheckoutSessionTaxIDCollectionParams           `form:"tax_id_collection"`
}

// A Session can be expired when it is in one of these statuses: open
//
// After it expires, a customer can't complete a Session and customers loading the Session see a message saying the Session is expired.
type CheckoutSessionExpireParams struct {
	Params `form:"*"`
}

// When retrieving a Checkout Session, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type CheckoutSessionListLineItemsParams struct {
	ListParams `form:"*"`
	Session    *string `form:"-"` // Included in URL
}

// When set, configuration used to recover the Checkout Session on expiry.
type CheckoutSessionAfterExpirationRecovery struct {
	AllowPromotionCodes bool   `json:"allow_promotion_codes"`
	Enabled             bool   `json:"enabled"`
	ExpiresAt           int64  `json:"expires_at"`
	URL                 string `json:"url"`
}

// When set, provides configuration for actions to take if this Checkout Session expires.
type CheckoutSessionAfterExpiration struct {
	Recovery *CheckoutSessionAfterExpirationRecovery `json:"recovery"`
}
type CheckoutSessionAutomaticTax struct {
	Enabled bool                              `json:"enabled"`
	Status  CheckoutSessionAutomaticTaxStatus `json:"status"`
}

// Results of `consent_collection` for this session.
type CheckoutSessionConsent struct {
	Promotions CheckoutSessionConsentPromotions `json:"promotions"`
}

// When set, provides configuration for the Checkout Session to gather active consent from customers.
type CheckoutSessionConsentCollection struct {
	Promotions CheckoutSessionConsentCollectionPromotions `json:"promotions"`
}

// The customer's tax IDs at time of checkout.
type CheckoutSessionCustomerDetailsTaxIDs struct {
	Type  CheckoutSessionCustomerDetailsTaxIDsType `json:"type"`
	Value string                                   `json:"value"`
}

// The customer details including the customer's tax exempt status and the customer's tax IDs. Only present on Sessions in `payment` or `subscription` mode.
type CheckoutSessionCustomerDetails struct {
	Email     string                                  `json:"email"`
	Phone     string                                  `json:"phone"`
	TaxExempt CheckoutSessionCustomerDetailsTaxExempt `json:"tax_exempt"`
	TaxIDs    []*CheckoutSessionCustomerDetailsTaxIDs `json:"tax_ids"`
}
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptions struct {
	CustomMandateURL    string                                                                    `json:"custom_mandate_url"`
	DefaultFor          []CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor    `json:"default_for"`
	IntervalDescription string                                                                    `json:"interval_description"`
	PaymentSchedule     CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule `json:"payment_schedule"`
	TransactionType     CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}
type CheckoutSessionPaymentMethodOptionsACSSDebit struct {
	Currency           string                                                         `json:"currency"`
	MandateOptions     *CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptions    `json:"mandate_options"`
	VerificationMethod CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}
type CheckoutSessionPaymentMethodOptionsBoleto struct {
	ExpiresAfterDays int64 `json:"expires_after_days"`
}
type CheckoutSessionPaymentMethodOptionsOXXO struct {
	ExpiresAfterDays int64 `json:"expires_after_days"`
}

// Payment-method-specific configuration for the PaymentIntent or SetupIntent of this CheckoutSession.
type CheckoutSessionPaymentMethodOptions struct {
	ACSSDebit *CheckoutSessionPaymentMethodOptionsACSSDebit `json:"acss_debit"`
	Boleto    *CheckoutSessionPaymentMethodOptionsBoleto    `json:"boleto"`
	OXXO      *CheckoutSessionPaymentMethodOptionsOXXO      `json:"oxxo"`
}
type CheckoutSessionPhoneNumberCollection struct {
	Enabled bool `json:"enabled"`
}

// When set, provides configuration for Checkout to collect a shipping address from a customer.
type CheckoutSessionShippingAddressCollection struct {
	AllowedCountries []string `json:"allowed_countries"`
}

// The shipping rate options applied to this Session.
type CheckoutSessionShippingOption struct {
	ShippingAmount int64         `json:"shipping_amount"`
	ShippingRate   *ShippingRate `json:"shipping_rate"`
}
type CheckoutSessionTaxIDCollection struct {
	Enabled bool `json:"enabled"`
}

// The aggregated line item discounts.
type CheckoutSessionTotalDetailsBreakdownDiscount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// The aggregated line item tax amounts by rate.
type CheckoutSessionTotalDetailsBreakdownTax struct {
	Amount  int64    `json:"amount"`
	Rate    *TaxRate `json:"rate"`
	TaxRate *TaxRate `json:"tax_rate"` // Do not use: use `Rate`
}
type CheckoutSessionTotalDetailsBreakdown struct {
	Discounts []*CheckoutSessionTotalDetailsBreakdownDiscount `json:"discounts"`
	Taxes     []*CheckoutSessionTotalDetailsBreakdownTax      `json:"taxes"`
}

// Tax and discount details for the computed total amount.
type CheckoutSessionTotalDetails struct {
	AmountDiscount int64                                 `json:"amount_discount"`
	AmountShipping int64                                 `json:"amount_shipping"`
	AmountTax      int64                                 `json:"amount_tax"`
	Breakdown      *CheckoutSessionTotalDetailsBreakdown `json:"breakdown"`
}

// A Checkout Session represents your customer's session as they pay for
// one-time purchases or subscriptions through [Checkout](https://stripe.com/docs/payments/checkout)
// or [Payment Links](https://stripe.com/docs/payments/payment-links). We recommend creating a
// new Session each time your customer attempts to pay.
//
// Once payment is successful, the Checkout Session will contain a reference
// to the [Customer](https://stripe.com/docs/api/customers), and either the successful
// [PaymentIntent](https://stripe.com/docs/api/payment_intents) or an active
// [Subscription](https://stripe.com/docs/api/subscriptions).
//
// You can create a Checkout Session on your server and pass its ID to the
// client to begin Checkout.
//
// Related guide: [Checkout Server Quickstart](https://stripe.com/docs/payments/checkout/api).
type CheckoutSession struct {
	APIResource
	AfterExpiration           *CheckoutSessionAfterExpiration           `json:"after_expiration"`
	AllowPromotionCodes       bool                                      `json:"allow_promotion_codes"`
	AmountSubtotal            int64                                     `json:"amount_subtotal"`
	AmountTotal               int64                                     `json:"amount_total"`
	AutomaticTax              *CheckoutSessionAutomaticTax              `json:"automatic_tax"`
	BillingAddressCollection  CheckoutSessionBillingAddressCollection   `json:"billing_address_collection"`
	CancelURL                 string                                    `json:"cancel_url"`
	ClientReferenceID         string                                    `json:"client_reference_id"`
	Consent                   *CheckoutSessionConsent                   `json:"consent"`
	ConsentCollection         *CheckoutSessionConsentCollection         `json:"consent_collection"`
	Currency                  Currency                                  `json:"currency"`
	Customer                  *Customer                                 `json:"customer"`
	CustomerDetails           *CheckoutSessionCustomerDetails           `json:"customer_details"`
	CustomerEmail             string                                    `json:"customer_email"`
	Deleted                   bool                                      `json:"deleted"`
	ExpiresAt                 int64                                     `json:"expires_at"`
	ID                        string                                    `json:"id"`
	LineItems                 *LineItemList                             `json:"line_items"`
	Livemode                  bool                                      `json:"livemode"`
	Locale                    string                                    `json:"locale"`
	Metadata                  map[string]string                         `json:"metadata"`
	Mode                      CheckoutSessionMode                       `json:"mode"`
	Object                    string                                    `json:"object"`
	PaymentIntent             *PaymentIntent                            `json:"payment_intent"`
	PaymentMethodOptions      *CheckoutSessionPaymentMethodOptions      `json:"payment_method_options"`
	PaymentMethodTypes        []string                                  `json:"payment_method_types"`
	PaymentStatus             CheckoutSessionPaymentStatus              `json:"payment_status"`
	PhoneNumberCollection     *CheckoutSessionPhoneNumberCollection     `json:"phone_number_collection"`
	RecoveredFrom             string                                    `json:"recovered_from"`
	SetupIntent               *SetupIntent                              `json:"setup_intent"`
	Shipping                  *ShippingDetails                          `json:"shipping"`
	ShippingAddressCollection *CheckoutSessionShippingAddressCollection `json:"shipping_address_collection"`
	ShippingOptions           []*CheckoutSessionShippingOption          `json:"shipping_options"`
	ShippingRate              *ShippingRate                             `json:"shipping_rate"`
	Status                    CheckoutSessionStatus                     `json:"status"`
	SubmitType                CheckoutSessionSubmitType                 `json:"submit_type"`
	Subscription              *Subscription                             `json:"subscription"`
	SuccessURL                string                                    `json:"success_url"`
	TaxIDCollection           *CheckoutSessionTaxIDCollection           `json:"tax_id_collection"`
	TotalDetails              *CheckoutSessionTotalDetails              `json:"total_details"`
	URL                       string                                    `json:"url"`
}

// UnmarshalJSON handles deserialization of a CheckoutSession.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *CheckoutSession) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type checkoutSession CheckoutSession
	var v checkoutSession
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = CheckoutSession(v)
	return nil
}

// CheckoutSessionList is a list of Sessions as retrieved from a list endpoint.
type CheckoutSessionList struct {
	APIResource
	ListMeta
	Data []*CheckoutSession `json:"data"`
}
