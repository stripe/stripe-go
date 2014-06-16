package stripe

import "os"

var key string

func init() {
	key = os.Getenv("STRIPE_KEY")

	if len(key) == 0 {
		panic("STRIPE_KEY environment variable is not set, but is needed to run tests!\n")
	}
}
