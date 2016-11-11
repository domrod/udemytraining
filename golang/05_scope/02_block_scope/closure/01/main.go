package main

import (
	"fmt"
)

var x = 42

func main() {

	fmt.Println(x)
	{
		fmt.Println(x)
		y := 45 // This variable is only accessible inside the braces
		fmt.Println(y)
	}

	// fmt.Println(y) // This does not work

}
