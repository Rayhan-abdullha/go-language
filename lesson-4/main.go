package main

import (
	"fmt"

	mathLib "example.com/math"
)

var (
	a = 10
	b = 20
)

func main() {
	if a < b {
		mathLib.Sum(a, b)
	} else {
		fmt.Print("A does not greater than b")
	}

	fmt.Print(mathLib.C)
	printNumber(100)
}

func printNumber(num int) {
	fmt.Println(num)
}
