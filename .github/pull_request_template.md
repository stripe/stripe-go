<!--
# Title Suggestions

- If this branch is in-progress, start the title with [wip].
- Reference the Github Issue ID if it exists
- Provide a short descriptive title to help the reviewer understand the PR context

For example:
[wip] #121: Support SalesForce Contact to Stripe Customer in the data mapper

# General PR Guidelines

- Keep PRs focused and incremental
- Avoid grouping uncoupled changes into a single PR
- CircleCI changes must be isolated to an individual PR
-->

### Summary

<!-- What does the code do? What have you changed? -->

#### Salesforce Sandbox Login

URL:
Username:
Password:

### Motivation

<!-- Why are you making this change? Link to GitHub issue(s) if they exist  -->

### Pre-flight checklist

[Full code review guidelines are here.](https://docs.google.com/document/d/1m87sHnXA25hhKMmR-0C8dmBtTYLQXDkz4FeBH9mxcro/edit#heading=h.cz1xkpga7giy)

Before requested a review from Stripe engineering, ensure all of the below items are complete:

- [ ] Remove any code that is unneeded/commented out, including unnecessary log entries.
- [ ] Prefer top-level conditional checks vs nested `if`s
- [ ] Prefer loops over an array of keys over duplicating assignment code.
- [ ] Prefer small, well-named methods, to large methods.
- [ ] Ensure there are no secrets/keys in the code
- [ ] Ensure that any future changes documented in code use a `TODO` prefix
- [ ] All constants should be uppercase and, if possible, located in a centralized `constants.cls`
- [ ] All variable and test names are descriptive
- [ ] Avoid any variable or method names like `order2`. Always name the object what it represents.
- [ ] Ensure all `debugger` entries and other debug-only breakpoints are removed
- [ ] There are no magic numbers. Descriptive constants are used instead.
- [ ] Use global constants for magic strings instead of duplicating the string across the codebase
- [ ] Don't use instance variables unless they are required. Prefer `const` or locally scoped variables.
- [ ] Automated tests are all passing
- [ ] Code review from Appiphony
- [ ] Make sure this PR is installed on a sandbox account which Stripe has access to and indicate which sandbox account is used in the PR.
- [ ] Ensure there there are no merge conflicts
- [ ] Do not check in developer specific configuration files (eg. sdfx-config.json / DS_Store files)
- [ ] Ensure you're not accidentally overriding changes or deleting files in main.

### Test plan

<!-- How did you test this change? What were you unable to test?  -->

- [ ] All changes in PR are covered by tests
- [ ] Failures and edge cases tested
