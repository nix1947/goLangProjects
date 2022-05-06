package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("You find yourself in a dimly lit caveran.")
	var command = ("Walk outside")
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)

	// Logical operators
	var room = "cave"

	if room == "cave" {
		fmt.Println("You find yourself in a dimly lit caveran.")
	} else if room == "entrance" {
		fmt.Println("There is a averan entrance here and a path to the east.")
	} else if room == "mountain" {
		fmt.Println("There is a cliff here. A path leads west down the mountain.")
	} else {
		fmt.Println("Everything is white")
	}

	// Loops
	var count = 10
	for count > 1 {
		fmt.Println("Currently counting", count)
		time.Sleep(time.Second)
		count--
	}

	// Infinite loop
	var degrees = 0

	for {
		fmt.Println(degrees)
		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}

		}
	}

	// Guessing program in go
	var counting = 10
	for counting > 0 {
		fmt.Println(counting)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		counting--
	}

	if count == 0 {
		fmt.Println("Litoff!")
	} else {
		fmt.Println("Launched failed")
	}

	// var count = 10 is equivalent

	var count1 = 0
	for count1 = 10; count1 > 0; count1-- {
		fmt.Println("Count1", count1)

	}
	fmt.Println("Count1", count1)

	// Short declaration of IF statement
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactice")
	} // Here num is no longer in scope

	// Short declaration of switch
	fmt.Println("Running short switch block")
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")

	case 1:
		fmt.Println("SpaceX")

	case 2:
		fmt.Println("Virgin Galactice")
	default:
		fmt.Println("Random spaceline #", num)
	} // Here no scoping of num

	// type float64 is double precision
	// All the function in math package use double precision
	// float32 is a single precision

}
