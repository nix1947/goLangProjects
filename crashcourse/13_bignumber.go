package main

import (
	"fmt"
	"math/big"
)

// Use big package

func main() {

	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(864000)
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away.")
	seconds := new(big.Int)
	// Use Division method
	seconds.Div(distance, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println("That is", days, "days of travel at light speed.")
	// What are two ways to make a big.Int with the number 86,400?
	//	Construct a big.Int with the NewInt function:
	secondsPerDay = big.NewInt(86400)
	// Or use the SetString method:
	secondsPerDay1 := new(big.Int)
	secondsPerDay.SetString("86400", 10)
	fmt.Println(secondsPerDay, secondsPerDay1)

}
