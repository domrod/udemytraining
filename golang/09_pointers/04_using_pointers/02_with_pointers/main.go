package main

import "fmt"

func zero(y *int) {
	*y = 0
}

func main() {

	x := 5
	zero(&x)             // Memory address of x is transfered
	fmt.Println("x:", x) // Prints the value obtained in func zero by DEREFERENCING: 0
}
