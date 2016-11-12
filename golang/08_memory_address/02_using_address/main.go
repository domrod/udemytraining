package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {

	var meters, yards float64
	fmt.Print("Enter distance in meters:")
	fmt.Scan(&meters) // NOTA => Scan puts the value read from STDIN at the MEMORY ADDRESS of the used variable!
	yards = meters * metersToYards
	fmt.Println("Distance in yards is:", yards)

}
