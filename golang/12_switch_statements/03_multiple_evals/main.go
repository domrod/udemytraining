package main

import "fmt"

func main() {

	switch "Jenny" {
	case "Dom", "Tim":
		fmt.Println("Hi Dom or Tim")
	case "Ida", "Jenny":
		fmt.Println("Hi Ida or Jenny")
	default:
		fmt.Println("Hi someone")
	}
}
