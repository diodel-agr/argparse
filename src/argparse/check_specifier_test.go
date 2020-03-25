package argparse

import (
	"fmt"
	"testing"
)

func TestCheckSpecifier(t *testing.T) {
	fmt.Println("\n=== CheckSpecifier ===")
	// ==== CheckSlice test. ==== //
	/* ][i */
	l := SplitSpecifier("][i")
	result, err := CheckSpecifier(l)
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
	result, err = CheckSpecifier(l)
	if result != true {
		fmt.Println("Result should be true for []i test. Error:", err)
		t.Error("Result should be true for []i test. Error: " + err)
	} else {
		fmt.Println("Test []i OK.")
	}
	// ==== CheckSliceSize test. ==== //
	/* []-10 */
	l = SplitSpecifier("[]-10")
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for []-10 test")
		t.Error("Result should be false for []-10 test")
	} else if result == false && err != "Slice size is negative: -10" {
		fmt.Println("Error should be: Slice size is negative: -10. Error:", err)
		t.Errorf("Error should be: Slice size is negative: -10. Error:" + err)
	} else {
		fmt.Println("Test []-10 OK.")
	}
	/* []-i */
	l = SplitSpecifier("[]-i")
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for []-10 test")
		t.Error("Result should be false for []-10 test")
	} else if result == false && err != "Invalid slice size: -" {
		fmt.Println("Error should be: Invalid slice size: -. Error:", err)
		t.Errorf("Error should be: Invalid slice size: -. Error:" + err)
	} else {
		fmt.Println("Test []-i OK.")
	}
	/* []00s */
	s := "[]00s"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for ", s, "test")
		t.Error("Result should be false for " + s + "test")
	} else if result == false && err != "Slice size is zero: 00" {
		fmt.Println("Error should be: Slice size is zero: 00. Error:", err)
		t.Errorf("Error should be: Slice size is zero: 00. Error:" + err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	/* []1.5i */
	s = "[]1.5s"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for ", s, "test")
		t.Error("Result should be false for " + s + "test")
	} else if result == false && err != "Invalid slice size: 1.5" {
		fmt.Println("Error should be: Invalid slice size: 1.5. Error:", err)
		t.Errorf("Error should be: Invalid slice size: 1.5. Error:" + err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	/* []12c64 */
	s = "[]12c64"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result != true {
		fmt.Println("Result should be true for ", s, "test. Error:", err)
		t.Error("Result should be true for " + s + "test. Error: " + err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	// ==== checkPointerType test. ==== //
	/* *f32 */
	s = "*f32"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result != true {
		fmt.Println("Result should be true for ", s, "test. Error:", err)
		t.Error("Result should be true for " + s + "test. Error: " + err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	/* []*s */
	s = "[]*s"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result != true {
		fmt.Println("Result should be true for ", s, "test. Error:", err)
		t.Error("Result should be true for " + s + "test. Error: " + err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	// ==== checkType test. ==== //
	/* ""-> false, "Missing type." */
	result, err = CheckSpecifier(SplitSpecifier(""))
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
	l = SplitSpecifier("[]32*32")
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
	/* *d */
	l = SplitSpecifier("*d")
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false *d")
		t.Error("Result should be false *d")
	} else if result == false && err != "Unknown type: d" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test *d OK.")
	}
	// ==== checkBitSize test. ==== //
	/* f- */
	s = "f-"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for", s, "test.")
		t.Error("Result should be false for " + s + "test.")
	} else if result == false && err != "Invalid bit size: -" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	/* ui-32 */
	s = "ui-32"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for", s, "test.")
		t.Error("Result should be false for " + s + "test.")
	} else if result == false && err != "Bit size is negative: -32" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	/* []c0 */
	s = "[]c0"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for", s, "test.")
		t.Error("Result should be false for " + s + "test.")
	} else if result == false && err != "Bit size is zero: 0" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	/* f1.1 */
	s = "f1.1"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for", s, "test.")
		t.Error("Result should be false for " + s + "test.")
	} else if result == false && err != "Invalid bit size: 1.1" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	/* []c44 */
	s = "[]c44"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for", s, "test.")
		t.Error("Result should be false for " + s + "test.")
	} else if result == false && err != "Invalit bit size: 44. Valid values are: 8, 16, 32, 64, 128." {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	// ==== checkTypeBitSize test ==== //
	/* r32 */
	s = "[]r32"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for", s, "test.")
		t.Error("Result should be false for " + s + "test.")
	} else if result == false && err != "For the rune, byte, bool or string type, you cannot specify any bit size" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	/* c8 */
	s = "c8"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for", s, "test.")
		t.Error("Result should be false for " + s + "test.")
	} else if result == false && err != "For the complex type, the valid bit size are 64 or 128" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
	/* f16 */
	s = "f16"
	l = SplitSpecifier(s)
	result, err = CheckSpecifier(l)
	if result == true {
		fmt.Println("Result should be false for", s, "test.")
		t.Error("Result should be false for " + s + "test.")
	} else if result == false && err != "For the float type, the valid bit sizes are 32 or 64" {
		fmt.Println(err)
		t.Errorf(err)
	} else {
		fmt.Println("Test", s, "OK.")
	}
}
