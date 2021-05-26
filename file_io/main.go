package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type S struct {
	Name string
	Age  int
}

// ./output 3 4 8 3  6 3 5 > text.txt
func main() {
	fmt.Println("printing a line to standard output")

	fmt.Printf("my struct: %+v\n", S{Name: "Uzama", Age: 23})

	fmt.Printf("hello %q\n", "hi")

	flag := true

	fmt.Printf("value: %t\n", flag)

	flt := 4.56785

	fmt.Printf("value: %.3f\n", flt)

	fmt.Printf("|%5d|%05d|%-5d|%5d|\n", 4, 5, 6, 7)

	s := []int{4, 5, 6}

	fmt.Printf("value: %v, type: %[1]T, special: %#[1]v\n", s)
	fmt.Printf("value: %v, type: %[1]T, special: %#[1]v\n", S{Name: "Uzama", Age: 23})
	fmt.Printf("value: %q, type: %[1]T, special: %#[1]q\n", [3]rune{'a', 'e', 't'})
	fmt.Printf("value: %q, type: %[1]T, special: %#[1]q\n", S{Name: "Uzama", Age: 23})
	fmt.Printf("value: %v, type: %[1]T, special: %#[1]v\n", map[string]int{"a": 1, "n": 2})
	fmt.Printf("value: %+v, type: %[1]T, special: %#[1]v\n", map[string]int{"a": 1, "n": 2})
	fmt.Printf("value: %q, type: %[1]T, special: %#[1]q\n", map[string]int{"a": 1, "n": 2})

	fmt.Printf("value: %q\n", []byte("Hello World"))

	fmt.Fprintln(os.Stderr, "printing to error output")

	// if arg := os.Args; len(arg) > 1 {
	// 	min, _ := strconv.Atoi(arg[1])
	// 	max, _ := strconv.Atoi(arg[1])

	// 	for i := 2; i < len(arg); i++ {
	// 		n, _ := strconv.Atoi(arg[i])

	// 		if n < min {
	// 			min = n
	// 		}

	// 		if n > max {
	// 			max = n
	// 		}
	// 	}

	// 	fmt.Printf("Max: %d, Min: %d\n", max, min)
	// }

	// io.WriteString(os.Stderr, fmt.Sprintln("write to std err"))

	// f := os.Stdin

	// defer f.Close()

	// scanner := bufio.NewScanner(f)

	// for scanner.Scan() {
	// 	fmt.Println(">> ", scanner.Text())
	// }

	fmt.Printf("%15s %15s %15s %15s\n", "Line Count", "Word Count", "Character Count", "Filename")

	for _, fName := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fName)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}

		fmt.Printf("%15d %15d %15d %15s\n", lc, wc, cc, fName)

		file.Close()
	}

}
