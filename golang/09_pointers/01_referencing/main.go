package main

import "fmt"

func main() {

	a := 42

	fmt.Println("a:", a)
	fmt.Println("a address:", &a)

	b := &a // b is referenced to 'a memory address'
	// This is called REFERENCING
	// b is a pointer to the memory address of an int
	// It could have been written the following way:
	// var b *int = &a
	fmt.Println("b:", b)             // Prints 'a memory address'
	fmt.Printf("b is of type %T", b) // Prints type *int (int pointer)
}
