package product

import (
	"fmt"
	"testing"

	stripe "github.com/stripe-internal/stripe-go"
)

func init() {
	stripe.Key = "McI72P18xiznF25AbdpLcaLurWfUMVEy"
}

func TestProduct(t *testing.T) {
	pr, err := Get("pr_5inYraUi3syZA6")
	fmt.Printf("%+v %+v", pr, err)
}
