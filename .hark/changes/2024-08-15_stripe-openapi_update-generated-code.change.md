---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1904
is_stripe_api_change: true
released_in_version: 79.8.0
---

* Add support for `AuthorizationCode` on `ChargePaymentMethodDetailsCard`
* Add support for `Wallet` on `ChargePaymentMethodDetailsCardPresent`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `ConfirmationTokenPaymentMethodPreviewCardPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, and `PaymentMethodCardPresent`
* Add support for `MandateOptions` on `PaymentIntentConfirmPaymentMethodOptionsBacsDebitParams`, `PaymentIntentPaymentMethodOptionsBacsDebitParams`, and `PaymentIntentPaymentMethodOptionsBacsDebit`
* Add support for `BACSDebit` on `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for `Chips` on `TreasuryOutboundPaymentTrackingDetailsUsDomesticWireParams`, `TreasuryOutboundPaymentTrackingDetailsUsDomesticWire`, `TreasuryOutboundTransferTrackingDetailsUsDomesticWireParams`, and `TreasuryOutboundTransferTrackingDetailsUsDomesticWire`
