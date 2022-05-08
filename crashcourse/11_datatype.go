package main

import (
	"fmt"
	"math"
)

func main() {
	// Default values
	var price float64
	// default value is zero
	fmt.Println(price)

	// Type Range Storage
	// int8: –128 to 127 8-bit (one byte) uint8 0 to 255
	// int16: –32,768 to 32,767 16-bit (two bytes)
	// uint16: 0 to 65535
	// int32: –2,147,483,648 to 2,147,483,647 32-bit (four bytes)
	// uint32: 0 to 4,294,967,295
	// int64 –9,223,372,036,854,775,808 to  9,223,372,036,854,775,807 64-bit (eight bytes)
	// uint64 0 to 18,446,744,073,709,551,615

	var year = 2018
	fmt.Printf("Type of year is %T and value is %[1]v \n", year, year)

	a := "text"
	fmt.Printf("Type of a is %T for %[1]v\n", a)

	b := 42
	fmt.Printf("Type of b is %T for %[1]v\n", b)

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
	//To minimize rounding errors, it is recommend that you perform multiplication before division.

	// 8-bit integer has a range of 0-255
	var red uint8 = 255
	red++ // overflow and it wrap to zero
	fmt.Printf("The overflow value of red is %v\n", red)

	var number int8 = 127
	number++
	// %b gives binary value
	// display in 8bit
	fmt.Printf("The overflow value %08b of number is %v\n", number, number)

	const max_int_value = math.MaxInt
	fmt.Printf("The maximum int value from math.maxInt is%v\n", max_int_value)

}
