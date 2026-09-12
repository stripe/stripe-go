---
title: Regenerate with decimal_string enabled for v2 APIs
pr_url: https://github.com/stripe/stripe-go/pull/2310
is_breaking: true
released_in_version: 85.0.0
---

- V2 API decimal fields changed type from `string` to `float64` with `json:",string"` and `form:",high_precision"` struct tags. Code that passes these fields as `string` will need to use `float64` instead. Affected fields:
  - **V2CoreAccountIdentityIndividualRelationship**: `PercentOwnership`
  - **V2CoreAccountPersonRelationship**: `PercentOwnership`
  - Params: `V2CoreAccountIdentityIndividualRelationshipParams`, `V2CoreAccountCreateIdentityIndividualRelationshipParams`, `V2CoreAccountUpdateIdentityIndividualRelationshipParams`, `V2CoreAccountTokenIdentityIndividualRelationshipParams`, `V2CoreAccountTokenCreateIdentityIndividualRelationshipParams`, `V2CoreAccountsPersonRelationshipParams`, `V2CoreAccountsPersonCreateRelationshipParams`, `V2CoreAccountsPersonUpdateRelationshipParams`, `V2CoreAccountsPersonTokenRelationshipParams`, `V2CoreAccountsPersonTokenCreateRelationshipParams`
