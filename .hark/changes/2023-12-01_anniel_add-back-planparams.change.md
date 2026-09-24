---
title: Add back PlanParams.ProductID
pr_url: https://github.com/stripe/stripe-go/pull/1777
released_in_version: 76.8.0
---

* Add back `PlanParams.ProductID`, which was mistakenly removed starting in v73.0.0. `ProductID` allows creation of a plan for an existing product by serializing `product` as a string .
