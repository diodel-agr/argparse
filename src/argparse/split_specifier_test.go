package argparse

import (
	"container/list"
	"fmt"
	"testing"
)

/* i -> [- - - i -] */
func testSplitSpec0(l *list.List) bool {
	expected := list.New()
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("i")
	expected.PushBack("")
	return CompareLists(l, expected)
}

/* ui16 -> [- - - ui 16] */
func testSplitSpec1(l *list.List) bool {
	expected := list.New()
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("ui")
	expected.PushBack("16")
	return CompareLists(l, expected)
}

/* []f32 -> [[] - - f 32] */
func testSplitSpec2(l *list.List) bool {
	expected := list.New()
	expected.PushBack("[]")
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("f")
	expected.PushBack("32")
	return CompareLists(l, expected)
}

/* []*i64 -> [[] - * i 64] */
func testSplitSpec3(l *list.List) bool {
	expected := list.New()
	expected.PushBack("[]")
	expected.PushBack("")
	expected.PushBack("*")
	expected.PushBack("i")
	expected.PushBack("64")
	return CompareLists(l, expected)
}

/* []10*ui32 -> [[] 10 * ui 32] */
func testSplitSpec4(l *list.List) bool {
	expected := list.New()
	expected.PushBack("[]")
	expected.PushBack("10")
	expected.PushBack("*")
	expected.PushBack("ui")
	expected.PushBack("32")
	return CompareLists(l, expected)
}

/* []5s -> [[] 5 - s -] */
func testSplitSpec5(l *list.List) bool {
	expected := list.New()
	expected.PushBack("[]")
	expected.PushBack("5")
	expected.PushBack("")
	expected.PushBack("s")
	expected.PushBack("")
	return CompareLists(l, expected)
}

/* *by -> [- - * by -] */
func testSplitSpec6(l *list.List) bool {
	expected := list.New()
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("*")
	expected.PushBack("by")
	expected.PushBack("")
	return CompareLists(l, expected)
}

/* []10b -> [[] 10 - b -] */
func testSplitSpec7(l *list.List) bool {
	expected := list.New()
	expected.PushBack("[]")
	expected.PushBack("10")
	expected.PushBack("")
	expected.PushBack("b")
	expected.PushBack("")
	return CompareLists(l, expected)
}

/* []3*c32 -> [[] 3 * c 32] */
func testSplitSpec8(l *list.List) bool {
	expected := list.New()
	expected.PushBack("[]")
	expected.PushBack("3")
	expected.PushBack("*")
	expected.PushBack("c")
	expected.PushBack("32")
	return CompareLists(l, expected)
}

/* c64 -> [- - - c 64] */
func testSplitSpec9(l *list.List) bool {
	expected := list.New()
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("c")
	expected.PushBack("64")
	return CompareLists(l, expected)
}

/* "" -> [- - - - -] */
func testSplitSpec10(l *list.List) bool {
	expected := list.New()
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("")
	expected.PushBack("")
	return CompareLists(l, expected)
}

func TestSplitSpecifier(t *testing.T) {
	fmt.Println("\n=== SplitSpecifier ===")
	/* i -> ["" "" "" i "" ] */
	result := SplitSpecifier("i")
	if testSplitSpec0(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestSplitSpecifier %s", "i")
	} else {
		fmt.Println("Test i OK.")
	}
	/* ui16 -> [- - - ui 16] */
	result = SplitSpecifier("ui16")
	if testSplitSpec1(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestSplitSpecifier %s", "ui16")
	} else {
		fmt.Println("Test ui16 OK.")
	}
	/* []f32 -> [[] - - f 32] */
	result = SplitSpecifier("[]f32")
	if testSplitSpec2(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestSplitSpecifier %s", "[]f32")
	} else {
		fmt.Println("Test []f32 OK.")
	}
	/* []*i64 -> [[] - * i 64] */
	result = SplitSpecifier("[]*i64")
	if testSplitSpec3(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestSplitSpecifier %s", "[]*i64")
	} else {
		fmt.Println("Test []*i64 OK.")
	}
	/* []10*ui32 -> [[] 10 * ui 32] */
	result = SplitSpecifier("[]10*ui32")
	if testSplitSpec4(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestSplitSpecifier %s", "[]10*ui32")
	} else {
		fmt.Println("Test []10*ui32 OK.")
	}
	/* []5s -> [[] 5 - s -] */
	result = SplitSpecifier("[]5s")
	if testSplitSpec5(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestSplitSpecifier %s", "[]5s")
	} else {
		fmt.Println("Test []5s OK.")
	}
	/* *by -> [- - * by -] */
	result = SplitSpecifier("*by")
	if testSplitSpec6(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestSplitSpecifier %s", "*by")
	} else {
		fmt.Println("Test *by OK.")
	}
	/* []10b -> [[] 10 - b -] */
	result = SplitSpecifier("[]10b")
	if testSplitSpec7(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestSplitSpecifier %s", "[]10b")
	} else {
		fmt.Println("Test []10b OK.")
	}
	/* c64 -> [- - - c 64] */
	result = SplitSpecifier("c64")
	if testSplitSpec9(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestSplitSpecifier %s", "c64")
	} else {
		fmt.Println("Test c64 OK.")
	}
	/* "" -> [- - - - -] */
	result = SplitSpecifier("")
	if testSplitSpec10(result) == false {
		PrintList(result)
		t.Errorf("Invalid TestSplitSpecifier %s", "")
	} else {
		fmt.Println("Test \"\" OK.")
	}
}
