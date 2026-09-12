---
title: Pull in V2 FinancialAccount changes for June release
pr_url: https://github.com/stripe/stripe-go/pull/2084
is_stripe_api_change: true
released_in_version: 82.4.0-beta.2
---

* Add support for `Close` and `New` methods on resource `V2MoneyManagementFinancialAccount`
* Add support for new value `storer` on enum `V2CoreAccount.AppliedConfigurations`
* Add support for `Storer` on `V2CoreAccountConfigurationParams` and `V2CoreAccountConfiguration`
* Add support for new values `financial_addresses.bank_accounts`, `holds_currencies.gbp`, `inbound_transfers.financial_accounts`, `outbound_payments.bank_accounts`, `outbound_payments.cards`, `outbound_payments.financial_accounts`, `outbound_transfers.bank_accounts`, and `outbound_transfers.financial_accounts` on enum `V2CoreAccountRequirementsEntriesImpactRestrictsCapabilities.Capability`
* Add support for new value `storer` on enum `V2CoreAccountRequirementsEntriesImpactRestrictsCapabilities.Configuration`
* Add support for `StatusDetails` on `V2MoneyManagementFinancialAccount`
* Add support for `Status` on `V2MoneyManagementFinancialAccountListParams`
* Add support for thin events `V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent` and `V2CoreAccountIncludingConfigurationStorerUpdatedEvent` with related object `V2CoreAccount`
* Add support for error types `AlreadyExistsError` and `NonZeroBalanceError`
