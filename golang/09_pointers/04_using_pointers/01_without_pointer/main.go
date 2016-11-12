package main

import "fmt"

func zero(y int) {
	y = 0
}

func main() {

	x := 5
	zero(x)
	fmt.Println("x:", x) // Still prints 5 because Go transfers are done per VALUE
}
