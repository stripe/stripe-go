all: checkin

checkin:
	go test -run "TestCheckin*" ./client
test:
	go test ./... -p=1
build:
	go build ./...
