package main

import (
	"fmt"
	"time"
)

func calc(num int) {
	for i := 0; i < 10; i++ {
		fmt.Println("calc = ", i+num)
	}
}
func calc1(num int) {
	for i := 0; i < 10; i++ {
		fmt.Println("calc1 = ", i)
	}
}

func main() {
	go calc(1)
	go calc1(3)

	time.Sleep(5 * time.Second)
}
