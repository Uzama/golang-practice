package main

import "fmt"

var a int

var b int = 1

var c = 1

var d = 1.0

var (
	x, y int
	z    = "Hello"
	s    string
)

type Apple string

type S interface{}

func main() {
	fmt.Printf("value: %v, type: %[1]T\n", a)
	fmt.Printf("value: %v, type: %[1]T\n", b)
	fmt.Printf("value: %v, type: %[1]T\n", c)
	fmt.Printf("value: %v, type: %[1]T\n", d)
	fmt.Printf("value: %v, type: %[1]T\n", x)
	fmt.Printf("value: %v, type: %[1]T\n", y)
	fmt.Printf("value: %v, type: %[1]T\n", z)
	fmt.Printf("value: %v, type: %[1]T\n", s)

	var a Apple = "I am Apple"

	fmt.Printf("value: %v, type: %[1]T\n", a)

	b := 2

	if b := 9; b%2 != 0 {
		b := 9
		fmt.Println(b)
	}

	fmt.Println(b)

	var s S
	// val, ok := s.(Apple)

	// fmt.Println(val, ok)

	fmt.Printf("%T\n", s)

	f := Hello

	f(4)

	d := Hi

	f = d

	f(3)

	m := bat{Name: ""}

	n := ball{Name: ""}

	o := wicket{}

	fmt.Println(m, n, o)

	// fmt.Println(m == n)

	var u Apple

	v := "Hello"

	// u = v

	u = "HI"

	u = Apple(v)

	fmt.Println(u)

}

func Hello(a int) {
	fmt.Println(a)
}

func Hi(a int) {

}

type ball struct {
	Name string
	Age  int
}

type bat struct {
	Name string
	Age  int
}

type wicket struct {
	Hello string
	Age   int
}
