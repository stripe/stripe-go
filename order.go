//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The status of the most recent automated tax calculation for this Order.
type OrderAutomaticTaxStatus string

// List of values that OrderAutomaticTaxStatus can take
const (
	OrderAutomaticTaxStatusComplete               OrderAutomaticTaxStatus = "complete"
	OrderAutomaticTaxStatusFailed                 OrderAutomaticTaxStatus = "failed"
	OrderAutomaticTaxStatusRequiresLocationInputs OrderAutomaticTaxStatus = "requires_location_inputs"
)

// The type of credit to apply to the order, only `gift_card` currently supported.
type OrderCreditType string

// List of values that OrderCreditType can take
const (
	OrderCreditTypeGiftCard OrderCreditType = "gift_card"
)

// Payment schedule for the mandate.
type OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule string

// List of values that OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule can take
const (
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleCombined OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "combined"
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleInterval OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "interval"
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleSporadic OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "sporadic"
)

// Transaction type of the mandate.
type OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsACSSDebitSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsACSSDebitSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitSetupFutureUsageNone       OrderPaymentSettingsPaymentMethodOptionsACSSDebitSetupFutureUsage = "none"
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitSetupFutureUsageOffSession OrderPaymentSettingsPaymentMethodOptionsACSSDebitSetupFutureUsage = "off_session"
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitSetupFutureUsageOnSession  OrderPaymentSettingsPaymentMethodOptionsACSSDebitSetupFutureUsage = "on_session"
)

// Bank account verification method.
type OrderPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that OrderPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     OrderPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodInstant       OrderPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	OrderPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits OrderPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// Controls when the funds will be captured from the customer's account.
type OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethod string

// List of values that OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethod can take
const (
	OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethodAutomatic      OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethod = "automatic"
	OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethodAutomaticAsync OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethod = "automatic_async"
	OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethodManual         OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with the payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the order's Customer, if present, after the order's PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
//
// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
type OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpaySetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpaySetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpaySetupFutureUsageNone OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpaySetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsAlipaySetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsAlipaySetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsAlipaySetupFutureUsageNone       OrderPaymentSettingsPaymentMethodOptionsAlipaySetupFutureUsage = "none"
	OrderPaymentSettingsPaymentMethodOptionsAlipaySetupFutureUsageOffSession OrderPaymentSettingsPaymentMethodOptionsAlipaySetupFutureUsage = "off_session"
)

// Preferred language of the Bancontact authorization page that the customer is redirected to.
type OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguageDE OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage = "de"
	OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguageEN OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage = "en"
	OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguageFR OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage = "fr"
	OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguageNL OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage = "nl"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsBancontactSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsBancontactSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsBancontactSetupFutureUsageNone       OrderPaymentSettingsPaymentMethodOptionsBancontactSetupFutureUsage = "none"
	OrderPaymentSettingsPaymentMethodOptionsBancontactSetupFutureUsageOffSession OrderPaymentSettingsPaymentMethodOptionsBancontactSetupFutureUsage = "off_session"
)

// Controls when the funds will be captured from the customer's account.
type OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethod string

// List of values that OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethod can take
const (
	OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethodAutomatic      OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethod = "automatic"
	OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethodAutomaticAsync OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethod = "automatic_async"
	OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethodManual         OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with the payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the order's Customer, if present, after the order's PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
//
// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
type OrderPaymentSettingsPaymentMethodOptionsCardSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsCardSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsCardSetupFutureUsageNone       OrderPaymentSettingsPaymentMethodOptionsCardSetupFutureUsage = "none"
	OrderPaymentSettingsPaymentMethodOptionsCardSetupFutureUsageOffSession OrderPaymentSettingsPaymentMethodOptionsCardSetupFutureUsage = "off_session"
	OrderPaymentSettingsPaymentMethodOptionsCardSetupFutureUsageOnSession  OrderPaymentSettingsPaymentMethodOptionsCardSetupFutureUsage = "on_session"
)

// The desired country code of the bank account information. Permitted values include: `BE`, `DE`, `ES`, `FR`, `IE`, or `NL`.
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry string

// List of values that OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry can take
const (
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryBE OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "BE"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryDE OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "DE"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryES OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "ES"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryFR OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "FR"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryIE OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "IE"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountryNL OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry = "NL"
)

// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
//
// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType string

// List of values that OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType can take
const (
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeABA      OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "aba"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeIBAN     OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "iban"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSEPA     OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "sepa"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSortCode OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "sort_code"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSpei     OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "spei"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSwift    OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "swift"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeZengin   OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType = "zengin"
)

// The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType string

// List of values that OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType can take
const (
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeEUBankTransfer OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "eu_bank_transfer"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeGBBankTransfer OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "gb_bank_transfer"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeJPBankTransfer OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "jp_bank_transfer"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeMXBankTransfer OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "mx_bank_transfer"
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeUSBankTransfer OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType = "us_bank_transfer"
)

// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType string

// List of values that OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType can take
const (
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceFundingTypeBankTransfer OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType = "bank_transfer"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceSetupFutureUsageNone OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsIDEALSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsIDEALSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsIDEALSetupFutureUsageNone       OrderPaymentSettingsPaymentMethodOptionsIDEALSetupFutureUsage = "none"
	OrderPaymentSettingsPaymentMethodOptionsIDEALSetupFutureUsageOffSession OrderPaymentSettingsPaymentMethodOptionsIDEALSetupFutureUsage = "off_session"
)

// Controls when the funds will be captured from the customer's account.
type OrderPaymentSettingsPaymentMethodOptionsKlarnaCaptureMethod string

// List of values that OrderPaymentSettingsPaymentMethodOptionsKlarnaCaptureMethod can take
const (
	OrderPaymentSettingsPaymentMethodOptionsKlarnaCaptureMethodManual OrderPaymentSettingsPaymentMethodOptionsKlarnaCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsKlarnaSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsKlarnaSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsKlarnaSetupFutureUsageNone OrderPaymentSettingsPaymentMethodOptionsKlarnaSetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type OrderPaymentSettingsPaymentMethodOptionsLinkCaptureMethod string

// List of values that OrderPaymentSettingsPaymentMethodOptionsLinkCaptureMethod can take
const (
	OrderPaymentSettingsPaymentMethodOptionsLinkCaptureMethodManual OrderPaymentSettingsPaymentMethodOptionsLinkCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsLinkSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsLinkSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsLinkSetupFutureUsageNone       OrderPaymentSettingsPaymentMethodOptionsLinkSetupFutureUsage = "none"
	OrderPaymentSettingsPaymentMethodOptionsLinkSetupFutureUsageOffSession OrderPaymentSettingsPaymentMethodOptionsLinkSetupFutureUsage = "off_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsOXXOSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsOXXOSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsOXXOSetupFutureUsageNone OrderPaymentSettingsPaymentMethodOptionsOXXOSetupFutureUsage = "none"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsP24SetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsP24SetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsP24SetupFutureUsageNone OrderPaymentSettingsPaymentMethodOptionsP24SetupFutureUsage = "none"
)

// Controls when the funds will be captured from the customer's account.
type OrderPaymentSettingsPaymentMethodOptionsPaypalCaptureMethod string

// List of values that OrderPaymentSettingsPaymentMethodOptionsPaypalCaptureMethod can take
const (
	OrderPaymentSettingsPaymentMethodOptionsPaypalCaptureMethodManual OrderPaymentSettingsPaymentMethodOptionsPaypalCaptureMethod = "manual"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsPaypalSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsPaypalSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsPaypalSetupFutureUsageNone       OrderPaymentSettingsPaymentMethodOptionsPaypalSetupFutureUsage = "none"
	OrderPaymentSettingsPaymentMethodOptionsPaypalSetupFutureUsageOffSession OrderPaymentSettingsPaymentMethodOptionsPaypalSetupFutureUsage = "off_session"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsSEPADebitSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsSEPADebitSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsSEPADebitSetupFutureUsageNone       OrderPaymentSettingsPaymentMethodOptionsSEPADebitSetupFutureUsage = "none"
	OrderPaymentSettingsPaymentMethodOptionsSEPADebitSetupFutureUsageOffSession OrderPaymentSettingsPaymentMethodOptionsSEPADebitSetupFutureUsage = "off_session"
	OrderPaymentSettingsPaymentMethodOptionsSEPADebitSetupFutureUsageOnSession  OrderPaymentSettingsPaymentMethodOptionsSEPADebitSetupFutureUsage = "on_session"
)

// Preferred language of the SOFORT authorization page that the customer is redirected to.
type OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguageDE OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguage = "de"
	OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguageEN OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguage = "en"
	OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguageES OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguage = "es"
	OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguageFR OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguage = "fr"
	OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguageIT OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguage = "it"
	OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguageNL OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguage = "nl"
	OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguagePL OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguage = "pl"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsSofortSetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsSofortSetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsSofortSetupFutureUsageNone       OrderPaymentSettingsPaymentMethodOptionsSofortSetupFutureUsage = "none"
	OrderPaymentSettingsPaymentMethodOptionsSofortSetupFutureUsageOffSession OrderPaymentSettingsPaymentMethodOptionsSofortSetupFutureUsage = "off_session"
)

