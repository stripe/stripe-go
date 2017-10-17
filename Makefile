all: test bench vet check-gofmt

bench:
	go test -race -bench . -run "Benchmark" ./form

build:
	go build ./...

check-gofmt:
	scripts/check_gofmt.sh

test:
	go test -race ./...

vet:
	go vet ./...
