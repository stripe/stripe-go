---
title: ", [#1933](https://github.com/stripe/stripe-go/pull/1933) Remove rawrequests Post, Get, and Delete in favor of rawrequests.Client"
pr_link: https://github.com/stripe/stripe-go/pull/1929
released_in_version: 80.2.0
---

* The individual `rawrequests` functions for Post, Get, and Delete methods are removed in favor of the client model which allows local configuration of backend and api key, which enables more flexible calls to new/preview/unsupported APIs.