// The client type that the end customer will pay from
type OrderPaymentSettingsPaymentMethodOptionsWeChatPayClient string

// List of values that OrderPaymentSettingsPaymentMethodOptionsWeChatPayClient can take
const (
	OrderPaymentSettingsPaymentMethodOptionsWeChatPayClientAndroid OrderPaymentSettingsPaymentMethodOptionsWeChatPayClient = "android"
	OrderPaymentSettingsPaymentMethodOptionsWeChatPayClientIOS     OrderPaymentSettingsPaymentMethodOptionsWeChatPayClient = "ios"
	OrderPaymentSettingsPaymentMethodOptionsWeChatPayClientWeb     OrderPaymentSettingsPaymentMethodOptionsWeChatPayClient = "web"
)

// Indicates that you intend to make future payments with this PaymentIntent's payment method.
//
// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
//
// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
type OrderPaymentSettingsPaymentMethodOptionsWeChatPaySetupFutureUsage string

// List of values that OrderPaymentSettingsPaymentMethodOptionsWeChatPaySetupFutureUsage can take
const (
	OrderPaymentSettingsPaymentMethodOptionsWeChatPaySetupFutureUsageNone OrderPaymentSettingsPaymentMethodOptionsWeChatPaySetupFutureUsage = "none"
)

// The list of [payment method types](https://stripe.com/docs/payments/payment-methods/overview) to provide to the order's PaymentIntent. Do not include this attribute if you prefer to manage your payment methods from the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).
type OrderPaymentSettingsPaymentMethodType string

// List of values that OrderPaymentSettingsPaymentMethodType can take
const (
	OrderPaymentSettingsPaymentMethodTypeACSSDebit        OrderPaymentSettingsPaymentMethodType = "acss_debit"
	OrderPaymentSettingsPaymentMethodTypeAfterpayClearpay OrderPaymentSettingsPaymentMethodType = "afterpay_clearpay"
	OrderPaymentSettingsPaymentMethodTypeAlipay           OrderPaymentSettingsPaymentMethodType = "alipay"
	OrderPaymentSettingsPaymentMethodTypeAUBECSDebit      OrderPaymentSettingsPaymentMethodType = "au_becs_debit"
	OrderPaymentSettingsPaymentMethodTypeBACSDebit        OrderPaymentSettingsPaymentMethodType = "bacs_debit"
	OrderPaymentSettingsPaymentMethodTypeBancontact       OrderPaymentSettingsPaymentMethodType = "bancontact"
	OrderPaymentSettingsPaymentMethodTypeCard             OrderPaymentSettingsPaymentMethodType = "card"
	OrderPaymentSettingsPaymentMethodTypeCustomerBalance  OrderPaymentSettingsPaymentMethodType = "customer_balance"
	OrderPaymentSettingsPaymentMethodTypeEPS              OrderPaymentSettingsPaymentMethodType = "eps"
	OrderPaymentSettingsPaymentMethodTypeFPX              OrderPaymentSettingsPaymentMethodType = "fpx"
	OrderPaymentSettingsPaymentMethodTypeGiropay          OrderPaymentSettingsPaymentMethodType = "giropay"
	OrderPaymentSettingsPaymentMethodTypeGrabpay          OrderPaymentSettingsPaymentMethodType = "grabpay"
	OrderPaymentSettingsPaymentMethodTypeIDEAL            OrderPaymentSettingsPaymentMethodType = "ideal"
	OrderPaymentSettingsPaymentMethodTypeKlarna           OrderPaymentSettingsPaymentMethodType = "klarna"
	OrderPaymentSettingsPaymentMethodTypeLink             OrderPaymentSettingsPaymentMethodType = "link"
	OrderPaymentSettingsPaymentMethodTypeOXXO             OrderPaymentSettingsPaymentMethodType = "oxxo"
	OrderPaymentSettingsPaymentMethodTypeP24              OrderPaymentSettingsPaymentMethodType = "p24"
	OrderPaymentSettingsPaymentMethodTypePaypal           OrderPaymentSettingsPaymentMethodType = "paypal"
	OrderPaymentSettingsPaymentMethodTypeSEPADebit        OrderPaymentSettingsPaymentMethodType = "sepa_debit"
	OrderPaymentSettingsPaymentMethodTypeSofort           OrderPaymentSettingsPaymentMethodType = "sofort"
	OrderPaymentSettingsPaymentMethodTypeWeChatPay        OrderPaymentSettingsPaymentMethodType = "wechat_pay"
)

// The status of the underlying payment associated with this order, if any. Null when the order is `open`.
type OrderPaymentStatus string

