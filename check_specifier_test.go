package main

import (
	"fmt"
	"testing"
)

func TestCheckSpecifier(t *testing.T) {
	fmt.Println("\n=== Test CheckSpecifier ===")
	/* ""-> false, "Missing type." */
	result, err := CheckSpecifier(SplitSpecifier(""))
	if result == true {
		fmt.Println("Result should be false (\"\")")
		t.Error("Result should be false (\"\")")
	} else if result == false && err != "Missing type." {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test \"\" OK.")
	}
	/* []32*32 -> false, "Missing type." */
	l := SplitSpecifier("[]32*32")
	PrintList(l)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false []32*32")
		t.Error("Result should be false []32*32")
	} else if result == false && err != "Missing type." {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test []i OK.")
	}
	/* ][i */
	l = SplitSpecifier("][i")
	PrintList(l)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for ][i test")
		t.Error("Result should be false for ][i test")
	} else if result == false && err != "Invalid slice specifier: ][" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test []32*32 OK.")
	}
	/* [32c32 */
	l = SplitSpecifier("[32c32")
	PrintList(l)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for [32c32 test")
		t.Error("Result should be false for [32c32 test")
	} else if result == false && err != "Invalid slice specifier: [" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test [32c32 OK.")
	}
	/* ]]32c32 */
	l = SplitSpecifier("]]32c32")
	PrintList(l)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for ]]32c32 test")
		t.Error("Result should be false for ]]32c32 test")
	} else if result == false && err != "Invalid slice specifier: ]]" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test ]]32c32 OK.")
	}
	/* [[[[32c32 */
	l = SplitSpecifier("[[[[32c32")
	PrintList(l)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for [[[[32c32 test")
		t.Error("Result should be false for [[[[32c32 test")
	} else if result == false && err != "Invalid slice specifier: [[[[" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test [[[[32c32 OK.")
	}
	/* []i */
	l = SplitSpecifier("[]i")
	PrintList(l)
	result, err = CheckSpecifier(l)
	if result != true {
		fmt.Println("Result should be true for []i test. Error:", err)
		t.Error("Result should be true for []i test. Error: " + err)
	} else {
		fmt.Println("Test []i OK.")
	}
}
