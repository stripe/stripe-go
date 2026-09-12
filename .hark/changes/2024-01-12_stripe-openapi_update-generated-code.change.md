---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1795
is_stripe_api_change: true
released_in_version: 76.14.0-beta.1
---

* Add support for `Amount` on `ChargeCapturePaymentDetailsFlightSegmentsParams`, `ChargePaymentDetailsFlightSegmentsParams`, `PaymentIntentCapturePaymentDetailsFlightSegmentsParams`, `PaymentIntentConfirmPaymentDetailsFlightSegmentsParams`, and `PaymentIntentPaymentDetailsFlightSegmentsParams`
* Add support for `NumberOfRooms` and `RoomClass` on `ChargeCapturePaymentDetailsLodgingParams`, `ChargePaymentDetailsLodgingParams`, `PaymentIntentCapturePaymentDetailsLodgingParams`, `PaymentIntentConfirmPaymentDetailsLodgingParams`, and `PaymentIntentPaymentDetailsLodgingParams`
* Add support for `BuyButton` on `CustomerSessionComponentsParams` and `CustomerSessionComponents`
* Add support for new values `high_risk_industry`, `insufficient_margin_ratio`, `insufficient_operating_profit`, `insufficient_reserves`, `insufficient_time_in_network`, `lacking_cash_account`, and `poor_payment_history_with_platform` on enum `IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasons`
* Add support for new values `high_risk_industry`, `insufficient_margin_ratio`, `insufficient_operating_profit`, `insufficient_reserves`, `insufficient_time_in_network`, and `lacking_cash_account` on enums `IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasons` and `IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasons`
