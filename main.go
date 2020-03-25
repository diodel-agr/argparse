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
		i := 0
		for e := listOfVariables.Front(); e != nil; e = e.Next() {
			fmt.Println(i, ": ", e.Value, "(", reflect.TypeOf(e.Value), ")")
			i++
		}
	}
}

// ./argparse []3i[]4s[]2f32 10 20 30 salut ce mai faci 3.14259 2.71
