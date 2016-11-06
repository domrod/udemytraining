package main

import (
	"fmt"
)

func main() {
	const limit int = 200
	for i := 0; i < limit; i++ {
		fmt.Printf("Decimal: %d \t - Binary: %b \t - Hexadecimal: %#X \n", i, i, i)
	}
}
