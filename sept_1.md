### Pre-work

- [x] Refactor existing amendment work to make it easier to modify for the new structure
- [x] Duplicate price if duplicate price ID exists on the subscription
  - We want to use the same price as much as we can, but we should't optimize for this

### Order Amendment:

- [x] Need to think on immutability gaurentees: we should prob look at data from SF when recreating old line items even if it changes.
- [x] New data structure to tie a termination item to all previous lines
  - Termination line will only contain a single reference
  - Recursively determine all references
  - Enforce LIFO queue on quantity termination
- [x] Ensure all metered billing items only have a quantity of 1
- [x] Conditionally use collasping system, only on termination
- [ ] Need to update lots of docs, although this can be punted (much harder to do later though...)

### Prorated Order Amendments

- [x] extract out one-time invoicing into separate method
- [ ] Refactor to pull out order amendment helpers into separate module
- [x] Helper to detect if order is prorated
  - If there are all new line items on the order, then it doesn't matter
  - If the order start date does NOT align with the next billing date, then it's definitely prorated
  - If the order start date does align, does the subscription term align with the billing cycle
- [x] Will CF always have co-termed amendments?
  - Yes, sounds like this may be a CPQ thing
  - [ ] enforce this in code
- [x] Split unit price on line to get bill cycle amount
- [ ] Use `end_date` vs iterations in initial phase and rip out associated logic
- [ ] Improve custom price field detection to be more reliable
- [ ] Think on "backend" prorations which is not something we thought about
  - We can create two phases on the initial subscription
  - Validate bug we encountered when prototyping
- [x] Sync w/CF on custom price field and if it's needed
  - They need this and are set on using it
- [x] Accept webhooks
- [x] Test calculating normal billing cycle price when order line is prorated
- [ ] Create invoice based on webhook
- [ ] Test new API endpoints from billing
- [ ] Test clock based tests to validate things are working properly
