package main

import "fmt"

func main() {

	x := 0
	fmt.Print("Enter a integer value:")
	fmt.Scan(&x)

	r := x % 2  // The remainder of division of x per 2
	if r == 0 { // Nota: use r == 0 (condition) and NOT r = 0 (assign value)
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "is odd")
	}
}
