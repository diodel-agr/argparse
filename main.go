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
	if listOfVariables != nil {
		fmt.Println("The result: ")
		i := 0
		for e := listOfVariables.Front(); e != nil; e = e.Next() {
			fmt.Println(i, ": ", e.Value, "(", reflect.TypeOf(e.Value), ")")
			i++
		}
	}
}

/* change the codition at the ParseArgList for loop */
/* BUG: ./argparse []i has 0 values so the loop does not execute to insert the emplty slice. */
