---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2489
semver_level: major
is_stripe_api_change: true
---

* ⚠️ Remove support for resource `V2BillingCadence`
* Add support for `New` and `Refill` test helper methods on resource `CapitalFinancingOffer`
* ⚠️ Remove support for `Cancel`, `Get`, `List`, `New`, and `Update` methods on resource `V2BillingCadence`
* Add support for `ExportTaxTransactions`, `ProductTaxCodeSelector`, and `TaxThresholdMonitoring` on `AccountSessionComponents`
* Add support for `Custom` on `CouponParams` and `Coupon`
* Add support for `Type` on `CouponParams` and `CustomerCustomerBalanceTransactionParams`
* Add support for new value `custom` on enum `Coupon.Type`
* Add support for `AppliedToInvoice` on `CustomerCustomerBalanceTransactionParams`
* Add support for `CustomerBalance` on `Invoice` and `QuotePreviewInvoice`
* Add support for `Connect`, `DataExport`, `EventDestination`, `Radar`, and `UserAccess` on `V2IamActivityLogDetails`
* Add support for new values `app`, `connect`, `data_export`, `event_destination`, `radar`, and `user_access` on enum `V2IamActivityLogDetails.Type`
* Add support for new values `apps.install`, `v2.core.event_destination`, and `webhook_endpoint` on enum `V2IamActivityLogRelatedObject.Type`
* Add support for new values `app_installed`, `app_uninstalled`, `application_authorized`, `application_created`, `application_deauthorized`, `connect_express_branding_updated`, `connect_oauth_disabled`, `connect_oauth_enabled`, `connect_settings_updated`, `connect_site_link_updated`, `data_export_requested`, `event_destination_added`, `event_destination_removed`, `event_destination_updated`, `payout_schedule_updated`, `radar_adaptive_3ds_disabled`, `radar_dynamic_risk_thresholds_disabled`, `radar_early_fraud_warning_disabled`, `radar_for_setup_intents_disabled`, `radar_fraud_dispute_protection_disabled`, `radar_rule_added`, `radar_rule_disabled`, `radar_rule_enabled`, `radar_rule_removed`, `radar_rule_updated`, and `user_access_started` on enum `V2IamActivityLog.Type`
* Add support for `ApplePay` on `V2MoneyManagementOutboundSetupIntentPayoutMethodDataParams` and `V2MoneyManagementPayoutMethod`
* Add support for new value `apple_pay` on enum `V2MoneyManagementPayoutMethod.Type`
* Add support for `Network` and `OriginatingBankAccount` on `V2MoneyManagementReceivedDebitBankTransfer`
* Add support for `Created` on `V2MoneyManagementReceivedDebitListParams`
* Add support for `LogsFrom` on `V2IamActivityLogListParams`
* Add support for `Severity` on `EventsV2CoreHealthEventGenerationFailureResolvedEvent`
* Add support for error code `inactive_api_operation_confirmation_required` on `QuotePreviewInvoiceLastFinalizationError`
