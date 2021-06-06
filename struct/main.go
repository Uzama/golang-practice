package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string
	Age     int
	Manager *Employee
	DOB     string
}

type S1 struct {
	Name string `json:"n"`
}

type S2 struct {
	Name string `json:"nam"`
}

func (e Employee) print() {
	manager := fmt.Sprintf("%+v", e.Manager)

	fmt.Printf("Name: %s, Age: %d, Manager, %s, DOB: %s\n", e.Name, e.Age, manager, e.DOB)
}

func main() {

	company := map[int]*Employee{}

	company[0] = &Employee{Name: "Zaid", Age: 24}

	company[1] = &Employee{}
	company[1].Name = "Uzama"
	company[1].Age = 24
	company[1].Manager = company[0]

	company[0].print()
	company[1].print()

	vehicle := struct {
		name  string
		model int
	}{
		"Car",
		34,
	}

	v := &vehicle

	fmt.Printf("%+v\n", vehicle)
	fmt.Println(v)

	a := struct {
		name string
	}{"Uzama"}

	b := struct {
		name string
	}{"Zaid"}

	fmt.Println(a, b)

	a = b

	fmt.Println(a, b)

	s1 := S1{
		"Hi",
	}

	s2 := S2{
		"Hello",
	}

	fmt.Println(s1, s2)

	s1 = S1(s2)
	fmt.Println(s1, s2)

	var isPresent map[int]struct{}

	fmt.Println(isPresent[5])

	done := make(chan struct{})

	fmt.Println(done)

	r := Response{
		Page: 1,
	}

	h := Response{Page: 3, Words: nil}

	fmt.Printf("%#v\n", r)
	fmt.Printf("%#v\n", h)

	data, _ := json.Marshal(r)

	fmt.Println(string(data))

	var k Response

	json.Unmarshal(data, &k)

	fmt.Printf("%#v\n", k)
}

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}
