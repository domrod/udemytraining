package main

import (
	"fmt"
	"os"
)

func main() {

	var a, b int64
	fmt.Print("Enter a small number: ")
	fmt.Scanf("%v\n", &a)
	fmt.Print("Enter a large number (larger than the previous one): ")
	fmt.Scanf("%v\n", &b)
	if a > b {
		fmt.Printf("%v is larger than %v\n", a, b)
		os.Exit(0)
	}
	fmt.Printf("The remainder of %v divided by %v is: %v\n", b, a, b%a)
}
