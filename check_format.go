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

// NewSpecifier - function used to create a Specifier object.
func NewSpecifier(l *list.List) *Specifier {
	slice := false
	pointerType := false
	sliceSize := uint(0)
	theType := ""
	bitSize := uint(0)
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
		result, _ := strconv.Atoi(str)
		sliceSize = uint(result)
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
		result, _ := strconv.Atoi(str)
		bitSize = uint(result)
	}
	// create the result.
	return &Specifier{slice, pointerType, sliceSize, bitSize, theType}
}

func (s Specifier) toString() string {
	result := ""
	// slice
	if s.slice == true {
		result = result + "[]"
	}
	// sliceSize.
	if s.sliceSize != 0 {
		result = result + strconv.Itoa(int(s.sliceSize))
	}
	// pointerType.
	if s.pointerType == true {
		result = result + "*"
	}
	// the type.
	result = result + s.theType
	// bitSize.
	if s.bitSize != 0 {
		result = result + strconv.Itoa(int(s.bitSize))
	}
	return result
}

// isTypeCompatible - function used to check if the type of two slices are compatible.
// The incompatible types: [int, uint, rune, byte, float, complex] and [boolean, string]
func (s Specifier) isTypeCompatible(s2 *Specifier) bool {
	if s.theType == "i" || s.theType == "ui" || s.theType == "f" || s.theType == "c" || s.theType == "r" || s.theType == "by" {
		// s1 is integer, float or complex.
		if s2.theType == "i" || s2.theType == "ui" || s2.theType == "f" || s2.theType == "c" || s2.theType == "r" || s2.theType == "by" {
			// s2 is also integer, float or complex.
			return false
		}
	}
	if s.theType == "b" || s.theType == "s" {
		// s1 is boolean or string.
		if s2.theType == "b" || s2.theType == "s" {
			// s2 is also boolean or string.
			return false
		}
	}
	return true
}

// sameSpecifier - function used to check whether the s1 specifier is is compatible with the s2 specifier.
// The basic rule is that 2 types are not compatible if they are both slices and neither or them does not have the size specified.
// For instance, if the cmd line arguments are []i[]i 10 20 30 40, []i and []ui are not compatible because there is no way to decide
// when the elements of one slice ends and one begins. The possible results are [10], [20, 30, 40] or [10, 20], [30, 40] or ...
// []i and []2i are compatible.
// []i and i are compatible.
// []i and []f32 are compatible.
// []i and []ui are not compatible.
func (s Specifier) isCompatible(s2 *Specifier) bool {
	if s.slice == false || s2.slice == false {
		return true // one of them is not slice.
	}
	// both specifiers are slices.
	// check if either of the slices have a size specified.
	if s.sliceSize != 0 || s2.sliceSize != 0 {
		return true
	}
	// both specifiers are slices and don't have their size specified.
	// check if the 2 types are compatible.
	return s.isTypeCompatible(s2)
}

// checkCompatibility - function used to check if the types from a format are compatible.
// Example: []i[]i is not valid because you can't tell when an array ends and one begins.
func checkCompatibility(sl *[]Specifier) (bool, string) {
	length := len(*sl)
	if length <= 1 {
		return true, ""
	}
	s1 := (*sl)[0]
	for i := 1; i < length; i++ {
		s2 := (*sl)[i]
		if s1.isCompatible(&s2) == false {
			return false, "Incompatible types: " + s1.toString() + " and " + s2.toString()
		}
		s1 = s2
	}

	return true, ""
}

// CheckFormat - function used to check if the format given is valid.
// @format: the format passed on the command line.
// @return: true if the format is valid, false otherwise.
// This function brakes the @format into a list of specifiers and checks
// wether the specifiers are valid and are compatible.
// Example: []i[]i is not valid because the specifiers are not compatible.
// The program will not know when one array ends and the other begins.
func CheckFormat(format string) (bool, *[]Specifier, string) {
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
			return result, nil, err
		}
		// create and add a new specifier to the specifier slice.
		sl = append(sl, *NewSpecifier(slist))
	}
	// check if the specifiers are compatible.
	result, err := checkCompatibility(&sl)
	if result == false {
		fmt.Println(err)
		return false, nil, err
	}
	// all fine, return true.
	return true, &sl, ""
}
