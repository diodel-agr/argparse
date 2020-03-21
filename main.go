package main

import (
	"fmt"
	"reflect"
)

func main() {
	listOfVariables := ParseArgList()
	fmt.Println("The result: ")
	i := 0
	for e := listOfVariables.Front(); e != nil; e = e.Next() {
		fmt.Println(i, ": ", e.Value, "(", reflect.TypeOf(e.Value), ")")
		i++
	}
}

// []iuii[]i is not compatible.
