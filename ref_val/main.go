package main

import "fmt"

type Widget struct {
	ID    int
	Count int
}

func main() {
	w := Widget{ID: 1, Count: 2}
	fmt.Println(w)

	// v := Expend(w)
	// fmt.Println(v)

	f1(&w)
	fmt.Println(w)

	a := []Widget{w, {ID: 2}}

	for _, s := range a {
		s.Count++
	}

	fmt.Println(a)

	for i := range a {
		a[i].Count++
	}

	fmt.Println(a)

	b := []int{1}
	b = append(b, 6)
	fmt.Println(b, cap(b), len(b))

	b = update(b)

	fmt.Println(b, cap(b), len(b))

	users := []user{
		{"Uzama", 0},
		{"Zaid", 1},
	}

	uzama := &users[0]
	sah := user{"Sah", 7}

	users = append(users, sah)
	fmt.Println(users)
	addUser(uzama)

	fmt.Println(users)

	items := [][2]byte{{2, 3}, {5, 6}, {9, 8}}

	t := [][]byte{}

	for _, item := range items {
		i := make([]byte, len(item))
		copy(i, item[:])
		t = append(t, i)
	}

	fmt.Println(t)
}

type user struct {
	name  string
	count int
}

func addUser(u *user) {
	u.count++
}

func update(a []int) []int {
	a = append(a, 2)
	return a
}

func Expend(w Widget) Widget {
	w.Count++
	return w
}

func f1(w *Widget) {
	w.Count++
	fmt.Println(w)
	f2(w)
}

func f2(w *Widget) {
	w.Count++
	fmt.Println(w)
	f3(*w)
}

func f3(w Widget) {
	w.Count++
	fmt.Println(w)
	f4(&w)
}

func f4(w *Widget) {
	w.Count++
	fmt.Println(w)
}
