all: test bench vet lint check-gofmt

bench:
	go test -race -bench . -run "Benchmark" ./form

build:
	go build ./...

check-gofmt:
	scripts/check_gofmt.sh

# We're trying to get linting activated, but there are so many errors that
# we're doing so incrementally instead of pushing it through as part of one
# giant patch.
#
# As new packages get cleaned up, add them here so that we don't regress. When
# all packages are here, switch to a `./...` instead of linting every package
# individually.
lint:
	golint -set_exit_status ./account
	golint -set_exit_status ./charge

test:
	go test -race ./...

vet:
	go vet ./...

coverage:
	# go currently cannot create coverage profiles when testing multiple packages, so we test each package
	# independently. This issue should be fixed in Go 1.10 (https://github.com/golang/go/issues/6909).
	go list ./... | xargs -n1 -I {} -P 4 go test -covermode=count -coverprofile=../../../{}/profile.coverprofile {}

clean:
	find . -name \*.coverprofile -delete
