package main

import (
	"container/list"
	"fmt"
	"strconv"
)

// numericVar - function used to convert the the string @s to int.
// @s: the string representing a numeric type.
// @return the integer.
func numericVar(s string) (interface{}, string) {
	result, err := strconv.Atoi(s)
	if err != nil {
		return nil, "Error converting from " + s + " to "
	}
	return result, ""
}

// createVariable - function used to create the variable corresponding on the specifier @s type.
// @s: specifier.
// @v: the string containing the values to convert.
// @return: the variable.
func createVariable(s Specifier, v []string) (interface{}, string) {
	// 2 variable care vor tine adresa unor functii.
	//var varType func(string) (interface{}, string)
	//var varconv func(string) (interface{}, string)
	// in urmatorul switch vom seta adresele functiilor
	// apoi apelam functiile si intoarcem rezultatul.
	switch s.theType {
	case "i": // int
	case "ui": // unsigned int
	case "r": // rune (int32)
	case "by": // byte (uint8)
		//varType = numericVar
	case "f": // float
		//
	case "c": // complex
		//
	case "b": // boolean
		//
	case "s": // string

	}
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
