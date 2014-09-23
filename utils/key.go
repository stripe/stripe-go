// Package utils provides internal utilities
package utils

import "os"

func GetTestKey() string {
	key := os.Getenv("STRIPE_KEY")

	if len(key) == 0 {
		panic("STRIPE_KEY environment variable is not set, but is needed to run tests!\n")
	}

	return key
}
