package main

import "fmt"

func main() {
	// Define a closure that adds a specific number to its input
	add := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}

	// Create a closure that adds 5
	add5 := add(5)
	// Use the closure to add 5 to different numbers
	fmt.Println(add5(10))
	fmt.Println(add5(20))
	fmt.Println(add5(30))
}
