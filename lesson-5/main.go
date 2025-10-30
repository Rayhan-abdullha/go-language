package main

import "fmt"

var a = 10

func main() {
	age := 20
	if age > 18 {
		a := 30
		fmt.Println("Inside if block, a =", a) // This will print 30
	}
	fmt.Println("Outside if block, a =", a) // This will print 10
}
