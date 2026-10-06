package main

import "fmt"

func main() {
	age := 23
	name := "arkan"

	// Print
	fmt.Print("hello, ")
	fmt.Print("World!\n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello world")
	fmt.Println("goodbye world")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v\n", age, name)
	fmt.Printf("my age is %q and my name is %q\n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %0.1f points!\n", 255.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println("the saved string is:", str)
}
