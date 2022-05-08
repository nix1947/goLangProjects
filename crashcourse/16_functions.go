package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func main() {
	// functions are the building blocks
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)

	fmt.Print(kelvin, " k is", celsius, "C")
}
