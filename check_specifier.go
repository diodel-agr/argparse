package main

import (
	"container/list"
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

func checkSlice(slice string) (bool, string) {
	if slice != "" && slice != "[]" {
		return false, ("Invalid slice specifier: " + slice) // TESTED.
	}
	return true, ""
}

func checkSliceSize(size string) (bool, string) {
	if size != "" {
		sizeInt, err := strconv.Atoi(size)
		if err != nil {
			return false, ("Invalid slice size: " + size)
		} else if sizeInt < 0 {
			return false, ("Slice size is negative: " + size)
		} else if sizeInt == 0 {
			return false, ("Slice size is zero: " + size)
		}
	}
	return true, ""
}

func checkPointerType(ptype string) (bool, string) {
	if ptype != "" && ptype != "*" {
		return false, ("Invalid pointer character: " + ptype)
	}
	return true, ""
}

func checkType(t string) (bool, string) {
	types := []string{"i", "ui", "r", "by", "f", "c", "b", "s"}
	length := len(types)

	for i := 0; i < length; i++ {
		if types[i] == t {
			return true, ""
		}
	}
	if t == "" {
		return false, "Missing type." // TESTED.
	}
	return false, ("Unknown type: " + t)
}

func checkBitSize(bsize string) (bool, string) {
	values := []int{8, 16, 32, 64}
	length := len(values)

	if bsize == "" {
		return true, ""
	}

	val, err := strconv.Atoi(bsize)
	if err != nil {
		return false, ("Invalid bit size: " + bsize)
	}
	for i := 0; i < length; i++ {
		if values[i] == val {
			return true, ""
		}
	}
	return false, "Invalit bit size: " + bsize + ". Valid values are: 8, 16, 32, 64."
}

// CheckSpecifier - function used to cerify if the elements of a specifier are valid.
// There are several rules:
// - the array specifier ( [] ) cannot be written like "[" or "]" or "][" or any other form.
// - the slice size has to be positive integer. Negative or real numbers are not permitted.
// - the pointer type has to be "*".
// - the type has to be a valid type.
// - the bit size has to be a positive integer and has to be in the list {8, 16, 32, 64}.
func CheckSpecifier(spec *list.List) (bool, string) {
	functions := []func(string) (bool, string){checkSlice, checkSliceSize, checkPointerType, checkType, checkBitSize}
	length := len(functions)

	el := spec.Front()

	for i := 0; i < length; i++ {
		result, err := functions[i](el.Value.(string))
		if result == false {
			return false, err
		}
		el = el.Next()
	}
	return true, ""
}
