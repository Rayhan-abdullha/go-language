package main

import "fmt"

func add(x int, y int, fn func(a int, b int) int) int {
	return fn(x, y)
}
func cb(a int, b int) int {
	return a + b
}
func main() {
	res := add(1, 2, cb)
	fmt.Println(res)
}
