package main

import (
	"fmt"
)

func main() {

	var a float64
	var b string
	var c int32
	var d bool
	// Nota: %#v instead of %v gives output in Golang format
	fmt.Printf("%#v \n", a)
	fmt.Printf("%#v \n", b)
	fmt.Printf("%#v \n", c)
	fmt.Printf("%#v \n", d)

}
