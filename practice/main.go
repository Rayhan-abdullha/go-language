package main

const email string = "coder@gmail.com"

// func getLength(str string) int {
// 	l := 0
// 	for i := 0; i < len(str); i++ {
// 		l++
// 	}
// 	// this is while loop
// 	for {
// 		if l > 99 {
// 			break
// 		}
// 		l++
// 	}
// 	return l
// }

// func getCoords(a int, b int) (x int, y int) {
// 	return
// }

// func calculator(a, b int) (mul, div int, err error) {
// 	if b == 0 {
// 		return 0, 0, errors.New("Can't divide by zero")
// 	}
// 	mul = a * b
// 	div = a / b
// 	return mul, div, nil
// }

type Obj struct {
	name  string
	value int
}

func main() {
	// mul, _, err := calculator(10, 0)
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }
	// res := mul
	// println(res)

	// mathlib.Add(3, 3)
	// closure.Closure()

	// user := Obj{
	// 	name:  "Gopher",
	// 	value: 42,
	// }
	// userClone := &user
	// userClone.name = "Changed"
	// println(user.name)

	bankSystem()

}

func bankSystem() {
	type Account struct {
		accountNumber string
		balance       float64
	}
	var createAccount = func(accountNumber string, initialBalance float64) *Account {
		return &Account{
			accountNumber: accountNumber,
			balance:       initialBalance,
		}
	}

	var deposit = func(acc *Account, amount float64) {
		acc.balance += amount
	}
	var withdraw = func(acc *Account, amount float64) bool {
		if amount > acc.balance {
			return false
		}
		acc.balance -= amount
		return true
	}
	var getBalance = func(acc *Account) float64 {
		return acc.balance
	}
	account := createAccount("123456", 1000.0)
	deposit(account, 500.0)
	success := withdraw(account, 200.0)
	if success {
		println("Withdrawal successful")
	}
	balance := getBalance(account)
	println("Current balance:", balance)
}
