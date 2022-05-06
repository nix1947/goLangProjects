package main

import "fmt"

func main() {
	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Look before your leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}

	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You head futher up the mountain.")
	// comma separated list of possible values
	case "go cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern")
		// fallthrough execute the remaining body
		fallthrough

	case "read sign":
		fmt.Println("The sign reads 'no minors'. ")
	default:
		fmt.Println("Didn't quite get that.")

	}
}
