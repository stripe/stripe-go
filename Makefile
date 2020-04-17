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
	golint -set_exit_status ./...

test:
	go run scripts/test_with_stripe_mock/main.go -race ./...

vet:
	go vet ./...

coverage:
	go run scripts/test_with_stripe_mock/main.go -covermode=count -coverprofile=combined.coverprofile ./...

clean:
	find . -name \*.coverprofile -delete
