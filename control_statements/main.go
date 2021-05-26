package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	a := 3
	b := 4

	if a == b {
		fmt.Println("a is equal to be")
	} else {
		fmt.Println("a is not equal to be")
	}

	if flag := doSomething(b); flag {
		fmt.Println("Did somthing")
	}

	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
		fmt.Print(math.Pow(float64(i), 2), " \n")
	}

	arr := []int{4, 5, 6, 3}
	for i := range arr {
		fmt.Println(i, " - ", arr[i])
	}

	for index, value := range arr {
		fmt.Println(index, "-", value)
	}

	for _, value := range arr {
		fmt.Print(value, " ")
	}
	fmt.Println()

	myMap := map[string]int{
		"Uzama": 34,
		"Zaid":  34,
	}

	for key := range myMap {
		fmt.Println(key, myMap[key])
	}

	for key, val := range myMap {
		fmt.Println(key, val)
	}

	i, j := 0, 5

	for {
		fmt.Println(i, j)
		if i > j {
			break
		}

		i++
	}

	start := time.Now()
	k := 0
	for {

		Ages := []int{22, 34, 21, 56, 78}
		record := map[string]int{
			"Uzama":   22,
			"Zaid":    34,
			"Mohamed": 17,
			"Fathima": 45,
			"Zulaiha": 78,
		}

		// var found []string

	outer:
		for _, age := range Ages {
			for _, val := range record {
				if age == val {
					// found = append(found, key)
					continue outer
				}
			}
			// fmt.Println(age, "not found")
		}
		if k > 10000000 {
			break
		}
		k++

	}
	fmt.Println(float64(time.Since(start)) / float64(k))

	switch day := time.Now().Weekday(); day {
	case 0, 6:
		fmt.Println("WeekEnds")
	case 1, 2, 3, 4, 5:
		fmt.Println("WeekDay")
	default:
		fmt.Println(day)
	}

	day := time.Now().Weekday()

	switch {
	case day == 0 || day == 6:
		fmt.Println("Week end")

	case day >= 1 && day <= 5:
		fmt.Println("Week day")

	default:
		fmt.Println("Wrong input")

	}
}

func doSomething(a int) bool {
	return a%2 == 0
}
