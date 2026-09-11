---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1765
is_stripe_api_change: true
released_in_version: 76.6.0-beta.1
---

* Add support for `IssuingCard` and `IssuingCardsList` on `AccountSessionComponentsParams`
* Add support for `EventDetails` and `Subscription` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentPaymentDetailsParams`, and `PaymentIntentPaymentDetails`
* Add support for `Affiliate` and `Delivery` on `ChargeCapturePaymentDetailsCarRentalParams`, `ChargeCapturePaymentDetailsFlightParams`, `ChargeCapturePaymentDetailsLodgingParams`, `ChargePaymentDetailsCarRentalParams`, `ChargePaymentDetailsFlightParams`, `ChargePaymentDetailsLodgingParams`, `PaymentIntentCapturePaymentDetailsCarRentalParams`, `PaymentIntentCapturePaymentDetailsFlightParams`, `PaymentIntentCapturePaymentDetailsLodgingParams`, `PaymentIntentConfirmPaymentDetailsCarRentalParams`, `PaymentIntentConfirmPaymentDetailsFlightParams`, `PaymentIntentConfirmPaymentDetailsLodgingParams`, `PaymentIntentPaymentDetailsCarRentalParams`, `PaymentIntentPaymentDetailsCarRental`, `PaymentIntentPaymentDetailsFlightParams`, and `PaymentIntentPaymentDetailsLodgingParams`
* Add support for `Drivers` on `ChargeCapturePaymentDetailsCarRentalParams`, `ChargePaymentDetailsCarRentalParams`, `PaymentIntentCapturePaymentDetailsCarRentalParams`, `PaymentIntentConfirmPaymentDetailsCarRentalParams`, `PaymentIntentPaymentDetailsCarRentalParams`, and `PaymentIntentPaymentDetailsCarRental`
* Add support for `Passengers` on `ChargeCapturePaymentDetailsFlightParams`, `ChargeCapturePaymentDetailsLodgingParams`, `ChargePaymentDetailsFlightParams`, `ChargePaymentDetailsLodgingParams`, `PaymentIntentCapturePaymentDetailsFlightParams`, `PaymentIntentCapturePaymentDetailsLodgingParams`, `PaymentIntentConfirmPaymentDetailsFlightParams`, `PaymentIntentConfirmPaymentDetailsLodgingParams`, `PaymentIntentPaymentDetailsFlightParams`, and `PaymentIntentPaymentDetailsLodgingParams`
* Add support for `Created` on `CustomerSession`
