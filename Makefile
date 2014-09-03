all: checkin

checkin:
	go test -v -run "TestCheckin*"
test:
	go test ./... -v
build:
	go build ./...
