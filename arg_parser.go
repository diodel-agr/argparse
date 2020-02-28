package main

import (
	"container/list"
	"fmt"
	"os"
	"strings"
)

// PrintList - function used to print the elements of a list.List object.
func PrintList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

// GetSpec - function used to return the string or the "-". Useful for the SplitSpecifier function.
func GetSpec(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

// BrakeFormat - function used to brake the format string into a list of specifiers.
// @format: the format.
// @return: list of strings representing the types.
// Possible types: i8, i16, i32, i64, ui8, ui16, ui32, ui64
//                 i, ui, r, by, f32, f64, c32, c64, b, s.
// Each type can be expressed as pointers, as an array, like []i
// (array of integers) or []5i (array of 5 integers).
func BrakeFormat(format string) *list.List {
	result := list.New()
	numbers := "0123456789y"
	types := "irbfcs"
	length := len(format)
	for i := 0; i < length; {
		j := i
		for (j+1 < length) && (strings.Index(types, format[j:j+1]) == -1) {
			j++
		}
		for (j+1 < length) && (strings.Index(numbers, format[j+1:j+2]) != -1) {
			j++
		}
		j++
		result.PushBack(format[i:j])
		i = j
	}
	return result
}

// SplitSpecifier - function used to split a specifier into its elements.
// A specifier has the following general format: []n*tx where,
// [] specifies whether the type is a slice or not.
// n is the slice size.
// * specifies if it is a pointer type.
// t is the type.
// x is the type size in bits.
//
// [] n * t x
// a  b c d e
func SplitSpecifier(t string) *list.List {
	letters := "abcdefghijklmnopqrstuvwxyz*" // "uirbfcs"
	a, b, c, d, e := 0, 0, 0, 0, 0           // the index of the specifier elements.
	j := 0
	length := len(t)
	/* check if it is array. */
	for j < length && ((t[j] == ']') || (t[j] == '[')) {
		j++
	}
	if j != a {
		b, c, d, e = j, j, j, j // j now points to the index of the array size.
	}
	/* find array size */
	for j < length && (strings.Index(letters, t[j:j+1]) == -1) {
		j++
	}
	if j != b {
		c, d, e = j, j, j // j now points to the index of the '*' or the type.
	}
	/* find if it is pointer type. */
	for j < length && t[j] == '*' {
		j++
	}
	if j != c {
		d, e = j, j
	}
	/* find the type. */
	for j < length && strings.Index(letters, t[j:j+1]) != -1 {
		j++
	}
	if j != d {
		e = j // j now points to the bits size.
	}
	slice := t[a:b]
	sliceSize := t[b:c]
	pointerType := t[c:d]
	theType := t[d:e]
	bitSize := t[e:]
	/* return result. */
	result := list.New()
	result.PushBack(slice)
	result.PushBack(sliceSize)
	result.PushBack(pointerType)
	result.PushBack(theType)
	result.PushBack(bitSize)
	return result
}

// ParseArgList - function used to parse the arguments from the commamd line.
// @input: the arguments from os.Args[1:]
// @return: [][]list of variables.
func ParseArgList() *list.List {
	argv := os.Args[1:]
	if len(argv) == 0 {
		fmt.Println("No command line arguments.")
		return nil
	} else {
		format := argv[0]
		flist := BrakeFormat(format)
		PrintList(flist)
		for l := flist.Front(); l != nil; l = l.Next() {
			SplitSpecifier(l.Value.(string))
		}
	}
	return nil
}
