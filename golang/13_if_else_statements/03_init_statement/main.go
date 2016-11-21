package main

import "fmt"

func main() {

	c := true
	if food := "ham"; c { // On this line, food is initialized on the fly while the condition if c is tested
		// Note that the scope of "food" will be only between these braces
		fmt.Println(food)
	}

}
