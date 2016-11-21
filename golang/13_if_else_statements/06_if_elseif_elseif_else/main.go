package main

import "fmt"

func main() {

	if false {
		fmt.Println("First statement")
	} else if false {
		fmt.Println("Second statement")
	} else if true {
		fmt.Println("Third statement")
	} else {
		fmt.Println("What else than false or true?")
	}

}
