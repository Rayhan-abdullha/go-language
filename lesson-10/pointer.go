package main

import "fmt"

// pass by reference
func pointer1(arr1 *[3]int) {
	arr1[0] = 1000
	fmt.Println(arr1)
}
func main() {
	arr := [3]int{1, 2, 4}
	pointer1(&arr)
	fmt.Println(arr)

	// x := 20
	// p := &x
	// println("x = ",x)
	// *p = 30
	// println("Address: ", p)
	// println("Value of pointer address - ", *p)
	// println("x = ", x)
}
