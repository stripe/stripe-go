---
title: Fix `BalanceTransactionSource` `UnmarshalJSON` for when `BalanceTransactionSource.Type == "transfer_reversal"` (previously, we were checking if `Type == "reversal"`, which was always false)
section: ⚠️ Changed
released_in_version: 73.0.0
---
