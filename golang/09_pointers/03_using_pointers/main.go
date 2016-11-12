package main

import "fmt"

func main() {

	a := 42

	fmt.Println("a:", a) // Prints 42
	fmt.Println("a address:", &a)

	var b *int = &a
	fmt.Println("b:", b)
	fmt.Println("*b:", *b)

	*b = 24 // Changing the value at "a" memory address

	fmt.Println("a is now:", a) // Prints 24

}
