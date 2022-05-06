package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "hello"
	f := `Do you like my hat?`
	g := 'M'

	// V for variables
	// go doc fmt.Println
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	// -15v create 15 space right
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	// The lightSpeed constant can’t be changed. If you try to assign it a new value, the Go
	// compiler will report the error “cannot assign to lightSpeed.”
	const lightSpeed = 299792 // Km/sec
	var distance = 56000000   // km
	fmt.Println(distance/lightSpeed, "seconds")

	// Declaring multiple variables at once

	var (
		firstName = "manoj"
		lastName  = "gautam"
	)

	fmt.Printf("Your first name is %v and lastname is %v", firstName, lastName)

}
