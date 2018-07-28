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
	golint -set_exit_status ./applepaydomain
	golint -set_exit_status ./balance
	golint -set_exit_status ./bankaccount
	golint -set_exit_status ./bitcoinreceiver
	golint -set_exit_status ./bitcointransaction
	golint -set_exit_status ./card
	golint -set_exit_status ./charge
	golint -set_exit_status ./countryspec
	golint -set_exit_status ./coupon
	golint -set_exit_status ./customer
	golint -set_exit_status ./discount
	golint -set_exit_status ./dispute
	golint -set_exit_status ./ephemeralkey
	golint -set_exit_status ./event
	golint -set_exit_status ./exchangerate
	golint -set_exit_status ./fee
	golint -set_exit_status ./feerefund
	golint -set_exit_status ./fileupload
	golint -set_exit_status ./invoice
	golint -set_exit_status ./invoiceitem
	golint -set_exit_status ./issuerfraudrecord
	golint -set_exit_status ./issuing/authorization
	golint -set_exit_status ./issuing/card
	golint -set_exit_status ./issuing/cardholder
	golint -set_exit_status ./issuing/dispute
	golint -set_exit_status ./issuing/transaction
	golint -set_exit_status ./loginlink
	golint -set_exit_status ./order
	golint -set_exit_status ./orderreturn
	golint -set_exit_status ./paymentintent
	golint -set_exit_status ./paymentsource
	golint -set_exit_status ./payout
	golint -set_exit_status ./plan
	golint -set_exit_status ./product
	golint -set_exit_status ./recipient
	golint -set_exit_status ./recipienttransfer
	golint -set_exit_status ./refund
	golint -set_exit_status ./reversal
	golint -set_exit_status ./sigma/scheduledqueryrun
	golint -set_exit_status ./sku
	golint -set_exit_status ./source
	golint -set_exit_status ./sourcetransaction
	golint -set_exit_status ./sub
	golint -set_exit_status ./subitem
	golint -set_exit_status ./threedsecure
	golint -set_exit_status ./token
	golint -set_exit_status ./topup
	golint -set_exit_status ./transfer
	golint -set_exit_status ./usagerecord

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
