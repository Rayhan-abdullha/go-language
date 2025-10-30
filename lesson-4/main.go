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
	printNmber(100)
}

func printNmber(num int) {
	fmt.Println(num)
}
