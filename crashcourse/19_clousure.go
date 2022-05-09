package main

import "fmt"

// Anonymous function also called function literal
// function without a name are anonymous function

// Function literals are closures
// variable having function value are closures
// Assigning a anonymous function to a variable
// and then use that variable like any other function
// Closures keep references to variables in the surrounding scope.

func main() {
	// Declaring anonymous function
	var f = func() {
		fmt.Println("Anonymous function")
	}

	f()

	// Declare anonymous function and call on time
	func(name string) {
		fmt.Println("Your name is", name)
	}("Manoj Gautam")

	//
}
