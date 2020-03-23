package main

import (
	"fmt"
	"reflect"
)

// i, ui, r, by
// f
// c
// b, s

func main() {
	listOfVariables := ParseArgList()
	fmt.Println("The result: ")
	i := 0
	for e := listOfVariables.Front(); e != nil; e = e.Next() {
		fmt.Println(i, ": ", e.Value, "(", reflect.TypeOf(e.Value), ")")
		i++
	}
}
