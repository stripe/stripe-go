### Pre-work

- [x] Refactor existing amendment work to make it easier to modify for the new structure
- [x] Duplicate price if duplicate price ID exists on the subscription
  - We want to use the same price as much as we can, but we should't optimize for this
- [ ] Fill in test templates to offer more feedback when refactoring

### Order Amendment:

- [x] Need to think on immutability gaurentees: we should prob look at data from SF when recreating old line items even if it changes.
- [x] New data structure to tie a termination item to all previous lines
  - Termination line will only contain a single reference
  - Recursively determine all references
  - Enforce LIFO queue on quantity termination
- [x] Ensure all metered billing items only have a quantity of 1
- [x] Conditionally use collasping system, only on termination
- Update & expand all tests
- Need to update lots of docs, although this can be punted (much harder to do later though...)

### Prorated Order Amendments

- Think on "backend" prorations which is not something we thought about
  - We can create two phases on the initial subscription
- Sync w/CF on custom price field and if it's needed
- Split unit price on line to get bill cycle amount
- Create invoice based on webhook
- Test new API endpoints from billing
