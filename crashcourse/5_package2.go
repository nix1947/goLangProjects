package main

import "fmt"

func main() {
	name := "manoj"

	charArray := []rune(name)

	for i := 0; i < len(charArray); i++ {
		// Converting rune to String data type
		fmt.Print(string(charArray[i]))
	}

}
