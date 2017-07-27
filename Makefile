# This is a testmode key for this test account:
#
#     ID:    cuD9Rwx8pgmRZRpVe02lsuR9cwp2Bzf7
#     Email: test+bindings@stripe.com
#
STRIPE_KEY=tGN0bIwXnHdwOa85VABjPdSn8nWY7G7I

all: checkin vet check-gofmt

check-gofmt:
	scripts/check_gofmt.sh

checkin:
	STRIPE_KEY=$(STRIPE_KEY) go test -run "TestCheckin*" ./...

test:
	STRIPE_KEY=$(STRIPE_KEY) go test ./... -p=1

build:
	go build ./...

vet:
	go vet ./...
