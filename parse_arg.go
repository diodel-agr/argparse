package main

import (
	"container/list"
	"fmt"
	"strconv"
)

// integerVar - function used to convert the the string @s to int.
// @s: the string representing a numeric type.
// @return the integer or an error message.
func (s Specifier) integerVar(str string, ls *list.List) string {
	if s.theType == "i" || s.theType == "ui" { // int or uint.
		// check bit size.

	} else if s.theType == "r" { // rune, alias int32.

	} else if s.theType == "by" { // byte, alias int8.

	}
	result, err := strconv.ParseInt(str, 10, s.bitSize)
	if err != nil {
		return nil, "Error converting from string " + s + " to integer"
	}
	return result, ""
}

// floatVar - function used to convert @s string to a float variable of bit size @size.
// @s: the string to convert.
// @size: the bit size.
// @return: the float variable or an error message.
func (s Specifier) floatVar(str string, ls *list.List) string {
	result, err := strconv.ParseFloat(str, s.bitSize)
	if err != nil {
		return "Error converting string " + str + " to float"
	}
	ls.PushBack(result)
	return ""
}

func (s Specifier) booleanVar(str string, ls *list.List) string {
	if str == "true" {
		ls.PushBack(true)
	} else {
		ls.PushBack(false)
	}
	return ""
}

func (s Specifier) stringVar(str string, ls *list.List) string {
	ls.PushBack(str)
	return ""
}

// convertSimple - function used to convert the contents of the @s string to the type defined by the @s specifier.
// @s: the specifier.
// @val: the string to convert.
// @return: the created variable os a error message.
func (s Specifier) convertSimple(str string, ls *list.List) string {
	// check if the type is an integer type.
	if s.theType == "i" || s.theType == "ui" || s.theType == "r" || s.theType == "by" {
		return s.integerVar(str)
	} else if s.theType == "f" {
		return s.floatVar(str)
	} else if s.theType == "b" {
		return s.booleanVar(str)
	} else if s.theType == "s" {
		return s.stringVar(str)
	} else if s.theType == "c" {
		return nil, "Complex type in Specifier.Convert(string) function"
	}
	return nil, "Unknown type in Specifier.Convert(string) function"
}

// convertSlice - function used to convert the @str slice which should store a complex variable.
// @s: the specifier.
// @str: slice containing the values to convert.
// @return: the complex number.
func (s Specifier) convertSlice(str []string, ls *list.List) string {
	if s.theType == "c" {
		// modify bit size.
		s.bitSize = s.bitSize << 1
		fct := func() {
			s.bitSize = s.bitSize >> 1
		}
		defer fct()
		// convert using float function.
		re, err := s.floatVar(str[0])
		if err != "" {
			return nil, err
		}
		im, err := s.floatVar(str[1])
		if err != "" {
			return nil, err
		}
		// return result.
		if s.bitSize == 64 { // 128 bit complex.
			return complex128(complex(re.(float64), im.(float64))), ""
		}
		// 64 bit complex.
		return complex64(complex(re.(float32), im.(float32))), ""
	}
	return nil, "Non-complex type in Specifier.Convert([]string) function"
}

// Convert - function used to wrap the convertion function.
// @s: the specifier.
// @str: the string to convert.
// @return: the variable or the error message.
func (s Specifier) Convert(str []string, ls *list.List) string {
	if s.theType == "c" {
		return s.convertSlice(str, ls)
	}
	return s.convertSimple(str[0], ls)
}

// createVariable - function used to create the variable corresponding on the specifier @s type.
// @s: specifier.
// @v: the string containing the values to convert.
// @return: the variable.
// i/ui -> {0, 8, 16, 32, 64}
// r -> rune direct (int32)
// by -> byte direct (int8)
// f -> {32, 64}
// c -> {64, 128}
// b -> boolean
// s -> string
func createVariable(s Specifier, v []string) (interface{}, string) {
	// 2 variable care vor tine adresa unor functii.
	var varType func(string, int) (interface{}, string)
	var varConv func(string) (interface{}, string)
	// in urmatorul switch vom seta adresele functiilor
	// select the type to which to parse the string: integer, float or string.
	switch s.theType {
	case "i": // int
		varType = integerVar
	case "ui": // unsigned int
		varType = integerVar
	case "r": // rune (int32)
		varType = integerVar
	case "by": // byte (uint8)
		varType = integerVar
	case "f": // float
		varType = floatVar
	case "c": // complex
		varType = floatVar
	case "b": // boolean
		varType = booleanVar
	case "s": // string
		varType = stringVar
	default:
		return nil, "Unknown specifier type: " + s.theType
	}
	// select the exact type and bit size to which to convert.
	return nil, ""
}

// parseArguments - function that reads the arguments from the command line and creates the variables.
// @slist: the list of specifiers.
// @argv: the slice of arguments from the command line.
// @return: the list of variables or the error string in case of error.
func parseArguments(slist []Specifier, argv []string) (*list.List, string) {
	varList := list.New() // the result.
	si := 0               // specifier index.
	ai := 0               // argument index.
	slen := len(slist)    // specifier list length.
	alen := len(argv)     // argument list length.
	for si < slen {
		spec := slist[si] // current specifier.
		if spec.slice == false {
			// variable.
			result, err := createVariable(spec, argv[ai:ai+1])
			if err != "" {
				return nil, err
			}
			varList.PushBack(result)
		} else if spec.sliceSize == 0 {
			// unsized slice
		} else {
			// sized slice.
		}
	}
	if ai != alen {
		fmt.Println("The given format did not parsed all arguments!")
	}
	return nil, ""
}
