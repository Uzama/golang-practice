package main

import (
	"fmt"
	"os"
)

var a int

var (
	b = 2
	c = 3.45
)

const (
	q = 1
	w = 2 * 234
	z = w < 44

	l uint8 = 0x09
	g uint8 = l & 0x7
)

// run the program using command: ./output < num.txt
func main() {
	d := 6 // only inside function
	e := 6.4

	fmt.Println(a, b, c, d)
	fmt.Printf("%4d: %7[1]T %[1]v\n", d)
	fmt.Printf("%.2f: %7[1]T %[1]v\n", e)
	fmt.Printf("%.2f: %7[1]T %[1]v\n", c)
	d = int(c)
	fmt.Printf("%4d: %7[1]T %[1]v\n", d)
	e = float64(d)
	fmt.Printf("%.2f: %7[1]T %[1]v\n", e)

	fmt.Println("\n-------------")

	var sum float64
	var n int

	for {
		var val float64

		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("average is: ", sum/float64(n))
}
