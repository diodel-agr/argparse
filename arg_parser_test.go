package main

import (
	"container/list"
	"testing"
)

func CompareLists(a, b list.List) bool {
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

func TestBrakeFormat(t *testing.T) {
	/* iisf -> [i, i, s, f ] */
	result := BrakeFormat("iisf")
	expected := list.New()
	expected.PushBack("i")
	expected.PushBack("i")
	expected.PushBack("s")
	expected.PushBack("f")
	if CompareLists(*result, *expected) == false {
		t.Errorf("Invalid TestBrakeFormat %s", "iisf")
	}
}
