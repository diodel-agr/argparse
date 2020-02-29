package main

import (
	"fmt"
	"testing"
)

func TestCheckFormat(t *testing.T) {
	fmt.Println("\n=== CheckFormat ===")
	/* iif32 */
	format := "iif32"
	result, _, err := CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []ii[]f32 */
	fmt.Println()
	format = "[]ii[]f32"
	result, _, err = CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []10f32is */
	format = "[]10f32is"
	result, _, err = CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* *riis */
	fmt.Println()
	format = "*riis"
	result, _, err = CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []10*ui32 */
	fmt.Println()
	format = "[]10*ui32"
	result, _, err = CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []-19ui -> error. */
	fmt.Println()
	format = "[]-19ui"
	result, _, err = CheckFormat(format)
	if result == true {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else if result == false && err != "Slice size is negative: -19" {
		fmt.Println(err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
}
