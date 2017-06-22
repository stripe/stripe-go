all: test bench vet check-gofmt

bench:
	go test -bench . -run "Benchmark" ./form

build:
	go build ./...

check-gofmt:
	scripts/check_gofmt.sh

test:
	go test ./...

vet:
	go vet ./...
