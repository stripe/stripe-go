---
title: Make unmarshalJSONVerbose unexported and context-aware
pr_link: https://github.com/stripe/stripe-go/pull/2301
is_breaking: true
released_in_version: 85.0.0
---

- Unexported `BackendImplementation.UnmarshalJSONVerbose` (now `unmarshalJSONVerbose`) and added a `context.Context` parameter for proper context propagation in error logging.
