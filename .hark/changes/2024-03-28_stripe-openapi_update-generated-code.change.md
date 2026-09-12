---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1830
is_stripe_api_change: true
released_in_version: 76.23.0
---

* Add support for new resources `Billing.MeterEventAdjustment`, `Billing.MeterEvent`, and `Billing.Meter`
* Add support for `Deactivate`, `Get`, `List`, `New`, `Reactivate`, and `Update` methods on resource `Meter`
* Add support for `New` method on resources `MeterEventAdjustment` and `MeterEvent`
* Add support for `AmazonPayPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for new value `verification_failed_representative_authority` on enums `AccountFutureRequirementsErrorsCode`, `AccountRequirementsErrorsCode`, `BankAccountFutureRequirementsErrorsCode`, and `BankAccountRequirementsErrorsCode`
* Add support for `DestinationOnBehalfOfChargeManagement` on `AccountSessionComponentsPaymentDetailsFeaturesParams`, `AccountSessionComponentsPaymentDetailsFeatures`, `AccountSessionComponentsPaymentsFeaturesParams`, and `AccountSessionComponentsPaymentsFeatures`
* Add support for `Mandate` on `ChargePaymentMethodDetailsUsBankAccount`, `TreasuryInboundTransferOriginPaymentMethodDetailsUsBankAccount`, `TreasuryOutboundPaymentDestinationPaymentMethodDetailsUsBankAccount`, and `TreasuryOutboundTransferDestinationPaymentMethodDetailsUsBankAccount`
* Add support for `SecondLine` on `IssuingCardParams`
* Add support for `Meter` on `PlanParams`, `Plan`, `PriceListRecurringParams`, `PriceRecurringParams`, and `PriceRecurring`
