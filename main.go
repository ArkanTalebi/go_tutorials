package main

import "fmt"

var someName = "hello"

// otherName := "World"

func main() {
	// strings
	var nameOne string = "Arkan"
	var nameTwo = "Majid"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Ali"
	nameThree = "Matin"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"

	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 127
	var numTwo int8 = -128
	var numThree uint8 = 255

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = -1.15
	var scoreTwo float64 = 12450987234112.7
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
