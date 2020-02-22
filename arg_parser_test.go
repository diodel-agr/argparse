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
	for i := 0; i < a.Len(); i++ {
		if a.Len() != b.Len() {
			return false
		}
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

func TestBrakeFormat(t *testing.T) {
	fmt.Println("TestBrakeFormat test function.")
	/* iisf -> [i, i, s, f ] */
	result := BrakeFormat("iisf")
	if test_BrakeFormat01(result) == false {
		t.Errorf("Invalid TestBrakeFormat %s", "iisf")
	} else {
		fmt.Println("Test iisf OK.")
	}
	/* []isi32[]10f -> [ []i, s, i32, []10f ] */
	result = BrakeFormat("[]isi32[]10f")
	if test_BrakeFormat01(result) == false {
		t.Errorf("Invalid TestBrakeFormat %s", "[]isi32[]10f")
	} else {
		fmt.Println("Test []isi32[]10f OK.")
	}
}
