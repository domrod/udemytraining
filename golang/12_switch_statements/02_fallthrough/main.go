package main

import "fmt"

func main() {

	name := "Dun"
	switch name {
	case "Dim":
		fmt.Println("Hi Dim")
	case "Dun":
		fmt.Println("Hi Dun")
		fallthrough // Executes the next statement even if it is not checked by the case statement
	case "Dom":
		fmt.Println("Hi Dom")
	default:
		fmt.Println("I don't know you", name)
	}

}
