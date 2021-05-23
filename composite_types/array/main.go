package main

import "fmt"

func main() {
	var a [3]int
	var b = [4]int{4, 5, 3, 3}
	var c = [...]int{2, 3, 4}

	d := a
	a[2] = 7

	var e = [...]int{4, 4, 4}

	e = c

	f := e[1:3]
	// f = append(f, 4)
	// f = append(f, 4)
	// f = append(f, 4)

	f[0] = 4567

	fmt.Printf("value:%v, capacity:%d, length:%d\n", a, cap(a), len(a))
	fmt.Printf("value:%v, capacity:%d, length:%d\n", b, cap(b), len(b))
	fmt.Printf("value:%v, capacity:%d, length:%d\n", c, cap(c), len(c))
	fmt.Printf("value:%v, capacity:%d, length:%d\n", d, cap(d), len(d))
	fmt.Printf("value:%v, capacity:%d, length:%d\n", e, cap(e), len(e))

	f = append(f, 4)
	fmt.Printf("%T, %[1]v\n", f)

}
