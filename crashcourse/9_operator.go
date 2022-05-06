package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Go don't support preincrement operator

	var age = 41
	age = age + 1
	age += 1
	age++

	// Generate random num between 1 and 9
	var randomNumber = rand.Intn(10)
	fmt.Printf("The random number generated is %v", randomNumber)

	rand.Seed(12)
	var anotherRandomNum = rand.Intn(20)
	fmt.Printf("Another random number with is %v\n", anotherRandomNum)

}
