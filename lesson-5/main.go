package main

import "fmt"

var a = 10

func main() {
	// age := 20
	// if age > 18 {
	// 	a := 30
	// 	fmt.Println("Inside if block, a =", a) // This will print 30
	// }
	print(a)

	func(a, b int) {
		print(a + b)
	}(3, 3)
}

func print(val any) {
	fmt.Println(val)
}

func init() {
	print(a)
	a = 30
}
