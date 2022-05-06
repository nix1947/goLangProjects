package main

import "fmt"

func main() {
	// Default values
	var price float64
	// default value is zero
	fmt.Println(price)

	third := 1.0 / 3

	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	//3f, 3 digit after decimal
	fmt.Printf("%.3f\n", third)
	//2f indicate: 2 digit after decimal
	//%4 indicate width, which indicate, minimum number
	// of character to display
	fmt.Printf("%4.2f\n", third)
	// Padded with space  2 digit after decimal
	fmt.Printf("%10.2f\n", third)
	//To minimize rounding errors, we recommend that you perform multiplication before
	// division.

}
