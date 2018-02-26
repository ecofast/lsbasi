// part1 project main.go
package main

import (
	"fmt"
	"part1/interpreter"
)

func main() {
	fmt.Println("Let's Build A Simple Interpreter - Part 1")

	parser := interpreter.New()
	s := ""
	for {
		if n, err := fmt.Scan(&s); n == 0 || err != nil {
			return
		}
		fmt.Println(parser.Parse(s))
	}
}
