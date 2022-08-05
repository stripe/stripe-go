all: test bench vet lint check-api-clients check-gofmt

bench:
	go test -race -bench . -run "Benchmark" ./form

build:
	go build ./...

check-api-clients:
	go run scripts/check_api_clients/main.go

check-gofmt:
	scripts/check_gofmt.sh

lint:
	staticcheck

test:
	go run scripts/test_with_stripe_mock/main.go -race ./...

vet:
	go vet ./...

coverage:
	go run scripts/test_with_stripe_mock/main.go -covermode=count -coverprofile=combined.coverprofile ./...

clean:
	find . -name \*.coverprofile -delete

MAJOR_VERSION := $(shell echo $(VERSION) | sed 's/\..*//')
update-version:
	@echo "$(VERSION)" > VERSION
	@perl -pi -e 's|const clientversion = "[.\d\-\w]+"|const clientversion = "$(VERSION)"|' stripe.go
	@perl -pi -e 's|github.com/stripe/stripe-go/v\d+|github.com/stripe/stripe-go/v$(MAJOR_VERSION)|' README.md
	@perl -pi -e 's|github.com/stripe/stripe-go/v\d+|github.com/stripe/stripe-go/v$(MAJOR_VERSION)|' go.mod
	@find . -name '*.go' -exec perl -pi -e 's|github.com/stripe/stripe-go/v\d+|github.com/stripe/stripe-go/v$(MAJOR_VERSION)|' {} +

codegen-format:
	go fmt ./...

.PHONY: codegen-format update-version
