package main

type S struct {
	Name string
	Age  int
}

func main() {
	a := [...]int{4, 5, 3, 1, 2}

	b := []*int{}

	for i := range a {
		b = append(b, &a[i])
	}

	for range b {
	}

	_ = process()
}

func get() *S {
	return &S{
		Name: "Uzama",
		Age:  45,
	}
}

func process() *S {
	g := get()

	s := S{}

	s.Name = g.Name
	s.Age = g.Age

	return &s
}
