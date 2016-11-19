package main

import "fmt"

func main() {

	i := 0

	for { // a priori infinite loop
		i++ // increments loop
		if i%2 == 0 {
			continue // goes directly to the next loop cycle, without executing the following lines
		}
		fmt.Println(i)
		if i >= 10 { // leaving loop condition
			break
		}

	}
}
