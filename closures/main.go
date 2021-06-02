package main

import "fmt"

func main() {
	f, g := fib(), fib()

	for i := f(); i < 5; i = f() {
		fmt.Println(i)
	}

	for i := g(); i < 5; i = g() {
		fmt.Println(i)
	}

	s := []int{4, 2, 6, 7, 3, 1, 0, 7}
	fmt.Println(s)
	sort(s, func(i, j int) bool {
		return s[i] > s[j]
	})

	fmt.Println(s)

	for i := 0; i < 5; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}

		do(v)
	}

	fun := make([]func(), 4)

	for i := 0; i < 4; i++ {
		i := i
		fun[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		fun[i]()
	}

}

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

func sort(s []int, f func(i, j int) bool) {
	for i := range s {
		for j := i; j < len(s)-1; j++ {
			if f(i, j) {
				temp := s[i]
				s[i] = s[j]
				s[j] = temp
			}
		}

	}
}

func do(d func()) {
	d()
}
