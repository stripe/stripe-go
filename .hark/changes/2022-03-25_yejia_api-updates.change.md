---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1437
is_stripe_api_change: true
released_in_version: 72.96.0
---

* Add support for PayNow and US Bank Accounts Debits payments
    * **Charge** ([API ref](https://stripe.com/docs/api/charges/object#charge_object-payment_method_details))
        * Add support for `PayNow` and `USBankAccount` on `ChargePaymentMethodDetails`
    * **Mandate** ([API ref](https://stripe.com/docs/api/mandates/object#mandate_object-payment_method_details))
        * Add support for `USBankAccount` on `MandatePaymentMethodDetails`
    * **Payment Intent** ([API ref](https://stripe.com/docs/api/payment_intents/object#payment_intent_object-payment_method_options))
        * Add support for `PayNow` and `USBankAccount` on `PaymentIntentPaymentMethodOptions`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentConfirmPaymentMethodDataParams`, and `PaymentIntentConfirmPaymentMethodOptionsParams`
        * Add support for `PayNowDisplayQRCode` on `PaymentIntentNextAction`
    * **Setup Intent** ([API ref](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-payment_method_options))
        * Add support for `USBankAccount` on `SetupIntentPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptions`, and `SetupIntentConfirmPaymentMethodOptionsParams`
    * **Setup Attempt** ([API ref](https://stripe.com/docs/api/setup_attempts/object#setup_attempt_object-payment_method_details))
        * Add support for `USBankAccount` on `SetupAttemptPaymentMethodDetails`
    * **Payment Method** ([API ref](https://stripe.com/docs/api/payment_methods/object#payment_method_object-paynow))
        * Add support for `PayNow` and `USBankAccount` on `PaymentMethod` and `PaymentMethodParams`
        * Add support for new values `paynow` and `us_bank_account` on enum `PaymentMethodType`
    * **Checkout Session** ([API ref](https://stripe.com/docs/api/checkout/sessions/create#create_checkout_session-payment_method_types))
        * Add support for `USBankAccount` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
    * **Invoice** ([API ref](https://stripe.com/docs/api/invoices/object#invoice_object-payment_settings-payment_method_types))
        * Add support for `USBankAccount` on `InvoicePaymentSettingsPaymentMethodOptions` and `InvoicePaymentSettingsPaymentMethodOptionsParams`
        * Add support for new values `paynow` and `us_bank_account` on enum `InvoicePaymentSettingsPaymentMethodTypes`
    * **Subscription** ([API ref](https://stripe.com/docs/api/subscriptions/object#subscription_object-payment_settings-payment_method_types))
        * Add support for `USBankAccount` on `SubscriptionPaymentSettingsPaymentMethodOptions` and `SubscriptionPaymentSettingsPaymentMethodOptionsParams`
        * Add support for new values `paynow` and `us_bank_account` on enum `SubscriptionPaymentSettingsPaymentMethodTypes`
    * **Account capabilities** ([API ref](https://stripe.com/docs/api/accounts/object#account_object-capabilities))
      * Add support for `PayNowPayments` and `USBankAccountAchPayments` on `AccountCapabilities` and `AccountCapabilitiesParams`
* Add support for `FailureBalanceTransaction` on `Charge`
* Add support for `TestClock` on `SubscriptionListParams`
* Add support for `CaptureMethod` on `PaymentIntentConfirmPaymentMethodOptionsAfterpayClearpayParams`, `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentConfirmPaymentMethodOptionsKlarnaParams`, `PaymentIntentPaymentMethodOptionsAfterpayClearpayParams`, `PaymentIntentPaymentMethodOptionsAfterpayClearpay`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCard`, `PaymentIntentPaymentMethodOptionsKlarnaParams`, `PaymentIntentPaymentMethodOptionsKlarna`, and `PaymentIntentTypeSpecificPaymentMethodOptionsClient`
* Add additional support for verify microdeposits on Payment Intent and Setup Intent ([API ref](https://stripe.com/docs/api/payment_intents/verify_microdeposits))
    * Add support for `DescriptorCode` on `PaymentIntentVerifyMicrodepositsParams` and `SetupIntentVerifyMicrodepositsParams`
    * Add support for `MicrodepositType` on `PaymentIntentNextActionVerifyWithMicrodeposits` and `SetupIntentNextActionVerifyWithMicrodeposits`
* Add case for `ConnectCollectionTransfer` on `BalanceTransactionSource` `UnmarshalJSON` (fixes #1392)
* Add missing `PayoutFailureCode`s (fixes #1438)
