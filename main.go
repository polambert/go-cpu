
package main

import (
	"fmt"
	"bufio"
	"os"

	"strings"
	"strconv"

	"io/ioutil"
)

func input(p string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(p)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	return text
}

func GetBytes(args []string) []uint16 {
	data := ""
	if len(args) >= 2 {
		filename := args[1]
		b, _ := ioutil.ReadFile(filename)
		data = string(b)
	} else {
		fmt.Println("BYTECODE INTERPRETER:")
		fmt.Println("PLEASE ENTER A STRING OF BYTES TO INTERPRET")

		data = input("> ")
	}

	// Get an array of bytes
	// Separated by spaces
	fields := strings.Fields(data)
	insts := make([]uint16, len(fields))
	for i := 0; i < len(fields); i++ {
		n, _ := strconv.Atoi(fields[i])
		insts[i] = uint16(n)
	}

	return insts
}

func main() {
	insts := GetBytes(os.Args)

	// Load the interpreter and run it

	RunBytes(insts)
}
