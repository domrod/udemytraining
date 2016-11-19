package main

import "fmt"

func main() {

	name := "Dum"
	switch name {
	case "Dim":
		fmt.Println("Hi my friend", name)
	case "Dun":
		fmt.Println("Hi my colleague", name)
	case "Dom":
		fmt.Println("My name is", name)
	default:
		fmt.Println("I don't know you", name)
	}

	// NOTA: there is no need for a "break" fallthrough
}
