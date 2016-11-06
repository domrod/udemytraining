package main

import (
	"fmt"
	// Importing a self-defined package
	"github.com/domrod/udemytraining/golang/03_package/stringutil"
)

func main() {

	fmt.Println(stringutil.Reverse("!oG olleH"))
	fmt.Println("My name is:", stringutil.MyName)

}
