package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Print(i, "")
	}
	fmt.Println()
	// while loop
	j := 1
	for j <= 10 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()
	// do while loop
	k := 1
	for {
		fmt.Print(k, " ")
		k++
		if k > 10 {
			break
		}
	}
	fmt.Println()
	// for loop with continue
	for l := 1; l <= 10; l++ {
		if l%2 == 0 {
			continue
		}
		fmt.Print(l, " ")
	}
	fmt.Println()
	// for loop with break
	for m := 1; m <= 10; m++ {
		if m > 5 {
			break
		}
		fmt.Print(m, " ")
	}
	fmt.Println()
	// for loop with nested loop
	for n := 1; n <= 5; n++ {
		for o := 1; o <= 5; o++ {
			fmt.Print("( ", n, ", ", o, " ) ")
		}
		fmt.Println()
	}
	fmt.Println()
}
