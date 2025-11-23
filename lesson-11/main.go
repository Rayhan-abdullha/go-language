package main

import (
	"fmt"
)

func main() {
	// ---------- BASIC DATA TYPES ----------

	// Boolean
	var isActive bool = true
	fmt.Println("Boolean:", isActive)

	// Integer types
	var i int = 42
	var i8 int8 = -128
	var i16 int16 = 32767
	var i32 int32 = -2147483648
	var i64 int64 = 9223372036854775807
	fmt.Println("Integers:", i, i8, i16, i32, i64)

	// Unsigned integers
	var u uint = 42
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	fmt.Println("Unsigned Integers:", u, u8, u16, u32, u64)

	// Floating point
	var f32 float32 = 3.14
	var f64 float64 = 2.718281828
	fmt.Println("Floats:", f32, f64)

	// Complex numbers
	var c1 complex64 = 1 + 2i
	var c2 complex128 = 3 + 4i
	fmt.Println("Complex:", c1, c2)
	fmt.Println("Real part:", real(c1), "Imaginary part:", imag(c1))

	// String
	var name string = "Rayhan"
	fmt.Println("String:", name)
	fmt.Println("Length of string:", len(name))
	fmt.Println("First character (byte):", name[0])
	multi := `Hello
World`
	fmt.Println("Multiline String:", multi)

	// Rune and Byte
	var b byte = 'A'   // alias for uint8
	var r rune = '♥'   // alias for int32 (Unicode)
	fmt.Println("Byte:", b, "Rune:", r)

	// ---------- DERIVED / COMPOSITE TYPES ----------

	// Array
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("Slice:", slice)

	// Map
	person := map[string]string{
		"name": "Rayhan",
		"city": "Dhaka",
	}
	fmt.Println("Map:", person)
	fmt.Println("Name:", person["name"])

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Rayhan", Age: 26}
	fmt.Println("Struct:", p)

	// Pointer
	var num int = 100
	var ptr *int = &num
	fmt.Println("Pointer Address:", ptr)
	fmt.Println("Pointer Value:", *ptr)

	// Function type
	func add(a, b int) int {
		return a + b
	}
	var operation func(int, int) int = func(a, b int) int {
		return a + b
	}
	fmt.Println("Function Type:", operation(5, 7))

	// Interface
	type Shape interface {
		Area() float64
	}
	type Circle struct {
		radius float64
	}
	func (c Circle) Area() float64 {
		return 3.14 * c.radius * c.radius
	}
	circle := Circle{radius: 5}
	fmt.Println("Interface Example - Circle Area:", circle.Area())

	// Channel (communication between goroutines)
	ch := make(chan int)
	go func() { ch <- 42 }()
	fmt.Println("Channel:", <-ch)

	// ---------- CUSTOM & SPECIAL TYPES ----------

	type ID int
	var userID ID = 123
	fmt.Println("Custom Type:", userID)

	type MyString = string
	var greeting MyString = "Hello, Golang!"
	fmt.Println("Type Alias:", greeting)

	// Error type example
	var err error = nil
	fmt.Println("Error Type:", err)
}
