package main

// This exemple shows closure in action in a not so simple case

import "fmt"

func wrapper() func() int { // A function type may be a function also (like types such as int or string)!
	// Then, this function returns a function, not a variable or a string.
	x := 0
	return func() int { // The closure of wrapper function start here
		x++
		return x
	} // and ends here
}

func wrapper2() int {
	y := 0
	{
		y++
		return y
	}
}

func main() {

	increment := wrapper() // increment is assigned to this function
	// Nota => We could have written it the following way:
	// var increment func() int
	// increment = wrapper()
	// But the "shorthand" is better to read
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2 => x is not reset to 0 because it is out of the scope of "return func() int"

	increment2 := wrapper2()
	fmt.Println(increment2) // 1
	fmt.Println(increment2) // 1

}