// List of values that OrderPaymentStatus can take
const (
	OrderPaymentStatusCanceled              OrderPaymentStatus = "canceled"
	OrderPaymentStatusComplete              OrderPaymentStatus = "complete"
	OrderPaymentStatusNotRequired           OrderPaymentStatus = "not_required"
	OrderPaymentStatusProcessing            OrderPaymentStatus = "processing"
	OrderPaymentStatusRequiresAction        OrderPaymentStatus = "requires_action"
	OrderPaymentStatusRequiresCapture       OrderPaymentStatus = "requires_capture"
	OrderPaymentStatusRequiresConfirmation  OrderPaymentStatus = "requires_confirmation"
	OrderPaymentStatusRequiresPaymentMethod OrderPaymentStatus = "requires_payment_method"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type OrderShippingCostTaxTaxabilityReason string

// List of values that OrderShippingCostTaxTaxabilityReason can take
const (
	OrderShippingCostTaxTaxabilityReasonCustomerExempt       OrderShippingCostTaxTaxabilityReason = "customer_exempt"
	OrderShippingCostTaxTaxabilityReasonNotCollecting        OrderShippingCostTaxTaxabilityReason = "not_collecting"
	OrderShippingCostTaxTaxabilityReasonNotSubjectToTax      OrderShippingCostTaxTaxabilityReason = "not_subject_to_tax"
	OrderShippingCostTaxTaxabilityReasonNotSupported         OrderShippingCostTaxTaxabilityReason = "not_supported"
	OrderShippingCostTaxTaxabilityReasonPortionProductExempt OrderShippingCostTaxTaxabilityReason = "portion_product_exempt"
	OrderShippingCostTaxTaxabilityReasonPortionReducedRated  OrderShippingCostTaxTaxabilityReason = "portion_reduced_rated"
	OrderShippingCostTaxTaxabilityReasonPortionStandardRated OrderShippingCostTaxTaxabilityReason = "portion_standard_rated"
	OrderShippingCostTaxTaxabilityReasonProductExempt        OrderShippingCostTaxTaxabilityReason = "product_exempt"
	OrderShippingCostTaxTaxabilityReasonProductExemptHoliday OrderShippingCostTaxTaxabilityReason = "product_exempt_holiday"
	OrderShippingCostTaxTaxabilityReasonProportionallyRated  OrderShippingCostTaxTaxabilityReason = "proportionally_rated"
	OrderShippingCostTaxTaxabilityReasonReducedRated         OrderShippingCostTaxTaxabilityReason = "reduced_rated"
	OrderShippingCostTaxTaxabilityReasonReverseCharge        OrderShippingCostTaxTaxabilityReason = "reverse_charge"
	OrderShippingCostTaxTaxabilityReasonStandardRated        OrderShippingCostTaxTaxabilityReason = "standard_rated"
	OrderShippingCostTaxTaxabilityReasonTaxableBasisReduced  OrderShippingCostTaxTaxabilityReason = "taxable_basis_reduced"
	OrderShippingCostTaxTaxabilityReasonZeroRated            OrderShippingCostTaxTaxabilityReason = "zero_rated"
)

// The overall status of the order.
type OrderStatus string

// List of values that OrderStatus can take
const (
	OrderStatusCanceled   OrderStatus = "canceled"
	OrderStatusComplete   OrderStatus = "complete"
	OrderStatusOpen       OrderStatus = "open"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusSubmitted  OrderStatus = "submitted"
)

// Describes the purchaser's tax exemption status. One of `none`, `exempt`, or `reverse`.
type OrderTaxDetailsTaxExempt string

// List of values that OrderTaxDetailsTaxExempt can take
const (
	OrderTaxDetailsTaxExemptExempt  OrderTaxDetailsTaxExempt = "exempt"
	OrderTaxDetailsTaxExemptNone    OrderTaxDetailsTaxExempt = "none"
	OrderTaxDetailsTaxExemptReverse OrderTaxDetailsTaxExempt = "reverse"
)

// The type of the tax ID, one of `ad_nrt`, `ar_cuit`, `eu_vat`, `bo_tin`, `br_cnpj`, `br_cpf`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `eu_oss_vat`, `pe_ruc`, `ro_tin`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, `vn_tin`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, or `unknown`
type OrderTaxDetailsTaxIDType string

// List of values that OrderTaxDetailsTaxIDType can take
const (
	OrderTaxDetailsTaxIDTypeADNRT    OrderTaxDetailsTaxIDType = "ad_nrt"
	OrderTaxDetailsTaxIDTypeAETRN    OrderTaxDetailsTaxIDType = "ae_trn"
	OrderTaxDetailsTaxIDTypeARCUIT   OrderTaxDetailsTaxIDType = "ar_cuit"
	OrderTaxDetailsTaxIDTypeAUABN    OrderTaxDetailsTaxIDType = "au_abn"
	OrderTaxDetailsTaxIDTypeAUARN    OrderTaxDetailsTaxIDType = "au_arn"
	OrderTaxDetailsTaxIDTypeBGUIC    OrderTaxDetailsTaxIDType = "bg_uic"
	OrderTaxDetailsTaxIDTypeBOTIN    OrderTaxDetailsTaxIDType = "bo_tin"
	OrderTaxDetailsTaxIDTypeBRCNPJ   OrderTaxDetailsTaxIDType = "br_cnpj"
	OrderTaxDetailsTaxIDTypeBRCPF    OrderTaxDetailsTaxIDType = "br_cpf"
	OrderTaxDetailsTaxIDTypeCABN     OrderTaxDetailsTaxIDType = "ca_bn"
	OrderTaxDetailsTaxIDTypeCAGSTHST OrderTaxDetailsTaxIDType = "ca_gst_hst"
	OrderTaxDetailsTaxIDTypeCAPSTBC  OrderTaxDetailsTaxIDType = "ca_pst_bc"
	OrderTaxDetailsTaxIDTypeCAPSTMB  OrderTaxDetailsTaxIDType = "ca_pst_mb"
	OrderTaxDetailsTaxIDTypeCAPSTSK  OrderTaxDetailsTaxIDType = "ca_pst_sk"
	OrderTaxDetailsTaxIDTypeCAQST    OrderTaxDetailsTaxIDType = "ca_qst"
	OrderTaxDetailsTaxIDTypeCHVAT    OrderTaxDetailsTaxIDType = "ch_vat"
	OrderTaxDetailsTaxIDTypeCLTIN    OrderTaxDetailsTaxIDType = "cl_tin"
	OrderTaxDetailsTaxIDTypeCNTIN    OrderTaxDetailsTaxIDType = "cn_tin"
	OrderTaxDetailsTaxIDTypeCONIT    OrderTaxDetailsTaxIDType = "co_nit"
	OrderTaxDetailsTaxIDTypeCRTIN    OrderTaxDetailsTaxIDType = "cr_tin"
	OrderTaxDetailsTaxIDTypeDORCN    OrderTaxDetailsTaxIDType = "do_rcn"
	OrderTaxDetailsTaxIDTypeECRUC    OrderTaxDetailsTaxIDType = "ec_ruc"
	OrderTaxDetailsTaxIDTypeEGTIN    OrderTaxDetailsTaxIDType = "eg_tin"
	OrderTaxDetailsTaxIDTypeESCIF    OrderTaxDetailsTaxIDType = "es_cif"
	OrderTaxDetailsTaxIDTypeEUOSSVAT OrderTaxDetailsTaxIDType = "eu_oss_vat"
	OrderTaxDetailsTaxIDTypeEUVAT    OrderTaxDetailsTaxIDType = "eu_vat"
	OrderTaxDetailsTaxIDTypeGBVAT    OrderTaxDetailsTaxIDType = "gb_vat"
	OrderTaxDetailsTaxIDTypeGEVAT    OrderTaxDetailsTaxIDType = "ge_vat"
	OrderTaxDetailsTaxIDTypeHKBR     OrderTaxDetailsTaxIDType = "hk_br"
	OrderTaxDetailsTaxIDTypeHUTIN    OrderTaxDetailsTaxIDType = "hu_tin"
	OrderTaxDetailsTaxIDTypeIDNPWP   OrderTaxDetailsTaxIDType = "id_npwp"
	OrderTaxDetailsTaxIDTypeILVAT    OrderTaxDetailsTaxIDType = "il_vat"
	OrderTaxDetailsTaxIDTypeINGST    OrderTaxDetailsTaxIDType = "in_gst"
	OrderTaxDetailsTaxIDTypeISVAT    OrderTaxDetailsTaxIDType = "is_vat"
	OrderTaxDetailsTaxIDTypeJPCN     OrderTaxDetailsTaxIDType = "jp_cn"
	OrderTaxDetailsTaxIDTypeJPRN     OrderTaxDetailsTaxIDType = "jp_rn"
	OrderTaxDetailsTaxIDTypeJPTRN    OrderTaxDetailsTaxIDType = "jp_trn"
	OrderTaxDetailsTaxIDTypeKEPIN    OrderTaxDetailsTaxIDType = "ke_pin"
	OrderTaxDetailsTaxIDTypeKRBRN    OrderTaxDetailsTaxIDType = "kr_brn"
	OrderTaxDetailsTaxIDTypeLIUID    OrderTaxDetailsTaxIDType = "li_uid"
	OrderTaxDetailsTaxIDTypeMXRFC    OrderTaxDetailsTaxIDType = "mx_rfc"
	OrderTaxDetailsTaxIDTypeMYFRP    OrderTaxDetailsTaxIDType = "my_frp"
	OrderTaxDetailsTaxIDTypeMYITN    OrderTaxDetailsTaxIDType = "my_itn"
	OrderTaxDetailsTaxIDTypeMYSST    OrderTaxDetailsTaxIDType = "my_sst"
	OrderTaxDetailsTaxIDTypeNOVAT    OrderTaxDetailsTaxIDType = "no_vat"
	OrderTaxDetailsTaxIDTypeNZGST    OrderTaxDetailsTaxIDType = "nz_gst"
	OrderTaxDetailsTaxIDTypePERUC    OrderTaxDetailsTaxIDType = "pe_ruc"
	OrderTaxDetailsTaxIDTypePHTIN    OrderTaxDetailsTaxIDType = "ph_tin"
	OrderTaxDetailsTaxIDTypeROTIN    OrderTaxDetailsTaxIDType = "ro_tin"
	OrderTaxDetailsTaxIDTypeRSPIB    OrderTaxDetailsTaxIDType = "rs_pib"
	OrderTaxDetailsTaxIDTypeRUINN    OrderTaxDetailsTaxIDType = "ru_inn"
	OrderTaxDetailsTaxIDTypeRUKPP    OrderTaxDetailsTaxIDType = "ru_kpp"
	OrderTaxDetailsTaxIDTypeSAVAT    OrderTaxDetailsTaxIDType = "sa_vat"
	OrderTaxDetailsTaxIDTypeSGGST    OrderTaxDetailsTaxIDType = "sg_gst"
	OrderTaxDetailsTaxIDTypeSGUEN    OrderTaxDetailsTaxIDType = "sg_uen"
	OrderTaxDetailsTaxIDTypeSITIN    OrderTaxDetailsTaxIDType = "si_tin"
	OrderTaxDetailsTaxIDTypeSVNIT    OrderTaxDetailsTaxIDType = "sv_nit"
	OrderTaxDetailsTaxIDTypeTHVAT    OrderTaxDetailsTaxIDType = "th_vat"
	OrderTaxDetailsTaxIDTypeTRTIN    OrderTaxDetailsTaxIDType = "tr_tin"
	OrderTaxDetailsTaxIDTypeTWVAT    OrderTaxDetailsTaxIDType = "tw_vat"
	OrderTaxDetailsTaxIDTypeUAVAT    OrderTaxDetailsTaxIDType = "ua_vat"
	OrderTaxDetailsTaxIDTypeUnknown  OrderTaxDetailsTaxIDType = "unknown"
	OrderTaxDetailsTaxIDTypeUSEIN    OrderTaxDetailsTaxIDType = "us_ein"
	OrderTaxDetailsTaxIDTypeUYRUC    OrderTaxDetailsTaxIDType = "uy_ruc"
	OrderTaxDetailsTaxIDTypeVERIF    OrderTaxDetailsTaxIDType = "ve_rif"
	OrderTaxDetailsTaxIDTypeVNTIN    OrderTaxDetailsTaxIDType = "vn_tin"
	OrderTaxDetailsTaxIDTypeZAVAT    OrderTaxDetailsTaxIDType = "za_vat"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type OrderTotalDetailsBreakdownTaxTaxabilityReason string

// List of values that OrderTotalDetailsBreakdownTaxTaxabilityReason can take
const (
	OrderTotalDetailsBreakdownTaxTaxabilityReasonCustomerExempt       OrderTotalDetailsBreakdownTaxTaxabilityReason = "customer_exempt"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonNotCollecting        OrderTotalDetailsBreakdownTaxTaxabilityReason = "not_collecting"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonNotSubjectToTax      OrderTotalDetailsBreakdownTaxTaxabilityReason = "not_subject_to_tax"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonNotSupported         OrderTotalDetailsBreakdownTaxTaxabilityReason = "not_supported"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonPortionProductExempt OrderTotalDetailsBreakdownTaxTaxabilityReason = "portion_product_exempt"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonPortionReducedRated  OrderTotalDetailsBreakdownTaxTaxabilityReason = "portion_reduced_rated"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonPortionStandardRated OrderTotalDetailsBreakdownTaxTaxabilityReason = "portion_standard_rated"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonProductExempt        OrderTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonProductExemptHoliday OrderTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt_holiday"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonProportionallyRated  OrderTotalDetailsBreakdownTaxTaxabilityReason = "proportionally_rated"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonReducedRated         OrderTotalDetailsBreakdownTaxTaxabilityReason = "reduced_rated"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonReverseCharge        OrderTotalDetailsBreakdownTaxTaxabilityReason = "reverse_charge"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonStandardRated        OrderTotalDetailsBreakdownTaxTaxabilityReason = "standard_rated"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonTaxableBasisReduced  OrderTotalDetailsBreakdownTaxTaxabilityReason = "taxable_basis_reduced"
	OrderTotalDetailsBreakdownTaxTaxabilityReasonZeroRated            OrderTotalDetailsBreakdownTaxTaxabilityReason = "zero_rated"
)

// Settings for automatic tax calculation for this order.
type OrderAutomaticTaxParams struct {
	// Enable automatic tax calculation which will automatically compute tax rates on this order.
	Enabled *bool `form:"enabled"`
}

// Billing details for the customer. If a customer is provided, this will be automatically populated with values from that customer if override values are not provided.
type OrderBillingDetailsParams struct {
	// The billing address provided by the customer.
	Address *AddressParams `form:"address"`
	// The billing email provided by the customer.
	Email *string `form:"email"`
	// The billing name provided by the customer.
	Name *string `form:"name"`
	// The billing phone number provided by the customer.
	Phone *string `form:"phone"`
}

// The credits to apply to the order, only `gift_card` currently supported.
type OrderCreditParams struct {
	// The gift card to apply to the order.
	GiftCard *string `form:"gift_card"`
	// The type of credit to apply to the order, only `gift_card` currently supported.
	Type *string `form:"type"`
}

// The coupons, promotion codes, and/or discounts to apply to the order.
type OrderDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
	// ID of the promotion code to create a new discount for.
	PromotionCode *string `form:"promotion_code"`
}

// The discounts applied to this line item.
type OrderLineItemDiscountParams struct {
	// ID of the coupon to create a new discount for.
	Coupon *string `form:"coupon"`
	// ID of an existing discount on the object (or one of its ancestors) to reuse.
	Discount *string `form:"discount"`
}

// Data used to generate a new Price object inline.
//
// The `price_data` parameter is an alternative to using the `product` or `price` parameters. If you create products upfront and configure a `Product.default_price`, pass the `product` parameter when creating a line item. If you prefer not to define products upfront, or if you charge variable prices, pass the `price_data` parameter to describe the price for this line item.
//
// Each time you pass `price_data` we create a Price for the product. This Price is hidden in both the Dashboard and API lists and cannot be reused.
type OrderLineItemPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// ID of the product this price belongs to.
	//
	// Use this to implement a variable-pricing model in your integration. This is required if `product_data` is not specifed.
	Product *string `form:"product"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// The dimensions of this product for shipping purposes.
type OrderLineItemProductDataPackageDimensionsParams struct {
	// Height, in inches. Maximum precision is 2 decimal places.
	Height *float64 `form:"height"`
	// Length, in inches. Maximum precision is 2 decimal places.
	Length *float64 `form:"length"`
	// Weight, in ounces. Maximum precision is 2 decimal places.
	Weight *float64 `form:"weight"`
	// Width, in inches. Maximum precision is 2 decimal places.
	Width *float64 `form:"width"`
}

// Defines a Product inline and adds it to the Order.
//
// `product_data` is an alternative to the `product` parameter. If you created a Product upfront, use the `product` parameter to refer to the existing Product. But if you prefer not to create Products upfront, pass the `product_data` parameter to define a Product inline as part of configuring the Order.
//
// `product_data` automatically creates a Product, just as if you had manually created the Product. If a Product with the same ID already exists, then `product_data` re-uses it to avoid duplicates.
type OrderLineItemProductDataParams struct {
	// The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description"`
	// A unique identifier for this product.
	//
	// `product_data` automatically creates a Product with this ID. If a Product with the same ID already exists, then `product_data` re-uses it to avoid duplicates. If any of the fields in the existing Product are different from the values in `product_data`, `product_data` updates the existing Product with the new information. So set `product_data[id]` to the same string every time you sell the same product, but don't re-use the same string for different products.
	ID *string `form:"id"`
	// A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name"`
	// The dimensions of this product for shipping purposes.
	PackageDimensions *OrderLineItemProductDataPackageDimensionsParams `form:"package_dimensions"`
	// Whether this product is shipped (i.e., physical goods).
	Shippable *bool `form:"shippable"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
	// A URL of a publicly-accessible webpage for this product.
	URL *string `form:"url"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *OrderLineItemProductDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A list of line items the customer is ordering. Each line item includes information about the product, the quantity, and the resulting cost.
type OrderLineItemParams struct {
	// The description for the line item. Will default to the name of the associated product.
	Description *string `form:"description"`
	// The discounts applied to this line item.
	Discounts []*OrderLineItemDiscountParams `form:"discounts"`
	// The ID of an existing line item on the order.
	ID *string `form:"id"`
	// The ID of a [Price](https://stripe.com/docs/api/prices) to add to the Order.
	//
	// The `price` parameter is an alternative to using the `product` parameter. If each of your products are sold at a single price, you can set `Product.default_price` and then pass the `product` parameter when creating a line item. If your products are sold at several possible prices, use the `price` parameter to explicitly specify which one to use.
	Price *string `form:"price"`
	// Data used to generate a new Price object inline.
	//
	// The `price_data` parameter is an alternative to using the `product` or `price` parameters. If you create products upfront and configure a `Product.default_price`, pass the `product` parameter when creating a line item. If you prefer not to define products upfront, or if you charge variable prices, pass the `price_data` parameter to describe the price for this line item.
	//
	// Each time you pass `price_data` we create a Price for the product. This Price is hidden in both the Dashboard and API lists and cannot be reused.
	PriceData *OrderLineItemPriceDataParams `form:"price_data"`
	// The ID of a [Product](https://stripe.com/docs/api/products) to add to the Order.
	//
	// The product must have a `default_price` specified. Otherwise, specify the price by passing the `price` or `price_data` parameter.
	Product *string `form:"product"`
	// Defines a Product inline and adds it to the Order.
	//
	// `product_data` is an alternative to the `product` parameter. If you created a Product upfront, use the `product` parameter to refer to the existing Product. But if you prefer not to create Products upfront, pass the `product_data` parameter to define a Product inline as part of configuring the Order.
	//
	// `product_data` automatically creates a Product, just as if you had manually created the Product. If a Product with the same ID already exists, then `product_data` re-uses it to avoid duplicates.
	ProductData *OrderLineItemProductDataParams `form:"product_data"`
	// The quantity of the line item.
	Quantity *int64 `form:"quantity"`
	// The tax rates applied to this line item.
	TaxRates []*string `form:"tax_rates"`
}

// Additional fields for Mandate creation
type OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// A URL for custom mandate text to render during confirmation step.
	// The URL will be rendered with additional GET parameters `payment_intent` and `payment_intent_client_secret` when confirming a Payment Intent,
	// or `setup_intent` and `setup_intent_client_secret` when confirming a Setup Intent.
	CustomMandateURL *string `form:"custom_mandate_url"`
	// Description of the mandate interval. Only required if 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription *string `form:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule *string `form:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type"`
}

