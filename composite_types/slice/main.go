package main

import "fmt"

func main() {
	var a []int
	var b = []int{1, 2}

	a = b
	c := b

	c = append(c, 8)

	e := a
	a = append(a, 10)
	a[0] = 80
	c[0] = 9

	c = append(c, 89)
	c = append(c, 89)

	b[1] = 9000

	f := c[1:5]

	f = append(f, 90)
	f[4] = 23
	f[0] = 0

	c = append(c, 876)
	f = append(f, 33)
	f[4] = -3

	f = append(f, -9)
	f = append(f, -9)

	c[1] = 34847584

	g := f[1:4]
	g = append(g, 31)

	f = append(f, 78)

	fmt.Printf("value:%v, capacity:%d, length:%d\n", a, cap(a), len(a))
	fmt.Printf("value:%v, capacity:%d, length:%d\n", b, cap(b), len(b))
	fmt.Printf("value:%v, capacity:%d, length:%d\n", c, cap(c), len(c))
	fmt.Printf("value:%v, capacity:%d, length:%d\n", e, cap(e), len(e))
	fmt.Printf("value:%v, capacity:%d, length:%d\n", f, cap(f), len(f))
	fmt.Printf("value:%v, capacity:%d, length:%d\n", g, cap(g), len(g))

	w := [...]int{1, 2, 3}
	x := []int{40, 10, 70}

	y, z := do(w, x)
	z[0] = 38484
	fmt.Println(w, x, y, z)
	fmt.Printf("%T, %T \n", w, x)
}

func do(a [3]int, b []int) ([]int, []int) {
	a[0] = 4
	b[0] = 3

	c := make([]int, 5)
	c[4] = 78
	copy(c, b)

	d := []int{4, 5}

	return c, d
}
