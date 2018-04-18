// part4 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"part4/interpreter"
	"strings"
)

func main() {
	fmt.Println("Let's Build A Simple Interpreter - Part 4")

	parser := interpreter.New()
	reader := bufio.NewReader(os.Stdin)
	for {
		if s, err := reader.ReadString('\n'); err == nil {
			fmt.Println(parser.Parse(strings.TrimSpace(s)))
			continue
		}
		break
	}
}
