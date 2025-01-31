package main

import "fmt"

func main() {

	// if else block example
	var a, b int
	a = 10
	b = 20
	if a > b {
		fmt.Println("a is greater than b by", a-b)
	} else {
		fmt.Println("b is greater than a by", b-a)
	}

	// if else if block example
	var c, d int
	c = 10
	d = 10

	// esle and esle if block must be in the same line of end { of if block
	if c > d {
		fmt.Println("c is greater than d by", c-d)
	} else if c < d {
		fmt.Println("d is greater than c by", d-c)
	} else {
		fmt.Println("c is equal to d")
	}

	// nested if block example
	var e, f, g int
	e = 10
	f = 20
	g = 30
	if e > f {
		if e > g {
			fmt.Println("e is greater than f and g")
		} else {
			fmt.Println("g is greater than e and f")
		}
	} else {
		if f > g {
			fmt.Println("f is greater than e and g")
		} else {
			fmt.Println("g is greater than e and f")
		}
	}

	// switch block example with character
	var ch rune = 'e' // rune is an alias for int32

	// no need to use break statement in go lang
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("Vowel")
	case 'y':
		fmt.Println("Weird Vowel")
	default:
		fmt.Println("Consonant")
	}
	// using byte you have to type cast the character
	var ch1 byte = 'y'

	switch ch1 {
	case byte('a'), byte('e'), byte('i'), byte('o'), byte('u'):
		fmt.Println("Vowel")
	case byte('y'):
		fmt.Println("Weird Vowel")
	default:
		fmt.Println("Consonant")
	}
}
