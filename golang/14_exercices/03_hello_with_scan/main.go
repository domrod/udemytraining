package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var name string
	fmt.Printf("What is your name? ")
	// fmt.Scan(&name)
	in := bufio.NewReader(os.Stdin) // Better than Scan => allow strings with spaces
	name, _ := in.ReadString('\n')
	fmt.Printf("Hello %s", name)

}
