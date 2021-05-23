package main

import "fmt"

func main() {
	var m map[string]int
	p := make(map[string]int)

	fmt.Println(p, m)

	a := m["key"]
	b := p["key"]

	fmt.Println(a, b)

	m = make(map[string]int)
	m["a"] = 9
	p["b"] = 10

	fmt.Println(p, m)

	value, ok := p["key"]

	fmt.Println(value, ok)

	value, ok = p["a"]

	fmt.Println(value, ok)

	m["c"]++

	value, ok = m["c"]

	fmt.Println(value, ok)

	do(m)

	value, ok = m["d"]

	fmt.Println(value, ok)
	fmt.Println(len(m))

	var n map[[10]int]int
	fmt.Println(n == nil)

	o := map[int]int{}

	if value, ok := o[2]; ok {
		fmt.Println(value)
	}

	delete(o, 9) // if exists delete, else no change

}

func do(m map[string]int) map[string]int {
	m["d"] = 90

	return m
}
