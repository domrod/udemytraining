package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			fmt.Println(i, "is divisible by 3")
		} else {
			fmt.Println(i, "is not divisible by 3")
		}
	}
}
