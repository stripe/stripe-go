---
title: Make unmarshalJSONVerbose unexported and context-aware
pr_url: https://github.com/stripe/stripe-go/pull/2301
semver_level: major
released_in_version: 85.0.0
---

- Unexported `BackendImplementation.UnmarshalJSONVerbose` (now `unmarshalJSONVerbose`) and added a `context.Context` parameter for proper context propagation in error logging.
