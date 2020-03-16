package main

import (
	"container/list"
	"fmt"
	"strconv"
)

var (
	intBitSize = 64
)

func (s Specifier) convertToInt(str string, ls *list.List) string {
	// check bit size.
	switch s.bitSize {
	case 0:
		result, err := strconv.ParseInt(str, 10, intBitSize)
		if err != nil {
			return "Error: " + err.Error()
		}
		ls.PushBack(int(result))
	case 8:
		result, err := strconv.ParseInt(str, 10, 8)
		if err != nil {
			return "Error: " + err.Error()
		}
		ls.PushBack(int8(result))
	case 16:
		result, err := strconv.ParseInt(str, 10, 16)
		if err != nil {
			return "Error: " + err.Error()
		}
		ls.PushBack(int16(result))
	case 32:
		result, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return "Error: " + err.Error()
		}
		ls.PushBack(int32(result))
	case 64:
		result, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return "Error: " + err.Error()
		}
		ls.PushBack(int64(result))
		/*
			default:
				return "Unknown integer bit size: " + strconv.Itoa(s.bitSize)
		*/
	}
	return ""
}

func (s Specifier) convertToUInt(str string, ls *list.List) string {
	// check bit size.
	switch s.bitSize {
	case 0:
		result, err := strconv.ParseUint(str, 10, intBitSize)
		if err != nil {
			return "Error: " + err.Error()
		}
		ls.PushBack(uint(result))
	case 8:
		result, err := strconv.ParseUint(str, 10, 8)
		if err != nil {
			return "Error: " + err.Error()
		}
		ls.PushBack(uint8(result))
	case 16:
		result, err := strconv.ParseUint(str, 10, 16)
		if err != nil {
			return "Error: " + err.Error()
		}
		ls.PushBack(uint16(result))
	case 32:
		result, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return "Error: " + err.Error()
		}
		ls.PushBack(uint32(result))
	case 64:
		result, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return "Error: " + err.Error()
		}
		ls.PushBack(uint64(result))
		/*
			default:
				return "Unknown integer bit size: " + strconv.Itoa(s.bitSize)
		*/
	}
	return ""
}

// integerVar - function used to convert the the string @s to a integer type.
// @s: the string representing a numeric type.
// @return the integer or an error message.
func (s Specifier) integerVar(str string, ls *list.List) string {
	if s.theType == "i" { // int or uint.
		return s.convertToInt(str, ls)
	} else if s.theType == "ui" {
		return s.convertToUInt(str, ls)
	} else if s.theType == "r" { // rune, alias int32.
		result, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return "Error converting string to rune (int32). Error: " + err.Error()
		}
		ls.PushBack(rune(result))
	} else if s.theType == "by" { // byte, alias int8.
		result, err := strconv.ParseUint(str, 10, 8)
		if err != nil {
			return "Error converting string to byte (int8). Error: " + err.Error()
		}
		ls.PushBack(byte(result))
	}
	return ""
}

// floatVar - function used to convert @s string to a float variable of bit size @size.
// @s: the string to convert.
// @size: the bit size.
// @return: the float variable or an error message.
func (s Specifier) floatVar(str string, ls *list.List) string {
	result, err := strconv.ParseFloat(str, s.bitSize)
	if err != nil {
		return "Error: " + err.Error()
	}
	if s.bitSize == 32 {
		ls.PushBack(float32(result))
	} else {
		ls.PushBack(result)
	}
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

// convertSimple - function used to convert the contents of the @str string to the type defined by the @s specifier.
// @s: the specifier.
// @val: the string to convert.
// @return: the created variable os a error message.
func (s Specifier) convertSimple(str string, ls *list.List) string {
	// check if the type is an integer type.
	if s.theType == "i" || s.theType == "ui" || s.theType == "r" || s.theType == "by" {
		return s.integerVar(str, ls)
	} else if s.theType == "f" {
		return s.floatVar(str, ls)
	} else if s.theType == "b" {
		return s.booleanVar(str, ls)
	} //else if s.theType == "s" {
	return s.stringVar(str, ls)
}

// convertSlice - function used to convert the @str slice which should store a complex variable.
// @s: the specifier.
// @str: slice containing the values to convert.
// @return: the complex number.
func (s Specifier) convertComplex(str []string, ls *list.List) string {
	// modify bit size.
	s.bitSize = s.bitSize >> 1
	fct := func() {
		s.bitSize = s.bitSize << 1
	}
	defer fct()
	// convert.
	re, err := strconv.ParseFloat(str[0], s.bitSize)
	if err != nil {
		return "Error converting string to float(complex). Error: " + err.Error()
	}
	im, err := strconv.ParseFloat(str[1], s.bitSize)
	if err != nil {
		return "Error converting string to float(complex). Error: " + err.Error()
	}
	// create variable.
	if s.bitSize == 32 {
		result := complex64(complex(re, im))
		ls.PushBack(result)
	} else {
		result := complex128(complex(re, im))
		ls.PushBack(result)
	}
	return ""
}

// convert - function used to wrap the convertion function.
// @s: the specifier.
// @str: the string to convert.
// @return: the variable or the error message.
func (s Specifier) convert(str []string, ai *int, ls *list.List) string {
	if s.theType == "c" {
		result := s.convertComplex(str[*ai:(*ai)+2], ls)
		*ai = *ai + 2
		return result
	}
	result := s.convertSimple(str[*ai], ls)
	*ai = *ai + 1
	return result
}

// parseArguments - function that reads the arguments from the command line and creates the variables.
// @slist: the list of specifiers.
// @argv: the slice of arguments from the command line.
// @return: the list of variables or the error string in case of error.
func parseArguments(slist []Specifier, argv []string) (*list.List, string) {
	fmt.Println("Parsing format:", slist)
	fmt.Println("Parse arguments:", argv)
	varList := list.New() // the result.
	si := 0               // specifier index.
	ai := 0               // argument index.
	slen := len(slist)    // specifier list length.
	alen := len(argv)     // argument list length.
	for si < slen && ai < alen {
		spec := slist[si] // current specifier.
		if spec.slice == false {
			// simple variable.
			err := spec.convert(argv[:], &ai, varList)
			if err != "" {
				return nil, err
			}
		} else if spec.sliceSize == 0 {
			// unsized slice
		} else {
			// sized slice.
		}
		si++
	}
	fmt.Println(si, ai, slen, alen)
	if si != slen {
		return nil, "The given format did not parsed all arguments!"
	}
	if ai != alen {
		return nil, "The format has insuficient specifiers"
	}
	return varList, ""
}
