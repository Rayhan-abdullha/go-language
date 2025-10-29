package main

import "fmt"

func sum(num int, num1 int) {
	res := num + num1
	fmt.Println(res)
}
func add(a int, b int) int {
	res := a + b
	return res
}
func getNumbers(num int, num1 int) (int, int) {
	sum := num + num1
	mul := num * num1
	return sum, mul
}
func main() {
	// sum(10, 20)
	// res := add(2, 4)
	// fmt.Println(res)
	res1, res3 := getNumbers(10, 20)
	fmt.Println(res1, res3)
}
