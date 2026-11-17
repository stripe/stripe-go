---
title: Send an `Idempotency-Key` on write requests made without params
semver_level: patch
jira_tickets_closed:
- RUN_DEVSDK-3666
---

- `RawRequest` and the deprecated resource packages (e.g. `account.New(nil)`) now generate an `Idempotency-Key` for `POST`, `PUT`, `PATCH`, and `DELETE` requests when params is `nil`, and reuse it on retries. Previously these requests were sent and retried without a key.
