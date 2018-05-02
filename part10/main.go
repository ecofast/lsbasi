// part10 project main.go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"part10/interpreter"
	"strings"
	"unsafe"
)

func main() {
	fmt.Println("Let's Build A Simple Interpreter - Part 10")

	interdivter := interpreter.New()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("pls enter file name: ")
		if s, err := reader.ReadString('\n'); err == nil {
			fileName := strings.TrimSpace(s)
			if b, err := ioutil.ReadFile(fileName); err == nil {
				interdivter.Interpret(*(*string)(unsafe.Pointer(&b)))
				interdivter.PrintSymbolTable()
			} else {
				fmt.Printf("cannot find file: %s\n", fileName)
			}
			continue
		}
		break
	}
}
