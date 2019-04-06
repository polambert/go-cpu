
package main

import (
	"fmt"
	"bufio"
	"os"

	"strings"
	"strconv"
)

func input(p string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(p)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	return text
}

func main() {
	fmt.Println("BYTECODE INTERPRETER:")
	fmt.Println("PLEASE ENTER A STRING OF BYTES TO INTERPRET")

	str := input("> ")

	// Get an array of bytes
	// Separated by spaces
	fields := strings.Fields(str)
	insts := make([]uint16, len(fields))
	for i := 0; i < len(fields); i++ {
		n, _ := strconv.Atoi(fields[i])
		insts[i] = uint16(n)
	}

	// Load the interpreter and run it

	RunBytes(insts)
}
