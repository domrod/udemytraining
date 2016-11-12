package main

import "fmt"

func main() {
	a := 42

	fmt.Println("a:", a)
	fmt.Println("a address:", &a)
	fmt.Printf("a address (converted in decimal): %d \n", &a)
	fmt.Printf("a address (converted in binary): %b \n", &a)

}
