package main

import (
	"fmt"
)

// x is visible locally by all functions of this package.
// This is package level scope.
var x = 42

func foo() {
	fmt.Printf("x has %T type \n", x)
	// y is only in this function.
	// This is block level scope.
	y := "Hi"
	fmt.Printf("y = %#v \n", y)
}

func main() {
	fmt.Println("x:", x)
	foo()
}
