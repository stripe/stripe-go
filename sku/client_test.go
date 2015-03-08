package sku

import (
	"fmt"
	"testing"

	stripe "github.com/stripe-internal/stripe-go"
)

func init() {
	stripe.Key = "McI72P18xiznF25AbdpLcaLurWfUMVEy"
}

func TestSKU(t *testing.T) {
	params := &stripe.SKUParams{}
	params.Expand("product")

	pr, err := Get("sku_5inYqW3PUcurZi", params)
	fmt.Printf("%+v %+v", pr, err)
}
