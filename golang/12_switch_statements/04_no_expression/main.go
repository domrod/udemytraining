package main

import "fmt"

func main() {

	name := "Dom"
	switch {
	case name == "Dom":
		fmt.Println("Hi Dom")
		fallthrough
	case len(name) == 3:
		fmt.Printf("Hi %v. Your name's length is 3 \n", name)
	case name == "Ida", name == "Jenny":
		fmt.Println("hi Ida or Jenny")
	default:
		fmt.Println("Who are you?")

	}
}
