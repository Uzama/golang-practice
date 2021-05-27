package main

import "fmt"

func A() {
	fmt.Println("This is function A")
}

func main() {
	b := func() {
		fmt.Println("This is function B")
	}

	B(b)
	c := C()

	c()

	fmt.Println(D("Dei, "))

	fmt.Println(E(map[string]int{"a": 9}))

	m := map[int]int{
		1: 10,
		2: 20,
	}
	fmt.Println(m)
	do(&m)
	fmt.Println(m)

	val, err := doIt(7)

	fmt.Println(val, err.Error())

}

func B(f func()) {
	f()
}

func C() func() {

	fmt.Println("this is func C")

	return func() {
		fmt.Println("this is fun C return")
	}
}

func D(s string) string {
	return s + "hello"
}

func E(a map[string]int) map[string]int {

	a["a"] = 1
	return a
}

func do(m *map[int]int) {
	(*m)[1] = 100
	fmt.Println(*m)
	*m = make(map[int]int)
	(*m)[1] = 200
	fmt.Println(*m)
}

func doIt(i int) (int, error) {
	fmt.Println("started doIt")
	defer fmt.Println(D("hoi, "))
	defer fmt.Println("he")
	return i, fmt.Errorf("error occurred")
}
