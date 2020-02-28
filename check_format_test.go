package main

import (
	"fmt"
	"testing"
)

func TestCheckFormat(t *testing.T) {
	fmt.Println("\n=== CheckFormat ===")
	/* iif32 */
	format := "iif32"
	result, err := CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
}