// If paying by `acss_debit`, this sub-hash contains details about the ACSS Debit payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsACSSDebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// If paying by `afterpay_clearpay`, this sub-hash contains details about the AfterpayClearpay payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayParams struct {
	// Controls when the funds will be captured from the customer's account.
	//
	// If provided, this parameter will override the top-level `capture_method` when finalizing the payment with this payment method type.
	//
	// If `capture_method` is already set on the PaymentIntent, providing an empty value for this parameter will unset the stored value for this payment method type.
	CaptureMethod *string `form:"capture_method"`
	// An internal identifier or reference this payment corresponds to. The identifier is limited to 128 characters and may contain only letters, digits, underscores, backslashes and dashes.
	Reference *string `form:"reference"`
	// Indicates that you intend to make future payments with the payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the order's Customer, if present, after the order's PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// If paying by `alipay`, this sub-hash contains details about the Alipay payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsAlipayParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsBancontactParams struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage *string `form:"preferred_language"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsCardParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// Indicates that you intend to make future payments with the payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the order's Customer, if present, after the order's PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Configuration for the eu_bank_transfer funding type.
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information. Permitted values include: `BE`, `DE`, `ES`, `FR`, `IE`, or `NL`.
	Country *string `form:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for the eu_bank_transfer funding type.
	EUBankTransfer *OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer"`
	// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
	//
	// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
	RequestedAddressTypes []*string `form:"requested_address_types"`
	// The list of bank transfer types that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.
	Type *string `form:"type"`
}

