package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ./output hi bye < text.txt
func main() {
	s := "%Zaid Σ"

	fmt.Printf("%8T, %[1]v\n", s)
	fmt.Printf("%8T, %[1]v\n", s[3:])
	fmt.Printf("%8T, %[1]v\n", s[7])
	fmt.Printf("%8T, %[1]v\n", []rune(s))
	fmt.Printf("%8T, %[1]v\n", []byte(s))

	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
	fmt.Println(len([]byte(s)))

	d := "Σuza"
	i := 0

	for i < len(d) {
		fmt.Printf("%v\n", string(d[i]))
		i++
	}

	message := "Hello World"
	fmt.Printf("%s, %[1]T, %d\n", message, len(message))
	part := message[5:]
	message += " All"
	fmt.Printf("%s, %[1]T, %d\n", message, len(message))
	message += " ho"
	// part = append(part, "h")
	fmt.Println(part)
	fmt.Println(message)

	s = "a string"

	fmt.Println(strings.Contains(s, "s"))
	fmt.Println(strings.Contains(s, "u"))
	fmt.Println(strings.HasPrefix(s, "A"))
	fmt.Println(strings.HasSuffix(s, "g"))
	fmt.Println(strings.Index(s, "in"))
	d = s
	s = strings.ToUpper(s)

	fmt.Println(d, s)

	fmt.Println("\n----------------")

	// replace string

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enough arguments")
		os.Exit(-1)
	}

	old, newString := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, newString)

		fmt.Println(t)
	}
}
