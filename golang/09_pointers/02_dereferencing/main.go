package main

import "fmt"

func main() {

	a := 42

	fmt.Println("a:", a)
	fmt.Println("a address:", &a)

	var b *int = &a
	fmt.Println("b:", b) // Prints 'a memory address'
	fmt.Println(*b)      // Prints the VALUE stored at this address
	// This is called DEREFERENCING
}
