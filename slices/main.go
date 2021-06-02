package main

import "fmt"

func main() {
	var s []int
	var t = []int{}
	var u = make([]int, 5)
	var v = make([]int, 0, 5)

	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil)

	a := [3]int{1, 2, 3}
	b := a[0:1]

	fmt.Printf("a: %T, %[1]v, %d, %d\n", a, cap(a), len(a))
	fmt.Printf("b: %T, %[1]v, %d, %d\n", b, cap(b), len(b))

	c := b[:2]

	fmt.Printf("c: %T, %[1]v, %d, %d\n", c, cap(c), len(c))

	d := c[:1:1]

	fmt.Printf("d: %T, %[1]v, %d, %d\n", d, cap(d), len(d))

	f := [...]int{1, 2, 3, 7}

	g := f[0:1:1]
	h := f[0:2:2]

	fmt.Printf("f[%p] = %v\n", &f, f)
	fmt.Printf("g[%p] = %[1]v\n", g)
	fmt.Printf("h[%p] = %[1]v\n", h)

	h = append(h, 4)

	fmt.Printf("f[%p] = %v\n", &f, f)
	fmt.Printf("h[%p] = %[1]v\n", h)

	g = append(g, 10)

	fmt.Printf("f[%p] = %v\n", &f, f)
	fmt.Printf("g[%p] = %[1]v\n", g)
	fmt.Printf("h[%p] = %[1]v\n", h)
}