// If paying by `customer_balance`, this sub-hash contains details about the Customer Balance payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType *string `form:"funding_type"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// If paying by `ideal`, this sub-hash contains details about the iDEAL payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsIDEALParams struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// If paying by `klarna`, this sub-hash contains details about the Klarna payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsKlarnaParams struct {
	// Controls when the funds will be captured from the customer's account.
	//
	// If provided, this parameter will override the top-level `capture_method` when finalizing the payment with this payment method type.
	//
	// If `capture_method` is already set on the PaymentIntent, providing an empty value for this parameter will unset the stored value for this payment method type.
	CaptureMethod *string `form:"capture_method"`
	// Preferred language of the Klarna authorization page that the customer is redirected to
	PreferredLocale *string `form:"preferred_locale"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// If paying by `link`, this sub-hash contains details about the Link payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsLinkParams struct {
	// Controls when the funds will be captured from the customer's account.
	//
	// If provided, this parameter will override the top-level `capture_method` when finalizing the payment with this payment method type.
	//
	// If `capture_method` is already set on the PaymentIntent, providing an empty value for this parameter will unset the stored value for this payment method type.
	CaptureMethod *string `form:"capture_method"`
	// [Deprecated] This is a legacy parameter that no longer has any function.
	PersistentToken *string `form:"persistent_token"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// If paying by `oxxo`, this sub-hash contains details about the OXXO payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsOXXOParams struct {
	// The number of calendar days before an OXXO voucher expires. For example, if you create an OXXO voucher on Monday and you set expires_after_days to 2, the OXXO invoice will expire on Wednesday at 23:59 America/Mexico_City time.
	ExpiresAfterDays *int64 `form:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// If paying by `p24`, this sub-hash contains details about the P24 payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsP24Params struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Confirm that the payer has accepted the P24 terms and conditions.
	TOSShownAndAccepted *bool `form:"tos_shown_and_accepted"`
}

// If paying by `paypal`, this sub-hash contains details about the PayPal payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsPaypalParams struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod *string `form:"capture_method"`
	// [Preferred locale](https://stripe.com/docs/payments/paypal/supported-locales) of the PayPal checkout page that the customer is redirected to.
	PreferredLocale *string `form:"preferred_locale"`
	// A reference of the PayPal transaction visible to customer which is mapped to PayPal's invoice ID. This must be a globally unique ID if you have configured in your PayPal settings to block multiple payments per invoice ID.
	Reference *string `form:"reference"`
	// A reference of the PayPal transaction visible to customer which is mapped to PayPal's invoice ID. This must be a globally unique ID if you have configured in your PayPal settings to block multiple payments per invoice ID.
	ReferenceID *string `form:"reference_id"`
	// The risk correlation ID for an on-session payment using a saved PayPal payment method.
	RiskCorrelationID *string `form:"risk_correlation_id"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// Additional fields for Mandate creation
type OrderPaymentSettingsPaymentMethodOptionsSEPADebitMandateOptionsParams struct{}

// If paying by `sepa_debit`, this sub-hash contains details about the SEPA Debit payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsSEPADebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *OrderPaymentSettingsPaymentMethodOptionsSEPADebitMandateOptionsParams `form:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// If paying by `sofort`, this sub-hash contains details about the Sofort payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsSofortParams struct {
	// Language shown to the payer on redirect.
	PreferredLanguage *string `form:"preferred_language"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// If paying by `wechat_pay`, this sub-hash contains details about the WeChat Pay payment method options to pass to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsWeChatPayParams struct {
	// The app ID registered with WeChat Pay. Only required when client is ios or android.
	AppID *string `form:"app_id"`
	// The client type that the end customer will pay from
	Client *string `form:"client"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage *string `form:"setup_future_usage"`
}

// PaymentMethod-specific configuration to provide to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptionsParams struct {
	// If paying by `acss_debit`, this sub-hash contains details about the ACSS Debit payment method options to pass to the order's PaymentIntent.
	ACSSDebit *OrderPaymentSettingsPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	// If paying by `afterpay_clearpay`, this sub-hash contains details about the AfterpayClearpay payment method options to pass to the order's PaymentIntent.
	AfterpayClearpay *OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayParams `form:"afterpay_clearpay"`
	// If paying by `alipay`, this sub-hash contains details about the Alipay payment method options to pass to the order's PaymentIntent.
	Alipay *OrderPaymentSettingsPaymentMethodOptionsAlipayParams `form:"alipay"`
	// If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the order's PaymentIntent.
	Bancontact *OrderPaymentSettingsPaymentMethodOptionsBancontactParams `form:"bancontact"`
	// If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the order's PaymentIntent.
	Card *OrderPaymentSettingsPaymentMethodOptionsCardParams `form:"card"`
	// If paying by `customer_balance`, this sub-hash contains details about the Customer Balance payment method options to pass to the order's PaymentIntent.
	CustomerBalance *OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceParams `form:"customer_balance"`
	// If paying by `ideal`, this sub-hash contains details about the iDEAL payment method options to pass to the order's PaymentIntent.
	IDEAL *OrderPaymentSettingsPaymentMethodOptionsIDEALParams `form:"ideal"`
	// If paying by `klarna`, this sub-hash contains details about the Klarna payment method options to pass to the order's PaymentIntent.
	Klarna *OrderPaymentSettingsPaymentMethodOptionsKlarnaParams `form:"klarna"`
	// If paying by `link`, this sub-hash contains details about the Link payment method options to pass to the order's PaymentIntent.
	Link *OrderPaymentSettingsPaymentMethodOptionsLinkParams `form:"link"`
	// If paying by `oxxo`, this sub-hash contains details about the OXXO payment method options to pass to the order's PaymentIntent.
	OXXO *OrderPaymentSettingsPaymentMethodOptionsOXXOParams `form:"oxxo"`
	// If paying by `p24`, this sub-hash contains details about the P24 payment method options to pass to the order's PaymentIntent.
	P24 *OrderPaymentSettingsPaymentMethodOptionsP24Params `form:"p24"`
	// If paying by `paypal`, this sub-hash contains details about the PayPal payment method options to pass to the order's PaymentIntent.
	Paypal *OrderPaymentSettingsPaymentMethodOptionsPaypalParams `form:"paypal"`
	// If paying by `sepa_debit`, this sub-hash contains details about the SEPA Debit payment method options to pass to the order's PaymentIntent.
	SEPADebit *OrderPaymentSettingsPaymentMethodOptionsSEPADebitParams `form:"sepa_debit"`
	// If paying by `sofort`, this sub-hash contains details about the Sofort payment method options to pass to the order's PaymentIntent.
	Sofort *OrderPaymentSettingsPaymentMethodOptionsSofortParams `form:"sofort"`
	// If paying by `wechat_pay`, this sub-hash contains details about the WeChat Pay payment method options to pass to the order's PaymentIntent.
	WeChatPay *OrderPaymentSettingsPaymentMethodOptionsWeChatPayParams `form:"wechat_pay"`
}

