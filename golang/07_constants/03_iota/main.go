package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	d = iota // 0
	e        // 1
	f        // 2
)

const (
	_ = iota      // 0
	h = iota * 10 // 10
	i = iota * 20 // 40
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 < (1 * 10) 10 shift of 1 in binary from column 1
	MB = 1 << (iota * 10) // 1 << (2 * 10) 20 shift of 1 in binary from column 1
	GB = 1 << (iota * 10) // 1 << (3 * 10) 30 shift of 1 in binary from column 1
)

func main() {

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)

	fmt.Println("h:", h)
	fmt.Println("i:", i)

	fmt.Printf("KB - binary: %b \n", KB)
	fmt.Printf("KB - decimal: %d \n", KB)
	fmt.Printf("MB - binary: %b \n", MB)
	fmt.Printf("MB - decimal: %d \n", MB)
	fmt.Printf("GB - binary: %b \n", GB)
	fmt.Printf("GB - decimal: %d \n", GB)
}
