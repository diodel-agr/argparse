package main

import (
	"container/list"
	"fmt"
	"strconv"
)

// Specifier - structure used to store the information regarding a specifier.
type Specifier struct {
	slice       bool
	pointerType bool
	sliceSize   uint
	bitSize     uint
	theType     string
}

// NewSpecifier - function used to create aSpecifier object.
func NewSpecifier(l *list.List) *Specifier {
	slice := false
	pointerType := false
	sliceSize = 0
	theType = ""
	bitSize = 0
	el := l.Front()
	// slice.
	str := el.Value.(string)
	if str != "" {
		slice = true
	}
	el = el.Next()
	// sliceSize.
	str = el.Value.(string)
	if str != "" {
		// convert string to int.
		pointerType, _ = strconv.Atoi(str)
	}
	el = el.Next()
	// pointerType.
	str = el.Value.(string)
	if str != "" {
		pointerType = true
	}
	el = el.Next()
	// theType.
	theType = el.Value.(string)
	el = el.Next()
	// bitSize.
	str = el.Value.(string)
	if str != "" {
		bitSize, _ = strconv.Atoi(str)
	}
	// create the result.
	return &Specifier{slice, pointerType, sliceSize, bitSize, theType}
}

// CheckFormat - function used to check if the format given is valid.
// @format: the format passed on the command line.
// @return: true if the format is valid, false otherwise.
// This function brakes the @format into a list of specifiers and checks
// wether the specifiers are valid and are compatible.
// Example: []i[]i is not valid because the specifiers are not compatible.
// The program will not know when one array ends and the other begins.
func CheckFormat(format string) (bool, string) {
	fmt.Println("Checking format:", format)
	speclist := BrakeFormat(format) // speclist = the list of specifiers.
	sl := []Specifier{}
	fmt.Print("Specifier list: ")
	PrintList(speclist)
	// check if each specifier is valid.
	for spec := speclist.Front(); spec != nil; spec = spec.Next() {
		slist := SplitSpecifier(spec.Value.(string))
		fmt.Print("Split: ")
		PrintList(slist)
		result, err := CheckSpecifier(slist)
		if result == false {
			fmt.Println("Error at specifier", slist, ":", err)
			return result, err
		}
		// create and add a new specifier to the specifier slice.
		sl.PushBack()
	}
	// check if the specifiers are compatible.
	// return true.
	return true, ""
}
