---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1472
is_stripe_api_change: true
released_in_version: 72.113.0
---

* Add support for `Affirm`, `Bancontact`, `Card`, `Ideal`, `P24`, and `Sofort` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for `AUBECSDebit`, `AfterpayClearpay`, `BACSDebit`, `EPS`, `FPX`, `Giropay`, `Grabpay`, `Klarna`, `PayNow`, and `SepaDebit` on `CheckoutSessionPaymentMethodOptionsParams`
* Add support for `SetupFutureUsage` on `CheckoutSessionPaymentMethodOptionsAcssDebitParams`, `CheckoutSessionPaymentMethodOptionsAcssDebit`, `CheckoutSessionPaymentMethodOptionsAfterpayClearpay`, `CheckoutSessionPaymentMethodOptionsAlipayParams`, `CheckoutSessionPaymentMethodOptionsAlipay`, `CheckoutSessionPaymentMethodOptionsAuBecsDebit`, `CheckoutSessionPaymentMethodOptionsBacsDebit`, `CheckoutSessionPaymentMethodOptionsBoletoParams`, `CheckoutSessionPaymentMethodOptionsBoleto`, `CheckoutSessionPaymentMethodOptionsEps`, `CheckoutSessionPaymentMethodOptionsFpx`, `CheckoutSessionPaymentMethodOptionsGiropay`, `CheckoutSessionPaymentMethodOptionsGrabpay`, `CheckoutSessionPaymentMethodOptionsKlarna`, `CheckoutSessionPaymentMethodOptionsKonbiniParams`, `CheckoutSessionPaymentMethodOptionsKonbini`, `CheckoutSessionPaymentMethodOptionsOxxoParams`, `CheckoutSessionPaymentMethodOptionsOxxo`, `CheckoutSessionPaymentMethodOptionsPaynow`, `CheckoutSessionPaymentMethodOptionsSepaDebit`, `CheckoutSessionPaymentMethodOptionsUsBankAccountParams`, `CheckoutSessionPaymentMethodOptionsUsBankAccount`, and `CheckoutSessionPaymentMethodOptionsWechatPayParams`
* Add support for `AttachToSelf` on `SetupAttempt`, `SetupIntentListParams`, and `SetupIntentParams`
* Add support for `FlowDirections` on `SetupAttempt` and `SetupIntentParams`
