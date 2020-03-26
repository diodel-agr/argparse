package main

import (
	"argparse"
	"fmt"
	"reflect"
)

// i, ui, r, by
// f
// c
// b, s
func main() {
	listOfVariables := argparse.ParseArgList()
	if listOfVariables != nil {
		fmt.Println("The result: ")
		for i := 0; i < listOfVariables.Len(); i++ {
			el := argparse.GetVariableAt(listOfVariables, i)
			fmt.Println(i, ": ", el, "(", reflect.TypeOf(el), ")")
		}
	}
}

// ./argparse []3i[]4s[]2f32 10 20 30 salut ce mai faci 3.14259 2.71
