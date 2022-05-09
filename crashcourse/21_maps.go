package main

import (
	"fmt"
	"sort"
)

func main() {
	// use maps for unstructured data
	// Declare, access and iterate over maps
	// Explore some use of the versatilemap

	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the earth is%v \n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 446

	// Accessing a key that doesn't exist in the map
	// Return 0

	moon := temperature["moon"]
	fmt.Printf("value of moon key is %v", moon)

	temperature["moon"] = -10

	if moon, ok := temperature["moon"]; ok {
		fmt.Printf("The temp of moon is %v\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}

	// Maps aren't copied.
	// Primitive types are array are copied while passing to a function
	// or assigning to a variable
	// Maps aren't copied.
	// Use trailing comma aswell.
	// They are referenced.

	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	planetsMarkII := planets

	planets["Earth"] = "Whoops"
	fmt.Println(planetsMarkII["Earth"])

	// Deleting a key
	delete(planets, "Earth")

	fmt.Println("The value is deleted", planets["Earth"])

	// delete built-in function removes an element from the map
	// This behavior is similar to multiple slices that point to
	// the same underlying array.

	// Creating maps with make
	// Adv is make, is preallocate the space with last parameter size.

	myMap := make(map[float64]int, 8)
	mySlice := make([]string, 4)
	mySlice = append(mySlice, "apple", "orange", "pear", "banana")

	fmt.Println(myMap)
	fmt.Println(mySlice)

	// Sorting the value in array
	sort.StringSlice(mySlice).Sort()
	fmt.Println("Value of sorting is", mySlice)

	// Using maps to count things.

	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	for _, t := range temperatures {
		frequency[t]++

	}

	for t, num := range frequency {
		fmt.Printf("%+.2f occours %d times\n", t, num)
	}

}
