---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2236
is_stripe_api_change: true
released_in_version: 84.2.0-alpha.2
---

* Add support for new resource `TaxLocation`
* Add support for `Get`, `List`, and `New` methods on resource `TaxLocation`
* Add support for `Pause` method on resource `Subscription`
* Add support for `PerformanceLocation` on `CheckoutSessionLineItemPriceDataProductDataTaxDetailsParams`, `InvoiceAddLinesLinePriceDataProductDataTaxDetailsParams`, `InvoiceLineItemPriceDataProductDataTaxDetailsParams`, `InvoiceUpdateLinesLinePriceDataProductDataTaxDetailsParams`, `PaymentLinkLineItemPriceDataProductDataTaxDetailsParams`, `ProductTaxDetailsParams`, `TaxCalculationLineItemParams`, and `TaxCalculationLineItem`
* Add support for new value `performance` on enums `TaxCalculationLineItemTaxBreakdown.Sourcing`, `TaxCalculationShippingCostTaxBreakdown.Sourcing`, and `TaxTransactionShippingCostTaxBreakdown.Sourcing`
* Add support for new values `admissions_tax`, `attendance_tax`, `entertainment_tax`, `gross_receipts_tax`, `hospitality_tax`, `luxury_tax`, `resort_tax`, and `tourism_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationTaxBreakdownTaxRateDetails.TaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetails.TaxType`
* Change type of `DelegatedCheckoutRequestedSessionParams.Metadata` from `map(string: string)` to `emptyable(map(string: string))`
* Change type of `DelegatedCheckoutRequestedSessionParams.PaymentMethodData` from `payment_method_data` to `emptyable(payment_method_data)`
* Change type of `DelegatedCheckoutRequestedSessionParams.SharedMetadata` from `map(string: string)` to `emptyable(map(string: string))`
* Add support for `Subscription` on `InvoiceParentScheduleDetails` and `QuotePreviewInvoiceParentScheduleDetails`
* Change type of `PaymentIntentConfirmPaymentDetailsBenefitParams.FRMealVoucher`, `PaymentIntentPaymentDetailsBenefitParams.FRMealVoucher`, `SetupIntentConfirmSetupDetailsBenefitParams.FRMealVoucher`, and `SetupIntentSetupDetailsBenefitParams.FRMealVoucher` from `payment_details_benefit_fr_meal_voucher` to `emptyable(payment_details_benefit_fr_meal_voucher)`
* Add support for `TaxDetails` on `PlanProductParams` and `PriceProductDataParams`
* Add support for `ExternalReference` on `Plan` and `Price`
* Add support for new value `phase_start` on enums `QuoteSubscriptionData.PhaseEffectiveAt` and `QuoteSubscriptionDataOverrides.PhaseEffectiveAt`
* Remove support for value `line_start` from enums `QuoteSubscriptionData.PhaseEffectiveAt` and `QuoteSubscriptionDataOverrides.PhaseEffectiveAt`
* Add support for `AdmissionsTax`, `AttendanceTax`, `EntertainmentTax`, `GrossReceiptsTax`, `HospitalityTax`, `LuxuryTax`, `ResortTax`, and `TourismTax` on `TaxRegistrationCountryOptionsUs`
* Add support for new values `admissions_tax`, `attendance_tax`, `entertainment_tax`, `gross_receipts_tax`, `hospitality_tax`, `luxury_tax`, `resort_tax`, and `tourism_tax` on enum `TaxRegistrationCountryOptionsUs.Type`
* Add support for `Requirements` on `TaxCode`
