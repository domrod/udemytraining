package main

import (
	"fmt"
	// Importing a self-defined package
	"github.com/domrod/udemytraining/golang/03_package/stringutil"
)

func main() {

	// Imported function (visible since is has a upper case first letter)
	fmt.Println(stringutil.Reverse("!oG olleH"))
	// Imported variable (visible since is has a upper case first letter)
	fmt.Println("My name is:", stringutil.MyName)

}
