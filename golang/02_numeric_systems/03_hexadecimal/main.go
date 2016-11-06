package main

import (
	"fmt"
)

func main() {
	// Nota: For hexadecimal, 2 ways => %x or %#x (with 0x (%#x) or 0X (%#X) notation)
	// fmt.Printf("Decimal: %d - Binary: %b - Hexadecimal: %x \n", 42, 42, 42)
	fmt.Printf("Decimal: %d \t - Binary: %b \t - Hexadecimal: %#X \n", 42, 42, 42)
}
