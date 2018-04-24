// part7 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"part7/interpreter"
	"strings"
)

func main() {
	fmt.Println("Let's Build A Simple Interpreter - Part 7")

	i := interpreter.New()
	reader := bufio.NewReader(os.Stdin)
	for {
		if s, err := reader.ReadString('\n'); err == nil {
			fmt.Println(i.Interpret(strings.TrimSpace(s)))
			continue
		}
		break
	}
}
