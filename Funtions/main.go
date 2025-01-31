package main

import "fmt"

func area(a, b int) int {
	return a * b
}

func changeValue(a int) {
	a = 20
	fmt.Println("Value of x inside function : ", a)
}

func changeValueRef(a *int) {
	*a = 20
	fmt.Println("Value of x inside function : ", *a)
}

func changeArray(arr [3]int) {
	arr[0] = 10
	fmt.Println("Array inside function : ", arr)
}

func changeSlice(slice []int) {
	slice[0] = 10
	fmt.Println("Slice inside function : ", slice)
}

func returnArray() []int {
	var arr [3]int = [3]int{1, 2, 3}
	return arr[:]
}
func main() {
	var a, b int
	fmt.Print("Enter height , width : ")
	fmt.Scanln(&a, &b)
	fmt.Println("Area of rectangle is : ", area(a, b))

	// pass by value example
	x := 10
	fmt.Println("Value of x before calling function : ", x)
	changeValue(x)
	fmt.Println("Value of x after calling function : ", x)

	// pass by reference example
	y := 10
	fmt.Println("Value of y before calling function : ", y)
	changeValueRef(&y)
	fmt.Println("Value of y after calling function : ", y)

	// create an array and pass it to a function
	arr := [3]int{1, 2, 3}
	fmt.Println("Array before passing to function : ", arr)
	changeArray(arr)
	fmt.Println("Array after passing to function : ", arr)
	// arrays are passed by value default

	// to pass by reference use slices
	slice := []int{1, 2, 3}
	fmt.Println("Slice before passing to function : ", slice)
	changeSlice(slice)
	fmt.Println("Slice after passing to function : ", slice)

	// return array from a function
	arr1 := returnArray()
	fmt.Println("Array returned from function : ", arr1, fmt.Sprintf("Type : %T", arr1))
}
