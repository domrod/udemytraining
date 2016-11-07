package main

import (
	"fmt"

	"github.com/domrod/udemytraining/golang/05_scope/01_package_scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName, "(printed directly with the variable)")
	vis.Printnames()
}
