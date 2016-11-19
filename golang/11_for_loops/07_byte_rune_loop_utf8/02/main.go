package main

import "fmt"

func main() {
	for i := 50; i < 140; i++ {
		fmt.Printf("%v - %v - %v - %v \n", i, string(i), []byte(string(i)), []rune(string(i)))
	}

	r := 'â' // NOTA: this is a rune, NOT the string "â"

	fmt.Printf("'â' is %v and is of type %T \n", r, r)
	fmt.Printf("rune('â') is %v and is of type %T \n", rune(r), rune(r))

	s := "â" // this is the string representation

	fmt.Printf("\"%v\" is %v and is of type %T \n", s, s, s)
	// fmt.Printf(`"%v" is %v and is of type %T`, s, s, s) // a way to avoid back slashes
	// fmt.Println("") // "\n" can't be included in the previous Printf

}
