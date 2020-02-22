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
func test_BrakeFormat01(l *list.List) bool {
	expected := list.New()
	expected.PushBack("i")
	expected.PushBack("i")
	expected.PushBack("s")
	expected.PushBack("f")
	return CompareLists(l, expected)
}

/* []isi32[]10f */
func test_BrakeFormat02(l *list.List) bool {
	expected := list.New()
	expected.PushBack("[]i")
	expected.PushBack("s")
	expected.PushBack("i32")
	expected.PushBack("[]10f")
	return CompareLists(l, expected)
}

/* []ibyrc32s[]10byi */
func test_BrakeFormat03(l *list.List) bool {
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
func test_BrakeFormat04(l *list.List) bool {
	expected := list.New()
	expected.PushBack("by")
	expected.PushBack("ui")
	return CompareLists(l, expected)
}

func TestBrakeFormat(t *testing.T) {
	fmt.Println("TestBrakeFormat test function.")
	/* iisf -> [i, i, s, f ] */
	result := BrakeFormat("iisf")
	if test_BrakeFormat01(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestBrakeFormat %s", "iisf")
	} else {
		fmt.Println("Test iisf OK.")
	}
	/* []isi32[]10f -> [ []i, s, i32, []10f ] */
	result = BrakeFormat("[]isi32[]10f")
	if test_BrakeFormat02(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestBrakeFormat %s", "[]isi32[]10f")
	} else {
		fmt.Println("Test []isi32[]10f OK.")
	}
	/* []ibyrc32s[]10byi */
	result = BrakeFormat("[]ibyrc32s[]10byi")
	if test_BrakeFormat03(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestBrakeFormat %s.", "[]ibyrc32s[]10byi")
	} else {
		fmt.Println("Test []ibyrc32s[]10byi OK.")
	}
	/**/
	result = BrakeFormat("byui")
	if test_BrakeFormat04(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestBrakeFormat %s.", "byui")
	} else {
		fmt.Println("Test byui OK.")
	}
}
