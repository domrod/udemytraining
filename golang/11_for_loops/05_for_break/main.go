package main

import "fmt"

func main() {

	i := 0

	for { // infinite loop (no a priori condition)
		fmt.Println(i)
		if i >= 10 { // condition to leave the loop
			break
		}
		i++
	}
}
