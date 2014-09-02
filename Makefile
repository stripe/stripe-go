all: build

test:
	go test ./... -v
build:
	go build ./...
