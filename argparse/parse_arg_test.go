package argparse

import (
	"fmt"
	"testing"
)

func testParseArg01(format string, argv []string) string {
	_, slist, err0 := CheckFormat(format)
	if err0 != "" {
		return err0
	}
	l, err1 := parseArguments(*slist, argv)
	if l != nil {
		fmt.Print("Parsing result: ")
		PrintList(l)
	}
	return err1
}

func TestParseArguments(t *testing.T) {
	fmt.Println("\n=== parseArguments ===")
	// int
	// i 10 -> OK
	format := "i"
	argv := []string{"10"}
	err := testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *i 10 -> OK
	format = "*i"
	argv = []string{"10"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// i 12.6 -> err0
	fmt.Println()
	format = "i"
	argv = []string{"12.6"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// i8 100 -> OK
	fmt.Println()
	format = "i8"
	argv = []string{"100"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *i8 100 -> OK
	fmt.Println()
	format = "*i8"
	argv = []string{"100"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// i8 123.4 -> err0
	fmt.Println()
	format = "i8"
	argv = []string{"123.4"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// i16 1024 -> OK
	fmt.Println()
	format = "i16"
	argv = []string{"1024"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *i16 1024 -> OK
	fmt.Println()
	format = "*i16"
	argv = []string{"1024"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// i16 0.10 -> err
	fmt.Println()
	format = "i16"
	argv = []string{"0.10"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// i32 17563 -> OK
	fmt.Println()
	format = "i32"
	argv = []string{"17563"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *i32 17563 -> OK
	fmt.Println()
	format = "*i32"
	argv = []string{"17563"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// i32 100.432 -> err
	fmt.Println()
	format = "i32"
	argv = []string{"100.432"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// i64 123456789 -> OK
	fmt.Println()
	format = "i64"
	argv = []string{"123456789"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *i64 123456789 -> OK
	fmt.Println()
	format = "*i64"
	argv = []string{"123456789"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// i64 10.23456 -> err
	fmt.Println()
	format = "i64"
	argv = []string{"10.23456"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// uint
	// ui 10 -> OK
	format = "ui"
	argv = []string{"10"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *ui 10 -> OK
	format = "*ui"
	argv = []string{"10"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// ui 12.6 -> err0
	fmt.Println()
	format = "ui"
	argv = []string{"12.6"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// ui8 100 -> OK
	fmt.Println()
	format = "ui8"
	argv = []string{"100"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *ui8 100 -> OK
	fmt.Println()
	format = "*ui8"
	argv = []string{"100"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// ui8 123.4 -> err0
	fmt.Println()
	format = "ui8"
	argv = []string{"123.4"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// ui16 1024 -> OK
	fmt.Println()
	format = "ui16"
	argv = []string{"1024"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *ui16 1024 -> OK
	fmt.Println()
	format = "*ui16"
	argv = []string{"1024"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// ui16 0.10 -> err
	fmt.Println()
	format = "ui16"
	argv = []string{"0.10"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// ui32 17563 -> OK
	fmt.Println()
	format = "ui32"
	argv = []string{"17563"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *ui32 17563 -> OK
	fmt.Println()
	format = "*ui32"
	argv = []string{"17563"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// ui32 100.432 -> err
	fmt.Println()
	format = "ui32"
	argv = []string{"100.432"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// ui64 123456789 -> OK
	fmt.Println()
	format = "ui64"
	argv = []string{"123456789"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *ui64 123456789 -> OK
	fmt.Println()
	format = "*ui64"
	argv = []string{"123456789"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// ui64 10.23456 -> err
	fmt.Println()
	format = "ui64"
	argv = []string{"10.23456"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// rune
	// r 12456 -> OK
	fmt.Println()
	format = "r"
	argv = []string{"12456"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *r 12456 -> OK
	fmt.Println()
	format = "*r"
	argv = []string{"12456"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// r 10.23456 -> err
	fmt.Println()
	format = "r"
	argv = []string{"10.23456"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// byte
	// by 124 -> OK
	fmt.Println()
	format = "by"
	argv = []string{"124"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *by 124 -> OK
	fmt.Println()
	format = "*by"
	argv = []string{"124"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// by 73.256 -> err
	fmt.Println()
	format = "by"
	argv = []string{"73.256"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// float
	// f32 125.2908 -> OK
	fmt.Println()
	format = "f32"
	argv = []string{"125.2908"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *f32 125.2908 -> OK
	fmt.Println()
	format = "*f32"
	argv = []string{"125.2908"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// f32 73.256d -> err
	fmt.Println()
	format = "f32"
	argv = []string{"73.256d"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// f64 12596.290822 -> OK
	fmt.Println()
	format = "f64"
	argv = []string{"12596.290822"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *f64 12596.290822 -> OK
	fmt.Println()
	format = "*f64"
	argv = []string{"12596.290822"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// f64 73.256d -> err
	fmt.Println()
	format = "f64"
	argv = []string{"73.256d"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// boolean
	// b true -> OK
	fmt.Println()
	format = "b"
	argv = []string{"true"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *b true -> OK
	fmt.Println()
	format = "*b"
	argv = []string{"true"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// b false -> OK
	fmt.Println()
	format = "b"
	argv = []string{"false"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// string
	// s false -> OK
	fmt.Println()
	format = "s"
	argv = []string{"salutare"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *s false -> OK
	fmt.Println()
	format = "*s"
	argv = []string{"salutare"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// complex
	// c64 10.23 15.222 -> OK
	fmt.Println()
	format = "c64"
	argv = []string{"10.23", "15.222"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// *c64 10.23 15.222 -> OK
	fmt.Println()
	format = "*c64"
	argv = []string{"10.23", "15.222"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// c128 134.923 93.123 -> OK
	format = "c128"
	argv = []string{"134.923", "93.123"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// c128 134.923s 93.123e -> err
	format = "c128"
	argv = []string{"134.923s", "93.123e"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// c128 134.923 93.123e -> err
	format = "c128"
	argv = []string{"134.923", "93.123e"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// iii32 10 20 30
	format = "iii32"
	argv = []string{"10", "20", "30"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// i8i16f32 10 200 123.56
	format = "i8i16f32"
	argv = []string{"10", "200", "123.56"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// i8sf64c128s 100 salut 3.14159 12.4 15.3 merge
	format = "i8sf64c128s"
	argv = []string{"100", "salut", "3.14159", "12.4", "15.3", "merge"}
	err = testParseArg01(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// i8i64f 13 15 -> invalid format
	format = "i8i64f32"
	argv = []string{"13", "15"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// i8i64f 13 15 13.2 113 -> insuficient specifiers
	format = "i8i64f32"
	argv = []string{"13", "15", "13.2", "113"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	//////////////////////////////////////////////////////////////////
	// unsized slice test.
	fmt.Println("\n=== Unsized slice test ===")
	// []ic128 134.923 93.123 -> ok
	format = "[]ic128"
	argv = []string{"134.923", "93.123e"}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	format = "[]i[]ui8"
	argv = []string{}
	err = testParseArg01(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
}
