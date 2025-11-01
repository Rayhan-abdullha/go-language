package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printUser(usr User) {
	usr.Age = 130
	usr.Name = "Rayhan Abdullah"
}

// receiver function
func (usr User) printUser1() {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}
func (usr User) printUser2(a int) {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
	fmt.Println(a)
}
func main() {
	var user User
	user = User{
		Name: "Rayhan Abdullah",
		Age:  28,
	}
	// fmt.Println("befor = ", user)
	// printUser(user)
	// fmt.Println("after changed = ", user)

	// receiver function call
	user.printUser1()
	user.printUser2(10)
}
