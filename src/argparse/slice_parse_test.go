package argparse

import (
	"fmt"
	"testing"
)

func testSliceParse(format string, argv []string) string {
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

func TestConvertSizedSlice(t *testing.T) {
	fmt.Println("\n=== parseArguments[sized slice] ===")
	// int
	// []1i 10 -> OK
	format := "[]1i"
	argv := []string{"10"}
	err := testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid format %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*i 10 -> OK
	format = "[]1*i"
	argv = []string{"10"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []2i 12.6 -> err0
	fmt.Println()
	format = "[]2i"
	argv = []string{"12.6"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1i8 100 -> OK
	format = "[]1i8"
	argv = []string{"100"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid format %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*i8 99 -> OK
	format = "[]1*i8"
	argv = []string{"99"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1i8 100 -> OK
	fmt.Println()
	format = "[]1i8"
	argv = []string{"100"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*i8 100 -> OK
	fmt.Println()
	format = "[]1*i8"
	argv = []string{"100"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1i8 123.4 -> err0
	fmt.Println()
	format = "[]1i8"
	argv = []string{"123.4"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1i16 1024 -> OK
	fmt.Println()
	format = "[]1i16"
	argv = []string{"1024"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*i16 1024 -> OK
	fmt.Println()
	format = "[]1*i16"
	argv = []string{"1024"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1i16 0.10 -> err
	fmt.Println()
	format = "[]1i16"
	argv = []string{"0.10"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1i32 17563 -> OK
	fmt.Println()
	format = "[]1i32"
	argv = []string{"17563"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*i32 17563 -> OK
	fmt.Println()
	format = "[]1*i32"
	argv = []string{"17563"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1i32 100.432 -> err
	fmt.Println()
	format = "[]1i32"
	argv = []string{"100.432"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1i64 123456789 -> OK
	fmt.Println()
	format = "[]1i64"
	argv = []string{"123456789"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*i64 123456789 -> OK
	fmt.Println()
	format = "[]1*i64"
	argv = []string{"123456789"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1i64 10.23456 -> err
	fmt.Println()
	format = "[]1i64"
	argv = []string{"10.23456"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// uint
	// []1ui 10 -> OK
	format = "[]1ui"
	argv = []string{"10"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*ui 10 -> OK
	format = "[]1*ui"
	argv = []string{"10"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1ui 12.6 -> err0
	fmt.Println()
	format = "[]1ui"
	argv = []string{"12.6"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1ui8 100 -> OK
	fmt.Println()
	format = "[]1ui8"
	argv = []string{"100"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*ui8 100 -> OK
	fmt.Println()
	format = "[]1*ui8"
	argv = []string{"100"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1ui8 123.4 -> err0
	fmt.Println()
	format = "[]1ui8"
	argv = []string{"123.4"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1ui16 1024 -> OK
	fmt.Println()
	format = "[]1ui16"
	argv = []string{"1024"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*ui16 1024 -> OK
	fmt.Println()
	format = "[]1*ui16"
	argv = []string{"1024"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1ui16 0.10 -> err
	fmt.Println()
	format = "[]1ui16"
	argv = []string{"0.10"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1ui32 17563 -> OK
	fmt.Println()
	format = "[]1ui32"
	argv = []string{"17563"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*ui32 17563 -> OK
	fmt.Println()
	format = "[]1*ui32"
	argv = []string{"17563"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1ui32 100.432 -> err
	fmt.Println()
	format = "[]1ui32"
	argv = []string{"100.432"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1ui64 123456789 -> OK
	fmt.Println()
	format = "[]1ui64"
	argv = []string{"123456789"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*ui64 123456789 -> OK
	fmt.Println()
	format = "[]1*ui64"
	argv = []string{"123456789"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1ui64 10.23456 -> err
	fmt.Println()
	format = "[]1ui64"
	argv = []string{"10.23456"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// rune
	// []1r 12456 -> OK
	fmt.Println()
	format = "[]1r"
	argv = []string{"12456"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*r 12456 -> OK
	fmt.Println()
	format = "[]1*r"
	argv = []string{"12456"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1r 10.23456 -> err
	fmt.Println()
	format = "[]1r"
	argv = []string{"10.23456"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// byte
	// []1by 124 -> OK
	fmt.Println()
	format = "[]1by"
	argv = []string{"124"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*by 124 -> OK
	fmt.Println()
	format = "[]1*by"
	argv = []string{"124"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1by 73.256 -> err
	fmt.Println()
	format = "[]1by"
	argv = []string{"73.256"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// float
	// []1f32 125.2908 -> OK
	fmt.Println()
	format = "[]1f32"
	argv = []string{"125.2908"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*f32 125.2908 -> OK
	fmt.Println()
	format = "[]1*f32"
	argv = []string{"125.2908"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1f32 73.256d -> err
	fmt.Println()
	format = "[]1f32"
	argv = []string{"73.256d"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1f64 12596.290822 -> OK
	fmt.Println()
	format = "[]1f64"
	argv = []string{"12596.290822"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*f64 12596.290822 -> OK
	fmt.Println()
	format = "[]1*f64"
	argv = []string{"12596.290822"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1f64 73.256d -> err
	fmt.Println()
	format = "[]1f64"
	argv = []string{"73.256d"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// boolean
	// []1b true -> OK
	fmt.Println()
	format = "[]1b"
	argv = []string{"true"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*b true -> OK
	fmt.Println()
	format = "[]1*b"
	argv = []string{"true"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1b false -> OK
	fmt.Println()
	format = "[]1b"
	argv = []string{"false"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// string
	// []1s false -> OK
	fmt.Println()
	format = "[]1s"
	argv = []string{"salutare"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*s false -> OK
	fmt.Println()
	format = "[]1*s"
	argv = []string{"salutare"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// complex
	// []1c64 10.23 15.222 -> OK
	fmt.Println()
	format = "[]1c64"
	argv = []string{"10.23", "15.222"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*c64 10.23 15.222 -> OK
	fmt.Println()
	format = "[]1*c64"
	argv = []string{"10.23", "15.222"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1c128 134.923 93.123 -> OK
	fmt.Println()
	format = "[]1c128"
	argv = []string{"134.923", "93.123"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1*c128 134.923 93.123 -> OK
	fmt.Println()
	format = "[]1*c128"
	argv = []string{"134.923", "93.123"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1c128 134.923s 93.123e -> err
	fmt.Println()
	format = "[]1c128"
	argv = []string{"134.923s", "93.123e"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []1c128 134.923 93.123e -> err
	fmt.Println()
	format = "[]1c128"
	argv = []string{"134.923", "93.123e"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []2uiiii32 3 4 10 20 30
	fmt.Println()
	format = "[]2uiiii32"
	argv = []string{"3", "4", "10", "20", "30"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []3f32 10 200 123.56
	fmt.Println()
	format = "[]3f32"
	argv = []string{"10", "200", "123.56"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// i8sf64c128s 100 salut 3.14159 12.4 15.3 merge
	fmt.Println()
	format = "i8sf64c128s"
	argv = []string{"100", "salut", "3.14159", "12.4", "15.3", "merge"}
	err = testSliceParse(format, argv)
	if err != "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println("Test", format, argv, "OK.")
	}
	// []2f 13 15 -> invalid format
	fmt.Println()
	format = "[]2f32[]2i"
	argv = []string{"13", "15"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
	// []2i8f64 13 15 13.2 113 -> insuficient specifiers
	fmt.Println()
	format = "[]2i8f64"
	argv = []string{"13", "15", "13.2", "113"}
	err = testSliceParse(format, argv)
	if err == "" {
		fmt.Println("Format", format, " error:", err)
		t.Errorf("Invalid CheckFormat %s. Error: %s.", format, err)
	} else {
		fmt.Println(err)
		fmt.Println("Test", format, argv, "OK.")
	}
}
