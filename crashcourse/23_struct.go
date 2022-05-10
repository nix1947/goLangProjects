package main

import "fmt"

//
// Structures group related values
//  together, making it simpler
// and less error-prone to pass them  around

// Structured are copied. which make the member
// variable value change independently
// Assignment makes a copy
// Field may be listed in any order
// Fields are optional taking on the zero value if not listed.
// No change are required when reordering or adding fields
// to the structure declaration

type location struct {
	lat  float64
	long float64
}

func main() {

	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = 12.3
	curiosity.long = 13.0

	fmt.Println(curiosity.lat)
	fmt.Println(curiosity.long)

	opportunity := location{lat: 12.3, long: 14.3}
	fmt.Println(opportunity)
}
