package main

import "fmt"

func main() {
	// Declare a single variable
	var a int
	a = 10
	fmt.Println("a:", a)

	// Declare and initialize a single variable
	var b int = 20
	fmt.Println("b:", b)

	// Type inference
	var c = 30
	fmt.Println("c:", c)

	// Short variable declaration with a single variable
	d := int(40) // implicit type is int
	fmt.Println("d:", d)

	// Declare multiple variables
	var e, f, g int
	e, f, g = 50, 60, 70
	fmt.Println("e:", e, "f:", f, "g:", g)

	// Declare and initialize multiple variables
	var h, i, j int = 80, 90, 100
	fmt.Println("h:", h, "i:", i, "j:", j)

	// Mixed variable types
	var k, l, m = 110, "hello", 3.14
	fmt.Println("k:", k, "l:", l, "m:", m)

	// Short variable declaration with multiple variables
	n, o, p := 120, "world", 6.28
	fmt.Println("n:", n, "o:", o, "p:", p)
}
