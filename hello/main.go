package main

import (
	"fmt"
	"os"
	"practice/hello/greet"
)

func main() {
	if len(os.Args) > 1 {
		message := greet.Say(os.Args[1:])
		fmt.Println(message)
	} else {
		fmt.Println("argument not provided")
	}
}
