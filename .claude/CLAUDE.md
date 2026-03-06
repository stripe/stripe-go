# stripe-go

## Testing

- Run all tests: `just test`
- Run tests for a single package: `just test ./invoice`
- Lint: `just lint` (uses go vet + staticcheck)
- Benchmarks: `just bench`

## Formatting

- Format: `just format` (uses gofmt + goimports)
- Format check: `just format-check`

## Key Locations

- Core client, HTTP handling, request logic, headers, retries: `stripe.go`
- Service-based client: `stripe_client.go`
- Tests: `stripe_test.go`

Note: this SDK uses a flat package layout. Most resource files (e.g. `customer.go`, `invoice.go`) are at the repo root alongside the core code.

## Generated Code

- Files containing `File generated from our OpenAPI spec` at the top are generated; do not edit. Similarly, any code block starting with `The beginning of the section generated from our OpenAPI spec` is generated and should not be edited directly.
  - If something in a generated file/range needs to be updated, add a summary of the change to your report but don't attempt to edit it directly.
- Resource files at the root (e.g. `customer.go`, `invoice.go`) are largely generated.
- `stripe.go` and `stripe_client.go` contain both generated and non-generated sections — look for section markers.

## Conventions

- Uses Go's standard `net/http` package
- Flat package structure (everything in package `stripe`)
- Major version in import path (e.g. `github.com/stripe/stripe-go/v82`) — `just format` handles updating these
- All code must run on all supported Go versions (full list in the test section of @.github/workflows/ci.yml)

### Comments

- Comments MUST only be used to:
  1. Document a function
  2. Explain the WHY of a piece of code
  3. Explain a particularly complicated piece of code
- Comments NEVER should be used to:
  1. Say what used to be there. That's no longer relevant!
  2. Explain the WHAT of a piece of code (unless it's very non-obvious)

It's ok not to put comments in a function if their addition wouldn't meaningfully clarify anything.
