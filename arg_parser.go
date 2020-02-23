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

// BrakeFormat - function used to brake the format string into a list of formats.
// @fmt: the format.
// @return: list of strings representing the types.
// Possible types: i8, i16, i32, i64, ui8, ui16, ui32, ui64
//                   i, ui, r, by, f32, f64, c32, c64, b, s.
// Each type can be expressed as an array, like []i (array of integers) or
// []5i (array of 5 integers).
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

// CheckFormat - function used to check if the format given is valid.
// @fmt: the format.
// @return: true if the format is valid, false otherwise.
// The types from the @flist argument has to be valid types.
// Example: []i[]i is not valid because the program will not know when one array ends and the other begins.
func CheckFormat(flist *list.List) bool {
	/*
		    isArray := false
			arraySize := 0
			numbers := "0123456789"
			   if c == '[' { // array type.
			       isArray = true
			   } else if c == ']' && isArray == false {
			       // format error.
			       panic(errors.New("Invalid format:", format))
			   } else if string.Contains(c, numbers) == true { // number detected.
			       if isArray == false {
			           // format error.
			           panic(errors.New("Invalid format:", format))
			       } else if arraySize == 0 && c == '0' {
			           panic(errors.New("Invalid format:", format, "Array size begins with or is 0."))
			       }
			       // increase array size.
			       arraySize = arraySize * 10 + (c - '0')
			   } else { // a type is expected.
	*/
	return true
}

// ParseArgList - function used to parse the arguments from the commamd line.
// @input: the arguments from os.Args[1:]
// @return: [][]list of variables.
func ParseArgList() *[]list.List {
	argv := os.Args[1:]
	if len(argv) == 0 {
		fmt.Println("No command line arguments.")
		return nil
	} else {
		format := argv[0]
		flist := BrakeFormat(format)
		PrintList(flist)
	}
	return nil
}
