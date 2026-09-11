---
title: Build SDK w/ V2 OpenAPI spec
pr_link: https://github.com/stripe/stripe-go/pull/2114
is_breaking: true
released_in_version: 83.0.0
---

- ⚠️ The delete methods for v2 APIs (the ones in the `V2` prefix) now return a `V2DeletedObject` which has the id of the object that has been deleted and a string representing the type of the object that has been deleted.
- ⚠️ Nullable properties on objects returned by v2 APIs now have the `omitempty` annotation
