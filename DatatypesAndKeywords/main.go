package main

import "fmt"

func main() {
	// basic keywords
	// break, default, func, interface, select, case, defer, go, map, struct, chan, else, goto, package, switch, const, fallthrough, if, range, type, continue, for, import, return, var
	// chan is used to communicate between goroutines
	// defer is used to delay the execution of a function until the surrounding function returns
	// go is used to start a goroutine
	// interface is used to define a set of methods that a type must implement
	// map is used to create a map
	// package is used to create a package, which is a way to group related functions and types together
	// struct is used to create a new struct, which is a collection of fields
	// type is used to create a new type
	// var is used to declare a variable

	// Data Types and Keywords
	// Go is a statically typed language, which means that the type of a variable is known at compile time.
	// Go has the following built-in data types:
	// bool
	var a bool = true
	fmt.Println("Blooean Datatype example", a, " is of type", fmt.Sprintf("%T", a))
	// string
	var b string = "Hello, World!"
	fmt.Println("String Datatype example", b, " is of type", fmt.Sprintf("%T", b))
	// int, int8, int16, int32, int64
	var c int = 42 // int is 32 or 64 bits depending on the platform
	fmt.Println("Integer Datatype example", c, " is of type", fmt.Sprintf("%T", c))
	// caluation limits
	// int8: -128 to 127 "2^7 to 2^7-1
	// int16: -32768 to 32767 "2^15	to 2^15-1"
	// int32: -2147483648 to 2147483647 "2^31 to 2^31-1"
	// int64: -9223372036854775808 to 9223372036854775807 "2^63 to 2^63-1"

	// uint, uint8, uint16, uint32, uint64, uintptr
	var d uint = 42 // uint is 32 or 64 bits depending on the platform
	fmt.Println("Unsigned Integer Datatype example", d, " is of type", fmt.Sprintf("%T", d))
	// caluation limits
	// uint8: 0 to 255 "0 to 2^8-1"
	// uint16: 0 to 65535 "0 to 2^16-1"
	// uint32: 0 to 4294967295 "0 to 2^32-1"
	// uint64: 0 to 18446744073709551615 "0 to 2^64-1"
	// uintptr: an unsigned integer large enough to store the uninterpreted bits of a pointer value

	// byte (alias for uint8)
	var e byte = 42
	fmt.Println("Byte Datatype example", e, " is of type", fmt.Sprintf("%T", e))
	// rune (alias for int32)
	var f rune = 42
	fmt.Println("Rune Datatype example", f, " is of type", fmt.Sprintf("%T", f))
	// float32, float64
	var g float32 = 3.14 // float32 is 32 bits
	fmt.Println("Float Datatype example", g, " is of type", fmt.Sprintf("%T", g))
	var h float64 = 3.14 // float64 is 64 bits
	fmt.Println("Float Datatype example", h, " is of type", fmt.Sprintf("%T", h))
	// caluation limits
	// float32: 1.18E-38 to 3.4E+38
	// float64: 2.23E-308 to 1.80E+308

	// complex64, complex128
	var i complex64 = 1 + 2i // complex64 is 32 bits
	fmt.Println("Complex Datatype example", i, " is of type", fmt.Sprintf("%T", i))
	var j complex128 = 1 + 2i // complex128 is 64 bits
	fmt.Println("Complex Datatype example", j, " is of type", fmt.Sprintf("%T", j))
	// caluation limits
	// complex64: real and imaginary parts are float32
	// complex128: real and imaginary parts are float64

	// character type
	var ch rune = 'a'
	fmt.Println("Character Datatype example", ch, " is of type", fmt.Sprintf("%T", ch))
	// caluation limits
	// rune: represents a Unicode code point
	// rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.
	// The type is used in Go to represent a Unicode code point, which is a 32-bit representation of a Unicode character.
	// character type with byte
	var ch1 byte = 'a'
	fmt.Println("Character Datatype example", ch1, " is of type", fmt.Sprintf("%T", ch1))
	// byte: represents a ASCII code point
	// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is used, by convention, to distinguish byte values from 8-bit unsigned integer values.
	// The type is used in Go to represent ASCII characters.

	// string type
	var k string = "Hello, World!"
	fmt.Println("String Datatype example", k, " is of type", fmt.Sprintf("%T", k))

	// Pointer type
	var l int = 42
	var m *int = &l
	fmt.Println("Pointer Datatype example", m, " is of type", fmt.Sprintf("%T", m))

	// Array type
	var n [3]int = [3]int{1, 2, 3}
	fmt.Println("Array Datatype example", n, " is of type", fmt.Sprintf("%T", n))

	// Slice type
	var o []int = []int{1, 2, 3}
	fmt.Println("Slice Datatype example", o, " is of type", fmt.Sprintf("%T", o))

	// Map type
	var p map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}

	fmt.Println("Map Datatype example", p, " is of type", fmt.Sprintf("%T", p))

	// Struct type
	type Person struct {
		name string
		age  int
	}
	var q Person = Person{name: "John", age: 30}
	fmt.Println("Struct Datatype example", q, " is of type", fmt.Sprintf("%T", q))

	// Function type
	var t = func() {
		fmt.Println("Hello, World!")
	}
	fmt.Println("Function Datatype example", t, " is of type", fmt.Sprintf("%T", t))

	// Channel type
	var u chan int = make(chan int)
	fmt.Println("Channel Datatype example", u, " is of type", fmt.Sprintf("%T", u))

	// Nil type
	var v *int = nil
	fmt.Println("Nil Datatype example", v, " is of type", fmt.Sprintf("%T", v))

	// Error type
	var w error = nil
	fmt.Println("Error Datatype example", w, " is of type", fmt.Sprintf("%T", w))

	// Constants
	const x int = 42
	fmt.Println("Constant Datatype example", x, " is of type", fmt.Sprintf("%T", x))

	// Constants can be typed or untyped
	const y = 42
	fmt.Println("Constant Datatype example", y, " is of type", fmt.Sprintf("%T", y))

	// Constants can be shadowed
	const z int = 42
	{
		const z = 21
		fmt.Println("Constant Datatype in block example", z, " is of type", fmt.Sprintf("%T", z))
	}
	fmt.Println("Constant Datatype example", z, " is of type", fmt.Sprintf("%T", z))

	// iota
	const (
		aa = iota
		bb = iota
		cc = iota
	)
	fmt.Println("Iota Datatype example", aa, " is of type", fmt.Sprintf("%T", aa))
	fmt.Println("Iota Datatype example", bb, " is of type", fmt.Sprintf("%T", bb))
	fmt.Println("Iota Datatype example", cc, " is of type", fmt.Sprintf("%T", cc))

	// iota can be used in const blocks to simplify definitions of incrementing numbers
	const (
		dd = iota
		ee
		ff
	)
	fmt.Println("Iota Datatype example", dd, " is of type", fmt.Sprintf("%T", dd))
	fmt.Println("Iota Datatype example", ee, " is of type", fmt.Sprintf("%T", ee))
	fmt.Println("Iota Datatype example", ff, " is of type", fmt.Sprintf("%T", ff))

}
