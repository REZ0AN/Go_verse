package main

import "fmt"

func main() {
	// Output Formatting

	// Println print with a newline
	fmt.Println("Hello, World!")

	// Printf print with a format
	fmt.Printf("Hello, %s!\n", "World")

	// Sprintf return a formatted string
	s := fmt.Sprintf("Hello, %s!", "World")

	// Errorf print with a format and return as error
	err := fmt.Errorf("hello, %s", "world")

	fmt.Println(s, err)

	// Print print without newline
	fmt.Print("Hello, World!")

	// Fprint print to a file
	// fmt.Fprint(file, "Hello, World!")

	// Fprintf print to a file with a format
	// fmt.Fprintf(file, "Hello, %s!", "World")

	// Fprintln print to a file with a newline
	// fmt.Fprintln(file, "Hello, World!")

	// Printf format specifiers
	// %v the value in a default format
	// %+v the value in a Go-syntax format
	// %#v a Go-syntax representation of the value
	// %T a Go-syntax representation of the type of the value
	// %t the word true or false
	// %b the base 2 representation of the value
	// %c the character represented by the value
	// %d the base 10 representation of the value
	// %o the base 8 representation of the value
	// %O the base 8 representation of the value with 0o prefix
	// %q a double-quoted character literal safely escaped with Go syntax
	// %x the base 16 representation of the value, with lower-case letters for a-f
	// %X the base 16 representation of the value, with upper-case letters for A-F
	// %U the Unicode format: U+1234; same as "U+%04X"
	// %e the scientific notation of the value
	// %E the scientific notation of the value with an upper-case E
	// %f the decimal point but no exponent, e.g. 123.456
	// %F synonym for %f
	// %g the shortest representation of the value: %e for large exponents, %f otherwise

	fmt.Printf("Value: %v\n", 42)
	fmt.Printf("Go-syntax: %+v\n", 42)
	fmt.Printf("Go-syntax: %#v\n", 42)
	fmt.Printf("Type: %T\n", 42)
	fmt.Printf("Boolean: %t\n", true)
	fmt.Printf("Base 2: %b\n", 42)
	fmt.Printf("Character: %c\n", 42)
	fmt.Printf("Base 10: %d\n", 42)
	fmt.Printf("Base 8: %o\n", 42)
	fmt.Printf("Base 8: %O\n", 42)
	fmt.Printf("Quoted: %q\n", 42)
	fmt.Printf("Base 16: %x\n", 42)
	fmt.Printf("Base 16: %X\n", 42)
	fmt.Printf("Unicode: %U\n", 42)
	fmt.Printf("Scientific: %e\n", 42.0)
	fmt.Printf("Scientific: %E\n", 42.0)
	fmt.Printf("Decimal: %f\n", 42.0)
	// variations of %f
	fmt.Printf("Decimal: %f\n", 42.0)
	fmt.Printf("Decimal: %6.4f\n", 42.043434334)
	fmt.Printf("Decimal: %2.2f\n", 432.03433)
	fmt.Printf("Decimal: %5.2f\n", 42.0)
	fmt.Printf("Decimal: %F\n", 42.0)
	fmt.Printf("Shortest: %g\n", 42.0)

}
