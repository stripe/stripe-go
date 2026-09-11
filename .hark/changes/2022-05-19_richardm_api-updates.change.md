---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1463
is_stripe_api_change: true
released_in_version: 72.108.0
---

* Add support for new resources `Treasury.CreditReversal`, `Treasury.DebitReversal`, `Treasury.FinancialAccountFeatures`, `Treasury.FinancialAccount`, `Treasury.FlowDetails`, `Treasury.InboundTransfer`, `Treasury.OutboundPayment`, `Treasury.OutboundTransfer`, `Treasury.ReceivedCredit`, `Treasury.ReceivedDebit`, `Treasury.TransactionEntry`, and `Treasury.Transaction`
* Add support for `RetrievePaymentMethod` method on resource `Customer`
* Add support for `ListOwners` and `List` methods on resource `FinancialConnections.Account`
* Change type of `BillingPortalSessionReturnUrl` from `string` to `nullable(string)`
* Add support for `AUBECSDebit`, `AfterpayClearpay`, `BACSDebit`, `EPS`, `FPX`, `Giropay`, `Grabpay`, `Klarna`, `PayNow`, and `SepaDebit` on `CheckoutSessionPaymentMethodOptions`
* Add support for `Treasury` on `IssuingAuthorization`, `IssuingDisputeParams`, `IssuingDispute`, and `IssuingTransaction`
* Add support for `FinancialAccount` on `IssuingCardParams` and `IssuingCard`
* Add support for `ClientSecret` on `Order`
* Add support for `Networks` on `PaymentIntentConfirmPaymentMethodOptionsUsBankAccountParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountParams`, `PaymentMethodUsBankAccount`, `SetupIntentConfirmPaymentMethodOptionsUsBankAccountParams`, and `SetupIntentPaymentMethodOptionsUsBankAccountParams`
* Add support for `AttachToSelf` and `FlowDirections` on `SetupIntent`
* Add support for `SaveDefaultPaymentMethod` on `SubscriptionPaymentSettingsParams` and `SubscriptionPaymentSettings`
* Add support for `CZK` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
