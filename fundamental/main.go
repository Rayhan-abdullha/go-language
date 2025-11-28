package main

func main() {
	// var a = "initial value"
	// var b string
	// b = "changed value"
	// fmt.Println(a, b)

	// var b, c = 2, 4
	// fmt.Println(b, c)

	//========== loop =========//

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := range 10 {
	// 	fmt.Println(i)
	// }

	// for {
	// 	// infinity loop
	// }

	// i := 2
	// for i > 1 {
	// 	if i > 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// arr := [5]int{2, 3, 4, 5, 6}

	// for index, value := range arr {
	// 	fmt.Println("index:", index, "value:", value)
	// }

	//========== if-else =========//
	// a := 5
	// if a < 10 {
	// 	fmt.Println("a is less than 10")
	// } else if a < 20 {
	// 	fmt.Println("a is less than 20")
	// } else {
	// 	fmt.Println("a is greater than or equal to 20")
	// }

	// ========== switch =========//
	// day := 3
	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// default:
	// 	fmt.Println("Another day")
	// }

	// ========== arr =========//
	// arr := [5]int{2, 3, 4, 5, 6}
	// fmt.Println(arr)

	// arr = [...]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl:", arr)

	// ========== slice =========//
	// slice := []int{}
	// s1 := make([]int, 3, 10)
	// s1 = append(s1, 10, 200, 320)
	// fmt.Println("size ", len(s1))
	// fmt.Println("cap = ", cap(s1))
	// fmt.Println("main slice = ", s1)
	// fmt.Println(s1[3:5])
	// s1 = s1[:len(s1)-1]
	// fmt.Println(s1)
	// fmt.Println("end of main func")

	//========== map =========//
	// m1 := map[string]int{
	// 	"rayhan":  28,
	// 	"abdullah": 30,
	// }
	// fmt.Println(m1)

	m := make(map[string]any)
	m["name"] = "rayhan abdullah"
	m["age"] = 28

	// delete(m, "age")
	// if m["age"] == nil {
	// 	fmt.Println("age key is not found")
	// } else {
	// 	fmt.Println("age:", m["age"])
	// }

	// value, ok := m["salary"]
	// fmt.Println(ok)
	// if ok {
	// 	fmt.Println("Salary:", value)
	// } else {
	// 	fmt.Println("salary not found")
	// }

}
