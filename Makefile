all: test

test:
	go test ./... -v
build:
	go build ./...
