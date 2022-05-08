package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	peace := "peace"
	var country = "Nepal"
	var year string = "2002"
	fmt.Printf("%v has been declerated a %v country in %v", country, peace, year)

	// Multiline string
	fmt.Println(`
		This is a multiple line string that 
		can span multiple lines with 
		\n escape sequences
	`)

	fmt.Printf("%v and printed using escape character", `I am mul
	ti \n line string\n`)

	// Character code points, rune, bytes
	// Numeric value: code points assigned by Unicode consortium
	// Code points are represented by rune
	// rune: is an alias for the int32 type
	// byte: is an alias for uint8 type

	var pi rune = 960 // 960 is a pi
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	// This print value
	fmt.Printf("%v, %v, %v, %v\n", pi, alpha, omega, bang)

	// To print character use %c
	fmt.Printf("%c, %c, %c, %c\n", pi, alpha, omega, bang)

	// Casting to rune or numeric value
	var grade rune = 'A'
	var star byte = '*'
	fmt.Printf("%v\n", grade) // 65
	// Character literals can also be used with the byte alias:
	fmt.Printf("%v\n", star) // 97

	// Accessing string items.

	// How many characters does ASCII encode?
	// 128 charactres

	message := "Shalom"
	c := message[5]
	fmt.Printf("%c\n", c)

	for i := 0; i < len(message); i++ {
		character := message[i]
		fmt.Printf("%c\n", character)
	}

	// Decoding unicode string

	var question = "¿Cómo estás?"
	fmt.Println(len(question), "Length in bytes")
	fmt.Println(utf8.RuneCountInString(question), "total runes in question string")
	var dataType, size = utf8.DecodeRuneInString(question)
	fmt.Printf("First rune is: %c which is taking %v bytes in memory\n", dataType, size)

	// Looping using range
	fmt.Println("Using range function in string")
	for index, char := range question {
		fmt.Println(index, string(char))
	}

	// Calculating rune and its bytes in english Alphabets
	var englishAlphabets = "abcdefghijklmnopqrstuvwxyz¿"

	for _, alphabet := range englishAlphabets {
		var _, size = utf8.DecodeLastRuneInString(string(alphabet))
		fmt.Printf("rune character :%c, rune value:%v, consumedByte: %v\n", alphabet, alphabet, size)
	}
}