// Provides configuration for completing a transfer for the order after it is paid.
type OrderPaymentSettingsTransferDataParams struct {
	// The amount that will be transferred automatically when the order is paid. If no amount is set, the full amount is transferred. There cannot be any line items with recurring prices when using this field.
	Amount *int64 `form:"amount"`
	// ID of the Connected account receiving the transfer.
	Destination *string `form:"destination"`
}

// Settings describing how the order should configure generated PaymentIntents.
type OrderPaymentSettingsParams struct {
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account.
	ApplicationFeeAmount *int64 `form:"application_fee_amount"`
	// PaymentMethod-specific configuration to provide to the order's PaymentIntent.
	PaymentMethodOptions *OrderPaymentSettingsPaymentMethodOptionsParams `form:"payment_method_options"`
	// The list of [payment method types](https://stripe.com/docs/payments/payment-methods/overview) to provide to the order's PaymentIntent. Do not include this attribute if you prefer to manage your payment methods from the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).
	PaymentMethodTypes []*string `form:"payment_method_types"`
	// The URL to redirect the customer to after they authenticate their payment.
	ReturnURL *string `form:"return_url"`
	// For non-card charges, you can use this value as the complete description that appears on your customers' statements. Must contain at least one letter, maximum 22 characters.
	StatementDescriptor *string `form:"statement_descriptor"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the prefix (shortened descriptor) or statement descriptor that's set on the account to form the complete statement descriptor. Maximum 22 characters for the concatenated descriptor.
	StatementDescriptorSuffix *string `form:"statement_descriptor_suffix"`
	// Provides configuration for completing a transfer for the order after it is paid.
	TransferData *OrderPaymentSettingsTransferDataParams `form:"transfer_data"`
}

// Payment information associated with the order, including payment settings.
type OrderPaymentParams struct {
	// Settings describing how the order should configure generated PaymentIntents.
	Settings *OrderPaymentSettingsParams `form:"settings"`
}

// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
type OrderShippingCostShippingRateDataDeliveryEstimateMaximumParams struct {
	// A unit of time.
	Unit *string `form:"unit"`
	// Must be greater than 0.
	Value *int64 `form:"value"`
}

// The lower bound of the estimated range. If empty, represents no lower bound.
type OrderShippingCostShippingRateDataDeliveryEstimateMinimumParams struct {
	// A unit of time.
	Unit *string `form:"unit"`
	// Must be greater than 0.
	Value *int64 `form:"value"`
}

// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
type OrderShippingCostShippingRateDataDeliveryEstimateParams struct {
	// The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.
	Maximum *OrderShippingCostShippingRateDataDeliveryEstimateMaximumParams `form:"maximum"`
	// The lower bound of the estimated range. If empty, represents no lower bound.
	Minimum *OrderShippingCostShippingRateDataDeliveryEstimateMinimumParams `form:"minimum"`
}

// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type OrderShippingCostShippingRateDataFixedAmountCurrencyOptionsParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
}

// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
type OrderShippingCostShippingRateDataFixedAmountParams struct {
	// A non-negative integer in cents representing how much to charge.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions map[string]*OrderShippingCostShippingRateDataFixedAmountCurrencyOptionsParams `form:"currency_options"`
}

// Parameters to create a new ad-hoc shipping rate for this order.
type OrderShippingCostShippingRateDataParams struct {
	// The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DeliveryEstimate *OrderShippingCostShippingRateDataDeliveryEstimateParams `form:"delivery_estimate"`
	// The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DisplayName *string `form:"display_name"`
	// Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.
	FixedAmount *OrderShippingCostShippingRateDataFixedAmountParams `form:"fixed_amount"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior *string `form:"tax_behavior"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.
	TaxCode *string `form:"tax_code"`
	// The type of calculation to use on the shipping rate. Can only be `fixed_amount` for now.
	Type *string `form:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *OrderShippingCostShippingRateDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Settings for the customer cost of shipping for this order.
type OrderShippingCostParams struct {
	// The ID of the shipping rate to use for this order.
	ShippingRate *string `form:"shipping_rate"`
	// Parameters to create a new ad-hoc shipping rate for this order.
	ShippingRateData *OrderShippingCostShippingRateDataParams `form:"shipping_rate_data"`
}

// Shipping details for the order.
type OrderShippingDetailsParams struct {
	// The shipping address for the order.
	Address *AddressParams `form:"address"`
	// The name of the recipient of the order.
	Name *string `form:"name"`
	// The phone number (including extension) for the recipient of the order.
	Phone *string `form:"phone"`
}

// The purchaser's tax IDs to be used for this order.
type OrderTaxDetailsTaxIDParams struct {
	// Type of the tax ID, one of `ad_nrt`, `ae_trn`, `ar_cuit`, `au_abn`, `au_arn`, `bg_uic`, `bo_tin`, `br_cnpj`, `br_cpf`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `ch_vat`, `cl_tin`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `eg_tin`, `es_cif`, `eu_oss_vat`, `eu_vat`, `gb_vat`, `ge_vat`, `hk_br`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `ke_pin`, `kr_brn`, `li_uid`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `no_vat`, `nz_gst`, `pe_ruc`, `ph_tin`, `ro_tin`, `rs_pib`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `sv_nit`, `th_vat`, `tr_tin`, `tw_vat`, `ua_vat`, `us_ein`, `uy_ruc`, `ve_rif`, `vn_tin`, or `za_vat`
	Type *string `form:"type"`
	// Value of the tax ID.
	Value *string `form:"value"`
}

// Additional tax details about the purchaser to be used for this order.
type OrderTaxDetailsParams struct {
	// The purchaser's tax exemption status. One of `none`, `exempt`, or `reverse`.
	TaxExempt *string `form:"tax_exempt"`
	// The purchaser's tax IDs to be used for this order.
	TaxIDs []*OrderTaxDetailsTaxIDParams `form:"tax_ids"`
}

// Creates a new open order object.
type OrderParams struct {
	Params `form:"*"`
	// Settings for automatic tax calculation for this order.
	AutomaticTax *OrderAutomaticTaxParams `form:"automatic_tax"`
	// Billing details for the customer. If a customer is provided, this will be automatically populated with values from that customer if override values are not provided.
	BillingDetails *OrderBillingDetailsParams `form:"billing_details"`
	// The credits to apply to the order, only `gift_card` currently supported. Pass the empty string `""` to unset this field.
	Credits []*OrderCreditParams `form:"credits"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The customer associated with this order.
	Customer *string `form:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// The coupons, promotion codes, and/or discounts to apply to the order. Pass the empty string `""` to unset this field.
	Discounts []*OrderDiscountParams `form:"discounts"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The IP address of the purchaser for this order.
	IPAddress *string `form:"ip_address"`
	// A list of line items the customer is ordering. Each line item includes information about the product, the quantity, and the resulting cost.
	LineItems []*OrderLineItemParams `form:"line_items"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Payment information associated with the order, including payment settings.
	Payment *OrderPaymentParams `form:"payment"`
	// Settings for the customer cost of shipping for this order.
	ShippingCost *OrderShippingCostParams `form:"shipping_cost"`
	// Shipping details for the order.
	ShippingDetails *OrderShippingDetailsParams `form:"shipping_details"`
	// Additional tax details about the purchaser to be used for this order.
	TaxDetails *OrderTaxDetailsParams `form:"tax_details"`
}

// AddExpand appends a new field to expand.
func (p *OrderParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *OrderParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Returns a list of your orders. The orders are returned sorted by creation date, with the most recently created orders appearing first.
type OrderListParams struct {
	ListParams `form:"*"`
	// Only return orders for the given customer.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *OrderListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Submitting an Order transitions the status to processing and creates a PaymentIntent object so the order can be paid. If the Order has an amount_total of 0, no PaymentIntent object will be created. Once the order is submitted, its contents cannot be changed, unless the [reopen](https://stripe.com/docs/api#reopen_order) method is called.
type OrderSubmitParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// `expected_total` should always be set to the order's `amount_total` field. If they don't match, submitting the order will fail. This helps detect race conditions where something else concurrently modifies the order.
	ExpectedTotal *int64 `form:"expected_total"`
}

// AddExpand appends a new field to expand.
func (p *OrderSubmitParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Cancels the order as well as the payment intent if one is attached.
type OrderCancelParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *OrderCancelParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Reopens a submitted order.
type OrderReopenParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *OrderReopenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// When retrieving an order, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type OrderListLineItemsParams struct {
	ListParams `form:"*"`
	ID         *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *OrderListLineItemsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type OrderAutomaticTax struct {
	// Whether Stripe automatically computes tax on this Order.
	Enabled bool `json:"enabled"`
	// The status of the most recent automated tax calculation for this Order.
	Status OrderAutomaticTaxStatus `json:"status"`
}

// Customer billing details associated with the order.
type OrderBillingDetails struct {
	// Billing address for the order.
	Address *Address `json:"address"`
	// Email address for the order.
	Email string `json:"email"`
	// Full name for the order.
	Name string `json:"name"`
	// Billing phone number for the order (including extension).
	Phone string `json:"phone"`
}

// Details for a gift card.
type OrderCreditGiftCard struct {
	// The token of the gift card applied to the order
	Card string `json:"card"`
}

// The credits applied to the Order. At most 10 credits can be applied to an Order.
type OrderCredit struct {
	// The amount of this credit to apply to the order.
	Amount int64 `json:"amount"`
	// Details for a gift card.
	GiftCard *OrderCreditGiftCard `json:"gift_card"`
	// Line items on this order that are ineligible for this credit
	IneligibleLineItems []string `json:"ineligible_line_items"`
	// The type of credit to apply to the order, only `gift_card` currently supported.
	Type OrderCreditType `json:"type"`
}

// Indicates whether order has been opted into using [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods) to manage payment method types.
type OrderPaymentSettingsAutomaticPaymentMethods struct {
	// Whether this Order has been opted into managing payment method types via the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).
	Enabled bool `json:"enabled"`
}
type OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions struct {
	// A URL for custom mandate text
	CustomMandateURL string `json:"custom_mandate_url"`
	// Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription string `json:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule `json:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}
type OrderPaymentSettingsPaymentMethodOptionsACSSDebit struct {
	MandateOptions *OrderPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsACSSDebitSetupFutureUsage `json:"setup_future_usage"`
	// Bank account verification method.
	VerificationMethod OrderPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}
type OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpay struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethod `json:"capture_method"`
	// Order identifier shown to the user in Afterpay's online portal. We recommend using a value that helps you answer any questions a customer might have about the payment. The identifier is limited to 128 characters and may contain only letters, digits, underscores, backslashes and dashes.
	Reference string `json:"reference"`
	// Indicates that you intend to make future payments with the payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the order's Customer, if present, after the order's PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpaySetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsAlipay struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsAlipaySetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsBancontact struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage OrderPaymentSettingsPaymentMethodOptionsBancontactPreferredLanguage `json:"preferred_language"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsBancontactSetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsCard struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethod `json:"capture_method"`
	// Indicates that you intend to make future payments with the payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the order's Customer, if present, after the order's PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	//
	// If `setup_future_usage` is already set and you are performing a request using a publishable key, you may only update the value from `on_session` to `off_session`.
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsCardSetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer struct {
	// The desired country code of the bank account information. Permitted values include: `BE`, `DE`, `ES`, `FR`, `IE`, or `NL`.
	Country OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferCountry `json:"country"`
}
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransfer struct {
	EUBankTransfer *OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransfer `json:"eu_bank_transfer"`
	// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.
	//
	// Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
	RequestedAddressTypes []OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressType `json:"requested_address_types"`
	// The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.
	Type OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType `json:"type"`
}
type OrderPaymentSettingsPaymentMethodOptionsCustomerBalance struct {
	BankTransfer *OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransfer `json:"bank_transfer"`
	// The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.
	FundingType OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceFundingType `json:"funding_type"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceSetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsIDEAL struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsIDEALSetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsKlarna struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod OrderPaymentSettingsPaymentMethodOptionsKlarnaCaptureMethod `json:"capture_method"`
	// Preferred locale of the Klarna checkout page that the customer is redirected to.
	PreferredLocale string `json:"preferred_locale"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsKlarnaSetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsLink struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod OrderPaymentSettingsPaymentMethodOptionsLinkCaptureMethod `json:"capture_method"`
	// [Deprecated] This is a legacy parameter that no longer has any function.
	PersistentToken string `json:"persistent_token"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsLinkSetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsOXXO struct {
	// The number of calendar days before an OXXO invoice expires. For example, if you create an OXXO invoice on Monday and you set expires_after_days to 2, the OXXO invoice will expire on Wednesday at 23:59 America/Mexico_City time.
	ExpiresAfterDays int64 `json:"expires_after_days"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsOXXOSetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsP24 struct {
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsP24SetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsPaypal struct {
	// Controls when the funds will be captured from the customer's account.
	CaptureMethod OrderPaymentSettingsPaymentMethodOptionsPaypalCaptureMethod `json:"capture_method"`
	// Preferred locale of the PayPal checkout page that the customer is redirected to.
	PreferredLocale string `json:"preferred_locale"`
	// A reference of the PayPal transaction visible to customer which is mapped to PayPal's invoice ID. This must be a globally unique ID if you have configured in your PayPal settings to block multiple payments per invoice ID.
	Reference string `json:"reference"`
	// A reference of the PayPal transaction visible to customer which is mapped to PayPal's invoice ID. This must be a globally unique ID if you have configured in your PayPal settings to block multiple payments per invoice ID.
	ReferenceID string `json:"reference_id"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsPaypalSetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsSEPADebitMandateOptions struct{}
type OrderPaymentSettingsPaymentMethodOptionsSEPADebit struct {
	MandateOptions *OrderPaymentSettingsPaymentMethodOptionsSEPADebitMandateOptions `json:"mandate_options"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsSEPADebitSetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsSofort struct {
	// Preferred language of the SOFORT authorization page that the customer is redirected to.
	PreferredLanguage OrderPaymentSettingsPaymentMethodOptionsSofortPreferredLanguage `json:"preferred_language"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsSofortSetupFutureUsage `json:"setup_future_usage"`
}
type OrderPaymentSettingsPaymentMethodOptionsWeChatPay struct {
	// The app ID registered with WeChat Pay. Only required when client is ios or android.
	AppID string `json:"app_id"`
	// The client type that the end customer will pay from
	Client OrderPaymentSettingsPaymentMethodOptionsWeChatPayClient `json:"client"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.
	//
	// Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.
	//
	// When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage OrderPaymentSettingsPaymentMethodOptionsWeChatPaySetupFutureUsage `json:"setup_future_usage"`
}

// PaymentMethod-specific configuration to provide to the order's PaymentIntent.
type OrderPaymentSettingsPaymentMethodOptions struct {
	ACSSDebit        *OrderPaymentSettingsPaymentMethodOptionsACSSDebit        `json:"acss_debit"`
	AfterpayClearpay *OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpay `json:"afterpay_clearpay"`
	Alipay           *OrderPaymentSettingsPaymentMethodOptionsAlipay           `json:"alipay"`
	Bancontact       *OrderPaymentSettingsPaymentMethodOptionsBancontact       `json:"bancontact"`
	Card             *OrderPaymentSettingsPaymentMethodOptionsCard             `json:"card"`
	CustomerBalance  *OrderPaymentSettingsPaymentMethodOptionsCustomerBalance  `json:"customer_balance"`
	IDEAL            *OrderPaymentSettingsPaymentMethodOptionsIDEAL            `json:"ideal"`
	Klarna           *OrderPaymentSettingsPaymentMethodOptionsKlarna           `json:"klarna"`
	Link             *OrderPaymentSettingsPaymentMethodOptionsLink             `json:"link"`
	OXXO             *OrderPaymentSettingsPaymentMethodOptionsOXXO             `json:"oxxo"`
	P24              *OrderPaymentSettingsPaymentMethodOptionsP24              `json:"p24"`
	Paypal           *OrderPaymentSettingsPaymentMethodOptionsPaypal           `json:"paypal"`
	SEPADebit        *OrderPaymentSettingsPaymentMethodOptionsSEPADebit        `json:"sepa_debit"`
	Sofort           *OrderPaymentSettingsPaymentMethodOptionsSofort           `json:"sofort"`
	WeChatPay        *OrderPaymentSettingsPaymentMethodOptionsWeChatPay        `json:"wechat_pay"`
}

// Provides configuration for completing a transfer for the order after it is paid.
type OrderPaymentSettingsTransferData struct {
	// The amount that will be transferred automatically when the order is paid. If no amount is set, the full amount is transferred. There cannot be any line items with recurring prices when using this field.
	Amount int64 `json:"amount"`
	// ID of the Connected account receiving the transfer.
	Destination *Account `json:"destination"`
}

// Settings describing how the order should configure generated PaymentIntents.
type OrderPaymentSettings struct {
	// The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account.
	ApplicationFeeAmount int64 `json:"application_fee_amount"`
	// Indicates whether order has been opted into using [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods) to manage payment method types.
	AutomaticPaymentMethods *OrderPaymentSettingsAutomaticPaymentMethods `json:"automatic_payment_methods"`
	// PaymentMethod-specific configuration to provide to the order's PaymentIntent.
	PaymentMethodOptions *OrderPaymentSettingsPaymentMethodOptions `json:"payment_method_options"`
	// The list of [payment method types](https://stripe.com/docs/payments/payment-methods/overview) to provide to the order's PaymentIntent. Do not include this attribute if you prefer to manage your payment methods from the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).
	PaymentMethodTypes []OrderPaymentSettingsPaymentMethodType `json:"payment_method_types"`
	// The URL to redirect the customer to after they authenticate their payment.
	ReturnURL string `json:"return_url"`
	// For non-card charges, you can use this value as the complete description that appears on your customers' statements. Must contain at least one letter, maximum 22 characters.
	StatementDescriptor string `json:"statement_descriptor"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the prefix (shortened descriptor) or statement descriptor that's set on the account to form the complete statement descriptor. Maximum 22 characters for the concatenated descriptor.
	StatementDescriptorSuffix string `json:"statement_descriptor_suffix"`
	// Provides configuration for completing a transfer for the order after it is paid.
	TransferData *OrderPaymentSettingsTransferData `json:"transfer_data"`
}
type OrderPayment struct {
	// ID of the payment intent associated with this order. Null when the order is `open`.
	PaymentIntent *PaymentIntent `json:"payment_intent"`
	// Settings describing how the order should configure generated PaymentIntents.
	Settings *OrderPaymentSettings `json:"settings"`
	// The status of the underlying payment associated with this order, if any. Null when the order is `open`.
	Status OrderPaymentStatus `json:"status"`
}

// The taxes applied to the shipping rate.
type OrderShippingCostTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.
	//
	// Related guide: [Tax rates](https://stripe.com/docs/billing/taxes/tax-rates)
	Rate *TaxRate `json:"rate"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason OrderShippingCostTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
}

// The details of the customer cost of shipping, including the customer chosen ShippingRate.
type OrderShippingCost struct {
	// Total shipping cost before any discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total tax amount applied due to shipping costs. If no tax was applied, defaults to 0.
	AmountTax int64 `json:"amount_tax"`
	// Total shipping cost after discounts and taxes are applied.
	AmountTotal int64 `json:"amount_total"`
	// The ID of the ShippingRate for this order.
	ShippingRate *ShippingRate `json:"shipping_rate"`
	// The taxes applied to the shipping rate.
	Taxes []*OrderShippingCostTax `json:"taxes"`
}

// Customer shipping information associated with the order.
type OrderShippingDetails struct {
	// Recipient shipping address. Required if the order includes products that are shippable.
	Address *Address `json:"address"`
	// Recipient name.
	Name string `json:"name"`
	// Recipient phone (including extension).
	Phone string `json:"phone"`
}

// The purchaser's tax IDs to be used in calculation of tax for this Order.
type OrderTaxDetailsTaxID struct {
	// The type of the tax ID, one of `ad_nrt`, `ar_cuit`, `eu_vat`, `bo_tin`, `br_cnpj`, `br_cpf`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `eu_oss_vat`, `pe_ruc`, `ro_tin`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, `vn_tin`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, or `unknown`
	Type OrderTaxDetailsTaxIDType `json:"type"`
	// The value of the tax ID.
	Value string `json:"value"`
}
type OrderTaxDetails struct {
	// Describes the purchaser's tax exemption status. One of `none`, `exempt`, or `reverse`.
	TaxExempt OrderTaxDetailsTaxExempt `json:"tax_exempt"`
	// The purchaser's tax IDs to be used in calculation of tax for this Order.
	TaxIDs []*OrderTaxDetailsTaxID `json:"tax_ids"`
}

// The aggregated discounts.
type OrderTotalDetailsBreakdownDiscount struct {
	// The amount discounted.
	Amount int64 `json:"amount"`
	// A discount represents the actual application of a [coupon](https://stripe.com/docs/api#coupons) or [promotion code](https://stripe.com/docs/api#promotion_codes).
	// It contains information about when the discount began, when it will end, and what it is applied to.
	//
	// Related guide: [Applying discounts to subscriptions](https://stripe.com/docs/billing/subscriptions/discounts)
	Discount *Discount `json:"discount"`
}

// The aggregated tax amounts by rate.
type OrderTotalDetailsBreakdownTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.
	//
	// Related guide: [Tax rates](https://stripe.com/docs/billing/taxes/tax-rates)
	Rate *TaxRate `json:"rate"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason OrderTotalDetailsBreakdownTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
}
type OrderTotalDetailsBreakdown struct {
	// The aggregated discounts.
	Discounts []*OrderTotalDetailsBreakdownDiscount `json:"discounts"`
	// The aggregated tax amounts by rate.
	Taxes []*OrderTotalDetailsBreakdownTax `json:"taxes"`
}
type OrderTotalDetails struct {
	AmountCredit int64 `json:"amount_credit"`
	// This is the sum of all the discounts.
	AmountDiscount int64 `json:"amount_discount"`
	// This is the sum of all the shipping amounts.
	AmountShipping int64 `json:"amount_shipping"`
	// This is the sum of all the tax amounts.
	AmountTax int64                       `json:"amount_tax"`
	Breakdown *OrderTotalDetailsBreakdown `json:"breakdown"`
}

// An Order describes a purchase being made by a customer, including the
// products & quantities being purchased, the order status, the payment information,
// and the billing/shipping details.
//
// Related guide: [Orders overview](https://stripe.com/docs/orders)
type Order struct {
	APIResource
	AmountRemaining int64 `json:"amount_remaining"`
	// Order cost before any discounts or taxes are applied. A positive integer representing the subtotal of the order in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency).
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total order cost after discounts and taxes are applied. A positive integer representing the cost of the order in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). To submit an order, the total must be either 0 or at least $0.50 USD or [equivalent in charge currency](https://stripe.com/docs/currencies#minimum-and-maximum-charge-amounts).
	AmountTotal int64 `json:"amount_total"`
	// ID of the Connect application that created the Order, if any.
	Application  *Application       `json:"application"`
	AutomaticTax *OrderAutomaticTax `json:"automatic_tax"`
	// Customer billing details associated with the order.
	BillingDetails *OrderBillingDetails `json:"billing_details"`
	// The client secret of this Order. Used for client-side retrieval using a publishable key.
	//
	// The client secret can be used to complete a payment for an Order from your frontend. It should not be stored, logged, embedded in URLs, or exposed to anyone other than the customer. Make sure that you have TLS enabled on any page that includes the client secret.
	//
	// Refer to our docs for [creating and processing an order](https://stripe.com/docs/orders-beta/create-and-process) to learn about how client_secret should be handled.
	ClientSecret string `json:"client_secret"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The credits applied to the Order. At most 10 credits can be applied to an Order.
	Credits []*OrderCredit `json:"credits"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The customer which this orders belongs to.
	Customer *Customer `json:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// The discounts applied to the order. Use `expand[]=discounts` to expand each discount.
	Discounts []*Discount `json:"discounts"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A recent IP address of the purchaser used for tax reporting and tax location inference.
	IPAddress string `json:"ip_address"`
	// A list of line items the customer is ordering. Each line item includes information about the product, the quantity, and the resulting cost. There is a maximum of 100 line items.
	LineItems *LineItemList `json:"line_items"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object  string        `json:"object"`
	Payment *OrderPayment `json:"payment"`
	// The details of the customer cost of shipping, including the customer chosen ShippingRate.
	ShippingCost *OrderShippingCost `json:"shipping_cost"`
	// Customer shipping information associated with the order.
	ShippingDetails *OrderShippingDetails `json:"shipping_details"`
	// The overall status of the order.
	Status       OrderStatus        `json:"status"`
	TaxDetails   *OrderTaxDetails   `json:"tax_details"`
	TotalDetails *OrderTotalDetails `json:"total_details"`
}

// OrderList is a list of Orders as retrieved from a list endpoint.
type OrderList struct {
	APIResource
	ListMeta
	Data []*Order `json:"data"`
}
