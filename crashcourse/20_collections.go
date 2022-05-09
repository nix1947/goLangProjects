package main

import (
	"fmt"
	"sort"
	"strings"
)

func hyperSpace(worlds []string) {
	// This function modify the worlds
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {

	// Defining array

	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	earth := planets[2]
	fmt.Println(earth)

	// initilize array
	// Trailing comma is required
	drawfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	for planet := 0; planet < len(drawfs); planet++ {
		fmt.Println(drawfs[planet])
	}

	// Printing using range
	// With the range keyword, the loop is simpler and avoids mistakes like going out of bounds (

	for i, planet := range drawfs {
		fmt.Println(i, planet)
	}

	// Passing an array to a function makes a complete copy of its contents
	planets2 := planets
	planets2[1] = "Whoops"
	fmt.Println(planets[1])
	fmt.Println(planets2[1])
	fmt.Println(drawfs)

	// Array without specifying the size.
	arr2 := [...]int{9, 7, 6, 4, 5, 3, 2, 4}
	fmt.Println(arr2)

	// Multidimensional arrays.
	// Arrays of arrays

	var board [8][8]string

	for i, _ := range board {
		for j, _ := range board {
			board[i][j] = "X"

		}
	}
	fmt.Println(board)

	// Slice an arrays
	terrestrial := planets[0:4]
	terrestrial1 := planets[:4] // index start from 0 to 3
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]
	allPlanets := planets[:] // no change

	fmt.Println(allPlanets, terrestrial, gasGiants, iceGiants, terrestrial1)

	// Modify the slice.
	iceGiants[1] = iceGiants[1] + "- Modified"
	fmt.Println(iceGiants[1])

	// Slices are more versatile than arrays in other ways too
	// Slices having any size can passed to function
	fruits := []string{"Apple   ", "Banana    "}
	fmt.Println(fruits)

	// This function modify the array and
	// Passing fixed size array throws error
	// hyperSpace(planets): throws error

	// creating slice
	sliceFruits := fruits[:]
	hyperSpace(sliceFruits)
	fmt.Println(sliceFruits)

	// Slice with methods.
	// with defined type, we can methods for that type.
	// methods can attached to the type.
	// sort package in the std lib declare a StringSlice type
	type StringSlice []string

	//Attach sort method in StringSlice

	// this is a slice
	colors := []string{"red", "blue", "green"}
	// colors = colors[:]

	sort.StringSlice(colors).Sort()
	fmt.Println(colors)

	// Appending in array
	// for appending array should be slice
	colors = append(colors, "purple")

	// can also pass multiple elements
	colors = append(colors, "magenta", "Aqua")

	// Dump slice length, capacity and contents
	dump := func(label string, slice []string) {
		fmt.Printf("%v: length %v, capacity %v %v \n", label, len(slice), cap(slice), slice)
	}

	fmt.Println("----Using, cap, len in array-------")
	dump("colors", colors)
	dump("colors[1:2]", colors[1:2])

	// Three index slicing
	// TODO:

	myFruits := []string{"apple", "banana", "orange", "Pear", "Strawberry", "watermelon", "Guava", "Mango"}

	myFavFruits := myFruits[0:4:4] // 4: length, 4: capacity
	myFavFruits = append(myFavFruits, "customFruit")
	fmt.Println(myFavFruits)

	// prealocate slice with make.
	// max size of myfavGames is 10 and it can't hold more size than 10
	myFavGames := make([]string, 0, 10)

	for i := 0; i < cap(myFavGames); i++ {
		myFavGames = append(myFavGames, string(i))
		fmt.Println(myFavGames)
	}

	// What is the benefit of making slices with make
	// Preallocating with make can set an initial capacity, thereby avoiding additional allocations and copies to enlarge the underlying array

	// Declaring variadic function
	// Printf and append are variadic function as they accept a variable number of arguments.
	// ...string indicate variable number of argument.
	terraform := func(prefix string, worlds ...string) []string {
		newWorlds := make([]string, len(worlds))
		for i := range worlds {
			newWorlds[i] = prefix + " " + worlds[i]
		}

		return newWorlds
	}

	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)

	ourPlanets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", ourPlanets...)

	fmt.Print(newPlanets)

}
