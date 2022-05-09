package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 273.5
}

// Passing a function inside function
// func_name, func(), return_type should declared
func measureRealTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v K \n", k)
		time.Sleep(time.Second)
	}
}

// Using type function to short the code
type sensor func() kelvin

// declaring function type with row argument
// and two string return type
type getRowFn func(row int) (string, string)

// sensor type can used in function signature below.
func measureRealTemp(samples int, myFunc sensor) {
	return
}

// use of getRowFn
func drawTable(sample int, myFunc getRowFn) {

}

func main() {
	// Assign function to variable
	// Pass function to function
	// write function that create function

	// Assigning function to variable
	sensor := fakeSensor

	fmt.Println(sensor())

	// Again assigning a func to a variable
	sensor = realSensor
	fmt.Println(sensor())

	//  Pass function to function
	measureRealTemperature(3, fakeSensor)

	// Declearing function types
	// Possible to declare a new type for function
	// Just declare the function and assign to variable a
	// and use that variable as a type

	sensorFunction := realSensor // sensorFunction is realSensor type function
	fmt.Println(sensorFunction())

}
