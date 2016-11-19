package main

import "fmt"

func main() {

	i := 0

	for { // there is no condition, so this loop lasts for ever
		fmt.Println(i)
		i++
	}
}
