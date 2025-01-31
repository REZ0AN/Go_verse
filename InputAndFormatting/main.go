package main

import "fmt"

func main() {
	// Declare a variable to store the user's name
	var name string

	// Ask the user for their name
	fmt.Println("What is your name?")
	fmt.Scan(&name)

	// Print a greeting to the user
	fmt.Println("Hello", name)

	// Declare a variable to store the user's age
	var age int
	fmt.Print("How old are you? => ")
	fmt.Scan(&age)

	// Print the user's age
	fmt.Println("You are", age, "years old")

	// Input a string first then an integer in a single scanf
	var name2 string
	var age2 int
	fmt.Print("Enter your name and age: ")
	// use Scanf to input a string and an integer
	fmt.Scanf("%s %d", &name2, &age2)
	fmt.Println("Hello", name2, "you are", age2, "years old")
	// explain how go lang handle flushing the buffer when using Scanf
	// Scanf does not flush the buffer, so it will read the newline character from the previous input
	// To fix this, use a buffer flusher like Scanln
	// fmt.Scanln()
	// fmt.Scanf("%s %d", &name2, &age2)
	// fmt.Println("Hello", name2, "you are", age2, "years old")
	// fmt.Scanln()
	// fmt.Scanf("%s %d", &name2, &age2)

	fmt.Scanf("%s", &name2)
	fmt.Scanf("%d", &age2)

	fmt.Println("Hello", name2, "you are", age2, "years old")
}
