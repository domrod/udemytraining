package main

import (
	"fmt"
)

func main() {
	const limit = 256
	for i := 0; i < limit; i++ {
		fmt.Printf("Decimal: %d \t Binary: %b \t Hexadecimal: %#x \t UTF8: %q \n", i, i, i, i)
	}

}
