package main

import "fmt"

type contact struct {
	title     string
	firstname string
	lastname  string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // This is an asserts. It checks the type of x.
	case int:
		fmt.Printf("%v is of type int \n", x)
	case string:
		fmt.Printf("%v is of type string \n", x)
	case contact:
		fmt.Printf("%v is of type contact \n", x)
	default:
		fmt.Printf("%v is of an unknown type \n", x)
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Dom")

	var contact1 = contact{"M.", "John", "Doe"} // Note the setup in "JSON" format
	SwitchOnType(contact1)
	SwitchOnType(contact1.title)
	SwitchOnType(contact1.firstname)
	SwitchOnType(contact1.lastname)

	// var contact2 contact
	// contact2.title = "Mrs."
	// contact2.firstname = "Jenny"
	// contact2.lastname = "Smith"
	contact2 := contact{"Mrs.", "Jenny", "Smith"}
	fmt.Printf("Last contact entered is %#v\n", contact2)

}
