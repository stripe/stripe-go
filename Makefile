all: test bench vet lint check-gofmt

bench:
	go test -race -bench . -run "Benchmark" ./form

build:
	go build ./...

check-gofmt:
	scripts/check_gofmt.sh

lint:
# Golint won't run on some Go versions. See `.travis.yml`.
ifndef SKIP_GOLINT
	golint -set_exit_status ./...
else
	# No Golint. Skipping Linting.
endif

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
