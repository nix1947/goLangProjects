package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Convert btn numeric, string, boolean types.
	// type coercing
	// Conversion between types is explicit to avoid ambuiguity
	// strconv package provides function for converting string to and from other types.
	age := 41
	marsAge := float64(age)
	fmt.Println(marsAge)

	// string conversion

	var pi rune = 960
	fmt.Println(string(pi))

	// Integer to Ascii, integer to ascii conversion
	var countdown = 10
	fmt.Println(strconv.Itoa(countdown))

	// Ascii to integer package
	ascii, err := strconv.Atoi("500")
	fmt.Printf("Integer value of ASCII is %v", ascii)

	if err != nil {
		fmt.Println("Something went wrong")
	}

	// USE strconv.IOTA and Sprintf to convert number to a string

	// Convert number to string
	str := fmt.Sprintf("%v", countdown)
	fmt.Println(str)

	// converting boolean values.
	fmt.Println("#####Boolean to String#######")
	launch := false
	launchToString := fmt.Sprintf("%v", launch)
	fmt.Println(launchToString)

	var yesNo string

	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}

	fmt.Println("Ready for launch ", yesNo)

	// Converting a string to boolean
	fmt.Println("#####Converting a string to a boolean###")
	var boolean string = "no"
	var convertedBooleanValue bool = false
	convertedBooleanValue = (boolean == "no")
	fmt.Println(convertedBooleanValue)

	var stringValue string = "true"
	var stringToBool bool = true

	stringToBool = (stringValue == "true")
	stringToBool = (stringValue == "yes")
	stringToBool = (stringValue == "1")

	fmt.Println(stringToBool)

	// OR using switch statement

	switch stringValue {
	case "true", "yes", "1":
		stringToBool = true
	}

	// Handling out of range while type conversion
	var bh float64 = 33767
	var h = int16(bh) // overflow err
	fmt.Println("Overflow should occurred, but go in it reset back to starting value", h)

	if bh < math.MinInt16 || bh > math.MaxInt16 {
		// Can be only casted bh to int16
		h = int16(h)
		fmt.Println("Casted without error", h)
	}

}
