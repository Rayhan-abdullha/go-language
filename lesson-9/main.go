package main

func calculator(a int, b int) (mul int, div int) {
	mul = a * b
	div = a / b
	return
}
func main() {
	// arr := [5]int{5, 2, 9, 44, 8}
	// for i := range len(arr) {
	// 	// if i > 5 {
	// 	// 	continue
	// 	// }
	// 	println(arr[i])
	// }

	// 2D Array
	// nums := [2][2]int{{2, 3}, {3, 5}}
	// fmt.Println(nums)

	// for i := range 10 {
	// 	println(i)
	// }

	// we can create var inside if condition
	/*
		if age := 10; age > 18 {
			println("You can watch Zero movie")
		} else {
			println("Warning!")
		}
	*/

	/*
		switch time.Now().Weekday() {
		case time.Saturday, time.Friday:
			println("Office is closed, enjoy weekend")
		default:
			println("Working Day")
		}
	*/

	/*
		for i := range 5 {
			switch i {
			case 1:
				fmt.Println(i, " is not prime number")
			case 2:
				fmt.Println(i, " is prime number")
			case 3:
				fmt.Println(i, " is not prime number")
			case 4:
				fmt.Println(i, " is prime number")
			case 5:
				fmt.Println(i, " is not prime number")
				}
			}
	*/

	// type Address struct {
	// 	dist     string
	// 	division string
	// }
	// type User struct {
	// 	name    string
	// 	age     int
	// 	email   string
	// 	address Address
	// }
	// var user1 = User{
	// 	name:  "John",
	// 	age:   30,
	// 	email: "john@gmail.com",
	// 	address: Address{
	// 		dist:     "Dhaka",
	// 		division: "Dhaka",
	// 	},
	// }
	// fmt.Println(user1)
	// sum := func(nums ...int) int {
	// 	total := 0
	// 	for _, v := range nums {
	// 		total += v
	// 	}
	// 	return total
	// }
	// res := sum(1, 2, 3, 4, 5)
	// println(res)

	m, d := calculator(10, 2)
	println("Multiplication: ", m)
	println("Division: ", d)
}
