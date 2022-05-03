// Go program made of packages
// Program start to run from `main` package

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Setting the seed value to generate random number
	rand.Seed(100)
	randNum := rand.Intn(12)
	rand.Seed(12)
	anotherRandomNumber := rand.Intn(12)
	fmt.Println(randNum)
	fmt.Println(anotherRandomNumber)

}
