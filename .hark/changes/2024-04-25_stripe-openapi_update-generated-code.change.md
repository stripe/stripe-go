---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1852
is_stripe_api_change: true
released_in_version: 78.4.0
---

* Add support for `SetupFutureUsage` on `CheckoutSessionPaymentMethodOptionsAmazonPay`, `CheckoutSessionPaymentMethodOptionsRevolutPay`, `PaymentIntentPaymentMethodOptionsAmazonPay`, and `PaymentIntentPaymentMethodOptionsRevolutPay`
* Change type of `EntitlementsActiveEntitlementFeature` from `string` to `*EntitlementsFeature`
* Remove support for inadvertently released identity verification features `Email` and `Phone` on `IdentityVerificationSessionOptionsParams`
* Add support for new values `amazon_pay` and `revolut_pay` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Add support for `AmazonPay` and `RevolutPay` on `MandatePaymentMethodDetails` and `SetupAttemptPaymentMethodDetails`
* Add support for `EndingBefore`, `Limit`, and `StartingAfter` on `PaymentMethodConfigurationListParams`
* Add support for `Mobilepay` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
