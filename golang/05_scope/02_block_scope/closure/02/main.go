package main

import (
	"fmt"
)

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(x)           // 0
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}
