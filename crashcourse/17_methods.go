package main

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsius Converter

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	// methods are like function with additional behaviour.
	// use type keyword to declares a new type with name and underlying type.

	const degrees = 20
	var k kelvin = 294.3
	fmt.Println("Converting to celsius is")
	c := kelvinToCelsius(k)
	fmt.Print(k, "K is", c, "C is")

	// Here celsius type is an unique type
	var temperature celsius = degrees
	temperature += 10
	fmt.Println(temperature)

	var warmUp float64 = 10
	// INVALID OPERATION
	// temperature += warmUp
	temperature += celsius(warmUp)
	fmt.Println(temperature)

}
