---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1451
is_stripe_api_change: true
released_in_version: 72.102.0
---

* Add support for new resources `FundingInstructions` and `Terminal.Configuration`
* Add support for `CreateFundingInstructions` method on resource `Customer`
* Add support for `CustomerBalance` on `ChargePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, and `PaymentMethod`
* Add support for `CashBalance` on `CustomerParams`
* Add support for `AmountDetails` on `PaymentIntent`
* Add support for `DisplayBankTransferInstructions` on `PaymentIntentNextAction`
* Add support for new value `customer_balance` on enum `PaymentMethodType`
* Add support for `ConfigurationOverrides` on `TerminalLocationParams` and `TerminalLocation`
