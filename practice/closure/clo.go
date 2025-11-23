package closure

const PI = 3.14

var name = "Golang"

func closureFunc() func() {
	count := 0
	show := func() {
		count++
		println("Closure called", count, "times")
	}
	return show
}
func Closure() {
	f := closureFunc()
	f()
	f()
	k := closureFunc()
	k()
	f()
	k()
}
func init() {
	println("Init function in closure package")
}
