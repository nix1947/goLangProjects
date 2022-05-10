package main

import (
	"fmt"
	"sort"
)

// Set is a collection similar to an array
// Except that each element is guranteed to occour only once.
// Go doesn't provide set collection
// But can implemented usng map

func main() {

	var temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}

	fmt.Println(set)

	if set[-28] {
		fmt.Println("Set member")
	}

	// sort the sets,
	// To sort you need to convert set to slice.

	// make a slice having length of set
	unique := make([]float64, 0, len(set))

	for t := range set {
		unique = append(unique, t)
	}

	// Sort the unique
	sort.Float64s(unique)
	fmt.Println(unique)
}
