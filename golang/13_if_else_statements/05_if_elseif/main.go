package main

import "fmt"

func main() {

	if false {
		fmt.Println("First statement")
	} else if true {
		fmt.Println("Second statement")
	} else {
		fmt.Println("What else than false or true?")
	}
}
