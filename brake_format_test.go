package main

import (
	"container/list"
	"fmt"
	"testing"
)

func CompareLists(a, b *list.List) bool {
	if a.Len() != b.Len() {
		return false
	}
	ea := a.Front()
	eb := b.Front()
	for i := 0; i != a.Len(); i++ {
		if ea.Value != eb.Value {
			return false
		}
		ea = ea.Next()
		eb = eb.Next()
	}
	return true
}

/* iisf -> [i, i, s, f ] */
func testBrakeFormat01(l *list.List) bool {
	expected := list.New()
	expected.PushBack("i")
	expected.PushBack("i")
	expected.PushBack("s")
	expected.PushBack("f")
	return CompareLists(l, expected)
}

/* []isi32[]10f */
func testBrakeFormat02(l *list.List) bool {
	expected := list.New()
	expected.PushBack("[]i")
	expected.PushBack("s")
	expected.PushBack("i32")
	expected.PushBack("[]10f")
	return CompareLists(l, expected)
}

/* []ibyrc32s[]10byi */
func testBrakeFormat03(l *list.List) bool {
	expected := list.New()
	expected.PushBack("[]i")
	expected.PushBack("by")
	expected.PushBack("r")
	expected.PushBack("c32")
	expected.PushBack("s")
	expected.PushBack("[]10by")
	expected.PushBack("i")
	return CompareLists(l, expected)
}

/* byui -> [by, ui] */
func testBrakeFormat04(l *list.List) bool {
	expected := list.New()
	expected.PushBack("by")
	expected.PushBack("ui")
	return CompareLists(l, expected)
}

/* i8i16i32i64ui8ui16ui32ui64iuirbyf32f64c32c64bs ->
   [i8, i16, i32, i64, ui8, ui16, ui32, ui64, i, ui, r, by, f32, f64, c32, c64, b, s] */
func testBrakeFormat05(l *list.List) bool {
	expected := list.New()
	expected.PushBack("i8")
	expected.PushBack("i16")
	expected.PushBack("i32")
	expected.PushBack("i64")
	expected.PushBack("ui8")
	expected.PushBack("ui16")
	expected.PushBack("ui32")
	expected.PushBack("ui64")
	expected.PushBack("i")
	expected.PushBack("ui")
	expected.PushBack("r")
	expected.PushBack("by")
	expected.PushBack("f32")
	expected.PushBack("f64")
	expected.PushBack("c32")
	expected.PushBack("c64")
	expected.PushBack("b")
	expected.PushBack("s")
	return CompareLists(l, expected)
}

func TestBrakeFormat(t *testing.T) {
	fmt.Println("TestBrakeFormat test function.")
	/* iisf -> [i, i, s, f ] */
	result := BrakeFormat("iisf")
	if testBrakeFormat01(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestBrakeFormat %s", "iisf")
	} else {
		fmt.Println("Test iisf OK.")
	}
	/* []isi32[]10f -> [ []i, s, i32, []10f ] */
	result = BrakeFormat("[]isi32[]10f")
	if testBrakeFormat02(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestBrakeFormat %s", "[]isi32[]10f")
	} else {
		fmt.Println("Test []isi32[]10f OK.")
	}
	/* []ibyrc32s[]10byi */
	result = BrakeFormat("[]ibyrc32s[]10byi")
	if testBrakeFormat03(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestBrakeFormat %s.", "[]ibyrc32s[]10byi")
	} else {
		fmt.Println("Test []ibyrc32s[]10byi OK.")
	}
	/* byui -> [by, ui] */
	result = BrakeFormat("byui")
	if testBrakeFormat04(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestBrakeFormat %s.", "byui")
	} else {
		fmt.Println("Test byui OK.")
	}
	/* i8i16i32i64ui8ui16ui32ui64iuirbyf32f64c32c64bs ->
	   [i8, i16, i32, i64, ui8, ui16, ui32, ui64, i, ui, r, by, f32, f64, c32, c64, b, s] */
	result = BrakeFormat("i8i16i32i64ui8ui16ui32ui64iuirbyf32f64c32c64bs")
	if testBrakeFormat05(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestBrakeFormat %s.", "i8i16i32i64ui8ui16ui32ui64iuirbyf32f64c32c64bs")
	} else {
		fmt.Println("Test i8i16i32i64ui8ui16ui32ui64iuirbyf32f64c32c64bs OK.")
	}
}
