// part6 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"part6/interpreter"
	"strings"
)

func main() {
	fmt.Println("Let's Build A Simple Interpreter - Part 6")

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
