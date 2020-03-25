package argparse

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
	/* []ii[]f32 -> error */
	format = "[]ii[]f32"
	result, _, err = CheckFormat(format)
	if result == true {
		fmt.Println("Result for", format, " test should be false.")
		t.Error("Result for " + format + " test should be false. Error: " + err)
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
	format = "*riis"
	result, _, err = CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []10*ui32 */
	format = "[]10*ui32"
	result, _, err = CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []-19ui -> error. */
	format = "[]-19ui"
	result, _, err = CheckFormat(format)
	if result == true {
		fmt.Println("Result for", format, " test should be false.")
		t.Error("Result for " + format + " test should be false. Error: " + err)
	} else if result == false && err != "Slice size is negative: -19" {
		fmt.Println(err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []i[]f32 -> error. */
	format = "[]i[]f32"
	result, _, err = CheckFormat(format)
	if result == true {
		fmt.Println("Result for", format, " test should be false.")
		t.Error("Result for " + format + " test should be false. Error: " + err)
	} else if result == false && err != "Incompatible types: []i and []f32" {
		fmt.Println(err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []ui32[]3c128 -> ok. */
	format = "[]ui32[]3c128"
	result, _, err = CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []b[]s */
	format = "[]b[]s"
	result, _, err = CheckFormat(format)
	if result == true {
		fmt.Println("Result for", format, " test should be false.")
		t.Error("Result for " + format + " test should be false. Error: " + err)
	} else if result == false && err != "Incompatible types: []b and []s" {
		fmt.Println(err)
		t.Error(err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []ui8[]b -> ok. */
	format = "[]ui8[]b"
	result, _, err = CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be false.")
		t.Error("Result for " + format + " test should be false. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []ii[]i -> error. */
	format = "[]ii[]i"
	result, _, err = CheckFormat(format)
	if result == true {
		fmt.Println("Result for", format, " test should be false.")
		t.Error("Result for " + format + " test should be false. Error: " + err)
	} else if result == false && err != "Incompatible types: []i and []i" {
		fmt.Println(err)
		t.Error(err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []is[]i -> ok. */
	format = "[]is[]i"
	result, _, err = CheckFormat(format)
	if result == false {
		fmt.Println("Result for", format, " test should be true.")
		t.Error("Result for " + format + " test should be true. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* []i[]f[]c64 -> eror. */
	format = "[]i[]f32[]c64"
	result, _, err = CheckFormat(format)
	if result == true {
		fmt.Println("Result for", format, " test should be false.")
		t.Error("Result for " + format + " test should be false. Error: " + err)
	} else {
		fmt.Println("Test " + format + " OK.")
	}
	/* test Specifier::toString() */
	s := Specifier{true, true, 5, 32, "ui"}
	if s.toString() != "[]5*ui32" {
		fmt.Println("Invalid Specifier.toString()")
		t.Error("Invalid Specifier.toString()")
	} else {
		fmt.Println("Test Specifier::toString() OK.")
	}
}
